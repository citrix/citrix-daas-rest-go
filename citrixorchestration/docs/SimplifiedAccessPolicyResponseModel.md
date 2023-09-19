# SimplifiedAccessPolicyResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAnonymous** | **bool** | Whether anonymous connections are allowed to the delivery group. | 
**AllowHdxAccess** | **bool** | Whether HDX connections are allowed to the delivery group. CHANGE: was: public bool AllowsHdxAccess { get; set; } | 
**AllowMachineRestart** | Pointer to **NullableBool** | Whether users are allowed to self-service restart machines. | [optional] 
**AllowRdpAccess** | **bool** | Whether RDP connections are allowed to the delivery group. CHANGE: was: public bool AllowsRdpAccess { get; set; } | 
**ConnectNotViaNetScalerGatewayAllowed** | **bool** | Whether connections that do not come through NetScaler Gateway are allowed. CHANGE: was: public bool ConnectNotViaAccessGatewayAllowed { get; set; } | 
**ConnectViaNetScalerGatewayAllowed** | **bool** | Whether connections that come through NetScaler Gateway are allowed. CHANGE: was: public bool ConnectViaAccessGatewayAllowed { get; set; } | 
**IncludedSmartAccessFilterEnabled** | **bool** | Indicates whether the IncludedSmartAccessTags filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. CHANGE: was: public bool AccessGatewayConnectionsUseFilters { get; set; } | 
**IncludedSmartAccessTags** | Pointer to [**[]SmartAccessTagResponseModel**](SmartAccessTagResponseModel.md) | The SmartAccess tags which grant access to the delivery group, if any occur in those provided by NetScaler Gateway with the user&#39;s connection. | [optional] 
**IncludedUserFilterEnabled** | **bool** | Indicates whether the IncludedUsers filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | 
**IncludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The users and groups who are granted access to the delivery group. CHANGE: was: public ADAccount[] IncludedUsers { get; set; } | [optional] 
**ExcludedUserFilterEnabled** | **bool** | Indicates whether the ExcludedUsers filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | 
**ExcludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The users and groups who are denied access to the delivery group. | [optional] 

## Methods

### NewSimplifiedAccessPolicyResponseModel

`func NewSimplifiedAccessPolicyResponseModel(allowAnonymous bool, allowHdxAccess bool, allowRdpAccess bool, connectNotViaNetScalerGatewayAllowed bool, connectViaNetScalerGatewayAllowed bool, includedSmartAccessFilterEnabled bool, includedUserFilterEnabled bool, excludedUserFilterEnabled bool, ) *SimplifiedAccessPolicyResponseModel`

NewSimplifiedAccessPolicyResponseModel instantiates a new SimplifiedAccessPolicyResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedAccessPolicyResponseModelWithDefaults

`func NewSimplifiedAccessPolicyResponseModelWithDefaults() *SimplifiedAccessPolicyResponseModel`

NewSimplifiedAccessPolicyResponseModelWithDefaults instantiates a new SimplifiedAccessPolicyResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAnonymous

`func (o *SimplifiedAccessPolicyResponseModel) GetAllowAnonymous() bool`

GetAllowAnonymous returns the AllowAnonymous field if non-nil, zero value otherwise.

### GetAllowAnonymousOk

`func (o *SimplifiedAccessPolicyResponseModel) GetAllowAnonymousOk() (*bool, bool)`

GetAllowAnonymousOk returns a tuple with the AllowAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymous

`func (o *SimplifiedAccessPolicyResponseModel) SetAllowAnonymous(v bool)`

SetAllowAnonymous sets AllowAnonymous field to given value.


### GetAllowHdxAccess

`func (o *SimplifiedAccessPolicyResponseModel) GetAllowHdxAccess() bool`

GetAllowHdxAccess returns the AllowHdxAccess field if non-nil, zero value otherwise.

### GetAllowHdxAccessOk

`func (o *SimplifiedAccessPolicyResponseModel) GetAllowHdxAccessOk() (*bool, bool)`

