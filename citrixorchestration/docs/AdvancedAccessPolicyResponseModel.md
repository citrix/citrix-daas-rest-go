# AdvancedAccessPolicyResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id | [optional] 
**Uid** | Pointer to **NullableInt32** | &#x60;DEPRECATED&#x60; DEPRECATED. Use Id. | [optional] 
**Name** | Pointer to **NullableString** | Name of the advanced access policy. | [optional] 
**AllowHdxAccess** | Pointer to **bool** | Whether HDX connections are allowed to the delivery group. CHANGE: was: public bool AllowsHdxAccess { get; set; } | [optional] 
**AllowMachineRestart** | Pointer to **bool** | Whether users are allowed to self-service restart machines. | [optional] 
**AllowRdpAccess** | Pointer to **bool** | Whether RDP connections are allowed to the delivery group. CHANGE: was: public bool AllowsRdpAccess { get; set; } | [optional] 
**AllowedConnection** | Pointer to [**AllowedConnection**](AllowedConnection.md) |  | [optional] 
**AllowedUsers** | Pointer to [**AllowedUser**](AllowedUser.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Description. | [optional] 
**Enabled** | **bool** | Whether the advanced access policy is enabled.  If a policy is disabled it is not considered when evaluating whether a user may access the delivery group. | 
**IsBuiltIn** | Pointer to **bool** | Whether the access policy is a built-in policy.  If true , the policy is not allowed to remove or rename. | [optional] 
**IsBroken** | Pointer to **bool** | Whether the access policy is broken. | [optional] 
**ExcludedClientIPFilterEnabled** | Pointer to **bool** | Indicates whether the ExcludedClientIPs filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**ExcludedClientIPs** | Pointer to **[]string** | The client IPs which will be denied access to the delivery group. | [optional] 
**ExcludedClientNameFilterEnabled** | Pointer to **bool** | Indicates whether the ExcludedClientNames filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**ExcludedClientNames** | Pointer to **[]string** | The client names which will be denied access to the delivery group. | [optional] 
**ExcludedSmartAccessFilterEnabled** | Pointer to **bool** | Indicates whether the ExcludedSmartAccessTags filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**ExcludedSmartAccessTags** | Pointer to [**[]SmartAccessTagResponseModel**](SmartAccessTagResponseModel.md) | The SmartAccess tags which will be denied access to the delivery group, if any occur in those provided by NetScaler Gateway with the user&#39;s connection. | [optional] 
**ExcludedUserFilterEnabled** | Pointer to **bool** | Indicates whether the ExcludedUsers filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**ExcludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The users and groups who are denied access to the delivery group. | [optional] 
**IncludedClientIPFilterEnabled** | Pointer to **bool** | Indicates whether the IncludedClientIPs filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**IncludedClientIPs** | Pointer to **[]string** | The client IPs which will be allowed access to the delivery group. | [optional] 
**IncludedClientNameFilterEnabled** | Pointer to **bool** | Indicates whether the IncludedClientNames filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**IncludedClientNames** | Pointer to **[]string** | The client names which will be granted access to the delivery group. | [optional] 
**IncludedSmartAccessFilterEnabled** | Pointer to **bool** | Indicates whether the IncludedSmartAccessTags filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. CHANGE: was: public bool AccessGatewayConnectionsUseFilters { get; set; } | [optional] 
**IncludedSmartAccessTags** | Pointer to [**[]SmartAccessTagResponseModel**](SmartAccessTagResponseModel.md) | The SmartAccess tags which grant access to the delivery group, if any occur in those provided by NetScaler Gateway with the user&#39;s connection. | [optional] 
**IncludedSmartAccessFilterType** | Pointer to [**FilterMatchType**](FilterMatchType.md) |  | [optional] 
**IncludedUserFilterEnabled** | Pointer to **bool** | Indicates whether the IncludedUsers filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**IncludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The users and groups who are granted access to the delivery group. CHANGE: was: public ADAccount[] IncludedUsers { get; set; } | [optional] 

## Methods

### NewAdvancedAccessPolicyResponseModel

`func NewAdvancedAccessPolicyResponseModel(enabled bool, ) *AdvancedAccessPolicyResponseModel`

NewAdvancedAccessPolicyResponseModel instantiates a new AdvancedAccessPolicyResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedAccessPolicyResponseModelWithDefaults

`func NewAdvancedAccessPolicyResponseModelWithDefaults() *AdvancedAccessPolicyResponseModel`

NewAdvancedAccessPolicyResponseModelWithDefaults instantiates a new AdvancedAccessPolicyResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdvancedAccessPolicyResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvancedAccessPolicyResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvancedAccessPolicyResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvancedAccessPolicyResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AdvancedAccessPolicyResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AdvancedAccessPolicyResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *AdvancedAccessPolicyResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AdvancedAccessPolicyResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AdvancedAccessPolicyResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AdvancedAccessPolicyResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *AdvancedAccessPolicyResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *AdvancedAccessPolicyResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetName

`func (o *AdvancedAccessPolicyResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvancedAccessPolicyResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvancedAccessPolicyResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvancedAccessPolicyResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdvancedAccessPolicyResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdvancedAccessPolicyResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAllowHdxAccess

`func (o *AdvancedAccessPolicyResponseModel) GetAllowHdxAccess() bool`

GetAllowHdxAccess returns the AllowHdxAccess field if non-nil, zero value otherwise.

### GetAllowHdxAccessOk

`func (o *AdvancedAccessPolicyResponseModel) GetAllowHdxAccessOk() (*bool, bool)`

GetAllowHdxAccessOk returns a tuple with the AllowHdxAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHdxAccess

`func (o *AdvancedAccessPolicyResponseModel) SetAllowHdxAccess(v bool)`

SetAllowHdxAccess sets AllowHdxAccess field to given value.

### HasAllowHdxAccess

`func (o *AdvancedAccessPolicyResponseModel) HasAllowHdxAccess() bool`

HasAllowHdxAccess returns a boolean if a field has been set.

### GetAllowMachineRestart

`func (o *AdvancedAccessPolicyResponseModel) GetAllowMachineRestart() bool`

GetAllowMachineRestart returns the AllowMachineRestart field if non-nil, zero value otherwise.

### GetAllowMachineRestartOk

`func (o *AdvancedAccessPolicyResponseModel) GetAllowMachineRestartOk() (*bool, bool)`

GetAllowMachineRestartOk returns a tuple with the AllowMachineRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMachineRestart

`func (o *AdvancedAccessPolicyResponseModel) SetAllowMachineRestart(v bool)`

SetAllowMachineRestart sets AllowMachineRestart field to given value.

### HasAllowMachineRestart

`func (o *AdvancedAccessPolicyResponseModel) HasAllowMachineRestart() bool`

HasAllowMachineRestart returns a boolean if a field has been set.

### GetAllowRdpAccess

`func (o *AdvancedAccessPolicyResponseModel) GetAllowRdpAccess() bool`

GetAllowRdpAccess returns the AllowRdpAccess field if non-nil, zero value otherwise.

### GetAllowRdpAccessOk

`func (o *AdvancedAccessPolicyResponseModel) GetAllowRdpAccessOk() (*bool, bool)`

GetAllowRdpAccessOk returns a tuple with the AllowRdpAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRdpAccess

`func (o *AdvancedAccessPolicyResponseModel) SetAllowRdpAccess(v bool)`

SetAllowRdpAccess sets AllowRdpAccess field to given value.

### HasAllowRdpAccess

`func (o *AdvancedAccessPolicyResponseModel) HasAllowRdpAccess() bool`

HasAllowRdpAccess returns a boolean if a field has been set.

### GetAllowedConnection

`func (o *AdvancedAccessPolicyResponseModel) GetAllowedConnection() AllowedConnection`

GetAllowedConnection returns the AllowedConnection field if non-nil, zero value otherwise.

### GetAllowedConnectionOk

`func (o *AdvancedAccessPolicyResponseModel) GetAllowedConnectionOk() (*AllowedConnection, bool)`

GetAllowedConnectionOk returns a tuple with the AllowedConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedConnection

`func (o *AdvancedAccessPolicyResponseModel) SetAllowedConnection(v AllowedConnection)`

SetAllowedConnection sets AllowedConnection field to given value.

### HasAllowedConnection

`func (o *AdvancedAccessPolicyResponseModel) HasAllowedConnection() bool`

HasAllowedConnection returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *AdvancedAccessPolicyResponseModel) GetAllowedUsers() AllowedUser`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *AdvancedAccessPolicyResponseModel) GetAllowedUsersOk() (*AllowedUser, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *AdvancedAccessPolicyResponseModel) SetAllowedUsers(v AllowedUser)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *AdvancedAccessPolicyResponseModel) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetDescription

`func (o *AdvancedAccessPolicyResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvancedAccessPolicyResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvancedAccessPolicyResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvancedAccessPolicyResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AdvancedAccessPolicyResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AdvancedAccessPolicyResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *AdvancedAccessPolicyResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdvancedAccessPolicyResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdvancedAccessPolicyResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIsBuiltIn

`func (o *AdvancedAccessPolicyResponseModel) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *AdvancedAccessPolicyResponseModel) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *AdvancedAccessPolicyResponseModel) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.

### HasIsBuiltIn

`func (o *AdvancedAccessPolicyResponseModel) HasIsBuiltIn() bool`

HasIsBuiltIn returns a boolean if a field has been set.

### GetIsBroken

`func (o *AdvancedAccessPolicyResponseModel) GetIsBroken() bool`

GetIsBroken returns the IsBroken field if non-nil, zero value otherwise.

### GetIsBrokenOk

`func (o *AdvancedAccessPolicyResponseModel) GetIsBrokenOk() (*bool, bool)`

GetIsBrokenOk returns a tuple with the IsBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBroken

`func (o *AdvancedAccessPolicyResponseModel) SetIsBroken(v bool)`

SetIsBroken sets IsBroken field to given value.

### HasIsBroken

`func (o *AdvancedAccessPolicyResponseModel) HasIsBroken() bool`

