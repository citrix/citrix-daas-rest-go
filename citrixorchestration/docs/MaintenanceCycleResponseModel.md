# MaintenanceCycleResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceCycleName** | Pointer to **NullableString** | The name of the Maintenance Cycle. Name is globally unique. | [optional] 
**MaintenanceCycleId** | Pointer to **string** | The id of maintenance cycle. | [optional] 
**MaintenanceCycleDescription** | Pointer to **NullableString** | The description of the Maintenance Cycle. | [optional] 
**MaintenanceCycleStatus** | Pointer to [**MaintenanceCycleStatus**](MaintenanceCycleStatus.md) |  | [optional] 
**MaxDurationInMinutes** | Pointer to **int32** | The max duration allowed for maintenance cycle. | [optional] 
**ProvisioningSchemeName** | Pointer to **NullableString** | The name of provisioning scheme. | [optional] 
**ProvisioningSchemeUid** | Pointer to **NullableString** | The uid of provisioning scheme. | [optional] 
**PurgeDBAfterInDays** | Pointer to **NullableInt32** | Purge history data after a certain number of days. | [optional] 
**ScheduledCompletionTimeInUTC** | Pointer to **NullableTime** | Maintenance cycle complete time in UTC. | [optional] 
**ScheduledStartTimeInUTC** | Pointer to **NullableTime** | Maintenance cycle start time in UTC. | [optional] 
**ScheduleId** | Pointer to **NullableString** | The schedule Id of maintenance cycle. It&#39;s not empty only if the maintenance cycle created by scheudle. | [optional] 
**SessionWarningLogOffTitle** | Pointer to **NullableString** | The title of session log off warning.  | [optional] 
**SessionWarningLogOffMessage** | Pointer to **NullableString** | The message of session log off warning. | [optional] 
**SessionWarningTimeInMinutes** | Pointer to **NullableInt32** | The countdown duration for session log off.  | [optional] 
**Operation** | Pointer to [**MaintenanceCycleOperation**](MaintenanceCycleOperation.md) |  | [optional] 
**AllVM** | Pointer to **bool** | If maintenance cycle created base on machine catalog. | [optional] 
**VMNames** | Pointer to **[]string** | The vm names in maintenance cycle.  | [optional] 

## Methods

### NewMaintenanceCycleResponseModel

`func NewMaintenanceCycleResponseModel() *MaintenanceCycleResponseModel`

NewMaintenanceCycleResponseModel instantiates a new MaintenanceCycleResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceCycleResponseModelWithDefaults

`func NewMaintenanceCycleResponseModelWithDefaults() *MaintenanceCycleResponseModel`

NewMaintenanceCycleResponseModelWithDefaults instantiates a new MaintenanceCycleResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceCycleName

`func (o *MaintenanceCycleResponseModel) GetMaintenanceCycleName() string`

GetMaintenanceCycleName returns the MaintenanceCycleName field if non-nil, zero value otherwise.

### GetMaintenanceCycleNameOk

`func (o *MaintenanceCycleResponseModel) GetMaintenanceCycleNameOk() (*string, bool)`

GetMaintenanceCycleNameOk returns a tuple with the MaintenanceCycleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleName

`func (o *MaintenanceCycleResponseModel) SetMaintenanceCycleName(v string)`

SetMaintenanceCycleName sets MaintenanceCycleName field to given value.

### HasMaintenanceCycleName

`func (o *MaintenanceCycleResponseModel) HasMaintenanceCycleName() bool`

HasMaintenanceCycleName returns a boolean if a field has been set.

### SetMaintenanceCycleNameNil

`func (o *MaintenanceCycleResponseModel) SetMaintenanceCycleNameNil(b bool)`

 SetMaintenanceCycleNameNil sets the value for MaintenanceCycleName to be an explicit nil

### UnsetMaintenanceCycleName
`func (o *MaintenanceCycleResponseModel) UnsetMaintenanceCycleName()`

UnsetMaintenanceCycleName ensures that no value is present for MaintenanceCycleName, not even an explicit nil
### GetMaintenanceCycleId

`func (o *MaintenanceCycleResponseModel) GetMaintenanceCycleId() string`

GetMaintenanceCycleId returns the MaintenanceCycleId field if non-nil, zero value otherwise.

### GetMaintenanceCycleIdOk

`func (o *MaintenanceCycleResponseModel) GetMaintenanceCycleIdOk() (*string, bool)`

