# ZonesErrorWarningModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfErrors** | Pointer to **int32** |  | [optional] 
**NumberOfWarnings** | Pointer to **int32** |  | [optional] 
**ErrorWarning** | Pointer to [**[]ZoneErrorWarningModel**](ZoneErrorWarningModel.md) |  | [optional] 

## Methods

### NewZonesErrorWarningModel

`func NewZonesErrorWarningModel() *ZonesErrorWarningModel`

NewZonesErrorWarningModel instantiates a new ZonesErrorWarningModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonesErrorWarningModelWithDefaults

`func NewZonesErrorWarningModelWithDefaults() *ZonesErrorWarningModel`

NewZonesErrorWarningModelWithDefaults instantiates a new ZonesErrorWarningModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfErrors

`func (o *ZonesErrorWarningModel) GetNumberOfErrors() int32`

GetNumberOfErrors returns the NumberOfErrors field if non-nil, zero value otherwise.

### GetNumberOfErrorsOk

`func (o *ZonesErrorWarningModel) GetNumberOfErrorsOk() (*int32, bool)`

GetNumberOfErrorsOk returns a tuple with the NumberOfErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfErrors

`func (o *ZonesErrorWarningModel) SetNumberOfErrors(v int32)`

SetNumberOfErrors sets NumberOfErrors field to given value.

### HasNumberOfErrors

`func (o *ZonesErrorWarningModel) HasNumberOfErrors() bool`

HasNumberOfErrors returns a boolean if a field has been set.

### GetNumberOfWarnings

`func (o *ZonesErrorWarningModel) GetNumberOfWarnings() int32`

GetNumberOfWarnings returns the NumberOfWarnings field if non-nil, zero value otherwise.

### GetNumberOfWarningsOk

`func (o *ZonesErrorWarningModel) GetNumberOfWarningsOk() (*int32, bool)`

GetNumberOfWarningsOk returns a tuple with the NumberOfWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfWarnings

`func (o *ZonesErrorWarningModel) SetNumberOfWarnings(v int32)`

SetNumberOfWarnings sets NumberOfWarnings field to given value.

### HasNumberOfWarnings

`func (o *ZonesErrorWarningModel) HasNumberOfWarnings() bool`

HasNumberOfWarnings returns a boolean if a field has been set.

### GetErrorWarning

`func (o *ZonesErrorWarningModel) GetErrorWarning() []ZoneErrorWarningModel`

GetErrorWarning returns the ErrorWarning field if non-nil, zero value otherwise.

### GetErrorWarningOk

`func (o *ZonesErrorWarningModel) GetErrorWarningOk() (*[]ZoneErrorWarningModel, bool)`

GetErrorWarningOk returns a tuple with the ErrorWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorWarning

`func (o *ZonesErrorWarningModel) SetErrorWarning(v []ZoneErrorWarningModel)`

SetErrorWarning sets ErrorWarning field to given value.

### HasErrorWarning

`func (o *ZonesErrorWarningModel) HasErrorWarning() bool`

HasErrorWarning returns a boolean if a field has been set.

### SetErrorWarningNil

`func (o *ZonesErrorWarningModel) SetErrorWarningNil(b bool)`

 SetErrorWarningNil sets the value for ErrorWarning to be an explicit nil

### UnsetErrorWarning
`func (o *ZonesErrorWarningModel) UnsetErrorWarning()`

UnsetErrorWarning ensures that no value is present for ErrorWarning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


