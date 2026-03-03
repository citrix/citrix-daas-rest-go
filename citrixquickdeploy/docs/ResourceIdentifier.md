# ResourceIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**Parent** | Pointer to [**NullableResourceIdentifier**](ResourceIdentifier.md) |  | [optional] [readonly] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] [readonly] 
**Provider** | Pointer to **NullableString** |  | [optional] [readonly] 
**Location** | Pointer to [**NullableAzureLocation**](AzureLocation.md) |  | [optional] [readonly] 
**ResourceGroupName** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewResourceIdentifier

`func NewResourceIdentifier() *ResourceIdentifier`

NewResourceIdentifier instantiates a new ResourceIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceIdentifierWithDefaults

`func NewResourceIdentifierWithDefaults() *ResourceIdentifier`

NewResourceIdentifierWithDefaults instantiates a new ResourceIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *ResourceIdentifier) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceIdentifier) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceIdentifier) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceIdentifier) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetName

`func (o *ResourceIdentifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceIdentifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceIdentifier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceIdentifier) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ResourceIdentifier) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ResourceIdentifier) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParent

`func (o *ResourceIdentifier) GetParent() ResourceIdentifier`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ResourceIdentifier) GetParentOk() (*ResourceIdentifier, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ResourceIdentifier) SetParent(v ResourceIdentifier)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ResourceIdentifier) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *ResourceIdentifier) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *ResourceIdentifier) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetSubscriptionId

`func (o *ResourceIdentifier) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ResourceIdentifier) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ResourceIdentifier) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ResourceIdentifier) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *ResourceIdentifier) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *ResourceIdentifier) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetProvider

`func (o *ResourceIdentifier) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ResourceIdentifier) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ResourceIdentifier) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ResourceIdentifier) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *ResourceIdentifier) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *ResourceIdentifier) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetLocation

`func (o *ResourceIdentifier) GetLocation() AzureLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ResourceIdentifier) GetLocationOk() (*AzureLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ResourceIdentifier) SetLocation(v AzureLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ResourceIdentifier) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ResourceIdentifier) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ResourceIdentifier) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetResourceGroupName

`func (o *ResourceIdentifier) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ResourceIdentifier) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ResourceIdentifier) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *ResourceIdentifier) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *ResourceIdentifier) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *ResourceIdentifier) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


