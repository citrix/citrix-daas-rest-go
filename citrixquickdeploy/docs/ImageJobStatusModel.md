# ImageJobStatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | ID of the job | [optional] 
**JobType** | Pointer to [**NullableCatalogJobType**](CatalogJobType.md) | Type of the job | [optional] 
**OverallProgressPercent** | Pointer to **NullableInt32** | The percentage of progress of the job | [optional] 
**IsCancellable** | Pointer to **bool** | Indicator of whether the job is cancellable | [optional] 
**Parameters** | Pointer to **map[string]string** | Parameters for the job | [optional] 
**SubJobs** | Pointer to [**[]CatalogJobType**](CatalogJobType.md) | Subjobs of the job | [optional] 
**Status** | Pointer to [**NullableTemplateImageState**](TemplateImageState.md) | Current state of the job | [optional] 
**ResultLocation** | Pointer to **NullableString** | Uri for query result of the job | [optional] 
**Warnings** | Pointer to **[]string** | Warnings of the job | [optional] 
**Error** | Pointer to **NullableString** | Error occurred in the job | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Timestamp when the job is created | [optional] 
**StartedAt** | Pointer to **NullableTime** | Timestamp when the job started | [optional] 
**EndedAt** | Pointer to **NullableTime** | Timestamp when the job ended | [optional] 

## Methods

### NewImageJobStatusModel

`func NewImageJobStatusModel() *ImageJobStatusModel`

NewImageJobStatusModel instantiates a new ImageJobStatusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageJobStatusModelWithDefaults

`func NewImageJobStatusModelWithDefaults() *ImageJobStatusModel`

NewImageJobStatusModelWithDefaults instantiates a new ImageJobStatusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageJobStatusModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageJobStatusModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageJobStatusModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageJobStatusModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ImageJobStatusModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ImageJobStatusModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetJobType

`func (o *ImageJobStatusModel) GetJobType() CatalogJobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ImageJobStatusModel) GetJobTypeOk() (*CatalogJobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ImageJobStatusModel) SetJobType(v CatalogJobType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ImageJobStatusModel) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *ImageJobStatusModel) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *ImageJobStatusModel) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetOverallProgressPercent

`func (o *ImageJobStatusModel) GetOverallProgressPercent() int32`

GetOverallProgressPercent returns the OverallProgressPercent field if non-nil, zero value otherwise.

### GetOverallProgressPercentOk

`func (o *ImageJobStatusModel) GetOverallProgressPercentOk() (*int32, bool)`

GetOverallProgressPercentOk returns a tuple with the OverallProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallProgressPercent

`func (o *ImageJobStatusModel) SetOverallProgressPercent(v int32)`

SetOverallProgressPercent sets OverallProgressPercent field to given value.

### HasOverallProgressPercent

`func (o *ImageJobStatusModel) HasOverallProgressPercent() bool`

HasOverallProgressPercent returns a boolean if a field has been set.

### SetOverallProgressPercentNil

`func (o *ImageJobStatusModel) SetOverallProgressPercentNil(b bool)`

 SetOverallProgressPercentNil sets the value for OverallProgressPercent to be an explicit nil

### UnsetOverallProgressPercent
`func (o *ImageJobStatusModel) UnsetOverallProgressPercent()`

UnsetOverallProgressPercent ensures that no value is present for OverallProgressPercent, not even an explicit nil
### GetIsCancellable

`func (o *ImageJobStatusModel) GetIsCancellable() bool`

GetIsCancellable returns the IsCancellable field if non-nil, zero value otherwise.

### GetIsCancellableOk

`func (o *ImageJobStatusModel) GetIsCancellableOk() (*bool, bool)`

GetIsCancellableOk returns a tuple with the IsCancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancellable

`func (o *ImageJobStatusModel) SetIsCancellable(v bool)`

SetIsCancellable sets IsCancellable field to given value.

### HasIsCancellable

`func (o *ImageJobStatusModel) HasIsCancellable() bool`

HasIsCancellable returns a boolean if a field has been set.

### GetParameters

`func (o *ImageJobStatusModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ImageJobStatusModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ImageJobStatusModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ImageJobStatusModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *ImageJobStatusModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *ImageJobStatusModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetSubJobs

`func (o *ImageJobStatusModel) GetSubJobs() []CatalogJobType`

GetSubJobs returns the SubJobs field if non-nil, zero value otherwise.

### GetSubJobsOk

`func (o *ImageJobStatusModel) GetSubJobsOk() (*[]CatalogJobType, bool)`

GetSubJobsOk returns a tuple with the SubJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubJobs

`func (o *ImageJobStatusModel) SetSubJobs(v []CatalogJobType)`

SetSubJobs sets SubJobs field to given value.

### HasSubJobs

`func (o *ImageJobStatusModel) HasSubJobs() bool`

HasSubJobs returns a boolean if a field has been set.

### SetSubJobsNil

`func (o *ImageJobStatusModel) SetSubJobsNil(b bool)`

 SetSubJobsNil sets the value for SubJobs to be an explicit nil

### UnsetSubJobs
`func (o *ImageJobStatusModel) UnsetSubJobs()`

UnsetSubJobs ensures that no value is present for SubJobs, not even an explicit nil
### GetStatus

`func (o *ImageJobStatusModel) GetStatus() TemplateImageState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageJobStatusModel) GetStatusOk() (*TemplateImageState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageJobStatusModel) SetStatus(v TemplateImageState)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImageJobStatusModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ImageJobStatusModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ImageJobStatusModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetResultLocation

`func (o *ImageJobStatusModel) GetResultLocation() string`

GetResultLocation returns the ResultLocation field if non-nil, zero value otherwise.

### GetResultLocationOk

`func (o *ImageJobStatusModel) GetResultLocationOk() (*string, bool)`

GetResultLocationOk returns a tuple with the ResultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultLocation

`func (o *ImageJobStatusModel) SetResultLocation(v string)`

SetResultLocation sets ResultLocation field to given value.

### HasResultLocation

`func (o *ImageJobStatusModel) HasResultLocation() bool`

HasResultLocation returns a boolean if a field has been set.

### SetResultLocationNil

`func (o *ImageJobStatusModel) SetResultLocationNil(b bool)`

 SetResultLocationNil sets the value for ResultLocation to be an explicit nil

### UnsetResultLocation
`func (o *ImageJobStatusModel) UnsetResultLocation()`

UnsetResultLocation ensures that no value is present for ResultLocation, not even an explicit nil
### GetWarnings

`func (o *ImageJobStatusModel) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ImageJobStatusModel) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ImageJobStatusModel) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ImageJobStatusModel) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *ImageJobStatusModel) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *ImageJobStatusModel) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetError

`func (o *ImageJobStatusModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ImageJobStatusModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ImageJobStatusModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ImageJobStatusModel) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ImageJobStatusModel) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ImageJobStatusModel) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetCreatedAt

`func (o *ImageJobStatusModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageJobStatusModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageJobStatusModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ImageJobStatusModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ImageJobStatusModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ImageJobStatusModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetStartedAt

`func (o *ImageJobStatusModel) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ImageJobStatusModel) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ImageJobStatusModel) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ImageJobStatusModel) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *ImageJobStatusModel) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ImageJobStatusModel) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *ImageJobStatusModel) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ImageJobStatusModel) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ImageJobStatusModel) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ImageJobStatusModel) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *ImageJobStatusModel) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *ImageJobStatusModel) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


