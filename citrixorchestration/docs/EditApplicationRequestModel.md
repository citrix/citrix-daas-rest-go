# EditApplicationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationFolder** | Pointer to **NullableString** | The folder in which the application resides. If not specified, the value will not be changed. May be specified as either the folder Id or Path. If specified as a path, and the path does not exist, it will be automatically created. | [optional] 
**ApplicationGroups** | Pointer to **[]string** | Specifies one or more application groups which the application will be published to. If not specified, will not be changed. If specified, the entire list of application groups must be specified. To disassociate the application from all application groups, specify an empty array ([]). | [optional] 
**DeliveryGroups** | Pointer to [**[]PriorityRefRequestModel**](PriorityRefRequestModel.md) | Specifies one or more delivery groups which the application will be published to. If not specified, will not be changed. If specified, the entire list of delivery groups must be specified. To disassociate the application from all delivery groups, specify an empty array ([]). | [optional] 
**BrowserName** | Pointer to **NullableString** | Internal name of the application. If not specified, will not be changed. If specified, must be unique within the site.  For legacy reasons the maximum length of the browser name cannot exceed 38 bytes when encoded using MBCS encoding. | [optional] 
**Categories** | Pointer to **[]string** | Categories in which the application resides. If not specified, the values will not be changed. If specified, the entire list is required. | [optional] 
**ClientFolder** | Pointer to **NullableString** | The folder that the application belongs to as the user sees it. If not specified, the value will not be changed. To remove the client folder, specify the empty string (\&quot;\&quot;). | [optional] 
**CpuPriorityLevel** | Pointer to [**CpuPriorityLevel**](CpuPriorityLevel.md) |  | [optional] 
**Description** | Pointer to **NullableString** | The description of the application. If not specified, will not be changed. | [optional] 
**Enabled** | Pointer to **NullableBool** | Indicates whether or not this application can be launched. If not specified, will not be changed. | [optional] 
**DoNotEnumerate** | Pointer to **NullableBool** | Indicates whether or not this application is enumerable | [optional] 
**HomeZone** | Pointer to **NullableString** | Home zone preference used when launching this application. If not specified, will not be changed. To disassociate the application from a home zone, set this to the empty string (\&quot;\&quot;). Cannot be set for applications where ApplicationType is equal to InstalledOnClient or PublishedContent. Cannot be set if HomeZoneMode is equal to \&quot;HomeZoneMode.Ignore\&quot;. If HomeZoneMode is equal to \&quot;HomeZoneMode.Only\&quot; or \&quot;HomeZoneMode.Prefer\&quot;, the value must either have been previously set, or is required. May be specified as either the Id or Name of the zone. | [optional] 
**HomeZoneMode** | Pointer to [**HomeZoneMode**](HomeZoneMode.md) |  | [optional] 
**Icon** | Pointer to **NullableString** | Icon used for the application. If not specified, will not be changed. | [optional] 
**IconFromClient** | Pointer to **NullableBool** | Specifies whether the icon is gotten from user&#39;s computer at run time. If not specified, will not be changed. | [optional] 
**IncludedUserFilterEnabled** | Pointer to **NullableBool** | Specifies whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access polic(ies) is implicitly granted access to the application. If not specified, will not be changed. | [optional] 
**IncludedUsers** | Pointer to **[]string** | Specifies the included users filter of the application; that is, the users and groups who are explicitly granted access to the published application. If not specified, will not be changed. If specified, all users to be included must be specified. | [optional] 
**PackagedApplicationVisibility** | Pointer to [**PackagedApplicationVisibility**](PackagedApplicationVisibility.md) |  | [optional] 
**InstalledAppProperties** | Pointer to [**EditInstalledAppRequestModel**](EditInstalledAppRequestModel.md) |  | [optional] 
**PackagedAppProperties** | Pointer to [**AppVAppRequestModel**](AppVAppRequestModel.md) |  | [optional] 
**AppVAppProperties** | Pointer to [**AppVAppRequestModel**](AppVAppRequestModel.md) |  | [optional] 
**ContentLocation** | Pointer to **NullableString** | Location of published content. If not specified, will not be changed. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of the application. When set the property value equal to null/empty means to remove this property. | [optional] 
**MaxPerUserInstances** | Pointer to **NullableInt32** | Specifies the maximum allowed concurrently running instances of the application that an individual user can have. If not specified, will not be changed. | [optional] 
**MaxTotalInstances** | Pointer to **NullableInt32** | Specifies the maximum allowed total of concurrently running instances of the application within the site. If not specified, will not be changed. | [optional] 
**Name** | Pointer to **NullableString** | Name of the application.  Only seen by administrators. If not specified, will not be changed. If specified, must be unique within the site. | [optional] 
**PublishedName** | Pointer to **NullableString** | The name seen by end users who have access to the application. If not specified, will not be changed. | [optional] 
**ShortcutAddedToDesktop** | Pointer to **NullableBool** | Specifies whether or not a shortcut to the application should be placed on the user device. If not specified, will not be changed. | [optional] 
**ShortcutAddedToStartMenu** | Pointer to **NullableBool** | Specifies whether or not a shortcut to the application should be placed in the user&#39;s start menu on their user device. If not specified, will not be changed. | [optional] 
**StartMenuFolder** | Pointer to **NullableString** | Specifies the name of the start menu folder that holds the application shortcut (if any). If not specified, will not be changed. To clear a previously-set value, use the empty string (\&quot;\&quot;). | [optional] 
**Visible** | Pointer to **NullableBool** | Indicates whether or not this application is visible to users. If not specified, will not be changed. | [optional] 
**WaitForPrinterCreation** | Pointer to **NullableBool** | Specifies whether or not the session waits for the printers to be created before allowing the user to interact with the session. If not specified, will not be changed. | [optional] 

