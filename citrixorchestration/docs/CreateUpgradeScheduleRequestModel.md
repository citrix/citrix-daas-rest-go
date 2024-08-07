# CreateUpgradeScheduleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTimeUtc** | Pointer to **NullableTime** | UTC time to start the Vda upgrade. Must be a future time. If set to null, the upgrade should be started at once. | [optional] 
**DurationInHours** | **int32** | Timeout duration in hours. Valid range is 1 to 24. | 
**VDAComponentsAndFeaturesRequestModel** | Pointer to [**VDAComponentsSelectionValidationRequestModel**](VDAComponentsSelectionValidationRequestModel.md) |  | [optional] 
**VdaWorkstationPackageUri** | Pointer to **NullableString** | Custom location to download the VDA Workstation package from. Currently, only network shares (specified using a UNC path) are supported. | [optional] 
**VdaServerPackageUri** | Pointer to **NullableString** | Custom location to download the VDA Server package from. Currently, only network shares (specified using a UNC path) are supported. | [optional] 
**ConcurrencyLevel** | Pointer to **NullableInt32** | Limits the Concurrent upgrades that can happen at any time in Upgrade Window. | [optional] 
**FailureThreshold** | Pointer to **NullableInt32** | Limits the number of failures that can take place during a scheduled upgrade. | [optional] 
**LogoffOption** | Pointer to [**LogoffOption**](LogoffOption.md) |  | [optional] 

## Methods

### NewCreateUpgradeScheduleRequestModel

`func NewCreateUpgradeScheduleRequestModel(durationInHours int32, ) *CreateUpgradeScheduleRequestModel`

NewCreateUpgradeScheduleRequestModel instantiates a new CreateUpgradeScheduleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpgradeScheduleRequestModelWithDefaults

`func NewCreateUpgradeScheduleRequestModelWithDefaults() *CreateUpgradeScheduleRequestModel`

NewCreateUpgradeScheduleRequestModelWithDefaults instantiates a new CreateUpgradeScheduleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTimeUtc

`func (o *CreateUpgradeScheduleRequestModel) GetStartDateTimeUtc() time.Time`

GetStartDateTimeUtc returns the StartDateTimeUtc field if non-nil, zero value otherwise.

### GetStartDateTimeUtcOk

`func (o *CreateUpgradeScheduleRequestModel) GetStartDateTimeUtcOk() (*time.Time, bool)`

GetStartDateTimeUtcOk returns a tuple with the StartDateTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTimeUtc

`func (o *CreateUpgradeScheduleRequestModel) SetStartDateTimeUtc(v time.Time)`

SetStartDateTimeUtc sets StartDateTimeUtc field to given value.

### HasStartDateTimeUtc

`func (o *CreateUpgradeScheduleRequestModel) HasStartDateTimeUtc() bool`

HasStartDateTimeUtc returns a boolean if a field has been set.

### SetStartDateTimeUtcNil

`func (o *CreateUpgradeScheduleRequestModel) SetStartDateTimeUtcNil(b bool)`

 SetStartDateTimeUtcNil sets the value for StartDateTimeUtc to be an explicit nil

### UnsetStartDateTimeUtc
`func (o *CreateUpgradeScheduleRequestModel) UnsetStartDateTimeUtc()`

UnsetStartDateTimeUtc ensures that no value is present for StartDateTimeUtc, not even an explicit nil
### GetDurationInHours

`func (o *CreateUpgradeScheduleRequestModel) GetDurationInHours() int32`

GetDurationInHours returns the DurationInHours field if non-nil, zero value otherwise.

### GetDurationInHoursOk

`func (o *CreateUpgradeScheduleRequestModel) GetDurationInHoursOk() (*int32, bool)`

GetDurationInHoursOk returns a tuple with the DurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInHours

`func (o *CreateUpgradeScheduleRequestModel) SetDurationInHours(v int32)`

SetDurationInHours sets DurationInHours field to given value.


### GetVDAComponentsAndFeaturesRequestModel

`func (o *CreateUpgradeScheduleRequestModel) GetVDAComponentsAndFeaturesRequestModel() VDAComponentsSelectionValidationRequestModel`

GetVDAComponentsAndFeaturesRequestModel returns the VDAComponentsAndFeaturesRequestModel field if non-nil, zero value otherwise.

### GetVDAComponentsAndFeaturesRequestModelOk

`func (o *CreateUpgradeScheduleRequestModel) GetVDAComponentsAndFeaturesRequestModelOk() (*VDAComponentsSelectionValidationRequestModel, bool)`

GetVDAComponentsAndFeaturesRequestModelOk returns a tuple with the VDAComponentsAndFeaturesRequestModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDAComponentsAndFeaturesRequestModel

`func (o *CreateUpgradeScheduleRequestModel) SetVDAComponentsAndFeaturesRequestModel(v VDAComponentsSelectionValidationRequestModel)`

SetVDAComponentsAndFeaturesRequestModel sets VDAComponentsAndFeaturesRequestModel field to given value.

### HasVDAComponentsAndFeaturesRequestModel

`func (o *CreateUpgradeScheduleRequestModel) HasVDAComponentsAndFeaturesRequestModel() bool`

HasVDAComponentsAndFeaturesRequestModel returns a boolean if a field has been set.

### GetVdaWorkstationPackageUri

`func (o *CreateUpgradeScheduleRequestModel) GetVdaWorkstationPackageUri() string`

