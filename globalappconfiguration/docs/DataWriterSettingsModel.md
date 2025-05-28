# DataWriterSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CitrixCloudCustomerId** | Pointer to **string** |  | [optional] 
**SettingsChannel** | Pointer to [**SettingsChannel**](SettingsChannel.md) |  | [optional] 
**Url** | Pointer to [**Url**](Url.md) |  | [optional] 
**Workspace** | Pointer to [**TranslatedSettings**](TranslatedSettings.md) |  | [optional] 

## Methods

### NewDataWriterSettingsModel

`func NewDataWriterSettingsModel() *DataWriterSettingsModel`

NewDataWriterSettingsModel instantiates a new DataWriterSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWriterSettingsModelWithDefaults

`func NewDataWriterSettingsModelWithDefaults() *DataWriterSettingsModel`

NewDataWriterSettingsModelWithDefaults instantiates a new DataWriterSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitrixCloudCustomerId

`func (o *DataWriterSettingsModel) GetCitrixCloudCustomerId() string`

GetCitrixCloudCustomerId returns the CitrixCloudCustomerId field if non-nil, zero value otherwise.

### GetCitrixCloudCustomerIdOk

`func (o *DataWriterSettingsModel) GetCitrixCloudCustomerIdOk() (*string, bool)`

GetCitrixCloudCustomerIdOk returns a tuple with the CitrixCloudCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixCloudCustomerId

`func (o *DataWriterSettingsModel) SetCitrixCloudCustomerId(v string)`

SetCitrixCloudCustomerId sets CitrixCloudCustomerId field to given value.

### HasCitrixCloudCustomerId

`func (o *DataWriterSettingsModel) HasCitrixCloudCustomerId() bool`

HasCitrixCloudCustomerId returns a boolean if a field has been set.

### GetSettingsChannel

`func (o *DataWriterSettingsModel) GetSettingsChannel() SettingsChannel`

GetSettingsChannel returns the SettingsChannel field if non-nil, zero value otherwise.

### GetSettingsChannelOk

`func (o *DataWriterSettingsModel) GetSettingsChannelOk() (*SettingsChannel, bool)`

GetSettingsChannelOk returns a tuple with the SettingsChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsChannel

`func (o *DataWriterSettingsModel) SetSettingsChannel(v SettingsChannel)`

SetSettingsChannel sets SettingsChannel field to given value.

### HasSettingsChannel

`func (o *DataWriterSettingsModel) HasSettingsChannel() bool`

HasSettingsChannel returns a boolean if a field has been set.

### GetUrl

`func (o *DataWriterSettingsModel) GetUrl() Url`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DataWriterSettingsModel) GetUrlOk() (*Url, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DataWriterSettingsModel) SetUrl(v Url)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DataWriterSettingsModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWorkspace

`func (o *DataWriterSettingsModel) GetWorkspace() TranslatedSettings`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *DataWriterSettingsModel) GetWorkspaceOk() (*TranslatedSettings, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *DataWriterSettingsModel) SetWorkspace(v TranslatedSettings)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *DataWriterSettingsModel) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


