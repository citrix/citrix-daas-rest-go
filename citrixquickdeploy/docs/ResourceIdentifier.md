# ResourceIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Parent** | Pointer to [**ResourceIdentifier**](ResourceIdentifier.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] [readonly] 
**Provider** | Pointer to **string** |  | [optional] [readonly] 
**Location** | Pointer to [**AzureLocation**](AzureLocation.md) |  | [optional] 
**ResourceGroupName** | Pointer to **string** |  | [optional] [readonly] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


