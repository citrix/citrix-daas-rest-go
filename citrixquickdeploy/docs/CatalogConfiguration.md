# CatalogConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the catalog | 
**Name** | **string** | User configured name | 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) | Quantity of sessions supported per-machine. | [optional] 
**AllocationType** | Pointer to [**CatalogAllocationType**](CatalogAllocationType.md) | Indicates the manner in which machines are allocated to users | [optional] 
**PersistStaticAllocatedVmDisks** | Pointer to **bool** | Indicates if catalogs that use statically allocated machines will have the disk contents persisted after shutdown | [optional] 
**OfferingId** | **string** | The offeringId for the catalog to be used in Cloud Library operations | 
**OfferingIdApp** | **string** | The Application offeringId for the catalog to be used in Cloud Library operations | 
**OfferingIdDesktop** | **string** | The Desktop offeringId for the catalog to be used in Cloud Library operations | 
**DeliveryGroupId** | **string** | The Delivery Group Id | 
**Advanced** | **bool** | Indicates if this is an advanced object | 
**State** | [**CatalogOverallState**](CatalogOverallState.md) | Status of the catalog | 
**SubState** | Pointer to [**CatalogOverallSubState**](CatalogOverallSubState.md) | Sub Status of the catalog | [optional] 
**Warnings** | Pointer to [**[]CatalogWarning**](CatalogWarning.md) |  | [optional] 
**StatusMessage** | Pointer to **string** | Current status of the catalog | [optional] 
**ExtraInfo** | Pointer to **string** | The string to displayed in UI for extra information | [optional] 
**TransactionId** | Pointer to **string** | The transaction id of the catalog deployment | [optional] 
**VnetPeeringId** | Pointer to **string** | ID of the Vnet Peering associated with the catalog | [optional] 
**VnetPeeringName** | Pointer to **string** | Name of the Vnet Peering associated with the catalog | [optional] 
**VpnConnectionId** | Pointer to **string** | ID of the Vpn Connection associated with the catalog | [optional] 
**VpnConnectionName** | Pointer to **string** | Name of the Vpn Connection associated with the catalog | [optional] 
**SubscriptionId** | Pointer to **string** | Id of the Subscription that catalog VMs will be deployed to | [optional] 
**SubscriptionName** | **string** | Name of the Subscription that catalog VMs will be deployed to | 
**ResourceGroup** | **string** | Name of the resource group used for this catalog | 
**VdaResourceGroup** | Pointer to **string** | The resource group for the VDAs in Azure | [optional] 
**VdaProvisioningSchemeId** | Pointer to **string** | The resource groups for the VDAs in Azure | [optional] 
**AreMcsVdaResourceGroupsUsed** | Pointer to **bool** | The resource groups for the VDAs in Azure | [optional] 
**ResourceLocationId** | Pointer to **string** | ID of the Resource Location associated with the catalog | [optional] 
**Region** | **string** | Azure region where VMs are deployed for this catalog | 
**VNetName** | **string** | Name of the vnet assigned to the catalog | 
**Subnet** | Pointer to **string** | The subnet that is associated with the catalog&#39;s VNet | [optional] 
**DomainJoined** | Pointer to **bool** | The flag to indicate if the catalog is joined with customer domain | [optional] 
**DomainName** | Pointer to **string** | Name of the domain that the catalog&#39;s VMs will join | [optional] 
**DomainOU** | Pointer to **string** | OU of the domain we are joining | [optional] 
**DomainServiceAccount** | Pointer to **string** | Name of the service account that will perform domain join opperations | [optional] 
**VmTypeInstanceType** | Pointer to **string** | Type of the VM machines used to create VDAs | [optional] [readonly] 
**ImageId** | Pointer to **string** | ID of the image that is used by the catalog | [optional] 
**TemplateImageName** | Pointer to **string** | Name of the template image that we are using for this catalog | [optional] 
**TemplateImageOs** | Pointer to **string** | Os type of the template image that we are using for this catalog | [optional] 
**CitrixManaged** | Pointer to **bool** | Indicates that partner-tenant relationship exists if not null | [optional] 
**CspCustomer** | Pointer to **string** | Indicates that partner-tenant relationship exists if not null | [optional] 
**TotalMachinesInCatalog** | Pointer to **int32** | Maximum number of machines assigned to the catalog | [optional] 
**WriteBackCacheConfiguration** | Pointer to [**WbcConfig**](WbcConfig.md) | Indicates whether or not write back cache is enabled for the VMs created from this provisioning scheme. | [optional] 
**TaskCompletionPercentage** | Pointer to **int32** | Percentage complete the current task being performed on the catalog is at | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | Last time when the catalog was modified | [optional] 
**LastBackupTime** | Pointer to **time.Time** | Last backup time for the catalog&#39;s VDAs | [optional] [readonly] 
**IsRemotePcCatalog** | Pointer to **bool** | Indicates if this is a remote pc catalog | [optional] 
**IsAzureAdJoined** | Pointer to **bool** | Indicates if the machines in the catalog will be Azure AD joined | [optional] 
**IsSecureBrowserCatalog** | Pointer to **bool** | Indicates if the catalog is for Secure Browser service | [optional] [readonly] 
**OrganizationalUnits** | Pointer to [**[]RemotePCEnrollmentScopeResponseModel**](RemotePCEnrollmentScopeResponseModel.md) | List of OUs for remote pc | [optional] 
**SupportsHibernation** | Pointer to **bool** | Indicates whether machines in catalog support hibernation | [optional] 