## Methods

### NewEditApplicationRequestModel

`func NewEditApplicationRequestModel() *EditApplicationRequestModel`

NewEditApplicationRequestModel instantiates a new EditApplicationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditApplicationRequestModelWithDefaults

`func NewEditApplicationRequestModelWithDefaults() *EditApplicationRequestModel`

NewEditApplicationRequestModelWithDefaults instantiates a new EditApplicationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFolder

`func (o *EditApplicationRequestModel) GetApplicationFolder() string`

GetApplicationFolder returns the ApplicationFolder field if non-nil, zero value otherwise.

### GetApplicationFolderOk

`func (o *EditApplicationRequestModel) GetApplicationFolderOk() (*string, bool)`

GetApplicationFolderOk returns a tuple with the ApplicationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFolder

`func (o *EditApplicationRequestModel) SetApplicationFolder(v string)`

SetApplicationFolder sets ApplicationFolder field to given value.

### HasApplicationFolder

`func (o *EditApplicationRequestModel) HasApplicationFolder() bool`

HasApplicationFolder returns a boolean if a field has been set.

### SetApplicationFolderNil

`func (o *EditApplicationRequestModel) SetApplicationFolderNil(b bool)`

 SetApplicationFolderNil sets the value for ApplicationFolder to be an explicit nil

### UnsetApplicationFolder
`func (o *EditApplicationRequestModel) UnsetApplicationFolder()`

UnsetApplicationFolder ensures that no value is present for ApplicationFolder, not even an explicit nil
### GetApplicationGroups

`func (o *EditApplicationRequestModel) GetApplicationGroups() []string`

GetApplicationGroups returns the ApplicationGroups field if non-nil, zero value otherwise.

### GetApplicationGroupsOk

`func (o *EditApplicationRequestModel) GetApplicationGroupsOk() (*[]string, bool)`

GetApplicationGroupsOk returns a tuple with the ApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroups

`func (o *EditApplicationRequestModel) SetApplicationGroups(v []string)`

SetApplicationGroups sets ApplicationGroups field to given value.

### HasApplicationGroups

