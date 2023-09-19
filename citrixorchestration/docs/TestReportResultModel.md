# TestReportResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestName** | Pointer to **NullableString** | The localized display name of the test. | [optional] 
**TestDescription** | Pointer to **NullableString** | The description of the localized end user description of the test. | [optional] 
**TestStartTime** | Pointer to **NullableString** | Gets or sets the test start time. If the test has not yet been run, this may be null. | [optional] 
**FormattedTestStartTime** | Pointer to **NullableString** | Gets or sets the formatted test start time. If the test has not yet been run, this may be null. RFC 3339 compatible format. | [optional] 
**TestEndTime** | Pointer to **NullableString** | Gets or sets the test end time. If the test has not yet been run, this may be null. | [optional] 
**FormattedTestEndTime** | Pointer to **NullableString** | Gets or sets the formatted test end time. If the test has not yet been run, this may be null. RFC 3339 compatible format. | [optional] 
**TestServiceTarget** | Pointer to **NullableString** | Gets or sets the test service target, generally the name of the service the test is being run against. | [optional] 
**TestComponentStatus** | Pointer to **NullableString** | Gets or sets the overall status of this test run.  If the test was run against more than one controller, it is the result of or&#39;ing together of the statuses of each component result. | [optional] 
**TestScope** | Pointer to [**TestScope**](TestScope.md) |  | [optional] 
**TestComponents** | Pointer to [**[]TestComponentResultModel**](TestComponentResultModel.md) | Gets or sets the list of component results. Each machine that the test is run against will be represented by a component result and their statuses aggregated into TestComponentStatus | [optional] 

## Methods

### NewTestReportResultModel

`func NewTestReportResultModel() *TestReportResultModel`

NewTestReportResultModel instantiates a new TestReportResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestReportResultModelWithDefaults

`func NewTestReportResultModelWithDefaults() *TestReportResultModel`

NewTestReportResultModelWithDefaults instantiates a new TestReportResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestName

`func (o *TestReportResultModel) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *TestReportResultModel) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *TestReportResultModel) SetTestName(v string)`

SetTestName sets TestName field to given value.

### HasTestName

`func (o *TestReportResultModel) HasTestName() bool`

HasTestName returns a boolean if a field has been set.

### SetTestNameNil

`func (o *TestReportResultModel) SetTestNameNil(b bool)`

 SetTestNameNil sets the value for TestName to be an explicit nil

### UnsetTestName
`func (o *TestReportResultModel) UnsetTestName()`

UnsetTestName ensures that no value is present for TestName, not even an explicit nil
### GetTestDescription

`func (o *TestReportResultModel) GetTestDescription() string`

GetTestDescription returns the TestDescription field if non-nil, zero value otherwise.

### GetTestDescriptionOk

`func (o *TestReportResultModel) GetTestDescriptionOk() (*string, bool)`

GetTestDescriptionOk returns a tuple with the TestDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDescription

`func (o *TestReportResultModel) SetTestDescription(v string)`

SetTestDescription sets TestDescription field to given value.

### HasTestDescription

`func (o *TestReportResultModel) HasTestDescription() bool`

HasTestDescription returns a boolean if a field has been set.

### SetTestDescriptionNil

`func (o *TestReportResultModel) SetTestDescriptionNil(b bool)`

 SetTestDescriptionNil sets the value for TestDescription to be an explicit nil

### UnsetTestDescription
`func (o *TestReportResultModel) UnsetTestDescription()`

UnsetTestDescription ensures that no value is present for TestDescription, not even an explicit nil
### GetTestStartTime

`func (o *TestReportResultModel) GetTestStartTime() string`

GetTestStartTime returns the TestStartTime field if non-nil, zero value otherwise.

### GetTestStartTimeOk

`func (o *TestReportResultModel) GetTestStartTimeOk() (*string, bool)`

GetTestStartTimeOk returns a tuple with the TestStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStartTime

`func (o *TestReportResultModel) SetTestStartTime(v string)`

