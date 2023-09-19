# ADDomainResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Domain name | [optional] 
**ForestName** | Pointer to **NullableString** | Parent forest of this domain | [optional] 
**Tag** | Pointer to **map[string]interface{}** | Cookie for use by the IActiveDirectoryService implementation | [optional] 

## Methods

### NewADDomainResponseModel

`func NewADDomainResponseModel() *ADDomainResponseModel`

NewADDomainResponseModel instantiates a new ADDomainResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainResponseModelWithDefaults

`func NewADDomainResponseModelWithDefaults() *ADDomainResponseModel`

NewADDomainResponseModelWithDefaults instantiates a new ADDomainResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ADDomainResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ADDomainResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ADDomainResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ADDomainResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ADDomainResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ADDomainResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetForestName

`func (o *ADDomainResponseModel) GetForestName() string`

GetForestName returns the ForestName field if non-nil, zero value otherwise.

### GetForestNameOk

`func (o *ADDomainResponseModel) GetForestNameOk() (*string, bool)`

GetForestNameOk returns a tuple with the ForestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestName

`func (o *ADDomainResponseModel) SetForestName(v string)`

SetForestName sets ForestName field to given value.

### HasForestName

`func (o *ADDomainResponseModel) HasForestName() bool`

HasForestName returns a boolean if a field has been set.

### SetForestNameNil

`func (o *ADDomainResponseModel) SetForestNameNil(b bool)`

 SetForestNameNil sets the value for ForestName to be an explicit nil

### UnsetForestName
`func (o *ADDomainResponseModel) UnsetForestName()`

UnsetForestName ensures that no value is present for ForestName, not even an explicit nil
### GetTag

`func (o *ADDomainResponseModel) GetTag() map[string]interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ADDomainResponseModel) GetTagOk() (*map[string]interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ADDomainResponseModel) SetTag(v map[string]interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ADDomainResponseModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ADDomainResponseModel) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ADDomainResponseModel) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


