# AzureSubscriptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AzureSubscriptionOverview**](AzureSubscriptionOverview.md) | List of known subscriptions that have been associated with the customer | 
**Subscriptions** | Pointer to [**[]AzureSubscriptionOverview**](AzureSubscriptionOverview.md) | Alias of Items property for backward compatibility | [optional] 
**MaxCitrixManagedSubscriptions** | Pointer to **int32** | The maximum allowed Citrix Managed subscriptions for the customer | [optional] 
**UniqueName** | Pointer to **string** | The unique name of the user who provided the auth code to list subscriptions. Will only be present if an Azure Authorization Code was provided | [optional] 
**StaleData** | Pointer to **bool** |  | [optional] 

## Methods

### NewAzureSubscriptionsModel

`func NewAzureSubscriptionsModel(items []AzureSubscriptionOverview, ) *AzureSubscriptionsModel`

NewAzureSubscriptionsModel instantiates a new AzureSubscriptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSubscriptionsModelWithDefaults

`func NewAzureSubscriptionsModelWithDefaults() *AzureSubscriptionsModel`

NewAzureSubscriptionsModelWithDefaults instantiates a new AzureSubscriptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AzureSubscriptionsModel) GetItems() []AzureSubscriptionOverview`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AzureSubscriptionsModel) GetItemsOk() (*[]AzureSubscriptionOverview, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AzureSubscriptionsModel) SetItems(v []AzureSubscriptionOverview)`

SetItems sets Items field to given value.


### GetSubscriptions

`func (o *AzureSubscriptionsModel) GetSubscriptions() []AzureSubscriptionOverview`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AzureSubscriptionsModel) GetSubscriptionsOk() (*[]AzureSubscriptionOverview, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AzureSubscriptionsModel) SetSubscriptions(v []AzureSubscriptionOverview)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *AzureSubscriptionsModel) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetMaxCitrixManagedSubscriptions

`func (o *AzureSubscriptionsModel) GetMaxCitrixManagedSubscriptions() int32`

GetMaxCitrixManagedSubscriptions returns the MaxCitrixManagedSubscriptions field if non-nil, zero value otherwise.

### GetMaxCitrixManagedSubscriptionsOk

`func (o *AzureSubscriptionsModel) GetMaxCitrixManagedSubscriptionsOk() (*int32, bool)`

GetMaxCitrixManagedSubscriptionsOk returns a tuple with the MaxCitrixManagedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCitrixManagedSubscriptions

`func (o *AzureSubscriptionsModel) SetMaxCitrixManagedSubscriptions(v int32)`

SetMaxCitrixManagedSubscriptions sets MaxCitrixManagedSubscriptions field to given value.

### HasMaxCitrixManagedSubscriptions

`func (o *AzureSubscriptionsModel) HasMaxCitrixManagedSubscriptions() bool`

HasMaxCitrixManagedSubscriptions returns a boolean if a field has been set.

### GetUniqueName

`func (o *AzureSubscriptionsModel) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *AzureSubscriptionsModel) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *AzureSubscriptionsModel) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *AzureSubscriptionsModel) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetStaleData

`func (o *AzureSubscriptionsModel) GetStaleData() bool`

GetStaleData returns the StaleData field if non-nil, zero value otherwise.

### GetStaleDataOk

`func (o *AzureSubscriptionsModel) GetStaleDataOk() (*bool, bool)`

GetStaleDataOk returns a tuple with the StaleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleData

`func (o *AzureSubscriptionsModel) SetStaleData(v bool)`

SetStaleData sets StaleData field to given value.

### HasStaleData

`func (o *AzureSubscriptionsModel) HasStaleData() bool`

HasStaleData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


