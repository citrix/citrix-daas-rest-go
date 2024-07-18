# ProvisionedVirtualMachineDetailsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceleratedNetwork** | Pointer to **NullableBool** | Accelerated Network | [optional] 
**ADAccountName** | Pointer to **NullableString** | Active Directory Account Name | [optional] 
**VMSid** | Pointer to **NullableString** | Virtual Machine Sid | [optional] 
**AssignedImage** | Pointer to **NullableString** | Assigned Image | [optional] 
**BootedImage** | Pointer to **NullableString** | Booted Image | [optional] 
**CustomVmData** | Pointer to [**ProvisionedVirtualMachineCustomVmDataResponseModel**](ProvisionedVirtualMachineCustomVmDataResponseModel.md) |  | [optional] 
**CpuCount** | Pointer to **int32** | Cpu Count | [optional] 
**CreationDate** | Pointer to **NullableString** | Creation Date | [optional] 
**Domain** | Pointer to **NullableString** | Domain | [optional] 
**EnableIntegrityMonitoring** | Pointer to **NullableBool** | IntegrityMonitoring | [optional] 
**EnableSecureBoot** | Pointer to **NullableBool** | From ProvScheme VMMetaData | [optional] 
**EnableVTPM** | Pointer to **NullableBool** | From ProvScheme VMMetaData | [optional] 
**EncryptionAtHost** | Pointer to **NullableBool** | From ProvScheme VMMetaData | [optional] 
**Errors** | Pointer to **[]string** | The Errors of provisioned virtual machine  | [optional] 
**PluginId** | Pointer to **NullableString** | Broker Hypervisor Connection Type - Plugin Factory Name | [optional] 
**ImageOutOfDate** | Pointer to **bool** | Image Out Of Date, booted image name doesn&#39;t equal assigned image name | [optional] 
**MemoryMB** | Pointer to **int32** | MemoryMB | [optional] 
**HostingUnitUid** | Pointer to **string** | Hosting Unit Uid | [optional] 
**HypervisorConnectionUid** | Pointer to **string** | Hypervisor Connection Uid | [optional] 
**IdentityDiskId** | Pointer to **NullableString** | Identity Disk Id | [optional] 
**LastBootTime** | Pointer to **NullableString** | Last Boot Time | [optional] 
**ProvisioningSchemeType** | Pointer to **NullableString** | Provisioning Scheme Type: MCS, PVS | [optional] 
**ProvisioningSchemeUpdateRequested** | Pointer to **NullableString** | Provisioning Scheme Update Requested | [optional] 
**ProvisioningSchemeVersion** | Pointer to **int32** | Provisioning SchemeVersion | [optional] 
**UseFullDiskCloneProvisioning** | Pointer to **bool** | Use Full Disk Clone Provisioning | [optional] 
**VMId** | Pointer to **NullableString** | VMId, Azure - resourcegroup/VMName | [optional] 
**VMName** | Pointer to **NullableString** | VM Name on the hypervisor | [optional] 
**Warnings** | Pointer to [**[]ProvisionedVirtualMachineDetailsWarningReponseModel**](ProvisionedVirtualMachineDetailsWarningReponseModel.md) | Warnings, decode CustomVMData failed reason | [optional] 
**WindowsActivationStatus** | Pointer to **NullableString** | Windows Activation Status | [optional] 
**WindowsActivationStatusErrorCode** | Pointer to **NullableString** | Windows Activation Status Error Code | [optional] 
**WindowsActivationStatusVirtualMachineError** | Pointer to **NullableString** | Windows Activation Status Virtual Machine Error | [optional] 
**WindowsActivationTypeProvisionedVirtualMachine** | Pointer to **NullableString** | Windows Activation Type Provisioned Virtual Machine | [optional] 
**WriteBackCacheDiskSize** | Pointer to **int32** | Write Back Cache(WBC) Disk Size | [optional] 
**WriteBackCacheDriveLetter** | Pointer to **string** | Write Back Cache(WBC) Drive Letter | [optional] 
**WriteBackCacheMemorySize** | Pointer to **int32** | Write Back Cache(WBC) Memory Size | [optional] 

## Methods

### NewProvisionedVirtualMachineDetailsResponseModel

`func NewProvisionedVirtualMachineDetailsResponseModel() *ProvisionedVirtualMachineDetailsResponseModel`

