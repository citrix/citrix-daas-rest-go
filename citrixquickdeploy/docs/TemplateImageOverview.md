# TemplateImageOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the template image | 
**Name** | **string** | Name of the Template Image / VM | 
**SessionSupport** | [**TemplateImageSessionSupport**](TemplateImageSessionSupport.md) | Type of sessions that are supported by OS | 
**State** | [**TemplateImageState**](TemplateImageState.md) | State of the template image | 
**SubState** | Pointer to [**NullableTemplateImageSubState**](TemplateImageSubState.md) | Sub State of template image | [optional] 
**OsPlatform** | Pointer to [**NullableSupportedOperatingSystemType**](SupportedOperatingSystemType.md) | Type of operating system that will be imported | [optional] 
**OsName** | Pointer to **NullableString** | Shows name of image OS | [optional] 
**OsVersion** | Pointer to **NullableString** | Shows version of image OS | [optional] 
**LinuxDomainSupport** | Pointer to **NullableString** | For linux customers it shows domain configuration set in mcs.conf file | [optional] 
**LinuxRdpSupport** | Pointer to **NullableBool** | For linux customers it shows if rdp package was installed | [optional] 
**McsioSupport** | Pointer to **NullableBool** | For windows customers it shows if mcsio driver was installed | [optional] 
**VusSupport** | Pointer to **NullableBool** | For windows customers it shows if VDA update service was installed | [optional] 
**SccmAgentSupport** | Pointer to **NullableBool** | For windows customers it shows if sccm agent was installed | [optional] 
**SccmAgentVersion** | Pointer to **NullableString** | For windows customers it shows sccm agent version | [optional] 
**PublicIp** | Pointer to **NullableString** | Shows public ip address for image | [optional] 
**PrivateIp** | Pointer to **NullableString** | Shows private ip address for image | [optional] 
**DomainName** | Pointer to **NullableString** | Shows if image is currently domain joined and to which domain | [optional] 
**IsServerOs** | Pointer to **bool** | Indicates if the image is using a server based OS | [optional] 
**VdaVersion** | Pointer to **NullableString** | Shows version of image Virtual Desktop Agent | [optional] 
**StatusMessageId** | Pointer to **NullableString** | Status message enum related to verifying and enumerating the image | [optional] 
**Status** | Pointer to **NullableString** | Status message related to verifying and enumerating the image | [optional] 
**ExtraInfo** | Pointer to **NullableString** | The string to displayed in UI for extra information | [optional] 
**Notes** | Pointer to **NullableString** | Customer notes about template image | [optional] 
**TransactionId** | Pointer to **NullableString** | ID of the transaction that the image was verified on | [optional] 
**SubscriptionId** | Pointer to **NullableString** | Id of the Subscription where the image is stored (BYOA) | [optional] 
**SubscriptionName** | **string** | Name of the Subscription that catalog VMs will be deployed to | 
**ResourceGroup** | Pointer to **NullableString** | Name of the Azure Resource Group where the image is stored | [optional] 
**StorageAccount** | Pointer to **NullableString** | Name of the Storage Account where the image is stored | [optional] 
**Region** | **string** | Azure region where VMs are deployed for this catalog | 
**BuilderDomainName** | Pointer to **NullableString** | Name of the Domain the Image Builder will join | [optional] 
**BuilderConnectionId** | Pointer to **NullableString** | ID of the on-prem connection associated with the builder image | [optional] 
**BuilderVmName** | Pointer to **NullableString** | Name of VM that is being used by the builder image | [optional] 
**BuilderVmType** | Pointer to **NullableString** | Type of VM that is being used by the builder image | [optional] 
**BuilderVmDiskSize** | Pointer to **NullableString** | Size of disk of VM that is being used by the builder image | [optional] 
**BuilderAllowedIPs** | Pointer to **[]string** | Ip Addresses allowed to RDP | [optional] 
**HyperVGen** | Pointer to **NullableString** | The HyperVGeneration that should be set to either V1 or V2 | [optional] 
**VtpmEnabled** | Pointer to **bool** | Is TrustedLaunch VTPM supported in V2 gen | [optional] 
**SecureBootEnabled** | Pointer to **bool** | Is TrustedLaunch SecureBoot supported in V2 gen | [optional] 
**CitrixPrepared** | **bool** | Whether the image was prepared by Citrix, or provided by the customer | 
**CspCustomer** | Pointer to **NullableString** | Indicates that partner-tenant relationship exists if not null | [optional] 
**IsCmekEnabled** | Pointer to **bool** | Indicates if the image is using customer managed encryption keys | [optional] 
**CmekId** | Pointer to **NullableString** | Indicates if customer managed encryption key Id | [optional] 
**CmekName** | Pointer to **NullableString** | Indicates if customer managed encryption key name | [optional] 
**IsDeprecated** | Pointer to **bool** | Indicates if the image is deprecated and should not be used in new catalogs | [optional] 
**IsByoaImage** | Pointer to **NullableBool** | Indicates if the image is built or imported with BYOA | [optional] 
**LinkedCatalogs** | Pointer to **int32** | Number of catalogs that are using this image | [optional] 
**LinkedCatalogsNames** | Pointer to **[]string** | Names of catalogs that are using this image | [optional] 
**CreatedDate** | Pointer to **NullableTime** | Created from datastore. | [optional] 
**FinalizedDate** | Pointer to **NullableTime** | Image builder finalized date. | [optional] 
**Path** | Pointer to **NullableString** | Customer image path in Azure | [optional] 
**SbSessionVdaVersion** | Pointer to **NullableString** | Shows version of sbsession | [optional] 
**IsSecureBrowserImage** | Pointer to **NullableBool** | Indicates if the image is for Secure Browser | [optional] 
**DiskSizeInGB** | Pointer to **NullableInt32** | The size of the disk in GB | [optional] 
**StartedAt** | Pointer to **NullableTime** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **NullableInt32** | Estimated total time for the job to finish | [optional] 

