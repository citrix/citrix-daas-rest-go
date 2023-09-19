# TestComponentResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestComponentTarget** | Pointer to **NullableString** | Gets or sets the test component target, usually the machine the test was run on. | [optional] 
**TestComponentStartTime** | Pointer to **NullableString** | Gets or sets the test component start time. If the test has not yet been run, this may be null. | [optional] 
**FormattedTestComponentStartTime** | Pointer to **NullableString** | Gets or sets the formatted test component start time. If the test has not yet been run, this may be null. RFC 3339 compatible format. | [optional] 
**TestComponentEndTime** | Pointer to **NullableString** | Gets or sets the test component end time. If the test has not yet been run, this may be null. | [optional] 
**FormattedTestComponentEndTime** | Pointer to **NullableString** | Gets or sets the formatted test component end time. If the test has not yet been run, this may be null. RFC 3339 compatible format. | [optional] 
**TestComponentStatus** | Pointer to [**TestStatus**](TestStatus.md) |  | [optional] 
**ResultDetails** | Pointer to [**[]TestComponentResultDetailModel**](TestComponentResultDetailModel.md) | Gets or sets the result details, which is a list of additional information about the state of a machine when a test succeeds, but has warnings, or fails.  There may be more than one value for each severity, for example, multiple warnings about the state of the License server. The values must be localized. | [optional] 

## Methods

### NewTestComponentResultModel

`func NewTestComponentResultModel() *TestComponentResultModel`

NewTestComponentResultModel instantiates a new TestComponentResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestComponentResultModelWithDefaults

`func NewTestComponentResultModelWithDefaults() *TestComponentResultModel`

NewTestComponentResultModelWithDefaults instantiates a new TestComponentResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestComponentTarget

`func (o *TestComponentResultModel) GetTestComponentTarget() string`

GetTestComponentTarget returns the TestComponentTarget field if non-nil, zero value otherwise.

### GetTestComponentTargetOk

`func (o *TestComponentResultModel) GetTestComponentTargetOk() (*string, bool)`

GetTestComponentTargetOk returns a tuple with the TestComponentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestComponentTarget

`func (o *TestComponentResultModel) SetTestComponentTarget(v string)`

SetTestComponentTarget sets TestComponentTarget field to given value.

### HasTestComponentTarget

`func (o *TestComponentResultModel) HasTestComponentTarget() bool`

HasTestComponentTarget returns a boolean if a field has been set.

### SetTestComponentTargetNil

`func (o *TestComponentResultModel) SetTestComponentTargetNil(b bool)`

 SetTestComponentTargetNil sets the value for TestComponentTarget to be an explicit nil

### UnsetTestComponentTarget
`func (o *TestComponentResultModel) UnsetTestComponentTarget()`

UnsetTestComponentTarget ensures that no value is present for TestComponentTarget, not even an explicit nil
### GetTestComponentStartTime

`func (o *TestComponentResultModel) GetTestComponentStartTime() string`

GetTestComponentStartTime returns the TestComponentStartTime field if non-nil, zero value otherwise.

### GetTestComponentStartTimeOk

`func (o *TestComponentResultModel) GetTestComponentStartTimeOk() (*string, bool)`

GetTestComponentStartTimeOk returns a tuple with the TestComponentStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestComponentStartTime

`func (o *TestComponentResultModel) SetTestComponentStartTime(v string)`

SetTestComponentStartTime sets TestComponentStartTime field to given value.

### HasTestComponentStartTime

`func (o *TestComponentResultModel) HasTestComponentStartTime() bool`

HasTestComponentStartTime returns a boolean if a field has been set.

### SetTestComponentStartTimeNil

`func (o *TestComponentResultModel) SetTestComponentStartTimeNil(b bool)`

 SetTestComponentStartTimeNil sets the value for TestComponentStartTime to be an explicit nil

### UnsetTestComponentStartTime
`func (o *TestComponentResultModel) UnsetTestComponentStartTime()`

UnsetTestComponentStartTime ensures that no value is present for TestComponentStartTime, not even an explicit nil
### GetFormattedTestComponentStartTime

`func (o *TestComponentResultModel) GetFormattedTestComponentStartTime() string`

GetFormattedTestComponentStartTime returns the FormattedTestComponentStartTime field if non-nil, zero value otherwise.

### GetFormattedTestComponentStartTimeOk

`func (o *TestComponentResultModel) GetFormattedTestComponentStartTimeOk() (*string, bool)`

GetFormattedTestComponentStartTimeOk returns a tuple with the FormattedTestComponentStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTestComponentStartTime

`func (o *TestComponentResultModel) SetFormattedTestComponentStartTime(v string)`

SetFormattedTestComponentStartTime sets FormattedTestComponentStartTime field to given value.

### HasFormattedTestComponentStartTime

`func (o *TestComponentResultModel) HasFormattedTestComponentStartTime() bool`

HasFormattedTestComponentStartTime returns a boolean if a field has been set.

