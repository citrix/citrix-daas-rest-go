# AppLibPackageDiscoveryProfileRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryType** | [**AppLibPackageDiscoveryType**](AppLibPackageDiscoveryType.md) |  | 
**Name** | **string** | The name of the package discovery profile | 
**DesktopGroupUid** | **int32** | The UID of the DesktopGroup from which the broker will select a VDA to run the discovery. | 
**Path** | Pointer to **string** | The root directory where the discovery will run | [optional] 
**Recurse** | Pointer to **bool** | A value indicating whether or not to search the child directory tree. | [optional] [default to true]
**ManagementServer** | Pointer to **string** | The url of the App-V Management server that packages will be discovered from. | [optional] 
**PublishingServer** | Pointer to **string** | The url of the App-V Publishing server that packages will be discovered from. | [optional] 
**Username** | Pointer to **string** | The user name of the App-V Server Administrator. | [optional] 
**Password** | Pointer to **string** | The password of the App-V Server Administrator. | [optional] 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**DiscoverAppV** | Pointer to **bool** | A value indicating whether or not to search should Discover App-V Packages. | [optional] [default to true]
**DiscoverMsix** | Pointer to **bool** | A value indicating whether or not to search should Discover MSIX Packages. | [optional] [default to false]
**DiscoverAppAttach** | Pointer to **bool** | A value indicating whether or not to search should discover AppAttach Packages. | [optional] [default to false]
**AutomateDiscovery** | Pointer to **bool** | A value indicating whether or not the discovery should run automatically. | [optional] [default to false]
**AutoDiscoveryPeriod** | Pointer to [**AutoDiscoveryPeriod**](AutoDiscoveryPeriod.md) |  | [optional] 
**AutoDiscoveryCadence** | Pointer to **int32** | A value indicating how often to trigger the discovery. | [optional] 

## Methods

### NewAppLibPackageDiscoveryProfileRequestModel

`func NewAppLibPackageDiscoveryProfileRequestModel(discoveryType AppLibPackageDiscoveryType, name string, desktopGroupUid int32, ) *AppLibPackageDiscoveryProfileRequestModel`

NewAppLibPackageDiscoveryProfileRequestModel instantiates a new AppLibPackageDiscoveryProfileRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLibPackageDiscoveryProfileRequestModelWithDefaults

`func NewAppLibPackageDiscoveryProfileRequestModelWithDefaults() *AppLibPackageDiscoveryProfileRequestModel`

NewAppLibPackageDiscoveryProfileRequestModelWithDefaults instantiates a new AppLibPackageDiscoveryProfileRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryType

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetDiscoveryType() AppLibPackageDiscoveryType`

GetDiscoveryType returns the DiscoveryType field if non-nil, zero value otherwise.

### GetDiscoveryTypeOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetDiscoveryTypeOk() (*AppLibPackageDiscoveryType, bool)`

GetDiscoveryTypeOk returns a tuple with the DiscoveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryType

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetDiscoveryType(v AppLibPackageDiscoveryType)`

SetDiscoveryType sets DiscoveryType field to given value.


### GetName

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDesktopGroupUid

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetDesktopGroupUid() int32`

GetDesktopGroupUid returns the DesktopGroupUid field if non-nil, zero value otherwise.

### GetDesktopGroupUidOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetDesktopGroupUidOk() (*int32, bool)`

GetDesktopGroupUidOk returns a tuple with the DesktopGroupUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopGroupUid

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetDesktopGroupUid(v int32)`

SetDesktopGroupUid sets DesktopGroupUid field to given value.


### GetPath

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRecurse

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetRecurse() bool`

GetRecurse returns the Recurse field if non-nil, zero value otherwise.

### GetRecurseOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetRecurseOk() (*bool, bool)`

GetRecurseOk returns a tuple with the Recurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurse

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetRecurse(v bool)`

SetRecurse sets Recurse field to given value.

### HasRecurse

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasRecurse() bool`

HasRecurse returns a boolean if a field has been set.

### GetManagementServer

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetManagementServer() string`

GetManagementServer returns the ManagementServer field if non-nil, zero value otherwise.

### GetManagementServerOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetManagementServerOk() (*string, bool)`

