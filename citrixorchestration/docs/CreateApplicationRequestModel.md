# CreateApplicationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationFolder** | Pointer to **NullableString** | The application folder in which the application should be created. | [optional] 
**ApplicationType** | Pointer to [**ApplicationType**](ApplicationType.md) |  | [optional] 
**BrowserName** | Pointer to **NullableString** | Internal name of the application. | [optional] 
**ClientFolder** | Pointer to **NullableString** | Specifies the folder that the application belongs to as the user sees it. | [optional] 
**CpuPriorityLevel** | Pointer to [**CpuPriorityLevel**](CpuPriorityLevel.md) |  | [optional] 
**ApplicationGroups** | Pointer to **[]string** | Specifies one or more application groups which the application will be published to. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the application. | [optional] 
**Enabled** | Pointer to **NullableBool** | Specifies whether or not this application can be launched. | [optional] [default to true]
**DoNotEnumerate** | Pointer to **NullableBool** | Indicates whether or not this application is enumerable | [optional] [default to false]
**HomeZone** | Pointer to **NullableString** | Specifies a home zone preference used when launching this application. | [optional] 
**HomeZoneMode** | Pointer to [**HomeZoneMode**](HomeZoneMode.md) |  | [optional] 
**Icon** | Pointer to **NullableString** | Icon to use for the application. | [optional] 
**IconFromClient** | Pointer to **NullableBool** | Specifies Whether the Icon is from client or not. Only can be set to &#x60;true&#x60; if Application Type is InstalledOnClient | [optional] [default to false]
**IncludedUserFilterEnabled** | Pointer to **NullableBool** | Specifies whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access polic(ies) is implicitly granted access to the application. | [optional] [default to false]
**IncludedUsers** | Pointer to **[]string** | Specifies the included users filter of the application; that is, the users and groups who are explicitly granted access to the published application. | [optional] 
**InstalledAppProperties** | Pointer to [**CreateInstalledAppRequestModel**](CreateInstalledAppRequestModel.md) |  | [optional] 
**AppVAppProperties** | Pointer to [**AppVAppRequestModel**](AppVAppRequestModel.md) |  | [optional] 
**ContentLocation** | Pointer to **NullableString** | Location of published content. | [optional] 
**MaxPerUserInstances** | Pointer to **NullableInt32** | Specifies the maximum allowed concurrently running instances of the application that an individual user can have. | [optional] [default to 0]
**MaxTotalInstances** | Pointer to **NullableInt32** | Specifies the maximum allowed total ofconcurrently running instances of the application within the site. | [optional] [default to 0]
**Name** | **string** | Specifies the name of the application. | 
**PublishedName** | Pointer to **NullableString** | The name seen by end users who have access to the application. | [optional] 
**ShortcutAddedToDesktop** | Pointer to **NullableBool** | Specifies whether or not a shortcut to the application should be placed on the user device. | [optional] [default to false]
**ShortcutAddedToStartMenu** | Pointer to **NullableBool** | Specifies whether or not a shortcut to the application should be placed in the user&#39;s start menu on their user device. | [optional] [default to false]
**StartMenuFolder** | Pointer to **NullableString** | Specifies the name of the start menu folder that holds the application shortcut (if any). | [optional] 
**Visible** | Pointer to **NullableBool** | Specifies whether or not this application is visible to users. | [optional] [default to true]
**WaitForPrinterCreation** | Pointer to **NullableBool** | Specifies whether or not the session waits for the printers to be created before allowing the user to interact with the session. | [optional] [default to false]

## Methods

### NewCreateApplicationRequestModel

`func NewCreateApplicationRequestModel(name string, ) *CreateApplicationRequestModel`

NewCreateApplicationRequestModel instantiates a new CreateApplicationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationRequestModelWithDefaults

`func NewCreateApplicationRequestModelWithDefaults() *CreateApplicationRequestModel`

NewCreateApplicationRequestModelWithDefaults instantiates a new CreateApplicationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFolder

`func (o *CreateApplicationRequestModel) GetApplicationFolder() string`

GetApplicationFolder returns the ApplicationFolder field if non-nil, zero value otherwise.

### GetApplicationFolderOk

`func (o *CreateApplicationRequestModel) GetApplicationFolderOk() (*string, bool)`

GetApplicationFolderOk returns a tuple with the ApplicationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFolder

`func (o *CreateApplicationRequestModel) SetApplicationFolder(v string)`

SetApplicationFolder sets ApplicationFolder field to given value.

