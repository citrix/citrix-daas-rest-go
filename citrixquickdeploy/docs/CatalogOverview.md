# CatalogOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureSubscriptionId** | Pointer to **NullableString** | ID of the Azure Subscription linked to this catalog | [optional] [readonly] 
**Id** | **string** | Unique identifier of the catalog | 
**Name** | **string** | User configured name | 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) | Quantity of sessions supported per-machine. | [optional] 
**AllocationType** | Pointer to [**CatalogAllocationType**](CatalogAllocationType.md) | Indicates the manner in which machines are allocated to users | [optional] 
**PersistStaticAllocatedVmDisks** | Pointer to **NullableBool** | Indicates if catalogs that use statically allocated machines will have the disk contents persisted after shutdown | [optional] 
**OfferingId** | **string** | The offeringId for the catalog to be used in Cloud Library operations | 
**OfferingIdApp** | **string** | The Application offeringId for the catalog to be used in Cloud Library operations | 
**OfferingIdDesktop** | **string** | The Desktop offeringId for the catalog to be used in Cloud Library operations | 
**DeliveryGroupId** | **string** | The Delivery Group Id | 
**Advanced** | **bool** | Indicates if this is an advanced object | 
**State** | [**CatalogOverallState**](CatalogOverallState.md) | Status of the catalog | 
**SubState** | Pointer to [**NullableCatalogOverallSubState**](CatalogOverallSubState.md) | Sub Status of the catalog | [optional] [readonly] 
**Warnings** | Pointer to [**[]CatalogWarning**](CatalogWarning.md) |  | [optional] 
**StatusMessage** | Pointer to **NullableString** | Current status of the catalog | [optional] 
**ExtraInfo** | Pointer to **NullableString** | The string to displayed in UI for extra information | [optional] 
**TransactionId** | Pointer to **NullableString** | The transaction id of the catalog deployment | [optional] 
**VnetPeeringId** | Pointer to **NullableString** | ID of the Vnet Peering associated with the catalog | [optional] 
**VnetPeeringName** | Pointer to **NullableString** | Name of the Vnet Peering associated with the catalog | [optional] 
**VpnConnectionId** | Pointer to **NullableString** | ID of the Vpn Connection associated with the catalog | [optional] 
**VpnConnectionName** | Pointer to **NullableString** | Name of the Vpn Connection associated with the catalog | [optional] 
**SubscriptionId** | Pointer to **NullableString** | Id of the Subscription that catalog VMs will be deployed to | [optional] 
**SubscriptionName** | **string** | Name of the Subscription that catalog VMs will be deployed to | 
**ResourceGroup** | **string** | Name of the resource group used for this catalog | 
**TemplateSpecResourceGroup** | Pointer to **NullableString** | Resource Group for the Template Specs used by the catalog | [optional] [readonly] 
**VdaResourceGroup** | Pointer to **NullableString** | The resource group for the VDAs in Azure | [optional] 
**VdaProvisioningSchemeId** | Pointer to **NullableString** | The resource groups for the VDAs in Azure | [optional] 
**AreMcsVdaResourceGroupsUsed** | Pointer to **NullableBool** | The resource groups for the VDAs in Azure | [optional] 
**ResourceLocationId** | Pointer to **NullableString** | ID of the Resource Location associated with the catalog | [optional] 
**Region** | **string** | Azure region where VMs are deployed for this catalog | 
**VNetName** | **string** | Name of the vnet assigned to the catalog | [readonly] 
**Subnet** | Pointer to **NullableString** | The subnet that is associated with the catalog&#39;s VNet | [optional] 
**DomainJoined** | Pointer to **NullableBool** | The flag to indicate if the catalog is joined with customer domain | [optional] 
**DomainName** | Pointer to **NullableString** | Name of the domain that the catalog&#39;s VMs will join | [optional] 
**DomainOU** | Pointer to **NullableString** | OU of the domain we are joining | [optional] 
**DomainServiceAccount** | Pointer to **NullableString** | Name of the service account that will perform domain join opperations | [optional] 
**ServiceAccountUid** | Pointer to **NullableString** | Service account to associate to the IdentityPool.  Used for Pure Entra ID joined catalogs. | [optional] [readonly] 
**VmTypeInstanceType** | Pointer to **NullableString** | Type of the VM machines used to create VDAs | [optional] [readonly] 
**ImageId** | Pointer to **NullableString** | ID of the image that is used by the catalog | [optional] 
**TemplateImageName** | Pointer to **NullableString** | Name of the template image that we are using for this catalog | [optional] 
**TemplateImageOs** | Pointer to **NullableString** | Os type of the template image that we are using for this catalog | [optional] 
**CitrixManaged** | Pointer to **NullableBool** | Indicates that partner-tenant relationship exists if not null | [optional] 
**CspCustomer** | Pointer to **NullableString** | Indicates that partner-tenant relationship exists if not null | [optional] 
**TotalMachinesInCatalog** | Pointer to **NullableInt32** | Maximum number of machines assigned to the catalog | [optional] 
**NumOfUsers** | Pointer to **NullableInt32** | The number of users that the catalog should support (user defined value) | [optional] 
**MaxNumOfUsers** | Pointer to **NullableInt32** | The maximum number of users supported by the catalog (TotalMachines * Sessions per VM) | [optional] 
**WriteBackCacheConfiguration** | Pointer to [**NullableWbcConfig**](WbcConfig.md) | Indicates whether or not write back cache is enabled for the VMs created from this provisioning scheme. | [optional] 
**TaskCompletionPercentage** | Pointer to **NullableInt32** | Percentage complete the current task being performed on the catalog is at | [optional] 
**LastModifiedTime** | Pointer to **NullableTime** | Last time when the catalog was modified | [optional] 
**LastBackupTime** | Pointer to **NullableTime** | Last backup time for the catalog&#39;s VDAs | [optional] [readonly] 
**IsRemotePcCatalog** | Pointer to **bool** | Indicates if this is a remote pc catalog | [optional] 
**IsAzureAdJoined** | Pointer to **bool** | Indicates if the machines in the catalog will be Azure AD joined | [optional] 
**IsSecureBrowserCatalog** | Pointer to **bool** | Indicates if the catalog is for Secure Browser service | [optional] [readonly] 
**OrganizationalUnits** | Pointer to [**[]RemotePCEnrollmentScopeResponseModel**](RemotePCEnrollmentScopeResponseModel.md) | List of OUs for remote pc | [optional] 
**SupportsHibernation** | Pointer to **bool** | Indicates whether machines in catalog support hibernation | [optional] 
**EnableAcceleratedNetworking** | Pointer to **bool** | Specifies whether to enable accelerated networking on the VM NIC | [optional] 
**EnableEncryptionAtHost** | Pointer to **bool** | Indicates whether encryption at the host level is enabled. | [optional] 
**CatalogType** | Pointer to [**CatalogType**](CatalogType.md) | The type of catalog | [optional] [readonly] 

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

