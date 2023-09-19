# IdentityContainerResponseModel

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
**Name** | Pointer to **NullableString** | The name of the container. | [optional] 
**Type** | Pointer to [**IdentityContainerType**](IdentityContainerType.md) |  | [optional] 
**PropertiesFetched** | **int32** | Properties fetched and populated in the Container object.  This is a bitfield indicating the fetched properties. | 

## Methods

### NewIdentityContainerResponseModel

`func NewIdentityContainerResponseModel(propertiesFetched int32, ) *IdentityContainerResponseModel`

NewIdentityContainerResponseModel instantiates a new IdentityContainerResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityContainerResponseModelWithDefaults

`func NewIdentityContainerResponseModelWithDefaults() *IdentityContainerResponseModel`

NewIdentityContainerResponseModelWithDefaults instantiates a new IdentityContainerResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *IdentityContainerResponseModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentityContainerResponseModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentityContainerResponseModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdentityContainerResponseModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *IdentityContainerResponseModel) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *IdentityContainerResponseModel) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetForest

`func (o *IdentityContainerResponseModel) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *IdentityContainerResponseModel) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *IdentityContainerResponseModel) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *IdentityContainerResponseModel) HasForest() bool`

HasForest returns a boolean if a field has been set.

### SetForestNil

`func (o *IdentityContainerResponseModel) SetForestNil(b bool)`

 SetForestNil sets the value for Forest to be an explicit nil

### UnsetForest
`func (o *IdentityContainerResponseModel) UnsetForest()`

UnsetForest ensures that no value is present for Forest, not even an explicit nil
### GetGuid

`func (o *IdentityContainerResponseModel) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IdentityContainerResponseModel) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IdentityContainerResponseModel) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IdentityContainerResponseModel) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *IdentityContainerResponseModel) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *IdentityContainerResponseModel) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetDistinguishedName

`func (o *IdentityContainerResponseModel) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *IdentityContainerResponseModel) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *IdentityContainerResponseModel) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *IdentityContainerResponseModel) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *IdentityContainerResponseModel) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *IdentityContainerResponseModel) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetCanonicalName

`func (o *IdentityContainerResponseModel) GetCanonicalName() string`

GetCanonicalName returns the CanonicalName field if non-nil, zero value otherwise.

### GetCanonicalNameOk

`func (o *IdentityContainerResponseModel) GetCanonicalNameOk() (*string, bool)`

GetCanonicalNameOk returns a tuple with the CanonicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalName

`func (o *IdentityContainerResponseModel) SetCanonicalName(v string)`

SetCanonicalName sets CanonicalName field to given value.

### HasCanonicalName

`func (o *IdentityContainerResponseModel) HasCanonicalName() bool`

HasCanonicalName returns a boolean if a field has been set.

### SetCanonicalNameNil

`func (o *IdentityContainerResponseModel) SetCanonicalNameNil(b bool)`

 SetCanonicalNameNil sets the value for CanonicalName to be an explicit nil

### UnsetCanonicalName
`func (o *IdentityContainerResponseModel) UnsetCanonicalName()`

UnsetCanonicalName ensures that no value is present for CanonicalName, not even an explicit nil
### GetPossibleLookupFailure

`func (o *IdentityContainerResponseModel) GetPossibleLookupFailure() bool`

GetPossibleLookupFailure returns the PossibleLookupFailure field if non-nil, zero value otherwise.

### GetPossibleLookupFailureOk

`func (o *IdentityContainerResponseModel) GetPossibleLookupFailureOk() (*bool, bool)`

GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleLookupFailure

`func (o *IdentityContainerResponseModel) SetPossibleLookupFailure(v bool)`

SetPossibleLookupFailure sets PossibleLookupFailure field to given value.

### HasPossibleLookupFailure

`func (o *IdentityContainerResponseModel) HasPossibleLookupFailure() bool`

HasPossibleLookupFailure returns a boolean if a field has been set.

### GetDirectoryServer

`func (o *IdentityContainerResponseModel) GetDirectoryServer() string`

GetDirectoryServer returns the DirectoryServer field if non-nil, zero value otherwise.

### GetDirectoryServerOk

`func (o *IdentityContainerResponseModel) GetDirectoryServerOk() (*string, bool)`

GetDirectoryServerOk returns a tuple with the DirectoryServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServer

`func (o *IdentityContainerResponseModel) SetDirectoryServer(v string)`

SetDirectoryServer sets DirectoryServer field to given value.

### HasDirectoryServer

`func (o *IdentityContainerResponseModel) HasDirectoryServer() bool`

HasDirectoryServer returns a boolean if a field has been set.

### SetDirectoryServerNil

`func (o *IdentityContainerResponseModel) SetDirectoryServerNil(b bool)`

 SetDirectoryServerNil sets the value for DirectoryServer to be an explicit nil

### UnsetDirectoryServer
`func (o *IdentityContainerResponseModel) UnsetDirectoryServer()`

UnsetDirectoryServer ensures that no value is present for DirectoryServer, not even an explicit nil
### GetName

`func (o *IdentityContainerResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityContainerResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityContainerResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityContainerResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IdentityContainerResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdentityContainerResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *IdentityContainerResponseModel) GetType() IdentityContainerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityContainerResponseModel) GetTypeOk() (*IdentityContainerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityContainerResponseModel) SetType(v IdentityContainerType)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityContainerResponseModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPropertiesFetched

`func (o *IdentityContainerResponseModel) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *IdentityContainerResponseModel) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *IdentityContainerResponseModel) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


