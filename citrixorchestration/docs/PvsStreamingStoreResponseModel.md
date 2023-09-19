# PvsStreamingStoreResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreId** | Pointer to **NullableString** | Id of the PVS store. | [optional] 
**StoreName** | **string** | Name of the PVS store. | 
**SiteId** | Pointer to **NullableString** | Id of the PVS site. | [optional] 

## Methods

### NewPvsStreamingStoreResponseModel

`func NewPvsStreamingStoreResponseModel(storeName string, ) *PvsStreamingStoreResponseModel`

NewPvsStreamingStoreResponseModel instantiates a new PvsStreamingStoreResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvsStreamingStoreResponseModelWithDefaults

`func NewPvsStreamingStoreResponseModelWithDefaults() *PvsStreamingStoreResponseModel`

NewPvsStreamingStoreResponseModelWithDefaults instantiates a new PvsStreamingStoreResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *PvsStreamingStoreResponseModel) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *PvsStreamingStoreResponseModel) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *PvsStreamingStoreResponseModel) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *PvsStreamingStoreResponseModel) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### SetStoreIdNil

`func (o *PvsStreamingStoreResponseModel) SetStoreIdNil(b bool)`

 SetStoreIdNil sets the value for StoreId to be an explicit nil

### UnsetStoreId
`func (o *PvsStreamingStoreResponseModel) UnsetStoreId()`

UnsetStoreId ensures that no value is present for StoreId, not even an explicit nil
### GetStoreName

`func (o *PvsStreamingStoreResponseModel) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *PvsStreamingStoreResponseModel) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *PvsStreamingStoreResponseModel) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.


### GetSiteId

`func (o *PvsStreamingStoreResponseModel) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PvsStreamingStoreResponseModel) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PvsStreamingStoreResponseModel) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PvsStreamingStoreResponseModel) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *PvsStreamingStoreResponseModel) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *PvsStreamingStoreResponseModel) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


