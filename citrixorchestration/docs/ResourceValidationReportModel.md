# ResourceValidationReportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**ResourceValidationCategory**](ResourceValidationCategory.md) |  | [optional] 
**Result** | Pointer to [**ResourceValidationResult**](ResourceValidationResult.md) |  | [optional] 
**Violations** | Pointer to [**[]ResourceValidationViolationModel**](ResourceValidationViolationModel.md) | A list of violations. | [optional] 

## Methods

### NewResourceValidationReportModel

`func NewResourceValidationReportModel() *ResourceValidationReportModel`

NewResourceValidationReportModel instantiates a new ResourceValidationReportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceValidationReportModelWithDefaults

`func NewResourceValidationReportModelWithDefaults() *ResourceValidationReportModel`

NewResourceValidationReportModelWithDefaults instantiates a new ResourceValidationReportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ResourceValidationReportModel) GetCategory() ResourceValidationCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ResourceValidationReportModel) GetCategoryOk() (*ResourceValidationCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ResourceValidationReportModel) SetCategory(v ResourceValidationCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ResourceValidationReportModel) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetResult

`func (o *ResourceValidationReportModel) GetResult() ResourceValidationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ResourceValidationReportModel) GetResultOk() (*ResourceValidationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ResourceValidationReportModel) SetResult(v ResourceValidationResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ResourceValidationReportModel) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetViolations

`func (o *ResourceValidationReportModel) GetViolations() []ResourceValidationViolationModel`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *ResourceValidationReportModel) GetViolationsOk() (*[]ResourceValidationViolationModel, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *ResourceValidationReportModel) SetViolations(v []ResourceValidationViolationModel)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *ResourceValidationReportModel) HasViolations() bool`

HasViolations returns a boolean if a field has been set.

### SetViolationsNil

`func (o *ResourceValidationReportModel) SetViolationsNil(b bool)`

 SetViolationsNil sets the value for Violations to be an explicit nil

### UnsetViolations
`func (o *ResourceValidationReportModel) UnsetViolations()`

UnsetViolations ensures that no value is present for Violations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


