# IdpInstanceCreateConnectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdaptiveAuthenticationConnectionDetails** | Pointer to [**AdaptiveAuthenticationConnectionModel**](AdaptiveAuthenticationConnectionModel.md) |  | [optional] 
**AuthDomainName** | Pointer to **NullableString** |  | [optional] 
**AzureAd** | Pointer to [**AzureAdConnectionSettings**](AzureAdConnectionSettings.md) |  | [optional] 
**GatewayConnectionDetails** | Pointer to [**GatewayConnectionModel**](GatewayConnectionModel.md) |  | [optional] 
**GoogleConnectionModel** | Pointer to [**GoogleConnectionModel**](GoogleConnectionModel.md) |  | [optional] 
**IdentityProviderType** | **string** |  | 
**OktaConnectionModel** | Pointer to [**OktaConnectionModel**](OktaConnectionModel.md) |  | [optional] 
**SamlConnectionModel** | Pointer to [**SamlConnectionModel**](SamlConnectionModel.md) |  | [optional] 
**IdentityProviderNickname** | **string** |  | 

## Methods

### NewIdpInstanceCreateConnectModel

`func NewIdpInstanceCreateConnectModel(identityProviderType string, identityProviderNickname string, ) *IdpInstanceCreateConnectModel`

NewIdpInstanceCreateConnectModel instantiates a new IdpInstanceCreateConnectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpInstanceCreateConnectModelWithDefaults

`func NewIdpInstanceCreateConnectModelWithDefaults() *IdpInstanceCreateConnectModel`

NewIdpInstanceCreateConnectModelWithDefaults instantiates a new IdpInstanceCreateConnectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptiveAuthenticationConnectionDetails

`func (o *IdpInstanceCreateConnectModel) GetAdaptiveAuthenticationConnectionDetails() AdaptiveAuthenticationConnectionModel`

GetAdaptiveAuthenticationConnectionDetails returns the AdaptiveAuthenticationConnectionDetails field if non-nil, zero value otherwise.

### GetAdaptiveAuthenticationConnectionDetailsOk

`func (o *IdpInstanceCreateConnectModel) GetAdaptiveAuthenticationConnectionDetailsOk() (*AdaptiveAuthenticationConnectionModel, bool)`

GetAdaptiveAuthenticationConnectionDetailsOk returns a tuple with the AdaptiveAuthenticationConnectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveAuthenticationConnectionDetails

`func (o *IdpInstanceCreateConnectModel) SetAdaptiveAuthenticationConnectionDetails(v AdaptiveAuthenticationConnectionModel)`

SetAdaptiveAuthenticationConnectionDetails sets AdaptiveAuthenticationConnectionDetails field to given value.

### HasAdaptiveAuthenticationConnectionDetails

`func (o *IdpInstanceCreateConnectModel) HasAdaptiveAuthenticationConnectionDetails() bool`

HasAdaptiveAuthenticationConnectionDetails returns a boolean if a field has been set.

### GetAuthDomainName

`func (o *IdpInstanceCreateConnectModel) GetAuthDomainName() string`

GetAuthDomainName returns the AuthDomainName field if non-nil, zero value otherwise.

### GetAuthDomainNameOk

`func (o *IdpInstanceCreateConnectModel) GetAuthDomainNameOk() (*string, bool)`

GetAuthDomainNameOk returns a tuple with the AuthDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomainName

`func (o *IdpInstanceCreateConnectModel) SetAuthDomainName(v string)`

SetAuthDomainName sets AuthDomainName field to given value.

### HasAuthDomainName

`func (o *IdpInstanceCreateConnectModel) HasAuthDomainName() bool`

HasAuthDomainName returns a boolean if a field has been set.

### SetAuthDomainNameNil

`func (o *IdpInstanceCreateConnectModel) SetAuthDomainNameNil(b bool)`

 SetAuthDomainNameNil sets the value for AuthDomainName to be an explicit nil

### UnsetAuthDomainName
`func (o *IdpInstanceCreateConnectModel) UnsetAuthDomainName()`

UnsetAuthDomainName ensures that no value is present for AuthDomainName, not even an explicit nil
### GetAzureAd

`func (o *IdpInstanceCreateConnectModel) GetAzureAd() AzureAdConnectionSettings`

GetAzureAd returns the AzureAd field if non-nil, zero value otherwise.

### GetAzureAdOk

`func (o *IdpInstanceCreateConnectModel) GetAzureAdOk() (*AzureAdConnectionSettings, bool)`

GetAzureAdOk returns a tuple with the AzureAd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAd

`func (o *IdpInstanceCreateConnectModel) SetAzureAd(v AzureAdConnectionSettings)`

SetAzureAd sets AzureAd field to given value.

### HasAzureAd

`func (o *IdpInstanceCreateConnectModel) HasAzureAd() bool`

