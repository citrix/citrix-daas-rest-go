# ImageDefinitionsAndImageVersionsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The Uid of the image definition. | [optional] 
**Name** | Pointer to **NullableString** | The name of the image definition. | [optional] 
**CreationTime** | Pointer to **NullableString** | The create date for this image definition. | [optional] 
**Description** | Pointer to **NullableString** | Description. | [optional] 
**OsType** | [**OsType**](OsType.md) |  | 
**LatestVersion** | **int32** | The latest version for the image definition. | 
**VDASessionSupport** | [**SessionSupport**](SessionSupport.md) |  | 
**HypervisorConnections** | Pointer to [**[]ImageDefinitionHypervisorConnectionResponseModel**](ImageDefinitionHypervisorConnectionResponseModel.md) | The hypervisor connections on which image definition is associated. | [optional] 
**ImageVersions** | Pointer to [**[]ImageVersionResponseModel**](ImageVersionResponseModel.md) | The image versions associated with this image definition. | [optional] 

## Methods

### NewImageDefinitionsAndImageVersionsResponseModel

`func NewImageDefinitionsAndImageVersionsResponseModel(osType OsType, latestVersion int32, vDASessionSupport SessionSupport, ) *ImageDefinitionsAndImageVersionsResponseModel`

NewImageDefinitionsAndImageVersionsResponseModel instantiates a new ImageDefinitionsAndImageVersionsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDefinitionsAndImageVersionsResponseModelWithDefaults

`func NewImageDefinitionsAndImageVersionsResponseModelWithDefaults() *ImageDefinitionsAndImageVersionsResponseModel`

NewImageDefinitionsAndImageVersionsResponseModelWithDefaults instantiates a new ImageDefinitionsAndImageVersionsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageDefinitionsAndImageVersionsResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ImageDefinitionsAndImageVersionsResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageDefinitionsAndImageVersionsResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImageDefinitionsAndImageVersionsResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreationTime

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ImageDefinitionsAndImageVersionsResponseModel) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### SetCreationTimeNil

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetCreationTimeNil(b bool)`

 SetCreationTimeNil sets the value for CreationTime to be an explicit nil

### UnsetCreationTime
`func (o *ImageDefinitionsAndImageVersionsResponseModel) UnsetCreationTime()`

UnsetCreationTime ensures that no value is present for CreationTime, not even an explicit nil
### GetDescription

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageDefinitionsAndImageVersionsResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ImageDefinitionsAndImageVersionsResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOsType

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetOsType() OsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetOsTypeOk() (*OsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetOsType(v OsType)`

SetOsType sets OsType field to given value.


### GetLatestVersion

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetLatestVersion() int32`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetLatestVersionOk() (*int32, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetLatestVersion(v int32)`

SetLatestVersion sets LatestVersion field to given value.


### GetVDASessionSupport

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetVDASessionSupport() SessionSupport`

GetVDASessionSupport returns the VDASessionSupport field if non-nil, zero value otherwise.

### GetVDASessionSupportOk

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetVDASessionSupportOk() (*SessionSupport, bool)`

GetVDASessionSupportOk returns a tuple with the VDASessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDASessionSupport

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetVDASessionSupport(v SessionSupport)`

SetVDASessionSupport sets VDASessionSupport field to given value.


### GetHypervisorConnections

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetHypervisorConnections() []ImageDefinitionHypervisorConnectionResponseModel`

GetHypervisorConnections returns the HypervisorConnections field if non-nil, zero value otherwise.

### GetHypervisorConnectionsOk

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetHypervisorConnectionsOk() (*[]ImageDefinitionHypervisorConnectionResponseModel, bool)`

GetHypervisorConnectionsOk returns a tuple with the HypervisorConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnections

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetHypervisorConnections(v []ImageDefinitionHypervisorConnectionResponseModel)`

SetHypervisorConnections sets HypervisorConnections field to given value.

### HasHypervisorConnections

`func (o *ImageDefinitionsAndImageVersionsResponseModel) HasHypervisorConnections() bool`

HasHypervisorConnections returns a boolean if a field has been set.

### SetHypervisorConnectionsNil

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetHypervisorConnectionsNil(b bool)`

 SetHypervisorConnectionsNil sets the value for HypervisorConnections to be an explicit nil

### UnsetHypervisorConnections
`func (o *ImageDefinitionsAndImageVersionsResponseModel) UnsetHypervisorConnections()`

UnsetHypervisorConnections ensures that no value is present for HypervisorConnections, not even an explicit nil
### GetImageVersions

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetImageVersions() []ImageVersionResponseModel`

GetImageVersions returns the ImageVersions field if non-nil, zero value otherwise.

### GetImageVersionsOk

`func (o *ImageDefinitionsAndImageVersionsResponseModel) GetImageVersionsOk() (*[]ImageVersionResponseModel, bool)`

GetImageVersionsOk returns a tuple with the ImageVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersions

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetImageVersions(v []ImageVersionResponseModel)`

SetImageVersions sets ImageVersions field to given value.

### HasImageVersions

`func (o *ImageDefinitionsAndImageVersionsResponseModel) HasImageVersions() bool`

HasImageVersions returns a boolean if a field has been set.

### SetImageVersionsNil

`func (o *ImageDefinitionsAndImageVersionsResponseModel) SetImageVersionsNil(b bool)`

 SetImageVersionsNil sets the value for ImageVersions to be an explicit nil

### UnsetImageVersions
`func (o *ImageDefinitionsAndImageVersionsResponseModel) UnsetImageVersions()`

UnsetImageVersions ensures that no value is present for ImageVersions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


