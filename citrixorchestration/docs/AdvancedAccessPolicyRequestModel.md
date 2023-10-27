# AdvancedAccessPolicyRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the broker access policy. | [optional] 
**Name** | Pointer to **NullableString** | Name of the advanced access policy. | [optional] 
**AllowHdxAccess** | Pointer to **NullableBool** | Whether to allow connections using HDX. Was: AllowsHdxAccess | [optional] [default to true]
**AllowMachineRestart** | Pointer to **NullableBool** | Whether to allow users to self-service restart machines. | [optional] [default to true]
**AllowRdpAccess** | Pointer to **NullableBool** | Whether to allow connections using RDP. Was: AllowsRdpAccess | [optional] [default to true]
**AllowedConnection** | Pointer to [**AllowedConnection**](AllowedConnection.md) |  | [optional] 
**AllowedUsers** | Pointer to [**AllowedUser**](AllowedUser.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Description. | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the advanced access policy is enabled.  If a policy is disabled it is not considered when evaluating whether a user may access the delivery group. | [optional] [default to true]
**ExcludedClientIPFilterEnabled** | Pointer to **NullableBool** | Specifies whether the ExcludedClientIPs filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**ExcludedClientIPs** | Pointer to **[]string** | Specifies the client IPs which will be denied access to the delivery group. | [optional] 
**ExcludedClientNameFilterEnabled** | Pointer to **NullableBool** | Specifies whether the ExcludedClientNames filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**ExcludedClientNames** | Pointer to **[]string** | Specifies the client names which will be denied access to the delivery group. Validation: each specified item must have a maximum length and regex match | [optional] 
**ExcludedSmartAccessFilterEnabled** | Pointer to **NullableBool** | Specifies whether the ExcludedSmartAccessTags filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**ExcludedSmartAccessTags** | Pointer to [**[]SmartAccessTagRequestModel**](SmartAccessTagRequestModel.md) | Specifies the SmartAccess tags which will be denied access to the delivery group, if any occur in those provided by NetScaler Gateway with the user&#39;s connection. | [optional] 
**ExcludedUserFilterEnabled** | Pointer to **NullableBool** | Specifies whether the ExcludedUsers filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**ExcludedUsers** | Pointer to **[]string** | Specifies the users and groups who are denied access to the delivery group. | [optional] 
**IncludedClientIPFilterEnabled** | Pointer to **NullableBool** | Specifies whether the IncludedClientIPs filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**IncludedClientIPs** | Pointer to **[]string** | Specifies the client IPs which will be allowed access to the delivery group. | [optional] 
**IncludedClientNameFilterEnabled** | Pointer to **NullableBool** | Specifies whether the IncludedClientNames filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**IncludedClientNames** | Pointer to **[]string** | Specifies the client names which will be granted access to the delivery group. Validation: each specified item must have a maximum length and regex match | [optional] 
**IncludedSmartAccessFilterEnabled** | Pointer to **NullableBool** | Specifies whether the IncludedSmartAccessTags filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**IncludedSmartAccessTags** | Pointer to [**[]SmartAccessTagRequestModel**](SmartAccessTagRequestModel.md) | Specifies the SmartAccess tags which grant access to the delivery group, if any occur in those provided by NetScaler Gateway with the user&#39;s connection. | [optional] 
**IncludedSmartAccessFilterType** | Pointer to [**FilterMatchType**](FilterMatchType.md) |  | [optional] 
**IncludedUserFilterEnabled** | Pointer to **NullableBool** | Specifies whether the IncludedUsers filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**IncludedUsers** | Pointer to **[]string** | Specifies the users and groups who are granted access to the delivery group. | [optional] 

## Methods

### NewAdvancedAccessPolicyRequestModel

`func NewAdvancedAccessPolicyRequestModel() *AdvancedAccessPolicyRequestModel`

NewAdvancedAccessPolicyRequestModel instantiates a new AdvancedAccessPolicyRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedAccessPolicyRequestModelWithDefaults

`func NewAdvancedAccessPolicyRequestModelWithDefaults() *AdvancedAccessPolicyRequestModel`

NewAdvancedAccessPolicyRequestModelWithDefaults instantiates a new AdvancedAccessPolicyRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdvancedAccessPolicyRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvancedAccessPolicyRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvancedAccessPolicyRequestModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvancedAccessPolicyRequestModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AdvancedAccessPolicyRequestModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AdvancedAccessPolicyRequestModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *AdvancedAccessPolicyRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvancedAccessPolicyRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvancedAccessPolicyRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvancedAccessPolicyRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdvancedAccessPolicyRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdvancedAccessPolicyRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAllowHdxAccess

`func (o *AdvancedAccessPolicyRequestModel) GetAllowHdxAccess() bool`

GetAllowHdxAccess returns the AllowHdxAccess field if non-nil, zero value otherwise.

### GetAllowHdxAccessOk

`func (o *AdvancedAccessPolicyRequestModel) GetAllowHdxAccessOk() (*bool, bool)`

GetAllowHdxAccessOk returns a tuple with the AllowHdxAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHdxAccess

`func (o *AdvancedAccessPolicyRequestModel) SetAllowHdxAccess(v bool)`

SetAllowHdxAccess sets AllowHdxAccess field to given value.

### HasAllowHdxAccess

`func (o *AdvancedAccessPolicyRequestModel) HasAllowHdxAccess() bool`

HasAllowHdxAccess returns a boolean if a field has been set.

### SetAllowHdxAccessNil

`func (o *AdvancedAccessPolicyRequestModel) SetAllowHdxAccessNil(b bool)`

 SetAllowHdxAccessNil sets the value for AllowHdxAccess to be an explicit nil

### UnsetAllowHdxAccess
`func (o *AdvancedAccessPolicyRequestModel) UnsetAllowHdxAccess()`

UnsetAllowHdxAccess ensures that no value is present for AllowHdxAccess, not even an explicit nil
### GetAllowMachineRestart

`func (o *AdvancedAccessPolicyRequestModel) GetAllowMachineRestart() bool`

GetAllowMachineRestart returns the AllowMachineRestart field if non-nil, zero value otherwise.

### GetAllowMachineRestartOk

`func (o *AdvancedAccessPolicyRequestModel) GetAllowMachineRestartOk() (*bool, bool)`

GetAllowMachineRestartOk returns a tuple with the AllowMachineRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMachineRestart

`func (o *AdvancedAccessPolicyRequestModel) SetAllowMachineRestart(v bool)`

SetAllowMachineRestart sets AllowMachineRestart field to given value.

### HasAllowMachineRestart

`func (o *AdvancedAccessPolicyRequestModel) HasAllowMachineRestart() bool`

HasAllowMachineRestart returns a boolean if a field has been set.

### SetAllowMachineRestartNil

`func (o *AdvancedAccessPolicyRequestModel) SetAllowMachineRestartNil(b bool)`

 SetAllowMachineRestartNil sets the value for AllowMachineRestart to be an explicit nil

### UnsetAllowMachineRestart
`func (o *AdvancedAccessPolicyRequestModel) UnsetAllowMachineRestart()`

UnsetAllowMachineRestart ensures that no value is present for AllowMachineRestart, not even an explicit nil
### GetAllowRdpAccess

`func (o *AdvancedAccessPolicyRequestModel) GetAllowRdpAccess() bool`

GetAllowRdpAccess returns the AllowRdpAccess field if non-nil, zero value otherwise.

### GetAllowRdpAccessOk

`func (o *AdvancedAccessPolicyRequestModel) GetAllowRdpAccessOk() (*bool, bool)`

GetAllowRdpAccessOk returns a tuple with the AllowRdpAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRdpAccess

`func (o *AdvancedAccessPolicyRequestModel) SetAllowRdpAccess(v bool)`

SetAllowRdpAccess sets AllowRdpAccess field to given value.

### HasAllowRdpAccess

`func (o *AdvancedAccessPolicyRequestModel) HasAllowRdpAccess() bool`

HasAllowRdpAccess returns a boolean if a field has been set.

### SetAllowRdpAccessNil

`func (o *AdvancedAccessPolicyRequestModel) SetAllowRdpAccessNil(b bool)`

 SetAllowRdpAccessNil sets the value for AllowRdpAccess to be an explicit nil

### UnsetAllowRdpAccess
`func (o *AdvancedAccessPolicyRequestModel) UnsetAllowRdpAccess()`

UnsetAllowRdpAccess ensures that no value is present for AllowRdpAccess, not even an explicit nil
### GetAllowedConnection

`func (o *AdvancedAccessPolicyRequestModel) GetAllowedConnection() AllowedConnection`

GetAllowedConnection returns the AllowedConnection field if non-nil, zero value otherwise.

### GetAllowedConnectionOk

`func (o *AdvancedAccessPolicyRequestModel) GetAllowedConnectionOk() (*AllowedConnection, bool)`

GetAllowedConnectionOk returns a tuple with the AllowedConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedConnection

`func (o *AdvancedAccessPolicyRequestModel) SetAllowedConnection(v AllowedConnection)`

SetAllowedConnection sets AllowedConnection field to given value.

### HasAllowedConnection

`func (o *AdvancedAccessPolicyRequestModel) HasAllowedConnection() bool`

HasAllowedConnection returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *AdvancedAccessPolicyRequestModel) GetAllowedUsers() AllowedUser`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *AdvancedAccessPolicyRequestModel) GetAllowedUsersOk() (*AllowedUser, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *AdvancedAccessPolicyRequestModel) SetAllowedUsers(v AllowedUser)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *AdvancedAccessPolicyRequestModel) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetDescription

`func (o *AdvancedAccessPolicyRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvancedAccessPolicyRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvancedAccessPolicyRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvancedAccessPolicyRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AdvancedAccessPolicyRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AdvancedAccessPolicyRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *AdvancedAccessPolicyRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdvancedAccessPolicyRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdvancedAccessPolicyRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AdvancedAccessPolicyRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *AdvancedAccessPolicyRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *AdvancedAccessPolicyRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetExcludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedClientIPFilterEnabled() bool`

GetExcludedClientIPFilterEnabled returns the ExcludedClientIPFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedClientIPFilterEnabledOk

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedClientIPFilterEnabledOk() (*bool, bool)`

GetExcludedClientIPFilterEnabledOk returns a tuple with the ExcludedClientIPFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedClientIPFilterEnabled(v bool)`

SetExcludedClientIPFilterEnabled sets ExcludedClientIPFilterEnabled field to given value.

### HasExcludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) HasExcludedClientIPFilterEnabled() bool`

HasExcludedClientIPFilterEnabled returns a boolean if a field has been set.

### SetExcludedClientIPFilterEnabledNil

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedClientIPFilterEnabledNil(b bool)`

 SetExcludedClientIPFilterEnabledNil sets the value for ExcludedClientIPFilterEnabled to be an explicit nil

### UnsetExcludedClientIPFilterEnabled
`func (o *AdvancedAccessPolicyRequestModel) UnsetExcludedClientIPFilterEnabled()`

UnsetExcludedClientIPFilterEnabled ensures that no value is present for ExcludedClientIPFilterEnabled, not even an explicit nil
### GetExcludedClientIPs

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedClientIPs() []string`

GetExcludedClientIPs returns the ExcludedClientIPs field if non-nil, zero value otherwise.

### GetExcludedClientIPsOk

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedClientIPsOk() (*[]string, bool)`

GetExcludedClientIPsOk returns a tuple with the ExcludedClientIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientIPs

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedClientIPs(v []string)`

SetExcludedClientIPs sets ExcludedClientIPs field to given value.

### HasExcludedClientIPs

`func (o *AdvancedAccessPolicyRequestModel) HasExcludedClientIPs() bool`

HasExcludedClientIPs returns a boolean if a field has been set.

### SetExcludedClientIPsNil

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedClientIPsNil(b bool)`

 SetExcludedClientIPsNil sets the value for ExcludedClientIPs to be an explicit nil

### UnsetExcludedClientIPs
`func (o *AdvancedAccessPolicyRequestModel) UnsetExcludedClientIPs()`

UnsetExcludedClientIPs ensures that no value is present for ExcludedClientIPs, not even an explicit nil
### GetExcludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedClientNameFilterEnabled() bool`

GetExcludedClientNameFilterEnabled returns the ExcludedClientNameFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedClientNameFilterEnabledOk

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedClientNameFilterEnabledOk() (*bool, bool)`

GetExcludedClientNameFilterEnabledOk returns a tuple with the ExcludedClientNameFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedClientNameFilterEnabled(v bool)`

SetExcludedClientNameFilterEnabled sets ExcludedClientNameFilterEnabled field to given value.

### HasExcludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) HasExcludedClientNameFilterEnabled() bool`

HasExcludedClientNameFilterEnabled returns a boolean if a field has been set.

### SetExcludedClientNameFilterEnabledNil

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedClientNameFilterEnabledNil(b bool)`

 SetExcludedClientNameFilterEnabledNil sets the value for ExcludedClientNameFilterEnabled to be an explicit nil

### UnsetExcludedClientNameFilterEnabled
`func (o *AdvancedAccessPolicyRequestModel) UnsetExcludedClientNameFilterEnabled()`

UnsetExcludedClientNameFilterEnabled ensures that no value is present for ExcludedClientNameFilterEnabled, not even an explicit nil
### GetExcludedClientNames

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedClientNames() []string`

GetExcludedClientNames returns the ExcludedClientNames field if non-nil, zero value otherwise.

### GetExcludedClientNamesOk

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedClientNamesOk() (*[]string, bool)`

GetExcludedClientNamesOk returns a tuple with the ExcludedClientNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientNames

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedClientNames(v []string)`

SetExcludedClientNames sets ExcludedClientNames field to given value.

### HasExcludedClientNames

`func (o *AdvancedAccessPolicyRequestModel) HasExcludedClientNames() bool`

HasExcludedClientNames returns a boolean if a field has been set.

### SetExcludedClientNamesNil

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedClientNamesNil(b bool)`

 SetExcludedClientNamesNil sets the value for ExcludedClientNames to be an explicit nil

### UnsetExcludedClientNames
`func (o *AdvancedAccessPolicyRequestModel) UnsetExcludedClientNames()`

UnsetExcludedClientNames ensures that no value is present for ExcludedClientNames, not even an explicit nil
### GetExcludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedSmartAccessFilterEnabled() bool`

GetExcludedSmartAccessFilterEnabled returns the ExcludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedSmartAccessFilterEnabledOk

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetExcludedSmartAccessFilterEnabledOk returns a tuple with the ExcludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedSmartAccessFilterEnabled(v bool)`

SetExcludedSmartAccessFilterEnabled sets ExcludedSmartAccessFilterEnabled field to given value.

### HasExcludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) HasExcludedSmartAccessFilterEnabled() bool`

HasExcludedSmartAccessFilterEnabled returns a boolean if a field has been set.

### SetExcludedSmartAccessFilterEnabledNil

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedSmartAccessFilterEnabledNil(b bool)`

 SetExcludedSmartAccessFilterEnabledNil sets the value for ExcludedSmartAccessFilterEnabled to be an explicit nil

### UnsetExcludedSmartAccessFilterEnabled
`func (o *AdvancedAccessPolicyRequestModel) UnsetExcludedSmartAccessFilterEnabled()`

UnsetExcludedSmartAccessFilterEnabled ensures that no value is present for ExcludedSmartAccessFilterEnabled, not even an explicit nil
### GetExcludedSmartAccessTags

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedSmartAccessTags() []SmartAccessTagRequestModel`

GetExcludedSmartAccessTags returns the ExcludedSmartAccessTags field if non-nil, zero value otherwise.

### GetExcludedSmartAccessTagsOk

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedSmartAccessTagsOk() (*[]SmartAccessTagRequestModel, bool)`

GetExcludedSmartAccessTagsOk returns a tuple with the ExcludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSmartAccessTags

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedSmartAccessTags(v []SmartAccessTagRequestModel)`

SetExcludedSmartAccessTags sets ExcludedSmartAccessTags field to given value.

### HasExcludedSmartAccessTags

`func (o *AdvancedAccessPolicyRequestModel) HasExcludedSmartAccessTags() bool`

HasExcludedSmartAccessTags returns a boolean if a field has been set.

### SetExcludedSmartAccessTagsNil

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedSmartAccessTagsNil(b bool)`

 SetExcludedSmartAccessTagsNil sets the value for ExcludedSmartAccessTags to be an explicit nil

### UnsetExcludedSmartAccessTags
`func (o *AdvancedAccessPolicyRequestModel) UnsetExcludedSmartAccessTags()`

UnsetExcludedSmartAccessTags ensures that no value is present for ExcludedSmartAccessTags, not even an explicit nil
### GetExcludedUserFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.

### HasExcludedUserFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) HasExcludedUserFilterEnabled() bool`

HasExcludedUserFilterEnabled returns a boolean if a field has been set.

### SetExcludedUserFilterEnabledNil

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedUserFilterEnabledNil(b bool)`

 SetExcludedUserFilterEnabledNil sets the value for ExcludedUserFilterEnabled to be an explicit nil

### UnsetExcludedUserFilterEnabled
`func (o *AdvancedAccessPolicyRequestModel) UnsetExcludedUserFilterEnabled()`

UnsetExcludedUserFilterEnabled ensures that no value is present for ExcludedUserFilterEnabled, not even an explicit nil
### GetExcludedUsers

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedUsers() []string`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *AdvancedAccessPolicyRequestModel) GetExcludedUsersOk() (*[]string, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedUsers(v []string)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *AdvancedAccessPolicyRequestModel) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.

