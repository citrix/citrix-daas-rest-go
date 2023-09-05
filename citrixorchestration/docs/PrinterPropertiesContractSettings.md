# PrinterPropertiesContractSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverridePaperSize** | Pointer to **bool** |  | [optional] 
**PaperSize** | Pointer to [**PrinterPaperSize**](PrinterPaperSize.md) |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**OverrideCopyCount** | Pointer to **bool** |  | [optional] 
**Collated** | Pointer to **bool** |  | [optional] 
**CopyCount** | Pointer to **int32** |  | [optional] 
**OverrideScale** | Pointer to **bool** |  | [optional] 
**Scale** | Pointer to **int32** |  | [optional] 
**OverrideColor** | Pointer to **bool** |  | [optional] 
**Color** | Pointer to [**PrinterColorSetting**](PrinterColorSetting.md) |  | [optional] 
**OverridePrintQuality** | Pointer to **bool** |  | [optional] 
**PrintQuality** | Pointer to [**PrintQualitySetting**](PrintQualitySetting.md) |  | [optional] 
**XResolution** | Pointer to **int32** |  | [optional] 
**YResolution** | Pointer to **int32** |  | [optional] 
**OverrideOrientation** | Pointer to **bool** |  | [optional] 
**Orientation** | Pointer to [**PrintOrientationSetting**](PrintOrientationSetting.md) |  | [optional] 
**OverrideDuplex** | Pointer to **bool** |  | [optional] 
**Duplex** | Pointer to [**PrintDuplexSetting**](PrintDuplexSetting.md) |  | [optional] 
**OverrideFormName** | Pointer to **bool** |  | [optional] 
**FormName** | Pointer to **string** |  | [optional] 
**OverrideTrueTypeOption** | Pointer to **bool** |  | [optional] 
**TrueTypeOption** | Pointer to [**PrintTrueTypeOption**](PrintTrueTypeOption.md) |  | [optional] 

## Methods

### NewPrinterPropertiesContractSettings

`func NewPrinterPropertiesContractSettings() *PrinterPropertiesContractSettings`

NewPrinterPropertiesContractSettings instantiates a new PrinterPropertiesContractSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrinterPropertiesContractSettingsWithDefaults

`func NewPrinterPropertiesContractSettingsWithDefaults() *PrinterPropertiesContractSettings`

NewPrinterPropertiesContractSettingsWithDefaults instantiates a new PrinterPropertiesContractSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverridePaperSize

`func (o *PrinterPropertiesContractSettings) GetOverridePaperSize() bool`

GetOverridePaperSize returns the OverridePaperSize field if non-nil, zero value otherwise.

### GetOverridePaperSizeOk

`func (o *PrinterPropertiesContractSettings) GetOverridePaperSizeOk() (*bool, bool)`

GetOverridePaperSizeOk returns a tuple with the OverridePaperSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePaperSize

`func (o *PrinterPropertiesContractSettings) SetOverridePaperSize(v bool)`

SetOverridePaperSize sets OverridePaperSize field to given value.

### HasOverridePaperSize

`func (o *PrinterPropertiesContractSettings) HasOverridePaperSize() bool`

HasOverridePaperSize returns a boolean if a field has been set.

### GetPaperSize

`func (o *PrinterPropertiesContractSettings) GetPaperSize() PrinterPaperSize`

GetPaperSize returns the PaperSize field if non-nil, zero value otherwise.

### GetPaperSizeOk

`func (o *PrinterPropertiesContractSettings) GetPaperSizeOk() (*PrinterPaperSize, bool)`

GetPaperSizeOk returns a tuple with the PaperSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperSize

`func (o *PrinterPropertiesContractSettings) SetPaperSize(v PrinterPaperSize)`

SetPaperSize sets PaperSize field to given value.

### HasPaperSize

`func (o *PrinterPropertiesContractSettings) HasPaperSize() bool`

HasPaperSize returns a boolean if a field has been set.

### GetWidth