SetTestStartTime sets TestStartTime field to given value.

### HasTestStartTime

`func (o *TestReportResultModel) HasTestStartTime() bool`

HasTestStartTime returns a boolean if a field has been set.

### SetTestStartTimeNil

`func (o *TestReportResultModel) SetTestStartTimeNil(b bool)`

 SetTestStartTimeNil sets the value for TestStartTime to be an explicit nil

### UnsetTestStartTime
`func (o *TestReportResultModel) UnsetTestStartTime()`

UnsetTestStartTime ensures that no value is present for TestStartTime, not even an explicit nil
### GetFormattedTestStartTime

`func (o *TestReportResultModel) GetFormattedTestStartTime() string`

GetFormattedTestStartTime returns the FormattedTestStartTime field if non-nil, zero value otherwise.

### GetFormattedTestStartTimeOk

`func (o *TestReportResultModel) GetFormattedTestStartTimeOk() (*string, bool)`

GetFormattedTestStartTimeOk returns a tuple with the FormattedTestStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTestStartTime

`func (o *TestReportResultModel) SetFormattedTestStartTime(v string)`

SetFormattedTestStartTime sets FormattedTestStartTime field to given value.

### HasFormattedTestStartTime

`func (o *TestReportResultModel) HasFormattedTestStartTime() bool`

HasFormattedTestStartTime returns a boolean if a field has been set.

### SetFormattedTestStartTimeNil

`func (o *TestReportResultModel) SetFormattedTestStartTimeNil(b bool)`

 SetFormattedTestStartTimeNil sets the value for FormattedTestStartTime to be an explicit nil

### UnsetFormattedTestStartTime
`func (o *TestReportResultModel) UnsetFormattedTestStartTime()`

UnsetFormattedTestStartTime ensures that no value is present for FormattedTestStartTime, not even an explicit nil
### GetTestEndTime

`func (o *TestReportResultModel) GetTestEndTime() string`

GetTestEndTime returns the TestEndTime field if non-nil, zero value otherwise.

### GetTestEndTimeOk

`func (o *TestReportResultModel) GetTestEndTimeOk() (*string, bool)`

GetTestEndTimeOk returns a tuple with the TestEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestEndTime

`func (o *TestReportResultModel) SetTestEndTime(v string)`

SetTestEndTime sets TestEndTime field to given value.

### HasTestEndTime

`func (o *TestReportResultModel) HasTestEndTime() bool`

HasTestEndTime returns a boolean if a field has been set.

### SetTestEndTimeNil

`func (o *TestReportResultModel) SetTestEndTimeNil(b bool)`

 SetTestEndTimeNil sets the value for TestEndTime to be an explicit nil

### UnsetTestEndTime
`func (o *TestReportResultModel) UnsetTestEndTime()`

UnsetTestEndTime ensures that no value is present for TestEndTime, not even an explicit nil
### GetFormattedTestEndTime

`func (o *TestReportResultModel) GetFormattedTestEndTime() string`

GetFormattedTestEndTime returns the FormattedTestEndTime field if non-nil, zero value otherwise.

### GetFormattedTestEndTimeOk

`func (o *TestReportResultModel) GetFormattedTestEndTimeOk() (*string, bool)`

GetFormattedTestEndTimeOk returns a tuple with the FormattedTestEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTestEndTime

`func (o *TestReportResultModel) SetFormattedTestEndTime(v string)`

SetFormattedTestEndTime sets FormattedTestEndTime field to given value.

### HasFormattedTestEndTime

`func (o *TestReportResultModel) HasFormattedTestEndTime() bool`

HasFormattedTestEndTime returns a boolean if a field has been set.

### SetFormattedTestEndTimeNil

`func (o *TestReportResultModel) SetFormattedTestEndTimeNil(b bool)`

 SetFormattedTestEndTimeNil sets the value for FormattedTestEndTime to be an explicit nil

### UnsetFormattedTestEndTime
`func (o *TestReportResultModel) UnsetFormattedTestEndTime()`