## Methods

### NewCatalogConfiguration

`func NewCatalogConfiguration(id string, name string, offeringId string, offeringIdApp string, offeringIdDesktop string, deliveryGroupId string, advanced bool, state CatalogOverallState, subscriptionName string, resourceGroup string, region string, vNetName string, ) *CatalogConfiguration`

NewCatalogConfiguration instantiates a new CatalogConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogConfigurationWithDefaults

`func NewCatalogConfigurationWithDefaults() *CatalogConfiguration`

NewCatalogConfigurationWithDefaults instantiates a new CatalogConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CatalogConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetSessionSupport

`func (o *CatalogConfiguration) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *CatalogConfiguration) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *CatalogConfiguration) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *CatalogConfiguration) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.

### GetAllocationType

`func (o *CatalogConfiguration) GetAllocationType() CatalogAllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *CatalogConfiguration) GetAllocationTypeOk() (*CatalogAllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *CatalogConfiguration) SetAllocationType(v CatalogAllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *CatalogConfiguration) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetPersistStaticAllocatedVmDisks

`func (o *CatalogConfiguration) GetPersistStaticAllocatedVmDisks() bool`

GetPersistStaticAllocatedVmDisks returns the PersistStaticAllocatedVmDisks field if non-nil, zero value otherwise.

### GetPersistStaticAllocatedVmDisksOk

`func (o *CatalogConfiguration) GetPersistStaticAllocatedVmDisksOk() (*bool, bool)`

GetPersistStaticAllocatedVmDisksOk returns a tuple with the PersistStaticAllocatedVmDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistStaticAllocatedVmDisks

`func (o *CatalogConfiguration) SetPersistStaticAllocatedVmDisks(v bool)`

SetPersistStaticAllocatedVmDisks sets PersistStaticAllocatedVmDisks field to given value.

### HasPersistStaticAllocatedVmDisks

`func (o *CatalogConfiguration) HasPersistStaticAllocatedVmDisks() bool`

HasPersistStaticAllocatedVmDisks returns a boolean if a field has been set.

### GetOfferingId

`func (o *CatalogConfiguration) GetOfferingId() string`

GetOfferingId returns the OfferingId field if non-nil, zero value otherwise.

### GetOfferingIdOk

`func (o *CatalogConfiguration) GetOfferingIdOk() (*string, bool)`

GetOfferingIdOk returns a tuple with the OfferingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferingId

`func (o *CatalogConfiguration) SetOfferingId(v string)`

SetOfferingId sets OfferingId field to given value.


### GetOfferingIdApp

`func (o *CatalogConfiguration) GetOfferingIdApp() string`

GetOfferingIdApp returns the OfferingIdApp field if non-nil, zero value otherwise.

### GetOfferingIdAppOk

`func (o *CatalogConfiguration) GetOfferingIdAppOk() (*string, bool)`

GetOfferingIdAppOk returns a tuple with the OfferingIdApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferingIdApp

`func (o *CatalogConfiguration) SetOfferingIdApp(v string)`

SetOfferingIdApp sets OfferingIdApp field to given value.


### GetOfferingIdDesktop

`func (o *CatalogConfiguration) GetOfferingIdDesktop() string`

GetOfferingIdDesktop returns the OfferingIdDesktop field if non-nil, zero value otherwise.

### GetOfferingIdDesktopOk

`func (o *CatalogConfiguration) GetOfferingIdDesktopOk() (*string, bool)`

GetOfferingIdDesktopOk returns a tuple with the OfferingIdDesktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferingIdDesktop

`func (o *CatalogConfiguration) SetOfferingIdDesktop(v string)`

SetOfferingIdDesktop sets OfferingIdDesktop field to given value.


### GetDeliveryGroupId

`func (o *CatalogConfiguration) GetDeliveryGroupId() string`

GetDeliveryGroupId returns the DeliveryGroupId field if non-nil, zero value otherwise.

### GetDeliveryGroupIdOk

`func (o *CatalogConfiguration) GetDeliveryGroupIdOk() (*string, bool)`

GetDeliveryGroupIdOk returns a tuple with the DeliveryGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroupId

`func (o *CatalogConfiguration) SetDeliveryGroupId(v string)`

SetDeliveryGroupId sets DeliveryGroupId field to given value.


### GetAdvanced

`func (o *CatalogConfiguration) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *CatalogConfiguration) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *CatalogConfiguration) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.


### GetState

`func (o *CatalogConfiguration) GetState() CatalogOverallState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CatalogConfiguration) GetStateOk() (*CatalogOverallState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CatalogConfiguration) SetState(v CatalogOverallState)`

SetState sets State field to given value.


### GetSubState

`func (o *CatalogConfiguration) GetSubState() CatalogOverallSubState`

GetSubState returns the SubState field if non-nil, zero value otherwise.

### GetSubStateOk

`func (o *CatalogConfiguration) GetSubStateOk() (*CatalogOverallSubState, bool)`

GetSubStateOk returns a tuple with the SubState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubState

`func (o *CatalogConfiguration) SetSubState(v CatalogOverallSubState)`

SetSubState sets SubState field to given value.

### HasSubState

`func (o *CatalogConfiguration) HasSubState() bool`

HasSubState returns a boolean if a field has been set.

### GetWarnings

`func (o *CatalogConfiguration) GetWarnings() []CatalogWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CatalogConfiguration) GetWarningsOk() (*[]CatalogWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CatalogConfiguration) SetWarnings(v []CatalogWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CatalogConfiguration) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetStatusMessage

`func (o *CatalogConfiguration) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *CatalogConfiguration) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *CatalogConfiguration) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *CatalogConfiguration) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetExtraInfo

`func (o *CatalogConfiguration) GetExtraInfo() string`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *CatalogConfiguration) GetExtraInfoOk() (*string, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *CatalogConfiguration) SetExtraInfo(v string)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *CatalogConfiguration) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetTransactionId

`func (o *CatalogConfiguration) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CatalogConfiguration) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CatalogConfiguration) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CatalogConfiguration) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetVnetPeeringId

`func (o *CatalogConfiguration) GetVnetPeeringId() string`

GetVnetPeeringId returns the VnetPeeringId field if non-nil, zero value otherwise.

### GetVnetPeeringIdOk

`func (o *CatalogConfiguration) GetVnetPeeringIdOk() (*string, bool)`

