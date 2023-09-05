# CollectionModelOfBackupRestoreHistoryInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]BackupRestoreHistoryInformation**](BackupRestoreHistoryInformation.md) | List of items. | 
**ContinuationToken** | Pointer to **string** | If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter. | [optional] 
**TotalItems** | Pointer to **int32** | Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to &#x60;$search&#x60; APIs. | [optional] 

## Methods

### NewCollectionModelOfBackupRestoreHistoryInformation

`func NewCollectionModelOfBackupRestoreHistoryInformation(items []BackupRestoreHistoryInformation, ) *CollectionModelOfBackupRestoreHistoryInformation`

NewCollectionModelOfBackupRestoreHistoryInformation instantiates a new CollectionModelOfBackupRestoreHistoryInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionModelOfBackupRestoreHistoryInformationWithDefaults

`func NewCollectionModelOfBackupRestoreHistoryInformationWithDefaults() *CollectionModelOfBackupRestoreHistoryInformation`

NewCollectionModelOfBackupRestoreHistoryInformationWithDefaults instantiates a new CollectionModelOfBackupRestoreHistoryInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CollectionModelOfBackupRestoreHistoryInformation) GetItems() []BackupRestoreHistoryInformation`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CollectionModelOfBackupRestoreHistoryInformation) GetItemsOk() (*[]BackupRestoreHistoryInformation, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CollectionModelOfBackupRestoreHistoryInformation) SetItems(v []BackupRestoreHistoryInformation)`

SetItems sets Items field to given value.


### GetContinuationToken

`func (o *CollectionModelOfBackupRestoreHistoryInformation) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *CollectionModelOfBackupRestoreHistoryInformation) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *CollectionModelOfBackupRestoreHistoryInformation) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *CollectionModelOfBackupRestoreHistoryInformation) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetTotalItems

`func (o *CollectionModelOfBackupRestoreHistoryInformation) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *CollectionModelOfBackupRestoreHistoryInformation) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *CollectionModelOfBackupRestoreHistoryInformation) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *CollectionModelOfBackupRestoreHistoryInformation) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


