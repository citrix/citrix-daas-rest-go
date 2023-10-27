# PolicyIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicySetGuid** | Pointer to **string** | GUID of the policy set. | [optional] 
**PolicyGuid** | Pointer to **string** | GUID of the policy. | [optional] 
**PolicyName** | Pointer to **NullableString** | Name of the policy. | [optional] 

## Methods

### NewPolicyIdentity

`func NewPolicyIdentity() *PolicyIdentity`

NewPolicyIdentity instantiates a new PolicyIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyIdentityWithDefaults

`func NewPolicyIdentityWithDefaults() *PolicyIdentity`

NewPolicyIdentityWithDefaults instantiates a new PolicyIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicySetGuid

`func (o *PolicyIdentity) GetPolicySetGuid() string`

GetPolicySetGuid returns the PolicySetGuid field if non-nil, zero value otherwise.

### GetPolicySetGuidOk

`func (o *PolicyIdentity) GetPolicySetGuidOk() (*string, bool)`

GetPolicySetGuidOk returns a tuple with the PolicySetGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetGuid

`func (o *PolicyIdentity) SetPolicySetGuid(v string)`

SetPolicySetGuid sets PolicySetGuid field to given value.

### HasPolicySetGuid

`func (o *PolicyIdentity) HasPolicySetGuid() bool`

HasPolicySetGuid returns a boolean if a field has been set.

### GetPolicyGuid

`func (o *PolicyIdentity) GetPolicyGuid() string`

GetPolicyGuid returns the PolicyGuid field if non-nil, zero value otherwise.

### GetPolicyGuidOk

`func (o *PolicyIdentity) GetPolicyGuidOk() (*string, bool)`

GetPolicyGuidOk returns a tuple with the PolicyGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGuid

`func (o *PolicyIdentity) SetPolicyGuid(v string)`

SetPolicyGuid sets PolicyGuid field to given value.

### HasPolicyGuid

`func (o *PolicyIdentity) HasPolicyGuid() bool`

HasPolicyGuid returns a boolean if a field has been set.

### GetPolicyName

`func (o *PolicyIdentity) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyIdentity) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyIdentity) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicyIdentity) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *PolicyIdentity) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *PolicyIdentity) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


