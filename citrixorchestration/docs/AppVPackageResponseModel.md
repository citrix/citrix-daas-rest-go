# AppVPackageResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppVApplications** | [**[]AppVApplicationResponseModel**](AppVApplicationResponseModel.md) | List of applications in the package. | 
**Description** | Pointer to **string** | Description of the package. | [optional] 
**Exists** | **bool** | Indicates whether the package exists. | 
**LibraryUid** | Pointer to **int32** | Id of the library which the package is in. | [optional] 
**LibraryName** | Pointer to **string** | Name of the library which the package is in. | [optional] 
**LibraryDescription** | Pointer to **string** | Description of the library which the package is in. | [optional] 
**Name** | **string** | Package name. | 
**Id** | **string** | Id of package. | 
**Path** | **string** | Path to package. | 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED.  Use Id. | [optional] 
**Version** | **string** | Package version. | 
**VersionGuid** | **string** | Package version id. | 
**SourceName** | Pointer to **string** | Gets or sets SourceName | [optional] 
**SourceType** | Pointer to [**AppVSource**](AppVSource.md) |  | [optional] 
**SourceTypeName** | Pointer to **string** | Gets or sets SourceType name | [optional] 
**OrderNumber** | Pointer to **int32** | Gets or sets OrderNumber | [optional] 
**ExplicitInclusion** | Pointer to **bool** | Gets or sets a value indicating whether ExplicitInclusion | [optional] 
**NumOfBrokerApplications** | Pointer to **int32** | Numbers of broker applications use the package | [optional] 
**NumOfBrokerDeliveryGroups** | Pointer to **int32** | Numbers of broker deliveryGroups associate the package | [optional] 

## Methods

### NewAppVPackageResponseModel

`func NewAppVPackageResponseModel(appVApplications []AppVApplicationResponseModel, exists bool, name string, id string, path string, version string, versionGuid string, ) *AppVPackageResponseModel`

NewAppVPackageResponseModel instantiates a new AppVPackageResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVPackageResponseModelWithDefaults

`func NewAppVPackageResponseModelWithDefaults() *AppVPackageResponseModel`

NewAppVPackageResponseModelWithDefaults instantiates a new AppVPackageResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppVApplications

`func (o *AppVPackageResponseModel) GetAppVApplications() []AppVApplicationResponseModel`

GetAppVApplications returns the AppVApplications field if non-nil, zero value otherwise.

### GetAppVApplicationsOk

`func (o *AppVPackageResponseModel) GetAppVApplicationsOk() (*[]AppVApplicationResponseModel, bool)`

GetAppVApplicationsOk returns a tuple with the AppVApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVApplications

`func (o *AppVPackageResponseModel) SetAppVApplications(v []AppVApplicationResponseModel)`

SetAppVApplications sets AppVApplications field to given value.


### GetDescription

