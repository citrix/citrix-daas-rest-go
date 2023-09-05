# CreateDeliveryGroupRequestModelSimpleAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAnonymous** | Pointer to **bool** | Whether to allow anonymous connections to the delivery group. | [optional] 
**AllowHdxAccess** | Pointer to **bool** | Whether to allow connections using HDX. Was: AllowsHdxAccess | [optional] [default to true]
**AllowMachineRestart** | Pointer to **bool** | Whether to allow users to self-service restart machines. | [optional] [default to true]
**AllowRdpAccess** | Pointer to **bool** | Whether to allow connections using RDP. Was: AllowsRdpAccess | [optional] [default to true]
**ConnectNotViaNetScalerGatewayAllowed** | Pointer to **bool** | Whether to allow connections that do not come through NetScaler Gateway. | [optional] [default to true]
**ConnectViaNetScalerGatewayAllowed** | Pointer to **bool** | Whether to allow connections that come through NetScaler Gateway. | [optional] [default to true]
**IncludedSmartAccessFilterEnabled** | Pointer to **bool** | Specifies whether the IncludedSmartAccessTags filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] [default to true]
**IncludedSmartAccessTags** | Pointer to [**[]SmartAccessTagRequestModel**](SmartAccessTagRequestModel.md) | Specifies the SmartAccess tags which grant access to the delivery group, if any occur in those provided by NetScaler Gateway with the user&#39;s connection. | [optional] 
**IncludedUserFilterEnabled** | Pointer to **bool** | Specifies whether the IncludedUsers filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] [default to false]
**IncludedUsers** | Pointer to **[]string** | Specifies the users and groups who are granted access to the delivery group. | [optional] 
**ExcludedUserFilterEnabled** | Pointer to **bool** | Specifies whether the ExcludedUsers filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] [default to false]
**ExcludedUsers** | Pointer to **[]string** | Specifies the users and groups who are denied access to the delivery group. | [optional] 

## Methods

### NewCreateDeliveryGroupRequestModelSimpleAccessPolicy

`func NewCreateDeliveryGroupRequestModelSimpleAccessPolicy() *CreateDeliveryGroupRequestModelSimpleAccessPolicy`

NewCreateDeliveryGroupRequestModelSimpleAccessPolicy instantiates a new CreateDeliveryGroupRequestModelSimpleAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeliveryGroupRequestModelSimpleAccessPolicyWithDefaults

`func NewCreateDeliveryGroupRequestModelSimpleAccessPolicyWithDefaults() *CreateDeliveryGroupRequestModelSimpleAccessPolicy`

NewCreateDeliveryGroupRequestModelSimpleAccessPolicyWithDefaults instantiates a new CreateDeliveryGroupRequestModelSimpleAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAnonymous

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetAllowAnonymous() bool`

GetAllowAnonymous returns the AllowAnonymous field if non-nil, zero value otherwise.

### GetAllowAnonymousOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetAllowAnonymousOk() (*bool, bool)`

GetAllowAnonymousOk returns a tuple with the AllowAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymous

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetAllowAnonymous(v bool)`

SetAllowAnonymous sets AllowAnonymous field to given value.

### HasAllowAnonymous

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasAllowAnonymous() bool`

HasAllowAnonymous returns a boolean if a field has been set.

### GetAllowHdxAccess

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetAllowHdxAccess() bool`

GetAllowHdxAccess returns the AllowHdxAccess field if non-nil, zero value otherwise.

### GetAllowHdxAccessOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetAllowHdxAccessOk() (*bool, bool)`

GetAllowHdxAccessOk returns a tuple with the AllowHdxAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHdxAccess

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetAllowHdxAccess(v bool)`

SetAllowHdxAccess sets AllowHdxAccess field to given value.

### HasAllowHdxAccess

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasAllowHdxAccess() bool`

HasAllowHdxAccess returns a boolean if a field has been set.

### GetAllowMachineRestart

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetAllowMachineRestart() bool`

GetAllowMachineRestart returns the AllowMachineRestart field if non-nil, zero value otherwise.

### GetAllowMachineRestartOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetAllowMachineRestartOk() (*bool, bool)`

GetAllowMachineRestartOk returns a tuple with the AllowMachineRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMachineRestart

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetAllowMachineRestart(v bool)`

SetAllowMachineRestart sets AllowMachineRestart field to given value.

### HasAllowMachineRestart

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasAllowMachineRestart() bool`

HasAllowMachineRestart returns a boolean if a field has been set.

