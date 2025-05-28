# AccountResourcesItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedicatedTenancySupportEnabled** | Pointer to **bool** | Specifies if dedicated tenancy support is enabled | [optional] 
**DedicatedTenancyManagementCidrRange** | Pointer to **NullableString** | The CIDR range used in dedicated tenancy | [optional] 
**ResourceType** | [**AwsAccountResourceType**](AwsAccountResourceType.md) | AWS Resource type | 
**AccountType** | [**AccountType**](AccountType.md) | The type of account the resource is associated with | 
**AccountId** | Pointer to **NullableString** | Account Id | [optional] 
**ResourceId** | Pointer to **NullableString** | Resource Id | [optional] 
**ImageId** | Pointer to **NullableString** | Image Id | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** | Security Group Description | [optional] 
**Status** | Pointer to [**NullableAwsEdcSubnetStatus**](AwsEdcSubnetStatus.md) | Subnet Status  Enum values AVAILABLE, PENDING | [optional] 
**Location** | Pointer to **NullableString** | Image Location | [optional] 
**ImageType** | Pointer to [**NullableAwsEdcAmiImageType**](AwsEdcAmiImageType.md) | Image Type  Enum values KERNEL, MACHINE, RAMDISK | [optional] 
**PlatformType** | Pointer to [**NullableAwsEdcAmiImagePlatform**](AwsEdcAmiImagePlatform.md) | Vpc CIDR Block  Enum values WINDOWS | [optional] 
**BootMode** | Pointer to [**NullableAwsEdcAmiImageBootMode**](AwsEdcAmiImageBootMode.md) | Image Boot Type  Enum values LEGACYBIOS, UEFI | [optional] 
**ArchitectureType** | Pointer to [**NullableAwsEdcAmiImageArchitecture**](AwsEdcAmiImageArchitecture.md) | Image Architecture  Enum values ARM64, I386, X86_64 | [optional] 
**VirtualizationType** | Pointer to [**NullableAwsEdcAmiImageVirtualization**](AwsEdcAmiImageVirtualization.md) | Image Virtualization  Enum values HVM, PARAVIRTUAL | [optional] 
**Public** | Pointer to **NullableBool** | Image IsPublic | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Vpc Tags | [optional] 
**DirectoryId** | Pointer to **NullableString** | Directory ID | [optional] 
**Domain** | Pointer to **NullableString** | Directory Domain | [optional] 
**DefaultOU** | Pointer to **NullableString** | Directory Organizational Unit | [optional] 
**Type** | Pointer to [**NullableAwsEdcDirectoryType**](AwsEdcDirectoryType.md) | Directory Type  Enum values ADCONNECTOR, MICROSOFTAD, SIMPLEAD, SHAREDMICROSOFTAD | [optional] 
**Size** | Pointer to [**NullableAwsEdcDirectorySize**](AwsEdcDirectorySize.md) | Directory Size  Enum values SMALL, LARGE | [optional] 
**RegistrationStatus** | Pointer to [**NullableAwsEdcDirectoryRegistrationStatus**](AwsEdcDirectoryRegistrationStatus.md) | Directory Registration State  Enum values REGISTERED, REGISTERING, DEREGISTERED, DEREGISTERING, ERROR | [optional] 
**Tenancy** | Pointer to [**NullableAwsEdcDirectoryTenancy**](AwsEdcDirectoryTenancy.md) | Directory Registration Tenancy  Enum values SHARED, DEDICATED | [optional] 
**VpcId** | Pointer to **NullableString** | Security Group VpcId | [optional] 
**SecurityGroupId** | Pointer to **NullableString** | Security Group Id | [optional] 
**SubnetIds** | Pointer to **[]string** | Directory associated Subnets | [optional] 
**VpcTenancy** | Pointer to [**NullableAwsEdcVpcTenancy**](AwsEdcVpcTenancy.md) | Vpc Tenancy  Enum values DEFAULT, DEDICATED, HOST | [optional] 
**CidrBlock** | Pointer to **NullableString** | Subnet CIDR Block | [optional] 
**SubnetId** | Pointer to **NullableString** | Subnet Id | [optional] 
**SubnetArn** | Pointer to **NullableString** | Subnet Status | [optional] 
**AwsRegion** | Pointer to **NullableString** | The AWS region where the subnet is located | [optional] 
**AvailabilityZoneId** | Pointer to **NullableString** | Subnet Availability Zone Id | [optional] 
**AvailabilityZone** | Pointer to **NullableString** | Subnet Availability Zone | [optional] 
**IsSupportedByAwsManagedCapacity** | Pointer to **bool** | Indicates if the subnet can be used for Managed Capacity deployments | [optional] [readonly] 
**RegionNames** | Pointer to **[]string** | The names of the regions | [optional] 
**Regions** | Pointer to [**[]AwsWorkspaceRegion**](AwsWorkspaceRegion.md) | The list of regions including their names, codes and supported az ids | [optional] 
**InboundPermissions** | Pointer to [**[]IpPermission**](IpPermission.md) | Security Group Inbound Permissions | [optional] 
**OutboundPermissions** | Pointer to [**[]IpPermission**](IpPermission.md) | Security Group Outbound Permissions | [optional] 
**FileContent** | Pointer to **NullableString** | The file data as an array of bytes | [optional] 
**FileName** | Pointer to **NullableString** | Suggested file name | [optional] 
**ContentType** | Pointer to **NullableString** | The type of content in the file | [optional] 
**VCpus** | Pointer to **NullableInt32** |  | [optional] 
**MemoryInMB** | Pointer to **NullableInt64** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**KeyName** | Pointer to **NullableString** | Name of Key Pari | [optional] 

