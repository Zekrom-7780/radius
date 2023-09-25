//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20231001preview

const (
	moduleName = "v20231001preview"
	moduleVersion = "v0.0.1"
)

// AWSCredentialKind - AWS credential kind
type AWSCredentialKind string

const (
	// AWSCredentialKindAccessKey - The AWS Access Key credential
	AWSCredentialKindAccessKey AWSCredentialKind = "AccessKey"
)

// PossibleAWSCredentialKindValues returns the possible values for the AWSCredentialKind const type.
func PossibleAWSCredentialKindValues() []AWSCredentialKind {
	return []AWSCredentialKind{	
		AWSCredentialKindAccessKey,
	}
}

// AzureCredentialKind - Azure credential kinds supported.
type AzureCredentialKind string

const (
	// AzureCredentialKindServicePrincipal - The Service Principal Credential
	AzureCredentialKindServicePrincipal AzureCredentialKind = "ServicePrincipal"
)

// PossibleAzureCredentialKindValues returns the possible values for the AzureCredentialKind const type.
func PossibleAzureCredentialKindValues() []AzureCredentialKind {
	return []AzureCredentialKind{	
		AzureCredentialKindServicePrincipal,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication CreatedByType = "Application"
	CreatedByTypeKey CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{	
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// CredentialStorageKind - Credential store kinds supported.
type CredentialStorageKind string

const (
	// CredentialStorageKindInternal - Internal credential storage
	CredentialStorageKindInternal CredentialStorageKind = "Internal"
)

// PossibleCredentialStorageKindValues returns the possible values for the CredentialStorageKind const type.
func PossibleCredentialStorageKindValues() []CredentialStorageKind {
	return []CredentialStorageKind{	
		CredentialStorageKindInternal,
	}
}

// PlaneKind - Plane kinds supported.
type PlaneKind string

const (
	// PlaneKindAWS - AWS Plane
	PlaneKindAWS PlaneKind = "AWS"
	// PlaneKindAzure - Azure Plane
	PlaneKindAzure PlaneKind = "Azure"
	// PlaneKindUCPNative - UCP Native Plane
	PlaneKindUCPNative PlaneKind = "UCPNative"
)

// PossiblePlaneKindValues returns the possible values for the PlaneKind const type.
func PossiblePlaneKindValues() []PlaneKind {
	return []PlaneKind{	
		PlaneKindAWS,
		PlaneKindAzure,
		PlaneKindUCPNative,
	}
}

// ProvisioningState - Provisioning state of the portable resource at the time the operation was called
type ProvisioningState string

const (
	// ProvisioningStateAccepted - The resource create request has been accepted
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - The resource is being deleted
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - The resource is being provisioned
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - The resource is updating
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{	
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// Versions - Supported API versions for Universal Control Plane resource provider.
type Versions string

const (
	// VersionsV20231001Preview - 2023-10-01-preview
	VersionsV20231001Preview Versions = "2023-10-01-preview"
)

// PossibleVersionsValues returns the possible values for the Versions const type.
func PossibleVersionsValues() []Versions {
	return []Versions{	
		VersionsV20231001Preview,
	}
}

