# CreateMaintenanceCycleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceCycleName** | Pointer to **NullableString** | The name of the Maintenance Cycle. Name is globally unique. | [optional] 
**MaintenanceCycleDescription** | Pointer to **NullableString** | The description of the Maintenance Cycle. | [optional] 
**ProvisioningSchemeName** | Pointer to **NullableString** | The name of provisioning scheme. | [optional] 
**ProvisioningSchemeUid** | Pointer to **NullableString** | The uid of provisioning scheme. | [optional] 
**ScheduledStartTimeInUTC** | Pointer to **NullableTime** | Maintenance cycle start time in UTC. | [optional] 
**MaxDurationInMinutes** | **int32** | The max duration allowed for maintenance cycle. | 
**PurgeDBAfterInDays** | Pointer to **NullableInt32** | Purge history data after a certain number of days. | [optional] 
**SessionWarningLogOffTitle** | Pointer to **NullableString** | The title of session log off warning.  | [optional] 
**SessionWarningLogOffMessage** | Pointer to **NullableString** | The message of session log off warning. | [optional] 
**SessionWarningTimeInMinutes** | Pointer to **NullableInt32** | The countdown duration for session log off.  | [optional] 
**Operation** | [**MaintenanceCycleOperation2**](MaintenanceCycleOperation2.md) |  | 
**VMNames** | Pointer to **[]string** | The vm names in maintenance cycle.  | [optional] 

## Methods

### NewCreateMaintenanceCycleRequestModel

`func NewCreateMaintenanceCycleRequestModel(maxDurationInMinutes int32, operation MaintenanceCycleOperation2, ) *CreateMaintenanceCycleRequestModel`

NewCreateMaintenanceCycleRequestModel instantiates a new CreateMaintenanceCycleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMaintenanceCycleRequestModelWithDefaults

`func NewCreateMaintenanceCycleRequestModelWithDefaults() *CreateMaintenanceCycleRequestModel`

NewCreateMaintenanceCycleRequestModelWithDefaults instantiates a new CreateMaintenanceCycleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceCycleName

`func (o *CreateMaintenanceCycleRequestModel) GetMaintenanceCycleName() string`

GetMaintenanceCycleName returns the MaintenanceCycleName field if non-nil, zero value otherwise.

### GetMaintenanceCycleNameOk

`func (o *CreateMaintenanceCycleRequestModel) GetMaintenanceCycleNameOk() (*string, bool)`

GetMaintenanceCycleNameOk returns a tuple with the MaintenanceCycleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleName

`func (o *CreateMaintenanceCycleRequestModel) SetMaintenanceCycleName(v string)`

SetMaintenanceCycleName sets MaintenanceCycleName field to given value.

### HasMaintenanceCycleName

`func (o *CreateMaintenanceCycleRequestModel) HasMaintenanceCycleName() bool`

HasMaintenanceCycleName returns a boolean if a field has been set.

### SetMaintenanceCycleNameNil

`func (o *CreateMaintenanceCycleRequestModel) SetMaintenanceCycleNameNil(b bool)`

 SetMaintenanceCycleNameNil sets the value for MaintenanceCycleName to be an explicit nil

### UnsetMaintenanceCycleName
`func (o *CreateMaintenanceCycleRequestModel) UnsetMaintenanceCycleName()`

UnsetMaintenanceCycleName ensures that no value is present for MaintenanceCycleName, not even an explicit nil
### GetMaintenanceCycleDescription

`func (o *CreateMaintenanceCycleRequestModel) GetMaintenanceCycleDescription() string`

GetMaintenanceCycleDescription returns the MaintenanceCycleDescription field if non-nil, zero value otherwise.

### GetMaintenanceCycleDescriptionOk

`func (o *CreateMaintenanceCycleRequestModel) GetMaintenanceCycleDescriptionOk() (*string, bool)`

GetMaintenanceCycleDescriptionOk returns a tuple with the MaintenanceCycleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCycleDescription

`func (o *CreateMaintenanceCycleRequestModel) SetMaintenanceCycleDescription(v string)`

SetMaintenanceCycleDescription sets MaintenanceCycleDescription field to given value.

### HasMaintenanceCycleDescription

`func (o *CreateMaintenanceCycleRequestModel) HasMaintenanceCycleDescription() bool`

HasMaintenanceCycleDescription returns a boolean if a field has been set.

### SetMaintenanceCycleDescriptionNil

