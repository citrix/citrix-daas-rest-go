# MachineApplicationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The id of the application. Used to be: Uuid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | [optional] 
**Uid** | Pointer to **NullableInt32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED. Use Id. | [optional] 
**ApplicationFolder** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**ApplicationType** | Pointer to [**ApplicationType**](ApplicationType.md) |  | [optional] 
**PackagedApplicationType** | Pointer to [**PackagedApplicationType**](PackagedApplicationType.md) |  | [optional] 
**PackagedApplicationVisibility** | Pointer to [**PackagedApplicationVisibility**](PackagedApplicationVisibility.md) |  | [optional] 
**ClientFolder** | Pointer to **NullableString** | The folder that the application belongs to as the user sees it. | [optional] 
**ContainerScopes** | Pointer to [**[]ContainerScopeResponseModel**](ContainerScopeResponseModel.md) | Delegated admin scopes in which the containers of the application reside. | [optional] 
**Description** | Pointer to **NullableString** | The description of the application. | [optional] 
**DoNotEnumerate** | Pointer to **bool** | Indicates whether this application is enumerable | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this application can be launched. | [optional] 
**IconId** | Pointer to **NullableString** | The id of the icon used for the application. Used to be: IconUid (and it was not globally unique) Needs to be globally unique Might be constructed from site ID + internal Uid | [optional] 
**InstalledAppProperties** | Pointer to [**InstalledAppResponseModel**](InstalledAppResponseModel.md) |  | [optional] 
**PackagedAppProperties** | Pointer to [**AppVAppResponseModel**](AppVAppResponseModel.md) |  | [optional] 
**AppVAppProperties** | Pointer to [**AppVAppResponseModel**](AppVAppResponseModel.md) |  | [optional] 
**ContentLocation** | Pointer to **NullableString** | Location of published content. | [optional] 
**Name** | Pointer to **NullableString** | Name of the application.  Only seen by administrators. | [optional] 
**PublishedName** | Pointer to **NullableString** | The name seen by end users who have access to the application. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of application. | [optional] 
**Visible** | Pointer to **bool** | Indicates whether this application is visible to users. | [optional] 
**SharingKind** | Pointer to [**SharingKind**](SharingKind.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with this application. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the application is assigned to.  If &#x60;null&#x60;, the application is not assigned to any tenants, and may be used by any tenant. | [optional] 
**CloudWorkspaceManaged** | Pointer to **bool** | Indicates whether users are managed in the Citrix Cloud Library, or within Studio. | [optional] 
**NumAssociatedDeliveryGroups** | Pointer to **int32** | Number of delivery groups that the application is associated with. | [optional] 
**NumAssociatedApplicationGroups** | Pointer to **int32** | Number of application groups that the application is associated with. | [optional] 
**AssociatedDeliveryGroupUuids** | Pointer to **[]string** | Delivery group Uuids that the application is associated with. | [optional] 
**AssociatedApplicationGroupUuids** | Pointer to **[]string** | Application group Uuids that the application is associated with. | [optional] 
**ZoneId** | Pointer to **NullableString** | Application Zone info. | [optional] 
**Published** | Pointer to **bool** | Indicates whether the application is published on the machine. | [optional] 
**InUse** | Pointer to **bool** | Indicates whether the application is in use on the machine. | [optional] 

## Methods

### NewMachineApplicationResponseModel

`func NewMachineApplicationResponseModel() *MachineApplicationResponseModel`

NewMachineApplicationResponseModel instantiates a new MachineApplicationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineApplicationResponseModelWithDefaults

`func NewMachineApplicationResponseModelWithDefaults() *MachineApplicationResponseModel`

NewMachineApplicationResponseModelWithDefaults instantiates a new MachineApplicationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MachineApplicationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineApplicationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineApplicationResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MachineApplicationResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MachineApplicationResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MachineApplicationResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *MachineApplicationResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MachineApplicationResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MachineApplicationResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MachineApplicationResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *MachineApplicationResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *MachineApplicationResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetApplicationFolder

`func (o *MachineApplicationResponseModel) GetApplicationFolder() RefResponseModel`

GetApplicationFolder returns the ApplicationFolder field if non-nil, zero value otherwise.

### GetApplicationFolderOk

`func (o *MachineApplicationResponseModel) GetApplicationFolderOk() (*RefResponseModel, bool)`

GetApplicationFolderOk returns a tuple with the ApplicationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFolder

`func (o *MachineApplicationResponseModel) SetApplicationFolder(v RefResponseModel)`

SetApplicationFolder sets ApplicationFolder field to given value.

### HasApplicationFolder

`func (o *MachineApplicationResponseModel) HasApplicationFolder() bool`

HasApplicationFolder returns a boolean if a field has been set.

### GetApplicationType

`func (o *MachineApplicationResponseModel) GetApplicationType() ApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *MachineApplicationResponseModel) GetApplicationTypeOk() (*ApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *MachineApplicationResponseModel) SetApplicationType(v ApplicationType)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *MachineApplicationResponseModel) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetPackagedApplicationType

`func (o *MachineApplicationResponseModel) GetPackagedApplicationType() PackagedApplicationType`

GetPackagedApplicationType returns the PackagedApplicationType field if non-nil, zero value otherwise.

### GetPackagedApplicationTypeOk

`func (o *MachineApplicationResponseModel) GetPackagedApplicationTypeOk() (*PackagedApplicationType, bool)`

GetPackagedApplicationTypeOk returns a tuple with the PackagedApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagedApplicationType

`func (o *MachineApplicationResponseModel) SetPackagedApplicationType(v PackagedApplicationType)`

SetPackagedApplicationType sets PackagedApplicationType field to given value.

### HasPackagedApplicationType

`func (o *MachineApplicationResponseModel) HasPackagedApplicationType() bool`

HasPackagedApplicationType returns a boolean if a field has been set.

### GetPackagedApplicationVisibility

`func (o *MachineApplicationResponseModel) GetPackagedApplicationVisibility() PackagedApplicationVisibility`

GetPackagedApplicationVisibility returns the PackagedApplicationVisibility field if non-nil, zero value otherwise.

### GetPackagedApplicationVisibilityOk

`func (o *MachineApplicationResponseModel) GetPackagedApplicationVisibilityOk() (*PackagedApplicationVisibility, bool)`

GetPackagedApplicationVisibilityOk returns a tuple with the PackagedApplicationVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagedApplicationVisibility

`func (o *MachineApplicationResponseModel) SetPackagedApplicationVisibility(v PackagedApplicationVisibility)`

SetPackagedApplicationVisibility sets PackagedApplicationVisibility field to given value.

### HasPackagedApplicationVisibility

`func (o *MachineApplicationResponseModel) HasPackagedApplicationVisibility() bool`

HasPackagedApplicationVisibility returns a boolean if a field has been set.

### GetClientFolder

`func (o *MachineApplicationResponseModel) GetClientFolder() string`

GetClientFolder returns the ClientFolder field if non-nil, zero value otherwise.

### GetClientFolderOk

`func (o *MachineApplicationResponseModel) GetClientFolderOk() (*string, bool)`

GetClientFolderOk returns a tuple with the ClientFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFolder

`func (o *MachineApplicationResponseModel) SetClientFolder(v string)`

SetClientFolder sets ClientFolder field to given value.

### HasClientFolder

`func (o *MachineApplicationResponseModel) HasClientFolder() bool`

HasClientFolder returns a boolean if a field has been set.

### SetClientFolderNil

`func (o *MachineApplicationResponseModel) SetClientFolderNil(b bool)`

 SetClientFolderNil sets the value for ClientFolder to be an explicit nil

### UnsetClientFolder
`func (o *MachineApplicationResponseModel) UnsetClientFolder()`

UnsetClientFolder ensures that no value is present for ClientFolder, not even an explicit nil
### GetContainerScopes

`func (o *MachineApplicationResponseModel) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *MachineApplicationResponseModel) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *MachineApplicationResponseModel) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.