### SetAzureSubscriptionIdNil

`func (o *CatalogOverview) SetAzureSubscriptionIdNil(b bool)`

 SetAzureSubscriptionIdNil sets the value for AzureSubscriptionId to be an explicit nil

### UnsetAzureSubscriptionId
`func (o *CatalogOverview) UnsetAzureSubscriptionId()`

UnsetAzureSubscriptionId ensures that no value is present for AzureSubscriptionId, not even an explicit nil
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

### SetPersistStaticAllocatedVmDisksNil

`func (o *CatalogOverview) SetPersistStaticAllocatedVmDisksNil(b bool)`

 SetPersistStaticAllocatedVmDisksNil sets the value for PersistStaticAllocatedVmDisks to be an explicit nil

### UnsetPersistStaticAllocatedVmDisks
`func (o *CatalogOverview) UnsetPersistStaticAllocatedVmDisks()`

UnsetPersistStaticAllocatedVmDisks ensures that no value is present for PersistStaticAllocatedVmDisks, not even an explicit nil
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

### SetSubStateNil

`func (o *CatalogOverview) SetSubStateNil(b bool)`

 SetSubStateNil sets the value for SubState to be an explicit nil

### UnsetSubState
`func (o *CatalogOverview) UnsetSubState()`

UnsetSubState ensures that no value is present for SubState, not even an explicit nil
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

### SetWarningsNil

`func (o *CatalogOverview) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *CatalogOverview) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
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

### SetStatusMessageNil

`func (o *CatalogOverview) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *CatalogOverview) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
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

### SetExtraInfoNil

