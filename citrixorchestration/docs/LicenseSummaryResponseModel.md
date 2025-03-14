# LicenseSummaryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseInformation** | Pointer to [**LicenseInformationModel**](LicenseInformationModel.md) |  | [optional] 
**LicenseInventories** | Pointer to [**[]LicenseInventoryModel**](LicenseInventoryModel.md) | Get license inventory list | [optional] 
**LicenseLasInventories** | Pointer to [**[]LicenseInventoryModel**](LicenseInventoryModel.md) | Get license Las inventory list | [optional] 
**NormalLicenses** | Pointer to **NullableInt32** | The number of the available licenses | [optional] 
**OverdraftLicenses** | Pointer to **NullableInt32** | The number of the overdraft licenses | [optional] 
**LicensesUsage** | Pointer to **NullableInt32** | The usage of the licenses | [optional] 
**CertAccepted** | Pointer to **NullableBool** | Indicates if the certificate is verified | [optional] 

## Methods

### NewLicenseSummaryResponseModel

`func NewLicenseSummaryResponseModel() *LicenseSummaryResponseModel`

NewLicenseSummaryResponseModel instantiates a new LicenseSummaryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseSummaryResponseModelWithDefaults

`func NewLicenseSummaryResponseModelWithDefaults() *LicenseSummaryResponseModel`

NewLicenseSummaryResponseModelWithDefaults instantiates a new LicenseSummaryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseInformation

`func (o *LicenseSummaryResponseModel) GetLicenseInformation() LicenseInformationModel`

GetLicenseInformation returns the LicenseInformation field if non-nil, zero value otherwise.

### GetLicenseInformationOk

`func (o *LicenseSummaryResponseModel) GetLicenseInformationOk() (*LicenseInformationModel, bool)`

GetLicenseInformationOk returns a tuple with the LicenseInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseInformation

`func (o *LicenseSummaryResponseModel) SetLicenseInformation(v LicenseInformationModel)`

SetLicenseInformation sets LicenseInformation field to given value.

### HasLicenseInformation

`func (o *LicenseSummaryResponseModel) HasLicenseInformation() bool`

HasLicenseInformation returns a boolean if a field has been set.

### GetLicenseInventories

`func (o *LicenseSummaryResponseModel) GetLicenseInventories() []LicenseInventoryModel`

GetLicenseInventories returns the LicenseInventories field if non-nil, zero value otherwise.

### GetLicenseInventoriesOk

`func (o *LicenseSummaryResponseModel) GetLicenseInventoriesOk() (*[]LicenseInventoryModel, bool)`

GetLicenseInventoriesOk returns a tuple with the LicenseInventories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseInventories

`func (o *LicenseSummaryResponseModel) SetLicenseInventories(v []LicenseInventoryModel)`

SetLicenseInventories sets LicenseInventories field to given value.

### HasLicenseInventories

`func (o *LicenseSummaryResponseModel) HasLicenseInventories() bool`

HasLicenseInventories returns a boolean if a field has been set.

### SetLicenseInventoriesNil

`func (o *LicenseSummaryResponseModel) SetLicenseInventoriesNil(b bool)`

 SetLicenseInventoriesNil sets the value for LicenseInventories to be an explicit nil

### UnsetLicenseInventories
`func (o *LicenseSummaryResponseModel) UnsetLicenseInventories()`

UnsetLicenseInventories ensures that no value is present for LicenseInventories, not even an explicit nil
### GetLicenseLasInventories

`func (o *LicenseSummaryResponseModel) GetLicenseLasInventories() []LicenseInventoryModel`

GetLicenseLasInventories returns the LicenseLasInventories field if non-nil, zero value otherwise.

### GetLicenseLasInventoriesOk

`func (o *LicenseSummaryResponseModel) GetLicenseLasInventoriesOk() (*[]LicenseInventoryModel, bool)`

GetLicenseLasInventoriesOk returns a tuple with the LicenseLasInventories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseLasInventories

`func (o *LicenseSummaryResponseModel) SetLicenseLasInventories(v []LicenseInventoryModel)`

SetLicenseLasInventories sets LicenseLasInventories field to given value.

### HasLicenseLasInventories

`func (o *LicenseSummaryResponseModel) HasLicenseLasInventories() bool`

HasLicenseLasInventories returns a boolean if a field has been set.

### SetLicenseLasInventoriesNil

`func (o *LicenseSummaryResponseModel) SetLicenseLasInventoriesNil(b bool)`

 SetLicenseLasInventoriesNil sets the value for LicenseLasInventories to be an explicit nil

### UnsetLicenseLasInventories
`func (o *LicenseSummaryResponseModel) UnsetLicenseLasInventories()`

