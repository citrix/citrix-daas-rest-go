# ApplicationDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the application. Used to be: Uuid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | 
**Uid** | Pointer to **NullableInt32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED. Use Id. | [optional] 
**ApplicationFolder** | [**RefResponseModel**](RefResponseModel.md) |  | 
**ApplicationType** | [**ApplicationType**](ApplicationType.md) |  | 
**ClientFolder** | Pointer to **NullableString** | The folder that the application belongs to as the user sees it. | [optional] 
**ContainerScopes** | [**[]ContainerScopeResponseModel**](ContainerScopeResponseModel.md) | Delegated admin scopes in which the containers of the application reside. | 
**Description** | Pointer to **NullableString** | The description of the application. | [optional] 
**DoNotEnumerate** | Pointer to **bool** | Indicates whether or not this application is enumerable | [optional] 
**Enabled** | **bool** | Indicates whether or not this application can be launched. | 
**IconId** | **string** | Id of the icon used for the application. Used to be: IconUid (and it was not globally unique) Needs to be globally unique Might be constructed from site ID + internal Uid | 
**InstalledAppProperties** | Pointer to [**InstalledAppResponseModel**](InstalledAppResponseModel.md) |  | [optional] 
**AppVAppProperties** | Pointer to [**AppVAppResponseModel**](AppVAppResponseModel.md) |  | [optional] 
**ContentLocation** | Pointer to **NullableString** | Location of published content. | [optional] 
**Name** | **string** | Name of the application.  Only seen by administrators. | 
**PublishedName** | **string** | The name seen by end users who have access to the application. | 
**Visible** | **bool** | Indicates whether or not this application is visible to users. | 
**SharingKind** | [**SharingKind**](SharingKind.md) |  | 
**Tags** | Pointer to **[]string** | Tags associated with this application. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the application is assigned to.  If &#x60;null&#x60;, the application is not assigned to any tenants, and may be used by any tenant. | [optional] 
**CloudWorkspaceManaged** | Pointer to **bool** | Indicates whether users are managed in the Citrix Cloud Library, or within Studio. | [optional] 
**NumAssociatedDeliveryGroups** | Pointer to **int32** | Number of delivery groups that the application is associated with. | [optional] 
**NumAssociatedApplicationGroups** | Pointer to **int32** | Number of application groups that the application is associated with. | [optional] 
**AssociatedDeliveryGroupUuids** | Pointer to **[]string** | Delivery group Uuids that the application is associated with. | [optional] 
**AssociatedApplicationGroupUuids** | Pointer to **[]string** | Application group Uuids that the application is associated with. | [optional] 
**ZoneId** | Pointer to **NullableString** | Application Zone info. | [optional] 
**BrowserName** | **string** | Internal name of the application. | 
**Categories** | Pointer to **[]string** | Categories in which the application resides. | [optional] 
**ConfiguredFtas** | Pointer to [**[]FtaResponseModel**](FtaResponseModel.md) | All file types which are associated with the application. | [optional] 
**CpuPriorityLevel** | [**CpuPriorityLevel**](CpuPriorityLevel.md) |  | 
**HomeZone** | [**RefResponseModel**](RefResponseModel.md) |  | 
**HomeZoneMode** | [**HomeZoneMode**](HomeZoneMode.md) |  | 
**IncludedUserFilterEnabled** | **bool** | Indicates whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access polic(ies) is implicitly granted access to the application. | 
**IconFromClient** | Pointer to **NullableBool** | Specifies whether the icon is gotten from user&#39;s computer at run time. If not specified, will not be changed. | [optional] 
**IncludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The included users filter of the application; that is, the users and groups who are explicitly granted access to the published application. | [optional] 
**MaxPerUserInstances** | Pointer to **NullableInt32** | The maximum allowed concurrently running instances of the application that an individual user can have. | [optional] 
**MaxTotalInstances** | Pointer to **NullableInt32** | The maximum allowed total of concurrently running instances of the application within the site. | [optional] 
**ShortcutAddedToDesktop** | Pointer to **NullableBool** | Indicates whether or not a shortcut to the application should be placed on the user device. | [optional] 
**ShortcutAddedToStartMenu** | Pointer to **NullableBool** | Indicates whether or not a shortcut to the application should be placed in the user&#39;s start menu on their user device. | [optional] 
**StartMenuFolder** | Pointer to **NullableString** | Name of the start menu folder that holds the application shortcut (if any). | [optional] 
**WaitForPrinterCreation** | Pointer to **NullableBool** | Indicates whether or not the session waits for the printers to be created before allowing the user to interact with the session. | [optional] 

