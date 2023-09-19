# AppAccessPolicyResponseModel

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

### NewAppAccessPolicyResponseModel

`func NewAppAccessPolicyResponseModel(excludedUserFilterEnabled bool, includedUserFilterEnabled bool, ) *AppAccessPolicyResponseModel`

NewAppAccessPolicyResponseModel instantiates a new AppAccessPolicyResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccessPolicyResponseModelWithDefaults

`func NewAppAccessPolicyResponseModelWithDefaults() *AppAccessPolicyResponseModel`

NewAppAccessPolicyResponseModelWithDefaults instantiates a new AppAccessPolicyResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AppAccessPolicyResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppAccessPolicyResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppAccessPolicyResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AppAccessPolicyResponseModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExcludedUserFilterEnabled

`func (o *AppAccessPolicyResponseModel) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *AppAccessPolicyResponseModel) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *AppAccessPolicyResponseModel) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.


### GetExcludedUsers

`func (o *AppAccessPolicyResponseModel) GetExcludedUsers() []IdentityUserResponseModel`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *AppAccessPolicyResponseModel) GetExcludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *AppAccessPolicyResponseModel) SetExcludedUsers(v []IdentityUserResponseModel)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *AppAccessPolicyResponseModel) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.

### SetExcludedUsersNil

`func (o *AppAccessPolicyResponseModel) SetExcludedUsersNil(b bool)`

 SetExcludedUsersNil sets the value for ExcludedUsers to be an explicit nil

### UnsetExcludedUsers
`func (o *AppAccessPolicyResponseModel) UnsetExcludedUsers()`

UnsetExcludedUsers ensures that no value is present for ExcludedUsers, not even an explicit nil
### GetIncludedUserFilterEnabled

`func (o *AppAccessPolicyResponseModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *AppAccessPolicyResponseModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *AppAccessPolicyResponseModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.


### GetIncludedUsers

`func (o *AppAccessPolicyResponseModel) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *AppAccessPolicyResponseModel) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *AppAccessPolicyResponseModel) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *AppAccessPolicyResponseModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *AppAccessPolicyResponseModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *AppAccessPolicyResponseModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil
### GetLeasingBehavior

`func (o *AppAccessPolicyResponseModel) GetLeasingBehavior() LeasingBehavior`

GetLeasingBehavior returns the LeasingBehavior field if non-nil, zero value otherwise.

### GetLeasingBehaviorOk

`func (o *AppAccessPolicyResponseModel) GetLeasingBehaviorOk() (*LeasingBehavior, bool)`

GetLeasingBehaviorOk returns a tuple with the LeasingBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasingBehavior

`func (o *AppAccessPolicyResponseModel) SetLeasingBehavior(v LeasingBehavior)`

SetLeasingBehavior sets LeasingBehavior field to given value.

### HasLeasingBehavior

`func (o *AppAccessPolicyResponseModel) HasLeasingBehavior() bool`

HasLeasingBehavior returns a boolean if a field has been set.

### GetSessionReconnection

`func (o *AppAccessPolicyResponseModel) GetSessionReconnection() SessionReconnection`

GetSessionReconnection returns the SessionReconnection field if non-nil, zero value otherwise.

### GetSessionReconnectionOk

`func (o *AppAccessPolicyResponseModel) GetSessionReconnectionOk() (*SessionReconnection, bool)`

GetSessionReconnectionOk returns a tuple with the SessionReconnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReconnection

`func (o *AppAccessPolicyResponseModel) SetSessionReconnection(v SessionReconnection)`

SetSessionReconnection sets SessionReconnection field to given value.

### HasSessionReconnection

`func (o *AppAccessPolicyResponseModel) HasSessionReconnection() bool`

HasSessionReconnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


