# EditApplicationGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminFolder** | Pointer to **NullableString** | The folder in which the application group resides. If not specified, the value will not be changed. May be specified as either the folder Id or Path. If specified as a path, and the path does not exist, it will be automatically created. | [optional] 
**Name** | Pointer to **NullableString** | Name of the application group.   If not specified, will not be changed. | [optional] 
**Description** | Pointer to **NullableString** | Description of the application group.   If not specified, will not be changed.  To remove a previously set description, use the empty string (\&quot;\&quot;). | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the application group&#39;s applications can be launched by end users.   If not specified, will not be changed. | [optional] 
**RestrictToTag** | Pointer to **NullableString** | Optional tag that may be used further to restrict which machines may be used for launching the application group&#39;s applications.   A machine may be used by an application group if either the application group has no tag restriction or the application group does have a tag restriction and the machine is tagged with the same tag.   If not specified, will not be changed.  To remove a previously set tag restriction, use the empty string (\&quot;\&quot;). | [optional] 
**Scopes** | Pointer to **[]string** | Administrative scopes which the application group should be a part of.   If not specified, will not be changed.   The \&quot;All\&quot; scope, and any tenant scopes, are implicit.  To remove from all non-implicit scopes, specify an empty array ([]).   Specifying tenant scopes is equivalent to specifying the  property and is subject to the same constraints. | [optional] 
**Tenants** | Pointer to **[]string** | Tenants to associate with the application group. | [optional] 
**SessionSharingEnabled** | Pointer to **NullableBool** | Whether the application group&#39;s applications can share sessions with applications that are not a member of this application group.   If not specified, will not be changed. | [optional] 
**IncludedUserFilterEnabled** | Pointer to **NullableBool** | Specifies whether the  filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery groups&#39; access polic(ies) is implicitly granted access to the applications within the application group.   If not specified, will not be changed. | [optional] 
**IncludedUsers** | Pointer to **[]string** | Specifies the included users filter of the application group; that is, the users and groups who are explicitly granted access to the applications in the application group.   If not specified, will not be changed.   If specified, all desired users must be listed.  To remove all users, specify an empty array ([]). | [optional] 
**DeliveryGroups** | Pointer to [**[]PriorityRefRequestModel**](PriorityRefRequestModel.md) | List of delivery groups to associate with the application group.   If not specified, will not be changed.   If specified, all desired delivery groups must be listed.  The application group must be associated with at least one delivery group. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of the application group. When set the property value equal to null/empty means to remove this property. | [optional] 

## Methods

### NewEditApplicationGroupRequestModel

`func NewEditApplicationGroupRequestModel() *EditApplicationGroupRequestModel`

NewEditApplicationGroupRequestModel instantiates a new EditApplicationGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditApplicationGroupRequestModelWithDefaults

`func NewEditApplicationGroupRequestModelWithDefaults() *EditApplicationGroupRequestModel`

NewEditApplicationGroupRequestModelWithDefaults instantiates a new EditApplicationGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminFolder

`func (o *EditApplicationGroupRequestModel) GetAdminFolder() string`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *EditApplicationGroupRequestModel) GetAdminFolderOk() (*string, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *EditApplicationGroupRequestModel) SetAdminFolder(v string)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *EditApplicationGroupRequestModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### SetAdminFolderNil

`func (o *EditApplicationGroupRequestModel) SetAdminFolderNil(b bool)`

 SetAdminFolderNil sets the value for AdminFolder to be an explicit nil

### UnsetAdminFolder
`func (o *EditApplicationGroupRequestModel) UnsetAdminFolder()`

UnsetAdminFolder ensures that no value is present for AdminFolder, not even an explicit nil
### GetName

