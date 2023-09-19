# ImageDefinitionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name associated with the image definition. | [optional] 
**Id** | Pointer to **string** | The Id of the image definition. | [optional] 
**OsType** | Pointer to **NullableString** | The os type of image definition. | [optional] 
**HostingUnitName** | Pointer to **NullableString** | The hosting unit name of image definition. | [optional] 
**HostingUnitUid** | Pointer to **string** | The hosting unit id of image definition. | [optional] 
**Description** | Pointer to **NullableString** | The description of the image definition. | [optional] 
**UseWriteBackCache** | Pointer to **bool** | The UseWriteBackCache of the image definition. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The custom properties of the image definition. Currently only resource group. | [optional] 

## Methods

### NewImageDefinitionResponseModel

`func NewImageDefinitionResponseModel() *ImageDefinitionResponseModel`

NewImageDefinitionResponseModel instantiates a new ImageDefinitionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDefinitionResponseModelWithDefaults

`func NewImageDefinitionResponseModelWithDefaults() *ImageDefinitionResponseModel`

NewImageDefinitionResponseModelWithDefaults instantiates a new ImageDefinitionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageDefinitionResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageDefinitionResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageDefinitionResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageDefinitionResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ImageDefinitionResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImageDefinitionResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *ImageDefinitionResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageDefinitionResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageDefinitionResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageDefinitionResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOsType

`func (o *ImageDefinitionResponseModel) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ImageDefinitionResponseModel) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ImageDefinitionResponseModel) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ImageDefinitionResponseModel) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *ImageDefinitionResponseModel) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *ImageDefinitionResponseModel) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetHostingUnitName

`func (o *ImageDefinitionResponseModel) GetHostingUnitName() string`

GetHostingUnitName returns the HostingUnitName field if non-nil, zero value otherwise.

### GetHostingUnitNameOk

`func (o *ImageDefinitionResponseModel) GetHostingUnitNameOk() (*string, bool)`

GetHostingUnitNameOk returns a tuple with the HostingUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnitName

`func (o *ImageDefinitionResponseModel) SetHostingUnitName(v string)`

SetHostingUnitName sets HostingUnitName field to given value.

### HasHostingUnitName

`func (o *ImageDefinitionResponseModel) HasHostingUnitName() bool`

HasHostingUnitName returns a boolean if a field has been set.

### SetHostingUnitNameNil

`func (o *ImageDefinitionResponseModel) SetHostingUnitNameNil(b bool)`

 SetHostingUnitNameNil sets the value for HostingUnitName to be an explicit nil

### UnsetHostingUnitName
`func (o *ImageDefinitionResponseModel) UnsetHostingUnitName()`

UnsetHostingUnitName ensures that no value is present for HostingUnitName, not even an explicit nil
### GetHostingUnitUid

`func (o *ImageDefinitionResponseModel) GetHostingUnitUid() string`

GetHostingUnitUid returns the HostingUnitUid field if non-nil, zero value otherwise.

### GetHostingUnitUidOk

`func (o *ImageDefinitionResponseModel) GetHostingUnitUidOk() (*string, bool)`

GetHostingUnitUidOk returns a tuple with the HostingUnitUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnitUid

`func (o *ImageDefinitionResponseModel) SetHostingUnitUid(v string)`

SetHostingUnitUid sets HostingUnitUid field to given value.

### HasHostingUnitUid

`func (o *ImageDefinitionResponseModel) HasHostingUnitUid() bool`

HasHostingUnitUid returns a boolean if a field has been set.

### GetDescription

`func (o *ImageDefinitionResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageDefinitionResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageDefinitionResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageDefinitionResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ImageDefinitionResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ImageDefinitionResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUseWriteBackCache

`func (o *ImageDefinitionResponseModel) GetUseWriteBackCache() bool`

GetUseWriteBackCache returns the UseWriteBackCache field if non-nil, zero value otherwise.

### GetUseWriteBackCacheOk

`func (o *ImageDefinitionResponseModel) GetUseWriteBackCacheOk() (*bool, bool)`

GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWriteBackCache

`func (o *ImageDefinitionResponseModel) SetUseWriteBackCache(v bool)`

SetUseWriteBackCache sets UseWriteBackCache field to given value.

### HasUseWriteBackCache

`func (o *ImageDefinitionResponseModel) HasUseWriteBackCache() bool`

HasUseWriteBackCache returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ImageDefinitionResponseModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ImageDefinitionResponseModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ImageDefinitionResponseModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ImageDefinitionResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *ImageDefinitionResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *ImageDefinitionResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


