# AzureVmOperationInputModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmConfig** | Pointer to [**AzureVMBase**](AzureVMBase.md) |  | [optional] 
**VmOperation** | Pointer to [**AzureVmOperation**](AzureVmOperation.md) |  | [optional] 

## Methods

### NewAzureVmOperationInputModel

`func NewAzureVmOperationInputModel() *AzureVmOperationInputModel`

NewAzureVmOperationInputModel instantiates a new AzureVmOperationInputModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVmOperationInputModelWithDefaults

`func NewAzureVmOperationInputModelWithDefaults() *AzureVmOperationInputModel`

NewAzureVmOperationInputModelWithDefaults instantiates a new AzureVmOperationInputModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmConfig

`func (o *AzureVmOperationInputModel) GetVmConfig() AzureVMBase`

GetVmConfig returns the VmConfig field if non-nil, zero value otherwise.

### GetVmConfigOk

`func (o *AzureVmOperationInputModel) GetVmConfigOk() (*AzureVMBase, bool)`

GetVmConfigOk returns a tuple with the VmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmConfig

`func (o *AzureVmOperationInputModel) SetVmConfig(v AzureVMBase)`

SetVmConfig sets VmConfig field to given value.

### HasVmConfig

`func (o *AzureVmOperationInputModel) HasVmConfig() bool`

HasVmConfig returns a boolean if a field has been set.

### GetVmOperation

`func (o *AzureVmOperationInputModel) GetVmOperation() AzureVmOperation`

GetVmOperation returns the VmOperation field if non-nil, zero value otherwise.

### GetVmOperationOk

`func (o *AzureVmOperationInputModel) GetVmOperationOk() (*AzureVmOperation, bool)`

GetVmOperationOk returns a tuple with the VmOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmOperation

`func (o *AzureVmOperationInputModel) SetVmOperation(v AzureVmOperation)`

SetVmOperation sets VmOperation field to given value.

### HasVmOperation

`func (o *AzureVmOperationInputModel) HasVmOperation() bool`

HasVmOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


