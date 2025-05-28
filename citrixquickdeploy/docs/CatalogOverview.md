# CatalogOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureSubscriptionId** | Pointer to **string** | ID of the Azure Subscription linked to this catalog | [optional] [readonly] 
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

### NewCatalogOverview

`func NewCatalogOverview(id string, name string, offeringId string, offeringIdApp string, offeringIdDesktop string, deliveryGroupId string, advanced bool, state CatalogOverallState, subscriptionName string, resourceGroup string, region string, vNetName string, ) *CatalogOverview`

NewCatalogOverview instantiates a new CatalogOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOverviewWithDefaults

`func NewCatalogOverviewWithDefaults() *CatalogOverview`

NewCatalogOverviewWithDefaults instantiates a new CatalogOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureSubscriptionId

`func (o *CatalogOverview) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *CatalogOverview) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *CatalogOverview) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *CatalogOverview) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.

### GetId

`func (o *CatalogOverview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogOverview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogOverview) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CatalogOverview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogOverview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogOverview) SetName(v string)`

SetName sets Name field to given value.


### GetSessionSupport

`func (o *CatalogOverview) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *CatalogOverview) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *CatalogOverview) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *CatalogOverview) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.

### GetAllocationType

`func (o *CatalogOverview) GetAllocationType() CatalogAllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *CatalogOverview) GetAllocationTypeOk() (*CatalogAllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *CatalogOverview) SetAllocationType(v CatalogAllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *CatalogOverview) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetPersistStaticAllocatedVmDisks

`func (o *CatalogOverview) GetPersistStaticAllocatedVmDisks() bool`

GetPersistStaticAllocatedVmDisks returns the PersistStaticAllocatedVmDisks field if non-nil, zero value otherwise.

### GetPersistStaticAllocatedVmDisksOk

`func (o *CatalogOverview) GetPersistStaticAllocatedVmDisksOk() (*bool, bool)`

GetPersistStaticAllocatedVmDisksOk returns a tuple with the PersistStaticAllocatedVmDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistStaticAllocatedVmDisks

`func (o *CatalogOverview) SetPersistStaticAllocatedVmDisks(v bool)`

SetPersistStaticAllocatedVmDisks sets PersistStaticAllocatedVmDisks field to given value.

### HasPersistStaticAllocatedVmDisks

`func (o *CatalogOverview) HasPersistStaticAllocatedVmDisks() bool`

HasPersistStaticAllocatedVmDisks returns a boolean if a field has been set.

### GetOfferingId

`func (o *CatalogOverview) GetOfferingId() string`

GetOfferingId returns the OfferingId field if non-nil, zero value otherwise.

### GetOfferingIdOk

`func (o *CatalogOverview) GetOfferingIdOk() (*string, bool)`

GetOfferingIdOk returns a tuple with the OfferingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferingId

`func (o *CatalogOverview) SetOfferingId(v string)`

SetOfferingId sets OfferingId field to given value.


### GetOfferingIdApp

`func (o *CatalogOverview) GetOfferingIdApp() string`

GetOfferingIdApp returns the OfferingIdApp field if non-nil, zero value otherwise.

### GetOfferingIdAppOk

`func (o *CatalogOverview) GetOfferingIdAppOk() (*string, bool)`

GetOfferingIdAppOk returns a tuple with the OfferingIdApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferingIdApp

`func (o *CatalogOverview) SetOfferingIdApp(v string)`

SetOfferingIdApp sets OfferingIdApp field to given value.


### GetOfferingIdDesktop

`func (o *CatalogOverview) GetOfferingIdDesktop() string`

GetOfferingIdDesktop returns the OfferingIdDesktop field if non-nil, zero value otherwise.

### GetOfferingIdDesktopOk

`func (o *CatalogOverview) GetOfferingIdDesktopOk() (*string, bool)`

GetOfferingIdDesktopOk returns a tuple with the OfferingIdDesktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferingIdDesktop

`func (o *CatalogOverview) SetOfferingIdDesktop(v string)`

SetOfferingIdDesktop sets OfferingIdDesktop field to given value.


### GetDeliveryGroupId

`func (o *CatalogOverview) GetDeliveryGroupId() string`

GetDeliveryGroupId returns the DeliveryGroupId field if non-nil, zero value otherwise.

### GetDeliveryGroupIdOk

`func (o *CatalogOverview) GetDeliveryGroupIdOk() (*string, bool)`

GetDeliveryGroupIdOk returns a tuple with the DeliveryGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroupId

`func (o *CatalogOverview) SetDeliveryGroupId(v string)`

SetDeliveryGroupId sets DeliveryGroupId field to given value.


### GetAdvanced

`func (o *CatalogOverview) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *CatalogOverview) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *CatalogOverview) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.


