# MachineGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchFilterGroups** | Pointer to [**[]MachineAndSessionSearchFilterGroupRequestModel**](MachineAndSessionSearchFilterGroupRequestModel.md) | List of advanced search filter groups. | [optional] 
**SearchFilterGroupsType** | Pointer to [**MachineAndSessionSearchFilterGroupsType**](MachineAndSessionSearchFilterGroupsType.md) |  | [optional] 
**MachinePropertyName** | Pointer to **NullableString** | The  property name of the broker machine for grouping | [optional] 

## Methods

### NewMachineGroupRequestModel

`func NewMachineGroupRequestModel() *MachineGroupRequestModel`

NewMachineGroupRequestModel instantiates a new MachineGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineGroupRequestModelWithDefaults

`func NewMachineGroupRequestModelWithDefaults() *MachineGroupRequestModel`

NewMachineGroupRequestModelWithDefaults instantiates a new MachineGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchFilterGroups

`func (o *MachineGroupRequestModel) GetSearchFilterGroups() []MachineAndSessionSearchFilterGroupRequestModel`

GetSearchFilterGroups returns the SearchFilterGroups field if non-nil, zero value otherwise.

### GetSearchFilterGroupsOk

`func (o *MachineGroupRequestModel) GetSearchFilterGroupsOk() (*[]MachineAndSessionSearchFilterGroupRequestModel, bool)`

GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroups

`func (o *MachineGroupRequestModel) SetSearchFilterGroups(v []MachineAndSessionSearchFilterGroupRequestModel)`

SetSearchFilterGroups sets SearchFilterGroups field to given value.

### HasSearchFilterGroups

`func (o *MachineGroupRequestModel) HasSearchFilterGroups() bool`

HasSearchFilterGroups returns a boolean if a field has been set.

### SetSearchFilterGroupsNil

`func (o *MachineGroupRequestModel) SetSearchFilterGroupsNil(b bool)`

 SetSearchFilterGroupsNil sets the value for SearchFilterGroups to be an explicit nil

### UnsetSearchFilterGroups
`func (o *MachineGroupRequestModel) UnsetSearchFilterGroups()`

UnsetSearchFilterGroups ensures that no value is present for SearchFilterGroups, not even an explicit nil
### GetSearchFilterGroupsType

`func (o *MachineGroupRequestModel) GetSearchFilterGroupsType() MachineAndSessionSearchFilterGroupsType`

GetSearchFilterGroupsType returns the SearchFilterGroupsType field if non-nil, zero value otherwise.

### GetSearchFilterGroupsTypeOk

`func (o *MachineGroupRequestModel) GetSearchFilterGroupsTypeOk() (*MachineAndSessionSearchFilterGroupsType, bool)`

GetSearchFilterGroupsTypeOk returns a tuple with the SearchFilterGroupsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterGroupsType

`func (o *MachineGroupRequestModel) SetSearchFilterGroupsType(v MachineAndSessionSearchFilterGroupsType)`

SetSearchFilterGroupsType sets SearchFilterGroupsType field to given value.

### HasSearchFilterGroupsType

`func (o *MachineGroupRequestModel) HasSearchFilterGroupsType() bool`

HasSearchFilterGroupsType returns a boolean if a field has been set.

### GetMachinePropertyName

`func (o *MachineGroupRequestModel) GetMachinePropertyName() string`

GetMachinePropertyName returns the MachinePropertyName field if non-nil, zero value otherwise.

### GetMachinePropertyNameOk

`func (o *MachineGroupRequestModel) GetMachinePropertyNameOk() (*string, bool)`

GetMachinePropertyNameOk returns a tuple with the MachinePropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinePropertyName

`func (o *MachineGroupRequestModel) SetMachinePropertyName(v string)`

SetMachinePropertyName sets MachinePropertyName field to given value.

### HasMachinePropertyName

`func (o *MachineGroupRequestModel) HasMachinePropertyName() bool`

HasMachinePropertyName returns a boolean if a field has been set.

### SetMachinePropertyNameNil

`func (o *MachineGroupRequestModel) SetMachinePropertyNameNil(b bool)`

 SetMachinePropertyNameNil sets the value for MachinePropertyName to be an explicit nil

### UnsetMachinePropertyName
`func (o *MachineGroupRequestModel) UnsetMachinePropertyName()`

UnsetMachinePropertyName ensures that no value is present for MachinePropertyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


