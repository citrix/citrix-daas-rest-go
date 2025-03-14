# MaintenanceCycleDetailsResponseModel

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

### NewMaintenanceCycleDetailsResponseModel

`func NewMaintenanceCycleDetailsResponseModel() *MaintenanceCycleDetailsResponseModel`

NewMaintenanceCycleDetailsResponseModel instantiates a new MaintenanceCycleDetailsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceCycleDetailsResponseModelWithDefaults

`func NewMaintenanceCycleDetailsResponseModelWithDefaults() *MaintenanceCycleDetailsResponseModel`

NewMaintenanceCycleDetailsResponseModelWithDefaults instantiates a new MaintenanceCycleDetailsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceCycleName

`func (o *MaintenanceCycleDetailsResponseModel) GetMaintenanceCycleName() string`

GetMaintenanceCycleName returns the MaintenanceCycleName field if non-nil, zero value otherwise.

### GetMaintenanceCycleNameOk

`func (o *MaintenanceCycleDetailsResponseModel) GetMaintenanceCycleNameOk() (*string, bool)`

GetMaintenanceCycleNameOk returns a tuple with the MaintenanceCycleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleName

`func (o *MaintenanceCycleDetailsResponseModel) SetMaintenanceCycleName(v string)`

SetMaintenanceCycleName sets MaintenanceCycleName field to given value.

### HasMaintenanceCycleName

`func (o *MaintenanceCycleDetailsResponseModel) HasMaintenanceCycleName() bool`

HasMaintenanceCycleName returns a boolean if a field has been set.

### SetMaintenanceCycleNameNil

`func (o *MaintenanceCycleDetailsResponseModel) SetMaintenanceCycleNameNil(b bool)`

 SetMaintenanceCycleNameNil sets the value for MaintenanceCycleName to be an explicit nil

### UnsetMaintenanceCycleName
`func (o *MaintenanceCycleDetailsResponseModel) UnsetMaintenanceCycleName()`

UnsetMaintenanceCycleName ensures that no value is present for MaintenanceCycleName, not even an explicit nil
### GetMaintenanceCycleId

`func (o *MaintenanceCycleDetailsResponseModel) GetMaintenanceCycleId() string`

GetMaintenanceCycleId returns the MaintenanceCycleId field if non-nil, zero value otherwise.

### GetMaintenanceCycleIdOk

`func (o *MaintenanceCycleDetailsResponseModel) GetMaintenanceCycleIdOk() (*string, bool)`

GetMaintenanceCycleIdOk returns a tuple with the MaintenanceCycleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleId

`func (o *MaintenanceCycleDetailsResponseModel) SetMaintenanceCycleId(v string)`

SetMaintenanceCycleId sets MaintenanceCycleId field to given value.

### HasMaintenanceCycleId

`func (o *MaintenanceCycleDetailsResponseModel) HasMaintenanceCycleId() bool`

HasMaintenanceCycleId returns a boolean if a field has been set.

### GetMaintenanceCycleDescription

`func (o *MaintenanceCycleDetailsResponseModel) GetMaintenanceCycleDescription() string`

GetMaintenanceCycleDescription returns the MaintenanceCycleDescription field if non-nil, zero value otherwise.

### GetMaintenanceCycleDescriptionOk

`func (o *MaintenanceCycleDetailsResponseModel) GetMaintenanceCycleDescriptionOk() (*string, bool)`

GetMaintenanceCycleDescriptionOk returns a tuple with the MaintenanceCycleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleDescription

`func (o *MaintenanceCycleDetailsResponseModel) SetMaintenanceCycleDescription(v string)`

SetMaintenanceCycleDescription sets MaintenanceCycleDescription field to given value.

### HasMaintenanceCycleDescription

`func (o *MaintenanceCycleDetailsResponseModel) HasMaintenanceCycleDescription() bool`

HasMaintenanceCycleDescription returns a boolean if a field has been set.

### SetMaintenanceCycleDescriptionNil

`func (o *MaintenanceCycleDetailsResponseModel) SetMaintenanceCycleDescriptionNil(b bool)`

 SetMaintenanceCycleDescriptionNil sets the value for MaintenanceCycleDescription to be an explicit nil

