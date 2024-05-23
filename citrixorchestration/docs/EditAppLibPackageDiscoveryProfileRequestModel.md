# EditAppLibPackageDiscoveryProfileRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the package discovery profile | [optional] 
**DesktopGroupUid** | Pointer to **NullableInt32** | The UID of the DesktopGroup from which the broker will select a VDA to run the discovery. | [optional] 
**Path** | Pointer to **NullableString** | The root directory where the discovery will run | [optional] 
**Recurse** | Pointer to **NullableBool** | A value indicating whether or not to search the child directory tree. | [optional] 
**ManagementServer** | Pointer to **NullableString** | The url of the App-V Management server that packages will be discovered from. | [optional] 
**PublishingServer** | Pointer to **NullableString** | The url of the App-V Publishing server that packages will be discovered from. | [optional] 
**Username** | Pointer to **NullableString** | The user name of the App-V Server Administrator. | [optional] 
**Password** | Pointer to **NullableString** | The password of the App-V Server Administrator. | [optional] 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**DiscoverAppV** | Pointer to **NullableBool** | A value indicating whether or not to search should Discover App-V Packages. | [optional] 
**DiscoverMsix** | Pointer to **NullableBool** | A value indicating whether or not to search should Discover MSIX Packages. | [optional] 
**DiscoverAppAttach** | Pointer to **NullableBool** | A value indicating whether or not to search should discover AppAttach Packages. | [optional] 
**AutomateDiscovery** | Pointer to **NullableBool** | A value indicating whether or not the discovery should run automatically. | [optional] 
**AutoDiscoveryPeriod** | Pointer to [**AutoDiscoveryPeriod**](AutoDiscoveryPeriod.md) |  | [optional] 
**AutoDiscoveryCadence** | Pointer to **NullableInt32** | A value indicating how often to trigger the discovery. | [optional] 

## Methods

### NewEditAppLibPackageDiscoveryProfileRequestModel

`func NewEditAppLibPackageDiscoveryProfileRequestModel() *EditAppLibPackageDiscoveryProfileRequestModel`

NewEditAppLibPackageDiscoveryProfileRequestModel instantiates a new EditAppLibPackageDiscoveryProfileRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAppLibPackageDiscoveryProfileRequestModelWithDefaults

`func NewEditAppLibPackageDiscoveryProfileRequestModelWithDefaults() *EditAppLibPackageDiscoveryProfileRequestModel`

NewEditAppLibPackageDiscoveryProfileRequestModelWithDefaults instantiates a new EditAppLibPackageDiscoveryProfileRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDesktopGroupUid

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDesktopGroupUid() int32`

GetDesktopGroupUid returns the DesktopGroupUid field if non-nil, zero value otherwise.

### GetDesktopGroupUidOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDesktopGroupUidOk() (*int32, bool)`

GetDesktopGroupUidOk returns a tuple with the DesktopGroupUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopGroupUid

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDesktopGroupUid(v int32)`

SetDesktopGroupUid sets DesktopGroupUid field to given value.

### HasDesktopGroupUid

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasDesktopGroupUid() bool`

HasDesktopGroupUid returns a boolean if a field has been set.

### SetDesktopGroupUidNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDesktopGroupUidNil(b bool)`

 SetDesktopGroupUidNil sets the value for DesktopGroupUid to be an explicit nil

### UnsetDesktopGroupUid
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetDesktopGroupUid()`

UnsetDesktopGroupUid ensures that no value is present for DesktopGroupUid, not even an explicit nil
### GetPath

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRecurse

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetRecurse() bool`

GetRecurse returns the Recurse field if non-nil, zero value otherwise.

### GetRecurseOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetRecurseOk() (*bool, bool)`

GetRecurseOk returns a tuple with the Recurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurse

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetRecurse(v bool)`

SetRecurse sets Recurse field to given value.

### HasRecurse

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasRecurse() bool`

HasRecurse returns a boolean if a field has been set.

### SetRecurseNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetRecurseNil(b bool)`

 SetRecurseNil sets the value for Recurse to be an explicit nil

### UnsetRecurse
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetRecurse()`

UnsetRecurse ensures that no value is present for Recurse, not even an explicit nil
### GetManagementServer

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetManagementServer() string`

GetManagementServer returns the ManagementServer field if non-nil, zero value otherwise.

### GetManagementServerOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetManagementServerOk() (*string, bool)`

GetManagementServerOk returns a tuple with the ManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServer

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetManagementServer(v string)`

SetManagementServer sets ManagementServer field to given value.

### HasManagementServer

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasManagementServer() bool`

HasManagementServer returns a boolean if a field has been set.

### SetManagementServerNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetManagementServerNil(b bool)`

 SetManagementServerNil sets the value for ManagementServer to be an explicit nil

### UnsetManagementServer
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetManagementServer()`

UnsetManagementServer ensures that no value is present for ManagementServer, not even an explicit nil
### GetPublishingServer

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.

### HasPublishingServer

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasPublishingServer() bool`

HasPublishingServer returns a boolean if a field has been set.

### SetPublishingServerNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPublishingServerNil(b bool)`

 SetPublishingServerNil sets the value for PublishingServer to be an explicit nil

### UnsetPublishingServer
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetPublishingServer()`

UnsetPublishingServer ensures that no value is present for PublishingServer, not even an explicit nil
### GetUsername

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordFormat

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.

### GetDiscoverAppV

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppV() bool`

