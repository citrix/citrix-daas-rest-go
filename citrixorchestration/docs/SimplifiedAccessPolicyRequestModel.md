# SimplifiedAccessPolicyRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAnonymous** | Pointer to **NullableBool** | Whether to allow anonymous connections to the delivery group. | [optional] 
**AllowHdxAccess** | Pointer to **NullableBool** | Whether to allow connections using HDX. Was: AllowsHdxAccess | [optional] [default to true]
**AllowMachineRestart** | Pointer to **NullableBool** | Whether to allow users to self-service restart machines. | [optional] [default to true]
**AllowRdpAccess** | Pointer to **NullableBool** | Whether to allow connections using RDP. Was: AllowsRdpAccess | [optional] [default to true]
**ConnectNotViaNetScalerGatewayAllowed** | Pointer to **NullableBool** | Whether to allow connections that do not come through NetScaler Gateway. | [optional] [default to true]
**ConnectViaNetScalerGatewayAllowed** | Pointer to **NullableBool** | Whether to allow connections that come through NetScaler Gateway. | [optional] [default to true]
**IncludedSmartAccessFilterEnabled** | Pointer to **NullableBool** | Specifies whether the IncludedSmartAccessTags filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] [default to true]
**IncludedSmartAccessTags** | Pointer to [**[]SmartAccessTagRequestModel**](SmartAccessTagRequestModel.md) | Specifies the SmartAccess tags which grant access to the delivery group, if any occur in those provided by NetScaler Gateway with the user&#39;s connection. | [optional] 
**IncludedUserFilterEnabled** | Pointer to **NullableBool** | Specifies whether the IncludedUsers filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] [default to false]
**IncludedUsers** | Pointer to **[]string** | Specifies the users and groups who are granted access to the delivery group. | [optional] 
**ExcludedUserFilterEnabled** | Pointer to **NullableBool** | Specifies whether the ExcludedUsers filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] [default to false]
**ExcludedUsers** | Pointer to **[]string** | Specifies the users and groups who are denied access to the delivery group. | [optional] 

## Methods

### NewSimplifiedAccessPolicyRequestModel

`func NewSimplifiedAccessPolicyRequestModel() *SimplifiedAccessPolicyRequestModel`

NewSimplifiedAccessPolicyRequestModel instantiates a new SimplifiedAccessPolicyRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedAccessPolicyRequestModelWithDefaults

`func NewSimplifiedAccessPolicyRequestModelWithDefaults() *SimplifiedAccessPolicyRequestModel`

NewSimplifiedAccessPolicyRequestModelWithDefaults instantiates a new SimplifiedAccessPolicyRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAnonymous

`func (o *SimplifiedAccessPolicyRequestModel) GetAllowAnonymous() bool`

GetAllowAnonymous returns the AllowAnonymous field if non-nil, zero value otherwise.

### GetAllowAnonymousOk

`func (o *SimplifiedAccessPolicyRequestModel) GetAllowAnonymousOk() (*bool, bool)`

GetAllowAnonymousOk returns a tuple with the AllowAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymous

`func (o *SimplifiedAccessPolicyRequestModel) SetAllowAnonymous(v bool)`

SetAllowAnonymous sets AllowAnonymous field to given value.

### HasAllowAnonymous

`func (o *SimplifiedAccessPolicyRequestModel) HasAllowAnonymous() bool`

HasAllowAnonymous returns a boolean if a field has been set.

### SetAllowAnonymousNil

`func (o *SimplifiedAccessPolicyRequestModel) SetAllowAnonymousNil(b bool)`

 SetAllowAnonymousNil sets the value for AllowAnonymous to be an explicit nil

### UnsetAllowAnonymous
`func (o *SimplifiedAccessPolicyRequestModel) UnsetAllowAnonymous()`

UnsetAllowAnonymous ensures that no value is present for AllowAnonymous, not even an explicit nil
### GetAllowHdxAccess

`func (o *SimplifiedAccessPolicyRequestModel) GetAllowHdxAccess() bool`

GetAllowHdxAccess returns the AllowHdxAccess field if non-nil, zero value otherwise.

### GetAllowHdxAccessOk

`func (o *SimplifiedAccessPolicyRequestModel) GetAllowHdxAccessOk() (*bool, bool)`

GetAllowHdxAccessOk returns a tuple with the AllowHdxAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHdxAccess

`func (o *SimplifiedAccessPolicyRequestModel) SetAllowHdxAccess(v bool)`

