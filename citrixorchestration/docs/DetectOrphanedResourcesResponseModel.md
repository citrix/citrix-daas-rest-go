# DetectOrphanedResourcesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hypervisor** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**DetectScope** | Pointer to [**DetectScope**](DetectScope.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** | The owner of the detect report | [optional] 
**ReportTime** | Pointer to **NullableString** | The time of the detect report created | [optional] 
**DetectStart** | Pointer to **NullableString** | The time of detect start. | [optional] 
**DetectEnd** | Pointer to **NullableString** | The time of detect end. | [optional] 
**NumDetected** | Pointer to **int32** | The number of orphaned resources detected.  | [optional] 
**ErrorMsg** | Pointer to **NullableString** | Error message if orphaned resources detected failed.  | [optional] 
**IsFailed** | Pointer to **bool** | If orphaned resources detected failed | [optional] 
**DetectResults** | Pointer to [**[]DetectOrphanedResourcesResultModel**](DetectOrphanedResourcesResultModel.md) | The detect results | [optional] 

## Methods

### NewDetectOrphanedResourcesResponseModel

`func NewDetectOrphanedResourcesResponseModel() *DetectOrphanedResourcesResponseModel`

NewDetectOrphanedResourcesResponseModel instantiates a new DetectOrphanedResourcesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectOrphanedResourcesResponseModelWithDefaults

`func NewDetectOrphanedResourcesResponseModelWithDefaults() *DetectOrphanedResourcesResponseModel`

NewDetectOrphanedResourcesResponseModelWithDefaults instantiates a new DetectOrphanedResourcesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHypervisor

`func (o *DetectOrphanedResourcesResponseModel) GetHypervisor() RefResponseModel`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *DetectOrphanedResourcesResponseModel) GetHypervisorOk() (*RefResponseModel, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *DetectOrphanedResourcesResponseModel) SetHypervisor(v RefResponseModel)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *DetectOrphanedResourcesResponseModel) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetDetectScope

`func (o *DetectOrphanedResourcesResponseModel) GetDetectScope() DetectScope`

GetDetectScope returns the DetectScope field if non-nil, zero value otherwise.

### GetDetectScopeOk

`func (o *DetectOrphanedResourcesResponseModel) GetDetectScopeOk() (*DetectScope, bool)`

GetDetectScopeOk returns a tuple with the DetectScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectScope

`func (o *DetectOrphanedResourcesResponseModel) SetDetectScope(v DetectScope)`

SetDetectScope sets DetectScope field to given value.

### HasDetectScope

`func (o *DetectOrphanedResourcesResponseModel) HasDetectScope() bool`

HasDetectScope returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DetectOrphanedResourcesResponseModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DetectOrphanedResourcesResponseModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DetectOrphanedResourcesResponseModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DetectOrphanedResourcesResponseModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *DetectOrphanedResourcesResponseModel) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *DetectOrphanedResourcesResponseModel) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetReportTime

`func (o *DetectOrphanedResourcesResponseModel) GetReportTime() string`

GetReportTime returns the ReportTime field if non-nil, zero value otherwise.

### GetReportTimeOk

`func (o *DetectOrphanedResourcesResponseModel) GetReportTimeOk() (*string, bool)`

GetReportTimeOk returns a tuple with the ReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTime

`func (o *DetectOrphanedResourcesResponseModel) SetReportTime(v string)`

SetReportTime sets ReportTime field to given value.

### HasReportTime

`func (o *DetectOrphanedResourcesResponseModel) HasReportTime() bool`

HasReportTime returns a boolean if a field has been set.

### SetReportTimeNil

`func (o *DetectOrphanedResourcesResponseModel) SetReportTimeNil(b bool)`

 SetReportTimeNil sets the value for ReportTime to be an explicit nil

### UnsetReportTime
`func (o *DetectOrphanedResourcesResponseModel) UnsetReportTime()`

UnsetReportTime ensures that no value is present for ReportTime, not even an explicit nil
### GetDetectStart

`func (o *DetectOrphanedResourcesResponseModel) GetDetectStart() string`

GetDetectStart returns the DetectStart field if non-nil, zero value otherwise.

### GetDetectStartOk

`func (o *DetectOrphanedResourcesResponseModel) GetDetectStartOk() (*string, bool)`

GetDetectStartOk returns a tuple with the DetectStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectStart

`func (o *DetectOrphanedResourcesResponseModel) SetDetectStart(v string)`

SetDetectStart sets DetectStart field to given value.

### HasDetectStart

`func (o *DetectOrphanedResourcesResponseModel) HasDetectStart() bool`

HasDetectStart returns a boolean if a field has been set.

### SetDetectStartNil