UnsetLicenseLasInventories ensures that no value is present for LicenseLasInventories, not even an explicit nil
### GetNormalLicenses

`func (o *LicenseSummaryResponseModel) GetNormalLicenses() int32`

GetNormalLicenses returns the NormalLicenses field if non-nil, zero value otherwise.

### GetNormalLicensesOk

`func (o *LicenseSummaryResponseModel) GetNormalLicensesOk() (*int32, bool)`

GetNormalLicensesOk returns a tuple with the NormalLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalLicenses

`func (o *LicenseSummaryResponseModel) SetNormalLicenses(v int32)`

SetNormalLicenses sets NormalLicenses field to given value.

### HasNormalLicenses

`func (o *LicenseSummaryResponseModel) HasNormalLicenses() bool`

HasNormalLicenses returns a boolean if a field has been set.

### SetNormalLicensesNil

`func (o *LicenseSummaryResponseModel) SetNormalLicensesNil(b bool)`

 SetNormalLicensesNil sets the value for NormalLicenses to be an explicit nil

### UnsetNormalLicenses
`func (o *LicenseSummaryResponseModel) UnsetNormalLicenses()`

UnsetNormalLicenses ensures that no value is present for NormalLicenses, not even an explicit nil
### GetOverdraftLicenses

`func (o *LicenseSummaryResponseModel) GetOverdraftLicenses() int32`

GetOverdraftLicenses returns the OverdraftLicenses field if non-nil, zero value otherwise.

### GetOverdraftLicensesOk

`func (o *LicenseSummaryResponseModel) GetOverdraftLicensesOk() (*int32, bool)`

GetOverdraftLicensesOk returns a tuple with the OverdraftLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftLicenses

`func (o *LicenseSummaryResponseModel) SetOverdraftLicenses(v int32)`

SetOverdraftLicenses sets OverdraftLicenses field to given value.

### HasOverdraftLicenses

`func (o *LicenseSummaryResponseModel) HasOverdraftLicenses() bool`

HasOverdraftLicenses returns a boolean if a field has been set.

### SetOverdraftLicensesNil

`func (o *LicenseSummaryResponseModel) SetOverdraftLicensesNil(b bool)`

 SetOverdraftLicensesNil sets the value for OverdraftLicenses to be an explicit nil

### UnsetOverdraftLicenses
`func (o *LicenseSummaryResponseModel) UnsetOverdraftLicenses()`

UnsetOverdraftLicenses ensures that no value is present for OverdraftLicenses, not even an explicit nil
### GetLicensesUsage

`func (o *LicenseSummaryResponseModel) GetLicensesUsage() int32`

GetLicensesUsage returns the LicensesUsage field if non-nil, zero value otherwise.

### GetLicensesUsageOk

`func (o *LicenseSummaryResponseModel) GetLicensesUsageOk() (*int32, bool)`

GetLicensesUsageOk returns a tuple with the LicensesUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensesUsage

`func (o *LicenseSummaryResponseModel) SetLicensesUsage(v int32)`

SetLicensesUsage sets LicensesUsage field to given value.

### HasLicensesUsage

`func (o *LicenseSummaryResponseModel) HasLicensesUsage() bool`

HasLicensesUsage returns a boolean if a field has been set.

### SetLicensesUsageNil

`func (o *LicenseSummaryResponseModel) SetLicensesUsageNil(b bool)`

 SetLicensesUsageNil sets the value for LicensesUsage to be an explicit nil

### UnsetLicensesUsage
`func (o *LicenseSummaryResponseModel) UnsetLicensesUsage()`

UnsetLicensesUsage ensures that no value is present for LicensesUsage, not even an explicit nil
### GetCertAccepted

`func (o *LicenseSummaryResponseModel) GetCertAccepted() bool`

GetCertAccepted returns the CertAccepted field if non-nil, zero value otherwise.

### GetCertAcceptedOk

`func (o *LicenseSummaryResponseModel) GetCertAcceptedOk() (*bool, bool)`

GetCertAcceptedOk returns a tuple with the CertAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAccepted

`func (o *LicenseSummaryResponseModel) SetCertAccepted(v bool)`

SetCertAccepted sets CertAccepted field to given value.

### HasCertAccepted

`func (o *LicenseSummaryResponseModel) HasCertAccepted() bool`

HasCertAccepted returns a boolean if a field has been set.

### SetCertAcceptedNil

`func (o *LicenseSummaryResponseModel) SetCertAcceptedNil(b bool)`

 SetCertAcceptedNil sets the value for CertAccepted to be an explicit nil

### UnsetCertAccepted
`func (o *LicenseSummaryResponseModel) UnsetCertAccepted()`

UnsetCertAccepted ensures that no value is present for CertAccepted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


