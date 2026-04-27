# AzurePoolSubscriptionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | Id of the tenant being onboarded | [optional] 
**SubscriptionId** | Pointer to **NullableString** | Id of the subscription being onboarded | [optional] 
**UnifiedEntraId** | Pointer to **bool** | Indicates if this subscriptions uses the new unified Entra ID | [optional] 
**State** | Pointer to [**DirectoryState**](DirectoryState.md) | Current state of subscription onboarding process | [optional] 
**ErrorDetails** | Pointer to **NullableString** | Failure details if any | [optional] 

## Methods

### NewAzurePoolSubscriptionDetails

`func NewAzurePoolSubscriptionDetails() *AzurePoolSubscriptionDetails`

NewAzurePoolSubscriptionDetails instantiates a new AzurePoolSubscriptionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurePoolSubscriptionDetailsWithDefaults

`func NewAzurePoolSubscriptionDetailsWithDefaults() *AzurePoolSubscriptionDetails`

NewAzurePoolSubscriptionDetailsWithDefaults instantiates a new AzurePoolSubscriptionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzurePoolSubscriptionDetails) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzurePoolSubscriptionDetails) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzurePoolSubscriptionDetails) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzurePoolSubscriptionDetails) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AzurePoolSubscriptionDetails) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AzurePoolSubscriptionDetails) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetSubscriptionId

`func (o *AzurePoolSubscriptionDetails) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzurePoolSubscriptionDetails) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzurePoolSubscriptionDetails) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzurePoolSubscriptionDetails) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *AzurePoolSubscriptionDetails) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AzurePoolSubscriptionDetails) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetUnifiedEntraId

`func (o *AzurePoolSubscriptionDetails) GetUnifiedEntraId() bool`

GetUnifiedEntraId returns the UnifiedEntraId field if non-nil, zero value otherwise.

### GetUnifiedEntraIdOk

`func (o *AzurePoolSubscriptionDetails) GetUnifiedEntraIdOk() (*bool, bool)`

GetUnifiedEntraIdOk returns a tuple with the UnifiedEntraId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedEntraId

`func (o *AzurePoolSubscriptionDetails) SetUnifiedEntraId(v bool)`

SetUnifiedEntraId sets UnifiedEntraId field to given value.

### HasUnifiedEntraId

`func (o *AzurePoolSubscriptionDetails) HasUnifiedEntraId() bool`

HasUnifiedEntraId returns a boolean if a field has been set.

### GetState

`func (o *AzurePoolSubscriptionDetails) GetState() DirectoryState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AzurePoolSubscriptionDetails) GetStateOk() (*DirectoryState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AzurePoolSubscriptionDetails) SetState(v DirectoryState)`

SetState sets State field to given value.

### HasState

`func (o *AzurePoolSubscriptionDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetErrorDetails

`func (o *AzurePoolSubscriptionDetails) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *AzurePoolSubscriptionDetails) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *AzurePoolSubscriptionDetails) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *AzurePoolSubscriptionDetails) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### SetErrorDetailsNil

`func (o *AzurePoolSubscriptionDetails) SetErrorDetailsNil(b bool)`

 SetErrorDetailsNil sets the value for ErrorDetails to be an explicit nil

### UnsetErrorDetails
`func (o *AzurePoolSubscriptionDetails) UnsetErrorDetails()`

UnsetErrorDetails ensures that no value is present for ErrorDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


