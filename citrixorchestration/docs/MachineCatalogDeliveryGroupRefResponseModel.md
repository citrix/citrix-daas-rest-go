# MachineCatalogDeliveryGroupRefResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the object. Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | [optional] 
**Uid** | Pointer to **NullableInt32** | DEPRECATED. Use Id. | [optional] 
**Name** | Pointer to **NullableString** | Name of the object. | [optional] 
**Priority** | Pointer to **NullableInt32** | RemotePC machine allocation priority for delivery group. | [optional] 
**NumAssigned** | **int32** | Number of machines assigned from the machine catalog into the delivery group. | 
**Total** | **int32** | Total number of machines in the delivery group. | 

## Methods

### NewMachineCatalogDeliveryGroupRefResponseModel

`func NewMachineCatalogDeliveryGroupRefResponseModel(numAssigned int32, total int32, ) *MachineCatalogDeliveryGroupRefResponseModel`

NewMachineCatalogDeliveryGroupRefResponseModel instantiates a new MachineCatalogDeliveryGroupRefResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogDeliveryGroupRefResponseModelWithDefaults

`func NewMachineCatalogDeliveryGroupRefResponseModelWithDefaults() *MachineCatalogDeliveryGroupRefResponseModel`

NewMachineCatalogDeliveryGroupRefResponseModelWithDefaults instantiates a new MachineCatalogDeliveryGroupRefResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineCatalogDeliveryGroupRefResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MachineCatalogDeliveryGroupRefResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MachineCatalogDeliveryGroupRefResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MachineCatalogDeliveryGroupRefResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MachineCatalogDeliveryGroupRefResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MachineCatalogDeliveryGroupRefResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *MachineCatalogDeliveryGroupRefResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *MachineCatalogDeliveryGroupRefResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetName

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineCatalogDeliveryGroupRefResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineCatalogDeliveryGroupRefResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MachineCatalogDeliveryGroupRefResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MachineCatalogDeliveryGroupRefResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPriority

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MachineCatalogDeliveryGroupRefResponseModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MachineCatalogDeliveryGroupRefResponseModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *MachineCatalogDeliveryGroupRefResponseModel) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *MachineCatalogDeliveryGroupRefResponseModel) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetNumAssigned

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetNumAssigned() int32`

GetNumAssigned returns the NumAssigned field if non-nil, zero value otherwise.

### GetNumAssignedOk

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetNumAssignedOk() (*int32, bool)`

GetNumAssignedOk returns a tuple with the NumAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssigned

`func (o *MachineCatalogDeliveryGroupRefResponseModel) SetNumAssigned(v int32)`

SetNumAssigned sets NumAssigned field to given value.


### GetTotal

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MachineCatalogDeliveryGroupRefResponseModel) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MachineCatalogDeliveryGroupRefResponseModel) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


