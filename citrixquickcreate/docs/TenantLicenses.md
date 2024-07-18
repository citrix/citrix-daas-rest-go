# TenantLicenses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | AAD Tenant of which the listed users are part of | [optional] 
**Users** | Pointer to [**[]CoreUser**](CoreUser.md) | List of users to assign/revoke license for in the provided AAD tenant | [optional] 

## Methods

### NewTenantLicenses

`func NewTenantLicenses() *TenantLicenses`

NewTenantLicenses instantiates a new TenantLicenses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantLicensesWithDefaults

`func NewTenantLicensesWithDefaults() *TenantLicenses`

NewTenantLicensesWithDefaults instantiates a new TenantLicenses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantLicenses) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantLicenses) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantLicenses) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantLicenses) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantLicenses) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantLicenses) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUsers

`func (o *TenantLicenses) GetUsers() []CoreUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *TenantLicenses) GetUsersOk() (*[]CoreUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *TenantLicenses) SetUsers(v []CoreUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *TenantLicenses) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *TenantLicenses) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *TenantLicenses) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


