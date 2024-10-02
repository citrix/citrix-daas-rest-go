# SystemOptimizationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Identity of this system optimization model and the identity of the site to which this configuration applies | [optional] [readonly] 
**EnableCpuSpikeProtection** | Pointer to **bool** | Lets you lower the CPU priority of processes that overload your CPU | [optional] 
**EnableCpuAutoProtection** | Pointer to **bool** | Lets you lower the CPU priority of processes automatically | [optional] 
**CpuUsageLimitOfSpikeProtection** | Pointer to **int32** | CPU usage limit (%) for CPU spike protection | [optional] 
**CpuUsageLimitSampleTimeOfSpikeProtection** | Pointer to **int32** | Limit sample time (s) for CPU spike protection | [optional] 
**IdlePriorityTimeOfSpikeProtection** | Pointer to **int32** | Idle priority time (s) for CPU spike protection | [optional] 
**EnableLimitCpuCoreUsage** | Pointer to **bool** | Lets you limit processes that trigger CPU spike protection to a specified number of logical processors | [optional] 
**CpuCoreLimitOfSpikeProtection** | Pointer to **int32** | CPU core usage limit | [optional] 
**EnableIntelligentCpuOptimization** | Pointer to **bool** | Enable agent to intelligently optimize the CPU priority of processes that trigger CPU spike protection | [optional] 
**EnableIntelligentIoOptimization** | Pointer to **bool** | Enable agent to intelligently optimize the I/O priority of processes that trigger CPU spike protection | [optional] 
**ExcludeProcessesFromCpuSpikeProtection** | Pointer to **bool** | Lets you add or remove processes from an exclusion list for CPU spike protection by executable name | [optional] 
**ProcessesExcludedFromCpuSpikeProtection** | Pointer to **[]string** | Processes excluded from CPU spike protection | [optional] 
**EnableProcessCpuPriority** | Pointer to **bool** | Lets you configure process CPU priority manually | [optional] 
**ProcessesOfCpuPriority** | Pointer to [**[]ProcessCpuPriority**](ProcessCpuPriority.md) | CPU priorities specified for the processes | [optional] 
**EnableProcessAffinity** | Pointer to **bool** | Lets you define how many \&quot;logical processors\&quot; for a process to use | [optional] 
**ProcessesOfCpuAffinity** | Pointer to [**[]ProcessAffinity**](ProcessAffinity.md) | Logical processors specified for the processes | [optional] 
**EnableProcessClamping** | Pointer to **bool** | Lets you prevent processes using more than a certain percentage of the CPU&#39;s processing power | [optional] 
**ProcessesOfCpuClamping** | Pointer to [**[]ProcessClamping**](ProcessClamping.md) | CPU clamping (%) specified for the processes | [optional] 
**EnableMemoryWorkingSetOptimization** | Pointer to **bool** | Lets you force applications that have been idle for a certain time to release excess memory until they are no longer idle | [optional] 
**IdleSampleTimeOfMemoryWorkingSetOptimization** | Pointer to **int32** | Time (min) for which an application must be idle before it is forced to release excess memory | [optional] 
**IdleStateLimitOfMemoryWorkingSetOptimization** | Pointer to **int32** | CPU usage (%) under which a process is considered to be idle | [optional] 
**ExcludeProcessesFromMemoryWorkingSetOptimization** | Pointer to **bool** | Lets you exclude processes from memory management by name | [optional] 
**ProcessesExcludedFromMemoryWorkingSetOptimization** | Pointer to **[]string** | Processes excluded from working set optimization | [optional] 
**EnableProcessIoPriority** | Pointer to **bool** | Lets you configure process I/O priority manually | [optional] 
**ProcessesOfIoPriority** | Pointer to [**[]ProcessIoPriority**](ProcessIoPriority.md) | I/O priorities specified for the processes | [optional] 
**EnableFastLogoff** | Pointer to **bool** | Enables fast logoff for all users in this configuration set | [optional] 
**ExcludeGroupsFromFastLogoff** | Pointer to **bool** | Lets you exclude specific groups of users from Fast Logoff | [optional] 
**GroupsExcludedFromFastLogoff** | Pointer to **[]string** | The groups you exclude from Fast Logoff | [optional] 

