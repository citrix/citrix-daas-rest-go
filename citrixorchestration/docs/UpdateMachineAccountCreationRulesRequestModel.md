# UpdateMachineAccountCreationRulesRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamingScheme** | Pointer to **string** | Defines the template name for AD accounts created in the identity pool.   If this is not specified, it will remain unchanged.   If the provisioning scheme is configured with a NamingScheme already, and this value is set to an empty string, the provisioning scheme will be reconfigured so that it will no longer automatically create machine accounts. | [optional] 
**NamingSchemeType** | Pointer to [**NamingSchemeType**](NamingSchemeType.md) |  | [optional] 
**OU** | Pointer to **string** | The OU that computer accounts will be created into.   If not specified, will not be changed.   Cannot be specified if  is set to an empty string.   If  was not previously set, but is being set now, then use the default account container specified by AD. This is the &#x60;Computers&#x60; container for out-of-the-box installations of AD. | [optional] 
**Domain** | Pointer to **string** | The AD domain name for the pool. Specify this in FQDN format; for example, MyDomain.com.   If not specified, will not be changed.   Cannot be specified if  is set to an empty string.   If  was not previously set, but is being set now, this property is required. | [optional] 
**NextValue** | Pointer to **string** | Defines the next value that will be used if creating new AD accounts.   If not specified, will not be changed.   Cannot be specified if  is set to an empty string.   If  was not previously set, but is being set now, the default is a sequence of &#x60;0&#x60;s or &#x60;A&#x60;s, depending on the . | [optional] 

## Methods

### NewUpdateMachineAccountCreationRulesRequestModel

`func NewUpdateMachineAccountCreationRulesRequestModel() *UpdateMachineAccountCreationRulesRequestModel`

NewUpdateMachineAccountCreationRulesRequestModel instantiates a new UpdateMachineAccountCreationRulesRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMachineAccountCreationRulesRequestModelWithDefaults

`func NewUpdateMachineAccountCreationRulesRequestModelWithDefaults() *UpdateMachineAccountCreationRulesRequestModel`

NewUpdateMachineAccountCreationRulesRequestModelWithDefaults instantiates a new UpdateMachineAccountCreationRulesRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamingScheme

`func (o *UpdateMachineAccountCreationRulesRequestModel) GetNamingScheme() string`

GetNamingScheme returns the NamingScheme field if non-nil, zero value otherwise.

### GetNamingSchemeOk

`func (o *UpdateMachineAccountCreationRulesRequestModel) GetNamingSchemeOk() (*string, bool)`

GetNamingSchemeOk returns a tuple with the NamingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingScheme

`func (o *UpdateMachineAccountCreationRulesRequestModel) SetNamingScheme(v string)`

SetNamingScheme sets NamingScheme field to given value.

### HasNamingScheme

`func (o *UpdateMachineAccountCreationRulesRequestModel) HasNamingScheme() bool`

HasNamingScheme returns a boolean if a field has been set.

### GetNamingSchemeType

`func (o *UpdateMachineAccountCreationRulesRequestModel) GetNamingSchemeType() NamingSchemeType`

GetNamingSchemeType returns the NamingSchemeType field if non-nil, zero value otherwise.

### GetNamingSchemeTypeOk

`func (o *UpdateMachineAccountCreationRulesRequestModel) GetNamingSchemeTypeOk() (*NamingSchemeType, bool)`

GetNamingSchemeTypeOk returns a tuple with the NamingSchemeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchemeType

`func (o *UpdateMachineAccountCreationRulesRequestModel) SetNamingSchemeType(v NamingSchemeType)`

SetNamingSchemeType sets NamingSchemeType field to given value.

### HasNamingSchemeType

`func (o *UpdateMachineAccountCreationRulesRequestModel) HasNamingSchemeType() bool`

HasNamingSchemeType returns a boolean if a field has been set.

### GetOU

`func (o *UpdateMachineAccountCreationRulesRequestModel) GetOU() string`

GetOU returns the OU field if non-nil, zero value otherwise.

### GetOUOk

`func (o *UpdateMachineAccountCreationRulesRequestModel) GetOUOk() (*string, bool)`

GetOUOk returns a tuple with the OU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOU

`func (o *UpdateMachineAccountCreationRulesRequestModel) SetOU(v string)`

SetOU sets OU field to given value.

### HasOU

`func (o *UpdateMachineAccountCreationRulesRequestModel) HasOU() bool`

HasOU returns a boolean if a field has been set.

### GetDomain

`func (o *UpdateMachineAccountCreationRulesRequestModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpdateMachineAccountCreationRulesRequestModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpdateMachineAccountCreationRulesRequestModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UpdateMachineAccountCreationRulesRequestModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetNextValue

`func (o *UpdateMachineAccountCreationRulesRequestModel) GetNextValue() string`

GetNextValue returns the NextValue field if non-nil, zero value otherwise.

### GetNextValueOk

`func (o *UpdateMachineAccountCreationRulesRequestModel) GetNextValueOk() (*string, bool)`

GetNextValueOk returns a tuple with the NextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextValue

`func (o *UpdateMachineAccountCreationRulesRequestModel) SetNextValue(v string)`

SetNextValue sets NextValue field to given value.

### HasNextValue

`func (o *UpdateMachineAccountCreationRulesRequestModel) HasNextValue() bool`

HasNextValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


