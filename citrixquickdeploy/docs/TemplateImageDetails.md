# TemplateImageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalogs** | [**[]TemplateImageAssociatedCatalog**](TemplateImageAssociatedCatalog.md) | Catalogs that are using this image | 
**Applications** | [**[]Application**](Application.md) | Applications that have been discovered in this template | 
**Id** | **string** | Id of the template image | 
**Name** | **string** | Name of the Template Image / VM | 
**SessionSupport** | [**TemplateImageSessionSupport**](TemplateImageSessionSupport.md) | Type of sessions that are supported by OS | 
**State** | [**TemplateImageState**](TemplateImageState.md) | State of the template image | 
**SubState** | Pointer to [**TemplateImageSubState**](TemplateImageSubState.md) | Sub State of template image | [optional] 
**OsPlatform** | Pointer to [**SupportedOperatingSystemType**](SupportedOperatingSystemType.md) | Type of operating system that will be imported | [optional] 
**OsName** | Pointer to **string** | Shows name of image OS | [optional] 
**OsVersion** | Pointer to **string** | Shows version of image OS | [optional] 
**LinuxDomainSupport** | Pointer to **string** | For linux customers it shows domain configuration set in mcs.conf file | [optional] 
**LinuxRdpSupport** | Pointer to **bool** | For linux customers it shows if rdp package was installed | [optional] 
**McsioSupport** | Pointer to **bool** | For windows customers it shows if mcsio driver was installed | [optional] 
**VusSupport** | Pointer to **bool** | For windows customers it shows if VDA update service was installed | [optional] 
**PublicIp** | Pointer to **string** | Shows public ip address for image | [optional] 
**PrivateIp** | Pointer to **string** | Shows private ip address for image | [optional] 
**DomainName** | Pointer to **string** | Shows if image is currently domain joined and to which domain | [optional] 
**IsServerOs** | Pointer to **bool** | Indicates if the image is using a server based OS | [optional] 
**VdaVersion** | Pointer to **string** | Shows version of image Virtual Desktop Agent | [optional] 
**StatusMessageId** | Pointer to **string** | Status message enum related to verifying and enumerating the image | [optional] 
**Status** | Pointer to **string** | Status message related to verifying and enumerating the image | [optional] 
**ExtraInfo** | Pointer to **string** | The string to displayed in UI for extra information | [optional] 
**Notes** | Pointer to **string** | Customer notes about template image | [optional] 
**TransactionId** | Pointer to **string** | ID of the transaction that the image was verified on | [optional] 
**SubscriptionId** | Pointer to **string** | Id of the Subscription where the image is stored (BYOA) | [optional] 
**SubscriptionName** | **string** | Name of the Subscription that catalog VMs will be deployed to | 
**ResourceGroup** | Pointer to **string** | Name of the Azure Resource Group where the image is stored | [optional] 
**StorageAccount** | Pointer to **string** | Name of the Storage Account where the image is stored | [optional] 
**Region** | **string** | Azure region where VMs are deployed for this catalog | 
**BuilderDomainName** | Pointer to **string** | Name of the Domain the Image Builder will join | [optional] 
**BuilderConnectionId** | Pointer to **string** | ID of the on-prem connection associated with the builder image | [optional] 
**BuilderVmName** | Pointer to **string** | Name of VM that is being used by the builder image | [optional] 
**BuilderVmType** | Pointer to **string** | Type of VM that is being used by the builder image | [optional] 
**BuilderVmDiskSize** | Pointer to **string** | Size of disk of VM that is being used by the builder image | [optional] 
**BuilderAllowedIPs** | Pointer to **[]string** | Ip Addresses allowed to RDP | [optional] 
**HyperVGen** | Pointer to **string** | The HyperVGeneration that should be set to either V1 or V2 | [optional] 
**VtpmEnabled** | Pointer to **bool** | Is TrustedLaunch VTPM supported in V2 gen | [optional] 
**SecureBootEnabled** | Pointer to **bool** | Is TrustedLaunch SecureBoot supported in V2 gen | [optional] 
**CitrixPrepared** | **bool** | Whether the image was prepared by Citrix, or provided by the customer | 
**CspCustomer** | Pointer to **string** | Indicates that partner-tenant relationship exists if not null | [optional] 
**IsCmekEnabled** | Pointer to **bool** | Indicates if the image is using customer managed encryption keys | [optional] 
**CmekId** | Pointer to **string** | Indicates if customer managed encryption key Id | [optional] 
**IsDeprecated** | Pointer to **bool** | Indicates if the image is deprecated and should not be used in new catalogs | [optional] 
**IsByoaImage** | Pointer to **bool** | Indicates if the image is built or imported with BYOA | [optional] 
**LinkedCatalogs** | Pointer to **int32** | Number of catalogs that are using this image | [optional] 
**LinkedCatalogsNames** | Pointer to **[]string** | Names of catalogs that are using this image | [optional] 
**CreatedDate** | Pointer to **time.Time** | Created from datastore. | [optional] 
**FinalizedDate** | Pointer to **time.Time** | Image builder finalized date. | [optional] 
**Path** | Pointer to **string** | Customer image path in Azure | [optional] 
**SbSessionVdaVersion** | Pointer to **string** | Shows version of sbsession | [optional] 
**IsSecureBrowserImage** | Pointer to **bool** | Indicates if the image is for Secure Browser | [optional] 
**StartedAt** | Pointer to **time.Time** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **int32** | Estimated total time for the job to finish | [optional] 

