# AzureStorageFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageAccount** | **string** | Name of the storage account the file is located within | 
**Name** | **string** | Name of the file | 
**Uri** | **string** | Path of the file within the storage account | 

## Methods

### NewAzureStorageFile

`func NewAzureStorageFile(storageAccount string, name string, uri string, ) *AzureStorageFile`

NewAzureStorageFile instantiates a new AzureStorageFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageFileWithDefaults

`func NewAzureStorageFileWithDefaults() *AzureStorageFile`

NewAzureStorageFileWithDefaults instantiates a new AzureStorageFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageAccount

`func (o *AzureStorageFile) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AzureStorageFile) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AzureStorageFile) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.


### GetName

`func (o *AzureStorageFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureStorageFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureStorageFile) SetName(v string)`

SetName sets Name field to given value.


### GetUri

`func (o *AzureStorageFile) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AzureStorageFile) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AzureStorageFile) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


