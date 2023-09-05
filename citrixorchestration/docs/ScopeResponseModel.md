# ScopeResponseModel

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

### NewScopeResponseModel

`func NewScopeResponseModel(id string, name string, isBuiltIn bool, isAllScope bool, isTenantScope bool, ) *ScopeResponseModel`

NewScopeResponseModel instantiates a new ScopeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeResponseModelWithDefaults

`func NewScopeResponseModelWithDefaults() *ScopeResponseModel`

NewScopeResponseModelWithDefaults instantiates a new ScopeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScopeResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScopeResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScopeResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetUid

`func (o *ScopeResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ScopeResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ScopeResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ScopeResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *ScopeResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScopeResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScopeResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScopeResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScopeResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScopeResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScopeResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsBuiltIn

`func (o *ScopeResponseModel) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *ScopeResponseModel) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *ScopeResponseModel) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.


### GetIsAllScope

`func (o *ScopeResponseModel) GetIsAllScope() bool`

GetIsAllScope returns the IsAllScope field if non-nil, zero value otherwise.

### GetIsAllScopeOk

`func (o *ScopeResponseModel) GetIsAllScopeOk() (*bool, bool)`

GetIsAllScopeOk returns a tuple with the IsAllScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllScope

`func (o *ScopeResponseModel) SetIsAllScope(v bool)`

SetIsAllScope sets IsAllScope field to given value.


### GetIsTenantScope

`func (o *ScopeResponseModel) GetIsTenantScope() bool`

GetIsTenantScope returns the IsTenantScope field if non-nil, zero value otherwise.

### GetIsTenantScopeOk

`func (o *ScopeResponseModel) GetIsTenantScopeOk() (*bool, bool)`

GetIsTenantScopeOk returns a tuple with the IsTenantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTenantScope

`func (o *ScopeResponseModel) SetIsTenantScope(v bool)`

SetIsTenantScope sets IsTenantScope field to given value.


### GetTenantId

`func (o *ScopeResponseModel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ScopeResponseModel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ScopeResponseModel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ScopeResponseModel) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTenantName

`func (o *ScopeResponseModel) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *ScopeResponseModel) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *ScopeResponseModel) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *ScopeResponseModel) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


