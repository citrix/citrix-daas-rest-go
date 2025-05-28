# RemoveAzureADUserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email of the account to remove. | [optional] 
**CspCustomerId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null. | [optional] 
**CspSiteId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null. | [optional] 

## Methods

### NewRemoveAzureADUserModel

`func NewRemoveAzureADUserModel() *RemoveAzureADUserModel`

NewRemoveAzureADUserModel instantiates a new RemoveAzureADUserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveAzureADUserModelWithDefaults

`func NewRemoveAzureADUserModelWithDefaults() *RemoveAzureADUserModel`

NewRemoveAzureADUserModelWithDefaults instantiates a new RemoveAzureADUserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RemoveAzureADUserModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RemoveAzureADUserModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RemoveAzureADUserModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RemoveAzureADUserModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCspCustomerId

`func (o *RemoveAzureADUserModel) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *RemoveAzureADUserModel) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *RemoveAzureADUserModel) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *RemoveAzureADUserModel) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### GetCspSiteId

`func (o *RemoveAzureADUserModel) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *RemoveAzureADUserModel) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *RemoveAzureADUserModel) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *RemoveAzureADUserModel) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


