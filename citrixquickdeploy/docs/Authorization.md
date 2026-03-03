# Authorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAuthorization

`func NewAuthorization() *Authorization`

NewAuthorization instantiates a new Authorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationWithDefaults

`func NewAuthorizationWithDefaults() *Authorization`

NewAuthorizationWithDefaults instantiates a new Authorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Authorization) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Authorization) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Authorization) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Authorization) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *Authorization) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *Authorization) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetScope

`func (o *Authorization) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Authorization) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Authorization) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Authorization) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *Authorization) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *Authorization) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