### GetState

`func (o *CatalogOverview) GetState() CatalogOverallState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CatalogOverview) GetStateOk() (*CatalogOverallState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CatalogOverview) SetState(v CatalogOverallState)`

SetState sets State field to given value.


### GetSubState

`func (o *CatalogOverview) GetSubState() CatalogOverallSubState`

GetSubState returns the SubState field if non-nil, zero value otherwise.

### GetSubStateOk

`func (o *CatalogOverview) GetSubStateOk() (*CatalogOverallSubState, bool)`

GetSubStateOk returns a tuple with the SubState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubState

`func (o *CatalogOverview) SetSubState(v CatalogOverallSubState)`

SetSubState sets SubState field to given value.

### HasSubState

`func (o *CatalogOverview) HasSubState() bool`

HasSubState returns a boolean if a field has been set.

### GetWarnings

`func (o *CatalogOverview) GetWarnings() []CatalogWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CatalogOverview) GetWarningsOk() (*[]CatalogWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CatalogOverview) SetWarnings(v []CatalogWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CatalogOverview) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetStatusMessage

`func (o *CatalogOverview) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *CatalogOverview) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *CatalogOverview) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *CatalogOverview) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetExtraInfo

`func (o *CatalogOverview) GetExtraInfo() string`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *CatalogOverview) GetExtraInfoOk() (*string, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *CatalogOverview) SetExtraInfo(v string)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *CatalogOverview) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetTransactionId

`func (o *CatalogOverview) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CatalogOverview) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CatalogOverview) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CatalogOverview) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetVnetPeeringId

`func (o *CatalogOverview) GetVnetPeeringId() string`

GetVnetPeeringId returns the VnetPeeringId field if non-nil, zero value otherwise.

### GetVnetPeeringIdOk

`func (o *CatalogOverview) GetVnetPeeringIdOk() (*string, bool)`

GetVnetPeeringIdOk returns a tuple with the VnetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetPeeringId

`func (o *CatalogOverview) SetVnetPeeringId(v string)`

SetVnetPeeringId sets VnetPeeringId field to given value.

### HasVnetPeeringId

`func (o *CatalogOverview) HasVnetPeeringId() bool`

HasVnetPeeringId returns a boolean if a field has been set.

### GetVnetPeeringName

`func (o *CatalogOverview) GetVnetPeeringName() string`

GetVnetPeeringName returns the VnetPeeringName field if non-nil, zero value otherwise.

### GetVnetPeeringNameOk

`func (o *CatalogOverview) GetVnetPeeringNameOk() (*string, bool)`

GetVnetPeeringNameOk returns a tuple with the VnetPeeringName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetPeeringName

`func (o *CatalogOverview) SetVnetPeeringName(v string)`

SetVnetPeeringName sets VnetPeeringName field to given value.

### HasVnetPeeringName

`func (o *CatalogOverview) HasVnetPeeringName() bool`

HasVnetPeeringName returns a boolean if a field has been set.

### GetVpnConnectionId

`func (o *CatalogOverview) GetVpnConnectionId() string`

GetVpnConnectionId returns the VpnConnectionId field if non-nil, zero value otherwise.

### GetVpnConnectionIdOk

`func (o *CatalogOverview) GetVpnConnectionIdOk() (*string, bool)`

GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionId

`func (o *CatalogOverview) SetVpnConnectionId(v string)`

SetVpnConnectionId sets VpnConnectionId field to given value.

### HasVpnConnectionId

`func (o *CatalogOverview) HasVpnConnectionId() bool`

HasVpnConnectionId returns a boolean if a field has been set.

### GetVpnConnectionName

`func (o *CatalogOverview) GetVpnConnectionName() string`

GetVpnConnectionName returns the VpnConnectionName field if non-nil, zero value otherwise.

### GetVpnConnectionNameOk

`func (o *CatalogOverview) GetVpnConnectionNameOk() (*string, bool)`

GetVpnConnectionNameOk returns a tuple with the VpnConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionName

`func (o *CatalogOverview) SetVpnConnectionName(v string)`

SetVpnConnectionName sets VpnConnectionName field to given value.

### HasVpnConnectionName

`func (o *CatalogOverview) HasVpnConnectionName() bool`

HasVpnConnectionName returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CatalogOverview) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CatalogOverview) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CatalogOverview) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CatalogOverview) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionName

`func (o *CatalogOverview) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *CatalogOverview) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *CatalogOverview) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.


