# ApplicationResponseModel

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

## Methods

### NewApplicationResponseModel

`func NewApplicationResponseModel(id string, applicationFolder RefResponseModel, applicationType ApplicationType, containerScopes []ContainerScopeResponseModel, enabled bool, iconId string, name string, publishedName string, visible bool, sharingKind SharingKind, ) *ApplicationResponseModel`

NewApplicationResponseModel instantiates a new ApplicationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseModelWithDefaults

`func NewApplicationResponseModelWithDefaults() *ApplicationResponseModel`

NewApplicationResponseModelWithDefaults instantiates a new ApplicationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *ApplicationResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApplicationResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApplicationResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ApplicationResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ApplicationResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ApplicationResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetApplicationFolder

`func (o *ApplicationResponseModel) GetApplicationFolder() RefResponseModel`

GetApplicationFolder returns the ApplicationFolder field if non-nil, zero value otherwise.

### GetApplicationFolderOk

`func (o *ApplicationResponseModel) GetApplicationFolderOk() (*RefResponseModel, bool)`

GetApplicationFolderOk returns a tuple with the ApplicationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFolder

`func (o *ApplicationResponseModel) SetApplicationFolder(v RefResponseModel)`

SetApplicationFolder sets ApplicationFolder field to given value.


### GetApplicationType

`func (o *ApplicationResponseModel) GetApplicationType() ApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *ApplicationResponseModel) GetApplicationTypeOk() (*ApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *ApplicationResponseModel) SetApplicationType(v ApplicationType)`

SetApplicationType sets ApplicationType field to given value.


### GetClientFolder

`func (o *ApplicationResponseModel) GetClientFolder() string`

GetClientFolder returns the ClientFolder field if non-nil, zero value otherwise.

### GetClientFolderOk

`func (o *ApplicationResponseModel) GetClientFolderOk() (*string, bool)`

GetClientFolderOk returns a tuple with the ClientFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFolder

`func (o *ApplicationResponseModel) SetClientFolder(v string)`

SetClientFolder sets ClientFolder field to given value.

### HasClientFolder

`func (o *ApplicationResponseModel) HasClientFolder() bool`

HasClientFolder returns a boolean if a field has been set.

### SetClientFolderNil

`func (o *ApplicationResponseModel) SetClientFolderNil(b bool)`

 SetClientFolderNil sets the value for ClientFolder to be an explicit nil

### UnsetClientFolder
`func (o *ApplicationResponseModel) UnsetClientFolder()`

UnsetClientFolder ensures that no value is present for ClientFolder, not even an explicit nil
### GetContainerScopes

