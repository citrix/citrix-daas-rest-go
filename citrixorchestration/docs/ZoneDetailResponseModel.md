# ZoneDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the zone. | 
**Name** | **string** | Name of the zone. | 
**Description** | Pointer to **NullableString** | Description of the zone. | [optional] 
**IsPrimary** | **bool** | Indicates whether the zone is the primary zone for the site.  Resources in the site default to the primary zone if not set otherwise. | 
**IsHealthy** | Pointer to **bool** | The flag to indicate whether the zone is healthy. | [optional] 
**LastStateChangeTimeInUtc** | Pointer to **NullableString** | The LastStateChangeTimeInUtc of the zone. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of Zone. | [optional] 
**ResourceLocation** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**Controllers** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | List of controllers in the zone. | [optional] 

## Methods

### NewZoneDetailResponseModel

`func NewZoneDetailResponseModel(id string, name string, isPrimary bool, ) *ZoneDetailResponseModel`

NewZoneDetailResponseModel instantiates a new ZoneDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneDetailResponseModelWithDefaults

`func NewZoneDetailResponseModelWithDefaults() *ZoneDetailResponseModel`

NewZoneDetailResponseModelWithDefaults instantiates a new ZoneDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ZoneDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ZoneDetailResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZoneDetailResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZoneDetailResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ZoneDetailResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ZoneDetailResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ZoneDetailResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsPrimary

`func (o *ZoneDetailResponseModel) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *ZoneDetailResponseModel) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *ZoneDetailResponseModel) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.


### GetIsHealthy

`func (o *ZoneDetailResponseModel) GetIsHealthy() bool`

GetIsHealthy returns the IsHealthy field if non-nil, zero value otherwise.

### GetIsHealthyOk

`func (o *ZoneDetailResponseModel) GetIsHealthyOk() (*bool, bool)`

GetIsHealthyOk returns a tuple with the IsHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHealthy

`func (o *ZoneDetailResponseModel) SetIsHealthy(v bool)`

SetIsHealthy sets IsHealthy field to given value.

### HasIsHealthy

`func (o *ZoneDetailResponseModel) HasIsHealthy() bool`

HasIsHealthy returns a boolean if a field has been set.

### GetLastStateChangeTimeInUtc

`func (o *ZoneDetailResponseModel) GetLastStateChangeTimeInUtc() string`

GetLastStateChangeTimeInUtc returns the LastStateChangeTimeInUtc field if non-nil, zero value otherwise.

### GetLastStateChangeTimeInUtcOk

`func (o *ZoneDetailResponseModel) GetLastStateChangeTimeInUtcOk() (*string, bool)`

GetLastStateChangeTimeInUtcOk returns a tuple with the LastStateChangeTimeInUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateChangeTimeInUtc

`func (o *ZoneDetailResponseModel) SetLastStateChangeTimeInUtc(v string)`

SetLastStateChangeTimeInUtc sets LastStateChangeTimeInUtc field to given value.

### HasLastStateChangeTimeInUtc

`func (o *ZoneDetailResponseModel) HasLastStateChangeTimeInUtc() bool`

HasLastStateChangeTimeInUtc returns a boolean if a field has been set.

### SetLastStateChangeTimeInUtcNil

`func (o *ZoneDetailResponseModel) SetLastStateChangeTimeInUtcNil(b bool)`

 SetLastStateChangeTimeInUtcNil sets the value for LastStateChangeTimeInUtc to be an explicit nil

### UnsetLastStateChangeTimeInUtc
`func (o *ZoneDetailResponseModel) UnsetLastStateChangeTimeInUtc()`

UnsetLastStateChangeTimeInUtc ensures that no value is present for LastStateChangeTimeInUtc, not even an explicit nil
### GetMetadata

`func (o *ZoneDetailResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ZoneDetailResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ZoneDetailResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ZoneDetailResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ZoneDetailResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ZoneDetailResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetResourceLocation

`func (o *ZoneDetailResponseModel) GetResourceLocation() RefResponseModel`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *ZoneDetailResponseModel) GetResourceLocationOk() (*RefResponseModel, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocation

`func (o *ZoneDetailResponseModel) SetResourceLocation(v RefResponseModel)`

SetResourceLocation sets ResourceLocation field to given value.

### HasResourceLocation

`func (o *ZoneDetailResponseModel) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.

### GetControllers

`func (o *ZoneDetailResponseModel) GetControllers() []RefResponseModel`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *ZoneDetailResponseModel) GetControllersOk() (*[]RefResponseModel, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *ZoneDetailResponseModel) SetControllers(v []RefResponseModel)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *ZoneDetailResponseModel) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### SetControllersNil

`func (o *ZoneDetailResponseModel) SetControllersNil(b bool)`

 SetControllersNil sets the value for Controllers to be an explicit nil

### UnsetControllers
`func (o *ZoneDetailResponseModel) UnsetControllers()`

UnsetControllers ensures that no value is present for Controllers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


