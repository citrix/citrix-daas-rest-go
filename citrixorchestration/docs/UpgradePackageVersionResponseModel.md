# UpgradePackageVersionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeType** | [**VdaUpgradeType**](VdaUpgradeType.md) |  | 
**UpgradePackageVersion** | **string** | Latest released package version for the upgrade type. | 

## Methods

### NewUpgradePackageVersionResponseModel

`func NewUpgradePackageVersionResponseModel(upgradeType VdaUpgradeType, upgradePackageVersion string, ) *UpgradePackageVersionResponseModel`

NewUpgradePackageVersionResponseModel instantiates a new UpgradePackageVersionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradePackageVersionResponseModelWithDefaults

`func NewUpgradePackageVersionResponseModelWithDefaults() *UpgradePackageVersionResponseModel`

NewUpgradePackageVersionResponseModelWithDefaults instantiates a new UpgradePackageVersionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeType

`func (o *UpgradePackageVersionResponseModel) GetUpgradeType() VdaUpgradeType`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### GetUpgradeTypeOk

`func (o *UpgradePackageVersionResponseModel) GetUpgradeTypeOk() (*VdaUpgradeType, bool)`

GetUpgradeTypeOk returns a tuple with the UpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeType

`func (o *UpgradePackageVersionResponseModel) SetUpgradeType(v VdaUpgradeType)`

SetUpgradeType sets UpgradeType field to given value.


### GetUpgradePackageVersion

`func (o *UpgradePackageVersionResponseModel) GetUpgradePackageVersion() string`

GetUpgradePackageVersion returns the UpgradePackageVersion field if non-nil, zero value otherwise.

### GetUpgradePackageVersionOk

`func (o *UpgradePackageVersionResponseModel) GetUpgradePackageVersionOk() (*string, bool)`

GetUpgradePackageVersionOk returns a tuple with the UpgradePackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePackageVersion

`func (o *UpgradePackageVersionResponseModel) SetUpgradePackageVersion(v string)`

SetUpgradePackageVersion sets UpgradePackageVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


