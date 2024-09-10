# AdministratorAccessScopeChoicesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Checkable** | [**BooleanPolicyProperty**](BooleanPolicyProperty.md) |  | 

## Methods

### NewAdministratorAccessScopeChoicesModel

`func NewAdministratorAccessScopeChoicesModel(name string, checkable BooleanPolicyProperty, ) *AdministratorAccessScopeChoicesModel`

NewAdministratorAccessScopeChoicesModel instantiates a new AdministratorAccessScopeChoicesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorAccessScopeChoicesModelWithDefaults

`func NewAdministratorAccessScopeChoicesModelWithDefaults() *AdministratorAccessScopeChoicesModel`

NewAdministratorAccessScopeChoicesModelWithDefaults instantiates a new AdministratorAccessScopeChoicesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdministratorAccessScopeChoicesModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdministratorAccessScopeChoicesModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdministratorAccessScopeChoicesModel) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *AdministratorAccessScopeChoicesModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdministratorAccessScopeChoicesModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdministratorAccessScopeChoicesModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AdministratorAccessScopeChoicesModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AdministratorAccessScopeChoicesModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AdministratorAccessScopeChoicesModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetCheckable

`func (o *AdministratorAccessScopeChoicesModel) GetCheckable() BooleanPolicyProperty`

GetCheckable returns the Checkable field if non-nil, zero value otherwise.

### GetCheckableOk

`func (o *AdministratorAccessScopeChoicesModel) GetCheckableOk() (*BooleanPolicyProperty, bool)`

GetCheckableOk returns a tuple with the Checkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckable

`func (o *AdministratorAccessScopeChoicesModel) SetCheckable(v BooleanPolicyProperty)`

SetCheckable sets Checkable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


