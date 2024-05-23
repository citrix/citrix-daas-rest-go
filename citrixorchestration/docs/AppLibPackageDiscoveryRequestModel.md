# AppLibPackageDiscoveryRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryType** | Pointer to [**AppLibPackageDiscoveryType**](AppLibPackageDiscoveryType.md) |  | [optional] 
**DiscoveryProfileUid** | Pointer to **NullableInt32** | the UID of AppLib package discovery profile | [optional] 
**DesktopGroupUid** | Pointer to **NullableInt32** | The UID of the DesktopGroup from which the broker will select a VDA to run the discover session. | [optional] 
**Path** | Pointer to **NullableString** | The root directory where the discovery session will run. | [optional] 
**Recurse** | Pointer to **NullableBool** | A value indicating whether or not to search the child directory tree. | [optional] [default to true]
**DiscoverAppV** | Pointer to **NullableBool** | A value indicating whether or not to discover App-V packages | [optional] [default to false]
**DiscoverMsiX** | Pointer to **NullableBool** | A value indicating whether or not to discover MSIX packages | [optional] [default to false]
**DiscoverAppAttach** | Pointer to **NullableBool** | A value indicating whether or not to discover AppAttach packages | [optional] [default to false]
**ManagementServer** | Pointer to **NullableString** | The url of the App-V Management Server that packages will be discovered from. | [optional] 
**PublishingServer** | Pointer to **NullableString** | A value indicating whether or not to discover AppAttach packages | [optional] 
**Username** | Pointer to **NullableString** | the username of App-V server admin | [optional] 
**Password** | Pointer to **NullableString** | the password of App-V server admin | [optional] 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 

## Methods

### NewAppLibPackageDiscoveryRequestModel

`func NewAppLibPackageDiscoveryRequestModel() *AppLibPackageDiscoveryRequestModel`

NewAppLibPackageDiscoveryRequestModel instantiates a new AppLibPackageDiscoveryRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLibPackageDiscoveryRequestModelWithDefaults

`func NewAppLibPackageDiscoveryRequestModelWithDefaults() *AppLibPackageDiscoveryRequestModel`

NewAppLibPackageDiscoveryRequestModelWithDefaults instantiates a new AppLibPackageDiscoveryRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryType

`func (o *AppLibPackageDiscoveryRequestModel) GetDiscoveryType() AppLibPackageDiscoveryType`

GetDiscoveryType returns the DiscoveryType field if non-nil, zero value otherwise.

### GetDiscoveryTypeOk

`func (o *AppLibPackageDiscoveryRequestModel) GetDiscoveryTypeOk() (*AppLibPackageDiscoveryType, bool)`

GetDiscoveryTypeOk returns a tuple with the DiscoveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryType

`func (o *AppLibPackageDiscoveryRequestModel) SetDiscoveryType(v AppLibPackageDiscoveryType)`

SetDiscoveryType sets DiscoveryType field to given value.

### HasDiscoveryType

`func (o *AppLibPackageDiscoveryRequestModel) HasDiscoveryType() bool`

HasDiscoveryType returns a boolean if a field has been set.

### GetDiscoveryProfileUid

`func (o *AppLibPackageDiscoveryRequestModel) GetDiscoveryProfileUid() int32`

GetDiscoveryProfileUid returns the DiscoveryProfileUid field if non-nil, zero value otherwise.

### GetDiscoveryProfileUidOk

`func (o *AppLibPackageDiscoveryRequestModel) GetDiscoveryProfileUidOk() (*int32, bool)`

GetDiscoveryProfileUidOk returns a tuple with the DiscoveryProfileUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryProfileUid

`func (o *AppLibPackageDiscoveryRequestModel) SetDiscoveryProfileUid(v int32)`

SetDiscoveryProfileUid sets DiscoveryProfileUid field to given value.

### HasDiscoveryProfileUid

`func (o *AppLibPackageDiscoveryRequestModel) HasDiscoveryProfileUid() bool`

HasDiscoveryProfileUid returns a boolean if a field has been set.

### SetDiscoveryProfileUidNil

`func (o *AppLibPackageDiscoveryRequestModel) SetDiscoveryProfileUidNil(b bool)`

 SetDiscoveryProfileUidNil sets the value for DiscoveryProfileUid to be an explicit nil

### UnsetDiscoveryProfileUid
`func (o *AppLibPackageDiscoveryRequestModel) UnsetDiscoveryProfileUid()`

UnsetDiscoveryProfileUid ensures that no value is present for DiscoveryProfileUid, not even an explicit nil
### GetDesktopGroupUid

`func (o *AppLibPackageDiscoveryRequestModel) GetDesktopGroupUid() int32`

