# LicenseInventoryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseProductName** | Pointer to [**LicenseProduct**](LicenseProduct.md) |  | [optional] 
**LocalizedLicenseProductName** | Pointer to **NullableString** | Localized license name of the product | [optional] 
**LicenseExpirationDate** | Pointer to **NullableString** | The expired date of the license | [optional] 
**LicenseSubscriptionAdvantageDate** | Pointer to **NullableString** | The subscription advantage date of the license | [optional] 
**LicenseType** | Pointer to **NullableString** | The type of the license | [optional] 
**LicensesInUse** | Pointer to **int32** | The number of the Licenses in use | [optional] 
**LicensesAvailable** | Pointer to **int32** | The number of the available licenses | [optional] 
**LicenseOverdraft** | Pointer to **int32** | The number of the overdraft licenses | [optional] 
**LicenseModel** | Pointer to [**LicenseModel**](LicenseModel.md) |  | [optional] 
**ProductEdition** | Pointer to [**ProductEdition**](ProductEdition.md) |  | [optional] 
**LocalizedLicenseProductEdition** | Pointer to **NullableString** | Localized license name of the product edition | [optional] 

## Methods

### NewLicenseInventoryModel

`func NewLicenseInventoryModel() *LicenseInventoryModel`

NewLicenseInventoryModel instantiates a new LicenseInventoryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseInventoryModelWithDefaults

`func NewLicenseInventoryModelWithDefaults() *LicenseInventoryModel`

NewLicenseInventoryModelWithDefaults instantiates a new LicenseInventoryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseProductName

`func (o *LicenseInventoryModel) GetLicenseProductName() LicenseProduct`

GetLicenseProductName returns the LicenseProductName field if non-nil, zero value otherwise.

### GetLicenseProductNameOk

`func (o *LicenseInventoryModel) GetLicenseProductNameOk() (*LicenseProduct, bool)`

GetLicenseProductNameOk returns a tuple with the LicenseProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseProductName

`func (o *LicenseInventoryModel) SetLicenseProductName(v LicenseProduct)`

SetLicenseProductName sets LicenseProductName field to given value.

### HasLicenseProductName

`func (o *LicenseInventoryModel) HasLicenseProductName() bool`

HasLicenseProductName returns a boolean if a field has been set.

### GetLocalizedLicenseProductName

`func (o *LicenseInventoryModel) GetLocalizedLicenseProductName() string`

GetLocalizedLicenseProductName returns the LocalizedLicenseProductName field if non-nil, zero value otherwise.

### GetLocalizedLicenseProductNameOk

`func (o *LicenseInventoryModel) GetLocalizedLicenseProductNameOk() (*string, bool)`

GetLocalizedLicenseProductNameOk returns a tuple with the LocalizedLicenseProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedLicenseProductName

`func (o *LicenseInventoryModel) SetLocalizedLicenseProductName(v string)`

SetLocalizedLicenseProductName sets LocalizedLicenseProductName field to given value.

### HasLocalizedLicenseProductName

`func (o *LicenseInventoryModel) HasLocalizedLicenseProductName() bool`

HasLocalizedLicenseProductName returns a boolean if a field has been set.

### SetLocalizedLicenseProductNameNil

`func (o *LicenseInventoryModel) SetLocalizedLicenseProductNameNil(b bool)`

 SetLocalizedLicenseProductNameNil sets the value for LocalizedLicenseProductName to be an explicit nil

### UnsetLocalizedLicenseProductName
`func (o *LicenseInventoryModel) UnsetLocalizedLicenseProductName()`

UnsetLocalizedLicenseProductName ensures that no value is present for LocalizedLicenseProductName, not even an explicit nil
### GetLicenseExpirationDate

`func (o *LicenseInventoryModel) GetLicenseExpirationDate() string`

GetLicenseExpirationDate returns the LicenseExpirationDate field if non-nil, zero value otherwise.

### GetLicenseExpirationDateOk

