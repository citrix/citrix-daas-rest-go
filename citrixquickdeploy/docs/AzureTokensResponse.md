# AzureTokensResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | ID of the apps the tokens are applicable to | [optional] 
**Tokens** | Pointer to [**[]AzureToken**](AzureToken.md) | The tokens to the specified resources | [optional] 

## Methods

### NewAzureTokensResponse

`func NewAzureTokensResponse() *AzureTokensResponse`

NewAzureTokensResponse instantiates a new AzureTokensResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureTokensResponseWithDefaults

`func NewAzureTokensResponseWithDefaults() *AzureTokensResponse`

NewAzureTokensResponseWithDefaults instantiates a new AzureTokensResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AzureTokensResponse) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AzureTokensResponse) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AzureTokensResponse) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AzureTokensResponse) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetTokens

`func (o *AzureTokensResponse) GetTokens() []AzureToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *AzureTokensResponse) GetTokensOk() (*[]AzureToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *AzureTokensResponse) SetTokens(v []AzureToken)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *AzureTokensResponse) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


