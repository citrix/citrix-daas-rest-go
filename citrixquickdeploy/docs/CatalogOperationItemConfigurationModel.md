# CatalogOperationItemConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**OperationType** | [**CatalogOperationType**](CatalogOperationType.md) |  | 
**Parameters** | **map[string]map[string]interface{}** |  | 
**SupportsXae** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewCatalogOperationItemConfigurationModel

`func NewCatalogOperationItemConfigurationModel(name string, description string, operationType CatalogOperationType, parameters map[string]map[string]interface{}, ) *CatalogOperationItemConfigurationModel`

NewCatalogOperationItemConfigurationModel instantiates a new CatalogOperationItemConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOperationItemConfigurationModelWithDefaults

`func NewCatalogOperationItemConfigurationModelWithDefaults() *CatalogOperationItemConfigurationModel`

NewCatalogOperationItemConfigurationModelWithDefaults instantiates a new CatalogOperationItemConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogOperationItemConfigurationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogOperationItemConfigurationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogOperationItemConfigurationModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CatalogOperationItemConfigurationModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogOperationItemConfigurationModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogOperationItemConfigurationModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOperationType

`func (o *CatalogOperationItemConfigurationModel) GetOperationType() CatalogOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *CatalogOperationItemConfigurationModel) GetOperationTypeOk() (*CatalogOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *CatalogOperationItemConfigurationModel) SetOperationType(v CatalogOperationType)`

SetOperationType sets OperationType field to given value.


### GetParameters

`func (o *CatalogOperationItemConfigurationModel) GetParameters() map[string]map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CatalogOperationItemConfigurationModel) GetParametersOk() (*map[string]map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CatalogOperationItemConfigurationModel) SetParameters(v map[string]map[string]interface{})`

SetParameters sets Parameters field to given value.


### GetSupportsXae

`func (o *CatalogOperationItemConfigurationModel) GetSupportsXae() bool`

GetSupportsXae returns the SupportsXae field if non-nil, zero value otherwise.

### GetSupportsXaeOk

`func (o *CatalogOperationItemConfigurationModel) GetSupportsXaeOk() (*bool, bool)`

GetSupportsXaeOk returns a tuple with the SupportsXae field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsXae

`func (o *CatalogOperationItemConfigurationModel) SetSupportsXae(v bool)`

SetSupportsXae sets SupportsXae field to given value.

### HasSupportsXae

`func (o *CatalogOperationItemConfigurationModel) HasSupportsXae() bool`

HasSupportsXae returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