`func (o *EditApplicationGroupRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditApplicationGroupRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditApplicationGroupRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditApplicationGroupRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditApplicationGroupRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditApplicationGroupRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *EditApplicationGroupRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditApplicationGroupRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditApplicationGroupRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditApplicationGroupRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EditApplicationGroupRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditApplicationGroupRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *EditApplicationGroupRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EditApplicationGroupRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EditApplicationGroupRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EditApplicationGroupRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *EditApplicationGroupRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *EditApplicationGroupRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetRestrictToTag

`func (o *EditApplicationGroupRequestModel) GetRestrictToTag() string`

GetRestrictToTag returns the RestrictToTag field if non-nil, zero value otherwise.

### GetRestrictToTagOk

`func (o *EditApplicationGroupRequestModel) GetRestrictToTagOk() (*string, bool)`

GetRestrictToTagOk returns a tuple with the RestrictToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToTag

`func (o *EditApplicationGroupRequestModel) SetRestrictToTag(v string)`

SetRestrictToTag sets RestrictToTag field to given value.

### HasRestrictToTag

`func (o *EditApplicationGroupRequestModel) HasRestrictToTag() bool`

HasRestrictToTag returns a boolean if a field has been set.

### SetRestrictToTagNil

`func (o *EditApplicationGroupRequestModel) SetRestrictToTagNil(b bool)`

 SetRestrictToTagNil sets the value for RestrictToTag to be an explicit nil

### UnsetRestrictToTag
`func (o *EditApplicationGroupRequestModel) UnsetRestrictToTag()`

UnsetRestrictToTag ensures that no value is present for RestrictToTag, not even an explicit nil
### GetScopes

`func (o *EditApplicationGroupRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *EditApplicationGroupRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *EditApplicationGroupRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *EditApplicationGroupRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *EditApplicationGroupRequestModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *EditApplicationGroupRequestModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *EditApplicationGroupRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *EditApplicationGroupRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *EditApplicationGroupRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *EditApplicationGroupRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *EditApplicationGroupRequestModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *EditApplicationGroupRequestModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetSessionSharingEnabled

`func (o *EditApplicationGroupRequestModel) GetSessionSharingEnabled() bool`

GetSessionSharingEnabled returns the SessionSharingEnabled field if non-nil, zero value otherwise.

### GetSessionSharingEnabledOk

`func (o *EditApplicationGroupRequestModel) GetSessionSharingEnabledOk() (*bool, bool)`

GetSessionSharingEnabledOk returns a tuple with the SessionSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSharingEnabled

`func (o *EditApplicationGroupRequestModel) SetSessionSharingEnabled(v bool)`

SetSessionSharingEnabled sets SessionSharingEnabled field to given value.

### HasSessionSharingEnabled

`func (o *EditApplicationGroupRequestModel) HasSessionSharingEnabled() bool`

HasSessionSharingEnabled returns a boolean if a field has been set.

### SetSessionSharingEnabledNil

`func (o *EditApplicationGroupRequestModel) SetSessionSharingEnabledNil(b bool)`

 SetSessionSharingEnabledNil sets the value for SessionSharingEnabled to be an explicit nil

### UnsetSessionSharingEnabled
`func (o *EditApplicationGroupRequestModel) UnsetSessionSharingEnabled()`

UnsetSessionSharingEnabled ensures that no value is present for SessionSharingEnabled, not even an explicit nil
### GetIncludedUserFilterEnabled

`func (o *EditApplicationGroupRequestModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *EditApplicationGroupRequestModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *EditApplicationGroupRequestModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *EditApplicationGroupRequestModel) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### SetIncludedUserFilterEnabledNil

`func (o *EditApplicationGroupRequestModel) SetIncludedUserFilterEnabledNil(b bool)`

 SetIncludedUserFilterEnabledNil sets the value for IncludedUserFilterEnabled to be an explicit nil

### UnsetIncludedUserFilterEnabled
`func (o *EditApplicationGroupRequestModel) UnsetIncludedUserFilterEnabled()`

UnsetIncludedUserFilterEnabled ensures that no value is present for IncludedUserFilterEnabled, not even an explicit nil
### GetIncludedUsers

`func (o *EditApplicationGroupRequestModel) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *EditApplicationGroupRequestModel) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *EditApplicationGroupRequestModel) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *EditApplicationGroupRequestModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *EditApplicationGroupRequestModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *EditApplicationGroupRequestModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil
### GetDeliveryGroups

`func (o *EditApplicationGroupRequestModel) GetDeliveryGroups() []PriorityRefRequestModel`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *EditApplicationGroupRequestModel) GetDeliveryGroupsOk() (*[]PriorityRefRequestModel, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *EditApplicationGroupRequestModel) SetDeliveryGroups(v []PriorityRefRequestModel)`

SetDeliveryGroups sets DeliveryGroups field to given value.

### HasDeliveryGroups

`func (o *EditApplicationGroupRequestModel) HasDeliveryGroups() bool`

HasDeliveryGroups returns a boolean if a field has been set.

### SetDeliveryGroupsNil

`func (o *EditApplicationGroupRequestModel) SetDeliveryGroupsNil(b bool)`

 SetDeliveryGroupsNil sets the value for DeliveryGroups to be an explicit nil

### UnsetDeliveryGroups
`func (o *EditApplicationGroupRequestModel) UnsetDeliveryGroups()`

UnsetDeliveryGroups ensures that no value is present for DeliveryGroups, not even an explicit nil
### GetMetadata

`func (o *EditApplicationGroupRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EditApplicationGroupRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EditApplicationGroupRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EditApplicationGroupRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EditApplicationGroupRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EditApplicationGroupRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


