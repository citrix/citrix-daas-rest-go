# ResourceAccessPolicyRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the resource access policy rule. | [optional] 
**Description** | Pointer to **NullableString** | Description of the resource access policy rule. | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the resource access policy rule is enabled. If a policy rule is disabled it is not considered when evaluating resource access. | [optional] [default to true]
**ExcludedSmartAccessFilterEnabled** | Pointer to **NullableBool** | Specifies whether the ExcludedSmartAccessTags filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**ExcludedSmartAccessTags** | Pointer to [**[]SmartAccessTagRequestModel**](SmartAccessTagRequestModel.md) | Specifies the SmartAccess tags which will be denied access to the resource. | [optional] 
**IncludedSmartAccessFilterEnabled** | Pointer to **NullableBool** | Specifies whether the IncludedSmartAccessTags filter is enabled.  If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**IncludedSmartAccessTags** | Pointer to [**[]SmartAccessTagRequestModel**](SmartAccessTagRequestModel.md) | Specifies the SmartAccess tags which grant access to the resource. | [optional] 
**IncludedSmartAccessFilterType** | Pointer to [**FilterMatchType**](FilterMatchType.md) |  | [optional] 
**ResourceType** | Pointer to [**ResourceAccessPolicyResourceType**](ResourceAccessPolicyResourceType.md) |  | [optional] 
**TagUid** | Pointer to **NullableInt32** | The UID of the tag associated with this policy rule. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Additional metadata for the resource access policy rule. | [optional] 

## Methods

### NewResourceAccessPolicyRequestModel

`func NewResourceAccessPolicyRequestModel() *ResourceAccessPolicyRequestModel`

NewResourceAccessPolicyRequestModel instantiates a new ResourceAccessPolicyRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAccessPolicyRequestModelWithDefaults

`func NewResourceAccessPolicyRequestModelWithDefaults() *ResourceAccessPolicyRequestModel`

NewResourceAccessPolicyRequestModelWithDefaults instantiates a new ResourceAccessPolicyRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceAccessPolicyRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceAccessPolicyRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceAccessPolicyRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceAccessPolicyRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ResourceAccessPolicyRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ResourceAccessPolicyRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ResourceAccessPolicyRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceAccessPolicyRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceAccessPolicyRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceAccessPolicyRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResourceAccessPolicyRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResourceAccessPolicyRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *ResourceAccessPolicyRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ResourceAccessPolicyRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ResourceAccessPolicyRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ResourceAccessPolicyRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *ResourceAccessPolicyRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ResourceAccessPolicyRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetExcludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyRequestModel) GetExcludedSmartAccessFilterEnabled() bool`

GetExcludedSmartAccessFilterEnabled returns the ExcludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedSmartAccessFilterEnabledOk

`func (o *ResourceAccessPolicyRequestModel) GetExcludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetExcludedSmartAccessFilterEnabledOk returns a tuple with the ExcludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyRequestModel) SetExcludedSmartAccessFilterEnabled(v bool)`

SetExcludedSmartAccessFilterEnabled sets ExcludedSmartAccessFilterEnabled field to given value.

### HasExcludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyRequestModel) HasExcludedSmartAccessFilterEnabled() bool`

HasExcludedSmartAccessFilterEnabled returns a boolean if a field has been set.

### SetExcludedSmartAccessFilterEnabledNil

`func (o *ResourceAccessPolicyRequestModel) SetExcludedSmartAccessFilterEnabledNil(b bool)`

 SetExcludedSmartAccessFilterEnabledNil sets the value for ExcludedSmartAccessFilterEnabled to be an explicit nil

### UnsetExcludedSmartAccessFilterEnabled
`func (o *ResourceAccessPolicyRequestModel) UnsetExcludedSmartAccessFilterEnabled()`

UnsetExcludedSmartAccessFilterEnabled ensures that no value is present for ExcludedSmartAccessFilterEnabled, not even an explicit nil
### GetExcludedSmartAccessTags

`func (o *ResourceAccessPolicyRequestModel) GetExcludedSmartAccessTags() []SmartAccessTagRequestModel`

GetExcludedSmartAccessTags returns the ExcludedSmartAccessTags field if non-nil, zero value otherwise.

### GetExcludedSmartAccessTagsOk

`func (o *ResourceAccessPolicyRequestModel) GetExcludedSmartAccessTagsOk() (*[]SmartAccessTagRequestModel, bool)`

GetExcludedSmartAccessTagsOk returns a tuple with the ExcludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSmartAccessTags

`func (o *ResourceAccessPolicyRequestModel) SetExcludedSmartAccessTags(v []SmartAccessTagRequestModel)`

SetExcludedSmartAccessTags sets ExcludedSmartAccessTags field to given value.

### HasExcludedSmartAccessTags

`func (o *ResourceAccessPolicyRequestModel) HasExcludedSmartAccessTags() bool`

HasExcludedSmartAccessTags returns a boolean if a field has been set.

### SetExcludedSmartAccessTagsNil

`func (o *ResourceAccessPolicyRequestModel) SetExcludedSmartAccessTagsNil(b bool)`

 SetExcludedSmartAccessTagsNil sets the value for ExcludedSmartAccessTags to be an explicit nil

### UnsetExcludedSmartAccessTags
`func (o *ResourceAccessPolicyRequestModel) UnsetExcludedSmartAccessTags()`

