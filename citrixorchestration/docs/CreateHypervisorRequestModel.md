# CreateHypervisorRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionDetails** | [**HypervisorConnectionDetailRequestModel**](HypervisorConnectionDetailRequestModel.md) |  | 
**PoolDetails** | Pointer to [**CreateHypervisorResourcePoolRequestModel**](CreateHypervisorResourcePoolRequestModel.md) |  | [optional] 

## Methods

### NewCreateHypervisorRequestModel

`func NewCreateHypervisorRequestModel(connectionDetails HypervisorConnectionDetailRequestModel, ) *CreateHypervisorRequestModel`

NewCreateHypervisorRequestModel instantiates a new CreateHypervisorRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHypervisorRequestModelWithDefaults

`func NewCreateHypervisorRequestModelWithDefaults() *CreateHypervisorRequestModel`

NewCreateHypervisorRequestModelWithDefaults instantiates a new CreateHypervisorRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionDetails

`func (o *CreateHypervisorRequestModel) GetConnectionDetails() HypervisorConnectionDetailRequestModel`

GetConnectionDetails returns the ConnectionDetails field if non-nil, zero value otherwise.

### GetConnectionDetailsOk

`func (o *CreateHypervisorRequestModel) GetConnectionDetailsOk() (*HypervisorConnectionDetailRequestModel, bool)`

GetConnectionDetailsOk returns a tuple with the ConnectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDetails

`func (o *CreateHypervisorRequestModel) SetConnectionDetails(v HypervisorConnectionDetailRequestModel)`

SetConnectionDetails sets ConnectionDetails field to given value.


### GetPoolDetails

`func (o *CreateHypervisorRequestModel) GetPoolDetails() CreateHypervisorResourcePoolRequestModel`

GetPoolDetails returns the PoolDetails field if non-nil, zero value otherwise.

### GetPoolDetailsOk

`func (o *CreateHypervisorRequestModel) GetPoolDetailsOk() (*CreateHypervisorResourcePoolRequestModel, bool)`

GetPoolDetailsOk returns a tuple with the PoolDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolDetails

`func (o *CreateHypervisorRequestModel) SetPoolDetails(v CreateHypervisorResourcePoolRequestModel)`

SetPoolDetails sets PoolDetails field to given value.

### HasPoolDetails

`func (o *CreateHypervisorRequestModel) HasPoolDetails() bool`

HasPoolDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


