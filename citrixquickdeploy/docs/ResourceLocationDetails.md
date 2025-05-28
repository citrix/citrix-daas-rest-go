# ResourceLocationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | Pointer to **string** | ID of the Resource Location | [optional] 
**CspCustomer** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**Name** | Pointer to **string** | Name of the Resource Location | [optional] 
**State** | Pointer to [**ResourceLocationState**](ResourceLocationState.md) | State of the Resource Location | [optional] 
**Region** | Pointer to **string** | Region where resources associated with the RL is located | [optional] 
**RegionId** | Pointer to **string** | Id of the Region where resources associated with the RL is located | [optional] 
**IsCitrixManaged** | Pointer to **bool** | Indicates if this RL is associated with Citrix Managed resources | [optional] 
**IsForNonDomainJoinedVms** | Pointer to **bool** | Indicates if the RL is used by non-domain joined catalogs | [optional] 
**DomainName** | Pointer to **string** | Domain name of the AD domain catalogs in this RL are associated with | [optional] 
**AzureSubscriptionId** | Pointer to **string** | ID of the Azure subscription associated with the Resource Location | [optional] 
**AzureVnet** | Pointer to **string** | Name of the Azure Virtual Network the RL is associated with | [optional] 
**OnPremConnection** | Pointer to **string** | Name of the on-prem connection associated with | [optional] 
**Connectors** | Pointer to [**[]ConnectorDetails**](ConnectorDetails.md) | List of connectors configured for the RL | [optional] 
**Jobs** | Pointer to [**[]ResourceLocationJob**](ResourceLocationJob.md) | List of jobs performed by the Resource Location. | [optional] 
**AssociatedCatalogs** | Pointer to [**[]AssociatedCatalog**](AssociatedCatalog.md) | List of catalogs that are associate with the Resource Location. | [optional] 
**ConnectorResourceGroup** | Pointer to **string** | The most recently used resource group for connectors of a BYOA RL | [optional] 
**VnetResourceGroup** | Pointer to **string** | The resource group containing the vnet the RL is associated with | [optional] 
**OrganizationalUnit** | Pointer to **string** | The most recently used OU for the connectors of a BYOA RL | [optional] 
**IsSecureBrowser** | Pointer to **bool** | Indicates if the Resource Location is for Secure Browser | [optional] 
**IsForConnectorlessCatalogs** | Pointer to **bool** | Indicates if the Resource Location is for connectorless catalogs | [optional] 
**IsSupportingAddingConnectors** | Pointer to **bool** | Indicates if connectorless Resource Location can add connectors | [optional] 
**NatGateway** | Pointer to [**NatGatewayModelOverview**](NatGatewayModelOverview.md) | Overview model of the NatGateway | [optional] 

## Methods

### NewResourceLocationDetails

`func NewResourceLocationDetails() *ResourceLocationDetails`

NewResourceLocationDetails instantiates a new ResourceLocationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLocationDetailsWithDefaults

`func NewResourceLocationDetailsWithDefaults() *ResourceLocationDetails`

NewResourceLocationDetailsWithDefaults instantiates a new ResourceLocationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *ResourceLocationDetails) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ResourceLocationDetails) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ResourceLocationDetails) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ResourceLocationDetails) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetCspCustomer

`func (o *ResourceLocationDetails) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *ResourceLocationDetails) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *ResourceLocationDetails) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *ResourceLocationDetails) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### GetName

`func (o *ResourceLocationDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceLocationDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceLocationDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceLocationDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *ResourceLocationDetails) GetState() ResourceLocationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResourceLocationDetails) GetStateOk() (*ResourceLocationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResourceLocationDetails) SetState(v ResourceLocationState)`

SetState sets State field to given value.

### HasState

`func (o *ResourceLocationDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRegion

`func (o *ResourceLocationDetails) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ResourceLocationDetails) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ResourceLocationDetails) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ResourceLocationDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionId

`func (o *ResourceLocationDetails) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ResourceLocationDetails) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ResourceLocationDetails) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ResourceLocationDetails) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetIsCitrixManaged

`func (o *ResourceLocationDetails) GetIsCitrixManaged() bool`

GetIsCitrixManaged returns the IsCitrixManaged field if non-nil, zero value otherwise.

### GetIsCitrixManagedOk

`func (o *ResourceLocationDetails) GetIsCitrixManagedOk() (*bool, bool)`

GetIsCitrixManagedOk returns a tuple with the IsCitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCitrixManaged

`func (o *ResourceLocationDetails) SetIsCitrixManaged(v bool)`

SetIsCitrixManaged sets IsCitrixManaged field to given value.

