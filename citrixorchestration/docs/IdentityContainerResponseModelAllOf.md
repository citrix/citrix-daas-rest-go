# IdentityContainerResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the container. | [optional] 
**Type** | Pointer to [**IdentityContainerType**](IdentityContainerType.md) |  | [optional] 
**PropertiesFetched** | **int32** | Properties fetched and populated in the Container object.  This is a bitfield indicating the fetched properties. | 

## Methods

### NewIdentityContainerResponseModelAllOf

`func NewIdentityContainerResponseModelAllOf(propertiesFetched int32, ) *IdentityContainerResponseModelAllOf`

NewIdentityContainerResponseModelAllOf instantiates a new IdentityContainerResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityContainerResponseModelAllOfWithDefaults

`func NewIdentityContainerResponseModelAllOfWithDefaults() *IdentityContainerResponseModelAllOf`

NewIdentityContainerResponseModelAllOfWithDefaults instantiates a new IdentityContainerResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdentityContainerResponseModelAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityContainerResponseModelAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityContainerResponseModelAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityContainerResponseModelAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IdentityContainerResponseModelAllOf) GetType() IdentityContainerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityContainerResponseModelAllOf) GetTypeOk() (*IdentityContainerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityContainerResponseModelAllOf) SetType(v IdentityContainerType)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityContainerResponseModelAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPropertiesFetched

`func (o *IdentityContainerResponseModelAllOf) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *IdentityContainerResponseModelAllOf) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *IdentityContainerResponseModelAllOf) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


