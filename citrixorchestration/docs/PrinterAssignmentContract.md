# PrinterAssignmentContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPrinterOption** | Pointer to [**DefaultPrinterOption**](DefaultPrinterOption.md) |  | [optional] 
**SpecificDefaultPrinter** | Pointer to **NullableString** | The custom printer. If specified, the default printer is not used. | [optional] 
**SessionPrinters** | Pointer to [**[]PrinterPropertiesContract**](PrinterPropertiesContract.md) | Session printers. | [optional] 
**Filters** | Pointer to **[]string** | Client names or IPs. | [optional] 

## Methods

### NewPrinterAssignmentContract

`func NewPrinterAssignmentContract() *PrinterAssignmentContract`

NewPrinterAssignmentContract instantiates a new PrinterAssignmentContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrinterAssignmentContractWithDefaults

`func NewPrinterAssignmentContractWithDefaults() *PrinterAssignmentContract`

NewPrinterAssignmentContractWithDefaults instantiates a new PrinterAssignmentContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPrinterOption

`func (o *PrinterAssignmentContract) GetDefaultPrinterOption() DefaultPrinterOption`

GetDefaultPrinterOption returns the DefaultPrinterOption field if non-nil, zero value otherwise.

### GetDefaultPrinterOptionOk

`func (o *PrinterAssignmentContract) GetDefaultPrinterOptionOk() (*DefaultPrinterOption, bool)`

GetDefaultPrinterOptionOk returns a tuple with the DefaultPrinterOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrinterOption

`func (o *PrinterAssignmentContract) SetDefaultPrinterOption(v DefaultPrinterOption)`

SetDefaultPrinterOption sets DefaultPrinterOption field to given value.

### HasDefaultPrinterOption

`func (o *PrinterAssignmentContract) HasDefaultPrinterOption() bool`

HasDefaultPrinterOption returns a boolean if a field has been set.

### GetSpecificDefaultPrinter

`func (o *PrinterAssignmentContract) GetSpecificDefaultPrinter() string`

GetSpecificDefaultPrinter returns the SpecificDefaultPrinter field if non-nil, zero value otherwise.

### GetSpecificDefaultPrinterOk

`func (o *PrinterAssignmentContract) GetSpecificDefaultPrinterOk() (*string, bool)`

GetSpecificDefaultPrinterOk returns a tuple with the SpecificDefaultPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificDefaultPrinter

`func (o *PrinterAssignmentContract) SetSpecificDefaultPrinter(v string)`

SetSpecificDefaultPrinter sets SpecificDefaultPrinter field to given value.

### HasSpecificDefaultPrinter

`func (o *PrinterAssignmentContract) HasSpecificDefaultPrinter() bool`

HasSpecificDefaultPrinter returns a boolean if a field has been set.

### SetSpecificDefaultPrinterNil

`func (o *PrinterAssignmentContract) SetSpecificDefaultPrinterNil(b bool)`

 SetSpecificDefaultPrinterNil sets the value for SpecificDefaultPrinter to be an explicit nil

### UnsetSpecificDefaultPrinter
`func (o *PrinterAssignmentContract) UnsetSpecificDefaultPrinter()`

UnsetSpecificDefaultPrinter ensures that no value is present for SpecificDefaultPrinter, not even an explicit nil
### GetSessionPrinters

`func (o *PrinterAssignmentContract) GetSessionPrinters() []PrinterPropertiesContract`

GetSessionPrinters returns the SessionPrinters field if non-nil, zero value otherwise.

### GetSessionPrintersOk

`func (o *PrinterAssignmentContract) GetSessionPrintersOk() (*[]PrinterPropertiesContract, bool)`

GetSessionPrintersOk returns a tuple with the SessionPrinters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPrinters

`func (o *PrinterAssignmentContract) SetSessionPrinters(v []PrinterPropertiesContract)`

SetSessionPrinters sets SessionPrinters field to given value.

### HasSessionPrinters

`func (o *PrinterAssignmentContract) HasSessionPrinters() bool`

HasSessionPrinters returns a boolean if a field has been set.

### SetSessionPrintersNil

`func (o *PrinterAssignmentContract) SetSessionPrintersNil(b bool)`

 SetSessionPrintersNil sets the value for SessionPrinters to be an explicit nil

### UnsetSessionPrinters
`func (o *PrinterAssignmentContract) UnsetSessionPrinters()`

UnsetSessionPrinters ensures that no value is present for SessionPrinters, not even an explicit nil
### GetFilters

`func (o *PrinterAssignmentContract) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PrinterAssignmentContract) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PrinterAssignmentContract) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PrinterAssignmentContract) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *PrinterAssignmentContract) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *PrinterAssignmentContract) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


