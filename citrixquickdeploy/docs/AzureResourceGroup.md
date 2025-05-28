# AzureResourceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | ID of the subscription the Resource Group is associated with | 
**Name** | **string** | Name of the Resource Group | 
**Location** | **string** | Region where resources within this group are located | 
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAzureResourceGroup

`func NewAzureResourceGroup(subscriptionId string, name string, location string, ) *AzureResourceGroup`

NewAzureResourceGroup instantiates a new AzureResourceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureResourceGroupWithDefaults

`func NewAzureResourceGroupWithDefaults() *AzureResourceGroup`

NewAzureResourceGroupWithDefaults instantiates a new AzureResourceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AzureResourceGroup) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureResourceGroup) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureResourceGroup) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetName

`func (o *AzureResourceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureResourceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureResourceGroup) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *AzureResourceGroup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureResourceGroup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureResourceGroup) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetId

`func (o *AzureResourceGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureResourceGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureResourceGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureResourceGroup) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


