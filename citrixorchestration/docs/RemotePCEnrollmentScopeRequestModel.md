# RemotePCEnrollmentScopeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OU** | **string** | Specifies the DN of an AD container containing machines allowed to enroll as remote PCs. | 
**IncludeSubfolders** | Pointer to **bool** | Indicates whether machines in subfolders of OU are allowed to enroll as remote PCs. | [optional] [default to false]
**IsOrganizationalUnit** | Pointer to **bool** | Indicates whether this objet is for a OU or for a machine | [optional] [default to false]
**MachinesExcluded** | Pointer to **[]string** | Machines which are explicitly excluded from matching the enrollment scope. | [optional] 
**MachinesIncluded** | Pointer to **[]string** | Machines which are included in the enrollment scope. | [optional] 
**AssignedUsers** | Pointer to **[]string** | The user(s) to whom this machine will be assigned. Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time. | [optional] 

## Methods

### NewRemotePCEnrollmentScopeRequestModel

`func NewRemotePCEnrollmentScopeRequestModel(oU string, ) *RemotePCEnrollmentScopeRequestModel`

NewRemotePCEnrollmentScopeRequestModel instantiates a new RemotePCEnrollmentScopeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemotePCEnrollmentScopeRequestModelWithDefaults

`func NewRemotePCEnrollmentScopeRequestModelWithDefaults() *RemotePCEnrollmentScopeRequestModel`

NewRemotePCEnrollmentScopeRequestModelWithDefaults instantiates a new RemotePCEnrollmentScopeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOU

`func (o *RemotePCEnrollmentScopeRequestModel) GetOU() string`

GetOU returns the OU field if non-nil, zero value otherwise.

### GetOUOk

`func (o *RemotePCEnrollmentScopeRequestModel) GetOUOk() (*string, bool)`

GetOUOk returns a tuple with the OU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOU

`func (o *RemotePCEnrollmentScopeRequestModel) SetOU(v string)`

SetOU sets OU field to given value.


### GetIncludeSubfolders

`func (o *RemotePCEnrollmentScopeRequestModel) GetIncludeSubfolders() bool`

GetIncludeSubfolders returns the IncludeSubfolders field if non-nil, zero value otherwise.

### GetIncludeSubfoldersOk

`func (o *RemotePCEnrollmentScopeRequestModel) GetIncludeSubfoldersOk() (*bool, bool)`

GetIncludeSubfoldersOk returns a tuple with the IncludeSubfolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubfolders

`func (o *RemotePCEnrollmentScopeRequestModel) SetIncludeSubfolders(v bool)`

SetIncludeSubfolders sets IncludeSubfolders field to given value.

### HasIncludeSubfolders

`func (o *RemotePCEnrollmentScopeRequestModel) HasIncludeSubfolders() bool`

HasIncludeSubfolders returns a boolean if a field has been set.

### GetIsOrganizationalUnit

`func (o *RemotePCEnrollmentScopeRequestModel) GetIsOrganizationalUnit() bool`

GetIsOrganizationalUnit returns the IsOrganizationalUnit field if non-nil, zero value otherwise.

### GetIsOrganizationalUnitOk

`func (o *RemotePCEnrollmentScopeRequestModel) GetIsOrganizationalUnitOk() (*bool, bool)`

GetIsOrganizationalUnitOk returns a tuple with the IsOrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizationalUnit

`func (o *RemotePCEnrollmentScopeRequestModel) SetIsOrganizationalUnit(v bool)`

SetIsOrganizationalUnit sets IsOrganizationalUnit field to given value.

### HasIsOrganizationalUnit

`func (o *RemotePCEnrollmentScopeRequestModel) HasIsOrganizationalUnit() bool`

HasIsOrganizationalUnit returns a boolean if a field has been set.

### GetMachinesExcluded

`func (o *RemotePCEnrollmentScopeRequestModel) GetMachinesExcluded() []string`

GetMachinesExcluded returns the MachinesExcluded field if non-nil, zero value otherwise.

### GetMachinesExcludedOk

`func (o *RemotePCEnrollmentScopeRequestModel) GetMachinesExcludedOk() (*[]string, bool)`

GetMachinesExcludedOk returns a tuple with the MachinesExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinesExcluded

`func (o *RemotePCEnrollmentScopeRequestModel) SetMachinesExcluded(v []string)`

SetMachinesExcluded sets MachinesExcluded field to given value.

### HasMachinesExcluded

`func (o *RemotePCEnrollmentScopeRequestModel) HasMachinesExcluded() bool`

HasMachinesExcluded returns a boolean if a field has been set.

### GetMachinesIncluded

`func (o *RemotePCEnrollmentScopeRequestModel) GetMachinesIncluded() []string`

GetMachinesIncluded returns the MachinesIncluded field if non-nil, zero value otherwise.

### GetMachinesIncludedOk

`func (o *RemotePCEnrollmentScopeRequestModel) GetMachinesIncludedOk() (*[]string, bool)`

GetMachinesIncludedOk returns a tuple with the MachinesIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinesIncluded

`func (o *RemotePCEnrollmentScopeRequestModel) SetMachinesIncluded(v []string)`

SetMachinesIncluded sets MachinesIncluded field to given value.

### HasMachinesIncluded

`func (o *RemotePCEnrollmentScopeRequestModel) HasMachinesIncluded() bool`

HasMachinesIncluded returns a boolean if a field has been set.

### GetAssignedUsers

`func (o *RemotePCEnrollmentScopeRequestModel) GetAssignedUsers() []string`

GetAssignedUsers returns the AssignedUsers field if non-nil, zero value otherwise.

### GetAssignedUsersOk

`func (o *RemotePCEnrollmentScopeRequestModel) GetAssignedUsersOk() (*[]string, bool)`

GetAssignedUsersOk returns a tuple with the AssignedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUsers

`func (o *RemotePCEnrollmentScopeRequestModel) SetAssignedUsers(v []string)`

SetAssignedUsers sets AssignedUsers field to given value.

### HasAssignedUsers

`func (o *RemotePCEnrollmentScopeRequestModel) HasAssignedUsers() bool`

HasAssignedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


