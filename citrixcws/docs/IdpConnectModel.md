# IdpConnectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthDomainName** | Pointer to **NullableString** |  | [optional] 
**IdentityProviderId** | **string** |  | 
**AdaptiveAuthenticationConnectionDetails** | Pointer to [**AdaptiveAuthenticationConnectionModel**](AdaptiveAuthenticationConnectionModel.md) |  | [optional] 
**AdOtpDetails** | Pointer to [**AdOtpConnectionModel**](AdOtpConnectionModel.md) |  | [optional] 
**GoogleConnectionModel** | Pointer to [**GoogleConnectionModel**](GoogleConnectionModel.md) |  | [optional] 
**GatewayConnectionDetails** | Pointer to [**GatewayConnectionModel**](GatewayConnectionModel.md) |  | [optional] 
**OktaConnectionModel** | Pointer to [**OktaConnectionModel**](OktaConnectionModel.md) |  | [optional] 
**SamlConnectionModel** | Pointer to [**SamlConnectionModel**](SamlConnectionModel.md) |  | [optional] 

## Methods

### NewIdpConnectModel

`func NewIdpConnectModel(identityProviderId string, ) *IdpConnectModel`

NewIdpConnectModel instantiates a new IdpConnectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpConnectModelWithDefaults

`func NewIdpConnectModelWithDefaults() *IdpConnectModel`

NewIdpConnectModelWithDefaults instantiates a new IdpConnectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthDomainName

`func (o *IdpConnectModel) GetAuthDomainName() string`

GetAuthDomainName returns the AuthDomainName field if non-nil, zero value otherwise.

### GetAuthDomainNameOk

`func (o *IdpConnectModel) GetAuthDomainNameOk() (*string, bool)`

GetAuthDomainNameOk returns a tuple with the AuthDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomainName

`func (o *IdpConnectModel) SetAuthDomainName(v string)`

SetAuthDomainName sets AuthDomainName field to given value.

### HasAuthDomainName

`func (o *IdpConnectModel) HasAuthDomainName() bool`

HasAuthDomainName returns a boolean if a field has been set.

### SetAuthDomainNameNil

`func (o *IdpConnectModel) SetAuthDomainNameNil(b bool)`

 SetAuthDomainNameNil sets the value for AuthDomainName to be an explicit nil

### UnsetAuthDomainName
`func (o *IdpConnectModel) UnsetAuthDomainName()`

UnsetAuthDomainName ensures that no value is present for AuthDomainName, not even an explicit nil
### GetIdentityProviderId

`func (o *IdpConnectModel) GetIdentityProviderId() string`

GetIdentityProviderId returns the IdentityProviderId field if non-nil, zero value otherwise.

### GetIdentityProviderIdOk

`func (o *IdpConnectModel) GetIdentityProviderIdOk() (*string, bool)`

GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderId

`func (o *IdpConnectModel) SetIdentityProviderId(v string)`

SetIdentityProviderId sets IdentityProviderId field to given value.


### GetAdaptiveAuthenticationConnectionDetails

`func (o *IdpConnectModel) GetAdaptiveAuthenticationConnectionDetails() AdaptiveAuthenticationConnectionModel`

GetAdaptiveAuthenticationConnectionDetails returns the AdaptiveAuthenticationConnectionDetails field if non-nil, zero value otherwise.

### GetAdaptiveAuthenticationConnectionDetailsOk

`func (o *IdpConnectModel) GetAdaptiveAuthenticationConnectionDetailsOk() (*AdaptiveAuthenticationConnectionModel, bool)`

GetAdaptiveAuthenticationConnectionDetailsOk returns a tuple with the AdaptiveAuthenticationConnectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveAuthenticationConnectionDetails

`func (o *IdpConnectModel) SetAdaptiveAuthenticationConnectionDetails(v AdaptiveAuthenticationConnectionModel)`

SetAdaptiveAuthenticationConnectionDetails sets AdaptiveAuthenticationConnectionDetails field to given value.

### HasAdaptiveAuthenticationConnectionDetails

`func (o *IdpConnectModel) HasAdaptiveAuthenticationConnectionDetails() bool`

