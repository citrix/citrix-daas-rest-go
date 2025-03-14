# UpdateMaintenanceCycleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceCycleName** | Pointer to **NullableString** | The name of the Maintenance Cycle. Name is globally unique. | [optional] 
**MaintenanceCycleDescription** | Pointer to **NullableString** | The description of the Maintenance Cycle. | [optional] 
**MaintenanceCycleUid** | **string** | The maintenance cycle uid. | 
**ProvisioningSchemeUid** | Pointer to **NullableString** | The uid of provisioning scheme. | [optional] 
**ScheduledStartTimeInUTC** | Pointer to **NullableTime** | Maintenance cycle start time in UTC. | [optional] 
**MaxDurationInMinutes** | **int32** | The max duration allowed for maintenance cycle. | 
**PurgeDBAfterInDays** | Pointer to **NullableInt32** | Purge history data after a certain number of days. | [optional] 
**SessionWarningLogOffTitle** | Pointer to **NullableString** | The title of session log off warning.  | [optional] 
**SessionWarningLogOffMessage** | Pointer to **NullableString** | The message of session log off warning. | [optional] 
**SessionWarningTimeInMinutes** | Pointer to **NullableInt32** | The countdown duration for session log off.  | [optional] 
**VMNames** | Pointer to **[]string** | The vm names in maintenance cycle.  | [optional] 

## Methods

### NewUpdateMaintenanceCycleRequestModel

`func NewUpdateMaintenanceCycleRequestModel(maintenanceCycleUid string, maxDurationInMinutes int32, ) *UpdateMaintenanceCycleRequestModel`

NewUpdateMaintenanceCycleRequestModel instantiates a new UpdateMaintenanceCycleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMaintenanceCycleRequestModelWithDefaults

`func NewUpdateMaintenanceCycleRequestModelWithDefaults() *UpdateMaintenanceCycleRequestModel`

NewUpdateMaintenanceCycleRequestModelWithDefaults instantiates a new UpdateMaintenanceCycleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceCycleName

`func (o *UpdateMaintenanceCycleRequestModel) GetMaintenanceCycleName() string`

GetMaintenanceCycleName returns the MaintenanceCycleName field if non-nil, zero value otherwise.

### GetMaintenanceCycleNameOk

`func (o *UpdateMaintenanceCycleRequestModel) GetMaintenanceCycleNameOk() (*string, bool)`

GetMaintenanceCycleNameOk returns a tuple with the MaintenanceCycleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleName

`func (o *UpdateMaintenanceCycleRequestModel) SetMaintenanceCycleName(v string)`

SetMaintenanceCycleName sets MaintenanceCycleName field to given value.

### HasMaintenanceCycleName

`func (o *UpdateMaintenanceCycleRequestModel) HasMaintenanceCycleName() bool`

HasMaintenanceCycleName returns a boolean if a field has been set.

### SetMaintenanceCycleNameNil

`func (o *UpdateMaintenanceCycleRequestModel) SetMaintenanceCycleNameNil(b bool)`

 SetMaintenanceCycleNameNil sets the value for MaintenanceCycleName to be an explicit nil

### UnsetMaintenanceCycleName
`func (o *UpdateMaintenanceCycleRequestModel) UnsetMaintenanceCycleName()`

UnsetMaintenanceCycleName ensures that no value is present for MaintenanceCycleName, not even an explicit nil
### GetMaintenanceCycleDescription

`func (o *UpdateMaintenanceCycleRequestModel) GetMaintenanceCycleDescription() string`

GetMaintenanceCycleDescription returns the MaintenanceCycleDescription field if non-nil, zero value otherwise.

### GetMaintenanceCycleDescriptionOk

`func (o *UpdateMaintenanceCycleRequestModel) GetMaintenanceCycleDescriptionOk() (*string, bool)`

GetMaintenanceCycleDescriptionOk returns a tuple with the MaintenanceCycleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleDescription

`func (o *UpdateMaintenanceCycleRequestModel) SetMaintenanceCycleDescription(v string)`

SetMaintenanceCycleDescription sets MaintenanceCycleDescription field to given value.

### HasMaintenanceCycleDescription