## Methods

### NewAccountResourcesItemsInner

`func NewAccountResourcesItemsInner(resourceType AwsAccountResourceType, accountType AccountType, ) *AccountResourcesItemsInner`

NewAccountResourcesItemsInner instantiates a new AccountResourcesItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResourcesItemsInnerWithDefaults

`func NewAccountResourcesItemsInnerWithDefaults() *AccountResourcesItemsInner`

NewAccountResourcesItemsInnerWithDefaults instantiates a new AccountResourcesItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedTenancySupportEnabled

`func (o *AccountResourcesItemsInner) GetDedicatedTenancySupportEnabled() bool`

GetDedicatedTenancySupportEnabled returns the DedicatedTenancySupportEnabled field if non-nil, zero value otherwise.

### GetDedicatedTenancySupportEnabledOk

`func (o *AccountResourcesItemsInner) GetDedicatedTenancySupportEnabledOk() (*bool, bool)`

GetDedicatedTenancySupportEnabledOk returns a tuple with the DedicatedTenancySupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTenancySupportEnabled

`func (o *AccountResourcesItemsInner) SetDedicatedTenancySupportEnabled(v bool)`

SetDedicatedTenancySupportEnabled sets DedicatedTenancySupportEnabled field to given value.

### HasDedicatedTenancySupportEnabled

`func (o *AccountResourcesItemsInner) HasDedicatedTenancySupportEnabled() bool`

HasDedicatedTenancySupportEnabled returns a boolean if a field has been set.

### GetDedicatedTenancyManagementCidrRange

`func (o *AccountResourcesItemsInner) GetDedicatedTenancyManagementCidrRange() string`

GetDedicatedTenancyManagementCidrRange returns the DedicatedTenancyManagementCidrRange field if non-nil, zero value otherwise.

### GetDedicatedTenancyManagementCidrRangeOk

`func (o *AccountResourcesItemsInner) GetDedicatedTenancyManagementCidrRangeOk() (*string, bool)`

GetDedicatedTenancyManagementCidrRangeOk returns a tuple with the DedicatedTenancyManagementCidrRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTenancyManagementCidrRange

`func (o *AccountResourcesItemsInner) SetDedicatedTenancyManagementCidrRange(v string)`

SetDedicatedTenancyManagementCidrRange sets DedicatedTenancyManagementCidrRange field to given value.

### HasDedicatedTenancyManagementCidrRange

`func (o *AccountResourcesItemsInner) HasDedicatedTenancyManagementCidrRange() bool`

HasDedicatedTenancyManagementCidrRange returns a boolean if a field has been set.

### SetDedicatedTenancyManagementCidrRangeNil

`func (o *AccountResourcesItemsInner) SetDedicatedTenancyManagementCidrRangeNil(b bool)`

 SetDedicatedTenancyManagementCidrRangeNil sets the value for DedicatedTenancyManagementCidrRange to be an explicit nil

### UnsetDedicatedTenancyManagementCidrRange
`func (o *AccountResourcesItemsInner) UnsetDedicatedTenancyManagementCidrRange()`

UnsetDedicatedTenancyManagementCidrRange ensures that no value is present for DedicatedTenancyManagementCidrRange, not even an explicit nil
### GetResourceType

