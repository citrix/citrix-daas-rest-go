# UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebootDuration** | Pointer to **int32** | Approximate maximum duration over which the reboot cycle runs, in minutes. | [optional] [default to -1]
**WarningDuration** | Pointer to **int32** | Time in minutes prior to a machine reboot at which a warning message is displayed in all user sessions on that machine. | [optional] [default to 0]
**WarningTitle** | Pointer to **string** | The window title used when showing the warning message in user sessions on a machine scheduled for reboot. | [optional] 
**WarningMessage** | Pointer to **string** | Warning message displayed in user sessions on a machine scheduled for a reboot.  The optional pattern &#39;%m%&#39; is replaced by the number of minutes until the reboot. | [optional] 
**WarningRepeatInterval** | Pointer to **int32** | Number of minutes to wait before showing the reboot warning message again. | [optional] [default to 0]

## Methods

### NewUpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions

`func NewUpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions() *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions`

NewUpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions instantiates a new UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMachineCatalogProvisioningSchemeRequestModelRebootOptionsWithDefaults

`func NewUpdateMachineCatalogProvisioningSchemeRequestModelRebootOptionsWithDefaults() *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions`

NewUpdateMachineCatalogProvisioningSchemeRequestModelRebootOptionsWithDefaults instantiates a new UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRebootDuration

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) GetRebootDuration() int32`

GetRebootDuration returns the RebootDuration field if non-nil, zero value otherwise.

### GetRebootDurationOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) GetRebootDurationOk() (*int32, bool)`

GetRebootDurationOk returns a tuple with the RebootDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootDuration

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) SetRebootDuration(v int32)`

SetRebootDuration sets RebootDuration field to given value.

### HasRebootDuration

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) HasRebootDuration() bool`

HasRebootDuration returns a boolean if a field has been set.

### GetWarningDuration

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) GetWarningDuration() int32`

GetWarningDuration returns the WarningDuration field if non-nil, zero value otherwise.

### GetWarningDurationOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) GetWarningDurationOk() (*int32, bool)`

GetWarningDurationOk returns a tuple with the WarningDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningDuration

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) SetWarningDuration(v int32)`

SetWarningDuration sets WarningDuration field to given value.

### HasWarningDuration

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) HasWarningDuration() bool`

HasWarningDuration returns a boolean if a field has been set.

### GetWarningTitle

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) GetWarningTitle() string`

GetWarningTitle returns the WarningTitle field if non-nil, zero value otherwise.

### GetWarningTitleOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) GetWarningTitleOk() (*string, bool)`

GetWarningTitleOk returns a tuple with the WarningTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningTitle

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) SetWarningTitle(v string)`

SetWarningTitle sets WarningTitle field to given value.

### HasWarningTitle

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) HasWarningTitle() bool`

HasWarningTitle returns a boolean if a field has been set.

### GetWarningMessage

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) GetWarningMessage() string`

GetWarningMessage returns the WarningMessage field if non-nil, zero value otherwise.

### GetWarningMessageOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) GetWarningMessageOk() (*string, bool)`

GetWarningMessageOk returns a tuple with the WarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessage

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) SetWarningMessage(v string)`

SetWarningMessage sets WarningMessage field to given value.

### HasWarningMessage

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) HasWarningMessage() bool`

HasWarningMessage returns a boolean if a field has been set.

### GetWarningRepeatInterval

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) GetWarningRepeatInterval() int32`

GetWarningRepeatInterval returns the WarningRepeatInterval field if non-nil, zero value otherwise.

### GetWarningRepeatIntervalOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) GetWarningRepeatIntervalOk() (*int32, bool)`

GetWarningRepeatIntervalOk returns a tuple with the WarningRepeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningRepeatInterval

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) SetWarningRepeatInterval(v int32)`

SetWarningRepeatInterval sets WarningRepeatInterval field to given value.

### HasWarningRepeatInterval

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions) HasWarningRepeatInterval() bool`

HasWarningRepeatInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