### SetExcludedUsersNil

`func (o *AdvancedAccessPolicyRequestModel) SetExcludedUsersNil(b bool)`

 SetExcludedUsersNil sets the value for ExcludedUsers to be an explicit nil

### UnsetExcludedUsers
`func (o *AdvancedAccessPolicyRequestModel) UnsetExcludedUsers()`

UnsetExcludedUsers ensures that no value is present for ExcludedUsers, not even an explicit nil
### GetIncludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedClientIPFilterEnabled() bool`

GetIncludedClientIPFilterEnabled returns the IncludedClientIPFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedClientIPFilterEnabledOk

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedClientIPFilterEnabledOk() (*bool, bool)`

GetIncludedClientIPFilterEnabledOk returns a tuple with the IncludedClientIPFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedClientIPFilterEnabled(v bool)`

SetIncludedClientIPFilterEnabled sets IncludedClientIPFilterEnabled field to given value.

### HasIncludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) HasIncludedClientIPFilterEnabled() bool`

HasIncludedClientIPFilterEnabled returns a boolean if a field has been set.

### SetIncludedClientIPFilterEnabledNil

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedClientIPFilterEnabledNil(b bool)`

 SetIncludedClientIPFilterEnabledNil sets the value for IncludedClientIPFilterEnabled to be an explicit nil

### UnsetIncludedClientIPFilterEnabled
`func (o *AdvancedAccessPolicyRequestModel) UnsetIncludedClientIPFilterEnabled()`

UnsetIncludedClientIPFilterEnabled ensures that no value is present for IncludedClientIPFilterEnabled, not even an explicit nil
### GetIncludedClientIPs

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedClientIPs() []string`

