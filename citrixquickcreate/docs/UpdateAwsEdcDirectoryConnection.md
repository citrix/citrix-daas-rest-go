# UpdateAwsEdcDirectoryConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultOU** | Pointer to **NullableString** | Default OU | [optional] 
**EnableWorkDocs** | Pointer to **NullableBool** | Enable Work Docs | [optional] 
**UserEnabledAsLocalAdministrator** | Pointer to **NullableBool** | Enable Local Administrator | [optional] 
**EnableInternetAccess** | Pointer to **NullableBool** | Enable Internet Access | [optional] 
**EnableMaintananceMode** | Pointer to **NullableBool** | Enable Maintanance Mode | [optional] 
**SecurityGroupId** | Pointer to **NullableString** | Custom Security Group Id | [optional] 

## Methods

### NewUpdateAwsEdcDirectoryConnection

`func NewUpdateAwsEdcDirectoryConnection() *UpdateAwsEdcDirectoryConnection`

NewUpdateAwsEdcDirectoryConnection instantiates a new UpdateAwsEdcDirectoryConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAwsEdcDirectoryConnectionWithDefaults

`func NewUpdateAwsEdcDirectoryConnectionWithDefaults() *UpdateAwsEdcDirectoryConnection`

NewUpdateAwsEdcDirectoryConnectionWithDefaults instantiates a new UpdateAwsEdcDirectoryConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultOU

`func (o *UpdateAwsEdcDirectoryConnection) GetDefaultOU() string`

GetDefaultOU returns the DefaultOU field if non-nil, zero value otherwise.

### GetDefaultOUOk

`func (o *UpdateAwsEdcDirectoryConnection) GetDefaultOUOk() (*string, bool)`

GetDefaultOUOk returns a tuple with the DefaultOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOU

`func (o *UpdateAwsEdcDirectoryConnection) SetDefaultOU(v string)`

SetDefaultOU sets DefaultOU field to given value.

### HasDefaultOU

`func (o *UpdateAwsEdcDirectoryConnection) HasDefaultOU() bool`

HasDefaultOU returns a boolean if a field has been set.

### SetDefaultOUNil

`func (o *UpdateAwsEdcDirectoryConnection) SetDefaultOUNil(b bool)`

 SetDefaultOUNil sets the value for DefaultOU to be an explicit nil

### UnsetDefaultOU
`func (o *UpdateAwsEdcDirectoryConnection) UnsetDefaultOU()`

UnsetDefaultOU ensures that no value is present for DefaultOU, not even an explicit nil
### GetEnableWorkDocs

`func (o *UpdateAwsEdcDirectoryConnection) GetEnableWorkDocs() bool`

GetEnableWorkDocs returns the EnableWorkDocs field if non-nil, zero value otherwise.

### GetEnableWorkDocsOk

`func (o *UpdateAwsEdcDirectoryConnection) GetEnableWorkDocsOk() (*bool, bool)`

GetEnableWorkDocsOk returns a tuple with the EnableWorkDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWorkDocs

`func (o *UpdateAwsEdcDirectoryConnection) SetEnableWorkDocs(v bool)`

SetEnableWorkDocs sets EnableWorkDocs field to given value.

### HasEnableWorkDocs

`func (o *UpdateAwsEdcDirectoryConnection) HasEnableWorkDocs() bool`

HasEnableWorkDocs returns a boolean if a field has been set.

### SetEnableWorkDocsNil

`func (o *UpdateAwsEdcDirectoryConnection) SetEnableWorkDocsNil(b bool)`

 SetEnableWorkDocsNil sets the value for EnableWorkDocs to be an explicit nil

### UnsetEnableWorkDocs
`func (o *UpdateAwsEdcDirectoryConnection) UnsetEnableWorkDocs()`

UnsetEnableWorkDocs ensures that no value is present for EnableWorkDocs, not even an explicit nil
### GetUserEnabledAsLocalAdministrator

`func (o *UpdateAwsEdcDirectoryConnection) GetUserEnabledAsLocalAdministrator() bool`

GetUserEnabledAsLocalAdministrator returns the UserEnabledAsLocalAdministrator field if non-nil, zero value otherwise.

### GetUserEnabledAsLocalAdministratorOk

`func (o *UpdateAwsEdcDirectoryConnection) GetUserEnabledAsLocalAdministratorOk() (*bool, bool)`