## Methods

### NewSystemOptimizationModel

`func NewSystemOptimizationModel() *SystemOptimizationModel`

NewSystemOptimizationModel instantiates a new SystemOptimizationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemOptimizationModelWithDefaults

`func NewSystemOptimizationModelWithDefaults() *SystemOptimizationModel`

NewSystemOptimizationModelWithDefaults instantiates a new SystemOptimizationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemOptimizationModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemOptimizationModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemOptimizationModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SystemOptimizationModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnableCpuSpikeProtection

`func (o *SystemOptimizationModel) GetEnableCpuSpikeProtection() bool`

GetEnableCpuSpikeProtection returns the EnableCpuSpikeProtection field if non-nil, zero value otherwise.

### GetEnableCpuSpikeProtectionOk

`func (o *SystemOptimizationModel) GetEnableCpuSpikeProtectionOk() (*bool, bool)`

GetEnableCpuSpikeProtectionOk returns a tuple with the EnableCpuSpikeProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCpuSpikeProtection

`func (o *SystemOptimizationModel) SetEnableCpuSpikeProtection(v bool)`

SetEnableCpuSpikeProtection sets EnableCpuSpikeProtection field to given value.

### HasEnableCpuSpikeProtection

`func (o *SystemOptimizationModel) HasEnableCpuSpikeProtection() bool`

HasEnableCpuSpikeProtection returns a boolean if a field has been set.

### GetEnableCpuAutoProtection

`func (o *SystemOptimizationModel) GetEnableCpuAutoProtection() bool`

GetEnableCpuAutoProtection returns the EnableCpuAutoProtection field if non-nil, zero value otherwise.

### GetEnableCpuAutoProtectionOk

`func (o *SystemOptimizationModel) GetEnableCpuAutoProtectionOk() (*bool, bool)`

GetEnableCpuAutoProtectionOk returns a tuple with the EnableCpuAutoProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCpuAutoProtection

`func (o *SystemOptimizationModel) SetEnableCpuAutoProtection(v bool)`

SetEnableCpuAutoProtection sets EnableCpuAutoProtection field to given value.

### HasEnableCpuAutoProtection

`func (o *SystemOptimizationModel) HasEnableCpuAutoProtection() bool`

HasEnableCpuAutoProtection returns a boolean if a field has been set.

### GetCpuUsageLimitOfSpikeProtection

`func (o *SystemOptimizationModel) GetCpuUsageLimitOfSpikeProtection() int32`

GetCpuUsageLimitOfSpikeProtection returns the CpuUsageLimitOfSpikeProtection field if non-nil, zero value otherwise.

### GetCpuUsageLimitOfSpikeProtectionOk

`func (o *SystemOptimizationModel) GetCpuUsageLimitOfSpikeProtectionOk() (*int32, bool)`

GetCpuUsageLimitOfSpikeProtectionOk returns a tuple with the CpuUsageLimitOfSpikeProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageLimitOfSpikeProtection

`func (o *SystemOptimizationModel) SetCpuUsageLimitOfSpikeProtection(v int32)`

SetCpuUsageLimitOfSpikeProtection sets CpuUsageLimitOfSpikeProtection field to given value.

### HasCpuUsageLimitOfSpikeProtection

`func (o *SystemOptimizationModel) HasCpuUsageLimitOfSpikeProtection() bool`

HasCpuUsageLimitOfSpikeProtection returns a boolean if a field has been set.

### GetCpuUsageLimitSampleTimeOfSpikeProtection

`func (o *SystemOptimizationModel) GetCpuUsageLimitSampleTimeOfSpikeProtection() int32`

GetCpuUsageLimitSampleTimeOfSpikeProtection returns the CpuUsageLimitSampleTimeOfSpikeProtection field if non-nil, zero value otherwise.

### GetCpuUsageLimitSampleTimeOfSpikeProtectionOk

