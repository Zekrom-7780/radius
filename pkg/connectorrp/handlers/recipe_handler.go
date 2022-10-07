// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/project-radius/radius/pkg/armrpc/api/conv"
	"github.com/project-radius/radius/pkg/azure/armauth"
	"github.com/project-radius/radius/pkg/azure/clients"
	"github.com/project-radius/radius/pkg/connectorrp/datamodel"
	coreDatamodel "github.com/project-radius/radius/pkg/corerp/datamodel"
	"github.com/project-radius/radius/pkg/radlogger"
	"github.com/project-radius/radius/pkg/resourcemodel"
	ucpresources "github.com/project-radius/radius/pkg/ucp/resources"
	"oras.land/oras-go/v2/content"
	"oras.land/oras-go/v2/registry/remote"
)

const deploymentPrefix = "recipe"

// RecipeHandler is an interface to deploy and delete recipe resources
//
//go:generate mockgen -destination=./mock_recipe_handler.go -package=handlers -self_package github.com/project-radius/radius/pkg/connectorrp/handlers github.com/project-radius/radius/pkg/connectorrp/handlers RecipeHandler
type RecipeHandler interface {
	DeployRecipe(ctx context.Context, recipe datamodel.RecipeProperties, providers coreDatamodel.ProviderProperties) ([]string, error)
	Delete(ctx context.Context, id string, apiVersion string) error
	GetResource(ctx context.Context, provider, id, apiVersion string) (map[string]interface{}, error)
}

func NewRecipeHandler(arm *armauth.ArmConfig) RecipeHandler {
	return &azureRecipeHandler{
		arm: arm,
	}
}

type azureRecipeHandler struct {
	arm *armauth.ArmConfig
}

