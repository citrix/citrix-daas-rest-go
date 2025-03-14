# UpdateImageVersionResourcePoolsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePools** | Pointer to **[]string** | List of references to resource pools. This will be deprecated and replaced by ImageVersionResourcePools in future. | [optional] 
**ImageVersionResourcePools** | Pointer to [**[]ImageVersionResourcePoolRequestModel**](ImageVersionResourcePoolRequestModel.md) | List of references to image version resource pools. This will replace ResourcePools. | [optional] 

## Methods

### NewUpdateImageVersionResourcePoolsRequestModel

`func NewUpdateImageVersionResourcePoolsRequestModel() *UpdateImageVersionResourcePoolsRequestModel`

NewUpdateImageVersionResourcePoolsRequestModel instantiates a new UpdateImageVersionResourcePoolsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateImageVersionResourcePoolsRequestModelWithDefaults

`func NewUpdateImageVersionResourcePoolsRequestModelWithDefaults() *UpdateImageVersionResourcePoolsRequestModel`

NewUpdateImageVersionResourcePoolsRequestModelWithDefaults instantiates a new UpdateImageVersionResourcePoolsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePools

`func (o *UpdateImageVersionResourcePoolsRequestModel) GetResourcePools() []string`

GetResourcePools returns the ResourcePools field if non-nil, zero value otherwise.

### GetResourcePoolsOk

`func (o *UpdateImageVersionResourcePoolsRequestModel) GetResourcePoolsOk() (*[]string, bool)`

GetResourcePoolsOk returns a tuple with the ResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePools

`func (o *UpdateImageVersionResourcePoolsRequestModel) SetResourcePools(v []string)`

SetResourcePools sets ResourcePools field to given value.

### HasResourcePools

`func (o *UpdateImageVersionResourcePoolsRequestModel) HasResourcePools() bool`

HasResourcePools returns a boolean if a field has been set.

### SetResourcePoolsNil

`func (o *UpdateImageVersionResourcePoolsRequestModel) SetResourcePoolsNil(b bool)`

 SetResourcePoolsNil sets the value for ResourcePools to be an explicit nil

### UnsetResourcePools
`func (o *UpdateImageVersionResourcePoolsRequestModel) UnsetResourcePools()`

UnsetResourcePools ensures that no value is present for ResourcePools, not even an explicit nil
### GetImageVersionResourcePools

`func (o *UpdateImageVersionResourcePoolsRequestModel) GetImageVersionResourcePools() []ImageVersionResourcePoolRequestModel`

GetImageVersionResourcePools returns the ImageVersionResourcePools field if non-nil, zero value otherwise.

### GetImageVersionResourcePoolsOk

`func (o *UpdateImageVersionResourcePoolsRequestModel) GetImageVersionResourcePoolsOk() (*[]ImageVersionResourcePoolRequestModel, bool)`

GetImageVersionResourcePoolsOk returns a tuple with the ImageVersionResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersionResourcePools

`func (o *UpdateImageVersionResourcePoolsRequestModel) SetImageVersionResourcePools(v []ImageVersionResourcePoolRequestModel)`

SetImageVersionResourcePools sets ImageVersionResourcePools field to given value.

### HasImageVersionResourcePools

`func (o *UpdateImageVersionResourcePoolsRequestModel) HasImageVersionResourcePools() bool`

HasImageVersionResourcePools returns a boolean if a field has been set.

### SetImageVersionResourcePoolsNil

`func (o *UpdateImageVersionResourcePoolsRequestModel) SetImageVersionResourcePoolsNil(b bool)`

 SetImageVersionResourcePoolsNil sets the value for ImageVersionResourcePools to be an explicit nil

### UnsetImageVersionResourcePools
`func (o *UpdateImageVersionResourcePoolsRequestModel) UnsetImageVersionResourcePools()`

UnsetImageVersionResourcePools ensures that no value is present for ImageVersionResourcePools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


