# CustomReportInfoRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | End date of the report. When CustomReportDateRangeType is Custom, start date should be provided. | [optional] 
**EndDate** | Pointer to **string** | End date of the report. When CustomReportDateRangeType is Custom, end date should be provided. | [optional] 
**CustomReportFormat** | [**CustomReportFormatType**](CustomReportFormatType.md) |  | 
**CustomReportDateRange** | [**CustomReportDateRangeType**](CustomReportDateRangeType.md) |  | 

## Methods

### NewCustomReportInfoRequestModel

`func NewCustomReportInfoRequestModel(customReportFormat CustomReportFormatType, customReportDateRange CustomReportDateRangeType, ) *CustomReportInfoRequestModel`

NewCustomReportInfoRequestModel instantiates a new CustomReportInfoRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomReportInfoRequestModelWithDefaults

`func NewCustomReportInfoRequestModelWithDefaults() *CustomReportInfoRequestModel`

NewCustomReportInfoRequestModelWithDefaults instantiates a new CustomReportInfoRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *CustomReportInfoRequestModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CustomReportInfoRequestModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CustomReportInfoRequestModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CustomReportInfoRequestModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CustomReportInfoRequestModel) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CustomReportInfoRequestModel) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CustomReportInfoRequestModel) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CustomReportInfoRequestModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCustomReportFormat

`func (o *CustomReportInfoRequestModel) GetCustomReportFormat() CustomReportFormatType`

GetCustomReportFormat returns the CustomReportFormat field if non-nil, zero value otherwise.

### GetCustomReportFormatOk

`func (o *CustomReportInfoRequestModel) GetCustomReportFormatOk() (*CustomReportFormatType, bool)`

GetCustomReportFormatOk returns a tuple with the CustomReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomReportFormat

`func (o *CustomReportInfoRequestModel) SetCustomReportFormat(v CustomReportFormatType)`

SetCustomReportFormat sets CustomReportFormat field to given value.


### GetCustomReportDateRange

`func (o *CustomReportInfoRequestModel) GetCustomReportDateRange() CustomReportDateRangeType`

GetCustomReportDateRange returns the CustomReportDateRange field if non-nil, zero value otherwise.

### GetCustomReportDateRangeOk

`func (o *CustomReportInfoRequestModel) GetCustomReportDateRangeOk() (*CustomReportDateRangeType, bool)`

GetCustomReportDateRangeOk returns a tuple with the CustomReportDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomReportDateRange

`func (o *CustomReportInfoRequestModel) SetCustomReportDateRange(v CustomReportDateRangeType)`

SetCustomReportDateRange sets CustomReportDateRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


