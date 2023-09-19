# ApplicationGroupDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the application group. Used to be: Uuid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | 
**Uid** | Pointer to **NullableInt32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED. Use Id. | [optional] 
**ContainerScopes** | [**[]ContainerScopeResponseModel**](ContainerScopeResponseModel.md) | Delegated admin scopes in which the containers of the application group reside. | 
**Description** | Pointer to **NullableString** | Description of the application group. As with other facets of application groups, the description is not visible to end users. | [optional] 
**Enabled** | **bool** | Indicates whether or not the applications in this application group can be launched. | 
**Name** | Pointer to **NullableString** | Simple administrative name of application group within parent admin folder (if any). This property  is not guaranteed unique across all application groups. | [optional] 
**FullName** | Pointer to **NullableString** | Name of this application group. The name uniquely identifies the application group within the site. As with other facets of application groups, the name is not visible to end users. | [optional] 
**Scopes** | [**[]ScopeResponseModel**](ScopeResponseModel.md) | The list of the delegated admin scopes to which the application group belongs. | 
**Tags** | Pointer to **[]string** | Tags associated with this application group. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the application group is assigned to.  If &#x60;null&#x60;, the application group is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**NumApplications** | **int32** | Number of applications present in the application group. | 
**NumMachines** | **int32** | Number of machines capable of hosting the applications in the application group. | 
**NumMachinesWithTag** | Pointer to **int32** | Total number of machines across all desktop groups on  which the application group is published, and which are tagged with the tag given by the RestrictToTag property. | [optional] 
**RestrictToTag** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**DeliveryGroups** | [**[]ApplicationGroupDeliveryGroupRefResponseModel**](ApplicationGroupDeliveryGroupRefResponseModel.md) | Delivery groups associated with the application group. | 
**AdminFolder** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**ApplicationCount** | **int32** | Number of applications in the group. | 
**IncludedUsersFilterEnabled** | **bool** | Indicates whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access polic(ies) is implicitly granted access to the applications in the application group. | 
**IncludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The included users filter of the application group; that is, the users and groups who are explicitly granted access to the applications in the application group. | [optional] 
**SessionSharingEnabled** | **bool** | Whether applications in this application group can share sessions with applications in other application groups (or with applications that are not a member of an application group). | 
**TotalMachines** | **int32** | Total number of machines across all delivery groups on which the application group is published. | 

## Methods

### NewApplicationGroupDetailResponseModel

`func NewApplicationGroupDetailResponseModel(id string, containerScopes []ContainerScopeResponseModel, enabled bool, scopes []ScopeResponseModel, numApplications int32, numMachines int32, deliveryGroups []ApplicationGroupDeliveryGroupRefResponseModel, applicationCount int32, includedUsersFilterEnabled bool, sessionSharingEnabled bool, totalMachines int32, ) *ApplicationGroupDetailResponseModel`

NewApplicationGroupDetailResponseModel instantiates a new ApplicationGroupDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGroupDetailResponseModelWithDefaults

`func NewApplicationGroupDetailResponseModelWithDefaults() *ApplicationGroupDetailResponseModel`

NewApplicationGroupDetailResponseModelWithDefaults instantiates a new ApplicationGroupDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationGroupDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationGroupDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationGroupDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *ApplicationGroupDetailResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApplicationGroupDetailResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApplicationGroupDetailResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ApplicationGroupDetailResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ApplicationGroupDetailResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ApplicationGroupDetailResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetContainerScopes

`func (o *ApplicationGroupDetailResponseModel) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *ApplicationGroupDetailResponseModel) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *ApplicationGroupDetailResponseModel) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.


### GetDescription

`func (o *ApplicationGroupDetailResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationGroupDetailResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationGroupDetailResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationGroupDetailResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationGroupDetailResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationGroupDetailResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *ApplicationGroupDetailResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationGroupDetailResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationGroupDetailResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *ApplicationGroupDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationGroupDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationGroupDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationGroupDetailResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApplicationGroupDetailResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApplicationGroupDetailResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *ApplicationGroupDetailResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ApplicationGroupDetailResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ApplicationGroupDetailResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ApplicationGroupDetailResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ApplicationGroupDetailResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ApplicationGroupDetailResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetScopes

`func (o *ApplicationGroupDetailResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationGroupDetailResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationGroupDetailResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetTags

`func (o *ApplicationGroupDetailResponseModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationGroupDetailResponseModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationGroupDetailResponseModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationGroupDetailResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApplicationGroupDetailResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApplicationGroupDetailResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTenants

`func (o *ApplicationGroupDetailResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApplicationGroupDetailResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApplicationGroupDetailResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApplicationGroupDetailResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *ApplicationGroupDetailResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *ApplicationGroupDetailResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetNumApplications

`func (o *ApplicationGroupDetailResponseModel) GetNumApplications() int32`

GetNumApplications returns the NumApplications field if non-nil, zero value otherwise.

### GetNumApplicationsOk

`func (o *ApplicationGroupDetailResponseModel) GetNumApplicationsOk() (*int32, bool)`

GetNumApplicationsOk returns a tuple with the NumApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumApplications

`func (o *ApplicationGroupDetailResponseModel) SetNumApplications(v int32)`

SetNumApplications sets NumApplications field to given value.


### GetNumMachines

`func (o *ApplicationGroupDetailResponseModel) GetNumMachines() int32`

GetNumMachines returns the NumMachines field if non-nil, zero value otherwise.

### GetNumMachinesOk

`func (o *ApplicationGroupDetailResponseModel) GetNumMachinesOk() (*int32, bool)`

GetNumMachinesOk returns a tuple with the NumMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachines

`func (o *ApplicationGroupDetailResponseModel) SetNumMachines(v int32)`

SetNumMachines sets NumMachines field to given value.


### GetNumMachinesWithTag

`func (o *ApplicationGroupDetailResponseModel) GetNumMachinesWithTag() int32`

GetNumMachinesWithTag returns the NumMachinesWithTag field if non-nil, zero value otherwise.

### GetNumMachinesWithTagOk

`func (o *ApplicationGroupDetailResponseModel) GetNumMachinesWithTagOk() (*int32, bool)`

GetNumMachinesWithTagOk returns a tuple with the NumMachinesWithTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachinesWithTag

`func (o *ApplicationGroupDetailResponseModel) SetNumMachinesWithTag(v int32)`

SetNumMachinesWithTag sets NumMachinesWithTag field to given value.

### HasNumMachinesWithTag

`func (o *ApplicationGroupDetailResponseModel) HasNumMachinesWithTag() bool`

HasNumMachinesWithTag returns a boolean if a field has been set.

### GetRestrictToTag

`func (o *ApplicationGroupDetailResponseModel) GetRestrictToTag() RefResponseModel`

GetRestrictToTag returns the RestrictToTag field if non-nil, zero value otherwise.

### GetRestrictToTagOk

`func (o *ApplicationGroupDetailResponseModel) GetRestrictToTagOk() (*RefResponseModel, bool)`

GetRestrictToTagOk returns a tuple with the RestrictToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToTag

`func (o *ApplicationGroupDetailResponseModel) SetRestrictToTag(v RefResponseModel)`

SetRestrictToTag sets RestrictToTag field to given value.

### HasRestrictToTag

`func (o *ApplicationGroupDetailResponseModel) HasRestrictToTag() bool`

HasRestrictToTag returns a boolean if a field has been set.

### GetDeliveryGroups

`func (o *ApplicationGroupDetailResponseModel) GetDeliveryGroups() []ApplicationGroupDeliveryGroupRefResponseModel`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *ApplicationGroupDetailResponseModel) GetDeliveryGroupsOk() (*[]ApplicationGroupDeliveryGroupRefResponseModel, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *ApplicationGroupDetailResponseModel) SetDeliveryGroups(v []ApplicationGroupDeliveryGroupRefResponseModel)`

SetDeliveryGroups sets DeliveryGroups field to given value.


### GetAdminFolder

`func (o *ApplicationGroupDetailResponseModel) GetAdminFolder() RefResponseModel`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *ApplicationGroupDetailResponseModel) GetAdminFolderOk() (*RefResponseModel, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *ApplicationGroupDetailResponseModel) SetAdminFolder(v RefResponseModel)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *ApplicationGroupDetailResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.

### GetApplicationCount

`func (o *ApplicationGroupDetailResponseModel) GetApplicationCount() int32`

GetApplicationCount returns the ApplicationCount field if non-nil, zero value otherwise.

### GetApplicationCountOk

`func (o *ApplicationGroupDetailResponseModel) GetApplicationCountOk() (*int32, bool)`

GetApplicationCountOk returns a tuple with the ApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCount

`func (o *ApplicationGroupDetailResponseModel) SetApplicationCount(v int32)`

SetApplicationCount sets ApplicationCount field to given value.


### GetIncludedUsersFilterEnabled

`func (o *ApplicationGroupDetailResponseModel) GetIncludedUsersFilterEnabled() bool`

GetIncludedUsersFilterEnabled returns the IncludedUsersFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUsersFilterEnabledOk

`func (o *ApplicationGroupDetailResponseModel) GetIncludedUsersFilterEnabledOk() (*bool, bool)`

GetIncludedUsersFilterEnabledOk returns a tuple with the IncludedUsersFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsersFilterEnabled

`func (o *ApplicationGroupDetailResponseModel) SetIncludedUsersFilterEnabled(v bool)`

SetIncludedUsersFilterEnabled sets IncludedUsersFilterEnabled field to given value.


### GetIncludedUsers

`func (o *ApplicationGroupDetailResponseModel) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *ApplicationGroupDetailResponseModel) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *ApplicationGroupDetailResponseModel) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *ApplicationGroupDetailResponseModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *ApplicationGroupDetailResponseModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *ApplicationGroupDetailResponseModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil
### GetSessionSharingEnabled

`func (o *ApplicationGroupDetailResponseModel) GetSessionSharingEnabled() bool`

GetSessionSharingEnabled returns the SessionSharingEnabled field if non-nil, zero value otherwise.

### GetSessionSharingEnabledOk

`func (o *ApplicationGroupDetailResponseModel) GetSessionSharingEnabledOk() (*bool, bool)`

GetSessionSharingEnabledOk returns a tuple with the SessionSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSharingEnabled

`func (o *ApplicationGroupDetailResponseModel) SetSessionSharingEnabled(v bool)`

SetSessionSharingEnabled sets SessionSharingEnabled field to given value.


### GetTotalMachines

`func (o *ApplicationGroupDetailResponseModel) GetTotalMachines() int32`

GetTotalMachines returns the TotalMachines field if non-nil, zero value otherwise.

### GetTotalMachinesOk

`func (o *ApplicationGroupDetailResponseModel) GetTotalMachinesOk() (*int32, bool)`

GetTotalMachinesOk returns a tuple with the TotalMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMachines

`func (o *ApplicationGroupDetailResponseModel) SetTotalMachines(v int32)`

SetTotalMachines sets TotalMachines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


