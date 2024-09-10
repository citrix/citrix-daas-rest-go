# AdministratorAccessPolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ServiceName** | **string** |  | 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Checkable** | [**BooleanPolicyProperty**](BooleanPolicyProperty.md) |  | 
**ScopeChoices** | Pointer to [**AdministratorAccessScopeChoices**](AdministratorAccessScopeChoices.md) |  | [optional] 

## Methods

### NewAdministratorAccessPolicyModel

`func NewAdministratorAccessPolicyModel(name string, serviceName string, checkable BooleanPolicyProperty, ) *AdministratorAccessPolicyModel`

NewAdministratorAccessPolicyModel instantiates a new AdministratorAccessPolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorAccessPolicyModelWithDefaults

`func NewAdministratorAccessPolicyModelWithDefaults() *AdministratorAccessPolicyModel`

NewAdministratorAccessPolicyModelWithDefaults instantiates a new AdministratorAccessPolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdministratorAccessPolicyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdministratorAccessPolicyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdministratorAccessPolicyModel) SetName(v string)`

SetName sets Name field to given value.


### GetServiceName

`func (o *AdministratorAccessPolicyModel) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *AdministratorAccessPolicyModel) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *AdministratorAccessPolicyModel) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetDisplayName

`func (o *AdministratorAccessPolicyModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdministratorAccessPolicyModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdministratorAccessPolicyModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AdministratorAccessPolicyModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AdministratorAccessPolicyModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AdministratorAccessPolicyModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetCheckable

`func (o *AdministratorAccessPolicyModel) GetCheckable() BooleanPolicyProperty`

GetCheckable returns the Checkable field if non-nil, zero value otherwise.

### GetCheckableOk

`func (o *AdministratorAccessPolicyModel) GetCheckableOk() (*BooleanPolicyProperty, bool)`

GetCheckableOk returns a tuple with the Checkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckable

`func (o *AdministratorAccessPolicyModel) SetCheckable(v BooleanPolicyProperty)`

SetCheckable sets Checkable field to given value.


### GetScopeChoices

`func (o *AdministratorAccessPolicyModel) GetScopeChoices() AdministratorAccessScopeChoices`

GetScopeChoices returns the ScopeChoices field if non-nil, zero value otherwise.

### GetScopeChoicesOk

`func (o *AdministratorAccessPolicyModel) GetScopeChoicesOk() (*AdministratorAccessScopeChoices, bool)`

GetScopeChoicesOk returns a tuple with the ScopeChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeChoices

`func (o *AdministratorAccessPolicyModel) SetScopeChoices(v AdministratorAccessScopeChoices)`

SetScopeChoices sets ScopeChoices field to given value.

### HasScopeChoices

`func (o *AdministratorAccessPolicyModel) HasScopeChoices() bool`

HasScopeChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


