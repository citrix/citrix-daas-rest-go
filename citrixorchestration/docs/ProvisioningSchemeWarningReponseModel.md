# ProvisioningSchemeWarningReponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ProvSchemeWarningType**](ProvSchemeWarningType.md) |  | [optional] 
**Message** | Pointer to **NullableString** | Message associated with warning | [optional] 

## Methods

### NewProvisioningSchemeWarningReponseModel

`func NewProvisioningSchemeWarningReponseModel() *ProvisioningSchemeWarningReponseModel`

NewProvisioningSchemeWarningReponseModel instantiates a new ProvisioningSchemeWarningReponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeWarningReponseModelWithDefaults

`func NewProvisioningSchemeWarningReponseModelWithDefaults() *ProvisioningSchemeWarningReponseModel`

NewProvisioningSchemeWarningReponseModelWithDefaults instantiates a new ProvisioningSchemeWarningReponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProvisioningSchemeWarningReponseModel) GetType() ProvSchemeWarningType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProvisioningSchemeWarningReponseModel) GetTypeOk() (*ProvSchemeWarningType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProvisioningSchemeWarningReponseModel) SetType(v ProvSchemeWarningType)`

SetType sets Type field to given value.

### HasType

`func (o *ProvisioningSchemeWarningReponseModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *ProvisioningSchemeWarningReponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProvisioningSchemeWarningReponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProvisioningSchemeWarningReponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProvisioningSchemeWarningReponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ProvisioningSchemeWarningReponseModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProvisioningSchemeWarningReponseModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


