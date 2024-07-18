# AwsEdcAccountResourceRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionNames** | Pointer to **[]string** | The names of the regions | [optional] 
**Regions** | Pointer to [**[]AwsWorkspaceRegion**](AwsWorkspaceRegion.md) | The list of regions including their names, codes and supported az ids | [optional] 

## Methods

### NewAwsEdcAccountResourceRegion

`func NewAwsEdcAccountResourceRegion() *AwsEdcAccountResourceRegion`

NewAwsEdcAccountResourceRegion instantiates a new AwsEdcAccountResourceRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcAccountResourceRegionWithDefaults

`func NewAwsEdcAccountResourceRegionWithDefaults() *AwsEdcAccountResourceRegion`

NewAwsEdcAccountResourceRegionWithDefaults instantiates a new AwsEdcAccountResourceRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionNames

`func (o *AwsEdcAccountResourceRegion) GetRegionNames() []string`

GetRegionNames returns the RegionNames field if non-nil, zero value otherwise.

### GetRegionNamesOk

`func (o *AwsEdcAccountResourceRegion) GetRegionNamesOk() (*[]string, bool)`

GetRegionNamesOk returns a tuple with the RegionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNames

`func (o *AwsEdcAccountResourceRegion) SetRegionNames(v []string)`

SetRegionNames sets RegionNames field to given value.

### HasRegionNames

`func (o *AwsEdcAccountResourceRegion) HasRegionNames() bool`

HasRegionNames returns a boolean if a field has been set.

### SetRegionNamesNil

`func (o *AwsEdcAccountResourceRegion) SetRegionNamesNil(b bool)`

 SetRegionNamesNil sets the value for RegionNames to be an explicit nil

### UnsetRegionNames
`func (o *AwsEdcAccountResourceRegion) UnsetRegionNames()`

UnsetRegionNames ensures that no value is present for RegionNames, not even an explicit nil
### GetRegions

`func (o *AwsEdcAccountResourceRegion) GetRegions() []AwsWorkspaceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AwsEdcAccountResourceRegion) GetRegionsOk() (*[]AwsWorkspaceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AwsEdcAccountResourceRegion) SetRegions(v []AwsWorkspaceRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AwsEdcAccountResourceRegion) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *AwsEdcAccountResourceRegion) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *AwsEdcAccountResourceRegion) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


