# CreateAppVIsolationGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of IsolationGroup | 
**Description** | Pointer to **NullableString** | Description of IsolationGroup | [optional] 
**IncludedAppVPackages** | Pointer to [**[]AppVIsolationGroupPackageRequestModel**](AppVIsolationGroupPackageRequestModel.md) | Included AppV packages | [optional] 

## Methods

### NewCreateAppVIsolationGroupRequestModel

`func NewCreateAppVIsolationGroupRequestModel(name string, ) *CreateAppVIsolationGroupRequestModel`

NewCreateAppVIsolationGroupRequestModel instantiates a new CreateAppVIsolationGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppVIsolationGroupRequestModelWithDefaults

`func NewCreateAppVIsolationGroupRequestModelWithDefaults() *CreateAppVIsolationGroupRequestModel`

NewCreateAppVIsolationGroupRequestModelWithDefaults instantiates a new CreateAppVIsolationGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAppVIsolationGroupRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAppVIsolationGroupRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAppVIsolationGroupRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAppVIsolationGroupRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAppVIsolationGroupRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAppVIsolationGroupRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAppVIsolationGroupRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateAppVIsolationGroupRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateAppVIsolationGroupRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIncludedAppVPackages

`func (o *CreateAppVIsolationGroupRequestModel) GetIncludedAppVPackages() []AppVIsolationGroupPackageRequestModel`

GetIncludedAppVPackages returns the IncludedAppVPackages field if non-nil, zero value otherwise.

### GetIncludedAppVPackagesOk

`func (o *CreateAppVIsolationGroupRequestModel) GetIncludedAppVPackagesOk() (*[]AppVIsolationGroupPackageRequestModel, bool)`

GetIncludedAppVPackagesOk returns a tuple with the IncludedAppVPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedAppVPackages

`func (o *CreateAppVIsolationGroupRequestModel) SetIncludedAppVPackages(v []AppVIsolationGroupPackageRequestModel)`

SetIncludedAppVPackages sets IncludedAppVPackages field to given value.

### HasIncludedAppVPackages

`func (o *CreateAppVIsolationGroupRequestModel) HasIncludedAppVPackages() bool`

HasIncludedAppVPackages returns a boolean if a field has been set.

### SetIncludedAppVPackagesNil

`func (o *CreateAppVIsolationGroupRequestModel) SetIncludedAppVPackagesNil(b bool)`

 SetIncludedAppVPackagesNil sets the value for IncludedAppVPackages to be an explicit nil

### UnsetIncludedAppVPackages
`func (o *CreateAppVIsolationGroupRequestModel) UnsetIncludedAppVPackages()`

UnsetIncludedAppVPackages ensures that no value is present for IncludedAppVPackages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