## Methods

### NewApplicationDetailResponseModel

`func NewApplicationDetailResponseModel(id string, applicationFolder RefResponseModel, applicationType ApplicationType, containerScopes []ContainerScopeResponseModel, enabled bool, iconId string, name string, publishedName string, visible bool, sharingKind SharingKind, browserName string, cpuPriorityLevel CpuPriorityLevel, homeZone RefResponseModel, homeZoneMode HomeZoneMode, includedUserFilterEnabled bool, ) *ApplicationDetailResponseModel`

NewApplicationDetailResponseModel instantiates a new ApplicationDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDetailResponseModelWithDefaults

`func NewApplicationDetailResponseModelWithDefaults() *ApplicationDetailResponseModel`

NewApplicationDetailResponseModelWithDefaults instantiates a new ApplicationDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *ApplicationDetailResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApplicationDetailResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApplicationDetailResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ApplicationDetailResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ApplicationDetailResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ApplicationDetailResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetApplicationFolder

`func (o *ApplicationDetailResponseModel) GetApplicationFolder() RefResponseModel`

GetApplicationFolder returns the ApplicationFolder field if non-nil, zero value otherwise.

### GetApplicationFolderOk

`func (o *ApplicationDetailResponseModel) GetApplicationFolderOk() (*RefResponseModel, bool)`

GetApplicationFolderOk returns a tuple with the ApplicationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFolder

`func (o *ApplicationDetailResponseModel) SetApplicationFolder(v RefResponseModel)`

SetApplicationFolder sets ApplicationFolder field to given value.


### GetApplicationType

`func (o *ApplicationDetailResponseModel) GetApplicationType() ApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *ApplicationDetailResponseModel) GetApplicationTypeOk() (*ApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *ApplicationDetailResponseModel) SetApplicationType(v ApplicationType)`

SetApplicationType sets ApplicationType field to given value.


### GetClientFolder

`func (o *ApplicationDetailResponseModel) GetClientFolder() string`

GetClientFolder returns the ClientFolder field if non-nil, zero value otherwise.

### GetClientFolderOk

`func (o *ApplicationDetailResponseModel) GetClientFolderOk() (*string, bool)`

GetClientFolderOk returns a tuple with the ClientFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFolder

`func (o *ApplicationDetailResponseModel) SetClientFolder(v string)`

SetClientFolder sets ClientFolder field to given value.

### HasClientFolder

`func (o *ApplicationDetailResponseModel) HasClientFolder() bool`

HasClientFolder returns a boolean if a field has been set.

### SetClientFolderNil

`func (o *ApplicationDetailResponseModel) SetClientFolderNil(b bool)`

 SetClientFolderNil sets the value for ClientFolder to be an explicit nil

### UnsetClientFolder
`func (o *ApplicationDetailResponseModel) UnsetClientFolder()`

UnsetClientFolder ensures that no value is present for ClientFolder, not even an explicit nil
### GetContainerScopes

