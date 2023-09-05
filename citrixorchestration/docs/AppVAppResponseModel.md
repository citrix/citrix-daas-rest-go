# AppVAppResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the App-V application | [optional] 
**PackageId** | **string** | App-V package identifier. | 
**PackageName** | **string** | App-V package name. | 
**PackageVersion** | **string** | App-V package version. | 
**PackageVersionId** | **string** | App-V package version Id. | 
**PublishingServer** | Pointer to **string** | App-V publishing server. | [optional] 
**ManagementServer** | Pointer to **string** | App-V Management server. | [optional] 
**SequenceLocation** | **string** | App-V sequence location. | 
**ServerMachineConfigurationUid** | Pointer to **string** | App-V server machine configuration Uid. | [optional] 
**TargetInPackage** | **bool** | Whether the target is in an App-V package. | 
**SourceTypeName** | Pointer to **string** | Application  Source Type Name. | [optional] 

## Methods

### NewAppVAppResponseModel

`func NewAppVAppResponseModel(packageId string, packageName string, packageVersion string, packageVersionId string, sequenceLocation string, targetInPackage bool, ) *AppVAppResponseModel`

NewAppVAppResponseModel instantiates a new AppVAppResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVAppResponseModelWithDefaults

`func NewAppVAppResponseModelWithDefaults() *AppVAppResponseModel`

NewAppVAppResponseModelWithDefaults instantiates a new AppVAppResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppVAppResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppVAppResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppVAppResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppVAppResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPackageId

`func (o *AppVAppResponseModel) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *AppVAppResponseModel) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *AppVAppResponseModel) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageName

`func (o *AppVAppResponseModel) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *AppVAppResponseModel) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *AppVAppResponseModel) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetPackageVersion

`func (o *AppVAppResponseModel) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *AppVAppResponseModel) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *AppVAppResponseModel) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetPackageVersionId

`func (o *AppVAppResponseModel) GetPackageVersionId() string`

GetPackageVersionId returns the PackageVersionId field if non-nil, zero value otherwise.

### GetPackageVersionIdOk

`func (o *AppVAppResponseModel) GetPackageVersionIdOk() (*string, bool)`

GetPackageVersionIdOk returns a tuple with the PackageVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersionId

`func (o *AppVAppResponseModel) SetPackageVersionId(v string)`

SetPackageVersionId sets PackageVersionId field to given value.


### GetPublishingServer

`func (o *AppVAppResponseModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *AppVAppResponseModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *AppVAppResponseModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.

### HasPublishingServer

`func (o *AppVAppResponseModel) HasPublishingServer() bool`

HasPublishingServer returns a boolean if a field has been set.

### GetManagementServer

`func (o *AppVAppResponseModel) GetManagementServer() string`

GetManagementServer returns the ManagementServer field if non-nil, zero value otherwise.

### GetManagementServerOk

`func (o *AppVAppResponseModel) GetManagementServerOk() (*string, bool)`

GetManagementServerOk returns a tuple with the ManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServer

`func (o *AppVAppResponseModel) SetManagementServer(v string)`

SetManagementServer sets ManagementServer field to given value.

### HasManagementServer

`func (o *AppVAppResponseModel) HasManagementServer() bool`

HasManagementServer returns a boolean if a field has been set.

### GetSequenceLocation

`func (o *AppVAppResponseModel) GetSequenceLocation() string`

GetSequenceLocation returns the SequenceLocation field if non-nil, zero value otherwise.

### GetSequenceLocationOk

`func (o *AppVAppResponseModel) GetSequenceLocationOk() (*string, bool)`

GetSequenceLocationOk returns a tuple with the SequenceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceLocation

`func (o *AppVAppResponseModel) SetSequenceLocation(v string)`

SetSequenceLocation sets SequenceLocation field to given value.


### GetServerMachineConfigurationUid

`func (o *AppVAppResponseModel) GetServerMachineConfigurationUid() string`

GetServerMachineConfigurationUid returns the ServerMachineConfigurationUid field if non-nil, zero value otherwise.

### GetServerMachineConfigurationUidOk

`func (o *AppVAppResponseModel) GetServerMachineConfigurationUidOk() (*string, bool)`

GetServerMachineConfigurationUidOk returns a tuple with the ServerMachineConfigurationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMachineConfigurationUid

`func (o *AppVAppResponseModel) SetServerMachineConfigurationUid(v string)`

SetServerMachineConfigurationUid sets ServerMachineConfigurationUid field to given value.

### HasServerMachineConfigurationUid

`func (o *AppVAppResponseModel) HasServerMachineConfigurationUid() bool`

HasServerMachineConfigurationUid returns a boolean if a field has been set.

### GetTargetInPackage

`func (o *AppVAppResponseModel) GetTargetInPackage() bool`

GetTargetInPackage returns the TargetInPackage field if non-nil, zero value otherwise.

### GetTargetInPackageOk

`func (o *AppVAppResponseModel) GetTargetInPackageOk() (*bool, bool)`

GetTargetInPackageOk returns a tuple with the TargetInPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetInPackage

`func (o *AppVAppResponseModel) SetTargetInPackage(v bool)`

SetTargetInPackage sets TargetInPackage field to given value.


### GetSourceTypeName

`func (o *AppVAppResponseModel) GetSourceTypeName() string`

GetSourceTypeName returns the SourceTypeName field if non-nil, zero value otherwise.

### GetSourceTypeNameOk

`func (o *AppVAppResponseModel) GetSourceTypeNameOk() (*string, bool)`

GetSourceTypeNameOk returns a tuple with the SourceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypeName

`func (o *AppVAppResponseModel) SetSourceTypeName(v string)`

SetSourceTypeName sets SourceTypeName field to given value.

### HasSourceTypeName

`func (o *AppVAppResponseModel) HasSourceTypeName() bool`

HasSourceTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