`func (o *EditApplicationRequestModel) HasApplicationGroups() bool`

HasApplicationGroups returns a boolean if a field has been set.

### SetApplicationGroupsNil

`func (o *EditApplicationRequestModel) SetApplicationGroupsNil(b bool)`

 SetApplicationGroupsNil sets the value for ApplicationGroups to be an explicit nil

### UnsetApplicationGroups
`func (o *EditApplicationRequestModel) UnsetApplicationGroups()`

UnsetApplicationGroups ensures that no value is present for ApplicationGroups, not even an explicit nil
### GetDeliveryGroups

`func (o *EditApplicationRequestModel) GetDeliveryGroups() []PriorityRefRequestModel`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *EditApplicationRequestModel) GetDeliveryGroupsOk() (*[]PriorityRefRequestModel, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *EditApplicationRequestModel) SetDeliveryGroups(v []PriorityRefRequestModel)`

SetDeliveryGroups sets DeliveryGroups field to given value.

### HasDeliveryGroups

`func (o *EditApplicationRequestModel) HasDeliveryGroups() bool`

HasDeliveryGroups returns a boolean if a field has been set.

### SetDeliveryGroupsNil

`func (o *EditApplicationRequestModel) SetDeliveryGroupsNil(b bool)`

 SetDeliveryGroupsNil sets the value for DeliveryGroups to be an explicit nil

### UnsetDeliveryGroups
`func (o *EditApplicationRequestModel) UnsetDeliveryGroups()`

UnsetDeliveryGroups ensures that no value is present for DeliveryGroups, not even an explicit nil
### GetBrowserName

`func (o *EditApplicationRequestModel) GetBrowserName() string`

GetBrowserName returns the BrowserName field if non-nil, zero value otherwise.

### GetBrowserNameOk

`func (o *EditApplicationRequestModel) GetBrowserNameOk() (*string, bool)`

GetBrowserNameOk returns a tuple with the BrowserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserName

`func (o *EditApplicationRequestModel) SetBrowserName(v string)`

SetBrowserName sets BrowserName field to given value.

### HasBrowserName

`func (o *EditApplicationRequestModel) HasBrowserName() bool`

HasBrowserName returns a boolean if a field has been set.

### SetBrowserNameNil

`func (o *EditApplicationRequestModel) SetBrowserNameNil(b bool)`

 SetBrowserNameNil sets the value for BrowserName to be an explicit nil

### UnsetBrowserName
`func (o *EditApplicationRequestModel) UnsetBrowserName()`

UnsetBrowserName ensures that no value is present for BrowserName, not even an explicit nil
### GetCategories

`func (o *EditApplicationRequestModel) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *EditApplicationRequestModel) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *EditApplicationRequestModel) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *EditApplicationRequestModel) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *EditApplicationRequestModel) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *EditApplicationRequestModel) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetClientFolder

`func (o *EditApplicationRequestModel) GetClientFolder() string`

GetClientFolder returns the ClientFolder field if non-nil, zero value otherwise.

### GetClientFolderOk

`func (o *EditApplicationRequestModel) GetClientFolderOk() (*string, bool)`

GetClientFolderOk returns a tuple with the ClientFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFolder

`func (o *EditApplicationRequestModel) SetClientFolder(v string)`

SetClientFolder sets ClientFolder field to given value.

### HasClientFolder

`func (o *EditApplicationRequestModel) HasClientFolder() bool`

HasClientFolder returns a boolean if a field has been set.

### SetClientFolderNil

`func (o *EditApplicationRequestModel) SetClientFolderNil(b bool)`

 SetClientFolderNil sets the value for ClientFolder to be an explicit nil

### UnsetClientFolder
`func (o *EditApplicationRequestModel) UnsetClientFolder()`

UnsetClientFolder ensures that no value is present for ClientFolder, not even an explicit nil
### GetCpuPriorityLevel

`func (o *EditApplicationRequestModel) GetCpuPriorityLevel() CpuPriorityLevel`