## Methods

### NewTemplateImageOverview

`func NewTemplateImageOverview(id string, name string, sessionSupport TemplateImageSessionSupport, state TemplateImageState, subscriptionName string, region string, citrixPrepared bool, ) *TemplateImageOverview`

NewTemplateImageOverview instantiates a new TemplateImageOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateImageOverviewWithDefaults

`func NewTemplateImageOverviewWithDefaults() *TemplateImageOverview`

NewTemplateImageOverviewWithDefaults instantiates a new TemplateImageOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateImageOverview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateImageOverview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateImageOverview) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TemplateImageOverview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateImageOverview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateImageOverview) SetName(v string)`

SetName sets Name field to given value.


### GetSessionSupport

`func (o *TemplateImageOverview) GetSessionSupport() TemplateImageSessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *TemplateImageOverview) GetSessionSupportOk() (*TemplateImageSessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *TemplateImageOverview) SetSessionSupport(v TemplateImageSessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetState

`func (o *TemplateImageOverview) GetState() TemplateImageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TemplateImageOverview) GetStateOk() (*TemplateImageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TemplateImageOverview) SetState(v TemplateImageState)`

SetState sets State field to given value.


### GetSubState

`func (o *TemplateImageOverview) GetSubState() TemplateImageSubState`

GetSubState returns the SubState field if non-nil, zero value otherwise.

### GetSubStateOk

`func (o *TemplateImageOverview) GetSubStateOk() (*TemplateImageSubState, bool)`

GetSubStateOk returns a tuple with the SubState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubState

`func (o *TemplateImageOverview) SetSubState(v TemplateImageSubState)`

SetSubState sets SubState field to given value.

### HasSubState

`func (o *TemplateImageOverview) HasSubState() bool`

HasSubState returns a boolean if a field has been set.

### SetSubStateNil

`func (o *TemplateImageOverview) SetSubStateNil(b bool)`

 SetSubStateNil sets the value for SubState to be an explicit nil

### UnsetSubState
`func (o *TemplateImageOverview) UnsetSubState()`

UnsetSubState ensures that no value is present for SubState, not even an explicit nil
### GetOsPlatform

`func (o *TemplateImageOverview) GetOsPlatform() SupportedOperatingSystemType`

GetOsPlatform returns the OsPlatform field if non-nil, zero value otherwise.

### GetOsPlatformOk

`func (o *TemplateImageOverview) GetOsPlatformOk() (*SupportedOperatingSystemType, bool)`

GetOsPlatformOk returns a tuple with the OsPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPlatform

`func (o *TemplateImageOverview) SetOsPlatform(v SupportedOperatingSystemType)`

SetOsPlatform sets OsPlatform field to given value.

### HasOsPlatform

`func (o *TemplateImageOverview) HasOsPlatform() bool`

HasOsPlatform returns a boolean if a field has been set.

### SetOsPlatformNil

`func (o *TemplateImageOverview) SetOsPlatformNil(b bool)`

 SetOsPlatformNil sets the value for OsPlatform to be an explicit nil

### UnsetOsPlatform
`func (o *TemplateImageOverview) UnsetOsPlatform()`

UnsetOsPlatform ensures that no value is present for OsPlatform, not even an explicit nil
### GetOsName

`func (o *TemplateImageOverview) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *TemplateImageOverview) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *TemplateImageOverview) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *TemplateImageOverview) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### SetOsNameNil

`func (o *TemplateImageOverview) SetOsNameNil(b bool)`

 SetOsNameNil sets the value for OsName to be an explicit nil

### UnsetOsName
`func (o *TemplateImageOverview) UnsetOsName()`

UnsetOsName ensures that no value is present for OsName, not even an explicit nil
### GetOsVersion

`func (o *TemplateImageOverview) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *TemplateImageOverview) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *TemplateImageOverview) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *TemplateImageOverview) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *TemplateImageOverview) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *TemplateImageOverview) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetLinuxDomainSupport

`func (o *TemplateImageOverview) GetLinuxDomainSupport() string`

GetLinuxDomainSupport returns the LinuxDomainSupport field if non-nil, zero value otherwise.

### GetLinuxDomainSupportOk

`func (o *TemplateImageOverview) GetLinuxDomainSupportOk() (*string, bool)`

GetLinuxDomainSupportOk returns a tuple with the LinuxDomainSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxDomainSupport

`func (o *TemplateImageOverview) SetLinuxDomainSupport(v string)`

SetLinuxDomainSupport sets LinuxDomainSupport field to given value.

### HasLinuxDomainSupport

`func (o *TemplateImageOverview) HasLinuxDomainSupport() bool`

HasLinuxDomainSupport returns a boolean if a field has been set.

### SetLinuxDomainSupportNil

`func (o *TemplateImageOverview) SetLinuxDomainSupportNil(b bool)`

 SetLinuxDomainSupportNil sets the value for LinuxDomainSupport to be an explicit nil

### UnsetLinuxDomainSupport
`func (o *TemplateImageOverview) UnsetLinuxDomainSupport()`

UnsetLinuxDomainSupport ensures that no value is present for LinuxDomainSupport, not even an explicit nil
### GetLinuxRdpSupport

`func (o *TemplateImageOverview) GetLinuxRdpSupport() bool`

GetLinuxRdpSupport returns the LinuxRdpSupport field if non-nil, zero value otherwise.

### GetLinuxRdpSupportOk

`func (o *TemplateImageOverview) GetLinuxRdpSupportOk() (*bool, bool)`

GetLinuxRdpSupportOk returns a tuple with the LinuxRdpSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxRdpSupport

`func (o *TemplateImageOverview) SetLinuxRdpSupport(v bool)`

SetLinuxRdpSupport sets LinuxRdpSupport field to given value.

### HasLinuxRdpSupport

`func (o *TemplateImageOverview) HasLinuxRdpSupport() bool`

HasLinuxRdpSupport returns a boolean if a field has been set.

### SetLinuxRdpSupportNil

`func (o *TemplateImageOverview) SetLinuxRdpSupportNil(b bool)`

 SetLinuxRdpSupportNil sets the value for LinuxRdpSupport to be an explicit nil

### UnsetLinuxRdpSupport
`func (o *TemplateImageOverview) UnsetLinuxRdpSupport()`

UnsetLinuxRdpSupport ensures that no value is present for LinuxRdpSupport, not even an explicit nil
### GetMcsioSupport

`func (o *TemplateImageOverview) GetMcsioSupport() bool`

GetMcsioSupport returns the McsioSupport field if non-nil, zero value otherwise.

### GetMcsioSupportOk

`func (o *TemplateImageOverview) GetMcsioSupportOk() (*bool, bool)`

GetMcsioSupportOk returns a tuple with the McsioSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsioSupport

`func (o *TemplateImageOverview) SetMcsioSupport(v bool)`

SetMcsioSupport sets McsioSupport field to given value.

### HasMcsioSupport

`func (o *TemplateImageOverview) HasMcsioSupport() bool`

HasMcsioSupport returns a boolean if a field has been set.

### SetMcsioSupportNil

`func (o *TemplateImageOverview) SetMcsioSupportNil(b bool)`

 SetMcsioSupportNil sets the value for McsioSupport to be an explicit nil

### UnsetMcsioSupport
`func (o *TemplateImageOverview) UnsetMcsioSupport()`

UnsetMcsioSupport ensures that no value is present for McsioSupport, not even an explicit nil
### GetVusSupport

`func (o *TemplateImageOverview) GetVusSupport() bool`

GetVusSupport returns the VusSupport field if non-nil, zero value otherwise.

### GetVusSupportOk

`func (o *TemplateImageOverview) GetVusSupportOk() (*bool, bool)`

GetVusSupportOk returns a tuple with the VusSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVusSupport

`func (o *TemplateImageOverview) SetVusSupport(v bool)`

SetVusSupport sets VusSupport field to given value.

### HasVusSupport

`func (o *TemplateImageOverview) HasVusSupport() bool`

HasVusSupport returns a boolean if a field has been set.

### SetVusSupportNil

`func (o *TemplateImageOverview) SetVusSupportNil(b bool)`

 SetVusSupportNil sets the value for VusSupport to be an explicit nil

### UnsetVusSupport
`func (o *TemplateImageOverview) UnsetVusSupport()`

UnsetVusSupport ensures that no value is present for VusSupport, not even an explicit nil
### GetSccmAgentSupport

`func (o *TemplateImageOverview) GetSccmAgentSupport() bool`

GetSccmAgentSupport returns the SccmAgentSupport field if non-nil, zero value otherwise.

### GetSccmAgentSupportOk

`func (o *TemplateImageOverview) GetSccmAgentSupportOk() (*bool, bool)`

GetSccmAgentSupportOk returns a tuple with the SccmAgentSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSccmAgentSupport

`func (o *TemplateImageOverview) SetSccmAgentSupport(v bool)`

SetSccmAgentSupport sets SccmAgentSupport field to given value.

### HasSccmAgentSupport

`func (o *TemplateImageOverview) HasSccmAgentSupport() bool`

HasSccmAgentSupport returns a boolean if a field has been set.

### SetSccmAgentSupportNil

`func (o *TemplateImageOverview) SetSccmAgentSupportNil(b bool)`

 SetSccmAgentSupportNil sets the value for SccmAgentSupport to be an explicit nil

### UnsetSccmAgentSupport
`func (o *TemplateImageOverview) UnsetSccmAgentSupport()`

UnsetSccmAgentSupport ensures that no value is present for SccmAgentSupport, not even an explicit nil
### GetSccmAgentVersion

`func (o *TemplateImageOverview) GetSccmAgentVersion() string`

GetSccmAgentVersion returns the SccmAgentVersion field if non-nil, zero value otherwise.

### GetSccmAgentVersionOk

`func (o *TemplateImageOverview) GetSccmAgentVersionOk() (*string, bool)`

GetSccmAgentVersionOk returns a tuple with the SccmAgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSccmAgentVersion

`func (o *TemplateImageOverview) SetSccmAgentVersion(v string)`

SetSccmAgentVersion sets SccmAgentVersion field to given value.

### HasSccmAgentVersion

`func (o *TemplateImageOverview) HasSccmAgentVersion() bool`

HasSccmAgentVersion returns a boolean if a field has been set.

### SetSccmAgentVersionNil

`func (o *TemplateImageOverview) SetSccmAgentVersionNil(b bool)`

 SetSccmAgentVersionNil sets the value for SccmAgentVersion to be an explicit nil

### UnsetSccmAgentVersion
`func (o *TemplateImageOverview) UnsetSccmAgentVersion()`

UnsetSccmAgentVersion ensures that no value is present for SccmAgentVersion, not even an explicit nil
### GetPublicIp

`func (o *TemplateImageOverview) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *TemplateImageOverview) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *TemplateImageOverview) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *TemplateImageOverview) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIpNil

`func (o *TemplateImageOverview) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *TemplateImageOverview) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetPrivateIp

`func (o *TemplateImageOverview) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *TemplateImageOverview) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *TemplateImageOverview) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *TemplateImageOverview) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *TemplateImageOverview) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *TemplateImageOverview) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetDomainName

`func (o *TemplateImageOverview) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *TemplateImageOverview) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *TemplateImageOverview) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *TemplateImageOverview) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *TemplateImageOverview) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *TemplateImageOverview) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetIsServerOs

`func (o *TemplateImageOverview) GetIsServerOs() bool`

GetIsServerOs returns the IsServerOs field if non-nil, zero value otherwise.

### GetIsServerOsOk

`func (o *TemplateImageOverview) GetIsServerOsOk() (*bool, bool)`

GetIsServerOsOk returns a tuple with the IsServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServerOs

`func (o *TemplateImageOverview) SetIsServerOs(v bool)`

SetIsServerOs sets IsServerOs field to given value.

### HasIsServerOs

`func (o *TemplateImageOverview) HasIsServerOs() bool`

HasIsServerOs returns a boolean if a field has been set.

### GetVdaVersion

`func (o *TemplateImageOverview) GetVdaVersion() string`

GetVdaVersion returns the VdaVersion field if non-nil, zero value otherwise.

### GetVdaVersionOk

`func (o *TemplateImageOverview) GetVdaVersionOk() (*string, bool)`

GetVdaVersionOk returns a tuple with the VdaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaVersion

`func (o *TemplateImageOverview) SetVdaVersion(v string)`

SetVdaVersion sets VdaVersion field to given value.

### HasVdaVersion

`func (o *TemplateImageOverview) HasVdaVersion() bool`

HasVdaVersion returns a boolean if a field has been set.

### SetVdaVersionNil

`func (o *TemplateImageOverview) SetVdaVersionNil(b bool)`

 SetVdaVersionNil sets the value for VdaVersion to be an explicit nil

### UnsetVdaVersion
`func (o *TemplateImageOverview) UnsetVdaVersion()`

UnsetVdaVersion ensures that no value is present for VdaVersion, not even an explicit nil
### GetStatusMessageId

`func (o *TemplateImageOverview) GetStatusMessageId() string`

GetStatusMessageId returns the StatusMessageId field if non-nil, zero value otherwise.

### GetStatusMessageIdOk

`func (o *TemplateImageOverview) GetStatusMessageIdOk() (*string, bool)`

GetStatusMessageIdOk returns a tuple with the StatusMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessageId

`func (o *TemplateImageOverview) SetStatusMessageId(v string)`

SetStatusMessageId sets StatusMessageId field to given value.

### HasStatusMessageId

`func (o *TemplateImageOverview) HasStatusMessageId() bool`

HasStatusMessageId returns a boolean if a field has been set.

### SetStatusMessageIdNil

`func (o *TemplateImageOverview) SetStatusMessageIdNil(b bool)`

 SetStatusMessageIdNil sets the value for StatusMessageId to be an explicit nil

### UnsetStatusMessageId
`func (o *TemplateImageOverview) UnsetStatusMessageId()`

UnsetStatusMessageId ensures that no value is present for StatusMessageId, not even an explicit nil
### GetStatus

`func (o *TemplateImageOverview) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateImageOverview) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateImageOverview) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TemplateImageOverview) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TemplateImageOverview) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TemplateImageOverview) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetExtraInfo

`func (o *TemplateImageOverview) GetExtraInfo() string`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *TemplateImageOverview) GetExtraInfoOk() (*string, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *TemplateImageOverview) SetExtraInfo(v string)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *TemplateImageOverview) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### SetExtraInfoNil

`func (o *TemplateImageOverview) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *TemplateImageOverview) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil
### GetNotes

`func (o *TemplateImageOverview) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TemplateImageOverview) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TemplateImageOverview) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TemplateImageOverview) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TemplateImageOverview) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TemplateImageOverview) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetTransactionId

`func (o *TemplateImageOverview) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TemplateImageOverview) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TemplateImageOverview) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TemplateImageOverview) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *TemplateImageOverview) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *TemplateImageOverview) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetSubscriptionId

`func (o *TemplateImageOverview) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *TemplateImageOverview) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *TemplateImageOverview) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *TemplateImageOverview) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *TemplateImageOverview) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *TemplateImageOverview) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetSubscriptionName

`func (o *TemplateImageOverview) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *TemplateImageOverview) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *TemplateImageOverview) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.


### GetResourceGroup

`func (o *TemplateImageOverview) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *TemplateImageOverview) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *TemplateImageOverview) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *TemplateImageOverview) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### SetResourceGroupNil

`func (o *TemplateImageOverview) SetResourceGroupNil(b bool)`

 SetResourceGroupNil sets the value for ResourceGroup to be an explicit nil

### UnsetResourceGroup
`func (o *TemplateImageOverview) UnsetResourceGroup()`

UnsetResourceGroup ensures that no value is present for ResourceGroup, not even an explicit nil
### GetStorageAccount

`func (o *TemplateImageOverview) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *TemplateImageOverview) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *TemplateImageOverview) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *TemplateImageOverview) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### SetStorageAccountNil

`func (o *TemplateImageOverview) SetStorageAccountNil(b bool)`

 SetStorageAccountNil sets the value for StorageAccount to be an explicit nil

### UnsetStorageAccount
`func (o *TemplateImageOverview) UnsetStorageAccount()`

UnsetStorageAccount ensures that no value is present for StorageAccount, not even an explicit nil
### GetRegion

