# PlatformSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedTo** | Pointer to **[]string** |  | [optional] 
**AssignmentPriority** | Pointer to **int32** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**[]CategorySettings**](CategorySettings.md) |  | [optional] 
**UserOverride** | Pointer to **bool** |  | [optional] 

## Methods

### NewPlatformSettings

`func NewPlatformSettings() *PlatformSettings`

NewPlatformSettings instantiates a new PlatformSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformSettingsWithDefaults

`func NewPlatformSettingsWithDefaults() *PlatformSettings`

NewPlatformSettingsWithDefaults instantiates a new PlatformSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedTo

`func (o *PlatformSettings) GetAssignedTo() []string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *PlatformSettings) GetAssignedToOk() (*[]string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *PlatformSettings) SetAssignedTo(v []string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *PlatformSettings) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetAssignmentPriority

`func (o *PlatformSettings) GetAssignmentPriority() int32`

GetAssignmentPriority returns the AssignmentPriority field if non-nil, zero value otherwise.

### GetAssignmentPriorityOk

`func (o *PlatformSettings) GetAssignmentPriorityOk() (*int32, bool)`

GetAssignmentPriorityOk returns a tuple with the AssignmentPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentPriority

`func (o *PlatformSettings) SetAssignmentPriority(v int32)`

SetAssignmentPriority sets AssignmentPriority field to given value.

### HasAssignmentPriority

`func (o *PlatformSettings) HasAssignmentPriority() bool`

HasAssignmentPriority returns a boolean if a field has been set.

### GetCategory

`func (o *PlatformSettings) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PlatformSettings) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PlatformSettings) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PlatformSettings) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSettings

`func (o *PlatformSettings) GetSettings() []CategorySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PlatformSettings) GetSettingsOk() (*[]CategorySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PlatformSettings) SetSettings(v []CategorySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PlatformSettings) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetUserOverride

`func (o *PlatformSettings) GetUserOverride() bool`

GetUserOverride returns the UserOverride field if non-nil, zero value otherwise.

### GetUserOverrideOk

`func (o *PlatformSettings) GetUserOverrideOk() (*bool, bool)`

GetUserOverrideOk returns a tuple with the UserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOverride

`func (o *PlatformSettings) SetUserOverride(v bool)`

SetUserOverride sets UserOverride field to given value.

### HasUserOverride

`func (o *PlatformSettings) HasUserOverride() bool`

HasUserOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


