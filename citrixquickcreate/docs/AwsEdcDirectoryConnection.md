# AwsEdcDirectoryConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryId** | Pointer to **NullableString** | Directory Id | [optional] 
**RegistrationStatus** | Pointer to [**AwsEdcDirectoryRegistrationStatus**](AwsEdcDirectoryRegistrationStatus.md) | Registration status of the directory | [optional] 
**DirectoryName** | Pointer to **NullableString** | Directory Name | [optional] 
**VpcId** | Pointer to **NullableString** | Directory Vpc | [optional] 
**Subnet1Id** | Pointer to **NullableString** | First Subnet Id | [optional] 
**Subnet2Id** | Pointer to **NullableString** | Second Subnet Id | [optional] 
**Tenancy** | Pointer to [**NullableAwsEdcDirectoryTenancy**](AwsEdcDirectoryTenancy.md) | Tenancy of directory  Enum values SHARED, DEDICATED | [optional] 
**EnableWorkDocs** | Pointer to **NullableBool** | Enable Work Docs | [optional] 
**UserEnabledAsLocalAdministrator** | Pointer to **NullableBool** | Enable Local Administrator | [optional] 
**EnableInternetAccess** | Pointer to **NullableBool** | Enable Internet Access | [optional] 
**EnableMaintananceMode** | Pointer to **NullableBool** | Enable Maintanance Mode | [optional] 
**SecurityGroupId** | Pointer to **NullableString** | Custom Security Group Id | [optional] 
**DefaultOU** | Pointer to **NullableString** | Custom Organizational Unit | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Any error message | [optional] 

## Methods

### NewAwsEdcDirectoryConnection

`func NewAwsEdcDirectoryConnection() *AwsEdcDirectoryConnection`

NewAwsEdcDirectoryConnection instantiates a new AwsEdcDirectoryConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcDirectoryConnectionWithDefaults

`func NewAwsEdcDirectoryConnectionWithDefaults() *AwsEdcDirectoryConnection`

NewAwsEdcDirectoryConnectionWithDefaults instantiates a new AwsEdcDirectoryConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryId

`func (o *AwsEdcDirectoryConnection) GetDirectoryId() string`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *AwsEdcDirectoryConnection) GetDirectoryIdOk() (*string, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *AwsEdcDirectoryConnection) SetDirectoryId(v string)`

SetDirectoryId sets DirectoryId field to given value.

### HasDirectoryId

`func (o *AwsEdcDirectoryConnection) HasDirectoryId() bool`

HasDirectoryId returns a boolean if a field has been set.

### SetDirectoryIdNil

`func (o *AwsEdcDirectoryConnection) SetDirectoryIdNil(b bool)`

 SetDirectoryIdNil sets the value for DirectoryId to be an explicit nil

### UnsetDirectoryId
`func (o *AwsEdcDirectoryConnection) UnsetDirectoryId()`

UnsetDirectoryId ensures that no value is present for DirectoryId, not even an explicit nil
### GetRegistrationStatus

`func (o *AwsEdcDirectoryConnection) GetRegistrationStatus() AwsEdcDirectoryRegistrationStatus`

GetRegistrationStatus returns the RegistrationStatus field if non-nil, zero value otherwise.

### GetRegistrationStatusOk

`func (o *AwsEdcDirectoryConnection) GetRegistrationStatusOk() (*AwsEdcDirectoryRegistrationStatus, bool)`

GetRegistrationStatusOk returns a tuple with the RegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationStatus

`func (o *AwsEdcDirectoryConnection) SetRegistrationStatus(v AwsEdcDirectoryRegistrationStatus)`

SetRegistrationStatus sets RegistrationStatus field to given value.

### HasRegistrationStatus

`func (o *AwsEdcDirectoryConnection) HasRegistrationStatus() bool`

HasRegistrationStatus returns a boolean if a field has been set.

### GetDirectoryName

`func (o *AwsEdcDirectoryConnection) GetDirectoryName() string`

GetDirectoryName returns the DirectoryName field if non-nil, zero value otherwise.

### GetDirectoryNameOk

`func (o *AwsEdcDirectoryConnection) GetDirectoryNameOk() (*string, bool)`

GetDirectoryNameOk returns a tuple with the DirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryName

`func (o *AwsEdcDirectoryConnection) SetDirectoryName(v string)`

SetDirectoryName sets DirectoryName field to given value.

### HasDirectoryName

`func (o *AwsEdcDirectoryConnection) HasDirectoryName() bool`

HasDirectoryName returns a boolean if a field has been set.

### SetDirectoryNameNil

`func (o *AwsEdcDirectoryConnection) SetDirectoryNameNil(b bool)`

 SetDirectoryNameNil sets the value for DirectoryName to be an explicit nil

### UnsetDirectoryName
`func (o *AwsEdcDirectoryConnection) UnsetDirectoryName()`

UnsetDirectoryName ensures that no value is present for DirectoryName, not even an explicit nil
### GetVpcId

`func (o *AwsEdcDirectoryConnection) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsEdcDirectoryConnection) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsEdcDirectoryConnection) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsEdcDirectoryConnection) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *AwsEdcDirectoryConnection) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AwsEdcDirectoryConnection) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetSubnet1Id