GetCpuPriorityLevel returns the CpuPriorityLevel field if non-nil, zero value otherwise.

### GetCpuPriorityLevelOk

`func (o *EditApplicationRequestModel) GetCpuPriorityLevelOk() (*CpuPriorityLevel, bool)`

GetCpuPriorityLevelOk returns a tuple with the CpuPriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPriorityLevel

`func (o *EditApplicationRequestModel) SetCpuPriorityLevel(v CpuPriorityLevel)`

SetCpuPriorityLevel sets CpuPriorityLevel field to given value.

### HasCpuPriorityLevel

`func (o *EditApplicationRequestModel) HasCpuPriorityLevel() bool`

HasCpuPriorityLevel returns a boolean if a field has been set.

### GetDescription

`func (o *EditApplicationRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditApplicationRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditApplicationRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditApplicationRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EditApplicationRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditApplicationRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *EditApplicationRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EditApplicationRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EditApplicationRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EditApplicationRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *EditApplicationRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *EditApplicationRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetDoNotEnumerate

`func (o *EditApplicationRequestModel) GetDoNotEnumerate() bool`

GetDoNotEnumerate returns the DoNotEnumerate field if non-nil, zero value otherwise.

### GetDoNotEnumerateOk

`func (o *EditApplicationRequestModel) GetDoNotEnumerateOk() (*bool, bool)`

GetDoNotEnumerateOk returns a tuple with the DoNotEnumerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotEnumerate

`func (o *EditApplicationRequestModel) SetDoNotEnumerate(v bool)`

SetDoNotEnumerate sets DoNotEnumerate field to given value.

### HasDoNotEnumerate

`func (o *EditApplicationRequestModel) HasDoNotEnumerate() bool`

HasDoNotEnumerate returns a boolean if a field has been set.

### SetDoNotEnumerateNil

`func (o *EditApplicationRequestModel) SetDoNotEnumerateNil(b bool)`

 SetDoNotEnumerateNil sets the value for DoNotEnumerate to be an explicit nil

### UnsetDoNotEnumerate
`func (o *EditApplicationRequestModel) UnsetDoNotEnumerate()`

UnsetDoNotEnumerate ensures that no value is present for DoNotEnumerate, not even an explicit nil
### GetHomeZone

`func (o *EditApplicationRequestModel) GetHomeZone() string`

GetHomeZone returns the HomeZone field if non-nil, zero value otherwise.

### GetHomeZoneOk

`func (o *EditApplicationRequestModel) GetHomeZoneOk() (*string, bool)`

GetHomeZoneOk returns a tuple with the HomeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeZone

`func (o *EditApplicationRequestModel) SetHomeZone(v string)`

SetHomeZone sets HomeZone field to given value.

### HasHomeZone

`func (o *EditApplicationRequestModel) HasHomeZone() bool`

HasHomeZone returns a boolean if a field has been set.

### SetHomeZoneNil

`func (o *EditApplicationRequestModel) SetHomeZoneNil(b bool)`

 SetHomeZoneNil sets the value for HomeZone to be an explicit nil

### UnsetHomeZone
`func (o *EditApplicationRequestModel) UnsetHomeZone()`

UnsetHomeZone ensures that no value is present for HomeZone, not even an explicit nil
### GetHomeZoneMode

`func (o *EditApplicationRequestModel) GetHomeZoneMode() HomeZoneMode`

GetHomeZoneMode returns the HomeZoneMode field if non-nil, zero value otherwise.

### GetHomeZoneModeOk

`func (o *EditApplicationRequestModel) GetHomeZoneModeOk() (*HomeZoneMode, bool)`

GetHomeZoneModeOk returns a tuple with the HomeZoneMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeZoneMode

`func (o *EditApplicationRequestModel) SetHomeZoneMode(v HomeZoneMode)`

SetHomeZoneMode sets HomeZoneMode field to given value.

### HasHomeZoneMode

`func (o *EditApplicationRequestModel) HasHomeZoneMode() bool`

HasHomeZoneMode returns a boolean if a field has been set.

### GetIcon

`func (o *EditApplicationRequestModel) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EditApplicationRequestModel) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EditApplicationRequestModel) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *EditApplicationRequestModel) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *EditApplicationRequestModel) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *EditApplicationRequestModel) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetIconFromClient

`func (o *EditApplicationRequestModel) GetIconFromClient() bool`

GetIconFromClient returns the IconFromClient field if non-nil, zero value otherwise.

### GetIconFromClientOk

`func (o *EditApplicationRequestModel) GetIconFromClientOk() (*bool, bool)`

GetIconFromClientOk returns a tuple with the IconFromClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFromClient

`func (o *EditApplicationRequestModel) SetIconFromClient(v bool)`

SetIconFromClient sets IconFromClient field to given value.

### HasIconFromClient

`func (o *EditApplicationRequestModel) HasIconFromClient() bool`

HasIconFromClient returns a boolean if a field has been set.

### SetIconFromClientNil

`func (o *EditApplicationRequestModel) SetIconFromClientNil(b bool)`

 SetIconFromClientNil sets the value for IconFromClient to be an explicit nil

### UnsetIconFromClient
`func (o *EditApplicationRequestModel) UnsetIconFromClient()`

UnsetIconFromClient ensures that no value is present for IconFromClient, not even an explicit nil
### GetIncludedUserFilterEnabled

`func (o *EditApplicationRequestModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *EditApplicationRequestModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *EditApplicationRequestModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *EditApplicationRequestModel) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### SetIncludedUserFilterEnabledNil

`func (o *EditApplicationRequestModel) SetIncludedUserFilterEnabledNil(b bool)`

 SetIncludedUserFilterEnabledNil sets the value for IncludedUserFilterEnabled to be an explicit nil

### UnsetIncludedUserFilterEnabled
`func (o *EditApplicationRequestModel) UnsetIncludedUserFilterEnabled()`

UnsetIncludedUserFilterEnabled ensures that no value is present for IncludedUserFilterEnabled, not even an explicit nil
### GetIncludedUsers

`func (o *EditApplicationRequestModel) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *EditApplicationRequestModel) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *EditApplicationRequestModel) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *EditApplicationRequestModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *EditApplicationRequestModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *EditApplicationRequestModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil
### GetPackagedApplicationVisibility

`func (o *EditApplicationRequestModel) GetPackagedApplicationVisibility() PackagedApplicationVisibility`

GetPackagedApplicationVisibility returns the PackagedApplicationVisibility field if non-nil, zero value otherwise.

### GetPackagedApplicationVisibilityOk

`func (o *EditApplicationRequestModel) GetPackagedApplicationVisibilityOk() (*PackagedApplicationVisibility, bool)`

GetPackagedApplicationVisibilityOk returns a tuple with the PackagedApplicationVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagedApplicationVisibility

`func (o *EditApplicationRequestModel) SetPackagedApplicationVisibility(v PackagedApplicationVisibility)`

SetPackagedApplicationVisibility sets PackagedApplicationVisibility field to given value.

### HasPackagedApplicationVisibility

`func (o *EditApplicationRequestModel) HasPackagedApplicationVisibility() bool`

HasPackagedApplicationVisibility returns a boolean if a field has been set.

### GetInstalledAppProperties

`func (o *EditApplicationRequestModel) GetInstalledAppProperties() EditInstalledAppRequestModel`

GetInstalledAppProperties returns the InstalledAppProperties field if non-nil, zero value otherwise.

### GetInstalledAppPropertiesOk

`func (o *EditApplicationRequestModel) GetInstalledAppPropertiesOk() (*EditInstalledAppRequestModel, bool)`

GetInstalledAppPropertiesOk returns a tuple with the InstalledAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAppProperties

`func (o *EditApplicationRequestModel) SetInstalledAppProperties(v EditInstalledAppRequestModel)`

