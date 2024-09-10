# IdpConnectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpType** | Pointer to **NullableString** |  | [optional] 
**ConnectedApp** | Pointer to **NullableString** |  | [optional] 
**ConnectionStatus** | Pointer to **bool** |  | [optional] 
**IdpProperties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewIdpConnectionModel

`func NewIdpConnectionModel() *IdpConnectionModel`

NewIdpConnectionModel instantiates a new IdpConnectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpConnectionModelWithDefaults

`func NewIdpConnectionModelWithDefaults() *IdpConnectionModel`

NewIdpConnectionModelWithDefaults instantiates a new IdpConnectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpType

`func (o *IdpConnectionModel) GetIdpType() string`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *IdpConnectionModel) GetIdpTypeOk() (*string, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *IdpConnectionModel) SetIdpType(v string)`

SetIdpType sets IdpType field to given value.

### HasIdpType

`func (o *IdpConnectionModel) HasIdpType() bool`

HasIdpType returns a boolean if a field has been set.

### SetIdpTypeNil

`func (o *IdpConnectionModel) SetIdpTypeNil(b bool)`

 SetIdpTypeNil sets the value for IdpType to be an explicit nil

### UnsetIdpType
`func (o *IdpConnectionModel) UnsetIdpType()`

UnsetIdpType ensures that no value is present for IdpType, not even an explicit nil
### GetConnectedApp

`func (o *IdpConnectionModel) GetConnectedApp() string`

GetConnectedApp returns the ConnectedApp field if non-nil, zero value otherwise.

### GetConnectedAppOk

`func (o *IdpConnectionModel) GetConnectedAppOk() (*string, bool)`

GetConnectedAppOk returns a tuple with the ConnectedApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedApp

`func (o *IdpConnectionModel) SetConnectedApp(v string)`

SetConnectedApp sets ConnectedApp field to given value.

### HasConnectedApp

`func (o *IdpConnectionModel) HasConnectedApp() bool`

HasConnectedApp returns a boolean if a field has been set.

### SetConnectedAppNil

`func (o *IdpConnectionModel) SetConnectedAppNil(b bool)`

 SetConnectedAppNil sets the value for ConnectedApp to be an explicit nil

### UnsetConnectedApp
`func (o *IdpConnectionModel) UnsetConnectedApp()`

UnsetConnectedApp ensures that no value is present for ConnectedApp, not even an explicit nil
### GetConnectionStatus

`func (o *IdpConnectionModel) GetConnectionStatus() bool`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *IdpConnectionModel) GetConnectionStatusOk() (*bool, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *IdpConnectionModel) SetConnectionStatus(v bool)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *IdpConnectionModel) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetIdpProperties

`func (o *IdpConnectionModel) GetIdpProperties() map[string]string`

GetIdpProperties returns the IdpProperties field if non-nil, zero value otherwise.

### GetIdpPropertiesOk

`func (o *IdpConnectionModel) GetIdpPropertiesOk() (*map[string]string, bool)`

GetIdpPropertiesOk returns a tuple with the IdpProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpProperties

`func (o *IdpConnectionModel) SetIdpProperties(v map[string]string)`

SetIdpProperties sets IdpProperties field to given value.

### HasIdpProperties

`func (o *IdpConnectionModel) HasIdpProperties() bool`

HasIdpProperties returns a boolean if a field has been set.

### SetIdpPropertiesNil

`func (o *IdpConnectionModel) SetIdpPropertiesNil(b bool)`

 SetIdpPropertiesNil sets the value for IdpProperties to be an explicit nil

### UnsetIdpProperties
`func (o *IdpConnectionModel) UnsetIdpProperties()`

UnsetIdpProperties ensures that no value is present for IdpProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