`func (o *SystemOptimizationModel) GetCpuUsageLimitSampleTimeOfSpikeProtectionOk() (*int32, bool)`

GetCpuUsageLimitSampleTimeOfSpikeProtectionOk returns a tuple with the CpuUsageLimitSampleTimeOfSpikeProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageLimitSampleTimeOfSpikeProtection

`func (o *SystemOptimizationModel) SetCpuUsageLimitSampleTimeOfSpikeProtection(v int32)`

SetCpuUsageLimitSampleTimeOfSpikeProtection sets CpuUsageLimitSampleTimeOfSpikeProtection field to given value.

### HasCpuUsageLimitSampleTimeOfSpikeProtection

`func (o *SystemOptimizationModel) HasCpuUsageLimitSampleTimeOfSpikeProtection() bool`

HasCpuUsageLimitSampleTimeOfSpikeProtection returns a boolean if a field has been set.

### GetIdlePriorityTimeOfSpikeProtection

`func (o *SystemOptimizationModel) GetIdlePriorityTimeOfSpikeProtection() int32`

GetIdlePriorityTimeOfSpikeProtection returns the IdlePriorityTimeOfSpikeProtection field if non-nil, zero value otherwise.

### GetIdlePriorityTimeOfSpikeProtectionOk

`func (o *SystemOptimizationModel) GetIdlePriorityTimeOfSpikeProtectionOk() (*int32, bool)`

GetIdlePriorityTimeOfSpikeProtectionOk returns a tuple with the IdlePriorityTimeOfSpikeProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdlePriorityTimeOfSpikeProtection

`func (o *SystemOptimizationModel) SetIdlePriorityTimeOfSpikeProtection(v int32)`

SetIdlePriorityTimeOfSpikeProtection sets IdlePriorityTimeOfSpikeProtection field to given value.

### HasIdlePriorityTimeOfSpikeProtection

`func (o *SystemOptimizationModel) HasIdlePriorityTimeOfSpikeProtection() bool`

HasIdlePriorityTimeOfSpikeProtection returns a boolean if a field has been set.

### GetEnableLimitCpuCoreUsage

`func (o *SystemOptimizationModel) GetEnableLimitCpuCoreUsage() bool`

GetEnableLimitCpuCoreUsage returns the EnableLimitCpuCoreUsage field if non-nil, zero value otherwise.

### GetEnableLimitCpuCoreUsageOk

`func (o *SystemOptimizationModel) GetEnableLimitCpuCoreUsageOk() (*bool, bool)`

GetEnableLimitCpuCoreUsageOk returns a tuple with the EnableLimitCpuCoreUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLimitCpuCoreUsage

`func (o *SystemOptimizationModel) SetEnableLimitCpuCoreUsage(v bool)`

SetEnableLimitCpuCoreUsage sets EnableLimitCpuCoreUsage field to given value.

### HasEnableLimitCpuCoreUsage

`func (o *SystemOptimizationModel) HasEnableLimitCpuCoreUsage() bool`

HasEnableLimitCpuCoreUsage returns a boolean if a field has been set.

### GetCpuCoreLimitOfSpikeProtection

`func (o *SystemOptimizationModel) GetCpuCoreLimitOfSpikeProtection() int32`

GetCpuCoreLimitOfSpikeProtection returns the CpuCoreLimitOfSpikeProtection field if non-nil, zero value otherwise.

### GetCpuCoreLimitOfSpikeProtectionOk

`func (o *SystemOptimizationModel) GetCpuCoreLimitOfSpikeProtectionOk() (*int32, bool)`

GetCpuCoreLimitOfSpikeProtectionOk returns a tuple with the CpuCoreLimitOfSpikeProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCoreLimitOfSpikeProtection

`func (o *SystemOptimizationModel) SetCpuCoreLimitOfSpikeProtection(v int32)`

SetCpuCoreLimitOfSpikeProtection sets CpuCoreLimitOfSpikeProtection field to given value.

### HasCpuCoreLimitOfSpikeProtection