### GetAllowRdpAccess

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetAllowRdpAccess() bool`

GetAllowRdpAccess returns the AllowRdpAccess field if non-nil, zero value otherwise.

### GetAllowRdpAccessOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetAllowRdpAccessOk() (*bool, bool)`

GetAllowRdpAccessOk returns a tuple with the AllowRdpAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRdpAccess

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetAllowRdpAccess(v bool)`

SetAllowRdpAccess sets AllowRdpAccess field to given value.

### HasAllowRdpAccess

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasAllowRdpAccess() bool`

HasAllowRdpAccess returns a boolean if a field has been set.

### GetConnectNotViaNetScalerGatewayAllowed

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetConnectNotViaNetScalerGatewayAllowed() bool`

GetConnectNotViaNetScalerGatewayAllowed returns the ConnectNotViaNetScalerGatewayAllowed field if non-nil, zero value otherwise.

### GetConnectNotViaNetScalerGatewayAllowedOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetConnectNotViaNetScalerGatewayAllowedOk() (*bool, bool)`

GetConnectNotViaNetScalerGatewayAllowedOk returns a tuple with the ConnectNotViaNetScalerGatewayAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectNotViaNetScalerGatewayAllowed

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetConnectNotViaNetScalerGatewayAllowed(v bool)`

SetConnectNotViaNetScalerGatewayAllowed sets ConnectNotViaNetScalerGatewayAllowed field to given value.

### HasConnectNotViaNetScalerGatewayAllowed

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasConnectNotViaNetScalerGatewayAllowed() bool`

HasConnectNotViaNetScalerGatewayAllowed returns a boolean if a field has been set.

### GetConnectViaNetScalerGatewayAllowed

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetConnectViaNetScalerGatewayAllowed() bool`

GetConnectViaNetScalerGatewayAllowed returns the ConnectViaNetScalerGatewayAllowed field if non-nil, zero value otherwise.

### GetConnectViaNetScalerGatewayAllowedOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetConnectViaNetScalerGatewayAllowedOk() (*bool, bool)`

GetConnectViaNetScalerGatewayAllowedOk returns a tuple with the ConnectViaNetScalerGatewayAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectViaNetScalerGatewayAllowed

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetConnectViaNetScalerGatewayAllowed(v bool)`

SetConnectViaNetScalerGatewayAllowed sets ConnectViaNetScalerGatewayAllowed field to given value.

### HasConnectViaNetScalerGatewayAllowed

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasConnectViaNetScalerGatewayAllowed() bool`

HasConnectViaNetScalerGatewayAllowed returns a boolean if a field has been set.

### GetIncludedSmartAccessFilterEnabled

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetIncludedSmartAccessFilterEnabled() bool`

GetIncludedSmartAccessFilterEnabled returns the IncludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterEnabledOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetIncludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetIncludedSmartAccessFilterEnabledOk returns a tuple with the IncludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterEnabled

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetIncludedSmartAccessFilterEnabled(v bool)`

SetIncludedSmartAccessFilterEnabled sets IncludedSmartAccessFilterEnabled field to given value.

### HasIncludedSmartAccessFilterEnabled

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasIncludedSmartAccessFilterEnabled() bool`

HasIncludedSmartAccessFilterEnabled returns a boolean if a field has been set.

### GetIncludedSmartAccessTags

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetIncludedSmartAccessTags() []SmartAccessTagRequestModel`

GetIncludedSmartAccessTags returns the IncludedSmartAccessTags field if non-nil, zero value otherwise.

### GetIncludedSmartAccessTagsOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetIncludedSmartAccessTagsOk() (*[]SmartAccessTagRequestModel, bool)`

GetIncludedSmartAccessTagsOk returns a tuple with the IncludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessTags

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetIncludedSmartAccessTags(v []SmartAccessTagRequestModel)`

SetIncludedSmartAccessTags sets IncludedSmartAccessTags field to given value.

### HasIncludedSmartAccessTags

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasIncludedSmartAccessTags() bool`

HasIncludedSmartAccessTags returns a boolean if a field has been set.

### GetIncludedUserFilterEnabled

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### GetIncludedUsers

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### GetExcludedUserFilterEnabled

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.

### HasExcludedUserFilterEnabled

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasExcludedUserFilterEnabled() bool`

HasExcludedUserFilterEnabled returns a boolean if a field has been set.

### GetExcludedUsers

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetExcludedUsers() []string`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) GetExcludedUsersOk() (*[]string, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) SetExcludedUsers(v []string)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *CreateDeliveryGroupRequestModelSimpleAccessPolicy) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


