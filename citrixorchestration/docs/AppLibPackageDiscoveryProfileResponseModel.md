# AppLibPackageDiscoveryProfileResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** | Gets or sets a unique persistent identifier for the PackageDiscoveryProfile. | [optional] 
**Name** | Pointer to **string** | Name for the PackageDiscoveryProfile. | [optional] 
**DesktopGroupUid** | Pointer to **int32** | The UID of the DesktopGroup from which the broker will select a VDA to run the discovery. | [optional] 
**Path** | Pointer to **string** | The path to the root directory where the discovery will run. | [optional] 
**Recurse** | Pointer to **bool** | A value indicating whether or not to search the child directory tree. | [optional] 
**ManagementServer** | Pointer to **string** | The url of the App-V Management server that packages will be discovered from. | [optional] 
**PublishingServer** | Pointer to **string** | The url of the App-V Pubishing server that packages will be discovered from. | [optional] 
**Username** | Pointer to **string** | The username of the App-V server administrator. | [optional] 
**DiscoverAppV** | Pointer to **bool** | A value indicating whether or not to search should Discover App-V Packages. | [optional] 
**DiscoverMsix** | Pointer to **bool** | A value indicating whether or not to search should Discover Msix Packages. | [optional] 
**DiscoverAppAttach** | Pointer to **bool** | A value indicating whether or not to search should Discover AppAttach Packages. | [optional] 
**AutomateDiscovery** | Pointer to **bool** | A value indicating whether or not the discovery should run automatically. | [optional] 
**AutoDiscoveryPeriod** | Pointer to **string** | A value indicating what time period should used to trigger the discovery. | [optional] 
**AutoDiscoveryCadence** | Pointer to **int32** | A value indicating how often to trigger the discovery. | [optional] 

## Methods

### NewAppLibPackageDiscoveryProfileResponseModel

`func NewAppLibPackageDiscoveryProfileResponseModel() *AppLibPackageDiscoveryProfileResponseModel`

NewAppLibPackageDiscoveryProfileResponseModel instantiates a new AppLibPackageDiscoveryProfileResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLibPackageDiscoveryProfileResponseModelWithDefaults

`func NewAppLibPackageDiscoveryProfileResponseModelWithDefaults() *AppLibPackageDiscoveryProfileResponseModel`

NewAppLibPackageDiscoveryProfileResponseModelWithDefaults instantiates a new AppLibPackageDiscoveryProfileResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDesktopGroupUid

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetDesktopGroupUid() int32`

GetDesktopGroupUid returns the DesktopGroupUid field if non-nil, zero value otherwise.

### GetDesktopGroupUidOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetDesktopGroupUidOk() (*int32, bool)`

GetDesktopGroupUidOk returns a tuple with the DesktopGroupUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopGroupUid

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetDesktopGroupUid(v int32)`

SetDesktopGroupUid sets DesktopGroupUid field to given value.

### HasDesktopGroupUid

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasDesktopGroupUid() bool`

HasDesktopGroupUid returns a boolean if a field has been set.

### GetPath

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRecurse

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetRecurse() bool`

GetRecurse returns the Recurse field if non-nil, zero value otherwise.

### GetRecurseOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetRecurseOk() (*bool, bool)`

GetRecurseOk returns a tuple with the Recurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurse

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetRecurse(v bool)`

SetRecurse sets Recurse field to given value.

### HasRecurse

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasRecurse() bool`

HasRecurse returns a boolean if a field has been set.

### GetManagementServer

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetManagementServer() string`

GetManagementServer returns the ManagementServer field if non-nil, zero value otherwise.

### GetManagementServerOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetManagementServerOk() (*string, bool)`

GetManagementServerOk returns a tuple with the ManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServer

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetManagementServer(v string)`

SetManagementServer sets ManagementServer field to given value.

### HasManagementServer

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasManagementServer() bool`

HasManagementServer returns a boolean if a field has been set.

### GetPublishingServer

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.

### HasPublishingServer

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasPublishingServer() bool`

HasPublishingServer returns a boolean if a field has been set.

### GetUsername

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDiscoverAppV

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetDiscoverAppV() bool`