UnsetFormattedTestEndTime ensures that no value is present for FormattedTestEndTime, not even an explicit nil
### GetTestServiceTarget

`func (o *TestReportResultModel) GetTestServiceTarget() string`

GetTestServiceTarget returns the TestServiceTarget field if non-nil, zero value otherwise.

### GetTestServiceTargetOk

`func (o *TestReportResultModel) GetTestServiceTargetOk() (*string, bool)`

GetTestServiceTargetOk returns a tuple with the TestServiceTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestServiceTarget

`func (o *TestReportResultModel) SetTestServiceTarget(v string)`

SetTestServiceTarget sets TestServiceTarget field to given value.

### HasTestServiceTarget

`func (o *TestReportResultModel) HasTestServiceTarget() bool`

HasTestServiceTarget returns a boolean if a field has been set.

### SetTestServiceTargetNil

`func (o *TestReportResultModel) SetTestServiceTargetNil(b bool)`

 SetTestServiceTargetNil sets the value for TestServiceTarget to be an explicit nil

### UnsetTestServiceTarget
`func (o *TestReportResultModel) UnsetTestServiceTarget()`

UnsetTestServiceTarget ensures that no value is present for TestServiceTarget, not even an explicit nil
### GetTestComponentStatus

`func (o *TestReportResultModel) GetTestComponentStatus() string`

GetTestComponentStatus returns the TestComponentStatus field if non-nil, zero value otherwise.

### GetTestComponentStatusOk

`func (o *TestReportResultModel) GetTestComponentStatusOk() (*string, bool)`

GetTestComponentStatusOk returns a tuple with the TestComponentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestComponentStatus

`func (o *TestReportResultModel) SetTestComponentStatus(v string)`

SetTestComponentStatus sets TestComponentStatus field to given value.

### HasTestComponentStatus

`func (o *TestReportResultModel) HasTestComponentStatus() bool`

HasTestComponentStatus returns a boolean if a field has been set.

### SetTestComponentStatusNil

`func (o *TestReportResultModel) SetTestComponentStatusNil(b bool)`

 SetTestComponentStatusNil sets the value for TestComponentStatus to be an explicit nil

### UnsetTestComponentStatus
`func (o *TestReportResultModel) UnsetTestComponentStatus()`

UnsetTestComponentStatus ensures that no value is present for TestComponentStatus, not even an explicit nil
### GetTestScope

`func (o *TestReportResultModel) GetTestScope() TestScope`

GetTestScope returns the TestScope field if non-nil, zero value otherwise.

### GetTestScopeOk

`func (o *TestReportResultModel) GetTestScopeOk() (*TestScope, bool)`

GetTestScopeOk returns a tuple with the TestScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestScope

`func (o *TestReportResultModel) SetTestScope(v TestScope)`

SetTestScope sets TestScope field to given value.

### HasTestScope

`func (o *TestReportResultModel) HasTestScope() bool`

HasTestScope returns a boolean if a field has been set.

### GetTestComponents

`func (o *TestReportResultModel) GetTestComponents() []TestComponentResultModel`

GetTestComponents returns the TestComponents field if non-nil, zero value otherwise.

### GetTestComponentsOk

`func (o *TestReportResultModel) GetTestComponentsOk() (*[]TestComponentResultModel, bool)`

GetTestComponentsOk returns a tuple with the TestComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestComponents

`func (o *TestReportResultModel) SetTestComponents(v []TestComponentResultModel)`

SetTestComponents sets TestComponents field to given value.

### HasTestComponents

`func (o *TestReportResultModel) HasTestComponents() bool`

HasTestComponents returns a boolean if a field has been set.

### SetTestComponentsNil

`func (o *TestReportResultModel) SetTestComponentsNil(b bool)`

 SetTestComponentsNil sets the value for TestComponents to be an explicit nil

### UnsetTestComponents
`func (o *TestReportResultModel) UnsetTestComponents()`

UnsetTestComponents ensures that no value is present for TestComponents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


