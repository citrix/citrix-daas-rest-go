# HypervisorAllResourceSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionDetail** | [**HypervisorConnectionDetailRequestModel**](HypervisorConnectionDetailRequestModel.md) |  | 
**ResourceSearchRequest** | [**HypervisorResourceSearchRequestModel**](HypervisorResourceSearchRequestModel.md) |  | 

## Methods

### NewHypervisorAllResourceSearchRequestModel

`func NewHypervisorAllResourceSearchRequestModel(connectionDetail HypervisorConnectionDetailRequestModel, resourceSearchRequest HypervisorResourceSearchRequestModel, ) *HypervisorAllResourceSearchRequestModel`

NewHypervisorAllResourceSearchRequestModel instantiates a new HypervisorAllResourceSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorAllResourceSearchRequestModelWithDefaults

`func NewHypervisorAllResourceSearchRequestModelWithDefaults() *HypervisorAllResourceSearchRequestModel`

NewHypervisorAllResourceSearchRequestModelWithDefaults instantiates a new HypervisorAllResourceSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionDetail

`func (o *HypervisorAllResourceSearchRequestModel) GetConnectionDetail() HypervisorConnectionDetailRequestModel`

GetConnectionDetail returns the ConnectionDetail field if non-nil, zero value otherwise.

### GetConnectionDetailOk

`func (o *HypervisorAllResourceSearchRequestModel) GetConnectionDetailOk() (*HypervisorConnectionDetailRequestModel, bool)`

GetConnectionDetailOk returns a tuple with the ConnectionDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDetail

`func (o *HypervisorAllResourceSearchRequestModel) SetConnectionDetail(v HypervisorConnectionDetailRequestModel)`

SetConnectionDetail sets ConnectionDetail field to given value.


### GetResourceSearchRequest

`func (o *HypervisorAllResourceSearchRequestModel) GetResourceSearchRequest() HypervisorResourceSearchRequestModel`

GetResourceSearchRequest returns the ResourceSearchRequest field if non-nil, zero value otherwise.

### GetResourceSearchRequestOk

`func (o *HypervisorAllResourceSearchRequestModel) GetResourceSearchRequestOk() (*HypervisorResourceSearchRequestModel, bool)`

GetResourceSearchRequestOk returns a tuple with the ResourceSearchRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSearchRequest

`func (o *HypervisorAllResourceSearchRequestModel) SetResourceSearchRequest(v HypervisorResourceSearchRequestModel)`

SetResourceSearchRequest sets ResourceSearchRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


