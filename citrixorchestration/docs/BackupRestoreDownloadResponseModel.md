# BackupRestoreDownloadResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**BackupRestoreActionStartStatus**](BackupRestoreActionStartStatus.md) |  | [optional] 
**BackupBinaryData** | Pointer to **NullableString** | Downloaded binary data; may be complete or partial | [optional] 
**BackupSize** | Pointer to **int64** | Total backup size | [optional] 
**SizeTransferred** | Pointer to **int32** | Size transferred in this response | [optional] 
**BackupMemberCount** | Pointer to **int32** | Number of members in downloaded zip when more than a backup is included (such as an associated restore failure) | [optional] 
**ContinuationToken** | Pointer to **NullableString** | Continuation token when this is a partial download. When empty then the download is complete. | [optional] 

## Methods

### NewBackupRestoreDownloadResponseModel

`func NewBackupRestoreDownloadResponseModel() *BackupRestoreDownloadResponseModel`

NewBackupRestoreDownloadResponseModel instantiates a new BackupRestoreDownloadResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreDownloadResponseModelWithDefaults

`func NewBackupRestoreDownloadResponseModelWithDefaults() *BackupRestoreDownloadResponseModel`

NewBackupRestoreDownloadResponseModelWithDefaults instantiates a new BackupRestoreDownloadResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *BackupRestoreDownloadResponseModel) GetResult() BackupRestoreActionStartStatus`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BackupRestoreDownloadResponseModel) GetResultOk() (*BackupRestoreActionStartStatus, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BackupRestoreDownloadResponseModel) SetResult(v BackupRestoreActionStartStatus)`

SetResult sets Result field to given value.

### HasResult

`func (o *BackupRestoreDownloadResponseModel) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetBackupBinaryData

`func (o *BackupRestoreDownloadResponseModel) GetBackupBinaryData() string`

GetBackupBinaryData returns the BackupBinaryData field if non-nil, zero value otherwise.

### GetBackupBinaryDataOk

`func (o *BackupRestoreDownloadResponseModel) GetBackupBinaryDataOk() (*string, bool)`

GetBackupBinaryDataOk returns a tuple with the BackupBinaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupBinaryData

`func (o *BackupRestoreDownloadResponseModel) SetBackupBinaryData(v string)`

SetBackupBinaryData sets BackupBinaryData field to given value.

### HasBackupBinaryData

`func (o *BackupRestoreDownloadResponseModel) HasBackupBinaryData() bool`

HasBackupBinaryData returns a boolean if a field has been set.

### SetBackupBinaryDataNil

`func (o *BackupRestoreDownloadResponseModel) SetBackupBinaryDataNil(b bool)`

 SetBackupBinaryDataNil sets the value for BackupBinaryData to be an explicit nil

### UnsetBackupBinaryData
`func (o *BackupRestoreDownloadResponseModel) UnsetBackupBinaryData()`

UnsetBackupBinaryData ensures that no value is present for BackupBinaryData, not even an explicit nil
### GetBackupSize

`func (o *BackupRestoreDownloadResponseModel) GetBackupSize() int64`

GetBackupSize returns the BackupSize field if non-nil, zero value otherwise.

### GetBackupSizeOk

`func (o *BackupRestoreDownloadResponseModel) GetBackupSizeOk() (*int64, bool)`

GetBackupSizeOk returns a tuple with the BackupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSize

`func (o *BackupRestoreDownloadResponseModel) SetBackupSize(v int64)`

SetBackupSize sets BackupSize field to given value.

### HasBackupSize

`func (o *BackupRestoreDownloadResponseModel) HasBackupSize() bool`

HasBackupSize returns a boolean if a field has been set.

### GetSizeTransferred

`func (o *BackupRestoreDownloadResponseModel) GetSizeTransferred() int32`

GetSizeTransferred returns the SizeTransferred field if non-nil, zero value otherwise.

### GetSizeTransferredOk

`func (o *BackupRestoreDownloadResponseModel) GetSizeTransferredOk() (*int32, bool)`

GetSizeTransferredOk returns a tuple with the SizeTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTransferred

`func (o *BackupRestoreDownloadResponseModel) SetSizeTransferred(v int32)`

SetSizeTransferred sets SizeTransferred field to given value.

### HasSizeTransferred

`func (o *BackupRestoreDownloadResponseModel) HasSizeTransferred() bool`

HasSizeTransferred returns a boolean if a field has been set.

### GetBackupMemberCount

`func (o *BackupRestoreDownloadResponseModel) GetBackupMemberCount() int32`

GetBackupMemberCount returns the BackupMemberCount field if non-nil, zero value otherwise.

### GetBackupMemberCountOk

`func (o *BackupRestoreDownloadResponseModel) GetBackupMemberCountOk() (*int32, bool)`

GetBackupMemberCountOk returns a tuple with the BackupMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMemberCount

`func (o *BackupRestoreDownloadResponseModel) SetBackupMemberCount(v int32)`

SetBackupMemberCount sets BackupMemberCount field to given value.

### HasBackupMemberCount

`func (o *BackupRestoreDownloadResponseModel) HasBackupMemberCount() bool`

HasBackupMemberCount returns a boolean if a field has been set.

### GetContinuationToken

`func (o *BackupRestoreDownloadResponseModel) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *BackupRestoreDownloadResponseModel) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *BackupRestoreDownloadResponseModel) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *BackupRestoreDownloadResponseModel) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *BackupRestoreDownloadResponseModel) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *BackupRestoreDownloadResponseModel) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