### UnsetMaintenanceCycleDescription
`func (o *MaintenanceCycleDetailsResponseModel) UnsetMaintenanceCycleDescription()`

UnsetMaintenanceCycleDescription ensures that no value is present for MaintenanceCycleDescription, not even an explicit nil
### GetMaintenanceCycleStatus

`func (o *MaintenanceCycleDetailsResponseModel) GetMaintenanceCycleStatus() MaintenanceCycleStatus`

GetMaintenanceCycleStatus returns the MaintenanceCycleStatus field if non-nil, zero value otherwise.

### GetMaintenanceCycleStatusOk

`func (o *MaintenanceCycleDetailsResponseModel) GetMaintenanceCycleStatusOk() (*MaintenanceCycleStatus, bool)`

GetMaintenanceCycleStatusOk returns a tuple with the MaintenanceCycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleStatus

`func (o *MaintenanceCycleDetailsResponseModel) SetMaintenanceCycleStatus(v MaintenanceCycleStatus)`

SetMaintenanceCycleStatus sets MaintenanceCycleStatus field to given value.

### HasMaintenanceCycleStatus

`func (o *MaintenanceCycleDetailsResponseModel) HasMaintenanceCycleStatus() bool`

HasMaintenanceCycleStatus returns a boolean if a field has been set.

### GetMaxDurationInMinutes

`func (o *MaintenanceCycleDetailsResponseModel) GetMaxDurationInMinutes() int32`

GetMaxDurationInMinutes returns the MaxDurationInMinutes field if non-nil, zero value otherwise.

### GetMaxDurationInMinutesOk

`func (o *MaintenanceCycleDetailsResponseModel) GetMaxDurationInMinutesOk() (*int32, bool)`

GetMaxDurationInMinutesOk returns a tuple with the MaxDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationInMinutes

`func (o *MaintenanceCycleDetailsResponseModel) SetMaxDurationInMinutes(v int32)`

SetMaxDurationInMinutes sets MaxDurationInMinutes field to given value.

### HasMaxDurationInMinutes

`func (o *MaintenanceCycleDetailsResponseModel) HasMaxDurationInMinutes() bool`

HasMaxDurationInMinutes returns a boolean if a field has been set.

### GetProvisioningSchemeName

`func (o *MaintenanceCycleDetailsResponseModel) GetProvisioningSchemeName() string`

GetProvisioningSchemeName returns the ProvisioningSchemeName field if non-nil, zero value otherwise.

### GetProvisioningSchemeNameOk

`func (o *MaintenanceCycleDetailsResponseModel) GetProvisioningSchemeNameOk() (*string, bool)`

GetProvisioningSchemeNameOk returns a tuple with the ProvisioningSchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeName

`func (o *MaintenanceCycleDetailsResponseModel) SetProvisioningSchemeName(v string)`

SetProvisioningSchemeName sets ProvisioningSchemeName field to given value.

### HasProvisioningSchemeName

`func (o *MaintenanceCycleDetailsResponseModel) HasProvisioningSchemeName() bool`

HasProvisioningSchemeName returns a boolean if a field has been set.

### SetProvisioningSchemeNameNil

`func (o *MaintenanceCycleDetailsResponseModel) SetProvisioningSchemeNameNil(b bool)`

 SetProvisioningSchemeNameNil sets the value for ProvisioningSchemeName to be an explicit nil

### UnsetProvisioningSchemeName
`func (o *MaintenanceCycleDetailsResponseModel) UnsetProvisioningSchemeName()`

UnsetProvisioningSchemeName ensures that no value is present for ProvisioningSchemeName, not even an explicit nil
### GetProvisioningSchemeUid

`func (o *MaintenanceCycleDetailsResponseModel) GetProvisioningSchemeUid() string`

GetProvisioningSchemeUid returns the ProvisioningSchemeUid field if non-nil, zero value otherwise.

### GetProvisioningSchemeUidOk

`func (o *MaintenanceCycleDetailsResponseModel) GetProvisioningSchemeUidOk() (*string, bool)`

GetProvisioningSchemeUidOk returns a tuple with the ProvisioningSchemeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeUid

`func (o *MaintenanceCycleDetailsResponseModel) SetProvisioningSchemeUid(v string)`

SetProvisioningSchemeUid sets ProvisioningSchemeUid field to given value.