### SetFormattedTestComponentStartTimeNil

`func (o *TestComponentResultModel) SetFormattedTestComponentStartTimeNil(b bool)`

 SetFormattedTestComponentStartTimeNil sets the value for FormattedTestComponentStartTime to be an explicit nil

### UnsetFormattedTestComponentStartTime
`func (o *TestComponentResultModel) UnsetFormattedTestComponentStartTime()`

UnsetFormattedTestComponentStartTime ensures that no value is present for FormattedTestComponentStartTime, not even an explicit nil
### GetTestComponentEndTime

`func (o *TestComponentResultModel) GetTestComponentEndTime() string`

GetTestComponentEndTime returns the TestComponentEndTime field if non-nil, zero value otherwise.

### GetTestComponentEndTimeOk

`func (o *TestComponentResultModel) GetTestComponentEndTimeOk() (*string, bool)`

GetTestComponentEndTimeOk returns a tuple with the TestComponentEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestComponentEndTime

`func (o *TestComponentResultModel) SetTestComponentEndTime(v string)`

SetTestComponentEndTime sets TestComponentEndTime field to given value.

### HasTestComponentEndTime

`func (o *TestComponentResultModel) HasTestComponentEndTime() bool`

HasTestComponentEndTime returns a boolean if a field has been set.

### SetTestComponentEndTimeNil

`func (o *TestComponentResultModel) SetTestComponentEndTimeNil(b bool)`

 SetTestComponentEndTimeNil sets the value for TestComponentEndTime to be an explicit nil

### UnsetTestComponentEndTime
`func (o *TestComponentResultModel) UnsetTestComponentEndTime()`

UnsetTestComponentEndTime ensures that no value is present for TestComponentEndTime, not even an explicit nil
### GetFormattedTestComponentEndTime

`func (o *TestComponentResultModel) GetFormattedTestComponentEndTime() string`

GetFormattedTestComponentEndTime returns the FormattedTestComponentEndTime field if non-nil, zero value otherwise.

### GetFormattedTestComponentEndTimeOk

`func (o *TestComponentResultModel) GetFormattedTestComponentEndTimeOk() (*string, bool)`

GetFormattedTestComponentEndTimeOk returns a tuple with the FormattedTestComponentEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTestComponentEndTime

`func (o *TestComponentResultModel) SetFormattedTestComponentEndTime(v string)`

SetFormattedTestComponentEndTime sets FormattedTestComponentEndTime field to given value.

### HasFormattedTestComponentEndTime

`func (o *TestComponentResultModel) HasFormattedTestComponentEndTime() bool`

HasFormattedTestComponentEndTime returns a boolean if a field has been set.

### SetFormattedTestComponentEndTimeNil

`func (o *TestComponentResultModel) SetFormattedTestComponentEndTimeNil(b bool)`

 SetFormattedTestComponentEndTimeNil sets the value for FormattedTestComponentEndTime to be an explicit nil

### UnsetFormattedTestComponentEndTime
`func (o *TestComponentResultModel) UnsetFormattedTestComponentEndTime()`

UnsetFormattedTestComponentEndTime ensures that no value is present for FormattedTestComponentEndTime, not even an explicit nil
### GetTestComponentStatus

`func (o *TestComponentResultModel) GetTestComponentStatus() TestStatus`

GetTestComponentStatus returns the TestComponentStatus field if non-nil, zero value otherwise.

### GetTestComponentStatusOk

`func (o *TestComponentResultModel) GetTestComponentStatusOk() (*TestStatus, bool)`

GetTestComponentStatusOk returns a tuple with the TestComponentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestComponentStatus

`func (o *TestComponentResultModel) SetTestComponentStatus(v TestStatus)`

SetTestComponentStatus sets TestComponentStatus field to given value.

### HasTestComponentStatus

`func (o *TestComponentResultModel) HasTestComponentStatus() bool`

HasTestComponentStatus returns a boolean if a field has been set.

### GetResultDetails

`func (o *TestComponentResultModel) GetResultDetails() []TestComponentResultDetailModel`

GetResultDetails returns the ResultDetails field if non-nil, zero value otherwise.

### GetResultDetailsOk

`func (o *TestComponentResultModel) GetResultDetailsOk() (*[]TestComponentResultDetailModel, bool)`

GetResultDetailsOk returns a tuple with the ResultDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultDetails

`func (o *TestComponentResultModel) SetResultDetails(v []TestComponentResultDetailModel)`

SetResultDetails sets ResultDetails field to given value.

### HasResultDetails

`func (o *TestComponentResultModel) HasResultDetails() bool`

HasResultDetails returns a boolean if a field has been set.

### SetResultDetailsNil

`func (o *TestComponentResultModel) SetResultDetailsNil(b bool)`

 SetResultDetailsNil sets the value for ResultDetails to be an explicit nil

### UnsetResultDetails
`func (o *TestComponentResultModel) UnsetResultDetails()`

UnsetResultDetails ensures that no value is present for ResultDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


