# CreateApplicationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationFolder** | Pointer to **string** | The application folder in which the application should be created. | [optional] 
**ApplicationType** | Pointer to [**ApplicationType**](ApplicationType.md) |  | [optional] 
**BrowserName** | Pointer to **string** | Internal name of the application. | [optional] 
**ClientFolder** | Pointer to **string** | Specifies the folder that the application belongs to as the user sees it. | [optional] 
**CpuPriorityLevel** | Pointer to [**CpuPriorityLevel**](CpuPriorityLevel.md) |  | [optional] 
**ApplicationGroups** | Pointer to **[]string** | Specifies one or more application groups which the application will be published to. | [optional] 
**Description** | Pointer to **string** | Specifies the description of the application. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether or not this application can be launched. | [optional] [default to true]
**DoNotEnumerate** | Pointer to **bool** | Indicates whether or not this application is enumerable | [optional] [default to false]
**HomeZone** | Pointer to **string** | Specifies a home zone preference used when launching this application. | [optional] 
**HomeZoneMode** | Pointer to [**HomeZoneMode**](HomeZoneMode.md) |  | [optional] 
**Icon** | Pointer to **string** | Icon to use for the application. | [optional] 
**IconFromClient** | Pointer to **bool** | Specifies Whether the Icon is from client or not. Only can be set to &#x60;true&#x60; if Application Type is InstalledOnClient | [optional] [default to false]
**IncludedUserFilterEnabled** | Pointer to **bool** | Specifies whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access polic(ies) is implicitly granted access to the application. | [optional] [default to false]
**IncludedUsers** | Pointer to **[]string** | Specifies the included users filter of the application; that is, the users and groups who are explicitly granted access to the published application. | [optional] 
**InstalledAppProperties** | Pointer to [**CreateApplicationRequestModelInstalledAppProperties**](CreateApplicationRequestModelInstalledAppProperties.md) |  | [optional] 
**AppVAppProperties** | Pointer to [**CreateApplicationRequestModelAppVAppProperties**](CreateApplicationRequestModelAppVAppProperties.md) |  | [optional] 
**ContentLocation** | Pointer to **string** | Location of published content. | [optional] 
**MaxPerUserInstances** | Pointer to **int32** | Specifies the maximum allowed concurrently running instances of the application that an individual user can have. | [optional] [default to 0]
**MaxTotalInstances** | Pointer to **int32** | Specifies the maximum allowed total ofconcurrently running instances of the application within the site. | [optional] [default to 0]
**Name** | **string** | Specifies the name of the application. | 
**PublishedName** | Pointer to **string** | The name seen by end users who have access to the application. | [optional] 
**ShortcutAddedToDesktop** | Pointer to **bool** | Specifies whether or not a shortcut to the application should be placed on the user device. | [optional] [default to false]
**ShortcutAddedToStartMenu** | Pointer to **bool** | Specifies whether or not a shortcut to the application should be placed in the user&#39;s start menu on their user device. | [optional] [default to false]
**StartMenuFolder** | Pointer to **string** | Specifies the name of the start menu folder that holds the application shortcut (if any). | [optional] 
**Visible** | Pointer to **bool** | Specifies whether or not this application is visible to users. | [optional] [default to true]
**WaitForPrinterCreation** | Pointer to **bool** | Specifies whether or not the session waits for the printers to be created before allowing the user to interact with the session. | [optional] [default to false]

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

### GetInstalledAppProperties

`func (o *CreateApplicationRequestModel) GetInstalledAppProperties() CreateApplicationRequestModelInstalledAppProperties`

GetInstalledAppProperties returns the InstalledAppProperties field if non-nil, zero value otherwise.

### GetInstalledAppPropertiesOk

`func (o *CreateApplicationRequestModel) GetInstalledAppPropertiesOk() (*CreateApplicationRequestModelInstalledAppProperties, bool)`

GetInstalledAppPropertiesOk returns a tuple with the InstalledAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAppProperties

`func (o *CreateApplicationRequestModel) SetInstalledAppProperties(v CreateApplicationRequestModelInstalledAppProperties)`

SetInstalledAppProperties sets InstalledAppProperties field to given value.

### HasInstalledAppProperties

`func (o *CreateApplicationRequestModel) HasInstalledAppProperties() bool`

HasInstalledAppProperties returns a boolean if a field has been set.

### GetAppVAppProperties

`func (o *CreateApplicationRequestModel) GetAppVAppProperties() CreateApplicationRequestModelAppVAppProperties`

GetAppVAppProperties returns the AppVAppProperties field if non-nil, zero value otherwise.

### GetAppVAppPropertiesOk

`func (o *CreateApplicationRequestModel) GetAppVAppPropertiesOk() (*CreateApplicationRequestModelAppVAppProperties, bool)`

GetAppVAppPropertiesOk returns a tuple with the AppVAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVAppProperties

`func (o *CreateApplicationRequestModel) SetAppVAppProperties(v CreateApplicationRequestModelAppVAppProperties)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