GetUserEnabledAsLocalAdministratorOk returns a tuple with the UserEnabledAsLocalAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnabledAsLocalAdministrator

`func (o *UpdateAwsEdcDirectoryConnection) SetUserEnabledAsLocalAdministrator(v bool)`

SetUserEnabledAsLocalAdministrator sets UserEnabledAsLocalAdministrator field to given value.

### HasUserEnabledAsLocalAdministrator

`func (o *UpdateAwsEdcDirectoryConnection) HasUserEnabledAsLocalAdministrator() bool`

HasUserEnabledAsLocalAdministrator returns a boolean if a field has been set.

### SetUserEnabledAsLocalAdministratorNil

`func (o *UpdateAwsEdcDirectoryConnection) SetUserEnabledAsLocalAdministratorNil(b bool)`

 SetUserEnabledAsLocalAdministratorNil sets the value for UserEnabledAsLocalAdministrator to be an explicit nil

### UnsetUserEnabledAsLocalAdministrator
`func (o *UpdateAwsEdcDirectoryConnection) UnsetUserEnabledAsLocalAdministrator()`

UnsetUserEnabledAsLocalAdministrator ensures that no value is present for UserEnabledAsLocalAdministrator, not even an explicit nil
### GetEnableInternetAccess

`func (o *UpdateAwsEdcDirectoryConnection) GetEnableInternetAccess() bool`

GetEnableInternetAccess returns the EnableInternetAccess field if non-nil, zero value otherwise.

### GetEnableInternetAccessOk

`func (o *UpdateAwsEdcDirectoryConnection) GetEnableInternetAccessOk() (*bool, bool)`

GetEnableInternetAccessOk returns a tuple with the EnableInternetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternetAccess

`func (o *UpdateAwsEdcDirectoryConnection) SetEnableInternetAccess(v bool)`

SetEnableInternetAccess sets EnableInternetAccess field to given value.

### HasEnableInternetAccess

`func (o *UpdateAwsEdcDirectoryConnection) HasEnableInternetAccess() bool`

HasEnableInternetAccess returns a boolean if a field has been set.

### SetEnableInternetAccessNil

`func (o *UpdateAwsEdcDirectoryConnection) SetEnableInternetAccessNil(b bool)`

 SetEnableInternetAccessNil sets the value for EnableInternetAccess to be an explicit nil

### UnsetEnableInternetAccess
`func (o *UpdateAwsEdcDirectoryConnection) UnsetEnableInternetAccess()`

UnsetEnableInternetAccess ensures that no value is present for EnableInternetAccess, not even an explicit nil
### GetEnableMaintananceMode

`func (o *UpdateAwsEdcDirectoryConnection) GetEnableMaintananceMode() bool`

GetEnableMaintananceMode returns the EnableMaintananceMode field if non-nil, zero value otherwise.

### GetEnableMaintananceModeOk

`func (o *UpdateAwsEdcDirectoryConnection) GetEnableMaintananceModeOk() (*bool, bool)`

GetEnableMaintananceModeOk returns a tuple with the EnableMaintananceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMaintananceMode

`func (o *UpdateAwsEdcDirectoryConnection) SetEnableMaintananceMode(v bool)`

SetEnableMaintananceMode sets EnableMaintananceMode field to given value.

### HasEnableMaintananceMode

`func (o *UpdateAwsEdcDirectoryConnection) HasEnableMaintananceMode() bool`

HasEnableMaintananceMode returns a boolean if a field has been set.

### SetEnableMaintananceModeNil

`func (o *UpdateAwsEdcDirectoryConnection) SetEnableMaintananceModeNil(b bool)`

 SetEnableMaintananceModeNil sets the value for EnableMaintananceMode to be an explicit nil

### UnsetEnableMaintananceMode
`func (o *UpdateAwsEdcDirectoryConnection) UnsetEnableMaintananceMode()`

UnsetEnableMaintananceMode ensures that no value is present for EnableMaintananceMode, not even an explicit nil
### GetSecurityGroupId

`func (o *UpdateAwsEdcDirectoryConnection) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *UpdateAwsEdcDirectoryConnection) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *UpdateAwsEdcDirectoryConnection) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *UpdateAwsEdcDirectoryConnection) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupIdNil

`func (o *UpdateAwsEdcDirectoryConnection) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *UpdateAwsEdcDirectoryConnection) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