`func (o *AccountResourcesItemsInner) GetResourceType() AwsAccountResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AccountResourcesItemsInner) GetResourceTypeOk() (*AwsAccountResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AccountResourcesItemsInner) SetResourceType(v AwsAccountResourceType)`

SetResourceType sets ResourceType field to given value.


### GetAccountType

`func (o *AccountResourcesItemsInner) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountResourcesItemsInner) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountResourcesItemsInner) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetAccountId

`func (o *AccountResourcesItemsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountResourcesItemsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountResourcesItemsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountResourcesItemsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *AccountResourcesItemsInner) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *AccountResourcesItemsInner) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetResourceId

`func (o *AccountResourcesItemsInner) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AccountResourcesItemsInner) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AccountResourcesItemsInner) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AccountResourcesItemsInner) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AccountResourcesItemsInner) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AccountResourcesItemsInner) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetImageId

`func (o *AccountResourcesItemsInner) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AccountResourcesItemsInner) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AccountResourcesItemsInner) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AccountResourcesItemsInner) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *AccountResourcesItemsInner) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *AccountResourcesItemsInner) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetName

`func (o *AccountResourcesItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountResourcesItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountResourcesItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountResourcesItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AccountResourcesItemsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AccountResourcesItemsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AccountResourcesItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountResourcesItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountResourcesItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountResourcesItemsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccountResourcesItemsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccountResourcesItemsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *AccountResourcesItemsInner) GetStatus() AwsEdcSubnetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountResourcesItemsInner) GetStatusOk() (*AwsEdcSubnetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountResourcesItemsInner) SetStatus(v AwsEdcSubnetStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountResourcesItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AccountResourcesItemsInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AccountResourcesItemsInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLocation

`func (o *AccountResourcesItemsInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AccountResourcesItemsInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AccountResourcesItemsInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AccountResourcesItemsInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AccountResourcesItemsInner) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AccountResourcesItemsInner) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetImageType

`func (o *AccountResourcesItemsInner) GetImageType() AwsEdcAmiImageType`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *AccountResourcesItemsInner) GetImageTypeOk() (*AwsEdcAmiImageType, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *AccountResourcesItemsInner) SetImageType(v AwsEdcAmiImageType)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *AccountResourcesItemsInner) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### SetImageTypeNil

`func (o *AccountResourcesItemsInner) SetImageTypeNil(b bool)`

 SetImageTypeNil sets the value for ImageType to be an explicit nil

### UnsetImageType
`func (o *AccountResourcesItemsInner) UnsetImageType()`

UnsetImageType ensures that no value is present for ImageType, not even an explicit nil
### GetPlatformType

`func (o *AccountResourcesItemsInner) GetPlatformType() AwsEdcAmiImagePlatform`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *AccountResourcesItemsInner) GetPlatformTypeOk() (*AwsEdcAmiImagePlatform, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *AccountResourcesItemsInner) SetPlatformType(v AwsEdcAmiImagePlatform)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *AccountResourcesItemsInner) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformTypeNil

`func (o *AccountResourcesItemsInner) SetPlatformTypeNil(b bool)`

 SetPlatformTypeNil sets the value for PlatformType to be an explicit nil

### UnsetPlatformType
`func (o *AccountResourcesItemsInner) UnsetPlatformType()`

UnsetPlatformType ensures that no value is present for PlatformType, not even an explicit nil
### GetBootMode

`func (o *AccountResourcesItemsInner) GetBootMode() AwsEdcAmiImageBootMode`

GetBootMode returns the BootMode field if non-nil, zero value otherwise.

### GetBootModeOk

`func (o *AccountResourcesItemsInner) GetBootModeOk() (*AwsEdcAmiImageBootMode, bool)`

GetBootModeOk returns a tuple with the BootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootMode

`func (o *AccountResourcesItemsInner) SetBootMode(v AwsEdcAmiImageBootMode)`

SetBootMode sets BootMode field to given value.

### HasBootMode

`func (o *AccountResourcesItemsInner) HasBootMode() bool`

HasBootMode returns a boolean if a field has been set.

### SetBootModeNil

`func (o *AccountResourcesItemsInner) SetBootModeNil(b bool)`

 SetBootModeNil sets the value for BootMode to be an explicit nil

### UnsetBootMode
`func (o *AccountResourcesItemsInner) UnsetBootMode()`

UnsetBootMode ensures that no value is present for BootMode, not even an explicit nil
### GetArchitectureType

`func (o *AccountResourcesItemsInner) GetArchitectureType() AwsEdcAmiImageArchitecture`

