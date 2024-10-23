# MachineModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Identity of machine-level AD object | [optional] [readonly] 
**Sid** | **string** | SID of machine-level AD object | 
**Name** | **string** | Distingushed name of machine-level AD object | 
**Type** | **string** | Type of machine-level AD object | 
**Description** | Pointer to **string** | Description of machine-level AD object | [optional] 
**SiteId** | **int64** | Identity of site to which this machine-level AD object belongs | 
**Enabled** | Pointer to **bool** | If this nachine-level AD object is enabled | [optional] 
**Priority** | Pointer to **int64** | Priority of machine-level AD object | [optional] 

## Methods

### NewMachineModel

`func NewMachineModel(sid string, name string, type_ string, siteId int64, ) *MachineModel`

NewMachineModel instantiates a new MachineModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineModelWithDefaults

`func NewMachineModelWithDefaults() *MachineModel`

NewMachineModelWithDefaults instantiates a new MachineModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MachineModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MachineModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSid

`func (o *MachineModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MachineModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MachineModel) SetSid(v string)`

SetSid sets Sid field to given value.


### GetName

`func (o *MachineModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MachineModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MachineModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MachineModel) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *MachineModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MachineModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MachineModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MachineModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSiteId

`func (o *MachineModel) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MachineModel) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MachineModel) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.


### GetEnabled

`func (o *MachineModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MachineModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MachineModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MachineModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *MachineModel) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MachineModel) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MachineModel) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MachineModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


