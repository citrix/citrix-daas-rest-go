# CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UcOid** | **string** |  | 
**Type** | **string** |  | 
**AccessType** | **string** |  | 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**ProviderType** | **string** |  | 
**ProviderId** | **string** |  | 
**ProviderProperties** | Pointer to **map[string]string** |  | [optional] 
**EmailPreferences** | Pointer to **NullableString** |  | [optional] 
**AuthDomain** | **string** |  | 
**Policies** | Pointer to [**[]CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel**](CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel.md) |  | [optional] 

## Methods

### NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel

`func NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel(ucOid string, type_ string, accessType string, providerType string, providerId string, authDomain string, ) *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel`

NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel instantiates a new CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorModelWithDefaults

`func NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorModelWithDefaults() *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel`

NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorModelWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUcOid

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetUcOid() string`

GetUcOid returns the UcOid field if non-nil, zero value otherwise.

### GetUcOidOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetUcOidOk() (*string, bool)`

GetUcOidOk returns a tuple with the UcOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcOid

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetUcOid(v string)`

SetUcOid sets UcOid field to given value.


### GetType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetType(v string)`

SetType sets Type field to given value.


### GetAccessType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetFirstName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetProviderType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetProviderId

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetProviderProperties

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetProviderProperties() map[string]string`

GetProviderProperties returns the ProviderProperties field if non-nil, zero value otherwise.

### GetProviderPropertiesOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetProviderPropertiesOk() (*map[string]string, bool)`

GetProviderPropertiesOk returns a tuple with the ProviderProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderProperties

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetProviderProperties(v map[string]string)`

SetProviderProperties sets ProviderProperties field to given value.

### HasProviderProperties

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) HasProviderProperties() bool`

HasProviderProperties returns a boolean if a field has been set.

### SetProviderPropertiesNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetProviderPropertiesNil(b bool)`

 SetProviderPropertiesNil sets the value for ProviderProperties to be an explicit nil

### UnsetProviderProperties
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) UnsetProviderProperties()`

UnsetProviderProperties ensures that no value is present for ProviderProperties, not even an explicit nil
### GetEmailPreferences

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetEmailPreferences() string`

GetEmailPreferences returns the EmailPreferences field if non-nil, zero value otherwise.

### GetEmailPreferencesOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetEmailPreferencesOk() (*string, bool)`

GetEmailPreferencesOk returns a tuple with the EmailPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPreferences

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetEmailPreferences(v string)`

SetEmailPreferences sets EmailPreferences field to given value.

### HasEmailPreferences

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) HasEmailPreferences() bool`

HasEmailPreferences returns a boolean if a field has been set.

### SetEmailPreferencesNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetEmailPreferencesNil(b bool)`

 SetEmailPreferencesNil sets the value for EmailPreferences to be an explicit nil

### UnsetEmailPreferences
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) UnsetEmailPreferences()`

UnsetEmailPreferences ensures that no value is present for EmailPreferences, not even an explicit nil
### GetAuthDomain

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetAuthDomain() string`

GetAuthDomain returns the AuthDomain field if non-nil, zero value otherwise.

### GetAuthDomainOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetAuthDomainOk() (*string, bool)`

GetAuthDomainOk returns a tuple with the AuthDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomain

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetAuthDomain(v string)`

SetAuthDomain sets AuthDomain field to given value.


### GetPolicies

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetPolicies() []CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) GetPoliciesOk() (*[]CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetPolicies(v []CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