## Methods

### NewTemplateImageDetails

`func NewTemplateImageDetails(catalogs []TemplateImageAssociatedCatalog, applications []Application, id string, name string, sessionSupport TemplateImageSessionSupport, state TemplateImageState, subscriptionName string, region string, citrixPrepared bool, ) *TemplateImageDetails`

NewTemplateImageDetails instantiates a new TemplateImageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateImageDetailsWithDefaults

`func NewTemplateImageDetailsWithDefaults() *TemplateImageDetails`

NewTemplateImageDetailsWithDefaults instantiates a new TemplateImageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogs

`func (o *TemplateImageDetails) GetCatalogs() []TemplateImageAssociatedCatalog`

GetCatalogs returns the Catalogs field if non-nil, zero value otherwise.

### GetCatalogsOk

`func (o *TemplateImageDetails) GetCatalogsOk() (*[]TemplateImageAssociatedCatalog, bool)`

GetCatalogsOk returns a tuple with the Catalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogs

`func (o *TemplateImageDetails) SetCatalogs(v []TemplateImageAssociatedCatalog)`

SetCatalogs sets Catalogs field to given value.


### GetApplications

`func (o *TemplateImageDetails) GetApplications() []Application`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *TemplateImageDetails) GetApplicationsOk() (*[]Application, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *TemplateImageDetails) SetApplications(v []Application)`

SetApplications sets Applications field to given value.


### GetId

`func (o *TemplateImageDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateImageDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateImageDetails) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TemplateImageDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateImageDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateImageDetails) SetName(v string)`

SetName sets Name field to given value.


### GetSessionSupport

`func (o *TemplateImageDetails) GetSessionSupport() TemplateImageSessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *TemplateImageDetails) GetSessionSupportOk() (*TemplateImageSessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *TemplateImageDetails) SetSessionSupport(v TemplateImageSessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetState

`func (o *TemplateImageDetails) GetState() TemplateImageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TemplateImageDetails) GetStateOk() (*TemplateImageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TemplateImageDetails) SetState(v TemplateImageState)`

SetState sets State field to given value.


### GetSubState

`func (o *TemplateImageDetails) GetSubState() TemplateImageSubState`

GetSubState returns the SubState field if non-nil, zero value otherwise.

### GetSubStateOk

`func (o *TemplateImageDetails) GetSubStateOk() (*TemplateImageSubState, bool)`

GetSubStateOk returns a tuple with the SubState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubState

`func (o *TemplateImageDetails) SetSubState(v TemplateImageSubState)`

SetSubState sets SubState field to given value.

### HasSubState

`func (o *TemplateImageDetails) HasSubState() bool`

HasSubState returns a boolean if a field has been set.

### GetOsPlatform

`func (o *TemplateImageDetails) GetOsPlatform() SupportedOperatingSystemType`

GetOsPlatform returns the OsPlatform field if non-nil, zero value otherwise.

### GetOsPlatformOk

`func (o *TemplateImageDetails) GetOsPlatformOk() (*SupportedOperatingSystemType, bool)`

GetOsPlatformOk returns a tuple with the OsPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPlatform

`func (o *TemplateImageDetails) SetOsPlatform(v SupportedOperatingSystemType)`

SetOsPlatform sets OsPlatform field to given value.

### HasOsPlatform

`func (o *TemplateImageDetails) HasOsPlatform() bool`

HasOsPlatform returns a boolean if a field has been set.

### GetOsName

`func (o *TemplateImageDetails) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *TemplateImageDetails) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *TemplateImageDetails) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *TemplateImageDetails) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *TemplateImageDetails) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *TemplateImageDetails) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *TemplateImageDetails) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *TemplateImageDetails) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetLinuxDomainSupport

`func (o *TemplateImageDetails) GetLinuxDomainSupport() string`

GetLinuxDomainSupport returns the LinuxDomainSupport field if non-nil, zero value otherwise.

### GetLinuxDomainSupportOk

`func (o *TemplateImageDetails) GetLinuxDomainSupportOk() (*string, bool)`

GetLinuxDomainSupportOk returns a tuple with the LinuxDomainSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxDomainSupport

`func (o *TemplateImageDetails) SetLinuxDomainSupport(v string)`

SetLinuxDomainSupport sets LinuxDomainSupport field to given value.

### HasLinuxDomainSupport

`func (o *TemplateImageDetails) HasLinuxDomainSupport() bool`

HasLinuxDomainSupport returns a boolean if a field has been set.

### GetLinuxRdpSupport

`func (o *TemplateImageDetails) GetLinuxRdpSupport() bool`

GetLinuxRdpSupport returns the LinuxRdpSupport field if non-nil, zero value otherwise.

### GetLinuxRdpSupportOk

`func (o *TemplateImageDetails) GetLinuxRdpSupportOk() (*bool, bool)`

GetLinuxRdpSupportOk returns a tuple with the LinuxRdpSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxRdpSupport

`func (o *TemplateImageDetails) SetLinuxRdpSupport(v bool)`

SetLinuxRdpSupport sets LinuxRdpSupport field to given value.

### HasLinuxRdpSupport

`func (o *TemplateImageDetails) HasLinuxRdpSupport() bool`

HasLinuxRdpSupport returns a boolean if a field has been set.

### GetMcsioSupport

`func (o *TemplateImageDetails) GetMcsioSupport() bool`

GetMcsioSupport returns the McsioSupport field if non-nil, zero value otherwise.

### GetMcsioSupportOk

`func (o *TemplateImageDetails) GetMcsioSupportOk() (*bool, bool)`

GetMcsioSupportOk returns a tuple with the McsioSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsioSupport

`func (o *TemplateImageDetails) SetMcsioSupport(v bool)`

SetMcsioSupport sets McsioSupport field to given value.

### HasMcsioSupport

`func (o *TemplateImageDetails) HasMcsioSupport() bool`

HasMcsioSupport returns a boolean if a field has been set.

### GetVusSupport

`func (o *TemplateImageDetails) GetVusSupport() bool`

GetVusSupport returns the VusSupport field if non-nil, zero value otherwise.

### GetVusSupportOk

`func (o *TemplateImageDetails) GetVusSupportOk() (*bool, bool)`

GetVusSupportOk returns a tuple with the VusSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVusSupport

`func (o *TemplateImageDetails) SetVusSupport(v bool)`

SetVusSupport sets VusSupport field to given value.

### HasVusSupport

`func (o *TemplateImageDetails) HasVusSupport() bool`

HasVusSupport returns a boolean if a field has been set.

### GetPublicIp

`func (o *TemplateImageDetails) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *TemplateImageDetails) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *TemplateImageDetails) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *TemplateImageDetails) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPrivateIp

`func (o *TemplateImageDetails) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *TemplateImageDetails) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *TemplateImageDetails) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *TemplateImageDetails) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetDomainName

`func (o *TemplateImageDetails) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *TemplateImageDetails) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *TemplateImageDetails) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *TemplateImageDetails) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetIsServerOs

`func (o *TemplateImageDetails) GetIsServerOs() bool`

GetIsServerOs returns the IsServerOs field if non-nil, zero value otherwise.

### GetIsServerOsOk

`func (o *TemplateImageDetails) GetIsServerOsOk() (*bool, bool)`

GetIsServerOsOk returns a tuple with the IsServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServerOs

`func (o *TemplateImageDetails) SetIsServerOs(v bool)`

SetIsServerOs sets IsServerOs field to given value.

### HasIsServerOs

`func (o *TemplateImageDetails) HasIsServerOs() bool`

HasIsServerOs returns a boolean if a field has been set.

### GetVdaVersion

`func (o *TemplateImageDetails) GetVdaVersion() string`

GetVdaVersion returns the VdaVersion field if non-nil, zero value otherwise.

### GetVdaVersionOk

`func (o *TemplateImageDetails) GetVdaVersionOk() (*string, bool)`

GetVdaVersionOk returns a tuple with the VdaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaVersion

`func (o *TemplateImageDetails) SetVdaVersion(v string)`

SetVdaVersion sets VdaVersion field to given value.

### HasVdaVersion

`func (o *TemplateImageDetails) HasVdaVersion() bool`

HasVdaVersion returns a boolean if a field has been set.

### GetStatusMessageId

`func (o *TemplateImageDetails) GetStatusMessageId() string`

GetStatusMessageId returns the StatusMessageId field if non-nil, zero value otherwise.

### GetStatusMessageIdOk

`func (o *TemplateImageDetails) GetStatusMessageIdOk() (*string, bool)`

GetStatusMessageIdOk returns a tuple with the StatusMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessageId

`func (o *TemplateImageDetails) SetStatusMessageId(v string)`

SetStatusMessageId sets StatusMessageId field to given value.

### HasStatusMessageId

`func (o *TemplateImageDetails) HasStatusMessageId() bool`

HasStatusMessageId returns a boolean if a field has been set.

### GetStatus

`func (o *TemplateImageDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateImageDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateImageDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TemplateImageDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExtraInfo

`func (o *TemplateImageDetails) GetExtraInfo() string`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *TemplateImageDetails) GetExtraInfoOk() (*string, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *TemplateImageDetails) SetExtraInfo(v string)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *TemplateImageDetails) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetNotes

`func (o *TemplateImageDetails) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TemplateImageDetails) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TemplateImageDetails) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TemplateImageDetails) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTransactionId

`func (o *TemplateImageDetails) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TemplateImageDetails) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TemplateImageDetails) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TemplateImageDetails) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *TemplateImageDetails) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *TemplateImageDetails) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *TemplateImageDetails) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *TemplateImageDetails) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionName

`func (o *TemplateImageDetails) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *TemplateImageDetails) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *TemplateImageDetails) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.


### GetResourceGroup

`func (o *TemplateImageDetails) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *TemplateImageDetails) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *TemplateImageDetails) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *TemplateImageDetails) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetStorageAccount

`func (o *TemplateImageDetails) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *TemplateImageDetails) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *TemplateImageDetails) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *TemplateImageDetails) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### GetRegion

