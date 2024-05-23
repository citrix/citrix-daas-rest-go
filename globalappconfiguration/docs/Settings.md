# Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppSettings** | Pointer to [**AppSettings**](AppSettings.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UseForAppConfig** | Pointer to **bool** |  | [optional] 

## Methods

### NewSettings

`func NewSettings() *Settings`

NewSettings instantiates a new Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsWithDefaults

`func NewSettingsWithDefaults() *Settings`

NewSettingsWithDefaults instantiates a new Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppSettings

`func (o *Settings) GetAppSettings() AppSettings`

GetAppSettings returns the AppSettings field if non-nil, zero value otherwise.

### GetAppSettingsOk

`func (o *Settings) GetAppSettingsOk() (*AppSettings, bool)`

GetAppSettingsOk returns a tuple with the AppSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSettings

`func (o *Settings) SetAppSettings(v AppSettings)`

SetAppSettings sets AppSettings field to given value.

### HasAppSettings

`func (o *Settings) HasAppSettings() bool`

HasAppSettings returns a boolean if a field has been set.

### GetDescription

`func (o *Settings) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Settings) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Settings) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Settings) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *Settings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Settings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Settings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Settings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUseForAppConfig

`func (o *Settings) GetUseForAppConfig() bool`

GetUseForAppConfig returns the UseForAppConfig field if non-nil, zero value otherwise.

### GetUseForAppConfigOk

`func (o *Settings) GetUseForAppConfigOk() (*bool, bool)`

GetUseForAppConfigOk returns a tuple with the UseForAppConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForAppConfig

`func (o *Settings) SetUseForAppConfig(v bool)`

SetUseForAppConfig sets UseForAppConfig field to given value.

### HasUseForAppConfig

`func (o *Settings) HasUseForAppConfig() bool`

HasUseForAppConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


