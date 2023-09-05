# ADDomainModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Domain name | [optional] 
**ForestName** | Pointer to **string** | Parent forest of this domain | [optional] 
**DirectoryEntry** | Pointer to [**ADDomainModelDirectoryEntry**](ADDomainModelDirectoryEntry.md) |  | [optional] 
**Tag** | Pointer to **map[string]interface{}** | Cookie for use by the IActiveDirectoryService implementation | [optional] 

## Methods

### NewADDomainModel

`func NewADDomainModel() *ADDomainModel`

NewADDomainModel instantiates a new ADDomainModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainModelWithDefaults

`func NewADDomainModelWithDefaults() *ADDomainModel`

NewADDomainModelWithDefaults instantiates a new ADDomainModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ADDomainModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ADDomainModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ADDomainModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ADDomainModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetForestName

`func (o *ADDomainModel) GetForestName() string`

GetForestName returns the ForestName field if non-nil, zero value otherwise.

### GetForestNameOk

`func (o *ADDomainModel) GetForestNameOk() (*string, bool)`

GetForestNameOk returns a tuple with the ForestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestName

`func (o *ADDomainModel) SetForestName(v string)`

SetForestName sets ForestName field to given value.

### HasForestName

`func (o *ADDomainModel) HasForestName() bool`

HasForestName returns a boolean if a field has been set.

### GetDirectoryEntry

`func (o *ADDomainModel) GetDirectoryEntry() ADDomainModelDirectoryEntry`

GetDirectoryEntry returns the DirectoryEntry field if non-nil, zero value otherwise.

### GetDirectoryEntryOk

`func (o *ADDomainModel) GetDirectoryEntryOk() (*ADDomainModelDirectoryEntry, bool)`

GetDirectoryEntryOk returns a tuple with the DirectoryEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryEntry

`func (o *ADDomainModel) SetDirectoryEntry(v ADDomainModelDirectoryEntry)`

SetDirectoryEntry sets DirectoryEntry field to given value.

### HasDirectoryEntry

`func (o *ADDomainModel) HasDirectoryEntry() bool`

HasDirectoryEntry returns a boolean if a field has been set.

### GetTag

`func (o *ADDomainModel) GetTag() map[string]interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ADDomainModel) GetTagOk() (*map[string]interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ADDomainModel) SetTag(v map[string]interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ADDomainModel) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