`func (o *AwsEdcDirectoryConnection) GetSubnet1Id() string`

GetSubnet1Id returns the Subnet1Id field if non-nil, zero value otherwise.

### GetSubnet1IdOk

`func (o *AwsEdcDirectoryConnection) GetSubnet1IdOk() (*string, bool)`

GetSubnet1IdOk returns a tuple with the Subnet1Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet1Id

`func (o *AwsEdcDirectoryConnection) SetSubnet1Id(v string)`

SetSubnet1Id sets Subnet1Id field to given value.

### HasSubnet1Id

`func (o *AwsEdcDirectoryConnection) HasSubnet1Id() bool`

HasSubnet1Id returns a boolean if a field has been set.

### SetSubnet1IdNil

`func (o *AwsEdcDirectoryConnection) SetSubnet1IdNil(b bool)`

 SetSubnet1IdNil sets the value for Subnet1Id to be an explicit nil

### UnsetSubnet1Id
`func (o *AwsEdcDirectoryConnection) UnsetSubnet1Id()`

UnsetSubnet1Id ensures that no value is present for Subnet1Id, not even an explicit nil
### GetSubnet2Id

`func (o *AwsEdcDirectoryConnection) GetSubnet2Id() string`

GetSubnet2Id returns the Subnet2Id field if non-nil, zero value otherwise.

### GetSubnet2IdOk

`func (o *AwsEdcDirectoryConnection) GetSubnet2IdOk() (*string, bool)`

GetSubnet2IdOk returns a tuple with the Subnet2Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet2Id

`func (o *AwsEdcDirectoryConnection) SetSubnet2Id(v string)`

SetSubnet2Id sets Subnet2Id field to given value.

### HasSubnet2Id

`func (o *AwsEdcDirectoryConnection) HasSubnet2Id() bool`

HasSubnet2Id returns a boolean if a field has been set.

### SetSubnet2IdNil

`func (o *AwsEdcDirectoryConnection) SetSubnet2IdNil(b bool)`

 SetSubnet2IdNil sets the value for Subnet2Id to be an explicit nil

### UnsetSubnet2Id
`func (o *AwsEdcDirectoryConnection) UnsetSubnet2Id()`

UnsetSubnet2Id ensures that no value is present for Subnet2Id, not even an explicit nil
### GetTenancy

`func (o *AwsEdcDirectoryConnection) GetTenancy() AwsEdcDirectoryTenancy`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *AwsEdcDirectoryConnection) GetTenancyOk() (*AwsEdcDirectoryTenancy, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *AwsEdcDirectoryConnection) SetTenancy(v AwsEdcDirectoryTenancy)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *AwsEdcDirectoryConnection) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### SetTenancyNil