`func (o *SystemOptimizationModel) HasCpuCoreLimitOfSpikeProtection() bool`

HasCpuCoreLimitOfSpikeProtection returns a boolean if a field has been set.

### GetEnableIntelligentCpuOptimization

`func (o *SystemOptimizationModel) GetEnableIntelligentCpuOptimization() bool`

GetEnableIntelligentCpuOptimization returns the EnableIntelligentCpuOptimization field if non-nil, zero value otherwise.

### GetEnableIntelligentCpuOptimizationOk

`func (o *SystemOptimizationModel) GetEnableIntelligentCpuOptimizationOk() (*bool, bool)`

GetEnableIntelligentCpuOptimizationOk returns a tuple with the EnableIntelligentCpuOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntelligentCpuOptimization

`func (o *SystemOptimizationModel) SetEnableIntelligentCpuOptimization(v bool)`

SetEnableIntelligentCpuOptimization sets EnableIntelligentCpuOptimization field to given value.

### HasEnableIntelligentCpuOptimization

`func (o *SystemOptimizationModel) HasEnableIntelligentCpuOptimization() bool`

HasEnableIntelligentCpuOptimization returns a boolean if a field has been set.

### GetEnableIntelligentIoOptimization

`func (o *SystemOptimizationModel) GetEnableIntelligentIoOptimization() bool`

GetEnableIntelligentIoOptimization returns the EnableIntelligentIoOptimization field if non-nil, zero value otherwise.

### GetEnableIntelligentIoOptimizationOk

`func (o *SystemOptimizationModel) GetEnableIntelligentIoOptimizationOk() (*bool, bool)`

GetEnableIntelligentIoOptimizationOk returns a tuple with the EnableIntelligentIoOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntelligentIoOptimization

`func (o *SystemOptimizationModel) SetEnableIntelligentIoOptimization(v bool)`

SetEnableIntelligentIoOptimization sets EnableIntelligentIoOptimization field to given value.

### HasEnableIntelligentIoOptimization

`func (o *SystemOptimizationModel) HasEnableIntelligentIoOptimization() bool`

HasEnableIntelligentIoOptimization returns a boolean if a field has been set.

### GetExcludeProcessesFromCpuSpikeProtection

`func (o *SystemOptimizationModel) GetExcludeProcessesFromCpuSpikeProtection() bool`

GetExcludeProcessesFromCpuSpikeProtection returns the ExcludeProcessesFromCpuSpikeProtection field if non-nil, zero value otherwise.

### GetExcludeProcessesFromCpuSpikeProtectionOk

`func (o *SystemOptimizationModel) GetExcludeProcessesFromCpuSpikeProtectionOk() (*bool, bool)`

GetExcludeProcessesFromCpuSpikeProtectionOk returns a tuple with the ExcludeProcessesFromCpuSpikeProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeProcessesFromCpuSpikeProtection

`func (o *SystemOptimizationModel) SetExcludeProcessesFromCpuSpikeProtection(v bool)`

SetExcludeProcessesFromCpuSpikeProtection sets ExcludeProcessesFromCpuSpikeProtection field to given value.

### HasExcludeProcessesFromCpuSpikeProtection

`func (o *SystemOptimizationModel) HasExcludeProcessesFromCpuSpikeProtection() bool`

HasExcludeProcessesFromCpuSpikeProtection returns a boolean if a field has been set.

### GetProcessesExcludedFromCpuSpikeProtection

`func (o *SystemOptimizationModel) GetProcessesExcludedFromCpuSpikeProtection() []string`

GetProcessesExcludedFromCpuSpikeProtection returns the ProcessesExcludedFromCpuSpikeProtection field if non-nil, zero value otherwise.

### GetProcessesExcludedFromCpuSpikeProtectionOk

`func (o *SystemOptimizationModel) GetProcessesExcludedFromCpuSpikeProtectionOk() (*[]string, bool)`

GetProcessesExcludedFromCpuSpikeProtectionOk returns a tuple with the ProcessesExcludedFromCpuSpikeProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessesExcludedFromCpuSpikeProtection

`func (o *SystemOptimizationModel) SetProcessesExcludedFromCpuSpikeProtection(v []string)`

SetProcessesExcludedFromCpuSpikeProtection sets ProcessesExcludedFromCpuSpikeProtection field to given value.

### HasProcessesExcludedFromCpuSpikeProtection

`func (o *SystemOptimizationModel) HasProcessesExcludedFromCpuSpikeProtection() bool`

HasProcessesExcludedFromCpuSpikeProtection returns a boolean if a field has been set.

### GetEnableProcessCpuPriority

`func (o *SystemOptimizationModel) GetEnableProcessCpuPriority() bool`

GetEnableProcessCpuPriority returns the EnableProcessCpuPriority field if non-nil, zero value otherwise.

### GetEnableProcessCpuPriorityOk

`func (o *SystemOptimizationModel) GetEnableProcessCpuPriorityOk() (*bool, bool)`

GetEnableProcessCpuPriorityOk returns a tuple with the EnableProcessCpuPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProcessCpuPriority

`func (o *SystemOptimizationModel) SetEnableProcessCpuPriority(v bool)`

SetEnableProcessCpuPriority sets EnableProcessCpuPriority field to given value.

### HasEnableProcessCpuPriority

`func (o *SystemOptimizationModel) HasEnableProcessCpuPriority() bool`

HasEnableProcessCpuPriority returns a boolean if a field has been set.

### GetProcessesOfCpuPriority

`func (o *SystemOptimizationModel) GetProcessesOfCpuPriority() []ProcessCpuPriority`

GetProcessesOfCpuPriority returns the ProcessesOfCpuPriority field if non-nil, zero value otherwise.

### GetProcessesOfCpuPriorityOk

`func (o *SystemOptimizationModel) GetProcessesOfCpuPriorityOk() (*[]ProcessCpuPriority, bool)`

GetProcessesOfCpuPriorityOk returns a tuple with the ProcessesOfCpuPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessesOfCpuPriority

`func (o *SystemOptimizationModel) SetProcessesOfCpuPriority(v []ProcessCpuPriority)`

SetProcessesOfCpuPriority sets ProcessesOfCpuPriority field to given value.

### HasProcessesOfCpuPriority

`func (o *SystemOptimizationModel) HasProcessesOfCpuPriority() bool`

HasProcessesOfCpuPriority returns a boolean if a field has been set.

### GetEnableProcessAffinity

`func (o *SystemOptimizationModel) GetEnableProcessAffinity() bool`

GetEnableProcessAffinity returns the EnableProcessAffinity field if non-nil, zero value otherwise.

### GetEnableProcessAffinityOk

`func (o *SystemOptimizationModel) GetEnableProcessAffinityOk() (*bool, bool)`

GetEnableProcessAffinityOk returns a tuple with the EnableProcessAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProcessAffinity

`func (o *SystemOptimizationModel) SetEnableProcessAffinity(v bool)`

SetEnableProcessAffinity sets EnableProcessAffinity field to given value.

### HasEnableProcessAffinity

`func (o *SystemOptimizationModel) HasEnableProcessAffinity() bool`

HasEnableProcessAffinity returns a boolean if a field has been set.

### GetProcessesOfCpuAffinity

`func (o *SystemOptimizationModel) GetProcessesOfCpuAffinity() []ProcessAffinity`

GetProcessesOfCpuAffinity returns the ProcessesOfCpuAffinity field if non-nil, zero value otherwise.

### GetProcessesOfCpuAffinityOk

`func (o *SystemOptimizationModel) GetProcessesOfCpuAffinityOk() (*[]ProcessAffinity, bool)`

GetProcessesOfCpuAffinityOk returns a tuple with the ProcessesOfCpuAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessesOfCpuAffinity

`func (o *SystemOptimizationModel) SetProcessesOfCpuAffinity(v []ProcessAffinity)`

SetProcessesOfCpuAffinity sets ProcessesOfCpuAffinity field to given value.

