# DetectOrphanedResourcesResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The resource id.  | [optional] 
**Type** | Pointer to **NullableString** | The type of resource.  | [optional] 
**ProvisioningSchemeId** | Pointer to **NullableString** | The provisioning scheme Id.  | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of orphaned resource | [optional] 

## Methods

### NewDetectOrphanedResourcesResultModel

`func NewDetectOrphanedResourcesResultModel() *DetectOrphanedResourcesResultModel`

NewDetectOrphanedResourcesResultModel instantiates a new DetectOrphanedResourcesResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectOrphanedResourcesResultModelWithDefaults

`func NewDetectOrphanedResourcesResultModelWithDefaults() *DetectOrphanedResourcesResultModel`

NewDetectOrphanedResourcesResultModelWithDefaults instantiates a new DetectOrphanedResourcesResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetectOrphanedResourcesResultModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetectOrphanedResourcesResultModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetectOrphanedResourcesResultModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DetectOrphanedResourcesResultModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DetectOrphanedResourcesResultModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DetectOrphanedResourcesResultModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *DetectOrphanedResourcesResultModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DetectOrphanedResourcesResultModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DetectOrphanedResourcesResultModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DetectOrphanedResourcesResultModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DetectOrphanedResourcesResultModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DetectOrphanedResourcesResultModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetProvisioningSchemeId

`func (o *DetectOrphanedResourcesResultModel) GetProvisioningSchemeId() string`

GetProvisioningSchemeId returns the ProvisioningSchemeId field if non-nil, zero value otherwise.

### GetProvisioningSchemeIdOk

`func (o *DetectOrphanedResourcesResultModel) GetProvisioningSchemeIdOk() (*string, bool)`

GetProvisioningSchemeIdOk returns a tuple with the ProvisioningSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeId

`func (o *DetectOrphanedResourcesResultModel) SetProvisioningSchemeId(v string)`

SetProvisioningSchemeId sets ProvisioningSchemeId field to given value.

### HasProvisioningSchemeId

`func (o *DetectOrphanedResourcesResultModel) HasProvisioningSchemeId() bool`

HasProvisioningSchemeId returns a boolean if a field has been set.

### SetProvisioningSchemeIdNil

`func (o *DetectOrphanedResourcesResultModel) SetProvisioningSchemeIdNil(b bool)`

 SetProvisioningSchemeIdNil sets the value for ProvisioningSchemeId to be an explicit nil

### UnsetProvisioningSchemeId
`func (o *DetectOrphanedResourcesResultModel) UnsetProvisioningSchemeId()`

UnsetProvisioningSchemeId ensures that no value is present for ProvisioningSchemeId, not even an explicit nil
### GetMetadata

`func (o *DetectOrphanedResourcesResultModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DetectOrphanedResourcesResultModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DetectOrphanedResourcesResultModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DetectOrphanedResourcesResultModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DetectOrphanedResourcesResultModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DetectOrphanedResourcesResultModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