`func (o *TemplateImageDetails) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TemplateImageDetails) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TemplateImageDetails) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetBuilderDomainName

`func (o *TemplateImageDetails) GetBuilderDomainName() string`

GetBuilderDomainName returns the BuilderDomainName field if non-nil, zero value otherwise.

### GetBuilderDomainNameOk

`func (o *TemplateImageDetails) GetBuilderDomainNameOk() (*string, bool)`

GetBuilderDomainNameOk returns a tuple with the BuilderDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderDomainName

`func (o *TemplateImageDetails) SetBuilderDomainName(v string)`

SetBuilderDomainName sets BuilderDomainName field to given value.

### HasBuilderDomainName

`func (o *TemplateImageDetails) HasBuilderDomainName() bool`

HasBuilderDomainName returns a boolean if a field has been set.

### GetBuilderConnectionId

`func (o *TemplateImageDetails) GetBuilderConnectionId() string`

GetBuilderConnectionId returns the BuilderConnectionId field if non-nil, zero value otherwise.

### GetBuilderConnectionIdOk

`func (o *TemplateImageDetails) GetBuilderConnectionIdOk() (*string, bool)`

GetBuilderConnectionIdOk returns a tuple with the BuilderConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderConnectionId

`func (o *TemplateImageDetails) SetBuilderConnectionId(v string)`

SetBuilderConnectionId sets BuilderConnectionId field to given value.

### HasBuilderConnectionId

`func (o *TemplateImageDetails) HasBuilderConnectionId() bool`

HasBuilderConnectionId returns a boolean if a field has been set.

### GetBuilderVmName

`func (o *TemplateImageDetails) GetBuilderVmName() string`

GetBuilderVmName returns the BuilderVmName field if non-nil, zero value otherwise.

### GetBuilderVmNameOk

`func (o *TemplateImageDetails) GetBuilderVmNameOk() (*string, bool)`

GetBuilderVmNameOk returns a tuple with the BuilderVmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderVmName

`func (o *TemplateImageDetails) SetBuilderVmName(v string)`

SetBuilderVmName sets BuilderVmName field to given value.

### HasBuilderVmName

`func (o *TemplateImageDetails) HasBuilderVmName() bool`

HasBuilderVmName returns a boolean if a field has been set.

### GetBuilderVmType

`func (o *TemplateImageDetails) GetBuilderVmType() string`

GetBuilderVmType returns the BuilderVmType field if non-nil, zero value otherwise.

### GetBuilderVmTypeOk

`func (o *TemplateImageDetails) GetBuilderVmTypeOk() (*string, bool)`

GetBuilderVmTypeOk returns a tuple with the BuilderVmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderVmType

`func (o *TemplateImageDetails) SetBuilderVmType(v string)`

SetBuilderVmType sets BuilderVmType field to given value.

### HasBuilderVmType

`func (o *TemplateImageDetails) HasBuilderVmType() bool`

HasBuilderVmType returns a boolean if a field has been set.

### GetBuilderVmDiskSize

`func (o *TemplateImageDetails) GetBuilderVmDiskSize() string`

GetBuilderVmDiskSize returns the BuilderVmDiskSize field if non-nil, zero value otherwise.

### GetBuilderVmDiskSizeOk

`func (o *TemplateImageDetails) GetBuilderVmDiskSizeOk() (*string, bool)`

GetBuilderVmDiskSizeOk returns a tuple with the BuilderVmDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderVmDiskSize

`func (o *TemplateImageDetails) SetBuilderVmDiskSize(v string)`

SetBuilderVmDiskSize sets BuilderVmDiskSize field to given value.

### HasBuilderVmDiskSize

`func (o *TemplateImageDetails) HasBuilderVmDiskSize() bool`

HasBuilderVmDiskSize returns a boolean if a field has been set.

### GetBuilderAllowedIPs

`func (o *TemplateImageDetails) GetBuilderAllowedIPs() []string`

GetBuilderAllowedIPs returns the BuilderAllowedIPs field if non-nil, zero value otherwise.

### GetBuilderAllowedIPsOk

`func (o *TemplateImageDetails) GetBuilderAllowedIPsOk() (*[]string, bool)`

GetBuilderAllowedIPsOk returns a tuple with the BuilderAllowedIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderAllowedIPs

`func (o *TemplateImageDetails) SetBuilderAllowedIPs(v []string)`

SetBuilderAllowedIPs sets BuilderAllowedIPs field to given value.

### HasBuilderAllowedIPs

`func (o *TemplateImageDetails) HasBuilderAllowedIPs() bool`

HasBuilderAllowedIPs returns a boolean if a field has been set.

### GetHyperVGen

`func (o *TemplateImageDetails) GetHyperVGen() string`

GetHyperVGen returns the HyperVGen field if non-nil, zero value otherwise.

### GetHyperVGenOk

`func (o *TemplateImageDetails) GetHyperVGenOk() (*string, bool)`

GetHyperVGenOk returns a tuple with the HyperVGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperVGen

`func (o *TemplateImageDetails) SetHyperVGen(v string)`

SetHyperVGen sets HyperVGen field to given value.

### HasHyperVGen

`func (o *TemplateImageDetails) HasHyperVGen() bool`

HasHyperVGen returns a boolean if a field has been set.

### GetVtpmEnabled

`func (o *TemplateImageDetails) GetVtpmEnabled() bool`

GetVtpmEnabled returns the VtpmEnabled field if non-nil, zero value otherwise.

### GetVtpmEnabledOk

`func (o *TemplateImageDetails) GetVtpmEnabledOk() (*bool, bool)`

GetVtpmEnabledOk returns a tuple with the VtpmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpmEnabled

`func (o *TemplateImageDetails) SetVtpmEnabled(v bool)`

SetVtpmEnabled sets VtpmEnabled field to given value.

### HasVtpmEnabled

`func (o *TemplateImageDetails) HasVtpmEnabled() bool`

HasVtpmEnabled returns a boolean if a field has been set.

### GetSecureBootEnabled

`func (o *TemplateImageDetails) GetSecureBootEnabled() bool`

GetSecureBootEnabled returns the SecureBootEnabled field if non-nil, zero value otherwise.

### GetSecureBootEnabledOk

`func (o *TemplateImageDetails) GetSecureBootEnabledOk() (*bool, bool)`

GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBootEnabled

`func (o *TemplateImageDetails) SetSecureBootEnabled(v bool)`

SetSecureBootEnabled sets SecureBootEnabled field to given value.

### HasSecureBootEnabled

`func (o *TemplateImageDetails) HasSecureBootEnabled() bool`

HasSecureBootEnabled returns a boolean if a field has been set.

### GetCitrixPrepared

`func (o *TemplateImageDetails) GetCitrixPrepared() bool`

GetCitrixPrepared returns the CitrixPrepared field if non-nil, zero value otherwise.

### GetCitrixPreparedOk

`func (o *TemplateImageDetails) GetCitrixPreparedOk() (*bool, bool)`

GetCitrixPreparedOk returns a tuple with the CitrixPrepared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixPrepared

`func (o *TemplateImageDetails) SetCitrixPrepared(v bool)`

SetCitrixPrepared sets CitrixPrepared field to given value.


### GetCspCustomer

`func (o *TemplateImageDetails) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *TemplateImageDetails) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *TemplateImageDetails) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *TemplateImageDetails) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### GetIsCmekEnabled

