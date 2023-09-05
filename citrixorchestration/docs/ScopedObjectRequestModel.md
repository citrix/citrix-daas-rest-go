# ScopedObjectRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | [**ScopedObjectType**](ScopedObjectType.md) |  | 
**Object** | **string** | Name or Id of the scoped object. | 

## Methods

### NewScopedObjectRequestModel

`func NewScopedObjectRequestModel(objectType ScopedObjectType, object string, ) *ScopedObjectRequestModel`

NewScopedObjectRequestModel instantiates a new ScopedObjectRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedObjectRequestModelWithDefaults

`func NewScopedObjectRequestModelWithDefaults() *ScopedObjectRequestModel`

NewScopedObjectRequestModelWithDefaults instantiates a new ScopedObjectRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *ScopedObjectRequestModel) GetObjectType() ScopedObjectType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ScopedObjectRequestModel) GetObjectTypeOk() (*ScopedObjectType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ScopedObjectRequestModel) SetObjectType(v ScopedObjectType)`

SetObjectType sets ObjectType field to given value.


### GetObject

`func (o *ScopedObjectRequestModel) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ScopedObjectRequestModel) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ScopedObjectRequestModel) SetObject(v string)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