HasIsBroken returns a boolean if a field has been set.

### GetExcludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedClientIPFilterEnabled() bool`

GetExcludedClientIPFilterEnabled returns the ExcludedClientIPFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedClientIPFilterEnabledOk

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedClientIPFilterEnabledOk() (*bool, bool)`

GetExcludedClientIPFilterEnabledOk returns a tuple with the ExcludedClientIPFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedClientIPFilterEnabled(v bool)`

SetExcludedClientIPFilterEnabled sets ExcludedClientIPFilterEnabled field to given value.

### HasExcludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) HasExcludedClientIPFilterEnabled() bool`

HasExcludedClientIPFilterEnabled returns a boolean if a field has been set.

### GetExcludedClientIPs

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedClientIPs() []string`

GetExcludedClientIPs returns the ExcludedClientIPs field if non-nil, zero value otherwise.

### GetExcludedClientIPsOk

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedClientIPsOk() (*[]string, bool)`

GetExcludedClientIPsOk returns a tuple with the ExcludedClientIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientIPs

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedClientIPs(v []string)`

SetExcludedClientIPs sets ExcludedClientIPs field to given value.

### HasExcludedClientIPs

`func (o *AdvancedAccessPolicyResponseModel) HasExcludedClientIPs() bool`

HasExcludedClientIPs returns a boolean if a field has been set.

### SetExcludedClientIPsNil

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedClientIPsNil(b bool)`

 SetExcludedClientIPsNil sets the value for ExcludedClientIPs to be an explicit nil

### UnsetExcludedClientIPs
`func (o *AdvancedAccessPolicyResponseModel) UnsetExcludedClientIPs()`

UnsetExcludedClientIPs ensures that no value is present for ExcludedClientIPs, not even an explicit nil
### GetExcludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedClientNameFilterEnabled() bool`

GetExcludedClientNameFilterEnabled returns the ExcludedClientNameFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedClientNameFilterEnabledOk

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedClientNameFilterEnabledOk() (*bool, bool)`

GetExcludedClientNameFilterEnabledOk returns a tuple with the ExcludedClientNameFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedClientNameFilterEnabled(v bool)`

SetExcludedClientNameFilterEnabled sets ExcludedClientNameFilterEnabled field to given value.

### HasExcludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) HasExcludedClientNameFilterEnabled() bool`

HasExcludedClientNameFilterEnabled returns a boolean if a field has been set.

### GetExcludedClientNames

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedClientNames() []string`

GetExcludedClientNames returns the ExcludedClientNames field if non-nil, zero value otherwise.

### GetExcludedClientNamesOk

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedClientNamesOk() (*[]string, bool)`

GetExcludedClientNamesOk returns a tuple with the ExcludedClientNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClientNames

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedClientNames(v []string)`

SetExcludedClientNames sets ExcludedClientNames field to given value.

### HasExcludedClientNames

`func (o *AdvancedAccessPolicyResponseModel) HasExcludedClientNames() bool`

HasExcludedClientNames returns a boolean if a field has been set.

### SetExcludedClientNamesNil

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedClientNamesNil(b bool)`

 SetExcludedClientNamesNil sets the value for ExcludedClientNames to be an explicit nil

### UnsetExcludedClientNames
`func (o *AdvancedAccessPolicyResponseModel) UnsetExcludedClientNames()`

UnsetExcludedClientNames ensures that no value is present for ExcludedClientNames, not even an explicit nil
### GetExcludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedSmartAccessFilterEnabled() bool`

GetExcludedSmartAccessFilterEnabled returns the ExcludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedSmartAccessFilterEnabledOk

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetExcludedSmartAccessFilterEnabledOk returns a tuple with the ExcludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedSmartAccessFilterEnabled(v bool)`

SetExcludedSmartAccessFilterEnabled sets ExcludedSmartAccessFilterEnabled field to given value.

### HasExcludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) HasExcludedSmartAccessFilterEnabled() bool`

HasExcludedSmartAccessFilterEnabled returns a boolean if a field has been set.

### GetExcludedSmartAccessTags

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedSmartAccessTags() []SmartAccessTagResponseModel`

GetExcludedSmartAccessTags returns the ExcludedSmartAccessTags field if non-nil, zero value otherwise.

### GetExcludedSmartAccessTagsOk

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedSmartAccessTagsOk() (*[]SmartAccessTagResponseModel, bool)`

GetExcludedSmartAccessTagsOk returns a tuple with the ExcludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSmartAccessTags

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedSmartAccessTags(v []SmartAccessTagResponseModel)`

SetExcludedSmartAccessTags sets ExcludedSmartAccessTags field to given value.

### HasExcludedSmartAccessTags

`func (o *AdvancedAccessPolicyResponseModel) HasExcludedSmartAccessTags() bool`

HasExcludedSmartAccessTags returns a boolean if a field has been set.

### SetExcludedSmartAccessTagsNil

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedSmartAccessTagsNil(b bool)`

 SetExcludedSmartAccessTagsNil sets the value for ExcludedSmartAccessTags to be an explicit nil

### UnsetExcludedSmartAccessTags
`func (o *AdvancedAccessPolicyResponseModel) UnsetExcludedSmartAccessTags()`

UnsetExcludedSmartAccessTags ensures that no value is present for ExcludedSmartAccessTags, not even an explicit nil
### GetExcludedUserFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.

### HasExcludedUserFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) HasExcludedUserFilterEnabled() bool`

HasExcludedUserFilterEnabled returns a boolean if a field has been set.

### GetExcludedUsers

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedUsers() []IdentityUserResponseModel`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *AdvancedAccessPolicyResponseModel) GetExcludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedUsers(v []IdentityUserResponseModel)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *AdvancedAccessPolicyResponseModel) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.

### SetExcludedUsersNil

`func (o *AdvancedAccessPolicyResponseModel) SetExcludedUsersNil(b bool)`

 SetExcludedUsersNil sets the value for ExcludedUsers to be an explicit nil

### UnsetExcludedUsers
`func (o *AdvancedAccessPolicyResponseModel) UnsetExcludedUsers()`

UnsetExcludedUsers ensures that no value is present for ExcludedUsers, not even an explicit nil
### GetIncludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedClientIPFilterEnabled() bool`

GetIncludedClientIPFilterEnabled returns the IncludedClientIPFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedClientIPFilterEnabledOk

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedClientIPFilterEnabledOk() (*bool, bool)`

GetIncludedClientIPFilterEnabledOk returns a tuple with the IncludedClientIPFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedClientIPFilterEnabled(v bool)`

SetIncludedClientIPFilterEnabled sets IncludedClientIPFilterEnabled field to given value.

### HasIncludedClientIPFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) HasIncludedClientIPFilterEnabled() bool`

HasIncludedClientIPFilterEnabled returns a boolean if a field has been set.

### GetIncludedClientIPs

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedClientIPs() []string`

GetIncludedClientIPs returns the IncludedClientIPs field if non-nil, zero value otherwise.

### GetIncludedClientIPsOk

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedClientIPsOk() (*[]string, bool)`

GetIncludedClientIPsOk returns a tuple with the IncludedClientIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientIPs

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedClientIPs(v []string)`

SetIncludedClientIPs sets IncludedClientIPs field to given value.

### HasIncludedClientIPs

`func (o *AdvancedAccessPolicyResponseModel) HasIncludedClientIPs() bool`

HasIncludedClientIPs returns a boolean if a field has been set.

### SetIncludedClientIPsNil

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedClientIPsNil(b bool)`

 SetIncludedClientIPsNil sets the value for IncludedClientIPs to be an explicit nil

### UnsetIncludedClientIPs
`func (o *AdvancedAccessPolicyResponseModel) UnsetIncludedClientIPs()`

UnsetIncludedClientIPs ensures that no value is present for IncludedClientIPs, not even an explicit nil
### GetIncludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedClientNameFilterEnabled() bool`

