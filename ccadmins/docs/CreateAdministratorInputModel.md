# CreateAdministratorInputModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**AdministratorAccessType**](AdministratorAccessType.md) |  | 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**ExternalProviderId** | Pointer to **string** |  | [optional] 
**ExternalUserId** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Policies** | Pointer to [**[]AdministratorAccessPolicyModel**](AdministratorAccessPolicyModel.md) |  | [optional] 
**ProviderType** | [**AdministratorProviderType**](AdministratorProviderType.md) |  | 
**Type** | **string** |  | 

## Methods

### NewCreateAdministratorInputModel

`func NewCreateAdministratorInputModel(accessType AdministratorAccessType, providerType AdministratorProviderType, type_ string, ) *CreateAdministratorInputModel`

NewCreateAdministratorInputModel instantiates a new CreateAdministratorInputModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdministratorInputModelWithDefaults

`func NewCreateAdministratorInputModelWithDefaults() *CreateAdministratorInputModel`

NewCreateAdministratorInputModelWithDefaults instantiates a new CreateAdministratorInputModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *CreateAdministratorInputModel) GetAccessType() AdministratorAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CreateAdministratorInputModel) GetAccessTypeOk() (*AdministratorAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CreateAdministratorInputModel) SetAccessType(v AdministratorAccessType)`

SetAccessType sets AccessType field to given value.


### GetDisplayName

`func (o *CreateAdministratorInputModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateAdministratorInputModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateAdministratorInputModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateAdministratorInputModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CreateAdministratorInputModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateAdministratorInputModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmail

`func (o *CreateAdministratorInputModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateAdministratorInputModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateAdministratorInputModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateAdministratorInputModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateAdministratorInputModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateAdministratorInputModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetExternalProviderId

`func (o *CreateAdministratorInputModel) GetExternalProviderId() string`

GetExternalProviderId returns the ExternalProviderId field if non-nil, zero value otherwise.

### GetExternalProviderIdOk

`func (o *CreateAdministratorInputModel) GetExternalProviderIdOk() (*string, bool)`

GetExternalProviderIdOk returns a tuple with the ExternalProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderId

`func (o *CreateAdministratorInputModel) SetExternalProviderId(v string)`

SetExternalProviderId sets ExternalProviderId field to given value.

### HasExternalProviderId

`func (o *CreateAdministratorInputModel) HasExternalProviderId() bool`

HasExternalProviderId returns a boolean if a field has been set.

### GetExternalUserId

`func (o *CreateAdministratorInputModel) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *CreateAdministratorInputModel) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *CreateAdministratorInputModel) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *CreateAdministratorInputModel) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetFirstName

`func (o *CreateAdministratorInputModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateAdministratorInputModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateAdministratorInputModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CreateAdministratorInputModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *CreateAdministratorInputModel) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CreateAdministratorInputModel) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *CreateAdministratorInputModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateAdministratorInputModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateAdministratorInputModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CreateAdministratorInputModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *CreateAdministratorInputModel) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CreateAdministratorInputModel) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPolicies

`func (o *CreateAdministratorInputModel) GetPolicies() []AdministratorAccessPolicyModel`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CreateAdministratorInputModel) GetPoliciesOk() (*[]AdministratorAccessPolicyModel, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CreateAdministratorInputModel) SetPolicies(v []AdministratorAccessPolicyModel)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CreateAdministratorInputModel) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *CreateAdministratorInputModel) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *CreateAdministratorInputModel) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil
### GetProviderType

`func (o *CreateAdministratorInputModel) GetProviderType() AdministratorProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *CreateAdministratorInputModel) GetProviderTypeOk() (*AdministratorProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *CreateAdministratorInputModel) SetProviderType(v AdministratorProviderType)`

SetProviderType sets ProviderType field to given value.


### GetType

`func (o *CreateAdministratorInputModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAdministratorInputModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAdministratorInputModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