// DeployRecipe deploys the recipe template fetched from the provided recipe TemplatePath using the providers scope.
// Currently the implementation assumes TemplatePath is location of an ARM JSON template in Azure Container Registry.
// Returns resource IDs of the resources deployed by the template
func (handler *azureRecipeHandler) DeployRecipe(ctx context.Context, recipe datamodel.RecipeProperties, providers coreDatamodel.ProviderProperties) (deployedResources []string, err error) {
	if recipe.TemplatePath == "" {
		return nil, fmt.Errorf("recipe template path cannot be empty")
	}
	if providers == (coreDatamodel.ProviderProperties{}) {
		return nil, conv.NewClientErrInvalidRequest(fmt.Sprintf("failed to deploy recipe %q. Environment provider scope is required to deploy connector recipes.", recipe.Name))
	}
	subscriptionID, resourceGroup, err := parseAzureProvider(&providers)
	if err != nil {
		return nil, err
	}

	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldResourceGroup, resourceGroup,
		radlogger.LogFieldSubscriptionID, subscriptionID,
	)
	logger.Info(fmt.Sprintf("Deploying recipe: %q, template: %q", recipe.Name, recipe.TemplatePath))

	registryRepo, tag := strings.Split(recipe.TemplatePath, ":")[0], strings.Split(recipe.TemplatePath, ":")[1]
	// get the recipe from ACR
	// client to the ACR repository in the templatePath
	repo, err := remote.NewRepository(registryRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to create client to registry %s", err.Error())
	}

	digest, err := getDigestFromManifest(ctx, repo, tag)
	if err != nil {
		return nil, conv.NewClientErrInvalidRequest(fmt.Sprintf("failed to fetch template from the path %q for recipe %q: %s", recipe.TemplatePath, recipe.Name, err.Error()))
	}

	recipeBytes, err := getRecipeBytes(ctx, repo, digest)
	if err != nil {
		return nil, conv.NewClientErrInvalidRequest(fmt.Sprintf("failed to fetch template from the path %q for recipe %q: %s", recipe.TemplatePath, recipe.Name, err.Error()))
	}

	recipeData := make(map[string]interface{})
	err = json.Unmarshal(recipeBytes, &recipeData)
	if err != nil {
		return nil, err
	}

	// Using ARM deployment client to deploy ARM JSON template fetched from ACR
	dClient := clients.NewDeploymentsClient(subscriptionID, handler.arm.Auth)
	deploymentName := deploymentPrefix + strconv.FormatInt(time.Now().UnixNano(), 10)
	dplResp, err := dClient.CreateOrUpdate(
		ctx,
		resourceGroup,
		deploymentName,
		resources.Deployment{
			Properties: &resources.DeploymentProperties{
				Template: recipeData,
				Mode:     resources.DeploymentModeIncremental,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	err = dplResp.WaitForCompletionRef(ctx, dClient.BaseClient.Client)
	if err != nil {
		return nil, err
	}

	resp, err := dplResp.Result(dClient)
	if err != nil {
		return nil, err
	}

	// return error if the Provisioning is not success
	if resp.Properties.ProvisioningState != resources.ProvisioningStateSucceeded {
		return nil, fmt.Errorf("failed to deploy the recipe %q, template path: %q, deployment: %q", recipe.Name, recipe.TemplatePath, deploymentName)
	}

	// Get list of output resources deployed
	for _, id := range *resp.Properties.OutputResources {
		deployedResources = append(deployedResources, *id.ID)
	}

	return deployedResources, nil
}

// getDigestFromManifest gets the layers digest from the manifest
func getDigestFromManifest(ctx context.Context, repo *remote.Repository, tag string) (string, error) {
	// resolves a manifest descriptor with a Tag reference
	descriptor, err := repo.Resolve(ctx, tag)
	if err != nil {
		return "", err
	}
	// get the manifest data
	rc, err := repo.Fetch(ctx, descriptor)
	if err != nil {
		return "", err
	}
	defer rc.Close()
	manifestBlob, err := content.ReadAll(rc, descriptor)
	if err != nil {
		return "", err
	}
	// create the manifest map to get the digest of the layer
	var manifest map[string]interface{}
	err = json.Unmarshal(manifestBlob, &manifest)
	if err != nil {
		return "", err
	}
	// get the layers digest to fetch the blob
	layer, ok := manifest["layers"].([]interface{})[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to decode the layers from manifest")
	}
	layerDigest, ok := layer["digest"].(string)
	if !ok {
		return "", fmt.Errorf("failed to decode the layers digest from manifest")
	}
	return layerDigest, nil
}

// getRecipeBytes fetches the recipe ARM JSON using the layers digest
func getRecipeBytes(ctx context.Context, repo *remote.Repository, layerDigest string) ([]byte, error) {
	// resolves a layer blob descriptor with a digest reference
	descriptor, err := repo.Blobs().Resolve(ctx, layerDigest)
	if err != nil {
		return nil, err
	}
	// get the layer data
	rc, err := repo.Fetch(ctx, descriptor)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	pulledBlob, err := content.ReadAll(rc, descriptor)
	if err != nil {
		return nil, err
	}
	return pulledBlob, nil
}

// parseAzureProvider parses the scope to get the subscriptionID and resourceGroup
func parseAzureProvider(providers *coreDatamodel.ProviderProperties) (subscriptionID string, resourceGroup string, err error) {
	if providers.Azure == (coreDatamodel.ProviderPropertiesAzure{}) {
		return "", "", conv.NewClientErrInvalidRequest("environment does not contain Azure provider scope required to deploy recipes on Azure")
	}
	// valid scope: "/subscriptions/test-sub/resourceGroups/test-group"
	scope := strings.Split(providers.Azure.Scope, "/")
	if len(scope) != 5 {
		return "", "", conv.NewClientErrInvalidRequest(fmt.Sprintf("invalid azure scope. Valid scope eg: %q", "/subscriptions/<subscriptionID>/resourceGroups/<resourceGroup>"))
	}
	subscriptionID = scope[2]
	resourceGroup = scope[4]
	if subscriptionID == "" || resourceGroup == "" {
		return "", "", conv.NewClientErrInvalidRequest("subscriptionID and resourceGroup must be provided to deploy connector recipes to Azure")
	}
	return
}

func (handler *azureRecipeHandler) Delete(ctx context.Context, id string, apiVersion string) error {
	logger := radlogger.GetLogger(ctx)
	parsed, err := ucpresources.ParseResource(id)
	if err != nil {
		return err
	}

	rc := clients.NewGenericResourceClient(parsed.FindScope(ucpresources.SubscriptionsSegment), handler.arm.Auth)
	_, err = rc.DeleteByID(ctx, id, apiVersion)
	if err != nil {
		if !clients.Is404Error(err) {
			return fmt.Errorf("failed to delete resource %q: %w", id, err)
		}
		logger.Info(fmt.Sprintf("Recipe resource %s does not exist: %s", id, err.Error()))
	}
	return nil
}

func (handler *azureRecipeHandler) GetResource(ctx context.Context, provider, id, apiVersion string) (resource map[string]interface{}, err error) {
	if provider == resourcemodel.ProviderAzure {
		resource, err := GetByID(ctx, handler.arm.Auth, id, apiVersion)
		if err != nil {
			return nil, err
		}

		// Return the serialized resource so renderers can use it for computed values.
		serialized, err := SerializeResource(*resource)
		if err != nil {
			return nil, err
		}

		return serialized, nil
	}

	return map[string]interface{}{}, nil
}