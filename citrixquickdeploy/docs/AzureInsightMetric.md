# AzureInsightMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**LocalizableString**](LocalizableString.md) |  | [optional] 
**Unit** | Pointer to [**Unit**](Unit.md) |  | [optional] 
**Data** | Pointer to [**[]MetricData**](MetricData.md) |  | [optional] 

## Methods

### NewAzureInsightMetric

`func NewAzureInsightMetric() *AzureInsightMetric`

NewAzureInsightMetric instantiates a new AzureInsightMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureInsightMetricWithDefaults

`func NewAzureInsightMetricWithDefaults() *AzureInsightMetric`

NewAzureInsightMetricWithDefaults instantiates a new AzureInsightMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureInsightMetric) GetName() LocalizableString`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureInsightMetric) GetNameOk() (*LocalizableString, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureInsightMetric) SetName(v LocalizableString)`

SetName sets Name field to given value.

### HasName

`func (o *AzureInsightMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnit

`func (o *AzureInsightMetric) GetUnit() Unit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AzureInsightMetric) GetUnitOk() (*Unit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AzureInsightMetric) SetUnit(v Unit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AzureInsightMetric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetData

`func (o *AzureInsightMetric) GetData() []MetricData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AzureInsightMetric) GetDataOk() (*[]MetricData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AzureInsightMetric) SetData(v []MetricData)`

SetData sets Data field to given value.

### HasData

`func (o *AzureInsightMetric) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