`func (o *UpdateMaintenanceCycleRequestModel) HasMaintenanceCycleDescription() bool`

HasMaintenanceCycleDescription returns a boolean if a field has been set.

### SetMaintenanceCycleDescriptionNil

`func (o *UpdateMaintenanceCycleRequestModel) SetMaintenanceCycleDescriptionNil(b bool)`

 SetMaintenanceCycleDescriptionNil sets the value for MaintenanceCycleDescription to be an explicit nil

### UnsetMaintenanceCycleDescription
`func (o *UpdateMaintenanceCycleRequestModel) UnsetMaintenanceCycleDescription()`

UnsetMaintenanceCycleDescription ensures that no value is present for MaintenanceCycleDescription, not even an explicit nil
### GetMaintenanceCycleUid

`func (o *UpdateMaintenanceCycleRequestModel) GetMaintenanceCycleUid() string`

GetMaintenanceCycleUid returns the MaintenanceCycleUid field if non-nil, zero value otherwise.

### GetMaintenanceCycleUidOk

`func (o *UpdateMaintenanceCycleRequestModel) GetMaintenanceCycleUidOk() (*string, bool)`

GetMaintenanceCycleUidOk returns a tuple with the MaintenanceCycleUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleUid

`func (o *UpdateMaintenanceCycleRequestModel) SetMaintenanceCycleUid(v string)`

SetMaintenanceCycleUid sets MaintenanceCycleUid field to given value.


### GetProvisioningSchemeUid

`func (o *UpdateMaintenanceCycleRequestModel) GetProvisioningSchemeUid() string`

GetProvisioningSchemeUid returns the ProvisioningSchemeUid field if non-nil, zero value otherwise.

### GetProvisioningSchemeUidOk

`func (o *UpdateMaintenanceCycleRequestModel) GetProvisioningSchemeUidOk() (*string, bool)`

GetProvisioningSchemeUidOk returns a tuple with the ProvisioningSchemeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeUid

`func (o *UpdateMaintenanceCycleRequestModel) SetProvisioningSchemeUid(v string)`

SetProvisioningSchemeUid sets ProvisioningSchemeUid field to given value.

### HasProvisioningSchemeUid

`func (o *UpdateMaintenanceCycleRequestModel) HasProvisioningSchemeUid() bool`

HasProvisioningSchemeUid returns a boolean if a field has been set.

### SetProvisioningSchemeUidNil

`func (o *UpdateMaintenanceCycleRequestModel) SetProvisioningSchemeUidNil(b bool)`

 SetProvisioningSchemeUidNil sets the value for ProvisioningSchemeUid to be an explicit nil

### UnsetProvisioningSchemeUid
`func (o *UpdateMaintenanceCycleRequestModel) UnsetProvisioningSchemeUid()`

UnsetProvisioningSchemeUid ensures that no value is present for ProvisioningSchemeUid, not even an explicit nil
### GetScheduledStartTimeInUTC

`func (o *UpdateMaintenanceCycleRequestModel) GetScheduledStartTimeInUTC() time.Time`

GetScheduledStartTimeInUTC returns the ScheduledStartTimeInUTC field if non-nil, zero value otherwise.

### GetScheduledStartTimeInUTCOk

`func (o *UpdateMaintenanceCycleRequestModel) GetScheduledStartTimeInUTCOk() (*time.Time, bool)`

GetScheduledStartTimeInUTCOk returns a tuple with the ScheduledStartTimeInUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTimeInUTC

`func (o *UpdateMaintenanceCycleRequestModel) SetScheduledStartTimeInUTC(v time.Time)`

SetScheduledStartTimeInUTC sets ScheduledStartTimeInUTC field to given value.

### HasScheduledStartTimeInUTC

`func (o *UpdateMaintenanceCycleRequestModel) HasScheduledStartTimeInUTC() bool`

HasScheduledStartTimeInUTC returns a boolean if a field has been set.

### SetScheduledStartTimeInUTCNil

`func (o *UpdateMaintenanceCycleRequestModel) SetScheduledStartTimeInUTCNil(b bool)`

 SetScheduledStartTimeInUTCNil sets the value for ScheduledStartTimeInUTC to be an explicit nil

