# ApplicationGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the application group. Used to be: Uuid Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED. Use Id. | [optional] 
**ContainerScopes** | [**[]ContainerScopeResponseModel**](ContainerScopeResponseModel.md) | Delegated admin scopes in which the containers of the application group reside. | 
**Description** | Pointer to **string** | Description of the application group. As with other facets of application groups, the description is not visible to end users. | [optional] 
**Enabled** | **bool** | Indicates whether or not the applications in this application group can be launched. | 
**Name** | Pointer to **string** | Simple administrative name of application group within parent admin folder (if any). This property  is not guaranteed unique across all application groups. | [optional] 
**FullName** | Pointer to **string** | Name of this application group. The name uniquely identifies the application group within the site. As with other facets of application groups, the name is not visible to end users. | [optional] 
**Scopes** | [**[]ScopeResponseModel**](ScopeResponseModel.md) | The list of the delegated admin scopes to which the application group belongs. | 
**Tags** | Pointer to **[]string** | Tags associated with this application group. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the application group is assigned to.  If &#x60;null&#x60;, the application group is not assigned to tenants, and may be used by any tenant, including future added tenants. | [optional] 
**NumApplications** | **int32** | Number of applications present in the application group. | 
**NumMachines** | **int32** | Number of machines capable of hosting the applications in the application group. | 
**NumMachinesWithTag** | Pointer to **int32** | Total number of machines across all desktop groups on  which the application group is published, and which are tagged with the tag given by the RestrictToTag property. | [optional] 
**RestrictToTag** | Pointer to [**ApplicationGroupResponseModelRestrictToTag**](ApplicationGroupResponseModelRestrictToTag.md) |  | [optional] 
**DeliveryGroups** | [**[]ApplicationGroupDeliveryGroupRefResponseModel**](ApplicationGroupDeliveryGroupRefResponseModel.md) | Delivery groups associated with the application group. | 
**AdminFolder** | Pointer to [**ApplicationGroupResponseModelAdminFolder**](ApplicationGroupResponseModelAdminFolder.md) |  | [optional] 

## Methods

### NewApplicationGroupResponseModel

`func NewApplicationGroupResponseModel(id string, containerScopes []ContainerScopeResponseModel, enabled bool, scopes []ScopeResponseModel, numApplications int32, numMachines int32, deliveryGroups []ApplicationGroupDeliveryGroupRefResponseModel, ) *ApplicationGroupResponseModel`

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

`func (o *ApplicationGroupResponseModel) GetRestrictToTag() ApplicationGroupResponseModelRestrictToTag`

GetRestrictToTag returns the RestrictToTag field if non-nil, zero value otherwise.

### GetRestrictToTagOk

`func (o *ApplicationGroupResponseModel) GetRestrictToTagOk() (*ApplicationGroupResponseModelRestrictToTag, bool)`

GetRestrictToTagOk returns a tuple with the RestrictToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToTag

`func (o *ApplicationGroupResponseModel) SetRestrictToTag(v ApplicationGroupResponseModelRestrictToTag)`

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


### GetAdminFolder

`func (o *ApplicationGroupResponseModel) GetAdminFolder() ApplicationGroupResponseModelAdminFolder`

GetAdminFolder returns the AdminFolder field if non-nil, zero value otherwise.

### GetAdminFolderOk

`func (o *ApplicationGroupResponseModel) GetAdminFolderOk() (*ApplicationGroupResponseModelAdminFolder, bool)`

GetAdminFolderOk returns a tuple with the AdminFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFolder

`func (o *ApplicationGroupResponseModel) SetAdminFolder(v ApplicationGroupResponseModelAdminFolder)`

SetAdminFolder sets AdminFolder field to given value.

### HasAdminFolder

`func (o *ApplicationGroupResponseModel) HasAdminFolder() bool`

HasAdminFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