`func (o *AwsEdcDirectoryConnection) SetTenancyNil(b bool)`

 SetTenancyNil sets the value for Tenancy to be an explicit nil

### UnsetTenancy
`func (o *AwsEdcDirectoryConnection) UnsetTenancy()`

UnsetTenancy ensures that no value is present for Tenancy, not even an explicit nil
### GetEnableWorkDocs

`func (o *AwsEdcDirectoryConnection) GetEnableWorkDocs() bool`

GetEnableWorkDocs returns the EnableWorkDocs field if non-nil, zero value otherwise.

### GetEnableWorkDocsOk

`func (o *AwsEdcDirectoryConnection) GetEnableWorkDocsOk() (*bool, bool)`

GetEnableWorkDocsOk returns a tuple with the EnableWorkDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWorkDocs

`func (o *AwsEdcDirectoryConnection) SetEnableWorkDocs(v bool)`

SetEnableWorkDocs sets EnableWorkDocs field to given value.

### HasEnableWorkDocs

`func (o *AwsEdcDirectoryConnection) HasEnableWorkDocs() bool`

HasEnableWorkDocs returns a boolean if a field has been set.

### SetEnableWorkDocsNil

`func (o *AwsEdcDirectoryConnection) SetEnableWorkDocsNil(b bool)`

 SetEnableWorkDocsNil sets the value for EnableWorkDocs to be an explicit nil

### UnsetEnableWorkDocs
`func (o *AwsEdcDirectoryConnection) UnsetEnableWorkDocs()`

UnsetEnableWorkDocs ensures that no value is present for EnableWorkDocs, not even an explicit nil
### GetUserEnabledAsLocalAdministrator

`func (o *AwsEdcDirectoryConnection) GetUserEnabledAsLocalAdministrator() bool`

GetUserEnabledAsLocalAdministrator returns the UserEnabledAsLocalAdministrator field if non-nil, zero value otherwise.

### GetUserEnabledAsLocalAdministratorOk

`func (o *AwsEdcDirectoryConnection) GetUserEnabledAsLocalAdministratorOk() (*bool, bool)`

GetUserEnabledAsLocalAdministratorOk returns a tuple with the UserEnabledAsLocalAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnabledAsLocalAdministrator

`func (o *AwsEdcDirectoryConnection) SetUserEnabledAsLocalAdministrator(v bool)`

SetUserEnabledAsLocalAdministrator sets UserEnabledAsLocalAdministrator field to given value.

### HasUserEnabledAsLocalAdministrator

`func (o *AwsEdcDirectoryConnection) HasUserEnabledAsLocalAdministrator() bool`

HasUserEnabledAsLocalAdministrator returns a boolean if a field has been set.

### SetUserEnabledAsLocalAdministratorNil

`func (o *AwsEdcDirectoryConnection) SetUserEnabledAsLocalAdministratorNil(b bool)`

 SetUserEnabledAsLocalAdministratorNil sets the value for UserEnabledAsLocalAdministrator to be an explicit nil

### UnsetUserEnabledAsLocalAdministrator
`func (o *AwsEdcDirectoryConnection) UnsetUserEnabledAsLocalAdministrator()`

UnsetUserEnabledAsLocalAdministrator ensures that no value is present for UserEnabledAsLocalAdministrator, not even an explicit nil
### GetEnableInternetAccess

`func (o *AwsEdcDirectoryConnection) GetEnableInternetAccess() bool`

GetEnableInternetAccess returns the EnableInternetAccess field if non-nil, zero value otherwise.

### GetEnableInternetAccessOk

`func (o *AwsEdcDirectoryConnection) GetEnableInternetAccessOk() (*bool, bool)`

GetEnableInternetAccessOk returns a tuple with the EnableInternetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternetAccess

