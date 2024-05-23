# GetAllDiscoveryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]DiscoveryRecordModel**](DiscoveryRecordModel.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAllDiscoveryResponse

`func NewGetAllDiscoveryResponse() *GetAllDiscoveryResponse`

NewGetAllDiscoveryResponse instantiates a new GetAllDiscoveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllDiscoveryResponseWithDefaults

`func NewGetAllDiscoveryResponseWithDefaults() *GetAllDiscoveryResponse`

NewGetAllDiscoveryResponseWithDefaults instantiates a new GetAllDiscoveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetAllDiscoveryResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetAllDiscoveryResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetAllDiscoveryResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetAllDiscoveryResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *GetAllDiscoveryResponse) GetItems() []DiscoveryRecordModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetAllDiscoveryResponse) GetItemsOk() (*[]DiscoveryRecordModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetAllDiscoveryResponse) SetItems(v []DiscoveryRecordModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetAllDiscoveryResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *GetAllDiscoveryResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetAllDiscoveryResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetAllDiscoveryResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetAllDiscoveryResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


