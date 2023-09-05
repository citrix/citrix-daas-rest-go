# LicenseSummaryResponseModelLicenseInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseServerWslAddress** | Pointer to **string** | The WSL Address of the license server, the format is https://[FQDN:Port] | [optional] 
**LicenseServerLicensingAddress** | Pointer to **string** | The address of License Address | [optional] 
**LicenseServer** | Pointer to **string** | The address of license server without the schema | [optional] 
**LicensePort** | Pointer to **int32** | The port of license server | [optional] 
**LicensingBurnInDate** | Pointer to **string** | The subscription advantage date of the license | [optional] 
**LicenseProduct** | Pointer to [**LicenseProduct**](LicenseProduct.md) |  | [optional] 
**ProductEdition** | Pointer to [**ProductEdition**](ProductEdition.md) |  | [optional] 
**LicenseModel** | Pointer to [**LicenseModel**](LicenseModel.md) |  | [optional] 
**OutOfBoxGracePeriodActive** | Pointer to **bool** | Indicate if trial license is used | [optional] 
**OutOfBoxGracePeriodHoursLeft** | Pointer to **int32** | Indicate remaining hours for the trial period | [optional] 

## Methods

### NewLicenseSummaryResponseModelLicenseInformation

`func NewLicenseSummaryResponseModelLicenseInformation() *LicenseSummaryResponseModelLicenseInformation`

NewLicenseSummaryResponseModelLicenseInformation instantiates a new LicenseSummaryResponseModelLicenseInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseSummaryResponseModelLicenseInformationWithDefaults

`func NewLicenseSummaryResponseModelLicenseInformationWithDefaults() *LicenseSummaryResponseModelLicenseInformation`

NewLicenseSummaryResponseModelLicenseInformationWithDefaults instantiates a new LicenseSummaryResponseModelLicenseInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseServerWslAddress

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicenseServerWslAddress() string`

GetLicenseServerWslAddress returns the LicenseServerWslAddress field if non-nil, zero value otherwise.

### GetLicenseServerWslAddressOk

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicenseServerWslAddressOk() (*string, bool)`

GetLicenseServerWslAddressOk returns a tuple with the LicenseServerWslAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServerWslAddress

`func (o *LicenseSummaryResponseModelLicenseInformation) SetLicenseServerWslAddress(v string)`

SetLicenseServerWslAddress sets LicenseServerWslAddress field to given value.

### HasLicenseServerWslAddress

`func (o *LicenseSummaryResponseModelLicenseInformation) HasLicenseServerWslAddress() bool`

HasLicenseServerWslAddress returns a boolean if a field has been set.

### GetLicenseServerLicensingAddress

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicenseServerLicensingAddress() string`

GetLicenseServerLicensingAddress returns the LicenseServerLicensingAddress field if non-nil, zero value otherwise.

### GetLicenseServerLicensingAddressOk

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicenseServerLicensingAddressOk() (*string, bool)`

GetLicenseServerLicensingAddressOk returns a tuple with the LicenseServerLicensingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServerLicensingAddress

`func (o *LicenseSummaryResponseModelLicenseInformation) SetLicenseServerLicensingAddress(v string)`

SetLicenseServerLicensingAddress sets LicenseServerLicensingAddress field to given value.

### HasLicenseServerLicensingAddress

`func (o *LicenseSummaryResponseModelLicenseInformation) HasLicenseServerLicensingAddress() bool`

HasLicenseServerLicensingAddress returns a boolean if a field has been set.

### GetLicenseServer

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicenseServer() string`

GetLicenseServer returns the LicenseServer field if non-nil, zero value otherwise.

### GetLicenseServerOk

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicenseServerOk() (*string, bool)`

GetLicenseServerOk returns a tuple with the LicenseServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServer

`func (o *LicenseSummaryResponseModelLicenseInformation) SetLicenseServer(v string)`

SetLicenseServer sets LicenseServer field to given value.

### HasLicenseServer

`func (o *LicenseSummaryResponseModelLicenseInformation) HasLicenseServer() bool`

HasLicenseServer returns a boolean if a field has been set.

### GetLicensePort

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicensePort() int32`

GetLicensePort returns the LicensePort field if non-nil, zero value otherwise.

### GetLicensePortOk

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicensePortOk() (*int32, bool)`

GetLicensePortOk returns a tuple with the LicensePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePort

`func (o *LicenseSummaryResponseModelLicenseInformation) SetLicensePort(v int32)`

SetLicensePort sets LicensePort field to given value.

### HasLicensePort

`func (o *LicenseSummaryResponseModelLicenseInformation) HasLicensePort() bool`

HasLicensePort returns a boolean if a field has been set.

