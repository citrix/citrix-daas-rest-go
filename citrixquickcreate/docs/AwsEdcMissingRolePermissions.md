# AwsEdcMissingRolePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to **NullableString** | Determines whether the permissions are Allowed or Denied | [optional] 
**Action** | Pointer to **[]string** | List of permissions | [optional] 
**Resource** | Pointer to **NullableString** | Target resources for the permissions | [optional] 
**Condition** | Pointer to [**NullableConditionalExpression**](ConditionalExpression.md) | Condition statements | [optional] 

## Methods

### NewAwsEdcMissingRolePermissions

`func NewAwsEdcMissingRolePermissions() *AwsEdcMissingRolePermissions`

NewAwsEdcMissingRolePermissions instantiates a new AwsEdcMissingRolePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcMissingRolePermissionsWithDefaults

`func NewAwsEdcMissingRolePermissionsWithDefaults() *AwsEdcMissingRolePermissions`

NewAwsEdcMissingRolePermissionsWithDefaults instantiates a new AwsEdcMissingRolePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *AwsEdcMissingRolePermissions) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *AwsEdcMissingRolePermissions) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *AwsEdcMissingRolePermissions) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *AwsEdcMissingRolePermissions) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### SetEffectNil

`func (o *AwsEdcMissingRolePermissions) SetEffectNil(b bool)`

 SetEffectNil sets the value for Effect to be an explicit nil

### UnsetEffect
`func (o *AwsEdcMissingRolePermissions) UnsetEffect()`

UnsetEffect ensures that no value is present for Effect, not even an explicit nil
### GetAction

`func (o *AwsEdcMissingRolePermissions) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AwsEdcMissingRolePermissions) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AwsEdcMissingRolePermissions) SetAction(v []string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AwsEdcMissingRolePermissions) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *AwsEdcMissingRolePermissions) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *AwsEdcMissingRolePermissions) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetResource

`func (o *AwsEdcMissingRolePermissions) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AwsEdcMissingRolePermissions) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AwsEdcMissingRolePermissions) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AwsEdcMissingRolePermissions) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *AwsEdcMissingRolePermissions) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *AwsEdcMissingRolePermissions) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetCondition

`func (o *AwsEdcMissingRolePermissions) GetCondition() ConditionalExpression`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AwsEdcMissingRolePermissions) GetConditionOk() (*ConditionalExpression, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AwsEdcMissingRolePermissions) SetCondition(v ConditionalExpression)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AwsEdcMissingRolePermissions) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *AwsEdcMissingRolePermissions) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *AwsEdcMissingRolePermissions) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