### GetResourceGroup

`func (o *CatalogOverview) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *CatalogOverview) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *CatalogOverview) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetVdaResourceGroup

`func (o *CatalogOverview) GetVdaResourceGroup() string`

GetVdaResourceGroup returns the VdaResourceGroup field if non-nil, zero value otherwise.

### GetVdaResourceGroupOk

`func (o *CatalogOverview) GetVdaResourceGroupOk() (*string, bool)`

GetVdaResourceGroupOk returns a tuple with the VdaResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaResourceGroup

`func (o *CatalogOverview) SetVdaResourceGroup(v string)`

SetVdaResourceGroup sets VdaResourceGroup field to given value.

### HasVdaResourceGroup

`func (o *CatalogOverview) HasVdaResourceGroup() bool`

HasVdaResourceGroup returns a boolean if a field has been set.

### GetVdaProvisioningSchemeId

`func (o *CatalogOverview) GetVdaProvisioningSchemeId() string`

GetVdaProvisioningSchemeId returns the VdaProvisioningSchemeId field if non-nil, zero value otherwise.

### GetVdaProvisioningSchemeIdOk

`func (o *CatalogOverview) GetVdaProvisioningSchemeIdOk() (*string, bool)`

GetVdaProvisioningSchemeIdOk returns a tuple with the VdaProvisioningSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaProvisioningSchemeId

`func (o *CatalogOverview) SetVdaProvisioningSchemeId(v string)`

SetVdaProvisioningSchemeId sets VdaProvisioningSchemeId field to given value.

### HasVdaProvisioningSchemeId

`func (o *CatalogOverview) HasVdaProvisioningSchemeId() bool`

HasVdaProvisioningSchemeId returns a boolean if a field has been set.

### GetAreMcsVdaResourceGroupsUsed

`func (o *CatalogOverview) GetAreMcsVdaResourceGroupsUsed() bool`

GetAreMcsVdaResourceGroupsUsed returns the AreMcsVdaResourceGroupsUsed field if non-nil, zero value otherwise.

### GetAreMcsVdaResourceGroupsUsedOk

`func (o *CatalogOverview) GetAreMcsVdaResourceGroupsUsedOk() (*bool, bool)`

GetAreMcsVdaResourceGroupsUsedOk returns a tuple with the AreMcsVdaResourceGroupsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreMcsVdaResourceGroupsUsed

`func (o *CatalogOverview) SetAreMcsVdaResourceGroupsUsed(v bool)`

SetAreMcsVdaResourceGroupsUsed sets AreMcsVdaResourceGroupsUsed field to given value.

### HasAreMcsVdaResourceGroupsUsed

`func (o *CatalogOverview) HasAreMcsVdaResourceGroupsUsed() bool`

HasAreMcsVdaResourceGroupsUsed returns a boolean if a field has been set.

### GetResourceLocationId

`func (o *CatalogOverview) GetResourceLocationId() string`

GetResourceLocationId returns the ResourceLocationId field if non-nil, zero value otherwise.

### GetResourceLocationIdOk

`func (o *CatalogOverview) GetResourceLocationIdOk() (*string, bool)`

GetResourceLocationIdOk returns a tuple with the ResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocationId

`func (o *CatalogOverview) SetResourceLocationId(v string)`

SetResourceLocationId sets ResourceLocationId field to given value.

### HasResourceLocationId

`func (o *CatalogOverview) HasResourceLocationId() bool`

HasResourceLocationId returns a boolean if a field has been set.

### GetRegion

`func (o *CatalogOverview) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CatalogOverview) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CatalogOverview) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVNetName

`func (o *CatalogOverview) GetVNetName() string`