`func (o *ApplicationDetailResponseModel) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *ApplicationDetailResponseModel) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *ApplicationDetailResponseModel) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.


### GetDescription

`func (o *ApplicationDetailResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationDetailResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationDetailResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationDetailResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationDetailResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationDetailResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDoNotEnumerate

`func (o *ApplicationDetailResponseModel) GetDoNotEnumerate() bool`

GetDoNotEnumerate returns the DoNotEnumerate field if non-nil, zero value otherwise.

### GetDoNotEnumerateOk

`func (o *ApplicationDetailResponseModel) GetDoNotEnumerateOk() (*bool, bool)`

GetDoNotEnumerateOk returns a tuple with the DoNotEnumerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotEnumerate

`func (o *ApplicationDetailResponseModel) SetDoNotEnumerate(v bool)`

SetDoNotEnumerate sets DoNotEnumerate field to given value.

### HasDoNotEnumerate

`func (o *ApplicationDetailResponseModel) HasDoNotEnumerate() bool`

HasDoNotEnumerate returns a boolean if a field has been set.

### GetEnabled

`func (o *ApplicationDetailResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationDetailResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationDetailResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIconId

`func (o *ApplicationDetailResponseModel) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *ApplicationDetailResponseModel) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *ApplicationDetailResponseModel) SetIconId(v string)`

SetIconId sets IconId field to given value.


### GetInstalledAppProperties

`func (o *ApplicationDetailResponseModel) GetInstalledAppProperties() InstalledAppResponseModel`

GetInstalledAppProperties returns the InstalledAppProperties field if non-nil, zero value otherwise.

### GetInstalledAppPropertiesOk

`func (o *ApplicationDetailResponseModel) GetInstalledAppPropertiesOk() (*InstalledAppResponseModel, bool)`

GetInstalledAppPropertiesOk returns a tuple with the InstalledAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAppProperties

`func (o *ApplicationDetailResponseModel) SetInstalledAppProperties(v InstalledAppResponseModel)`

SetInstalledAppProperties sets InstalledAppProperties field to given value.

### HasInstalledAppProperties

`func (o *ApplicationDetailResponseModel) HasInstalledAppProperties() bool`

HasInstalledAppProperties returns a boolean if a field has been set.

### GetAppVAppProperties

`func (o *ApplicationDetailResponseModel) GetAppVAppProperties() AppVAppResponseModel`

GetAppVAppProperties returns the AppVAppProperties field if non-nil, zero value otherwise.

### GetAppVAppPropertiesOk

`func (o *ApplicationDetailResponseModel) GetAppVAppPropertiesOk() (*AppVAppResponseModel, bool)`

GetAppVAppPropertiesOk returns a tuple with the AppVAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVAppProperties

`func (o *ApplicationDetailResponseModel) SetAppVAppProperties(v AppVAppResponseModel)`

SetAppVAppProperties sets AppVAppProperties field to given value.

### HasAppVAppProperties

`func (o *ApplicationDetailResponseModel) HasAppVAppProperties() bool`

HasAppVAppProperties returns a boolean if a field has been set.

### GetContentLocation

`func (o *ApplicationDetailResponseModel) GetContentLocation() string`

GetContentLocation returns the ContentLocation field if non-nil, zero value otherwise.

### GetContentLocationOk

`func (o *ApplicationDetailResponseModel) GetContentLocationOk() (*string, bool)`

GetContentLocationOk returns a tuple with the ContentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLocation

`func (o *ApplicationDetailResponseModel) SetContentLocation(v string)`

SetContentLocation sets ContentLocation field to given value.

### HasContentLocation

`func (o *ApplicationDetailResponseModel) HasContentLocation() bool`

HasContentLocation returns a boolean if a field has been set.

### SetContentLocationNil

`func (o *ApplicationDetailResponseModel) SetContentLocationNil(b bool)`

 SetContentLocationNil sets the value for ContentLocation to be an explicit nil

### UnsetContentLocation
`func (o *ApplicationDetailResponseModel) UnsetContentLocation()`

UnsetContentLocation ensures that no value is present for ContentLocation, not even an explicit nil
### GetName

`func (o *ApplicationDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetPublishedName

`func (o *ApplicationDetailResponseModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *ApplicationDetailResponseModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *ApplicationDetailResponseModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.


### GetVisible

`func (o *ApplicationDetailResponseModel) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ApplicationDetailResponseModel) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ApplicationDetailResponseModel) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetSharingKind

`func (o *ApplicationDetailResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *ApplicationDetailResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *ApplicationDetailResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.


### GetTags

`func (o *ApplicationDetailResponseModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationDetailResponseModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationDetailResponseModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationDetailResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApplicationDetailResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApplicationDetailResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTenants

`func (o *ApplicationDetailResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApplicationDetailResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApplicationDetailResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApplicationDetailResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *ApplicationDetailResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *ApplicationDetailResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetCloudWorkspaceManaged

`func (o *ApplicationDetailResponseModel) GetCloudWorkspaceManaged() bool`

GetCloudWorkspaceManaged returns the CloudWorkspaceManaged field if non-nil, zero value otherwise.

### GetCloudWorkspaceManagedOk

`func (o *ApplicationDetailResponseModel) GetCloudWorkspaceManagedOk() (*bool, bool)`

GetCloudWorkspaceManagedOk returns a tuple with the CloudWorkspaceManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudWorkspaceManaged

`func (o *ApplicationDetailResponseModel) SetCloudWorkspaceManaged(v bool)`

SetCloudWorkspaceManaged sets CloudWorkspaceManaged field to given value.

### HasCloudWorkspaceManaged

`func (o *ApplicationDetailResponseModel) HasCloudWorkspaceManaged() bool`

HasCloudWorkspaceManaged returns a boolean if a field has been set.

### GetNumAssociatedDeliveryGroups

`func (o *ApplicationDetailResponseModel) GetNumAssociatedDeliveryGroups() int32`

GetNumAssociatedDeliveryGroups returns the NumAssociatedDeliveryGroups field if non-nil, zero value otherwise.

### GetNumAssociatedDeliveryGroupsOk

`func (o *ApplicationDetailResponseModel) GetNumAssociatedDeliveryGroupsOk() (*int32, bool)`

GetNumAssociatedDeliveryGroupsOk returns a tuple with the NumAssociatedDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssociatedDeliveryGroups

`func (o *ApplicationDetailResponseModel) SetNumAssociatedDeliveryGroups(v int32)`

SetNumAssociatedDeliveryGroups sets NumAssociatedDeliveryGroups field to given value.

### HasNumAssociatedDeliveryGroups

`func (o *ApplicationDetailResponseModel) HasNumAssociatedDeliveryGroups() bool`

HasNumAssociatedDeliveryGroups returns a boolean if a field has been set.

### GetNumAssociatedApplicationGroups

`func (o *ApplicationDetailResponseModel) GetNumAssociatedApplicationGroups() int32`

GetNumAssociatedApplicationGroups returns the NumAssociatedApplicationGroups field if non-nil, zero value otherwise.

### GetNumAssociatedApplicationGroupsOk

`func (o *ApplicationDetailResponseModel) GetNumAssociatedApplicationGroupsOk() (*int32, bool)`

GetNumAssociatedApplicationGroupsOk returns a tuple with the NumAssociatedApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssociatedApplicationGroups

`func (o *ApplicationDetailResponseModel) SetNumAssociatedApplicationGroups(v int32)`

SetNumAssociatedApplicationGroups sets NumAssociatedApplicationGroups field to given value.

### HasNumAssociatedApplicationGroups

`func (o *ApplicationDetailResponseModel) HasNumAssociatedApplicationGroups() bool`

HasNumAssociatedApplicationGroups returns a boolean if a field has been set.

### GetAssociatedDeliveryGroupUuids

`func (o *ApplicationDetailResponseModel) GetAssociatedDeliveryGroupUuids() []string`

GetAssociatedDeliveryGroupUuids returns the AssociatedDeliveryGroupUuids field if non-nil, zero value otherwise.

### GetAssociatedDeliveryGroupUuidsOk

`func (o *ApplicationDetailResponseModel) GetAssociatedDeliveryGroupUuidsOk() (*[]string, bool)`

GetAssociatedDeliveryGroupUuidsOk returns a tuple with the AssociatedDeliveryGroupUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDeliveryGroupUuids

`func (o *ApplicationDetailResponseModel) SetAssociatedDeliveryGroupUuids(v []string)`

SetAssociatedDeliveryGroupUuids sets AssociatedDeliveryGroupUuids field to given value.

### HasAssociatedDeliveryGroupUuids

`func (o *ApplicationDetailResponseModel) HasAssociatedDeliveryGroupUuids() bool`

HasAssociatedDeliveryGroupUuids returns a boolean if a field has been set.

### SetAssociatedDeliveryGroupUuidsNil

`func (o *ApplicationDetailResponseModel) SetAssociatedDeliveryGroupUuidsNil(b bool)`

 SetAssociatedDeliveryGroupUuidsNil sets the value for AssociatedDeliveryGroupUuids to be an explicit nil

### UnsetAssociatedDeliveryGroupUuids
`func (o *ApplicationDetailResponseModel) UnsetAssociatedDeliveryGroupUuids()`

UnsetAssociatedDeliveryGroupUuids ensures that no value is present for AssociatedDeliveryGroupUuids, not even an explicit nil
### GetAssociatedApplicationGroupUuids

`func (o *ApplicationDetailResponseModel) GetAssociatedApplicationGroupUuids() []string`

GetAssociatedApplicationGroupUuids returns the AssociatedApplicationGroupUuids field if non-nil, zero value otherwise.

### GetAssociatedApplicationGroupUuidsOk

`func (o *ApplicationDetailResponseModel) GetAssociatedApplicationGroupUuidsOk() (*[]string, bool)`

GetAssociatedApplicationGroupUuidsOk returns a tuple with the AssociatedApplicationGroupUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedApplicationGroupUuids

`func (o *ApplicationDetailResponseModel) SetAssociatedApplicationGroupUuids(v []string)`

SetAssociatedApplicationGroupUuids sets AssociatedApplicationGroupUuids field to given value.

### HasAssociatedApplicationGroupUuids

`func (o *ApplicationDetailResponseModel) HasAssociatedApplicationGroupUuids() bool`

HasAssociatedApplicationGroupUuids returns a boolean if a field has been set.

### SetAssociatedApplicationGroupUuidsNil

`func (o *ApplicationDetailResponseModel) SetAssociatedApplicationGroupUuidsNil(b bool)`

 SetAssociatedApplicationGroupUuidsNil sets the value for AssociatedApplicationGroupUuids to be an explicit nil

### UnsetAssociatedApplicationGroupUuids
`func (o *ApplicationDetailResponseModel) UnsetAssociatedApplicationGroupUuids()`

UnsetAssociatedApplicationGroupUuids ensures that no value is present for AssociatedApplicationGroupUuids, not even an explicit nil
### GetZoneId

`func (o *ApplicationDetailResponseModel) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ApplicationDetailResponseModel) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ApplicationDetailResponseModel) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ApplicationDetailResponseModel) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *ApplicationDetailResponseModel) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *ApplicationDetailResponseModel) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetBrowserName

`func (o *ApplicationDetailResponseModel) GetBrowserName() string`

GetBrowserName returns the BrowserName field if non-nil, zero value otherwise.

### GetBrowserNameOk

`func (o *ApplicationDetailResponseModel) GetBrowserNameOk() (*string, bool)`

GetBrowserNameOk returns a tuple with the BrowserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserName

`func (o *ApplicationDetailResponseModel) SetBrowserName(v string)`

SetBrowserName sets BrowserName field to given value.


### GetCategories

`func (o *ApplicationDetailResponseModel) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ApplicationDetailResponseModel) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ApplicationDetailResponseModel) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ApplicationDetailResponseModel) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *ApplicationDetailResponseModel) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *ApplicationDetailResponseModel) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetConfiguredFtas

`func (o *ApplicationDetailResponseModel) GetConfiguredFtas() []FtaResponseModel`

GetConfiguredFtas returns the ConfiguredFtas field if non-nil, zero value otherwise.

### GetConfiguredFtasOk

`func (o *ApplicationDetailResponseModel) GetConfiguredFtasOk() (*[]FtaResponseModel, bool)`

GetConfiguredFtasOk returns a tuple with the ConfiguredFtas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredFtas

`func (o *ApplicationDetailResponseModel) SetConfiguredFtas(v []FtaResponseModel)`

SetConfiguredFtas sets ConfiguredFtas field to given value.

### HasConfiguredFtas

`func (o *ApplicationDetailResponseModel) HasConfiguredFtas() bool`

HasConfiguredFtas returns a boolean if a field has been set.

### SetConfiguredFtasNil

`func (o *ApplicationDetailResponseModel) SetConfiguredFtasNil(b bool)`

 SetConfiguredFtasNil sets the value for ConfiguredFtas to be an explicit nil

### UnsetConfiguredFtas
`func (o *ApplicationDetailResponseModel) UnsetConfiguredFtas()`

UnsetConfiguredFtas ensures that no value is present for ConfiguredFtas, not even an explicit nil
### GetCpuPriorityLevel

`func (o *ApplicationDetailResponseModel) GetCpuPriorityLevel() CpuPriorityLevel`

GetCpuPriorityLevel returns the CpuPriorityLevel field if non-nil, zero value otherwise.

### GetCpuPriorityLevelOk

`func (o *ApplicationDetailResponseModel) GetCpuPriorityLevelOk() (*CpuPriorityLevel, bool)`

GetCpuPriorityLevelOk returns a tuple with the CpuPriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPriorityLevel

`func (o *ApplicationDetailResponseModel) SetCpuPriorityLevel(v CpuPriorityLevel)`

SetCpuPriorityLevel sets CpuPriorityLevel field to given value.


### GetHomeZone

`func (o *ApplicationDetailResponseModel) GetHomeZone() RefResponseModel`

GetHomeZone returns the HomeZone field if non-nil, zero value otherwise.

### GetHomeZoneOk

`func (o *ApplicationDetailResponseModel) GetHomeZoneOk() (*RefResponseModel, bool)`

GetHomeZoneOk returns a tuple with the HomeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeZone

`func (o *ApplicationDetailResponseModel) SetHomeZone(v RefResponseModel)`

SetHomeZone sets HomeZone field to given value.


### GetHomeZoneMode

`func (o *ApplicationDetailResponseModel) GetHomeZoneMode() HomeZoneMode`

GetHomeZoneMode returns the HomeZoneMode field if non-nil, zero value otherwise.

### GetHomeZoneModeOk

`func (o *ApplicationDetailResponseModel) GetHomeZoneModeOk() (*HomeZoneMode, bool)`

GetHomeZoneModeOk returns a tuple with the HomeZoneMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeZoneMode

`func (o *ApplicationDetailResponseModel) SetHomeZoneMode(v HomeZoneMode)`

SetHomeZoneMode sets HomeZoneMode field to given value.


### GetIncludedUserFilterEnabled

`func (o *ApplicationDetailResponseModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *ApplicationDetailResponseModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *ApplicationDetailResponseModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.


### GetIconFromClient

`func (o *ApplicationDetailResponseModel) GetIconFromClient() bool`

GetIconFromClient returns the IconFromClient field if non-nil, zero value otherwise.

### GetIconFromClientOk

`func (o *ApplicationDetailResponseModel) GetIconFromClientOk() (*bool, bool)`

GetIconFromClientOk returns a tuple with the IconFromClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFromClient

`func (o *ApplicationDetailResponseModel) SetIconFromClient(v bool)`

SetIconFromClient sets IconFromClient field to given value.

### HasIconFromClient

`func (o *ApplicationDetailResponseModel) HasIconFromClient() bool`

HasIconFromClient returns a boolean if a field has been set.

### SetIconFromClientNil

`func (o *ApplicationDetailResponseModel) SetIconFromClientNil(b bool)`

 SetIconFromClientNil sets the value for IconFromClient to be an explicit nil

### UnsetIconFromClient
`func (o *ApplicationDetailResponseModel) UnsetIconFromClient()`

UnsetIconFromClient ensures that no value is present for IconFromClient, not even an explicit nil
### GetIncludedUsers

`func (o *ApplicationDetailResponseModel) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *ApplicationDetailResponseModel) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *ApplicationDetailResponseModel) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *ApplicationDetailResponseModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *ApplicationDetailResponseModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *ApplicationDetailResponseModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil
### GetMaxPerUserInstances

`func (o *ApplicationDetailResponseModel) GetMaxPerUserInstances() int32`

GetMaxPerUserInstances returns the MaxPerUserInstances field if non-nil, zero value otherwise.

### GetMaxPerUserInstancesOk

`func (o *ApplicationDetailResponseModel) GetMaxPerUserInstancesOk() (*int32, bool)`

GetMaxPerUserInstancesOk returns a tuple with the MaxPerUserInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerUserInstances

`func (o *ApplicationDetailResponseModel) SetMaxPerUserInstances(v int32)`

SetMaxPerUserInstances sets MaxPerUserInstances field to given value.

### HasMaxPerUserInstances

`func (o *ApplicationDetailResponseModel) HasMaxPerUserInstances() bool`

HasMaxPerUserInstances returns a boolean if a field has been set.

### SetMaxPerUserInstancesNil

`func (o *ApplicationDetailResponseModel) SetMaxPerUserInstancesNil(b bool)`

 SetMaxPerUserInstancesNil sets the value for MaxPerUserInstances to be an explicit nil

### UnsetMaxPerUserInstances
`func (o *ApplicationDetailResponseModel) UnsetMaxPerUserInstances()`

UnsetMaxPerUserInstances ensures that no value is present for MaxPerUserInstances, not even an explicit nil
### GetMaxTotalInstances

`func (o *ApplicationDetailResponseModel) GetMaxTotalInstances() int32`

GetMaxTotalInstances returns the MaxTotalInstances field if non-nil, zero value otherwise.

### GetMaxTotalInstancesOk

`func (o *ApplicationDetailResponseModel) GetMaxTotalInstancesOk() (*int32, bool)`

GetMaxTotalInstancesOk returns a tuple with the MaxTotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalInstances

`func (o *ApplicationDetailResponseModel) SetMaxTotalInstances(v int32)`

SetMaxTotalInstances sets MaxTotalInstances field to given value.

### HasMaxTotalInstances

`func (o *ApplicationDetailResponseModel) HasMaxTotalInstances() bool`

HasMaxTotalInstances returns a boolean if a field has been set.

### SetMaxTotalInstancesNil

`func (o *ApplicationDetailResponseModel) SetMaxTotalInstancesNil(b bool)`

 SetMaxTotalInstancesNil sets the value for MaxTotalInstances to be an explicit nil

### UnsetMaxTotalInstances
`func (o *ApplicationDetailResponseModel) UnsetMaxTotalInstances()`

UnsetMaxTotalInstances ensures that no value is present for MaxTotalInstances, not even an explicit nil
### GetShortcutAddedToDesktop

`func (o *ApplicationDetailResponseModel) GetShortcutAddedToDesktop() bool`

GetShortcutAddedToDesktop returns the ShortcutAddedToDesktop field if non-nil, zero value otherwise.

### GetShortcutAddedToDesktopOk

`func (o *ApplicationDetailResponseModel) GetShortcutAddedToDesktopOk() (*bool, bool)`

GetShortcutAddedToDesktopOk returns a tuple with the ShortcutAddedToDesktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutAddedToDesktop

`func (o *ApplicationDetailResponseModel) SetShortcutAddedToDesktop(v bool)`

SetShortcutAddedToDesktop sets ShortcutAddedToDesktop field to given value.

### HasShortcutAddedToDesktop

`func (o *ApplicationDetailResponseModel) HasShortcutAddedToDesktop() bool`

HasShortcutAddedToDesktop returns a boolean if a field has been set.

### SetShortcutAddedToDesktopNil

`func (o *ApplicationDetailResponseModel) SetShortcutAddedToDesktopNil(b bool)`

 SetShortcutAddedToDesktopNil sets the value for ShortcutAddedToDesktop to be an explicit nil

### UnsetShortcutAddedToDesktop
`func (o *ApplicationDetailResponseModel) UnsetShortcutAddedToDesktop()`

UnsetShortcutAddedToDesktop ensures that no value is present for ShortcutAddedToDesktop, not even an explicit nil
### GetShortcutAddedToStartMenu

`func (o *ApplicationDetailResponseModel) GetShortcutAddedToStartMenu() bool`

GetShortcutAddedToStartMenu returns the ShortcutAddedToStartMenu field if non-nil, zero value otherwise.

### GetShortcutAddedToStartMenuOk

`func (o *ApplicationDetailResponseModel) GetShortcutAddedToStartMenuOk() (*bool, bool)`

GetShortcutAddedToStartMenuOk returns a tuple with the ShortcutAddedToStartMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutAddedToStartMenu

`func (o *ApplicationDetailResponseModel) SetShortcutAddedToStartMenu(v bool)`

SetShortcutAddedToStartMenu sets ShortcutAddedToStartMenu field to given value.

### HasShortcutAddedToStartMenu

`func (o *ApplicationDetailResponseModel) HasShortcutAddedToStartMenu() bool`

HasShortcutAddedToStartMenu returns a boolean if a field has been set.

### SetShortcutAddedToStartMenuNil

`func (o *ApplicationDetailResponseModel) SetShortcutAddedToStartMenuNil(b bool)`

 SetShortcutAddedToStartMenuNil sets the value for ShortcutAddedToStartMenu to be an explicit nil

### UnsetShortcutAddedToStartMenu
`func (o *ApplicationDetailResponseModel) UnsetShortcutAddedToStartMenu()`

UnsetShortcutAddedToStartMenu ensures that no value is present for ShortcutAddedToStartMenu, not even an explicit nil
### GetStartMenuFolder

`func (o *ApplicationDetailResponseModel) GetStartMenuFolder() string`

GetStartMenuFolder returns the StartMenuFolder field if non-nil, zero value otherwise.

### GetStartMenuFolderOk

`func (o *ApplicationDetailResponseModel) GetStartMenuFolderOk() (*string, bool)`

GetStartMenuFolderOk returns a tuple with the StartMenuFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMenuFolder

`func (o *ApplicationDetailResponseModel) SetStartMenuFolder(v string)`

SetStartMenuFolder sets StartMenuFolder field to given value.

### HasStartMenuFolder

`func (o *ApplicationDetailResponseModel) HasStartMenuFolder() bool`

HasStartMenuFolder returns a boolean if a field has been set.

### SetStartMenuFolderNil

`func (o *ApplicationDetailResponseModel) SetStartMenuFolderNil(b bool)`

 SetStartMenuFolderNil sets the value for StartMenuFolder to be an explicit nil

### UnsetStartMenuFolder
`func (o *ApplicationDetailResponseModel) UnsetStartMenuFolder()`

UnsetStartMenuFolder ensures that no value is present for StartMenuFolder, not even an explicit nil
### GetWaitForPrinterCreation

`func (o *ApplicationDetailResponseModel) GetWaitForPrinterCreation() bool`

GetWaitForPrinterCreation returns the WaitForPrinterCreation field if non-nil, zero value otherwise.

### GetWaitForPrinterCreationOk

`func (o *ApplicationDetailResponseModel) GetWaitForPrinterCreationOk() (*bool, bool)`

GetWaitForPrinterCreationOk returns a tuple with the WaitForPrinterCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForPrinterCreation

`func (o *ApplicationDetailResponseModel) SetWaitForPrinterCreation(v bool)`

SetWaitForPrinterCreation sets WaitForPrinterCreation field to given value.

### HasWaitForPrinterCreation

`func (o *ApplicationDetailResponseModel) HasWaitForPrinterCreation() bool`

HasWaitForPrinterCreation returns a boolean if a field has been set.

### SetWaitForPrinterCreationNil

`func (o *ApplicationDetailResponseModel) SetWaitForPrinterCreationNil(b bool)`

 SetWaitForPrinterCreationNil sets the value for WaitForPrinterCreation to be an explicit nil

### UnsetWaitForPrinterCreation
`func (o *ApplicationDetailResponseModel) UnsetWaitForPrinterCreation()`

UnsetWaitForPrinterCreation ensures that no value is present for WaitForPrinterCreation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


