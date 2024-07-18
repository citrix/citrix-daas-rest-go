# DeleteMachines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**MachineIds** | **[]string** | The list of machine ids to be deleted | 
**Credentials** | Pointer to [**NullableDeleteMachinesCredentials**](DeleteMachinesCredentials.md) |  | [optional] 

## Methods

### NewDeleteMachines

`func NewDeleteMachines(accountType AccountType, machineIds []string, ) *DeleteMachines`

NewDeleteMachines instantiates a new DeleteMachines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMachinesWithDefaults

`func NewDeleteMachinesWithDefaults() *DeleteMachines`

NewDeleteMachinesWithDefaults instantiates a new DeleteMachines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *DeleteMachines) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *DeleteMachines) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *DeleteMachines) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetMachineIds

`func (o *DeleteMachines) GetMachineIds() []string`

GetMachineIds returns the MachineIds field if non-nil, zero value otherwise.

### GetMachineIdsOk

`func (o *DeleteMachines) GetMachineIdsOk() (*[]string, bool)`

GetMachineIdsOk returns a tuple with the MachineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineIds

`func (o *DeleteMachines) SetMachineIds(v []string)`

SetMachineIds sets MachineIds field to given value.


### GetCredentials

`func (o *DeleteMachines) GetCredentials() DeleteMachinesCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DeleteMachines) GetCredentialsOk() (*DeleteMachinesCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DeleteMachines) SetCredentials(v DeleteMachinesCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *DeleteMachines) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *DeleteMachines) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *DeleteMachines) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