GetVnetPeeringIdOk returns a tuple with the VnetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetPeeringId

`func (o *CatalogConfiguration) SetVnetPeeringId(v string)`

SetVnetPeeringId sets VnetPeeringId field to given value.

### HasVnetPeeringId

`func (o *CatalogConfiguration) HasVnetPeeringId() bool`

HasVnetPeeringId returns a boolean if a field has been set.

### GetVnetPeeringName

`func (o *CatalogConfiguration) GetVnetPeeringName() string`

GetVnetPeeringName returns the VnetPeeringName field if non-nil, zero value otherwise.

### GetVnetPeeringNameOk

`func (o *CatalogConfiguration) GetVnetPeeringNameOk() (*string, bool)`

GetVnetPeeringNameOk returns a tuple with the VnetPeeringName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetPeeringName

`func (o *CatalogConfiguration) SetVnetPeeringName(v string)`

SetVnetPeeringName sets VnetPeeringName field to given value.

### HasVnetPeeringName

`func (o *CatalogConfiguration) HasVnetPeeringName() bool`

HasVnetPeeringName returns a boolean if a field has been set.

### GetVpnConnectionId

`func (o *CatalogConfiguration) GetVpnConnectionId() string`

GetVpnConnectionId returns the VpnConnectionId field if non-nil, zero value otherwise.

### GetVpnConnectionIdOk

`func (o *CatalogConfiguration) GetVpnConnectionIdOk() (*string, bool)`

GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionId

`func (o *CatalogConfiguration) SetVpnConnectionId(v string)`

SetVpnConnectionId sets VpnConnectionId field to given value.

### HasVpnConnectionId

`func (o *CatalogConfiguration) HasVpnConnectionId() bool`

HasVpnConnectionId returns a boolean if a field has been set.

### GetVpnConnectionName

`func (o *CatalogConfiguration) GetVpnConnectionName() string`

GetVpnConnectionName returns the VpnConnectionName field if non-nil, zero value otherwise.

### GetVpnConnectionNameOk

`func (o *CatalogConfiguration) GetVpnConnectionNameOk() (*string, bool)`

GetVpnConnectionNameOk returns a tuple with the VpnConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionName

`func (o *CatalogConfiguration) SetVpnConnectionName(v string)`

SetVpnConnectionName sets VpnConnectionName field to given value.

### HasVpnConnectionName

`func (o *CatalogConfiguration) HasVpnConnectionName() bool`

HasVpnConnectionName returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CatalogConfiguration) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CatalogConfiguration) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CatalogConfiguration) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CatalogConfiguration) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionName

`func (o *CatalogConfiguration) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *CatalogConfiguration) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *CatalogConfiguration) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.


### GetResourceGroup

`func (o *CatalogConfiguration) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *CatalogConfiguration) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *CatalogConfiguration) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetVdaResourceGroup

`func (o *CatalogConfiguration) GetVdaResourceGroup() string`

GetVdaResourceGroup returns the VdaResourceGroup field if non-nil, zero value otherwise.

### GetVdaResourceGroupOk

`func (o *CatalogConfiguration) GetVdaResourceGroupOk() (*string, bool)`

GetVdaResourceGroupOk returns a tuple with the VdaResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaResourceGroup

`func (o *CatalogConfiguration) SetVdaResourceGroup(v string)`

SetVdaResourceGroup sets VdaResourceGroup field to given value.

### HasVdaResourceGroup

`func (o *CatalogConfiguration) HasVdaResourceGroup() bool`

HasVdaResourceGroup returns a boolean if a field has been set.

### GetVdaProvisioningSchemeId

`func (o *CatalogConfiguration) GetVdaProvisioningSchemeId() string`

GetVdaProvisioningSchemeId returns the VdaProvisioningSchemeId field if non-nil, zero value otherwise.

### GetVdaProvisioningSchemeIdOk

`func (o *CatalogConfiguration) GetVdaProvisioningSchemeIdOk() (*string, bool)`

