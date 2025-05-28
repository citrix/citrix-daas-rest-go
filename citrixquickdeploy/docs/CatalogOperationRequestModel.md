# CatalogOperationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | Pointer to [**CatalogOperationType**](CatalogOperationType.md) |  | [optional] 
**Parameters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewCatalogOperationRequestModel

`func NewCatalogOperationRequestModel() *CatalogOperationRequestModel`

NewCatalogOperationRequestModel instantiates a new CatalogOperationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOperationRequestModelWithDefaults

`func NewCatalogOperationRequestModelWithDefaults() *CatalogOperationRequestModel`

NewCatalogOperationRequestModelWithDefaults instantiates a new CatalogOperationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *CatalogOperationRequestModel) GetOperationType() CatalogOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *CatalogOperationRequestModel) GetOperationTypeOk() (*CatalogOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *CatalogOperationRequestModel) SetOperationType(v CatalogOperationType)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *CatalogOperationRequestModel) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetParameters

`func (o *CatalogOperationRequestModel) GetParameters() map[string]map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CatalogOperationRequestModel) GetParametersOk() (*map[string]map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CatalogOperationRequestModel) SetParameters(v map[string]map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CatalogOperationRequestModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


