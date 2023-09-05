# ProvisioningSchemeResponseModelMachineAccountCreationRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamingScheme** | Pointer to **string** | The template name for AD accounts created in the identity pool. | [optional] 
**NamingSchemeType** | Pointer to [**NamingSchemeType**](NamingSchemeType.md) |  | [optional] 
**OU** | Pointer to **string** | The OU that computer accounts will be created into. | [optional] 
**Domain** | Pointer to [**MachineAccountCreationRulesResponseModelDomain**](MachineAccountCreationRulesResponseModelDomain.md) |  | [optional] 
**NextValue** | Pointer to **string** | The next value that will be used if creating new AD accounts. | [optional] 

## Methods

### NewProvisioningSchemeResponseModelMachineAccountCreationRules

`func NewProvisioningSchemeResponseModelMachineAccountCreationRules() *ProvisioningSchemeResponseModelMachineAccountCreationRules`

NewProvisioningSchemeResponseModelMachineAccountCreationRules instantiates a new ProvisioningSchemeResponseModelMachineAccountCreationRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeResponseModelMachineAccountCreationRulesWithDefaults

`func NewProvisioningSchemeResponseModelMachineAccountCreationRulesWithDefaults() *ProvisioningSchemeResponseModelMachineAccountCreationRules`

NewProvisioningSchemeResponseModelMachineAccountCreationRulesWithDefaults instantiates a new ProvisioningSchemeResponseModelMachineAccountCreationRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamingScheme

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) GetNamingScheme() string`

GetNamingScheme returns the NamingScheme field if non-nil, zero value otherwise.

### GetNamingSchemeOk

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) GetNamingSchemeOk() (*string, bool)`

GetNamingSchemeOk returns a tuple with the NamingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingScheme

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) SetNamingScheme(v string)`

SetNamingScheme sets NamingScheme field to given value.

### HasNamingScheme

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) HasNamingScheme() bool`

HasNamingScheme returns a boolean if a field has been set.

### GetNamingSchemeType

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) GetNamingSchemeType() NamingSchemeType`

GetNamingSchemeType returns the NamingSchemeType field if non-nil, zero value otherwise.

### GetNamingSchemeTypeOk

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) GetNamingSchemeTypeOk() (*NamingSchemeType, bool)`

GetNamingSchemeTypeOk returns a tuple with the NamingSchemeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchemeType

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) SetNamingSchemeType(v NamingSchemeType)`

SetNamingSchemeType sets NamingSchemeType field to given value.

### HasNamingSchemeType

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) HasNamingSchemeType() bool`

HasNamingSchemeType returns a boolean if a field has been set.

### GetOU

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) GetOU() string`

GetOU returns the OU field if non-nil, zero value otherwise.

### GetOUOk

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) GetOUOk() (*string, bool)`

GetOUOk returns a tuple with the OU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOU

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) SetOU(v string)`

SetOU sets OU field to given value.

### HasOU

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) HasOU() bool`

HasOU returns a boolean if a field has been set.

### GetDomain

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) GetDomain() MachineAccountCreationRulesResponseModelDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) GetDomainOk() (*MachineAccountCreationRulesResponseModelDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) SetDomain(v MachineAccountCreationRulesResponseModelDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetNextValue

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) GetNextValue() string`

GetNextValue returns the NextValue field if non-nil, zero value otherwise.

### GetNextValueOk

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) GetNextValueOk() (*string, bool)`

GetNextValueOk returns a tuple with the NextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextValue

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) SetNextValue(v string)`

SetNextValue sets NextValue field to given value.

### HasNextValue

`func (o *ProvisioningSchemeResponseModelMachineAccountCreationRules) HasNextValue() bool`

HasNextValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