`func (o *CatalogOverview) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *CatalogOverview) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil
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

### SetTransactionIdNil

`func (o *CatalogOverview) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *CatalogOverview) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
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

### SetVnetPeeringIdNil

`func (o *CatalogOverview) SetVnetPeeringIdNil(b bool)`

 SetVnetPeeringIdNil sets the value for VnetPeeringId to be an explicit nil

### UnsetVnetPeeringId
`func (o *CatalogOverview) UnsetVnetPeeringId()`

UnsetVnetPeeringId ensures that no value is present for VnetPeeringId, not even an explicit nil
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

### SetVnetPeeringNameNil

`func (o *CatalogOverview) SetVnetPeeringNameNil(b bool)`

 SetVnetPeeringNameNil sets the value for VnetPeeringName to be an explicit nil

### UnsetVnetPeeringName
`func (o *CatalogOverview) UnsetVnetPeeringName()`

UnsetVnetPeeringName ensures that no value is present for VnetPeeringName, not even an explicit nil
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

### SetVpnConnectionIdNil

`func (o *CatalogOverview) SetVpnConnectionIdNil(b bool)`

 SetVpnConnectionIdNil sets the value for VpnConnectionId to be an explicit nil

### UnsetVpnConnectionId
`func (o *CatalogOverview) UnsetVpnConnectionId()`

UnsetVpnConnectionId ensures that no value is present for VpnConnectionId, not even an explicit nil
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

### SetVpnConnectionNameNil

`func (o *CatalogOverview) SetVpnConnectionNameNil(b bool)`

 SetVpnConnectionNameNil sets the value for VpnConnectionName to be an explicit nil

### UnsetVpnConnectionName
`func (o *CatalogOverview) UnsetVpnConnectionName()`

UnsetVpnConnectionName ensures that no value is present for VpnConnectionName, not even an explicit nil
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

### SetSubscriptionIdNil

`func (o *CatalogOverview) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CatalogOverview) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
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


### GetTemplateSpecResourceGroup

`func (o *CatalogOverview) GetTemplateSpecResourceGroup() string`

GetTemplateSpecResourceGroup returns the TemplateSpecResourceGroup field if non-nil, zero value otherwise.

### GetTemplateSpecResourceGroupOk

`func (o *CatalogOverview) GetTemplateSpecResourceGroupOk() (*string, bool)`

GetTemplateSpecResourceGroupOk returns a tuple with the TemplateSpecResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSpecResourceGroup

`func (o *CatalogOverview) SetTemplateSpecResourceGroup(v string)`

SetTemplateSpecResourceGroup sets TemplateSpecResourceGroup field to given value.

### HasTemplateSpecResourceGroup

`func (o *CatalogOverview) HasTemplateSpecResourceGroup() bool`

HasTemplateSpecResourceGroup returns a boolean if a field has been set.

### SetTemplateSpecResourceGroupNil

`func (o *CatalogOverview) SetTemplateSpecResourceGroupNil(b bool)`

 SetTemplateSpecResourceGroupNil sets the value for TemplateSpecResourceGroup to be an explicit nil

### UnsetTemplateSpecResourceGroup
`func (o *CatalogOverview) UnsetTemplateSpecResourceGroup()`

UnsetTemplateSpecResourceGroup ensures that no value is present for TemplateSpecResourceGroup, not even an explicit nil
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

### SetVdaResourceGroupNil

`func (o *CatalogOverview) SetVdaResourceGroupNil(b bool)`

 SetVdaResourceGroupNil sets the value for VdaResourceGroup to be an explicit nil

### UnsetVdaResourceGroup
`func (o *CatalogOverview) UnsetVdaResourceGroup()`

UnsetVdaResourceGroup ensures that no value is present for VdaResourceGroup, not even an explicit nil
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

### SetVdaProvisioningSchemeIdNil

`func (o *CatalogOverview) SetVdaProvisioningSchemeIdNil(b bool)`

 SetVdaProvisioningSchemeIdNil sets the value for VdaProvisioningSchemeId to be an explicit nil

### UnsetVdaProvisioningSchemeId
`func (o *CatalogOverview) UnsetVdaProvisioningSchemeId()`

UnsetVdaProvisioningSchemeId ensures that no value is present for VdaProvisioningSchemeId, not even an explicit nil
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

### SetAreMcsVdaResourceGroupsUsedNil

`func (o *CatalogOverview) SetAreMcsVdaResourceGroupsUsedNil(b bool)`

 SetAreMcsVdaResourceGroupsUsedNil sets the value for AreMcsVdaResourceGroupsUsed to be an explicit nil

### UnsetAreMcsVdaResourceGroupsUsed
`func (o *CatalogOverview) UnsetAreMcsVdaResourceGroupsUsed()`

UnsetAreMcsVdaResourceGroupsUsed ensures that no value is present for AreMcsVdaResourceGroupsUsed, not even an explicit nil
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

### SetResourceLocationIdNil

`func (o *CatalogOverview) SetResourceLocationIdNil(b bool)`

 SetResourceLocationIdNil sets the value for ResourceLocationId to be an explicit nil

### UnsetResourceLocationId
`func (o *CatalogOverview) UnsetResourceLocationId()`

UnsetResourceLocationId ensures that no value is present for ResourceLocationId, not even an explicit nil
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

### SetSubnetNil

`func (o *CatalogOverview) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *CatalogOverview) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
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

