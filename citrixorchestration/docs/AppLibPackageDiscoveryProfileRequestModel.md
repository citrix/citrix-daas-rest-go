# AppLibPackageDiscoveryProfileRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryType** | [**AppLibPackageDiscoveryType**](AppLibPackageDiscoveryType.md) |  | 
**Name** | **string** | The name of the package discovery profile | 
**DesktopGroupUid** | **int32** | The UID of the DesktopGroup from which the broker will select a VDA to run the discovery. | 
**Path** | Pointer to **NullableString** | The root directory where the discovery will run | [optional] 
**Recurse** | Pointer to **NullableBool** | A value indicating whether or not to search the child directory tree. | [optional] [default to true]
**ManagementServer** | Pointer to **NullableString** | The url of the App-V Management server that packages will be discovered from. | [optional] 
**PublishingServer** | Pointer to **NullableString** | The url of the App-V Publishing server that packages will be discovered from. | [optional] 
**Username** | Pointer to **NullableString** | The user name of the App-V Server Administrator. | [optional] 
**Password** | Pointer to **NullableString** | The password of the App-V Server Administrator. | [optional] 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**DiscoverAppV** | Pointer to **NullableBool** | A value indicating whether or not to search should Discover App-V Packages. | [optional] [default to true]
**DiscoverMsix** | Pointer to **NullableBool** | A value indicating whether or not to search should Discover MSIX Packages. | [optional] [default to false]
**DiscoverAppAttach** | Pointer to **NullableBool** | A value indicating whether or not to search should discover AppAttach Packages. | [optional] [default to false]
**AutomateDiscovery** | Pointer to **NullableBool** | A value indicating whether or not the discovery should run automatically. | [optional] [default to false]
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

### SetPathNil

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *AppLibPackageDiscoveryProfileRequestModel) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
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

### SetRecurseNil

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetRecurseNil(b bool)`

 SetRecurseNil sets the value for Recurse to be an explicit nil

### UnsetRecurse
`func (o *AppLibPackageDiscoveryProfileRequestModel) UnsetRecurse()`

UnsetRecurse ensures that no value is present for Recurse, not even an explicit nil
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

### SetManagementServerNil

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetManagementServerNil(b bool)`

 SetManagementServerNil sets the value for ManagementServer to be an explicit nil

### UnsetManagementServer
`func (o *AppLibPackageDiscoveryProfileRequestModel) UnsetManagementServer()`

UnsetManagementServer ensures that no value is present for ManagementServer, not even an explicit nil
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

### SetPublishingServerNil

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetPublishingServerNil(b bool)`

 SetPublishingServerNil sets the value for PublishingServer to be an explicit nil

### UnsetPublishingServer
`func (o *AppLibPackageDiscoveryProfileRequestModel) UnsetPublishingServer()`

UnsetPublishingServer ensures that no value is present for PublishingServer, not even an explicit nil
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

### SetUsernameNil

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AppLibPackageDiscoveryProfileRequestModel) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
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

### SetPasswordNil

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AppLibPackageDiscoveryProfileRequestModel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
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

### SetDiscoverAppVNil

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppVNil(b bool)`

 SetDiscoverAppVNil sets the value for DiscoverAppV to be an explicit nil

### UnsetDiscoverAppV
`func (o *AppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverAppV()`

UnsetDiscoverAppV ensures that no value is present for DiscoverAppV, not even an explicit nil
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

### SetDiscoverMsixNil

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetDiscoverMsixNil(b bool)`

 SetDiscoverMsixNil sets the value for DiscoverMsix to be an explicit nil

### UnsetDiscoverMsix
`func (o *AppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverMsix()`

UnsetDiscoverMsix ensures that no value is present for DiscoverMsix, not even an explicit nil
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

### SetDiscoverAppAttachNil

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppAttachNil(b bool)`

 SetDiscoverAppAttachNil sets the value for DiscoverAppAttach to be an explicit nil

### UnsetDiscoverAppAttach
`func (o *AppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverAppAttach()`

UnsetDiscoverAppAttach ensures that no value is present for DiscoverAppAttach, not even an explicit nil
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

### SetAutomateDiscoveryNil

`func (o *AppLibPackageDiscoveryProfileRequestModel) SetAutomateDiscoveryNil(b bool)`

 SetAutomateDiscoveryNil sets the value for AutomateDiscovery to be an explicit nil

### UnsetAutomateDiscovery
`func (o *AppLibPackageDiscoveryProfileRequestModel) UnsetAutomateDiscovery()`

UnsetAutomateDiscovery ensures that no value is present for AutomateDiscovery, not even an explicit nil
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


