# GetAllProfileResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]ProfileResponseModel**](ProfileResponseModel.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAllProfileResponseModel

`func NewGetAllProfileResponseModel() *GetAllProfileResponseModel`

NewGetAllProfileResponseModel instantiates a new GetAllProfileResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllProfileResponseModelWithDefaults

`func NewGetAllProfileResponseModelWithDefaults() *GetAllProfileResponseModel`

NewGetAllProfileResponseModelWithDefaults instantiates a new GetAllProfileResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetAllProfileResponseModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetAllProfileResponseModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetAllProfileResponseModel) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetAllProfileResponseModel) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *GetAllProfileResponseModel) GetItems() []ProfileResponseModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetAllProfileResponseModel) GetItemsOk() (*[]ProfileResponseModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetAllProfileResponseModel) SetItems(v []ProfileResponseModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetAllProfileResponseModel) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *GetAllProfileResponseModel) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetAllProfileResponseModel) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetAllProfileResponseModel) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetAllProfileResponseModel) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


