# CreateApplicationGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminFolder** | Pointer to **NullableString** | The admin folder in which the application group should be created. | [optional] 
**Name** | **string** | Name of the application group to create. | 
**Description** | Pointer to **NullableString** | Description of the application group. | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the application group&#39;s applications can be launched by end users. | [optional] [default to true]
**RestrictToTag** | Pointer to **NullableString** | Optional tag that may be used further to restrict which machines may be used for launching the application group&#39;s applications.   A machine may be used by an application group if either the application group has no tag restriction or the application group does have a tag restriction and the machine is tagged with the same tag. | [optional] 
**Scopes** | Pointer to **[]string** | Administrative scopes which the newly created application group will be a part of. | [optional] 
**Tenants** | Pointer to **[]string** | Tenant(s) to associate the application group with. | [optional] 
**SessionSharingEnabled** | Pointer to **NullableBool** | Whether the application group&#39;s applications can share sessions with applications that are not a member of this application group. | [optional] [default to true]
**IncludedUserFilterEnabled** | Pointer to **NullableBool** | Specifies whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery groups&#39; access polic(ies) is implicitly granted access to the applications within the application group. | [optional] [default to false]
**IncludedUsers** | Pointer to **[]string** | Specifies the included users filter of the application group; that is, the users and groups who are explicitly granted access to the applications in the application group. | [optional] 
**DeliveryGroups** | Pointer to [**[]PriorityRefRequestModel**](PriorityRefRequestModel.md) | List of delivery groups to associate with the application group. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of the new application group. | [optional] 

## Methods

### NewCreateApplicationGroupRequestModel

`func NewCreateApplicationGroupRequestModel(name string, ) *CreateApplicationGroupRequestModel`

NewCreateApplicationGroupRequestModel instantiates a new CreateApplicationGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationGroupRequestModelWithDefaults

`func NewCreateApplicationGroupRequestModelWithDefaults() *CreateApplicationGroupRequestModel`

NewCreateApplicationGroupRequestModelWithDefaults instantiates a new CreateApplicationGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminFolder

`func (o *CreateApplicationGroupRequestModel) GetAdminFolder() string`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *CreateApplicationGroupRequestModel) GetAdminFolderOk() (*string, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *CreateApplicationGroupRequestModel) SetAdminFolder(v string)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *CreateApplicationGroupRequestModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### SetAdminFolderNil

`func (o *CreateApplicationGroupRequestModel) SetAdminFolderNil(b bool)`

 SetAdminFolderNil sets the value for AdminFolder to be an explicit nil

### UnsetAdminFolder
`func (o *CreateApplicationGroupRequestModel) UnsetAdminFolder()`

UnsetAdminFolder ensures that no value is present for AdminFolder, not even an explicit nil
### GetName

