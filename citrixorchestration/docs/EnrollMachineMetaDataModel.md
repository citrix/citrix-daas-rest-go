# EnrollMachineMetaDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostedMachineId** | Pointer to **NullableString** | Presented if such an id could be found from the enrolled vda machine locally | [optional] 
**MacAddresses** | Pointer to **[]string** | The enrolled machine&#39;s MAC addresses  | [optional] 

## Methods

### NewEnrollMachineMetaDataModel

`func NewEnrollMachineMetaDataModel() *EnrollMachineMetaDataModel`

NewEnrollMachineMetaDataModel instantiates a new EnrollMachineMetaDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollMachineMetaDataModelWithDefaults

`func NewEnrollMachineMetaDataModelWithDefaults() *EnrollMachineMetaDataModel`

NewEnrollMachineMetaDataModelWithDefaults instantiates a new EnrollMachineMetaDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostedMachineId

`func (o *EnrollMachineMetaDataModel) GetHostedMachineId() string`

GetHostedMachineId returns the HostedMachineId field if non-nil, zero value otherwise.

### GetHostedMachineIdOk

`func (o *EnrollMachineMetaDataModel) GetHostedMachineIdOk() (*string, bool)`

GetHostedMachineIdOk returns a tuple with the HostedMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedMachineId

`func (o *EnrollMachineMetaDataModel) SetHostedMachineId(v string)`

SetHostedMachineId sets HostedMachineId field to given value.

### HasHostedMachineId

`func (o *EnrollMachineMetaDataModel) HasHostedMachineId() bool`

HasHostedMachineId returns a boolean if a field has been set.

### SetHostedMachineIdNil

`func (o *EnrollMachineMetaDataModel) SetHostedMachineIdNil(b bool)`

 SetHostedMachineIdNil sets the value for HostedMachineId to be an explicit nil

### UnsetHostedMachineId
`func (o *EnrollMachineMetaDataModel) UnsetHostedMachineId()`

UnsetHostedMachineId ensures that no value is present for HostedMachineId, not even an explicit nil
### GetMacAddresses

`func (o *EnrollMachineMetaDataModel) GetMacAddresses() []string`

GetMacAddresses returns the MacAddresses field if non-nil, zero value otherwise.

### GetMacAddressesOk

`func (o *EnrollMachineMetaDataModel) GetMacAddressesOk() (*[]string, bool)`

GetMacAddressesOk returns a tuple with the MacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddresses

`func (o *EnrollMachineMetaDataModel) SetMacAddresses(v []string)`

SetMacAddresses sets MacAddresses field to given value.

### HasMacAddresses

`func (o *EnrollMachineMetaDataModel) HasMacAddresses() bool`

HasMacAddresses returns a boolean if a field has been set.

### SetMacAddressesNil

`func (o *EnrollMachineMetaDataModel) SetMacAddressesNil(b bool)`

 SetMacAddressesNil sets the value for MacAddresses to be an explicit nil

### UnsetMacAddresses
`func (o *EnrollMachineMetaDataModel) UnsetMacAddresses()`

UnsetMacAddresses ensures that no value is present for MacAddresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


