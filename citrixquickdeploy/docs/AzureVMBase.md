# AzureVMBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroupName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to [**CatalogResourceType**](CatalogResourceType.md) |  | [optional] 

## Methods

### NewAzureVMBase

`func NewAzureVMBase() *AzureVMBase`

NewAzureVMBase instantiates a new AzureVMBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVMBaseWithDefaults

`func NewAzureVMBaseWithDefaults() *AzureVMBase`

NewAzureVMBaseWithDefaults instantiates a new AzureVMBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroupName

`func (o *AzureVMBase) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureVMBase) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureVMBase) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureVMBase) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetName

`func (o *AzureVMBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureVMBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureVMBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureVMBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceType

`func (o *AzureVMBase) GetResourceType() CatalogResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AzureVMBase) GetResourceTypeOk() (*CatalogResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AzureVMBase) SetResourceType(v CatalogResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AzureVMBase) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