GetMaintenanceCycleIdOk returns a tuple with the MaintenanceCycleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleId

`func (o *MaintenanceCycleResponseModel) SetMaintenanceCycleId(v string)`

SetMaintenanceCycleId sets MaintenanceCycleId field to given value.

### HasMaintenanceCycleId

`func (o *MaintenanceCycleResponseModel) HasMaintenanceCycleId() bool`

HasMaintenanceCycleId returns a boolean if a field has been set.

### GetMaintenanceCycleDescription

`func (o *MaintenanceCycleResponseModel) GetMaintenanceCycleDescription() string`

GetMaintenanceCycleDescription returns the MaintenanceCycleDescription field if non-nil, zero value otherwise.

### GetMaintenanceCycleDescriptionOk

`func (o *MaintenanceCycleResponseModel) GetMaintenanceCycleDescriptionOk() (*string, bool)`

GetMaintenanceCycleDescriptionOk returns a tuple with the MaintenanceCycleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleDescription

`func (o *MaintenanceCycleResponseModel) SetMaintenanceCycleDescription(v string)`

SetMaintenanceCycleDescription sets MaintenanceCycleDescription field to given value.

### HasMaintenanceCycleDescription

`func (o *MaintenanceCycleResponseModel) HasMaintenanceCycleDescription() bool`

HasMaintenanceCycleDescription returns a boolean if a field has been set.

### SetMaintenanceCycleDescriptionNil

`func (o *MaintenanceCycleResponseModel) SetMaintenanceCycleDescriptionNil(b bool)`

 SetMaintenanceCycleDescriptionNil sets the value for MaintenanceCycleDescription to be an explicit nil

### UnsetMaintenanceCycleDescription
`func (o *MaintenanceCycleResponseModel) UnsetMaintenanceCycleDescription()`

UnsetMaintenanceCycleDescription ensures that no value is present for MaintenanceCycleDescription, not even an explicit nil
### GetMaintenanceCycleStatus

`func (o *MaintenanceCycleResponseModel) GetMaintenanceCycleStatus() MaintenanceCycleStatus`

GetMaintenanceCycleStatus returns the MaintenanceCycleStatus field if non-nil, zero value otherwise.

### GetMaintenanceCycleStatusOk

`func (o *MaintenanceCycleResponseModel) GetMaintenanceCycleStatusOk() (*MaintenanceCycleStatus, bool)`

GetMaintenanceCycleStatusOk returns a tuple with the MaintenanceCycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleStatus

`func (o *MaintenanceCycleResponseModel) SetMaintenanceCycleStatus(v MaintenanceCycleStatus)`

SetMaintenanceCycleStatus sets MaintenanceCycleStatus field to given value.

### HasMaintenanceCycleStatus

`func (o *MaintenanceCycleResponseModel) HasMaintenanceCycleStatus() bool`

HasMaintenanceCycleStatus returns a boolean if a field has been set.

### GetMaxDurationInMinutes

`func (o *MaintenanceCycleResponseModel) GetMaxDurationInMinutes() int32`

GetMaxDurationInMinutes returns the MaxDurationInMinutes field if non-nil, zero value otherwise.

### GetMaxDurationInMinutesOk

`func (o *MaintenanceCycleResponseModel) GetMaxDurationInMinutesOk() (*int32, bool)`

GetMaxDurationInMinutesOk returns a tuple with the MaxDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationInMinutes

`func (o *MaintenanceCycleResponseModel) SetMaxDurationInMinutes(v int32)`

SetMaxDurationInMinutes sets MaxDurationInMinutes field to given value.

### HasMaxDurationInMinutes

`func (o *MaintenanceCycleResponseModel) HasMaxDurationInMinutes() bool`

HasMaxDurationInMinutes returns a boolean if a field has been set.

### GetProvisioningSchemeName

`func (o *MaintenanceCycleResponseModel) GetProvisioningSchemeName() string`

GetProvisioningSchemeName returns the ProvisioningSchemeName field if non-nil, zero value otherwise.

### GetProvisioningSchemeNameOk

`func (o *MaintenanceCycleResponseModel) GetProvisioningSchemeNameOk() (*string, bool)`

GetProvisioningSchemeNameOk returns a tuple with the ProvisioningSchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeName

`func (o *MaintenanceCycleResponseModel) SetProvisioningSchemeName(v string)`

SetProvisioningSchemeName sets ProvisioningSchemeName field to given value.

### HasProvisioningSchemeName

`func (o *MaintenanceCycleResponseModel) HasProvisioningSchemeName() bool`