`func (o *AwsEdcDirectoryConnection) SetEnableInternetAccess(v bool)`

SetEnableInternetAccess sets EnableInternetAccess field to given value.

### HasEnableInternetAccess

`func (o *AwsEdcDirectoryConnection) HasEnableInternetAccess() bool`

HasEnableInternetAccess returns a boolean if a field has been set.

### SetEnableInternetAccessNil

`func (o *AwsEdcDirectoryConnection) SetEnableInternetAccessNil(b bool)`

 SetEnableInternetAccessNil sets the value for EnableInternetAccess to be an explicit nil

### UnsetEnableInternetAccess
`func (o *AwsEdcDirectoryConnection) UnsetEnableInternetAccess()`

UnsetEnableInternetAccess ensures that no value is present for EnableInternetAccess, not even an explicit nil
### GetEnableMaintananceMode

`func (o *AwsEdcDirectoryConnection) GetEnableMaintananceMode() bool`

GetEnableMaintananceMode returns the EnableMaintananceMode field if non-nil, zero value otherwise.

### GetEnableMaintananceModeOk

`func (o *AwsEdcDirectoryConnection) GetEnableMaintananceModeOk() (*bool, bool)`

GetEnableMaintananceModeOk returns a tuple with the EnableMaintananceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMaintananceMode

`func (o *AwsEdcDirectoryConnection) SetEnableMaintananceMode(v bool)`

SetEnableMaintananceMode sets EnableMaintananceMode field to given value.

### HasEnableMaintananceMode

`func (o *AwsEdcDirectoryConnection) HasEnableMaintananceMode() bool`

HasEnableMaintananceMode returns a boolean if a field has been set.

### SetEnableMaintananceModeNil

`func (o *AwsEdcDirectoryConnection) SetEnableMaintananceModeNil(b bool)`

 SetEnableMaintananceModeNil sets the value for EnableMaintananceMode to be an explicit nil

### UnsetEnableMaintananceMode
`func (o *AwsEdcDirectoryConnection) UnsetEnableMaintananceMode()`

UnsetEnableMaintananceMode ensures that no value is present for EnableMaintananceMode, not even an explicit nil
### GetSecurityGroupId

`func (o *AwsEdcDirectoryConnection) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *AwsEdcDirectoryConnection) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *AwsEdcDirectoryConnection) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *AwsEdcDirectoryConnection) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupIdNil

`func (o *AwsEdcDirectoryConnection) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *AwsEdcDirectoryConnection) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil
### GetDefaultOU

`func (o *AwsEdcDirectoryConnection) GetDefaultOU() string`

GetDefaultOU returns the DefaultOU field if non-nil, zero value otherwise.

### GetDefaultOUOk

`func (o *AwsEdcDirectoryConnection) GetDefaultOUOk() (*string, bool)`

GetDefaultOUOk returns a tuple with the DefaultOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOU

`func (o *AwsEdcDirectoryConnection) SetDefaultOU(v string)`

SetDefaultOU sets DefaultOU field to given value.

### HasDefaultOU

`func (o *AwsEdcDirectoryConnection) HasDefaultOU() bool`

HasDefaultOU returns a boolean if a field has been set.

### SetDefaultOUNil

`func (o *AwsEdcDirectoryConnection) SetDefaultOUNil(b bool)`

 SetDefaultOUNil sets the value for DefaultOU to be an explicit nil

### UnsetDefaultOU
`func (o *AwsEdcDirectoryConnection) UnsetDefaultOU()`

UnsetDefaultOU ensures that no value is present for DefaultOU, not even an explicit nil
### GetErrorMessage

`func (o *AwsEdcDirectoryConnection) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AwsEdcDirectoryConnection) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AwsEdcDirectoryConnection) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AwsEdcDirectoryConnection) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AwsEdcDirectoryConnection) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AwsEdcDirectoryConnection) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


