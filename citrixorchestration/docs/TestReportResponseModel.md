# TestReportResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestType** | Pointer to [**TestType**](TestType.md) |  | [optional] 
**CreatedBy** | Pointer to **string** | The owner of the test report | [optional] 
**ReportTime** | Pointer to **string** | The time of the test report created | [optional] 
**FormattedReportTime** | Pointer to **string** | The formatted time of the test report created RFC 3339 compatible format. | [optional] 
**TestResults** | Pointer to [**[]TestReportResultModel**](TestReportResultModel.md) | The test results | [optional] 

## Methods

### NewTestReportResponseModel

`func NewTestReportResponseModel() *TestReportResponseModel`

NewTestReportResponseModel instantiates a new TestReportResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestReportResponseModelWithDefaults

`func NewTestReportResponseModelWithDefaults() *TestReportResponseModel`

NewTestReportResponseModelWithDefaults instantiates a new TestReportResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestType

`func (o *TestReportResponseModel) GetTestType() TestType`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *TestReportResponseModel) GetTestTypeOk() (*TestType, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *TestReportResponseModel) SetTestType(v TestType)`

SetTestType sets TestType field to given value.

### HasTestType

`func (o *TestReportResponseModel) HasTestType() bool`

HasTestType returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TestReportResponseModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TestReportResponseModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TestReportResponseModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TestReportResponseModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetReportTime

`func (o *TestReportResponseModel) GetReportTime() string`

GetReportTime returns the ReportTime field if non-nil, zero value otherwise.

### GetReportTimeOk

`func (o *TestReportResponseModel) GetReportTimeOk() (*string, bool)`

GetReportTimeOk returns a tuple with the ReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTime

`func (o *TestReportResponseModel) SetReportTime(v string)`

SetReportTime sets ReportTime field to given value.

### HasReportTime

`func (o *TestReportResponseModel) HasReportTime() bool`

HasReportTime returns a boolean if a field has been set.

### GetFormattedReportTime

`func (o *TestReportResponseModel) GetFormattedReportTime() string`

GetFormattedReportTime returns the FormattedReportTime field if non-nil, zero value otherwise.

### GetFormattedReportTimeOk

`func (o *TestReportResponseModel) GetFormattedReportTimeOk() (*string, bool)`

GetFormattedReportTimeOk returns a tuple with the FormattedReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedReportTime

`func (o *TestReportResponseModel) SetFormattedReportTime(v string)`

SetFormattedReportTime sets FormattedReportTime field to given value.

### HasFormattedReportTime

`func (o *TestReportResponseModel) HasFormattedReportTime() bool`

HasFormattedReportTime returns a boolean if a field has been set.

### GetTestResults

`func (o *TestReportResponseModel) GetTestResults() []TestReportResultModel`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *TestReportResponseModel) GetTestResultsOk() (*[]TestReportResultModel, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *TestReportResponseModel) SetTestResults(v []TestReportResultModel)`

SetTestResults sets TestResults field to given value.

### HasTestResults

`func (o *TestReportResponseModel) HasTestResults() bool`

HasTestResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