SetInstalledAppProperties sets InstalledAppProperties field to given value.

### HasInstalledAppProperties

`func (o *EditApplicationRequestModel) HasInstalledAppProperties() bool`

HasInstalledAppProperties returns a boolean if a field has been set.

### GetPackagedAppProperties

`func (o *EditApplicationRequestModel) GetPackagedAppProperties() AppVAppRequestModel`

GetPackagedAppProperties returns the PackagedAppProperties field if non-nil, zero value otherwise.

### GetPackagedAppPropertiesOk

`func (o *EditApplicationRequestModel) GetPackagedAppPropertiesOk() (*AppVAppRequestModel, bool)`

GetPackagedAppPropertiesOk returns a tuple with the PackagedAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagedAppProperties

`func (o *EditApplicationRequestModel) SetPackagedAppProperties(v AppVAppRequestModel)`

SetPackagedAppProperties sets PackagedAppProperties field to given value.

### HasPackagedAppProperties

`func (o *EditApplicationRequestModel) HasPackagedAppProperties() bool`

HasPackagedAppProperties returns a boolean if a field has been set.

### GetAppVAppProperties

`func (o *EditApplicationRequestModel) GetAppVAppProperties() AppVAppRequestModel`

GetAppVAppProperties returns the AppVAppProperties field if non-nil, zero value otherwise.

### GetAppVAppPropertiesOk

`func (o *EditApplicationRequestModel) GetAppVAppPropertiesOk() (*AppVAppRequestModel, bool)`

GetAppVAppPropertiesOk returns a tuple with the AppVAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVAppProperties

`func (o *EditApplicationRequestModel) SetAppVAppProperties(v AppVAppRequestModel)`

SetAppVAppProperties sets AppVAppProperties field to given value.

### HasAppVAppProperties

`func (o *EditApplicationRequestModel) HasAppVAppProperties() bool`

HasAppVAppProperties returns a boolean if a field has been set.

### GetContentLocation

`func (o *EditApplicationRequestModel) GetContentLocation() string`

GetContentLocation returns the ContentLocation field if non-nil, zero value otherwise.

### GetContentLocationOk

`func (o *EditApplicationRequestModel) GetContentLocationOk() (*string, bool)`

GetContentLocationOk returns a tuple with the ContentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLocation

`func (o *EditApplicationRequestModel) SetContentLocation(v string)`

SetContentLocation sets ContentLocation field to given value.

### HasContentLocation

`func (o *EditApplicationRequestModel) HasContentLocation() bool`

HasContentLocation returns a boolean if a field has been set.

### SetContentLocationNil

`func (o *EditApplicationRequestModel) SetContentLocationNil(b bool)`

 SetContentLocationNil sets the value for ContentLocation to be an explicit nil

### UnsetContentLocation
`func (o *EditApplicationRequestModel) UnsetContentLocation()`

UnsetContentLocation ensures that no value is present for ContentLocation, not even an explicit nil
### GetMetadata