`func (o *TemplateImageOverview) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TemplateImageOverview) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TemplateImageOverview) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetBuilderDomainName

`func (o *TemplateImageOverview) GetBuilderDomainName() string`

GetBuilderDomainName returns the BuilderDomainName field if non-nil, zero value otherwise.

### GetBuilderDomainNameOk

`func (o *TemplateImageOverview) GetBuilderDomainNameOk() (*string, bool)`

GetBuilderDomainNameOk returns a tuple with the BuilderDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderDomainName

`func (o *TemplateImageOverview) SetBuilderDomainName(v string)`

SetBuilderDomainName sets BuilderDomainName field to given value.

### HasBuilderDomainName

`func (o *TemplateImageOverview) HasBuilderDomainName() bool`

HasBuilderDomainName returns a boolean if a field has been set.

### SetBuilderDomainNameNil

`func (o *TemplateImageOverview) SetBuilderDomainNameNil(b bool)`

 SetBuilderDomainNameNil sets the value for BuilderDomainName to be an explicit nil

### UnsetBuilderDomainName
`func (o *TemplateImageOverview) UnsetBuilderDomainName()`

UnsetBuilderDomainName ensures that no value is present for BuilderDomainName, not even an explicit nil
### GetBuilderConnectionId

`func (o *TemplateImageOverview) GetBuilderConnectionId() string`

GetBuilderConnectionId returns the BuilderConnectionId field if non-nil, zero value otherwise.

### GetBuilderConnectionIdOk

`func (o *TemplateImageOverview) GetBuilderConnectionIdOk() (*string, bool)`

GetBuilderConnectionIdOk returns a tuple with the BuilderConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderConnectionId

`func (o *TemplateImageOverview) SetBuilderConnectionId(v string)`

SetBuilderConnectionId sets BuilderConnectionId field to given value.

### HasBuilderConnectionId

`func (o *TemplateImageOverview) HasBuilderConnectionId() bool`

HasBuilderConnectionId returns a boolean if a field has been set.

### SetBuilderConnectionIdNil

`func (o *TemplateImageOverview) SetBuilderConnectionIdNil(b bool)`

 SetBuilderConnectionIdNil sets the value for BuilderConnectionId to be an explicit nil

### UnsetBuilderConnectionId
`func (o *TemplateImageOverview) UnsetBuilderConnectionId()`

UnsetBuilderConnectionId ensures that no value is present for BuilderConnectionId, not even an explicit nil
### GetBuilderVmName

`func (o *TemplateImageOverview) GetBuilderVmName() string`

GetBuilderVmName returns the BuilderVmName field if non-nil, zero value otherwise.

### GetBuilderVmNameOk

`func (o *TemplateImageOverview) GetBuilderVmNameOk() (*string, bool)`

GetBuilderVmNameOk returns a tuple with the BuilderVmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderVmName

`func (o *TemplateImageOverview) SetBuilderVmName(v string)`

SetBuilderVmName sets BuilderVmName field to given value.

### HasBuilderVmName

`func (o *TemplateImageOverview) HasBuilderVmName() bool`

HasBuilderVmName returns a boolean if a field has been set.

### SetBuilderVmNameNil

`func (o *TemplateImageOverview) SetBuilderVmNameNil(b bool)`

 SetBuilderVmNameNil sets the value for BuilderVmName to be an explicit nil

### UnsetBuilderVmName
`func (o *TemplateImageOverview) UnsetBuilderVmName()`

UnsetBuilderVmName ensures that no value is present for BuilderVmName, not even an explicit nil
### GetBuilderVmType

`func (o *TemplateImageOverview) GetBuilderVmType() string`

GetBuilderVmType returns the BuilderVmType field if non-nil, zero value otherwise.

### GetBuilderVmTypeOk

`func (o *TemplateImageOverview) GetBuilderVmTypeOk() (*string, bool)`

GetBuilderVmTypeOk returns a tuple with the BuilderVmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderVmType

`func (o *TemplateImageOverview) SetBuilderVmType(v string)`

SetBuilderVmType sets BuilderVmType field to given value.

### HasBuilderVmType

`func (o *TemplateImageOverview) HasBuilderVmType() bool`

HasBuilderVmType returns a boolean if a field has been set.

### SetBuilderVmTypeNil

`func (o *TemplateImageOverview) SetBuilderVmTypeNil(b bool)`

 SetBuilderVmTypeNil sets the value for BuilderVmType to be an explicit nil

### UnsetBuilderVmType
`func (o *TemplateImageOverview) UnsetBuilderVmType()`

UnsetBuilderVmType ensures that no value is present for BuilderVmType, not even an explicit nil
### GetBuilderVmDiskSize

`func (o *TemplateImageOverview) GetBuilderVmDiskSize() string`

GetBuilderVmDiskSize returns the BuilderVmDiskSize field if non-nil, zero value otherwise.

### GetBuilderVmDiskSizeOk

`func (o *TemplateImageOverview) GetBuilderVmDiskSizeOk() (*string, bool)`

GetBuilderVmDiskSizeOk returns a tuple with the BuilderVmDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderVmDiskSize

`func (o *TemplateImageOverview) SetBuilderVmDiskSize(v string)`

SetBuilderVmDiskSize sets BuilderVmDiskSize field to given value.

### HasBuilderVmDiskSize

`func (o *TemplateImageOverview) HasBuilderVmDiskSize() bool`

HasBuilderVmDiskSize returns a boolean if a field has been set.

### SetBuilderVmDiskSizeNil

`func (o *TemplateImageOverview) SetBuilderVmDiskSizeNil(b bool)`

 SetBuilderVmDiskSizeNil sets the value for BuilderVmDiskSize to be an explicit nil

### UnsetBuilderVmDiskSize
`func (o *TemplateImageOverview) UnsetBuilderVmDiskSize()`

UnsetBuilderVmDiskSize ensures that no value is present for BuilderVmDiskSize, not even an explicit nil
### GetBuilderAllowedIPs

`func (o *TemplateImageOverview) GetBuilderAllowedIPs() []string`

GetBuilderAllowedIPs returns the BuilderAllowedIPs field if non-nil, zero value otherwise.

### GetBuilderAllowedIPsOk

`func (o *TemplateImageOverview) GetBuilderAllowedIPsOk() (*[]string, bool)`

GetBuilderAllowedIPsOk returns a tuple with the BuilderAllowedIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderAllowedIPs

`func (o *TemplateImageOverview) SetBuilderAllowedIPs(v []string)`

SetBuilderAllowedIPs sets BuilderAllowedIPs field to given value.

### HasBuilderAllowedIPs

`func (o *TemplateImageOverview) HasBuilderAllowedIPs() bool`

HasBuilderAllowedIPs returns a boolean if a field has been set.

### SetBuilderAllowedIPsNil

`func (o *TemplateImageOverview) SetBuilderAllowedIPsNil(b bool)`

 SetBuilderAllowedIPsNil sets the value for BuilderAllowedIPs to be an explicit nil

### UnsetBuilderAllowedIPs
`func (o *TemplateImageOverview) UnsetBuilderAllowedIPs()`

UnsetBuilderAllowedIPs ensures that no value is present for BuilderAllowedIPs, not even an explicit nil
### GetHyperVGen

`func (o *TemplateImageOverview) GetHyperVGen() string`

GetHyperVGen returns the HyperVGen field if non-nil, zero value otherwise.

### GetHyperVGenOk

`func (o *TemplateImageOverview) GetHyperVGenOk() (*string, bool)`

GetHyperVGenOk returns a tuple with the HyperVGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperVGen

`func (o *TemplateImageOverview) SetHyperVGen(v string)`

SetHyperVGen sets HyperVGen field to given value.

### HasHyperVGen

`func (o *TemplateImageOverview) HasHyperVGen() bool`

HasHyperVGen returns a boolean if a field has been set.

### SetHyperVGenNil

`func (o *TemplateImageOverview) SetHyperVGenNil(b bool)`

 SetHyperVGenNil sets the value for HyperVGen to be an explicit nil

### UnsetHyperVGen
`func (o *TemplateImageOverview) UnsetHyperVGen()`

UnsetHyperVGen ensures that no value is present for HyperVGen, not even an explicit nil
### GetVtpmEnabled

`func (o *TemplateImageOverview) GetVtpmEnabled() bool`

GetVtpmEnabled returns the VtpmEnabled field if non-nil, zero value otherwise.

### GetVtpmEnabledOk

`func (o *TemplateImageOverview) GetVtpmEnabledOk() (*bool, bool)`

GetVtpmEnabledOk returns a tuple with the VtpmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpmEnabled

`func (o *TemplateImageOverview) SetVtpmEnabled(v bool)`

SetVtpmEnabled sets VtpmEnabled field to given value.

### HasVtpmEnabled

`func (o *TemplateImageOverview) HasVtpmEnabled() bool`

HasVtpmEnabled returns a boolean if a field has been set.

### GetSecureBootEnabled

`func (o *TemplateImageOverview) GetSecureBootEnabled() bool`

GetSecureBootEnabled returns the SecureBootEnabled field if non-nil, zero value otherwise.

### GetSecureBootEnabledOk

`func (o *TemplateImageOverview) GetSecureBootEnabledOk() (*bool, bool)`

GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBootEnabled

`func (o *TemplateImageOverview) SetSecureBootEnabled(v bool)`

SetSecureBootEnabled sets SecureBootEnabled field to given value.

### HasSecureBootEnabled

`func (o *TemplateImageOverview) HasSecureBootEnabled() bool`

HasSecureBootEnabled returns a boolean if a field has been set.

### GetCitrixPrepared

`func (o *TemplateImageOverview) GetCitrixPrepared() bool`

GetCitrixPrepared returns the CitrixPrepared field if non-nil, zero value otherwise.

### GetCitrixPreparedOk

`func (o *TemplateImageOverview) GetCitrixPreparedOk() (*bool, bool)`

GetCitrixPreparedOk returns a tuple with the CitrixPrepared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixPrepared

`func (o *TemplateImageOverview) SetCitrixPrepared(v bool)`

SetCitrixPrepared sets CitrixPrepared field to given value.


### GetCspCustomer

`func (o *TemplateImageOverview) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *TemplateImageOverview) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *TemplateImageOverview) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *TemplateImageOverview) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### SetCspCustomerNil

