# ApplicationDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserName** | **string** | Internal name of the application. | 
**Categories** | Pointer to **[]string** | Categories in which the application resides. | [optional] 
**ConfiguredFtas** | Pointer to [**[]FtaResponseModel**](FtaResponseModel.md) | All file types which are associated with the application. | [optional] 
**CpuPriorityLevel** | [**CpuPriorityLevel**](CpuPriorityLevel.md) |  | 
**HomeZone** | [**ApplicationDetailResponseModelAllOfHomeZone**](ApplicationDetailResponseModelAllOfHomeZone.md) |  | 
**HomeZoneMode** | [**HomeZoneMode**](HomeZoneMode.md) |  | 
**IncludedUserFilterEnabled** | **bool** | Indicates whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access polic(ies) is implicitly granted access to the application. | 
**IconFromClient** | Pointer to **bool** | Specifies whether the icon is gotten from user&#39;s computer at run time. If not specified, will not be changed. | [optional] 
**IncludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The included users filter of the application; that is, the users and groups who are explicitly granted access to the published application. | [optional] 
**MaxPerUserInstances** | Pointer to **int32** | The maximum allowed concurrently running instances of the application that an individual user can have. | [optional] 
**MaxTotalInstances** | Pointer to **int32** | The maximum allowed total of concurrently running instances of the application within the site. | [optional] 
**ShortcutAddedToDesktop** | Pointer to **bool** | Indicates whether or not a shortcut to the application should be placed on the user device. | [optional] 
**ShortcutAddedToStartMenu** | Pointer to **bool** | Indicates whether or not a shortcut to the application should be placed in the user&#39;s start menu on their user device. | [optional] 
**StartMenuFolder** | Pointer to **string** | Name of the start menu folder that holds the application shortcut (if any). | [optional] 
**WaitForPrinterCreation** | Pointer to **bool** | Indicates whether or not the session waits for the printers to be created before allowing the user to interact with the session. | [optional] 

## Methods

### NewApplicationDetailResponseModelAllOf

`func NewApplicationDetailResponseModelAllOf(browserName string, cpuPriorityLevel CpuPriorityLevel, homeZone ApplicationDetailResponseModelAllOfHomeZone, homeZoneMode HomeZoneMode, includedUserFilterEnabled bool, ) *ApplicationDetailResponseModelAllOf`

NewApplicationDetailResponseModelAllOf instantiates a new ApplicationDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDetailResponseModelAllOfWithDefaults

`func NewApplicationDetailResponseModelAllOfWithDefaults() *ApplicationDetailResponseModelAllOf`

NewApplicationDetailResponseModelAllOfWithDefaults instantiates a new ApplicationDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserName

`func (o *ApplicationDetailResponseModelAllOf) GetBrowserName() string`

GetBrowserName returns the BrowserName field if non-nil, zero value otherwise.

### GetBrowserNameOk

`func (o *ApplicationDetailResponseModelAllOf) GetBrowserNameOk() (*string, bool)`

GetBrowserNameOk returns a tuple with the BrowserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserName

`func (o *ApplicationDetailResponseModelAllOf) SetBrowserName(v string)`

SetBrowserName sets BrowserName field to given value.


### GetCategories

