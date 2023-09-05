# JobResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the job in progress. | 
**Type** | [**JobType**](JobType.md) |  | 
**OverallProgressPercent** | Pointer to **int32** | Overall progress percent; 0..100. May be null if the job does not support progress reporting. | [optional] 
**IsCancellable** | Pointer to **bool** | Indicates whether the job may be cancelled. DeleteJob | [optional] 
**Parameters** | [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Indicators to the caller about what object(s) the job is affecting.  Consult the documentation for APIs that initiate jobs to see the properties expected to be present. | 
**SubJobs** | Pointer to [**[]JobResponseModel**](JobResponseModel.md) | List of sub-jobs contained within the job. | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**ResultLocation** | Pointer to **string** | URL where the job results can be obtained. | [optional] 
**ErrorString** | Pointer to **string** | If a job or subjob failed, this will be the reason why the failure occurred, presented in a human-readable format. | [optional] 
**ErrorCode** | Pointer to [**JobErrorCode**](JobErrorCode.md) |  | [optional] 
**ErrorParameters** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | If a job or subjob failed, this will be information about related objects that were involved in the failure. | [optional] 
**CreationTime** | **string** | Time when the job was created. | 
**FormattedCreationTime** | **string** | Formatted time when the job was created. RFC 3339 compatible format. | 
**StartTime** | Pointer to **string** | Time when the job was started. Will be null if the job has not yet started. | [optional] 
**FormattedStartTime** | Pointer to **string** | Formatted time when the job was started. Will be null if the job has not yet started. RFC 3339 compatible format. | [optional] 
**EndTime** | Pointer to **string** | Time when the job was completed. Will be null if the job has not yet completed. | [optional] 
**FormattedEndTime** | Pointer to **string** | Formatted time when the job was completed. Will be null if the job has not yet completed. RFC 3339 compatible format. | [optional] 

## Methods

### NewJobResponseModel

`func NewJobResponseModel(id string, type_ JobType, parameters []NameValueStringPairModel, status JobStatus, creationTime string, formattedCreationTime string, ) *JobResponseModel`

NewJobResponseModel instantiates a new JobResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResponseModelWithDefaults

`func NewJobResponseModelWithDefaults() *JobResponseModel`

NewJobResponseModelWithDefaults instantiates a new JobResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *JobResponseModel) GetType() JobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobResponseModel) GetTypeOk() (*JobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobResponseModel) SetType(v JobType)`

SetType sets Type field to given value.


### GetOverallProgressPercent

`func (o *JobResponseModel) GetOverallProgressPercent() int32`

GetOverallProgressPercent returns the OverallProgressPercent field if non-nil, zero value otherwise.

### GetOverallProgressPercentOk

`func (o *JobResponseModel) GetOverallProgressPercentOk() (*int32, bool)`

GetOverallProgressPercentOk returns a tuple with the OverallProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallProgressPercent

`func (o *JobResponseModel) SetOverallProgressPercent(v int32)`

SetOverallProgressPercent sets OverallProgressPercent field to given value.

### HasOverallProgressPercent

`func (o *JobResponseModel) HasOverallProgressPercent() bool`

HasOverallProgressPercent returns a boolean if a field has been set.

### GetIsCancellable

`func (o *JobResponseModel) GetIsCancellable() bool`

GetIsCancellable returns the IsCancellable field if non-nil, zero value otherwise.

### GetIsCancellableOk

`func (o *JobResponseModel) GetIsCancellableOk() (*bool, bool)`

GetIsCancellableOk returns a tuple with the IsCancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancellable

`func (o *JobResponseModel) SetIsCancellable(v bool)`

SetIsCancellable sets IsCancellable field to given value.

### HasIsCancellable

`func (o *JobResponseModel) HasIsCancellable() bool`

HasIsCancellable returns a boolean if a field has been set.

### GetParameters

`func (o *JobResponseModel) GetParameters() []NameValueStringPairModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *JobResponseModel) GetParametersOk() (*[]NameValueStringPairModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *JobResponseModel) SetParameters(v []NameValueStringPairModel)`

SetParameters sets Parameters field to given value.


### GetSubJobs

`func (o *JobResponseModel) GetSubJobs() []JobResponseModel`

GetSubJobs returns the SubJobs field if non-nil, zero value otherwise.

### GetSubJobsOk

`func (o *JobResponseModel) GetSubJobsOk() (*[]JobResponseModel, bool)`

GetSubJobsOk returns a tuple with the SubJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubJobs

`func (o *JobResponseModel) SetSubJobs(v []JobResponseModel)`

SetSubJobs sets SubJobs field to given value.

### HasSubJobs

`func (o *JobResponseModel) HasSubJobs() bool`

HasSubJobs returns a boolean if a field has been set.

### GetStatus

`func (o *JobResponseModel) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobResponseModel) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobResponseModel) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetResultLocation

`func (o *JobResponseModel) GetResultLocation() string`

GetResultLocation returns the ResultLocation field if non-nil, zero value otherwise.

### GetResultLocationOk

`func (o *JobResponseModel) GetResultLocationOk() (*string, bool)`

GetResultLocationOk returns a tuple with the ResultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultLocation

`func (o *JobResponseModel) SetResultLocation(v string)`

SetResultLocation sets ResultLocation field to given value.

### HasResultLocation

`func (o *JobResponseModel) HasResultLocation() bool`

HasResultLocation returns a boolean if a field has been set.

### GetErrorString

`func (o *JobResponseModel) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *JobResponseModel) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *JobResponseModel) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *JobResponseModel) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetErrorCode

`func (o *JobResponseModel) GetErrorCode() JobErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *JobResponseModel) GetErrorCodeOk() (*JobErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *JobResponseModel) SetErrorCode(v JobErrorCode)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *JobResponseModel) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorParameters

`func (o *JobResponseModel) GetErrorParameters() []NameValueStringPairModel`

GetErrorParameters returns the ErrorParameters field if non-nil, zero value otherwise.

### GetErrorParametersOk

`func (o *JobResponseModel) GetErrorParametersOk() (*[]NameValueStringPairModel, bool)`

GetErrorParametersOk returns a tuple with the ErrorParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorParameters

`func (o *JobResponseModel) SetErrorParameters(v []NameValueStringPairModel)`

SetErrorParameters sets ErrorParameters field to given value.

### HasErrorParameters

`func (o *JobResponseModel) HasErrorParameters() bool`

HasErrorParameters returns a boolean if a field has been set.

### GetCreationTime

`func (o *JobResponseModel) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *JobResponseModel) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *JobResponseModel) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.


### GetFormattedCreationTime

`func (o *JobResponseModel) GetFormattedCreationTime() string`

GetFormattedCreationTime returns the FormattedCreationTime field if non-nil, zero value otherwise.

### GetFormattedCreationTimeOk

`func (o *JobResponseModel) GetFormattedCreationTimeOk() (*string, bool)`

GetFormattedCreationTimeOk returns a tuple with the FormattedCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCreationTime

`func (o *JobResponseModel) SetFormattedCreationTime(v string)`

SetFormattedCreationTime sets FormattedCreationTime field to given value.


### GetStartTime

`func (o *JobResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *JobResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *JobResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *JobResponseModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetFormattedStartTime

`func (o *JobResponseModel) GetFormattedStartTime() string`

GetFormattedStartTime returns the FormattedStartTime field if non-nil, zero value otherwise.

### GetFormattedStartTimeOk

`func (o *JobResponseModel) GetFormattedStartTimeOk() (*string, bool)`

GetFormattedStartTimeOk returns a tuple with the FormattedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedStartTime

`func (o *JobResponseModel) SetFormattedStartTime(v string)`

SetFormattedStartTime sets FormattedStartTime field to given value.

### HasFormattedStartTime

`func (o *JobResponseModel) HasFormattedStartTime() bool`

HasFormattedStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *JobResponseModel) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *JobResponseModel) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *JobResponseModel) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *JobResponseModel) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFormattedEndTime

`func (o *JobResponseModel) GetFormattedEndTime() string`

GetFormattedEndTime returns the FormattedEndTime field if non-nil, zero value otherwise.

### GetFormattedEndTimeOk

`func (o *JobResponseModel) GetFormattedEndTimeOk() (*string, bool)`

GetFormattedEndTimeOk returns a tuple with the FormattedEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedEndTime

`func (o *JobResponseModel) SetFormattedEndTime(v string)`

SetFormattedEndTime sets FormattedEndTime field to given value.

### HasFormattedEndTime

`func (o *JobResponseModel) HasFormattedEndTime() bool`

HasFormattedEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


