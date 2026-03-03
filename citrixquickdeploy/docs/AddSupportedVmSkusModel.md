# AddSupportedVmSkusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]SupportedAzureVmSku**](SupportedAzureVmSku.md) | List of Azure VM SKUs to add support for | [optional] 

## Methods

### NewAddSupportedVmSkusModel

`func NewAddSupportedVmSkusModel() *AddSupportedVmSkusModel`

NewAddSupportedVmSkusModel instantiates a new AddSupportedVmSkusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSupportedVmSkusModelWithDefaults

`func NewAddSupportedVmSkusModelWithDefaults() *AddSupportedVmSkusModel`

NewAddSupportedVmSkusModelWithDefaults instantiates a new AddSupportedVmSkusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AddSupportedVmSkusModel) GetItems() []SupportedAzureVmSku`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AddSupportedVmSkusModel) GetItemsOk() (*[]SupportedAzureVmSku, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AddSupportedVmSkusModel) SetItems(v []SupportedAzureVmSku)`

SetItems sets Items field to given value.

### HasItems

`func (o *AddSupportedVmSkusModel) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *AddSupportedVmSkusModel) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *AddSupportedVmSkusModel) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