`func (o *CreateMaintenanceCycleRequestModel) SetMaintenanceCycleDescriptionNil(b bool)`

 SetMaintenanceCycleDescriptionNil sets the value for MaintenanceCycleDescription to be an explicit nil

### UnsetMaintenanceCycleDescription
`func (o *CreateMaintenanceCycleRequestModel) UnsetMaintenanceCycleDescription()`

UnsetMaintenanceCycleDescription ensures that no value is present for MaintenanceCycleDescription, not even an explicit nil
### GetProvisioningSchemeName

`func (o *CreateMaintenanceCycleRequestModel) GetProvisioningSchemeName() string`

GetProvisioningSchemeName returns the ProvisioningSchemeName field if non-nil, zero value otherwise.

### GetProvisioningSchemeNameOk

`func (o *CreateMaintenanceCycleRequestModel) GetProvisioningSchemeNameOk() (*string, bool)`

GetProvisioningSchemeNameOk returns a tuple with the ProvisioningSchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeName

`func (o *CreateMaintenanceCycleRequestModel) SetProvisioningSchemeName(v string)`

SetProvisioningSchemeName sets ProvisioningSchemeName field to given value.

### HasProvisioningSchemeName

`func (o *CreateMaintenanceCycleRequestModel) HasProvisioningSchemeName() bool`

HasProvisioningSchemeName returns a boolean if a field has been set.

### SetProvisioningSchemeNameNil

`func (o *CreateMaintenanceCycleRequestModel) SetProvisioningSchemeNameNil(b bool)`

 SetProvisioningSchemeNameNil sets the value for ProvisioningSchemeName to be an explicit nil

### UnsetProvisioningSchemeName
`func (o *CreateMaintenanceCycleRequestModel) UnsetProvisioningSchemeName()`

UnsetProvisioningSchemeName ensures that no value is present for ProvisioningSchemeName, not even an explicit nil
### GetProvisioningSchemeUid

`func (o *CreateMaintenanceCycleRequestModel) GetProvisioningSchemeUid() string`

GetProvisioningSchemeUid returns the ProvisioningSchemeUid field if non-nil, zero value otherwise.

### GetProvisioningSchemeUidOk

`func (o *CreateMaintenanceCycleRequestModel) GetProvisioningSchemeUidOk() (*string, bool)`

GetProvisioningSchemeUidOk returns a tuple with the ProvisioningSchemeUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeUid

`func (o *CreateMaintenanceCycleRequestModel) SetProvisioningSchemeUid(v string)`

SetProvisioningSchemeUid sets ProvisioningSchemeUid field to given value.

### HasProvisioningSchemeUid

`func (o *CreateMaintenanceCycleRequestModel) HasProvisioningSchemeUid() bool`

HasProvisioningSchemeUid returns a boolean if a field has been set.

### SetProvisioningSchemeUidNil

`func (o *CreateMaintenanceCycleRequestModel) SetProvisioningSchemeUidNil(b bool)`

 SetProvisioningSchemeUidNil sets the value for ProvisioningSchemeUid to be an explicit nil

### UnsetProvisioningSchemeUid
`func (o *CreateMaintenanceCycleRequestModel) UnsetProvisioningSchemeUid()`

UnsetProvisioningSchemeUid ensures that no value is present for ProvisioningSchemeUid, not even an explicit nil
### GetScheduledStartTimeInUTC

`func (o *CreateMaintenanceCycleRequestModel) GetScheduledStartTimeInUTC() time.Time`

GetScheduledStartTimeInUTC returns the ScheduledStartTimeInUTC field if non-nil, zero value otherwise.

### GetScheduledStartTimeInUTCOk

`func (o *CreateMaintenanceCycleRequestModel) GetScheduledStartTimeInUTCOk() (*time.Time, bool)`

GetScheduledStartTimeInUTCOk returns a tuple with the ScheduledStartTimeInUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTimeInUTC

`func (o *CreateMaintenanceCycleRequestModel) SetScheduledStartTimeInUTC(v time.Time)`

SetScheduledStartTimeInUTC sets ScheduledStartTimeInUTC field to given value.

### HasScheduledStartTimeInUTC

`func (o *CreateMaintenanceCycleRequestModel) HasScheduledStartTimeInUTC() bool`

HasScheduledStartTimeInUTC returns a boolean if a field has been set.

### SetScheduledStartTimeInUTCNil

