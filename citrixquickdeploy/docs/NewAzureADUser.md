# NewAzureADUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Username for the account, should not include the domain | [optional] 
**FirstName** | Pointer to **string** | First name of the user | [optional] 
**LastName** | Pointer to **string** | Last name of the user | [optional] 
**CspCustomerId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


