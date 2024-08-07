# GpoTestPolicyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **NullableString** | Error code. | [optional] 
**PolicyName** | Pointer to **NullableString** | Policy name. | [optional] 
**Settings** | Pointer to [**[]GpoTestSettingData**](GpoTestSettingData.md) | Data for settings in the policy. Only settings whose values have errors may be listed. | [optional] 
**Filters** | Pointer to [**[]GpoTestFilterData**](GpoTestFilterData.md) | Data for filters in the policy. Only filters whose values have errors may be listed. | [optional] 

## Methods

### NewGpoTestPolicyData

`func NewGpoTestPolicyData() *GpoTestPolicyData`

NewGpoTestPolicyData instantiates a new GpoTestPolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpoTestPolicyDataWithDefaults

`func NewGpoTestPolicyDataWithDefaults() *GpoTestPolicyData`

NewGpoTestPolicyDataWithDefaults instantiates a new GpoTestPolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *GpoTestPolicyData) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GpoTestPolicyData) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GpoTestPolicyData) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GpoTestPolicyData) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *GpoTestPolicyData) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *GpoTestPolicyData) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetPolicyName

`func (o *GpoTestPolicyData) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *GpoTestPolicyData) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *GpoTestPolicyData) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *GpoTestPolicyData) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *GpoTestPolicyData) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *GpoTestPolicyData) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetSettings

`func (o *GpoTestPolicyData) GetSettings() []GpoTestSettingData`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *GpoTestPolicyData) GetSettingsOk() (*[]GpoTestSettingData, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *GpoTestPolicyData) SetSettings(v []GpoTestSettingData)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *GpoTestPolicyData) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *GpoTestPolicyData) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *GpoTestPolicyData) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetFilters

`func (o *GpoTestPolicyData) GetFilters() []GpoTestFilterData`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GpoTestPolicyData) GetFiltersOk() (*[]GpoTestFilterData, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GpoTestPolicyData) SetFilters(v []GpoTestFilterData)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GpoTestPolicyData) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *GpoTestPolicyData) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *GpoTestPolicyData) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


