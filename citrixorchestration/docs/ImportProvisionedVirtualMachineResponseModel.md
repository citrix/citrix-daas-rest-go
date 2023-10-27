# ImportProvisionedVirtualMachineResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ADAccountSid** | Pointer to **NullableString** | The Sid of the virtual machine. | [optional] 
**VMId** | Pointer to **NullableString** | The virtual machine Id. | [optional] 
**VMName** | Pointer to **NullableString** | The virtual machine name. | [optional] 
**FailureReason** | Pointer to **NullableString** | Failure reason. | [optional] 
**Status** | Pointer to [**ImportMachineStatus**](ImportMachineStatus.md) |  | [optional] 

## Methods

### NewImportProvisionedVirtualMachineResponseModel

`func NewImportProvisionedVirtualMachineResponseModel() *ImportProvisionedVirtualMachineResponseModel`

NewImportProvisionedVirtualMachineResponseModel instantiates a new ImportProvisionedVirtualMachineResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportProvisionedVirtualMachineResponseModelWithDefaults

`func NewImportProvisionedVirtualMachineResponseModelWithDefaults() *ImportProvisionedVirtualMachineResponseModel`

NewImportProvisionedVirtualMachineResponseModelWithDefaults instantiates a new ImportProvisionedVirtualMachineResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetADAccountSid

`func (o *ImportProvisionedVirtualMachineResponseModel) GetADAccountSid() string`

GetADAccountSid returns the ADAccountSid field if non-nil, zero value otherwise.

### GetADAccountSidOk

`func (o *ImportProvisionedVirtualMachineResponseModel) GetADAccountSidOk() (*string, bool)`

GetADAccountSidOk returns a tuple with the ADAccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADAccountSid

`func (o *ImportProvisionedVirtualMachineResponseModel) SetADAccountSid(v string)`

SetADAccountSid sets ADAccountSid field to given value.

### HasADAccountSid

`func (o *ImportProvisionedVirtualMachineResponseModel) HasADAccountSid() bool`

HasADAccountSid returns a boolean if a field has been set.

### SetADAccountSidNil

`func (o *ImportProvisionedVirtualMachineResponseModel) SetADAccountSidNil(b bool)`

 SetADAccountSidNil sets the value for ADAccountSid to be an explicit nil

### UnsetADAccountSid
`func (o *ImportProvisionedVirtualMachineResponseModel) UnsetADAccountSid()`

UnsetADAccountSid ensures that no value is present for ADAccountSid, not even an explicit nil
### GetVMId

`func (o *ImportProvisionedVirtualMachineResponseModel) GetVMId() string`

GetVMId returns the VMId field if non-nil, zero value otherwise.

### GetVMIdOk

`func (o *ImportProvisionedVirtualMachineResponseModel) GetVMIdOk() (*string, bool)`

GetVMIdOk returns a tuple with the VMId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMId

`func (o *ImportProvisionedVirtualMachineResponseModel) SetVMId(v string)`

SetVMId sets VMId field to given value.

### HasVMId

`func (o *ImportProvisionedVirtualMachineResponseModel) HasVMId() bool`

HasVMId returns a boolean if a field has been set.

### SetVMIdNil

`func (o *ImportProvisionedVirtualMachineResponseModel) SetVMIdNil(b bool)`

 SetVMIdNil sets the value for VMId to be an explicit nil

### UnsetVMId
`func (o *ImportProvisionedVirtualMachineResponseModel) UnsetVMId()`

UnsetVMId ensures that no value is present for VMId, not even an explicit nil
### GetVMName

`func (o *ImportProvisionedVirtualMachineResponseModel) GetVMName() string`

GetVMName returns the VMName field if non-nil, zero value otherwise.

### GetVMNameOk

`func (o *ImportProvisionedVirtualMachineResponseModel) GetVMNameOk() (*string, bool)`

GetVMNameOk returns a tuple with the VMName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMName

`func (o *ImportProvisionedVirtualMachineResponseModel) SetVMName(v string)`

SetVMName sets VMName field to given value.

### HasVMName

`func (o *ImportProvisionedVirtualMachineResponseModel) HasVMName() bool`

HasVMName returns a boolean if a field has been set.

### SetVMNameNil

`func (o *ImportProvisionedVirtualMachineResponseModel) SetVMNameNil(b bool)`

 SetVMNameNil sets the value for VMName to be an explicit nil

### UnsetVMName
`func (o *ImportProvisionedVirtualMachineResponseModel) UnsetVMName()`

UnsetVMName ensures that no value is present for VMName, not even an explicit nil
### GetFailureReason

`func (o *ImportProvisionedVirtualMachineResponseModel) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ImportProvisionedVirtualMachineResponseModel) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ImportProvisionedVirtualMachineResponseModel) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ImportProvisionedVirtualMachineResponseModel) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *ImportProvisionedVirtualMachineResponseModel) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *ImportProvisionedVirtualMachineResponseModel) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetStatus

`func (o *ImportProvisionedVirtualMachineResponseModel) GetStatus() ImportMachineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportProvisionedVirtualMachineResponseModel) GetStatusOk() (*ImportMachineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportProvisionedVirtualMachineResponseModel) SetStatus(v ImportMachineStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImportProvisionedVirtualMachineResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