GetArchitectureType returns the ArchitectureType field if non-nil, zero value otherwise.

### GetArchitectureTypeOk

`func (o *AccountResourcesItemsInner) GetArchitectureTypeOk() (*AwsEdcAmiImageArchitecture, bool)`

GetArchitectureTypeOk returns a tuple with the ArchitectureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectureType

`func (o *AccountResourcesItemsInner) SetArchitectureType(v AwsEdcAmiImageArchitecture)`

SetArchitectureType sets ArchitectureType field to given value.

### HasArchitectureType

`func (o *AccountResourcesItemsInner) HasArchitectureType() bool`

HasArchitectureType returns a boolean if a field has been set.

### SetArchitectureTypeNil

`func (o *AccountResourcesItemsInner) SetArchitectureTypeNil(b bool)`

 SetArchitectureTypeNil sets the value for ArchitectureType to be an explicit nil

### UnsetArchitectureType
`func (o *AccountResourcesItemsInner) UnsetArchitectureType()`

UnsetArchitectureType ensures that no value is present for ArchitectureType, not even an explicit nil
### GetVirtualizationType

`func (o *AccountResourcesItemsInner) GetVirtualizationType() AwsEdcAmiImageVirtualization`

GetVirtualizationType returns the VirtualizationType field if non-nil, zero value otherwise.

### GetVirtualizationTypeOk

`func (o *AccountResourcesItemsInner) GetVirtualizationTypeOk() (*AwsEdcAmiImageVirtualization, bool)`

GetVirtualizationTypeOk returns a tuple with the VirtualizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationType

`func (o *AccountResourcesItemsInner) SetVirtualizationType(v AwsEdcAmiImageVirtualization)`

SetVirtualizationType sets VirtualizationType field to given value.

### HasVirtualizationType

`func (o *AccountResourcesItemsInner) HasVirtualizationType() bool`

HasVirtualizationType returns a boolean if a field has been set.

### SetVirtualizationTypeNil

`func (o *AccountResourcesItemsInner) SetVirtualizationTypeNil(b bool)`

 SetVirtualizationTypeNil sets the value for VirtualizationType to be an explicit nil

### UnsetVirtualizationType
`func (o *AccountResourcesItemsInner) UnsetVirtualizationType()`

UnsetVirtualizationType ensures that no value is present for VirtualizationType, not even an explicit nil
### GetPublic

`func (o *AccountResourcesItemsInner) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *AccountResourcesItemsInner) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *AccountResourcesItemsInner) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *AccountResourcesItemsInner) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### SetPublicNil

`func (o *AccountResourcesItemsInner) SetPublicNil(b bool)`

 SetPublicNil sets the value for Public to be an explicit nil

### UnsetPublic
`func (o *AccountResourcesItemsInner) UnsetPublic()`

UnsetPublic ensures that no value is present for Public, not even an explicit nil
### GetTags

`func (o *AccountResourcesItemsInner) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccountResourcesItemsInner) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccountResourcesItemsInner) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccountResourcesItemsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AccountResourcesItemsInner) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AccountResourcesItemsInner) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDirectoryId

`func (o *AccountResourcesItemsInner) GetDirectoryId() string`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *AccountResourcesItemsInner) GetDirectoryIdOk() (*string, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *AccountResourcesItemsInner) SetDirectoryId(v string)`

SetDirectoryId sets DirectoryId field to given value.

### HasDirectoryId

`func (o *AccountResourcesItemsInner) HasDirectoryId() bool`

HasDirectoryId returns a boolean if a field has been set.

### SetDirectoryIdNil

`func (o *AccountResourcesItemsInner) SetDirectoryIdNil(b bool)`

 SetDirectoryIdNil sets the value for DirectoryId to be an explicit nil

### UnsetDirectoryId
`func (o *AccountResourcesItemsInner) UnsetDirectoryId()`

UnsetDirectoryId ensures that no value is present for DirectoryId, not even an explicit nil
### GetDomain

`func (o *AccountResourcesItemsInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AccountResourcesItemsInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AccountResourcesItemsInner) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AccountResourcesItemsInner) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *AccountResourcesItemsInner) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *AccountResourcesItemsInner) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetDefaultOU

`func (o *AccountResourcesItemsInner) GetDefaultOU() string`

GetDefaultOU returns the DefaultOU field if non-nil, zero value otherwise.

### GetDefaultOUOk

`func (o *AccountResourcesItemsInner) GetDefaultOUOk() (*string, bool)`