SetAllowHdxAccess sets AllowHdxAccess field to given value.

### HasAllowHdxAccess

`func (o *SimplifiedAccessPolicyRequestModel) HasAllowHdxAccess() bool`

HasAllowHdxAccess returns a boolean if a field has been set.

### SetAllowHdxAccessNil

`func (o *SimplifiedAccessPolicyRequestModel) SetAllowHdxAccessNil(b bool)`

 SetAllowHdxAccessNil sets the value for AllowHdxAccess to be an explicit nil

### UnsetAllowHdxAccess
`func (o *SimplifiedAccessPolicyRequestModel) UnsetAllowHdxAccess()`

UnsetAllowHdxAccess ensures that no value is present for AllowHdxAccess, not even an explicit nil
### GetAllowMachineRestart

`func (o *SimplifiedAccessPolicyRequestModel) GetAllowMachineRestart() bool`

GetAllowMachineRestart returns the AllowMachineRestart field if non-nil, zero value otherwise.

### GetAllowMachineRestartOk

`func (o *SimplifiedAccessPolicyRequestModel) GetAllowMachineRestartOk() (*bool, bool)`

GetAllowMachineRestartOk returns a tuple with the AllowMachineRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMachineRestart

`func (o *SimplifiedAccessPolicyRequestModel) SetAllowMachineRestart(v bool)`

SetAllowMachineRestart sets AllowMachineRestart field to given value.

### HasAllowMachineRestart

`func (o *SimplifiedAccessPolicyRequestModel) HasAllowMachineRestart() bool`

HasAllowMachineRestart returns a boolean if a field has been set.

### SetAllowMachineRestartNil

`func (o *SimplifiedAccessPolicyRequestModel) SetAllowMachineRestartNil(b bool)`

 SetAllowMachineRestartNil sets the value for AllowMachineRestart to be an explicit nil

### UnsetAllowMachineRestart
`func (o *SimplifiedAccessPolicyRequestModel) UnsetAllowMachineRestart()`

UnsetAllowMachineRestart ensures that no value is present for AllowMachineRestart, not even an explicit nil
### GetAllowRdpAccess

`func (o *SimplifiedAccessPolicyRequestModel) GetAllowRdpAccess() bool`

GetAllowRdpAccess returns the AllowRdpAccess field if non-nil, zero value otherwise.

### GetAllowRdpAccessOk

`func (o *SimplifiedAccessPolicyRequestModel) GetAllowRdpAccessOk() (*bool, bool)`

GetAllowRdpAccessOk returns a tuple with the AllowRdpAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRdpAccess

`func (o *SimplifiedAccessPolicyRequestModel) SetAllowRdpAccess(v bool)`

SetAllowRdpAccess sets AllowRdpAccess field to given value.

### HasAllowRdpAccess

`func (o *SimplifiedAccessPolicyRequestModel) HasAllowRdpAccess() bool`

HasAllowRdpAccess returns a boolean if a field has been set.

### SetAllowRdpAccessNil

`func (o *SimplifiedAccessPolicyRequestModel) SetAllowRdpAccessNil(b bool)`

 SetAllowRdpAccessNil sets the value for AllowRdpAccess to be an explicit nil

### UnsetAllowRdpAccess
`func (o *SimplifiedAccessPolicyRequestModel) UnsetAllowRdpAccess()`

UnsetAllowRdpAccess ensures that no value is present for AllowRdpAccess, not even an explicit nil
### GetConnectNotViaNetScalerGatewayAllowed

`func (o *SimplifiedAccessPolicyRequestModel) GetConnectNotViaNetScalerGatewayAllowed() bool`

GetConnectNotViaNetScalerGatewayAllowed returns the ConnectNotViaNetScalerGatewayAllowed field if non-nil, zero value otherwise.

### GetConnectNotViaNetScalerGatewayAllowedOk

`func (o *SimplifiedAccessPolicyRequestModel) GetConnectNotViaNetScalerGatewayAllowedOk() (*bool, bool)`

GetConnectNotViaNetScalerGatewayAllowedOk returns a tuple with the ConnectNotViaNetScalerGatewayAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectNotViaNetScalerGatewayAllowed

`func (o *SimplifiedAccessPolicyRequestModel) SetConnectNotViaNetScalerGatewayAllowed(v bool)`

SetConnectNotViaNetScalerGatewayAllowed sets ConnectNotViaNetScalerGatewayAllowed field to given value.

### HasConnectNotViaNetScalerGatewayAllowed