HasProvisioningSchemeName returns a boolean if a field has been set.

### SetProvisioningSchemeNameNil

`func (o *MaintenanceCycleResponseModel) SetProvisioningSchemeNameNil(b bool)`

 SetProvisioningSchemeNameNil sets the value for ProvisioningSchemeName to be an explicit nil

### UnsetProvisioningSchemeName
`func (o *MaintenanceCycleResponseModel) UnsetProvisioningSchemeName()`

UnsetProvisioningSchemeName ensures that no value is present for ProvisioningSchemeName, not even an explicit nil
### GetProvisioningSchemeUid

`func (o *MaintenanceCycleResponseModel) GetProvisioningSchemeUid() string`

GetProvisioningSchemeUid returns the ProvisioningSchemeUid field if non-nil, zero value otherwise.

### GetProvisioningSchemeUidOk

`func (o *MaintenanceCycleResponseModel) GetProvisioningSchemeUidOk() (*string, bool)`

GetProvisioningSchemeUidOk returns a tuple with the ProvisioningSchemeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeUid

`func (o *MaintenanceCycleResponseModel) SetProvisioningSchemeUid(v string)`

SetProvisioningSchemeUid sets ProvisioningSchemeUid field to given value.

### HasProvisioningSchemeUid

`func (o *MaintenanceCycleResponseModel) HasProvisioningSchemeUid() bool`

HasProvisioningSchemeUid returns a boolean if a field has been set.

### SetProvisioningSchemeUidNil

`func (o *MaintenanceCycleResponseModel) SetProvisioningSchemeUidNil(b bool)`

 SetProvisioningSchemeUidNil sets the value for ProvisioningSchemeUid to be an explicit nil

### UnsetProvisioningSchemeUid
`func (o *MaintenanceCycleResponseModel) UnsetProvisioningSchemeUid()`

UnsetProvisioningSchemeUid ensures that no value is present for ProvisioningSchemeUid, not even an explicit nil
### GetPurgeDBAfterInDays

`func (o *MaintenanceCycleResponseModel) GetPurgeDBAfterInDays() int32`

GetPurgeDBAfterInDays returns the PurgeDBAfterInDays field if non-nil, zero value otherwise.

### GetPurgeDBAfterInDaysOk

`func (o *MaintenanceCycleResponseModel) GetPurgeDBAfterInDaysOk() (*int32, bool)`

GetPurgeDBAfterInDaysOk returns a tuple with the PurgeDBAfterInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeDBAfterInDays

`func (o *MaintenanceCycleResponseModel) SetPurgeDBAfterInDays(v int32)`

SetPurgeDBAfterInDays sets PurgeDBAfterInDays field to given value.

### HasPurgeDBAfterInDays

`func (o *MaintenanceCycleResponseModel) HasPurgeDBAfterInDays() bool`

HasPurgeDBAfterInDays returns a boolean if a field has been set.

### SetPurgeDBAfterInDaysNil

`func (o *MaintenanceCycleResponseModel) SetPurgeDBAfterInDaysNil(b bool)`

 SetPurgeDBAfterInDaysNil sets the value for PurgeDBAfterInDays to be an explicit nil

### UnsetPurgeDBAfterInDays
`func (o *MaintenanceCycleResponseModel) UnsetPurgeDBAfterInDays()`

UnsetPurgeDBAfterInDays ensures that no value is present for PurgeDBAfterInDays, not even an explicit nil
### GetScheduledCompletionTimeInUTC

`func (o *MaintenanceCycleResponseModel) GetScheduledCompletionTimeInUTC() time.Time`

GetScheduledCompletionTimeInUTC returns the ScheduledCompletionTimeInUTC field if non-nil, zero value otherwise.

### GetScheduledCompletionTimeInUTCOk

`func (o *MaintenanceCycleResponseModel) GetScheduledCompletionTimeInUTCOk() (*time.Time, bool)`

GetScheduledCompletionTimeInUTCOk returns a tuple with the ScheduledCompletionTimeInUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCompletionTimeInUTC

`func (o *MaintenanceCycleResponseModel) SetScheduledCompletionTimeInUTC(v time.Time)`

SetScheduledCompletionTimeInUTC sets ScheduledCompletionTimeInUTC field to given value.

### HasScheduledCompletionTimeInUTC

`func (o *MaintenanceCycleResponseModel) HasScheduledCompletionTimeInUTC() bool`

