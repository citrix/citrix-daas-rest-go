# ApplicationDiscoveryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to **string** | Serialized string that is a list of machine shares, or a list of directories, or application properties that parsed from a selected executable file. | [optional] 
**Type** | Pointer to [**ApplicationDiscoveryItemType**](ApplicationDiscoveryItemType.md) |  | [optional] 

## Methods

### NewApplicationDiscoveryResponseModel

`func NewApplicationDiscoveryResponseModel() *ApplicationDiscoveryResponseModel`

NewApplicationDiscoveryResponseModel instantiates a new ApplicationDiscoveryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDiscoveryResponseModelWithDefaults

`func NewApplicationDiscoveryResponseModelWithDefaults() *ApplicationDiscoveryResponseModel`

NewApplicationDiscoveryResponseModelWithDefaults instantiates a new ApplicationDiscoveryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *ApplicationDiscoveryResponseModel) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ApplicationDiscoveryResponseModel) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ApplicationDiscoveryResponseModel) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *ApplicationDiscoveryResponseModel) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetType

`func (o *ApplicationDiscoveryResponseModel) GetType() ApplicationDiscoveryItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationDiscoveryResponseModel) GetTypeOk() (*ApplicationDiscoveryItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationDiscoveryResponseModel) SetType(v ApplicationDiscoveryItemType)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationDiscoveryResponseModel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