`func (o *ApplicationDetailResponseModelAllOf) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ApplicationDetailResponseModelAllOf) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ApplicationDetailResponseModelAllOf) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ApplicationDetailResponseModelAllOf) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetConfiguredFtas

`func (o *ApplicationDetailResponseModelAllOf) GetConfiguredFtas() []FtaResponseModel`

GetConfiguredFtas returns the ConfiguredFtas field if non-nil, zero value otherwise.

### GetConfiguredFtasOk

`func (o *ApplicationDetailResponseModelAllOf) GetConfiguredFtasOk() (*[]FtaResponseModel, bool)`

GetConfiguredFtasOk returns a tuple with the ConfiguredFtas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredFtas

`func (o *ApplicationDetailResponseModelAllOf) SetConfiguredFtas(v []FtaResponseModel)`

SetConfiguredFtas sets ConfiguredFtas field to given value.

### HasConfiguredFtas

`func (o *ApplicationDetailResponseModelAllOf) HasConfiguredFtas() bool`

HasConfiguredFtas returns a boolean if a field has been set.

### GetCpuPriorityLevel

`func (o *ApplicationDetailResponseModelAllOf) GetCpuPriorityLevel() CpuPriorityLevel`

GetCpuPriorityLevel returns the CpuPriorityLevel field if non-nil, zero value otherwise.

### GetCpuPriorityLevelOk

`func (o *ApplicationDetailResponseModelAllOf) GetCpuPriorityLevelOk() (*CpuPriorityLevel, bool)`

GetCpuPriorityLevelOk returns a tuple with the CpuPriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPriorityLevel

`func (o *ApplicationDetailResponseModelAllOf) SetCpuPriorityLevel(v CpuPriorityLevel)`

SetCpuPriorityLevel sets CpuPriorityLevel field to given value.


### GetHomeZone

`func (o *ApplicationDetailResponseModelAllOf) GetHomeZone() ApplicationDetailResponseModelAllOfHomeZone`

GetHomeZone returns the HomeZone field if non-nil, zero value otherwise.

### GetHomeZoneOk

`func (o *ApplicationDetailResponseModelAllOf) GetHomeZoneOk() (*ApplicationDetailResponseModelAllOfHomeZone, bool)`

GetHomeZoneOk returns a tuple with the HomeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeZone

`func (o *ApplicationDetailResponseModelAllOf) SetHomeZone(v ApplicationDetailResponseModelAllOfHomeZone)`

SetHomeZone sets HomeZone field to given value.


### GetHomeZoneMode

`func (o *ApplicationDetailResponseModelAllOf) GetHomeZoneMode() HomeZoneMode`

GetHomeZoneMode returns the HomeZoneMode field if non-nil, zero value otherwise.

### GetHomeZoneModeOk

`func (o *ApplicationDetailResponseModelAllOf) GetHomeZoneModeOk() (*HomeZoneMode, bool)`

GetHomeZoneModeOk returns a tuple with the HomeZoneMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeZoneMode

`func (o *ApplicationDetailResponseModelAllOf) SetHomeZoneMode(v HomeZoneMode)`

SetHomeZoneMode sets HomeZoneMode field to given value.


### GetIncludedUserFilterEnabled

`func (o *ApplicationDetailResponseModelAllOf) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *ApplicationDetailResponseModelAllOf) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *ApplicationDetailResponseModelAllOf) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.


### GetIconFromClient

`func (o *ApplicationDetailResponseModelAllOf) GetIconFromClient() bool`

GetIconFromClient returns the IconFromClient field if non-nil, zero value otherwise.

### GetIconFromClientOk

`func (o *ApplicationDetailResponseModelAllOf) GetIconFromClientOk() (*bool, bool)`

GetIconFromClientOk returns a tuple with the IconFromClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFromClient

`func (o *ApplicationDetailResponseModelAllOf) SetIconFromClient(v bool)`

SetIconFromClient sets IconFromClient field to given value.

### HasIconFromClient

`func (o *ApplicationDetailResponseModelAllOf) HasIconFromClient() bool`

HasIconFromClient returns a boolean if a field has been set.

### GetIncludedUsers

