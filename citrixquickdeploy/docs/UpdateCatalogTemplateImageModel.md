# UpdateCatalogTemplateImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | ID of the master image to update the catalog with | 
**CitrixPrepared** | Pointer to **bool** | Whether the image was prepared by Citrix, or provided by the customer | [optional] 
**VdaUpdateDelay** | Pointer to **int32** | Number of minutes to delay updating the VDAs.  -1 &#x3D; Natural reboot (on next shutdown)  0 &#x3D; Immediate reboot (restart all machines now)  &gt;0 &#x3D; Distributed reboot (gradual rollout with warnings over specified minutes) | [optional] 
**WarningTitle** | Pointer to **NullableString** | The window title used when showing the warning message in user sessions.  Optional. If not specified, defaults to &#39;Help desk message&#39;. | [optional] 
**WarningMessage** | Pointer to **NullableString** | Custom warning message displayed to users before reboot.  The optional pattern &#39;%m%&#39; is replaced by the number of minutes until the reboot.  If not specified, uses the default message. | [optional] 

## Methods

### NewUpdateCatalogTemplateImageModel

`func NewUpdateCatalogTemplateImageModel(templateId string, ) *UpdateCatalogTemplateImageModel`

NewUpdateCatalogTemplateImageModel instantiates a new UpdateCatalogTemplateImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCatalogTemplateImageModelWithDefaults

`func NewUpdateCatalogTemplateImageModelWithDefaults() *UpdateCatalogTemplateImageModel`

NewUpdateCatalogTemplateImageModelWithDefaults instantiates a new UpdateCatalogTemplateImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *UpdateCatalogTemplateImageModel) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UpdateCatalogTemplateImageModel) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UpdateCatalogTemplateImageModel) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetCitrixPrepared

`func (o *UpdateCatalogTemplateImageModel) GetCitrixPrepared() bool`

GetCitrixPrepared returns the CitrixPrepared field if non-nil, zero value otherwise.

### GetCitrixPreparedOk

`func (o *UpdateCatalogTemplateImageModel) GetCitrixPreparedOk() (*bool, bool)`

GetCitrixPreparedOk returns a tuple with the CitrixPrepared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixPrepared

`func (o *UpdateCatalogTemplateImageModel) SetCitrixPrepared(v bool)`

SetCitrixPrepared sets CitrixPrepared field to given value.

### HasCitrixPrepared

`func (o *UpdateCatalogTemplateImageModel) HasCitrixPrepared() bool`

HasCitrixPrepared returns a boolean if a field has been set.

### GetVdaUpdateDelay

`func (o *UpdateCatalogTemplateImageModel) GetVdaUpdateDelay() int32`

GetVdaUpdateDelay returns the VdaUpdateDelay field if non-nil, zero value otherwise.

### GetVdaUpdateDelayOk

`func (o *UpdateCatalogTemplateImageModel) GetVdaUpdateDelayOk() (*int32, bool)`

GetVdaUpdateDelayOk returns a tuple with the VdaUpdateDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaUpdateDelay

`func (o *UpdateCatalogTemplateImageModel) SetVdaUpdateDelay(v int32)`

SetVdaUpdateDelay sets VdaUpdateDelay field to given value.

### HasVdaUpdateDelay

`func (o *UpdateCatalogTemplateImageModel) HasVdaUpdateDelay() bool`

HasVdaUpdateDelay returns a boolean if a field has been set.

### GetWarningTitle

`func (o *UpdateCatalogTemplateImageModel) GetWarningTitle() string`

GetWarningTitle returns the WarningTitle field if non-nil, zero value otherwise.

### GetWarningTitleOk

`func (o *UpdateCatalogTemplateImageModel) GetWarningTitleOk() (*string, bool)`

GetWarningTitleOk returns a tuple with the WarningTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningTitle

`func (o *UpdateCatalogTemplateImageModel) SetWarningTitle(v string)`

SetWarningTitle sets WarningTitle field to given value.

### HasWarningTitle

`func (o *UpdateCatalogTemplateImageModel) HasWarningTitle() bool`

HasWarningTitle returns a boolean if a field has been set.

### SetWarningTitleNil

`func (o *UpdateCatalogTemplateImageModel) SetWarningTitleNil(b bool)`

 SetWarningTitleNil sets the value for WarningTitle to be an explicit nil

### UnsetWarningTitle
`func (o *UpdateCatalogTemplateImageModel) UnsetWarningTitle()`

UnsetWarningTitle ensures that no value is present for WarningTitle, not even an explicit nil
### GetWarningMessage

`func (o *UpdateCatalogTemplateImageModel) GetWarningMessage() string`

GetWarningMessage returns the WarningMessage field if non-nil, zero value otherwise.

### GetWarningMessageOk

`func (o *UpdateCatalogTemplateImageModel) GetWarningMessageOk() (*string, bool)`

GetWarningMessageOk returns a tuple with the WarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessage

`func (o *UpdateCatalogTemplateImageModel) SetWarningMessage(v string)`

SetWarningMessage sets WarningMessage field to given value.

### HasWarningMessage

`func (o *UpdateCatalogTemplateImageModel) HasWarningMessage() bool`

HasWarningMessage returns a boolean if a field has been set.

### SetWarningMessageNil

`func (o *UpdateCatalogTemplateImageModel) SetWarningMessageNil(b bool)`

 SetWarningMessageNil sets the value for WarningMessage to be an explicit nil

### UnsetWarningMessage
`func (o *UpdateCatalogTemplateImageModel) UnsetWarningMessage()`

UnsetWarningMessage ensures that no value is present for WarningMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


