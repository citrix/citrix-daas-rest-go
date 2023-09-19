# IdentitySiteResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name associated with the site object. | [optional] 
**Guid** | Pointer to **NullableString** | The guid of the site. | [optional] 
**Subnets** | Pointer to [**[]IdentitySubnetResponseModel**](IdentitySubnetResponseModel.md) | The subnets in the site. | [optional] 
**PossibleLookupFailure** | Pointer to **bool** |     For individual identity lookup, usually an exception will be thrown.      | [optional] 
**PropertiesFetched** | **int32** | Properties fetched and populated in the Site object.  This is a bitfield indicating the fetched properties. | 

## Methods

### NewIdentitySiteResponseModel

`func NewIdentitySiteResponseModel(propertiesFetched int32, ) *IdentitySiteResponseModel`

NewIdentitySiteResponseModel instantiates a new IdentitySiteResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySiteResponseModelWithDefaults

`func NewIdentitySiteResponseModelWithDefaults() *IdentitySiteResponseModel`

NewIdentitySiteResponseModelWithDefaults instantiates a new IdentitySiteResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdentitySiteResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentitySiteResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentitySiteResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentitySiteResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IdentitySiteResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdentitySiteResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGuid

`func (o *IdentitySiteResponseModel) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IdentitySiteResponseModel) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IdentitySiteResponseModel) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IdentitySiteResponseModel) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *IdentitySiteResponseModel) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *IdentitySiteResponseModel) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetSubnets

`func (o *IdentitySiteResponseModel) GetSubnets() []IdentitySubnetResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *IdentitySiteResponseModel) GetSubnetsOk() (*[]IdentitySubnetResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *IdentitySiteResponseModel) SetSubnets(v []IdentitySubnetResponseModel)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *IdentitySiteResponseModel) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *IdentitySiteResponseModel) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *IdentitySiteResponseModel) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetPossibleLookupFailure

`func (o *IdentitySiteResponseModel) GetPossibleLookupFailure() bool`

GetPossibleLookupFailure returns the PossibleLookupFailure field if non-nil, zero value otherwise.

### GetPossibleLookupFailureOk

`func (o *IdentitySiteResponseModel) GetPossibleLookupFailureOk() (*bool, bool)`

GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleLookupFailure

`func (o *IdentitySiteResponseModel) SetPossibleLookupFailure(v bool)`

SetPossibleLookupFailure sets PossibleLookupFailure field to given value.

### HasPossibleLookupFailure

`func (o *IdentitySiteResponseModel) HasPossibleLookupFailure() bool`

HasPossibleLookupFailure returns a boolean if a field has been set.

### GetPropertiesFetched

`func (o *IdentitySiteResponseModel) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *IdentitySiteResponseModel) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *IdentitySiteResponseModel) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