`func (o *TemplateImageOverview) SetCspCustomerNil(b bool)`

 SetCspCustomerNil sets the value for CspCustomer to be an explicit nil

### UnsetCspCustomer
`func (o *TemplateImageOverview) UnsetCspCustomer()`

UnsetCspCustomer ensures that no value is present for CspCustomer, not even an explicit nil
### GetIsCmekEnabled

`func (o *TemplateImageOverview) GetIsCmekEnabled() bool`

GetIsCmekEnabled returns the IsCmekEnabled field if non-nil, zero value otherwise.

### GetIsCmekEnabledOk

`func (o *TemplateImageOverview) GetIsCmekEnabledOk() (*bool, bool)`

GetIsCmekEnabledOk returns a tuple with the IsCmekEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCmekEnabled

`func (o *TemplateImageOverview) SetIsCmekEnabled(v bool)`

SetIsCmekEnabled sets IsCmekEnabled field to given value.

### HasIsCmekEnabled

`func (o *TemplateImageOverview) HasIsCmekEnabled() bool`

HasIsCmekEnabled returns a boolean if a field has been set.

### GetCmekId

`func (o *TemplateImageOverview) GetCmekId() string`

GetCmekId returns the CmekId field if non-nil, zero value otherwise.

### GetCmekIdOk

`func (o *TemplateImageOverview) GetCmekIdOk() (*string, bool)`

GetCmekIdOk returns a tuple with the CmekId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmekId

`func (o *TemplateImageOverview) SetCmekId(v string)`

SetCmekId sets CmekId field to given value.

### HasCmekId

