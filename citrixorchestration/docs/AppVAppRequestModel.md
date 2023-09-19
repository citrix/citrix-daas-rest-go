# AppVAppRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | App-V application Id. | 
**Identifier** | Pointer to **NullableString** | App-V application identifier. | [optional] 
**PackageId** | **string** | App-V package identifier. | 
**PackageName** | **string** | App-V package name. | 
**PackageVersion** | **string** | App-V package version. | 
**PackageVersionId** | **string** | App-V package version Id. | 
**PublishingServer** | Pointer to **NullableString** | App-V publishing server. | [optional] 
**SequenceLocation** | **string** | App-V sequence location. | 
**ServerMachineConfigurationUid** | Pointer to **NullableString** | App-V server machine configuration Uid. | [optional] 
**TargetInPackage** | Pointer to **bool** | Specifies whether the target is in an App-V package. | [optional] 

## Methods

### NewAppVAppRequestModel

`func NewAppVAppRequestModel(id string, packageId string, packageName string, packageVersion string, packageVersionId string, sequenceLocation string, ) *AppVAppRequestModel`

NewAppVAppRequestModel instantiates a new AppVAppRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVAppRequestModelWithDefaults

`func NewAppVAppRequestModelWithDefaults() *AppVAppRequestModel`

NewAppVAppRequestModelWithDefaults instantiates a new AppVAppRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppVAppRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppVAppRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppVAppRequestModel) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *AppVAppRequestModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AppVAppRequestModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AppVAppRequestModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AppVAppRequestModel) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *AppVAppRequestModel) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *AppVAppRequestModel) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetPackageId

`func (o *AppVAppRequestModel) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *AppVAppRequestModel) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *AppVAppRequestModel) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageName

`func (o *AppVAppRequestModel) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *AppVAppRequestModel) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *AppVAppRequestModel) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetPackageVersion

`func (o *AppVAppRequestModel) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *AppVAppRequestModel) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *AppVAppRequestModel) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetPackageVersionId

`func (o *AppVAppRequestModel) GetPackageVersionId() string`

GetPackageVersionId returns the PackageVersionId field if non-nil, zero value otherwise.

### GetPackageVersionIdOk

`func (o *AppVAppRequestModel) GetPackageVersionIdOk() (*string, bool)`

GetPackageVersionIdOk returns a tuple with the PackageVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersionId

`func (o *AppVAppRequestModel) SetPackageVersionId(v string)`

SetPackageVersionId sets PackageVersionId field to given value.


### GetPublishingServer

`func (o *AppVAppRequestModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *AppVAppRequestModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *AppVAppRequestModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.

### HasPublishingServer

`func (o *AppVAppRequestModel) HasPublishingServer() bool`

HasPublishingServer returns a boolean if a field has been set.

### SetPublishingServerNil

`func (o *AppVAppRequestModel) SetPublishingServerNil(b bool)`

 SetPublishingServerNil sets the value for PublishingServer to be an explicit nil

### UnsetPublishingServer
`func (o *AppVAppRequestModel) UnsetPublishingServer()`

UnsetPublishingServer ensures that no value is present for PublishingServer, not even an explicit nil
### GetSequenceLocation

`func (o *AppVAppRequestModel) GetSequenceLocation() string`

GetSequenceLocation returns the SequenceLocation field if non-nil, zero value otherwise.

### GetSequenceLocationOk

`func (o *AppVAppRequestModel) GetSequenceLocationOk() (*string, bool)`

GetSequenceLocationOk returns a tuple with the SequenceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceLocation

`func (o *AppVAppRequestModel) SetSequenceLocation(v string)`

SetSequenceLocation sets SequenceLocation field to given value.


### GetServerMachineConfigurationUid

`func (o *AppVAppRequestModel) GetServerMachineConfigurationUid() string`

GetServerMachineConfigurationUid returns the ServerMachineConfigurationUid field if non-nil, zero value otherwise.

### GetServerMachineConfigurationUidOk

`func (o *AppVAppRequestModel) GetServerMachineConfigurationUidOk() (*string, bool)`

GetServerMachineConfigurationUidOk returns a tuple with the ServerMachineConfigurationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMachineConfigurationUid

`func (o *AppVAppRequestModel) SetServerMachineConfigurationUid(v string)`

SetServerMachineConfigurationUid sets ServerMachineConfigurationUid field to given value.

### HasServerMachineConfigurationUid

`func (o *AppVAppRequestModel) HasServerMachineConfigurationUid() bool`

HasServerMachineConfigurationUid returns a boolean if a field has been set.

### SetServerMachineConfigurationUidNil

`func (o *AppVAppRequestModel) SetServerMachineConfigurationUidNil(b bool)`

 SetServerMachineConfigurationUidNil sets the value for ServerMachineConfigurationUid to be an explicit nil

### UnsetServerMachineConfigurationUid
`func (o *AppVAppRequestModel) UnsetServerMachineConfigurationUid()`

UnsetServerMachineConfigurationUid ensures that no value is present for ServerMachineConfigurationUid, not even an explicit nil
### GetTargetInPackage

`func (o *AppVAppRequestModel) GetTargetInPackage() bool`

GetTargetInPackage returns the TargetInPackage field if non-nil, zero value otherwise.

### GetTargetInPackageOk

`func (o *AppVAppRequestModel) GetTargetInPackageOk() (*bool, bool)`

GetTargetInPackageOk returns a tuple with the TargetInPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetInPackage

`func (o *AppVAppRequestModel) SetTargetInPackage(v bool)`

SetTargetInPackage sets TargetInPackage field to given value.

### HasTargetInPackage

`func (o *AppVAppRequestModel) HasTargetInPackage() bool`

HasTargetInPackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


