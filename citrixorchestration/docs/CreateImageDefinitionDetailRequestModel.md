# CreateImageDefinitionDetailRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDefinitionName** | **string** | The name associated with the image definition. | 
**ImageDefinitionDescription** | Pointer to **NullableString** | The description associated with the image definition. | [optional] 
**HostingUnit** | **string** | The name or id of the HostingUnit the image definition is held on. | 
**OsType** | **string** | OS Type of the image definition. OsType | 
**UseWriteBackCache** | Pointer to **NullableBool** | Indicates whether or not write back cache is enabled for the VMs created from this image definition. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image definition that are specific to the target hosting infrastructure. | [optional] 

## Methods

### NewCreateImageDefinitionDetailRequestModel

`func NewCreateImageDefinitionDetailRequestModel(imageDefinitionName string, hostingUnit string, osType string, ) *CreateImageDefinitionDetailRequestModel`

NewCreateImageDefinitionDetailRequestModel instantiates a new CreateImageDefinitionDetailRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageDefinitionDetailRequestModelWithDefaults

`func NewCreateImageDefinitionDetailRequestModelWithDefaults() *CreateImageDefinitionDetailRequestModel`

NewCreateImageDefinitionDetailRequestModelWithDefaults instantiates a new CreateImageDefinitionDetailRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDefinitionName

`func (o *CreateImageDefinitionDetailRequestModel) GetImageDefinitionName() string`

GetImageDefinitionName returns the ImageDefinitionName field if non-nil, zero value otherwise.

### GetImageDefinitionNameOk

`func (o *CreateImageDefinitionDetailRequestModel) GetImageDefinitionNameOk() (*string, bool)`

GetImageDefinitionNameOk returns a tuple with the ImageDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinitionName

`func (o *CreateImageDefinitionDetailRequestModel) SetImageDefinitionName(v string)`

SetImageDefinitionName sets ImageDefinitionName field to given value.


### GetImageDefinitionDescription

`func (o *CreateImageDefinitionDetailRequestModel) GetImageDefinitionDescription() string`

GetImageDefinitionDescription returns the ImageDefinitionDescription field if non-nil, zero value otherwise.

### GetImageDefinitionDescriptionOk

`func (o *CreateImageDefinitionDetailRequestModel) GetImageDefinitionDescriptionOk() (*string, bool)`

GetImageDefinitionDescriptionOk returns a tuple with the ImageDefinitionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinitionDescription

`func (o *CreateImageDefinitionDetailRequestModel) SetImageDefinitionDescription(v string)`

SetImageDefinitionDescription sets ImageDefinitionDescription field to given value.

### HasImageDefinitionDescription

`func (o *CreateImageDefinitionDetailRequestModel) HasImageDefinitionDescription() bool`

HasImageDefinitionDescription returns a boolean if a field has been set.

### SetImageDefinitionDescriptionNil

`func (o *CreateImageDefinitionDetailRequestModel) SetImageDefinitionDescriptionNil(b bool)`

 SetImageDefinitionDescriptionNil sets the value for ImageDefinitionDescription to be an explicit nil

### UnsetImageDefinitionDescription
`func (o *CreateImageDefinitionDetailRequestModel) UnsetImageDefinitionDescription()`

UnsetImageDefinitionDescription ensures that no value is present for ImageDefinitionDescription, not even an explicit nil
### GetHostingUnit

`func (o *CreateImageDefinitionDetailRequestModel) GetHostingUnit() string`

GetHostingUnit returns the HostingUnit field if non-nil, zero value otherwise.

### GetHostingUnitOk

`func (o *CreateImageDefinitionDetailRequestModel) GetHostingUnitOk() (*string, bool)`

GetHostingUnitOk returns a tuple with the HostingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnit

`func (o *CreateImageDefinitionDetailRequestModel) SetHostingUnit(v string)`

SetHostingUnit sets HostingUnit field to given value.


### GetOsType

`func (o *CreateImageDefinitionDetailRequestModel) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CreateImageDefinitionDetailRequestModel) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CreateImageDefinitionDetailRequestModel) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetUseWriteBackCache

`func (o *CreateImageDefinitionDetailRequestModel) GetUseWriteBackCache() bool`

GetUseWriteBackCache returns the UseWriteBackCache field if non-nil, zero value otherwise.

### GetUseWriteBackCacheOk

`func (o *CreateImageDefinitionDetailRequestModel) GetUseWriteBackCacheOk() (*bool, bool)`

GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWriteBackCache

`func (o *CreateImageDefinitionDetailRequestModel) SetUseWriteBackCache(v bool)`

SetUseWriteBackCache sets UseWriteBackCache field to given value.

### HasUseWriteBackCache

`func (o *CreateImageDefinitionDetailRequestModel) HasUseWriteBackCache() bool`

HasUseWriteBackCache returns a boolean if a field has been set.

### SetUseWriteBackCacheNil

`func (o *CreateImageDefinitionDetailRequestModel) SetUseWriteBackCacheNil(b bool)`

 SetUseWriteBackCacheNil sets the value for UseWriteBackCache to be an explicit nil

### UnsetUseWriteBackCache
`func (o *CreateImageDefinitionDetailRequestModel) UnsetUseWriteBackCache()`

UnsetUseWriteBackCache ensures that no value is present for UseWriteBackCache, not even an explicit nil
### GetCustomProperties

`func (o *CreateImageDefinitionDetailRequestModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateImageDefinitionDetailRequestModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateImageDefinitionDetailRequestModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateImageDefinitionDetailRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *CreateImageDefinitionDetailRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *CreateImageDefinitionDetailRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