HasAdaptiveAuthenticationConnectionDetails returns a boolean if a field has been set.

### GetAdOtpDetails

`func (o *IdpConnectModel) GetAdOtpDetails() AdOtpConnectionModel`

GetAdOtpDetails returns the AdOtpDetails field if non-nil, zero value otherwise.

### GetAdOtpDetailsOk

`func (o *IdpConnectModel) GetAdOtpDetailsOk() (*AdOtpConnectionModel, bool)`

GetAdOtpDetailsOk returns a tuple with the AdOtpDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdOtpDetails

`func (o *IdpConnectModel) SetAdOtpDetails(v AdOtpConnectionModel)`

SetAdOtpDetails sets AdOtpDetails field to given value.

### HasAdOtpDetails

`func (o *IdpConnectModel) HasAdOtpDetails() bool`

HasAdOtpDetails returns a boolean if a field has been set.

### GetGoogleConnectionModel

`func (o *IdpConnectModel) GetGoogleConnectionModel() GoogleConnectionModel`

GetGoogleConnectionModel returns the GoogleConnectionModel field if non-nil, zero value otherwise.

### GetGoogleConnectionModelOk

`func (o *IdpConnectModel) GetGoogleConnectionModelOk() (*GoogleConnectionModel, bool)`

GetGoogleConnectionModelOk returns a tuple with the GoogleConnectionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleConnectionModel

`func (o *IdpConnectModel) SetGoogleConnectionModel(v GoogleConnectionModel)`

SetGoogleConnectionModel sets GoogleConnectionModel field to given value.

### HasGoogleConnectionModel

`func (o *IdpConnectModel) HasGoogleConnectionModel() bool`

HasGoogleConnectionModel returns a boolean if a field has been set.

### GetGatewayConnectionDetails

`func (o *IdpConnectModel) GetGatewayConnectionDetails() GatewayConnectionModel`

GetGatewayConnectionDetails returns the GatewayConnectionDetails field if non-nil, zero value otherwise.

### GetGatewayConnectionDetailsOk

`func (o *IdpConnectModel) GetGatewayConnectionDetailsOk() (*GatewayConnectionModel, bool)`

GetGatewayConnectionDetailsOk returns a tuple with the GatewayConnectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConnectionDetails

`func (o *IdpConnectModel) SetGatewayConnectionDetails(v GatewayConnectionModel)`

SetGatewayConnectionDetails sets GatewayConnectionDetails field to given value.

### HasGatewayConnectionDetails

`func (o *IdpConnectModel) HasGatewayConnectionDetails() bool`

HasGatewayConnectionDetails returns a boolean if a field has been set.

### GetOktaConnectionModel

`func (o *IdpConnectModel) GetOktaConnectionModel() OktaConnectionModel`

GetOktaConnectionModel returns the OktaConnectionModel field if non-nil, zero value otherwise.

### GetOktaConnectionModelOk

`func (o *IdpConnectModel) GetOktaConnectionModelOk() (*OktaConnectionModel, bool)`

GetOktaConnectionModelOk returns a tuple with the OktaConnectionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaConnectionModel

`func (o *IdpConnectModel) SetOktaConnectionModel(v OktaConnectionModel)`

SetOktaConnectionModel sets OktaConnectionModel field to given value.

### HasOktaConnectionModel

`func (o *IdpConnectModel) HasOktaConnectionModel() bool`

HasOktaConnectionModel returns a boolean if a field has been set.

### GetSamlConnectionModel

`func (o *IdpConnectModel) GetSamlConnectionModel() SamlConnectionModel`

GetSamlConnectionModel returns the SamlConnectionModel field if non-nil, zero value otherwise.

### GetSamlConnectionModelOk

`func (o *IdpConnectModel) GetSamlConnectionModelOk() (*SamlConnectionModel, bool)`

GetSamlConnectionModelOk returns a tuple with the SamlConnectionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlConnectionModel

`func (o *IdpConnectModel) SetSamlConnectionModel(v SamlConnectionModel)`

SetSamlConnectionModel sets SamlConnectionModel field to given value.

### HasSamlConnectionModel

`func (o *IdpConnectModel) HasSamlConnectionModel() bool`

HasSamlConnectionModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