### HasApplicationFolder

`func (o *CreateApplicationRequestModel) HasApplicationFolder() bool`

HasApplicationFolder returns a boolean if a field has been set.

### SetApplicationFolderNil

`func (o *CreateApplicationRequestModel) SetApplicationFolderNil(b bool)`

 SetApplicationFolderNil sets the value for ApplicationFolder to be an explicit nil

### UnsetApplicationFolder
`func (o *CreateApplicationRequestModel) UnsetApplicationFolder()`

UnsetApplicationFolder ensures that no value is present for ApplicationFolder, not even an explicit nil
### GetApplicationType

`func (o *CreateApplicationRequestModel) GetApplicationType() ApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *CreateApplicationRequestModel) GetApplicationTypeOk() (*ApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *CreateApplicationRequestModel) SetApplicationType(v ApplicationType)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *CreateApplicationRequestModel) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetBrowserName

`func (o *CreateApplicationRequestModel) GetBrowserName() string`

GetBrowserName returns the BrowserName field if non-nil, zero value otherwise.

### GetBrowserNameOk

`func (o *CreateApplicationRequestModel) GetBrowserNameOk() (*string, bool)`

GetBrowserNameOk returns a tuple with the BrowserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserName

`func (o *CreateApplicationRequestModel) SetBrowserName(v string)`

SetBrowserName sets BrowserName field to given value.

### HasBrowserName

`func (o *CreateApplicationRequestModel) HasBrowserName() bool`

HasBrowserName returns a boolean if a field has been set.

### SetBrowserNameNil

`func (o *CreateApplicationRequestModel) SetBrowserNameNil(b bool)`

 SetBrowserNameNil sets the value for BrowserName to be an explicit nil

### UnsetBrowserName
`func (o *CreateApplicationRequestModel) UnsetBrowserName()`

UnsetBrowserName ensures that no value is present for BrowserName, not even an explicit nil
### GetClientFolder

`func (o *CreateApplicationRequestModel) GetClientFolder() string`

GetClientFolder returns the ClientFolder field if non-nil, zero value otherwise.

### GetClientFolderOk

`func (o *CreateApplicationRequestModel) GetClientFolderOk() (*string, bool)`

GetClientFolderOk returns a tuple with the ClientFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFolder

`func (o *CreateApplicationRequestModel) SetClientFolder(v string)`

SetClientFolder sets ClientFolder field to given value.

### HasClientFolder

`func (o *CreateApplicationRequestModel) HasClientFolder() bool`

HasClientFolder returns a boolean if a field has been set.

### SetClientFolderNil

`func (o *CreateApplicationRequestModel) SetClientFolderNil(b bool)`

 SetClientFolderNil sets the value for ClientFolder to be an explicit nil

### UnsetClientFolder
`func (o *CreateApplicationRequestModel) UnsetClientFolder()`

UnsetClientFolder ensures that no value is present for ClientFolder, not even an explicit nil
### GetCpuPriorityLevel

`func (o *CreateApplicationRequestModel) GetCpuPriorityLevel() CpuPriorityLevel`

GetCpuPriorityLevel returns the CpuPriorityLevel field if non-nil, zero value otherwise.

### GetCpuPriorityLevelOk

`func (o *CreateApplicationRequestModel) GetCpuPriorityLevelOk() (*CpuPriorityLevel, bool)`

GetCpuPriorityLevelOk returns a tuple with the CpuPriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPriorityLevel

`func (o *CreateApplicationRequestModel) SetCpuPriorityLevel(v CpuPriorityLevel)`

SetCpuPriorityLevel sets CpuPriorityLevel field to given value.

### HasCpuPriorityLevel

`func (o *CreateApplicationRequestModel) HasCpuPriorityLevel() bool`

HasCpuPriorityLevel returns a boolean if a field has been set.

### GetApplicationGroups

`func (o *CreateApplicationRequestModel) GetApplicationGroups() []string`

GetApplicationGroups returns the ApplicationGroups field if non-nil, zero value otherwise.

### GetApplicationGroupsOk

`func (o *CreateApplicationRequestModel) GetApplicationGroupsOk() (*[]string, bool)`

GetApplicationGroupsOk returns a tuple with the ApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroups

`func (o *CreateApplicationRequestModel) SetApplicationGroups(v []string)`

SetApplicationGroups sets ApplicationGroups field to given value.

### HasApplicationGroups

`func (o *CreateApplicationRequestModel) HasApplicationGroups() bool`

HasApplicationGroups returns a boolean if a field has been set.

### SetApplicationGroupsNil

`func (o *CreateApplicationRequestModel) SetApplicationGroupsNil(b bool)`

 SetApplicationGroupsNil sets the value for ApplicationGroups to be an explicit nil

### UnsetApplicationGroups
`func (o *CreateApplicationRequestModel) UnsetApplicationGroups()`

UnsetApplicationGroups ensures that no value is present for ApplicationGroups, not even an explicit nil
### GetDescription

`func (o *CreateApplicationRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateApplicationRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateApplicationRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateApplicationRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateApplicationRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateApplicationRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *CreateApplicationRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateApplicationRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateApplicationRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateApplicationRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *CreateApplicationRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *CreateApplicationRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetDoNotEnumerate

`func (o *CreateApplicationRequestModel) GetDoNotEnumerate() bool`

GetDoNotEnumerate returns the DoNotEnumerate field if non-nil, zero value otherwise.

### GetDoNotEnumerateOk

`func (o *CreateApplicationRequestModel) GetDoNotEnumerateOk() (*bool, bool)`

GetDoNotEnumerateOk returns a tuple with the DoNotEnumerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotEnumerate

`func (o *CreateApplicationRequestModel) SetDoNotEnumerate(v bool)`

SetDoNotEnumerate sets DoNotEnumerate field to given value.

### HasDoNotEnumerate

`func (o *CreateApplicationRequestModel) HasDoNotEnumerate() bool`

HasDoNotEnumerate returns a boolean if a field has been set.

### SetDoNotEnumerateNil

`func (o *CreateApplicationRequestModel) SetDoNotEnumerateNil(b bool)`

 SetDoNotEnumerateNil sets the value for DoNotEnumerate to be an explicit nil

### UnsetDoNotEnumerate
`func (o *CreateApplicationRequestModel) UnsetDoNotEnumerate()`

UnsetDoNotEnumerate ensures that no value is present for DoNotEnumerate, not even an explicit nil
### GetHomeZone

`func (o *CreateApplicationRequestModel) GetHomeZone() string`

GetHomeZone returns the HomeZone field if non-nil, zero value otherwise.

### GetHomeZoneOk

`func (o *CreateApplicationRequestModel) GetHomeZoneOk() (*string, bool)`

GetHomeZoneOk returns a tuple with the HomeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeZone

`func (o *CreateApplicationRequestModel) SetHomeZone(v string)`

SetHomeZone sets HomeZone field to given value.

### HasHomeZone

`func (o *CreateApplicationRequestModel) HasHomeZone() bool`

HasHomeZone returns a boolean if a field has been set.

### SetHomeZoneNil

`func (o *CreateApplicationRequestModel) SetHomeZoneNil(b bool)`

 SetHomeZoneNil sets the value for HomeZone to be an explicit nil

### UnsetHomeZone
`func (o *CreateApplicationRequestModel) UnsetHomeZone()`

UnsetHomeZone ensures that no value is present for HomeZone, not even an explicit nil
### GetHomeZoneMode

`func (o *CreateApplicationRequestModel) GetHomeZoneMode() HomeZoneMode`

GetHomeZoneMode returns the HomeZoneMode field if non-nil, zero value otherwise.

### GetHomeZoneModeOk

`func (o *CreateApplicationRequestModel) GetHomeZoneModeOk() (*HomeZoneMode, bool)`

GetHomeZoneModeOk returns a tuple with the HomeZoneMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeZoneMode

`func (o *CreateApplicationRequestModel) SetHomeZoneMode(v HomeZoneMode)`

SetHomeZoneMode sets HomeZoneMode field to given value.

### HasHomeZoneMode

`func (o *CreateApplicationRequestModel) HasHomeZoneMode() bool`

HasHomeZoneMode returns a boolean if a field has been set.

### GetIcon

`func (o *CreateApplicationRequestModel) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateApplicationRequestModel) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateApplicationRequestModel) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateApplicationRequestModel) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *CreateApplicationRequestModel) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CreateApplicationRequestModel) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetIconFromClient

`func (o *CreateApplicationRequestModel) GetIconFromClient() bool`

GetIconFromClient returns the IconFromClient field if non-nil, zero value otherwise.

### GetIconFromClientOk

`func (o *CreateApplicationRequestModel) GetIconFromClientOk() (*bool, bool)`

GetIconFromClientOk returns a tuple with the IconFromClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFromClient

`func (o *CreateApplicationRequestModel) SetIconFromClient(v bool)`

SetIconFromClient sets IconFromClient field to given value.

### HasIconFromClient

`func (o *CreateApplicationRequestModel) HasIconFromClient() bool`

HasIconFromClient returns a boolean if a field has been set.

### SetIconFromClientNil

`func (o *CreateApplicationRequestModel) SetIconFromClientNil(b bool)`

 SetIconFromClientNil sets the value for IconFromClient to be an explicit nil

### UnsetIconFromClient
`func (o *CreateApplicationRequestModel) UnsetIconFromClient()`

UnsetIconFromClient ensures that no value is present for IconFromClient, not even an explicit nil
### GetIncludedUserFilterEnabled

`func (o *CreateApplicationRequestModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *CreateApplicationRequestModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *CreateApplicationRequestModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *CreateApplicationRequestModel) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### SetIncludedUserFilterEnabledNil

`func (o *CreateApplicationRequestModel) SetIncludedUserFilterEnabledNil(b bool)`

 SetIncludedUserFilterEnabledNil sets the value for IncludedUserFilterEnabled to be an explicit nil

### UnsetIncludedUserFilterEnabled
`func (o *CreateApplicationRequestModel) UnsetIncludedUserFilterEnabled()`

UnsetIncludedUserFilterEnabled ensures that no value is present for IncludedUserFilterEnabled, not even an explicit nil
### GetIncludedUsers

`func (o *CreateApplicationRequestModel) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *CreateApplicationRequestModel) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *CreateApplicationRequestModel) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *CreateApplicationRequestModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *CreateApplicationRequestModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *CreateApplicationRequestModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil
### GetInstalledAppProperties

`func (o *CreateApplicationRequestModel) GetInstalledAppProperties() CreateInstalledAppRequestModel`

GetInstalledAppProperties returns the InstalledAppProperties field if non-nil, zero value otherwise.

### GetInstalledAppPropertiesOk

`func (o *CreateApplicationRequestModel) GetInstalledAppPropertiesOk() (*CreateInstalledAppRequestModel, bool)`

GetInstalledAppPropertiesOk returns a tuple with the InstalledAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAppProperties

`func (o *CreateApplicationRequestModel) SetInstalledAppProperties(v CreateInstalledAppRequestModel)`

SetInstalledAppProperties sets InstalledAppProperties field to given value.

### HasInstalledAppProperties

`func (o *CreateApplicationRequestModel) HasInstalledAppProperties() bool`

HasInstalledAppProperties returns a boolean if a field has been set.

### GetAppVAppProperties

`func (o *CreateApplicationRequestModel) GetAppVAppProperties() AppVAppRequestModel`

GetAppVAppProperties returns the AppVAppProperties field if non-nil, zero value otherwise.

### GetAppVAppPropertiesOk

`func (o *CreateApplicationRequestModel) GetAppVAppPropertiesOk() (*AppVAppRequestModel, bool)`

GetAppVAppPropertiesOk returns a tuple with the AppVAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVAppProperties

`func (o *CreateApplicationRequestModel) SetAppVAppProperties(v AppVAppRequestModel)`

SetAppVAppProperties sets AppVAppProperties field to given value.

### HasAppVAppProperties

`func (o *CreateApplicationRequestModel) HasAppVAppProperties() bool`

HasAppVAppProperties returns a boolean if a field has been set.

### GetContentLocation

`func (o *CreateApplicationRequestModel) GetContentLocation() string`

GetContentLocation returns the ContentLocation field if non-nil, zero value otherwise.

### GetContentLocationOk

`func (o *CreateApplicationRequestModel) GetContentLocationOk() (*string, bool)`

GetContentLocationOk returns a tuple with the ContentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLocation

`func (o *CreateApplicationRequestModel) SetContentLocation(v string)`

SetContentLocation sets ContentLocation field to given value.

### HasContentLocation

`func (o *CreateApplicationRequestModel) HasContentLocation() bool`

HasContentLocation returns a boolean if a field has been set.

### SetContentLocationNil

`func (o *CreateApplicationRequestModel) SetContentLocationNil(b bool)`

 SetContentLocationNil sets the value for ContentLocation to be an explicit nil

### UnsetContentLocation
`func (o *CreateApplicationRequestModel) UnsetContentLocation()`

UnsetContentLocation ensures that no value is present for ContentLocation, not even an explicit nil
### GetMaxPerUserInstances

`func (o *CreateApplicationRequestModel) GetMaxPerUserInstances() int32`

GetMaxPerUserInstances returns the MaxPerUserInstances field if non-nil, zero value otherwise.

### GetMaxPerUserInstancesOk

`func (o *CreateApplicationRequestModel) GetMaxPerUserInstancesOk() (*int32, bool)`

GetMaxPerUserInstancesOk returns a tuple with the MaxPerUserInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerUserInstances

`func (o *CreateApplicationRequestModel) SetMaxPerUserInstances(v int32)`

SetMaxPerUserInstances sets MaxPerUserInstances field to given value.

### HasMaxPerUserInstances

`func (o *CreateApplicationRequestModel) HasMaxPerUserInstances() bool`

HasMaxPerUserInstances returns a boolean if a field has been set.

### SetMaxPerUserInstancesNil

`func (o *CreateApplicationRequestModel) SetMaxPerUserInstancesNil(b bool)`

 SetMaxPerUserInstancesNil sets the value for MaxPerUserInstances to be an explicit nil

### UnsetMaxPerUserInstances
`func (o *CreateApplicationRequestModel) UnsetMaxPerUserInstances()`

UnsetMaxPerUserInstances ensures that no value is present for MaxPerUserInstances, not even an explicit nil
### GetMaxTotalInstances

`func (o *CreateApplicationRequestModel) GetMaxTotalInstances() int32`

GetMaxTotalInstances returns the MaxTotalInstances field if non-nil, zero value otherwise.

### GetMaxTotalInstancesOk

`func (o *CreateApplicationRequestModel) GetMaxTotalInstancesOk() (*int32, bool)`

GetMaxTotalInstancesOk returns a tuple with the MaxTotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalInstances

`func (o *CreateApplicationRequestModel) SetMaxTotalInstances(v int32)`

SetMaxTotalInstances sets MaxTotalInstances field to given value.

### HasMaxTotalInstances

`func (o *CreateApplicationRequestModel) HasMaxTotalInstances() bool`

HasMaxTotalInstances returns a boolean if a field has been set.

### SetMaxTotalInstancesNil

`func (o *CreateApplicationRequestModel) SetMaxTotalInstancesNil(b bool)`

 SetMaxTotalInstancesNil sets the value for MaxTotalInstances to be an explicit nil

### UnsetMaxTotalInstances
`func (o *CreateApplicationRequestModel) UnsetMaxTotalInstances()`

UnsetMaxTotalInstances ensures that no value is present for MaxTotalInstances, not even an explicit nil
### GetName

`func (o *CreateApplicationRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApplicationRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApplicationRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetPublishedName

`func (o *CreateApplicationRequestModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *CreateApplicationRequestModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *CreateApplicationRequestModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *CreateApplicationRequestModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *CreateApplicationRequestModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *CreateApplicationRequestModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetShortcutAddedToDesktop

`func (o *CreateApplicationRequestModel) GetShortcutAddedToDesktop() bool`

GetShortcutAddedToDesktop returns the ShortcutAddedToDesktop field if non-nil, zero value otherwise.

### GetShortcutAddedToDesktopOk

`func (o *CreateApplicationRequestModel) GetShortcutAddedToDesktopOk() (*bool, bool)`

GetShortcutAddedToDesktopOk returns a tuple with the ShortcutAddedToDesktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutAddedToDesktop

`func (o *CreateApplicationRequestModel) SetShortcutAddedToDesktop(v bool)`

SetShortcutAddedToDesktop sets ShortcutAddedToDesktop field to given value.

### HasShortcutAddedToDesktop

`func (o *CreateApplicationRequestModel) HasShortcutAddedToDesktop() bool`

HasShortcutAddedToDesktop returns a boolean if a field has been set.

### SetShortcutAddedToDesktopNil

`func (o *CreateApplicationRequestModel) SetShortcutAddedToDesktopNil(b bool)`

 SetShortcutAddedToDesktopNil sets the value for ShortcutAddedToDesktop to be an explicit nil

### UnsetShortcutAddedToDesktop
`func (o *CreateApplicationRequestModel) UnsetShortcutAddedToDesktop()`

UnsetShortcutAddedToDesktop ensures that no value is present for ShortcutAddedToDesktop, not even an explicit nil
### GetShortcutAddedToStartMenu

`func (o *CreateApplicationRequestModel) GetShortcutAddedToStartMenu() bool`

GetShortcutAddedToStartMenu returns the ShortcutAddedToStartMenu field if non-nil, zero value otherwise.

### GetShortcutAddedToStartMenuOk

`func (o *CreateApplicationRequestModel) GetShortcutAddedToStartMenuOk() (*bool, bool)`

GetShortcutAddedToStartMenuOk returns a tuple with the ShortcutAddedToStartMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutAddedToStartMenu

`func (o *CreateApplicationRequestModel) SetShortcutAddedToStartMenu(v bool)`

SetShortcutAddedToStartMenu sets ShortcutAddedToStartMenu field to given value.

### HasShortcutAddedToStartMenu

`func (o *CreateApplicationRequestModel) HasShortcutAddedToStartMenu() bool`

HasShortcutAddedToStartMenu returns a boolean if a field has been set.

### SetShortcutAddedToStartMenuNil

`func (o *CreateApplicationRequestModel) SetShortcutAddedToStartMenuNil(b bool)`

 SetShortcutAddedToStartMenuNil sets the value for ShortcutAddedToStartMenu to be an explicit nil

### UnsetShortcutAddedToStartMenu
`func (o *CreateApplicationRequestModel) UnsetShortcutAddedToStartMenu()`

UnsetShortcutAddedToStartMenu ensures that no value is present for ShortcutAddedToStartMenu, not even an explicit nil
### GetStartMenuFolder

`func (o *CreateApplicationRequestModel) GetStartMenuFolder() string`

GetStartMenuFolder returns the StartMenuFolder field if non-nil, zero value otherwise.

### GetStartMenuFolderOk

`func (o *CreateApplicationRequestModel) GetStartMenuFolderOk() (*string, bool)`

GetStartMenuFolderOk returns a tuple with the StartMenuFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMenuFolder

`func (o *CreateApplicationRequestModel) SetStartMenuFolder(v string)`

SetStartMenuFolder sets StartMenuFolder field to given value.

### HasStartMenuFolder

`func (o *CreateApplicationRequestModel) HasStartMenuFolder() bool`

HasStartMenuFolder returns a boolean if a field has been set.

### SetStartMenuFolderNil

`func (o *CreateApplicationRequestModel) SetStartMenuFolderNil(b bool)`

 SetStartMenuFolderNil sets the value for StartMenuFolder to be an explicit nil

### UnsetStartMenuFolder
`func (o *CreateApplicationRequestModel) UnsetStartMenuFolder()`

UnsetStartMenuFolder ensures that no value is present for StartMenuFolder, not even an explicit nil
### GetVisible

`func (o *CreateApplicationRequestModel) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *CreateApplicationRequestModel) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *CreateApplicationRequestModel) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *CreateApplicationRequestModel) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### SetVisibleNil

