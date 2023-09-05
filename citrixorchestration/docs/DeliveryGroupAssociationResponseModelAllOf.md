# DeliveryGroupAssociationResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associated** | **bool** | Whether the delivery group was associated. | 
**Priority** | Pointer to **int32** | The priority of delivery group association. | [optional] 

## Methods

### NewDeliveryGroupAssociationResponseModelAllOf

`func NewDeliveryGroupAssociationResponseModelAllOf(associated bool, ) *DeliveryGroupAssociationResponseModelAllOf`

NewDeliveryGroupAssociationResponseModelAllOf instantiates a new DeliveryGroupAssociationResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupAssociationResponseModelAllOfWithDefaults

`func NewDeliveryGroupAssociationResponseModelAllOfWithDefaults() *DeliveryGroupAssociationResponseModelAllOf`

NewDeliveryGroupAssociationResponseModelAllOfWithDefaults instantiates a new DeliveryGroupAssociationResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociated

`func (o *DeliveryGroupAssociationResponseModelAllOf) GetAssociated() bool`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *DeliveryGroupAssociationResponseModelAllOf) GetAssociatedOk() (*bool, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *DeliveryGroupAssociationResponseModelAllOf) SetAssociated(v bool)`

SetAssociated sets Associated field to given value.


### GetPriority

`func (o *DeliveryGroupAssociationResponseModelAllOf) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DeliveryGroupAssociationResponseModelAllOf) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DeliveryGroupAssociationResponseModelAllOf) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DeliveryGroupAssociationResponseModelAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


