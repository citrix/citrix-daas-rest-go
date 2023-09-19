# MachineCatalogWarningResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MachineCatalogWarningType**](MachineCatalogWarningType.md) |  | 
**Subtype** | Pointer to [**MachineCatalogWarningSubtype**](MachineCatalogWarningSubtype.md) |  | [optional] 
**Message** | Pointer to **NullableString** | Message associated with warning | [optional] 

## Methods

### NewMachineCatalogWarningResponseModel

`func NewMachineCatalogWarningResponseModel(type_ MachineCatalogWarningType, ) *MachineCatalogWarningResponseModel`

NewMachineCatalogWarningResponseModel instantiates a new MachineCatalogWarningResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogWarningResponseModelWithDefaults

`func NewMachineCatalogWarningResponseModelWithDefaults() *MachineCatalogWarningResponseModel`

NewMachineCatalogWarningResponseModelWithDefaults instantiates a new MachineCatalogWarningResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MachineCatalogWarningResponseModel) GetType() MachineCatalogWarningType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MachineCatalogWarningResponseModel) GetTypeOk() (*MachineCatalogWarningType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MachineCatalogWarningResponseModel) SetType(v MachineCatalogWarningType)`

SetType sets Type field to given value.


### GetSubtype

`func (o *MachineCatalogWarningResponseModel) GetSubtype() MachineCatalogWarningSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MachineCatalogWarningResponseModel) GetSubtypeOk() (*MachineCatalogWarningSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MachineCatalogWarningResponseModel) SetSubtype(v MachineCatalogWarningSubtype)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *MachineCatalogWarningResponseModel) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetMessage

`func (o *MachineCatalogWarningResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MachineCatalogWarningResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MachineCatalogWarningResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MachineCatalogWarningResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MachineCatalogWarningResponseModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MachineCatalogWarningResponseModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


