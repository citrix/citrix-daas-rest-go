# ApplicationGroupDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationCount** | **int32** | Number of applications in the group. | 
**IncludedUsersFilterEnabled** | **bool** | Indicates whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access polic(ies) is implicitly granted access to the applications in the application group. | 
**IncludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The included users filter of the application group; that is, the users and groups who are explicitly granted access to the applications in the application group. | [optional] 
**SessionSharingEnabled** | **bool** | Whether applications in this application group can share sessions with applications in other application groups (or with applications that are not a member of an application group). | 
**TotalMachines** | **int32** | Total number of machines across all delivery groups on which the application group is published. | 

## Methods

### NewApplicationGroupDetailResponseModelAllOf

`func NewApplicationGroupDetailResponseModelAllOf(applicationCount int32, includedUsersFilterEnabled bool, sessionSharingEnabled bool, totalMachines int32, ) *ApplicationGroupDetailResponseModelAllOf`

NewApplicationGroupDetailResponseModelAllOf instantiates a new ApplicationGroupDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGroupDetailResponseModelAllOfWithDefaults

`func NewApplicationGroupDetailResponseModelAllOfWithDefaults() *ApplicationGroupDetailResponseModelAllOf`

NewApplicationGroupDetailResponseModelAllOfWithDefaults instantiates a new ApplicationGroupDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationCount

`func (o *ApplicationGroupDetailResponseModelAllOf) GetApplicationCount() int32`

GetApplicationCount returns the ApplicationCount field if non-nil, zero value otherwise.

### GetApplicationCountOk

`func (o *ApplicationGroupDetailResponseModelAllOf) GetApplicationCountOk() (*int32, bool)`

GetApplicationCountOk returns a tuple with the ApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCount

`func (o *ApplicationGroupDetailResponseModelAllOf) SetApplicationCount(v int32)`

SetApplicationCount sets ApplicationCount field to given value.


### GetIncludedUsersFilterEnabled

`func (o *ApplicationGroupDetailResponseModelAllOf) GetIncludedUsersFilterEnabled() bool`

GetIncludedUsersFilterEnabled returns the IncludedUsersFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUsersFilterEnabledOk

`func (o *ApplicationGroupDetailResponseModelAllOf) GetIncludedUsersFilterEnabledOk() (*bool, bool)`

GetIncludedUsersFilterEnabledOk returns a tuple with the IncludedUsersFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsersFilterEnabled

`func (o *ApplicationGroupDetailResponseModelAllOf) SetIncludedUsersFilterEnabled(v bool)`

SetIncludedUsersFilterEnabled sets IncludedUsersFilterEnabled field to given value.


### GetIncludedUsers

`func (o *ApplicationGroupDetailResponseModelAllOf) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *ApplicationGroupDetailResponseModelAllOf) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *ApplicationGroupDetailResponseModelAllOf) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *ApplicationGroupDetailResponseModelAllOf) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### GetSessionSharingEnabled

`func (o *ApplicationGroupDetailResponseModelAllOf) GetSessionSharingEnabled() bool`

GetSessionSharingEnabled returns the SessionSharingEnabled field if non-nil, zero value otherwise.

### GetSessionSharingEnabledOk

`func (o *ApplicationGroupDetailResponseModelAllOf) GetSessionSharingEnabledOk() (*bool, bool)`

GetSessionSharingEnabledOk returns a tuple with the SessionSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSharingEnabled

`func (o *ApplicationGroupDetailResponseModelAllOf) SetSessionSharingEnabled(v bool)`

SetSessionSharingEnabled sets SessionSharingEnabled field to given value.


### GetTotalMachines

`func (o *ApplicationGroupDetailResponseModelAllOf) GetTotalMachines() int32`

GetTotalMachines returns the TotalMachines field if non-nil, zero value otherwise.

### GetTotalMachinesOk

`func (o *ApplicationGroupDetailResponseModelAllOf) GetTotalMachinesOk() (*int32, bool)`

GetTotalMachinesOk returns a tuple with the TotalMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMachines

`func (o *ApplicationGroupDetailResponseModelAllOf) SetTotalMachines(v int32)`

SetTotalMachines sets TotalMachines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