GetIncludedClientIPs returns the IncludedClientIPs field if non-nil, zero value otherwise.

### GetIncludedClientIPsOk

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedClientIPsOk() (*[]string, bool)`

GetIncludedClientIPsOk returns a tuple with the IncludedClientIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientIPs

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedClientIPs(v []string)`

SetIncludedClientIPs sets IncludedClientIPs field to given value.

### HasIncludedClientIPs

`func (o *AdvancedAccessPolicyRequestModel) HasIncludedClientIPs() bool`

HasIncludedClientIPs returns a boolean if a field has been set.

### SetIncludedClientIPsNil

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedClientIPsNil(b bool)`

 SetIncludedClientIPsNil sets the value for IncludedClientIPs to be an explicit nil

### UnsetIncludedClientIPs
`func (o *AdvancedAccessPolicyRequestModel) UnsetIncludedClientIPs()`

UnsetIncludedClientIPs ensures that no value is present for IncludedClientIPs, not even an explicit nil
### GetIncludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedClientNameFilterEnabled() bool`

GetIncludedClientNameFilterEnabled returns the IncludedClientNameFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedClientNameFilterEnabledOk

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedClientNameFilterEnabledOk() (*bool, bool)`

GetIncludedClientNameFilterEnabledOk returns a tuple with the IncludedClientNameFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedClientNameFilterEnabled(v bool)`

