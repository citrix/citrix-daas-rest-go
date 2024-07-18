# ActiveTenantUserLicenses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | AAD Tenant of which the licensed users are part of | [optional] 
**LicensedUsers** | Pointer to **[]string** | User IDs of users with active licenses for which we want to revoke HDX License | [optional] 

## Methods

### NewActiveTenantUserLicenses

`func NewActiveTenantUserLicenses() *ActiveTenantUserLicenses`

NewActiveTenantUserLicenses instantiates a new ActiveTenantUserLicenses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveTenantUserLicensesWithDefaults

`func NewActiveTenantUserLicensesWithDefaults() *ActiveTenantUserLicenses`

NewActiveTenantUserLicensesWithDefaults instantiates a new ActiveTenantUserLicenses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ActiveTenantUserLicenses) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ActiveTenantUserLicenses) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ActiveTenantUserLicenses) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ActiveTenantUserLicenses) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ActiveTenantUserLicenses) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ActiveTenantUserLicenses) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetLicensedUsers

`func (o *ActiveTenantUserLicenses) GetLicensedUsers() []string`

GetLicensedUsers returns the LicensedUsers field if non-nil, zero value otherwise.

### GetLicensedUsersOk

`func (o *ActiveTenantUserLicenses) GetLicensedUsersOk() (*[]string, bool)`

GetLicensedUsersOk returns a tuple with the LicensedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedUsers

`func (o *ActiveTenantUserLicenses) SetLicensedUsers(v []string)`

SetLicensedUsers sets LicensedUsers field to given value.

### HasLicensedUsers

`func (o *ActiveTenantUserLicenses) HasLicensedUsers() bool`

HasLicensedUsers returns a boolean if a field has been set.

### SetLicensedUsersNil

`func (o *ActiveTenantUserLicenses) SetLicensedUsersNil(b bool)`

 SetLicensedUsersNil sets the value for LicensedUsers to be an explicit nil

### UnsetLicensedUsers
`func (o *ActiveTenantUserLicenses) UnsetLicensedUsers()`

UnsetLicensedUsers ensures that no value is present for LicensedUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