`func (o *TemplateImageDetails) GetIsCmekEnabled() bool`

GetIsCmekEnabled returns the IsCmekEnabled field if non-nil, zero value otherwise.

### GetIsCmekEnabledOk

`func (o *TemplateImageDetails) GetIsCmekEnabledOk() (*bool, bool)`

GetIsCmekEnabledOk returns a tuple with the IsCmekEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCmekEnabled

`func (o *TemplateImageDetails) SetIsCmekEnabled(v bool)`

SetIsCmekEnabled sets IsCmekEnabled field to given value.

### HasIsCmekEnabled

`func (o *TemplateImageDetails) HasIsCmekEnabled() bool`

HasIsCmekEnabled returns a boolean if a field has been set.

### GetCmekId

`func (o *TemplateImageDetails) GetCmekId() string`

GetCmekId returns the CmekId field if non-nil, zero value otherwise.

### GetCmekIdOk

`func (o *TemplateImageDetails) GetCmekIdOk() (*string, bool)`

GetCmekIdOk returns a tuple with the CmekId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmekId

`func (o *TemplateImageDetails) SetCmekId(v string)`

SetCmekId sets CmekId field to given value.

### HasCmekId

`func (o *TemplateImageDetails) HasCmekId() bool`

HasCmekId returns a boolean if a field has been set.

### GetIsDeprecated

`func (o *TemplateImageDetails) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *TemplateImageDetails) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *TemplateImageDetails) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *TemplateImageDetails) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.

### GetIsByoaImage

`func (o *TemplateImageDetails) GetIsByoaImage() bool`

GetIsByoaImage returns the IsByoaImage field if non-nil, zero value otherwise.

### GetIsByoaImageOk

`func (o *TemplateImageDetails) GetIsByoaImageOk() (*bool, bool)`

GetIsByoaImageOk returns a tuple with the IsByoaImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByoaImage

`func (o *TemplateImageDetails) SetIsByoaImage(v bool)`

SetIsByoaImage sets IsByoaImage field to given value.

### HasIsByoaImage

`func (o *TemplateImageDetails) HasIsByoaImage() bool`

HasIsByoaImage returns a boolean if a field has been set.

### GetLinkedCatalogs

`func (o *TemplateImageDetails) GetLinkedCatalogs() int32`

GetLinkedCatalogs returns the LinkedCatalogs field if non-nil, zero value otherwise.

### GetLinkedCatalogsOk

`func (o *TemplateImageDetails) GetLinkedCatalogsOk() (*int32, bool)`

GetLinkedCatalogsOk returns a tuple with the LinkedCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCatalogs

`func (o *TemplateImageDetails) SetLinkedCatalogs(v int32)`

SetLinkedCatalogs sets LinkedCatalogs field to given value.

### HasLinkedCatalogs

`func (o *TemplateImageDetails) HasLinkedCatalogs() bool`

HasLinkedCatalogs returns a boolean if a field has been set.

### GetLinkedCatalogsNames

`func (o *TemplateImageDetails) GetLinkedCatalogsNames() []string`

GetLinkedCatalogsNames returns the LinkedCatalogsNames field if non-nil, zero value otherwise.

### GetLinkedCatalogsNamesOk

`func (o *TemplateImageDetails) GetLinkedCatalogsNamesOk() (*[]string, bool)`

GetLinkedCatalogsNamesOk returns a tuple with the LinkedCatalogsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCatalogsNames

`func (o *TemplateImageDetails) SetLinkedCatalogsNames(v []string)`

SetLinkedCatalogsNames sets LinkedCatalogsNames field to given value.

### HasLinkedCatalogsNames

`func (o *TemplateImageDetails) HasLinkedCatalogsNames() bool`

HasLinkedCatalogsNames returns a boolean if a field has been set.

### GetCreatedDate

`func (o *TemplateImageDetails) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TemplateImageDetails) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TemplateImageDetails) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TemplateImageDetails) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetFinalizedDate

`func (o *TemplateImageDetails) GetFinalizedDate() time.Time`

GetFinalizedDate returns the FinalizedDate field if non-nil, zero value otherwise.

### GetFinalizedDateOk

`func (o *TemplateImageDetails) GetFinalizedDateOk() (*time.Time, bool)`

GetFinalizedDateOk returns a tuple with the FinalizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedDate

`func (o *TemplateImageDetails) SetFinalizedDate(v time.Time)`

SetFinalizedDate sets FinalizedDate field to given value.

### HasFinalizedDate

`func (o *TemplateImageDetails) HasFinalizedDate() bool`

HasFinalizedDate returns a boolean if a field has been set.

### GetPath

`func (o *TemplateImageDetails) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TemplateImageDetails) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TemplateImageDetails) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TemplateImageDetails) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSbSessionVdaVersion

`func (o *TemplateImageDetails) GetSbSessionVdaVersion() string`

GetSbSessionVdaVersion returns the SbSessionVdaVersion field if non-nil, zero value otherwise.

### GetSbSessionVdaVersionOk

`func (o *TemplateImageDetails) GetSbSessionVdaVersionOk() (*string, bool)`

GetSbSessionVdaVersionOk returns a tuple with the SbSessionVdaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbSessionVdaVersion

`func (o *TemplateImageDetails) SetSbSessionVdaVersion(v string)`

SetSbSessionVdaVersion sets SbSessionVdaVersion field to given value.

### HasSbSessionVdaVersion

`func (o *TemplateImageDetails) HasSbSessionVdaVersion() bool`

HasSbSessionVdaVersion returns a boolean if a field has been set.

### GetIsSecureBrowserImage

`func (o *TemplateImageDetails) GetIsSecureBrowserImage() bool`

GetIsSecureBrowserImage returns the IsSecureBrowserImage field if non-nil, zero value otherwise.

### GetIsSecureBrowserImageOk

`func (o *TemplateImageDetails) GetIsSecureBrowserImageOk() (*bool, bool)`

GetIsSecureBrowserImageOk returns a tuple with the IsSecureBrowserImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureBrowserImage

`func (o *TemplateImageDetails) SetIsSecureBrowserImage(v bool)`

SetIsSecureBrowserImage sets IsSecureBrowserImage field to given value.

### HasIsSecureBrowserImage

`func (o *TemplateImageDetails) HasIsSecureBrowserImage() bool`

HasIsSecureBrowserImage returns a boolean if a field has been set.

### GetStartedAt

`func (o *TemplateImageDetails) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TemplateImageDetails) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TemplateImageDetails) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TemplateImageDetails) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEstimatedTimeInMinute

`func (o *TemplateImageDetails) GetEstimatedTimeInMinute() int32`

GetEstimatedTimeInMinute returns the EstimatedTimeInMinute field if non-nil, zero value otherwise.

### GetEstimatedTimeInMinuteOk

`func (o *TemplateImageDetails) GetEstimatedTimeInMinuteOk() (*int32, bool)`

GetEstimatedTimeInMinuteOk returns a tuple with the EstimatedTimeInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeInMinute

`func (o *TemplateImageDetails) SetEstimatedTimeInMinute(v int32)`

SetEstimatedTimeInMinute sets EstimatedTimeInMinute field to given value.

### HasEstimatedTimeInMinute

`func (o *TemplateImageDetails) HasEstimatedTimeInMinute() bool`

HasEstimatedTimeInMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


