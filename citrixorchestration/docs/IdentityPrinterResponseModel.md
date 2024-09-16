# IdentityPrinterResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrintShareName** | Pointer to **NullableString** | Shared name of the printer. | [optional] 
**ServerName** | Pointer to **NullableString** | The printer server name. | [optional] 
**UncName** | Pointer to **NullableString** | The printer UNC name. | [optional] 
**DriverName** | Pointer to **NullableString** | The printer driver name. | [optional] 
**Location** | Pointer to **NullableString** | The printer location. | [optional] 
**Description** | Pointer to **NullableString** | The printer description. | [optional] 

## Methods

### NewIdentityPrinterResponseModel

`func NewIdentityPrinterResponseModel() *IdentityPrinterResponseModel`

NewIdentityPrinterResponseModel instantiates a new IdentityPrinterResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityPrinterResponseModelWithDefaults

`func NewIdentityPrinterResponseModelWithDefaults() *IdentityPrinterResponseModel`

NewIdentityPrinterResponseModelWithDefaults instantiates a new IdentityPrinterResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrintShareName

`func (o *IdentityPrinterResponseModel) GetPrintShareName() string`

GetPrintShareName returns the PrintShareName field if non-nil, zero value otherwise.

### GetPrintShareNameOk

`func (o *IdentityPrinterResponseModel) GetPrintShareNameOk() (*string, bool)`

GetPrintShareNameOk returns a tuple with the PrintShareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintShareName

`func (o *IdentityPrinterResponseModel) SetPrintShareName(v string)`

SetPrintShareName sets PrintShareName field to given value.

### HasPrintShareName

`func (o *IdentityPrinterResponseModel) HasPrintShareName() bool`

HasPrintShareName returns a boolean if a field has been set.

### SetPrintShareNameNil

`func (o *IdentityPrinterResponseModel) SetPrintShareNameNil(b bool)`

 SetPrintShareNameNil sets the value for PrintShareName to be an explicit nil

### UnsetPrintShareName
`func (o *IdentityPrinterResponseModel) UnsetPrintShareName()`

UnsetPrintShareName ensures that no value is present for PrintShareName, not even an explicit nil
### GetServerName

`func (o *IdentityPrinterResponseModel) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *IdentityPrinterResponseModel) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *IdentityPrinterResponseModel) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *IdentityPrinterResponseModel) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *IdentityPrinterResponseModel) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *IdentityPrinterResponseModel) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetUncName

`func (o *IdentityPrinterResponseModel) GetUncName() string`

GetUncName returns the UncName field if non-nil, zero value otherwise.

### GetUncNameOk

`func (o *IdentityPrinterResponseModel) GetUncNameOk() (*string, bool)`

GetUncNameOk returns a tuple with the UncName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncName

`func (o *IdentityPrinterResponseModel) SetUncName(v string)`

SetUncName sets UncName field to given value.

### HasUncName

`func (o *IdentityPrinterResponseModel) HasUncName() bool`

HasUncName returns a boolean if a field has been set.

### SetUncNameNil

`func (o *IdentityPrinterResponseModel) SetUncNameNil(b bool)`

 SetUncNameNil sets the value for UncName to be an explicit nil

### UnsetUncName
`func (o *IdentityPrinterResponseModel) UnsetUncName()`

UnsetUncName ensures that no value is present for UncName, not even an explicit nil
### GetDriverName

`func (o *IdentityPrinterResponseModel) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *IdentityPrinterResponseModel) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *IdentityPrinterResponseModel) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *IdentityPrinterResponseModel) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### SetDriverNameNil

`func (o *IdentityPrinterResponseModel) SetDriverNameNil(b bool)`

 SetDriverNameNil sets the value for DriverName to be an explicit nil

### UnsetDriverName
`func (o *IdentityPrinterResponseModel) UnsetDriverName()`

UnsetDriverName ensures that no value is present for DriverName, not even an explicit nil
### GetLocation

`func (o *IdentityPrinterResponseModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IdentityPrinterResponseModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IdentityPrinterResponseModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IdentityPrinterResponseModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *IdentityPrinterResponseModel) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *IdentityPrinterResponseModel) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetDescription

`func (o *IdentityPrinterResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityPrinterResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityPrinterResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityPrinterResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IdentityPrinterResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IdentityPrinterResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


