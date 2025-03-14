# LicenseInformationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseServerStatus** | Pointer to [**LicenseServerStatus**](LicenseServerStatus.md) |  | [optional] 
**LicenseServerWslAddress** | Pointer to **NullableString** | The WSL Address of the license server, the format is https://[FQDN:Port] | [optional] 
**LicenseServerLicensingAddress** | Pointer to **NullableString** | The address of License Address | [optional] 
**LicenseServer** | Pointer to **NullableString** | The address of license server without the schema | [optional] 
**LicensePort** | Pointer to **int32** | The port of license server | [optional] 
**LicensingBurnInDate** | Pointer to **NullableString** | The subscription advantage date of the license | [optional] 
**LicenseProduct** | Pointer to [**LicenseProduct**](LicenseProduct.md) |  | [optional] 
**ProductEdition** | Pointer to [**ProductEdition**](ProductEdition.md) |  | [optional] 
**LicenseModel** | Pointer to [**LicenseModel**](LicenseModel.md) |  | [optional] 
**OutOfBoxGracePeriodActive** | Pointer to **bool** | Indicate if trial license is used | [optional] 
**OutOfBoxGracePeriodHoursLeft** | Pointer to **NullableInt32** | Indicate remaining hours for the trial period | [optional] 
**UseLicenseActivationService** | Pointer to **NullableBool** | Indicate if License Activation Service is used | [optional] 

## Methods

### NewLicenseInformationModel

`func NewLicenseInformationModel() *LicenseInformationModel`

NewLicenseInformationModel instantiates a new LicenseInformationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseInformationModelWithDefaults

`func NewLicenseInformationModelWithDefaults() *LicenseInformationModel`

NewLicenseInformationModelWithDefaults instantiates a new LicenseInformationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseServerStatus

`func (o *LicenseInformationModel) GetLicenseServerStatus() LicenseServerStatus`

GetLicenseServerStatus returns the LicenseServerStatus field if non-nil, zero value otherwise.

### GetLicenseServerStatusOk

`func (o *LicenseInformationModel) GetLicenseServerStatusOk() (*LicenseServerStatus, bool)`

GetLicenseServerStatusOk returns a tuple with the LicenseServerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServerStatus

`func (o *LicenseInformationModel) SetLicenseServerStatus(v LicenseServerStatus)`

SetLicenseServerStatus sets LicenseServerStatus field to given value.

### HasLicenseServerStatus

`func (o *LicenseInformationModel) HasLicenseServerStatus() bool`

HasLicenseServerStatus returns a boolean if a field has been set.

### GetLicenseServerWslAddress

`func (o *LicenseInformationModel) GetLicenseServerWslAddress() string`

GetLicenseServerWslAddress returns the LicenseServerWslAddress field if non-nil, zero value otherwise.

### GetLicenseServerWslAddressOk

`func (o *LicenseInformationModel) GetLicenseServerWslAddressOk() (*string, bool)`

GetLicenseServerWslAddressOk returns a tuple with the LicenseServerWslAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServerWslAddress

`func (o *LicenseInformationModel) SetLicenseServerWslAddress(v string)`

SetLicenseServerWslAddress sets LicenseServerWslAddress field to given value.

### HasLicenseServerWslAddress

`func (o *LicenseInformationModel) HasLicenseServerWslAddress() bool`

HasLicenseServerWslAddress returns a boolean if a field has been set.

### SetLicenseServerWslAddressNil

`func (o *LicenseInformationModel) SetLicenseServerWslAddressNil(b bool)`

 SetLicenseServerWslAddressNil sets the value for LicenseServerWslAddress to be an explicit nil

### UnsetLicenseServerWslAddress
`func (o *LicenseInformationModel) UnsetLicenseServerWslAddress()`

UnsetLicenseServerWslAddress ensures that no value is present for LicenseServerWslAddress, not even an explicit nil
### GetLicenseServerLicensingAddress

`func (o *LicenseInformationModel) GetLicenseServerLicensingAddress() string`

GetLicenseServerLicensingAddress returns the LicenseServerLicensingAddress field if non-nil, zero value otherwise.

### GetLicenseServerLicensingAddressOk

`func (o *LicenseInformationModel) GetLicenseServerLicensingAddressOk() (*string, bool)`

GetLicenseServerLicensingAddressOk returns a tuple with the LicenseServerLicensingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServerLicensingAddress

`func (o *LicenseInformationModel) SetLicenseServerLicensingAddress(v string)`

SetLicenseServerLicensingAddress sets LicenseServerLicensingAddress field to given value.

### HasLicenseServerLicensingAddress

