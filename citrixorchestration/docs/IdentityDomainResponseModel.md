# IdentityDomainResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forest** | Pointer to **string** | The forest the domain belongs to. | [optional] 
**Parent** | Pointer to **string** | The parent of the domain. | [optional] 
**Name** | Pointer to **string** | The name of the domain. | [optional] 
**Children** | Pointer to **[]string** | The sub-domains associated with the Domain object. | [optional] 
**Sid** | Pointer to **string** | The Security Identifier associated with the Domain object. | [optional] 
**Guid** | Pointer to **string** | The GUID is unique across the enterprise and anywhere else. | [optional] 
**NetBiosName** | Pointer to **string** | The NETBiosName associated with the Domain object. | [optional] 
**DistinguishedName** | Pointer to **string** | The Distinguished name of the Domain object. Distinguished names (DNs) are unique and they unambiguously identify objects in the directory. | [optional] 
**Controllers** | Pointer to **[]string** | The discoverable controllers/directory servers for the domain. | [optional] 
**DefaultController** | Pointer to **string** | The default controller/directory server for the domain. | [optional] 
**TrustedDomains** | Pointer to **[]string** | Other Domains that have trusts with this Domain. | [optional] 
**UpnSuffixes** | Pointer to **[]string** | The UPN suffixes of the domain. | [optional] 
**PossibleLookupFailure** | Pointer to **bool** | Set to True during a bulk identity lookup operation. For individual identity lookup, usually an exception will be throw. | [optional] 
**PropertiesFetched** | **int32** | Properties fetched and populated in the Domain object.  This is a bitfield indicating the fetched properties. | 

## Methods

### NewIdentityDomainResponseModel

`func NewIdentityDomainResponseModel(propertiesFetched int32, ) *IdentityDomainResponseModel`

NewIdentityDomainResponseModel instantiates a new IdentityDomainResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityDomainResponseModelWithDefaults

`func NewIdentityDomainResponseModelWithDefaults() *IdentityDomainResponseModel`

NewIdentityDomainResponseModelWithDefaults instantiates a new IdentityDomainResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForest

`func (o *IdentityDomainResponseModel) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *IdentityDomainResponseModel) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *IdentityDomainResponseModel) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *IdentityDomainResponseModel) HasForest() bool`

HasForest returns a boolean if a field has been set.

### GetParent

`func (o *IdentityDomainResponseModel) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IdentityDomainResponseModel) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IdentityDomainResponseModel) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IdentityDomainResponseModel) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetName

`func (o *IdentityDomainResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityDomainResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityDomainResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityDomainResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChildren

`func (o *IdentityDomainResponseModel) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *IdentityDomainResponseModel) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *IdentityDomainResponseModel) SetChildren(v []string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *IdentityDomainResponseModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetSid

`func (o *IdentityDomainResponseModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IdentityDomainResponseModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IdentityDomainResponseModel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IdentityDomainResponseModel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetGuid

`func (o *IdentityDomainResponseModel) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IdentityDomainResponseModel) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IdentityDomainResponseModel) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IdentityDomainResponseModel) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetNetBiosName

`func (o *IdentityDomainResponseModel) GetNetBiosName() string`

GetNetBiosName returns the NetBiosName field if non-nil, zero value otherwise.

### GetNetBiosNameOk

`func (o *IdentityDomainResponseModel) GetNetBiosNameOk() (*string, bool)`

GetNetBiosNameOk returns a tuple with the NetBiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetBiosName

`func (o *IdentityDomainResponseModel) SetNetBiosName(v string)`

SetNetBiosName sets NetBiosName field to given value.

### HasNetBiosName

`func (o *IdentityDomainResponseModel) HasNetBiosName() bool`

HasNetBiosName returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *IdentityDomainResponseModel) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *IdentityDomainResponseModel) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *IdentityDomainResponseModel) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *IdentityDomainResponseModel) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetControllers

`func (o *IdentityDomainResponseModel) GetControllers() []string`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *IdentityDomainResponseModel) GetControllersOk() (*[]string, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *IdentityDomainResponseModel) SetControllers(v []string)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *IdentityDomainResponseModel) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetDefaultController

`func (o *IdentityDomainResponseModel) GetDefaultController() string`

GetDefaultController returns the DefaultController field if non-nil, zero value otherwise.

### GetDefaultControllerOk

`func (o *IdentityDomainResponseModel) GetDefaultControllerOk() (*string, bool)`

GetDefaultControllerOk returns a tuple with the DefaultController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultController

`func (o *IdentityDomainResponseModel) SetDefaultController(v string)`

SetDefaultController sets DefaultController field to given value.

### HasDefaultController

`func (o *IdentityDomainResponseModel) HasDefaultController() bool`

HasDefaultController returns a boolean if a field has been set.

### GetTrustedDomains

`func (o *IdentityDomainResponseModel) GetTrustedDomains() []string`

GetTrustedDomains returns the TrustedDomains field if non-nil, zero value otherwise.

### GetTrustedDomainsOk

`func (o *IdentityDomainResponseModel) GetTrustedDomainsOk() (*[]string, bool)`

GetTrustedDomainsOk returns a tuple with the TrustedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomains

`func (o *IdentityDomainResponseModel) SetTrustedDomains(v []string)`

SetTrustedDomains sets TrustedDomains field to given value.

### HasTrustedDomains

`func (o *IdentityDomainResponseModel) HasTrustedDomains() bool`

HasTrustedDomains returns a boolean if a field has been set.

### GetUpnSuffixes

`func (o *IdentityDomainResponseModel) GetUpnSuffixes() []string`

GetUpnSuffixes returns the UpnSuffixes field if non-nil, zero value otherwise.

### GetUpnSuffixesOk

`func (o *IdentityDomainResponseModel) GetUpnSuffixesOk() (*[]string, bool)`

GetUpnSuffixesOk returns a tuple with the UpnSuffixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpnSuffixes

`func (o *IdentityDomainResponseModel) SetUpnSuffixes(v []string)`

SetUpnSuffixes sets UpnSuffixes field to given value.

### HasUpnSuffixes

`func (o *IdentityDomainResponseModel) HasUpnSuffixes() bool`

HasUpnSuffixes returns a boolean if a field has been set.

### GetPossibleLookupFailure

`func (o *IdentityDomainResponseModel) GetPossibleLookupFailure() bool`

GetPossibleLookupFailure returns the PossibleLookupFailure field if non-nil, zero value otherwise.

### GetPossibleLookupFailureOk

`func (o *IdentityDomainResponseModel) GetPossibleLookupFailureOk() (*bool, bool)`

GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleLookupFailure

`func (o *IdentityDomainResponseModel) SetPossibleLookupFailure(v bool)`

SetPossibleLookupFailure sets PossibleLookupFailure field to given value.

### HasPossibleLookupFailure

`func (o *IdentityDomainResponseModel) HasPossibleLookupFailure() bool`

HasPossibleLookupFailure returns a boolean if a field has been set.

### GetPropertiesFetched

`func (o *IdentityDomainResponseModel) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *IdentityDomainResponseModel) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *IdentityDomainResponseModel) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