`func (o *CreateMaintenanceCycleRequestModel) SetScheduledStartTimeInUTCNil(b bool)`

 SetScheduledStartTimeInUTCNil sets the value for ScheduledStartTimeInUTC to be an explicit nil

### UnsetScheduledStartTimeInUTC
`func (o *CreateMaintenanceCycleRequestModel) UnsetScheduledStartTimeInUTC()`

UnsetScheduledStartTimeInUTC ensures that no value is present for ScheduledStartTimeInUTC, not even an explicit nil
### GetMaxDurationInMinutes

`func (o *CreateMaintenanceCycleRequestModel) GetMaxDurationInMinutes() int32`

GetMaxDurationInMinutes returns the MaxDurationInMinutes field if non-nil, zero value otherwise.

### GetMaxDurationInMinutesOk

`func (o *CreateMaintenanceCycleRequestModel) GetMaxDurationInMinutesOk() (*int32, bool)`

GetMaxDurationInMinutesOk returns a tuple with the MaxDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationInMinutes

`func (o *CreateMaintenanceCycleRequestModel) SetMaxDurationInMinutes(v int32)`

SetMaxDurationInMinutes sets MaxDurationInMinutes field to given value.


### GetPurgeDBAfterInDays

`func (o *CreateMaintenanceCycleRequestModel) GetPurgeDBAfterInDays() int32`

GetPurgeDBAfterInDays returns the PurgeDBAfterInDays field if non-nil, zero value otherwise.

### GetPurgeDBAfterInDaysOk

`func (o *CreateMaintenanceCycleRequestModel) GetPurgeDBAfterInDaysOk() (*int32, bool)`

GetPurgeDBAfterInDaysOk returns a tuple with the PurgeDBAfterInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeDBAfterInDays

`func (o *CreateMaintenanceCycleRequestModel) SetPurgeDBAfterInDays(v int32)`

SetPurgeDBAfterInDays sets PurgeDBAfterInDays field to given value.

### HasPurgeDBAfterInDays

`func (o *CreateMaintenanceCycleRequestModel) HasPurgeDBAfterInDays() bool`

HasPurgeDBAfterInDays returns a boolean if a field has been set.

### SetPurgeDBAfterInDaysNil

`func (o *CreateMaintenanceCycleRequestModel) SetPurgeDBAfterInDaysNil(b bool)`

 SetPurgeDBAfterInDaysNil sets the value for PurgeDBAfterInDays to be an explicit nil

### UnsetPurgeDBAfterInDays
`func (o *CreateMaintenanceCycleRequestModel) UnsetPurgeDBAfterInDays()`

UnsetPurgeDBAfterInDays ensures that no value is present for PurgeDBAfterInDays, not even an explicit nil
### GetSessionWarningLogOffTitle

`func (o *CreateMaintenanceCycleRequestModel) GetSessionWarningLogOffTitle() string`

GetSessionWarningLogOffTitle returns the SessionWarningLogOffTitle field if non-nil, zero value otherwise.

### GetSessionWarningLogOffTitleOk

`func (o *CreateMaintenanceCycleRequestModel) GetSessionWarningLogOffTitleOk() (*string, bool)`

GetSessionWarningLogOffTitleOk returns a tuple with the SessionWarningLogOffTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffTitle

`func (o *CreateMaintenanceCycleRequestModel) SetSessionWarningLogOffTitle(v string)`

SetSessionWarningLogOffTitle sets SessionWarningLogOffTitle field to given value.

### HasSessionWarningLogOffTitle

`func (o *CreateMaintenanceCycleRequestModel) HasSessionWarningLogOffTitle() bool`

HasSessionWarningLogOffTitle returns a boolean if a field has been set.

### SetSessionWarningLogOffTitleNil

`func (o *CreateMaintenanceCycleRequestModel) SetSessionWarningLogOffTitleNil(b bool)`

 SetSessionWarningLogOffTitleNil sets the value for SessionWarningLogOffTitle to be an explicit nil

### UnsetSessionWarningLogOffTitle
`func (o *CreateMaintenanceCycleRequestModel) UnsetSessionWarningLogOffTitle()`

UnsetSessionWarningLogOffTitle ensures that no value is present for SessionWarningLogOffTitle, not even an explicit nil
### GetSessionWarningLogOffMessage

`func (o *CreateMaintenanceCycleRequestModel) GetSessionWarningLogOffMessage() string`