`func (o *TemplateImageOverview) HasCmekId() bool`

HasCmekId returns a boolean if a field has been set.

### SetCmekIdNil

`func (o *TemplateImageOverview) SetCmekIdNil(b bool)`

 SetCmekIdNil sets the value for CmekId to be an explicit nil

### UnsetCmekId
`func (o *TemplateImageOverview) UnsetCmekId()`

UnsetCmekId ensures that no value is present for CmekId, not even an explicit nil
### GetCmekName

`func (o *TemplateImageOverview) GetCmekName() string`

GetCmekName returns the CmekName field if non-nil, zero value otherwise.

### GetCmekNameOk

`func (o *TemplateImageOverview) GetCmekNameOk() (*string, bool)`

GetCmekNameOk returns a tuple with the CmekName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmekName

`func (o *TemplateImageOverview) SetCmekName(v string)`

SetCmekName sets CmekName field to given value.

### HasCmekName

`func (o *TemplateImageOverview) HasCmekName() bool`

HasCmekName returns a boolean if a field has been set.

### SetCmekNameNil

`func (o *TemplateImageOverview) SetCmekNameNil(b bool)`

 SetCmekNameNil sets the value for CmekName to be an explicit nil

### UnsetCmekName
`func (o *TemplateImageOverview) UnsetCmekName()`

UnsetCmekName ensures that no value is present for CmekName, not even an explicit nil
### GetIsDeprecated

`func (o *TemplateImageOverview) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *TemplateImageOverview) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *TemplateImageOverview) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *TemplateImageOverview) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.

### GetIsByoaImage

`func (o *TemplateImageOverview) GetIsByoaImage() bool`

GetIsByoaImage returns the IsByoaImage field if non-nil, zero value otherwise.

### GetIsByoaImageOk

`func (o *TemplateImageOverview) GetIsByoaImageOk() (*bool, bool)`

GetIsByoaImageOk returns a tuple with the IsByoaImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByoaImage

`func (o *TemplateImageOverview) SetIsByoaImage(v bool)`

SetIsByoaImage sets IsByoaImage field to given value.

### HasIsByoaImage

`func (o *TemplateImageOverview) HasIsByoaImage() bool`

HasIsByoaImage returns a boolean if a field has been set.

### SetIsByoaImageNil

`func (o *TemplateImageOverview) SetIsByoaImageNil(b bool)`

 SetIsByoaImageNil sets the value for IsByoaImage to be an explicit nil

### UnsetIsByoaImage
`func (o *TemplateImageOverview) UnsetIsByoaImage()`

UnsetIsByoaImage ensures that no value is present for IsByoaImage, not even an explicit nil
### GetLinkedCatalogs

`func (o *TemplateImageOverview) GetLinkedCatalogs() int32`

GetLinkedCatalogs returns the LinkedCatalogs field if non-nil, zero value otherwise.

### GetLinkedCatalogsOk

`func (o *TemplateImageOverview) GetLinkedCatalogsOk() (*int32, bool)`

GetLinkedCatalogsOk returns a tuple with the LinkedCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCatalogs

`func (o *TemplateImageOverview) SetLinkedCatalogs(v int32)`

SetLinkedCatalogs sets LinkedCatalogs field to given value.

### HasLinkedCatalogs

`func (o *TemplateImageOverview) HasLinkedCatalogs() bool`

HasLinkedCatalogs returns a boolean if a field has been set.

### GetLinkedCatalogsNames

`func (o *TemplateImageOverview) GetLinkedCatalogsNames() []string`

GetLinkedCatalogsNames returns the LinkedCatalogsNames field if non-nil, zero value otherwise.

### GetLinkedCatalogsNamesOk

`func (o *TemplateImageOverview) GetLinkedCatalogsNamesOk() (*[]string, bool)`

GetLinkedCatalogsNamesOk returns a tuple with the LinkedCatalogsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCatalogsNames

`func (o *TemplateImageOverview) SetLinkedCatalogsNames(v []string)`

SetLinkedCatalogsNames sets LinkedCatalogsNames field to given value.

### HasLinkedCatalogsNames

`func (o *TemplateImageOverview) HasLinkedCatalogsNames() bool`

HasLinkedCatalogsNames returns a boolean if a field has been set.

### SetLinkedCatalogsNamesNil

`func (o *TemplateImageOverview) SetLinkedCatalogsNamesNil(b bool)`

 SetLinkedCatalogsNamesNil sets the value for LinkedCatalogsNames to be an explicit nil

### UnsetLinkedCatalogsNames
`func (o *TemplateImageOverview) UnsetLinkedCatalogsNames()`

UnsetLinkedCatalogsNames ensures that no value is present for LinkedCatalogsNames, not even an explicit nil
### GetCreatedDate

`func (o *TemplateImageOverview) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TemplateImageOverview) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TemplateImageOverview) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TemplateImageOverview) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TemplateImageOverview) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TemplateImageOverview) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetFinalizedDate

`func (o *TemplateImageOverview) GetFinalizedDate() time.Time`

GetFinalizedDate returns the FinalizedDate field if non-nil, zero value otherwise.

### GetFinalizedDateOk

`func (o *TemplateImageOverview) GetFinalizedDateOk() (*time.Time, bool)`

GetFinalizedDateOk returns a tuple with the FinalizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedDate

`func (o *TemplateImageOverview) SetFinalizedDate(v time.Time)`

SetFinalizedDate sets FinalizedDate field to given value.

### HasFinalizedDate

`func (o *TemplateImageOverview) HasFinalizedDate() bool`

HasFinalizedDate returns a boolean if a field has been set.

### SetFinalizedDateNil

`func (o *TemplateImageOverview) SetFinalizedDateNil(b bool)`

 SetFinalizedDateNil sets the value for FinalizedDate to be an explicit nil

