# ApplicationGroupDeliveryGroupRefResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the object. Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | [optional] 
**Uid** | Pointer to **NullableInt32** | DEPRECATED. Use Id. | [optional] 
**Name** | Pointer to **NullableString** | Name of the object. | [optional] 
**Priority** | Pointer to **int32** | Specifies the priority of the mapping between application and delivery group where lower numbers imply higher priority with zero being highest. | [optional] 
**NumMachines** | Pointer to **int32** | Number of machines within the delivery group that are capable of hosting the applications in the application group. | [optional] 
**NumTaggedMachines** | Pointer to **int32** | Number of tagged machines within the delivery group that are capable of hosting the applications in the application group. | [optional] 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) |  | [optional] 

## Methods

### NewApplicationGroupDeliveryGroupRefResponseModel

`func NewApplicationGroupDeliveryGroupRefResponseModel() *ApplicationGroupDeliveryGroupRefResponseModel`

NewApplicationGroupDeliveryGroupRefResponseModel instantiates a new ApplicationGroupDeliveryGroupRefResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGroupDeliveryGroupRefResponseModelWithDefaults

`func NewApplicationGroupDeliveryGroupRefResponseModelWithDefaults() *ApplicationGroupDeliveryGroupRefResponseModel`

NewApplicationGroupDeliveryGroupRefResponseModelWithDefaults instantiates a new ApplicationGroupDeliveryGroupRefResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ApplicationGroupDeliveryGroupRefResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ApplicationGroupDeliveryGroupRefResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetName

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApplicationGroupDeliveryGroupRefResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPriority

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetNumMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetNumMachines() int32`

GetNumMachines returns the NumMachines field if non-nil, zero value otherwise.

### GetNumMachinesOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetNumMachinesOk() (*int32, bool)`

GetNumMachinesOk returns a tuple with the NumMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) SetNumMachines(v int32)`

SetNumMachines sets NumMachines field to given value.

### HasNumMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) HasNumMachines() bool`

HasNumMachines returns a boolean if a field has been set.

### GetNumTaggedMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetNumTaggedMachines() int32`

GetNumTaggedMachines returns the NumTaggedMachines field if non-nil, zero value otherwise.

### GetNumTaggedMachinesOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetNumTaggedMachinesOk() (*int32, bool)`

GetNumTaggedMachinesOk returns a tuple with the NumTaggedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTaggedMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) SetNumTaggedMachines(v int32)`

SetNumTaggedMachines sets NumTaggedMachines field to given value.

### HasNumTaggedMachines

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) HasNumTaggedMachines() bool`

HasNumTaggedMachines returns a boolean if a field has been set.

### GetSessionSupport

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *ApplicationGroupDeliveryGroupRefResponseModel) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


