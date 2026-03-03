# FlowLogConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowLogId** | Pointer to **NullableString** | Resource ID of the flow log | [optional] 
**VnetPeeringId** | Pointer to **NullableString** | VNet Peering ID | [optional] 
**CitrixVnetId** | Pointer to **NullableString** | Citrix-managed VNet resource ID | [optional] 
**StorageAccountId** | Pointer to **NullableString** | Storage account resource ID | [optional] 
**IsEnabled** | Pointer to **bool** | Whether flow logs are enabled | [optional] 
**RetentionDays** | Pointer to **int32** | Retention days | [optional] 
**Region** | Pointer to **NullableString** | Storage account region | [optional] 
**EnabledAt** | Pointer to **NullableTime** | When flow logs were enabled | [optional] 

## Methods

### NewFlowLogConfigurationModel

`func NewFlowLogConfigurationModel() *FlowLogConfigurationModel`

NewFlowLogConfigurationModel instantiates a new FlowLogConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowLogConfigurationModelWithDefaults

`func NewFlowLogConfigurationModelWithDefaults() *FlowLogConfigurationModel`

NewFlowLogConfigurationModelWithDefaults instantiates a new FlowLogConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowLogId

`func (o *FlowLogConfigurationModel) GetFlowLogId() string`

GetFlowLogId returns the FlowLogId field if non-nil, zero value otherwise.

### GetFlowLogIdOk

`func (o *FlowLogConfigurationModel) GetFlowLogIdOk() (*string, bool)`

GetFlowLogIdOk returns a tuple with the FlowLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowLogId

`func (o *FlowLogConfigurationModel) SetFlowLogId(v string)`

SetFlowLogId sets FlowLogId field to given value.

### HasFlowLogId

`func (o *FlowLogConfigurationModel) HasFlowLogId() bool`

HasFlowLogId returns a boolean if a field has been set.

### SetFlowLogIdNil

`func (o *FlowLogConfigurationModel) SetFlowLogIdNil(b bool)`

 SetFlowLogIdNil sets the value for FlowLogId to be an explicit nil

### UnsetFlowLogId
`func (o *FlowLogConfigurationModel) UnsetFlowLogId()`

UnsetFlowLogId ensures that no value is present for FlowLogId, not even an explicit nil
### GetVnetPeeringId

`func (o *FlowLogConfigurationModel) GetVnetPeeringId() string`

GetVnetPeeringId returns the VnetPeeringId field if non-nil, zero value otherwise.

### GetVnetPeeringIdOk

`func (o *FlowLogConfigurationModel) GetVnetPeeringIdOk() (*string, bool)`

GetVnetPeeringIdOk returns a tuple with the VnetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetPeeringId

`func (o *FlowLogConfigurationModel) SetVnetPeeringId(v string)`

SetVnetPeeringId sets VnetPeeringId field to given value.

### HasVnetPeeringId

`func (o *FlowLogConfigurationModel) HasVnetPeeringId() bool`

HasVnetPeeringId returns a boolean if a field has been set.

### SetVnetPeeringIdNil

`func (o *FlowLogConfigurationModel) SetVnetPeeringIdNil(b bool)`

 SetVnetPeeringIdNil sets the value for VnetPeeringId to be an explicit nil

### UnsetVnetPeeringId
`func (o *FlowLogConfigurationModel) UnsetVnetPeeringId()`

UnsetVnetPeeringId ensures that no value is present for VnetPeeringId, not even an explicit nil
### GetCitrixVnetId

`func (o *FlowLogConfigurationModel) GetCitrixVnetId() string`

GetCitrixVnetId returns the CitrixVnetId field if non-nil, zero value otherwise.

### GetCitrixVnetIdOk

`func (o *FlowLogConfigurationModel) GetCitrixVnetIdOk() (*string, bool)`

GetCitrixVnetIdOk returns a tuple with the CitrixVnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixVnetId

`func (o *FlowLogConfigurationModel) SetCitrixVnetId(v string)`

SetCitrixVnetId sets CitrixVnetId field to given value.

### HasCitrixVnetId

`func (o *FlowLogConfigurationModel) HasCitrixVnetId() bool`

HasCitrixVnetId returns a boolean if a field has been set.

### SetCitrixVnetIdNil

`func (o *FlowLogConfigurationModel) SetCitrixVnetIdNil(b bool)`

 SetCitrixVnetIdNil sets the value for CitrixVnetId to be an explicit nil

### UnsetCitrixVnetId
`func (o *FlowLogConfigurationModel) UnsetCitrixVnetId()`

UnsetCitrixVnetId ensures that no value is present for CitrixVnetId, not even an explicit nil
### GetStorageAccountId

`func (o *FlowLogConfigurationModel) GetStorageAccountId() string`

GetStorageAccountId returns the StorageAccountId field if non-nil, zero value otherwise.

### GetStorageAccountIdOk

`func (o *FlowLogConfigurationModel) GetStorageAccountIdOk() (*string, bool)`

GetStorageAccountIdOk returns a tuple with the StorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountId

`func (o *FlowLogConfigurationModel) SetStorageAccountId(v string)`

SetStorageAccountId sets StorageAccountId field to given value.

### HasStorageAccountId

`func (o *FlowLogConfigurationModel) HasStorageAccountId() bool`

HasStorageAccountId returns a boolean if a field has been set.

### SetStorageAccountIdNil

`func (o *FlowLogConfigurationModel) SetStorageAccountIdNil(b bool)`

 SetStorageAccountIdNil sets the value for StorageAccountId to be an explicit nil

### UnsetStorageAccountId
`func (o *FlowLogConfigurationModel) UnsetStorageAccountId()`

UnsetStorageAccountId ensures that no value is present for StorageAccountId, not even an explicit nil
### GetIsEnabled

`func (o *FlowLogConfigurationModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FlowLogConfigurationModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FlowLogConfigurationModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FlowLogConfigurationModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetRetentionDays

`func (o *FlowLogConfigurationModel) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *FlowLogConfigurationModel) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *FlowLogConfigurationModel) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *FlowLogConfigurationModel) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetRegion

`func (o *FlowLogConfigurationModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FlowLogConfigurationModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FlowLogConfigurationModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FlowLogConfigurationModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *FlowLogConfigurationModel) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *FlowLogConfigurationModel) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetEnabledAt

`func (o *FlowLogConfigurationModel) GetEnabledAt() time.Time`

GetEnabledAt returns the EnabledAt field if non-nil, zero value otherwise.

### GetEnabledAtOk

`func (o *FlowLogConfigurationModel) GetEnabledAtOk() (*time.Time, bool)`

GetEnabledAtOk returns a tuple with the EnabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAt

`func (o *FlowLogConfigurationModel) SetEnabledAt(v time.Time)`

SetEnabledAt sets EnabledAt field to given value.

### HasEnabledAt

`func (o *FlowLogConfigurationModel) HasEnabledAt() bool`

HasEnabledAt returns a boolean if a field has been set.

### SetEnabledAtNil

`func (o *FlowLogConfigurationModel) SetEnabledAtNil(b bool)`

 SetEnabledAtNil sets the value for EnabledAt to be an explicit nil

### UnsetEnabledAt
`func (o *FlowLogConfigurationModel) UnsetEnabledAt()`

UnsetEnabledAt ensures that no value is present for EnabledAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


