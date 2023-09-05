# LicenseResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductCode** | **string** | The product code. | 
**ProductEdition** | **string** | The product edition. | 
**LicenseModel** | [**LicenseModel**](LicenseModel.md) |  | 

## Methods

### NewLicenseResponseModel

`func NewLicenseResponseModel(productCode string, productEdition string, licenseModel LicenseModel, ) *LicenseResponseModel`

NewLicenseResponseModel instantiates a new LicenseResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseResponseModelWithDefaults

`func NewLicenseResponseModelWithDefaults() *LicenseResponseModel`

NewLicenseResponseModelWithDefaults instantiates a new LicenseResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductCode

`func (o *LicenseResponseModel) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *LicenseResponseModel) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *LicenseResponseModel) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetProductEdition

`func (o *LicenseResponseModel) GetProductEdition() string`

GetProductEdition returns the ProductEdition field if non-nil, zero value otherwise.

### GetProductEditionOk

`func (o *LicenseResponseModel) GetProductEditionOk() (*string, bool)`

GetProductEditionOk returns a tuple with the ProductEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductEdition

`func (o *LicenseResponseModel) SetProductEdition(v string)`

SetProductEdition sets ProductEdition field to given value.


### GetLicenseModel

`func (o *LicenseResponseModel) GetLicenseModel() LicenseModel`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *LicenseResponseModel) GetLicenseModelOk() (*LicenseModel, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *LicenseResponseModel) SetLicenseModel(v LicenseModel)`

SetLicenseModel sets LicenseModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