GetDefaultOUOk returns a tuple with the DefaultOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOU

`func (o *AccountResourcesItemsInner) SetDefaultOU(v string)`

SetDefaultOU sets DefaultOU field to given value.

### HasDefaultOU

`func (o *AccountResourcesItemsInner) HasDefaultOU() bool`

HasDefaultOU returns a boolean if a field has been set.

### SetDefaultOUNil

`func (o *AccountResourcesItemsInner) SetDefaultOUNil(b bool)`

 SetDefaultOUNil sets the value for DefaultOU to be an explicit nil

### UnsetDefaultOU
`func (o *AccountResourcesItemsInner) UnsetDefaultOU()`

UnsetDefaultOU ensures that no value is present for DefaultOU, not even an explicit nil
### GetType

`func (o *AccountResourcesItemsInner) GetType() AwsEdcDirectoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountResourcesItemsInner) GetTypeOk() (*AwsEdcDirectoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountResourcesItemsInner) SetType(v AwsEdcDirectoryType)`

SetType sets Type field to given value.

### HasType

`func (o *AccountResourcesItemsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AccountResourcesItemsInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AccountResourcesItemsInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSize

`func (o *AccountResourcesItemsInner) GetSize() AwsEdcDirectorySize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AccountResourcesItemsInner) GetSizeOk() (*AwsEdcDirectorySize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AccountResourcesItemsInner) SetSize(v AwsEdcDirectorySize)`

SetSize sets Size field to given value.

### HasSize

`func (o *AccountResourcesItemsInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *AccountResourcesItemsInner) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *AccountResourcesItemsInner) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetRegistrationStatus

`func (o *AccountResourcesItemsInner) GetRegistrationStatus() AwsEdcDirectoryRegistrationStatus`

GetRegistrationStatus returns the RegistrationStatus field if non-nil, zero value otherwise.

### GetRegistrationStatusOk

`func (o *AccountResourcesItemsInner) GetRegistrationStatusOk() (*AwsEdcDirectoryRegistrationStatus, bool)`

GetRegistrationStatusOk returns a tuple with the RegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationStatus

`func (o *AccountResourcesItemsInner) SetRegistrationStatus(v AwsEdcDirectoryRegistrationStatus)`

SetRegistrationStatus sets RegistrationStatus field to given value.

### HasRegistrationStatus

`func (o *AccountResourcesItemsInner) HasRegistrationStatus() bool`

HasRegistrationStatus returns a boolean if a field has been set.

### SetRegistrationStatusNil

`func (o *AccountResourcesItemsInner) SetRegistrationStatusNil(b bool)`

 SetRegistrationStatusNil sets the value for RegistrationStatus to be an explicit nil

### UnsetRegistrationStatus
`func (o *AccountResourcesItemsInner) UnsetRegistrationStatus()`

UnsetRegistrationStatus ensures that no value is present for RegistrationStatus, not even an explicit nil
### GetTenancy

`func (o *AccountResourcesItemsInner) GetTenancy() AwsEdcDirectoryTenancy`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *AccountResourcesItemsInner) GetTenancyOk() (*AwsEdcDirectoryTenancy, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *AccountResourcesItemsInner) SetTenancy(v AwsEdcDirectoryTenancy)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *AccountResourcesItemsInner) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### SetTenancyNil

`func (o *AccountResourcesItemsInner) SetTenancyNil(b bool)`

 SetTenancyNil sets the value for Tenancy to be an explicit nil

### UnsetTenancy
`func (o *AccountResourcesItemsInner) UnsetTenancy()`

UnsetTenancy ensures that no value is present for Tenancy, not even an explicit nil
### GetVpcId

`func (o *AccountResourcesItemsInner) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AccountResourcesItemsInner) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AccountResourcesItemsInner) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AccountResourcesItemsInner) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *AccountResourcesItemsInner) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AccountResourcesItemsInner) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetSecurityGroupId

`func (o *AccountResourcesItemsInner) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *AccountResourcesItemsInner) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *AccountResourcesItemsInner) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *AccountResourcesItemsInner) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupIdNil

`func (o *AccountResourcesItemsInner) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *AccountResourcesItemsInner) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil
### GetSubnetIds

`func (o *AccountResourcesItemsInner) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *AccountResourcesItemsInner) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *AccountResourcesItemsInner) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.

### HasSubnetIds

