# SupportedPersonaAddOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to [**SupportedPersonaAddOnProperty**](SupportedPersonaAddOnProperty.md) | The property of the add-on. Eg. AllocationType | [optional] 
**Values** | Pointer to **map[string]int32** | The value of the add-on property. This is a dictionary where the key is the value of the property and the value is the add-on credits.  For Eg. if the property is AllocationType, the key could be \&quot;Random\&quot; and the value could be 10. | [optional] 

## Methods

### NewSupportedPersonaAddOn

`func NewSupportedPersonaAddOn() *SupportedPersonaAddOn`

NewSupportedPersonaAddOn instantiates a new SupportedPersonaAddOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedPersonaAddOnWithDefaults

`func NewSupportedPersonaAddOnWithDefaults() *SupportedPersonaAddOn`

NewSupportedPersonaAddOnWithDefaults instantiates a new SupportedPersonaAddOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *SupportedPersonaAddOn) GetProperty() SupportedPersonaAddOnProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SupportedPersonaAddOn) GetPropertyOk() (*SupportedPersonaAddOnProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SupportedPersonaAddOn) SetProperty(v SupportedPersonaAddOnProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SupportedPersonaAddOn) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValues

`func (o *SupportedPersonaAddOn) GetValues() map[string]int32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SupportedPersonaAddOn) GetValuesOk() (*map[string]int32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SupportedPersonaAddOn) SetValues(v map[string]int32)`

SetValues sets Values field to given value.

### HasValues

`func (o *SupportedPersonaAddOn) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *SupportedPersonaAddOn) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *SupportedPersonaAddOn) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