`func (o *SimplifiedAccessPolicyRequestModel) HasConnectNotViaNetScalerGatewayAllowed() bool`

HasConnectNotViaNetScalerGatewayAllowed returns a boolean if a field has been set.

### SetConnectNotViaNetScalerGatewayAllowedNil

`func (o *SimplifiedAccessPolicyRequestModel) SetConnectNotViaNetScalerGatewayAllowedNil(b bool)`

 SetConnectNotViaNetScalerGatewayAllowedNil sets the value for ConnectNotViaNetScalerGatewayAllowed to be an explicit nil

### UnsetConnectNotViaNetScalerGatewayAllowed
`func (o *SimplifiedAccessPolicyRequestModel) UnsetConnectNotViaNetScalerGatewayAllowed()`

UnsetConnectNotViaNetScalerGatewayAllowed ensures that no value is present for ConnectNotViaNetScalerGatewayAllowed, not even an explicit nil
### GetConnectViaNetScalerGatewayAllowed

`func (o *SimplifiedAccessPolicyRequestModel) GetConnectViaNetScalerGatewayAllowed() bool`

GetConnectViaNetScalerGatewayAllowed returns the ConnectViaNetScalerGatewayAllowed field if non-nil, zero value otherwise.

### GetConnectViaNetScalerGatewayAllowedOk

`func (o *SimplifiedAccessPolicyRequestModel) GetConnectViaNetScalerGatewayAllowedOk() (*bool, bool)`

GetConnectViaNetScalerGatewayAllowedOk returns a tuple with the ConnectViaNetScalerGatewayAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectViaNetScalerGatewayAllowed

`func (o *SimplifiedAccessPolicyRequestModel) SetConnectViaNetScalerGatewayAllowed(v bool)`

SetConnectViaNetScalerGatewayAllowed sets ConnectViaNetScalerGatewayAllowed field to given value.

### HasConnectViaNetScalerGatewayAllowed

`func (o *SimplifiedAccessPolicyRequestModel) HasConnectViaNetScalerGatewayAllowed() bool`

HasConnectViaNetScalerGatewayAllowed returns a boolean if a field has been set.

### SetConnectViaNetScalerGatewayAllowedNil

`func (o *SimplifiedAccessPolicyRequestModel) SetConnectViaNetScalerGatewayAllowedNil(b bool)`

 SetConnectViaNetScalerGatewayAllowedNil sets the value for ConnectViaNetScalerGatewayAllowed to be an explicit nil

### UnsetConnectViaNetScalerGatewayAllowed
`func (o *SimplifiedAccessPolicyRequestModel) UnsetConnectViaNetScalerGatewayAllowed()`

UnsetConnectViaNetScalerGatewayAllowed ensures that no value is present for ConnectViaNetScalerGatewayAllowed, not even an explicit nil
### GetIncludedSmartAccessFilterEnabled

`func (o *SimplifiedAccessPolicyRequestModel) GetIncludedSmartAccessFilterEnabled() bool`

GetIncludedSmartAccessFilterEnabled returns the IncludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterEnabledOk

`func (o *SimplifiedAccessPolicyRequestModel) GetIncludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetIncludedSmartAccessFilterEnabledOk returns a tuple with the IncludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterEnabled

`func (o *SimplifiedAccessPolicyRequestModel) SetIncludedSmartAccessFilterEnabled(v bool)`

SetIncludedSmartAccessFilterEnabled sets IncludedSmartAccessFilterEnabled field to given value.

### HasIncludedSmartAccessFilterEnabled

`func (o *SimplifiedAccessPolicyRequestModel) HasIncludedSmartAccessFilterEnabled() bool`

HasIncludedSmartAccessFilterEnabled returns a boolean if a field has been set.

### SetIncludedSmartAccessFilterEnabledNil

`func (o *SimplifiedAccessPolicyRequestModel) SetIncludedSmartAccessFilterEnabledNil(b bool)`

 SetIncludedSmartAccessFilterEnabledNil sets the value for IncludedSmartAccessFilterEnabled to be an explicit nil

### UnsetIncludedSmartAccessFilterEnabled
`func (o *SimplifiedAccessPolicyRequestModel) UnsetIncludedSmartAccessFilterEnabled()`

UnsetIncludedSmartAccessFilterEnabled ensures that no value is present for IncludedSmartAccessFilterEnabled, not even an explicit nil
### GetIncludedSmartAccessTags

`func (o *SimplifiedAccessPolicyRequestModel) GetIncludedSmartAccessTags() []SmartAccessTagRequestModel`

GetIncludedSmartAccessTags returns the IncludedSmartAccessTags field if non-nil, zero value otherwise.

### GetIncludedSmartAccessTagsOk

`func (o *SimplifiedAccessPolicyRequestModel) GetIncludedSmartAccessTagsOk() (*[]SmartAccessTagRequestModel, bool)`

GetIncludedSmartAccessTagsOk returns a tuple with the IncludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessTags

`func (o *SimplifiedAccessPolicyRequestModel) SetIncludedSmartAccessTags(v []SmartAccessTagRequestModel)`

SetIncludedSmartAccessTags sets IncludedSmartAccessTags field to given value.

### HasIncludedSmartAccessTags

`func (o *SimplifiedAccessPolicyRequestModel) HasIncludedSmartAccessTags() bool`

HasIncludedSmartAccessTags returns a boolean if a field has been set.

### SetIncludedSmartAccessTagsNil

`func (o *SimplifiedAccessPolicyRequestModel) SetIncludedSmartAccessTagsNil(b bool)`

 SetIncludedSmartAccessTagsNil sets the value for IncludedSmartAccessTags to be an explicit nil

### UnsetIncludedSmartAccessTags
`func (o *SimplifiedAccessPolicyRequestModel) UnsetIncludedSmartAccessTags()`

UnsetIncludedSmartAccessTags ensures that no value is present for IncludedSmartAccessTags, not even an explicit nil
### GetIncludedUserFilterEnabled

`func (o *SimplifiedAccessPolicyRequestModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *SimplifiedAccessPolicyRequestModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *SimplifiedAccessPolicyRequestModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *SimplifiedAccessPolicyRequestModel) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### SetIncludedUserFilterEnabledNil

`func (o *SimplifiedAccessPolicyRequestModel) SetIncludedUserFilterEnabledNil(b bool)`

 SetIncludedUserFilterEnabledNil sets the value for IncludedUserFilterEnabled to be an explicit nil

### UnsetIncludedUserFilterEnabled
`func (o *SimplifiedAccessPolicyRequestModel) UnsetIncludedUserFilterEnabled()`

UnsetIncludedUserFilterEnabled ensures that no value is present for IncludedUserFilterEnabled, not even an explicit nil
### GetIncludedUsers

`func (o *SimplifiedAccessPolicyRequestModel) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *SimplifiedAccessPolicyRequestModel) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *SimplifiedAccessPolicyRequestModel) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *SimplifiedAccessPolicyRequestModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *SimplifiedAccessPolicyRequestModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *SimplifiedAccessPolicyRequestModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil
### GetExcludedUserFilterEnabled

`func (o *SimplifiedAccessPolicyRequestModel) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *SimplifiedAccessPolicyRequestModel) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *SimplifiedAccessPolicyRequestModel) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.

### HasExcludedUserFilterEnabled

`func (o *SimplifiedAccessPolicyRequestModel) HasExcludedUserFilterEnabled() bool`

HasExcludedUserFilterEnabled returns a boolean if a field has been set.

### SetExcludedUserFilterEnabledNil

`func (o *SimplifiedAccessPolicyRequestModel) SetExcludedUserFilterEnabledNil(b bool)`

 SetExcludedUserFilterEnabledNil sets the value for ExcludedUserFilterEnabled to be an explicit nil

### UnsetExcludedUserFilterEnabled
`func (o *SimplifiedAccessPolicyRequestModel) UnsetExcludedUserFilterEnabled()`

UnsetExcludedUserFilterEnabled ensures that no value is present for ExcludedUserFilterEnabled, not even an explicit nil
### GetExcludedUsers

`func (o *SimplifiedAccessPolicyRequestModel) GetExcludedUsers() []string`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *SimplifiedAccessPolicyRequestModel) GetExcludedUsersOk() (*[]string, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *SimplifiedAccessPolicyRequestModel) SetExcludedUsers(v []string)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *SimplifiedAccessPolicyRequestModel) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.

### SetExcludedUsersNil

`func (o *SimplifiedAccessPolicyRequestModel) SetExcludedUsersNil(b bool)`

 SetExcludedUsersNil sets the value for ExcludedUsers to be an explicit nil

### UnsetExcludedUsers
`func (o *SimplifiedAccessPolicyRequestModel) UnsetExcludedUsers()`

UnsetExcludedUsers ensures that no value is present for ExcludedUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


