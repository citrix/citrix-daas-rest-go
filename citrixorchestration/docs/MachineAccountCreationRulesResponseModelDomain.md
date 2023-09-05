# MachineAccountCreationRulesResponseModelDomain

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

### NewMachineAccountCreationRulesResponseModelDomain

`func NewMachineAccountCreationRulesResponseModelDomain(propertiesFetched int32, ) *MachineAccountCreationRulesResponseModelDomain`

NewMachineAccountCreationRulesResponseModelDomain instantiates a new MachineAccountCreationRulesResponseModelDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAccountCreationRulesResponseModelDomainWithDefaults

`func NewMachineAccountCreationRulesResponseModelDomainWithDefaults() *MachineAccountCreationRulesResponseModelDomain`

NewMachineAccountCreationRulesResponseModelDomainWithDefaults instantiates a new MachineAccountCreationRulesResponseModelDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForest

`func (o *MachineAccountCreationRulesResponseModelDomain) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *MachineAccountCreationRulesResponseModelDomain) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *MachineAccountCreationRulesResponseModelDomain) HasForest() bool`

HasForest returns a boolean if a field has been set.

### GetParent

`func (o *MachineAccountCreationRulesResponseModelDomain) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *MachineAccountCreationRulesResponseModelDomain) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *MachineAccountCreationRulesResponseModelDomain) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetName

`func (o *MachineAccountCreationRulesResponseModelDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineAccountCreationRulesResponseModelDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineAccountCreationRulesResponseModelDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChildren

`func (o *MachineAccountCreationRulesResponseModelDomain) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MachineAccountCreationRulesResponseModelDomain) SetChildren(v []string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *MachineAccountCreationRulesResponseModelDomain) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetSid

`func (o *MachineAccountCreationRulesResponseModelDomain) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MachineAccountCreationRulesResponseModelDomain) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *MachineAccountCreationRulesResponseModelDomain) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetGuid

`func (o *MachineAccountCreationRulesResponseModelDomain) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *MachineAccountCreationRulesResponseModelDomain) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *MachineAccountCreationRulesResponseModelDomain) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetNetBiosName

`func (o *MachineAccountCreationRulesResponseModelDomain) GetNetBiosName() string`

GetNetBiosName returns the NetBiosName field if non-nil, zero value otherwise.

### GetNetBiosNameOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetNetBiosNameOk() (*string, bool)`

GetNetBiosNameOk returns a tuple with the NetBiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetBiosName

`func (o *MachineAccountCreationRulesResponseModelDomain) SetNetBiosName(v string)`

SetNetBiosName sets NetBiosName field to given value.

### HasNetBiosName

`func (o *MachineAccountCreationRulesResponseModelDomain) HasNetBiosName() bool`

HasNetBiosName returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *MachineAccountCreationRulesResponseModelDomain) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *MachineAccountCreationRulesResponseModelDomain) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *MachineAccountCreationRulesResponseModelDomain) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetControllers

`func (o *MachineAccountCreationRulesResponseModelDomain) GetControllers() []string`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetControllersOk() (*[]string, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *MachineAccountCreationRulesResponseModelDomain) SetControllers(v []string)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *MachineAccountCreationRulesResponseModelDomain) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetDefaultController

`func (o *MachineAccountCreationRulesResponseModelDomain) GetDefaultController() string`

GetDefaultController returns the DefaultController field if non-nil, zero value otherwise.

### GetDefaultControllerOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetDefaultControllerOk() (*string, bool)`

GetDefaultControllerOk returns a tuple with the DefaultController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultController

`func (o *MachineAccountCreationRulesResponseModelDomain) SetDefaultController(v string)`

SetDefaultController sets DefaultController field to given value.

### HasDefaultController

`func (o *MachineAccountCreationRulesResponseModelDomain) HasDefaultController() bool`

HasDefaultController returns a boolean if a field has been set.

### GetTrustedDomains

`func (o *MachineAccountCreationRulesResponseModelDomain) GetTrustedDomains() []string`

GetTrustedDomains returns the TrustedDomains field if non-nil, zero value otherwise.

### GetTrustedDomainsOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetTrustedDomainsOk() (*[]string, bool)`

GetTrustedDomainsOk returns a tuple with the TrustedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDomains

`func (o *MachineAccountCreationRulesResponseModelDomain) SetTrustedDomains(v []string)`

SetTrustedDomains sets TrustedDomains field to given value.

### HasTrustedDomains

`func (o *MachineAccountCreationRulesResponseModelDomain) HasTrustedDomains() bool`

HasTrustedDomains returns a boolean if a field has been set.

### GetUpnSuffixes

`func (o *MachineAccountCreationRulesResponseModelDomain) GetUpnSuffixes() []string`

GetUpnSuffixes returns the UpnSuffixes field if non-nil, zero value otherwise.

### GetUpnSuffixesOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetUpnSuffixesOk() (*[]string, bool)`

GetUpnSuffixesOk returns a tuple with the UpnSuffixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpnSuffixes

`func (o *MachineAccountCreationRulesResponseModelDomain) SetUpnSuffixes(v []string)`

SetUpnSuffixes sets UpnSuffixes field to given value.

### HasUpnSuffixes

`func (o *MachineAccountCreationRulesResponseModelDomain) HasUpnSuffixes() bool`

HasUpnSuffixes returns a boolean if a field has been set.

### GetPossibleLookupFailure

`func (o *MachineAccountCreationRulesResponseModelDomain) GetPossibleLookupFailure() bool`

GetPossibleLookupFailure returns the PossibleLookupFailure field if non-nil, zero value otherwise.

### GetPossibleLookupFailureOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetPossibleLookupFailureOk() (*bool, bool)`

GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleLookupFailure

`func (o *MachineAccountCreationRulesResponseModelDomain) SetPossibleLookupFailure(v bool)`

SetPossibleLookupFailure sets PossibleLookupFailure field to given value.

### HasPossibleLookupFailure

`func (o *MachineAccountCreationRulesResponseModelDomain) HasPossibleLookupFailure() bool`

HasPossibleLookupFailure returns a boolean if a field has been set.

### GetPropertiesFetched

`func (o *MachineAccountCreationRulesResponseModelDomain) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *MachineAccountCreationRulesResponseModelDomain) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *MachineAccountCreationRulesResponseModelDomain) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


