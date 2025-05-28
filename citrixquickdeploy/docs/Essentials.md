# Essentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertId** | Pointer to **string** |  | [optional] 
**AlertRule** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**SignalType** | Pointer to **string** |  | [optional] 
**MonitorCondition** | Pointer to **string** |  | [optional] 
**MonitoringService** | Pointer to **string** |  | [optional] 
**AlertTargetIDs** | Pointer to **[]string** |  | [optional] 
**OriginAlertId** | Pointer to **string** |  | [optional] 
**FiredDateTime** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EssentialsVersion** | Pointer to **string** |  | [optional] 
**AlertContextVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewEssentials

`func NewEssentials() *Essentials`

NewEssentials instantiates a new Essentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEssentialsWithDefaults

`func NewEssentialsWithDefaults() *Essentials`

NewEssentialsWithDefaults instantiates a new Essentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertId

`func (o *Essentials) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *Essentials) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *Essentials) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *Essentials) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetAlertRule

`func (o *Essentials) GetAlertRule() string`

GetAlertRule returns the AlertRule field if non-nil, zero value otherwise.

### GetAlertRuleOk

`func (o *Essentials) GetAlertRuleOk() (*string, bool)`

GetAlertRuleOk returns a tuple with the AlertRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRule

`func (o *Essentials) SetAlertRule(v string)`

SetAlertRule sets AlertRule field to given value.

### HasAlertRule

`func (o *Essentials) HasAlertRule() bool`

HasAlertRule returns a boolean if a field has been set.

### GetSeverity

`func (o *Essentials) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Essentials) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Essentials) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Essentials) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSignalType

`func (o *Essentials) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *Essentials) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *Essentials) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.

### HasSignalType

`func (o *Essentials) HasSignalType() bool`

HasSignalType returns a boolean if a field has been set.

### GetMonitorCondition

`func (o *Essentials) GetMonitorCondition() string`

GetMonitorCondition returns the MonitorCondition field if non-nil, zero value otherwise.

### GetMonitorConditionOk

`func (o *Essentials) GetMonitorConditionOk() (*string, bool)`

GetMonitorConditionOk returns a tuple with the MonitorCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorCondition

`func (o *Essentials) SetMonitorCondition(v string)`

SetMonitorCondition sets MonitorCondition field to given value.

### HasMonitorCondition

`func (o *Essentials) HasMonitorCondition() bool`

HasMonitorCondition returns a boolean if a field has been set.

### GetMonitoringService

`func (o *Essentials) GetMonitoringService() string`

GetMonitoringService returns the MonitoringService field if non-nil, zero value otherwise.

### GetMonitoringServiceOk

`func (o *Essentials) GetMonitoringServiceOk() (*string, bool)`

GetMonitoringServiceOk returns a tuple with the MonitoringService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringService

`func (o *Essentials) SetMonitoringService(v string)`

SetMonitoringService sets MonitoringService field to given value.

### HasMonitoringService

`func (o *Essentials) HasMonitoringService() bool`

HasMonitoringService returns a boolean if a field has been set.

### GetAlertTargetIDs

`func (o *Essentials) GetAlertTargetIDs() []string`

GetAlertTargetIDs returns the AlertTargetIDs field if non-nil, zero value otherwise.

### GetAlertTargetIDsOk

`func (o *Essentials) GetAlertTargetIDsOk() (*[]string, bool)`

GetAlertTargetIDsOk returns a tuple with the AlertTargetIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTargetIDs

`func (o *Essentials) SetAlertTargetIDs(v []string)`

SetAlertTargetIDs sets AlertTargetIDs field to given value.

### HasAlertTargetIDs

`func (o *Essentials) HasAlertTargetIDs() bool`

HasAlertTargetIDs returns a boolean if a field has been set.

### GetOriginAlertId

`func (o *Essentials) GetOriginAlertId() string`

GetOriginAlertId returns the OriginAlertId field if non-nil, zero value otherwise.

### GetOriginAlertIdOk

`func (o *Essentials) GetOriginAlertIdOk() (*string, bool)`

GetOriginAlertIdOk returns a tuple with the OriginAlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAlertId

`func (o *Essentials) SetOriginAlertId(v string)`

SetOriginAlertId sets OriginAlertId field to given value.

### HasOriginAlertId

`func (o *Essentials) HasOriginAlertId() bool`

HasOriginAlertId returns a boolean if a field has been set.

### GetFiredDateTime

`func (o *Essentials) GetFiredDateTime() time.Time`

GetFiredDateTime returns the FiredDateTime field if non-nil, zero value otherwise.

### GetFiredDateTimeOk

`func (o *Essentials) GetFiredDateTimeOk() (*time.Time, bool)`

GetFiredDateTimeOk returns a tuple with the FiredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiredDateTime

`func (o *Essentials) SetFiredDateTime(v time.Time)`

SetFiredDateTime sets FiredDateTime field to given value.

### HasFiredDateTime

`func (o *Essentials) HasFiredDateTime() bool`

HasFiredDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *Essentials) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Essentials) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Essentials) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Essentials) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEssentialsVersion

`func (o *Essentials) GetEssentialsVersion() string`

GetEssentialsVersion returns the EssentialsVersion field if non-nil, zero value otherwise.

### GetEssentialsVersionOk

`func (o *Essentials) GetEssentialsVersionOk() (*string, bool)`

GetEssentialsVersionOk returns a tuple with the EssentialsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssentialsVersion

`func (o *Essentials) SetEssentialsVersion(v string)`

SetEssentialsVersion sets EssentialsVersion field to given value.

### HasEssentialsVersion

`func (o *Essentials) HasEssentialsVersion() bool`

HasEssentialsVersion returns a boolean if a field has been set.

### GetAlertContextVersion

`func (o *Essentials) GetAlertContextVersion() string`

GetAlertContextVersion returns the AlertContextVersion field if non-nil, zero value otherwise.

### GetAlertContextVersionOk

`func (o *Essentials) GetAlertContextVersionOk() (*string, bool)`

GetAlertContextVersionOk returns a tuple with the AlertContextVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertContextVersion

`func (o *Essentials) SetAlertContextVersion(v string)`

SetAlertContextVersion sets AlertContextVersion field to given value.

### HasAlertContextVersion

`func (o *Essentials) HasAlertContextVersion() bool`

HasAlertContextVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


