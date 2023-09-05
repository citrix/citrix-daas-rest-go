# AppVIsolationGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of IsolationGroup | [optional] 
**Description** | Pointer to **string** | Description of IsolationGroup | [optional] 
**Version** | Pointer to **string** | Version of IsolationGroup | [optional] 
**Uid** | Pointer to **int32** | Uid of IsolationGroup | [optional] 
**AppVPackages** | Pointer to [**[]AppVPackageResponseModel**](AppVPackageResponseModel.md) | Contained AppV packages | [optional] 

## Methods

### NewAppVIsolationGroupResponseModel

`func NewAppVIsolationGroupResponseModel() *AppVIsolationGroupResponseModel`

NewAppVIsolationGroupResponseModel instantiates a new AppVIsolationGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVIsolationGroupResponseModelWithDefaults

`func NewAppVIsolationGroupResponseModelWithDefaults() *AppVIsolationGroupResponseModel`

NewAppVIsolationGroupResponseModelWithDefaults instantiates a new AppVIsolationGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppVIsolationGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppVIsolationGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppVIsolationGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppVIsolationGroupResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AppVIsolationGroupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppVIsolationGroupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppVIsolationGroupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppVIsolationGroupResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *AppVIsolationGroupResponseModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppVIsolationGroupResponseModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppVIsolationGroupResponseModel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppVIsolationGroupResponseModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUid

`func (o *AppVIsolationGroupResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AppVIsolationGroupResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AppVIsolationGroupResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AppVIsolationGroupResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetAppVPackages

`func (o *AppVIsolationGroupResponseModel) GetAppVPackages() []AppVPackageResponseModel`

GetAppVPackages returns the AppVPackages field if non-nil, zero value otherwise.

### GetAppVPackagesOk

`func (o *AppVIsolationGroupResponseModel) GetAppVPackagesOk() (*[]AppVPackageResponseModel, bool)`

GetAppVPackagesOk returns a tuple with the AppVPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVPackages

`func (o *AppVIsolationGroupResponseModel) SetAppVPackages(v []AppVPackageResponseModel)`

SetAppVPackages sets AppVPackages field to given value.

### HasAppVPackages

`func (o *AppVIsolationGroupResponseModel) HasAppVPackages() bool`

HasAppVPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


