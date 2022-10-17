//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

// AWSCredentialClientCreateOrUpdateOptions contains the optional parameters for the AWSCredentialClient.CreateOrUpdate method.
type AWSCredentialClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

type AWSCredentialProperties struct {
	// REQUIRED; Access key ID for AWS identity
	AccessKeyID *string `json:"accessKeyId,omitempty"`

	// REQUIRED; The kind of secret
	Kind *string `json:"kind,omitempty"`

	// REQUIRED; Secret Access Key for AWS identity
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`

	// REQUIRED
	Storage *CredentialResourcePropertiesStorage `json:"storage,omitempty"`
}

// GetCredentialResourceProperties implements the CredentialResourcePropertiesClassification interface for type AWSCredentialProperties.
func (a *AWSCredentialProperties) GetCredentialResourceProperties() *CredentialResourceProperties {
	return &CredentialResourceProperties{
		Kind: a.Kind,
		Storage: a.Storage,
	}
}

// AWSCredentialsClientDeleteOptions contains the optional parameters for the AWSCredentialsClient.Delete method.
type AWSCredentialsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AWSCredentialsClientGetOptions contains the optional parameters for the AWSCredentialsClient.Get method.
type AWSCredentialsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AWSCredentialsClientListOptions contains the optional parameters for the AWSCredentialsClient.List method.
type AWSCredentialsClientListOptions struct {
	// placeholder for future optional parameters
}

// AzureCredentialClientCreateOrUpdateOptions contains the optional parameters for the AzureCredentialClient.CreateOrUpdate
// method.
type AzureCredentialClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AzureCredentialsClientDeleteOptions contains the optional parameters for the AzureCredentialsClient.Delete method.
type AzureCredentialsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AzureCredentialsClientGetOptions contains the optional parameters for the AzureCredentialsClient.Get method.
type AzureCredentialsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureCredentialsClientListOptions contains the optional parameters for the AzureCredentialsClient.List method.
type AzureCredentialsClientListOptions struct {
	// placeholder for future optional parameters
}

type AzureServicePrincipalProperties struct {
	// REQUIRED; clientId when the CredentialKind is ServicePrincipal
	ClientID *string `json:"clientId,omitempty"`

	// REQUIRED; The kind of secret
	Kind *string `json:"kind,omitempty"`

	// REQUIRED; secret when the CredentialKind is ServicePrincipal
	Secret *string `json:"secret,omitempty"`

	// REQUIRED
	Storage *CredentialResourcePropertiesStorage `json:"storage,omitempty"`

	// REQUIRED; tenantId when the CredentialKind is ServicePrincipal
	TenantID *string `json:"tenantId,omitempty"`
}

// GetCredentialResourceProperties implements the CredentialResourcePropertiesClassification interface for type AzureServicePrincipalProperties.
func (a *AzureServicePrincipalProperties) GetCredentialResourceProperties() *CredentialResourceProperties {
	return &CredentialResourceProperties{
		Kind: a.Kind,
		Storage: a.Storage,
	}
}

// CredentialResource - Credential to a plane instance
type CredentialResource struct {
	// REQUIRED; Credential properties
	Properties CredentialResourcePropertiesClassification `json:"properties,omitempty"`
}

// CredentialResourceList - The list of credentials.
type CredentialResourceList struct {
	// The list of credentials.
	Value []*CredentialResource `json:"value,omitempty"`
}

// CredentialResourcePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetCredentialResourceProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AWSCredentialProperties, *AzureServicePrincipalProperties, *CredentialResourceProperties
type CredentialResourcePropertiesClassification interface {
	// GetCredentialResourceProperties returns the CredentialResourceProperties content of the underlying type.
	GetCredentialResourceProperties() *CredentialResourceProperties
}

// CredentialResourceProperties - Credential properties
type CredentialResourceProperties struct {
	// REQUIRED; The kind of secret
	Kind *string `json:"kind,omitempty"`

	// REQUIRED
	Storage *CredentialResourcePropertiesStorage `json:"storage,omitempty"`
}

// GetCredentialResourceProperties implements the CredentialResourcePropertiesClassification interface for type CredentialResourceProperties.
func (c *CredentialResourceProperties) GetCredentialResourceProperties() *CredentialResourceProperties { return c }

type CredentialResourcePropertiesStorage struct {
	// REQUIRED; credential store kinds supported.
	Kind *CredentialStorageKind `json:"kind,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// PlaneResource - UCP Plane.
type PlaneResource struct {
	// REQUIRED; UCP Plane properties
	Properties *PlaneResourceProperties `json:"properties,omitempty"`
}

// PlaneResourceList - The list of planes.
type PlaneResourceList struct {
	// The list of planes.
	Value []*PlaneResource `json:"value,omitempty"`
}

// PlaneResourceProperties - UCP Plane properties
type PlaneResourceProperties struct {
	// REQUIRED; Plane kinds supported.
	Kind *PlaneKind `json:"kind,omitempty"`

	// Resource Providers for UCP Native Plane
	ResourceProviders map[string]*string `json:"resourceProviders,omitempty"`

	// URL to forward requests to for non UCP Native Plane
	URL *string `json:"url,omitempty"`
}

// PlanesClientCreateOrUpdateOptions contains the optional parameters for the PlanesClient.CreateOrUpdate method.
type PlanesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PlanesClientDeleteOptions contains the optional parameters for the PlanesClient.Delete method.
type PlanesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PlanesClientGetOptions contains the optional parameters for the PlanesClient.Get method.
type PlanesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PlanesClientListOptions contains the optional parameters for the PlanesClient.List method.
type PlanesClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceGroupResource - UCP ResourceGroup.
type ResourceGroupResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceGroupResourceList - The list of resource groups.
type ResourceGroupResourceList struct {
	// The list of resource groups.
	Value []*ResourceGroupResource `json:"value,omitempty"`
}

// ResourceGroupsClientCreateOrUpdateOptions contains the optional parameters for the ResourceGroupsClient.CreateOrUpdate
// method.
type ResourceGroupsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ResourceGroupsClientDeleteOptions contains the optional parameters for the ResourceGroupsClient.Delete method.
type ResourceGroupsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ResourceGroupsClientGetOptions contains the optional parameters for the ResourceGroupsClient.Get method.
type ResourceGroupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ResourceGroupsClientListOptions contains the optional parameters for the ResourceGroupsClient.List method.
type ResourceGroupsClientListOptions struct {
	// placeholder for future optional parameters
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