GetAllowHdxAccessOk returns a tuple with the AllowHdxAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHdxAccess

`func (o *SimplifiedAccessPolicyResponseModel) SetAllowHdxAccess(v bool)`

SetAllowHdxAccess sets AllowHdxAccess field to given value.


### GetAllowMachineRestart

`func (o *SimplifiedAccessPolicyResponseModel) GetAllowMachineRestart() bool`

GetAllowMachineRestart returns the AllowMachineRestart field if non-nil, zero value otherwise.

### GetAllowMachineRestartOk

`func (o *SimplifiedAccessPolicyResponseModel) GetAllowMachineRestartOk() (*bool, bool)`

GetAllowMachineRestartOk returns a tuple with the AllowMachineRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMachineRestart

`func (o *SimplifiedAccessPolicyResponseModel) SetAllowMachineRestart(v bool)`

SetAllowMachineRestart sets AllowMachineRestart field to given value.

### HasAllowMachineRestart

`func (o *SimplifiedAccessPolicyResponseModel) HasAllowMachineRestart() bool`

HasAllowMachineRestart returns a boolean if a field has been set.

### SetAllowMachineRestartNil

`func (o *SimplifiedAccessPolicyResponseModel) SetAllowMachineRestartNil(b bool)`

 SetAllowMachineRestartNil sets the value for AllowMachineRestart to be an explicit nil

### UnsetAllowMachineRestart
`func (o *SimplifiedAccessPolicyResponseModel) UnsetAllowMachineRestart()`

UnsetAllowMachineRestart ensures that no value is present for AllowMachineRestart, not even an explicit nil
### GetAllowRdpAccess

`func (o *SimplifiedAccessPolicyResponseModel) GetAllowRdpAccess() bool`

GetAllowRdpAccess returns the AllowRdpAccess field if non-nil, zero value otherwise.

### GetAllowRdpAccessOk

`func (o *SimplifiedAccessPolicyResponseModel) GetAllowRdpAccessOk() (*bool, bool)`

GetAllowRdpAccessOk returns a tuple with the AllowRdpAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRdpAccess

`func (o *SimplifiedAccessPolicyResponseModel) SetAllowRdpAccess(v bool)`

SetAllowRdpAccess sets AllowRdpAccess field to given value.


### GetConnectNotViaNetScalerGatewayAllowed

`func (o *SimplifiedAccessPolicyResponseModel) GetConnectNotViaNetScalerGatewayAllowed() bool`

GetConnectNotViaNetScalerGatewayAllowed returns the ConnectNotViaNetScalerGatewayAllowed field if non-nil, zero value otherwise.

### GetConnectNotViaNetScalerGatewayAllowedOk

`func (o *SimplifiedAccessPolicyResponseModel) GetConnectNotViaNetScalerGatewayAllowedOk() (*bool, bool)`

GetConnectNotViaNetScalerGatewayAllowedOk returns a tuple with the ConnectNotViaNetScalerGatewayAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectNotViaNetScalerGatewayAllowed

`func (o *SimplifiedAccessPolicyResponseModel) SetConnectNotViaNetScalerGatewayAllowed(v bool)`

SetConnectNotViaNetScalerGatewayAllowed sets ConnectNotViaNetScalerGatewayAllowed field to given value.


### GetConnectViaNetScalerGatewayAllowed

`func (o *SimplifiedAccessPolicyResponseModel) GetConnectViaNetScalerGatewayAllowed() bool`

GetConnectViaNetScalerGatewayAllowed returns the ConnectViaNetScalerGatewayAllowed field if non-nil, zero value otherwise.

### GetConnectViaNetScalerGatewayAllowedOk

`func (o *SimplifiedAccessPolicyResponseModel) GetConnectViaNetScalerGatewayAllowedOk() (*bool, bool)`

GetConnectViaNetScalerGatewayAllowedOk returns a tuple with the ConnectViaNetScalerGatewayAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectViaNetScalerGatewayAllowed

`func (o *SimplifiedAccessPolicyResponseModel) SetConnectViaNetScalerGatewayAllowed(v bool)`

