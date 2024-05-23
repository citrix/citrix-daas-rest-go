# ImageVersionSpecWarningResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ImageVersionSpecWarningType**](ImageVersionSpecWarningType.md) |  | 
**Message** | Pointer to **NullableString** | Message associated with warning | [optional] 

## Methods

### NewImageVersionSpecWarningResponseModel

`func NewImageVersionSpecWarningResponseModel(type_ ImageVersionSpecWarningType, ) *ImageVersionSpecWarningResponseModel`

NewImageVersionSpecWarningResponseModel instantiates a new ImageVersionSpecWarningResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionSpecWarningResponseModelWithDefaults

`func NewImageVersionSpecWarningResponseModelWithDefaults() *ImageVersionSpecWarningResponseModel`

NewImageVersionSpecWarningResponseModelWithDefaults instantiates a new ImageVersionSpecWarningResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ImageVersionSpecWarningResponseModel) GetType() ImageVersionSpecWarningType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageVersionSpecWarningResponseModel) GetTypeOk() (*ImageVersionSpecWarningType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageVersionSpecWarningResponseModel) SetType(v ImageVersionSpecWarningType)`

SetType sets Type field to given value.


### GetMessage

`func (o *ImageVersionSpecWarningResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImageVersionSpecWarningResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImageVersionSpecWarningResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ImageVersionSpecWarningResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ImageVersionSpecWarningResponseModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ImageVersionSpecWarningResponseModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


