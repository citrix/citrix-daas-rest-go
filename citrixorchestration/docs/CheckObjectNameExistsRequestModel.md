# CheckObjectNameExistsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | [**ObjectType**](ObjectType.md) |  | 
**NameOrPath** | **string** | Name or path of the object to check. | 

## Methods

### NewCheckObjectNameExistsRequestModel

`func NewCheckObjectNameExistsRequestModel(objectType ObjectType, nameOrPath string, ) *CheckObjectNameExistsRequestModel`

NewCheckObjectNameExistsRequestModel instantiates a new CheckObjectNameExistsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckObjectNameExistsRequestModelWithDefaults

`func NewCheckObjectNameExistsRequestModelWithDefaults() *CheckObjectNameExistsRequestModel`

NewCheckObjectNameExistsRequestModelWithDefaults instantiates a new CheckObjectNameExistsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *CheckObjectNameExistsRequestModel) GetObjectType() ObjectType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CheckObjectNameExistsRequestModel) GetObjectTypeOk() (*ObjectType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CheckObjectNameExistsRequestModel) SetObjectType(v ObjectType)`

SetObjectType sets ObjectType field to given value.


### GetNameOrPath

`func (o *CheckObjectNameExistsRequestModel) GetNameOrPath() string`

GetNameOrPath returns the NameOrPath field if non-nil, zero value otherwise.

### GetNameOrPathOk

`func (o *CheckObjectNameExistsRequestModel) GetNameOrPathOk() (*string, bool)`

GetNameOrPathOk returns a tuple with the NameOrPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOrPath

`func (o *CheckObjectNameExistsRequestModel) SetNameOrPath(v string)`

SetNameOrPath sets NameOrPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


