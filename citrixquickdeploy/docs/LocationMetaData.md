# LocationMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionType** | Pointer to **string** | The type of the region. | [optional] [readonly] 
**RegionCategory** | Pointer to **string** | The category of the region. | [optional] [readonly] 
**GeographyGroup** | Pointer to **string** | The geography group of the location. | [optional] [readonly] 
**Longitude** | Pointer to **string** | The longitude of the location. | [optional] [readonly] 
**Latitude** | Pointer to **string** | The latitude of the location. | [optional] [readonly] 
**PhysicalLocation** | Pointer to **string** | The physical location of the Azure location. | [optional] [readonly] 
**PairedRegions** | Pointer to [**[]PairedRegion**](PairedRegion.md) | The regions paired to this region. | [optional] [readonly] 
**HomeLocation** | Pointer to **string** | The home location of an edge zone. | [optional] [readonly] 

## Methods

### NewLocationMetaData

`func NewLocationMetaData() *LocationMetaData`

NewLocationMetaData instantiates a new LocationMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationMetaDataWithDefaults

`func NewLocationMetaDataWithDefaults() *LocationMetaData`

NewLocationMetaDataWithDefaults instantiates a new LocationMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionType

`func (o *LocationMetaData) GetRegionType() string`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *LocationMetaData) GetRegionTypeOk() (*string, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *LocationMetaData) SetRegionType(v string)`

SetRegionType sets RegionType field to given value.

### HasRegionType

`func (o *LocationMetaData) HasRegionType() bool`

HasRegionType returns a boolean if a field has been set.

### GetRegionCategory

`func (o *LocationMetaData) GetRegionCategory() string`

GetRegionCategory returns the RegionCategory field if non-nil, zero value otherwise.

### GetRegionCategoryOk

`func (o *LocationMetaData) GetRegionCategoryOk() (*string, bool)`

GetRegionCategoryOk returns a tuple with the RegionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCategory

`func (o *LocationMetaData) SetRegionCategory(v string)`

SetRegionCategory sets RegionCategory field to given value.

### HasRegionCategory

`func (o *LocationMetaData) HasRegionCategory() bool`

HasRegionCategory returns a boolean if a field has been set.

### GetGeographyGroup

`func (o *LocationMetaData) GetGeographyGroup() string`

GetGeographyGroup returns the GeographyGroup field if non-nil, zero value otherwise.

### GetGeographyGroupOk

`func (o *LocationMetaData) GetGeographyGroupOk() (*string, bool)`

GetGeographyGroupOk returns a tuple with the GeographyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographyGroup

`func (o *LocationMetaData) SetGeographyGroup(v string)`

SetGeographyGroup sets GeographyGroup field to given value.

### HasGeographyGroup

`func (o *LocationMetaData) HasGeographyGroup() bool`

HasGeographyGroup returns a boolean if a field has been set.

### GetLongitude

`func (o *LocationMetaData) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocationMetaData) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocationMetaData) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *LocationMetaData) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *LocationMetaData) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocationMetaData) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocationMetaData) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *LocationMetaData) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetPhysicalLocation

`func (o *LocationMetaData) GetPhysicalLocation() string`

GetPhysicalLocation returns the PhysicalLocation field if non-nil, zero value otherwise.

### GetPhysicalLocationOk

`func (o *LocationMetaData) GetPhysicalLocationOk() (*string, bool)`

GetPhysicalLocationOk returns a tuple with the PhysicalLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalLocation

`func (o *LocationMetaData) SetPhysicalLocation(v string)`

SetPhysicalLocation sets PhysicalLocation field to given value.

### HasPhysicalLocation

`func (o *LocationMetaData) HasPhysicalLocation() bool`

HasPhysicalLocation returns a boolean if a field has been set.

### GetPairedRegions

`func (o *LocationMetaData) GetPairedRegions() []PairedRegion`

GetPairedRegions returns the PairedRegions field if non-nil, zero value otherwise.

### GetPairedRegionsOk

`func (o *LocationMetaData) GetPairedRegionsOk() (*[]PairedRegion, bool)`

GetPairedRegionsOk returns a tuple with the PairedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairedRegions

`func (o *LocationMetaData) SetPairedRegions(v []PairedRegion)`

SetPairedRegions sets PairedRegions field to given value.

### HasPairedRegions

`func (o *LocationMetaData) HasPairedRegions() bool`

HasPairedRegions returns a boolean if a field has been set.

### GetHomeLocation

`func (o *LocationMetaData) GetHomeLocation() string`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *LocationMetaData) GetHomeLocationOk() (*string, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *LocationMetaData) SetHomeLocation(v string)`

SetHomeLocation sets HomeLocation field to given value.

### HasHomeLocation

`func (o *LocationMetaData) HasHomeLocation() bool`

HasHomeLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


