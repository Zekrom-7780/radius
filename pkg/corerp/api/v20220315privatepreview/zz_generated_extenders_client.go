//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExtendersClient contains the methods for the Extenders group.
// Don't use this type directly, use NewExtendersClient() instead.
type ExtendersClient struct {
	host string
	rootScope string
	pl runtime.Pipeline
}

// NewExtendersClient creates a new instance of ExtendersClient with the specified values.
// rootScope - The scope in which the resource is present. For Azure resource this would be /subscriptions/{subscriptionID}/resourceGroups/{resourcegroupID}
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExtendersClient(rootScope string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExtendersClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ExtendersClient{
		rootScope: rootScope,
		host: ep,
pl: pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Extender resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// extenderName - The name of the Extender link resource
// extenderParameters - extender create parameters
// options - ExtendersClientCreateOrUpdateOptions contains the optional parameters for the ExtendersClient.CreateOrUpdate
// method.
func (client *ExtendersClient) CreateOrUpdate(ctx context.Context, extenderName string, extenderParameters ExtenderResource, options *ExtendersClientCreateOrUpdateOptions) (ExtendersClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, extenderName, extenderParameters, options)
	if err != nil {
		return ExtendersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtendersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ExtendersClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExtendersClient) createOrUpdateCreateRequest(ctx context.Context, extenderName string, extenderParameters ExtenderResource, options *ExtendersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/extenders/{extenderName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if extenderName == "" {
		return nil, errors.New("parameter extenderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extenderName}", url.PathEscape(extenderName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, extenderParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ExtendersClient) createOrUpdateHandleResponse(resp *http.Response) (ExtendersClientCreateOrUpdateResponse, error) {
	result := ExtendersClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtenderResponseResource); err != nil {
		return ExtendersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing extender resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// extenderName - The name of the Extender link resource
// options - ExtendersClientDeleteOptions contains the optional parameters for the ExtendersClient.Delete method.
func (client *ExtendersClient) Delete(ctx context.Context, extenderName string, options *ExtendersClientDeleteOptions) (ExtendersClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, extenderName, options)
	if err != nil {
		return ExtendersClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtendersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return ExtendersClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ExtendersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExtendersClient) deleteCreateRequest(ctx context.Context, extenderName string, options *ExtendersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/extenders/{extenderName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if extenderName == "" {
		return nil, errors.New("parameter extenderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extenderName}", url.PathEscape(extenderName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about a extender resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// extenderName - The name of the Extender link resource
// options - ExtendersClientGetOptions contains the optional parameters for the ExtendersClient.Get method.
func (client *ExtendersClient) Get(ctx context.Context, extenderName string, options *ExtendersClientGetOptions) (ExtendersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, extenderName, options)
	if err != nil {
		return ExtendersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtendersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtendersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtendersClient) getCreateRequest(ctx context.Context, extenderName string, options *ExtendersClientGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/extenders/{extenderName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if extenderName == "" {
		return nil, errors.New("parameter extenderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extenderName}", url.PathEscape(extenderName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtendersClient) getHandleResponse(resp *http.Response) (ExtendersClientGetResponse, error) {
	result := ExtendersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtenderResponseResource); err != nil {
		return ExtendersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByRootScopePager - Lists information about all extender resources in the given root scope
// Generated from API version 2022-03-15-privatepreview
// options - ExtendersClientListByRootScopeOptions contains the optional parameters for the ExtendersClient.ListByRootScope
// method.
func (client *ExtendersClient) NewListByRootScopePager(options *ExtendersClientListByRootScopeOptions) (*runtime.Pager[ExtendersClientListByRootScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ExtendersClientListByRootScopeResponse]{
		More: func(page ExtendersClientListByRootScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExtendersClientListByRootScopeResponse) (ExtendersClientListByRootScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRootScopeCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExtendersClientListByRootScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExtendersClientListByRootScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExtendersClientListByRootScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRootScopeHandleResponse(resp)
		},
	})
}

// listByRootScopeCreateRequest creates the ListByRootScope request.
func (client *ExtendersClient) listByRootScopeCreateRequest(ctx context.Context, options *ExtendersClientListByRootScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/extenders"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRootScopeHandleResponse handles the ListByRootScope response.
func (client *ExtendersClient) listByRootScopeHandleResponse(resp *http.Response) (ExtendersClientListByRootScopeResponse, error) {
	result := ExtendersClientListByRootScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtenderList); err != nil {
		return ExtendersClientListByRootScopeResponse{}, err
	}
	return result, nil
}

// ListSecrets - Lists secrets values for the specified Extender resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// extenderName - The name of the Extender link resource
// options - ExtendersClientListSecretsOptions contains the optional parameters for the ExtendersClient.ListSecrets method.
func (client *ExtendersClient) ListSecrets(ctx context.Context, extenderName string, options *ExtendersClientListSecretsOptions) (ExtendersClientListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, extenderName, options)
	if err != nil {
		return ExtendersClientListSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtendersClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtendersClientListSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *ExtendersClient) listSecretsCreateRequest(ctx context.Context, extenderName string, options *ExtendersClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/extenders/{extenderName}/listSecrets"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if extenderName == "" {
		return nil, errors.New("parameter extenderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extenderName}", url.PathEscape(extenderName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *ExtendersClient) listSecretsHandleResponse(resp *http.Response) (ExtendersClientListSecretsResponse, error) {
	result := ExtendersClientListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ExtendersClientListSecretsResponse{}, err
	}
	return result, nil
}
