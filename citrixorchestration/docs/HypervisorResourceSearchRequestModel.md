# HypervisorResourceSearchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The path to retrieve contents from. | [optional] 
**ResourceType** | Pointer to **string** | The type of resources to get. | [optional] 
**ForwardDirection** | Pointer to **bool** | True if returning items which begin at the ContinuationToken, false if ending at the ContinuationToken. | [optional] 
**NameContains** | Pointer to **string** | Only return items which contain this string in their Name. | [optional] 
**Tags** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | A list of key-value pairs, searching by a tag lets you filter the resources by a tag key or a tag value(value is optional). For example: [{\&quot;Name\&quot;: \&quot;Priority\&quot;, \&quot;Value\&quot;: \&quot;high\&quot;}, {\&quot;Name\&quot;: \&quot;Name\&quot;, \&quot;Value\&quot;: \&quot;test\&quot;}, {\&quot;Name\&quot;: \&quot;TestMachine\&quot;, \&quot;Value\&quot;: \&quot;\&quot;}, {\&quot;Name\&quot;: \&quot;ADBackup\&quot;, \&quot;Value\&quot;: \&quot;Not tagged\&quot;}] | [optional] 
**Id** | Pointer to **string** | Only return the item with the matching Id. | [optional] 

## Methods

### NewHypervisorResourceSearchRequestModel

`func NewHypervisorResourceSearchRequestModel() *HypervisorResourceSearchRequestModel`

NewHypervisorResourceSearchRequestModel instantiates a new HypervisorResourceSearchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourceSearchRequestModelWithDefaults

`func NewHypervisorResourceSearchRequestModelWithDefaults() *HypervisorResourceSearchRequestModel`

NewHypervisorResourceSearchRequestModelWithDefaults instantiates a new HypervisorResourceSearchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *HypervisorResourceSearchRequestModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HypervisorResourceSearchRequestModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HypervisorResourceSearchRequestModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HypervisorResourceSearchRequestModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetResourceType

`func (o *HypervisorResourceSearchRequestModel) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *HypervisorResourceSearchRequestModel) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *HypervisorResourceSearchRequestModel) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *HypervisorResourceSearchRequestModel) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetForwardDirection

`func (o *HypervisorResourceSearchRequestModel) GetForwardDirection() bool`

GetForwardDirection returns the ForwardDirection field if non-nil, zero value otherwise.

### GetForwardDirectionOk

`func (o *HypervisorResourceSearchRequestModel) GetForwardDirectionOk() (*bool, bool)`

GetForwardDirectionOk returns a tuple with the ForwardDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardDirection

`func (o *HypervisorResourceSearchRequestModel) SetForwardDirection(v bool)`

SetForwardDirection sets ForwardDirection field to given value.

### HasForwardDirection

`func (o *HypervisorResourceSearchRequestModel) HasForwardDirection() bool`

HasForwardDirection returns a boolean if a field has been set.

### GetNameContains

`func (o *HypervisorResourceSearchRequestModel) GetNameContains() string`

GetNameContains returns the NameContains field if non-nil, zero value otherwise.

### GetNameContainsOk

`func (o *HypervisorResourceSearchRequestModel) GetNameContainsOk() (*string, bool)`

GetNameContainsOk returns a tuple with the NameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameContains

`func (o *HypervisorResourceSearchRequestModel) SetNameContains(v string)`

SetNameContains sets NameContains field to given value.

### HasNameContains

`func (o *HypervisorResourceSearchRequestModel) HasNameContains() bool`

HasNameContains returns a boolean if a field has been set.

### GetTags

`func (o *HypervisorResourceSearchRequestModel) GetTags() []NameValueStringPairModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HypervisorResourceSearchRequestModel) GetTagsOk() (*[]NameValueStringPairModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HypervisorResourceSearchRequestModel) SetTags(v []NameValueStringPairModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HypervisorResourceSearchRequestModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetId

`func (o *HypervisorResourceSearchRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorResourceSearchRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorResourceSearchRequestModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorResourceSearchRequestModel) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


