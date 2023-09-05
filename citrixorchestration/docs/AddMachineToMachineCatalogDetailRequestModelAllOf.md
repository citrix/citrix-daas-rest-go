# AddMachineToMachineCatalogDetailRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PvsAddress** | Pointer to **string** | IP address of the PVS server to be used.  This only applies if the ProvisioningType is PVS. | [optional] 
**PvsDomain** | Pointer to **string** | The domain of the PVS server to be used. This only applies if the ProvisioningType is PVS. | [optional] 
**PvsCollectionIds** | Pointer to **[]string** | Collection IDs of PVS collections containing machines that should be added to the catalog.  This only applies if the ProvisioningType is PVS, and is required in that case.  Each item must be a valid PVS collection ID residing on the PVS server located at the specified . | [optional] 
**MachineAccountCreationRules** | Pointer to [**AddMachineToMachineCatalogDetailRequestModelAllOfMachineAccountCreationRules**](AddMachineToMachineCatalogDetailRequestModelAllOfMachineAccountCreationRules.md) |  | [optional] 
**AddAvailableMachineAccount** | Pointer to [**AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount**](AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount.md) |  | [optional] 

## Methods

### NewAddMachineToMachineCatalogDetailRequestModelAllOf

`func NewAddMachineToMachineCatalogDetailRequestModelAllOf() *AddMachineToMachineCatalogDetailRequestModelAllOf`

NewAddMachineToMachineCatalogDetailRequestModelAllOf instantiates a new AddMachineToMachineCatalogDetailRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMachineToMachineCatalogDetailRequestModelAllOfWithDefaults

`func NewAddMachineToMachineCatalogDetailRequestModelAllOfWithDefaults() *AddMachineToMachineCatalogDetailRequestModelAllOf`

NewAddMachineToMachineCatalogDetailRequestModelAllOfWithDefaults instantiates a new AddMachineToMachineCatalogDetailRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPvsAddress

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) GetPvsAddress() string`

GetPvsAddress returns the PvsAddress field if non-nil, zero value otherwise.

### GetPvsAddressOk

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) GetPvsAddressOk() (*string, bool)`

GetPvsAddressOk returns a tuple with the PvsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsAddress

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) SetPvsAddress(v string)`

SetPvsAddress sets PvsAddress field to given value.

### HasPvsAddress

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) HasPvsAddress() bool`

HasPvsAddress returns a boolean if a field has been set.

### GetPvsDomain

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) GetPvsDomain() string`

GetPvsDomain returns the PvsDomain field if non-nil, zero value otherwise.

### GetPvsDomainOk

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) GetPvsDomainOk() (*string, bool)`

GetPvsDomainOk returns a tuple with the PvsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsDomain

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) SetPvsDomain(v string)`

SetPvsDomain sets PvsDomain field to given value.

### HasPvsDomain

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) HasPvsDomain() bool`

HasPvsDomain returns a boolean if a field has been set.

### GetPvsCollectionIds

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) GetPvsCollectionIds() []string`

GetPvsCollectionIds returns the PvsCollectionIds field if non-nil, zero value otherwise.

### GetPvsCollectionIdsOk

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) GetPvsCollectionIdsOk() (*[]string, bool)`

GetPvsCollectionIdsOk returns a tuple with the PvsCollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsCollectionIds

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) SetPvsCollectionIds(v []string)`

SetPvsCollectionIds sets PvsCollectionIds field to given value.

### HasPvsCollectionIds

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) HasPvsCollectionIds() bool`

HasPvsCollectionIds returns a boolean if a field has been set.

### GetMachineAccountCreationRules

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) GetMachineAccountCreationRules() AddMachineToMachineCatalogDetailRequestModelAllOfMachineAccountCreationRules`

GetMachineAccountCreationRules returns the MachineAccountCreationRules field if non-nil, zero value otherwise.

### GetMachineAccountCreationRulesOk

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) GetMachineAccountCreationRulesOk() (*AddMachineToMachineCatalogDetailRequestModelAllOfMachineAccountCreationRules, bool)`

GetMachineAccountCreationRulesOk returns a tuple with the MachineAccountCreationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccountCreationRules

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) SetMachineAccountCreationRules(v AddMachineToMachineCatalogDetailRequestModelAllOfMachineAccountCreationRules)`

SetMachineAccountCreationRules sets MachineAccountCreationRules field to given value.

### HasMachineAccountCreationRules

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) HasMachineAccountCreationRules() bool`

HasMachineAccountCreationRules returns a boolean if a field has been set.

### GetAddAvailableMachineAccount

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) GetAddAvailableMachineAccount() AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount`

GetAddAvailableMachineAccount returns the AddAvailableMachineAccount field if non-nil, zero value otherwise.

### GetAddAvailableMachineAccountOk

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) GetAddAvailableMachineAccountOk() (*AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount, bool)`

GetAddAvailableMachineAccountOk returns a tuple with the AddAvailableMachineAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAvailableMachineAccount

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) SetAddAvailableMachineAccount(v AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount)`

SetAddAvailableMachineAccount sets AddAvailableMachineAccount field to given value.

### HasAddAvailableMachineAccount

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOf) HasAddAvailableMachineAccount() bool`

HasAddAvailableMachineAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