### HasProcessesOfCpuAffinity

`func (o *SystemOptimizationModel) HasProcessesOfCpuAffinity() bool`

HasProcessesOfCpuAffinity returns a boolean if a field has been set.

### GetEnableProcessClamping

`func (o *SystemOptimizationModel) GetEnableProcessClamping() bool`

GetEnableProcessClamping returns the EnableProcessClamping field if non-nil, zero value otherwise.

### GetEnableProcessClampingOk

`func (o *SystemOptimizationModel) GetEnableProcessClampingOk() (*bool, bool)`

GetEnableProcessClampingOk returns a tuple with the EnableProcessClamping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProcessClamping

`func (o *SystemOptimizationModel) SetEnableProcessClamping(v bool)`

SetEnableProcessClamping sets EnableProcessClamping field to given value.

### HasEnableProcessClamping

`func (o *SystemOptimizationModel) HasEnableProcessClamping() bool`

HasEnableProcessClamping returns a boolean if a field has been set.

### GetProcessesOfCpuClamping

`func (o *SystemOptimizationModel) GetProcessesOfCpuClamping() []ProcessClamping`

GetProcessesOfCpuClamping returns the ProcessesOfCpuClamping field if non-nil, zero value otherwise.

### GetProcessesOfCpuClampingOk

`func (o *SystemOptimizationModel) GetProcessesOfCpuClampingOk() (*[]ProcessClamping, bool)`

GetProcessesOfCpuClampingOk returns a tuple with the ProcessesOfCpuClamping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessesOfCpuClamping

`func (o *SystemOptimizationModel) SetProcessesOfCpuClamping(v []ProcessClamping)`

SetProcessesOfCpuClamping sets ProcessesOfCpuClamping field to given value.

### HasProcessesOfCpuClamping

`func (o *SystemOptimizationModel) HasProcessesOfCpuClamping() bool`

HasProcessesOfCpuClamping returns a boolean if a field has been set.

### GetEnableMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) GetEnableMemoryWorkingSetOptimization() bool`

GetEnableMemoryWorkingSetOptimization returns the EnableMemoryWorkingSetOptimization field if non-nil, zero value otherwise.

### GetEnableMemoryWorkingSetOptimizationOk

`func (o *SystemOptimizationModel) GetEnableMemoryWorkingSetOptimizationOk() (*bool, bool)`

GetEnableMemoryWorkingSetOptimizationOk returns a tuple with the EnableMemoryWorkingSetOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) SetEnableMemoryWorkingSetOptimization(v bool)`

SetEnableMemoryWorkingSetOptimization sets EnableMemoryWorkingSetOptimization field to given value.

### HasEnableMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) HasEnableMemoryWorkingSetOptimization() bool`

HasEnableMemoryWorkingSetOptimization returns a boolean if a field has been set.

### GetIdleSampleTimeOfMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) GetIdleSampleTimeOfMemoryWorkingSetOptimization() int32`

GetIdleSampleTimeOfMemoryWorkingSetOptimization returns the IdleSampleTimeOfMemoryWorkingSetOptimization field if non-nil, zero value otherwise.

### GetIdleSampleTimeOfMemoryWorkingSetOptimizationOk

`func (o *SystemOptimizationModel) GetIdleSampleTimeOfMemoryWorkingSetOptimizationOk() (*int32, bool)`

GetIdleSampleTimeOfMemoryWorkingSetOptimizationOk returns a tuple with the IdleSampleTimeOfMemoryWorkingSetOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleSampleTimeOfMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) SetIdleSampleTimeOfMemoryWorkingSetOptimization(v int32)`

SetIdleSampleTimeOfMemoryWorkingSetOptimization sets IdleSampleTimeOfMemoryWorkingSetOptimization field to given value.

### HasIdleSampleTimeOfMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) HasIdleSampleTimeOfMemoryWorkingSetOptimization() bool`

HasIdleSampleTimeOfMemoryWorkingSetOptimization returns a boolean if a field has been set.

### GetIdleStateLimitOfMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) GetIdleStateLimitOfMemoryWorkingSetOptimization() int32`