GetIncludedClientNameFilterEnabled returns the IncludedClientNameFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedClientNameFilterEnabledOk

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedClientNameFilterEnabledOk() (*bool, bool)`

GetIncludedClientNameFilterEnabledOk returns a tuple with the IncludedClientNameFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedClientNameFilterEnabled(v bool)`

SetIncludedClientNameFilterEnabled sets IncludedClientNameFilterEnabled field to given value.

### HasIncludedClientNameFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) HasIncludedClientNameFilterEnabled() bool`

HasIncludedClientNameFilterEnabled returns a boolean if a field has been set.

### GetIncludedClientNames

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedClientNames() []string`

GetIncludedClientNames returns the IncludedClientNames field if non-nil, zero value otherwise.

### GetIncludedClientNamesOk

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedClientNamesOk() (*[]string, bool)`

GetIncludedClientNamesOk returns a tuple with the IncludedClientNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedClientNames

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedClientNames(v []string)`

SetIncludedClientNames sets IncludedClientNames field to given value.

### HasIncludedClientNames

`func (o *AdvancedAccessPolicyResponseModel) HasIncludedClientNames() bool`

HasIncludedClientNames returns a boolean if a field has been set.

### SetIncludedClientNamesNil

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedClientNamesNil(b bool)`

 SetIncludedClientNamesNil sets the value for IncludedClientNames to be an explicit nil

### UnsetIncludedClientNames
`func (o *AdvancedAccessPolicyResponseModel) UnsetIncludedClientNames()`

UnsetIncludedClientNames ensures that no value is present for IncludedClientNames, not even an explicit nil
### GetIncludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedSmartAccessFilterEnabled() bool`

GetIncludedSmartAccessFilterEnabled returns the IncludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterEnabledOk

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetIncludedSmartAccessFilterEnabledOk returns a tuple with the IncludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedSmartAccessFilterEnabled(v bool)`

SetIncludedSmartAccessFilterEnabled sets IncludedSmartAccessFilterEnabled field to given value.

### HasIncludedSmartAccessFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) HasIncludedSmartAccessFilterEnabled() bool`

HasIncludedSmartAccessFilterEnabled returns a boolean if a field has been set.

### GetIncludedSmartAccessTags

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedSmartAccessTags() []SmartAccessTagResponseModel`

GetIncludedSmartAccessTags returns the IncludedSmartAccessTags field if non-nil, zero value otherwise.

### GetIncludedSmartAccessTagsOk

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedSmartAccessTagsOk() (*[]SmartAccessTagResponseModel, bool)`

GetIncludedSmartAccessTagsOk returns a tuple with the IncludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessTags

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedSmartAccessTags(v []SmartAccessTagResponseModel)`

SetIncludedSmartAccessTags sets IncludedSmartAccessTags field to given value.

### HasIncludedSmartAccessTags

`func (o *AdvancedAccessPolicyResponseModel) HasIncludedSmartAccessTags() bool`

HasIncludedSmartAccessTags returns a boolean if a field has been set.

### SetIncludedSmartAccessTagsNil

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedSmartAccessTagsNil(b bool)`

 SetIncludedSmartAccessTagsNil sets the value for IncludedSmartAccessTags to be an explicit nil

### UnsetIncludedSmartAccessTags
`func (o *AdvancedAccessPolicyResponseModel) UnsetIncludedSmartAccessTags()`

UnsetIncludedSmartAccessTags ensures that no value is present for IncludedSmartAccessTags, not even an explicit nil
### GetIncludedSmartAccessFilterType

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedSmartAccessFilterType() FilterMatchType`

GetIncludedSmartAccessFilterType returns the IncludedSmartAccessFilterType field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterTypeOk

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedSmartAccessFilterTypeOk() (*FilterMatchType, bool)`

GetIncludedSmartAccessFilterTypeOk returns a tuple with the IncludedSmartAccessFilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterType

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedSmartAccessFilterType(v FilterMatchType)`

SetIncludedSmartAccessFilterType sets IncludedSmartAccessFilterType field to given value.

### HasIncludedSmartAccessFilterType

`func (o *AdvancedAccessPolicyResponseModel) HasIncludedSmartAccessFilterType() bool`

HasIncludedSmartAccessFilterType returns a boolean if a field has been set.

### GetIncludedUserFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *AdvancedAccessPolicyResponseModel) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### GetIncludedUsers

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *AdvancedAccessPolicyResponseModel) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *AdvancedAccessPolicyResponseModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *AdvancedAccessPolicyResponseModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *AdvancedAccessPolicyResponseModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