NewProvisionedVirtualMachineDetailsResponseModel instantiates a new ProvisionedVirtualMachineDetailsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedVirtualMachineDetailsResponseModelWithDefaults

`func NewProvisionedVirtualMachineDetailsResponseModelWithDefaults() *ProvisionedVirtualMachineDetailsResponseModel`

NewProvisionedVirtualMachineDetailsResponseModelWithDefaults instantiates a new ProvisionedVirtualMachineDetailsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceleratedNetwork

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetAcceleratedNetwork() bool`

GetAcceleratedNetwork returns the AcceleratedNetwork field if non-nil, zero value otherwise.

### GetAcceleratedNetworkOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetAcceleratedNetworkOk() (*bool, bool)`

GetAcceleratedNetworkOk returns a tuple with the AcceleratedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratedNetwork

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetAcceleratedNetwork(v bool)`

SetAcceleratedNetwork sets AcceleratedNetwork field to given value.

### HasAcceleratedNetwork

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasAcceleratedNetwork() bool`

HasAcceleratedNetwork returns a boolean if a field has been set.

### SetAcceleratedNetworkNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetAcceleratedNetworkNil(b bool)`

 SetAcceleratedNetworkNil sets the value for AcceleratedNetwork to be an explicit nil

### UnsetAcceleratedNetwork
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetAcceleratedNetwork()`

UnsetAcceleratedNetwork ensures that no value is present for AcceleratedNetwork, not even an explicit nil
### GetADAccountName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetADAccountName() string`

GetADAccountName returns the ADAccountName field if non-nil, zero value otherwise.

### GetADAccountNameOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetADAccountNameOk() (*string, bool)`

GetADAccountNameOk returns a tuple with the ADAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADAccountName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetADAccountName(v string)`

SetADAccountName sets ADAccountName field to given value.

### HasADAccountName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasADAccountName() bool`

HasADAccountName returns a boolean if a field has been set.

### SetADAccountNameNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetADAccountNameNil(b bool)`

 SetADAccountNameNil sets the value for ADAccountName to be an explicit nil

### UnsetADAccountName
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetADAccountName()`

UnsetADAccountName ensures that no value is present for ADAccountName, not even an explicit nil
### GetVMSid

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetVMSid() string`

GetVMSid returns the VMSid field if non-nil, zero value otherwise.

### GetVMSidOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetVMSidOk() (*string, bool)`

GetVMSidOk returns a tuple with the VMSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMSid

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetVMSid(v string)`

SetVMSid sets VMSid field to given value.

### HasVMSid

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasVMSid() bool`

HasVMSid returns a boolean if a field has been set.

### SetVMSidNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetVMSidNil(b bool)`

 SetVMSidNil sets the value for VMSid to be an explicit nil

### UnsetVMSid
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetVMSid()`

UnsetVMSid ensures that no value is present for VMSid, not even an explicit nil
### GetAssignedImage

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetAssignedImage() string`

GetAssignedImage returns the AssignedImage field if non-nil, zero value otherwise.

### GetAssignedImageOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetAssignedImageOk() (*string, bool)`

GetAssignedImageOk returns a tuple with the AssignedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedImage

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetAssignedImage(v string)`

SetAssignedImage sets AssignedImage field to given value.

### HasAssignedImage

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasAssignedImage() bool`

HasAssignedImage returns a boolean if a field has been set.

### SetAssignedImageNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetAssignedImageNil(b bool)`

 SetAssignedImageNil sets the value for AssignedImage to be an explicit nil

### UnsetAssignedImage
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetAssignedImage()`

UnsetAssignedImage ensures that no value is present for AssignedImage, not even an explicit nil
### GetBootedImage

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetBootedImage() string`

GetBootedImage returns the BootedImage field if non-nil, zero value otherwise.

### GetBootedImageOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetBootedImageOk() (*string, bool)`

GetBootedImageOk returns a tuple with the BootedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootedImage

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetBootedImage(v string)`

SetBootedImage sets BootedImage field to given value.

### HasBootedImage

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasBootedImage() bool`

HasBootedImage returns a boolean if a field has been set.

### SetBootedImageNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetBootedImageNil(b bool)`

 SetBootedImageNil sets the value for BootedImage to be an explicit nil