### HasProvisioningSchemeUid

`func (o *MaintenanceCycleDetailsResponseModel) HasProvisioningSchemeUid() bool`

HasProvisioningSchemeUid returns a boolean if a field has been set.

### SetProvisioningSchemeUidNil

`func (o *MaintenanceCycleDetailsResponseModel) SetProvisioningSchemeUidNil(b bool)`

 SetProvisioningSchemeUidNil sets the value for ProvisioningSchemeUid to be an explicit nil

### UnsetProvisioningSchemeUid
`func (o *MaintenanceCycleDetailsResponseModel) UnsetProvisioningSchemeUid()`

UnsetProvisioningSchemeUid ensures that no value is present for ProvisioningSchemeUid, not even an explicit nil
### GetPurgeDBAfterInDays

`func (o *MaintenanceCycleDetailsResponseModel) GetPurgeDBAfterInDays() int32`

GetPurgeDBAfterInDays returns the PurgeDBAfterInDays field if non-nil, zero value otherwise.

### GetPurgeDBAfterInDaysOk

`func (o *MaintenanceCycleDetailsResponseModel) GetPurgeDBAfterInDaysOk() (*int32, bool)`

GetPurgeDBAfterInDaysOk returns a tuple with the PurgeDBAfterInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeDBAfterInDays

`func (o *MaintenanceCycleDetailsResponseModel) SetPurgeDBAfterInDays(v int32)`

SetPurgeDBAfterInDays sets PurgeDBAfterInDays field to given value.

### HasPurgeDBAfterInDays

`func (o *MaintenanceCycleDetailsResponseModel) HasPurgeDBAfterInDays() bool`

HasPurgeDBAfterInDays returns a boolean if a field has been set.

### SetPurgeDBAfterInDaysNil

`func (o *MaintenanceCycleDetailsResponseModel) SetPurgeDBAfterInDaysNil(b bool)`

 SetPurgeDBAfterInDaysNil sets the value for PurgeDBAfterInDays to be an explicit nil

### UnsetPurgeDBAfterInDays
`func (o *MaintenanceCycleDetailsResponseModel) UnsetPurgeDBAfterInDays()`

UnsetPurgeDBAfterInDays ensures that no value is present for PurgeDBAfterInDays, not even an explicit nil
### GetScheduledCompletionTimeInUTC

`func (o *MaintenanceCycleDetailsResponseModel) GetScheduledCompletionTimeInUTC() time.Time`

GetScheduledCompletionTimeInUTC returns the ScheduledCompletionTimeInUTC field if non-nil, zero value otherwise.

### GetScheduledCompletionTimeInUTCOk

`func (o *MaintenanceCycleDetailsResponseModel) GetScheduledCompletionTimeInUTCOk() (*time.Time, bool)`

GetScheduledCompletionTimeInUTCOk returns a tuple with the ScheduledCompletionTimeInUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCompletionTimeInUTC

`func (o *MaintenanceCycleDetailsResponseModel) SetScheduledCompletionTimeInUTC(v time.Time)`

SetScheduledCompletionTimeInUTC sets ScheduledCompletionTimeInUTC field to given value.

### HasScheduledCompletionTimeInUTC

`func (o *MaintenanceCycleDetailsResponseModel) HasScheduledCompletionTimeInUTC() bool`

HasScheduledCompletionTimeInUTC returns a boolean if a field has been set.

### SetScheduledCompletionTimeInUTCNil

`func (o *MaintenanceCycleDetailsResponseModel) SetScheduledCompletionTimeInUTCNil(b bool)`

 SetScheduledCompletionTimeInUTCNil sets the value for ScheduledCompletionTimeInUTC to be an explicit nil

### UnsetScheduledCompletionTimeInUTC
`func (o *MaintenanceCycleDetailsResponseModel) UnsetScheduledCompletionTimeInUTC()`

UnsetScheduledCompletionTimeInUTC ensures that no value is present for ScheduledCompletionTimeInUTC, not even an explicit nil
### GetScheduledStartTimeInUTC

`func (o *MaintenanceCycleDetailsResponseModel) GetScheduledStartTimeInUTC() time.Time`

GetScheduledStartTimeInUTC returns the ScheduledStartTimeInUTC field if non-nil, zero value otherwise.