GetManagementServerOk returns a tuple with the ManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServer

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetManagementServer(v string)`

SetManagementServer sets ManagementServer field to given value.

### HasManagementServer

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasManagementServer() bool`

HasManagementServer returns a boolean if a field has been set.

### GetPublishingServer

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.

### HasPublishingServer

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasPublishingServer() bool`

HasPublishingServer returns a boolean if a field has been set.

### GetUsername

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordFormat

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.

### GetDiscoverAppV

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppV() bool`

GetDiscoverAppV returns the DiscoverAppV field if non-nil, zero value otherwise.

### GetDiscoverAppVOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppVOk() (*bool, bool)`

GetDiscoverAppVOk returns a tuple with the DiscoverAppV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAppV

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppV(v bool)`

SetDiscoverAppV sets DiscoverAppV field to given value.

### HasDiscoverAppV

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasDiscoverAppV() bool`

HasDiscoverAppV returns a boolean if a field has been set.

### GetDiscoverMsix

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetDiscoverMsix() bool`

GetDiscoverMsix returns the DiscoverMsix field if non-nil, zero value otherwise.

### GetDiscoverMsixOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetDiscoverMsixOk() (*bool, bool)`

GetDiscoverMsixOk returns a tuple with the DiscoverMsix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverMsix

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetDiscoverMsix(v bool)`

SetDiscoverMsix sets DiscoverMsix field to given value.

### HasDiscoverMsix

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasDiscoverMsix() bool`

HasDiscoverMsix returns a boolean if a field has been set.

### GetDiscoverAppAttach

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppAttach() bool`

GetDiscoverAppAttach returns the DiscoverAppAttach field if non-nil, zero value otherwise.

### GetDiscoverAppAttachOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppAttachOk() (*bool, bool)`

GetDiscoverAppAttachOk returns a tuple with the DiscoverAppAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAppAttach

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppAttach(v bool)`

SetDiscoverAppAttach sets DiscoverAppAttach field to given value.

### HasDiscoverAppAttach

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasDiscoverAppAttach() bool`

HasDiscoverAppAttach returns a boolean if a field has been set.

### GetAutomateDiscovery

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetAutomateDiscovery() bool`

GetAutomateDiscovery returns the AutomateDiscovery field if non-nil, zero value otherwise.

### GetAutomateDiscoveryOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetAutomateDiscoveryOk() (*bool, bool)`

GetAutomateDiscoveryOk returns a tuple with the AutomateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomateDiscovery

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetAutomateDiscovery(v bool)`

SetAutomateDiscovery sets AutomateDiscovery field to given value.

### HasAutomateDiscovery

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasAutomateDiscovery() bool`

HasAutomateDiscovery returns a boolean if a field has been set.

### GetAutoDiscoveryPeriod

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryPeriod() AutoDiscoveryPeriod`

GetAutoDiscoveryPeriod returns the AutoDiscoveryPeriod field if non-nil, zero value otherwise.

### GetAutoDiscoveryPeriodOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryPeriodOk() (*AutoDiscoveryPeriod, bool)`

GetAutoDiscoveryPeriodOk returns a tuple with the AutoDiscoveryPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscoveryPeriod

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetAutoDiscoveryPeriod(v AutoDiscoveryPeriod)`

SetAutoDiscoveryPeriod sets AutoDiscoveryPeriod field to given value.

### HasAutoDiscoveryPeriod

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasAutoDiscoveryPeriod() bool`

HasAutoDiscoveryPeriod returns a boolean if a field has been set.

### GetAutoDiscoveryCadence

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryCadence() int32`

GetAutoDiscoveryCadence returns the AutoDiscoveryCadence field if non-nil, zero value otherwise.

### GetAutoDiscoveryCadenceOk

`func (o *AppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryCadenceOk() (*int32, bool)`

GetAutoDiscoveryCadenceOk returns a tuple with the AutoDiscoveryCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscoveryCadence

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetAutoDiscoveryCadence(v int32)`

SetAutoDiscoveryCadence sets AutoDiscoveryCadence field to given value.

### HasAutoDiscoveryCadence

`func (o *AppLibPackageDiscoveryProfileRequestModel) HasAutoDiscoveryCadence() bool`

HasAutoDiscoveryCadence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


