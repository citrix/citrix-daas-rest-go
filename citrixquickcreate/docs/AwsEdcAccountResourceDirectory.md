# AwsEdcAccountResourceDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryId** | Pointer to **NullableString** | Directory ID | [optional] 
**Name** | Pointer to **NullableString** | Directory Name | [optional] 
**Description** | Pointer to **NullableString** | Directory Description | [optional] 
**Domain** | Pointer to **NullableString** | Directory Domain | [optional] 
**DefaultOU** | Pointer to **NullableString** | Directory Organizational Unit | [optional] 
**Type** | Pointer to [**NullableAwsEdcDirectoryType**](AwsEdcDirectoryType.md) | Directory Type  Enum values ADCONNECTOR, MICROSOFTAD, SIMPLEAD, SHAREDMICROSOFTAD | [optional] 
**Size** | Pointer to [**NullableAwsEdcDirectorySize**](AwsEdcDirectorySize.md) | Directory Size  Enum values SMALL, LARGE | [optional] 
**Status** | Pointer to [**NullableAwsEdcDirectoryStatus**](AwsEdcDirectoryStatus.md) | Directory State  Enum values ACTIVE, CREATED, CREATING, DELETED, DELETING, FAILED, IMPARED, INOPERABLE, REQUESTED, RESTOREFAILED, RESTORING | [optional] 
**RegistrationStatus** | Pointer to [**NullableAwsEdcDirectoryRegistrationStatus**](AwsEdcDirectoryRegistrationStatus.md) | Directory Registration State  Enum values REGISTERED, REGISTERING, DEREGISTERED, DEREGISTERING, ERROR | [optional] 
**Tenancy** | Pointer to [**NullableAwsEdcDirectoryTenancy**](AwsEdcDirectoryTenancy.md) | Directory Registration Tenancy  Enum values SHARED, DEDICATED | [optional] 
**VpcId** | Pointer to **NullableString** | Directory associated VPC | [optional] 
**SecurityGroupId** | Pointer to **NullableString** | Directory associated Security Group | [optional] 
**SubnetIds** | Pointer to **[]string** | Directory associated Subnets | [optional] 

## Methods

### NewAwsEdcAccountResourceDirectory

`func NewAwsEdcAccountResourceDirectory() *AwsEdcAccountResourceDirectory`

NewAwsEdcAccountResourceDirectory instantiates a new AwsEdcAccountResourceDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcAccountResourceDirectoryWithDefaults

`func NewAwsEdcAccountResourceDirectoryWithDefaults() *AwsEdcAccountResourceDirectory`

NewAwsEdcAccountResourceDirectoryWithDefaults instantiates a new AwsEdcAccountResourceDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryId

`func (o *AwsEdcAccountResourceDirectory) GetDirectoryId() string`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *AwsEdcAccountResourceDirectory) GetDirectoryIdOk() (*string, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *AwsEdcAccountResourceDirectory) SetDirectoryId(v string)`

SetDirectoryId sets DirectoryId field to given value.

### HasDirectoryId

`func (o *AwsEdcAccountResourceDirectory) HasDirectoryId() bool`

HasDirectoryId returns a boolean if a field has been set.

### SetDirectoryIdNil

`func (o *AwsEdcAccountResourceDirectory) SetDirectoryIdNil(b bool)`

 SetDirectoryIdNil sets the value for DirectoryId to be an explicit nil

### UnsetDirectoryId
`func (o *AwsEdcAccountResourceDirectory) UnsetDirectoryId()`

UnsetDirectoryId ensures that no value is present for DirectoryId, not even an explicit nil
### GetName

`func (o *AwsEdcAccountResourceDirectory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsEdcAccountResourceDirectory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsEdcAccountResourceDirectory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsEdcAccountResourceDirectory) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsEdcAccountResourceDirectory) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsEdcAccountResourceDirectory) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AwsEdcAccountResourceDirectory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsEdcAccountResourceDirectory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsEdcAccountResourceDirectory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsEdcAccountResourceDirectory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AwsEdcAccountResourceDirectory) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AwsEdcAccountResourceDirectory) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomain

`func (o *AwsEdcAccountResourceDirectory) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AwsEdcAccountResourceDirectory) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AwsEdcAccountResourceDirectory) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AwsEdcAccountResourceDirectory) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *AwsEdcAccountResourceDirectory) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *AwsEdcAccountResourceDirectory) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetDefaultOU

`func (o *AwsEdcAccountResourceDirectory) GetDefaultOU() string`

GetDefaultOU returns the DefaultOU field if non-nil, zero value otherwise.

### GetDefaultOUOk

`func (o *AwsEdcAccountResourceDirectory) GetDefaultOUOk() (*string, bool)`

GetDefaultOUOk returns a tuple with the DefaultOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOU

`func (o *AwsEdcAccountResourceDirectory) SetDefaultOU(v string)`

SetDefaultOU sets DefaultOU field to given value.

### HasDefaultOU

`func (o *AwsEdcAccountResourceDirectory) HasDefaultOU() bool`

HasDefaultOU returns a boolean if a field has been set.

### SetDefaultOUNil

`func (o *AwsEdcAccountResourceDirectory) SetDefaultOUNil(b bool)`

 SetDefaultOUNil sets the value for DefaultOU to be an explicit nil

### UnsetDefaultOU
`func (o *AwsEdcAccountResourceDirectory) UnsetDefaultOU()`

UnsetDefaultOU ensures that no value is present for DefaultOU, not even an explicit nil
### GetType

`func (o *AwsEdcAccountResourceDirectory) GetType() AwsEdcDirectoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsEdcAccountResourceDirectory) GetTypeOk() (*AwsEdcDirectoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsEdcAccountResourceDirectory) SetType(v AwsEdcDirectoryType)`

SetType sets Type field to given value.

### HasType

`func (o *AwsEdcAccountResourceDirectory) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AwsEdcAccountResourceDirectory) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AwsEdcAccountResourceDirectory) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSize

`func (o *AwsEdcAccountResourceDirectory) GetSize() AwsEdcDirectorySize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AwsEdcAccountResourceDirectory) GetSizeOk() (*AwsEdcDirectorySize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AwsEdcAccountResourceDirectory) SetSize(v AwsEdcDirectorySize)`

SetSize sets Size field to given value.

### HasSize

`func (o *AwsEdcAccountResourceDirectory) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *AwsEdcAccountResourceDirectory) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *AwsEdcAccountResourceDirectory) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *AwsEdcAccountResourceDirectory) GetStatus() AwsEdcDirectoryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsEdcAccountResourceDirectory) GetStatusOk() (*AwsEdcDirectoryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsEdcAccountResourceDirectory) SetStatus(v AwsEdcDirectoryStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsEdcAccountResourceDirectory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AwsEdcAccountResourceDirectory) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AwsEdcAccountResourceDirectory) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRegistrationStatus

`func (o *AwsEdcAccountResourceDirectory) GetRegistrationStatus() AwsEdcDirectoryRegistrationStatus`

GetRegistrationStatus returns the RegistrationStatus field if non-nil, zero value otherwise.

### GetRegistrationStatusOk

`func (o *AwsEdcAccountResourceDirectory) GetRegistrationStatusOk() (*AwsEdcDirectoryRegistrationStatus, bool)`

GetRegistrationStatusOk returns a tuple with the RegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationStatus

`func (o *AwsEdcAccountResourceDirectory) SetRegistrationStatus(v AwsEdcDirectoryRegistrationStatus)`

SetRegistrationStatus sets RegistrationStatus field to given value.

### HasRegistrationStatus

`func (o *AwsEdcAccountResourceDirectory) HasRegistrationStatus() bool`

HasRegistrationStatus returns a boolean if a field has been set.

### SetRegistrationStatusNil

`func (o *AwsEdcAccountResourceDirectory) SetRegistrationStatusNil(b bool)`

 SetRegistrationStatusNil sets the value for RegistrationStatus to be an explicit nil

### UnsetRegistrationStatus
`func (o *AwsEdcAccountResourceDirectory) UnsetRegistrationStatus()`

UnsetRegistrationStatus ensures that no value is present for RegistrationStatus, not even an explicit nil
### GetTenancy

`func (o *AwsEdcAccountResourceDirectory) GetTenancy() AwsEdcDirectoryTenancy`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *AwsEdcAccountResourceDirectory) GetTenancyOk() (*AwsEdcDirectoryTenancy, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *AwsEdcAccountResourceDirectory) SetTenancy(v AwsEdcDirectoryTenancy)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *AwsEdcAccountResourceDirectory) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### SetTenancyNil

`func (o *AwsEdcAccountResourceDirectory) SetTenancyNil(b bool)`

 SetTenancyNil sets the value for Tenancy to be an explicit nil

### UnsetTenancy
`func (o *AwsEdcAccountResourceDirectory) UnsetTenancy()`

UnsetTenancy ensures that no value is present for Tenancy, not even an explicit nil
### GetVpcId

`func (o *AwsEdcAccountResourceDirectory) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsEdcAccountResourceDirectory) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsEdcAccountResourceDirectory) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsEdcAccountResourceDirectory) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *AwsEdcAccountResourceDirectory) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AwsEdcAccountResourceDirectory) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetSecurityGroupId

`func (o *AwsEdcAccountResourceDirectory) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *AwsEdcAccountResourceDirectory) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *AwsEdcAccountResourceDirectory) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *AwsEdcAccountResourceDirectory) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupIdNil

`func (o *AwsEdcAccountResourceDirectory) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *AwsEdcAccountResourceDirectory) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil
### GetSubnetIds

`func (o *AwsEdcAccountResourceDirectory) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *AwsEdcAccountResourceDirectory) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *AwsEdcAccountResourceDirectory) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.

### HasSubnetIds

`func (o *AwsEdcAccountResourceDirectory) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.

### SetSubnetIdsNil

`func (o *AwsEdcAccountResourceDirectory) SetSubnetIdsNil(b bool)`

 SetSubnetIdsNil sets the value for SubnetIds to be an explicit nil

### UnsetSubnetIds
`func (o *AwsEdcAccountResourceDirectory) UnsetSubnetIds()`

UnsetSubnetIds ensures that no value is present for SubnetIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


