# ADSCPConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the service connection point entry in active directory. | [optional] 
**Keywords** | Pointer to **[]string** | The value of keywords attribute of the service connection point entry in active directory. | [optional] 

## Methods

### NewADSCPConfigurationModel

`func NewADSCPConfigurationModel() *ADSCPConfigurationModel`

NewADSCPConfigurationModel instantiates a new ADSCPConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADSCPConfigurationModelWithDefaults

`func NewADSCPConfigurationModelWithDefaults() *ADSCPConfigurationModel`

NewADSCPConfigurationModelWithDefaults instantiates a new ADSCPConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ADSCPConfigurationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ADSCPConfigurationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ADSCPConfigurationModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ADSCPConfigurationModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ADSCPConfigurationModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ADSCPConfigurationModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKeywords

`func (o *ADSCPConfigurationModel) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *ADSCPConfigurationModel) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *ADSCPConfigurationModel) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *ADSCPConfigurationModel) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### SetKeywordsNil

`func (o *ADSCPConfigurationModel) SetKeywordsNil(b bool)`

 SetKeywordsNil sets the value for Keywords to be an explicit nil

### UnsetKeywords
`func (o *ADSCPConfigurationModel) UnsetKeywords()`

UnsetKeywords ensures that no value is present for Keywords, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


