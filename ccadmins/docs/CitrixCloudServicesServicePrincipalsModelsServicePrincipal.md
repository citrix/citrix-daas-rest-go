# CitrixCloudServicesServicePrincipalsModelsServicePrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UcOid** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**ServiceProfile** | Pointer to **NullableString** |  | [optional] 
**Creator** | Pointer to [**CitrixCloudServicesServicePrincipalsModelsCreator**](CitrixCloudServicesServicePrincipalsModelsCreator.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**AccessType** | Pointer to [**CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType**](CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType.md) |  | [optional] 
**Primary** | Pointer to [**CitrixCloudServicesServicePrincipalsModelsSecretMetadata**](CitrixCloudServicesServicePrincipalsModelsSecretMetadata.md) |  | [optional] 
**Secondary** | Pointer to [**CitrixCloudServicesServicePrincipalsModelsSecretMetadata**](CitrixCloudServicesServicePrincipalsModelsSecretMetadata.md) |  | [optional] 
**LastAccessedDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCitrixCloudServicesServicePrincipalsModelsServicePrincipal

`func NewCitrixCloudServicesServicePrincipalsModelsServicePrincipal() *CitrixCloudServicesServicePrincipalsModelsServicePrincipal`

NewCitrixCloudServicesServicePrincipalsModelsServicePrincipal instantiates a new CitrixCloudServicesServicePrincipalsModelsServicePrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesServicePrincipalsModelsServicePrincipalWithDefaults

`func NewCitrixCloudServicesServicePrincipalsModelsServicePrincipalWithDefaults() *CitrixCloudServicesServicePrincipalsModelsServicePrincipal`

NewCitrixCloudServicesServicePrincipalsModelsServicePrincipalWithDefaults instantiates a new CitrixCloudServicesServicePrincipalsModelsServicePrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUcOid

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetUcOid() string`

GetUcOid returns the UcOid field if non-nil, zero value otherwise.

### GetUcOidOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetUcOidOk() (*string, bool)`

GetUcOidOk returns a tuple with the UcOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcOid

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetUcOid(v string)`

SetUcOid sets UcOid field to given value.

### HasUcOid

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasUcOid() bool`

HasUcOid returns a boolean if a field has been set.

### SetUcOidNil

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetUcOidNil(b bool)`

 SetUcOidNil sets the value for UcOid to be an explicit nil

### UnsetUcOid
`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) UnsetUcOid()`

UnsetUcOid ensures that no value is present for UcOid, not even an explicit nil
### GetClientId

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetCustomerId

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetServiceProfile

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### SetServiceProfileNil

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetServiceProfileNil(b bool)`

 SetServiceProfileNil sets the value for ServiceProfile to be an explicit nil

### UnsetServiceProfile
`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) UnsetServiceProfile()`

UnsetServiceProfile ensures that no value is present for ServiceProfile, not even an explicit nil
### GetCreator

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetCreator() CitrixCloudServicesServicePrincipalsModelsCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetCreatorOk() (*CitrixCloudServicesServicePrincipalsModelsCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetCreator(v CitrixCloudServicesServicePrincipalsModelsCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCreatedDate

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetName

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccessType

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetAccessType() CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetAccessTypeOk() (*CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetAccessType(v CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetPrimary

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetPrimary() CitrixCloudServicesServicePrincipalsModelsSecretMetadata`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetPrimaryOk() (*CitrixCloudServicesServicePrincipalsModelsSecretMetadata, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetPrimary(v CitrixCloudServicesServicePrincipalsModelsSecretMetadata)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetSecondary

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetSecondary() CitrixCloudServicesServicePrincipalsModelsSecretMetadata`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetSecondaryOk() (*CitrixCloudServicesServicePrincipalsModelsSecretMetadata, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetSecondary(v CitrixCloudServicesServicePrincipalsModelsSecretMetadata)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetLastAccessedDate

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetLastAccessedDate() time.Time`

GetLastAccessedDate returns the LastAccessedDate field if non-nil, zero value otherwise.

### GetLastAccessedDateOk

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) GetLastAccessedDateOk() (*time.Time, bool)`

GetLastAccessedDateOk returns a tuple with the LastAccessedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedDate

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetLastAccessedDate(v time.Time)`

SetLastAccessedDate sets LastAccessedDate field to given value.

### HasLastAccessedDate

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) HasLastAccessedDate() bool`

HasLastAccessedDate returns a boolean if a field has been set.

### SetLastAccessedDateNil

`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) SetLastAccessedDateNil(b bool)`

 SetLastAccessedDateNil sets the value for LastAccessedDate to be an explicit nil

### UnsetLastAccessedDate
`func (o *CitrixCloudServicesServicePrincipalsModelsServicePrincipal) UnsetLastAccessedDate()`

UnsetLastAccessedDate ensures that no value is present for LastAccessedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


