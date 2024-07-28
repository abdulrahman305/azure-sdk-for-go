// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcomputefleet

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet"
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

// CachingTypes - Specifies the caching requirements.
type CachingTypes string

const (
	// CachingTypesNone - 'None' is default for Standard Storage
	CachingTypesNone CachingTypes = "None"
	// CachingTypesReadOnly - 'ReadOnly' is default for Premium Storage
	CachingTypesReadOnly CachingTypes = "ReadOnly"
	// CachingTypesReadWrite - 'ReadWrite' is default for OS Disk
	CachingTypesReadWrite CachingTypes = "ReadWrite"
)

// PossibleCachingTypesValues returns the possible values for the CachingTypes const type.
func PossibleCachingTypesValues() []CachingTypes {
	return []CachingTypes{
		CachingTypesNone,
		CachingTypesReadOnly,
		CachingTypesReadWrite,
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

// DeleteOptions - Specify what happens to the network interface when the VM is deleted
type DeleteOptions string

const (
	// DeleteOptionsDelete - Delete Option
	DeleteOptionsDelete DeleteOptions = "Delete"
	// DeleteOptionsDetach - Detach Option
	DeleteOptionsDetach DeleteOptions = "Detach"
)

// PossibleDeleteOptionsValues returns the possible values for the DeleteOptions const type.
func PossibleDeleteOptionsValues() []DeleteOptions {
	return []DeleteOptions{
		DeleteOptionsDelete,
		DeleteOptionsDetach,
	}
}

// DiffDiskOptions - Specifies the ephemeral disk option for operating system disk.
type DiffDiskOptions string

const (
	// DiffDiskOptionsLocal - Local Option.
	DiffDiskOptionsLocal DiffDiskOptions = "Local"
)

// PossibleDiffDiskOptionsValues returns the possible values for the DiffDiskOptions const type.
func PossibleDiffDiskOptionsValues() []DiffDiskOptions {
	return []DiffDiskOptions{
		DiffDiskOptionsLocal,
	}
}

// DiffDiskPlacement - Specifies the ephemeral disk placement for operating system disk. This property
// can be used by user in the request to choose the location i.e, cache disk or
// resource disk space for Ephemeral OS disk provisioning. For more information on
// Ephemeral OS disk size requirements, please refer Ephemeral OS disk size
// requirements for Windows VM at
// https://docs.microsoft.com/azure/virtual-machines/windows/ephemeral-os-disks#size-requirements
// and Linux VM at
// https://docs.microsoft.com/azure/virtual-machines/linux/ephemeral-os-disks#size-requirements
// Minimum api-version for NvmeDisk: 2024-03-01.
type DiffDiskPlacement string

const (
	// DiffDiskPlacementCacheDisk - CacheDisk option.
	DiffDiskPlacementCacheDisk DiffDiskPlacement = "CacheDisk"
	// DiffDiskPlacementNvmeDisk - NvmeDisk option.
	DiffDiskPlacementNvmeDisk DiffDiskPlacement = "NvmeDisk"
	// DiffDiskPlacementResourceDisk - Resource Disk option.
	DiffDiskPlacementResourceDisk DiffDiskPlacement = "ResourceDisk"
)

// PossibleDiffDiskPlacementValues returns the possible values for the DiffDiskPlacement const type.
func PossibleDiffDiskPlacementValues() []DiffDiskPlacement {
	return []DiffDiskPlacement{
		DiffDiskPlacementCacheDisk,
		DiffDiskPlacementNvmeDisk,
		DiffDiskPlacementResourceDisk,
	}
}

// DiskControllerTypes - Specifies the disk controller type configured for the VM and
// VirtualMachineScaleSet. This property is only supported for virtual machines
// whose operating system disk and VM sku supports Generation 2
// (https://docs.microsoft.com/en-us/azure/virtual-machines/generation-2), please
// check the HyperVGenerations capability returned as part of VM sku capabilities
// in the response of Microsoft.Compute SKUs api for the region contains V2
// (https://docs.microsoft.com/rest/api/compute/resourceskus/list). For more
// information about Disk Controller Types supported please refer to
// https://aka.ms/azure-diskcontrollertypes.
type DiskControllerTypes string

const (
	// DiskControllerTypesNVMe - NVMe disk type
	DiskControllerTypesNVMe DiskControllerTypes = "NVMe"
	// DiskControllerTypesSCSI - SCSI disk type
	DiskControllerTypesSCSI DiskControllerTypes = "SCSI"
)

// PossibleDiskControllerTypesValues returns the possible values for the DiskControllerTypes const type.
func PossibleDiskControllerTypesValues() []DiskControllerTypes {
	return []DiskControllerTypes{
		DiskControllerTypesNVMe,
		DiskControllerTypesSCSI,
	}
}

// DiskCreateOptionTypes - Specifies how the virtual machine should be created.
type DiskCreateOptionTypes string

const (
	// DiskCreateOptionTypesAttach - This value is used when you are using a specialized disk to create the virtual machine.
	DiskCreateOptionTypesAttach DiskCreateOptionTypes = "Attach"
	// DiskCreateOptionTypesCopy - This value is used to create a data disk from a snapshot or another disk.
	DiskCreateOptionTypesCopy DiskCreateOptionTypes = "Copy"
	// DiskCreateOptionTypesEmpty - This value is used when creating an empty data disk.
	DiskCreateOptionTypesEmpty DiskCreateOptionTypes = "Empty"
	// DiskCreateOptionTypesFromImage - This value is used when you are using an image to create the virtual machine.
	// If you are using a platform image, you also use the imageReference element
	// described above. If you are using a marketplace image, you also use the
	// plan element previously described.
	DiskCreateOptionTypesFromImage DiskCreateOptionTypes = "FromImage"
	// DiskCreateOptionTypesRestore - This value is used to create a data disk from a disk restore point.
	DiskCreateOptionTypesRestore DiskCreateOptionTypes = "Restore"
)

// PossibleDiskCreateOptionTypesValues returns the possible values for the DiskCreateOptionTypes const type.
func PossibleDiskCreateOptionTypesValues() []DiskCreateOptionTypes {
	return []DiskCreateOptionTypes{
		DiskCreateOptionTypesAttach,
		DiskCreateOptionTypesCopy,
		DiskCreateOptionTypesEmpty,
		DiskCreateOptionTypesFromImage,
		DiskCreateOptionTypesRestore,
	}
}

// DiskDeleteOptionTypes - Specifies the behavior of the managed disk when the VM gets deleted, for
// example whether the managed disk is deleted or detached. Supported values are:
// **Delete.** If this value is used, the managed disk is deleted when VM gets
// deleted. **Detach.** If this value is used, the managed disk is retained after
// VM gets deleted. Minimum api-version: 2021-03-01.
type DiskDeleteOptionTypes string

const (
	// DiskDeleteOptionTypesDelete - If this value is used, the managed disk is deleted when VM gets deleted.
	DiskDeleteOptionTypesDelete DiskDeleteOptionTypes = "Delete"
	// DiskDeleteOptionTypesDetach - If this value is used, the managed disk is retained after VM gets deleted.
	DiskDeleteOptionTypesDetach DiskDeleteOptionTypes = "Detach"
)

// PossibleDiskDeleteOptionTypesValues returns the possible values for the DiskDeleteOptionTypes const type.
func PossibleDiskDeleteOptionTypesValues() []DiskDeleteOptionTypes {
	return []DiskDeleteOptionTypes{
		DiskDeleteOptionTypesDelete,
		DiskDeleteOptionTypesDetach,
	}
}

// DomainNameLabelScopeTypes - The Domain name label scope.The concatenation of the hashed domain name label
// that generated according to the policy from domain name label scope and vm
// index will be the domain name labels of the PublicIPAddress resources that will
// be created
type DomainNameLabelScopeTypes string

const (
	// DomainNameLabelScopeTypesNoReuse - NoReuse type
	DomainNameLabelScopeTypesNoReuse DomainNameLabelScopeTypes = "NoReuse"
	// DomainNameLabelScopeTypesResourceGroupReuse - ResourceGroupReuse type
	DomainNameLabelScopeTypesResourceGroupReuse DomainNameLabelScopeTypes = "ResourceGroupReuse"
	// DomainNameLabelScopeTypesSubscriptionReuse - SubscriptionReuse type
	DomainNameLabelScopeTypesSubscriptionReuse DomainNameLabelScopeTypes = "SubscriptionReuse"
	// DomainNameLabelScopeTypesTenantReuse - TenantReuse type
	DomainNameLabelScopeTypesTenantReuse DomainNameLabelScopeTypes = "TenantReuse"
)

// PossibleDomainNameLabelScopeTypesValues returns the possible values for the DomainNameLabelScopeTypes const type.
func PossibleDomainNameLabelScopeTypesValues() []DomainNameLabelScopeTypes {
	return []DomainNameLabelScopeTypes{
		DomainNameLabelScopeTypesNoReuse,
		DomainNameLabelScopeTypesResourceGroupReuse,
		DomainNameLabelScopeTypesSubscriptionReuse,
		DomainNameLabelScopeTypesTenantReuse,
	}
}

// EvictionPolicy - Different kind of eviction policies
type EvictionPolicy string

const (
	// EvictionPolicyDeallocate - When evicted, the Spot VM will be deallocated/stopped
	EvictionPolicyDeallocate EvictionPolicy = "Deallocate"
	// EvictionPolicyDelete - When evicted, the Spot VM will be deleted and the corresponding capacity will be updated to reflect
	// this.
	EvictionPolicyDelete EvictionPolicy = "Delete"
)

// PossibleEvictionPolicyValues returns the possible values for the EvictionPolicy const type.
func PossibleEvictionPolicyValues() []EvictionPolicy {
	return []EvictionPolicy{
		EvictionPolicyDeallocate,
		EvictionPolicyDelete,
	}
}

// IPVersion - Available from Api-Version 2017-03-30 onwards, it represents whether the
// specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4. Possible
// values are: 'IPv4' and 'IPv6'.
type IPVersion string

const (
	// IPVersionIPv4 - IPv4 version
	IPVersionIPv4 IPVersion = "IPv4"
	// IPVersionIPv6 - IPv6 version
	IPVersionIPv6 IPVersion = "IPv6"
)

// PossibleIPVersionValues returns the possible values for the IPVersion const type.
func PossibleIPVersionValues() []IPVersion {
	return []IPVersion{
		IPVersionIPv4,
		IPVersionIPv6,
	}
}

// LinuxPatchAssessmentMode - Specifies the mode of VM Guest Patch Assessment for the IaaS virtual
// machine.<br /><br /> Possible values are:<br /><br /> **ImageDefault** - You
// control the timing of patch assessments on a virtual machine. <br /><br />
// **AutomaticByPlatform** - The platform will trigger periodic patch assessments.
// The property provisionVMAgent must be true.
type LinuxPatchAssessmentMode string

const (
	// LinuxPatchAssessmentModeAutomaticByPlatform - The platform will trigger periodic patch assessments.The property provisionVMAgent
	// must be true.
	LinuxPatchAssessmentModeAutomaticByPlatform LinuxPatchAssessmentMode = "AutomaticByPlatform"
	// LinuxPatchAssessmentModeImageDefault - You control the timing of patch assessments on a virtual machine.
	LinuxPatchAssessmentModeImageDefault LinuxPatchAssessmentMode = "ImageDefault"
)

// PossibleLinuxPatchAssessmentModeValues returns the possible values for the LinuxPatchAssessmentMode const type.
func PossibleLinuxPatchAssessmentModeValues() []LinuxPatchAssessmentMode {
	return []LinuxPatchAssessmentMode{
		LinuxPatchAssessmentModeAutomaticByPlatform,
		LinuxPatchAssessmentModeImageDefault,
	}
}

// LinuxVMGuestPatchAutomaticByPlatformRebootSetting - Specifies the reboot setting for all AutomaticByPlatform patch installation
// operations.
type LinuxVMGuestPatchAutomaticByPlatformRebootSetting string

const (
	// LinuxVMGuestPatchAutomaticByPlatformRebootSettingAlways - Always Reboot setting
	LinuxVMGuestPatchAutomaticByPlatformRebootSettingAlways LinuxVMGuestPatchAutomaticByPlatformRebootSetting = "Always"
	// LinuxVMGuestPatchAutomaticByPlatformRebootSettingIfRequired - IfRequired Reboot setting
	LinuxVMGuestPatchAutomaticByPlatformRebootSettingIfRequired LinuxVMGuestPatchAutomaticByPlatformRebootSetting = "IfRequired"
	// LinuxVMGuestPatchAutomaticByPlatformRebootSettingNever - Never Reboot setting
	LinuxVMGuestPatchAutomaticByPlatformRebootSettingNever LinuxVMGuestPatchAutomaticByPlatformRebootSetting = "Never"
	// LinuxVMGuestPatchAutomaticByPlatformRebootSettingUnknown - Unknown Reboot setting
	LinuxVMGuestPatchAutomaticByPlatformRebootSettingUnknown LinuxVMGuestPatchAutomaticByPlatformRebootSetting = "Unknown"
)

// PossibleLinuxVMGuestPatchAutomaticByPlatformRebootSettingValues returns the possible values for the LinuxVMGuestPatchAutomaticByPlatformRebootSetting const type.
func PossibleLinuxVMGuestPatchAutomaticByPlatformRebootSettingValues() []LinuxVMGuestPatchAutomaticByPlatformRebootSetting {
	return []LinuxVMGuestPatchAutomaticByPlatformRebootSetting{
		LinuxVMGuestPatchAutomaticByPlatformRebootSettingAlways,
		LinuxVMGuestPatchAutomaticByPlatformRebootSettingIfRequired,
		LinuxVMGuestPatchAutomaticByPlatformRebootSettingNever,
		LinuxVMGuestPatchAutomaticByPlatformRebootSettingUnknown,
	}
}

// LinuxVMGuestPatchMode - Specifies the mode of VM Guest Patching to IaaS virtual machine or virtual
// machines associated to virtual machine scale set with OrchestrationMode as
// Flexible.
type LinuxVMGuestPatchMode string

const (
	// LinuxVMGuestPatchModeAutomaticByPlatform - The virtual machine will be automatically updated by the platform. The property
	// provisionVMAgent must be true.
	LinuxVMGuestPatchModeAutomaticByPlatform LinuxVMGuestPatchMode = "AutomaticByPlatform"
	// LinuxVMGuestPatchModeImageDefault - The virtual machine's default patching configuration is used.
	LinuxVMGuestPatchModeImageDefault LinuxVMGuestPatchMode = "ImageDefault"
)

// PossibleLinuxVMGuestPatchModeValues returns the possible values for the LinuxVMGuestPatchMode const type.
func PossibleLinuxVMGuestPatchModeValues() []LinuxVMGuestPatchMode {
	return []LinuxVMGuestPatchMode{
		LinuxVMGuestPatchModeAutomaticByPlatform,
		LinuxVMGuestPatchModeImageDefault,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	// ManagedServiceIdentityTypeNone - No managed identity.
	ManagedServiceIdentityTypeNone ManagedServiceIdentityType = "None"
	// ManagedServiceIdentityTypeSystemAndUserAssigned - System and user assigned managed identity.
	ManagedServiceIdentityTypeSystemAndUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	// ManagedServiceIdentityTypeSystemAssigned - System assigned managed identity.
	ManagedServiceIdentityTypeSystemAssigned ManagedServiceIdentityType = "SystemAssigned"
	// ManagedServiceIdentityTypeUserAssigned - User assigned managed identity.
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAndUserAssigned,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// Mode - Specifies the mode that ProxyAgent will execute on if the feature is enabled.
// ProxyAgent will start to audit or monitor but not enforce access control over
// requests to host endpoints in Audit mode, while in Enforce mode it will enforce
// access control. The default value is Enforce mode.
type Mode string

const (
	// ModeAudit - Audit Mode
	ModeAudit Mode = "Audit"
	// ModeEnforce - Enforce Mode
	ModeEnforce Mode = "Enforce"
)

// PossibleModeValues returns the possible values for the Mode const type.
func PossibleModeValues() []Mode {
	return []Mode{
		ModeAudit,
		ModeEnforce,
	}
}

// NetworkAPIVersion - specifies the Microsoft.Network API version used when creating networking
// resources in the Network Interface Configurations for Virtual Machine Scale Set
// with orchestration mode 'Flexible'
type NetworkAPIVersion string

const (
	// NetworkAPIVersion20201101 - Initial version supported. Later versions are supported as well.
	NetworkAPIVersion20201101 NetworkAPIVersion = "2020-11-01"
)

// PossibleNetworkAPIVersionValues returns the possible values for the NetworkAPIVersion const type.
func PossibleNetworkAPIVersionValues() []NetworkAPIVersion {
	return []NetworkAPIVersion{
		NetworkAPIVersion20201101,
	}
}

// NetworkInterfaceAuxiliaryMode - Specifies whether the Auxiliary mode is enabled for the Network Interface
// resource.
type NetworkInterfaceAuxiliaryMode string

const (
	// NetworkInterfaceAuxiliaryModeAcceleratedConnections - AcceleratedConnections Mode
	NetworkInterfaceAuxiliaryModeAcceleratedConnections NetworkInterfaceAuxiliaryMode = "AcceleratedConnections"
	// NetworkInterfaceAuxiliaryModeFloating - Floating Mode
	NetworkInterfaceAuxiliaryModeFloating NetworkInterfaceAuxiliaryMode = "Floating"
	// NetworkInterfaceAuxiliaryModeNone - None Mode
	NetworkInterfaceAuxiliaryModeNone NetworkInterfaceAuxiliaryMode = "None"
)

// PossibleNetworkInterfaceAuxiliaryModeValues returns the possible values for the NetworkInterfaceAuxiliaryMode const type.
func PossibleNetworkInterfaceAuxiliaryModeValues() []NetworkInterfaceAuxiliaryMode {
	return []NetworkInterfaceAuxiliaryMode{
		NetworkInterfaceAuxiliaryModeAcceleratedConnections,
		NetworkInterfaceAuxiliaryModeFloating,
		NetworkInterfaceAuxiliaryModeNone,
	}
}

// NetworkInterfaceAuxiliarySKU - Specifies whether the Auxiliary sku is enabled for the Network Interface
// resource.
type NetworkInterfaceAuxiliarySKU string

const (
	// NetworkInterfaceAuxiliarySKUA1 - A1 sku
	NetworkInterfaceAuxiliarySKUA1 NetworkInterfaceAuxiliarySKU = "A1"
	// NetworkInterfaceAuxiliarySKUA2 - A2 sku
	NetworkInterfaceAuxiliarySKUA2 NetworkInterfaceAuxiliarySKU = "A2"
	// NetworkInterfaceAuxiliarySKUA4 - A4 sku
	NetworkInterfaceAuxiliarySKUA4 NetworkInterfaceAuxiliarySKU = "A4"
	// NetworkInterfaceAuxiliarySKUA8 - A8 sku
	NetworkInterfaceAuxiliarySKUA8 NetworkInterfaceAuxiliarySKU = "A8"
	// NetworkInterfaceAuxiliarySKUNone - no sku
	NetworkInterfaceAuxiliarySKUNone NetworkInterfaceAuxiliarySKU = "None"
)

// PossibleNetworkInterfaceAuxiliarySKUValues returns the possible values for the NetworkInterfaceAuxiliarySKU const type.
func PossibleNetworkInterfaceAuxiliarySKUValues() []NetworkInterfaceAuxiliarySKU {
	return []NetworkInterfaceAuxiliarySKU{
		NetworkInterfaceAuxiliarySKUA1,
		NetworkInterfaceAuxiliarySKUA2,
		NetworkInterfaceAuxiliarySKUA4,
		NetworkInterfaceAuxiliarySKUA8,
		NetworkInterfaceAuxiliarySKUNone,
	}
}

// OperatingSystemTypes - This property allows you to specify the type of the OS that is included in the
// disk if creating a VM from user-image or a specialized VHD. Possible values
// are: **Windows,** **Linux.**
type OperatingSystemTypes string

const (
	// OperatingSystemTypesLinux - Linux OS type
	OperatingSystemTypesLinux OperatingSystemTypes = "Linux"
	// OperatingSystemTypesWindows - Windows OS type
	OperatingSystemTypesWindows OperatingSystemTypes = "Windows"
)

// PossibleOperatingSystemTypesValues returns the possible values for the OperatingSystemTypes const type.
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return []OperatingSystemTypes{
		OperatingSystemTypesLinux,
		OperatingSystemTypesWindows,
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

// ProtocolTypes - Specifies the protocol of WinRM listener. Possible values are: **http,**
// **https.**
type ProtocolTypes string

const (
	// ProtocolTypesHTTP - Http protocol
	ProtocolTypesHTTP ProtocolTypes = "Http"
	// ProtocolTypesHTTPS - Https protocol
	ProtocolTypesHTTPS ProtocolTypes = "Https"
)

// PossibleProtocolTypesValues returns the possible values for the ProtocolTypes const type.
func PossibleProtocolTypesValues() []ProtocolTypes {
	return []ProtocolTypes{
		ProtocolTypesHTTP,
		ProtocolTypesHTTPS,
	}
}

// ProvisioningState - The status of the current operation.
type ProvisioningState string

const (
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating - Initial creation in progress.
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting - Deletion in progress.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateMigrating - Resource is being migrated from one subscription or resource group to another.
	ProvisioningStateMigrating ProvisioningState = "Migrating"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Update in progress.
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMigrating,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicIPAddressSKUName - Specify public IP sku name.
type PublicIPAddressSKUName string

const (
	// PublicIPAddressSKUNameBasic - Basic sku name
	PublicIPAddressSKUNameBasic PublicIPAddressSKUName = "Basic"
	// PublicIPAddressSKUNameStandard - Standard sku name
	PublicIPAddressSKUNameStandard PublicIPAddressSKUName = "Standard"
)

// PossiblePublicIPAddressSKUNameValues returns the possible values for the PublicIPAddressSKUName const type.
func PossiblePublicIPAddressSKUNameValues() []PublicIPAddressSKUName {
	return []PublicIPAddressSKUName{
		PublicIPAddressSKUNameBasic,
		PublicIPAddressSKUNameStandard,
	}
}

// PublicIPAddressSKUTier - Specify public IP sku tier
type PublicIPAddressSKUTier string

const (
	// PublicIPAddressSKUTierGlobal - Global sku tier
	PublicIPAddressSKUTierGlobal PublicIPAddressSKUTier = "Global"
	// PublicIPAddressSKUTierRegional - Regional sku tier
	PublicIPAddressSKUTierRegional PublicIPAddressSKUTier = "Regional"
)

// PossiblePublicIPAddressSKUTierValues returns the possible values for the PublicIPAddressSKUTier const type.
func PossiblePublicIPAddressSKUTierValues() []PublicIPAddressSKUTier {
	return []PublicIPAddressSKUTier{
		PublicIPAddressSKUTierGlobal,
		PublicIPAddressSKUTierRegional,
	}
}

// RegularPriorityAllocationStrategy - Regular VM Allocation strategy types for Compute Fleet
type RegularPriorityAllocationStrategy string

const (
	// RegularPriorityAllocationStrategyLowestPrice - Default. VM sizes distribution will be determined to optimize for price.
	RegularPriorityAllocationStrategyLowestPrice RegularPriorityAllocationStrategy = "LowestPrice"
	// RegularPriorityAllocationStrategyPrioritized - VM sizes distribution will be determined to optimize for the 'priority'
	// as specified for each vm size.
	RegularPriorityAllocationStrategyPrioritized RegularPriorityAllocationStrategy = "Prioritized"
)

// PossibleRegularPriorityAllocationStrategyValues returns the possible values for the RegularPriorityAllocationStrategy const type.
func PossibleRegularPriorityAllocationStrategyValues() []RegularPriorityAllocationStrategy {
	return []RegularPriorityAllocationStrategy{
		RegularPriorityAllocationStrategyLowestPrice,
		RegularPriorityAllocationStrategyPrioritized,
	}
}

// SecurityEncryptionTypes - Specifies the EncryptionType of the managed disk.
// **Note:** It can be set for only Confidential VMs.
type SecurityEncryptionTypes string

const (
	// SecurityEncryptionTypesDiskWithVMGuestState - EncryptionType of the managed disk is set to DiskWithVMGuestState for encryption
	// of the managed disk along with VMGuestState blob.
	SecurityEncryptionTypesDiskWithVMGuestState SecurityEncryptionTypes = "DiskWithVMGuestState"
	// SecurityEncryptionTypesNonPersistedTPM - EncryptionType of the managed disk is set to NonPersistedTPM for not persisting
	// firmware state in the VMGuestState blob.
	SecurityEncryptionTypesNonPersistedTPM SecurityEncryptionTypes = "NonPersistedTPM"
	// SecurityEncryptionTypesVMGuestStateOnly - EncryptionType of the managed disk is set to VMGuestStateOnly for encryption
	// of just the VMGuestState blob.
	SecurityEncryptionTypesVMGuestStateOnly SecurityEncryptionTypes = "VMGuestStateOnly"
)

// PossibleSecurityEncryptionTypesValues returns the possible values for the SecurityEncryptionTypes const type.
func PossibleSecurityEncryptionTypesValues() []SecurityEncryptionTypes {
	return []SecurityEncryptionTypes{
		SecurityEncryptionTypesDiskWithVMGuestState,
		SecurityEncryptionTypesNonPersistedTPM,
		SecurityEncryptionTypesVMGuestStateOnly,
	}
}

// SecurityTypes - Specifies the SecurityType of the virtual machine. It has to be set to any
// specified value to enable UefiSettings. The default behavior is: UefiSettings
// will not be enabled unless this property is set.
type SecurityTypes string

const (
	// SecurityTypesConfidentialVM - ConfidentialVM security type
	SecurityTypesConfidentialVM SecurityTypes = "ConfidentialVM"
	// SecurityTypesTrustedLaunch - TrustedLaunch security type
	SecurityTypesTrustedLaunch SecurityTypes = "TrustedLaunch"
)

// PossibleSecurityTypesValues returns the possible values for the SecurityTypes const type.
func PossibleSecurityTypesValues() []SecurityTypes {
	return []SecurityTypes{
		SecurityTypesConfidentialVM,
		SecurityTypesTrustedLaunch,
	}
}

// SettingNames - Specifies the name of the setting to which the content applies. Possible values
// are: FirstLogonCommands and AutoLogon.
type SettingNames string

const (
	// SettingNamesAutoLogon - AutoLogon setting
	SettingNamesAutoLogon SettingNames = "AutoLogon"
	// SettingNamesFirstLogonCommands - FirstLogonCommands setting
	SettingNamesFirstLogonCommands SettingNames = "FirstLogonCommands"
)

// PossibleSettingNamesValues returns the possible values for the SettingNames const type.
func PossibleSettingNamesValues() []SettingNames {
	return []SettingNames{
		SettingNamesAutoLogon,
		SettingNamesFirstLogonCommands,
	}
}

// SpotAllocationStrategy - Spot allocation strategy types for Compute Fleet
type SpotAllocationStrategy string

const (
	// SpotAllocationStrategyCapacityOptimized - VM sizes distribution will be determined to optimize for capacity.
	SpotAllocationStrategyCapacityOptimized SpotAllocationStrategy = "CapacityOptimized"
	// SpotAllocationStrategyLowestPrice - VM sizes distribution will be determined to optimize for price. Note: Capacity will
	// still be considered here but will be given much less weight.
	SpotAllocationStrategyLowestPrice SpotAllocationStrategy = "LowestPrice"
	// SpotAllocationStrategyPriceCapacityOptimized - Default. VM sizes distribution will be determined to optimize for both price
	// and capacity.
	SpotAllocationStrategyPriceCapacityOptimized SpotAllocationStrategy = "PriceCapacityOptimized"
)

// PossibleSpotAllocationStrategyValues returns the possible values for the SpotAllocationStrategy const type.
func PossibleSpotAllocationStrategyValues() []SpotAllocationStrategy {
	return []SpotAllocationStrategy{
		SpotAllocationStrategyCapacityOptimized,
		SpotAllocationStrategyLowestPrice,
		SpotAllocationStrategyPriceCapacityOptimized,
	}
}

// StorageAccountTypes - Specifies the storage account type for the managed disk. Managed OS disk
// storage account type can only be set when you create the scale set. NOTE:
// UltraSSD_LRS can only be used with data disks. It cannot be used with OS Disk.
// Standard_LRS uses Standard HDD. StandardSSD_LRS uses Standard SSD. Premium_LRS
// uses Premium SSD. UltraSSD_LRS uses Ultra disk. Premium_ZRS uses Premium SSD
// zone redundant storage. StandardSSD_ZRS uses Standard SSD zone redundant
// storage. For more information regarding disks supported for Windows Virtual
// Machines, refer to
// https://docs.microsoft.com/azure/virtual-machines/windows/disks-types and, for
// Linux Virtual Machines, refer to
// https://docs.microsoft.com/azure/virtual-machines/linux/disks-types
type StorageAccountTypes string

const (
	// StorageAccountTypesPremiumLRS - Premium_LRS option.
	StorageAccountTypesPremiumLRS StorageAccountTypes = "Premium_LRS"
	// StorageAccountTypesPremiumV2LRS - PremiumV2_LRS option.
	StorageAccountTypesPremiumV2LRS StorageAccountTypes = "PremiumV2_LRS"
	// StorageAccountTypesPremiumZRS - Premium_ZRS option.
	StorageAccountTypesPremiumZRS StorageAccountTypes = "Premium_ZRS"
	// StorageAccountTypesStandardLRS - Standard_LRS option.
	StorageAccountTypesStandardLRS StorageAccountTypes = "Standard_LRS"
	// StorageAccountTypesStandardSSDLRS - StandardSSD_LRS option.
	StorageAccountTypesStandardSSDLRS StorageAccountTypes = "StandardSSD_LRS"
	// StorageAccountTypesStandardSSDZRS - StandardSSD_ZRS option.
	StorageAccountTypesStandardSSDZRS StorageAccountTypes = "StandardSSD_ZRS"
	// StorageAccountTypesUltraSSDLRS - UltraSSD_LRS option.
	StorageAccountTypesUltraSSDLRS StorageAccountTypes = "UltraSSD_LRS"
)

// PossibleStorageAccountTypesValues returns the possible values for the StorageAccountTypes const type.
func PossibleStorageAccountTypesValues() []StorageAccountTypes {
	return []StorageAccountTypes{
		StorageAccountTypesPremiumLRS,
		StorageAccountTypesPremiumV2LRS,
		StorageAccountTypesPremiumZRS,
		StorageAccountTypesStandardLRS,
		StorageAccountTypesStandardSSDLRS,
		StorageAccountTypesStandardSSDZRS,
		StorageAccountTypesUltraSSDLRS,
	}
}

// WindowsPatchAssessmentMode - Specifies the mode of VM Guest patch assessment for the IaaS virtual machine.
type WindowsPatchAssessmentMode string

const (
	// WindowsPatchAssessmentModeAutomaticByPlatform - The platform will trigger periodic patch assessments. The property provisionVMAgent
	// must be true.
	WindowsPatchAssessmentModeAutomaticByPlatform WindowsPatchAssessmentMode = "AutomaticByPlatform"
	// WindowsPatchAssessmentModeImageDefault - You control the timing of patch assessments on a virtual machine.
	WindowsPatchAssessmentModeImageDefault WindowsPatchAssessmentMode = "ImageDefault"
)

// PossibleWindowsPatchAssessmentModeValues returns the possible values for the WindowsPatchAssessmentMode const type.
func PossibleWindowsPatchAssessmentModeValues() []WindowsPatchAssessmentMode {
	return []WindowsPatchAssessmentMode{
		WindowsPatchAssessmentModeAutomaticByPlatform,
		WindowsPatchAssessmentModeImageDefault,
	}
}

// WindowsVMGuestPatchAutomaticByPlatformRebootSetting - Specifies the reboot setting for all AutomaticByPlatform patch installation
// operations.
type WindowsVMGuestPatchAutomaticByPlatformRebootSetting string

const (
	// WindowsVMGuestPatchAutomaticByPlatformRebootSettingAlways - Always Reboot setting
	WindowsVMGuestPatchAutomaticByPlatformRebootSettingAlways WindowsVMGuestPatchAutomaticByPlatformRebootSetting = "Always"
	// WindowsVMGuestPatchAutomaticByPlatformRebootSettingIfRequired - IfRequired Reboot setting
	WindowsVMGuestPatchAutomaticByPlatformRebootSettingIfRequired WindowsVMGuestPatchAutomaticByPlatformRebootSetting = "IfRequired"
	// WindowsVMGuestPatchAutomaticByPlatformRebootSettingNever - Never Reboot setting
	WindowsVMGuestPatchAutomaticByPlatformRebootSettingNever WindowsVMGuestPatchAutomaticByPlatformRebootSetting = "Never"
	// WindowsVMGuestPatchAutomaticByPlatformRebootSettingUnknown - Unknown Reboot setting
	WindowsVMGuestPatchAutomaticByPlatformRebootSettingUnknown WindowsVMGuestPatchAutomaticByPlatformRebootSetting = "Unknown"
)

// PossibleWindowsVMGuestPatchAutomaticByPlatformRebootSettingValues returns the possible values for the WindowsVMGuestPatchAutomaticByPlatformRebootSetting const type.
func PossibleWindowsVMGuestPatchAutomaticByPlatformRebootSettingValues() []WindowsVMGuestPatchAutomaticByPlatformRebootSetting {
	return []WindowsVMGuestPatchAutomaticByPlatformRebootSetting{
		WindowsVMGuestPatchAutomaticByPlatformRebootSettingAlways,
		WindowsVMGuestPatchAutomaticByPlatformRebootSettingIfRequired,
		WindowsVMGuestPatchAutomaticByPlatformRebootSettingNever,
		WindowsVMGuestPatchAutomaticByPlatformRebootSettingUnknown,
	}
}

// WindowsVMGuestPatchMode - Specifies the mode of VM Guest Patching to IaaS virtual machine or virtual
// machines associated to virtual machine scale set with OrchestrationMode as
// Flexible.
type WindowsVMGuestPatchMode string

const (
	// WindowsVMGuestPatchModeAutomaticByOS - The virtual machine will automatically be updated by the OS.
	// The property WindowsConfiguration.enableAutomaticUpdates must be true.
	WindowsVMGuestPatchModeAutomaticByOS WindowsVMGuestPatchMode = "AutomaticByOS"
	// WindowsVMGuestPatchModeAutomaticByPlatform - The virtual machine will automatically updated by the platform. The properties
	// provisionVMAgent and WindowsConfiguration.enableAutomaticUpdates must be true.
	WindowsVMGuestPatchModeAutomaticByPlatform WindowsVMGuestPatchMode = "AutomaticByPlatform"
	// WindowsVMGuestPatchModeManual - You control the application of patches to a virtual machine.
	// You do this by applying patches manually inside the VM. In this mode,
	// automatic updates are disabled; the property WindowsConfiguration.enableAutomaticUpdates
	// must be false
	WindowsVMGuestPatchModeManual WindowsVMGuestPatchMode = "Manual"
)

// PossibleWindowsVMGuestPatchModeValues returns the possible values for the WindowsVMGuestPatchMode const type.
func PossibleWindowsVMGuestPatchModeValues() []WindowsVMGuestPatchMode {
	return []WindowsVMGuestPatchMode{
		WindowsVMGuestPatchModeAutomaticByOS,
		WindowsVMGuestPatchModeAutomaticByPlatform,
		WindowsVMGuestPatchModeManual,
	}
}