SetIncludedClientNameFilterEnabled sets IncludedClientNameFilterEnabled field to given value.

### HasIncludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) HasIncludedClientNameFilterEnabled() bool`

HasIncludedClientNameFilterEnabled returns a boolean if a field has been set.

### SetIncludedClientNameFilterEnabledNil

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedClientNameFilterEnabledNil(b bool)`

 SetIncludedClientNameFilterEnabledNil sets the value for IncludedClientNameFilterEnabled to be an explicit nil

### UnsetIncludedClientNameFilterEnabled
`func (o *AdvancedAccessPolicyRequestModel) UnsetIncludedClientNameFilterEnabled()`

UnsetIncludedClientNameFilterEnabled ensures that no value is present for IncludedClientNameFilterEnabled, not even an explicit nil
### GetIncludedClientNames

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedClientNames() []string`

GetIncludedClientNames returns the IncludedClientNames field if non-nil, zero value otherwise.

### GetIncludedClientNamesOk

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedClientNamesOk() (*[]string, bool)`

GetIncludedClientNamesOk returns a tuple with the IncludedClientNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientNames

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedClientNames(v []string)`

SetIncludedClientNames sets IncludedClientNames field to given value.

### HasIncludedClientNames

`func (o *AdvancedAccessPolicyRequestModel) HasIncludedClientNames() bool`

