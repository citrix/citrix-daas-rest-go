# DeleteVdaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineUids** | Pointer to **[]int32** | List of machine Uids to delete | [optional] 
**ServiceAccount** | Pointer to **string** | Service account to perform delete with | [optional] 
**ServiceAccountPassword** | Pointer to **string** | Customer&#39;s domain password | [optional] 

## Methods

### NewDeleteVdaModel

`func NewDeleteVdaModel() *DeleteVdaModel`

NewDeleteVdaModel instantiates a new DeleteVdaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVdaModelWithDefaults

`func NewDeleteVdaModelWithDefaults() *DeleteVdaModel`

NewDeleteVdaModelWithDefaults instantiates a new DeleteVdaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineUids

`func (o *DeleteVdaModel) GetMachineUids() []int32`

GetMachineUids returns the MachineUids field if non-nil, zero value otherwise.

### GetMachineUidsOk

`func (o *DeleteVdaModel) GetMachineUidsOk() (*[]int32, bool)`

GetMachineUidsOk returns a tuple with the MachineUids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineUids

`func (o *DeleteVdaModel) SetMachineUids(v []int32)`

SetMachineUids sets MachineUids field to given value.

### HasMachineUids

`func (o *DeleteVdaModel) HasMachineUids() bool`

HasMachineUids returns a boolean if a field has been set.

### GetServiceAccount

`func (o *DeleteVdaModel) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *DeleteVdaModel) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *DeleteVdaModel) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *DeleteVdaModel) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetServiceAccountPassword

`func (o *DeleteVdaModel) GetServiceAccountPassword() string`

GetServiceAccountPassword returns the ServiceAccountPassword field if non-nil, zero value otherwise.

### GetServiceAccountPasswordOk

`func (o *DeleteVdaModel) GetServiceAccountPasswordOk() (*string, bool)`

GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountPassword

`func (o *DeleteVdaModel) SetServiceAccountPassword(v string)`

SetServiceAccountPassword sets ServiceAccountPassword field to given value.

### HasServiceAccountPassword

`func (o *DeleteVdaModel) HasServiceAccountPassword() bool`

HasServiceAccountPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


