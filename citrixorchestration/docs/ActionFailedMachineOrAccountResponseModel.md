# ActionFailedMachineOrAccountResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Failed machine or account. | [optional] 
**ActionError** | Pointer to [**ActionError**](ActionError.md) |  | [optional] 

## Methods

### NewActionFailedMachineOrAccountResponseModel

`func NewActionFailedMachineOrAccountResponseModel() *ActionFailedMachineOrAccountResponseModel`

NewActionFailedMachineOrAccountResponseModel instantiates a new ActionFailedMachineOrAccountResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionFailedMachineOrAccountResponseModelWithDefaults

`func NewActionFailedMachineOrAccountResponseModelWithDefaults() *ActionFailedMachineOrAccountResponseModel`

NewActionFailedMachineOrAccountResponseModelWithDefaults instantiates a new ActionFailedMachineOrAccountResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ActionFailedMachineOrAccountResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionFailedMachineOrAccountResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionFailedMachineOrAccountResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionFailedMachineOrAccountResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ActionFailedMachineOrAccountResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ActionFailedMachineOrAccountResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetActionError

`func (o *ActionFailedMachineOrAccountResponseModel) GetActionError() ActionError`

GetActionError returns the ActionError field if non-nil, zero value otherwise.

### GetActionErrorOk

`func (o *ActionFailedMachineOrAccountResponseModel) GetActionErrorOk() (*ActionError, bool)`

GetActionErrorOk returns a tuple with the ActionError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionError

`func (o *ActionFailedMachineOrAccountResponseModel) SetActionError(v ActionError)`

SetActionError sets ActionError field to given value.

### HasActionError

`func (o *ActionFailedMachineOrAccountResponseModel) HasActionError() bool`

HasActionError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


