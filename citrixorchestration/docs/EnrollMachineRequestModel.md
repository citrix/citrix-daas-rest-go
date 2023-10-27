# EnrollMachineRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdaInstanceId** | Pointer to **NullableString** | The id for the vda&#39;s key pair registered with the MFA trust service. | [optional] 
**VdaInstanceName** | Pointer to **NullableString** | The instance name for the key. it is also used as the machine&#39;s name for non-domain joined vda. | [optional] 
**ServicePublicKey** | Pointer to **NullableString** | The vda&#39;s public service key to be registered with FMA Trust Service. | [optional] 
**MachineSid** | Pointer to **NullableString** | Real sid of AD machine for domain joined; or virtual sid if non-domain joined this parameter will be used to create machine | [optional] 
**MachineMetadata** | Pointer to [**EnrollMachineMetaDataModel**](EnrollMachineMetaDataModel.md) |  | [optional] 

## Methods

### NewEnrollMachineRequestModel

`func NewEnrollMachineRequestModel() *EnrollMachineRequestModel`

NewEnrollMachineRequestModel instantiates a new EnrollMachineRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollMachineRequestModelWithDefaults

`func NewEnrollMachineRequestModelWithDefaults() *EnrollMachineRequestModel`

NewEnrollMachineRequestModelWithDefaults instantiates a new EnrollMachineRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdaInstanceId

`func (o *EnrollMachineRequestModel) GetVdaInstanceId() string`

GetVdaInstanceId returns the VdaInstanceId field if non-nil, zero value otherwise.

### GetVdaInstanceIdOk

`func (o *EnrollMachineRequestModel) GetVdaInstanceIdOk() (*string, bool)`

GetVdaInstanceIdOk returns a tuple with the VdaInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaInstanceId

`func (o *EnrollMachineRequestModel) SetVdaInstanceId(v string)`

SetVdaInstanceId sets VdaInstanceId field to given value.

### HasVdaInstanceId

`func (o *EnrollMachineRequestModel) HasVdaInstanceId() bool`

HasVdaInstanceId returns a boolean if a field has been set.

### SetVdaInstanceIdNil

`func (o *EnrollMachineRequestModel) SetVdaInstanceIdNil(b bool)`

 SetVdaInstanceIdNil sets the value for VdaInstanceId to be an explicit nil

### UnsetVdaInstanceId
`func (o *EnrollMachineRequestModel) UnsetVdaInstanceId()`

UnsetVdaInstanceId ensures that no value is present for VdaInstanceId, not even an explicit nil
### GetVdaInstanceName

`func (o *EnrollMachineRequestModel) GetVdaInstanceName() string`

GetVdaInstanceName returns the VdaInstanceName field if non-nil, zero value otherwise.

### GetVdaInstanceNameOk

`func (o *EnrollMachineRequestModel) GetVdaInstanceNameOk() (*string, bool)`

GetVdaInstanceNameOk returns a tuple with the VdaInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaInstanceName

`func (o *EnrollMachineRequestModel) SetVdaInstanceName(v string)`

SetVdaInstanceName sets VdaInstanceName field to given value.

### HasVdaInstanceName

`func (o *EnrollMachineRequestModel) HasVdaInstanceName() bool`

HasVdaInstanceName returns a boolean if a field has been set.

### SetVdaInstanceNameNil

`func (o *EnrollMachineRequestModel) SetVdaInstanceNameNil(b bool)`

 SetVdaInstanceNameNil sets the value for VdaInstanceName to be an explicit nil

### UnsetVdaInstanceName
`func (o *EnrollMachineRequestModel) UnsetVdaInstanceName()`

UnsetVdaInstanceName ensures that no value is present for VdaInstanceName, not even an explicit nil
### GetServicePublicKey

`func (o *EnrollMachineRequestModel) GetServicePublicKey() string`

GetServicePublicKey returns the ServicePublicKey field if non-nil, zero value otherwise.

### GetServicePublicKeyOk

`func (o *EnrollMachineRequestModel) GetServicePublicKeyOk() (*string, bool)`

GetServicePublicKeyOk returns a tuple with the ServicePublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePublicKey

`func (o *EnrollMachineRequestModel) SetServicePublicKey(v string)`

SetServicePublicKey sets ServicePublicKey field to given value.

### HasServicePublicKey

`func (o *EnrollMachineRequestModel) HasServicePublicKey() bool`

HasServicePublicKey returns a boolean if a field has been set.

### SetServicePublicKeyNil

`func (o *EnrollMachineRequestModel) SetServicePublicKeyNil(b bool)`

 SetServicePublicKeyNil sets the value for ServicePublicKey to be an explicit nil

### UnsetServicePublicKey
`func (o *EnrollMachineRequestModel) UnsetServicePublicKey()`

UnsetServicePublicKey ensures that no value is present for ServicePublicKey, not even an explicit nil
### GetMachineSid

`func (o *EnrollMachineRequestModel) GetMachineSid() string`

GetMachineSid returns the MachineSid field if non-nil, zero value otherwise.

### GetMachineSidOk

`func (o *EnrollMachineRequestModel) GetMachineSidOk() (*string, bool)`

GetMachineSidOk returns a tuple with the MachineSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineSid

`func (o *EnrollMachineRequestModel) SetMachineSid(v string)`

SetMachineSid sets MachineSid field to given value.

### HasMachineSid

`func (o *EnrollMachineRequestModel) HasMachineSid() bool`

HasMachineSid returns a boolean if a field has been set.

### SetMachineSidNil

`func (o *EnrollMachineRequestModel) SetMachineSidNil(b bool)`

 SetMachineSidNil sets the value for MachineSid to be an explicit nil

### UnsetMachineSid
`func (o *EnrollMachineRequestModel) UnsetMachineSid()`

UnsetMachineSid ensures that no value is present for MachineSid, not even an explicit nil
### GetMachineMetadata

`func (o *EnrollMachineRequestModel) GetMachineMetadata() EnrollMachineMetaDataModel`

GetMachineMetadata returns the MachineMetadata field if non-nil, zero value otherwise.

### GetMachineMetadataOk

`func (o *EnrollMachineRequestModel) GetMachineMetadataOk() (*EnrollMachineMetaDataModel, bool)`

GetMachineMetadataOk returns a tuple with the MachineMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineMetadata

`func (o *EnrollMachineRequestModel) SetMachineMetadata(v EnrollMachineMetaDataModel)`

SetMachineMetadata sets MachineMetadata field to given value.

### HasMachineMetadata

`func (o *EnrollMachineRequestModel) HasMachineMetadata() bool`

HasMachineMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