`func (o *EditApplicationRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EditApplicationRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EditApplicationRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EditApplicationRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EditApplicationRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EditApplicationRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMaxPerUserInstances

`func (o *EditApplicationRequestModel) GetMaxPerUserInstances() int32`

GetMaxPerUserInstances returns the MaxPerUserInstances field if non-nil, zero value otherwise.

### GetMaxPerUserInstancesOk

`func (o *EditApplicationRequestModel) GetMaxPerUserInstancesOk() (*int32, bool)`

GetMaxPerUserInstancesOk returns a tuple with the MaxPerUserInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerUserInstances

`func (o *EditApplicationRequestModel) SetMaxPerUserInstances(v int32)`

SetMaxPerUserInstances sets MaxPerUserInstances field to given value.

### HasMaxPerUserInstances

`func (o *EditApplicationRequestModel) HasMaxPerUserInstances() bool`

HasMaxPerUserInstances returns a boolean if a field has been set.

### SetMaxPerUserInstancesNil

`func (o *EditApplicationRequestModel) SetMaxPerUserInstancesNil(b bool)`

 SetMaxPerUserInstancesNil sets the value for MaxPerUserInstances to be an explicit nil

### UnsetMaxPerUserInstances
`func (o *EditApplicationRequestModel) UnsetMaxPerUserInstances()`

UnsetMaxPerUserInstances ensures that no value is present for MaxPerUserInstances, not even an explicit nil
### GetMaxTotalInstances

`func (o *EditApplicationRequestModel) GetMaxTotalInstances() int32`

GetMaxTotalInstances returns the MaxTotalInstances field if non-nil, zero value otherwise.

### GetMaxTotalInstancesOk

`func (o *EditApplicationRequestModel) GetMaxTotalInstancesOk() (*int32, bool)`

GetMaxTotalInstancesOk returns a tuple with the MaxTotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalInstances

`func (o *EditApplicationRequestModel) SetMaxTotalInstances(v int32)`

SetMaxTotalInstances sets MaxTotalInstances field to given value.

### HasMaxTotalInstances

`func (o *EditApplicationRequestModel) HasMaxTotalInstances() bool`

HasMaxTotalInstances returns a boolean if a field has been set.

### SetMaxTotalInstancesNil

`func (o *EditApplicationRequestModel) SetMaxTotalInstancesNil(b bool)`

 SetMaxTotalInstancesNil sets the value for MaxTotalInstances to be an explicit nil

### UnsetMaxTotalInstances
`func (o *EditApplicationRequestModel) UnsetMaxTotalInstances()`

UnsetMaxTotalInstances ensures that no value is present for MaxTotalInstances, not even an explicit nil
### GetName

`func (o *EditApplicationRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditApplicationRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditApplicationRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditApplicationRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditApplicationRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditApplicationRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPublishedName

`func (o *EditApplicationRequestModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *EditApplicationRequestModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *EditApplicationRequestModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *EditApplicationRequestModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *EditApplicationRequestModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *EditApplicationRequestModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetShortcutAddedToDesktop

`func (o *EditApplicationRequestModel) GetShortcutAddedToDesktop() bool`

GetShortcutAddedToDesktop returns the ShortcutAddedToDesktop field if non-nil, zero value otherwise.

### GetShortcutAddedToDesktopOk

`func (o *EditApplicationRequestModel) GetShortcutAddedToDesktopOk() (*bool, bool)`

GetShortcutAddedToDesktopOk returns a tuple with the ShortcutAddedToDesktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutAddedToDesktop

`func (o *EditApplicationRequestModel) SetShortcutAddedToDesktop(v bool)`

SetShortcutAddedToDesktop sets ShortcutAddedToDesktop field to given value.

### HasShortcutAddedToDesktop

`func (o *EditApplicationRequestModel) HasShortcutAddedToDesktop() bool`

HasShortcutAddedToDesktop returns a boolean if a field has been set.

### SetShortcutAddedToDesktopNil

`func (o *EditApplicationRequestModel) SetShortcutAddedToDesktopNil(b bool)`

 SetShortcutAddedToDesktopNil sets the value for ShortcutAddedToDesktop to be an explicit nil

### UnsetShortcutAddedToDesktop
`func (o *EditApplicationRequestModel) UnsetShortcutAddedToDesktop()`

UnsetShortcutAddedToDesktop ensures that no value is present for ShortcutAddedToDesktop, not even an explicit nil
### GetShortcutAddedToStartMenu

`func (o *EditApplicationRequestModel) GetShortcutAddedToStartMenu() bool`

GetShortcutAddedToStartMenu returns the ShortcutAddedToStartMenu field if non-nil, zero value otherwise.

### GetShortcutAddedToStartMenuOk

`func (o *EditApplicationRequestModel) GetShortcutAddedToStartMenuOk() (*bool, bool)`

GetShortcutAddedToStartMenuOk returns a tuple with the ShortcutAddedToStartMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutAddedToStartMenu

`func (o *EditApplicationRequestModel) SetShortcutAddedToStartMenu(v bool)`

SetShortcutAddedToStartMenu sets ShortcutAddedToStartMenu field to given value.

### HasShortcutAddedToStartMenu

`func (o *EditApplicationRequestModel) HasShortcutAddedToStartMenu() bool`

HasShortcutAddedToStartMenu returns a boolean if a field has been set.

### SetShortcutAddedToStartMenuNil

`func (o *EditApplicationRequestModel) SetShortcutAddedToStartMenuNil(b bool)`

 SetShortcutAddedToStartMenuNil sets the value for ShortcutAddedToStartMenu to be an explicit nil

### UnsetShortcutAddedToStartMenu
`func (o *EditApplicationRequestModel) UnsetShortcutAddedToStartMenu()`

UnsetShortcutAddedToStartMenu ensures that no value is present for ShortcutAddedToStartMenu, not even an explicit nil
### GetStartMenuFolder

`func (o *EditApplicationRequestModel) GetStartMenuFolder() string`

GetStartMenuFolder returns the StartMenuFolder field if non-nil, zero value otherwise.

### GetStartMenuFolderOk

`func (o *EditApplicationRequestModel) GetStartMenuFolderOk() (*string, bool)`

GetStartMenuFolderOk returns a tuple with the StartMenuFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMenuFolder

`func (o *EditApplicationRequestModel) SetStartMenuFolder(v string)`

SetStartMenuFolder sets StartMenuFolder field to given value.

### HasStartMenuFolder

`func (o *EditApplicationRequestModel) HasStartMenuFolder() bool`

HasStartMenuFolder returns a boolean if a field has been set.

### SetStartMenuFolderNil

`func (o *EditApplicationRequestModel) SetStartMenuFolderNil(b bool)`

 SetStartMenuFolderNil sets the value for StartMenuFolder to be an explicit nil

### UnsetStartMenuFolder
`func (o *EditApplicationRequestModel) UnsetStartMenuFolder()`

UnsetStartMenuFolder ensures that no value is present for StartMenuFolder, not even an explicit nil
### GetVisible

`func (o *EditApplicationRequestModel) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *EditApplicationRequestModel) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *EditApplicationRequestModel) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *EditApplicationRequestModel) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### SetVisibleNil

