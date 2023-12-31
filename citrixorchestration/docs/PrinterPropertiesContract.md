# PrinterPropertiesContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **NullableString** | The FQDN of the printer name. | [optional] 
**Model** | Pointer to **NullableString** | Printer model name. | [optional] 
**Location** | Pointer to **NullableString** | The location. | [optional] 
**Settings** | Pointer to [**PrinterSettingsContract**](PrinterSettingsContract.md) |  | [optional] 
**Serialized** | Pointer to **NullableString** | The serialized data. This is here so that JSON serialization can work. This should not be used. | [optional] 

## Methods

### NewPrinterPropertiesContract

`func NewPrinterPropertiesContract() *PrinterPropertiesContract`

NewPrinterPropertiesContract instantiates a new PrinterPropertiesContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrinterPropertiesContractWithDefaults

`func NewPrinterPropertiesContractWithDefaults() *PrinterPropertiesContract`

NewPrinterPropertiesContractWithDefaults instantiates a new PrinterPropertiesContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PrinterPropertiesContract) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PrinterPropertiesContract) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PrinterPropertiesContract) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PrinterPropertiesContract) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *PrinterPropertiesContract) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *PrinterPropertiesContract) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetModel

`func (o *PrinterPropertiesContract) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PrinterPropertiesContract) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PrinterPropertiesContract) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PrinterPropertiesContract) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *PrinterPropertiesContract) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *PrinterPropertiesContract) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetLocation

`func (o *PrinterPropertiesContract) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PrinterPropertiesContract) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PrinterPropertiesContract) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PrinterPropertiesContract) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PrinterPropertiesContract) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PrinterPropertiesContract) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetSettings

`func (o *PrinterPropertiesContract) GetSettings() PrinterSettingsContract`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PrinterPropertiesContract) GetSettingsOk() (*PrinterSettingsContract, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PrinterPropertiesContract) SetSettings(v PrinterSettingsContract)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PrinterPropertiesContract) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSerialized

`func (o *PrinterPropertiesContract) GetSerialized() string`

GetSerialized returns the Serialized field if non-nil, zero value otherwise.

### GetSerializedOk

`func (o *PrinterPropertiesContract) GetSerializedOk() (*string, bool)`

GetSerializedOk returns a tuple with the Serialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialized

`func (o *PrinterPropertiesContract) SetSerialized(v string)`

SetSerialized sets Serialized field to given value.

### HasSerialized

`func (o *PrinterPropertiesContract) HasSerialized() bool`

HasSerialized returns a boolean if a field has been set.

### SetSerializedNil

`func (o *PrinterPropertiesContract) SetSerializedNil(b bool)`

 SetSerializedNil sets the value for Serialized to be an explicit nil

### UnsetSerialized
`func (o *PrinterPropertiesContract) UnsetSerialized()`

UnsetSerialized ensures that no value is present for Serialized, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


