# PolicyBodyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Policy name. | [optional] 
**IsEnabled** | Pointer to **bool** | Is policy enabled. | [optional] 
**Description** | Pointer to **NullableString** | Policy description. | [optional] 

## Methods

### NewPolicyBodyRequest

`func NewPolicyBodyRequest() *PolicyBodyRequest`

NewPolicyBodyRequest instantiates a new PolicyBodyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyBodyRequestWithDefaults

`func NewPolicyBodyRequestWithDefaults() *PolicyBodyRequest`

NewPolicyBodyRequestWithDefaults instantiates a new PolicyBodyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyBodyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyBodyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyBodyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyBodyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PolicyBodyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PolicyBodyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsEnabled

`func (o *PolicyBodyRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PolicyBodyRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PolicyBodyRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PolicyBodyRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyBodyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyBodyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyBodyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyBodyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicyBodyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicyBodyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


