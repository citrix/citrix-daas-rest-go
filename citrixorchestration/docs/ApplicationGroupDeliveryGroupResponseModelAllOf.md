# ApplicationGroupDeliveryGroupResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | **int32** | Specifies the priority of the mapping between application and delivery group where lower numbers imply higher priority with zero being highest. | 
**NumMachines** | **int32** | Number of machines within the delivery group that are capable of hosting the applications in the application group. | 

## Methods

### NewApplicationGroupDeliveryGroupResponseModelAllOf

`func NewApplicationGroupDeliveryGroupResponseModelAllOf(priority int32, numMachines int32, ) *ApplicationGroupDeliveryGroupResponseModelAllOf`

NewApplicationGroupDeliveryGroupResponseModelAllOf instantiates a new ApplicationGroupDeliveryGroupResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGroupDeliveryGroupResponseModelAllOfWithDefaults

`func NewApplicationGroupDeliveryGroupResponseModelAllOfWithDefaults() *ApplicationGroupDeliveryGroupResponseModelAllOf`

NewApplicationGroupDeliveryGroupResponseModelAllOfWithDefaults instantiates a new ApplicationGroupDeliveryGroupResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *ApplicationGroupDeliveryGroupResponseModelAllOf) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApplicationGroupDeliveryGroupResponseModelAllOf) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApplicationGroupDeliveryGroupResponseModelAllOf) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetNumMachines

`func (o *ApplicationGroupDeliveryGroupResponseModelAllOf) GetNumMachines() int32`

GetNumMachines returns the NumMachines field if non-nil, zero value otherwise.

### GetNumMachinesOk

`func (o *ApplicationGroupDeliveryGroupResponseModelAllOf) GetNumMachinesOk() (*int32, bool)`

GetNumMachinesOk returns a tuple with the NumMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachines

`func (o *ApplicationGroupDeliveryGroupResponseModelAllOf) SetNumMachines(v int32)`

SetNumMachines sets NumMachines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