GetDiscoverAppV returns the DiscoverAppV field if non-nil, zero value otherwise.

### GetDiscoverAppVOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppVOk() (*bool, bool)`

GetDiscoverAppVOk returns a tuple with the DiscoverAppV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAppV

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppV(v bool)`

SetDiscoverAppV sets DiscoverAppV field to given value.

### HasDiscoverAppV

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasDiscoverAppV() bool`

HasDiscoverAppV returns a boolean if a field has been set.

### SetDiscoverAppVNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppVNil(b bool)`

 SetDiscoverAppVNil sets the value for DiscoverAppV to be an explicit nil

### UnsetDiscoverAppV
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverAppV()`

UnsetDiscoverAppV ensures that no value is present for DiscoverAppV, not even an explicit nil
### GetDiscoverMsix

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverMsix() bool`

GetDiscoverMsix returns the DiscoverMsix field if non-nil, zero value otherwise.

### GetDiscoverMsixOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverMsixOk() (*bool, bool)`

GetDiscoverMsixOk returns a tuple with the DiscoverMsix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverMsix

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverMsix(v bool)`

SetDiscoverMsix sets DiscoverMsix field to given value.

### HasDiscoverMsix

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasDiscoverMsix() bool`

HasDiscoverMsix returns a boolean if a field has been set.

### SetDiscoverMsixNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverMsixNil(b bool)`

 SetDiscoverMsixNil sets the value for DiscoverMsix to be an explicit nil

### UnsetDiscoverMsix
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverMsix()`

UnsetDiscoverMsix ensures that no value is present for DiscoverMsix, not even an explicit nil
### GetDiscoverAppAttach

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppAttach() bool`

GetDiscoverAppAttach returns the DiscoverAppAttach field if non-nil, zero value otherwise.

### GetDiscoverAppAttachOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppAttachOk() (*bool, bool)`

GetDiscoverAppAttachOk returns a tuple with the DiscoverAppAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAppAttach

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppAttach(v bool)`

SetDiscoverAppAttach sets DiscoverAppAttach field to given value.

### HasDiscoverAppAttach

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasDiscoverAppAttach() bool`

HasDiscoverAppAttach returns a boolean if a field has been set.

### SetDiscoverAppAttachNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppAttachNil(b bool)`

 SetDiscoverAppAttachNil sets the value for DiscoverAppAttach to be an explicit nil

### UnsetDiscoverAppAttach
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverAppAttach()`

UnsetDiscoverAppAttach ensures that no value is present for DiscoverAppAttach, not even an explicit nil
### GetAutomateDiscovery

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutomateDiscovery() bool`

GetAutomateDiscovery returns the AutomateDiscovery field if non-nil, zero value otherwise.

### GetAutomateDiscoveryOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutomateDiscoveryOk() (*bool, bool)`

GetAutomateDiscoveryOk returns a tuple with the AutomateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomateDiscovery

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetAutomateDiscovery(v bool)`

SetAutomateDiscovery sets AutomateDiscovery field to given value.

### HasAutomateDiscovery

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasAutomateDiscovery() bool`

HasAutomateDiscovery returns a boolean if a field has been set.

### SetAutomateDiscoveryNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetAutomateDiscoveryNil(b bool)`

 SetAutomateDiscoveryNil sets the value for AutomateDiscovery to be an explicit nil

### UnsetAutomateDiscovery
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetAutomateDiscovery()`

UnsetAutomateDiscovery ensures that no value is present for AutomateDiscovery, not even an explicit nil
### GetAutoDiscoveryPeriod

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryPeriod() AutoDiscoveryPeriod`

GetAutoDiscoveryPeriod returns the AutoDiscoveryPeriod field if non-nil, zero value otherwise.

### GetAutoDiscoveryPeriodOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryPeriodOk() (*AutoDiscoveryPeriod, bool)`

GetAutoDiscoveryPeriodOk returns a tuple with the AutoDiscoveryPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscoveryPeriod

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetAutoDiscoveryPeriod(v AutoDiscoveryPeriod)`

SetAutoDiscoveryPeriod sets AutoDiscoveryPeriod field to given value.

### HasAutoDiscoveryPeriod

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasAutoDiscoveryPeriod() bool`

HasAutoDiscoveryPeriod returns a boolean if a field has been set.

### GetAutoDiscoveryCadence

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryCadence() int32`

GetAutoDiscoveryCadence returns the AutoDiscoveryCadence field if non-nil, zero value otherwise.

### GetAutoDiscoveryCadenceOk

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryCadenceOk() (*int32, bool)`

GetAutoDiscoveryCadenceOk returns a tuple with the AutoDiscoveryCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscoveryCadence

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetAutoDiscoveryCadence(v int32)`

SetAutoDiscoveryCadence sets AutoDiscoveryCadence field to given value.

### HasAutoDiscoveryCadence

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasAutoDiscoveryCadence() bool`

HasAutoDiscoveryCadence returns a boolean if a field has been set.

### SetAutoDiscoveryCadenceNil

`func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetAutoDiscoveryCadenceNil(b bool)`

 SetAutoDiscoveryCadenceNil sets the value for AutoDiscoveryCadence to be an explicit nil

### UnsetAutoDiscoveryCadence
`func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetAutoDiscoveryCadence()`

UnsetAutoDiscoveryCadence ensures that no value is present for AutoDiscoveryCadence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


