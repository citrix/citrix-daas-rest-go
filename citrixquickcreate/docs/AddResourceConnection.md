# AddResourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) | The type of provider associated with the account | 
**Name** | **string** | The name of the resource connection to create | 
**ZoneId** | Pointer to **NullableString** | Zone id | [optional] 
**ResourceLocationId** | Pointer to **NullableString** | The ID of the CC resource location  Ignored if ZoneId is specified | [optional] 
**ConnectionId** | Pointer to **NullableString** | The ID of the resource connection to be added | [optional] [readonly] 

## Methods

### NewAddResourceConnection

`func NewAddResourceConnection(accountType AccountType, name string, ) *AddResourceConnection`

NewAddResourceConnection instantiates a new AddResourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddResourceConnectionWithDefaults

`func NewAddResourceConnectionWithDefaults() *AddResourceConnection`

NewAddResourceConnectionWithDefaults instantiates a new AddResourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *AddResourceConnection) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AddResourceConnection) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AddResourceConnection) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetName

`func (o *AddResourceConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddResourceConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddResourceConnection) SetName(v string)`

SetName sets Name field to given value.


### GetZoneId

`func (o *AddResourceConnection) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddResourceConnection) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddResourceConnection) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *AddResourceConnection) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *AddResourceConnection) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *AddResourceConnection) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetResourceLocationId

`func (o *AddResourceConnection) GetResourceLocationId() string`

GetResourceLocationId returns the ResourceLocationId field if non-nil, zero value otherwise.

### GetResourceLocationIdOk

`func (o *AddResourceConnection) GetResourceLocationIdOk() (*string, bool)`

GetResourceLocationIdOk returns a tuple with the ResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocationId

`func (o *AddResourceConnection) SetResourceLocationId(v string)`

SetResourceLocationId sets ResourceLocationId field to given value.

### HasResourceLocationId

`func (o *AddResourceConnection) HasResourceLocationId() bool`

HasResourceLocationId returns a boolean if a field has been set.

### SetResourceLocationIdNil

`func (o *AddResourceConnection) SetResourceLocationIdNil(b bool)`

 SetResourceLocationIdNil sets the value for ResourceLocationId to be an explicit nil

### UnsetResourceLocationId
`func (o *AddResourceConnection) UnsetResourceLocationId()`

UnsetResourceLocationId ensures that no value is present for ResourceLocationId, not even an explicit nil
### GetConnectionId

`func (o *AddResourceConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *AddResourceConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *AddResourceConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *AddResourceConnection) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *AddResourceConnection) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *AddResourceConnection) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