`func (o *LicenseInventoryModel) GetLicenseExpirationDateOk() (*string, bool)`

GetLicenseExpirationDateOk returns a tuple with the LicenseExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpirationDate

`func (o *LicenseInventoryModel) SetLicenseExpirationDate(v string)`

SetLicenseExpirationDate sets LicenseExpirationDate field to given value.

### HasLicenseExpirationDate

`func (o *LicenseInventoryModel) HasLicenseExpirationDate() bool`

HasLicenseExpirationDate returns a boolean if a field has been set.

### SetLicenseExpirationDateNil

`func (o *LicenseInventoryModel) SetLicenseExpirationDateNil(b bool)`

 SetLicenseExpirationDateNil sets the value for LicenseExpirationDate to be an explicit nil

### UnsetLicenseExpirationDate
`func (o *LicenseInventoryModel) UnsetLicenseExpirationDate()`

UnsetLicenseExpirationDate ensures that no value is present for LicenseExpirationDate, not even an explicit nil
### GetLicenseSubscriptionAdvantageDate

`func (o *LicenseInventoryModel) GetLicenseSubscriptionAdvantageDate() string`

GetLicenseSubscriptionAdvantageDate returns the LicenseSubscriptionAdvantageDate field if non-nil, zero value otherwise.

### GetLicenseSubscriptionAdvantageDateOk

`func (o *LicenseInventoryModel) GetLicenseSubscriptionAdvantageDateOk() (*string, bool)`

GetLicenseSubscriptionAdvantageDateOk returns a tuple with the LicenseSubscriptionAdvantageDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseSubscriptionAdvantageDate

`func (o *LicenseInventoryModel) SetLicenseSubscriptionAdvantageDate(v string)`

SetLicenseSubscriptionAdvantageDate sets LicenseSubscriptionAdvantageDate field to given value.

### HasLicenseSubscriptionAdvantageDate

`func (o *LicenseInventoryModel) HasLicenseSubscriptionAdvantageDate() bool`

HasLicenseSubscriptionAdvantageDate returns a boolean if a field has been set.

### SetLicenseSubscriptionAdvantageDateNil

`func (o *LicenseInventoryModel) SetLicenseSubscriptionAdvantageDateNil(b bool)`

 SetLicenseSubscriptionAdvantageDateNil sets the value for LicenseSubscriptionAdvantageDate to be an explicit nil

### UnsetLicenseSubscriptionAdvantageDate
`func (o *LicenseInventoryModel) UnsetLicenseSubscriptionAdvantageDate()`

UnsetLicenseSubscriptionAdvantageDate ensures that no value is present for LicenseSubscriptionAdvantageDate, not even an explicit nil
### GetLicenseType

