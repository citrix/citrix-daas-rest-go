# ZoneResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the zone. | 
**Name** | **string** | Name of the zone. | 
**Description** | Pointer to **string** | Description of the zone. | [optional] 
**IsPrimary** | **bool** | Indicates whether the zone is the primary zone for the site.  Resources in the site default to the primary zone if not set otherwise. | 
**IsHealthy** | Pointer to **bool** | The flag to indicate whether the zone is healthy. | [optional] 
**LastStateChangeTimeInUtc** | Pointer to **string** | The LastStateChangeTimeInUtc of the zone. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of Zone. | [optional] 
**ResourceLocation** | Pointer to [**ZoneResponseModelResourceLocation**](ZoneResponseModelResourceLocation.md) |  | [optional] 

## Methods

### NewZoneResponseModel

`func NewZoneResponseModel(id string, name string, isPrimary bool, ) *ZoneResponseModel`

NewZoneResponseModel instantiates a new ZoneResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneResponseModelWithDefaults

`func NewZoneResponseModelWithDefaults() *ZoneResponseModel`

NewZoneResponseModelWithDefaults instantiates a new ZoneResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ZoneResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ZoneResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZoneResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZoneResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ZoneResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPrimary

`func (o *ZoneResponseModel) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *ZoneResponseModel) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *ZoneResponseModel) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.


### GetIsHealthy

`func (o *ZoneResponseModel) GetIsHealthy() bool`

GetIsHealthy returns the IsHealthy field if non-nil, zero value otherwise.

### GetIsHealthyOk

`func (o *ZoneResponseModel) GetIsHealthyOk() (*bool, bool)`

GetIsHealthyOk returns a tuple with the IsHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHealthy

`func (o *ZoneResponseModel) SetIsHealthy(v bool)`

SetIsHealthy sets IsHealthy field to given value.

### HasIsHealthy

`func (o *ZoneResponseModel) HasIsHealthy() bool`

HasIsHealthy returns a boolean if a field has been set.

### GetLastStateChangeTimeInUtc

`func (o *ZoneResponseModel) GetLastStateChangeTimeInUtc() string`

GetLastStateChangeTimeInUtc returns the LastStateChangeTimeInUtc field if non-nil, zero value otherwise.

### GetLastStateChangeTimeInUtcOk

`func (o *ZoneResponseModel) GetLastStateChangeTimeInUtcOk() (*string, bool)`

GetLastStateChangeTimeInUtcOk returns a tuple with the LastStateChangeTimeInUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateChangeTimeInUtc

`func (o *ZoneResponseModel) SetLastStateChangeTimeInUtc(v string)`

SetLastStateChangeTimeInUtc sets LastStateChangeTimeInUtc field to given value.

### HasLastStateChangeTimeInUtc

`func (o *ZoneResponseModel) HasLastStateChangeTimeInUtc() bool`

HasLastStateChangeTimeInUtc returns a boolean if a field has been set.

### GetMetadata

`func (o *ZoneResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ZoneResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ZoneResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ZoneResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResourceLocation

`func (o *ZoneResponseModel) GetResourceLocation() ZoneResponseModelResourceLocation`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *ZoneResponseModel) GetResourceLocationOk() (*ZoneResponseModelResourceLocation, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocation

`func (o *ZoneResponseModel) SetResourceLocation(v ZoneResponseModelResourceLocation)`

SetResourceLocation sets ResourceLocation field to given value.

### HasResourceLocation

`func (o *ZoneResponseModel) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