GetVdaWorkstationPackageUri returns the VdaWorkstationPackageUri field if non-nil, zero value otherwise.

### GetVdaWorkstationPackageUriOk

`func (o *CreateUpgradeScheduleRequestModel) GetVdaWorkstationPackageUriOk() (*string, bool)`

GetVdaWorkstationPackageUriOk returns a tuple with the VdaWorkstationPackageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaWorkstationPackageUri

`func (o *CreateUpgradeScheduleRequestModel) SetVdaWorkstationPackageUri(v string)`

SetVdaWorkstationPackageUri sets VdaWorkstationPackageUri field to given value.

### HasVdaWorkstationPackageUri

`func (o *CreateUpgradeScheduleRequestModel) HasVdaWorkstationPackageUri() bool`

HasVdaWorkstationPackageUri returns a boolean if a field has been set.

### SetVdaWorkstationPackageUriNil

`func (o *CreateUpgradeScheduleRequestModel) SetVdaWorkstationPackageUriNil(b bool)`

 SetVdaWorkstationPackageUriNil sets the value for VdaWorkstationPackageUri to be an explicit nil

### UnsetVdaWorkstationPackageUri
`func (o *CreateUpgradeScheduleRequestModel) UnsetVdaWorkstationPackageUri()`

UnsetVdaWorkstationPackageUri ensures that no value is present for VdaWorkstationPackageUri, not even an explicit nil
### GetVdaServerPackageUri

`func (o *CreateUpgradeScheduleRequestModel) GetVdaServerPackageUri() string`

GetVdaServerPackageUri returns the VdaServerPackageUri field if non-nil, zero value otherwise.

### GetVdaServerPackageUriOk

`func (o *CreateUpgradeScheduleRequestModel) GetVdaServerPackageUriOk() (*string, bool)`

GetVdaServerPackageUriOk returns a tuple with the VdaServerPackageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaServerPackageUri

`func (o *CreateUpgradeScheduleRequestModel) SetVdaServerPackageUri(v string)`

SetVdaServerPackageUri sets VdaServerPackageUri field to given value.

### HasVdaServerPackageUri

`func (o *CreateUpgradeScheduleRequestModel) HasVdaServerPackageUri() bool`

HasVdaServerPackageUri returns a boolean if a field has been set.

### SetVdaServerPackageUriNil

`func (o *CreateUpgradeScheduleRequestModel) SetVdaServerPackageUriNil(b bool)`

 SetVdaServerPackageUriNil sets the value for VdaServerPackageUri to be an explicit nil

### UnsetVdaServerPackageUri
`func (o *CreateUpgradeScheduleRequestModel) UnsetVdaServerPackageUri()`

UnsetVdaServerPackageUri ensures that no value is present for VdaServerPackageUri, not even an explicit nil
### GetConcurrencyLevel

`func (o *CreateUpgradeScheduleRequestModel) GetConcurrencyLevel() int32`

GetConcurrencyLevel returns the ConcurrencyLevel field if non-nil, zero value otherwise.

### GetConcurrencyLevelOk

`func (o *CreateUpgradeScheduleRequestModel) GetConcurrencyLevelOk() (*int32, bool)`

GetConcurrencyLevelOk returns a tuple with the ConcurrencyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyLevel

`func (o *CreateUpgradeScheduleRequestModel) SetConcurrencyLevel(v int32)`

SetConcurrencyLevel sets ConcurrencyLevel field to given value.

### HasConcurrencyLevel

`func (o *CreateUpgradeScheduleRequestModel) HasConcurrencyLevel() bool`

HasConcurrencyLevel returns a boolean if a field has been set.

### SetConcurrencyLevelNil

`func (o *CreateUpgradeScheduleRequestModel) SetConcurrencyLevelNil(b bool)`

 SetConcurrencyLevelNil sets the value for ConcurrencyLevel to be an explicit nil

### UnsetConcurrencyLevel
`func (o *CreateUpgradeScheduleRequestModel) UnsetConcurrencyLevel()`

UnsetConcurrencyLevel ensures that no value is present for ConcurrencyLevel, not even an explicit nil
### GetFailureThreshold

`func (o *CreateUpgradeScheduleRequestModel) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *CreateUpgradeScheduleRequestModel) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *CreateUpgradeScheduleRequestModel) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *CreateUpgradeScheduleRequestModel) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.

### SetFailureThresholdNil

`func (o *CreateUpgradeScheduleRequestModel) SetFailureThresholdNil(b bool)`

 SetFailureThresholdNil sets the value for FailureThreshold to be an explicit nil

### UnsetFailureThreshold
`func (o *CreateUpgradeScheduleRequestModel) UnsetFailureThreshold()`

UnsetFailureThreshold ensures that no value is present for FailureThreshold, not even an explicit nil
### GetLogoffOption

`func (o *CreateUpgradeScheduleRequestModel) GetLogoffOption() LogoffOption`

GetLogoffOption returns the LogoffOption field if non-nil, zero value otherwise.

### GetLogoffOptionOk

`func (o *CreateUpgradeScheduleRequestModel) GetLogoffOptionOk() (*LogoffOption, bool)`

GetLogoffOptionOk returns a tuple with the LogoffOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffOption

`func (o *CreateUpgradeScheduleRequestModel) SetLogoffOption(v LogoffOption)`

SetLogoffOption sets LogoffOption field to given value.

### HasLogoffOption

`func (o *CreateUpgradeScheduleRequestModel) HasLogoffOption() bool`

HasLogoffOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


