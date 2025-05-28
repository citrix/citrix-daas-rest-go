# CatalogTemplateImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | ID of the Template image to configure for the catalog | 
**CitrixPrepared** | Pointer to **bool** | Whether the image was prepared by Citrix, or provided by the customer | [optional] 

## Methods

### NewCatalogTemplateImageModel

`func NewCatalogTemplateImageModel(templateId string, ) *CatalogTemplateImageModel`

NewCatalogTemplateImageModel instantiates a new CatalogTemplateImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogTemplateImageModelWithDefaults

`func NewCatalogTemplateImageModelWithDefaults() *CatalogTemplateImageModel`

NewCatalogTemplateImageModelWithDefaults instantiates a new CatalogTemplateImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *CatalogTemplateImageModel) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CatalogTemplateImageModel) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CatalogTemplateImageModel) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetCitrixPrepared

`func (o *CatalogTemplateImageModel) GetCitrixPrepared() bool`

GetCitrixPrepared returns the CitrixPrepared field if non-nil, zero value otherwise.

### GetCitrixPreparedOk

`func (o *CatalogTemplateImageModel) GetCitrixPreparedOk() (*bool, bool)`

GetCitrixPreparedOk returns a tuple with the CitrixPrepared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixPrepared

`func (o *CatalogTemplateImageModel) SetCitrixPrepared(v bool)`

SetCitrixPrepared sets CitrixPrepared field to given value.

### HasCitrixPrepared

`func (o *CatalogTemplateImageModel) HasCitrixPrepared() bool`

HasCitrixPrepared returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


