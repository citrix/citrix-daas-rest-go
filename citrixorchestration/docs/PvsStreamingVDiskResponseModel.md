# PvsStreamingVDiskResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ADPasswordEnabled** | Pointer to **NullableBool** | Is AD password enabled. | [optional] 
**DiskLocatorId** | Pointer to **NullableString** | Id of the PVS vDisk. | [optional] 
**DiskLocatorName** | **string** | Name of the PVS vDisk. | 
**SiteId** | Pointer to **NullableString** | Id of the PVS site. | [optional] 
**SiteName** | Pointer to **NullableString** | Name of the PVS site. | [optional] 
**StoreId** | Pointer to **NullableString** | Id of the PVS store. | [optional] 
**StoreName** | Pointer to **NullableString** | Name of the PVS store. | [optional] 
**StreamingEnabled** | Pointer to **NullableBool** | Is streaming enabled. | [optional] 
**WriteCacheSize** | Pointer to **NullableInt32** | Write cache size of the vDisk. | [optional] 
**WriteCacheType** | Pointer to **NullableInt32** | Write cache type of the vDisk. | [optional] 

## Methods

### NewPvsStreamingVDiskResponseModel

`func NewPvsStreamingVDiskResponseModel(diskLocatorName string, ) *PvsStreamingVDiskResponseModel`

NewPvsStreamingVDiskResponseModel instantiates a new PvsStreamingVDiskResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvsStreamingVDiskResponseModelWithDefaults

`func NewPvsStreamingVDiskResponseModelWithDefaults() *PvsStreamingVDiskResponseModel`

NewPvsStreamingVDiskResponseModelWithDefaults instantiates a new PvsStreamingVDiskResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetADPasswordEnabled

`func (o *PvsStreamingVDiskResponseModel) GetADPasswordEnabled() bool`

GetADPasswordEnabled returns the ADPasswordEnabled field if non-nil, zero value otherwise.

### GetADPasswordEnabledOk

`func (o *PvsStreamingVDiskResponseModel) GetADPasswordEnabledOk() (*bool, bool)`

GetADPasswordEnabledOk returns a tuple with the ADPasswordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADPasswordEnabled

`func (o *PvsStreamingVDiskResponseModel) SetADPasswordEnabled(v bool)`

SetADPasswordEnabled sets ADPasswordEnabled field to given value.

### HasADPasswordEnabled

`func (o *PvsStreamingVDiskResponseModel) HasADPasswordEnabled() bool`

HasADPasswordEnabled returns a boolean if a field has been set.

### SetADPasswordEnabledNil

`func (o *PvsStreamingVDiskResponseModel) SetADPasswordEnabledNil(b bool)`

 SetADPasswordEnabledNil sets the value for ADPasswordEnabled to be an explicit nil

### UnsetADPasswordEnabled
`func (o *PvsStreamingVDiskResponseModel) UnsetADPasswordEnabled()`

UnsetADPasswordEnabled ensures that no value is present for ADPasswordEnabled, not even an explicit nil
### GetDiskLocatorId

`func (o *PvsStreamingVDiskResponseModel) GetDiskLocatorId() string`

GetDiskLocatorId returns the DiskLocatorId field if non-nil, zero value otherwise.

### GetDiskLocatorIdOk

`func (o *PvsStreamingVDiskResponseModel) GetDiskLocatorIdOk() (*string, bool)`

GetDiskLocatorIdOk returns a tuple with the DiskLocatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLocatorId

`func (o *PvsStreamingVDiskResponseModel) SetDiskLocatorId(v string)`

SetDiskLocatorId sets DiskLocatorId field to given value.

### HasDiskLocatorId

`func (o *PvsStreamingVDiskResponseModel) HasDiskLocatorId() bool`

HasDiskLocatorId returns a boolean if a field has been set.

### SetDiskLocatorIdNil

`func (o *PvsStreamingVDiskResponseModel) SetDiskLocatorIdNil(b bool)`

 SetDiskLocatorIdNil sets the value for DiskLocatorId to be an explicit nil

### UnsetDiskLocatorId
`func (o *PvsStreamingVDiskResponseModel) UnsetDiskLocatorId()`

UnsetDiskLocatorId ensures that no value is present for DiskLocatorId, not even an explicit nil
### GetDiskLocatorName

`func (o *PvsStreamingVDiskResponseModel) GetDiskLocatorName() string`

GetDiskLocatorName returns the DiskLocatorName field if non-nil, zero value otherwise.

### GetDiskLocatorNameOk

`func (o *PvsStreamingVDiskResponseModel) GetDiskLocatorNameOk() (*string, bool)`

GetDiskLocatorNameOk returns a tuple with the DiskLocatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLocatorName

`func (o *PvsStreamingVDiskResponseModel) SetDiskLocatorName(v string)`

SetDiskLocatorName sets DiskLocatorName field to given value.


### GetSiteId

