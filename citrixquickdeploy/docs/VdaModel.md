# VdaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MachineUid** | Pointer to **int32** |  | [optional] [readonly] 
**Sid** | Pointer to **string** |  | [optional] [readonly] 
**CatalogUid** | Pointer to **int32** |  | [optional] [readonly] 
**RegistrationState** | Pointer to **string** |  | [optional] [readonly] 
**SummaryState** | Pointer to **string** |  | [optional] [readonly] 
**Powerstate** | Pointer to **string** |  | [optional] [readonly] 
**SessionCount** | Pointer to **int32** |  | [optional] [readonly] 
**InMaintenanceMode** | Pointer to **bool** |  | [optional] [readonly] 
**LastDeregistrationTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastDeregistrationReason** | Pointer to **string** |  | [optional] [readonly] 
**VdaResourceGroup** | Pointer to **string** |  | [optional] 
**AllocationType** | Pointer to [**CatalogAllocationType**](CatalogAllocationType.md) |  | [optional] 
**AssociatedUsers** | Pointer to **[]string** |  | [optional] [readonly] 
**RestoreStatus** | Pointer to [**RestoreStatusModel**](RestoreStatusModel.md) |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**ImageOutOfDate** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewVdaModel

`func NewVdaModel() *VdaModel`

NewVdaModel instantiates a new VdaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdaModelWithDefaults

`func NewVdaModelWithDefaults() *VdaModel`

NewVdaModelWithDefaults instantiates a new VdaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *VdaModel) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *VdaModel) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *VdaModel) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *VdaModel) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetName

`func (o *VdaModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VdaModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VdaModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VdaModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMachineUid

`func (o *VdaModel) GetMachineUid() int32`

GetMachineUid returns the MachineUid field if non-nil, zero value otherwise.

### GetMachineUidOk

`func (o *VdaModel) GetMachineUidOk() (*int32, bool)`

GetMachineUidOk returns a tuple with the MachineUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineUid

`func (o *VdaModel) SetMachineUid(v int32)`

SetMachineUid sets MachineUid field to given value.

### HasMachineUid

`func (o *VdaModel) HasMachineUid() bool`

HasMachineUid returns a boolean if a field has been set.

### GetSid

`func (o *VdaModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VdaModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VdaModel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VdaModel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetCatalogUid

`func (o *VdaModel) GetCatalogUid() int32`

GetCatalogUid returns the CatalogUid field if non-nil, zero value otherwise.

### GetCatalogUidOk

`func (o *VdaModel) GetCatalogUidOk() (*int32, bool)`

GetCatalogUidOk returns a tuple with the CatalogUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogUid

`func (o *VdaModel) SetCatalogUid(v int32)`

SetCatalogUid sets CatalogUid field to given value.

### HasCatalogUid

`func (o *VdaModel) HasCatalogUid() bool`

HasCatalogUid returns a boolean if a field has been set.

### GetRegistrationState

`func (o *VdaModel) GetRegistrationState() string`

GetRegistrationState returns the RegistrationState field if non-nil, zero value otherwise.

### GetRegistrationStateOk

`func (o *VdaModel) GetRegistrationStateOk() (*string, bool)`

GetRegistrationStateOk returns a tuple with the RegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationState

`func (o *VdaModel) SetRegistrationState(v string)`

SetRegistrationState sets RegistrationState field to given value.

### HasRegistrationState

`func (o *VdaModel) HasRegistrationState() bool`

HasRegistrationState returns a boolean if a field has been set.

### GetSummaryState

`func (o *VdaModel) GetSummaryState() string`

GetSummaryState returns the SummaryState field if non-nil, zero value otherwise.

### GetSummaryStateOk

`func (o *VdaModel) GetSummaryStateOk() (*string, bool)`

GetSummaryStateOk returns a tuple with the SummaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryState

`func (o *VdaModel) SetSummaryState(v string)`

SetSummaryState sets SummaryState field to given value.

### HasSummaryState

`func (o *VdaModel) HasSummaryState() bool`

HasSummaryState returns a boolean if a field has been set.

### GetPowerstate

`func (o *VdaModel) GetPowerstate() string`

GetPowerstate returns the Powerstate field if non-nil, zero value otherwise.

### GetPowerstateOk

`func (o *VdaModel) GetPowerstateOk() (*string, bool)`

GetPowerstateOk returns a tuple with the Powerstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerstate

`func (o *VdaModel) SetPowerstate(v string)`

SetPowerstate sets Powerstate field to given value.

### HasPowerstate

`func (o *VdaModel) HasPowerstate() bool`

HasPowerstate returns a boolean if a field has been set.

### GetSessionCount

`func (o *VdaModel) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *VdaModel) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *VdaModel) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *VdaModel) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *VdaModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *VdaModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *VdaModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *VdaModel) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.

### GetLastDeregistrationTime

`func (o *VdaModel) GetLastDeregistrationTime() time.Time`

GetLastDeregistrationTime returns the LastDeregistrationTime field if non-nil, zero value otherwise.

### GetLastDeregistrationTimeOk

`func (o *VdaModel) GetLastDeregistrationTimeOk() (*time.Time, bool)`

GetLastDeregistrationTimeOk returns a tuple with the LastDeregistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeregistrationTime

`func (o *VdaModel) SetLastDeregistrationTime(v time.Time)`

SetLastDeregistrationTime sets LastDeregistrationTime field to given value.

### HasLastDeregistrationTime

`func (o *VdaModel) HasLastDeregistrationTime() bool`

HasLastDeregistrationTime returns a boolean if a field has been set.

### GetLastDeregistrationReason

`func (o *VdaModel) GetLastDeregistrationReason() string`

GetLastDeregistrationReason returns the LastDeregistrationReason field if non-nil, zero value otherwise.

### GetLastDeregistrationReasonOk

`func (o *VdaModel) GetLastDeregistrationReasonOk() (*string, bool)`

GetLastDeregistrationReasonOk returns a tuple with the LastDeregistrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeregistrationReason

`func (o *VdaModel) SetLastDeregistrationReason(v string)`

SetLastDeregistrationReason sets LastDeregistrationReason field to given value.

### HasLastDeregistrationReason

`func (o *VdaModel) HasLastDeregistrationReason() bool`

HasLastDeregistrationReason returns a boolean if a field has been set.

### GetVdaResourceGroup

`func (o *VdaModel) GetVdaResourceGroup() string`

GetVdaResourceGroup returns the VdaResourceGroup field if non-nil, zero value otherwise.

### GetVdaResourceGroupOk

`func (o *VdaModel) GetVdaResourceGroupOk() (*string, bool)`

GetVdaResourceGroupOk returns a tuple with the VdaResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaResourceGroup

`func (o *VdaModel) SetVdaResourceGroup(v string)`

SetVdaResourceGroup sets VdaResourceGroup field to given value.

### HasVdaResourceGroup

`func (o *VdaModel) HasVdaResourceGroup() bool`

HasVdaResourceGroup returns a boolean if a field has been set.

### GetAllocationType

`func (o *VdaModel) GetAllocationType() CatalogAllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *VdaModel) GetAllocationTypeOk() (*CatalogAllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *VdaModel) SetAllocationType(v CatalogAllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *VdaModel) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetAssociatedUsers

`func (o *VdaModel) GetAssociatedUsers() []string`

GetAssociatedUsers returns the AssociatedUsers field if non-nil, zero value otherwise.

### GetAssociatedUsersOk

`func (o *VdaModel) GetAssociatedUsersOk() (*[]string, bool)`

GetAssociatedUsersOk returns a tuple with the AssociatedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedUsers

`func (o *VdaModel) SetAssociatedUsers(v []string)`

SetAssociatedUsers sets AssociatedUsers field to given value.

### HasAssociatedUsers

`func (o *VdaModel) HasAssociatedUsers() bool`

HasAssociatedUsers returns a boolean if a field has been set.

### GetRestoreStatus

`func (o *VdaModel) GetRestoreStatus() RestoreStatusModel`

GetRestoreStatus returns the RestoreStatus field if non-nil, zero value otherwise.

### GetRestoreStatusOk

`func (o *VdaModel) GetRestoreStatusOk() (*RestoreStatusModel, bool)`

GetRestoreStatusOk returns a tuple with the RestoreStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreStatus

`func (o *VdaModel) SetRestoreStatus(v RestoreStatusModel)`

SetRestoreStatus sets RestoreStatus field to given value.

### HasRestoreStatus

`func (o *VdaModel) HasRestoreStatus() bool`

HasRestoreStatus returns a boolean if a field has been set.

### GetIpAddress

`func (o *VdaModel) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VdaModel) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VdaModel) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VdaModel) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetImageOutOfDate

`func (o *VdaModel) GetImageOutOfDate() bool`

GetImageOutOfDate returns the ImageOutOfDate field if non-nil, zero value otherwise.

### GetImageOutOfDateOk

`func (o *VdaModel) GetImageOutOfDateOk() (*bool, bool)`

GetImageOutOfDateOk returns a tuple with the ImageOutOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOutOfDate

`func (o *VdaModel) SetImageOutOfDate(v bool)`

SetImageOutOfDate sets ImageOutOfDate field to given value.

### HasImageOutOfDate

`func (o *VdaModel) HasImageOutOfDate() bool`

HasImageOutOfDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


