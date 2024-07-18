# CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AccessType** | **string** |  | 
**ProviderType** | **string** |  | 
**Policies** | Pointer to [**[]CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel**](CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel.md) |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**ExternalProviderId** | Pointer to **NullableString** |  | [optional] 
**ExternalUserId** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel

`func NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel(type_ string, accessType string, providerType string, ) *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel`

NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel instantiates a new CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModelWithDefaults

`func NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModelWithDefaults() *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel`

NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModelWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetType(v string)`

SetType sets Type field to given value.


### GetAccessType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetProviderType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetPolicies

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetPolicies() []CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetPoliciesOk() (*[]CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetPolicies(v []CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil
### GetEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetExternalProviderId

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetExternalProviderId() string`

GetExternalProviderId returns the ExternalProviderId field if non-nil, zero value otherwise.

### GetExternalProviderIdOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetExternalProviderIdOk() (*string, bool)`

GetExternalProviderIdOk returns a tuple with the ExternalProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderId

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetExternalProviderId(v string)`

SetExternalProviderId sets ExternalProviderId field to given value.

### HasExternalProviderId

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) HasExternalProviderId() bool`

HasExternalProviderId returns a boolean if a field has been set.

### SetExternalProviderIdNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetExternalProviderIdNil(b bool)`

 SetExternalProviderIdNil sets the value for ExternalProviderId to be an explicit nil

### UnsetExternalProviderId
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) UnsetExternalProviderId()`

UnsetExternalProviderId ensures that no value is present for ExternalProviderId, not even an explicit nil
### GetExternalUserId

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### SetExternalUserIdNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetExternalUserIdNil(b bool)`

 SetExternalUserIdNil sets the value for ExternalUserId to be an explicit nil

### UnsetExternalUserId
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) UnsetExternalUserId()`

UnsetExternalUserId ensures that no value is present for ExternalUserId, not even an explicit nil
### GetFirstName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


