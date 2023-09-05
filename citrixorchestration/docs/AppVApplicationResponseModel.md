# AppVApplicationResponseModel

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

## Methods

### NewAppVApplicationResponseModel

`func NewAppVApplicationResponseModel(id string, name string, packageId string, packageName string, packageVersion string, packageVersionId string, publishingServer string, ) *AppVApplicationResponseModel`

NewAppVApplicationResponseModel instantiates a new AppVApplicationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVApplicationResponseModelWithDefaults

`func NewAppVApplicationResponseModelWithDefaults() *AppVApplicationResponseModel`

NewAppVApplicationResponseModelWithDefaults instantiates a new AppVApplicationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppVApplicationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppVApplicationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppVApplicationResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetTarget

`func (o *AppVApplicationResponseModel) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AppVApplicationResponseModel) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AppVApplicationResponseModel) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AppVApplicationResponseModel) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetIdentifier

`func (o *AppVApplicationResponseModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AppVApplicationResponseModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AppVApplicationResponseModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AppVApplicationResponseModel) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetServerMachineConfigurationUid

`func (o *AppVApplicationResponseModel) GetServerMachineConfigurationUid() string`

GetServerMachineConfigurationUid returns the ServerMachineConfigurationUid field if non-nil, zero value otherwise.

### GetServerMachineConfigurationUidOk

`func (o *AppVApplicationResponseModel) GetServerMachineConfigurationUidOk() (*string, bool)`

GetServerMachineConfigurationUidOk returns a tuple with the ServerMachineConfigurationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMachineConfigurationUid

`func (o *AppVApplicationResponseModel) SetServerMachineConfigurationUid(v string)`

SetServerMachineConfigurationUid sets ServerMachineConfigurationUid field to given value.

### HasServerMachineConfigurationUid

`func (o *AppVApplicationResponseModel) HasServerMachineConfigurationUid() bool`

HasServerMachineConfigurationUid returns a boolean if a field has been set.

### GetName

`func (o *AppVApplicationResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppVApplicationResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppVApplicationResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetPackageId

`func (o *AppVApplicationResponseModel) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *AppVApplicationResponseModel) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *AppVApplicationResponseModel) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageName

`func (o *AppVApplicationResponseModel) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *AppVApplicationResponseModel) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *AppVApplicationResponseModel) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetPackageVersion

`func (o *AppVApplicationResponseModel) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *AppVApplicationResponseModel) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *AppVApplicationResponseModel) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetPackageVersionId

`func (o *AppVApplicationResponseModel) GetPackageVersionId() string`

GetPackageVersionId returns the PackageVersionId field if non-nil, zero value otherwise.

### GetPackageVersionIdOk

`func (o *AppVApplicationResponseModel) GetPackageVersionIdOk() (*string, bool)`

GetPackageVersionIdOk returns a tuple with the PackageVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersionId

`func (o *AppVApplicationResponseModel) SetPackageVersionId(v string)`

SetPackageVersionId sets PackageVersionId field to given value.


### GetPublishingServer

`func (o *AppVApplicationResponseModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *AppVApplicationResponseModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *AppVApplicationResponseModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.


### GetUid

`func (o *AppVApplicationResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AppVApplicationResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AppVApplicationResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AppVApplicationResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetBrokerApplicationNames

`func (o *AppVApplicationResponseModel) GetBrokerApplicationNames() []string`

GetBrokerApplicationNames returns the BrokerApplicationNames field if non-nil, zero value otherwise.

### GetBrokerApplicationNamesOk

`func (o *AppVApplicationResponseModel) GetBrokerApplicationNamesOk() (*[]string, bool)`

GetBrokerApplicationNamesOk returns a tuple with the BrokerApplicationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerApplicationNames

`func (o *AppVApplicationResponseModel) SetBrokerApplicationNames(v []string)`

SetBrokerApplicationNames sets BrokerApplicationNames field to given value.

### HasBrokerApplicationNames

`func (o *AppVApplicationResponseModel) HasBrokerApplicationNames() bool`

HasBrokerApplicationNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


