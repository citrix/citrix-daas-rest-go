# AddAwsEdcDirectoryConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryId** | Pointer to **NullableString** | Directory Id | [optional] 
**Subnet1Id** | Pointer to **NullableString** | First Subnet Id | [optional] 
**Subnet2Id** | Pointer to **NullableString** | Second Subnet Id | [optional] 
**Tenancy** | Pointer to [**NullableTenancy**](Tenancy.md) |  | [optional] 
**EnableWorkDocs** | Pointer to **bool** | Enable Work Docs | [optional] 
**UserEnabledAsLocalAdministrator** | Pointer to **bool** | Enable Local Administrator | [optional] 
**SecurityGroupId** | Pointer to **NullableString** | The identifier of the security group | [optional] 
**DefaultOu** | Pointer to **NullableString** | The default OU for workspace directories | [optional] 
**EnableMaintenanceMode** | Pointer to **bool** | Indicates if maintenance mode is enabled for workspaces | [optional] 

## Methods

### NewAddAwsEdcDirectoryConnection

`func NewAddAwsEdcDirectoryConnection() *AddAwsEdcDirectoryConnection`

NewAddAwsEdcDirectoryConnection instantiates a new AddAwsEdcDirectoryConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAwsEdcDirectoryConnectionWithDefaults

`func NewAddAwsEdcDirectoryConnectionWithDefaults() *AddAwsEdcDirectoryConnection`

NewAddAwsEdcDirectoryConnectionWithDefaults instantiates a new AddAwsEdcDirectoryConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryId

`func (o *AddAwsEdcDirectoryConnection) GetDirectoryId() string`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *AddAwsEdcDirectoryConnection) GetDirectoryIdOk() (*string, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *AddAwsEdcDirectoryConnection) SetDirectoryId(v string)`

SetDirectoryId sets DirectoryId field to given value.

### HasDirectoryId

`func (o *AddAwsEdcDirectoryConnection) HasDirectoryId() bool`

HasDirectoryId returns a boolean if a field has been set.

### SetDirectoryIdNil

`func (o *AddAwsEdcDirectoryConnection) SetDirectoryIdNil(b bool)`

 SetDirectoryIdNil sets the value for DirectoryId to be an explicit nil

### UnsetDirectoryId
`func (o *AddAwsEdcDirectoryConnection) UnsetDirectoryId()`

UnsetDirectoryId ensures that no value is present for DirectoryId, not even an explicit nil
### GetSubnet1Id

`func (o *AddAwsEdcDirectoryConnection) GetSubnet1Id() string`

GetSubnet1Id returns the Subnet1Id field if non-nil, zero value otherwise.

### GetSubnet1IdOk

`func (o *AddAwsEdcDirectoryConnection) GetSubnet1IdOk() (*string, bool)`

GetSubnet1IdOk returns a tuple with the Subnet1Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet1Id

`func (o *AddAwsEdcDirectoryConnection) SetSubnet1Id(v string)`

SetSubnet1Id sets Subnet1Id field to given value.

### HasSubnet1Id

`func (o *AddAwsEdcDirectoryConnection) HasSubnet1Id() bool`

HasSubnet1Id returns a boolean if a field has been set.

### SetSubnet1IdNil

`func (o *AddAwsEdcDirectoryConnection) SetSubnet1IdNil(b bool)`

 SetSubnet1IdNil sets the value for Subnet1Id to be an explicit nil

### UnsetSubnet1Id
`func (o *AddAwsEdcDirectoryConnection) UnsetSubnet1Id()`

UnsetSubnet1Id ensures that no value is present for Subnet1Id, not even an explicit nil
### GetSubnet2Id

`func (o *AddAwsEdcDirectoryConnection) GetSubnet2Id() string`

GetSubnet2Id returns the Subnet2Id field if non-nil, zero value otherwise.

### GetSubnet2IdOk

`func (o *AddAwsEdcDirectoryConnection) GetSubnet2IdOk() (*string, bool)`

GetSubnet2IdOk returns a tuple with the Subnet2Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet2Id

`func (o *AddAwsEdcDirectoryConnection) SetSubnet2Id(v string)`

SetSubnet2Id sets Subnet2Id field to given value.

### HasSubnet2Id

`func (o *AddAwsEdcDirectoryConnection) HasSubnet2Id() bool`

HasSubnet2Id returns a boolean if a field has been set.

### SetSubnet2IdNil

`func (o *AddAwsEdcDirectoryConnection) SetSubnet2IdNil(b bool)`

 SetSubnet2IdNil sets the value for Subnet2Id to be an explicit nil

### UnsetSubnet2Id
`func (o *AddAwsEdcDirectoryConnection) UnsetSubnet2Id()`

UnsetSubnet2Id ensures that no value is present for Subnet2Id, not even an explicit nil
### GetTenancy

`func (o *AddAwsEdcDirectoryConnection) GetTenancy() Tenancy`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *AddAwsEdcDirectoryConnection) GetTenancyOk() (*Tenancy, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *AddAwsEdcDirectoryConnection) SetTenancy(v Tenancy)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *AddAwsEdcDirectoryConnection) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### SetTenancyNil

`func (o *AddAwsEdcDirectoryConnection) SetTenancyNil(b bool)`

 SetTenancyNil sets the value for Tenancy to be an explicit nil

### UnsetTenancy
`func (o *AddAwsEdcDirectoryConnection) UnsetTenancy()`

UnsetTenancy ensures that no value is present for Tenancy, not even an explicit nil
### GetEnableWorkDocs

`func (o *AddAwsEdcDirectoryConnection) GetEnableWorkDocs() bool`

GetEnableWorkDocs returns the EnableWorkDocs field if non-nil, zero value otherwise.

### GetEnableWorkDocsOk

`func (o *AddAwsEdcDirectoryConnection) GetEnableWorkDocsOk() (*bool, bool)`

GetEnableWorkDocsOk returns a tuple with the EnableWorkDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWorkDocs

`func (o *AddAwsEdcDirectoryConnection) SetEnableWorkDocs(v bool)`

SetEnableWorkDocs sets EnableWorkDocs field to given value.

### HasEnableWorkDocs

`func (o *AddAwsEdcDirectoryConnection) HasEnableWorkDocs() bool`

HasEnableWorkDocs returns a boolean if a field has been set.

### GetUserEnabledAsLocalAdministrator

`func (o *AddAwsEdcDirectoryConnection) GetUserEnabledAsLocalAdministrator() bool`

GetUserEnabledAsLocalAdministrator returns the UserEnabledAsLocalAdministrator field if non-nil, zero value otherwise.

### GetUserEnabledAsLocalAdministratorOk

`func (o *AddAwsEdcDirectoryConnection) GetUserEnabledAsLocalAdministratorOk() (*bool, bool)`

GetUserEnabledAsLocalAdministratorOk returns a tuple with the UserEnabledAsLocalAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnabledAsLocalAdministrator

`func (o *AddAwsEdcDirectoryConnection) SetUserEnabledAsLocalAdministrator(v bool)`

SetUserEnabledAsLocalAdministrator sets UserEnabledAsLocalAdministrator field to given value.

### HasUserEnabledAsLocalAdministrator

`func (o *AddAwsEdcDirectoryConnection) HasUserEnabledAsLocalAdministrator() bool`

HasUserEnabledAsLocalAdministrator returns a boolean if a field has been set.

### GetSecurityGroupId

`func (o *AddAwsEdcDirectoryConnection) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *AddAwsEdcDirectoryConnection) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *AddAwsEdcDirectoryConnection) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *AddAwsEdcDirectoryConnection) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupIdNil

`func (o *AddAwsEdcDirectoryConnection) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *AddAwsEdcDirectoryConnection) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil
### GetDefaultOu

`func (o *AddAwsEdcDirectoryConnection) GetDefaultOu() string`

GetDefaultOu returns the DefaultOu field if non-nil, zero value otherwise.

### GetDefaultOuOk

`func (o *AddAwsEdcDirectoryConnection) GetDefaultOuOk() (*string, bool)`

GetDefaultOuOk returns a tuple with the DefaultOu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOu

`func (o *AddAwsEdcDirectoryConnection) SetDefaultOu(v string)`

SetDefaultOu sets DefaultOu field to given value.

### HasDefaultOu

`func (o *AddAwsEdcDirectoryConnection) HasDefaultOu() bool`

HasDefaultOu returns a boolean if a field has been set.

### SetDefaultOuNil

`func (o *AddAwsEdcDirectoryConnection) SetDefaultOuNil(b bool)`

 SetDefaultOuNil sets the value for DefaultOu to be an explicit nil

### UnsetDefaultOu
`func (o *AddAwsEdcDirectoryConnection) UnsetDefaultOu()`

UnsetDefaultOu ensures that no value is present for DefaultOu, not even an explicit nil
### GetEnableMaintenanceMode

`func (o *AddAwsEdcDirectoryConnection) GetEnableMaintenanceMode() bool`

GetEnableMaintenanceMode returns the EnableMaintenanceMode field if non-nil, zero value otherwise.

### GetEnableMaintenanceModeOk

`func (o *AddAwsEdcDirectoryConnection) GetEnableMaintenanceModeOk() (*bool, bool)`

GetEnableMaintenanceModeOk returns a tuple with the EnableMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMaintenanceMode

`func (o *AddAwsEdcDirectoryConnection) SetEnableMaintenanceMode(v bool)`

SetEnableMaintenanceMode sets EnableMaintenanceMode field to given value.

### HasEnableMaintenanceMode

`func (o *AddAwsEdcDirectoryConnection) HasEnableMaintenanceMode() bool`

HasEnableMaintenanceMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


