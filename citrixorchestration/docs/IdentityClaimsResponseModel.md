# IdentityClaimsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oid** | Pointer to **string** | Oid claim of the identity. | [optional] 
**Sid** | Pointer to **string** | Sid claim of the identity. | [optional] 
**Upn** | Pointer to **string** | Upn claim of the identity. | [optional] 
**Email** | Pointer to **string** | Email claim of the identity. | [optional] 

## Methods

### NewIdentityClaimsResponseModel

`func NewIdentityClaimsResponseModel() *IdentityClaimsResponseModel`

NewIdentityClaimsResponseModel instantiates a new IdentityClaimsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityClaimsResponseModelWithDefaults

`func NewIdentityClaimsResponseModelWithDefaults() *IdentityClaimsResponseModel`

NewIdentityClaimsResponseModelWithDefaults instantiates a new IdentityClaimsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOid

`func (o *IdentityClaimsResponseModel) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *IdentityClaimsResponseModel) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *IdentityClaimsResponseModel) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *IdentityClaimsResponseModel) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetSid

`func (o *IdentityClaimsResponseModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IdentityClaimsResponseModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IdentityClaimsResponseModel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IdentityClaimsResponseModel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetUpn

`func (o *IdentityClaimsResponseModel) GetUpn() string`

GetUpn returns the Upn field if non-nil, zero value otherwise.

### GetUpnOk

`func (o *IdentityClaimsResponseModel) GetUpnOk() (*string, bool)`

GetUpnOk returns a tuple with the Upn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpn

`func (o *IdentityClaimsResponseModel) SetUpn(v string)`

SetUpn sets Upn field to given value.

### HasUpn

`func (o *IdentityClaimsResponseModel) HasUpn() bool`

HasUpn returns a boolean if a field has been set.

### GetEmail

`func (o *IdentityClaimsResponseModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityClaimsResponseModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityClaimsResponseModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdentityClaimsResponseModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