### UnsetFinalizedDate
`func (o *TemplateImageOverview) UnsetFinalizedDate()`

UnsetFinalizedDate ensures that no value is present for FinalizedDate, not even an explicit nil
### GetPath

`func (o *TemplateImageOverview) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TemplateImageOverview) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TemplateImageOverview) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TemplateImageOverview) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *TemplateImageOverview) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *TemplateImageOverview) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSbSessionVdaVersion

`func (o *TemplateImageOverview) GetSbSessionVdaVersion() string`

GetSbSessionVdaVersion returns the SbSessionVdaVersion field if non-nil, zero value otherwise.

### GetSbSessionVdaVersionOk

`func (o *TemplateImageOverview) GetSbSessionVdaVersionOk() (*string, bool)`

GetSbSessionVdaVersionOk returns a tuple with the SbSessionVdaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbSessionVdaVersion

`func (o *TemplateImageOverview) SetSbSessionVdaVersion(v string)`

SetSbSessionVdaVersion sets SbSessionVdaVersion field to given value.

### HasSbSessionVdaVersion

`func (o *TemplateImageOverview) HasSbSessionVdaVersion() bool`

HasSbSessionVdaVersion returns a boolean if a field has been set.

### SetSbSessionVdaVersionNil

`func (o *TemplateImageOverview) SetSbSessionVdaVersionNil(b bool)`

 SetSbSessionVdaVersionNil sets the value for SbSessionVdaVersion to be an explicit nil

### UnsetSbSessionVdaVersion
`func (o *TemplateImageOverview) UnsetSbSessionVdaVersion()`

UnsetSbSessionVdaVersion ensures that no value is present for SbSessionVdaVersion, not even an explicit nil
### GetIsSecureBrowserImage

`func (o *TemplateImageOverview) GetIsSecureBrowserImage() bool`

GetIsSecureBrowserImage returns the IsSecureBrowserImage field if non-nil, zero value otherwise.

### GetIsSecureBrowserImageOk

`func (o *TemplateImageOverview) GetIsSecureBrowserImageOk() (*bool, bool)`

GetIsSecureBrowserImageOk returns a tuple with the IsSecureBrowserImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureBrowserImage

`func (o *TemplateImageOverview) SetIsSecureBrowserImage(v bool)`

SetIsSecureBrowserImage sets IsSecureBrowserImage field to given value.

### HasIsSecureBrowserImage

`func (o *TemplateImageOverview) HasIsSecureBrowserImage() bool`

HasIsSecureBrowserImage returns a boolean if a field has been set.

### SetIsSecureBrowserImageNil

`func (o *TemplateImageOverview) SetIsSecureBrowserImageNil(b bool)`

 SetIsSecureBrowserImageNil sets the value for IsSecureBrowserImage to be an explicit nil

### UnsetIsSecureBrowserImage
`func (o *TemplateImageOverview) UnsetIsSecureBrowserImage()`

UnsetIsSecureBrowserImage ensures that no value is present for IsSecureBrowserImage, not even an explicit nil
### GetDiskSizeInGB

`func (o *TemplateImageOverview) GetDiskSizeInGB() int32`

GetDiskSizeInGB returns the DiskSizeInGB field if non-nil, zero value otherwise.

### GetDiskSizeInGBOk

`func (o *TemplateImageOverview) GetDiskSizeInGBOk() (*int32, bool)`

GetDiskSizeInGBOk returns a tuple with the DiskSizeInGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeInGB

`func (o *TemplateImageOverview) SetDiskSizeInGB(v int32)`

SetDiskSizeInGB sets DiskSizeInGB field to given value.

### HasDiskSizeInGB

`func (o *TemplateImageOverview) HasDiskSizeInGB() bool`

HasDiskSizeInGB returns a boolean if a field has been set.

### SetDiskSizeInGBNil

`func (o *TemplateImageOverview) SetDiskSizeInGBNil(b bool)`

 SetDiskSizeInGBNil sets the value for DiskSizeInGB to be an explicit nil

### UnsetDiskSizeInGB
`func (o *TemplateImageOverview) UnsetDiskSizeInGB()`

UnsetDiskSizeInGB ensures that no value is present for DiskSizeInGB, not even an explicit nil
### GetStartedAt

`func (o *TemplateImageOverview) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TemplateImageOverview) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TemplateImageOverview) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TemplateImageOverview) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *TemplateImageOverview) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *TemplateImageOverview) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEstimatedTimeInMinute

`func (o *TemplateImageOverview) GetEstimatedTimeInMinute() int32`

GetEstimatedTimeInMinute returns the EstimatedTimeInMinute field if non-nil, zero value otherwise.

### GetEstimatedTimeInMinuteOk

`func (o *TemplateImageOverview) GetEstimatedTimeInMinuteOk() (*int32, bool)`

GetEstimatedTimeInMinuteOk returns a tuple with the EstimatedTimeInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeInMinute

`func (o *TemplateImageOverview) SetEstimatedTimeInMinute(v int32)`

SetEstimatedTimeInMinute sets EstimatedTimeInMinute field to given value.

### HasEstimatedTimeInMinute

`func (o *TemplateImageOverview) HasEstimatedTimeInMinute() bool`

HasEstimatedTimeInMinute returns a boolean if a field has been set.

### SetEstimatedTimeInMinuteNil

`func (o *TemplateImageOverview) SetEstimatedTimeInMinuteNil(b bool)`

 SetEstimatedTimeInMinuteNil sets the value for EstimatedTimeInMinute to be an explicit nil

### UnsetEstimatedTimeInMinute
`func (o *TemplateImageOverview) UnsetEstimatedTimeInMinute()`

UnsetEstimatedTimeInMinute ensures that no value is present for EstimatedTimeInMinute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