`func (o *ApplicationResponseModel) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *ApplicationResponseModel) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *ApplicationResponseModel) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.


### GetDescription

`func (o *ApplicationResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDoNotEnumerate

`func (o *ApplicationResponseModel) GetDoNotEnumerate() bool`

GetDoNotEnumerate returns the DoNotEnumerate field if non-nil, zero value otherwise.

### GetDoNotEnumerateOk

`func (o *ApplicationResponseModel) GetDoNotEnumerateOk() (*bool, bool)`

GetDoNotEnumerateOk returns a tuple with the DoNotEnumerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotEnumerate

`func (o *ApplicationResponseModel) SetDoNotEnumerate(v bool)`

SetDoNotEnumerate sets DoNotEnumerate field to given value.

### HasDoNotEnumerate

`func (o *ApplicationResponseModel) HasDoNotEnumerate() bool`

HasDoNotEnumerate returns a boolean if a field has been set.

### GetEnabled

`func (o *ApplicationResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIconId

`func (o *ApplicationResponseModel) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *ApplicationResponseModel) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *ApplicationResponseModel) SetIconId(v string)`

SetIconId sets IconId field to given value.


### GetInstalledAppProperties

`func (o *ApplicationResponseModel) GetInstalledAppProperties() InstalledAppResponseModel`

GetInstalledAppProperties returns the InstalledAppProperties field if non-nil, zero value otherwise.

### GetInstalledAppPropertiesOk

`func (o *ApplicationResponseModel) GetInstalledAppPropertiesOk() (*InstalledAppResponseModel, bool)`

GetInstalledAppPropertiesOk returns a tuple with the InstalledAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAppProperties

`func (o *ApplicationResponseModel) SetInstalledAppProperties(v InstalledAppResponseModel)`

SetInstalledAppProperties sets InstalledAppProperties field to given value.

### HasInstalledAppProperties

`func (o *ApplicationResponseModel) HasInstalledAppProperties() bool`

HasInstalledAppProperties returns a boolean if a field has been set.

### GetAppVAppProperties

`func (o *ApplicationResponseModel) GetAppVAppProperties() AppVAppResponseModel`

GetAppVAppProperties returns the AppVAppProperties field if non-nil, zero value otherwise.

### GetAppVAppPropertiesOk

`func (o *ApplicationResponseModel) GetAppVAppPropertiesOk() (*AppVAppResponseModel, bool)`

GetAppVAppPropertiesOk returns a tuple with the AppVAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVAppProperties

`func (o *ApplicationResponseModel) SetAppVAppProperties(v AppVAppResponseModel)`

SetAppVAppProperties sets AppVAppProperties field to given value.

### HasAppVAppProperties

`func (o *ApplicationResponseModel) HasAppVAppProperties() bool`

HasAppVAppProperties returns a boolean if a field has been set.

### GetContentLocation

`func (o *ApplicationResponseModel) GetContentLocation() string`

GetContentLocation returns the ContentLocation field if non-nil, zero value otherwise.

### GetContentLocationOk

`func (o *ApplicationResponseModel) GetContentLocationOk() (*string, bool)`

GetContentLocationOk returns a tuple with the ContentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLocation

`func (o *ApplicationResponseModel) SetContentLocation(v string)`

SetContentLocation sets ContentLocation field to given value.

### HasContentLocation

`func (o *ApplicationResponseModel) HasContentLocation() bool`

HasContentLocation returns a boolean if a field has been set.

### SetContentLocationNil

`func (o *ApplicationResponseModel) SetContentLocationNil(b bool)`

 SetContentLocationNil sets the value for ContentLocation to be an explicit nil

### UnsetContentLocation
`func (o *ApplicationResponseModel) UnsetContentLocation()`

UnsetContentLocation ensures that no value is present for ContentLocation, not even an explicit nil
### GetName

`func (o *ApplicationResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetPublishedName

`func (o *ApplicationResponseModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *ApplicationResponseModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *ApplicationResponseModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.


### GetVisible

`func (o *ApplicationResponseModel) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ApplicationResponseModel) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ApplicationResponseModel) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetSharingKind

`func (o *ApplicationResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *ApplicationResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *ApplicationResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.


### GetTags

`func (o *ApplicationResponseModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationResponseModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationResponseModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApplicationResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApplicationResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTenants

`func (o *ApplicationResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApplicationResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApplicationResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApplicationResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *ApplicationResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *ApplicationResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetCloudWorkspaceManaged

`func (o *ApplicationResponseModel) GetCloudWorkspaceManaged() bool`

GetCloudWorkspaceManaged returns the CloudWorkspaceManaged field if non-nil, zero value otherwise.

### GetCloudWorkspaceManagedOk

`func (o *ApplicationResponseModel) GetCloudWorkspaceManagedOk() (*bool, bool)`

GetCloudWorkspaceManagedOk returns a tuple with the CloudWorkspaceManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudWorkspaceManaged

`func (o *ApplicationResponseModel) SetCloudWorkspaceManaged(v bool)`

SetCloudWorkspaceManaged sets CloudWorkspaceManaged field to given value.

### HasCloudWorkspaceManaged

`func (o *ApplicationResponseModel) HasCloudWorkspaceManaged() bool`

HasCloudWorkspaceManaged returns a boolean if a field has been set.

### GetNumAssociatedDeliveryGroups

`func (o *ApplicationResponseModel) GetNumAssociatedDeliveryGroups() int32`

GetNumAssociatedDeliveryGroups returns the NumAssociatedDeliveryGroups field if non-nil, zero value otherwise.

### GetNumAssociatedDeliveryGroupsOk

`func (o *ApplicationResponseModel) GetNumAssociatedDeliveryGroupsOk() (*int32, bool)`

GetNumAssociatedDeliveryGroupsOk returns a tuple with the NumAssociatedDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssociatedDeliveryGroups

`func (o *ApplicationResponseModel) SetNumAssociatedDeliveryGroups(v int32)`

SetNumAssociatedDeliveryGroups sets NumAssociatedDeliveryGroups field to given value.

### HasNumAssociatedDeliveryGroups

`func (o *ApplicationResponseModel) HasNumAssociatedDeliveryGroups() bool`

HasNumAssociatedDeliveryGroups returns a boolean if a field has been set.

### GetNumAssociatedApplicationGroups

`func (o *ApplicationResponseModel) GetNumAssociatedApplicationGroups() int32`

GetNumAssociatedApplicationGroups returns the NumAssociatedApplicationGroups field if non-nil, zero value otherwise.

### GetNumAssociatedApplicationGroupsOk

`func (o *ApplicationResponseModel) GetNumAssociatedApplicationGroupsOk() (*int32, bool)`

GetNumAssociatedApplicationGroupsOk returns a tuple with the NumAssociatedApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssociatedApplicationGroups

`func (o *ApplicationResponseModel) SetNumAssociatedApplicationGroups(v int32)`

SetNumAssociatedApplicationGroups sets NumAssociatedApplicationGroups field to given value.

### HasNumAssociatedApplicationGroups

`func (o *ApplicationResponseModel) HasNumAssociatedApplicationGroups() bool`

HasNumAssociatedApplicationGroups returns a boolean if a field has been set.

### GetAssociatedDeliveryGroupUuids

`func (o *ApplicationResponseModel) GetAssociatedDeliveryGroupUuids() []string`

GetAssociatedDeliveryGroupUuids returns the AssociatedDeliveryGroupUuids field if non-nil, zero value otherwise.

### GetAssociatedDeliveryGroupUuidsOk

`func (o *ApplicationResponseModel) GetAssociatedDeliveryGroupUuidsOk() (*[]string, bool)`

GetAssociatedDeliveryGroupUuidsOk returns a tuple with the AssociatedDeliveryGroupUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDeliveryGroupUuids

`func (o *ApplicationResponseModel) SetAssociatedDeliveryGroupUuids(v []string)`

SetAssociatedDeliveryGroupUuids sets AssociatedDeliveryGroupUuids field to given value.

### HasAssociatedDeliveryGroupUuids

`func (o *ApplicationResponseModel) HasAssociatedDeliveryGroupUuids() bool`

HasAssociatedDeliveryGroupUuids returns a boolean if a field has been set.

### SetAssociatedDeliveryGroupUuidsNil

`func (o *ApplicationResponseModel) SetAssociatedDeliveryGroupUuidsNil(b bool)`

 SetAssociatedDeliveryGroupUuidsNil sets the value for AssociatedDeliveryGroupUuids to be an explicit nil

### UnsetAssociatedDeliveryGroupUuids
`func (o *ApplicationResponseModel) UnsetAssociatedDeliveryGroupUuids()`

UnsetAssociatedDeliveryGroupUuids ensures that no value is present for AssociatedDeliveryGroupUuids, not even an explicit nil
### GetAssociatedApplicationGroupUuids

`func (o *ApplicationResponseModel) GetAssociatedApplicationGroupUuids() []string`

GetAssociatedApplicationGroupUuids returns the AssociatedApplicationGroupUuids field if non-nil, zero value otherwise.

### GetAssociatedApplicationGroupUuidsOk

`func (o *ApplicationResponseModel) GetAssociatedApplicationGroupUuidsOk() (*[]string, bool)`

GetAssociatedApplicationGroupUuidsOk returns a tuple with the AssociatedApplicationGroupUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedApplicationGroupUuids

`func (o *ApplicationResponseModel) SetAssociatedApplicationGroupUuids(v []string)`

SetAssociatedApplicationGroupUuids sets AssociatedApplicationGroupUuids field to given value.

### HasAssociatedApplicationGroupUuids

`func (o *ApplicationResponseModel) HasAssociatedApplicationGroupUuids() bool`

HasAssociatedApplicationGroupUuids returns a boolean if a field has been set.

### SetAssociatedApplicationGroupUuidsNil

`func (o *ApplicationResponseModel) SetAssociatedApplicationGroupUuidsNil(b bool)`

 SetAssociatedApplicationGroupUuidsNil sets the value for AssociatedApplicationGroupUuids to be an explicit nil

### UnsetAssociatedApplicationGroupUuids
`func (o *ApplicationResponseModel) UnsetAssociatedApplicationGroupUuids()`

UnsetAssociatedApplicationGroupUuids ensures that no value is present for AssociatedApplicationGroupUuids, not even an explicit nil
### GetZoneId

`func (o *ApplicationResponseModel) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ApplicationResponseModel) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ApplicationResponseModel) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ApplicationResponseModel) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *ApplicationResponseModel) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *ApplicationResponseModel) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


