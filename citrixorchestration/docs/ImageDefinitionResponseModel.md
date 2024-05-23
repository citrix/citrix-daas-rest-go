# ImageDefinitionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **NullableString** | the name of the image definition | [optional] 
**CreationTime** | Pointer to **NullableString** | the create date for this image definition | [optional] 
**Description** | Pointer to **NullableString** | description | [optional] 
**OsType** | [**OsType**](OsType.md) |  | 
**LatestVersion** | **int32** | the latest version for the image definition | 
**VDASessionSupport** | [**SessionSupport**](SessionSupport.md) |  | 

## Methods

### NewImageDefinitionResponseModel

`func NewImageDefinitionResponseModel(id string, osType OsType, latestVersion int32, vDASessionSupport SessionSupport, ) *ImageDefinitionResponseModel`

NewImageDefinitionResponseModel instantiates a new ImageDefinitionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDefinitionResponseModelWithDefaults

`func NewImageDefinitionResponseModelWithDefaults() *ImageDefinitionResponseModel`

NewImageDefinitionResponseModelWithDefaults instantiates a new ImageDefinitionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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
### GetCreationTime

`func (o *ImageDefinitionResponseModel) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ImageDefinitionResponseModel) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ImageDefinitionResponseModel) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ImageDefinitionResponseModel) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### SetCreationTimeNil

`func (o *ImageDefinitionResponseModel) SetCreationTimeNil(b bool)`

 SetCreationTimeNil sets the value for CreationTime to be an explicit nil

### UnsetCreationTime
`func (o *ImageDefinitionResponseModel) UnsetCreationTime()`

UnsetCreationTime ensures that no value is present for CreationTime, not even an explicit nil
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
### GetOsType

`func (o *ImageDefinitionResponseModel) GetOsType() OsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ImageDefinitionResponseModel) GetOsTypeOk() (*OsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ImageDefinitionResponseModel) SetOsType(v OsType)`

SetOsType sets OsType field to given value.


### GetLatestVersion

`func (o *ImageDefinitionResponseModel) GetLatestVersion() int32`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *ImageDefinitionResponseModel) GetLatestVersionOk() (*int32, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *ImageDefinitionResponseModel) SetLatestVersion(v int32)`

SetLatestVersion sets LatestVersion field to given value.


### GetVDASessionSupport

`func (o *ImageDefinitionResponseModel) GetVDASessionSupport() SessionSupport`

GetVDASessionSupport returns the VDASessionSupport field if non-nil, zero value otherwise.

### GetVDASessionSupportOk

`func (o *ImageDefinitionResponseModel) GetVDASessionSupportOk() (*SessionSupport, bool)`

GetVDASessionSupportOk returns a tuple with the VDASessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDASessionSupport

`func (o *ImageDefinitionResponseModel) SetVDASessionSupport(v SessionSupport)`

SetVDASessionSupport sets VDASessionSupport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