### HasContainerScopes

`func (o *MachineApplicationResponseModel) HasContainerScopes() bool`

HasContainerScopes returns a boolean if a field has been set.

### SetContainerScopesNil

`func (o *MachineApplicationResponseModel) SetContainerScopesNil(b bool)`

 SetContainerScopesNil sets the value for ContainerScopes to be an explicit nil

### UnsetContainerScopes
`func (o *MachineApplicationResponseModel) UnsetContainerScopes()`

UnsetContainerScopes ensures that no value is present for ContainerScopes, not even an explicit nil
### GetDescription

`func (o *MachineApplicationResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MachineApplicationResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MachineApplicationResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MachineApplicationResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MachineApplicationResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MachineApplicationResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDoNotEnumerate

`func (o *MachineApplicationResponseModel) GetDoNotEnumerate() bool`

GetDoNotEnumerate returns the DoNotEnumerate field if non-nil, zero value otherwise.

### GetDoNotEnumerateOk

`func (o *MachineApplicationResponseModel) GetDoNotEnumerateOk() (*bool, bool)`

GetDoNotEnumerateOk returns a tuple with the DoNotEnumerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotEnumerate

`func (o *MachineApplicationResponseModel) SetDoNotEnumerate(v bool)`

SetDoNotEnumerate sets DoNotEnumerate field to given value.

### HasDoNotEnumerate

`func (o *MachineApplicationResponseModel) HasDoNotEnumerate() bool`

HasDoNotEnumerate returns a boolean if a field has been set.

### GetEnabled

`func (o *MachineApplicationResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MachineApplicationResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MachineApplicationResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MachineApplicationResponseModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIconId

`func (o *MachineApplicationResponseModel) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *MachineApplicationResponseModel) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *MachineApplicationResponseModel) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *MachineApplicationResponseModel) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### SetIconIdNil

`func (o *MachineApplicationResponseModel) SetIconIdNil(b bool)`

 SetIconIdNil sets the value for IconId to be an explicit nil

### UnsetIconId
`func (o *MachineApplicationResponseModel) UnsetIconId()`

UnsetIconId ensures that no value is present for IconId, not even an explicit nil
### GetInstalledAppProperties

`func (o *MachineApplicationResponseModel) GetInstalledAppProperties() InstalledAppResponseModel`

GetInstalledAppProperties returns the InstalledAppProperties field if non-nil, zero value otherwise.

### GetInstalledAppPropertiesOk

`func (o *MachineApplicationResponseModel) GetInstalledAppPropertiesOk() (*InstalledAppResponseModel, bool)`

GetInstalledAppPropertiesOk returns a tuple with the InstalledAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAppProperties

`func (o *MachineApplicationResponseModel) SetInstalledAppProperties(v InstalledAppResponseModel)`

SetInstalledAppProperties sets InstalledAppProperties field to given value.

### HasInstalledAppProperties

`func (o *MachineApplicationResponseModel) HasInstalledAppProperties() bool`

HasInstalledAppProperties returns a boolean if a field has been set.

### GetPackagedAppProperties

`func (o *MachineApplicationResponseModel) GetPackagedAppProperties() AppVAppResponseModel`

GetPackagedAppProperties returns the PackagedAppProperties field if non-nil, zero value otherwise.

### GetPackagedAppPropertiesOk

`func (o *MachineApplicationResponseModel) GetPackagedAppPropertiesOk() (*AppVAppResponseModel, bool)`

GetPackagedAppPropertiesOk returns a tuple with the PackagedAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagedAppProperties

`func (o *MachineApplicationResponseModel) SetPackagedAppProperties(v AppVAppResponseModel)`

SetPackagedAppProperties sets PackagedAppProperties field to given value.

### HasPackagedAppProperties

`func (o *MachineApplicationResponseModel) HasPackagedAppProperties() bool`

HasPackagedAppProperties returns a boolean if a field has been set.

### GetAppVAppProperties

`func (o *MachineApplicationResponseModel) GetAppVAppProperties() AppVAppResponseModel`

GetAppVAppProperties returns the AppVAppProperties field if non-nil, zero value otherwise.

### GetAppVAppPropertiesOk

`func (o *MachineApplicationResponseModel) GetAppVAppPropertiesOk() (*AppVAppResponseModel, bool)`

GetAppVAppPropertiesOk returns a tuple with the AppVAppProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVAppProperties

`func (o *MachineApplicationResponseModel) SetAppVAppProperties(v AppVAppResponseModel)`

SetAppVAppProperties sets AppVAppProperties field to given value.

### HasAppVAppProperties

`func (o *MachineApplicationResponseModel) HasAppVAppProperties() bool`

HasAppVAppProperties returns a boolean if a field has been set.

### GetContentLocation

`func (o *MachineApplicationResponseModel) GetContentLocation() string`

GetContentLocation returns the ContentLocation field if non-nil, zero value otherwise.

### GetContentLocationOk

`func (o *MachineApplicationResponseModel) GetContentLocationOk() (*string, bool)`

GetContentLocationOk returns a tuple with the ContentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLocation

`func (o *MachineApplicationResponseModel) SetContentLocation(v string)`

SetContentLocation sets ContentLocation field to given value.

### HasContentLocation

`func (o *MachineApplicationResponseModel) HasContentLocation() bool`

HasContentLocation returns a boolean if a field has been set.

### SetContentLocationNil

`func (o *MachineApplicationResponseModel) SetContentLocationNil(b bool)`

 SetContentLocationNil sets the value for ContentLocation to be an explicit nil

### UnsetContentLocation
`func (o *MachineApplicationResponseModel) UnsetContentLocation()`

UnsetContentLocation ensures that no value is present for ContentLocation, not even an explicit nil
### GetName

`func (o *MachineApplicationResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineApplicationResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineApplicationResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineApplicationResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MachineApplicationResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MachineApplicationResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPublishedName

`func (o *MachineApplicationResponseModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *MachineApplicationResponseModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *MachineApplicationResponseModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *MachineApplicationResponseModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *MachineApplicationResponseModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *MachineApplicationResponseModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetMetadata

`func (o *MachineApplicationResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MachineApplicationResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MachineApplicationResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MachineApplicationResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *MachineApplicationResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *MachineApplicationResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetVisible

`func (o *MachineApplicationResponseModel) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *MachineApplicationResponseModel) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *MachineApplicationResponseModel) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *MachineApplicationResponseModel) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetSharingKind

`func (o *MachineApplicationResponseModel) GetSharingKind() SharingKind`

GetSharingKind returns the SharingKind field if non-nil, zero value otherwise.

### GetSharingKindOk

`func (o *MachineApplicationResponseModel) GetSharingKindOk() (*SharingKind, bool)`

GetSharingKindOk returns a tuple with the SharingKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKind

`func (o *MachineApplicationResponseModel) SetSharingKind(v SharingKind)`

SetSharingKind sets SharingKind field to given value.

### HasSharingKind

`func (o *MachineApplicationResponseModel) HasSharingKind() bool`

HasSharingKind returns a boolean if a field has been set.

### GetTags

`func (o *MachineApplicationResponseModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MachineApplicationResponseModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MachineApplicationResponseModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MachineApplicationResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MachineApplicationResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MachineApplicationResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTenants

`func (o *MachineApplicationResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *MachineApplicationResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *MachineApplicationResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *MachineApplicationResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *MachineApplicationResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *MachineApplicationResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetCloudWorkspaceManaged

`func (o *MachineApplicationResponseModel) GetCloudWorkspaceManaged() bool`

GetCloudWorkspaceManaged returns the CloudWorkspaceManaged field if non-nil, zero value otherwise.

### GetCloudWorkspaceManagedOk

`func (o *MachineApplicationResponseModel) GetCloudWorkspaceManagedOk() (*bool, bool)`

GetCloudWorkspaceManagedOk returns a tuple with the CloudWorkspaceManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudWorkspaceManaged

`func (o *MachineApplicationResponseModel) SetCloudWorkspaceManaged(v bool)`

SetCloudWorkspaceManaged sets CloudWorkspaceManaged field to given value.

### HasCloudWorkspaceManaged

`func (o *MachineApplicationResponseModel) HasCloudWorkspaceManaged() bool`

HasCloudWorkspaceManaged returns a boolean if a field has been set.

### GetNumAssociatedDeliveryGroups

`func (o *MachineApplicationResponseModel) GetNumAssociatedDeliveryGroups() int32`

GetNumAssociatedDeliveryGroups returns the NumAssociatedDeliveryGroups field if non-nil, zero value otherwise.

### GetNumAssociatedDeliveryGroupsOk

`func (o *MachineApplicationResponseModel) GetNumAssociatedDeliveryGroupsOk() (*int32, bool)`

GetNumAssociatedDeliveryGroupsOk returns a tuple with the NumAssociatedDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssociatedDeliveryGroups

`func (o *MachineApplicationResponseModel) SetNumAssociatedDeliveryGroups(v int32)`

SetNumAssociatedDeliveryGroups sets NumAssociatedDeliveryGroups field to given value.

### HasNumAssociatedDeliveryGroups

`func (o *MachineApplicationResponseModel) HasNumAssociatedDeliveryGroups() bool`

HasNumAssociatedDeliveryGroups returns a boolean if a field has been set.

### GetNumAssociatedApplicationGroups

`func (o *MachineApplicationResponseModel) GetNumAssociatedApplicationGroups() int32`

GetNumAssociatedApplicationGroups returns the NumAssociatedApplicationGroups field if non-nil, zero value otherwise.

### GetNumAssociatedApplicationGroupsOk

`func (o *MachineApplicationResponseModel) GetNumAssociatedApplicationGroupsOk() (*int32, bool)`

GetNumAssociatedApplicationGroupsOk returns a tuple with the NumAssociatedApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssociatedApplicationGroups

`func (o *MachineApplicationResponseModel) SetNumAssociatedApplicationGroups(v int32)`

SetNumAssociatedApplicationGroups sets NumAssociatedApplicationGroups field to given value.

### HasNumAssociatedApplicationGroups

`func (o *MachineApplicationResponseModel) HasNumAssociatedApplicationGroups() bool`

HasNumAssociatedApplicationGroups returns a boolean if a field has been set.

### GetAssociatedDeliveryGroupUuids

`func (o *MachineApplicationResponseModel) GetAssociatedDeliveryGroupUuids() []string`

GetAssociatedDeliveryGroupUuids returns the AssociatedDeliveryGroupUuids field if non-nil, zero value otherwise.

### GetAssociatedDeliveryGroupUuidsOk

`func (o *MachineApplicationResponseModel) GetAssociatedDeliveryGroupUuidsOk() (*[]string, bool)`

GetAssociatedDeliveryGroupUuidsOk returns a tuple with the AssociatedDeliveryGroupUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDeliveryGroupUuids

`func (o *MachineApplicationResponseModel) SetAssociatedDeliveryGroupUuids(v []string)`

SetAssociatedDeliveryGroupUuids sets AssociatedDeliveryGroupUuids field to given value.

### HasAssociatedDeliveryGroupUuids

`func (o *MachineApplicationResponseModel) HasAssociatedDeliveryGroupUuids() bool`

HasAssociatedDeliveryGroupUuids returns a boolean if a field has been set.

### SetAssociatedDeliveryGroupUuidsNil

`func (o *MachineApplicationResponseModel) SetAssociatedDeliveryGroupUuidsNil(b bool)`

 SetAssociatedDeliveryGroupUuidsNil sets the value for AssociatedDeliveryGroupUuids to be an explicit nil

### UnsetAssociatedDeliveryGroupUuids
`func (o *MachineApplicationResponseModel) UnsetAssociatedDeliveryGroupUuids()`

UnsetAssociatedDeliveryGroupUuids ensures that no value is present for AssociatedDeliveryGroupUuids, not even an explicit nil
### GetAssociatedApplicationGroupUuids

`func (o *MachineApplicationResponseModel) GetAssociatedApplicationGroupUuids() []string`

GetAssociatedApplicationGroupUuids returns the AssociatedApplicationGroupUuids field if non-nil, zero value otherwise.

### GetAssociatedApplicationGroupUuidsOk

`func (o *MachineApplicationResponseModel) GetAssociatedApplicationGroupUuidsOk() (*[]string, bool)`

GetAssociatedApplicationGroupUuidsOk returns a tuple with the AssociatedApplicationGroupUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedApplicationGroupUuids

`func (o *MachineApplicationResponseModel) SetAssociatedApplicationGroupUuids(v []string)`

SetAssociatedApplicationGroupUuids sets AssociatedApplicationGroupUuids field to given value.

### HasAssociatedApplicationGroupUuids

`func (o *MachineApplicationResponseModel) HasAssociatedApplicationGroupUuids() bool`

HasAssociatedApplicationGroupUuids returns a boolean if a field has been set.

### SetAssociatedApplicationGroupUuidsNil

`func (o *MachineApplicationResponseModel) SetAssociatedApplicationGroupUuidsNil(b bool)`

 SetAssociatedApplicationGroupUuidsNil sets the value for AssociatedApplicationGroupUuids to be an explicit nil

### UnsetAssociatedApplicationGroupUuids
`func (o *MachineApplicationResponseModel) UnsetAssociatedApplicationGroupUuids()`

UnsetAssociatedApplicationGroupUuids ensures that no value is present for AssociatedApplicationGroupUuids, not even an explicit nil
### GetZoneId

`func (o *MachineApplicationResponseModel) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *MachineApplicationResponseModel) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *MachineApplicationResponseModel) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *MachineApplicationResponseModel) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *MachineApplicationResponseModel) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *MachineApplicationResponseModel) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetPublished

`func (o *MachineApplicationResponseModel) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *MachineApplicationResponseModel) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *MachineApplicationResponseModel) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *MachineApplicationResponseModel) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetInUse

`func (o *MachineApplicationResponseModel) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *MachineApplicationResponseModel) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *MachineApplicationResponseModel) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *MachineApplicationResponseModel) HasInUse() bool`

HasInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