GetVdaProvisioningSchemeIdOk returns a tuple with the VdaProvisioningSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaProvisioningSchemeId

`func (o *CatalogConfiguration) SetVdaProvisioningSchemeId(v string)`

SetVdaProvisioningSchemeId sets VdaProvisioningSchemeId field to given value.

### HasVdaProvisioningSchemeId

`func (o *CatalogConfiguration) HasVdaProvisioningSchemeId() bool`

HasVdaProvisioningSchemeId returns a boolean if a field has been set.

### GetAreMcsVdaResourceGroupsUsed

`func (o *CatalogConfiguration) GetAreMcsVdaResourceGroupsUsed() bool`

GetAreMcsVdaResourceGroupsUsed returns the AreMcsVdaResourceGroupsUsed field if non-nil, zero value otherwise.

### GetAreMcsVdaResourceGroupsUsedOk

`func (o *CatalogConfiguration) GetAreMcsVdaResourceGroupsUsedOk() (*bool, bool)`

GetAreMcsVdaResourceGroupsUsedOk returns a tuple with the AreMcsVdaResourceGroupsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreMcsVdaResourceGroupsUsed

`func (o *CatalogConfiguration) SetAreMcsVdaResourceGroupsUsed(v bool)`

SetAreMcsVdaResourceGroupsUsed sets AreMcsVdaResourceGroupsUsed field to given value.

### HasAreMcsVdaResourceGroupsUsed

`func (o *CatalogConfiguration) HasAreMcsVdaResourceGroupsUsed() bool`

HasAreMcsVdaResourceGroupsUsed returns a boolean if a field has been set.

### GetResourceLocationId

`func (o *CatalogConfiguration) GetResourceLocationId() string`

GetResourceLocationId returns the ResourceLocationId field if non-nil, zero value otherwise.

### GetResourceLocationIdOk

`func (o *CatalogConfiguration) GetResourceLocationIdOk() (*string, bool)`

GetResourceLocationIdOk returns a tuple with the ResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocationId

`func (o *CatalogConfiguration) SetResourceLocationId(v string)`

SetResourceLocationId sets ResourceLocationId field to given value.

### HasResourceLocationId

`func (o *CatalogConfiguration) HasResourceLocationId() bool`

HasResourceLocationId returns a boolean if a field has been set.

### GetRegion

`func (o *CatalogConfiguration) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CatalogConfiguration) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CatalogConfiguration) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVNetName

`func (o *CatalogConfiguration) GetVNetName() string`

GetVNetName returns the VNetName field if non-nil, zero value otherwise.

### GetVNetNameOk

`func (o *CatalogConfiguration) GetVNetNameOk() (*string, bool)`

GetVNetNameOk returns a tuple with the VNetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNetName

`func (o *CatalogConfiguration) SetVNetName(v string)`

SetVNetName sets VNetName field to given value.


### GetSubnet

`func (o *CatalogConfiguration) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CatalogConfiguration) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CatalogConfiguration) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *CatalogConfiguration) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetDomainJoined

`func (o *CatalogConfiguration) GetDomainJoined() bool`

GetDomainJoined returns the DomainJoined field if non-nil, zero value otherwise.

### GetDomainJoinedOk

`func (o *CatalogConfiguration) GetDomainJoinedOk() (*bool, bool)`

GetDomainJoinedOk returns a tuple with the DomainJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainJoined

`func (o *CatalogConfiguration) SetDomainJoined(v bool)`

SetDomainJoined sets DomainJoined field to given value.

### HasDomainJoined

`func (o *CatalogConfiguration) HasDomainJoined() bool`

HasDomainJoined returns a boolean if a field has been set.

### GetDomainName

