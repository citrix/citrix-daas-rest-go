# AppAccessPolicyRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether the app access policy is enabled. | [optional] 
**ExcludedUserFilterEnabled** | Pointer to **bool** | Specifies whether the ExcludedUsers filter is enabled.  If the filter is disabled then any user entries in the filter are ignored when determining which applications are available for a user. | [optional] 
**ExcludedUsers** | Pointer to **[]string** | Specifies the excluded users filter for applications published on the delivery group; that is, the users and groups who are explicitly denied access to ALL applications published on the delivery group. | [optional] 
**IncludedUserFilterEnabled** | Pointer to **bool** | Specifies whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access policy may access applications published on the delivery group. | [optional] 
**IncludedUsers** | Pointer to **[]string** | Specifies the included users filter for applications published on the delivery group; that is, the users and groups who are explicitly granted access to applications published on the delivery group. | [optional] 
**LeasingBehavior** | Pointer to [**LeasingBehavior**](LeasingBehavior.md) |  | [optional] 
**SessionReconnection** | Pointer to [**SessionReconnection**](SessionReconnection.md) |  | [optional] 

## Methods

### NewAppAccessPolicyRequestModel

`func NewAppAccessPolicyRequestModel() *AppAccessPolicyRequestModel`

NewAppAccessPolicyRequestModel instantiates a new AppAccessPolicyRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccessPolicyRequestModelWithDefaults

`func NewAppAccessPolicyRequestModelWithDefaults() *AppAccessPolicyRequestModel`

NewAppAccessPolicyRequestModelWithDefaults instantiates a new AppAccessPolicyRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AppAccessPolicyRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppAccessPolicyRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppAccessPolicyRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AppAccessPolicyRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExcludedUserFilterEnabled

`func (o *AppAccessPolicyRequestModel) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *AppAccessPolicyRequestModel) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *AppAccessPolicyRequestModel) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.

### HasExcludedUserFilterEnabled

`func (o *AppAccessPolicyRequestModel) HasExcludedUserFilterEnabled() bool`

HasExcludedUserFilterEnabled returns a boolean if a field has been set.

### GetExcludedUsers

`func (o *AppAccessPolicyRequestModel) GetExcludedUsers() []string`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *AppAccessPolicyRequestModel) GetExcludedUsersOk() (*[]string, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *AppAccessPolicyRequestModel) SetExcludedUsers(v []string)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *AppAccessPolicyRequestModel) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.

### GetIncludedUserFilterEnabled

`func (o *AppAccessPolicyRequestModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *AppAccessPolicyRequestModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *AppAccessPolicyRequestModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *AppAccessPolicyRequestModel) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### GetIncludedUsers

`func (o *AppAccessPolicyRequestModel) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *AppAccessPolicyRequestModel) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *AppAccessPolicyRequestModel) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *AppAccessPolicyRequestModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### GetLeasingBehavior

`func (o *AppAccessPolicyRequestModel) GetLeasingBehavior() LeasingBehavior`

GetLeasingBehavior returns the LeasingBehavior field if non-nil, zero value otherwise.

### GetLeasingBehaviorOk

`func (o *AppAccessPolicyRequestModel) GetLeasingBehaviorOk() (*LeasingBehavior, bool)`

GetLeasingBehaviorOk returns a tuple with the LeasingBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasingBehavior

`func (o *AppAccessPolicyRequestModel) SetLeasingBehavior(v LeasingBehavior)`

SetLeasingBehavior sets LeasingBehavior field to given value.

### HasLeasingBehavior

`func (o *AppAccessPolicyRequestModel) HasLeasingBehavior() bool`

HasLeasingBehavior returns a boolean if a field has been set.

### GetSessionReconnection

`func (o *AppAccessPolicyRequestModel) GetSessionReconnection() SessionReconnection`

GetSessionReconnection returns the SessionReconnection field if non-nil, zero value otherwise.

### GetSessionReconnectionOk

`func (o *AppAccessPolicyRequestModel) GetSessionReconnectionOk() (*SessionReconnection, bool)`

GetSessionReconnectionOk returns a tuple with the SessionReconnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReconnection

`func (o *AppAccessPolicyRequestModel) SetSessionReconnection(v SessionReconnection)`

SetSessionReconnection sets SessionReconnection field to given value.

### HasSessionReconnection

`func (o *AppAccessPolicyRequestModel) HasSessionReconnection() bool`

HasSessionReconnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