HasIncludedClientNames returns a boolean if a field has been set.

### SetIncludedClientNamesNil

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedClientNamesNil(b bool)`

 SetIncludedClientNamesNil sets the value for IncludedClientNames to be an explicit nil

### UnsetIncludedClientNames
`func (o *AdvancedAccessPolicyRequestModel) UnsetIncludedClientNames()`

UnsetIncludedClientNames ensures that no value is present for IncludedClientNames, not even an explicit nil
### GetIncludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedSmartAccessFilterEnabled() bool`

GetIncludedSmartAccessFilterEnabled returns the IncludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterEnabledOk

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetIncludedSmartAccessFilterEnabledOk returns a tuple with the IncludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedSmartAccessFilterEnabled(v bool)`

SetIncludedSmartAccessFilterEnabled sets IncludedSmartAccessFilterEnabled field to given value.

### HasIncludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) HasIncludedSmartAccessFilterEnabled() bool`

HasIncludedSmartAccessFilterEnabled returns a boolean if a field has been set.

### SetIncludedSmartAccessFilterEnabledNil

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedSmartAccessFilterEnabledNil(b bool)`

 SetIncludedSmartAccessFilterEnabledNil sets the value for IncludedSmartAccessFilterEnabled to be an explicit nil

### UnsetIncludedSmartAccessFilterEnabled
`func (o *AdvancedAccessPolicyRequestModel) UnsetIncludedSmartAccessFilterEnabled()`

UnsetIncludedSmartAccessFilterEnabled ensures that no value is present for IncludedSmartAccessFilterEnabled, not even an explicit nil
### GetIncludedSmartAccessTags

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedSmartAccessTags() []SmartAccessTagRequestModel`

