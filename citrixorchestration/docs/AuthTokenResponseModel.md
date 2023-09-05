# AuthTokenResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The FMA token. | [optional] 
**Principal** | Pointer to **string** | The user principal. | [optional] 
**UserId** | Pointer to **string** | The user id. | [optional] 
**CustomerId** | Pointer to **string** | The customer id. | [optional] 
**ExpiresAt** | Pointer to **string** | The token will expire at the UTC time. | [optional] 

## Methods

### NewAuthTokenResponseModel

`func NewAuthTokenResponseModel() *AuthTokenResponseModel`

NewAuthTokenResponseModel instantiates a new AuthTokenResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenResponseModelWithDefaults

`func NewAuthTokenResponseModelWithDefaults() *AuthTokenResponseModel`

NewAuthTokenResponseModelWithDefaults instantiates a new AuthTokenResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AuthTokenResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthTokenResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthTokenResponseModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthTokenResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetPrincipal

`func (o *AuthTokenResponseModel) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *AuthTokenResponseModel) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *AuthTokenResponseModel) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *AuthTokenResponseModel) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetUserId

`func (o *AuthTokenResponseModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthTokenResponseModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthTokenResponseModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthTokenResponseModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AuthTokenResponseModel) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AuthTokenResponseModel) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AuthTokenResponseModel) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AuthTokenResponseModel) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AuthTokenResponseModel) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AuthTokenResponseModel) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AuthTokenResponseModel) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AuthTokenResponseModel) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


