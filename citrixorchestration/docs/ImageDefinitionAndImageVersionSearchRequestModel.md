# ImageDefinitionAndImageVersionSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | Pointer to **NullableString** | Resource pool in which to search for the image definition and image version. | [optional] 
**VDASessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) |  | [optional] 
**OsType** | Pointer to [**OsType**](OsType.md) |  | [optional] 

## Methods

### NewImageDefinitionAndImageVersionSearchRequestModel

`func NewImageDefinitionAndImageVersionSearchRequestModel() *ImageDefinitionAndImageVersionSearchRequestModel`

NewImageDefinitionAndImageVersionSearchRequestModel instantiates a new ImageDefinitionAndImageVersionSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDefinitionAndImageVersionSearchRequestModelWithDefaults

`func NewImageDefinitionAndImageVersionSearchRequestModelWithDefaults() *ImageDefinitionAndImageVersionSearchRequestModel`

NewImageDefinitionAndImageVersionSearchRequestModelWithDefaults instantiates a new ImageDefinitionAndImageVersionSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePool

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### SetResourcePoolNil

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) SetResourcePoolNil(b bool)`

 SetResourcePoolNil sets the value for ResourcePool to be an explicit nil

### UnsetResourcePool
`func (o *ImageDefinitionAndImageVersionSearchRequestModel) UnsetResourcePool()`

UnsetResourcePool ensures that no value is present for ResourcePool, not even an explicit nil
### GetVDASessionSupport

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) GetVDASessionSupport() SessionSupport`

GetVDASessionSupport returns the VDASessionSupport field if non-nil, zero value otherwise.

### GetVDASessionSupportOk

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) GetVDASessionSupportOk() (*SessionSupport, bool)`

GetVDASessionSupportOk returns a tuple with the VDASessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDASessionSupport

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) SetVDASessionSupport(v SessionSupport)`

SetVDASessionSupport sets VDASessionSupport field to given value.

### HasVDASessionSupport

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) HasVDASessionSupport() bool`

HasVDASessionSupport returns a boolean if a field has been set.

### GetOsType

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) GetOsType() OsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) GetOsTypeOk() (*OsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) SetOsType(v OsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ImageDefinitionAndImageVersionSearchRequestModel) HasOsType() bool`

HasOsType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