### SetDomainJoinedNil

`func (o *CatalogOverview) SetDomainJoinedNil(b bool)`

 SetDomainJoinedNil sets the value for DomainJoined to be an explicit nil

### UnsetDomainJoined
`func (o *CatalogOverview) UnsetDomainJoined()`

UnsetDomainJoined ensures that no value is present for DomainJoined, not even an explicit nil
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

### SetDomainNameNil

`func (o *CatalogOverview) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *CatalogOverview) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
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

### SetDomainOUNil

`func (o *CatalogOverview) SetDomainOUNil(b bool)`

 SetDomainOUNil sets the value for DomainOU to be an explicit nil

### UnsetDomainOU
`func (o *CatalogOverview) UnsetDomainOU()`

UnsetDomainOU ensures that no value is present for DomainOU, not even an explicit nil
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

### SetDomainServiceAccountNil

`func (o *CatalogOverview) SetDomainServiceAccountNil(b bool)`

 SetDomainServiceAccountNil sets the value for DomainServiceAccount to be an explicit nil

### UnsetDomainServiceAccount
`func (o *CatalogOverview) UnsetDomainServiceAccount()`

UnsetDomainServiceAccount ensures that no value is present for DomainServiceAccount, not even an explicit nil
### GetServiceAccountUid

`func (o *CatalogOverview) GetServiceAccountUid() string`

GetServiceAccountUid returns the ServiceAccountUid field if non-nil, zero value otherwise.

### GetServiceAccountUidOk

`func (o *CatalogOverview) GetServiceAccountUidOk() (*string, bool)`

GetServiceAccountUidOk returns a tuple with the ServiceAccountUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountUid

`func (o *CatalogOverview) SetServiceAccountUid(v string)`

SetServiceAccountUid sets ServiceAccountUid field to given value.

### HasServiceAccountUid

`func (o *CatalogOverview) HasServiceAccountUid() bool`

HasServiceAccountUid returns a boolean if a field has been set.

### SetServiceAccountUidNil

`func (o *CatalogOverview) SetServiceAccountUidNil(b bool)`

 SetServiceAccountUidNil sets the value for ServiceAccountUid to be an explicit nil

### UnsetServiceAccountUid
`func (o *CatalogOverview) UnsetServiceAccountUid()`

UnsetServiceAccountUid ensures that no value is present for ServiceAccountUid, not even an explicit nil
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

### SetVmTypeInstanceTypeNil

`func (o *CatalogOverview) SetVmTypeInstanceTypeNil(b bool)`

 SetVmTypeInstanceTypeNil sets the value for VmTypeInstanceType to be an explicit nil

### UnsetVmTypeInstanceType
`func (o *CatalogOverview) UnsetVmTypeInstanceType()`

UnsetVmTypeInstanceType ensures that no value is present for VmTypeInstanceType, not even an explicit nil
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

### SetImageIdNil

`func (o *CatalogOverview) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *CatalogOverview) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
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

### SetTemplateImageNameNil

`func (o *CatalogOverview) SetTemplateImageNameNil(b bool)`

 SetTemplateImageNameNil sets the value for TemplateImageName to be an explicit nil

### UnsetTemplateImageName
`func (o *CatalogOverview) UnsetTemplateImageName()`