`func (o *LicenseInformationModel) HasLicenseServerLicensingAddress() bool`

HasLicenseServerLicensingAddress returns a boolean if a field has been set.

### SetLicenseServerLicensingAddressNil

`func (o *LicenseInformationModel) SetLicenseServerLicensingAddressNil(b bool)`

 SetLicenseServerLicensingAddressNil sets the value for LicenseServerLicensingAddress to be an explicit nil

### UnsetLicenseServerLicensingAddress
`func (o *LicenseInformationModel) UnsetLicenseServerLicensingAddress()`

UnsetLicenseServerLicensingAddress ensures that no value is present for LicenseServerLicensingAddress, not even an explicit nil
### GetLicenseServer

`func (o *LicenseInformationModel) GetLicenseServer() string`

GetLicenseServer returns the LicenseServer field if non-nil, zero value otherwise.

### GetLicenseServerOk

`func (o *LicenseInformationModel) GetLicenseServerOk() (*string, bool)`

GetLicenseServerOk returns a tuple with the LicenseServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServer

`func (o *LicenseInformationModel) SetLicenseServer(v string)`

SetLicenseServer sets LicenseServer field to given value.

### HasLicenseServer

`func (o *LicenseInformationModel) HasLicenseServer() bool`

HasLicenseServer returns a boolean if a field has been set.

### SetLicenseServerNil

`func (o *LicenseInformationModel) SetLicenseServerNil(b bool)`

 SetLicenseServerNil sets the value for LicenseServer to be an explicit nil

### UnsetLicenseServer
`func (o *LicenseInformationModel) UnsetLicenseServer()`

UnsetLicenseServer ensures that no value is present for LicenseServer, not even an explicit nil
### GetLicensePort

`func (o *LicenseInformationModel) GetLicensePort() int32`

GetLicensePort returns the LicensePort field if non-nil, zero value otherwise.

### GetLicensePortOk

`func (o *LicenseInformationModel) GetLicensePortOk() (*int32, bool)`

GetLicensePortOk returns a tuple with the LicensePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePort

`func (o *LicenseInformationModel) SetLicensePort(v int32)`

SetLicensePort sets LicensePort field to given value.

### HasLicensePort

`func (o *LicenseInformationModel) HasLicensePort() bool`

HasLicensePort returns a boolean if a field has been set.

### GetLicensingBurnInDate

`func (o *LicenseInformationModel) GetLicensingBurnInDate() string`

GetLicensingBurnInDate returns the LicensingBurnInDate field if non-nil, zero value otherwise.

### GetLicensingBurnInDateOk

`func (o *LicenseInformationModel) GetLicensingBurnInDateOk() (*string, bool)`

GetLicensingBurnInDateOk returns a tuple with the LicensingBurnInDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingBurnInDate

`func (o *LicenseInformationModel) SetLicensingBurnInDate(v string)`

SetLicensingBurnInDate sets LicensingBurnInDate field to given value.

### HasLicensingBurnInDate

`func (o *LicenseInformationModel) HasLicensingBurnInDate() bool`

HasLicensingBurnInDate returns a boolean if a field has been set.

### SetLicensingBurnInDateNil

`func (o *LicenseInformationModel) SetLicensingBurnInDateNil(b bool)`

 SetLicensingBurnInDateNil sets the value for LicensingBurnInDate to be an explicit nil

### UnsetLicensingBurnInDate
`func (o *LicenseInformationModel) UnsetLicensingBurnInDate()`

UnsetLicensingBurnInDate ensures that no value is present for LicensingBurnInDate, not even an explicit nil
### GetLicenseProduct

`func (o *LicenseInformationModel) GetLicenseProduct() LicenseProduct`

GetLicenseProduct returns the LicenseProduct field if non-nil, zero value otherwise.

### GetLicenseProductOk

`func (o *LicenseInformationModel) GetLicenseProductOk() (*LicenseProduct, bool)`

GetLicenseProductOk returns a tuple with the LicenseProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseProduct

`func (o *LicenseInformationModel) SetLicenseProduct(v LicenseProduct)`

SetLicenseProduct sets LicenseProduct field to given value.

### HasLicenseProduct

`func (o *LicenseInformationModel) HasLicenseProduct() bool`

HasLicenseProduct returns a boolean if a field has been set.

### GetProductEdition

`func (o *LicenseInformationModel) GetProductEdition() ProductEdition`

GetProductEdition returns the ProductEdition field if non-nil, zero value otherwise.

### GetProductEditionOk

`func (o *LicenseInformationModel) GetProductEditionOk() (*ProductEdition, bool)`

GetProductEditionOk returns a tuple with the ProductEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductEdition

`func (o *LicenseInformationModel) SetProductEdition(v ProductEdition)`

SetProductEdition sets ProductEdition field to given value.

### HasProductEdition

`func (o *LicenseInformationModel) HasProductEdition() bool`

HasProductEdition returns a boolean if a field has been set.

### GetLicenseModel

`func (o *LicenseInformationModel) GetLicenseModel() LicenseModel`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *LicenseInformationModel) GetLicenseModelOk() (*LicenseModel, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *LicenseInformationModel) SetLicenseModel(v LicenseModel)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *LicenseInformationModel) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetOutOfBoxGracePeriodActive

`func (o *LicenseInformationModel) GetOutOfBoxGracePeriodActive() bool`

GetOutOfBoxGracePeriodActive returns the OutOfBoxGracePeriodActive field if non-nil, zero value otherwise.

### GetOutOfBoxGracePeriodActiveOk

`func (o *LicenseInformationModel) GetOutOfBoxGracePeriodActiveOk() (*bool, bool)`

GetOutOfBoxGracePeriodActiveOk returns a tuple with the OutOfBoxGracePeriodActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBoxGracePeriodActive

`func (o *LicenseInformationModel) SetOutOfBoxGracePeriodActive(v bool)`

SetOutOfBoxGracePeriodActive sets OutOfBoxGracePeriodActive field to given value.

### HasOutOfBoxGracePeriodActive

`func (o *LicenseInformationModel) HasOutOfBoxGracePeriodActive() bool`

HasOutOfBoxGracePeriodActive returns a boolean if a field has been set.

### GetOutOfBoxGracePeriodHoursLeft

`func (o *LicenseInformationModel) GetOutOfBoxGracePeriodHoursLeft() int32`

GetOutOfBoxGracePeriodHoursLeft returns the OutOfBoxGracePeriodHoursLeft field if non-nil, zero value otherwise.

### GetOutOfBoxGracePeriodHoursLeftOk

`func (o *LicenseInformationModel) GetOutOfBoxGracePeriodHoursLeftOk() (*int32, bool)`

GetOutOfBoxGracePeriodHoursLeftOk returns a tuple with the OutOfBoxGracePeriodHoursLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBoxGracePeriodHoursLeft

`func (o *LicenseInformationModel) SetOutOfBoxGracePeriodHoursLeft(v int32)`

SetOutOfBoxGracePeriodHoursLeft sets OutOfBoxGracePeriodHoursLeft field to given value.

### HasOutOfBoxGracePeriodHoursLeft

`func (o *LicenseInformationModel) HasOutOfBoxGracePeriodHoursLeft() bool`

HasOutOfBoxGracePeriodHoursLeft returns a boolean if a field has been set.

### SetOutOfBoxGracePeriodHoursLeftNil

`func (o *LicenseInformationModel) SetOutOfBoxGracePeriodHoursLeftNil(b bool)`

 SetOutOfBoxGracePeriodHoursLeftNil sets the value for OutOfBoxGracePeriodHoursLeft to be an explicit nil

### UnsetOutOfBoxGracePeriodHoursLeft
`func (o *LicenseInformationModel) UnsetOutOfBoxGracePeriodHoursLeft()`

UnsetOutOfBoxGracePeriodHoursLeft ensures that no value is present for OutOfBoxGracePeriodHoursLeft, not even an explicit nil
### GetUseLicenseActivationService

`func (o *LicenseInformationModel) GetUseLicenseActivationService() bool`

GetUseLicenseActivationService returns the UseLicenseActivationService field if non-nil, zero value otherwise.

### GetUseLicenseActivationServiceOk

`func (o *LicenseInformationModel) GetUseLicenseActivationServiceOk() (*bool, bool)`

GetUseLicenseActivationServiceOk returns a tuple with the UseLicenseActivationService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLicenseActivationService

`func (o *LicenseInformationModel) SetUseLicenseActivationService(v bool)`

SetUseLicenseActivationService sets UseLicenseActivationService field to given value.

### HasUseLicenseActivationService

`func (o *LicenseInformationModel) HasUseLicenseActivationService() bool`

HasUseLicenseActivationService returns a boolean if a field has been set.

### SetUseLicenseActivationServiceNil

`func (o *LicenseInformationModel) SetUseLicenseActivationServiceNil(b bool)`

 SetUseLicenseActivationServiceNil sets the value for UseLicenseActivationService to be an explicit nil

### UnsetUseLicenseActivationService
`func (o *LicenseInformationModel) UnsetUseLicenseActivationService()`

UnsetUseLicenseActivationService ensures that no value is present for UseLicenseActivationService, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


