# BackupRestoreRemoveHistoryRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **time.Time** | Start date to remove history | 
**EndDate** | **time.Time** | End date to remove history | 

## Methods

### NewBackupRestoreRemoveHistoryRequestModel

`func NewBackupRestoreRemoveHistoryRequestModel(startDate time.Time, endDate time.Time, ) *BackupRestoreRemoveHistoryRequestModel`

NewBackupRestoreRemoveHistoryRequestModel instantiates a new BackupRestoreRemoveHistoryRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreRemoveHistoryRequestModelWithDefaults

`func NewBackupRestoreRemoveHistoryRequestModelWithDefaults() *BackupRestoreRemoveHistoryRequestModel`

NewBackupRestoreRemoveHistoryRequestModelWithDefaults instantiates a new BackupRestoreRemoveHistoryRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *BackupRestoreRemoveHistoryRequestModel) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BackupRestoreRemoveHistoryRequestModel) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BackupRestoreRemoveHistoryRequestModel) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *BackupRestoreRemoveHistoryRequestModel) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BackupRestoreRemoveHistoryRequestModel) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BackupRestoreRemoveHistoryRequestModel) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