UnsetExcludedSmartAccessTags ensures that no value is present for ExcludedSmartAccessTags, not even an explicit nil
### GetIncludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyRequestModel) GetIncludedSmartAccessFilterEnabled() bool`

GetIncludedSmartAccessFilterEnabled returns the IncludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterEnabledOk

`func (o *ResourceAccessPolicyRequestModel) GetIncludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetIncludedSmartAccessFilterEnabledOk returns a tuple with the IncludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyRequestModel) SetIncludedSmartAccessFilterEnabled(v bool)`

SetIncludedSmartAccessFilterEnabled sets IncludedSmartAccessFilterEnabled field to given value.

### HasIncludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyRequestModel) HasIncludedSmartAccessFilterEnabled() bool`

HasIncludedSmartAccessFilterEnabled returns a boolean if a field has been set.

### SetIncludedSmartAccessFilterEnabledNil

`func (o *ResourceAccessPolicyRequestModel) SetIncludedSmartAccessFilterEnabledNil(b bool)`

 SetIncludedSmartAccessFilterEnabledNil sets the value for IncludedSmartAccessFilterEnabled to be an explicit nil

### UnsetIncludedSmartAccessFilterEnabled
`func (o *ResourceAccessPolicyRequestModel) UnsetIncludedSmartAccessFilterEnabled()`

UnsetIncludedSmartAccessFilterEnabled ensures that no value is present for IncludedSmartAccessFilterEnabled, not even an explicit nil
### GetIncludedSmartAccessTags

`func (o *ResourceAccessPolicyRequestModel) GetIncludedSmartAccessTags() []SmartAccessTagRequestModel`

GetIncludedSmartAccessTags returns the IncludedSmartAccessTags field if non-nil, zero value otherwise.

### GetIncludedSmartAccessTagsOk

`func (o *ResourceAccessPolicyRequestModel) GetIncludedSmartAccessTagsOk() (*[]SmartAccessTagRequestModel, bool)`

GetIncludedSmartAccessTagsOk returns a tuple with the IncludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessTags

`func (o *ResourceAccessPolicyRequestModel) SetIncludedSmartAccessTags(v []SmartAccessTagRequestModel)`

SetIncludedSmartAccessTags sets IncludedSmartAccessTags field to given value.

### HasIncludedSmartAccessTags

`func (o *ResourceAccessPolicyRequestModel) HasIncludedSmartAccessTags() bool`

HasIncludedSmartAccessTags returns a boolean if a field has been set.

### SetIncludedSmartAccessTagsNil

`func (o *ResourceAccessPolicyRequestModel) SetIncludedSmartAccessTagsNil(b bool)`

 SetIncludedSmartAccessTagsNil sets the value for IncludedSmartAccessTags to be an explicit nil

### UnsetIncludedSmartAccessTags
`func (o *ResourceAccessPolicyRequestModel) UnsetIncludedSmartAccessTags()`

UnsetIncludedSmartAccessTags ensures that no value is present for IncludedSmartAccessTags, not even an explicit nil
### GetIncludedSmartAccessFilterType

`func (o *ResourceAccessPolicyRequestModel) GetIncludedSmartAccessFilterType() FilterMatchType`

GetIncludedSmartAccessFilterType returns the IncludedSmartAccessFilterType field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterTypeOk

`func (o *ResourceAccessPolicyRequestModel) GetIncludedSmartAccessFilterTypeOk() (*FilterMatchType, bool)`

GetIncludedSmartAccessFilterTypeOk returns a tuple with the IncludedSmartAccessFilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterType

`func (o *ResourceAccessPolicyRequestModel) SetIncludedSmartAccessFilterType(v FilterMatchType)`

SetIncludedSmartAccessFilterType sets IncludedSmartAccessFilterType field to given value.

### HasIncludedSmartAccessFilterType

`func (o *ResourceAccessPolicyRequestModel) HasIncludedSmartAccessFilterType() bool`

HasIncludedSmartAccessFilterType returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourceAccessPolicyRequestModel) GetResourceType() ResourceAccessPolicyResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceAccessPolicyRequestModel) GetResourceTypeOk() (*ResourceAccessPolicyResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceAccessPolicyRequestModel) SetResourceType(v ResourceAccessPolicyResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceAccessPolicyRequestModel) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTagUid

`func (o *ResourceAccessPolicyRequestModel) GetTagUid() int32`

GetTagUid returns the TagUid field if non-nil, zero value otherwise.

### GetTagUidOk

`func (o *ResourceAccessPolicyRequestModel) GetTagUidOk() (*int32, bool)`

GetTagUidOk returns a tuple with the TagUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagUid

`func (o *ResourceAccessPolicyRequestModel) SetTagUid(v int32)`

SetTagUid sets TagUid field to given value.

### HasTagUid

`func (o *ResourceAccessPolicyRequestModel) HasTagUid() bool`

HasTagUid returns a boolean if a field has been set.

### SetTagUidNil

`func (o *ResourceAccessPolicyRequestModel) SetTagUidNil(b bool)`

 SetTagUidNil sets the value for TagUid to be an explicit nil

### UnsetTagUid
`func (o *ResourceAccessPolicyRequestModel) UnsetTagUid()`

UnsetTagUid ensures that no value is present for TagUid, not even an explicit nil
### GetMetadata

`func (o *ResourceAccessPolicyRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceAccessPolicyRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceAccessPolicyRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceAccessPolicyRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ResourceAccessPolicyRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ResourceAccessPolicyRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