HasScheduledCompletionTimeInUTC returns a boolean if a field has been set.

### SetScheduledCompletionTimeInUTCNil

`func (o *MaintenanceCycleResponseModel) SetScheduledCompletionTimeInUTCNil(b bool)`

 SetScheduledCompletionTimeInUTCNil sets the value for ScheduledCompletionTimeInUTC to be an explicit nil

### UnsetScheduledCompletionTimeInUTC
`func (o *MaintenanceCycleResponseModel) UnsetScheduledCompletionTimeInUTC()`

UnsetScheduledCompletionTimeInUTC ensures that no value is present for ScheduledCompletionTimeInUTC, not even an explicit nil
### GetScheduledStartTimeInUTC

`func (o *MaintenanceCycleResponseModel) GetScheduledStartTimeInUTC() time.Time`

GetScheduledStartTimeInUTC returns the ScheduledStartTimeInUTC field if non-nil, zero value otherwise.

### GetScheduledStartTimeInUTCOk

`func (o *MaintenanceCycleResponseModel) GetScheduledStartTimeInUTCOk() (*time.Time, bool)`

GetScheduledStartTimeInUTCOk returns a tuple with the ScheduledStartTimeInUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTimeInUTC

`func (o *MaintenanceCycleResponseModel) SetScheduledStartTimeInUTC(v time.Time)`

SetScheduledStartTimeInUTC sets ScheduledStartTimeInUTC field to given value.

### HasScheduledStartTimeInUTC

`func (o *MaintenanceCycleResponseModel) HasScheduledStartTimeInUTC() bool`

HasScheduledStartTimeInUTC returns a boolean if a field has been set.

### SetScheduledStartTimeInUTCNil

`func (o *MaintenanceCycleResponseModel) SetScheduledStartTimeInUTCNil(b bool)`

 SetScheduledStartTimeInUTCNil sets the value for ScheduledStartTimeInUTC to be an explicit nil

### UnsetScheduledStartTimeInUTC
`func (o *MaintenanceCycleResponseModel) UnsetScheduledStartTimeInUTC()`

UnsetScheduledStartTimeInUTC ensures that no value is present for ScheduledStartTimeInUTC, not even an explicit nil
### GetScheduleId