SetConnectViaNetScalerGatewayAllowed sets ConnectViaNetScalerGatewayAllowed field to given value.


### GetIncludedSmartAccessFilterEnabled

`func (o *SimplifiedAccessPolicyResponseModel) GetIncludedSmartAccessFilterEnabled() bool`

GetIncludedSmartAccessFilterEnabled returns the IncludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterEnabledOk

`func (o *SimplifiedAccessPolicyResponseModel) GetIncludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetIncludedSmartAccessFilterEnabledOk returns a tuple with the IncludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterEnabled

`func (o *SimplifiedAccessPolicyResponseModel) SetIncludedSmartAccessFilterEnabled(v bool)`

SetIncludedSmartAccessFilterEnabled sets IncludedSmartAccessFilterEnabled field to given value.


### GetIncludedSmartAccessTags

`func (o *SimplifiedAccessPolicyResponseModel) GetIncludedSmartAccessTags() []SmartAccessTagResponseModel`

GetIncludedSmartAccessTags returns the IncludedSmartAccessTags field if non-nil, zero value otherwise.

### GetIncludedSmartAccessTagsOk

`func (o *SimplifiedAccessPolicyResponseModel) GetIncludedSmartAccessTagsOk() (*[]SmartAccessTagResponseModel, bool)`

GetIncludedSmartAccessTagsOk returns a tuple with the IncludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessTags

`func (o *SimplifiedAccessPolicyResponseModel) SetIncludedSmartAccessTags(v []SmartAccessTagResponseModel)`

SetIncludedSmartAccessTags sets IncludedSmartAccessTags field to given value.

### HasIncludedSmartAccessTags

`func (o *SimplifiedAccessPolicyResponseModel) HasIncludedSmartAccessTags() bool`

HasIncludedSmartAccessTags returns a boolean if a field has been set.

### SetIncludedSmartAccessTagsNil

`func (o *SimplifiedAccessPolicyResponseModel) SetIncludedSmartAccessTagsNil(b bool)`

 SetIncludedSmartAccessTagsNil sets the value for IncludedSmartAccessTags to be an explicit nil

### UnsetIncludedSmartAccessTags
`func (o *SimplifiedAccessPolicyResponseModel) UnsetIncludedSmartAccessTags()`

UnsetIncludedSmartAccessTags ensures that no value is present for IncludedSmartAccessTags, not even an explicit nil
### GetIncludedUserFilterEnabled

`func (o *SimplifiedAccessPolicyResponseModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *SimplifiedAccessPolicyResponseModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *SimplifiedAccessPolicyResponseModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.


### GetIncludedUsers

`func (o *SimplifiedAccessPolicyResponseModel) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *SimplifiedAccessPolicyResponseModel) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *SimplifiedAccessPolicyResponseModel) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *SimplifiedAccessPolicyResponseModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *SimplifiedAccessPolicyResponseModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *SimplifiedAccessPolicyResponseModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil
### GetExcludedUserFilterEnabled

`func (o *SimplifiedAccessPolicyResponseModel) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *SimplifiedAccessPolicyResponseModel) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *SimplifiedAccessPolicyResponseModel) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.


### GetExcludedUsers

`func (o *SimplifiedAccessPolicyResponseModel) GetExcludedUsers() []IdentityUserResponseModel`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *SimplifiedAccessPolicyResponseModel) GetExcludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *SimplifiedAccessPolicyResponseModel) SetExcludedUsers(v []IdentityUserResponseModel)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *SimplifiedAccessPolicyResponseModel) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.

### SetExcludedUsersNil

`func (o *SimplifiedAccessPolicyResponseModel) SetExcludedUsersNil(b bool)`

 SetExcludedUsersNil sets the value for ExcludedUsers to be an explicit nil

### UnsetExcludedUsers
`func (o *SimplifiedAccessPolicyResponseModel) UnsetExcludedUsers()`

UnsetExcludedUsers ensures that no value is present for ExcludedUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


