# TemplateImageUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL of the image | [optional] 
**AzureSasUrlExpiryTime** | **time.Time** | Time for which the Url would be valid | 

## Methods

### NewTemplateImageUrl

`func NewTemplateImageUrl(azureSasUrlExpiryTime time.Time, ) *TemplateImageUrl`

NewTemplateImageUrl instantiates a new TemplateImageUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateImageUrlWithDefaults

`func NewTemplateImageUrlWithDefaults() *TemplateImageUrl`

NewTemplateImageUrlWithDefaults instantiates a new TemplateImageUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *TemplateImageUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TemplateImageUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TemplateImageUrl) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TemplateImageUrl) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAzureSasUrlExpiryTime

`func (o *TemplateImageUrl) GetAzureSasUrlExpiryTime() time.Time`

GetAzureSasUrlExpiryTime returns the AzureSasUrlExpiryTime field if non-nil, zero value otherwise.

### GetAzureSasUrlExpiryTimeOk

`func (o *TemplateImageUrl) GetAzureSasUrlExpiryTimeOk() (*time.Time, bool)`

GetAzureSasUrlExpiryTimeOk returns a tuple with the AzureSasUrlExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSasUrlExpiryTime

`func (o *TemplateImageUrl) SetAzureSasUrlExpiryTime(v time.Time)`

SetAzureSasUrlExpiryTime sets AzureSasUrlExpiryTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