`func (o *ApplicationDetailResponseModelAllOf) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *ApplicationDetailResponseModelAllOf) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *ApplicationDetailResponseModelAllOf) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *ApplicationDetailResponseModelAllOf) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### GetMaxPerUserInstances

`func (o *ApplicationDetailResponseModelAllOf) GetMaxPerUserInstances() int32`

GetMaxPerUserInstances returns the MaxPerUserInstances field if non-nil, zero value otherwise.

### GetMaxPerUserInstancesOk

`func (o *ApplicationDetailResponseModelAllOf) GetMaxPerUserInstancesOk() (*int32, bool)`

GetMaxPerUserInstancesOk returns a tuple with the MaxPerUserInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerUserInstances

`func (o *ApplicationDetailResponseModelAllOf) SetMaxPerUserInstances(v int32)`

SetMaxPerUserInstances sets MaxPerUserInstances field to given value.

### HasMaxPerUserInstances

`func (o *ApplicationDetailResponseModelAllOf) HasMaxPerUserInstances() bool`

HasMaxPerUserInstances returns a boolean if a field has been set.

### GetMaxTotalInstances

`func (o *ApplicationDetailResponseModelAllOf) GetMaxTotalInstances() int32`

GetMaxTotalInstances returns the MaxTotalInstances field if non-nil, zero value otherwise.

### GetMaxTotalInstancesOk

`func (o *ApplicationDetailResponseModelAllOf) GetMaxTotalInstancesOk() (*int32, bool)`

GetMaxTotalInstancesOk returns a tuple with the MaxTotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalInstances

`func (o *ApplicationDetailResponseModelAllOf) SetMaxTotalInstances(v int32)`

SetMaxTotalInstances sets MaxTotalInstances field to given value.

### HasMaxTotalInstances

`func (o *ApplicationDetailResponseModelAllOf) HasMaxTotalInstances() bool`

HasMaxTotalInstances returns a boolean if a field has been set.

### GetShortcutAddedToDesktop

`func (o *ApplicationDetailResponseModelAllOf) GetShortcutAddedToDesktop() bool`

GetShortcutAddedToDesktop returns the ShortcutAddedToDesktop field if non-nil, zero value otherwise.

### GetShortcutAddedToDesktopOk

`func (o *ApplicationDetailResponseModelAllOf) GetShortcutAddedToDesktopOk() (*bool, bool)`

GetShortcutAddedToDesktopOk returns a tuple with the ShortcutAddedToDesktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutAddedToDesktop

`func (o *ApplicationDetailResponseModelAllOf) SetShortcutAddedToDesktop(v bool)`

SetShortcutAddedToDesktop sets ShortcutAddedToDesktop field to given value.

### HasShortcutAddedToDesktop

`func (o *ApplicationDetailResponseModelAllOf) HasShortcutAddedToDesktop() bool`

HasShortcutAddedToDesktop returns a boolean if a field has been set.

### GetShortcutAddedToStartMenu

`func (o *ApplicationDetailResponseModelAllOf) GetShortcutAddedToStartMenu() bool`

GetShortcutAddedToStartMenu returns the ShortcutAddedToStartMenu field if non-nil, zero value otherwise.

### GetShortcutAddedToStartMenuOk

`func (o *ApplicationDetailResponseModelAllOf) GetShortcutAddedToStartMenuOk() (*bool, bool)`

GetShortcutAddedToStartMenuOk returns a tuple with the ShortcutAddedToStartMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutAddedToStartMenu

`func (o *ApplicationDetailResponseModelAllOf) SetShortcutAddedToStartMenu(v bool)`

SetShortcutAddedToStartMenu sets ShortcutAddedToStartMenu field to given value.

### HasShortcutAddedToStartMenu

`func (o *ApplicationDetailResponseModelAllOf) HasShortcutAddedToStartMenu() bool`

HasShortcutAddedToStartMenu returns a boolean if a field has been set.

### GetStartMenuFolder

`func (o *ApplicationDetailResponseModelAllOf) GetStartMenuFolder() string`

GetStartMenuFolder returns the StartMenuFolder field if non-nil, zero value otherwise.

### GetStartMenuFolderOk

`func (o *ApplicationDetailResponseModelAllOf) GetStartMenuFolderOk() (*string, bool)`

GetStartMenuFolderOk returns a tuple with the StartMenuFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMenuFolder

`func (o *ApplicationDetailResponseModelAllOf) SetStartMenuFolder(v string)`

SetStartMenuFolder sets StartMenuFolder field to given value.

### HasStartMenuFolder

`func (o *ApplicationDetailResponseModelAllOf) HasStartMenuFolder() bool`

HasStartMenuFolder returns a boolean if a field has been set.

### GetWaitForPrinterCreation

`func (o *ApplicationDetailResponseModelAllOf) GetWaitForPrinterCreation() bool`

GetWaitForPrinterCreation returns the WaitForPrinterCreation field if non-nil, zero value otherwise.

### GetWaitForPrinterCreationOk

`func (o *ApplicationDetailResponseModelAllOf) GetWaitForPrinterCreationOk() (*bool, bool)`

GetWaitForPrinterCreationOk returns a tuple with the WaitForPrinterCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForPrinterCreation

`func (o *ApplicationDetailResponseModelAllOf) SetWaitForPrinterCreation(v bool)`

SetWaitForPrinterCreation sets WaitForPrinterCreation field to given value.

### HasWaitForPrinterCreation

`func (o *ApplicationDetailResponseModelAllOf) HasWaitForPrinterCreation() bool`

HasWaitForPrinterCreation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


