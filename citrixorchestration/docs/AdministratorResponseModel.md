# AdministratorResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**IdentityUserResponseModel**](IdentityUserResponseModel.md) |  | 
**ScopesAndRoles** | [**[]AdministratorRightResponseModel**](AdministratorRightResponseModel.md) | Rights associated with the administrator. CHANGE: was public ScopeRolePair[] ScopesAndRoles { get; set; } | 
**Enabled** | **bool** | Indicates whether the administrator is enabled. Disabled administrators cannot administer the site UNLESS they are a member of a different user group which is granted access by a different administrator record. | 

## Methods

### NewAdministratorResponseModel

`func NewAdministratorResponseModel(user IdentityUserResponseModel, scopesAndRoles []AdministratorRightResponseModel, enabled bool, ) *AdministratorResponseModel`

NewAdministratorResponseModel instantiates a new AdministratorResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorResponseModelWithDefaults

`func NewAdministratorResponseModelWithDefaults() *AdministratorResponseModel`

NewAdministratorResponseModelWithDefaults instantiates a new AdministratorResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *AdministratorResponseModel) GetUser() IdentityUserResponseModel`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AdministratorResponseModel) GetUserOk() (*IdentityUserResponseModel, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AdministratorResponseModel) SetUser(v IdentityUserResponseModel)`

SetUser sets User field to given value.


### GetScopesAndRoles

`func (o *AdministratorResponseModel) GetScopesAndRoles() []AdministratorRightResponseModel`

GetScopesAndRoles returns the ScopesAndRoles field if non-nil, zero value otherwise.

### GetScopesAndRolesOk

`func (o *AdministratorResponseModel) GetScopesAndRolesOk() (*[]AdministratorRightResponseModel, bool)`

GetScopesAndRolesOk returns a tuple with the ScopesAndRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesAndRoles

`func (o *AdministratorResponseModel) SetScopesAndRoles(v []AdministratorRightResponseModel)`

SetScopesAndRoles sets ScopesAndRoles field to given value.


### GetEnabled

`func (o *AdministratorResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdministratorResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdministratorResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