`func (o *PrinterPropertiesContractSettings) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PrinterPropertiesContractSettings) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PrinterPropertiesContractSettings) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PrinterPropertiesContractSettings) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *PrinterPropertiesContractSettings) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PrinterPropertiesContractSettings) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PrinterPropertiesContractSettings) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *PrinterPropertiesContractSettings) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetOverrideCopyCount

`func (o *PrinterPropertiesContractSettings) GetOverrideCopyCount() bool`

GetOverrideCopyCount returns the OverrideCopyCount field if non-nil, zero value otherwise.

### GetOverrideCopyCountOk

`func (o *PrinterPropertiesContractSettings) GetOverrideCopyCountOk() (*bool, bool)`

GetOverrideCopyCountOk returns a tuple with the OverrideCopyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideCopyCount

`func (o *PrinterPropertiesContractSettings) SetOverrideCopyCount(v bool)`

SetOverrideCopyCount sets OverrideCopyCount field to given value.

### HasOverrideCopyCount

`func (o *PrinterPropertiesContractSettings) HasOverrideCopyCount() bool`

HasOverrideCopyCount returns a boolean if a field has been set.

### GetCollated

`func (o *PrinterPropertiesContractSettings) GetCollated() bool`

GetCollated returns the Collated field if non-nil, zero value otherwise.

### GetCollatedOk

`func (o *PrinterPropertiesContractSettings) GetCollatedOk() (*bool, bool)`

GetCollatedOk returns a tuple with the Collated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollated

`func (o *PrinterPropertiesContractSettings) SetCollated(v bool)`

SetCollated sets Collated field to given value.

### HasCollated

`func (o *PrinterPropertiesContractSettings) HasCollated() bool`

HasCollated returns a boolean if a field has been set.

### GetCopyCount

`func (o *PrinterPropertiesContractSettings) GetCopyCount() int32`

GetCopyCount returns the CopyCount field if non-nil, zero value otherwise.

### GetCopyCountOk

`func (o *PrinterPropertiesContractSettings) GetCopyCountOk() (*int32, bool)`

GetCopyCountOk returns a tuple with the CopyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyCount

`func (o *PrinterPropertiesContractSettings) SetCopyCount(v int32)`

SetCopyCount sets CopyCount field to given value.

### HasCopyCount

`func (o *PrinterPropertiesContractSettings) HasCopyCount() bool`

HasCopyCount returns a boolean if a field has been set.

### GetOverrideScale

`func (o *PrinterPropertiesContractSettings) GetOverrideScale() bool`

GetOverrideScale returns the OverrideScale field if non-nil, zero value otherwise.

### GetOverrideScaleOk

`func (o *PrinterPropertiesContractSettings) GetOverrideScaleOk() (*bool, bool)`

GetOverrideScaleOk returns a tuple with the OverrideScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideScale

`func (o *PrinterPropertiesContractSettings) SetOverrideScale(v bool)`

SetOverrideScale sets OverrideScale field to given value.

### HasOverrideScale

`func (o *PrinterPropertiesContractSettings) HasOverrideScale() bool`

HasOverrideScale returns a boolean if a field has been set.

### GetScale

`func (o *PrinterPropertiesContractSettings) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *PrinterPropertiesContractSettings) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *PrinterPropertiesContractSettings) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *PrinterPropertiesContractSettings) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetOverrideColor

`func (o *PrinterPropertiesContractSettings) GetOverrideColor() bool`

GetOverrideColor returns the OverrideColor field if non-nil, zero value otherwise.

### GetOverrideColorOk

`func (o *PrinterPropertiesContractSettings) GetOverrideColorOk() (*bool, bool)`

GetOverrideColorOk returns a tuple with the OverrideColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideColor

`func (o *PrinterPropertiesContractSettings) SetOverrideColor(v bool)`

SetOverrideColor sets OverrideColor field to given value.

### HasOverrideColor

`func (o *PrinterPropertiesContractSettings) HasOverrideColor() bool`

HasOverrideColor returns a boolean if a field has been set.

### GetColor

`func (o *PrinterPropertiesContractSettings) GetColor() PrinterColorSetting`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PrinterPropertiesContractSettings) GetColorOk() (*PrinterColorSetting, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PrinterPropertiesContractSettings) SetColor(v PrinterColorSetting)`

SetColor sets Color field to given value.

### HasColor

`func (o *PrinterPropertiesContractSettings) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetOverridePrintQuality

`func (o *PrinterPropertiesContractSettings) GetOverridePrintQuality() bool`

GetOverridePrintQuality returns the OverridePrintQuality field if non-nil, zero value otherwise.

### GetOverridePrintQualityOk

`func (o *PrinterPropertiesContractSettings) GetOverridePrintQualityOk() (*bool, bool)`

GetOverridePrintQualityOk returns a tuple with the OverridePrintQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePrintQuality

`func (o *PrinterPropertiesContractSettings) SetOverridePrintQuality(v bool)`

SetOverridePrintQuality sets OverridePrintQuality field to given value.

### HasOverridePrintQuality

`func (o *PrinterPropertiesContractSettings) HasOverridePrintQuality() bool`

HasOverridePrintQuality returns a boolean if a field has been set.

### GetPrintQuality

`func (o *PrinterPropertiesContractSettings) GetPrintQuality() PrintQualitySetting`

GetPrintQuality returns the PrintQuality field if non-nil, zero value otherwise.

### GetPrintQualityOk

`func (o *PrinterPropertiesContractSettings) GetPrintQualityOk() (*PrintQualitySetting, bool)`

GetPrintQualityOk returns a tuple with the PrintQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintQuality

`func (o *PrinterPropertiesContractSettings) SetPrintQuality(v PrintQualitySetting)`

SetPrintQuality sets PrintQuality field to given value.

### HasPrintQuality

`func (o *PrinterPropertiesContractSettings) HasPrintQuality() bool`

HasPrintQuality returns a boolean if a field has been set.

### GetXResolution

`func (o *PrinterPropertiesContractSettings) GetXResolution() int32`

GetXResolution returns the XResolution field if non-nil, zero value otherwise.

### GetXResolutionOk

`func (o *PrinterPropertiesContractSettings) GetXResolutionOk() (*int32, bool)`

GetXResolutionOk returns a tuple with the XResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXResolution

`func (o *PrinterPropertiesContractSettings) SetXResolution(v int32)`

SetXResolution sets XResolution field to given value.

### HasXResolution

`func (o *PrinterPropertiesContractSettings) HasXResolution() bool`

HasXResolution returns a boolean if a field has been set.

### GetYResolution

`func (o *PrinterPropertiesContractSettings) GetYResolution() int32`

GetYResolution returns the YResolution field if non-nil, zero value otherwise.

### GetYResolutionOk

`func (o *PrinterPropertiesContractSettings) GetYResolutionOk() (*int32, bool)`

GetYResolutionOk returns a tuple with the YResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYResolution

`func (o *PrinterPropertiesContractSettings) SetYResolution(v int32)`

SetYResolution sets YResolution field to given value.

### HasYResolution

`func (o *PrinterPropertiesContractSettings) HasYResolution() bool`

HasYResolution returns a boolean if a field has been set.

### GetOverrideOrientation

`func (o *PrinterPropertiesContractSettings) GetOverrideOrientation() bool`

GetOverrideOrientation returns the OverrideOrientation field if non-nil, zero value otherwise.

### GetOverrideOrientationOk

`func (o *PrinterPropertiesContractSettings) GetOverrideOrientationOk() (*bool, bool)`

GetOverrideOrientationOk returns a tuple with the OverrideOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideOrientation

`func (o *PrinterPropertiesContractSettings) SetOverrideOrientation(v bool)`

SetOverrideOrientation sets OverrideOrientation field to given value.

### HasOverrideOrientation

`func (o *PrinterPropertiesContractSettings) HasOverrideOrientation() bool`

HasOverrideOrientation returns a boolean if a field has been set.

### GetOrientation

`func (o *PrinterPropertiesContractSettings) GetOrientation() PrintOrientationSetting`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *PrinterPropertiesContractSettings) GetOrientationOk() (*PrintOrientationSetting, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *PrinterPropertiesContractSettings) SetOrientation(v PrintOrientationSetting)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *PrinterPropertiesContractSettings) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetOverrideDuplex

`func (o *PrinterPropertiesContractSettings) GetOverrideDuplex() bool`