`func (o *CatalogConfiguration) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CatalogConfiguration) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CatalogConfiguration) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CatalogConfiguration) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainOU

`func (o *CatalogConfiguration) GetDomainOU() string`

GetDomainOU returns the DomainOU field if non-nil, zero value otherwise.

### GetDomainOUOk

`func (o *CatalogConfiguration) GetDomainOUOk() (*string, bool)`

GetDomainOUOk returns a tuple with the DomainOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainOU

`func (o *CatalogConfiguration) SetDomainOU(v string)`

SetDomainOU sets DomainOU field to given value.

### HasDomainOU

`func (o *CatalogConfiguration) HasDomainOU() bool`

HasDomainOU returns a boolean if a field has been set.

### GetDomainServiceAccount

`func (o *CatalogConfiguration) GetDomainServiceAccount() string`

GetDomainServiceAccount returns the DomainServiceAccount field if non-nil, zero value otherwise.

### GetDomainServiceAccountOk

`func (o *CatalogConfiguration) GetDomainServiceAccountOk() (*string, bool)`

GetDomainServiceAccountOk returns a tuple with the DomainServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainServiceAccount

`func (o *CatalogConfiguration) SetDomainServiceAccount(v string)`

SetDomainServiceAccount sets DomainServiceAccount field to given value.

### HasDomainServiceAccount

`func (o *CatalogConfiguration) HasDomainServiceAccount() bool`

HasDomainServiceAccount returns a boolean if a field has been set.

### GetVmTypeInstanceType

`func (o *CatalogConfiguration) GetVmTypeInstanceType() string`

GetVmTypeInstanceType returns the VmTypeInstanceType field if non-nil, zero value otherwise.

### GetVmTypeInstanceTypeOk

`func (o *CatalogConfiguration) GetVmTypeInstanceTypeOk() (*string, bool)`

GetVmTypeInstanceTypeOk returns a tuple with the VmTypeInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTypeInstanceType

`func (o *CatalogConfiguration) SetVmTypeInstanceType(v string)`

SetVmTypeInstanceType sets VmTypeInstanceType field to given value.

### HasVmTypeInstanceType

`func (o *CatalogConfiguration) HasVmTypeInstanceType() bool`

HasVmTypeInstanceType returns a boolean if a field has been set.

### GetImageId

`func (o *CatalogConfiguration) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CatalogConfiguration) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CatalogConfiguration) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CatalogConfiguration) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetTemplateImageName

`func (o *CatalogConfiguration) GetTemplateImageName() string`

GetTemplateImageName returns the TemplateImageName field if non-nil, zero value otherwise.

### GetTemplateImageNameOk

`func (o *CatalogConfiguration) GetTemplateImageNameOk() (*string, bool)`

GetTemplateImageNameOk returns a tuple with the TemplateImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateImageName

`func (o *CatalogConfiguration) SetTemplateImageName(v string)`

SetTemplateImageName sets TemplateImageName field to given value.

### HasTemplateImageName

`func (o *CatalogConfiguration) HasTemplateImageName() bool`

HasTemplateImageName returns a boolean if a field has been set.

### GetTemplateImageOs

`func (o *CatalogConfiguration) GetTemplateImageOs() string`

GetTemplateImageOs returns the TemplateImageOs field if non-nil, zero value otherwise.

### GetTemplateImageOsOk

`func (o *CatalogConfiguration) GetTemplateImageOsOk() (*string, bool)`

GetTemplateImageOsOk returns a tuple with the TemplateImageOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateImageOs

`func (o *CatalogConfiguration) SetTemplateImageOs(v string)`

SetTemplateImageOs sets TemplateImageOs field to given value.

### HasTemplateImageOs

`func (o *CatalogConfiguration) HasTemplateImageOs() bool`

HasTemplateImageOs returns a boolean if a field has been set.

### GetCitrixManaged

`func (o *CatalogConfiguration) GetCitrixManaged() bool`

GetCitrixManaged returns the CitrixManaged field if non-nil, zero value otherwise.

### GetCitrixManagedOk

`func (o *CatalogConfiguration) GetCitrixManagedOk() (*bool, bool)`

GetCitrixManagedOk returns a tuple with the CitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManaged

`func (o *CatalogConfiguration) SetCitrixManaged(v bool)`

SetCitrixManaged sets CitrixManaged field to given value.

### HasCitrixManaged

`func (o *CatalogConfiguration) HasCitrixManaged() bool`

HasCitrixManaged returns a boolean if a field has been set.

### GetCspCustomer

`func (o *CatalogConfiguration) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *CatalogConfiguration) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *CatalogConfiguration) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *CatalogConfiguration) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### GetTotalMachinesInCatalog

`func (o *CatalogConfiguration) GetTotalMachinesInCatalog() int32`

GetTotalMachinesInCatalog returns the TotalMachinesInCatalog field if non-nil, zero value otherwise.

### GetTotalMachinesInCatalogOk

`func (o *CatalogConfiguration) GetTotalMachinesInCatalogOk() (*int32, bool)`

GetTotalMachinesInCatalogOk returns a tuple with the TotalMachinesInCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMachinesInCatalog

`func (o *CatalogConfiguration) SetTotalMachinesInCatalog(v int32)`

SetTotalMachinesInCatalog sets TotalMachinesInCatalog field to given value.

### HasTotalMachinesInCatalog

`func (o *CatalogConfiguration) HasTotalMachinesInCatalog() bool`

HasTotalMachinesInCatalog returns a boolean if a field has been set.

### GetWriteBackCacheConfiguration

`func (o *CatalogConfiguration) GetWriteBackCacheConfiguration() WbcConfig`

GetWriteBackCacheConfiguration returns the WriteBackCacheConfiguration field if non-nil, zero value otherwise.

### GetWriteBackCacheConfigurationOk

`func (o *CatalogConfiguration) GetWriteBackCacheConfigurationOk() (*WbcConfig, bool)`

GetWriteBackCacheConfigurationOk returns a tuple with the WriteBackCacheConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheConfiguration

`func (o *CatalogConfiguration) SetWriteBackCacheConfiguration(v WbcConfig)`

SetWriteBackCacheConfiguration sets WriteBackCacheConfiguration field to given value.

### HasWriteBackCacheConfiguration

`func (o *CatalogConfiguration) HasWriteBackCacheConfiguration() bool`

HasWriteBackCacheConfiguration returns a boolean if a field has been set.

### GetTaskCompletionPercentage

`func (o *CatalogConfiguration) GetTaskCompletionPercentage() int32`

GetTaskCompletionPercentage returns the TaskCompletionPercentage field if non-nil, zero value otherwise.

### GetTaskCompletionPercentageOk

`func (o *CatalogConfiguration) GetTaskCompletionPercentageOk() (*int32, bool)`

GetTaskCompletionPercentageOk returns a tuple with the TaskCompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCompletionPercentage

`func (o *CatalogConfiguration) SetTaskCompletionPercentage(v int32)`

SetTaskCompletionPercentage sets TaskCompletionPercentage field to given value.

### HasTaskCompletionPercentage

`func (o *CatalogConfiguration) HasTaskCompletionPercentage() bool`

HasTaskCompletionPercentage returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CatalogConfiguration) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CatalogConfiguration) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CatalogConfiguration) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *CatalogConfiguration) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetLastBackupTime

`func (o *CatalogConfiguration) GetLastBackupTime() time.Time`

GetLastBackupTime returns the LastBackupTime field if non-nil, zero value otherwise.

### GetLastBackupTimeOk

`func (o *CatalogConfiguration) GetLastBackupTimeOk() (*time.Time, bool)`

GetLastBackupTimeOk returns a tuple with the LastBackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupTime

`func (o *CatalogConfiguration) SetLastBackupTime(v time.Time)`

SetLastBackupTime sets LastBackupTime field to given value.

### HasLastBackupTime

`func (o *CatalogConfiguration) HasLastBackupTime() bool`

HasLastBackupTime returns a boolean if a field has been set.

### GetIsRemotePcCatalog

`func (o *CatalogConfiguration) GetIsRemotePcCatalog() bool`

GetIsRemotePcCatalog returns the IsRemotePcCatalog field if non-nil, zero value otherwise.

### GetIsRemotePcCatalogOk

`func (o *CatalogConfiguration) GetIsRemotePcCatalogOk() (*bool, bool)`

GetIsRemotePcCatalogOk returns a tuple with the IsRemotePcCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePcCatalog

`func (o *CatalogConfiguration) SetIsRemotePcCatalog(v bool)`

SetIsRemotePcCatalog sets IsRemotePcCatalog field to given value.

### HasIsRemotePcCatalog

`func (o *CatalogConfiguration) HasIsRemotePcCatalog() bool`

HasIsRemotePcCatalog returns a boolean if a field has been set.

### GetIsAzureAdJoined

`func (o *CatalogConfiguration) GetIsAzureAdJoined() bool`

GetIsAzureAdJoined returns the IsAzureAdJoined field if non-nil, zero value otherwise.

### GetIsAzureAdJoinedOk

`func (o *CatalogConfiguration) GetIsAzureAdJoinedOk() (*bool, bool)`

GetIsAzureAdJoinedOk returns a tuple with the IsAzureAdJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzureAdJoined

`func (o *CatalogConfiguration) SetIsAzureAdJoined(v bool)`

SetIsAzureAdJoined sets IsAzureAdJoined field to given value.

### HasIsAzureAdJoined

`func (o *CatalogConfiguration) HasIsAzureAdJoined() bool`

HasIsAzureAdJoined returns a boolean if a field has been set.

### GetIsSecureBrowserCatalog

`func (o *CatalogConfiguration) GetIsSecureBrowserCatalog() bool`

GetIsSecureBrowserCatalog returns the IsSecureBrowserCatalog field if non-nil, zero value otherwise.

### GetIsSecureBrowserCatalogOk

`func (o *CatalogConfiguration) GetIsSecureBrowserCatalogOk() (*bool, bool)`

GetIsSecureBrowserCatalogOk returns a tuple with the IsSecureBrowserCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureBrowserCatalog

`func (o *CatalogConfiguration) SetIsSecureBrowserCatalog(v bool)`

SetIsSecureBrowserCatalog sets IsSecureBrowserCatalog field to given value.

### HasIsSecureBrowserCatalog

`func (o *CatalogConfiguration) HasIsSecureBrowserCatalog() bool`

HasIsSecureBrowserCatalog returns a boolean if a field has been set.

### GetOrganizationalUnits

`func (o *CatalogConfiguration) GetOrganizationalUnits() []RemotePCEnrollmentScopeResponseModel`

GetOrganizationalUnits returns the OrganizationalUnits field if non-nil, zero value otherwise.

### GetOrganizationalUnitsOk

`func (o *CatalogConfiguration) GetOrganizationalUnitsOk() (*[]RemotePCEnrollmentScopeResponseModel, bool)`

GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnits

`func (o *CatalogConfiguration) SetOrganizationalUnits(v []RemotePCEnrollmentScopeResponseModel)`

SetOrganizationalUnits sets OrganizationalUnits field to given value.

### HasOrganizationalUnits

`func (o *CatalogConfiguration) HasOrganizationalUnits() bool`

HasOrganizationalUnits returns a boolean if a field has been set.

### GetSupportsHibernation

`func (o *CatalogConfiguration) GetSupportsHibernation() bool`

GetSupportsHibernation returns the SupportsHibernation field if non-nil, zero value otherwise.

### GetSupportsHibernationOk

`func (o *CatalogConfiguration) GetSupportsHibernationOk() (*bool, bool)`

GetSupportsHibernationOk returns a tuple with the SupportsHibernation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsHibernation

`func (o *CatalogConfiguration) SetSupportsHibernation(v bool)`

SetSupportsHibernation sets SupportsHibernation field to given value.

### HasSupportsHibernation

`func (o *CatalogConfiguration) HasSupportsHibernation() bool`

HasSupportsHibernation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


