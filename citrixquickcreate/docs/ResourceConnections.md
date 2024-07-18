# ResourceConnections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AwsEdcDirectoryConnection**](AwsEdcDirectoryConnection.md) | Enumerable of HostingUnit | [optional] 

## Methods

### NewResourceConnections

`func NewResourceConnections() *ResourceConnections`

NewResourceConnections instantiates a new ResourceConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConnectionsWithDefaults

`func NewResourceConnectionsWithDefaults() *ResourceConnections`

NewResourceConnectionsWithDefaults instantiates a new ResourceConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ResourceConnections) GetItems() []AwsEdcDirectoryConnection`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResourceConnections) GetItemsOk() (*[]AwsEdcDirectoryConnection, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResourceConnections) SetItems(v []AwsEdcDirectoryConnection)`

SetItems sets Items field to given value.

### HasItems

`func (o *ResourceConnections) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *ResourceConnections) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *ResourceConnections) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


