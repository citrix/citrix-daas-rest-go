# IdentityBaseResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | The Domain the object belongs to. NOTE: When doing a forest level search, the domain property is populated only if the object contains a SID and the SID was requested as property to be fetched. | [optional] 
**Forest** | Pointer to **NullableString** | The forest the object belongs to. | [optional] 
**Guid** | Pointer to **NullableString** | The GUID of the object, the GUID is unique across the enterprise and anywhere else. | [optional] 
**DistinguishedName** | Pointer to **NullableString** | The Distinguished name of the object. Distinguished names (DNs) are unique and they unambiguously identify objects in the directory. | [optional] 
**CanonicalName** | Pointer to **NullableString** | The unique distinguished name of the object in canonical format. | [optional] 
**PossibleLookupFailure** | Pointer to **bool** | Set to True during a bulk identity lookup operation. For individual identity lookup, usually an exception will be thrown. | [optional] 
**DirectoryServer** | Pointer to **NullableString** | Domain controller that performed the add/update operation on this object | [optional] 

## Methods

### NewIdentityBaseResponseModel

`func NewIdentityBaseResponseModel() *IdentityBaseResponseModel`

NewIdentityBaseResponseModel instantiates a new IdentityBaseResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityBaseResponseModelWithDefaults

`func NewIdentityBaseResponseModelWithDefaults() *IdentityBaseResponseModel`

NewIdentityBaseResponseModelWithDefaults instantiates a new IdentityBaseResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *IdentityBaseResponseModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentityBaseResponseModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentityBaseResponseModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdentityBaseResponseModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *IdentityBaseResponseModel) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *IdentityBaseResponseModel) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetForest

`func (o *IdentityBaseResponseModel) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *IdentityBaseResponseModel) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *IdentityBaseResponseModel) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *IdentityBaseResponseModel) HasForest() bool`

HasForest returns a boolean if a field has been set.

### SetForestNil

`func (o *IdentityBaseResponseModel) SetForestNil(b bool)`

 SetForestNil sets the value for Forest to be an explicit nil

### UnsetForest
`func (o *IdentityBaseResponseModel) UnsetForest()`

UnsetForest ensures that no value is present for Forest, not even an explicit nil
### GetGuid

`func (o *IdentityBaseResponseModel) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IdentityBaseResponseModel) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IdentityBaseResponseModel) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IdentityBaseResponseModel) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *IdentityBaseResponseModel) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *IdentityBaseResponseModel) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetDistinguishedName

`func (o *IdentityBaseResponseModel) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *IdentityBaseResponseModel) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *IdentityBaseResponseModel) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *IdentityBaseResponseModel) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *IdentityBaseResponseModel) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *IdentityBaseResponseModel) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetCanonicalName

`func (o *IdentityBaseResponseModel) GetCanonicalName() string`

GetCanonicalName returns the CanonicalName field if non-nil, zero value otherwise.

### GetCanonicalNameOk

`func (o *IdentityBaseResponseModel) GetCanonicalNameOk() (*string, bool)`

GetCanonicalNameOk returns a tuple with the CanonicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalName

`func (o *IdentityBaseResponseModel) SetCanonicalName(v string)`

SetCanonicalName sets CanonicalName field to given value.

### HasCanonicalName

`func (o *IdentityBaseResponseModel) HasCanonicalName() bool`

HasCanonicalName returns a boolean if a field has been set.

### SetCanonicalNameNil

`func (o *IdentityBaseResponseModel) SetCanonicalNameNil(b bool)`

 SetCanonicalNameNil sets the value for CanonicalName to be an explicit nil

### UnsetCanonicalName
`func (o *IdentityBaseResponseModel) UnsetCanonicalName()`

UnsetCanonicalName ensures that no value is present for CanonicalName, not even an explicit nil
### GetPossibleLookupFailure

`func (o *IdentityBaseResponseModel) GetPossibleLookupFailure() bool`

GetPossibleLookupFailure returns the PossibleLookupFailure field if non-nil, zero value otherwise.

### GetPossibleLookupFailureOk

`func (o *IdentityBaseResponseModel) GetPossibleLookupFailureOk() (*bool, bool)`

GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleLookupFailure

`func (o *IdentityBaseResponseModel) SetPossibleLookupFailure(v bool)`

SetPossibleLookupFailure sets PossibleLookupFailure field to given value.

### HasPossibleLookupFailure

`func (o *IdentityBaseResponseModel) HasPossibleLookupFailure() bool`

HasPossibleLookupFailure returns a boolean if a field has been set.

### GetDirectoryServer

`func (o *IdentityBaseResponseModel) GetDirectoryServer() string`

GetDirectoryServer returns the DirectoryServer field if non-nil, zero value otherwise.

### GetDirectoryServerOk

`func (o *IdentityBaseResponseModel) GetDirectoryServerOk() (*string, bool)`

GetDirectoryServerOk returns a tuple with the DirectoryServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServer

`func (o *IdentityBaseResponseModel) SetDirectoryServer(v string)`

SetDirectoryServer sets DirectoryServer field to given value.

### HasDirectoryServer

`func (o *IdentityBaseResponseModel) HasDirectoryServer() bool`

HasDirectoryServer returns a boolean if a field has been set.

### SetDirectoryServerNil

`func (o *IdentityBaseResponseModel) SetDirectoryServerNil(b bool)`

 SetDirectoryServerNil sets the value for DirectoryServer to be an explicit nil

### UnsetDirectoryServer
`func (o *IdentityBaseResponseModel) UnsetDirectoryServer()`

UnsetDirectoryServer ensures that no value is present for DirectoryServer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


