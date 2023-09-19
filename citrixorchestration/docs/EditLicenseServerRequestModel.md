# EditLicenseServerRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseServer** | **string** | License server address | 
**Thumbprint** | Pointer to **NullableString** | The certificate thumbprint of the license server | [optional] 

## Methods

### NewEditLicenseServerRequestModel

`func NewEditLicenseServerRequestModel(licenseServer string, ) *EditLicenseServerRequestModel`

NewEditLicenseServerRequestModel instantiates a new EditLicenseServerRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditLicenseServerRequestModelWithDefaults

`func NewEditLicenseServerRequestModelWithDefaults() *EditLicenseServerRequestModel`

NewEditLicenseServerRequestModelWithDefaults instantiates a new EditLicenseServerRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseServer

`func (o *EditLicenseServerRequestModel) GetLicenseServer() string`

GetLicenseServer returns the LicenseServer field if non-nil, zero value otherwise.

### GetLicenseServerOk

`func (o *EditLicenseServerRequestModel) GetLicenseServerOk() (*string, bool)`

GetLicenseServerOk returns a tuple with the LicenseServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServer

`func (o *EditLicenseServerRequestModel) SetLicenseServer(v string)`

SetLicenseServer sets LicenseServer field to given value.


### GetThumbprint

`func (o *EditLicenseServerRequestModel) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *EditLicenseServerRequestModel) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *EditLicenseServerRequestModel) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *EditLicenseServerRequestModel) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *EditLicenseServerRequestModel) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *EditLicenseServerRequestModel) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