GetIncludedSmartAccessTags returns the IncludedSmartAccessTags field if non-nil, zero value otherwise.

### GetIncludedSmartAccessTagsOk

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedSmartAccessTagsOk() (*[]SmartAccessTagRequestModel, bool)`

GetIncludedSmartAccessTagsOk returns a tuple with the IncludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessTags

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedSmartAccessTags(v []SmartAccessTagRequestModel)`

SetIncludedSmartAccessTags sets IncludedSmartAccessTags field to given value.

### HasIncludedSmartAccessTags

`func (o *AdvancedAccessPolicyRequestModel) HasIncludedSmartAccessTags() bool`

HasIncludedSmartAccessTags returns a boolean if a field has been set.

### SetIncludedSmartAccessTagsNil

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedSmartAccessTagsNil(b bool)`

 SetIncludedSmartAccessTagsNil sets the value for IncludedSmartAccessTags to be an explicit nil

### UnsetIncludedSmartAccessTags
`func (o *AdvancedAccessPolicyRequestModel) UnsetIncludedSmartAccessTags()`

UnsetIncludedSmartAccessTags ensures that no value is present for IncludedSmartAccessTags, not even an explicit nil
### GetIncludedSmartAccessFilterType

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedSmartAccessFilterType() FilterMatchType`

GetIncludedSmartAccessFilterType returns the IncludedSmartAccessFilterType field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterTypeOk

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedSmartAccessFilterTypeOk() (*FilterMatchType, bool)`

GetIncludedSmartAccessFilterTypeOk returns a tuple with the IncludedSmartAccessFilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterType

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedSmartAccessFilterType(v FilterMatchType)`

SetIncludedSmartAccessFilterType sets IncludedSmartAccessFilterType field to given value.

### HasIncludedSmartAccessFilterType

`func (o *AdvancedAccessPolicyRequestModel) HasIncludedSmartAccessFilterType() bool`

HasIncludedSmartAccessFilterType returns a boolean if a field has been set.

### GetIncludedUserFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *AdvancedAccessPolicyRequestModel) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### SetIncludedUserFilterEnabledNil

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedUserFilterEnabledNil(b bool)`

 SetIncludedUserFilterEnabledNil sets the value for IncludedUserFilterEnabled to be an explicit nil

### UnsetIncludedUserFilterEnabled
`func (o *AdvancedAccessPolicyRequestModel) UnsetIncludedUserFilterEnabled()`

UnsetIncludedUserFilterEnabled ensures that no value is present for IncludedUserFilterEnabled, not even an explicit nil
### GetIncludedUsers

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *AdvancedAccessPolicyRequestModel) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *AdvancedAccessPolicyRequestModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *AdvancedAccessPolicyRequestModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *AdvancedAccessPolicyRequestModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