`func (o *EditApplicationRequestModel) SetVisibleNil(b bool)`

 SetVisibleNil sets the value for Visible to be an explicit nil

### UnsetVisible
`func (o *EditApplicationRequestModel) UnsetVisible()`

UnsetVisible ensures that no value is present for Visible, not even an explicit nil
### GetWaitForPrinterCreation

`func (o *EditApplicationRequestModel) GetWaitForPrinterCreation() bool`

GetWaitForPrinterCreation returns the WaitForPrinterCreation field if non-nil, zero value otherwise.

### GetWaitForPrinterCreationOk

`func (o *EditApplicationRequestModel) GetWaitForPrinterCreationOk() (*bool, bool)`

GetWaitForPrinterCreationOk returns a tuple with the WaitForPrinterCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForPrinterCreation

`func (o *EditApplicationRequestModel) SetWaitForPrinterCreation(v bool)`

SetWaitForPrinterCreation sets WaitForPrinterCreation field to given value.

### HasWaitForPrinterCreation

`func (o *EditApplicationRequestModel) HasWaitForPrinterCreation() bool`

HasWaitForPrinterCreation returns a boolean if a field has been set.

### SetWaitForPrinterCreationNil

`func (o *EditApplicationRequestModel) SetWaitForPrinterCreationNil(b bool)`

 SetWaitForPrinterCreationNil sets the value for WaitForPrinterCreation to be an explicit nil

### UnsetWaitForPrinterCreation
`func (o *EditApplicationRequestModel) UnsetWaitForPrinterCreation()`

UnsetWaitForPrinterCreation ensures that no value is present for WaitForPrinterCreation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