`func (o *AccountResourcesItemsInner) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.

### SetSubnetIdsNil

`func (o *AccountResourcesItemsInner) SetSubnetIdsNil(b bool)`

 SetSubnetIdsNil sets the value for SubnetIds to be an explicit nil

### UnsetSubnetIds
`func (o *AccountResourcesItemsInner) UnsetSubnetIds()`

UnsetSubnetIds ensures that no value is present for SubnetIds, not even an explicit nil
### GetVpcTenancy

`func (o *AccountResourcesItemsInner) GetVpcTenancy() AwsEdcVpcTenancy`

GetVpcTenancy returns the VpcTenancy field if non-nil, zero value otherwise.

### GetVpcTenancyOk

`func (o *AccountResourcesItemsInner) GetVpcTenancyOk() (*AwsEdcVpcTenancy, bool)`

GetVpcTenancyOk returns a tuple with the VpcTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcTenancy

`func (o *AccountResourcesItemsInner) SetVpcTenancy(v AwsEdcVpcTenancy)`

SetVpcTenancy sets VpcTenancy field to given value.

### HasVpcTenancy

`func (o *AccountResourcesItemsInner) HasVpcTenancy() bool`

HasVpcTenancy returns a boolean if a field has been set.

### SetVpcTenancyNil

`func (o *AccountResourcesItemsInner) SetVpcTenancyNil(b bool)`

 SetVpcTenancyNil sets the value for VpcTenancy to be an explicit nil

### UnsetVpcTenancy
`func (o *AccountResourcesItemsInner) UnsetVpcTenancy()`

UnsetVpcTenancy ensures that no value is present for VpcTenancy, not even an explicit nil
### GetCidrBlock

`func (o *AccountResourcesItemsInner) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AccountResourcesItemsInner) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AccountResourcesItemsInner) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *AccountResourcesItemsInner) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### SetCidrBlockNil

`func (o *AccountResourcesItemsInner) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *AccountResourcesItemsInner) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
### GetSubnetId

`func (o *AccountResourcesItemsInner) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AccountResourcesItemsInner) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AccountResourcesItemsInner) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *AccountResourcesItemsInner) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *AccountResourcesItemsInner) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *AccountResourcesItemsInner) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetSubnetArn

`func (o *AccountResourcesItemsInner) GetSubnetArn() string`

GetSubnetArn returns the SubnetArn field if non-nil, zero value otherwise.

### GetSubnetArnOk

`func (o *AccountResourcesItemsInner) GetSubnetArnOk() (*string, bool)`

GetSubnetArnOk returns a tuple with the SubnetArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetArn

`func (o *AccountResourcesItemsInner) SetSubnetArn(v string)`

SetSubnetArn sets SubnetArn field to given value.

### HasSubnetArn

`func (o *AccountResourcesItemsInner) HasSubnetArn() bool`

HasSubnetArn returns a boolean if a field has been set.

### SetSubnetArnNil

`func (o *AccountResourcesItemsInner) SetSubnetArnNil(b bool)`

 SetSubnetArnNil sets the value for SubnetArn to be an explicit nil

### UnsetSubnetArn
`func (o *AccountResourcesItemsInner) UnsetSubnetArn()`

UnsetSubnetArn ensures that no value is present for SubnetArn, not even an explicit nil
### GetAwsRegion

`func (o *AccountResourcesItemsInner) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AccountResourcesItemsInner) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AccountResourcesItemsInner) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *AccountResourcesItemsInner) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *AccountResourcesItemsInner) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AccountResourcesItemsInner) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetAvailabilityZoneId

`func (o *AccountResourcesItemsInner) GetAvailabilityZoneId() string`

GetAvailabilityZoneId returns the AvailabilityZoneId field if non-nil, zero value otherwise.

### GetAvailabilityZoneIdOk

`func (o *AccountResourcesItemsInner) GetAvailabilityZoneIdOk() (*string, bool)`

GetAvailabilityZoneIdOk returns a tuple with the AvailabilityZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneId

`func (o *AccountResourcesItemsInner) SetAvailabilityZoneId(v string)`

SetAvailabilityZoneId sets AvailabilityZoneId field to given value.

### HasAvailabilityZoneId

`func (o *AccountResourcesItemsInner) HasAvailabilityZoneId() bool`

HasAvailabilityZoneId returns a boolean if a field has been set.

### SetAvailabilityZoneIdNil

`func (o *AccountResourcesItemsInner) SetAvailabilityZoneIdNil(b bool)`

 SetAvailabilityZoneIdNil sets the value for AvailabilityZoneId to be an explicit nil

### UnsetAvailabilityZoneId
`func (o *AccountResourcesItemsInner) UnsetAvailabilityZoneId()`

UnsetAvailabilityZoneId ensures that no value is present for AvailabilityZoneId, not even an explicit nil
### GetAvailabilityZone

`func (o *AccountResourcesItemsInner) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AccountResourcesItemsInner) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AccountResourcesItemsInner) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AccountResourcesItemsInner) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *AccountResourcesItemsInner) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *AccountResourcesItemsInner) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetIsSupportedByAwsManagedCapacity

`func (o *AccountResourcesItemsInner) GetIsSupportedByAwsManagedCapacity() bool`

GetIsSupportedByAwsManagedCapacity returns the IsSupportedByAwsManagedCapacity field if non-nil, zero value otherwise.

### GetIsSupportedByAwsManagedCapacityOk

`func (o *AccountResourcesItemsInner) GetIsSupportedByAwsManagedCapacityOk() (*bool, bool)`

GetIsSupportedByAwsManagedCapacityOk returns a tuple with the IsSupportedByAwsManagedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupportedByAwsManagedCapacity

`func (o *AccountResourcesItemsInner) SetIsSupportedByAwsManagedCapacity(v bool)`

SetIsSupportedByAwsManagedCapacity sets IsSupportedByAwsManagedCapacity field to given value.

### HasIsSupportedByAwsManagedCapacity

`func (o *AccountResourcesItemsInner) HasIsSupportedByAwsManagedCapacity() bool`

HasIsSupportedByAwsManagedCapacity returns a boolean if a field has been set.

### GetRegionNames

`func (o *AccountResourcesItemsInner) GetRegionNames() []string`

GetRegionNames returns the RegionNames field if non-nil, zero value otherwise.

### GetRegionNamesOk

`func (o *AccountResourcesItemsInner) GetRegionNamesOk() (*[]string, bool)`

GetRegionNamesOk returns a tuple with the RegionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNames

`func (o *AccountResourcesItemsInner) SetRegionNames(v []string)`

SetRegionNames sets RegionNames field to given value.

### HasRegionNames

`func (o *AccountResourcesItemsInner) HasRegionNames() bool`

HasRegionNames returns a boolean if a field has been set.

### SetRegionNamesNil

`func (o *AccountResourcesItemsInner) SetRegionNamesNil(b bool)`

 SetRegionNamesNil sets the value for RegionNames to be an explicit nil

### UnsetRegionNames
`func (o *AccountResourcesItemsInner) UnsetRegionNames()`

UnsetRegionNames ensures that no value is present for RegionNames, not even an explicit nil
### GetRegions

`func (o *AccountResourcesItemsInner) GetRegions() []AwsWorkspaceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AccountResourcesItemsInner) GetRegionsOk() (*[]AwsWorkspaceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AccountResourcesItemsInner) SetRegions(v []AwsWorkspaceRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AccountResourcesItemsInner) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *AccountResourcesItemsInner) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *AccountResourcesItemsInner) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetInboundPermissions

`func (o *AccountResourcesItemsInner) GetInboundPermissions() []IpPermission`

GetInboundPermissions returns the InboundPermissions field if non-nil, zero value otherwise.

### GetInboundPermissionsOk

`func (o *AccountResourcesItemsInner) GetInboundPermissionsOk() (*[]IpPermission, bool)`

GetInboundPermissionsOk returns a tuple with the InboundPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundPermissions

`func (o *AccountResourcesItemsInner) SetInboundPermissions(v []IpPermission)`

SetInboundPermissions sets InboundPermissions field to given value.

### HasInboundPermissions

`func (o *AccountResourcesItemsInner) HasInboundPermissions() bool`

HasInboundPermissions returns a boolean if a field has been set.

### SetInboundPermissionsNil

`func (o *AccountResourcesItemsInner) SetInboundPermissionsNil(b bool)`

 SetInboundPermissionsNil sets the value for InboundPermissions to be an explicit nil

### UnsetInboundPermissions
`func (o *AccountResourcesItemsInner) UnsetInboundPermissions()`

UnsetInboundPermissions ensures that no value is present for InboundPermissions, not even an explicit nil
### GetOutboundPermissions

`func (o *AccountResourcesItemsInner) GetOutboundPermissions() []IpPermission`

GetOutboundPermissions returns the OutboundPermissions field if non-nil, zero value otherwise.

### GetOutboundPermissionsOk

`func (o *AccountResourcesItemsInner) GetOutboundPermissionsOk() (*[]IpPermission, bool)`

GetOutboundPermissionsOk returns a tuple with the OutboundPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundPermissions

`func (o *AccountResourcesItemsInner) SetOutboundPermissions(v []IpPermission)`

SetOutboundPermissions sets OutboundPermissions field to given value.

### HasOutboundPermissions

`func (o *AccountResourcesItemsInner) HasOutboundPermissions() bool`

HasOutboundPermissions returns a boolean if a field has been set.

### SetOutboundPermissionsNil

`func (o *AccountResourcesItemsInner) SetOutboundPermissionsNil(b bool)`

 SetOutboundPermissionsNil sets the value for OutboundPermissions to be an explicit nil

### UnsetOutboundPermissions
`func (o *AccountResourcesItemsInner) UnsetOutboundPermissions()`

UnsetOutboundPermissions ensures that no value is present for OutboundPermissions, not even an explicit nil
### GetFileContent

`func (o *AccountResourcesItemsInner) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *AccountResourcesItemsInner) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *AccountResourcesItemsInner) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.

