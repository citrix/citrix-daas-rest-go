# AppVApplicationDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the App-V application. Was: AppVAppId | 
**Target** | Pointer to **string** | Target of the App-V application. | [optional] 
**Identifier** | Pointer to **string** | Identifier of the App-V application. | [optional] 
**ServerMachineConfigurationUid** | Pointer to **string** | Internal property.  If &#x60;null&#x60;, indicates that the application type is AppLibraryAppV rather than AppV. | [optional] 
**Name** | **string** | App name. | 
**PackageId** | **string** | App package id. | 
**PackageName** | **string** | App package name. | 
**PackageVersion** | **string** | App package version. | 
**PackageVersionId** | **string** | App package version id. | 
**PublishingServer** | **string** | App publishing server. | 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED.  Use Id. | [optional] 
**BrokerApplicationNames** | Pointer to **[]string** | Delivered broker application names. | [optional] 
**SequenceLocation** | **string** | App sequence location. | 
**TargetInPackage** | **bool** | Whether the app target is in a package. | 
**Users** | [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | Users associated with the App-V application on the management server. | 
**WorkingDirectory** | **string** | Working directory for the App-V application as defined on the management server. | 

## Methods

### NewAppVApplicationDetailResponseModel

`func NewAppVApplicationDetailResponseModel(id string, name string, packageId string, packageName string, packageVersion string, packageVersionId string, publishingServer string, sequenceLocation string, targetInPackage bool, users []IdentityUserResponseModel, workingDirectory string, ) *AppVApplicationDetailResponseModel`

NewAppVApplicationDetailResponseModel instantiates a new AppVApplicationDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVApplicationDetailResponseModelWithDefaults

`func NewAppVApplicationDetailResponseModelWithDefaults() *AppVApplicationDetailResponseModel`

NewAppVApplicationDetailResponseModelWithDefaults instantiates a new AppVApplicationDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppVApplicationDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppVApplicationDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppVApplicationDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetTarget

`func (o *AppVApplicationDetailResponseModel) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AppVApplicationDetailResponseModel) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AppVApplicationDetailResponseModel) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AppVApplicationDetailResponseModel) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetIdentifier

`func (o *AppVApplicationDetailResponseModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AppVApplicationDetailResponseModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AppVApplicationDetailResponseModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AppVApplicationDetailResponseModel) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetServerMachineConfigurationUid

`func (o *AppVApplicationDetailResponseModel) GetServerMachineConfigurationUid() string`

GetServerMachineConfigurationUid returns the ServerMachineConfigurationUid field if non-nil, zero value otherwise.

### GetServerMachineConfigurationUidOk

`func (o *AppVApplicationDetailResponseModel) GetServerMachineConfigurationUidOk() (*string, bool)`

GetServerMachineConfigurationUidOk returns a tuple with the ServerMachineConfigurationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMachineConfigurationUid

`func (o *AppVApplicationDetailResponseModel) SetServerMachineConfigurationUid(v string)`

SetServerMachineConfigurationUid sets ServerMachineConfigurationUid field to given value.

### HasServerMachineConfigurationUid

`func (o *AppVApplicationDetailResponseModel) HasServerMachineConfigurationUid() bool`

HasServerMachineConfigurationUid returns a boolean if a field has been set.

### GetName

`func (o *AppVApplicationDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppVApplicationDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppVApplicationDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetPackageId

`func (o *AppVApplicationDetailResponseModel) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *AppVApplicationDetailResponseModel) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *AppVApplicationDetailResponseModel) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageName

`func (o *AppVApplicationDetailResponseModel) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *AppVApplicationDetailResponseModel) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *AppVApplicationDetailResponseModel) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetPackageVersion

`func (o *AppVApplicationDetailResponseModel) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *AppVApplicationDetailResponseModel) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *AppVApplicationDetailResponseModel) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetPackageVersionId

`func (o *AppVApplicationDetailResponseModel) GetPackageVersionId() string`

GetPackageVersionId returns the PackageVersionId field if non-nil, zero value otherwise.

### GetPackageVersionIdOk

`func (o *AppVApplicationDetailResponseModel) GetPackageVersionIdOk() (*string, bool)`

GetPackageVersionIdOk returns a tuple with the PackageVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersionId

`func (o *AppVApplicationDetailResponseModel) SetPackageVersionId(v string)`

SetPackageVersionId sets PackageVersionId field to given value.


### GetPublishingServer

`func (o *AppVApplicationDetailResponseModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *AppVApplicationDetailResponseModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *AppVApplicationDetailResponseModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.


### GetUid

`func (o *AppVApplicationDetailResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AppVApplicationDetailResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AppVApplicationDetailResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AppVApplicationDetailResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetBrokerApplicationNames

`func (o *AppVApplicationDetailResponseModel) GetBrokerApplicationNames() []string`

GetBrokerApplicationNames returns the BrokerApplicationNames field if non-nil, zero value otherwise.

### GetBrokerApplicationNamesOk

`func (o *AppVApplicationDetailResponseModel) GetBrokerApplicationNamesOk() (*[]string, bool)`

GetBrokerApplicationNamesOk returns a tuple with the BrokerApplicationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerApplicationNames

`func (o *AppVApplicationDetailResponseModel) SetBrokerApplicationNames(v []string)`

SetBrokerApplicationNames sets BrokerApplicationNames field to given value.

### HasBrokerApplicationNames

`func (o *AppVApplicationDetailResponseModel) HasBrokerApplicationNames() bool`

HasBrokerApplicationNames returns a boolean if a field has been set.

### GetSequenceLocation

`func (o *AppVApplicationDetailResponseModel) GetSequenceLocation() string`

GetSequenceLocation returns the SequenceLocation field if non-nil, zero value otherwise.

### GetSequenceLocationOk

`func (o *AppVApplicationDetailResponseModel) GetSequenceLocationOk() (*string, bool)`

GetSequenceLocationOk returns a tuple with the SequenceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceLocation

`func (o *AppVApplicationDetailResponseModel) SetSequenceLocation(v string)`

SetSequenceLocation sets SequenceLocation field to given value.


### GetTargetInPackage

`func (o *AppVApplicationDetailResponseModel) GetTargetInPackage() bool`

GetTargetInPackage returns the TargetInPackage field if non-nil, zero value otherwise.

### GetTargetInPackageOk

`func (o *AppVApplicationDetailResponseModel) GetTargetInPackageOk() (*bool, bool)`

GetTargetInPackageOk returns a tuple with the TargetInPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetInPackage

`func (o *AppVApplicationDetailResponseModel) SetTargetInPackage(v bool)`

SetTargetInPackage sets TargetInPackage field to given value.


### GetUsers

`func (o *AppVApplicationDetailResponseModel) GetUsers() []IdentityUserResponseModel`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AppVApplicationDetailResponseModel) GetUsersOk() (*[]IdentityUserResponseModel, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AppVApplicationDetailResponseModel) SetUsers(v []IdentityUserResponseModel)`

SetUsers sets Users field to given value.


### GetWorkingDirectory

`func (o *AppVApplicationDetailResponseModel) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *AppVApplicationDetailResponseModel) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *AppVApplicationDetailResponseModel) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