GetOverrideDuplex returns the OverrideDuplex field if non-nil, zero value otherwise.

### GetOverrideDuplexOk

`func (o *PrinterPropertiesContractSettings) GetOverrideDuplexOk() (*bool, bool)`

GetOverrideDuplexOk returns a tuple with the OverrideDuplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDuplex

`func (o *PrinterPropertiesContractSettings) SetOverrideDuplex(v bool)`

SetOverrideDuplex sets OverrideDuplex field to given value.

### HasOverrideDuplex

`func (o *PrinterPropertiesContractSettings) HasOverrideDuplex() bool`

HasOverrideDuplex returns a boolean if a field has been set.

### GetDuplex

`func (o *PrinterPropertiesContractSettings) GetDuplex() PrintDuplexSetting`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *PrinterPropertiesContractSettings) GetDuplexOk() (*PrintDuplexSetting, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *PrinterPropertiesContractSettings) SetDuplex(v PrintDuplexSetting)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *PrinterPropertiesContractSettings) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetOverrideFormName

`func (o *PrinterPropertiesContractSettings) GetOverrideFormName() bool`

GetOverrideFormName returns the OverrideFormName field if non-nil, zero value otherwise.

### GetOverrideFormNameOk

`func (o *PrinterPropertiesContractSettings) GetOverrideFormNameOk() (*bool, bool)`

GetOverrideFormNameOk returns a tuple with the OverrideFormName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideFormName

`func (o *PrinterPropertiesContractSettings) SetOverrideFormName(v bool)`

SetOverrideFormName sets OverrideFormName field to given value.

### HasOverrideFormName

`func (o *PrinterPropertiesContractSettings) HasOverrideFormName() bool`

HasOverrideFormName returns a boolean if a field has been set.

### GetFormName

`func (o *PrinterPropertiesContractSettings) GetFormName() string`

GetFormName returns the FormName field if non-nil, zero value otherwise.

### GetFormNameOk

`func (o *PrinterPropertiesContractSettings) GetFormNameOk() (*string, bool)`

GetFormNameOk returns a tuple with the FormName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormName

`func (o *PrinterPropertiesContractSettings) SetFormName(v string)`

SetFormName sets FormName field to given value.

### HasFormName

`func (o *PrinterPropertiesContractSettings) HasFormName() bool`

HasFormName returns a boolean if a field has been set.

### GetOverrideTrueTypeOption

`func (o *PrinterPropertiesContractSettings) GetOverrideTrueTypeOption() bool`

GetOverrideTrueTypeOption returns the OverrideTrueTypeOption field if non-nil, zero value otherwise.

### GetOverrideTrueTypeOptionOk

`func (o *PrinterPropertiesContractSettings) GetOverrideTrueTypeOptionOk() (*bool, bool)`

GetOverrideTrueTypeOptionOk returns a tuple with the OverrideTrueTypeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTrueTypeOption

`func (o *PrinterPropertiesContractSettings) SetOverrideTrueTypeOption(v bool)`

SetOverrideTrueTypeOption sets OverrideTrueTypeOption field to given value.

### HasOverrideTrueTypeOption

`func (o *PrinterPropertiesContractSettings) HasOverrideTrueTypeOption() bool`

HasOverrideTrueTypeOption returns a boolean if a field has been set.

### GetTrueTypeOption

`func (o *PrinterPropertiesContractSettings) GetTrueTypeOption() PrintTrueTypeOption`

GetTrueTypeOption returns the TrueTypeOption field if non-nil, zero value otherwise.

### GetTrueTypeOptionOk

`func (o *PrinterPropertiesContractSettings) GetTrueTypeOptionOk() (*PrintTrueTypeOption, bool)`

GetTrueTypeOptionOk returns a tuple with the TrueTypeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueTypeOption

`func (o *PrinterPropertiesContractSettings) SetTrueTypeOption(v PrintTrueTypeOption)`

SetTrueTypeOption sets TrueTypeOption field to given value.

### HasTrueTypeOption

`func (o *PrinterPropertiesContractSettings) HasTrueTypeOption() bool`

HasTrueTypeOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


