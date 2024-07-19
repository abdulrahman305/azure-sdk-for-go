// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmongocluster

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
	moduleVersion = "v0.1.0"
)

// ActionType - Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	// ActionTypeInternal - Actions are for internal-only APIs.
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CheckNameAvailabilityReason - Possible reasons for a name not being available.
type CheckNameAvailabilityReason string

const (
	// CheckNameAvailabilityReasonAlreadyExists - Name already exists.
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	// CheckNameAvailabilityReasonInvalid - Name is invalid.
	CheckNameAvailabilityReasonInvalid CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// CreateMode - The mode that the Mongo Cluster is created with.
type CreateMode string

const (
	// CreateModeDefault - Create a new mongo cluster.
	CreateModeDefault CreateMode = "Default"
	// CreateModePointInTimeRestore - Create a mongo cluster from a restore point-in-time.
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeDefault,
		CreateModePointInTimeRestore,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
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

// MongoClusterStatus - The status of the Mongo cluster resource.
type MongoClusterStatus string

const (
	// MongoClusterStatusDropping - The mongo cluster resource is being dropped.
	MongoClusterStatusDropping MongoClusterStatus = "Dropping"
	// MongoClusterStatusProvisioning - The mongo cluster resource is being provisioned.
	MongoClusterStatusProvisioning MongoClusterStatus = "Provisioning"
	// MongoClusterStatusReady - The mongo cluster resource is ready for use.
	MongoClusterStatusReady MongoClusterStatus = "Ready"
	// MongoClusterStatusStarting - The mongo cluster resource is being started.
	MongoClusterStatusStarting MongoClusterStatus = "Starting"
	// MongoClusterStatusStopped - The mongo cluster resource is stopped.
	MongoClusterStatusStopped MongoClusterStatus = "Stopped"
	// MongoClusterStatusStopping - The mongo cluster resource is being stopped.
	MongoClusterStatusStopping MongoClusterStatus = "Stopping"
	// MongoClusterStatusUpdating - The mongo cluster resource is being updated.
	MongoClusterStatusUpdating MongoClusterStatus = "Updating"
)

// PossibleMongoClusterStatusValues returns the possible values for the MongoClusterStatus const type.
func PossibleMongoClusterStatusValues() []MongoClusterStatus {
	return []MongoClusterStatus{
		MongoClusterStatusDropping,
		MongoClusterStatusProvisioning,
		MongoClusterStatusReady,
		MongoClusterStatusStarting,
		MongoClusterStatusStopped,
		MongoClusterStatusStopping,
		MongoClusterStatusUpdating,
	}
}

// NodeKind - The kind of the node on the cluster.
type NodeKind string

const (
	// NodeKindShard - The node is a shard kind.
	NodeKindShard NodeKind = "Shard"
)

// PossibleNodeKindValues returns the possible values for the NodeKind const type.
func PossibleNodeKindValues() []NodeKind {
	return []NodeKind{
		NodeKindShard,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	// OriginSystem - Indicates the operation is initiated by a system.
	OriginSystem Origin = "system"
	// OriginUser - Indicates the operation is initiated by a user.
	OriginUser Origin = "user"
	// OriginUserSystem - Indicates the operation is initiated by a user or system.
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// PrivateEndpointConnectionProvisioningStateCreating - Connection is being created
	PrivateEndpointConnectionProvisioningStateCreating PrivateEndpointConnectionProvisioningState = "Creating"
	// PrivateEndpointConnectionProvisioningStateDeleting - Connection is being deleted
	PrivateEndpointConnectionProvisioningStateDeleting PrivateEndpointConnectionProvisioningState = "Deleting"
	// PrivateEndpointConnectionProvisioningStateFailed - Connection provisioning has failed
	PrivateEndpointConnectionProvisioningStateFailed PrivateEndpointConnectionProvisioningState = "Failed"
	// PrivateEndpointConnectionProvisioningStateSucceeded - Connection has been provisioned
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	// PrivateEndpointServiceConnectionStatusApproved - Connection approved
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	// PrivateEndpointServiceConnectionStatusPending - Connectionaiting for approval or rejection
	PrivateEndpointServiceConnectionStatusPending PrivateEndpointServiceConnectionStatus = "Pending"
	// PrivateEndpointServiceConnectionStatusRejected - Connection Rejected
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ProvisioningState - The provisioning state of the last accepted operation.
type ProvisioningState string

const (
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDropping - A drop operation is in-progress on the resource.
	ProvisioningStateDropping ProvisioningState = "Dropping"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateInProgress - An operation is in-progress on the resource.
	ProvisioningStateInProgress ProvisioningState = "InProgress"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - An update operation is in-progress on the resource.
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateDropping,
		ProvisioningStateFailed,
		ProvisioningStateInProgress,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - Whether or not public endpoint access is allowed for this Mongo cluster. Value is optional and default
// value is 'Enabled'
type PublicNetworkAccess string

const (
	// PublicNetworkAccessDisabled - If set, the private endpoints are the exclusive access method.
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	// PublicNetworkAccessEnabled - If set, mongo cluster can be accessed through private and public methods.
	PublicNetworkAccessEnabled PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}
