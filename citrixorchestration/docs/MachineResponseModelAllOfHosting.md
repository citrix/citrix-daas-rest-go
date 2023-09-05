# MachineResponseModelAllOfHosting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostedMachineId** | Pointer to **string** | Unique ID within the hosting unit of the target managed machine. | [optional] 
**HostedMachineName** | Pointer to **string** | The friendly name of a hosted machine as used by its hypervisor. This is not necessarily the DNS name of the machine. | [optional] 
**HostingServerName** | Pointer to **string** | DNS name of the hypervisor that is hosting the machine. | [optional] 
**LastHostingUpdateTime** | Pointer to **string** | Time of last update to any hosting data (such as power states) for this machine reported by the hypervisor connection. | [optional] 
**FormattedLastHostingUpdateTime** | Pointer to **string** | Formatted time of last update to any hosting data (such as power states) for this machine reported by the hypervisor connection. RFC 3339 compatible format. | [optional] 
**HypervisorConnection** | Pointer to [**MachineHostingResponseModelHypervisorConnection**](MachineHostingResponseModelHypervisorConnection.md) |  | [optional] 
**ImageOutOfDate** | Pointer to **bool** | Indicates whether the machine image matches the latest image specified for the machine catalog. | [optional] 

## Methods

### NewMachineResponseModelAllOfHosting

`func NewMachineResponseModelAllOfHosting() *MachineResponseModelAllOfHosting`

NewMachineResponseModelAllOfHosting instantiates a new MachineResponseModelAllOfHosting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineResponseModelAllOfHostingWithDefaults

`func NewMachineResponseModelAllOfHostingWithDefaults() *MachineResponseModelAllOfHosting`

NewMachineResponseModelAllOfHostingWithDefaults instantiates a new MachineResponseModelAllOfHosting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostedMachineId

`func (o *MachineResponseModelAllOfHosting) GetHostedMachineId() string`

GetHostedMachineId returns the HostedMachineId field if non-nil, zero value otherwise.

### GetHostedMachineIdOk

`func (o *MachineResponseModelAllOfHosting) GetHostedMachineIdOk() (*string, bool)`

GetHostedMachineIdOk returns a tuple with the HostedMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedMachineId

`func (o *MachineResponseModelAllOfHosting) SetHostedMachineId(v string)`

SetHostedMachineId sets HostedMachineId field to given value.

### HasHostedMachineId

`func (o *MachineResponseModelAllOfHosting) HasHostedMachineId() bool`

HasHostedMachineId returns a boolean if a field has been set.

### GetHostedMachineName

`func (o *MachineResponseModelAllOfHosting) GetHostedMachineName() string`

GetHostedMachineName returns the HostedMachineName field if non-nil, zero value otherwise.

### GetHostedMachineNameOk

`func (o *MachineResponseModelAllOfHosting) GetHostedMachineNameOk() (*string, bool)`

GetHostedMachineNameOk returns a tuple with the HostedMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedMachineName

`func (o *MachineResponseModelAllOfHosting) SetHostedMachineName(v string)`

SetHostedMachineName sets HostedMachineName field to given value.

### HasHostedMachineName

`func (o *MachineResponseModelAllOfHosting) HasHostedMachineName() bool`

HasHostedMachineName returns a boolean if a field has been set.

### GetHostingServerName

`func (o *MachineResponseModelAllOfHosting) GetHostingServerName() string`

GetHostingServerName returns the HostingServerName field if non-nil, zero value otherwise.

### GetHostingServerNameOk

`func (o *MachineResponseModelAllOfHosting) GetHostingServerNameOk() (*string, bool)`

GetHostingServerNameOk returns a tuple with the HostingServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingServerName

`func (o *MachineResponseModelAllOfHosting) SetHostingServerName(v string)`

SetHostingServerName sets HostingServerName field to given value.

### HasHostingServerName

`func (o *MachineResponseModelAllOfHosting) HasHostingServerName() bool`

HasHostingServerName returns a boolean if a field has been set.

### GetLastHostingUpdateTime

`func (o *MachineResponseModelAllOfHosting) GetLastHostingUpdateTime() string`

GetLastHostingUpdateTime returns the LastHostingUpdateTime field if non-nil, zero value otherwise.

### GetLastHostingUpdateTimeOk

`func (o *MachineResponseModelAllOfHosting) GetLastHostingUpdateTimeOk() (*string, bool)`

GetLastHostingUpdateTimeOk returns a tuple with the LastHostingUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHostingUpdateTime

`func (o *MachineResponseModelAllOfHosting) SetLastHostingUpdateTime(v string)`

SetLastHostingUpdateTime sets LastHostingUpdateTime field to given value.

### HasLastHostingUpdateTime

`func (o *MachineResponseModelAllOfHosting) HasLastHostingUpdateTime() bool`

HasLastHostingUpdateTime returns a boolean if a field has been set.

### GetFormattedLastHostingUpdateTime

`func (o *MachineResponseModelAllOfHosting) GetFormattedLastHostingUpdateTime() string`

GetFormattedLastHostingUpdateTime returns the FormattedLastHostingUpdateTime field if non-nil, zero value otherwise.

### GetFormattedLastHostingUpdateTimeOk

`func (o *MachineResponseModelAllOfHosting) GetFormattedLastHostingUpdateTimeOk() (*string, bool)`

GetFormattedLastHostingUpdateTimeOk returns a tuple with the FormattedLastHostingUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedLastHostingUpdateTime

`func (o *MachineResponseModelAllOfHosting) SetFormattedLastHostingUpdateTime(v string)`

SetFormattedLastHostingUpdateTime sets FormattedLastHostingUpdateTime field to given value.

### HasFormattedLastHostingUpdateTime

`func (o *MachineResponseModelAllOfHosting) HasFormattedLastHostingUpdateTime() bool`

HasFormattedLastHostingUpdateTime returns a boolean if a field has been set.

### GetHypervisorConnection

`func (o *MachineResponseModelAllOfHosting) GetHypervisorConnection() MachineHostingResponseModelHypervisorConnection`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *MachineResponseModelAllOfHosting) GetHypervisorConnectionOk() (*MachineHostingResponseModelHypervisorConnection, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *MachineResponseModelAllOfHosting) SetHypervisorConnection(v MachineHostingResponseModelHypervisorConnection)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *MachineResponseModelAllOfHosting) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### GetImageOutOfDate

`func (o *MachineResponseModelAllOfHosting) GetImageOutOfDate() bool`

GetImageOutOfDate returns the ImageOutOfDate field if non-nil, zero value otherwise.

### GetImageOutOfDateOk

`func (o *MachineResponseModelAllOfHosting) GetImageOutOfDateOk() (*bool, bool)`

GetImageOutOfDateOk returns a tuple with the ImageOutOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOutOfDate

`func (o *MachineResponseModelAllOfHosting) SetImageOutOfDate(v bool)`

SetImageOutOfDate sets ImageOutOfDate field to given value.

### HasImageOutOfDate

`func (o *MachineResponseModelAllOfHosting) HasImageOutOfDate() bool`

HasImageOutOfDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


