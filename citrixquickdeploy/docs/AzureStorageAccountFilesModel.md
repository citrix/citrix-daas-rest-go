# AzureStorageAccountFilesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageAccount** | **string** | Name of the Storage Account | 
**Items** | [**[]AzureStorageFile**](AzureStorageFile.md) | Files within the account | 
**Files** | Pointer to [**[]AzureStorageFile**](AzureStorageFile.md) | Alias of Items property for backward compatibility | [optional] 

## Methods

### NewAzureStorageAccountFilesModel

`func NewAzureStorageAccountFilesModel(storageAccount string, items []AzureStorageFile, ) *AzureStorageAccountFilesModel`

NewAzureStorageAccountFilesModel instantiates a new AzureStorageAccountFilesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageAccountFilesModelWithDefaults

`func NewAzureStorageAccountFilesModelWithDefaults() *AzureStorageAccountFilesModel`

NewAzureStorageAccountFilesModelWithDefaults instantiates a new AzureStorageAccountFilesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageAccount

`func (o *AzureStorageAccountFilesModel) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AzureStorageAccountFilesModel) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AzureStorageAccountFilesModel) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.


### GetItems

`func (o *AzureStorageAccountFilesModel) GetItems() []AzureStorageFile`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AzureStorageAccountFilesModel) GetItemsOk() (*[]AzureStorageFile, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AzureStorageAccountFilesModel) SetItems(v []AzureStorageFile)`

SetItems sets Items field to given value.


### GetFiles

`func (o *AzureStorageAccountFilesModel) GetFiles() []AzureStorageFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *AzureStorageAccountFilesModel) GetFilesOk() (*[]AzureStorageFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *AzureStorageAccountFilesModel) SetFiles(v []AzureStorageFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *AzureStorageAccountFilesModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