### UnsetScheduledStartTimeInUTC
`func (o *UpdateMaintenanceCycleRequestModel) UnsetScheduledStartTimeInUTC()`

UnsetScheduledStartTimeInUTC ensures that no value is present for ScheduledStartTimeInUTC, not even an explicit nil
### GetMaxDurationInMinutes

`func (o *UpdateMaintenanceCycleRequestModel) GetMaxDurationInMinutes() int32`

GetMaxDurationInMinutes returns the MaxDurationInMinutes field if non-nil, zero value otherwise.

### GetMaxDurationInMinutesOk

`func (o *UpdateMaintenanceCycleRequestModel) GetMaxDurationInMinutesOk() (*int32, bool)`

GetMaxDurationInMinutesOk returns a tuple with the MaxDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationInMinutes

`func (o *UpdateMaintenanceCycleRequestModel) SetMaxDurationInMinutes(v int32)`

SetMaxDurationInMinutes sets MaxDurationInMinutes field to given value.


### GetPurgeDBAfterInDays

`func (o *UpdateMaintenanceCycleRequestModel) GetPurgeDBAfterInDays() int32`

GetPurgeDBAfterInDays returns the PurgeDBAfterInDays field if non-nil, zero value otherwise.

### GetPurgeDBAfterInDaysOk

`func (o *UpdateMaintenanceCycleRequestModel) GetPurgeDBAfterInDaysOk() (*int32, bool)`

GetPurgeDBAfterInDaysOk returns a tuple with the PurgeDBAfterInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeDBAfterInDays

`func (o *UpdateMaintenanceCycleRequestModel) SetPurgeDBAfterInDays(v int32)`

SetPurgeDBAfterInDays sets PurgeDBAfterInDays field to given value.

### HasPurgeDBAfterInDays

`func (o *UpdateMaintenanceCycleRequestModel) HasPurgeDBAfterInDays() bool`

HasPurgeDBAfterInDays returns a boolean if a field has been set.

### SetPurgeDBAfterInDaysNil

`func (o *UpdateMaintenanceCycleRequestModel) SetPurgeDBAfterInDaysNil(b bool)`

 SetPurgeDBAfterInDaysNil sets the value for PurgeDBAfterInDays to be an explicit nil

### UnsetPurgeDBAfterInDays
`func (o *UpdateMaintenanceCycleRequestModel) UnsetPurgeDBAfterInDays()`

UnsetPurgeDBAfterInDays ensures that no value is present for PurgeDBAfterInDays, not even an explicit nil
### GetSessionWarningLogOffTitle

`func (o *UpdateMaintenanceCycleRequestModel) GetSessionWarningLogOffTitle() string`

GetSessionWarningLogOffTitle returns the SessionWarningLogOffTitle field if non-nil, zero value otherwise.

### GetSessionWarningLogOffTitleOk

`func (o *UpdateMaintenanceCycleRequestModel) GetSessionWarningLogOffTitleOk() (*string, bool)`

GetSessionWarningLogOffTitleOk returns a tuple with the SessionWarningLogOffTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffTitle

`func (o *UpdateMaintenanceCycleRequestModel) SetSessionWarningLogOffTitle(v string)`

SetSessionWarningLogOffTitle sets SessionWarningLogOffTitle field to given value.

### HasSessionWarningLogOffTitle

`func (o *UpdateMaintenanceCycleRequestModel) HasSessionWarningLogOffTitle() bool`

HasSessionWarningLogOffTitle returns a boolean if a field has been set.

### SetSessionWarningLogOffTitleNil

`func (o *UpdateMaintenanceCycleRequestModel) SetSessionWarningLogOffTitleNil(b bool)`

 SetSessionWarningLogOffTitleNil sets the value for SessionWarningLogOffTitle to be an explicit nil

### UnsetSessionWarningLogOffTitle
`func (o *UpdateMaintenanceCycleRequestModel) UnsetSessionWarningLogOffTitle()`

UnsetSessionWarningLogOffTitle ensures that no value is present for SessionWarningLogOffTitle, not even an explicit nil
### GetSessionWarningLogOffMessage

