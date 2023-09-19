# ResourceValidationViolationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**ResourceViolationLevel**](ResourceViolationLevel.md) |  | 
**Message** | **string** | A violation message. | 
**Detail** | Pointer to **NullableString** | The attached detail, could be null. | [optional] 
**RelativePath** | Pointer to **NullableString** | The relative path of the resource which owns this violation, could be null. | [optional] 

## Methods

### NewResourceValidationViolationModel

`func NewResourceValidationViolationModel(level ResourceViolationLevel, message string, ) *ResourceValidationViolationModel`

NewResourceValidationViolationModel instantiates a new ResourceValidationViolationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceValidationViolationModelWithDefaults

`func NewResourceValidationViolationModelWithDefaults() *ResourceValidationViolationModel`

NewResourceValidationViolationModelWithDefaults instantiates a new ResourceValidationViolationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ResourceValidationViolationModel) GetLevel() ResourceViolationLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ResourceValidationViolationModel) GetLevelOk() (*ResourceViolationLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ResourceValidationViolationModel) SetLevel(v ResourceViolationLevel)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *ResourceValidationViolationModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceValidationViolationModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceValidationViolationModel) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetail

`func (o *ResourceValidationViolationModel) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResourceValidationViolationModel) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResourceValidationViolationModel) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResourceValidationViolationModel) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ResourceValidationViolationModel) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ResourceValidationViolationModel) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetRelativePath

`func (o *ResourceValidationViolationModel) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *ResourceValidationViolationModel) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *ResourceValidationViolationModel) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *ResourceValidationViolationModel) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### SetRelativePathNil

`func (o *ResourceValidationViolationModel) SetRelativePathNil(b bool)`

 SetRelativePathNil sets the value for RelativePath to be an explicit nil

### UnsetRelativePath
`func (o *ResourceValidationViolationModel) UnsetRelativePath()`

UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