### GetLicensingBurnInDate

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicensingBurnInDate() string`

GetLicensingBurnInDate returns the LicensingBurnInDate field if non-nil, zero value otherwise.

### GetLicensingBurnInDateOk

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicensingBurnInDateOk() (*string, bool)`

GetLicensingBurnInDateOk returns a tuple with the LicensingBurnInDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingBurnInDate

`func (o *LicenseSummaryResponseModelLicenseInformation) SetLicensingBurnInDate(v string)`

SetLicensingBurnInDate sets LicensingBurnInDate field to given value.

### HasLicensingBurnInDate

`func (o *LicenseSummaryResponseModelLicenseInformation) HasLicensingBurnInDate() bool`

HasLicensingBurnInDate returns a boolean if a field has been set.

### GetLicenseProduct

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicenseProduct() LicenseProduct`

GetLicenseProduct returns the LicenseProduct field if non-nil, zero value otherwise.

### GetLicenseProductOk

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicenseProductOk() (*LicenseProduct, bool)`

GetLicenseProductOk returns a tuple with the LicenseProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseProduct

`func (o *LicenseSummaryResponseModelLicenseInformation) SetLicenseProduct(v LicenseProduct)`

SetLicenseProduct sets LicenseProduct field to given value.

### HasLicenseProduct

`func (o *LicenseSummaryResponseModelLicenseInformation) HasLicenseProduct() bool`

HasLicenseProduct returns a boolean if a field has been set.

### GetProductEdition

`func (o *LicenseSummaryResponseModelLicenseInformation) GetProductEdition() ProductEdition`

GetProductEdition returns the ProductEdition field if non-nil, zero value otherwise.

### GetProductEditionOk

`func (o *LicenseSummaryResponseModelLicenseInformation) GetProductEditionOk() (*ProductEdition, bool)`

GetProductEditionOk returns a tuple with the ProductEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductEdition

`func (o *LicenseSummaryResponseModelLicenseInformation) SetProductEdition(v ProductEdition)`

SetProductEdition sets ProductEdition field to given value.

### HasProductEdition

`func (o *LicenseSummaryResponseModelLicenseInformation) HasProductEdition() bool`

HasProductEdition returns a boolean if a field has been set.

### GetLicenseModel

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicenseModel() LicenseModel`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *LicenseSummaryResponseModelLicenseInformation) GetLicenseModelOk() (*LicenseModel, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *LicenseSummaryResponseModelLicenseInformation) SetLicenseModel(v LicenseModel)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *LicenseSummaryResponseModelLicenseInformation) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetOutOfBoxGracePeriodActive

`func (o *LicenseSummaryResponseModelLicenseInformation) GetOutOfBoxGracePeriodActive() bool`

GetOutOfBoxGracePeriodActive returns the OutOfBoxGracePeriodActive field if non-nil, zero value otherwise.

### GetOutOfBoxGracePeriodActiveOk

`func (o *LicenseSummaryResponseModelLicenseInformation) GetOutOfBoxGracePeriodActiveOk() (*bool, bool)`

GetOutOfBoxGracePeriodActiveOk returns a tuple with the OutOfBoxGracePeriodActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBoxGracePeriodActive

`func (o *LicenseSummaryResponseModelLicenseInformation) SetOutOfBoxGracePeriodActive(v bool)`

SetOutOfBoxGracePeriodActive sets OutOfBoxGracePeriodActive field to given value.

### HasOutOfBoxGracePeriodActive

`func (o *LicenseSummaryResponseModelLicenseInformation) HasOutOfBoxGracePeriodActive() bool`

HasOutOfBoxGracePeriodActive returns a boolean if a field has been set.

### GetOutOfBoxGracePeriodHoursLeft

`func (o *LicenseSummaryResponseModelLicenseInformation) GetOutOfBoxGracePeriodHoursLeft() int32`

GetOutOfBoxGracePeriodHoursLeft returns the OutOfBoxGracePeriodHoursLeft field if non-nil, zero value otherwise.

### GetOutOfBoxGracePeriodHoursLeftOk

`func (o *LicenseSummaryResponseModelLicenseInformation) GetOutOfBoxGracePeriodHoursLeftOk() (*int32, bool)`

GetOutOfBoxGracePeriodHoursLeftOk returns a tuple with the OutOfBoxGracePeriodHoursLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBoxGracePeriodHoursLeft

`func (o *LicenseSummaryResponseModelLicenseInformation) SetOutOfBoxGracePeriodHoursLeft(v int32)`

SetOutOfBoxGracePeriodHoursLeft sets OutOfBoxGracePeriodHoursLeft field to given value.

### HasOutOfBoxGracePeriodHoursLeft

`func (o *LicenseSummaryResponseModelLicenseInformation) HasOutOfBoxGracePeriodHoursLeft() bool`

HasOutOfBoxGracePeriodHoursLeft returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