### UnsetBootedImage
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetBootedImage()`

UnsetBootedImage ensures that no value is present for BootedImage, not even an explicit nil
### GetCustomVmData

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetCustomVmData() ProvisionedVirtualMachineCustomVmDataResponseModel`

GetCustomVmData returns the CustomVmData field if non-nil, zero value otherwise.

### GetCustomVmDataOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetCustomVmDataOk() (*ProvisionedVirtualMachineCustomVmDataResponseModel, bool)`

GetCustomVmDataOk returns a tuple with the CustomVmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVmData

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetCustomVmData(v ProvisionedVirtualMachineCustomVmDataResponseModel)`

SetCustomVmData sets CustomVmData field to given value.

### HasCustomVmData

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasCustomVmData() bool`

HasCustomVmData returns a boolean if a field has been set.

### GetCpuCount

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetCreationDate

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### SetCreationDateNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetCreationDateNil(b bool)`

 SetCreationDateNil sets the value for CreationDate to be an explicit nil

### UnsetCreationDate
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetCreationDate()`

UnsetCreationDate ensures that no value is present for CreationDate, not even an explicit nil
### GetDomain

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEnableIntegrityMonitoring

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetEnableIntegrityMonitoring() bool`

GetEnableIntegrityMonitoring returns the EnableIntegrityMonitoring field if non-nil, zero value otherwise.

### GetEnableIntegrityMonitoringOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetEnableIntegrityMonitoringOk() (*bool, bool)`

GetEnableIntegrityMonitoringOk returns a tuple with the EnableIntegrityMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntegrityMonitoring

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetEnableIntegrityMonitoring(v bool)`

SetEnableIntegrityMonitoring sets EnableIntegrityMonitoring field to given value.

### HasEnableIntegrityMonitoring

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasEnableIntegrityMonitoring() bool`

HasEnableIntegrityMonitoring returns a boolean if a field has been set.

### SetEnableIntegrityMonitoringNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetEnableIntegrityMonitoringNil(b bool)`

 SetEnableIntegrityMonitoringNil sets the value for EnableIntegrityMonitoring to be an explicit nil

### UnsetEnableIntegrityMonitoring
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetEnableIntegrityMonitoring()`

UnsetEnableIntegrityMonitoring ensures that no value is present for EnableIntegrityMonitoring, not even an explicit nil
### GetEnableSecureBoot

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetEnableSecureBoot() bool`

GetEnableSecureBoot returns the EnableSecureBoot field if non-nil, zero value otherwise.

### GetEnableSecureBootOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetEnableSecureBootOk() (*bool, bool)`

GetEnableSecureBootOk returns a tuple with the EnableSecureBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSecureBoot

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetEnableSecureBoot(v bool)`

SetEnableSecureBoot sets EnableSecureBoot field to given value.

### HasEnableSecureBoot

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasEnableSecureBoot() bool`

HasEnableSecureBoot returns a boolean if a field has been set.

### SetEnableSecureBootNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetEnableSecureBootNil(b bool)`

 SetEnableSecureBootNil sets the value for EnableSecureBoot to be an explicit nil

### UnsetEnableSecureBoot
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetEnableSecureBoot()`

UnsetEnableSecureBoot ensures that no value is present for EnableSecureBoot, not even an explicit nil
### GetEnableVTPM

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetEnableVTPM() bool`

GetEnableVTPM returns the EnableVTPM field if non-nil, zero value otherwise.

### GetEnableVTPMOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetEnableVTPMOk() (*bool, bool)`

GetEnableVTPMOk returns a tuple with the EnableVTPM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVTPM

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetEnableVTPM(v bool)`

SetEnableVTPM sets EnableVTPM field to given value.

### HasEnableVTPM

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasEnableVTPM() bool`

HasEnableVTPM returns a boolean if a field has been set.

### SetEnableVTPMNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetEnableVTPMNil(b bool)`

 SetEnableVTPMNil sets the value for EnableVTPM to be an explicit nil

### UnsetEnableVTPM
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetEnableVTPM()`

UnsetEnableVTPM ensures that no value is present for EnableVTPM, not even an explicit nil
### GetEncryptionAtHost

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetEncryptionAtHost() bool`

GetEncryptionAtHost returns the EncryptionAtHost field if non-nil, zero value otherwise.

### GetEncryptionAtHostOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetEncryptionAtHostOk() (*bool, bool)`

GetEncryptionAtHostOk returns a tuple with the EncryptionAtHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtHost

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetEncryptionAtHost(v bool)`

SetEncryptionAtHost sets EncryptionAtHost field to given value.

### HasEncryptionAtHost

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasEncryptionAtHost() bool`

HasEncryptionAtHost returns a boolean if a field has been set.

### SetEncryptionAtHostNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetEncryptionAtHostNil(b bool)`

 SetEncryptionAtHostNil sets the value for EncryptionAtHost to be an explicit nil

### UnsetEncryptionAtHost
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetEncryptionAtHost()`

UnsetEncryptionAtHost ensures that no value is present for EncryptionAtHost, not even an explicit nil
### GetErrors

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetPluginId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### SetPluginIdNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetPluginIdNil(b bool)`

 SetPluginIdNil sets the value for PluginId to be an explicit nil

### UnsetPluginId
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetPluginId()`

UnsetPluginId ensures that no value is present for PluginId, not even an explicit nil
### GetImageOutOfDate

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetImageOutOfDate() bool`

GetImageOutOfDate returns the ImageOutOfDate field if non-nil, zero value otherwise.

### GetImageOutOfDateOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetImageOutOfDateOk() (*bool, bool)`

GetImageOutOfDateOk returns a tuple with the ImageOutOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOutOfDate

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetImageOutOfDate(v bool)`

SetImageOutOfDate sets ImageOutOfDate field to given value.

### HasImageOutOfDate

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasImageOutOfDate() bool`

HasImageOutOfDate returns a boolean if a field has been set.

### GetMemoryMB

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetHostingUnitUid

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetHostingUnitUid() string`

GetHostingUnitUid returns the HostingUnitUid field if non-nil, zero value otherwise.

### GetHostingUnitUidOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetHostingUnitUidOk() (*string, bool)`

GetHostingUnitUidOk returns a tuple with the HostingUnitUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingUnitUid

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetHostingUnitUid(v string)`

SetHostingUnitUid sets HostingUnitUid field to given value.

### HasHostingUnitUid

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasHostingUnitUid() bool`

HasHostingUnitUid returns a boolean if a field has been set.

### GetHypervisorConnectionUid

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetHypervisorConnectionUid() string`

GetHypervisorConnectionUid returns the HypervisorConnectionUid field if non-nil, zero value otherwise.

### GetHypervisorConnectionUidOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetHypervisorConnectionUidOk() (*string, bool)`

GetHypervisorConnectionUidOk returns a tuple with the HypervisorConnectionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnectionUid

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetHypervisorConnectionUid(v string)`

SetHypervisorConnectionUid sets HypervisorConnectionUid field to given value.

### HasHypervisorConnectionUid

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasHypervisorConnectionUid() bool`

HasHypervisorConnectionUid returns a boolean if a field has been set.

### GetIdentityDiskId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetIdentityDiskId() string`

GetIdentityDiskId returns the IdentityDiskId field if non-nil, zero value otherwise.

### GetIdentityDiskIdOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetIdentityDiskIdOk() (*string, bool)`

GetIdentityDiskIdOk returns a tuple with the IdentityDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDiskId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetIdentityDiskId(v string)`

SetIdentityDiskId sets IdentityDiskId field to given value.

### HasIdentityDiskId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasIdentityDiskId() bool`

HasIdentityDiskId returns a boolean if a field has been set.

### SetIdentityDiskIdNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetIdentityDiskIdNil(b bool)`

 SetIdentityDiskIdNil sets the value for IdentityDiskId to be an explicit nil

### UnsetIdentityDiskId
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetIdentityDiskId()`

UnsetIdentityDiskId ensures that no value is present for IdentityDiskId, not even an explicit nil
### GetLastBootTime

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetLastBootTime() string`

GetLastBootTime returns the LastBootTime field if non-nil, zero value otherwise.

### GetLastBootTimeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetLastBootTimeOk() (*string, bool)`

GetLastBootTimeOk returns a tuple with the LastBootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBootTime

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetLastBootTime(v string)`

SetLastBootTime sets LastBootTime field to given value.

### HasLastBootTime

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasLastBootTime() bool`

HasLastBootTime returns a boolean if a field has been set.

### SetLastBootTimeNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetLastBootTimeNil(b bool)`

 SetLastBootTimeNil sets the value for LastBootTime to be an explicit nil

### UnsetLastBootTime
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetLastBootTime()`

UnsetLastBootTime ensures that no value is present for LastBootTime, not even an explicit nil
### GetProvisioningSchemeType

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeType() string`

GetProvisioningSchemeType returns the ProvisioningSchemeType field if non-nil, zero value otherwise.

### GetProvisioningSchemeTypeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeTypeOk() (*string, bool)`

GetProvisioningSchemeTypeOk returns a tuple with the ProvisioningSchemeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeType

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvisioningSchemeType(v string)`

SetProvisioningSchemeType sets ProvisioningSchemeType field to given value.

### HasProvisioningSchemeType

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasProvisioningSchemeType() bool`

HasProvisioningSchemeType returns a boolean if a field has been set.

### SetProvisioningSchemeTypeNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvisioningSchemeTypeNil(b bool)`

 SetProvisioningSchemeTypeNil sets the value for ProvisioningSchemeType to be an explicit nil

### UnsetProvisioningSchemeType
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetProvisioningSchemeType()`

UnsetProvisioningSchemeType ensures that no value is present for ProvisioningSchemeType, not even an explicit nil
### GetProvisioningSchemeUpdateRequested

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeUpdateRequested() string`

GetProvisioningSchemeUpdateRequested returns the ProvisioningSchemeUpdateRequested field if non-nil, zero value otherwise.

### GetProvisioningSchemeUpdateRequestedOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeUpdateRequestedOk() (*string, bool)`

GetProvisioningSchemeUpdateRequestedOk returns a tuple with the ProvisioningSchemeUpdateRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeUpdateRequested

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvisioningSchemeUpdateRequested(v string)`

SetProvisioningSchemeUpdateRequested sets ProvisioningSchemeUpdateRequested field to given value.

### HasProvisioningSchemeUpdateRequested

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasProvisioningSchemeUpdateRequested() bool`

HasProvisioningSchemeUpdateRequested returns a boolean if a field has been set.

### SetProvisioningSchemeUpdateRequestedNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvisioningSchemeUpdateRequestedNil(b bool)`

 SetProvisioningSchemeUpdateRequestedNil sets the value for ProvisioningSchemeUpdateRequested to be an explicit nil

### UnsetProvisioningSchemeUpdateRequested
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetProvisioningSchemeUpdateRequested()`

UnsetProvisioningSchemeUpdateRequested ensures that no value is present for ProvisioningSchemeUpdateRequested, not even an explicit nil
### GetProvisioningSchemeVersion

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeVersion() int32`

GetProvisioningSchemeVersion returns the ProvisioningSchemeVersion field if non-nil, zero value otherwise.

### GetProvisioningSchemeVersionOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeVersionOk() (*int32, bool)`

GetProvisioningSchemeVersionOk returns a tuple with the ProvisioningSchemeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeVersion

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvisioningSchemeVersion(v int32)`

SetProvisioningSchemeVersion sets ProvisioningSchemeVersion field to given value.

### HasProvisioningSchemeVersion

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasProvisioningSchemeVersion() bool`

HasProvisioningSchemeVersion returns a boolean if a field has been set.

### GetUseFullDiskCloneProvisioning

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetUseFullDiskCloneProvisioning() bool`

GetUseFullDiskCloneProvisioning returns the UseFullDiskCloneProvisioning field if non-nil, zero value otherwise.

### GetUseFullDiskCloneProvisioningOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetUseFullDiskCloneProvisioningOk() (*bool, bool)`

GetUseFullDiskCloneProvisioningOk returns a tuple with the UseFullDiskCloneProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFullDiskCloneProvisioning

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetUseFullDiskCloneProvisioning(v bool)`

SetUseFullDiskCloneProvisioning sets UseFullDiskCloneProvisioning field to given value.

### HasUseFullDiskCloneProvisioning

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasUseFullDiskCloneProvisioning() bool`

HasUseFullDiskCloneProvisioning returns a boolean if a field has been set.

### GetVMId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetVMId() string`

GetVMId returns the VMId field if non-nil, zero value otherwise.

### GetVMIdOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetVMIdOk() (*string, bool)`

GetVMIdOk returns a tuple with the VMId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetVMId(v string)`

SetVMId sets VMId field to given value.

### HasVMId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasVMId() bool`

