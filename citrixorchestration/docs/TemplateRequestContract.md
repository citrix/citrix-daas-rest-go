# TemplateRequestContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateName** | Pointer to **string** | Template name as identifier | [optional] 
**Description** | Pointer to **string** | Template description | [optional] 
**Settings** | Pointer to [**[]SettingRequestContract**](SettingRequestContract.md) | Settings | [optional] 
**ForceWrite** | Pointer to **bool** | True to force this change to the database even if changes have been made by another user | [optional] 

## Methods

### NewTemplateRequestContract

`func NewTemplateRequestContract() *TemplateRequestContract`

NewTemplateRequestContract instantiates a new TemplateRequestContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateRequestContractWithDefaults

`func NewTemplateRequestContractWithDefaults() *TemplateRequestContract`

NewTemplateRequestContractWithDefaults instantiates a new TemplateRequestContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateName

`func (o *TemplateRequestContract) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *TemplateRequestContract) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *TemplateRequestContract) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *TemplateRequestContract) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateRequestContract) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateRequestContract) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateRequestContract) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateRequestContract) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSettings

`func (o *TemplateRequestContract) GetSettings() []SettingRequestContract`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *TemplateRequestContract) GetSettingsOk() (*[]SettingRequestContract, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *TemplateRequestContract) SetSettings(v []SettingRequestContract)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *TemplateRequestContract) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetForceWrite

`func (o *TemplateRequestContract) GetForceWrite() bool`

GetForceWrite returns the ForceWrite field if non-nil, zero value otherwise.

### GetForceWriteOk

`func (o *TemplateRequestContract) GetForceWriteOk() (*bool, bool)`

GetForceWriteOk returns a tuple with the ForceWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceWrite

`func (o *TemplateRequestContract) SetForceWrite(v bool)`

SetForceWrite sets ForceWrite field to given value.

### HasForceWrite

`func (o *TemplateRequestContract) HasForceWrite() bool`

HasForceWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


