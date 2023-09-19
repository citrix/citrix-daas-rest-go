# ScopedObjectResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | [**RefResponseModel**](RefResponseModel.md) |  | 
**ObjectDescription** | Pointer to **NullableString** | Description of the scoped object, if it has one. | [optional] 
**ObjectType** | [**ScopedObjectType**](ScopedObjectType.md) |  | 

## Methods

### NewScopedObjectResponseModel

`func NewScopedObjectResponseModel(object RefResponseModel, objectType ScopedObjectType, ) *ScopedObjectResponseModel`

NewScopedObjectResponseModel instantiates a new ScopedObjectResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedObjectResponseModelWithDefaults

`func NewScopedObjectResponseModelWithDefaults() *ScopedObjectResponseModel`

NewScopedObjectResponseModelWithDefaults instantiates a new ScopedObjectResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ScopedObjectResponseModel) GetObject() RefResponseModel`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ScopedObjectResponseModel) GetObjectOk() (*RefResponseModel, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ScopedObjectResponseModel) SetObject(v RefResponseModel)`

SetObject sets Object field to given value.


### GetObjectDescription

`func (o *ScopedObjectResponseModel) GetObjectDescription() string`

GetObjectDescription returns the ObjectDescription field if non-nil, zero value otherwise.

### GetObjectDescriptionOk

`func (o *ScopedObjectResponseModel) GetObjectDescriptionOk() (*string, bool)`

GetObjectDescriptionOk returns a tuple with the ObjectDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDescription

`func (o *ScopedObjectResponseModel) SetObjectDescription(v string)`

SetObjectDescription sets ObjectDescription field to given value.

### HasObjectDescription

`func (o *ScopedObjectResponseModel) HasObjectDescription() bool`

HasObjectDescription returns a boolean if a field has been set.

### SetObjectDescriptionNil

`func (o *ScopedObjectResponseModel) SetObjectDescriptionNil(b bool)`

 SetObjectDescriptionNil sets the value for ObjectDescription to be an explicit nil

### UnsetObjectDescription
`func (o *ScopedObjectResponseModel) UnsetObjectDescription()`

UnsetObjectDescription ensures that no value is present for ObjectDescription, not even an explicit nil
### GetObjectType

`func (o *ScopedObjectResponseModel) GetObjectType() ScopedObjectType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ScopedObjectResponseModel) GetObjectTypeOk() (*ScopedObjectType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ScopedObjectResponseModel) SetObjectType(v ScopedObjectType)`

SetObjectType sets ObjectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


