# UpdateCatalogTemplateImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | ID of the master image to update the catalog with | 
**CitrixPrepared** | Pointer to **bool** | Whether the image was prepared by Citrix, or provided by the customer | [optional] 
**VdaUpdateDelay** | Pointer to **int32** | Number of minutes to delay updating the VDAs | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


