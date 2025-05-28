# AzureCostReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureCostWithCitrixMeters** | Pointer to **map[string]float64** | MeterCategory mapping to MeterCost for Azure Meters with defined Citrix Meters | [optional] 
**AzureCostWithoutCitrixMeters** | Pointer to **map[string]float64** | MeterCategory mapping to MeterCost for Azure Meters w/o defined Citrix Meters | [optional] 

## Methods

### NewAzureCostReport

`func NewAzureCostReport() *AzureCostReport`

NewAzureCostReport instantiates a new AzureCostReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCostReportWithDefaults

`func NewAzureCostReportWithDefaults() *AzureCostReport`

NewAzureCostReportWithDefaults instantiates a new AzureCostReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureCostWithCitrixMeters

`func (o *AzureCostReport) GetAzureCostWithCitrixMeters() map[string]float64`

GetAzureCostWithCitrixMeters returns the AzureCostWithCitrixMeters field if non-nil, zero value otherwise.

### GetAzureCostWithCitrixMetersOk

`func (o *AzureCostReport) GetAzureCostWithCitrixMetersOk() (*map[string]float64, bool)`

GetAzureCostWithCitrixMetersOk returns a tuple with the AzureCostWithCitrixMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCostWithCitrixMeters

`func (o *AzureCostReport) SetAzureCostWithCitrixMeters(v map[string]float64)`

SetAzureCostWithCitrixMeters sets AzureCostWithCitrixMeters field to given value.

### HasAzureCostWithCitrixMeters

`func (o *AzureCostReport) HasAzureCostWithCitrixMeters() bool`

HasAzureCostWithCitrixMeters returns a boolean if a field has been set.

### GetAzureCostWithoutCitrixMeters

`func (o *AzureCostReport) GetAzureCostWithoutCitrixMeters() map[string]float64`

GetAzureCostWithoutCitrixMeters returns the AzureCostWithoutCitrixMeters field if non-nil, zero value otherwise.

### GetAzureCostWithoutCitrixMetersOk

`func (o *AzureCostReport) GetAzureCostWithoutCitrixMetersOk() (*map[string]float64, bool)`

GetAzureCostWithoutCitrixMetersOk returns a tuple with the AzureCostWithoutCitrixMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCostWithoutCitrixMeters

`func (o *AzureCostReport) SetAzureCostWithoutCitrixMeters(v map[string]float64)`

SetAzureCostWithoutCitrixMeters sets AzureCostWithoutCitrixMeters field to given value.

### HasAzureCostWithoutCitrixMeters

`func (o *AzureCostReport) HasAzureCostWithoutCitrixMeters() bool`

HasAzureCostWithoutCitrixMeters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