UnsetTemplateImageName ensures that no value is present for TemplateImageName, not even an explicit nil
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

### SetTemplateImageOsNil

`func (o *CatalogOverview) SetTemplateImageOsNil(b bool)`

 SetTemplateImageOsNil sets the value for TemplateImageOs to be an explicit nil

### UnsetTemplateImageOs
`func (o *CatalogOverview) UnsetTemplateImageOs()`

UnsetTemplateImageOs ensures that no value is present for TemplateImageOs, not even an explicit nil
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

### SetCitrixManagedNil

`func (o *CatalogOverview) SetCitrixManagedNil(b bool)`

 SetCitrixManagedNil sets the value for CitrixManaged to be an explicit nil

### UnsetCitrixManaged
`func (o *CatalogOverview) UnsetCitrixManaged()`

UnsetCitrixManaged ensures that no value is present for CitrixManaged, not even an explicit nil
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

### SetCspCustomerNil

`func (o *CatalogOverview) SetCspCustomerNil(b bool)`

 SetCspCustomerNil sets the value for CspCustomer to be an explicit nil

### UnsetCspCustomer
`func (o *CatalogOverview) UnsetCspCustomer()`

UnsetCspCustomer ensures that no value is present for CspCustomer, not even an explicit nil
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

### SetTotalMachinesInCatalogNil

`func (o *CatalogOverview) SetTotalMachinesInCatalogNil(b bool)`

 SetTotalMachinesInCatalogNil sets the value for TotalMachinesInCatalog to be an explicit nil

### UnsetTotalMachinesInCatalog
`func (o *CatalogOverview) UnsetTotalMachinesInCatalog()`

UnsetTotalMachinesInCatalog ensures that no value is present for TotalMachinesInCatalog, not even an explicit nil
### GetNumOfUsers

`func (o *CatalogOverview) GetNumOfUsers() int32`

GetNumOfUsers returns the NumOfUsers field if non-nil, zero value otherwise.

### GetNumOfUsersOk

`func (o *CatalogOverview) GetNumOfUsersOk() (*int32, bool)`

GetNumOfUsersOk returns a tuple with the NumOfUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfUsers

`func (o *CatalogOverview) SetNumOfUsers(v int32)`

SetNumOfUsers sets NumOfUsers field to given value.

### HasNumOfUsers

`func (o *CatalogOverview) HasNumOfUsers() bool`

HasNumOfUsers returns a boolean if a field has been set.

### SetNumOfUsersNil

`func (o *CatalogOverview) SetNumOfUsersNil(b bool)`

 SetNumOfUsersNil sets the value for NumOfUsers to be an explicit nil

### UnsetNumOfUsers
`func (o *CatalogOverview) UnsetNumOfUsers()`

UnsetNumOfUsers ensures that no value is present for NumOfUsers, not even an explicit nil
### GetMaxNumOfUsers

`func (o *CatalogOverview) GetMaxNumOfUsers() int32`

GetMaxNumOfUsers returns the MaxNumOfUsers field if non-nil, zero value otherwise.

### GetMaxNumOfUsersOk

`func (o *CatalogOverview) GetMaxNumOfUsersOk() (*int32, bool)`

GetMaxNumOfUsersOk returns a tuple with the MaxNumOfUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOfUsers

`func (o *CatalogOverview) SetMaxNumOfUsers(v int32)`

SetMaxNumOfUsers sets MaxNumOfUsers field to given value.

### HasMaxNumOfUsers

`func (o *CatalogOverview) HasMaxNumOfUsers() bool`

HasMaxNumOfUsers returns a boolean if a field has been set.

### SetMaxNumOfUsersNil

`func (o *CatalogOverview) SetMaxNumOfUsersNil(b bool)`

 SetMaxNumOfUsersNil sets the value for MaxNumOfUsers to be an explicit nil

### UnsetMaxNumOfUsers
`func (o *CatalogOverview) UnsetMaxNumOfUsers()`

UnsetMaxNumOfUsers ensures that no value is present for MaxNumOfUsers, not even an explicit nil
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

### SetWriteBackCacheConfigurationNil

`func (o *CatalogOverview) SetWriteBackCacheConfigurationNil(b bool)`

 SetWriteBackCacheConfigurationNil sets the value for WriteBackCacheConfiguration to be an explicit nil

