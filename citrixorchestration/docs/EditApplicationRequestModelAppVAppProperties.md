# EditApplicationRequestModelAppVAppProperties

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

### NewEditApplicationRequestModelAppVAppProperties

`func NewEditApplicationRequestModelAppVAppProperties(id string, packageId string, packageName string, packageVersion string, packageVersionId string, sequenceLocation string, ) *EditApplicationRequestModelAppVAppProperties`

NewEditApplicationRequestModelAppVAppProperties instantiates a new EditApplicationRequestModelAppVAppProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditApplicationRequestModelAppVAppPropertiesWithDefaults

`func NewEditApplicationRequestModelAppVAppPropertiesWithDefaults() *EditApplicationRequestModelAppVAppProperties`

NewEditApplicationRequestModelAppVAppPropertiesWithDefaults instantiates a new EditApplicationRequestModelAppVAppProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditApplicationRequestModelAppVAppProperties) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditApplicationRequestModelAppVAppProperties) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditApplicationRequestModelAppVAppProperties) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *EditApplicationRequestModelAppVAppProperties) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EditApplicationRequestModelAppVAppProperties) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EditApplicationRequestModelAppVAppProperties) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *EditApplicationRequestModelAppVAppProperties) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetPackageId

`func (o *EditApplicationRequestModelAppVAppProperties) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *EditApplicationRequestModelAppVAppProperties) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *EditApplicationRequestModelAppVAppProperties) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageName

`func (o *EditApplicationRequestModelAppVAppProperties) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *EditApplicationRequestModelAppVAppProperties) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *EditApplicationRequestModelAppVAppProperties) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetPackageVersion

`func (o *EditApplicationRequestModelAppVAppProperties) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *EditApplicationRequestModelAppVAppProperties) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *EditApplicationRequestModelAppVAppProperties) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetPackageVersionId

`func (o *EditApplicationRequestModelAppVAppProperties) GetPackageVersionId() string`

GetPackageVersionId returns the PackageVersionId field if non-nil, zero value otherwise.

### GetPackageVersionIdOk

`func (o *EditApplicationRequestModelAppVAppProperties) GetPackageVersionIdOk() (*string, bool)`

GetPackageVersionIdOk returns a tuple with the PackageVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersionId

`func (o *EditApplicationRequestModelAppVAppProperties) SetPackageVersionId(v string)`

SetPackageVersionId sets PackageVersionId field to given value.


### GetPublishingServer

`func (o *EditApplicationRequestModelAppVAppProperties) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *EditApplicationRequestModelAppVAppProperties) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *EditApplicationRequestModelAppVAppProperties) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.

### HasPublishingServer

`func (o *EditApplicationRequestModelAppVAppProperties) HasPublishingServer() bool`

HasPublishingServer returns a boolean if a field has been set.

### GetSequenceLocation

`func (o *EditApplicationRequestModelAppVAppProperties) GetSequenceLocation() string`

GetSequenceLocation returns the SequenceLocation field if non-nil, zero value otherwise.

### GetSequenceLocationOk

`func (o *EditApplicationRequestModelAppVAppProperties) GetSequenceLocationOk() (*string, bool)`

GetSequenceLocationOk returns a tuple with the SequenceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceLocation

`func (o *EditApplicationRequestModelAppVAppProperties) SetSequenceLocation(v string)`

SetSequenceLocation sets SequenceLocation field to given value.


### GetServerMachineConfigurationUid

`func (o *EditApplicationRequestModelAppVAppProperties) GetServerMachineConfigurationUid() string`

GetServerMachineConfigurationUid returns the ServerMachineConfigurationUid field if non-nil, zero value otherwise.

### GetServerMachineConfigurationUidOk

`func (o *EditApplicationRequestModelAppVAppProperties) GetServerMachineConfigurationUidOk() (*string, bool)`

GetServerMachineConfigurationUidOk returns a tuple with the ServerMachineConfigurationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMachineConfigurationUid

`func (o *EditApplicationRequestModelAppVAppProperties) SetServerMachineConfigurationUid(v string)`

SetServerMachineConfigurationUid sets ServerMachineConfigurationUid field to given value.

### HasServerMachineConfigurationUid

`func (o *EditApplicationRequestModelAppVAppProperties) HasServerMachineConfigurationUid() bool`

HasServerMachineConfigurationUid returns a boolean if a field has been set.

### GetTargetInPackage

`func (o *EditApplicationRequestModelAppVAppProperties) GetTargetInPackage() bool`

GetTargetInPackage returns the TargetInPackage field if non-nil, zero value otherwise.

### GetTargetInPackageOk

`func (o *EditApplicationRequestModelAppVAppProperties) GetTargetInPackageOk() (*bool, bool)`

GetTargetInPackageOk returns a tuple with the TargetInPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetInPackage

`func (o *EditApplicationRequestModelAppVAppProperties) SetTargetInPackage(v bool)`

SetTargetInPackage sets TargetInPackage field to given value.

### HasTargetInPackage

`func (o *EditApplicationRequestModelAppVAppProperties) HasTargetInPackage() bool`

HasTargetInPackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