GetIdleStateLimitOfMemoryWorkingSetOptimization returns the IdleStateLimitOfMemoryWorkingSetOptimization field if non-nil, zero value otherwise.

### GetIdleStateLimitOfMemoryWorkingSetOptimizationOk

`func (o *SystemOptimizationModel) GetIdleStateLimitOfMemoryWorkingSetOptimizationOk() (*int32, bool)`

GetIdleStateLimitOfMemoryWorkingSetOptimizationOk returns a tuple with the IdleStateLimitOfMemoryWorkingSetOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleStateLimitOfMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) SetIdleStateLimitOfMemoryWorkingSetOptimization(v int32)`

SetIdleStateLimitOfMemoryWorkingSetOptimization sets IdleStateLimitOfMemoryWorkingSetOptimization field to given value.

### HasIdleStateLimitOfMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) HasIdleStateLimitOfMemoryWorkingSetOptimization() bool`

HasIdleStateLimitOfMemoryWorkingSetOptimization returns a boolean if a field has been set.

### GetExcludeProcessesFromMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) GetExcludeProcessesFromMemoryWorkingSetOptimization() bool`

GetExcludeProcessesFromMemoryWorkingSetOptimization returns the ExcludeProcessesFromMemoryWorkingSetOptimization field if non-nil, zero value otherwise.

### GetExcludeProcessesFromMemoryWorkingSetOptimizationOk

`func (o *SystemOptimizationModel) GetExcludeProcessesFromMemoryWorkingSetOptimizationOk() (*bool, bool)`

GetExcludeProcessesFromMemoryWorkingSetOptimizationOk returns a tuple with the ExcludeProcessesFromMemoryWorkingSetOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeProcessesFromMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) SetExcludeProcessesFromMemoryWorkingSetOptimization(v bool)`

SetExcludeProcessesFromMemoryWorkingSetOptimization sets ExcludeProcessesFromMemoryWorkingSetOptimization field to given value.

### HasExcludeProcessesFromMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) HasExcludeProcessesFromMemoryWorkingSetOptimization() bool`

HasExcludeProcessesFromMemoryWorkingSetOptimization returns a boolean if a field has been set.

### GetProcessesExcludedFromMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) GetProcessesExcludedFromMemoryWorkingSetOptimization() []string`

GetProcessesExcludedFromMemoryWorkingSetOptimization returns the ProcessesExcludedFromMemoryWorkingSetOptimization field if non-nil, zero value otherwise.

### GetProcessesExcludedFromMemoryWorkingSetOptimizationOk

`func (o *SystemOptimizationModel) GetProcessesExcludedFromMemoryWorkingSetOptimizationOk() (*[]string, bool)`

