# ImageVersionSpecRefResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Id of the image version specification. | 
**PreparationType** | [**PreparationType**](PreparationType.md) |  | 
**ResourcePool** | [**HypervisorResourcePoolRefResponseModel**](HypervisorResourcePoolRefResponseModel.md) |  | 

## Methods

### NewImageVersionSpecRefResponseModel

`func NewImageVersionSpecRefResponseModel(id string, preparationType PreparationType, resourcePool HypervisorResourcePoolRefResponseModel, ) *ImageVersionSpecRefResponseModel`

NewImageVersionSpecRefResponseModel instantiates a new ImageVersionSpecRefResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionSpecRefResponseModelWithDefaults

`func NewImageVersionSpecRefResponseModelWithDefaults() *ImageVersionSpecRefResponseModel`

NewImageVersionSpecRefResponseModelWithDefaults instantiates a new ImageVersionSpecRefResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageVersionSpecRefResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionSpecRefResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionSpecRefResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetPreparationType

`func (o *ImageVersionSpecRefResponseModel) GetPreparationType() PreparationType`

GetPreparationType returns the PreparationType field if non-nil, zero value otherwise.

### GetPreparationTypeOk

`func (o *ImageVersionSpecRefResponseModel) GetPreparationTypeOk() (*PreparationType, bool)`

GetPreparationTypeOk returns a tuple with the PreparationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationType

`func (o *ImageVersionSpecRefResponseModel) SetPreparationType(v PreparationType)`

SetPreparationType sets PreparationType field to given value.


### GetResourcePool

`func (o *ImageVersionSpecRefResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ImageVersionSpecRefResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ImageVersionSpecRefResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel)`

SetResourcePool sets ResourcePool field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


