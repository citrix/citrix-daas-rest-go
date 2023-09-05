# RemotePCEnrollmentScopeResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OU** | **string** | Specifies the DN of an AD container containing machines allowed to enroll as remote PCs. | 
**IncludeSubfolders** | Pointer to **bool** | Indicates whether machines in subfolders of OU are allowed to enroll as remote PCs. | [optional] [default to false]
**IsOrganizationalUnit** | Pointer to **bool** | Indicates whether this objet is for a OU or for a machine | [optional] [default to false]
**MachinesExcluded** | Pointer to **[]string** | Machines which are explicitly excluded from matching the enrollment scope. | [optional] 
**MachinesIncluded** | Pointer to **[]string** | Machines which are included in the enrollment scope. | [optional] 

## Methods

### NewRemotePCEnrollmentScopeResponseModel

`func NewRemotePCEnrollmentScopeResponseModel(oU string, ) *RemotePCEnrollmentScopeResponseModel`

NewRemotePCEnrollmentScopeResponseModel instantiates a new RemotePCEnrollmentScopeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemotePCEnrollmentScopeResponseModelWithDefaults

`func NewRemotePCEnrollmentScopeResponseModelWithDefaults() *RemotePCEnrollmentScopeResponseModel`

NewRemotePCEnrollmentScopeResponseModelWithDefaults instantiates a new RemotePCEnrollmentScopeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOU

`func (o *RemotePCEnrollmentScopeResponseModel) GetOU() string`

GetOU returns the OU field if non-nil, zero value otherwise.

### GetOUOk

`func (o *RemotePCEnrollmentScopeResponseModel) GetOUOk() (*string, bool)`

GetOUOk returns a tuple with the OU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOU

`func (o *RemotePCEnrollmentScopeResponseModel) SetOU(v string)`

SetOU sets OU field to given value.


### GetIncludeSubfolders

`func (o *RemotePCEnrollmentScopeResponseModel) GetIncludeSubfolders() bool`

GetIncludeSubfolders returns the IncludeSubfolders field if non-nil, zero value otherwise.

### GetIncludeSubfoldersOk

`func (o *RemotePCEnrollmentScopeResponseModel) GetIncludeSubfoldersOk() (*bool, bool)`

GetIncludeSubfoldersOk returns a tuple with the IncludeSubfolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubfolders

`func (o *RemotePCEnrollmentScopeResponseModel) SetIncludeSubfolders(v bool)`

SetIncludeSubfolders sets IncludeSubfolders field to given value.

### HasIncludeSubfolders

`func (o *RemotePCEnrollmentScopeResponseModel) HasIncludeSubfolders() bool`

HasIncludeSubfolders returns a boolean if a field has been set.

### GetIsOrganizationalUnit

`func (o *RemotePCEnrollmentScopeResponseModel) GetIsOrganizationalUnit() bool`

GetIsOrganizationalUnit returns the IsOrganizationalUnit field if non-nil, zero value otherwise.

### GetIsOrganizationalUnitOk

`func (o *RemotePCEnrollmentScopeResponseModel) GetIsOrganizationalUnitOk() (*bool, bool)`

GetIsOrganizationalUnitOk returns a tuple with the IsOrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizationalUnit

`func (o *RemotePCEnrollmentScopeResponseModel) SetIsOrganizationalUnit(v bool)`

SetIsOrganizationalUnit sets IsOrganizationalUnit field to given value.

### HasIsOrganizationalUnit

`func (o *RemotePCEnrollmentScopeResponseModel) HasIsOrganizationalUnit() bool`

HasIsOrganizationalUnit returns a boolean if a field has been set.

### GetMachinesExcluded

`func (o *RemotePCEnrollmentScopeResponseModel) GetMachinesExcluded() []string`

GetMachinesExcluded returns the MachinesExcluded field if non-nil, zero value otherwise.

### GetMachinesExcludedOk

`func (o *RemotePCEnrollmentScopeResponseModel) GetMachinesExcludedOk() (*[]string, bool)`

GetMachinesExcludedOk returns a tuple with the MachinesExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinesExcluded

`func (o *RemotePCEnrollmentScopeResponseModel) SetMachinesExcluded(v []string)`

SetMachinesExcluded sets MachinesExcluded field to given value.

### HasMachinesExcluded

`func (o *RemotePCEnrollmentScopeResponseModel) HasMachinesExcluded() bool`

HasMachinesExcluded returns a boolean if a field has been set.

### GetMachinesIncluded

`func (o *RemotePCEnrollmentScopeResponseModel) GetMachinesIncluded() []string`

GetMachinesIncluded returns the MachinesIncluded field if non-nil, zero value otherwise.

### GetMachinesIncludedOk

`func (o *RemotePCEnrollmentScopeResponseModel) GetMachinesIncludedOk() (*[]string, bool)`

GetMachinesIncludedOk returns a tuple with the MachinesIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinesIncluded

`func (o *RemotePCEnrollmentScopeResponseModel) SetMachinesIncluded(v []string)`

SetMachinesIncluded sets MachinesIncluded field to given value.

### HasMachinesIncluded

`func (o *RemotePCEnrollmentScopeResponseModel) HasMachinesIncluded() bool`

HasMachinesIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