HasVMId returns a boolean if a field has been set.

### SetVMIdNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetVMIdNil(b bool)`

 SetVMIdNil sets the value for VMId to be an explicit nil

### UnsetVMId
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetVMId()`

UnsetVMId ensures that no value is present for VMId, not even an explicit nil
### GetVMName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetVMName() string`

GetVMName returns the VMName field if non-nil, zero value otherwise.

### GetVMNameOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetVMNameOk() (*string, bool)`

GetVMNameOk returns a tuple with the VMName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetVMName(v string)`

SetVMName sets VMName field to given value.

### HasVMName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasVMName() bool`

HasVMName returns a boolean if a field has been set.

### SetVMNameNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetVMNameNil(b bool)`

 SetVMNameNil sets the value for VMName to be an explicit nil

### UnsetVMName
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetVMName()`

UnsetVMName ensures that no value is present for VMName, not even an explicit nil
### GetWarnings

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWarnings() []ProvisionedVirtualMachineDetailsWarningReponseModel`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWarningsOk() (*[]ProvisionedVirtualMachineDetailsWarningReponseModel, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWarnings(v []ProvisionedVirtualMachineDetailsWarningReponseModel)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetWindowsActivationStatus

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWindowsActivationStatus() string`

GetWindowsActivationStatus returns the WindowsActivationStatus field if non-nil, zero value otherwise.

### GetWindowsActivationStatusOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWindowsActivationStatusOk() (*string, bool)`

GetWindowsActivationStatusOk returns a tuple with the WindowsActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsActivationStatus

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWindowsActivationStatus(v string)`

SetWindowsActivationStatus sets WindowsActivationStatus field to given value.

### HasWindowsActivationStatus

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWindowsActivationStatus() bool`

HasWindowsActivationStatus returns a boolean if a field has been set.

### SetWindowsActivationStatusNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWindowsActivationStatusNil(b bool)`

 SetWindowsActivationStatusNil sets the value for WindowsActivationStatus to be an explicit nil

### UnsetWindowsActivationStatus
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetWindowsActivationStatus()`

UnsetWindowsActivationStatus ensures that no value is present for WindowsActivationStatus, not even an explicit nil
### GetWindowsActivationStatusErrorCode

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWindowsActivationStatusErrorCode() string`

GetWindowsActivationStatusErrorCode returns the WindowsActivationStatusErrorCode field if non-nil, zero value otherwise.

### GetWindowsActivationStatusErrorCodeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWindowsActivationStatusErrorCodeOk() (*string, bool)`

GetWindowsActivationStatusErrorCodeOk returns a tuple with the WindowsActivationStatusErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsActivationStatusErrorCode

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWindowsActivationStatusErrorCode(v string)`

SetWindowsActivationStatusErrorCode sets WindowsActivationStatusErrorCode field to given value.

### HasWindowsActivationStatusErrorCode

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWindowsActivationStatusErrorCode() bool`

HasWindowsActivationStatusErrorCode returns a boolean if a field has been set.

### SetWindowsActivationStatusErrorCodeNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWindowsActivationStatusErrorCodeNil(b bool)`

 SetWindowsActivationStatusErrorCodeNil sets the value for WindowsActivationStatusErrorCode to be an explicit nil

### UnsetWindowsActivationStatusErrorCode
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetWindowsActivationStatusErrorCode()`

UnsetWindowsActivationStatusErrorCode ensures that no value is present for WindowsActivationStatusErrorCode, not even an explicit nil
### GetWindowsActivationStatusVirtualMachineError

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWindowsActivationStatusVirtualMachineError() string`

GetWindowsActivationStatusVirtualMachineError returns the WindowsActivationStatusVirtualMachineError field if non-nil, zero value otherwise.

### GetWindowsActivationStatusVirtualMachineErrorOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWindowsActivationStatusVirtualMachineErrorOk() (*string, bool)`

GetWindowsActivationStatusVirtualMachineErrorOk returns a tuple with the WindowsActivationStatusVirtualMachineError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsActivationStatusVirtualMachineError

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWindowsActivationStatusVirtualMachineError(v string)`

SetWindowsActivationStatusVirtualMachineError sets WindowsActivationStatusVirtualMachineError field to given value.

### HasWindowsActivationStatusVirtualMachineError

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWindowsActivationStatusVirtualMachineError() bool`

HasWindowsActivationStatusVirtualMachineError returns a boolean if a field has been set.

### SetWindowsActivationStatusVirtualMachineErrorNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWindowsActivationStatusVirtualMachineErrorNil(b bool)`

 SetWindowsActivationStatusVirtualMachineErrorNil sets the value for WindowsActivationStatusVirtualMachineError to be an explicit nil

### UnsetWindowsActivationStatusVirtualMachineError
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetWindowsActivationStatusVirtualMachineError()`

UnsetWindowsActivationStatusVirtualMachineError ensures that no value is present for WindowsActivationStatusVirtualMachineError, not even an explicit nil
### GetWindowsActivationTypeProvisionedVirtualMachine

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWindowsActivationTypeProvisionedVirtualMachine() string`

GetWindowsActivationTypeProvisionedVirtualMachine returns the WindowsActivationTypeProvisionedVirtualMachine field if non-nil, zero value otherwise.

### GetWindowsActivationTypeProvisionedVirtualMachineOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWindowsActivationTypeProvisionedVirtualMachineOk() (*string, bool)`

GetWindowsActivationTypeProvisionedVirtualMachineOk returns a tuple with the WindowsActivationTypeProvisionedVirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsActivationTypeProvisionedVirtualMachine

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWindowsActivationTypeProvisionedVirtualMachine(v string)`

SetWindowsActivationTypeProvisionedVirtualMachine sets WindowsActivationTypeProvisionedVirtualMachine field to given value.

### HasWindowsActivationTypeProvisionedVirtualMachine

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWindowsActivationTypeProvisionedVirtualMachine() bool`

HasWindowsActivationTypeProvisionedVirtualMachine returns a boolean if a field has been set.

### SetWindowsActivationTypeProvisionedVirtualMachineNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWindowsActivationTypeProvisionedVirtualMachineNil(b bool)`

 SetWindowsActivationTypeProvisionedVirtualMachineNil sets the value for WindowsActivationTypeProvisionedVirtualMachine to be an explicit nil

### UnsetWindowsActivationTypeProvisionedVirtualMachine
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetWindowsActivationTypeProvisionedVirtualMachine()`

UnsetWindowsActivationTypeProvisionedVirtualMachine ensures that no value is present for WindowsActivationTypeProvisionedVirtualMachine, not even an explicit nil
### GetWriteBackCacheDiskSize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheDiskSize() int32`

GetWriteBackCacheDiskSize returns the WriteBackCacheDiskSize field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheDiskSizeOk() (*int32, bool)`

GetWriteBackCacheDiskSizeOk returns a tuple with the WriteBackCacheDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWriteBackCacheDiskSize(v int32)`

SetWriteBackCacheDiskSize sets WriteBackCacheDiskSize field to given value.

### HasWriteBackCacheDiskSize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWriteBackCacheDiskSize() bool`

HasWriteBackCacheDiskSize returns a boolean if a field has been set.

### GetWriteBackCacheDriveLetter

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheDriveLetter() string`

GetWriteBackCacheDriveLetter returns the WriteBackCacheDriveLetter field if non-nil, zero value otherwise.

### GetWriteBackCacheDriveLetterOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheDriveLetterOk() (*string, bool)`

GetWriteBackCacheDriveLetterOk returns a tuple with the WriteBackCacheDriveLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDriveLetter

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWriteBackCacheDriveLetter(v string)`

SetWriteBackCacheDriveLetter sets WriteBackCacheDriveLetter field to given value.

### HasWriteBackCacheDriveLetter

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWriteBackCacheDriveLetter() bool`

HasWriteBackCacheDriveLetter returns a boolean if a field has been set.

### GetWriteBackCacheMemorySize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheMemorySize() int32`

GetWriteBackCacheMemorySize returns the WriteBackCacheMemorySize field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheMemorySizeOk() (*int32, bool)`

GetWriteBackCacheMemorySizeOk returns a tuple with the WriteBackCacheMemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWriteBackCacheMemorySize(v int32)`

SetWriteBackCacheMemorySize sets WriteBackCacheMemorySize field to given value.

### HasWriteBackCacheMemorySize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWriteBackCacheMemorySize() bool`

HasWriteBackCacheMemorySize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