`func (o *CreateApplicationRequestModel) SetVisibleNil(b bool)`

 SetVisibleNil sets the value for Visible to be an explicit nil

### UnsetVisible
`func (o *CreateApplicationRequestModel) UnsetVisible()`

UnsetVisible ensures that no value is present for Visible, not even an explicit nil
### GetWaitForPrinterCreation

`func (o *CreateApplicationRequestModel) GetWaitForPrinterCreation() bool`

GetWaitForPrinterCreation returns the WaitForPrinterCreation field if non-nil, zero value otherwise.

### GetWaitForPrinterCreationOk

`func (o *CreateApplicationRequestModel) GetWaitForPrinterCreationOk() (*bool, bool)`

GetWaitForPrinterCreationOk returns a tuple with the WaitForPrinterCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForPrinterCreation

`func (o *CreateApplicationRequestModel) SetWaitForPrinterCreation(v bool)`

SetWaitForPrinterCreation sets WaitForPrinterCreation field to given value.

### HasWaitForPrinterCreation

`func (o *CreateApplicationRequestModel) HasWaitForPrinterCreation() bool`

HasWaitForPrinterCreation returns a boolean if a field has been set.

### SetWaitForPrinterCreationNil

`func (o *CreateApplicationRequestModel) SetWaitForPrinterCreationNil(b bool)`

 SetWaitForPrinterCreationNil sets the value for WaitForPrinterCreation to be an explicit nil

### UnsetWaitForPrinterCreation
`func (o *CreateApplicationRequestModel) UnsetWaitForPrinterCreation()`

UnsetWaitForPrinterCreation ensures that no value is present for WaitForPrinterCreation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


