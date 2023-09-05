# StoreFrontServerResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the StoreFront server. | 
**Name** | **string** | Name of the StoreFront server. | 
**Description** | Pointer to **string** | Description of the StoreFront server. | [optional] 
**Url** | **string** | Url of the StoreFront server. | 
**Enabled** | **bool** | Indicates whether the StoreFront server is enabled. Disabled StoreFront servers will not have their URLs added to hosted receiver. | 
**DesktopGroupRefCount** | Pointer to **int32** | DesktopGroupRefCount | [optional] 

## Methods

### NewStoreFrontServerResponseModel

`func NewStoreFrontServerResponseModel(id string, name string, url string, enabled bool, ) *StoreFrontServerResponseModel`

NewStoreFrontServerResponseModel instantiates a new StoreFrontServerResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreFrontServerResponseModelWithDefaults

`func NewStoreFrontServerResponseModelWithDefaults() *StoreFrontServerResponseModel`

NewStoreFrontServerResponseModelWithDefaults instantiates a new StoreFrontServerResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StoreFrontServerResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoreFrontServerResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoreFrontServerResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StoreFrontServerResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreFrontServerResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreFrontServerResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StoreFrontServerResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoreFrontServerResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoreFrontServerResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoreFrontServerResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *StoreFrontServerResponseModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StoreFrontServerResponseModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StoreFrontServerResponseModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *StoreFrontServerResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StoreFrontServerResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StoreFrontServerResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDesktopGroupRefCount

`func (o *StoreFrontServerResponseModel) GetDesktopGroupRefCount() int32`

GetDesktopGroupRefCount returns the DesktopGroupRefCount field if non-nil, zero value otherwise.

### GetDesktopGroupRefCountOk

`func (o *StoreFrontServerResponseModel) GetDesktopGroupRefCountOk() (*int32, bool)`

GetDesktopGroupRefCountOk returns a tuple with the DesktopGroupRefCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopGroupRefCount

`func (o *StoreFrontServerResponseModel) SetDesktopGroupRefCount(v int32)`

SetDesktopGroupRefCount sets DesktopGroupRefCount field to given value.

### HasDesktopGroupRefCount

`func (o *StoreFrontServerResponseModel) HasDesktopGroupRefCount() bool`

HasDesktopGroupRefCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