`func (o *CreateApplicationGroupRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApplicationGroupRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApplicationGroupRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateApplicationGroupRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateApplicationGroupRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateApplicationGroupRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateApplicationGroupRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateApplicationGroupRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateApplicationGroupRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *CreateApplicationGroupRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateApplicationGroupRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateApplicationGroupRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateApplicationGroupRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *CreateApplicationGroupRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *CreateApplicationGroupRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetRestrictToTag

`func (o *CreateApplicationGroupRequestModel) GetRestrictToTag() string`

GetRestrictToTag returns the RestrictToTag field if non-nil, zero value otherwise.

### GetRestrictToTagOk

`func (o *CreateApplicationGroupRequestModel) GetRestrictToTagOk() (*string, bool)`

GetRestrictToTagOk returns a tuple with the RestrictToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToTag

`func (o *CreateApplicationGroupRequestModel) SetRestrictToTag(v string)`

SetRestrictToTag sets RestrictToTag field to given value.

### HasRestrictToTag

`func (o *CreateApplicationGroupRequestModel) HasRestrictToTag() bool`

HasRestrictToTag returns a boolean if a field has been set.

### SetRestrictToTagNil

`func (o *CreateApplicationGroupRequestModel) SetRestrictToTagNil(b bool)`

 SetRestrictToTagNil sets the value for RestrictToTag to be an explicit nil

### UnsetRestrictToTag
`func (o *CreateApplicationGroupRequestModel) UnsetRestrictToTag()`

UnsetRestrictToTag ensures that no value is present for RestrictToTag, not even an explicit nil
### GetScopes

`func (o *CreateApplicationGroupRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateApplicationGroupRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateApplicationGroupRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateApplicationGroupRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *CreateApplicationGroupRequestModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *CreateApplicationGroupRequestModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *CreateApplicationGroupRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateApplicationGroupRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateApplicationGroupRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateApplicationGroupRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *CreateApplicationGroupRequestModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *CreateApplicationGroupRequestModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetSessionSharingEnabled

`func (o *CreateApplicationGroupRequestModel) GetSessionSharingEnabled() bool`

GetSessionSharingEnabled returns the SessionSharingEnabled field if non-nil, zero value otherwise.

### GetSessionSharingEnabledOk

`func (o *CreateApplicationGroupRequestModel) GetSessionSharingEnabledOk() (*bool, bool)`

GetSessionSharingEnabledOk returns a tuple with the SessionSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSharingEnabled

`func (o *CreateApplicationGroupRequestModel) SetSessionSharingEnabled(v bool)`

SetSessionSharingEnabled sets SessionSharingEnabled field to given value.

### HasSessionSharingEnabled

`func (o *CreateApplicationGroupRequestModel) HasSessionSharingEnabled() bool`

HasSessionSharingEnabled returns a boolean if a field has been set.

### SetSessionSharingEnabledNil

`func (o *CreateApplicationGroupRequestModel) SetSessionSharingEnabledNil(b bool)`

 SetSessionSharingEnabledNil sets the value for SessionSharingEnabled to be an explicit nil

### UnsetSessionSharingEnabled
`func (o *CreateApplicationGroupRequestModel) UnsetSessionSharingEnabled()`

UnsetSessionSharingEnabled ensures that no value is present for SessionSharingEnabled, not even an explicit nil
### GetIncludedUserFilterEnabled

`func (o *CreateApplicationGroupRequestModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *CreateApplicationGroupRequestModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *CreateApplicationGroupRequestModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *CreateApplicationGroupRequestModel) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### SetIncludedUserFilterEnabledNil

`func (o *CreateApplicationGroupRequestModel) SetIncludedUserFilterEnabledNil(b bool)`

 SetIncludedUserFilterEnabledNil sets the value for IncludedUserFilterEnabled to be an explicit nil

### UnsetIncludedUserFilterEnabled
`func (o *CreateApplicationGroupRequestModel) UnsetIncludedUserFilterEnabled()`

UnsetIncludedUserFilterEnabled ensures that no value is present for IncludedUserFilterEnabled, not even an explicit nil
### GetIncludedUsers

`func (o *CreateApplicationGroupRequestModel) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *CreateApplicationGroupRequestModel) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *CreateApplicationGroupRequestModel) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *CreateApplicationGroupRequestModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *CreateApplicationGroupRequestModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *CreateApplicationGroupRequestModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil
### GetDeliveryGroups

`func (o *CreateApplicationGroupRequestModel) GetDeliveryGroups() []PriorityRefRequestModel`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *CreateApplicationGroupRequestModel) GetDeliveryGroupsOk() (*[]PriorityRefRequestModel, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *CreateApplicationGroupRequestModel) SetDeliveryGroups(v []PriorityRefRequestModel)`

SetDeliveryGroups sets DeliveryGroups field to given value.

### HasDeliveryGroups

`func (o *CreateApplicationGroupRequestModel) HasDeliveryGroups() bool`

HasDeliveryGroups returns a boolean if a field has been set.

### SetDeliveryGroupsNil

`func (o *CreateApplicationGroupRequestModel) SetDeliveryGroupsNil(b bool)`

 SetDeliveryGroupsNil sets the value for DeliveryGroups to be an explicit nil

### UnsetDeliveryGroups
`func (o *CreateApplicationGroupRequestModel) UnsetDeliveryGroups()`

UnsetDeliveryGroups ensures that no value is present for DeliveryGroups, not even an explicit nil
### GetMetadata

`func (o *CreateApplicationGroupRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateApplicationGroupRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateApplicationGroupRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateApplicationGroupRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateApplicationGroupRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateApplicationGroupRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


