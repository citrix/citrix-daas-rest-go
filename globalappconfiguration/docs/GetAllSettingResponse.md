# GetAllSettingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]SettingsRecordModel**](SettingsRecordModel.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAllSettingResponse

`func NewGetAllSettingResponse() *GetAllSettingResponse`

NewGetAllSettingResponse instantiates a new GetAllSettingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllSettingResponseWithDefaults

`func NewGetAllSettingResponseWithDefaults() *GetAllSettingResponse`

NewGetAllSettingResponseWithDefaults instantiates a new GetAllSettingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetAllSettingResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetAllSettingResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetAllSettingResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetAllSettingResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *GetAllSettingResponse) GetItems() []SettingsRecordModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetAllSettingResponse) GetItemsOk() (*[]SettingsRecordModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetAllSettingResponse) SetItems(v []SettingsRecordModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetAllSettingResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *GetAllSettingResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetAllSettingResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetAllSettingResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetAllSettingResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