GetProcessesExcludedFromMemoryWorkingSetOptimizationOk returns a tuple with the ProcessesExcludedFromMemoryWorkingSetOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessesExcludedFromMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) SetProcessesExcludedFromMemoryWorkingSetOptimization(v []string)`

SetProcessesExcludedFromMemoryWorkingSetOptimization sets ProcessesExcludedFromMemoryWorkingSetOptimization field to given value.

### HasProcessesExcludedFromMemoryWorkingSetOptimization

`func (o *SystemOptimizationModel) HasProcessesExcludedFromMemoryWorkingSetOptimization() bool`

HasProcessesExcludedFromMemoryWorkingSetOptimization returns a boolean if a field has been set.

### GetEnableProcessIoPriority

`func (o *SystemOptimizationModel) GetEnableProcessIoPriority() bool`

GetEnableProcessIoPriority returns the EnableProcessIoPriority field if non-nil, zero value otherwise.

### GetEnableProcessIoPriorityOk

`func (o *SystemOptimizationModel) GetEnableProcessIoPriorityOk() (*bool, bool)`

GetEnableProcessIoPriorityOk returns a tuple with the EnableProcessIoPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProcessIoPriority

`func (o *SystemOptimizationModel) SetEnableProcessIoPriority(v bool)`

SetEnableProcessIoPriority sets EnableProcessIoPriority field to given value.

### HasEnableProcessIoPriority

`func (o *SystemOptimizationModel) HasEnableProcessIoPriority() bool`

HasEnableProcessIoPriority returns a boolean if a field has been set.

### GetProcessesOfIoPriority

`func (o *SystemOptimizationModel) GetProcessesOfIoPriority() []ProcessIoPriority`

GetProcessesOfIoPriority returns the ProcessesOfIoPriority field if non-nil, zero value otherwise.

### GetProcessesOfIoPriorityOk

`func (o *SystemOptimizationModel) GetProcessesOfIoPriorityOk() (*[]ProcessIoPriority, bool)`

GetProcessesOfIoPriorityOk returns a tuple with the ProcessesOfIoPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessesOfIoPriority

`func (o *SystemOptimizationModel) SetProcessesOfIoPriority(v []ProcessIoPriority)`

SetProcessesOfIoPriority sets ProcessesOfIoPriority field to given value.

### HasProcessesOfIoPriority

`func (o *SystemOptimizationModel) HasProcessesOfIoPriority() bool`

HasProcessesOfIoPriority returns a boolean if a field has been set.

### GetEnableFastLogoff

`func (o *SystemOptimizationModel) GetEnableFastLogoff() bool`

GetEnableFastLogoff returns the EnableFastLogoff field if non-nil, zero value otherwise.

### GetEnableFastLogoffOk

`func (o *SystemOptimizationModel) GetEnableFastLogoffOk() (*bool, bool)`

GetEnableFastLogoffOk returns a tuple with the EnableFastLogoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFastLogoff

`func (o *SystemOptimizationModel) SetEnableFastLogoff(v bool)`

SetEnableFastLogoff sets EnableFastLogoff field to given value.

### HasEnableFastLogoff

`func (o *SystemOptimizationModel) HasEnableFastLogoff() bool`

HasEnableFastLogoff returns a boolean if a field has been set.

### GetExcludeGroupsFromFastLogoff

`func (o *SystemOptimizationModel) GetExcludeGroupsFromFastLogoff() bool`

GetExcludeGroupsFromFastLogoff returns the ExcludeGroupsFromFastLogoff field if non-nil, zero value otherwise.

### GetExcludeGroupsFromFastLogoffOk

`func (o *SystemOptimizationModel) GetExcludeGroupsFromFastLogoffOk() (*bool, bool)`

GetExcludeGroupsFromFastLogoffOk returns a tuple with the ExcludeGroupsFromFastLogoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGroupsFromFastLogoff

`func (o *SystemOptimizationModel) SetExcludeGroupsFromFastLogoff(v bool)`

SetExcludeGroupsFromFastLogoff sets ExcludeGroupsFromFastLogoff field to given value.

### HasExcludeGroupsFromFastLogoff

`func (o *SystemOptimizationModel) HasExcludeGroupsFromFastLogoff() bool`

HasExcludeGroupsFromFastLogoff returns a boolean if a field has been set.

### GetGroupsExcludedFromFastLogoff

`func (o *SystemOptimizationModel) GetGroupsExcludedFromFastLogoff() []string`

GetGroupsExcludedFromFastLogoff returns the GroupsExcludedFromFastLogoff field if non-nil, zero value otherwise.

### GetGroupsExcludedFromFastLogoffOk

`func (o *SystemOptimizationModel) GetGroupsExcludedFromFastLogoffOk() (*[]string, bool)`

GetGroupsExcludedFromFastLogoffOk returns a tuple with the GroupsExcludedFromFastLogoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsExcludedFromFastLogoff

`func (o *SystemOptimizationModel) SetGroupsExcludedFromFastLogoff(v []string)`

SetGroupsExcludedFromFastLogoff sets GroupsExcludedFromFastLogoff field to given value.

### HasGroupsExcludedFromFastLogoff

`func (o *SystemOptimizationModel) HasGroupsExcludedFromFastLogoff() bool`

HasGroupsExcludedFromFastLogoff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


