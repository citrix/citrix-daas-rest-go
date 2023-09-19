# IdentityForestResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name associated with the forest object. | [optional] 
**RootDomain** | Pointer to **NullableString** | The root domain of the forest. | [optional] 
**Domains** | Pointer to **[]string** | The domains under the forest. | [optional] 
**Sites** | Pointer to **[]string** | The sites under the forest. | [optional] 
**TrustedForests** | Pointer to **[]string** | Other forests that have trust with this forest. | [optional] 
**PossibleLookupFailure** | Pointer to **bool** | Set to True during a bulk identity lookup operation. For individual identity lookup, usually an exception will be throw. | [optional] 
**PropertiesFetched** | **int32** | Properties fetched and populated in the Forest object.  This is a bitfield indicating the fetched properties. | 

## Methods

### NewIdentityForestResponseModel

`func NewIdentityForestResponseModel(propertiesFetched int32, ) *IdentityForestResponseModel`

NewIdentityForestResponseModel instantiates a new IdentityForestResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityForestResponseModelWithDefaults

`func NewIdentityForestResponseModelWithDefaults() *IdentityForestResponseModel`

NewIdentityForestResponseModelWithDefaults instantiates a new IdentityForestResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdentityForestResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityForestResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityForestResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityForestResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IdentityForestResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdentityForestResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRootDomain

`func (o *IdentityForestResponseModel) GetRootDomain() string`

GetRootDomain returns the RootDomain field if non-nil, zero value otherwise.

### GetRootDomainOk

`func (o *IdentityForestResponseModel) GetRootDomainOk() (*string, bool)`

GetRootDomainOk returns a tuple with the RootDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDomain

`func (o *IdentityForestResponseModel) SetRootDomain(v string)`

SetRootDomain sets RootDomain field to given value.

### HasRootDomain

`func (o *IdentityForestResponseModel) HasRootDomain() bool`

HasRootDomain returns a boolean if a field has been set.

### SetRootDomainNil

`func (o *IdentityForestResponseModel) SetRootDomainNil(b bool)`

 SetRootDomainNil sets the value for RootDomain to be an explicit nil

### UnsetRootDomain
`func (o *IdentityForestResponseModel) UnsetRootDomain()`

UnsetRootDomain ensures that no value is present for RootDomain, not even an explicit nil
### GetDomains

`func (o *IdentityForestResponseModel) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IdentityForestResponseModel) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IdentityForestResponseModel) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IdentityForestResponseModel) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomainsNil

`func (o *IdentityForestResponseModel) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *IdentityForestResponseModel) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil
### GetSites

`func (o *IdentityForestResponseModel) GetSites() []string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *IdentityForestResponseModel) GetSitesOk() (*[]string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *IdentityForestResponseModel) SetSites(v []string)`

SetSites sets Sites field to given value.

### HasSites

`func (o *IdentityForestResponseModel) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *IdentityForestResponseModel) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *IdentityForestResponseModel) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetTrustedForests

`func (o *IdentityForestResponseModel) GetTrustedForests() []string`

GetTrustedForests returns the TrustedForests field if non-nil, zero value otherwise.

### GetTrustedForestsOk

`func (o *IdentityForestResponseModel) GetTrustedForestsOk() (*[]string, bool)`

GetTrustedForestsOk returns a tuple with the TrustedForests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedForests

`func (o *IdentityForestResponseModel) SetTrustedForests(v []string)`

SetTrustedForests sets TrustedForests field to given value.

### HasTrustedForests

`func (o *IdentityForestResponseModel) HasTrustedForests() bool`

HasTrustedForests returns a boolean if a field has been set.

### SetTrustedForestsNil

`func (o *IdentityForestResponseModel) SetTrustedForestsNil(b bool)`

 SetTrustedForestsNil sets the value for TrustedForests to be an explicit nil

### UnsetTrustedForests
`func (o *IdentityForestResponseModel) UnsetTrustedForests()`

UnsetTrustedForests ensures that no value is present for TrustedForests, not even an explicit nil
### GetPossibleLookupFailure

`func (o *IdentityForestResponseModel) GetPossibleLookupFailure() bool`

GetPossibleLookupFailure returns the PossibleLookupFailure field if non-nil, zero value otherwise.

### GetPossibleLookupFailureOk

`func (o *IdentityForestResponseModel) GetPossibleLookupFailureOk() (*bool, bool)`

GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleLookupFailure

`func (o *IdentityForestResponseModel) SetPossibleLookupFailure(v bool)`

SetPossibleLookupFailure sets PossibleLookupFailure field to given value.

### HasPossibleLookupFailure

`func (o *IdentityForestResponseModel) HasPossibleLookupFailure() bool`

HasPossibleLookupFailure returns a boolean if a field has been set.

### GetPropertiesFetched

`func (o *IdentityForestResponseModel) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *IdentityForestResponseModel) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *IdentityForestResponseModel) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