GetVNetName returns the VNetName field if non-nil, zero value otherwise.

### GetVNetNameOk

`func (o *CatalogOverview) GetVNetNameOk() (*string, bool)`

GetVNetNameOk returns a tuple with the VNetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNetName

`func (o *CatalogOverview) SetVNetName(v string)`

SetVNetName sets VNetName field to given value.


### GetSubnet

`func (o *CatalogOverview) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CatalogOverview) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CatalogOverview) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *CatalogOverview) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetDomainJoined

`func (o *CatalogOverview) GetDomainJoined() bool`

GetDomainJoined returns the DomainJoined field if non-nil, zero value otherwise.

### GetDomainJoinedOk

`func (o *CatalogOverview) GetDomainJoinedOk() (*bool, bool)`

GetDomainJoinedOk returns a tuple with the DomainJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainJoined

`func (o *CatalogOverview) SetDomainJoined(v bool)`

SetDomainJoined sets DomainJoined field to given value.

### HasDomainJoined

`func (o *CatalogOverview) HasDomainJoined() bool`

HasDomainJoined returns a boolean if a field has been set.

### GetDomainName

`func (o *CatalogOverview) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CatalogOverview) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CatalogOverview) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CatalogOverview) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainOU

`func (o *CatalogOverview) GetDomainOU() string`

GetDomainOU returns the DomainOU field if non-nil, zero value otherwise.

### GetDomainOUOk

`func (o *CatalogOverview) GetDomainOUOk() (*string, bool)`

GetDomainOUOk returns a tuple with the DomainOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainOU

`func (o *CatalogOverview) SetDomainOU(v string)`

SetDomainOU sets DomainOU field to given value.

### HasDomainOU

`func (o *CatalogOverview) HasDomainOU() bool`

HasDomainOU returns a boolean if a field has been set.

### GetDomainServiceAccount

`func (o *CatalogOverview) GetDomainServiceAccount() string`

GetDomainServiceAccount returns the DomainServiceAccount field if non-nil, zero value otherwise.

### GetDomainServiceAccountOk

`func (o *CatalogOverview) GetDomainServiceAccountOk() (*string, bool)`

GetDomainServiceAccountOk returns a tuple with the DomainServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainServiceAccount

`func (o *CatalogOverview) SetDomainServiceAccount(v string)`

SetDomainServiceAccount sets DomainServiceAccount field to given value.

### HasDomainServiceAccount

`func (o *CatalogOverview) HasDomainServiceAccount() bool`

HasDomainServiceAccount returns a boolean if a field has been set.

### GetVmTypeInstanceType

`func (o *CatalogOverview) GetVmTypeInstanceType() string`

GetVmTypeInstanceType returns the VmTypeInstanceType field if non-nil, zero value otherwise.

### GetVmTypeInstanceTypeOk

`func (o *CatalogOverview) GetVmTypeInstanceTypeOk() (*string, bool)`

GetVmTypeInstanceTypeOk returns a tuple with the VmTypeInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTypeInstanceType

`func (o *CatalogOverview) SetVmTypeInstanceType(v string)`

SetVmTypeInstanceType sets VmTypeInstanceType field to given value.

### HasVmTypeInstanceType

`func (o *CatalogOverview) HasVmTypeInstanceType() bool`

HasVmTypeInstanceType returns a boolean if a field has been set.

### GetImageId