### HasFileContent

`func (o *AccountResourcesItemsInner) HasFileContent() bool`

HasFileContent returns a boolean if a field has been set.

### SetFileContentNil

`func (o *AccountResourcesItemsInner) SetFileContentNil(b bool)`

 SetFileContentNil sets the value for FileContent to be an explicit nil

### UnsetFileContent
`func (o *AccountResourcesItemsInner) UnsetFileContent()`

UnsetFileContent ensures that no value is present for FileContent, not even an explicit nil
### GetFileName

`func (o *AccountResourcesItemsInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AccountResourcesItemsInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AccountResourcesItemsInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AccountResourcesItemsInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *AccountResourcesItemsInner) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *AccountResourcesItemsInner) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetContentType

`func (o *AccountResourcesItemsInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AccountResourcesItemsInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AccountResourcesItemsInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *AccountResourcesItemsInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *AccountResourcesItemsInner) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *AccountResourcesItemsInner) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetVCpus

`func (o *AccountResourcesItemsInner) GetVCpus() int32`

GetVCpus returns the VCpus field if non-nil, zero value otherwise.

### GetVCpusOk

`func (o *AccountResourcesItemsInner) GetVCpusOk() (*int32, bool)`

GetVCpusOk returns a tuple with the VCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCpus

`func (o *AccountResourcesItemsInner) SetVCpus(v int32)`

SetVCpus sets VCpus field to given value.

### HasVCpus

`func (o *AccountResourcesItemsInner) HasVCpus() bool`

HasVCpus returns a boolean if a field has been set.

### SetVCpusNil

`func (o *AccountResourcesItemsInner) SetVCpusNil(b bool)`

 SetVCpusNil sets the value for VCpus to be an explicit nil

### UnsetVCpus
`func (o *AccountResourcesItemsInner) UnsetVCpus()`

UnsetVCpus ensures that no value is present for VCpus, not even an explicit nil
### GetMemoryInMB

`func (o *AccountResourcesItemsInner) GetMemoryInMB() int64`

GetMemoryInMB returns the MemoryInMB field if non-nil, zero value otherwise.

### GetMemoryInMBOk

`func (o *AccountResourcesItemsInner) GetMemoryInMBOk() (*int64, bool)`

GetMemoryInMBOk returns a tuple with the MemoryInMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMB

`func (o *AccountResourcesItemsInner) SetMemoryInMB(v int64)`

SetMemoryInMB sets MemoryInMB field to given value.

### HasMemoryInMB

`func (o *AccountResourcesItemsInner) HasMemoryInMB() bool`

HasMemoryInMB returns a boolean if a field has been set.

### SetMemoryInMBNil

`func (o *AccountResourcesItemsInner) SetMemoryInMBNil(b bool)`

 SetMemoryInMBNil sets the value for MemoryInMB to be an explicit nil

### UnsetMemoryInMB
`func (o *AccountResourcesItemsInner) UnsetMemoryInMB()`

UnsetMemoryInMB ensures that no value is present for MemoryInMB, not even an explicit nil
### GetIsDefault

`func (o *AccountResourcesItemsInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AccountResourcesItemsInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AccountResourcesItemsInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AccountResourcesItemsInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetKeyName

`func (o *AccountResourcesItemsInner) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *AccountResourcesItemsInner) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *AccountResourcesItemsInner) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *AccountResourcesItemsInner) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *AccountResourcesItemsInner) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *AccountResourcesItemsInner) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


