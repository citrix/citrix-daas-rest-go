# UpdateAppVIsolationGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of IsolationGroup | [optional] 
**Description** | Pointer to **NullableString** | Description of IsolationGroup | [optional] 
**IncludedAppVPackages** | Pointer to [**[]AppVIsolationGroupPackageRequestModel**](AppVIsolationGroupPackageRequestModel.md) | Included AppV packages | [optional] 

## Methods

### NewUpdateAppVIsolationGroupRequestModel

`func NewUpdateAppVIsolationGroupRequestModel() *UpdateAppVIsolationGroupRequestModel`

NewUpdateAppVIsolationGroupRequestModel instantiates a new UpdateAppVIsolationGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppVIsolationGroupRequestModelWithDefaults

`func NewUpdateAppVIsolationGroupRequestModelWithDefaults() *UpdateAppVIsolationGroupRequestModel`

NewUpdateAppVIsolationGroupRequestModelWithDefaults instantiates a new UpdateAppVIsolationGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAppVIsolationGroupRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAppVIsolationGroupRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAppVIsolationGroupRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAppVIsolationGroupRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateAppVIsolationGroupRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateAppVIsolationGroupRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateAppVIsolationGroupRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAppVIsolationGroupRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAppVIsolationGroupRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAppVIsolationGroupRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateAppVIsolationGroupRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateAppVIsolationGroupRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIncludedAppVPackages

`func (o *UpdateAppVIsolationGroupRequestModel) GetIncludedAppVPackages() []AppVIsolationGroupPackageRequestModel`

GetIncludedAppVPackages returns the IncludedAppVPackages field if non-nil, zero value otherwise.

### GetIncludedAppVPackagesOk

`func (o *UpdateAppVIsolationGroupRequestModel) GetIncludedAppVPackagesOk() (*[]AppVIsolationGroupPackageRequestModel, bool)`

GetIncludedAppVPackagesOk returns a tuple with the IncludedAppVPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedAppVPackages

`func (o *UpdateAppVIsolationGroupRequestModel) SetIncludedAppVPackages(v []AppVIsolationGroupPackageRequestModel)`

SetIncludedAppVPackages sets IncludedAppVPackages field to given value.

### HasIncludedAppVPackages

`func (o *UpdateAppVIsolationGroupRequestModel) HasIncludedAppVPackages() bool`

HasIncludedAppVPackages returns a boolean if a field has been set.

### SetIncludedAppVPackagesNil

`func (o *UpdateAppVIsolationGroupRequestModel) SetIncludedAppVPackagesNil(b bool)`

 SetIncludedAppVPackagesNil sets the value for IncludedAppVPackages to be an explicit nil

### UnsetIncludedAppVPackages
`func (o *UpdateAppVIsolationGroupRequestModel) UnsetIncludedAppVPackages()`

UnsetIncludedAppVPackages ensures that no value is present for IncludedAppVPackages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


