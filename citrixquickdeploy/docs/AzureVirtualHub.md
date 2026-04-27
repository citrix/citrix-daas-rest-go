# AzureVirtualHub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | ID of the Azure Subscription where the Virtual Hub is configured | 
**ResourceGroup** | **string** | Name of the Resource Group the Virtual Hub is associated with | 
**Name** | **string** | Name of the Virtual Hub | 
**ResourceId** | Pointer to **NullableString** | Azure resource ID of the Virtual Hub | [optional] 
**Location** | Pointer to **NullableString** | Azure region where the Virtual Hub is located | [optional] 
**VirtualWanId** | Pointer to **NullableString** | Resource ID of the parent Virtual WAN | [optional] 

## Methods

### NewAzureVirtualHub

`func NewAzureVirtualHub(subscriptionId string, resourceGroup string, name string, ) *AzureVirtualHub`

NewAzureVirtualHub instantiates a new AzureVirtualHub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVirtualHubWithDefaults

`func NewAzureVirtualHubWithDefaults() *AzureVirtualHub`

NewAzureVirtualHubWithDefaults instantiates a new AzureVirtualHub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AzureVirtualHub) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureVirtualHub) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureVirtualHub) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroup

`func (o *AzureVirtualHub) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureVirtualHub) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureVirtualHub) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetName

`func (o *AzureVirtualHub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureVirtualHub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureVirtualHub) SetName(v string)`

SetName sets Name field to given value.


### GetResourceId

`func (o *AzureVirtualHub) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AzureVirtualHub) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AzureVirtualHub) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AzureVirtualHub) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AzureVirtualHub) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AzureVirtualHub) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetLocation

`func (o *AzureVirtualHub) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureVirtualHub) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureVirtualHub) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AzureVirtualHub) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AzureVirtualHub) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AzureVirtualHub) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetVirtualWanId

`func (o *AzureVirtualHub) GetVirtualWanId() string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *AzureVirtualHub) GetVirtualWanIdOk() (*string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *AzureVirtualHub) SetVirtualWanId(v string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *AzureVirtualHub) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.

### SetVirtualWanIdNil

`func (o *AzureVirtualHub) SetVirtualWanIdNil(b bool)`

 SetVirtualWanIdNil sets the value for VirtualWanId to be an explicit nil

### UnsetVirtualWanId
`func (o *AzureVirtualHub) UnsetVirtualWanId()`

UnsetVirtualWanId ensures that no value is present for VirtualWanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