`func (o *MaintenanceCycleResponseModel) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *MaintenanceCycleResponseModel) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *MaintenanceCycleResponseModel) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *MaintenanceCycleResponseModel) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *MaintenanceCycleResponseModel) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *MaintenanceCycleResponseModel) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetSessionWarningLogOffTitle

`func (o *MaintenanceCycleResponseModel) GetSessionWarningLogOffTitle() string`

GetSessionWarningLogOffTitle returns the SessionWarningLogOffTitle field if non-nil, zero value otherwise.

### GetSessionWarningLogOffTitleOk

`func (o *MaintenanceCycleResponseModel) GetSessionWarningLogOffTitleOk() (*string, bool)`

GetSessionWarningLogOffTitleOk returns a tuple with the SessionWarningLogOffTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffTitle

`func (o *MaintenanceCycleResponseModel) SetSessionWarningLogOffTitle(v string)`

SetSessionWarningLogOffTitle sets SessionWarningLogOffTitle field to given value.

### HasSessionWarningLogOffTitle

`func (o *MaintenanceCycleResponseModel) HasSessionWarningLogOffTitle() bool`

HasSessionWarningLogOffTitle returns a boolean if a field has been set.

### SetSessionWarningLogOffTitleNil

`func (o *MaintenanceCycleResponseModel) SetSessionWarningLogOffTitleNil(b bool)`

 SetSessionWarningLogOffTitleNil sets the value for SessionWarningLogOffTitle to be an explicit nil

### UnsetSessionWarningLogOffTitle
`func (o *MaintenanceCycleResponseModel) UnsetSessionWarningLogOffTitle()`

UnsetSessionWarningLogOffTitle ensures that no value is present for SessionWarningLogOffTitle, not even an explicit nil
### GetSessionWarningLogOffMessage

`func (o *MaintenanceCycleResponseModel) GetSessionWarningLogOffMessage() string`

GetSessionWarningLogOffMessage returns the SessionWarningLogOffMessage field if non-nil, zero value otherwise.

### GetSessionWarningLogOffMessageOk

`func (o *MaintenanceCycleResponseModel) GetSessionWarningLogOffMessageOk() (*string, bool)`

GetSessionWarningLogOffMessageOk returns a tuple with the SessionWarningLogOffMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffMessage

`func (o *MaintenanceCycleResponseModel) SetSessionWarningLogOffMessage(v string)`

SetSessionWarningLogOffMessage sets SessionWarningLogOffMessage field to given value.

### HasSessionWarningLogOffMessage

`func (o *MaintenanceCycleResponseModel) HasSessionWarningLogOffMessage() bool`

HasSessionWarningLogOffMessage returns a boolean if a field has been set.

### SetSessionWarningLogOffMessageNil

`func (o *MaintenanceCycleResponseModel) SetSessionWarningLogOffMessageNil(b bool)`

 SetSessionWarningLogOffMessageNil sets the value for SessionWarningLogOffMessage to be an explicit nil

### UnsetSessionWarningLogOffMessage
`func (o *MaintenanceCycleResponseModel) UnsetSessionWarningLogOffMessage()`

UnsetSessionWarningLogOffMessage ensures that no value is present for SessionWarningLogOffMessage, not even an explicit nil
### GetSessionWarningTimeInMinutes

`func (o *MaintenanceCycleResponseModel) GetSessionWarningTimeInMinutes() int32`

GetSessionWarningTimeInMinutes returns the SessionWarningTimeInMinutes field if non-nil, zero value otherwise.

### GetSessionWarningTimeInMinutesOk

`func (o *MaintenanceCycleResponseModel) GetSessionWarningTimeInMinutesOk() (*int32, bool)`

GetSessionWarningTimeInMinutesOk returns a tuple with the SessionWarningTimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningTimeInMinutes

`func (o *MaintenanceCycleResponseModel) SetSessionWarningTimeInMinutes(v int32)`

SetSessionWarningTimeInMinutes sets SessionWarningTimeInMinutes field to given value.

### HasSessionWarningTimeInMinutes

`func (o *MaintenanceCycleResponseModel) HasSessionWarningTimeInMinutes() bool`

HasSessionWarningTimeInMinutes returns a boolean if a field has been set.

### SetSessionWarningTimeInMinutesNil

`func (o *MaintenanceCycleResponseModel) SetSessionWarningTimeInMinutesNil(b bool)`

 SetSessionWarningTimeInMinutesNil sets the value for SessionWarningTimeInMinutes to be an explicit nil

### UnsetSessionWarningTimeInMinutes
`func (o *MaintenanceCycleResponseModel) UnsetSessionWarningTimeInMinutes()`

UnsetSessionWarningTimeInMinutes ensures that no value is present for SessionWarningTimeInMinutes, not even an explicit nil
### GetOperation

`func (o *MaintenanceCycleResponseModel) GetOperation() MaintenanceCycleOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MaintenanceCycleResponseModel) GetOperationOk() (*MaintenanceCycleOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *MaintenanceCycleResponseModel) SetOperation(v MaintenanceCycleOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *MaintenanceCycleResponseModel) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetAllVM

`func (o *MaintenanceCycleResponseModel) GetAllVM() bool`

GetAllVM returns the AllVM field if non-nil, zero value otherwise.

### GetAllVMOk

`func (o *MaintenanceCycleResponseModel) GetAllVMOk() (*bool, bool)`

GetAllVMOk returns a tuple with the AllVM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllVM

`func (o *MaintenanceCycleResponseModel) SetAllVM(v bool)`

SetAllVM sets AllVM field to given value.

### HasAllVM

`func (o *MaintenanceCycleResponseModel) HasAllVM() bool`

HasAllVM returns a boolean if a field has been set.

### GetVMNames

`func (o *MaintenanceCycleResponseModel) GetVMNames() []string`

GetVMNames returns the VMNames field if non-nil, zero value otherwise.

### GetVMNamesOk

`func (o *MaintenanceCycleResponseModel) GetVMNamesOk() (*[]string, bool)`

GetVMNamesOk returns a tuple with the VMNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMNames

`func (o *MaintenanceCycleResponseModel) SetVMNames(v []string)`

SetVMNames sets VMNames field to given value.

### HasVMNames

`func (o *MaintenanceCycleResponseModel) HasVMNames() bool`

HasVMNames returns a boolean if a field has been set.

### SetVMNamesNil

`func (o *MaintenanceCycleResponseModel) SetVMNamesNil(b bool)`

 SetVMNamesNil sets the value for VMNames to be an explicit nil

### UnsetVMNames
`func (o *MaintenanceCycleResponseModel) UnsetVMNames()`

UnsetVMNames ensures that no value is present for VMNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