### HasIsCitrixManaged

`func (o *ResourceLocationDetails) HasIsCitrixManaged() bool`

HasIsCitrixManaged returns a boolean if a field has been set.

### GetIsForNonDomainJoinedVms

`func (o *ResourceLocationDetails) GetIsForNonDomainJoinedVms() bool`

GetIsForNonDomainJoinedVms returns the IsForNonDomainJoinedVms field if non-nil, zero value otherwise.

### GetIsForNonDomainJoinedVmsOk

`func (o *ResourceLocationDetails) GetIsForNonDomainJoinedVmsOk() (*bool, bool)`

GetIsForNonDomainJoinedVmsOk returns a tuple with the IsForNonDomainJoinedVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForNonDomainJoinedVms

`func (o *ResourceLocationDetails) SetIsForNonDomainJoinedVms(v bool)`

SetIsForNonDomainJoinedVms sets IsForNonDomainJoinedVms field to given value.

### HasIsForNonDomainJoinedVms

`func (o *ResourceLocationDetails) HasIsForNonDomainJoinedVms() bool`

HasIsForNonDomainJoinedVms returns a boolean if a field has been set.

### GetDomainName

`func (o *ResourceLocationDetails) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ResourceLocationDetails) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ResourceLocationDetails) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ResourceLocationDetails) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetAzureSubscriptionId

`func (o *ResourceLocationDetails) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *ResourceLocationDetails) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *ResourceLocationDetails) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *ResourceLocationDetails) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.

### GetAzureVnet

`func (o *ResourceLocationDetails) GetAzureVnet() string`

GetAzureVnet returns the AzureVnet field if non-nil, zero value otherwise.

### GetAzureVnetOk

`func (o *ResourceLocationDetails) GetAzureVnetOk() (*string, bool)`

GetAzureVnetOk returns a tuple with the AzureVnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVnet

`func (o *ResourceLocationDetails) SetAzureVnet(v string)`

SetAzureVnet sets AzureVnet field to given value.

### HasAzureVnet

`func (o *ResourceLocationDetails) HasAzureVnet() bool`

HasAzureVnet returns a boolean if a field has been set.

### GetOnPremConnection

`func (o *ResourceLocationDetails) GetOnPremConnection() string`

GetOnPremConnection returns the OnPremConnection field if non-nil, zero value otherwise.

### GetOnPremConnectionOk

`func (o *ResourceLocationDetails) GetOnPremConnectionOk() (*string, bool)`

GetOnPremConnectionOk returns a tuple with the OnPremConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremConnection

`func (o *ResourceLocationDetails) SetOnPremConnection(v string)`

SetOnPremConnection sets OnPremConnection field to given value.

### HasOnPremConnection

`func (o *ResourceLocationDetails) HasOnPremConnection() bool`

HasOnPremConnection returns a boolean if a field has been set.

### GetConnectors

`func (o *ResourceLocationDetails) GetConnectors() []ConnectorDetails`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *ResourceLocationDetails) GetConnectorsOk() (*[]ConnectorDetails, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *ResourceLocationDetails) SetConnectors(v []ConnectorDetails)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *ResourceLocationDetails) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetJobs

`func (o *ResourceLocationDetails) GetJobs() []ResourceLocationJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ResourceLocationDetails) GetJobsOk() (*[]ResourceLocationJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ResourceLocationDetails) SetJobs(v []ResourceLocationJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ResourceLocationDetails) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetAssociatedCatalogs

`func (o *ResourceLocationDetails) GetAssociatedCatalogs() []AssociatedCatalog`

GetAssociatedCatalogs returns the AssociatedCatalogs field if non-nil, zero value otherwise.

### GetAssociatedCatalogsOk

`func (o *ResourceLocationDetails) GetAssociatedCatalogsOk() (*[]AssociatedCatalog, bool)`

GetAssociatedCatalogsOk returns a tuple with the AssociatedCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedCatalogs

`func (o *ResourceLocationDetails) SetAssociatedCatalogs(v []AssociatedCatalog)`

SetAssociatedCatalogs sets AssociatedCatalogs field to given value.

### HasAssociatedCatalogs

`func (o *ResourceLocationDetails) HasAssociatedCatalogs() bool`

HasAssociatedCatalogs returns a boolean if a field has been set.

### GetConnectorResourceGroup

`func (o *ResourceLocationDetails) GetConnectorResourceGroup() string`

GetConnectorResourceGroup returns the ConnectorResourceGroup field if non-nil, zero value otherwise.

### GetConnectorResourceGroupOk

`func (o *ResourceLocationDetails) GetConnectorResourceGroupOk() (*string, bool)`

GetConnectorResourceGroupOk returns a tuple with the ConnectorResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorResourceGroup

`func (o *ResourceLocationDetails) SetConnectorResourceGroup(v string)`

SetConnectorResourceGroup sets ConnectorResourceGroup field to given value.

### HasConnectorResourceGroup

`func (o *ResourceLocationDetails) HasConnectorResourceGroup() bool`

HasConnectorResourceGroup returns a boolean if a field has been set.

### GetVnetResourceGroup

`func (o *ResourceLocationDetails) GetVnetResourceGroup() string`

GetVnetResourceGroup returns the VnetResourceGroup field if non-nil, zero value otherwise.

### GetVnetResourceGroupOk

`func (o *ResourceLocationDetails) GetVnetResourceGroupOk() (*string, bool)`

GetVnetResourceGroupOk returns a tuple with the VnetResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetResourceGroup

`func (o *ResourceLocationDetails) SetVnetResourceGroup(v string)`

SetVnetResourceGroup sets VnetResourceGroup field to given value.

### HasVnetResourceGroup

`func (o *ResourceLocationDetails) HasVnetResourceGroup() bool`

HasVnetResourceGroup returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *ResourceLocationDetails) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *ResourceLocationDetails) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *ResourceLocationDetails) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *ResourceLocationDetails) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetIsSecureBrowser

`func (o *ResourceLocationDetails) GetIsSecureBrowser() bool`

GetIsSecureBrowser returns the IsSecureBrowser field if non-nil, zero value otherwise.

### GetIsSecureBrowserOk

`func (o *ResourceLocationDetails) GetIsSecureBrowserOk() (*bool, bool)`

GetIsSecureBrowserOk returns a tuple with the IsSecureBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureBrowser

`func (o *ResourceLocationDetails) SetIsSecureBrowser(v bool)`

SetIsSecureBrowser sets IsSecureBrowser field to given value.

### HasIsSecureBrowser

`func (o *ResourceLocationDetails) HasIsSecureBrowser() bool`

HasIsSecureBrowser returns a boolean if a field has been set.

### GetIsForConnectorlessCatalogs

`func (o *ResourceLocationDetails) GetIsForConnectorlessCatalogs() bool`

GetIsForConnectorlessCatalogs returns the IsForConnectorlessCatalogs field if non-nil, zero value otherwise.

### GetIsForConnectorlessCatalogsOk

`func (o *ResourceLocationDetails) GetIsForConnectorlessCatalogsOk() (*bool, bool)`

GetIsForConnectorlessCatalogsOk returns a tuple with the IsForConnectorlessCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForConnectorlessCatalogs

`func (o *ResourceLocationDetails) SetIsForConnectorlessCatalogs(v bool)`

SetIsForConnectorlessCatalogs sets IsForConnectorlessCatalogs field to given value.

### HasIsForConnectorlessCatalogs

`func (o *ResourceLocationDetails) HasIsForConnectorlessCatalogs() bool`

HasIsForConnectorlessCatalogs returns a boolean if a field has been set.

### GetIsSupportingAddingConnectors

`func (o *ResourceLocationDetails) GetIsSupportingAddingConnectors() bool`

GetIsSupportingAddingConnectors returns the IsSupportingAddingConnectors field if non-nil, zero value otherwise.

### GetIsSupportingAddingConnectorsOk

`func (o *ResourceLocationDetails) GetIsSupportingAddingConnectorsOk() (*bool, bool)`

GetIsSupportingAddingConnectorsOk returns a tuple with the IsSupportingAddingConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupportingAddingConnectors

`func (o *ResourceLocationDetails) SetIsSupportingAddingConnectors(v bool)`

SetIsSupportingAddingConnectors sets IsSupportingAddingConnectors field to given value.

### HasIsSupportingAddingConnectors

`func (o *ResourceLocationDetails) HasIsSupportingAddingConnectors() bool`

HasIsSupportingAddingConnectors returns a boolean if a field has been set.

### GetNatGateway

`func (o *ResourceLocationDetails) GetNatGateway() NatGatewayModelOverview`

GetNatGateway returns the NatGateway field if non-nil, zero value otherwise.

### GetNatGatewayOk

`func (o *ResourceLocationDetails) GetNatGatewayOk() (*NatGatewayModelOverview, bool)`

GetNatGatewayOk returns a tuple with the NatGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatGateway

`func (o *ResourceLocationDetails) SetNatGateway(v NatGatewayModelOverview)`

SetNatGateway sets NatGateway field to given value.

### HasNatGateway

`func (o *ResourceLocationDetails) HasNatGateway() bool`

HasNatGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