`func (o *LicenseInventoryModel) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *LicenseInventoryModel) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *LicenseInventoryModel) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *LicenseInventoryModel) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### SetLicenseTypeNil

`func (o *LicenseInventoryModel) SetLicenseTypeNil(b bool)`

 SetLicenseTypeNil sets the value for LicenseType to be an explicit nil

### UnsetLicenseType
`func (o *LicenseInventoryModel) UnsetLicenseType()`

UnsetLicenseType ensures that no value is present for LicenseType, not even an explicit nil
### GetLicensesInUse

`func (o *LicenseInventoryModel) GetLicensesInUse() int32`

GetLicensesInUse returns the LicensesInUse field if non-nil, zero value otherwise.

### GetLicensesInUseOk

`func (o *LicenseInventoryModel) GetLicensesInUseOk() (*int32, bool)`

GetLicensesInUseOk returns a tuple with the LicensesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensesInUse

`func (o *LicenseInventoryModel) SetLicensesInUse(v int32)`

SetLicensesInUse sets LicensesInUse field to given value.

### HasLicensesInUse

`func (o *LicenseInventoryModel) HasLicensesInUse() bool`

HasLicensesInUse returns a boolean if a field has been set.

### GetLicensesAvailable

`func (o *LicenseInventoryModel) GetLicensesAvailable() int32`

GetLicensesAvailable returns the LicensesAvailable field if non-nil, zero value otherwise.

### GetLicensesAvailableOk

`func (o *LicenseInventoryModel) GetLicensesAvailableOk() (*int32, bool)`

GetLicensesAvailableOk returns a tuple with the LicensesAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensesAvailable

`func (o *LicenseInventoryModel) SetLicensesAvailable(v int32)`

SetLicensesAvailable sets LicensesAvailable field to given value.

### HasLicensesAvailable

`func (o *LicenseInventoryModel) HasLicensesAvailable() bool`

HasLicensesAvailable returns a boolean if a field has been set.

### GetLicenseOverdraft

`func (o *LicenseInventoryModel) GetLicenseOverdraft() int32`

GetLicenseOverdraft returns the LicenseOverdraft field if non-nil, zero value otherwise.

### GetLicenseOverdraftOk

`func (o *LicenseInventoryModel) GetLicenseOverdraftOk() (*int32, bool)`

GetLicenseOverdraftOk returns a tuple with the LicenseOverdraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseOverdraft

`func (o *LicenseInventoryModel) SetLicenseOverdraft(v int32)`

SetLicenseOverdraft sets LicenseOverdraft field to given value.

### HasLicenseOverdraft

`func (o *LicenseInventoryModel) HasLicenseOverdraft() bool`

HasLicenseOverdraft returns a boolean if a field has been set.

### GetLicenseModel

`func (o *LicenseInventoryModel) GetLicenseModel() LicenseModel`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *LicenseInventoryModel) GetLicenseModelOk() (*LicenseModel, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *LicenseInventoryModel) SetLicenseModel(v LicenseModel)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *LicenseInventoryModel) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetProductEdition

`func (o *LicenseInventoryModel) GetProductEdition() ProductEdition`

GetProductEdition returns the ProductEdition field if non-nil, zero value otherwise.

### GetProductEditionOk

`func (o *LicenseInventoryModel) GetProductEditionOk() (*ProductEdition, bool)`

GetProductEditionOk returns a tuple with the ProductEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductEdition

`func (o *LicenseInventoryModel) SetProductEdition(v ProductEdition)`

SetProductEdition sets ProductEdition field to given value.

### HasProductEdition

`func (o *LicenseInventoryModel) HasProductEdition() bool`

HasProductEdition returns a boolean if a field has been set.

### GetLocalizedLicenseProductEdition

`func (o *LicenseInventoryModel) GetLocalizedLicenseProductEdition() string`

GetLocalizedLicenseProductEdition returns the LocalizedLicenseProductEdition field if non-nil, zero value otherwise.

### GetLocalizedLicenseProductEditionOk

`func (o *LicenseInventoryModel) GetLocalizedLicenseProductEditionOk() (*string, bool)`

GetLocalizedLicenseProductEditionOk returns a tuple with the LocalizedLicenseProductEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedLicenseProductEdition

`func (o *LicenseInventoryModel) SetLocalizedLicenseProductEdition(v string)`

SetLocalizedLicenseProductEdition sets LocalizedLicenseProductEdition field to given value.

### HasLocalizedLicenseProductEdition

`func (o *LicenseInventoryModel) HasLocalizedLicenseProductEdition() bool`

HasLocalizedLicenseProductEdition returns a boolean if a field has been set.

### SetLocalizedLicenseProductEditionNil

`func (o *LicenseInventoryModel) SetLocalizedLicenseProductEditionNil(b bool)`

 SetLocalizedLicenseProductEditionNil sets the value for LocalizedLicenseProductEdition to be an explicit nil

### UnsetLocalizedLicenseProductEdition
`func (o *LicenseInventoryModel) UnsetLocalizedLicenseProductEdition()`

UnsetLocalizedLicenseProductEdition ensures that no value is present for LocalizedLicenseProductEdition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


