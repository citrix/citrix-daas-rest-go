# ResourceAccessPolicyResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Uid of the resouce access policy rule. | [optional] 
**Name** | Pointer to **NullableString** | Name of the resource access policy rule. | [optional] 
**Description** | Pointer to **NullableString** | Description of the resource access policy rule. | [optional] 
**Enabled** | Pointer to **bool** | Whether the resource access policy rule is enabled. If a policy rule is disabled it is not considered when evaluating resource access. | [optional] 
**ExcludedSmartAccessFilterEnabled** | Pointer to **bool** | Specifies whether the excluded SmartAccess filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**ExcludedSmartAccessTags** | Pointer to [**[]SmartAccessTagResponseModel**](SmartAccessTagResponseModel.md) | Specifies the SmartAccess tags which will be denied access to resources. | [optional] 
**IncludedSmartAccessFilterEnabled** | Pointer to **bool** | Specifies whether the included SmartAccess filter is enabled. If the filter is disabled, it is ignored when the access policy is evaluated. | [optional] 
**IncludedSmartAccessFilterType** | Pointer to [**FilterMatchType**](FilterMatchType.md) |  | [optional] 
**IncludedSmartAccessTags** | Pointer to [**[]SmartAccessTagResponseModel**](SmartAccessTagResponseModel.md) | Specifies the SmartAccess tags which will be allowed access to resources. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Additional metadata for the resource access policy rule. | [optional] 
**ResourceType** | Pointer to [**ResourceAccessPolicyResourceType**](ResourceAccessPolicyResourceType.md) |  | [optional] 
**TagName** | Pointer to **NullableString** | The name of the tag associated with this policy rule. | [optional] 
**TagUid** | Pointer to **NullableInt32** | The UID of the tag associated with this policy rule. | [optional] 

## Methods

### NewResourceAccessPolicyResponseModel

`func NewResourceAccessPolicyResponseModel() *ResourceAccessPolicyResponseModel`

NewResourceAccessPolicyResponseModel instantiates a new ResourceAccessPolicyResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAccessPolicyResponseModelWithDefaults

`func NewResourceAccessPolicyResponseModelWithDefaults() *ResourceAccessPolicyResponseModel`

NewResourceAccessPolicyResponseModelWithDefaults instantiates a new ResourceAccessPolicyResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceAccessPolicyResponseModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceAccessPolicyResponseModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceAccessPolicyResponseModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceAccessPolicyResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceAccessPolicyResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceAccessPolicyResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceAccessPolicyResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceAccessPolicyResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ResourceAccessPolicyResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ResourceAccessPolicyResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ResourceAccessPolicyResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceAccessPolicyResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceAccessPolicyResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceAccessPolicyResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResourceAccessPolicyResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResourceAccessPolicyResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *ResourceAccessPolicyResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ResourceAccessPolicyResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ResourceAccessPolicyResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ResourceAccessPolicyResponseModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExcludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyResponseModel) GetExcludedSmartAccessFilterEnabled() bool`

GetExcludedSmartAccessFilterEnabled returns the ExcludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedSmartAccessFilterEnabledOk

`func (o *ResourceAccessPolicyResponseModel) GetExcludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetExcludedSmartAccessFilterEnabledOk returns a tuple with the ExcludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyResponseModel) SetExcludedSmartAccessFilterEnabled(v bool)`

SetExcludedSmartAccessFilterEnabled sets ExcludedSmartAccessFilterEnabled field to given value.

### HasExcludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyResponseModel) HasExcludedSmartAccessFilterEnabled() bool`

HasExcludedSmartAccessFilterEnabled returns a boolean if a field has been set.

### GetExcludedSmartAccessTags

`func (o *ResourceAccessPolicyResponseModel) GetExcludedSmartAccessTags() []SmartAccessTagResponseModel`

GetExcludedSmartAccessTags returns the ExcludedSmartAccessTags field if non-nil, zero value otherwise.

### GetExcludedSmartAccessTagsOk

`func (o *ResourceAccessPolicyResponseModel) GetExcludedSmartAccessTagsOk() (*[]SmartAccessTagResponseModel, bool)`

GetExcludedSmartAccessTagsOk returns a tuple with the ExcludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSmartAccessTags

`func (o *ResourceAccessPolicyResponseModel) SetExcludedSmartAccessTags(v []SmartAccessTagResponseModel)`

SetExcludedSmartAccessTags sets ExcludedSmartAccessTags field to given value.

### HasExcludedSmartAccessTags

`func (o *ResourceAccessPolicyResponseModel) HasExcludedSmartAccessTags() bool`

HasExcludedSmartAccessTags returns a boolean if a field has been set.

### SetExcludedSmartAccessTagsNil

`func (o *ResourceAccessPolicyResponseModel) SetExcludedSmartAccessTagsNil(b bool)`

 SetExcludedSmartAccessTagsNil sets the value for ExcludedSmartAccessTags to be an explicit nil

### UnsetExcludedSmartAccessTags
`func (o *ResourceAccessPolicyResponseModel) UnsetExcludedSmartAccessTags()`

UnsetExcludedSmartAccessTags ensures that no value is present for ExcludedSmartAccessTags, not even an explicit nil
### GetIncludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyResponseModel) GetIncludedSmartAccessFilterEnabled() bool`

GetIncludedSmartAccessFilterEnabled returns the IncludedSmartAccessFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterEnabledOk

`func (o *ResourceAccessPolicyResponseModel) GetIncludedSmartAccessFilterEnabledOk() (*bool, bool)`

GetIncludedSmartAccessFilterEnabledOk returns a tuple with the IncludedSmartAccessFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyResponseModel) SetIncludedSmartAccessFilterEnabled(v bool)`

SetIncludedSmartAccessFilterEnabled sets IncludedSmartAccessFilterEnabled field to given value.

### HasIncludedSmartAccessFilterEnabled

`func (o *ResourceAccessPolicyResponseModel) HasIncludedSmartAccessFilterEnabled() bool`

HasIncludedSmartAccessFilterEnabled returns a boolean if a field has been set.

### GetIncludedSmartAccessFilterType

`func (o *ResourceAccessPolicyResponseModel) GetIncludedSmartAccessFilterType() FilterMatchType`

GetIncludedSmartAccessFilterType returns the IncludedSmartAccessFilterType field if non-nil, zero value otherwise.

### GetIncludedSmartAccessFilterTypeOk

`func (o *ResourceAccessPolicyResponseModel) GetIncludedSmartAccessFilterTypeOk() (*FilterMatchType, bool)`

GetIncludedSmartAccessFilterTypeOk returns a tuple with the IncludedSmartAccessFilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessFilterType

`func (o *ResourceAccessPolicyResponseModel) SetIncludedSmartAccessFilterType(v FilterMatchType)`

SetIncludedSmartAccessFilterType sets IncludedSmartAccessFilterType field to given value.

### HasIncludedSmartAccessFilterType

`func (o *ResourceAccessPolicyResponseModel) HasIncludedSmartAccessFilterType() bool`

HasIncludedSmartAccessFilterType returns a boolean if a field has been set.

### GetIncludedSmartAccessTags

`func (o *ResourceAccessPolicyResponseModel) GetIncludedSmartAccessTags() []SmartAccessTagResponseModel`

GetIncludedSmartAccessTags returns the IncludedSmartAccessTags field if non-nil, zero value otherwise.

### GetIncludedSmartAccessTagsOk

`func (o *ResourceAccessPolicyResponseModel) GetIncludedSmartAccessTagsOk() (*[]SmartAccessTagResponseModel, bool)`

GetIncludedSmartAccessTagsOk returns a tuple with the IncludedSmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSmartAccessTags

`func (o *ResourceAccessPolicyResponseModel) SetIncludedSmartAccessTags(v []SmartAccessTagResponseModel)`

SetIncludedSmartAccessTags sets IncludedSmartAccessTags field to given value.

### HasIncludedSmartAccessTags

`func (o *ResourceAccessPolicyResponseModel) HasIncludedSmartAccessTags() bool`

HasIncludedSmartAccessTags returns a boolean if a field has been set.

### SetIncludedSmartAccessTagsNil

`func (o *ResourceAccessPolicyResponseModel) SetIncludedSmartAccessTagsNil(b bool)`

 SetIncludedSmartAccessTagsNil sets the value for IncludedSmartAccessTags to be an explicit nil

### UnsetIncludedSmartAccessTags
`func (o *ResourceAccessPolicyResponseModel) UnsetIncludedSmartAccessTags()`

UnsetIncludedSmartAccessTags ensures that no value is present for IncludedSmartAccessTags, not even an explicit nil
### GetMetadata

`func (o *ResourceAccessPolicyResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceAccessPolicyResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceAccessPolicyResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceAccessPolicyResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ResourceAccessPolicyResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ResourceAccessPolicyResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetResourceType

`func (o *ResourceAccessPolicyResponseModel) GetResourceType() ResourceAccessPolicyResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceAccessPolicyResponseModel) GetResourceTypeOk() (*ResourceAccessPolicyResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceAccessPolicyResponseModel) SetResourceType(v ResourceAccessPolicyResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceAccessPolicyResponseModel) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTagName

`func (o *ResourceAccessPolicyResponseModel) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *ResourceAccessPolicyResponseModel) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *ResourceAccessPolicyResponseModel) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *ResourceAccessPolicyResponseModel) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### SetTagNameNil

`func (o *ResourceAccessPolicyResponseModel) SetTagNameNil(b bool)`

 SetTagNameNil sets the value for TagName to be an explicit nil

### UnsetTagName
`func (o *ResourceAccessPolicyResponseModel) UnsetTagName()`

UnsetTagName ensures that no value is present for TagName, not even an explicit nil
### GetTagUid

`func (o *ResourceAccessPolicyResponseModel) GetTagUid() int32`

GetTagUid returns the TagUid field if non-nil, zero value otherwise.

### GetTagUidOk

`func (o *ResourceAccessPolicyResponseModel) GetTagUidOk() (*int32, bool)`

GetTagUidOk returns a tuple with the TagUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagUid

`func (o *ResourceAccessPolicyResponseModel) SetTagUid(v int32)`

SetTagUid sets TagUid field to given value.

### HasTagUid

`func (o *ResourceAccessPolicyResponseModel) HasTagUid() bool`

HasTagUid returns a boolean if a field has been set.

### SetTagUidNil

`func (o *ResourceAccessPolicyResponseModel) SetTagUidNil(b bool)`

 SetTagUidNil sets the value for TagUid to be an explicit nil

### UnsetTagUid
`func (o *ResourceAccessPolicyResponseModel) UnsetTagUid()`

UnsetTagUid ensures that no value is present for TagUid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


