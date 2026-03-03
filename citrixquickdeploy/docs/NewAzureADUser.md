# NewAzureADUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Username for the account, should not include the domain | [optional] 
**FirstName** | Pointer to **NullableString** | First name of the user | [optional] 
**LastName** | Pointer to **NullableString** | Last name of the user | [optional] 
**CspCustomerId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 

## Methods

### NewNewAzureADUser

`func NewNewAzureADUser() *NewAzureADUser`

NewNewAzureADUser instantiates a new NewAzureADUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAzureADUserWithDefaults

`func NewNewAzureADUserWithDefaults() *NewAzureADUser`

NewNewAzureADUserWithDefaults instantiates a new NewAzureADUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *NewAzureADUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewAzureADUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewAzureADUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NewAzureADUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *NewAzureADUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *NewAzureADUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFirstName

`func (o *NewAzureADUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *NewAzureADUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *NewAzureADUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *NewAzureADUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *NewAzureADUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *NewAzureADUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *NewAzureADUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *NewAzureADUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *NewAzureADUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *NewAzureADUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *NewAzureADUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *NewAzureADUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCspCustomerId

`func (o *NewAzureADUser) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *NewAzureADUser) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *NewAzureADUser) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *NewAzureADUser) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### SetCspCustomerIdNil

`func (o *NewAzureADUser) SetCspCustomerIdNil(b bool)`

 SetCspCustomerIdNil sets the value for CspCustomerId to be an explicit nil

### UnsetCspCustomerId
`func (o *NewAzureADUser) UnsetCspCustomerId()`

UnsetCspCustomerId ensures that no value is present for CspCustomerId, not even an explicit nil
### GetCspSiteId

`func (o *NewAzureADUser) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *NewAzureADUser) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *NewAzureADUser) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *NewAzureADUser) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.

### SetCspSiteIdNil

`func (o *NewAzureADUser) SetCspSiteIdNil(b bool)`

 SetCspSiteIdNil sets the value for CspSiteId to be an explicit nil

### UnsetCspSiteId
`func (o *NewAzureADUser) UnsetCspSiteId()`

UnsetCspSiteId ensures that no value is present for CspSiteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


