# RunScriptOnVmModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogResourceType** | [**CatalogResourceType**](CatalogResourceType.md) | The type of catalog resource, either VDA or Connector. | 
**ScriptName** | **string** | The name of the script to be run on the vda. | 
**ScriptParameters** | Pointer to **map[string]string** | Parameters to be passed to the script. | [optional] 

## Methods

### NewRunScriptOnVmModel

`func NewRunScriptOnVmModel(catalogResourceType CatalogResourceType, scriptName string, ) *RunScriptOnVmModel`

NewRunScriptOnVmModel instantiates a new RunScriptOnVmModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunScriptOnVmModelWithDefaults

`func NewRunScriptOnVmModelWithDefaults() *RunScriptOnVmModel`

NewRunScriptOnVmModelWithDefaults instantiates a new RunScriptOnVmModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogResourceType

`func (o *RunScriptOnVmModel) GetCatalogResourceType() CatalogResourceType`

GetCatalogResourceType returns the CatalogResourceType field if non-nil, zero value otherwise.

### GetCatalogResourceTypeOk

`func (o *RunScriptOnVmModel) GetCatalogResourceTypeOk() (*CatalogResourceType, bool)`

GetCatalogResourceTypeOk returns a tuple with the CatalogResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogResourceType

`func (o *RunScriptOnVmModel) SetCatalogResourceType(v CatalogResourceType)`

SetCatalogResourceType sets CatalogResourceType field to given value.


### GetScriptName

`func (o *RunScriptOnVmModel) GetScriptName() string`

GetScriptName returns the ScriptName field if non-nil, zero value otherwise.

### GetScriptNameOk

`func (o *RunScriptOnVmModel) GetScriptNameOk() (*string, bool)`

GetScriptNameOk returns a tuple with the ScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptName

`func (o *RunScriptOnVmModel) SetScriptName(v string)`

SetScriptName sets ScriptName field to given value.


### GetScriptParameters

`func (o *RunScriptOnVmModel) GetScriptParameters() map[string]string`

GetScriptParameters returns the ScriptParameters field if non-nil, zero value otherwise.

### GetScriptParametersOk

`func (o *RunScriptOnVmModel) GetScriptParametersOk() (*map[string]string, bool)`

GetScriptParametersOk returns a tuple with the ScriptParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptParameters

`func (o *RunScriptOnVmModel) SetScriptParameters(v map[string]string)`

SetScriptParameters sets ScriptParameters field to given value.

### HasScriptParameters

`func (o *RunScriptOnVmModel) HasScriptParameters() bool`

HasScriptParameters returns a boolean if a field has been set.

### SetScriptParametersNil

`func (o *RunScriptOnVmModel) SetScriptParametersNil(b bool)`

 SetScriptParametersNil sets the value for ScriptParameters to be an explicit nil

### UnsetScriptParameters
`func (o *RunScriptOnVmModel) UnsetScriptParameters()`

UnsetScriptParameters ensures that no value is present for ScriptParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


