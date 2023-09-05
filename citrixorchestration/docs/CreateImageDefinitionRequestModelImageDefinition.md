# CreateImageDefinitionRequestModelImageDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDefinitionName** | **string** | The name associated with the image definition. | 
**ImageDefinitionDescription** | Pointer to **string** | The description associated with the image definition. | [optional] 
**HostingUnit** | **string** | The name or id of the HostingUnit the image definition is held on. | 
**OsType** | **string** | OS Type of the image definition. OsType | 
**UseWriteBackCache** | Pointer to **bool** | Indicates whether or not write back cache is enabled for the VMs created from this image definition. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image definition that are specific to the target hosting infrastructure. | [optional] 

## Methods

### NewCreateImageDefinitionRequestModelImageDefinition

`func NewCreateImageDefinitionRequestModelImageDefinition(imageDefinitionName string, hostingUnit string, osType string, ) *CreateImageDefinitionRequestModelImageDefinition`

NewCreateImageDefinitionRequestModelImageDefinition instantiates a new CreateImageDefinitionRequestModelImageDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageDefinitionRequestModelImageDefinitionWithDefaults

`func NewCreateImageDefinitionRequestModelImageDefinitionWithDefaults() *CreateImageDefinitionRequestModelImageDefinition`

NewCreateImageDefinitionRequestModelImageDefinitionWithDefaults instantiates a new CreateImageDefinitionRequestModelImageDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDefinitionName

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetImageDefinitionName() string`

GetImageDefinitionName returns the ImageDefinitionName field if non-nil, zero value otherwise.

### GetImageDefinitionNameOk

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetImageDefinitionNameOk() (*string, bool)`

GetImageDefinitionNameOk returns a tuple with the ImageDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinitionName

`func (o *CreateImageDefinitionRequestModelImageDefinition) SetImageDefinitionName(v string)`

SetImageDefinitionName sets ImageDefinitionName field to given value.


### GetImageDefinitionDescription

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetImageDefinitionDescription() string`

GetImageDefinitionDescription returns the ImageDefinitionDescription field if non-nil, zero value otherwise.

### GetImageDefinitionDescriptionOk

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetImageDefinitionDescriptionOk() (*string, bool)`

GetImageDefinitionDescriptionOk returns a tuple with the ImageDefinitionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinitionDescription

`func (o *CreateImageDefinitionRequestModelImageDefinition) SetImageDefinitionDescription(v string)`

SetImageDefinitionDescription sets ImageDefinitionDescription field to given value.

### HasImageDefinitionDescription

`func (o *CreateImageDefinitionRequestModelImageDefinition) HasImageDefinitionDescription() bool`

HasImageDefinitionDescription returns a boolean if a field has been set.

### GetHostingUnit

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetHostingUnit() string`

GetHostingUnit returns the HostingUnit field if non-nil, zero value otherwise.

### GetHostingUnitOk

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetHostingUnitOk() (*string, bool)`

GetHostingUnitOk returns a tuple with the HostingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnit

`func (o *CreateImageDefinitionRequestModelImageDefinition) SetHostingUnit(v string)`

SetHostingUnit sets HostingUnit field to given value.


### GetOsType

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CreateImageDefinitionRequestModelImageDefinition) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetUseWriteBackCache

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetUseWriteBackCache() bool`

GetUseWriteBackCache returns the UseWriteBackCache field if non-nil, zero value otherwise.

### GetUseWriteBackCacheOk

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetUseWriteBackCacheOk() (*bool, bool)`

GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWriteBackCache

`func (o *CreateImageDefinitionRequestModelImageDefinition) SetUseWriteBackCache(v bool)`

SetUseWriteBackCache sets UseWriteBackCache field to given value.

### HasUseWriteBackCache

`func (o *CreateImageDefinitionRequestModelImageDefinition) HasUseWriteBackCache() bool`

HasUseWriteBackCache returns a boolean if a field has been set.

### GetCustomProperties

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateImageDefinitionRequestModelImageDefinition) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateImageDefinitionRequestModelImageDefinition) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateImageDefinitionRequestModelImageDefinition) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


