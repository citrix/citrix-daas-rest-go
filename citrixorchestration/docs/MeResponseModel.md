# MeResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | Id of the logged-in user. Internally: this is the *sub* from the bearer token.  NOT the user_id. | 
**DisplayName** | **string** | Display name of the user. Internally: the *displayname* from the bearer token. | 
**ExpiryTime** | **string** | Time when the login token expires. RFC 3339 compatible format. Internally: the *exp* from the bearer token, converted to RFC 3339. | 
**RefreshExpirationTime** | **string** | Time after which the login token can no longer be refreshed. RFC 3339 compatible format. Internally: the *refresh_expiration* from the bearer token, converted to RFC 3339. | 
**VerifiedEmail** | Pointer to **string** | Verified email of the user. | [optional] 
**IsCspCustomer** | Pointer to **bool** | True if the current customer is a csp customer(partner/tenant) | [optional] 
**Customers** | [**[]MyCustomerResponseModel**](MyCustomerResponseModel.md) | List of customers and sites accessible by the user. | 

## Methods

### NewMeResponseModel

`func NewMeResponseModel(userId string, displayName string, expiryTime string, refreshExpirationTime string, customers []MyCustomerResponseModel, ) *MeResponseModel`

NewMeResponseModel instantiates a new MeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeResponseModelWithDefaults

`func NewMeResponseModelWithDefaults() *MeResponseModel`

NewMeResponseModelWithDefaults instantiates a new MeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *MeResponseModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MeResponseModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MeResponseModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetDisplayName

`func (o *MeResponseModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MeResponseModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MeResponseModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExpiryTime

`func (o *MeResponseModel) GetExpiryTime() string`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *MeResponseModel) GetExpiryTimeOk() (*string, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *MeResponseModel) SetExpiryTime(v string)`

SetExpiryTime sets ExpiryTime field to given value.


### GetRefreshExpirationTime

`func (o *MeResponseModel) GetRefreshExpirationTime() string`

GetRefreshExpirationTime returns the RefreshExpirationTime field if non-nil, zero value otherwise.

### GetRefreshExpirationTimeOk

`func (o *MeResponseModel) GetRefreshExpirationTimeOk() (*string, bool)`

GetRefreshExpirationTimeOk returns a tuple with the RefreshExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpirationTime

`func (o *MeResponseModel) SetRefreshExpirationTime(v string)`

SetRefreshExpirationTime sets RefreshExpirationTime field to given value.


### GetVerifiedEmail

`func (o *MeResponseModel) GetVerifiedEmail() string`

GetVerifiedEmail returns the VerifiedEmail field if non-nil, zero value otherwise.

### GetVerifiedEmailOk

`func (o *MeResponseModel) GetVerifiedEmailOk() (*string, bool)`

GetVerifiedEmailOk returns a tuple with the VerifiedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedEmail

`func (o *MeResponseModel) SetVerifiedEmail(v string)`

SetVerifiedEmail sets VerifiedEmail field to given value.

### HasVerifiedEmail

`func (o *MeResponseModel) HasVerifiedEmail() bool`

HasVerifiedEmail returns a boolean if a field has been set.

### GetIsCspCustomer

`func (o *MeResponseModel) GetIsCspCustomer() bool`

GetIsCspCustomer returns the IsCspCustomer field if non-nil, zero value otherwise.

### GetIsCspCustomerOk

`func (o *MeResponseModel) GetIsCspCustomerOk() (*bool, bool)`

GetIsCspCustomerOk returns a tuple with the IsCspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCspCustomer

`func (o *MeResponseModel) SetIsCspCustomer(v bool)`

SetIsCspCustomer sets IsCspCustomer field to given value.

### HasIsCspCustomer

`func (o *MeResponseModel) HasIsCspCustomer() bool`

HasIsCspCustomer returns a boolean if a field has been set.

### GetCustomers

`func (o *MeResponseModel) GetCustomers() []MyCustomerResponseModel`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *MeResponseModel) GetCustomersOk() (*[]MyCustomerResponseModel, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *MeResponseModel) SetCustomers(v []MyCustomerResponseModel)`

SetCustomers sets Customers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


