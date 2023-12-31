# DeliveryGroupAssociationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the object. Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | [optional] 
**Uid** | Pointer to **NullableInt32** | DEPRECATED. Use Id. | [optional] 
**Name** | Pointer to **NullableString** | Name of the object. | [optional] 
**Associated** | **bool** | Whether the delivery group was associated. | 
**Priority** | Pointer to **NullableInt32** | The priority of delivery group association. | [optional] 

## Methods

### NewDeliveryGroupAssociationResponseModel

`func NewDeliveryGroupAssociationResponseModel(associated bool, ) *DeliveryGroupAssociationResponseModel`

NewDeliveryGroupAssociationResponseModel instantiates a new DeliveryGroupAssociationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupAssociationResponseModelWithDefaults

`func NewDeliveryGroupAssociationResponseModelWithDefaults() *DeliveryGroupAssociationResponseModel`

NewDeliveryGroupAssociationResponseModelWithDefaults instantiates a new DeliveryGroupAssociationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliveryGroupAssociationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryGroupAssociationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryGroupAssociationResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeliveryGroupAssociationResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeliveryGroupAssociationResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeliveryGroupAssociationResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *DeliveryGroupAssociationResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DeliveryGroupAssociationResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DeliveryGroupAssociationResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DeliveryGroupAssociationResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *DeliveryGroupAssociationResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *DeliveryGroupAssociationResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetName

`func (o *DeliveryGroupAssociationResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeliveryGroupAssociationResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeliveryGroupAssociationResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeliveryGroupAssociationResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeliveryGroupAssociationResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeliveryGroupAssociationResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAssociated

`func (o *DeliveryGroupAssociationResponseModel) GetAssociated() bool`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *DeliveryGroupAssociationResponseModel) GetAssociatedOk() (*bool, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *DeliveryGroupAssociationResponseModel) SetAssociated(v bool)`

SetAssociated sets Associated field to given value.


### GetPriority

`func (o *DeliveryGroupAssociationResponseModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DeliveryGroupAssociationResponseModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DeliveryGroupAssociationResponseModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DeliveryGroupAssociationResponseModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *DeliveryGroupAssociationResponseModel) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *DeliveryGroupAssociationResponseModel) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


