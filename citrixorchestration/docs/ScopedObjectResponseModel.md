# ScopedObjectResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | [**ScopedObjectResponseModelObject**](ScopedObjectResponseModelObject.md) |  | 
**ObjectDescription** | Pointer to **string** | Description of the scoped object, if it has one. | [optional] 
**ObjectType** | [**ScopedObjectType**](ScopedObjectType.md) |  | 

## Methods

### NewScopedObjectResponseModel

`func NewScopedObjectResponseModel(object ScopedObjectResponseModelObject, objectType ScopedObjectType, ) *ScopedObjectResponseModel`

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

`func (o *ScopedObjectResponseModel) GetObject() ScopedObjectResponseModelObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ScopedObjectResponseModel) GetObjectOk() (*ScopedObjectResponseModelObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ScopedObjectResponseModel) SetObject(v ScopedObjectResponseModelObject)`

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


