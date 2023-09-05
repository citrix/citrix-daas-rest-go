# SendMessageRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Style** | Pointer to [**MessageStyle**](MessageStyle.md) |  | [optional] 
**Title** | Pointer to **string** | Text to display in the message box title bar. | [optional] 
**Text** | Pointer to **string** | The message to display. | [optional] 

## Methods

### NewSendMessageRequestModel

`func NewSendMessageRequestModel() *SendMessageRequestModel`

NewSendMessageRequestModel instantiates a new SendMessageRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageRequestModelWithDefaults

`func NewSendMessageRequestModelWithDefaults() *SendMessageRequestModel`

NewSendMessageRequestModelWithDefaults instantiates a new SendMessageRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStyle

`func (o *SendMessageRequestModel) GetStyle() MessageStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *SendMessageRequestModel) GetStyleOk() (*MessageStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *SendMessageRequestModel) SetStyle(v MessageStyle)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *SendMessageRequestModel) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetTitle

`func (o *SendMessageRequestModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendMessageRequestModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendMessageRequestModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SendMessageRequestModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetText

`func (o *SendMessageRequestModel) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SendMessageRequestModel) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SendMessageRequestModel) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SendMessageRequestModel) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


