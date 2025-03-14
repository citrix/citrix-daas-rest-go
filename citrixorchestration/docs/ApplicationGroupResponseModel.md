# ApplicationGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the application group. Used to be: Uuid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | [optional] 
**Uid** | Pointer to **NullableInt32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED. Use Id. | [optional] 
**ContainerScopes** | Pointer to [**[]ContainerScopeResponseModel**](ContainerScopeResponseModel.md) | Delegated admin scopes in which the containers of the application group reside. | [optional] 
**Description** | Pointer to **NullableString** | Description of the application group. As with other facets of application groups, the description is not visible to end users. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether or not the applications in this application group can be launched. | [optional] 
**Name** | Pointer to **NullableString** | Simple administrative name of application group within parent admin folder (if any). This property is not guaranteed unique across all application groups. | [optional] 
**FullName** | Pointer to **NullableString** | Name of this application group. The name uniquely identifies the application group within the site. As with other facets of application groups, the name is not visible to end users. | [optional] 
**Scopes** | Pointer to [**[]ScopeResponseModel**](ScopeResponseModel.md) | The list of the delegated admin scopes to which the application group belongs. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of application group. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with this application group. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the application group is assigned to.  If &#x60;null&#x60;, the application group is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**NumApplications** | Pointer to **int32** | Number of applications present in the application group. | [optional] 
**NumMachines** | Pointer to **int32** | Number of machines capable of hosting the applications in the application group. | [optional] 
**NumMachinesWithTag** | Pointer to **int32** | Total number of machines across all desktop groups on which the application group is published, and which are tagged with the tag given by the RestrictToTag property. | [optional] 
**RestrictToTag** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**DeliveryGroups** | Pointer to [**[]ApplicationGroupDeliveryGroupRefResponseModel**](ApplicationGroupDeliveryGroupRefResponseModel.md) | Delivery groups associated with the application group. | [optional] 
**AssociatedDeliveryGroupUids** | Pointer to **[]int32** | Delivery group uids associated with the application group. | [optional] 
**AdminFolder** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 

## Methods

### NewApplicationGroupResponseModel

`func NewApplicationGroupResponseModel() *ApplicationGroupResponseModel`

NewApplicationGroupResponseModel instantiates a new ApplicationGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGroupResponseModelWithDefaults

`func NewApplicationGroupResponseModelWithDefaults() *ApplicationGroupResponseModel`

NewApplicationGroupResponseModelWithDefaults instantiates a new ApplicationGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationGroupResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ApplicationGroupResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ApplicationGroupResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *ApplicationGroupResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApplicationGroupResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApplicationGroupResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ApplicationGroupResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ApplicationGroupResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ApplicationGroupResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetContainerScopes

`func (o *ApplicationGroupResponseModel) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *ApplicationGroupResponseModel) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *ApplicationGroupResponseModel) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.

### HasContainerScopes

`func (o *ApplicationGroupResponseModel) HasContainerScopes() bool`

HasContainerScopes returns a boolean if a field has been set.

### SetContainerScopesNil

`func (o *ApplicationGroupResponseModel) SetContainerScopesNil(b bool)`

 SetContainerScopesNil sets the value for ContainerScopes to be an explicit nil

### UnsetContainerScopes
`func (o *ApplicationGroupResponseModel) UnsetContainerScopes()`

UnsetContainerScopes ensures that no value is present for ContainerScopes, not even an explicit nil
### GetDescription

`func (o *ApplicationGroupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationGroupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationGroupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationGroupResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationGroupResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationGroupResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *ApplicationGroupResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationGroupResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationGroupResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplicationGroupResponseModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *ApplicationGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationGroupResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApplicationGroupResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApplicationGroupResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *ApplicationGroupResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ApplicationGroupResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ApplicationGroupResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ApplicationGroupResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ApplicationGroupResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ApplicationGroupResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetScopes

`func (o *ApplicationGroupResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationGroupResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationGroupResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApplicationGroupResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ApplicationGroupResponseModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ApplicationGroupResponseModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetMetadata

`func (o *ApplicationGroupResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationGroupResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationGroupResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationGroupResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ApplicationGroupResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ApplicationGroupResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetTags

`func (o *ApplicationGroupResponseModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationGroupResponseModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationGroupResponseModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationGroupResponseModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApplicationGroupResponseModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApplicationGroupResponseModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTenants

`func (o *ApplicationGroupResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApplicationGroupResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApplicationGroupResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApplicationGroupResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *ApplicationGroupResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *ApplicationGroupResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetNumApplications

`func (o *ApplicationGroupResponseModel) GetNumApplications() int32`

GetNumApplications returns the NumApplications field if non-nil, zero value otherwise.

### GetNumApplicationsOk

`func (o *ApplicationGroupResponseModel) GetNumApplicationsOk() (*int32, bool)`

GetNumApplicationsOk returns a tuple with the NumApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumApplications

`func (o *ApplicationGroupResponseModel) SetNumApplications(v int32)`

SetNumApplications sets NumApplications field to given value.

### HasNumApplications

`func (o *ApplicationGroupResponseModel) HasNumApplications() bool`

HasNumApplications returns a boolean if a field has been set.

### GetNumMachines

`func (o *ApplicationGroupResponseModel) GetNumMachines() int32`

GetNumMachines returns the NumMachines field if non-nil, zero value otherwise.

### GetNumMachinesOk

`func (o *ApplicationGroupResponseModel) GetNumMachinesOk() (*int32, bool)`

GetNumMachinesOk returns a tuple with the NumMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachines

`func (o *ApplicationGroupResponseModel) SetNumMachines(v int32)`

SetNumMachines sets NumMachines field to given value.

### HasNumMachines

`func (o *ApplicationGroupResponseModel) HasNumMachines() bool`

HasNumMachines returns a boolean if a field has been set.

### GetNumMachinesWithTag

`func (o *ApplicationGroupResponseModel) GetNumMachinesWithTag() int32`

GetNumMachinesWithTag returns the NumMachinesWithTag field if non-nil, zero value otherwise.

### GetNumMachinesWithTagOk

`func (o *ApplicationGroupResponseModel) GetNumMachinesWithTagOk() (*int32, bool)`

GetNumMachinesWithTagOk returns a tuple with the NumMachinesWithTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachinesWithTag

`func (o *ApplicationGroupResponseModel) SetNumMachinesWithTag(v int32)`

SetNumMachinesWithTag sets NumMachinesWithTag field to given value.

### HasNumMachinesWithTag

`func (o *ApplicationGroupResponseModel) HasNumMachinesWithTag() bool`

HasNumMachinesWithTag returns a boolean if a field has been set.

### GetRestrictToTag

`func (o *ApplicationGroupResponseModel) GetRestrictToTag() RefResponseModel`

GetRestrictToTag returns the RestrictToTag field if non-nil, zero value otherwise.

### GetRestrictToTagOk

`func (o *ApplicationGroupResponseModel) GetRestrictToTagOk() (*RefResponseModel, bool)`

GetRestrictToTagOk returns a tuple with the RestrictToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToTag

`func (o *ApplicationGroupResponseModel) SetRestrictToTag(v RefResponseModel)`

SetRestrictToTag sets RestrictToTag field to given value.

### HasRestrictToTag

`func (o *ApplicationGroupResponseModel) HasRestrictToTag() bool`

HasRestrictToTag returns a boolean if a field has been set.

### GetDeliveryGroups

`func (o *ApplicationGroupResponseModel) GetDeliveryGroups() []ApplicationGroupDeliveryGroupRefResponseModel`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *ApplicationGroupResponseModel) GetDeliveryGroupsOk() (*[]ApplicationGroupDeliveryGroupRefResponseModel, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *ApplicationGroupResponseModel) SetDeliveryGroups(v []ApplicationGroupDeliveryGroupRefResponseModel)`

SetDeliveryGroups sets DeliveryGroups field to given value.

### HasDeliveryGroups

`func (o *ApplicationGroupResponseModel) HasDeliveryGroups() bool`

HasDeliveryGroups returns a boolean if a field has been set.

### SetDeliveryGroupsNil

`func (o *ApplicationGroupResponseModel) SetDeliveryGroupsNil(b bool)`

 SetDeliveryGroupsNil sets the value for DeliveryGroups to be an explicit nil

### UnsetDeliveryGroups
`func (o *ApplicationGroupResponseModel) UnsetDeliveryGroups()`

UnsetDeliveryGroups ensures that no value is present for DeliveryGroups, not even an explicit nil
### GetAssociatedDeliveryGroupUids

`func (o *ApplicationGroupResponseModel) GetAssociatedDeliveryGroupUids() []int32`

GetAssociatedDeliveryGroupUids returns the AssociatedDeliveryGroupUids field if non-nil, zero value otherwise.

### GetAssociatedDeliveryGroupUidsOk

`func (o *ApplicationGroupResponseModel) GetAssociatedDeliveryGroupUidsOk() (*[]int32, bool)`

GetAssociatedDeliveryGroupUidsOk returns a tuple with the AssociatedDeliveryGroupUids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDeliveryGroupUids

`func (o *ApplicationGroupResponseModel) SetAssociatedDeliveryGroupUids(v []int32)`

SetAssociatedDeliveryGroupUids sets AssociatedDeliveryGroupUids field to given value.

### HasAssociatedDeliveryGroupUids

`func (o *ApplicationGroupResponseModel) HasAssociatedDeliveryGroupUids() bool`

HasAssociatedDeliveryGroupUids returns a boolean if a field has been set.

### SetAssociatedDeliveryGroupUidsNil

`func (o *ApplicationGroupResponseModel) SetAssociatedDeliveryGroupUidsNil(b bool)`

 SetAssociatedDeliveryGroupUidsNil sets the value for AssociatedDeliveryGroupUids to be an explicit nil

### UnsetAssociatedDeliveryGroupUids
`func (o *ApplicationGroupResponseModel) UnsetAssociatedDeliveryGroupUids()`

UnsetAssociatedDeliveryGroupUids ensures that no value is present for AssociatedDeliveryGroupUids, not even an explicit nil
### GetAdminFolder

`func (o *ApplicationGroupResponseModel) GetAdminFolder() RefResponseModel`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *ApplicationGroupResponseModel) GetAdminFolderOk() (*RefResponseModel, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *ApplicationGroupResponseModel) SetAdminFolder(v RefResponseModel)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *ApplicationGroupResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


