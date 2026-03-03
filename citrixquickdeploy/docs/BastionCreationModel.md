# BastionCreationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**IpAddressPrefix** | Pointer to **NullableString** |  | [optional] 
**ConnectionType** | Pointer to [**NullableOnPremConnectionType**](OnPremConnectionType.md) |  | [optional] 

## Methods

### NewBastionCreationModel

`func NewBastionCreationModel() *BastionCreationModel`

NewBastionCreationModel instantiates a new BastionCreationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBastionCreationModelWithDefaults

`func NewBastionCreationModelWithDefaults() *BastionCreationModel`

NewBastionCreationModelWithDefaults instantiates a new BastionCreationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *BastionCreationModel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BastionCreationModel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BastionCreationModel) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BastionCreationModel) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *BastionCreationModel) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *BastionCreationModel) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *BastionCreationModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BastionCreationModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BastionCreationModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BastionCreationModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *BastionCreationModel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *BastionCreationModel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetIpAddressPrefix

`func (o *BastionCreationModel) GetIpAddressPrefix() string`

GetIpAddressPrefix returns the IpAddressPrefix field if non-nil, zero value otherwise.

### GetIpAddressPrefixOk

`func (o *BastionCreationModel) GetIpAddressPrefixOk() (*string, bool)`

GetIpAddressPrefixOk returns a tuple with the IpAddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressPrefix

`func (o *BastionCreationModel) SetIpAddressPrefix(v string)`

SetIpAddressPrefix sets IpAddressPrefix field to given value.

### HasIpAddressPrefix

`func (o *BastionCreationModel) HasIpAddressPrefix() bool`

HasIpAddressPrefix returns a boolean if a field has been set.

### SetIpAddressPrefixNil

`func (o *BastionCreationModel) SetIpAddressPrefixNil(b bool)`

 SetIpAddressPrefixNil sets the value for IpAddressPrefix to be an explicit nil

### UnsetIpAddressPrefix
`func (o *BastionCreationModel) UnsetIpAddressPrefix()`

UnsetIpAddressPrefix ensures that no value is present for IpAddressPrefix, not even an explicit nil
### GetConnectionType

`func (o *BastionCreationModel) GetConnectionType() OnPremConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *BastionCreationModel) GetConnectionTypeOk() (*OnPremConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *BastionCreationModel) SetConnectionType(v OnPremConnectionType)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *BastionCreationModel) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### SetConnectionTypeNil

`func (o *BastionCreationModel) SetConnectionTypeNil(b bool)`

 SetConnectionTypeNil sets the value for ConnectionType to be an explicit nil

### UnsetConnectionType
`func (o *BastionCreationModel) UnsetConnectionType()`

UnsetConnectionType ensures that no value is present for ConnectionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


