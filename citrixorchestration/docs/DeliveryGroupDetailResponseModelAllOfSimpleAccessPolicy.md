# DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAnonymous** | **bool** | Whether anonymous connections are allowed to the delivery group. | 
**AllowHdxAccess** | **bool** | Whether HDX connections are allowed to the delivery group. CHANGE: was: public bool AllowsHdxAccess { get; set; } | 
**AllowMachineRestart** | Pointer to **bool** | Whether users are allowed to self-service restart machines. | [optional] 
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

### NewDeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy

`func NewDeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy(allowAnonymous bool, allowHdxAccess bool, allowRdpAccess bool, connectNotViaNetScalerGatewayAllowed bool, connectViaNetScalerGatewayAllowed bool, includedSmartAccessFilterEnabled bool, includedUserFilterEnabled bool, excludedUserFilterEnabled bool, ) *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy`

NewDeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy instantiates a new DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupDetailResponseModelAllOfSimpleAccessPolicyWithDefaults

`func NewDeliveryGroupDetailResponseModelAllOfSimpleAccessPolicyWithDefaults() *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy`

NewDeliveryGroupDetailResponseModelAllOfSimpleAccessPolicyWithDefaults instantiates a new DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAnonymous

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetAllowAnonymous() bool`

GetAllowAnonymous returns the AllowAnonymous field if non-nil, zero value otherwise.

### GetAllowAnonymousOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetAllowAnonymousOk() (*bool, bool)`

GetAllowAnonymousOk returns a tuple with the AllowAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymous

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetAllowAnonymous(v bool)`

SetAllowAnonymous sets AllowAnonymous field to given value.


### GetAllowHdxAccess

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetAllowHdxAccess() bool`

GetAllowHdxAccess returns the AllowHdxAccess field if non-nil, zero value otherwise.

### GetAllowHdxAccessOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetAllowHdxAccessOk() (*bool, bool)`

GetAllowHdxAccessOk returns a tuple with the AllowHdxAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHdxAccess

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetAllowHdxAccess(v bool)`

SetAllowHdxAccess sets AllowHdxAccess field to given value.


### GetAllowMachineRestart

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetAllowMachineRestart() bool`

GetAllowMachineRestart returns the AllowMachineRestart field if non-nil, zero value otherwise.

### GetAllowMachineRestartOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetAllowMachineRestartOk() (*bool, bool)`

GetAllowMachineRestartOk returns a tuple with the AllowMachineRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMachineRestart

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetAllowMachineRestart(v bool)`

SetAllowMachineRestart sets AllowMachineRestart field to given value.

### HasAllowMachineRestart

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) HasAllowMachineRestart() bool`

HasAllowMachineRestart returns a boolean if a field has been set.

### GetAllowRdpAccess

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetAllowRdpAccess() bool`

GetAllowRdpAccess returns the AllowRdpAccess field if non-nil, zero value otherwise.

### GetAllowRdpAccessOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetAllowRdpAccessOk() (*bool, bool)`

GetAllowRdpAccessOk returns a tuple with the AllowRdpAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRdpAccess

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetAllowRdpAccess(v bool)`

SetAllowRdpAccess sets AllowRdpAccess field to given value.


### GetConnectNotViaNetScalerGatewayAllowed

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetConnectNotViaNetScalerGatewayAllowed() bool`

GetConnectNotViaNetScalerGatewayAllowed returns the ConnectNotViaNetScalerGatewayAllowed field if non-nil, zero value otherwise.

### GetConnectNotViaNetScalerGatewayAllowedOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetConnectNotViaNetScalerGatewayAllowedOk() (*bool, bool)`

GetConnectNotViaNetScalerGatewayAllowedOk returns a tuple with the ConnectNotViaNetScalerGatewayAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectNotViaNetScalerGatewayAllowed

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetConnectNotViaNetScalerGatewayAllowed(v bool)`

SetConnectNotViaNetScalerGatewayAllowed sets ConnectNotViaNetScalerGatewayAllowed field to given value.


### GetConnectViaNetScalerGatewayAllowed

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetConnectViaNetScalerGatewayAllowed() bool`

GetConnectViaNetScalerGatewayAllowed returns the ConnectViaNetScalerGatewayAllowed field if non-nil, zero value otherwise.

### GetConnectViaNetScalerGatewayAllowedOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetConnectViaNetScalerGatewayAllowedOk() (*bool, bool)`

GetConnectViaNetScalerGatewayAllowedOk returns a tuple with the ConnectViaNetScalerGatewayAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectViaNetScalerGatewayAllowed

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetConnectViaNetScalerGatewayAllowed(v bool)`

SetConnectViaNetScalerGatewayAllowed sets ConnectViaNetScalerGatewayAllowed field to given value.


### GetIncludedSmartAccessFilterEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetIncludedSmartAccessFilterEnabled() bool`

GetIncludedSmartAccessFilterEnabled returns the IncludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterEnabledOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetIncludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetIncludedSmartAccessFilterEnabledOk returns a tuple with the IncludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetIncludedSmartAccessFilterEnabled(v bool)`

SetIncludedSmartAccessFilterEnabled sets IncludedSmartAccessFilterEnabled field to given value.


### GetIncludedSmartAccessTags

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetIncludedSmartAccessTags() []SmartAccessTagResponseModel`

GetIncludedSmartAccessTags returns the IncludedSmartAccessTags field if non-nil, zero value otherwise.

### GetIncludedSmartAccessTagsOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetIncludedSmartAccessTagsOk() (*[]SmartAccessTagResponseModel, bool)`

GetIncludedSmartAccessTagsOk returns a tuple with the IncludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessTags

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetIncludedSmartAccessTags(v []SmartAccessTagResponseModel)`

SetIncludedSmartAccessTags sets IncludedSmartAccessTags field to given value.

### HasIncludedSmartAccessTags

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) HasIncludedSmartAccessTags() bool`

HasIncludedSmartAccessTags returns a boolean if a field has been set.

### GetIncludedUserFilterEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.


### GetIncludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### GetExcludedUserFilterEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.


### GetExcludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetExcludedUsers() []IdentityUserResponseModel`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) GetExcludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) SetExcludedUsers(v []IdentityUserResponseModel)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfSimpleAccessPolicy) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


