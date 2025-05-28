# AzureToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **string** |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AzureTokenType**](AzureTokenType.md) |  | [optional] 

## Methods

### NewAzureToken

`func NewAzureToken() *AzureToken`

NewAzureToken instantiates a new AzureToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureTokenWithDefaults

`func NewAzureTokenWithDefaults() *AzureToken`

NewAzureTokenWithDefaults instantiates a new AzureToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *AzureToken) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AzureToken) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AzureToken) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AzureToken) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetAccessToken

`func (o *AzureToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AzureToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AzureToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AzureToken) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetType

`func (o *AzureToken) GetType() AzureTokenType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureToken) GetTypeOk() (*AzureTokenType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureToken) SetType(v AzureTokenType)`

SetType sets Type field to given value.

### HasType

`func (o *AzureToken) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


