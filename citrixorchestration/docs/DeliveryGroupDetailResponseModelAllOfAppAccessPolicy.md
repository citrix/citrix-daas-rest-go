# DeliveryGroupDetailResponseModelAllOfAppAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether this policy is enabled. | [optional] 
**ExcludedUserFilterEnabled** | **bool** | Indicates whether the ExcludedUsers filter is enabled.  If the filter is disabled then any user entries in the filter are ignored when determining which applications are available for a user. | 
**ExcludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The excluded users filter for applications published on the delivery group; that is, the users and groups who are explicitly denied access to ALL applications published on the delivery group. | [optional] 
**IncludedUserFilterEnabled** | **bool** | Indicates whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access policy may access applications published on the delivery group. | 
**IncludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The included users filter for applications published on the delivery group; that is, the users and groups who are explicitly granted access to applications published on the delivery group. | [optional] 
**LeasingBehavior** | Pointer to [**LeasingBehavior**](LeasingBehavior.md) |  | [optional] 
**SessionReconnection** | Pointer to [**SessionReconnection**](SessionReconnection.md) |  | [optional] 

## Methods

### NewDeliveryGroupDetailResponseModelAllOfAppAccessPolicy

`func NewDeliveryGroupDetailResponseModelAllOfAppAccessPolicy(excludedUserFilterEnabled bool, includedUserFilterEnabled bool, ) *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy`

NewDeliveryGroupDetailResponseModelAllOfAppAccessPolicy instantiates a new DeliveryGroupDetailResponseModelAllOfAppAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupDetailResponseModelAllOfAppAccessPolicyWithDefaults

`func NewDeliveryGroupDetailResponseModelAllOfAppAccessPolicyWithDefaults() *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy`

NewDeliveryGroupDetailResponseModelAllOfAppAccessPolicyWithDefaults instantiates a new DeliveryGroupDetailResponseModelAllOfAppAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExcludedUserFilterEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.


### GetExcludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetExcludedUsers() []IdentityUserResponseModel`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetExcludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) SetExcludedUsers(v []IdentityUserResponseModel)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.

### GetIncludedUserFilterEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.


### GetIncludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### GetLeasingBehavior

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetLeasingBehavior() LeasingBehavior`

GetLeasingBehavior returns the LeasingBehavior field if non-nil, zero value otherwise.

### GetLeasingBehaviorOk

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetLeasingBehaviorOk() (*LeasingBehavior, bool)`

GetLeasingBehaviorOk returns a tuple with the LeasingBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasingBehavior

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) SetLeasingBehavior(v LeasingBehavior)`

SetLeasingBehavior sets LeasingBehavior field to given value.

### HasLeasingBehavior

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) HasLeasingBehavior() bool`

HasLeasingBehavior returns a boolean if a field has been set.

### GetSessionReconnection

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetSessionReconnection() SessionReconnection`

GetSessionReconnection returns the SessionReconnection field if non-nil, zero value otherwise.

### GetSessionReconnectionOk

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) GetSessionReconnectionOk() (*SessionReconnection, bool)`

GetSessionReconnectionOk returns a tuple with the SessionReconnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReconnection

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) SetSessionReconnection(v SessionReconnection)`

SetSessionReconnection sets SessionReconnection field to given value.

### HasSessionReconnection

`func (o *DeliveryGroupDetailResponseModelAllOfAppAccessPolicy) HasSessionReconnection() bool`

HasSessionReconnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


