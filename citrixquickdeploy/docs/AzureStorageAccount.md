# AzureStorageAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | Azure subscription the storage account is associated with | [optional] 
**ResourceGroup** | **string** | Name of the Resource Group the Storage Account is associated with | 
**Location** | **string** | The location of the storage account | 
**Name** | **string** | Name of the storage account | 
**IsPremiumStorage** | Pointer to **bool** | Indicates if the storage account is using premium storage | [optional] 
**AccountKey** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureStorageAccount

`func NewAzureStorageAccount(resourceGroup string, location string, name string, ) *AzureStorageAccount`

NewAzureStorageAccount instantiates a new AzureStorageAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageAccountWithDefaults

`func NewAzureStorageAccountWithDefaults() *AzureStorageAccount`

NewAzureStorageAccountWithDefaults instantiates a new AzureStorageAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AzureStorageAccount) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureStorageAccount) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureStorageAccount) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureStorageAccount) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetResourceGroup

`func (o *AzureStorageAccount) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureStorageAccount) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureStorageAccount) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetLocation

`func (o *AzureStorageAccount) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureStorageAccount) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureStorageAccount) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *AzureStorageAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureStorageAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureStorageAccount) SetName(v string)`

SetName sets Name field to given value.


### GetIsPremiumStorage

`func (o *AzureStorageAccount) GetIsPremiumStorage() bool`

GetIsPremiumStorage returns the IsPremiumStorage field if non-nil, zero value otherwise.

### GetIsPremiumStorageOk

`func (o *AzureStorageAccount) GetIsPremiumStorageOk() (*bool, bool)`

GetIsPremiumStorageOk returns a tuple with the IsPremiumStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremiumStorage

`func (o *AzureStorageAccount) SetIsPremiumStorage(v bool)`

SetIsPremiumStorage sets IsPremiumStorage field to given value.

### HasIsPremiumStorage

`func (o *AzureStorageAccount) HasIsPremiumStorage() bool`

HasIsPremiumStorage returns a boolean if a field has been set.

### GetAccountKey

`func (o *AzureStorageAccount) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *AzureStorageAccount) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *AzureStorageAccount) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.

### HasAccountKey

`func (o *AzureStorageAccount) HasAccountKey() bool`

HasAccountKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