GetSessionWarningLogOffMessage returns the SessionWarningLogOffMessage field if non-nil, zero value otherwise.

### GetSessionWarningLogOffMessageOk

`func (o *CreateMaintenanceCycleRequestModel) GetSessionWarningLogOffMessageOk() (*string, bool)`

GetSessionWarningLogOffMessageOk returns a tuple with the SessionWarningLogOffMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffMessage

`func (o *CreateMaintenanceCycleRequestModel) SetSessionWarningLogOffMessage(v string)`

SetSessionWarningLogOffMessage sets SessionWarningLogOffMessage field to given value.

### HasSessionWarningLogOffMessage

`func (o *CreateMaintenanceCycleRequestModel) HasSessionWarningLogOffMessage() bool`

HasSessionWarningLogOffMessage returns a boolean if a field has been set.

### SetSessionWarningLogOffMessageNil

`func (o *CreateMaintenanceCycleRequestModel) SetSessionWarningLogOffMessageNil(b bool)`

 SetSessionWarningLogOffMessageNil sets the value for SessionWarningLogOffMessage to be an explicit nil

### UnsetSessionWarningLogOffMessage
`func (o *CreateMaintenanceCycleRequestModel) UnsetSessionWarningLogOffMessage()`

UnsetSessionWarningLogOffMessage ensures that no value is present for SessionWarningLogOffMessage, not even an explicit nil
### GetSessionWarningTimeInMinutes

`func (o *CreateMaintenanceCycleRequestModel) GetSessionWarningTimeInMinutes() int32`

GetSessionWarningTimeInMinutes returns the SessionWarningTimeInMinutes field if non-nil, zero value otherwise.

### GetSessionWarningTimeInMinutesOk

`func (o *CreateMaintenanceCycleRequestModel) GetSessionWarningTimeInMinutesOk() (*int32, bool)`

GetSessionWarningTimeInMinutesOk returns a tuple with the SessionWarningTimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningTimeInMinutes

`func (o *CreateMaintenanceCycleRequestModel) SetSessionWarningTimeInMinutes(v int32)`

SetSessionWarningTimeInMinutes sets SessionWarningTimeInMinutes field to given value.

### HasSessionWarningTimeInMinutes

`func (o *CreateMaintenanceCycleRequestModel) HasSessionWarningTimeInMinutes() bool`

HasSessionWarningTimeInMinutes returns a boolean if a field has been set.

### SetSessionWarningTimeInMinutesNil

`func (o *CreateMaintenanceCycleRequestModel) SetSessionWarningTimeInMinutesNil(b bool)`

 SetSessionWarningTimeInMinutesNil sets the value for SessionWarningTimeInMinutes to be an explicit nil

### UnsetSessionWarningTimeInMinutes
`func (o *CreateMaintenanceCycleRequestModel) UnsetSessionWarningTimeInMinutes()`

UnsetSessionWarningTimeInMinutes ensures that no value is present for SessionWarningTimeInMinutes, not even an explicit nil
### GetOperation

`func (o *CreateMaintenanceCycleRequestModel) GetOperation() MaintenanceCycleOperation2`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CreateMaintenanceCycleRequestModel) GetOperationOk() (*MaintenanceCycleOperation2, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CreateMaintenanceCycleRequestModel) SetOperation(v MaintenanceCycleOperation2)`

SetOperation sets Operation field to given value.


### GetVMNames

`func (o *CreateMaintenanceCycleRequestModel) GetVMNames() []string`

GetVMNames returns the VMNames field if non-nil, zero value otherwise.

### GetVMNamesOk

`func (o *CreateMaintenanceCycleRequestModel) GetVMNamesOk() (*[]string, bool)`

GetVMNamesOk returns a tuple with the VMNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMNames

`func (o *CreateMaintenanceCycleRequestModel) SetVMNames(v []string)`

SetVMNames sets VMNames field to given value.

### HasVMNames

`func (o *CreateMaintenanceCycleRequestModel) HasVMNames() bool`

HasVMNames returns a boolean if a field has been set.

### SetVMNamesNil

`func (o *CreateMaintenanceCycleRequestModel) SetVMNamesNil(b bool)`

 SetVMNamesNil sets the value for VMNames to be an explicit nil

### UnsetVMNames
`func (o *CreateMaintenanceCycleRequestModel) UnsetVMNames()`

UnsetVMNames ensures that no value is present for VMNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


