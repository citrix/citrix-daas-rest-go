# UpdatedCustomerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**AuthDomainNames** | Pointer to **[]string** |  | [optional] 
**AuthDomains** | Pointer to [**[]AuthDomainModel**](AuthDomainModel.md) |  | [optional] 
**Aliases** | Pointer to [**[]AliasModel**](AliasModel.md) |  | [optional] 

## Methods

### NewUpdatedCustomerModel

`func NewUpdatedCustomerModel() *UpdatedCustomerModel`

NewUpdatedCustomerModel instantiates a new UpdatedCustomerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedCustomerModelWithDefaults

`func NewUpdatedCustomerModelWithDefaults() *UpdatedCustomerModel`

NewUpdatedCustomerModelWithDefaults instantiates a new UpdatedCustomerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatedCustomerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatedCustomerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatedCustomerModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdatedCustomerModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UpdatedCustomerModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdatedCustomerModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *UpdatedCustomerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatedCustomerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatedCustomerModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatedCustomerModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdatedCustomerModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdatedCustomerModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAuthDomainNames

`func (o *UpdatedCustomerModel) GetAuthDomainNames() []string`

GetAuthDomainNames returns the AuthDomainNames field if non-nil, zero value otherwise.

### GetAuthDomainNamesOk

`func (o *UpdatedCustomerModel) GetAuthDomainNamesOk() (*[]string, bool)`

GetAuthDomainNamesOk returns a tuple with the AuthDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomainNames

`func (o *UpdatedCustomerModel) SetAuthDomainNames(v []string)`

SetAuthDomainNames sets AuthDomainNames field to given value.

### HasAuthDomainNames

`func (o *UpdatedCustomerModel) HasAuthDomainNames() bool`

HasAuthDomainNames returns a boolean if a field has been set.

### SetAuthDomainNamesNil

`func (o *UpdatedCustomerModel) SetAuthDomainNamesNil(b bool)`

 SetAuthDomainNamesNil sets the value for AuthDomainNames to be an explicit nil

### UnsetAuthDomainNames
`func (o *UpdatedCustomerModel) UnsetAuthDomainNames()`

UnsetAuthDomainNames ensures that no value is present for AuthDomainNames, not even an explicit nil
### GetAuthDomains

`func (o *UpdatedCustomerModel) GetAuthDomains() []AuthDomainModel`

GetAuthDomains returns the AuthDomains field if non-nil, zero value otherwise.

### GetAuthDomainsOk

`func (o *UpdatedCustomerModel) GetAuthDomainsOk() (*[]AuthDomainModel, bool)`

GetAuthDomainsOk returns a tuple with the AuthDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomains

`func (o *UpdatedCustomerModel) SetAuthDomains(v []AuthDomainModel)`

SetAuthDomains sets AuthDomains field to given value.

### HasAuthDomains

`func (o *UpdatedCustomerModel) HasAuthDomains() bool`

HasAuthDomains returns a boolean if a field has been set.

### SetAuthDomainsNil

`func (o *UpdatedCustomerModel) SetAuthDomainsNil(b bool)`

 SetAuthDomainsNil sets the value for AuthDomains to be an explicit nil

### UnsetAuthDomains
`func (o *UpdatedCustomerModel) UnsetAuthDomains()`

UnsetAuthDomains ensures that no value is present for AuthDomains, not even an explicit nil
### GetAliases

`func (o *UpdatedCustomerModel) GetAliases() []AliasModel`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *UpdatedCustomerModel) GetAliasesOk() (*[]AliasModel, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *UpdatedCustomerModel) SetAliases(v []AliasModel)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *UpdatedCustomerModel) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliasesNil

`func (o *UpdatedCustomerModel) SetAliasesNil(b bool)`

 SetAliasesNil sets the value for Aliases to be an explicit nil

### UnsetAliases
`func (o *UpdatedCustomerModel) UnsetAliases()`

UnsetAliases ensures that no value is present for Aliases, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


