# MachineAccountCreationRulesRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamingScheme** | **string** | Defines the template name for AD accounts created in the identity pool.  Required. | 
**NamingSchemeType** | Pointer to [**NamingSchemeType**](NamingSchemeType.md) |  | [optional] 
**OU** | Pointer to **string** | The OU that computer accounts will be created into.  Optional. | [optional] 
**Domain** | **string** | The AD domain name for the pool. Specify this in FQDN format; for example, MyDomain.com. Required. | 
**NextValue** | Pointer to **string** | Defines the next value that will be used if creating new AD accounts.  Optional. | [optional] 
**IdentityPoolId** | Pointer to **string** | Existing identity pool id | [optional] 

## Methods

### NewMachineAccountCreationRulesRequestModel

`func NewMachineAccountCreationRulesRequestModel(namingScheme string, domain string, ) *MachineAccountCreationRulesRequestModel`

NewMachineAccountCreationRulesRequestModel instantiates a new MachineAccountCreationRulesRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAccountCreationRulesRequestModelWithDefaults

`func NewMachineAccountCreationRulesRequestModelWithDefaults() *MachineAccountCreationRulesRequestModel`

NewMachineAccountCreationRulesRequestModelWithDefaults instantiates a new MachineAccountCreationRulesRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamingScheme

`func (o *MachineAccountCreationRulesRequestModel) GetNamingScheme() string`

GetNamingScheme returns the NamingScheme field if non-nil, zero value otherwise.

### GetNamingSchemeOk

`func (o *MachineAccountCreationRulesRequestModel) GetNamingSchemeOk() (*string, bool)`

GetNamingSchemeOk returns a tuple with the NamingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingScheme

`func (o *MachineAccountCreationRulesRequestModel) SetNamingScheme(v string)`

SetNamingScheme sets NamingScheme field to given value.


### GetNamingSchemeType

`func (o *MachineAccountCreationRulesRequestModel) GetNamingSchemeType() NamingSchemeType`

GetNamingSchemeType returns the NamingSchemeType field if non-nil, zero value otherwise.

### GetNamingSchemeTypeOk

`func (o *MachineAccountCreationRulesRequestModel) GetNamingSchemeTypeOk() (*NamingSchemeType, bool)`

GetNamingSchemeTypeOk returns a tuple with the NamingSchemeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchemeType

`func (o *MachineAccountCreationRulesRequestModel) SetNamingSchemeType(v NamingSchemeType)`

SetNamingSchemeType sets NamingSchemeType field to given value.

### HasNamingSchemeType

`func (o *MachineAccountCreationRulesRequestModel) HasNamingSchemeType() bool`

HasNamingSchemeType returns a boolean if a field has been set.

### GetOU

`func (o *MachineAccountCreationRulesRequestModel) GetOU() string`

GetOU returns the OU field if non-nil, zero value otherwise.

### GetOUOk

`func (o *MachineAccountCreationRulesRequestModel) GetOUOk() (*string, bool)`

GetOUOk returns a tuple with the OU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOU

`func (o *MachineAccountCreationRulesRequestModel) SetOU(v string)`

SetOU sets OU field to given value.

### HasOU

`func (o *MachineAccountCreationRulesRequestModel) HasOU() bool`

HasOU returns a boolean if a field has been set.

### GetDomain

`func (o *MachineAccountCreationRulesRequestModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *MachineAccountCreationRulesRequestModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *MachineAccountCreationRulesRequestModel) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetNextValue

`func (o *MachineAccountCreationRulesRequestModel) GetNextValue() string`

GetNextValue returns the NextValue field if non-nil, zero value otherwise.

### GetNextValueOk

`func (o *MachineAccountCreationRulesRequestModel) GetNextValueOk() (*string, bool)`

GetNextValueOk returns a tuple with the NextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextValue

`func (o *MachineAccountCreationRulesRequestModel) SetNextValue(v string)`

SetNextValue sets NextValue field to given value.

### HasNextValue

`func (o *MachineAccountCreationRulesRequestModel) HasNextValue() bool`

HasNextValue returns a boolean if a field has been set.

### GetIdentityPoolId

`func (o *MachineAccountCreationRulesRequestModel) GetIdentityPoolId() string`

GetIdentityPoolId returns the IdentityPoolId field if non-nil, zero value otherwise.

### GetIdentityPoolIdOk

`func (o *MachineAccountCreationRulesRequestModel) GetIdentityPoolIdOk() (*string, bool)`

GetIdentityPoolIdOk returns a tuple with the IdentityPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityPoolId

`func (o *MachineAccountCreationRulesRequestModel) SetIdentityPoolId(v string)`

SetIdentityPoolId sets IdentityPoolId field to given value.

### HasIdentityPoolId

`func (o *MachineAccountCreationRulesRequestModel) HasIdentityPoolId() bool`

HasIdentityPoolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