GetDesktopGroupUid returns the DesktopGroupUid field if non-nil, zero value otherwise.

### GetDesktopGroupUidOk

`func (o *AppLibPackageDiscoveryRequestModel) GetDesktopGroupUidOk() (*int32, bool)`

GetDesktopGroupUidOk returns a tuple with the DesktopGroupUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopGroupUid

`func (o *AppLibPackageDiscoveryRequestModel) SetDesktopGroupUid(v int32)`

SetDesktopGroupUid sets DesktopGroupUid field to given value.

### HasDesktopGroupUid

`func (o *AppLibPackageDiscoveryRequestModel) HasDesktopGroupUid() bool`

HasDesktopGroupUid returns a boolean if a field has been set.

### SetDesktopGroupUidNil

`func (o *AppLibPackageDiscoveryRequestModel) SetDesktopGroupUidNil(b bool)`

 SetDesktopGroupUidNil sets the value for DesktopGroupUid to be an explicit nil

### UnsetDesktopGroupUid
`func (o *AppLibPackageDiscoveryRequestModel) UnsetDesktopGroupUid()`

UnsetDesktopGroupUid ensures that no value is present for DesktopGroupUid, not even an explicit nil
### GetPath

`func (o *AppLibPackageDiscoveryRequestModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppLibPackageDiscoveryRequestModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppLibPackageDiscoveryRequestModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AppLibPackageDiscoveryRequestModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *AppLibPackageDiscoveryRequestModel) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *AppLibPackageDiscoveryRequestModel) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRecurse

`func (o *AppLibPackageDiscoveryRequestModel) GetRecurse() bool`

GetRecurse returns the Recurse field if non-nil, zero value otherwise.

### GetRecurseOk

`func (o *AppLibPackageDiscoveryRequestModel) GetRecurseOk() (*bool, bool)`

GetRecurseOk returns a tuple with the Recurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurse

`func (o *AppLibPackageDiscoveryRequestModel) SetRecurse(v bool)`

SetRecurse sets Recurse field to given value.

### HasRecurse

`func (o *AppLibPackageDiscoveryRequestModel) HasRecurse() bool`

HasRecurse returns a boolean if a field has been set.

### SetRecurseNil

`func (o *AppLibPackageDiscoveryRequestModel) SetRecurseNil(b bool)`

 SetRecurseNil sets the value for Recurse to be an explicit nil

### UnsetRecurse
`func (o *AppLibPackageDiscoveryRequestModel) UnsetRecurse()`

UnsetRecurse ensures that no value is present for Recurse, not even an explicit nil
### GetDiscoverAppV

`func (o *AppLibPackageDiscoveryRequestModel) GetDiscoverAppV() bool`

GetDiscoverAppV returns the DiscoverAppV field if non-nil, zero value otherwise.

### GetDiscoverAppVOk

`func (o *AppLibPackageDiscoveryRequestModel) GetDiscoverAppVOk() (*bool, bool)`

GetDiscoverAppVOk returns a tuple with the DiscoverAppV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAppV

`func (o *AppLibPackageDiscoveryRequestModel) SetDiscoverAppV(v bool)`

SetDiscoverAppV sets DiscoverAppV field to given value.

### HasDiscoverAppV

`func (o *AppLibPackageDiscoveryRequestModel) HasDiscoverAppV() bool`

HasDiscoverAppV returns a boolean if a field has been set.

### SetDiscoverAppVNil

`func (o *AppLibPackageDiscoveryRequestModel) SetDiscoverAppVNil(b bool)`

 SetDiscoverAppVNil sets the value for DiscoverAppV to be an explicit nil

### UnsetDiscoverAppV
`func (o *AppLibPackageDiscoveryRequestModel) UnsetDiscoverAppV()`

UnsetDiscoverAppV ensures that no value is present for DiscoverAppV, not even an explicit nil
### GetDiscoverMsiX

`func (o *AppLibPackageDiscoveryRequestModel) GetDiscoverMsiX() bool`

GetDiscoverMsiX returns the DiscoverMsiX field if non-nil, zero value otherwise.

### GetDiscoverMsiXOk

`func (o *AppLibPackageDiscoveryRequestModel) GetDiscoverMsiXOk() (*bool, bool)`

GetDiscoverMsiXOk returns a tuple with the DiscoverMsiX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverMsiX

`func (o *AppLibPackageDiscoveryRequestModel) SetDiscoverMsiX(v bool)`

SetDiscoverMsiX sets DiscoverMsiX field to given value.

### HasDiscoverMsiX

`func (o *AppLibPackageDiscoveryRequestModel) HasDiscoverMsiX() bool`

HasDiscoverMsiX returns a boolean if a field has been set.

### SetDiscoverMsiXNil

`func (o *AppLibPackageDiscoveryRequestModel) SetDiscoverMsiXNil(b bool)`

 SetDiscoverMsiXNil sets the value for DiscoverMsiX to be an explicit nil

### UnsetDiscoverMsiX
`func (o *AppLibPackageDiscoveryRequestModel) UnsetDiscoverMsiX()`

UnsetDiscoverMsiX ensures that no value is present for DiscoverMsiX, not even an explicit nil
### GetDiscoverAppAttach

`func (o *AppLibPackageDiscoveryRequestModel) GetDiscoverAppAttach() bool`

GetDiscoverAppAttach returns the DiscoverAppAttach field if non-nil, zero value otherwise.

### GetDiscoverAppAttachOk

`func (o *AppLibPackageDiscoveryRequestModel) GetDiscoverAppAttachOk() (*bool, bool)`

GetDiscoverAppAttachOk returns a tuple with the DiscoverAppAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverAppAttach

`func (o *AppLibPackageDiscoveryRequestModel) SetDiscoverAppAttach(v bool)`

SetDiscoverAppAttach sets DiscoverAppAttach field to given value.

### HasDiscoverAppAttach

`func (o *AppLibPackageDiscoveryRequestModel) HasDiscoverAppAttach() bool`

HasDiscoverAppAttach returns a boolean if a field has been set.

### SetDiscoverAppAttachNil

`func (o *AppLibPackageDiscoveryRequestModel) SetDiscoverAppAttachNil(b bool)`

 SetDiscoverAppAttachNil sets the value for DiscoverAppAttach to be an explicit nil

### UnsetDiscoverAppAttach
`func (o *AppLibPackageDiscoveryRequestModel) UnsetDiscoverAppAttach()`

UnsetDiscoverAppAttach ensures that no value is present for DiscoverAppAttach, not even an explicit nil
### GetManagementServer

`func (o *AppLibPackageDiscoveryRequestModel) GetManagementServer() string`

GetManagementServer returns the ManagementServer field if non-nil, zero value otherwise.

### GetManagementServerOk

`func (o *AppLibPackageDiscoveryRequestModel) GetManagementServerOk() (*string, bool)`

GetManagementServerOk returns a tuple with the ManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServer

`func (o *AppLibPackageDiscoveryRequestModel) SetManagementServer(v string)`

SetManagementServer sets ManagementServer field to given value.

### HasManagementServer

`func (o *AppLibPackageDiscoveryRequestModel) HasManagementServer() bool`

HasManagementServer returns a boolean if a field has been set.

### SetManagementServerNil

`func (o *AppLibPackageDiscoveryRequestModel) SetManagementServerNil(b bool)`

 SetManagementServerNil sets the value for ManagementServer to be an explicit nil

### UnsetManagementServer
`func (o *AppLibPackageDiscoveryRequestModel) UnsetManagementServer()`

UnsetManagementServer ensures that no value is present for ManagementServer, not even an explicit nil
### GetPublishingServer

`func (o *AppLibPackageDiscoveryRequestModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *AppLibPackageDiscoveryRequestModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *AppLibPackageDiscoveryRequestModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.

### HasPublishingServer

`func (o *AppLibPackageDiscoveryRequestModel) HasPublishingServer() bool`

HasPublishingServer returns a boolean if a field has been set.

### SetPublishingServerNil

`func (o *AppLibPackageDiscoveryRequestModel) SetPublishingServerNil(b bool)`

 SetPublishingServerNil sets the value for PublishingServer to be an explicit nil

### UnsetPublishingServer
`func (o *AppLibPackageDiscoveryRequestModel) UnsetPublishingServer()`

UnsetPublishingServer ensures that no value is present for PublishingServer, not even an explicit nil
### GetUsername

`func (o *AppLibPackageDiscoveryRequestModel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AppLibPackageDiscoveryRequestModel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AppLibPackageDiscoveryRequestModel) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AppLibPackageDiscoveryRequestModel) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *AppLibPackageDiscoveryRequestModel) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AppLibPackageDiscoveryRequestModel) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *AppLibPackageDiscoveryRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AppLibPackageDiscoveryRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AppLibPackageDiscoveryRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AppLibPackageDiscoveryRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AppLibPackageDiscoveryRequestModel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AppLibPackageDiscoveryRequestModel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordFormat

`func (o *AppLibPackageDiscoveryRequestModel) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *AppLibPackageDiscoveryRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *AppLibPackageDiscoveryRequestModel) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *AppLibPackageDiscoveryRequestModel) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


