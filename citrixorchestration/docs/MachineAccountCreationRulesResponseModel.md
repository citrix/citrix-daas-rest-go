# MachineAccountCreationRulesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamingScheme** | Pointer to **string** | The template name for AD accounts created in the identity pool. | [optional] 
**NamingSchemeType** | Pointer to [**NamingSchemeType**](NamingSchemeType.md) |  | [optional] 
**OU** | Pointer to **string** | The OU that computer accounts will be created into. | [optional] 
**Domain** | Pointer to [**MachineAccountCreationRulesResponseModelDomain**](MachineAccountCreationRulesResponseModelDomain.md) |  | [optional] 
**NextValue** | Pointer to **string** | The next value that will be used if creating new AD accounts. | [optional] 

## Methods

### NewMachineAccountCreationRulesResponseModel

`func NewMachineAccountCreationRulesResponseModel() *MachineAccountCreationRulesResponseModel`

NewMachineAccountCreationRulesResponseModel instantiates a new MachineAccountCreationRulesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAccountCreationRulesResponseModelWithDefaults

`func NewMachineAccountCreationRulesResponseModelWithDefaults() *MachineAccountCreationRulesResponseModel`

NewMachineAccountCreationRulesResponseModelWithDefaults instantiates a new MachineAccountCreationRulesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamingScheme

`func (o *MachineAccountCreationRulesResponseModel) GetNamingScheme() string`

GetNamingScheme returns the NamingScheme field if non-nil, zero value otherwise.

### GetNamingSchemeOk

`func (o *MachineAccountCreationRulesResponseModel) GetNamingSchemeOk() (*string, bool)`

GetNamingSchemeOk returns a tuple with the NamingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingScheme

`func (o *MachineAccountCreationRulesResponseModel) SetNamingScheme(v string)`

SetNamingScheme sets NamingScheme field to given value.

### HasNamingScheme

`func (o *MachineAccountCreationRulesResponseModel) HasNamingScheme() bool`

HasNamingScheme returns a boolean if a field has been set.

### GetNamingSchemeType

`func (o *MachineAccountCreationRulesResponseModel) GetNamingSchemeType() NamingSchemeType`

GetNamingSchemeType returns the NamingSchemeType field if non-nil, zero value otherwise.

### GetNamingSchemeTypeOk

`func (o *MachineAccountCreationRulesResponseModel) GetNamingSchemeTypeOk() (*NamingSchemeType, bool)`

GetNamingSchemeTypeOk returns a tuple with the NamingSchemeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchemeType

`func (o *MachineAccountCreationRulesResponseModel) SetNamingSchemeType(v NamingSchemeType)`

SetNamingSchemeType sets NamingSchemeType field to given value.

### HasNamingSchemeType

`func (o *MachineAccountCreationRulesResponseModel) HasNamingSchemeType() bool`

HasNamingSchemeType returns a boolean if a field has been set.

### GetOU

`func (o *MachineAccountCreationRulesResponseModel) GetOU() string`

GetOU returns the OU field if non-nil, zero value otherwise.

### GetOUOk

`func (o *MachineAccountCreationRulesResponseModel) GetOUOk() (*string, bool)`

GetOUOk returns a tuple with the OU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOU

`func (o *MachineAccountCreationRulesResponseModel) SetOU(v string)`

SetOU sets OU field to given value.

### HasOU

`func (o *MachineAccountCreationRulesResponseModel) HasOU() bool`

HasOU returns a boolean if a field has been set.

### GetDomain

`func (o *MachineAccountCreationRulesResponseModel) GetDomain() MachineAccountCreationRulesResponseModelDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *MachineAccountCreationRulesResponseModel) GetDomainOk() (*MachineAccountCreationRulesResponseModelDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *MachineAccountCreationRulesResponseModel) SetDomain(v MachineAccountCreationRulesResponseModelDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *MachineAccountCreationRulesResponseModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetNextValue

`func (o *MachineAccountCreationRulesResponseModel) GetNextValue() string`

GetNextValue returns the NextValue field if non-nil, zero value otherwise.

### GetNextValueOk

`func (o *MachineAccountCreationRulesResponseModel) GetNextValueOk() (*string, bool)`

GetNextValueOk returns a tuple with the NextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextValue

`func (o *MachineAccountCreationRulesResponseModel) SetNextValue(v string)`

SetNextValue sets NextValue field to given value.

### HasNextValue

`func (o *MachineAccountCreationRulesResponseModel) HasNextValue() bool`

HasNextValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


