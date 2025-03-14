# ImageVersionResourcePoolRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | Pointer to **NullableString** | Hypervisor resource pool Id or name with which image version associates | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image that are specific to the target hosting infrastructure. | [optional] 

## Methods

### NewImageVersionResourcePoolRequestModel

`func NewImageVersionResourcePoolRequestModel() *ImageVersionResourcePoolRequestModel`

NewImageVersionResourcePoolRequestModel instantiates a new ImageVersionResourcePoolRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionResourcePoolRequestModelWithDefaults

`func NewImageVersionResourcePoolRequestModelWithDefaults() *ImageVersionResourcePoolRequestModel`

NewImageVersionResourcePoolRequestModelWithDefaults instantiates a new ImageVersionResourcePoolRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePool

`func (o *ImageVersionResourcePoolRequestModel) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ImageVersionResourcePoolRequestModel) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ImageVersionResourcePoolRequestModel) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ImageVersionResourcePoolRequestModel) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### SetResourcePoolNil

`func (o *ImageVersionResourcePoolRequestModel) SetResourcePoolNil(b bool)`

 SetResourcePoolNil sets the value for ResourcePool to be an explicit nil

### UnsetResourcePool
`func (o *ImageVersionResourcePoolRequestModel) UnsetResourcePool()`

UnsetResourcePool ensures that no value is present for ResourcePool, not even an explicit nil
### GetCustomProperties

`func (o *ImageVersionResourcePoolRequestModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ImageVersionResourcePoolRequestModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ImageVersionResourcePoolRequestModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ImageVersionResourcePoolRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *ImageVersionResourcePoolRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *ImageVersionResourcePoolRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