GetDiscoverAppV returns the DiscoverAppV field if non-nil, zero value otherwise.

### GetDiscoverAppVOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetDiscoverAppVOk() (*bool, bool)`

GetDiscoverAppVOk returns a tuple with the DiscoverAppV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAppV

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetDiscoverAppV(v bool)`

SetDiscoverAppV sets DiscoverAppV field to given value.

### HasDiscoverAppV

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasDiscoverAppV() bool`

HasDiscoverAppV returns a boolean if a field has been set.

### GetDiscoverMsix

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetDiscoverMsix() bool`

GetDiscoverMsix returns the DiscoverMsix field if non-nil, zero value otherwise.

### GetDiscoverMsixOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetDiscoverMsixOk() (*bool, bool)`

GetDiscoverMsixOk returns a tuple with the DiscoverMsix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverMsix

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetDiscoverMsix(v bool)`

SetDiscoverMsix sets DiscoverMsix field to given value.

### HasDiscoverMsix

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasDiscoverMsix() bool`

HasDiscoverMsix returns a boolean if a field has been set.

### GetDiscoverAppAttach

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetDiscoverAppAttach() bool`

GetDiscoverAppAttach returns the DiscoverAppAttach field if non-nil, zero value otherwise.

### GetDiscoverAppAttachOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetDiscoverAppAttachOk() (*bool, bool)`

GetDiscoverAppAttachOk returns a tuple with the DiscoverAppAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAppAttach

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetDiscoverAppAttach(v bool)`

SetDiscoverAppAttach sets DiscoverAppAttach field to given value.

### HasDiscoverAppAttach

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasDiscoverAppAttach() bool`

HasDiscoverAppAttach returns a boolean if a field has been set.

### GetAutomateDiscovery

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetAutomateDiscovery() bool`

GetAutomateDiscovery returns the AutomateDiscovery field if non-nil, zero value otherwise.

### GetAutomateDiscoveryOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetAutomateDiscoveryOk() (*bool, bool)`

GetAutomateDiscoveryOk returns a tuple with the AutomateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomateDiscovery

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetAutomateDiscovery(v bool)`

SetAutomateDiscovery sets AutomateDiscovery field to given value.

### HasAutomateDiscovery

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasAutomateDiscovery() bool`

HasAutomateDiscovery returns a boolean if a field has been set.

### GetAutoDiscoveryPeriod

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetAutoDiscoveryPeriod() string`

GetAutoDiscoveryPeriod returns the AutoDiscoveryPeriod field if non-nil, zero value otherwise.

### GetAutoDiscoveryPeriodOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetAutoDiscoveryPeriodOk() (*string, bool)`

GetAutoDiscoveryPeriodOk returns a tuple with the AutoDiscoveryPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscoveryPeriod

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetAutoDiscoveryPeriod(v string)`

SetAutoDiscoveryPeriod sets AutoDiscoveryPeriod field to given value.

### HasAutoDiscoveryPeriod

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasAutoDiscoveryPeriod() bool`

HasAutoDiscoveryPeriod returns a boolean if a field has been set.

### GetAutoDiscoveryCadence

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetAutoDiscoveryCadence() int32`

GetAutoDiscoveryCadence returns the AutoDiscoveryCadence field if non-nil, zero value otherwise.

### GetAutoDiscoveryCadenceOk

`func (o *AppLibPackageDiscoveryProfileResponseModel) GetAutoDiscoveryCadenceOk() (*int32, bool)`

GetAutoDiscoveryCadenceOk returns a tuple with the AutoDiscoveryCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscoveryCadence

`func (o *AppLibPackageDiscoveryProfileResponseModel) SetAutoDiscoveryCadence(v int32)`

SetAutoDiscoveryCadence sets AutoDiscoveryCadence field to given value.

### HasAutoDiscoveryCadence

`func (o *AppLibPackageDiscoveryProfileResponseModel) HasAutoDiscoveryCadence() bool`

HasAutoDiscoveryCadence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