`func (o *UpdateMaintenanceCycleRequestModel) GetSessionWarningLogOffMessage() string`

GetSessionWarningLogOffMessage returns the SessionWarningLogOffMessage field if non-nil, zero value otherwise.

### GetSessionWarningLogOffMessageOk

`func (o *UpdateMaintenanceCycleRequestModel) GetSessionWarningLogOffMessageOk() (*string, bool)`

GetSessionWarningLogOffMessageOk returns a tuple with the SessionWarningLogOffMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffMessage

`func (o *UpdateMaintenanceCycleRequestModel) SetSessionWarningLogOffMessage(v string)`

SetSessionWarningLogOffMessage sets SessionWarningLogOffMessage field to given value.

### HasSessionWarningLogOffMessage

`func (o *UpdateMaintenanceCycleRequestModel) HasSessionWarningLogOffMessage() bool`

HasSessionWarningLogOffMessage returns a boolean if a field has been set.

### SetSessionWarningLogOffMessageNil

`func (o *UpdateMaintenanceCycleRequestModel) SetSessionWarningLogOffMessageNil(b bool)`

 SetSessionWarningLogOffMessageNil sets the value for SessionWarningLogOffMessage to be an explicit nil

### UnsetSessionWarningLogOffMessage
`func (o *UpdateMaintenanceCycleRequestModel) UnsetSessionWarningLogOffMessage()`

UnsetSessionWarningLogOffMessage ensures that no value is present for SessionWarningLogOffMessage, not even an explicit nil
### GetSessionWarningTimeInMinutes

`func (o *UpdateMaintenanceCycleRequestModel) GetSessionWarningTimeInMinutes() int32`

GetSessionWarningTimeInMinutes returns the SessionWarningTimeInMinutes field if non-nil, zero value otherwise.

### GetSessionWarningTimeInMinutesOk

`func (o *UpdateMaintenanceCycleRequestModel) GetSessionWarningTimeInMinutesOk() (*int32, bool)`

GetSessionWarningTimeInMinutesOk returns a tuple with the SessionWarningTimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningTimeInMinutes

`func (o *UpdateMaintenanceCycleRequestModel) SetSessionWarningTimeInMinutes(v int32)`

SetSessionWarningTimeInMinutes sets SessionWarningTimeInMinutes field to given value.

### HasSessionWarningTimeInMinutes

`func (o *UpdateMaintenanceCycleRequestModel) HasSessionWarningTimeInMinutes() bool`

HasSessionWarningTimeInMinutes returns a boolean if a field has been set.

### SetSessionWarningTimeInMinutesNil

`func (o *UpdateMaintenanceCycleRequestModel) SetSessionWarningTimeInMinutesNil(b bool)`

 SetSessionWarningTimeInMinutesNil sets the value for SessionWarningTimeInMinutes to be an explicit nil

### UnsetSessionWarningTimeInMinutes
`func (o *UpdateMaintenanceCycleRequestModel) UnsetSessionWarningTimeInMinutes()`

UnsetSessionWarningTimeInMinutes ensures that no value is present for SessionWarningTimeInMinutes, not even an explicit nil
### GetVMNames

`func (o *UpdateMaintenanceCycleRequestModel) GetVMNames() []string`

GetVMNames returns the VMNames field if non-nil, zero value otherwise.

### GetVMNamesOk

`func (o *UpdateMaintenanceCycleRequestModel) GetVMNamesOk() (*[]string, bool)`

GetVMNamesOk returns a tuple with the VMNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMNames

`func (o *UpdateMaintenanceCycleRequestModel) SetVMNames(v []string)`

SetVMNames sets VMNames field to given value.

### HasVMNames

`func (o *UpdateMaintenanceCycleRequestModel) HasVMNames() bool`

HasVMNames returns a boolean if a field has been set.

### SetVMNamesNil

`func (o *UpdateMaintenanceCycleRequestModel) SetVMNamesNil(b bool)`

 SetVMNamesNil sets the value for VMNames to be an explicit nil

### UnsetVMNames
`func (o *UpdateMaintenanceCycleRequestModel) UnsetVMNames()`

UnsetVMNames ensures that no value is present for VMNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


