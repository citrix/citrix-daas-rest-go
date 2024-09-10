# AuthDomainModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to **map[string]interface{}** |  | [optional] 
**IdpDisplayName** | Pointer to **NullableString** |  | [optional] 
**IdpProperties** | Pointer to **map[string]string** |  | [optional] 
**IdpType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Product** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAuthDomainModel

`func NewAuthDomainModel() *AuthDomainModel`

NewAuthDomainModel instantiates a new AuthDomainModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthDomainModelWithDefaults

`func NewAuthDomainModelWithDefaults() *AuthDomainModel`

NewAuthDomainModelWithDefaults instantiates a new AuthDomainModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *AuthDomainModel) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *AuthDomainModel) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *AuthDomainModel) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *AuthDomainModel) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *AuthDomainModel) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *AuthDomainModel) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetIdpDisplayName

`func (o *AuthDomainModel) GetIdpDisplayName() string`

GetIdpDisplayName returns the IdpDisplayName field if non-nil, zero value otherwise.

### GetIdpDisplayNameOk

`func (o *AuthDomainModel) GetIdpDisplayNameOk() (*string, bool)`

GetIdpDisplayNameOk returns a tuple with the IdpDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpDisplayName

`func (o *AuthDomainModel) SetIdpDisplayName(v string)`

SetIdpDisplayName sets IdpDisplayName field to given value.

### HasIdpDisplayName

`func (o *AuthDomainModel) HasIdpDisplayName() bool`

HasIdpDisplayName returns a boolean if a field has been set.

### SetIdpDisplayNameNil

`func (o *AuthDomainModel) SetIdpDisplayNameNil(b bool)`

 SetIdpDisplayNameNil sets the value for IdpDisplayName to be an explicit nil

### UnsetIdpDisplayName
`func (o *AuthDomainModel) UnsetIdpDisplayName()`

UnsetIdpDisplayName ensures that no value is present for IdpDisplayName, not even an explicit nil
### GetIdpProperties

`func (o *AuthDomainModel) GetIdpProperties() map[string]string`

GetIdpProperties returns the IdpProperties field if non-nil, zero value otherwise.

### GetIdpPropertiesOk

`func (o *AuthDomainModel) GetIdpPropertiesOk() (*map[string]string, bool)`

GetIdpPropertiesOk returns a tuple with the IdpProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpProperties

`func (o *AuthDomainModel) SetIdpProperties(v map[string]string)`

SetIdpProperties sets IdpProperties field to given value.

### HasIdpProperties

`func (o *AuthDomainModel) HasIdpProperties() bool`

HasIdpProperties returns a boolean if a field has been set.

### SetIdpPropertiesNil

`func (o *AuthDomainModel) SetIdpPropertiesNil(b bool)`

 SetIdpPropertiesNil sets the value for IdpProperties to be an explicit nil

### UnsetIdpProperties
`func (o *AuthDomainModel) UnsetIdpProperties()`

UnsetIdpProperties ensures that no value is present for IdpProperties, not even an explicit nil
### GetIdpType

`func (o *AuthDomainModel) GetIdpType() string`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *AuthDomainModel) GetIdpTypeOk() (*string, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *AuthDomainModel) SetIdpType(v string)`

SetIdpType sets IdpType field to given value.

### HasIdpType

`func (o *AuthDomainModel) HasIdpType() bool`

HasIdpType returns a boolean if a field has been set.

### SetIdpTypeNil

`func (o *AuthDomainModel) SetIdpTypeNil(b bool)`

 SetIdpTypeNil sets the value for IdpType to be an explicit nil

### UnsetIdpType
`func (o *AuthDomainModel) UnsetIdpType()`

UnsetIdpType ensures that no value is present for IdpType, not even an explicit nil
### GetName

`func (o *AuthDomainModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthDomainModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthDomainModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthDomainModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AuthDomainModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AuthDomainModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProduct

`func (o *AuthDomainModel) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AuthDomainModel) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AuthDomainModel) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AuthDomainModel) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *AuthDomainModel) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *AuthDomainModel) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


