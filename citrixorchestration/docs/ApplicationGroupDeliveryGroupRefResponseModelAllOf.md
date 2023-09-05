# ApplicationGroupDeliveryGroupRefResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | Specifies the priority of the mapping between application and delivery group where lower numbers imply higher priority with zero being highest. | [optional] 
**NumMachines** | Pointer to **int32** | Number of machines within the delivery group that are capable of hosting the applications in the application group. | [optional] 
**NumTaggedMachines** | Pointer to **int32** | Number of tagged machines within the delivery group that are capable of hosting the applications in the application group. | [optional] 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) |  | [optional] 

## Methods

### NewApplicationGroupDeliveryGroupRefResponseModelAllOf

`func NewApplicationGroupDeliveryGroupRefResponseModelAllOf() *ApplicationGroupDeliveryGroupRefResponseModelAllOf`

NewApplicationGroupDeliveryGroupRefResponseModelAllOf instantiates a new ApplicationGroupDeliveryGroupRefResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGroupDeliveryGroupRefResponseModelAllOfWithDefaults

`func NewApplicationGroupDeliveryGroupRefResponseModelAllOfWithDefaults() *ApplicationGroupDeliveryGroupRefResponseModelAllOf`

NewApplicationGroupDeliveryGroupRefResponseModelAllOfWithDefaults instantiates a new ApplicationGroupDeliveryGroupRefResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetNumMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) GetNumMachines() int32`

GetNumMachines returns the NumMachines field if non-nil, zero value otherwise.

### GetNumMachinesOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) GetNumMachinesOk() (*int32, bool)`

GetNumMachinesOk returns a tuple with the NumMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) SetNumMachines(v int32)`

SetNumMachines sets NumMachines field to given value.

### HasNumMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) HasNumMachines() bool`

HasNumMachines returns a boolean if a field has been set.

### GetNumTaggedMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) GetNumTaggedMachines() int32`

GetNumTaggedMachines returns the NumTaggedMachines field if non-nil, zero value otherwise.

### GetNumTaggedMachinesOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) GetNumTaggedMachinesOk() (*int32, bool)`

GetNumTaggedMachinesOk returns a tuple with the NumTaggedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTaggedMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) SetNumTaggedMachines(v int32)`

SetNumTaggedMachines sets NumTaggedMachines field to given value.

### HasNumTaggedMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) HasNumTaggedMachines() bool`

HasNumTaggedMachines returns a boolean if a field has been set.

### GetSessionSupport

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *ApplicationGroupDeliveryGroupRefResponseModelAllOf) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