### GetScheduledStartTimeInUTCOk

`func (o *MaintenanceCycleDetailsResponseModel) GetScheduledStartTimeInUTCOk() (*time.Time, bool)`

GetScheduledStartTimeInUTCOk returns a tuple with the ScheduledStartTimeInUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTimeInUTC

`func (o *MaintenanceCycleDetailsResponseModel) SetScheduledStartTimeInUTC(v time.Time)`

SetScheduledStartTimeInUTC sets ScheduledStartTimeInUTC field to given value.

### HasScheduledStartTimeInUTC

`func (o *MaintenanceCycleDetailsResponseModel) HasScheduledStartTimeInUTC() bool`

HasScheduledStartTimeInUTC returns a boolean if a field has been set.

### SetScheduledStartTimeInUTCNil

`func (o *MaintenanceCycleDetailsResponseModel) SetScheduledStartTimeInUTCNil(b bool)`

 SetScheduledStartTimeInUTCNil sets the value for ScheduledStartTimeInUTC to be an explicit nil

### UnsetScheduledStartTimeInUTC
`func (o *MaintenanceCycleDetailsResponseModel) UnsetScheduledStartTimeInUTC()`

UnsetScheduledStartTimeInUTC ensures that no value is present for ScheduledStartTimeInUTC, not even an explicit nil
### GetScheduleId

`func (o *MaintenanceCycleDetailsResponseModel) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *MaintenanceCycleDetailsResponseModel) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *MaintenanceCycleDetailsResponseModel) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *MaintenanceCycleDetailsResponseModel) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *MaintenanceCycleDetailsResponseModel) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *MaintenanceCycleDetailsResponseModel) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetSessionWarningLogOffTitle

`func (o *MaintenanceCycleDetailsResponseModel) GetSessionWarningLogOffTitle() string`

GetSessionWarningLogOffTitle returns the SessionWarningLogOffTitle field if non-nil, zero value otherwise.

### GetSessionWarningLogOffTitleOk

`func (o *MaintenanceCycleDetailsResponseModel) GetSessionWarningLogOffTitleOk() (*string, bool)`

GetSessionWarningLogOffTitleOk returns a tuple with the SessionWarningLogOffTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffTitle

`func (o *MaintenanceCycleDetailsResponseModel) SetSessionWarningLogOffTitle(v string)`

SetSessionWarningLogOffTitle sets SessionWarningLogOffTitle field to given value.

### HasSessionWarningLogOffTitle

`func (o *MaintenanceCycleDetailsResponseModel) HasSessionWarningLogOffTitle() bool`

HasSessionWarningLogOffTitle returns a boolean if a field has been set.

### SetSessionWarningLogOffTitleNil

`func (o *MaintenanceCycleDetailsResponseModel) SetSessionWarningLogOffTitleNil(b bool)`

 SetSessionWarningLogOffTitleNil sets the value for SessionWarningLogOffTitle to be an explicit nil

### UnsetSessionWarningLogOffTitle
`func (o *MaintenanceCycleDetailsResponseModel) UnsetSessionWarningLogOffTitle()`

UnsetSessionWarningLogOffTitle ensures that no value is present for SessionWarningLogOffTitle, not even an explicit nil
### GetSessionWarningLogOffMessage

`func (o *MaintenanceCycleDetailsResponseModel) GetSessionWarningLogOffMessage() string`

GetSessionWarningLogOffMessage returns the SessionWarningLogOffMessage field if non-nil, zero value otherwise.

### GetSessionWarningLogOffMessageOk

`func (o *MaintenanceCycleDetailsResponseModel) GetSessionWarningLogOffMessageOk() (*string, bool)`

GetSessionWarningLogOffMessageOk returns a tuple with the SessionWarningLogOffMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffMessage

`func (o *MaintenanceCycleDetailsResponseModel) SetSessionWarningLogOffMessage(v string)`

SetSessionWarningLogOffMessage sets SessionWarningLogOffMessage field to given value.

### HasSessionWarningLogOffMessage

`func (o *MaintenanceCycleDetailsResponseModel) HasSessionWarningLogOffMessage() bool`

HasSessionWarningLogOffMessage returns a boolean if a field has been set.

### SetSessionWarningLogOffMessageNil

`func (o *MaintenanceCycleDetailsResponseModel) SetSessionWarningLogOffMessageNil(b bool)`

 SetSessionWarningLogOffMessageNil sets the value for SessionWarningLogOffMessage to be an explicit nil

### UnsetSessionWarningLogOffMessage
`func (o *MaintenanceCycleDetailsResponseModel) UnsetSessionWarningLogOffMessage()`

UnsetSessionWarningLogOffMessage ensures that no value is present for SessionWarningLogOffMessage, not even an explicit nil
### GetSessionWarningTimeInMinutes

`func (o *MaintenanceCycleDetailsResponseModel) GetSessionWarningTimeInMinutes() int32`

GetSessionWarningTimeInMinutes returns the SessionWarningTimeInMinutes field if non-nil, zero value otherwise.

### GetSessionWarningTimeInMinutesOk

`func (o *MaintenanceCycleDetailsResponseModel) GetSessionWarningTimeInMinutesOk() (*int32, bool)`

GetSessionWarningTimeInMinutesOk returns a tuple with the SessionWarningTimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningTimeInMinutes

`func (o *MaintenanceCycleDetailsResponseModel) SetSessionWarningTimeInMinutes(v int32)`

SetSessionWarningTimeInMinutes sets SessionWarningTimeInMinutes field to given value.

### HasSessionWarningTimeInMinutes

`func (o *MaintenanceCycleDetailsResponseModel) HasSessionWarningTimeInMinutes() bool`

HasSessionWarningTimeInMinutes returns a boolean if a field has been set.

### SetSessionWarningTimeInMinutesNil

`func (o *MaintenanceCycleDetailsResponseModel) SetSessionWarningTimeInMinutesNil(b bool)`

 SetSessionWarningTimeInMinutesNil sets the value for SessionWarningTimeInMinutes to be an explicit nil

### UnsetSessionWarningTimeInMinutes
`func (o *MaintenanceCycleDetailsResponseModel) UnsetSessionWarningTimeInMinutes()`

UnsetSessionWarningTimeInMinutes ensures that no value is present for SessionWarningTimeInMinutes, not even an explicit nil
### GetOperation

`func (o *MaintenanceCycleDetailsResponseModel) GetOperation() MaintenanceCycleOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MaintenanceCycleDetailsResponseModel) GetOperationOk() (*MaintenanceCycleOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *MaintenanceCycleDetailsResponseModel) SetOperation(v MaintenanceCycleOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *MaintenanceCycleDetailsResponseModel) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetAllVM

`func (o *MaintenanceCycleDetailsResponseModel) GetAllVM() bool`

GetAllVM returns the AllVM field if non-nil, zero value otherwise.

### GetAllVMOk

`func (o *MaintenanceCycleDetailsResponseModel) GetAllVMOk() (*bool, bool)`

GetAllVMOk returns a tuple with the AllVM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllVM

`func (o *MaintenanceCycleDetailsResponseModel) SetAllVM(v bool)`

SetAllVM sets AllVM field to given value.

### HasAllVM

`func (o *MaintenanceCycleDetailsResponseModel) HasAllVM() bool`

HasAllVM returns a boolean if a field has been set.

### GetVMNames

`func (o *MaintenanceCycleDetailsResponseModel) GetVMNames() []string`

GetVMNames returns the VMNames field if non-nil, zero value otherwise.

### GetVMNamesOk

`func (o *MaintenanceCycleDetailsResponseModel) GetVMNamesOk() (*[]string, bool)`

GetVMNamesOk returns a tuple with the VMNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMNames

`func (o *MaintenanceCycleDetailsResponseModel) SetVMNames(v []string)`

SetVMNames sets VMNames field to given value.

### HasVMNames

`func (o *MaintenanceCycleDetailsResponseModel) HasVMNames() bool`

HasVMNames returns a boolean if a field has been set.

### SetVMNamesNil

`func (o *MaintenanceCycleDetailsResponseModel) SetVMNamesNil(b bool)`

 SetVMNamesNil sets the value for VMNames to be an explicit nil

### UnsetVMNames
`func (o *MaintenanceCycleDetailsResponseModel) UnsetVMNames()`

UnsetVMNames ensures that no value is present for VMNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


