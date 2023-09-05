# MachineCatalogDeliveryGroupRefResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | RemotePC machine allocation priority for delivery group. | [optional] 
**NumAssigned** | **int32** | Number of machines assigned from the machine catalog into the delivery group. | 
**Total** | **int32** | Total number of machines in the delivery group. | 

## Methods

### NewMachineCatalogDeliveryGroupRefResponseModelAllOf

`func NewMachineCatalogDeliveryGroupRefResponseModelAllOf(numAssigned int32, total int32, ) *MachineCatalogDeliveryGroupRefResponseModelAllOf`

NewMachineCatalogDeliveryGroupRefResponseModelAllOf instantiates a new MachineCatalogDeliveryGroupRefResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogDeliveryGroupRefResponseModelAllOfWithDefaults

`func NewMachineCatalogDeliveryGroupRefResponseModelAllOfWithDefaults() *MachineCatalogDeliveryGroupRefResponseModelAllOf`

NewMachineCatalogDeliveryGroupRefResponseModelAllOfWithDefaults instantiates a new MachineCatalogDeliveryGroupRefResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *MachineCatalogDeliveryGroupRefResponseModelAllOf) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MachineCatalogDeliveryGroupRefResponseModelAllOf) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MachineCatalogDeliveryGroupRefResponseModelAllOf) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MachineCatalogDeliveryGroupRefResponseModelAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetNumAssigned

`func (o *MachineCatalogDeliveryGroupRefResponseModelAllOf) GetNumAssigned() int32`

GetNumAssigned returns the NumAssigned field if non-nil, zero value otherwise.

### GetNumAssignedOk

`func (o *MachineCatalogDeliveryGroupRefResponseModelAllOf) GetNumAssignedOk() (*int32, bool)`

GetNumAssignedOk returns a tuple with the NumAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssigned

`func (o *MachineCatalogDeliveryGroupRefResponseModelAllOf) SetNumAssigned(v int32)`

SetNumAssigned sets NumAssigned field to given value.


### GetTotal

`func (o *MachineCatalogDeliveryGroupRefResponseModelAllOf) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MachineCatalogDeliveryGroupRefResponseModelAllOf) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MachineCatalogDeliveryGroupRefResponseModelAllOf) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