`func (o *DetectOrphanedResourcesResponseModel) SetDetectStartNil(b bool)`

 SetDetectStartNil sets the value for DetectStart to be an explicit nil

### UnsetDetectStart
`func (o *DetectOrphanedResourcesResponseModel) UnsetDetectStart()`

UnsetDetectStart ensures that no value is present for DetectStart, not even an explicit nil
### GetDetectEnd

`func (o *DetectOrphanedResourcesResponseModel) GetDetectEnd() string`

GetDetectEnd returns the DetectEnd field if non-nil, zero value otherwise.

### GetDetectEndOk

`func (o *DetectOrphanedResourcesResponseModel) GetDetectEndOk() (*string, bool)`

GetDetectEndOk returns a tuple with the DetectEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectEnd

`func (o *DetectOrphanedResourcesResponseModel) SetDetectEnd(v string)`

SetDetectEnd sets DetectEnd field to given value.

### HasDetectEnd

`func (o *DetectOrphanedResourcesResponseModel) HasDetectEnd() bool`

HasDetectEnd returns a boolean if a field has been set.

### SetDetectEndNil

`func (o *DetectOrphanedResourcesResponseModel) SetDetectEndNil(b bool)`

 SetDetectEndNil sets the value for DetectEnd to be an explicit nil

### UnsetDetectEnd
`func (o *DetectOrphanedResourcesResponseModel) UnsetDetectEnd()`

UnsetDetectEnd ensures that no value is present for DetectEnd, not even an explicit nil
### GetNumDetected

`func (o *DetectOrphanedResourcesResponseModel) GetNumDetected() int32`

GetNumDetected returns the NumDetected field if non-nil, zero value otherwise.

### GetNumDetectedOk

`func (o *DetectOrphanedResourcesResponseModel) GetNumDetectedOk() (*int32, bool)`

GetNumDetectedOk returns a tuple with the NumDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDetected

`func (o *DetectOrphanedResourcesResponseModel) SetNumDetected(v int32)`

SetNumDetected sets NumDetected field to given value.

### HasNumDetected

`func (o *DetectOrphanedResourcesResponseModel) HasNumDetected() bool`

HasNumDetected returns a boolean if a field has been set.

### GetErrorMsg

`func (o *DetectOrphanedResourcesResponseModel) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *DetectOrphanedResourcesResponseModel) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *DetectOrphanedResourcesResponseModel) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *DetectOrphanedResourcesResponseModel) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### SetErrorMsgNil

`func (o *DetectOrphanedResourcesResponseModel) SetErrorMsgNil(b bool)`

 SetErrorMsgNil sets the value for ErrorMsg to be an explicit nil

### UnsetErrorMsg
`func (o *DetectOrphanedResourcesResponseModel) UnsetErrorMsg()`

UnsetErrorMsg ensures that no value is present for ErrorMsg, not even an explicit nil
### GetIsFailed

`func (o *DetectOrphanedResourcesResponseModel) GetIsFailed() bool`

GetIsFailed returns the IsFailed field if non-nil, zero value otherwise.

### GetIsFailedOk

`func (o *DetectOrphanedResourcesResponseModel) GetIsFailedOk() (*bool, bool)`

GetIsFailedOk returns a tuple with the IsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFailed

`func (o *DetectOrphanedResourcesResponseModel) SetIsFailed(v bool)`

SetIsFailed sets IsFailed field to given value.

### HasIsFailed

`func (o *DetectOrphanedResourcesResponseModel) HasIsFailed() bool`

HasIsFailed returns a boolean if a field has been set.

### GetDetectResults

`func (o *DetectOrphanedResourcesResponseModel) GetDetectResults() []DetectOrphanedResourcesResultModel`

GetDetectResults returns the DetectResults field if non-nil, zero value otherwise.

### GetDetectResultsOk

`func (o *DetectOrphanedResourcesResponseModel) GetDetectResultsOk() (*[]DetectOrphanedResourcesResultModel, bool)`

GetDetectResultsOk returns a tuple with the DetectResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectResults

`func (o *DetectOrphanedResourcesResponseModel) SetDetectResults(v []DetectOrphanedResourcesResultModel)`

SetDetectResults sets DetectResults field to given value.

### HasDetectResults

`func (o *DetectOrphanedResourcesResponseModel) HasDetectResults() bool`

HasDetectResults returns a boolean if a field has been set.

### SetDetectResultsNil

`func (o *DetectOrphanedResourcesResponseModel) SetDetectResultsNil(b bool)`

 SetDetectResultsNil sets the value for DetectResults to be an explicit nil

### UnsetDetectResults
`func (o *DetectOrphanedResourcesResponseModel) UnsetDetectResults()`

UnsetDetectResults ensures that no value is present for DetectResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


