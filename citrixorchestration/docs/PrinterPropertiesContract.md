# PrinterPropertiesContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The FQDN of the printer name. | [optional] 
**Model** | Pointer to **string** | Printer model name. | [optional] 
**Location** | Pointer to **string** | The location. | [optional] 
**Settings** | Pointer to [**PrinterPropertiesContractSettings**](PrinterPropertiesContractSettings.md) |  | [optional] 
**Serialized** | Pointer to **string** | The serialized data. This is here so that JSON serialization can work. This should not be used. | [optional] 

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

### GetSettings

`func (o *PrinterPropertiesContract) GetSettings() PrinterPropertiesContractSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PrinterPropertiesContract) GetSettingsOk() (*PrinterPropertiesContractSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PrinterPropertiesContract) SetSettings(v PrinterPropertiesContractSettings)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