`func (o *PvsStreamingVDiskResponseModel) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PvsStreamingVDiskResponseModel) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PvsStreamingVDiskResponseModel) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PvsStreamingVDiskResponseModel) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *PvsStreamingVDiskResponseModel) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *PvsStreamingVDiskResponseModel) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetSiteName

`func (o *PvsStreamingVDiskResponseModel) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *PvsStreamingVDiskResponseModel) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *PvsStreamingVDiskResponseModel) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *PvsStreamingVDiskResponseModel) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *PvsStreamingVDiskResponseModel) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *PvsStreamingVDiskResponseModel) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetStoreId

`func (o *PvsStreamingVDiskResponseModel) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *PvsStreamingVDiskResponseModel) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *PvsStreamingVDiskResponseModel) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *PvsStreamingVDiskResponseModel) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### SetStoreIdNil

`func (o *PvsStreamingVDiskResponseModel) SetStoreIdNil(b bool)`

 SetStoreIdNil sets the value for StoreId to be an explicit nil

### UnsetStoreId
`func (o *PvsStreamingVDiskResponseModel) UnsetStoreId()`

UnsetStoreId ensures that no value is present for StoreId, not even an explicit nil
### GetStoreName

`func (o *PvsStreamingVDiskResponseModel) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *PvsStreamingVDiskResponseModel) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *PvsStreamingVDiskResponseModel) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *PvsStreamingVDiskResponseModel) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### SetStoreNameNil

`func (o *PvsStreamingVDiskResponseModel) SetStoreNameNil(b bool)`

 SetStoreNameNil sets the value for StoreName to be an explicit nil

### UnsetStoreName
`func (o *PvsStreamingVDiskResponseModel) UnsetStoreName()`

UnsetStoreName ensures that no value is present for StoreName, not even an explicit nil
### GetStreamingEnabled

`func (o *PvsStreamingVDiskResponseModel) GetStreamingEnabled() bool`

GetStreamingEnabled returns the StreamingEnabled field if non-nil, zero value otherwise.

### GetStreamingEnabledOk

`func (o *PvsStreamingVDiskResponseModel) GetStreamingEnabledOk() (*bool, bool)`

GetStreamingEnabledOk returns a tuple with the StreamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingEnabled

`func (o *PvsStreamingVDiskResponseModel) SetStreamingEnabled(v bool)`

SetStreamingEnabled sets StreamingEnabled field to given value.

### HasStreamingEnabled

`func (o *PvsStreamingVDiskResponseModel) HasStreamingEnabled() bool`

HasStreamingEnabled returns a boolean if a field has been set.

### SetStreamingEnabledNil

`func (o *PvsStreamingVDiskResponseModel) SetStreamingEnabledNil(b bool)`

 SetStreamingEnabledNil sets the value for StreamingEnabled to be an explicit nil

### UnsetStreamingEnabled
`func (o *PvsStreamingVDiskResponseModel) UnsetStreamingEnabled()`

UnsetStreamingEnabled ensures that no value is present for StreamingEnabled, not even an explicit nil
### GetWriteCacheSize

`func (o *PvsStreamingVDiskResponseModel) GetWriteCacheSize() int32`

GetWriteCacheSize returns the WriteCacheSize field if non-nil, zero value otherwise.

### GetWriteCacheSizeOk

`func (o *PvsStreamingVDiskResponseModel) GetWriteCacheSizeOk() (*int32, bool)`

GetWriteCacheSizeOk returns a tuple with the WriteCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCacheSize

`func (o *PvsStreamingVDiskResponseModel) SetWriteCacheSize(v int32)`

SetWriteCacheSize sets WriteCacheSize field to given value.

### HasWriteCacheSize

`func (o *PvsStreamingVDiskResponseModel) HasWriteCacheSize() bool`

HasWriteCacheSize returns a boolean if a field has been set.

### SetWriteCacheSizeNil

`func (o *PvsStreamingVDiskResponseModel) SetWriteCacheSizeNil(b bool)`

 SetWriteCacheSizeNil sets the value for WriteCacheSize to be an explicit nil

### UnsetWriteCacheSize
`func (o *PvsStreamingVDiskResponseModel) UnsetWriteCacheSize()`

UnsetWriteCacheSize ensures that no value is present for WriteCacheSize, not even an explicit nil
### GetWriteCacheType

`func (o *PvsStreamingVDiskResponseModel) GetWriteCacheType() int32`

GetWriteCacheType returns the WriteCacheType field if non-nil, zero value otherwise.

### GetWriteCacheTypeOk

`func (o *PvsStreamingVDiskResponseModel) GetWriteCacheTypeOk() (*int32, bool)`

GetWriteCacheTypeOk returns a tuple with the WriteCacheType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCacheType

`func (o *PvsStreamingVDiskResponseModel) SetWriteCacheType(v int32)`

SetWriteCacheType sets WriteCacheType field to given value.

### HasWriteCacheType

`func (o *PvsStreamingVDiskResponseModel) HasWriteCacheType() bool`

HasWriteCacheType returns a boolean if a field has been set.

### SetWriteCacheTypeNil

`func (o *PvsStreamingVDiskResponseModel) SetWriteCacheTypeNil(b bool)`

 SetWriteCacheTypeNil sets the value for WriteCacheType to be an explicit nil

### UnsetWriteCacheType
`func (o *PvsStreamingVDiskResponseModel) UnsetWriteCacheType()`

UnsetWriteCacheType ensures that no value is present for WriteCacheType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


