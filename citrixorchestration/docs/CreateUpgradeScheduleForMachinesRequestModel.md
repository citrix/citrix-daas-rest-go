# CreateUpgradeScheduleForMachinesRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineNameOrIds** | **[]string** | Machine list to create VDA upgrade schedules for. Item can be machine name or Id. | 
**StartDateTimeUtc** | Pointer to **NullableTime** | UTC time to start the Vda upgrade. Must be a future time. If set to null, the upgrade should be started at once. | [optional] 
**DurationInHours** | **int32** | Timeout duration in hours. Valid range is 1 to 24. | 
**VdaWorkstationPackageUri** | Pointer to **NullableString** | Custom location to download the VDA Workstation package from. Currently, only network shares (specified using a UNC path) are supported. | [optional] 
**VdaServerPackageUri** | Pointer to **NullableString** | Custom location to download the VDA Server package from. Currently, only network shares (specified using a UNC path) are supported. | [optional] 
**LogoffOption** | Pointer to [**LogoffOption**](LogoffOption.md) |  | [optional] 

## Methods

### NewCreateUpgradeScheduleForMachinesRequestModel

`func NewCreateUpgradeScheduleForMachinesRequestModel(machineNameOrIds []string, durationInHours int32, ) *CreateUpgradeScheduleForMachinesRequestModel`

NewCreateUpgradeScheduleForMachinesRequestModel instantiates a new CreateUpgradeScheduleForMachinesRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpgradeScheduleForMachinesRequestModelWithDefaults

`func NewCreateUpgradeScheduleForMachinesRequestModelWithDefaults() *CreateUpgradeScheduleForMachinesRequestModel`

NewCreateUpgradeScheduleForMachinesRequestModelWithDefaults instantiates a new CreateUpgradeScheduleForMachinesRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineNameOrIds

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetMachineNameOrIds() []string`

GetMachineNameOrIds returns the MachineNameOrIds field if non-nil, zero value otherwise.

### GetMachineNameOrIdsOk

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetMachineNameOrIdsOk() (*[]string, bool)`

GetMachineNameOrIdsOk returns a tuple with the MachineNameOrIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineNameOrIds

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetMachineNameOrIds(v []string)`

SetMachineNameOrIds sets MachineNameOrIds field to given value.


### GetStartDateTimeUtc

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetStartDateTimeUtc() time.Time`

GetStartDateTimeUtc returns the StartDateTimeUtc field if non-nil, zero value otherwise.

### GetStartDateTimeUtcOk

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetStartDateTimeUtcOk() (*time.Time, bool)`

GetStartDateTimeUtcOk returns a tuple with the StartDateTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTimeUtc

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetStartDateTimeUtc(v time.Time)`

SetStartDateTimeUtc sets StartDateTimeUtc field to given value.

### HasStartDateTimeUtc

`func (o *CreateUpgradeScheduleForMachinesRequestModel) HasStartDateTimeUtc() bool`

HasStartDateTimeUtc returns a boolean if a field has been set.

### SetStartDateTimeUtcNil

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetStartDateTimeUtcNil(b bool)`

 SetStartDateTimeUtcNil sets the value for StartDateTimeUtc to be an explicit nil

### UnsetStartDateTimeUtc
`func (o *CreateUpgradeScheduleForMachinesRequestModel) UnsetStartDateTimeUtc()`

UnsetStartDateTimeUtc ensures that no value is present for StartDateTimeUtc, not even an explicit nil
### GetDurationInHours

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetDurationInHours() int32`

GetDurationInHours returns the DurationInHours field if non-nil, zero value otherwise.

### GetDurationInHoursOk

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetDurationInHoursOk() (*int32, bool)`

GetDurationInHoursOk returns a tuple with the DurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInHours

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetDurationInHours(v int32)`

SetDurationInHours sets DurationInHours field to given value.


### GetVdaWorkstationPackageUri

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetVdaWorkstationPackageUri() string`

GetVdaWorkstationPackageUri returns the VdaWorkstationPackageUri field if non-nil, zero value otherwise.

### GetVdaWorkstationPackageUriOk

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetVdaWorkstationPackageUriOk() (*string, bool)`

GetVdaWorkstationPackageUriOk returns a tuple with the VdaWorkstationPackageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaWorkstationPackageUri

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetVdaWorkstationPackageUri(v string)`

SetVdaWorkstationPackageUri sets VdaWorkstationPackageUri field to given value.

### HasVdaWorkstationPackageUri

`func (o *CreateUpgradeScheduleForMachinesRequestModel) HasVdaWorkstationPackageUri() bool`

HasVdaWorkstationPackageUri returns a boolean if a field has been set.

### SetVdaWorkstationPackageUriNil

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetVdaWorkstationPackageUriNil(b bool)`

 SetVdaWorkstationPackageUriNil sets the value for VdaWorkstationPackageUri to be an explicit nil

### UnsetVdaWorkstationPackageUri
`func (o *CreateUpgradeScheduleForMachinesRequestModel) UnsetVdaWorkstationPackageUri()`

UnsetVdaWorkstationPackageUri ensures that no value is present for VdaWorkstationPackageUri, not even an explicit nil
### GetVdaServerPackageUri

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetVdaServerPackageUri() string`

GetVdaServerPackageUri returns the VdaServerPackageUri field if non-nil, zero value otherwise.

### GetVdaServerPackageUriOk

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetVdaServerPackageUriOk() (*string, bool)`

GetVdaServerPackageUriOk returns a tuple with the VdaServerPackageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaServerPackageUri

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetVdaServerPackageUri(v string)`

SetVdaServerPackageUri sets VdaServerPackageUri field to given value.

### HasVdaServerPackageUri

`func (o *CreateUpgradeScheduleForMachinesRequestModel) HasVdaServerPackageUri() bool`

HasVdaServerPackageUri returns a boolean if a field has been set.

### SetVdaServerPackageUriNil

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetVdaServerPackageUriNil(b bool)`

 SetVdaServerPackageUriNil sets the value for VdaServerPackageUri to be an explicit nil

### UnsetVdaServerPackageUri
`func (o *CreateUpgradeScheduleForMachinesRequestModel) UnsetVdaServerPackageUri()`

UnsetVdaServerPackageUri ensures that no value is present for VdaServerPackageUri, not even an explicit nil
### GetLogoffOption

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetLogoffOption() LogoffOption`

GetLogoffOption returns the LogoffOption field if non-nil, zero value otherwise.

### GetLogoffOptionOk

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetLogoffOptionOk() (*LogoffOption, bool)`

GetLogoffOptionOk returns a tuple with the LogoffOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffOption

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetLogoffOption(v LogoffOption)`

SetLogoffOption sets LogoffOption field to given value.

### HasLogoffOption

`func (o *CreateUpgradeScheduleForMachinesRequestModel) HasLogoffOption() bool`

HasLogoffOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


