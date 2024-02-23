# ProvisionedVirtualMachineDetailsWarningReponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ProvisionedVirtualMachineDetailsWarningType**](ProvisionedVirtualMachineDetailsWarningType.md) |  | [optional] 
**Message** | Pointer to **NullableString** | Message associated with warning | [optional] 

## Methods

### NewProvisionedVirtualMachineDetailsWarningReponseModel

`func NewProvisionedVirtualMachineDetailsWarningReponseModel() *ProvisionedVirtualMachineDetailsWarningReponseModel`

NewProvisionedVirtualMachineDetailsWarningReponseModel instantiates a new ProvisionedVirtualMachineDetailsWarningReponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedVirtualMachineDetailsWarningReponseModelWithDefaults

`func NewProvisionedVirtualMachineDetailsWarningReponseModelWithDefaults() *ProvisionedVirtualMachineDetailsWarningReponseModel`

NewProvisionedVirtualMachineDetailsWarningReponseModelWithDefaults instantiates a new ProvisionedVirtualMachineDetailsWarningReponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProvisionedVirtualMachineDetailsWarningReponseModel) GetType() ProvisionedVirtualMachineDetailsWarningType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProvisionedVirtualMachineDetailsWarningReponseModel) GetTypeOk() (*ProvisionedVirtualMachineDetailsWarningType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProvisionedVirtualMachineDetailsWarningReponseModel) SetType(v ProvisionedVirtualMachineDetailsWarningType)`

SetType sets Type field to given value.

### HasType

`func (o *ProvisionedVirtualMachineDetailsWarningReponseModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *ProvisionedVirtualMachineDetailsWarningReponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProvisionedVirtualMachineDetailsWarningReponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProvisionedVirtualMachineDetailsWarningReponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProvisionedVirtualMachineDetailsWarningReponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ProvisionedVirtualMachineDetailsWarningReponseModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProvisionedVirtualMachineDetailsWarningReponseModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


