# MachineIdentityPoolResponseModelDefaultOUDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Domain name | [optional] 
**ForestName** | Pointer to **string** | Parent forest of this domain | [optional] 
**DirectoryEntry** | Pointer to [**ADDomainModelDirectoryEntry**](ADDomainModelDirectoryEntry.md) |  | [optional] 
**Tag** | Pointer to **map[string]interface{}** | Cookie for use by the IActiveDirectoryService implementation | [optional] 

## Methods

### NewMachineIdentityPoolResponseModelDefaultOUDomain

`func NewMachineIdentityPoolResponseModelDefaultOUDomain() *MachineIdentityPoolResponseModelDefaultOUDomain`

NewMachineIdentityPoolResponseModelDefaultOUDomain instantiates a new MachineIdentityPoolResponseModelDefaultOUDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineIdentityPoolResponseModelDefaultOUDomainWithDefaults

`func NewMachineIdentityPoolResponseModelDefaultOUDomainWithDefaults() *MachineIdentityPoolResponseModelDefaultOUDomain`

NewMachineIdentityPoolResponseModelDefaultOUDomainWithDefaults instantiates a new MachineIdentityPoolResponseModelDefaultOUDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetForestName

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) GetForestName() string`

GetForestName returns the ForestName field if non-nil, zero value otherwise.

### GetForestNameOk

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) GetForestNameOk() (*string, bool)`

GetForestNameOk returns a tuple with the ForestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestName

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) SetForestName(v string)`

SetForestName sets ForestName field to given value.

### HasForestName

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) HasForestName() bool`

HasForestName returns a boolean if a field has been set.

### GetDirectoryEntry

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) GetDirectoryEntry() ADDomainModelDirectoryEntry`

GetDirectoryEntry returns the DirectoryEntry field if non-nil, zero value otherwise.

### GetDirectoryEntryOk

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) GetDirectoryEntryOk() (*ADDomainModelDirectoryEntry, bool)`

GetDirectoryEntryOk returns a tuple with the DirectoryEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryEntry

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) SetDirectoryEntry(v ADDomainModelDirectoryEntry)`

SetDirectoryEntry sets DirectoryEntry field to given value.

### HasDirectoryEntry

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) HasDirectoryEntry() bool`

HasDirectoryEntry returns a boolean if a field has been set.

### GetTag

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) GetTag() map[string]interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) GetTagOk() (*map[string]interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) SetTag(v map[string]interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *MachineIdentityPoolResponseModelDefaultOUDomain) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