`func (o *AppVPackageResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppVPackageResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppVPackageResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppVPackageResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExists

`func (o *AppVPackageResponseModel) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *AppVPackageResponseModel) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *AppVPackageResponseModel) SetExists(v bool)`

SetExists sets Exists field to given value.


### GetLibraryUid

`func (o *AppVPackageResponseModel) GetLibraryUid() int32`

GetLibraryUid returns the LibraryUid field if non-nil, zero value otherwise.

### GetLibraryUidOk

`func (o *AppVPackageResponseModel) GetLibraryUidOk() (*int32, bool)`

GetLibraryUidOk returns a tuple with the LibraryUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryUid

`func (o *AppVPackageResponseModel) SetLibraryUid(v int32)`

SetLibraryUid sets LibraryUid field to given value.

### HasLibraryUid

`func (o *AppVPackageResponseModel) HasLibraryUid() bool`

HasLibraryUid returns a boolean if a field has been set.

### GetLibraryName

`func (o *AppVPackageResponseModel) GetLibraryName() string`

GetLibraryName returns the LibraryName field if non-nil, zero value otherwise.

### GetLibraryNameOk

`func (o *AppVPackageResponseModel) GetLibraryNameOk() (*string, bool)`

GetLibraryNameOk returns a tuple with the LibraryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryName

`func (o *AppVPackageResponseModel) SetLibraryName(v string)`

SetLibraryName sets LibraryName field to given value.

### HasLibraryName

`func (o *AppVPackageResponseModel) HasLibraryName() bool`

HasLibraryName returns a boolean if a field has been set.

### GetLibraryDescription

`func (o *AppVPackageResponseModel) GetLibraryDescription() string`

GetLibraryDescription returns the LibraryDescription field if non-nil, zero value otherwise.

### GetLibraryDescriptionOk

`func (o *AppVPackageResponseModel) GetLibraryDescriptionOk() (*string, bool)`

GetLibraryDescriptionOk returns a tuple with the LibraryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryDescription

`func (o *AppVPackageResponseModel) SetLibraryDescription(v string)`

SetLibraryDescription sets LibraryDescription field to given value.

### HasLibraryDescription

`func (o *AppVPackageResponseModel) HasLibraryDescription() bool`

HasLibraryDescription returns a boolean if a field has been set.

### GetName

`func (o *AppVPackageResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppVPackageResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppVPackageResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *AppVPackageResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppVPackageResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppVPackageResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetPath

`func (o *AppVPackageResponseModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppVPackageResponseModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppVPackageResponseModel) SetPath(v string)`

SetPath sets Path field to given value.


### GetUid

`func (o *AppVPackageResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AppVPackageResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AppVPackageResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AppVPackageResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetVersion

`func (o *AppVPackageResponseModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppVPackageResponseModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppVPackageResponseModel) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersionGuid

`func (o *AppVPackageResponseModel) GetVersionGuid() string`

GetVersionGuid returns the VersionGuid field if non-nil, zero value otherwise.

### GetVersionGuidOk

`func (o *AppVPackageResponseModel) GetVersionGuidOk() (*string, bool)`

GetVersionGuidOk returns a tuple with the VersionGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGuid

`func (o *AppVPackageResponseModel) SetVersionGuid(v string)`

SetVersionGuid sets VersionGuid field to given value.


### GetSourceName

`func (o *AppVPackageResponseModel) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AppVPackageResponseModel) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AppVPackageResponseModel) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AppVPackageResponseModel) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSourceType

`func (o *AppVPackageResponseModel) GetSourceType() AppVSource`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AppVPackageResponseModel) GetSourceTypeOk() (*AppVSource, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AppVPackageResponseModel) SetSourceType(v AppVSource)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *AppVPackageResponseModel) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSourceTypeName

`func (o *AppVPackageResponseModel) GetSourceTypeName() string`

GetSourceTypeName returns the SourceTypeName field if non-nil, zero value otherwise.

### GetSourceTypeNameOk

`func (o *AppVPackageResponseModel) GetSourceTypeNameOk() (*string, bool)`

GetSourceTypeNameOk returns a tuple with the SourceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypeName

`func (o *AppVPackageResponseModel) SetSourceTypeName(v string)`

SetSourceTypeName sets SourceTypeName field to given value.

### HasSourceTypeName

`func (o *AppVPackageResponseModel) HasSourceTypeName() bool`

HasSourceTypeName returns a boolean if a field has been set.

### GetOrderNumber

`func (o *AppVPackageResponseModel) GetOrderNumber() int32`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *AppVPackageResponseModel) GetOrderNumberOk() (*int32, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *AppVPackageResponseModel) SetOrderNumber(v int32)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *AppVPackageResponseModel) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetExplicitInclusion

`func (o *AppVPackageResponseModel) GetExplicitInclusion() bool`

GetExplicitInclusion returns the ExplicitInclusion field if non-nil, zero value otherwise.

### GetExplicitInclusionOk

`func (o *AppVPackageResponseModel) GetExplicitInclusionOk() (*bool, bool)`

GetExplicitInclusionOk returns a tuple with the ExplicitInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitInclusion

`func (o *AppVPackageResponseModel) SetExplicitInclusion(v bool)`

SetExplicitInclusion sets ExplicitInclusion field to given value.

### HasExplicitInclusion

`func (o *AppVPackageResponseModel) HasExplicitInclusion() bool`

HasExplicitInclusion returns a boolean if a field has been set.

### GetNumOfBrokerApplications

`func (o *AppVPackageResponseModel) GetNumOfBrokerApplications() int32`

GetNumOfBrokerApplications returns the NumOfBrokerApplications field if non-nil, zero value otherwise.

### GetNumOfBrokerApplicationsOk

`func (o *AppVPackageResponseModel) GetNumOfBrokerApplicationsOk() (*int32, bool)`

GetNumOfBrokerApplicationsOk returns a tuple with the NumOfBrokerApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfBrokerApplications

`func (o *AppVPackageResponseModel) SetNumOfBrokerApplications(v int32)`

SetNumOfBrokerApplications sets NumOfBrokerApplications field to given value.

### HasNumOfBrokerApplications

`func (o *AppVPackageResponseModel) HasNumOfBrokerApplications() bool`

HasNumOfBrokerApplications returns a boolean if a field has been set.

### GetNumOfBrokerDeliveryGroups

`func (o *AppVPackageResponseModel) GetNumOfBrokerDeliveryGroups() int32`

GetNumOfBrokerDeliveryGroups returns the NumOfBrokerDeliveryGroups field if non-nil, zero value otherwise.

### GetNumOfBrokerDeliveryGroupsOk

`func (o *AppVPackageResponseModel) GetNumOfBrokerDeliveryGroupsOk() (*int32, bool)`

GetNumOfBrokerDeliveryGroupsOk returns a tuple with the NumOfBrokerDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfBrokerDeliveryGroups

`func (o *AppVPackageResponseModel) SetNumOfBrokerDeliveryGroups(v int32)`

SetNumOfBrokerDeliveryGroups sets NumOfBrokerDeliveryGroups field to given value.

### HasNumOfBrokerDeliveryGroups

`func (o *AppVPackageResponseModel) HasNumOfBrokerDeliveryGroups() bool`

HasNumOfBrokerDeliveryGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


