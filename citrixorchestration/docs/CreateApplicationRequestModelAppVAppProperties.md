# CreateApplicationRequestModelAppVAppProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | App-V application Id. | 
**Identifier** | Pointer to **string** | App-V application identifier. | [optional] 
**PackageId** | **string** | App-V package identifier. | 
**PackageName** | **string** | App-V package name. | 
**PackageVersion** | **string** | App-V package version. | 
**PackageVersionId** | **string** | App-V package version Id. | 
**PublishingServer** | Pointer to **string** | App-V publishing server. | [optional] 
**SequenceLocation** | **string** | App-V sequence location. | 
**ServerMachineConfigurationUid** | Pointer to **string** | App-V server machine configuration Uid. | [optional] 
**TargetInPackage** | Pointer to **bool** | Specifies whether the target is in an App-V package. | [optional] 

## Methods

### NewCreateApplicationRequestModelAppVAppProperties

`func NewCreateApplicationRequestModelAppVAppProperties(id string, packageId string, packageName string, packageVersion string, packageVersionId string, sequenceLocation string, ) *CreateApplicationRequestModelAppVAppProperties`

NewCreateApplicationRequestModelAppVAppProperties instantiates a new CreateApplicationRequestModelAppVAppProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationRequestModelAppVAppPropertiesWithDefaults

`func NewCreateApplicationRequestModelAppVAppPropertiesWithDefaults() *CreateApplicationRequestModelAppVAppProperties`

NewCreateApplicationRequestModelAppVAppPropertiesWithDefaults instantiates a new CreateApplicationRequestModelAppVAppProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateApplicationRequestModelAppVAppProperties) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateApplicationRequestModelAppVAppProperties) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateApplicationRequestModelAppVAppProperties) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *CreateApplicationRequestModelAppVAppProperties) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateApplicationRequestModelAppVAppProperties) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateApplicationRequestModelAppVAppProperties) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CreateApplicationRequestModelAppVAppProperties) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetPackageId

`func (o *CreateApplicationRequestModelAppVAppProperties) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CreateApplicationRequestModelAppVAppProperties) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CreateApplicationRequestModelAppVAppProperties) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageName

`func (o *CreateApplicationRequestModelAppVAppProperties) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *CreateApplicationRequestModelAppVAppProperties) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *CreateApplicationRequestModelAppVAppProperties) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetPackageVersion

`func (o *CreateApplicationRequestModelAppVAppProperties) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *CreateApplicationRequestModelAppVAppProperties) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *CreateApplicationRequestModelAppVAppProperties) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetPackageVersionId

`func (o *CreateApplicationRequestModelAppVAppProperties) GetPackageVersionId() string`

GetPackageVersionId returns the PackageVersionId field if non-nil, zero value otherwise.

### GetPackageVersionIdOk

`func (o *CreateApplicationRequestModelAppVAppProperties) GetPackageVersionIdOk() (*string, bool)`

GetPackageVersionIdOk returns a tuple with the PackageVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersionId

`func (o *CreateApplicationRequestModelAppVAppProperties) SetPackageVersionId(v string)`

SetPackageVersionId sets PackageVersionId field to given value.


### GetPublishingServer

`func (o *CreateApplicationRequestModelAppVAppProperties) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *CreateApplicationRequestModelAppVAppProperties) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *CreateApplicationRequestModelAppVAppProperties) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.

### HasPublishingServer

`func (o *CreateApplicationRequestModelAppVAppProperties) HasPublishingServer() bool`

HasPublishingServer returns a boolean if a field has been set.

### GetSequenceLocation

`func (o *CreateApplicationRequestModelAppVAppProperties) GetSequenceLocation() string`

GetSequenceLocation returns the SequenceLocation field if non-nil, zero value otherwise.

### GetSequenceLocationOk

`func (o *CreateApplicationRequestModelAppVAppProperties) GetSequenceLocationOk() (*string, bool)`

GetSequenceLocationOk returns a tuple with the SequenceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceLocation

`func (o *CreateApplicationRequestModelAppVAppProperties) SetSequenceLocation(v string)`

SetSequenceLocation sets SequenceLocation field to given value.


### GetServerMachineConfigurationUid

`func (o *CreateApplicationRequestModelAppVAppProperties) GetServerMachineConfigurationUid() string`

GetServerMachineConfigurationUid returns the ServerMachineConfigurationUid field if non-nil, zero value otherwise.

### GetServerMachineConfigurationUidOk

`func (o *CreateApplicationRequestModelAppVAppProperties) GetServerMachineConfigurationUidOk() (*string, bool)`

GetServerMachineConfigurationUidOk returns a tuple with the ServerMachineConfigurationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMachineConfigurationUid

`func (o *CreateApplicationRequestModelAppVAppProperties) SetServerMachineConfigurationUid(v string)`

SetServerMachineConfigurationUid sets ServerMachineConfigurationUid field to given value.

### HasServerMachineConfigurationUid

`func (o *CreateApplicationRequestModelAppVAppProperties) HasServerMachineConfigurationUid() bool`

HasServerMachineConfigurationUid returns a boolean if a field has been set.

### GetTargetInPackage

`func (o *CreateApplicationRequestModelAppVAppProperties) GetTargetInPackage() bool`

GetTargetInPackage returns the TargetInPackage field if non-nil, zero value otherwise.

### GetTargetInPackageOk

`func (o *CreateApplicationRequestModelAppVAppProperties) GetTargetInPackageOk() (*bool, bool)`

GetTargetInPackageOk returns a tuple with the TargetInPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetInPackage

`func (o *CreateApplicationRequestModelAppVAppProperties) SetTargetInPackage(v bool)`

SetTargetInPackage sets TargetInPackage field to given value.

### HasTargetInPackage

`func (o *CreateApplicationRequestModelAppVAppProperties) HasTargetInPackage() bool`

HasTargetInPackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