HasAzureAd returns a boolean if a field has been set.

### GetGatewayConnectionDetails

`func (o *IdpInstanceCreateConnectModel) GetGatewayConnectionDetails() GatewayConnectionModel`

GetGatewayConnectionDetails returns the GatewayConnectionDetails field if non-nil, zero value otherwise.

### GetGatewayConnectionDetailsOk

`func (o *IdpInstanceCreateConnectModel) GetGatewayConnectionDetailsOk() (*GatewayConnectionModel, bool)`

GetGatewayConnectionDetailsOk returns a tuple with the GatewayConnectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConnectionDetails

`func (o *IdpInstanceCreateConnectModel) SetGatewayConnectionDetails(v GatewayConnectionModel)`

SetGatewayConnectionDetails sets GatewayConnectionDetails field to given value.

### HasGatewayConnectionDetails

`func (o *IdpInstanceCreateConnectModel) HasGatewayConnectionDetails() bool`

HasGatewayConnectionDetails returns a boolean if a field has been set.

### GetGoogleConnectionModel

`func (o *IdpInstanceCreateConnectModel) GetGoogleConnectionModel() GoogleConnectionModel`

GetGoogleConnectionModel returns the GoogleConnectionModel field if non-nil, zero value otherwise.

### GetGoogleConnectionModelOk

`func (o *IdpInstanceCreateConnectModel) GetGoogleConnectionModelOk() (*GoogleConnectionModel, bool)`

GetGoogleConnectionModelOk returns a tuple with the GoogleConnectionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleConnectionModel

`func (o *IdpInstanceCreateConnectModel) SetGoogleConnectionModel(v GoogleConnectionModel)`

SetGoogleConnectionModel sets GoogleConnectionModel field to given value.

### HasGoogleConnectionModel

`func (o *IdpInstanceCreateConnectModel) HasGoogleConnectionModel() bool`

HasGoogleConnectionModel returns a boolean if a field has been set.

### GetIdentityProviderType

`func (o *IdpInstanceCreateConnectModel) GetIdentityProviderType() string`

GetIdentityProviderType returns the IdentityProviderType field if non-nil, zero value otherwise.

### GetIdentityProviderTypeOk

`func (o *IdpInstanceCreateConnectModel) GetIdentityProviderTypeOk() (*string, bool)`

GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderType

`func (o *IdpInstanceCreateConnectModel) SetIdentityProviderType(v string)`

SetIdentityProviderType sets IdentityProviderType field to given value.


### GetOktaConnectionModel

`func (o *IdpInstanceCreateConnectModel) GetOktaConnectionModel() OktaConnectionModel`

GetOktaConnectionModel returns the OktaConnectionModel field if non-nil, zero value otherwise.

### GetOktaConnectionModelOk

`func (o *IdpInstanceCreateConnectModel) GetOktaConnectionModelOk() (*OktaConnectionModel, bool)`

GetOktaConnectionModelOk returns a tuple with the OktaConnectionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaConnectionModel

`func (o *IdpInstanceCreateConnectModel) SetOktaConnectionModel(v OktaConnectionModel)`

SetOktaConnectionModel sets OktaConnectionModel field to given value.

### HasOktaConnectionModel

`func (o *IdpInstanceCreateConnectModel) HasOktaConnectionModel() bool`

HasOktaConnectionModel returns a boolean if a field has been set.

### GetSamlConnectionModel

`func (o *IdpInstanceCreateConnectModel) GetSamlConnectionModel() SamlConnectionModel`

GetSamlConnectionModel returns the SamlConnectionModel field if non-nil, zero value otherwise.

### GetSamlConnectionModelOk

`func (o *IdpInstanceCreateConnectModel) GetSamlConnectionModelOk() (*SamlConnectionModel, bool)`

GetSamlConnectionModelOk returns a tuple with the SamlConnectionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlConnectionModel

`func (o *IdpInstanceCreateConnectModel) SetSamlConnectionModel(v SamlConnectionModel)`

SetSamlConnectionModel sets SamlConnectionModel field to given value.

### HasSamlConnectionModel

`func (o *IdpInstanceCreateConnectModel) HasSamlConnectionModel() bool`

HasSamlConnectionModel returns a boolean if a field has been set.

### GetIdentityProviderNickname

`func (o *IdpInstanceCreateConnectModel) GetIdentityProviderNickname() string`

GetIdentityProviderNickname returns the IdentityProviderNickname field if non-nil, zero value otherwise.

### GetIdentityProviderNicknameOk

`func (o *IdpInstanceCreateConnectModel) GetIdentityProviderNicknameOk() (*string, bool)`

GetIdentityProviderNicknameOk returns a tuple with the IdentityProviderNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderNickname

`func (o *IdpInstanceCreateConnectModel) SetIdentityProviderNickname(v string)`

SetIdentityProviderNickname sets IdentityProviderNickname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