### UnsetWriteBackCacheConfiguration
`func (o *CatalogOverview) UnsetWriteBackCacheConfiguration()`

UnsetWriteBackCacheConfiguration ensures that no value is present for WriteBackCacheConfiguration, not even an explicit nil
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

### SetTaskCompletionPercentageNil

`func (o *CatalogOverview) SetTaskCompletionPercentageNil(b bool)`

 SetTaskCompletionPercentageNil sets the value for TaskCompletionPercentage to be an explicit nil

### UnsetTaskCompletionPercentage
`func (o *CatalogOverview) UnsetTaskCompletionPercentage()`

UnsetTaskCompletionPercentage ensures that no value is present for TaskCompletionPercentage, not even an explicit nil
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

### SetLastModifiedTimeNil

`func (o *CatalogOverview) SetLastModifiedTimeNil(b bool)`

 SetLastModifiedTimeNil sets the value for LastModifiedTime to be an explicit nil

### UnsetLastModifiedTime
`func (o *CatalogOverview) UnsetLastModifiedTime()`

UnsetLastModifiedTime ensures that no value is present for LastModifiedTime, not even an explicit nil
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

### SetLastBackupTimeNil

`func (o *CatalogOverview) SetLastBackupTimeNil(b bool)`

 SetLastBackupTimeNil sets the value for LastBackupTime to be an explicit nil

### UnsetLastBackupTime
`func (o *CatalogOverview) UnsetLastBackupTime()`

UnsetLastBackupTime ensures that no value is present for LastBackupTime, not even an explicit nil
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

### SetOrganizationalUnitsNil

`func (o *CatalogOverview) SetOrganizationalUnitsNil(b bool)`

 SetOrganizationalUnitsNil sets the value for OrganizationalUnits to be an explicit nil

### UnsetOrganizationalUnits
`func (o *CatalogOverview) UnsetOrganizationalUnits()`

UnsetOrganizationalUnits ensures that no value is present for OrganizationalUnits, not even an explicit nil
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

### GetEnableAcceleratedNetworking

`func (o *CatalogOverview) GetEnableAcceleratedNetworking() bool`

GetEnableAcceleratedNetworking returns the EnableAcceleratedNetworking field if non-nil, zero value otherwise.

### GetEnableAcceleratedNetworkingOk

`func (o *CatalogOverview) GetEnableAcceleratedNetworkingOk() (*bool, bool)`

GetEnableAcceleratedNetworkingOk returns a tuple with the EnableAcceleratedNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAcceleratedNetworking

`func (o *CatalogOverview) SetEnableAcceleratedNetworking(v bool)`

SetEnableAcceleratedNetworking sets EnableAcceleratedNetworking field to given value.

### HasEnableAcceleratedNetworking

`func (o *CatalogOverview) HasEnableAcceleratedNetworking() bool`

HasEnableAcceleratedNetworking returns a boolean if a field has been set.

### GetEnableEncryptionAtHost

`func (o *CatalogOverview) GetEnableEncryptionAtHost() bool`

GetEnableEncryptionAtHost returns the EnableEncryptionAtHost field if non-nil, zero value otherwise.

### GetEnableEncryptionAtHostOk

`func (o *CatalogOverview) GetEnableEncryptionAtHostOk() (*bool, bool)`

GetEnableEncryptionAtHostOk returns a tuple with the EnableEncryptionAtHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryptionAtHost

`func (o *CatalogOverview) SetEnableEncryptionAtHost(v bool)`

SetEnableEncryptionAtHost sets EnableEncryptionAtHost field to given value.

### HasEnableEncryptionAtHost

`func (o *CatalogOverview) HasEnableEncryptionAtHost() bool`

HasEnableEncryptionAtHost returns a boolean if a field has been set.

### GetCatalogType

`func (o *CatalogOverview) GetCatalogType() CatalogType`

GetCatalogType returns the CatalogType field if non-nil, zero value otherwise.

### GetCatalogTypeOk

`func (o *CatalogOverview) GetCatalogTypeOk() (*CatalogType, bool)`

GetCatalogTypeOk returns a tuple with the CatalogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogType

`func (o *CatalogOverview) SetCatalogType(v CatalogType)`

SetCatalogType sets CatalogType field to given value.

### HasCatalogType

`func (o *CatalogOverview) HasCatalogType() bool`

HasCatalogType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


