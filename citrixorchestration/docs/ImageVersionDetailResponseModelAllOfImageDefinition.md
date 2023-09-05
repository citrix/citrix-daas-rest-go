# ImageVersionDetailResponseModelAllOfImageDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name associated with the image definition. | [optional] 
**Id** | Pointer to **string** | The Id of the image definition. | [optional] 
**OsType** | Pointer to **string** | The os type of image definition. | [optional] 
**HostingUnitName** | Pointer to **string** | The hosting unit name of image definition. | [optional] 
**HostingUnitUid** | Pointer to **string** | The hosting unit id of image definition. | [optional] 
**Description** | Pointer to **string** | The description of the image definition. | [optional] 
**UseWriteBackCache** | Pointer to **bool** | The UseWriteBackCache of the image definition. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The custom properties of the image definition. Currently only resource group. | [optional] 

## Methods

### NewImageVersionDetailResponseModelAllOfImageDefinition

`func NewImageVersionDetailResponseModelAllOfImageDefinition() *ImageVersionDetailResponseModelAllOfImageDefinition`

NewImageVersionDetailResponseModelAllOfImageDefinition instantiates a new ImageVersionDetailResponseModelAllOfImageDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionDetailResponseModelAllOfImageDefinitionWithDefaults

`func NewImageVersionDetailResponseModelAllOfImageDefinitionWithDefaults() *ImageVersionDetailResponseModelAllOfImageDefinition`

NewImageVersionDetailResponseModelAllOfImageDefinitionWithDefaults instantiates a new ImageVersionDetailResponseModelAllOfImageDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOsType

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetHostingUnitName

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetHostingUnitName() string`

GetHostingUnitName returns the HostingUnitName field if non-nil, zero value otherwise.

### GetHostingUnitNameOk

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetHostingUnitNameOk() (*string, bool)`

GetHostingUnitNameOk returns a tuple with the HostingUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnitName

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) SetHostingUnitName(v string)`

SetHostingUnitName sets HostingUnitName field to given value.

### HasHostingUnitName

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) HasHostingUnitName() bool`

HasHostingUnitName returns a boolean if a field has been set.

### GetHostingUnitUid

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetHostingUnitUid() string`

GetHostingUnitUid returns the HostingUnitUid field if non-nil, zero value otherwise.

### GetHostingUnitUidOk

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetHostingUnitUidOk() (*string, bool)`

GetHostingUnitUidOk returns a tuple with the HostingUnitUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnitUid

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) SetHostingUnitUid(v string)`

SetHostingUnitUid sets HostingUnitUid field to given value.

### HasHostingUnitUid

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) HasHostingUnitUid() bool`

HasHostingUnitUid returns a boolean if a field has been set.

### GetDescription

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUseWriteBackCache

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetUseWriteBackCache() bool`

GetUseWriteBackCache returns the UseWriteBackCache field if non-nil, zero value otherwise.

### GetUseWriteBackCacheOk

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetUseWriteBackCacheOk() (*bool, bool)`

GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWriteBackCache

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) SetUseWriteBackCache(v bool)`

SetUseWriteBackCache sets UseWriteBackCache field to given value.

### HasUseWriteBackCache

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) HasUseWriteBackCache() bool`

HasUseWriteBackCache returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ImageVersionDetailResponseModelAllOfImageDefinition) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