`func (o *CatalogOverview) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CatalogOverview) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CatalogOverview) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CatalogOverview) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetTemplateImageName

`func (o *CatalogOverview) GetTemplateImageName() string`

GetTemplateImageName returns the TemplateImageName field if non-nil, zero value otherwise.

### GetTemplateImageNameOk

`func (o *CatalogOverview) GetTemplateImageNameOk() (*string, bool)`

GetTemplateImageNameOk returns a tuple with the TemplateImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateImageName

`func (o *CatalogOverview) SetTemplateImageName(v string)`

SetTemplateImageName sets TemplateImageName field to given value.

### HasTemplateImageName

`func (o *CatalogOverview) HasTemplateImageName() bool`

HasTemplateImageName returns a boolean if a field has been set.

### GetTemplateImageOs

`func (o *CatalogOverview) GetTemplateImageOs() string`

GetTemplateImageOs returns the TemplateImageOs field if non-nil, zero value otherwise.

### GetTemplateImageOsOk

`func (o *CatalogOverview) GetTemplateImageOsOk() (*string, bool)`

GetTemplateImageOsOk returns a tuple with the TemplateImageOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateImageOs

`func (o *CatalogOverview) SetTemplateImageOs(v string)`

SetTemplateImageOs sets TemplateImageOs field to given value.

### HasTemplateImageOs

`func (o *CatalogOverview) HasTemplateImageOs() bool`

HasTemplateImageOs returns a boolean if a field has been set.

### GetCitrixManaged

`func (o *CatalogOverview) GetCitrixManaged() bool`

GetCitrixManaged returns the CitrixManaged field if non-nil, zero value otherwise.

### GetCitrixManagedOk

`func (o *CatalogOverview) GetCitrixManagedOk() (*bool, bool)`

GetCitrixManagedOk returns a tuple with the CitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManaged

`func (o *CatalogOverview) SetCitrixManaged(v bool)`

SetCitrixManaged sets CitrixManaged field to given value.

### HasCitrixManaged

`func (o *CatalogOverview) HasCitrixManaged() bool`

HasCitrixManaged returns a boolean if a field has been set.

### GetCspCustomer

`func (o *CatalogOverview) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *CatalogOverview) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *CatalogOverview) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *CatalogOverview) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### GetTotalMachinesInCatalog

`func (o *CatalogOverview) GetTotalMachinesInCatalog() int32`

GetTotalMachinesInCatalog returns the TotalMachinesInCatalog field if non-nil, zero value otherwise.

### GetTotalMachinesInCatalogOk

`func (o *CatalogOverview) GetTotalMachinesInCatalogOk() (*int32, bool)`

GetTotalMachinesInCatalogOk returns a tuple with the TotalMachinesInCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMachinesInCatalog

`func (o *CatalogOverview) SetTotalMachinesInCatalog(v int32)`

SetTotalMachinesInCatalog sets TotalMachinesInCatalog field to given value.

### HasTotalMachinesInCatalog

`func (o *CatalogOverview) HasTotalMachinesInCatalog() bool`

HasTotalMachinesInCatalog returns a boolean if a field has been set.

### GetWriteBackCacheConfiguration

`func (o *CatalogOverview) GetWriteBackCacheConfiguration() WbcConfig`

GetWriteBackCacheConfiguration returns the WriteBackCacheConfiguration field if non-nil, zero value otherwise.

### GetWriteBackCacheConfigurationOk

`func (o *CatalogOverview) GetWriteBackCacheConfigurationOk() (*WbcConfig, bool)`

GetWriteBackCacheConfigurationOk returns a tuple with the WriteBackCacheConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheConfiguration

`func (o *CatalogOverview) SetWriteBackCacheConfiguration(v WbcConfig)`

SetWriteBackCacheConfiguration sets WriteBackCacheConfiguration field to given value.

### HasWriteBackCacheConfiguration

`func (o *CatalogOverview) HasWriteBackCacheConfiguration() bool`

HasWriteBackCacheConfiguration returns a boolean if a field has been set.

### GetTaskCompletionPercentage

`func (o *CatalogOverview) GetTaskCompletionPercentage() int32`

GetTaskCompletionPercentage returns the TaskCompletionPercentage field if non-nil, zero value otherwise.

### GetTaskCompletionPercentageOk

`func (o *CatalogOverview) GetTaskCompletionPercentageOk() (*int32, bool)`

GetTaskCompletionPercentageOk returns a tuple with the TaskCompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCompletionPercentage

`func (o *CatalogOverview) SetTaskCompletionPercentage(v int32)`

SetTaskCompletionPercentage sets TaskCompletionPercentage field to given value.

### HasTaskCompletionPercentage

`func (o *CatalogOverview) HasTaskCompletionPercentage() bool`

HasTaskCompletionPercentage returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CatalogOverview) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CatalogOverview) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CatalogOverview) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *CatalogOverview) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetLastBackupTime

`func (o *CatalogOverview) GetLastBackupTime() time.Time`

GetLastBackupTime returns the LastBackupTime field if non-nil, zero value otherwise.

### GetLastBackupTimeOk

`func (o *CatalogOverview) GetLastBackupTimeOk() (*time.Time, bool)`

GetLastBackupTimeOk returns a tuple with the LastBackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupTime

`func (o *CatalogOverview) SetLastBackupTime(v time.Time)`

SetLastBackupTime sets LastBackupTime field to given value.

### HasLastBackupTime

`func (o *CatalogOverview) HasLastBackupTime() bool`

HasLastBackupTime returns a boolean if a field has been set.

### GetIsRemotePcCatalog

`func (o *CatalogOverview) GetIsRemotePcCatalog() bool`

GetIsRemotePcCatalog returns the IsRemotePcCatalog field if non-nil, zero value otherwise.

### GetIsRemotePcCatalogOk

`func (o *CatalogOverview) GetIsRemotePcCatalogOk() (*bool, bool)`

GetIsRemotePcCatalogOk returns a tuple with the IsRemotePcCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotePcCatalog

`func (o *CatalogOverview) SetIsRemotePcCatalog(v bool)`

SetIsRemotePcCatalog sets IsRemotePcCatalog field to given value.

### HasIsRemotePcCatalog

`func (o *CatalogOverview) HasIsRemotePcCatalog() bool`

HasIsRemotePcCatalog returns a boolean if a field has been set.

### GetIsAzureAdJoined

`func (o *CatalogOverview) GetIsAzureAdJoined() bool`

GetIsAzureAdJoined returns the IsAzureAdJoined field if non-nil, zero value otherwise.

### GetIsAzureAdJoinedOk

`func (o *CatalogOverview) GetIsAzureAdJoinedOk() (*bool, bool)`

GetIsAzureAdJoinedOk returns a tuple with the IsAzureAdJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzureAdJoined

`func (o *CatalogOverview) SetIsAzureAdJoined(v bool)`

SetIsAzureAdJoined sets IsAzureAdJoined field to given value.

### HasIsAzureAdJoined

`func (o *CatalogOverview) HasIsAzureAdJoined() bool`

HasIsAzureAdJoined returns a boolean if a field has been set.

### GetIsSecureBrowserCatalog

`func (o *CatalogOverview) GetIsSecureBrowserCatalog() bool`

GetIsSecureBrowserCatalog returns the IsSecureBrowserCatalog field if non-nil, zero value otherwise.

### GetIsSecureBrowserCatalogOk

`func (o *CatalogOverview) GetIsSecureBrowserCatalogOk() (*bool, bool)`

GetIsSecureBrowserCatalogOk returns a tuple with the IsSecureBrowserCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureBrowserCatalog

`func (o *CatalogOverview) SetIsSecureBrowserCatalog(v bool)`

SetIsSecureBrowserCatalog sets IsSecureBrowserCatalog field to given value.

### HasIsSecureBrowserCatalog

`func (o *CatalogOverview) HasIsSecureBrowserCatalog() bool`

HasIsSecureBrowserCatalog returns a boolean if a field has been set.

### GetOrganizationalUnits

`func (o *CatalogOverview) GetOrganizationalUnits() []RemotePCEnrollmentScopeResponseModel`

GetOrganizationalUnits returns the OrganizationalUnits field if non-nil, zero value otherwise.

### GetOrganizationalUnitsOk

`func (o *CatalogOverview) GetOrganizationalUnitsOk() (*[]RemotePCEnrollmentScopeResponseModel, bool)`

GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnits

`func (o *CatalogOverview) SetOrganizationalUnits(v []RemotePCEnrollmentScopeResponseModel)`

SetOrganizationalUnits sets OrganizationalUnits field to given value.

### HasOrganizationalUnits

`func (o *CatalogOverview) HasOrganizationalUnits() bool`

HasOrganizationalUnits returns a boolean if a field has been set.

### GetSupportsHibernation

`func (o *CatalogOverview) GetSupportsHibernation() bool`

GetSupportsHibernation returns the SupportsHibernation field if non-nil, zero value otherwise.

### GetSupportsHibernationOk

`func (o *CatalogOverview) GetSupportsHibernationOk() (*bool, bool)`

GetSupportsHibernationOk returns a tuple with the SupportsHibernation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsHibernation

`func (o *CatalogOverview) SetSupportsHibernation(v bool)`

SetSupportsHibernation sets SupportsHibernation field to given value.

### HasSupportsHibernation

`func (o *CatalogOverview) HasSupportsHibernation() bool`

HasSupportsHibernation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


