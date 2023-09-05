# BackupRestoreInformationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeZone** | **string** | Set customer timezone to schedule backups in | 
**TimeZoneOffset** | **int32** | Customer&#39;s selected time zone offset in seconds | 
**Initialized** | **bool** | When true, indicates the customer has gone through the WebStudio initialization process When false, WebStudio will walk the customer through the initialization process to select the timezone and create the first backup schedule | 

## Methods

### NewBackupRestoreInformationRequestModel

`func NewBackupRestoreInformationRequestModel(timeZone string, timeZoneOffset int32, initialized bool, ) *BackupRestoreInformationRequestModel`

NewBackupRestoreInformationRequestModel instantiates a new BackupRestoreInformationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreInformationRequestModelWithDefaults

`func NewBackupRestoreInformationRequestModelWithDefaults() *BackupRestoreInformationRequestModel`

NewBackupRestoreInformationRequestModelWithDefaults instantiates a new BackupRestoreInformationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeZone

`func (o *BackupRestoreInformationRequestModel) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *BackupRestoreInformationRequestModel) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *BackupRestoreInformationRequestModel) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetTimeZoneOffset

`func (o *BackupRestoreInformationRequestModel) GetTimeZoneOffset() int32`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *BackupRestoreInformationRequestModel) GetTimeZoneOffsetOk() (*int32, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *BackupRestoreInformationRequestModel) SetTimeZoneOffset(v int32)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.


### GetInitialized

`func (o *BackupRestoreInformationRequestModel) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *BackupRestoreInformationRequestModel) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *BackupRestoreInformationRequestModel) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


