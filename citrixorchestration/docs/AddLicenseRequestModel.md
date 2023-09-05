# AddLicenseRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseFile** | **string** | Encoded Base64 for license file, and the format should be utf-8. | 
**FileName** | **string** | Licence file name to be added. | 
**OverrideFile** | **bool** | Indicate if override exist license file when the license file already exists. | 

## Methods

### NewAddLicenseRequestModel

`func NewAddLicenseRequestModel(licenseFile string, fileName string, overrideFile bool, ) *AddLicenseRequestModel`

NewAddLicenseRequestModel instantiates a new AddLicenseRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLicenseRequestModelWithDefaults

`func NewAddLicenseRequestModelWithDefaults() *AddLicenseRequestModel`

NewAddLicenseRequestModelWithDefaults instantiates a new AddLicenseRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseFile

`func (o *AddLicenseRequestModel) GetLicenseFile() string`

GetLicenseFile returns the LicenseFile field if non-nil, zero value otherwise.

### GetLicenseFileOk

`func (o *AddLicenseRequestModel) GetLicenseFileOk() (*string, bool)`

GetLicenseFileOk returns a tuple with the LicenseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFile

`func (o *AddLicenseRequestModel) SetLicenseFile(v string)`

SetLicenseFile sets LicenseFile field to given value.


### GetFileName

`func (o *AddLicenseRequestModel) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AddLicenseRequestModel) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AddLicenseRequestModel) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetOverrideFile

`func (o *AddLicenseRequestModel) GetOverrideFile() bool`

GetOverrideFile returns the OverrideFile field if non-nil, zero value otherwise.

### GetOverrideFileOk

`func (o *AddLicenseRequestModel) GetOverrideFileOk() (*bool, bool)`

GetOverrideFileOk returns a tuple with the OverrideFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideFile

`func (o *AddLicenseRequestModel) SetOverrideFile(v bool)`

SetOverrideFile sets OverrideFile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


