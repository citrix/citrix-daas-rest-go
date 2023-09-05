# AdministratorRightResponseModelScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the administrative scope. Used to be: ScopeId (and it was not globally unique) Needs to be globally unique Might be constructed from site ID + internal Uid | 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;Id&#39;/&gt;.&#x60; DEPRECATED.  Use Id. | [optional] 
**Name** | **string** | Name of the administrative scope. | 
**Description** | Pointer to **string** | Description of the administrative scope. | [optional] 
**IsBuiltIn** | **bool** | Whether the administrative scope is built-in. | 
**IsAllScope** | **bool** | Indicates the built-in \&quot;All\&quot; scope.  There will be exactly one scope with this property set to &#x60;true&#x60;. | 
**IsTenantScope** | **bool** | Whether the scope is created for CSP tenant. | 
**TenantId** | Pointer to **string** | Id of the CSP tenant. Valid when IsTenantScope is true. | [optional] 
**TenantName** | Pointer to **string** | Name of the CSP tenant. Valid when IsTenantScope is true. | [optional] 

## Methods

### NewAdministratorRightResponseModelScope

`func NewAdministratorRightResponseModelScope(id string, name string, isBuiltIn bool, isAllScope bool, isTenantScope bool, ) *AdministratorRightResponseModelScope`

NewAdministratorRightResponseModelScope instantiates a new AdministratorRightResponseModelScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorRightResponseModelScopeWithDefaults

`func NewAdministratorRightResponseModelScopeWithDefaults() *AdministratorRightResponseModelScope`

NewAdministratorRightResponseModelScopeWithDefaults instantiates a new AdministratorRightResponseModelScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdministratorRightResponseModelScope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdministratorRightResponseModelScope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdministratorRightResponseModelScope) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *AdministratorRightResponseModelScope) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AdministratorRightResponseModelScope) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AdministratorRightResponseModelScope) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AdministratorRightResponseModelScope) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *AdministratorRightResponseModelScope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdministratorRightResponseModelScope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdministratorRightResponseModelScope) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AdministratorRightResponseModelScope) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdministratorRightResponseModelScope) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdministratorRightResponseModelScope) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdministratorRightResponseModelScope) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsBuiltIn

`func (o *AdministratorRightResponseModelScope) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *AdministratorRightResponseModelScope) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *AdministratorRightResponseModelScope) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.


### GetIsAllScope

`func (o *AdministratorRightResponseModelScope) GetIsAllScope() bool`

GetIsAllScope returns the IsAllScope field if non-nil, zero value otherwise.

### GetIsAllScopeOk

`func (o *AdministratorRightResponseModelScope) GetIsAllScopeOk() (*bool, bool)`

GetIsAllScopeOk returns a tuple with the IsAllScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllScope

`func (o *AdministratorRightResponseModelScope) SetIsAllScope(v bool)`

SetIsAllScope sets IsAllScope field to given value.


### GetIsTenantScope

`func (o *AdministratorRightResponseModelScope) GetIsTenantScope() bool`

GetIsTenantScope returns the IsTenantScope field if non-nil, zero value otherwise.

### GetIsTenantScopeOk

`func (o *AdministratorRightResponseModelScope) GetIsTenantScopeOk() (*bool, bool)`

GetIsTenantScopeOk returns a tuple with the IsTenantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTenantScope

`func (o *AdministratorRightResponseModelScope) SetIsTenantScope(v bool)`

SetIsTenantScope sets IsTenantScope field to given value.


### GetTenantId

`func (o *AdministratorRightResponseModelScope) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AdministratorRightResponseModelScope) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AdministratorRightResponseModelScope) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AdministratorRightResponseModelScope) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTenantName

`func (o *AdministratorRightResponseModelScope) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *AdministratorRightResponseModelScope) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *AdministratorRightResponseModelScope) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *AdministratorRightResponseModelScope) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


