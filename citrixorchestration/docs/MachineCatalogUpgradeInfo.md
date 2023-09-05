# MachineCatalogUpgradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeType** | Pointer to [**VdaUpgradeType**](VdaUpgradeType.md) |  | [optional] 
**UpgradeState** | Pointer to [**VdaUpgradeState**](VdaUpgradeState.md) |  | [optional] 
**UpgradeScheduleStatus** | Pointer to [**VdaUpgradeScheduleStatus**](VdaUpgradeScheduleStatus.md) |  | [optional] 
**UpgradeOngoingMachinesCount** | Pointer to **int32** | Number of machines in the machine catalog with in-progress upgrades. | [optional] 
**UpgradeFailedMachinesCount** | Pointer to **int32** | Number of machines in the machine catalog with failed upgrades. | [optional] 

## Methods

### NewMachineCatalogUpgradeInfo

`func NewMachineCatalogUpgradeInfo() *MachineCatalogUpgradeInfo`

NewMachineCatalogUpgradeInfo instantiates a new MachineCatalogUpgradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogUpgradeInfoWithDefaults

`func NewMachineCatalogUpgradeInfoWithDefaults() *MachineCatalogUpgradeInfo`

NewMachineCatalogUpgradeInfoWithDefaults instantiates a new MachineCatalogUpgradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeType

`func (o *MachineCatalogUpgradeInfo) GetUpgradeType() VdaUpgradeType`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### GetUpgradeTypeOk

`func (o *MachineCatalogUpgradeInfo) GetUpgradeTypeOk() (*VdaUpgradeType, bool)`

GetUpgradeTypeOk returns a tuple with the UpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeType

`func (o *MachineCatalogUpgradeInfo) SetUpgradeType(v VdaUpgradeType)`

SetUpgradeType sets UpgradeType field to given value.

### HasUpgradeType

`func (o *MachineCatalogUpgradeInfo) HasUpgradeType() bool`

HasUpgradeType returns a boolean if a field has been set.

### GetUpgradeState

`func (o *MachineCatalogUpgradeInfo) GetUpgradeState() VdaUpgradeState`

GetUpgradeState returns the UpgradeState field if non-nil, zero value otherwise.

### GetUpgradeStateOk

`func (o *MachineCatalogUpgradeInfo) GetUpgradeStateOk() (*VdaUpgradeState, bool)`

GetUpgradeStateOk returns a tuple with the UpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeState

`func (o *MachineCatalogUpgradeInfo) SetUpgradeState(v VdaUpgradeState)`

SetUpgradeState sets UpgradeState field to given value.

### HasUpgradeState

`func (o *MachineCatalogUpgradeInfo) HasUpgradeState() bool`

HasUpgradeState returns a boolean if a field has been set.

### GetUpgradeScheduleStatus

`func (o *MachineCatalogUpgradeInfo) GetUpgradeScheduleStatus() VdaUpgradeScheduleStatus`

GetUpgradeScheduleStatus returns the UpgradeScheduleStatus field if non-nil, zero value otherwise.

### GetUpgradeScheduleStatusOk

`func (o *MachineCatalogUpgradeInfo) GetUpgradeScheduleStatusOk() (*VdaUpgradeScheduleStatus, bool)`

GetUpgradeScheduleStatusOk returns a tuple with the UpgradeScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeScheduleStatus

`func (o *MachineCatalogUpgradeInfo) SetUpgradeScheduleStatus(v VdaUpgradeScheduleStatus)`

SetUpgradeScheduleStatus sets UpgradeScheduleStatus field to given value.

### HasUpgradeScheduleStatus

`func (o *MachineCatalogUpgradeInfo) HasUpgradeScheduleStatus() bool`

HasUpgradeScheduleStatus returns a boolean if a field has been set.

### GetUpgradeOngoingMachinesCount

`func (o *MachineCatalogUpgradeInfo) GetUpgradeOngoingMachinesCount() int32`

GetUpgradeOngoingMachinesCount returns the UpgradeOngoingMachinesCount field if non-nil, zero value otherwise.

### GetUpgradeOngoingMachinesCountOk

`func (o *MachineCatalogUpgradeInfo) GetUpgradeOngoingMachinesCountOk() (*int32, bool)`

GetUpgradeOngoingMachinesCountOk returns a tuple with the UpgradeOngoingMachinesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOngoingMachinesCount

`func (o *MachineCatalogUpgradeInfo) SetUpgradeOngoingMachinesCount(v int32)`

SetUpgradeOngoingMachinesCount sets UpgradeOngoingMachinesCount field to given value.

### HasUpgradeOngoingMachinesCount

`func (o *MachineCatalogUpgradeInfo) HasUpgradeOngoingMachinesCount() bool`

HasUpgradeOngoingMachinesCount returns a boolean if a field has been set.

### GetUpgradeFailedMachinesCount

`func (o *MachineCatalogUpgradeInfo) GetUpgradeFailedMachinesCount() int32`

GetUpgradeFailedMachinesCount returns the UpgradeFailedMachinesCount field if non-nil, zero value otherwise.

### GetUpgradeFailedMachinesCountOk

`func (o *MachineCatalogUpgradeInfo) GetUpgradeFailedMachinesCountOk() (*int32, bool)`

GetUpgradeFailedMachinesCountOk returns a tuple with the UpgradeFailedMachinesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeFailedMachinesCount

`func (o *MachineCatalogUpgradeInfo) SetUpgradeFailedMachinesCount(v int32)`

SetUpgradeFailedMachinesCount sets UpgradeFailedMachinesCount field to given value.

### HasUpgradeFailedMachinesCount

`func (o *MachineCatalogUpgradeInfo) HasUpgradeFailedMachinesCount() bool`

HasUpgradeFailedMachinesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


