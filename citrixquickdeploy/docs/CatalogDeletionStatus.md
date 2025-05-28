# CatalogDeletionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeedsServiceAccountCredentials** | Pointer to **bool** | Indicates if service account credentials are needed | [optional] 
**PromptForResourceLocationDeletion** | Pointer to **bool** | Indicates if the catalog&#39;s resource location could be deleted with the catalog | [optional] 

## Methods

### NewCatalogDeletionStatus

`func NewCatalogDeletionStatus() *CatalogDeletionStatus`

NewCatalogDeletionStatus instantiates a new CatalogDeletionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogDeletionStatusWithDefaults

`func NewCatalogDeletionStatusWithDefaults() *CatalogDeletionStatus`

NewCatalogDeletionStatusWithDefaults instantiates a new CatalogDeletionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeedsServiceAccountCredentials

`func (o *CatalogDeletionStatus) GetNeedsServiceAccountCredentials() bool`

GetNeedsServiceAccountCredentials returns the NeedsServiceAccountCredentials field if non-nil, zero value otherwise.

### GetNeedsServiceAccountCredentialsOk

`func (o *CatalogDeletionStatus) GetNeedsServiceAccountCredentialsOk() (*bool, bool)`

GetNeedsServiceAccountCredentialsOk returns a tuple with the NeedsServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsServiceAccountCredentials

`func (o *CatalogDeletionStatus) SetNeedsServiceAccountCredentials(v bool)`

SetNeedsServiceAccountCredentials sets NeedsServiceAccountCredentials field to given value.

### HasNeedsServiceAccountCredentials

`func (o *CatalogDeletionStatus) HasNeedsServiceAccountCredentials() bool`

HasNeedsServiceAccountCredentials returns a boolean if a field has been set.

### GetPromptForResourceLocationDeletion

`func (o *CatalogDeletionStatus) GetPromptForResourceLocationDeletion() bool`

GetPromptForResourceLocationDeletion returns the PromptForResourceLocationDeletion field if non-nil, zero value otherwise.

### GetPromptForResourceLocationDeletionOk

`func (o *CatalogDeletionStatus) GetPromptForResourceLocationDeletionOk() (*bool, bool)`

GetPromptForResourceLocationDeletionOk returns a tuple with the PromptForResourceLocationDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptForResourceLocationDeletion

`func (o *CatalogDeletionStatus) SetPromptForResourceLocationDeletion(v bool)`

SetPromptForResourceLocationDeletion sets PromptForResourceLocationDeletion field to given value.

### HasPromptForResourceLocationDeletion

`func (o *CatalogDeletionStatus) HasPromptForResourceLocationDeletion() bool`

HasPromptForResourceLocationDeletion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


