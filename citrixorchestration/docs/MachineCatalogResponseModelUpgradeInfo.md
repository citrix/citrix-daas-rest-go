# MachineCatalogResponseModelUpgradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeType** | Pointer to [**VdaUpgradeType**](VdaUpgradeType.md) |  | [optional] 
**UpgradeState** | Pointer to [**VdaUpgradeState**](VdaUpgradeState.md) |  | [optional] 
**UpgradeScheduleStatus** | Pointer to [**VdaUpgradeScheduleStatus**](VdaUpgradeScheduleStatus.md) |  | [optional] 
**UpgradeOngoingMachinesCount** | Pointer to **int32** | Number of machines in the machine catalog with in-progress upgrades. | [optional] 
**UpgradeFailedMachinesCount** | Pointer to **int32** | Number of machines in the machine catalog with failed upgrades. | [optional] 

## Methods

### NewMachineCatalogResponseModelUpgradeInfo

`func NewMachineCatalogResponseModelUpgradeInfo() *MachineCatalogResponseModelUpgradeInfo`

NewMachineCatalogResponseModelUpgradeInfo instantiates a new MachineCatalogResponseModelUpgradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogResponseModelUpgradeInfoWithDefaults

`func NewMachineCatalogResponseModelUpgradeInfoWithDefaults() *MachineCatalogResponseModelUpgradeInfo`

NewMachineCatalogResponseModelUpgradeInfoWithDefaults instantiates a new MachineCatalogResponseModelUpgradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeType

`func (o *MachineCatalogResponseModelUpgradeInfo) GetUpgradeType() VdaUpgradeType`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### GetUpgradeTypeOk

`func (o *MachineCatalogResponseModelUpgradeInfo) GetUpgradeTypeOk() (*VdaUpgradeType, bool)`

GetUpgradeTypeOk returns a tuple with the UpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeType

`func (o *MachineCatalogResponseModelUpgradeInfo) SetUpgradeType(v VdaUpgradeType)`

SetUpgradeType sets UpgradeType field to given value.

### HasUpgradeType

`func (o *MachineCatalogResponseModelUpgradeInfo) HasUpgradeType() bool`

HasUpgradeType returns a boolean if a field has been set.

### GetUpgradeState

`func (o *MachineCatalogResponseModelUpgradeInfo) GetUpgradeState() VdaUpgradeState`

GetUpgradeState returns the UpgradeState field if non-nil, zero value otherwise.

### GetUpgradeStateOk

`func (o *MachineCatalogResponseModelUpgradeInfo) GetUpgradeStateOk() (*VdaUpgradeState, bool)`

GetUpgradeStateOk returns a tuple with the UpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeState

`func (o *MachineCatalogResponseModelUpgradeInfo) SetUpgradeState(v VdaUpgradeState)`

SetUpgradeState sets UpgradeState field to given value.

### HasUpgradeState

`func (o *MachineCatalogResponseModelUpgradeInfo) HasUpgradeState() bool`

HasUpgradeState returns a boolean if a field has been set.

### GetUpgradeScheduleStatus

`func (o *MachineCatalogResponseModelUpgradeInfo) GetUpgradeScheduleStatus() VdaUpgradeScheduleStatus`

GetUpgradeScheduleStatus returns the UpgradeScheduleStatus field if non-nil, zero value otherwise.

### GetUpgradeScheduleStatusOk

`func (o *MachineCatalogResponseModelUpgradeInfo) GetUpgradeScheduleStatusOk() (*VdaUpgradeScheduleStatus, bool)`

GetUpgradeScheduleStatusOk returns a tuple with the UpgradeScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeScheduleStatus

`func (o *MachineCatalogResponseModelUpgradeInfo) SetUpgradeScheduleStatus(v VdaUpgradeScheduleStatus)`

SetUpgradeScheduleStatus sets UpgradeScheduleStatus field to given value.

### HasUpgradeScheduleStatus

`func (o *MachineCatalogResponseModelUpgradeInfo) HasUpgradeScheduleStatus() bool`

HasUpgradeScheduleStatus returns a boolean if a field has been set.

### GetUpgradeOngoingMachinesCount

`func (o *MachineCatalogResponseModelUpgradeInfo) GetUpgradeOngoingMachinesCount() int32`

GetUpgradeOngoingMachinesCount returns the UpgradeOngoingMachinesCount field if non-nil, zero value otherwise.

### GetUpgradeOngoingMachinesCountOk

`func (o *MachineCatalogResponseModelUpgradeInfo) GetUpgradeOngoingMachinesCountOk() (*int32, bool)`

GetUpgradeOngoingMachinesCountOk returns a tuple with the UpgradeOngoingMachinesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOngoingMachinesCount

`func (o *MachineCatalogResponseModelUpgradeInfo) SetUpgradeOngoingMachinesCount(v int32)`

SetUpgradeOngoingMachinesCount sets UpgradeOngoingMachinesCount field to given value.

### HasUpgradeOngoingMachinesCount

`func (o *MachineCatalogResponseModelUpgradeInfo) HasUpgradeOngoingMachinesCount() bool`

HasUpgradeOngoingMachinesCount returns a boolean if a field has been set.

### GetUpgradeFailedMachinesCount

`func (o *MachineCatalogResponseModelUpgradeInfo) GetUpgradeFailedMachinesCount() int32`

GetUpgradeFailedMachinesCount returns the UpgradeFailedMachinesCount field if non-nil, zero value otherwise.

### GetUpgradeFailedMachinesCountOk

`func (o *MachineCatalogResponseModelUpgradeInfo) GetUpgradeFailedMachinesCountOk() (*int32, bool)`

GetUpgradeFailedMachinesCountOk returns a tuple with the UpgradeFailedMachinesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeFailedMachinesCount

`func (o *MachineCatalogResponseModelUpgradeInfo) SetUpgradeFailedMachinesCount(v int32)`

SetUpgradeFailedMachinesCount sets UpgradeFailedMachinesCount field to given value.

### HasUpgradeFailedMachinesCount

`func (o *MachineCatalogResponseModelUpgradeInfo) HasUpgradeFailedMachinesCount() bool`

HasUpgradeFailedMachinesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


