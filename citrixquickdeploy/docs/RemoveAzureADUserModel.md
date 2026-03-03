# RemoveAzureADUserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Email of the account to remove. | [optional] 
**CspCustomerId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists otherwise null. | [optional] 
**CspSiteId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists otherwise null. | [optional] 

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

### SetEmailNil

`func (o *RemoveAzureADUserModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *RemoveAzureADUserModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
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

### SetCspCustomerIdNil

`func (o *RemoveAzureADUserModel) SetCspCustomerIdNil(b bool)`

 SetCspCustomerIdNil sets the value for CspCustomerId to be an explicit nil

### UnsetCspCustomerId
`func (o *RemoveAzureADUserModel) UnsetCspCustomerId()`

UnsetCspCustomerId ensures that no value is present for CspCustomerId, not even an explicit nil
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

### SetCspSiteIdNil

`func (o *RemoveAzureADUserModel) SetCspSiteIdNil(b bool)`

 SetCspSiteIdNil sets the value for CspSiteId to be an explicit nil

### UnsetCspSiteId
`func (o *RemoveAzureADUserModel) UnsetCspSiteId()`

UnsetCspSiteId ensures that no value is present for CspSiteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


