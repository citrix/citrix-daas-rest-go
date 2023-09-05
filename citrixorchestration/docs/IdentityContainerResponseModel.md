# IdentityContainerResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The Domain the object belongs to. NOTE: When doing a forest level search, the domain property is populated only if the object contains a SID and the SID was requested as property to be fetched. | [optional] 
**Forest** | Pointer to **string** | The forest the object belongs to. | [optional] 
**Guid** | Pointer to **string** | The GUID of the object, the GUID is unique across the enterprise and anywhere else. | [optional] 
**DistinguishedName** | Pointer to **string** | The Distinguished name of the object. Distinguished names (DNs) are unique and they unambiguously identify objects in the directory. | [optional] 
**CanonicalName** | Pointer to **string** | The unique distinguished name of the object in canonical format. | [optional] 
**PossibleLookupFailure** | Pointer to **bool** | Set to True during a bulk identity lookup operation. For individual identity lookup, usually an exception will be thrown. | [optional] 
**DirectoryServer** | Pointer to **string** | Domain controller that performed the add/update operation on this object | [optional] 
**Name** | Pointer to **string** | The name of the container. | [optional] 
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


