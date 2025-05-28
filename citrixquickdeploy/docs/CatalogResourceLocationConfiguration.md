# CatalogResourceLocationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAssignedExistingResourceLocation** | Pointer to **bool** | Indicates if the catalog is assigned to an existing Resource Location, instead of needing a new one | [optional] 
**Name** | **string** | The desired name of the resource location that will be created for the catalog | 
**AzureResourceGroup** | Pointer to **string** | Name of the resource location where to provision the connector vdas | [optional] 
**OrganizationalUnit** | Pointer to **string** | Organization Unit associated with computer accounts added for the Resource Location | [optional] 
**ConnectivityMethod** | Pointer to [**ConnectivityType**](ConnectivityType.md) | Connectivity method for access to desktops and apps | [optional] 
**VmSize** | Pointer to **string** | The ID of the vm size for the connector | [optional] 

## Methods

### NewCatalogResourceLocationConfiguration

`func NewCatalogResourceLocationConfiguration(name string, ) *CatalogResourceLocationConfiguration`

NewCatalogResourceLocationConfiguration instantiates a new CatalogResourceLocationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogResourceLocationConfigurationWithDefaults

`func NewCatalogResourceLocationConfigurationWithDefaults() *CatalogResourceLocationConfiguration`

NewCatalogResourceLocationConfigurationWithDefaults instantiates a new CatalogResourceLocationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAssignedExistingResourceLocation

`func (o *CatalogResourceLocationConfiguration) GetIsAssignedExistingResourceLocation() bool`

GetIsAssignedExistingResourceLocation returns the IsAssignedExistingResourceLocation field if non-nil, zero value otherwise.

### GetIsAssignedExistingResourceLocationOk

`func (o *CatalogResourceLocationConfiguration) GetIsAssignedExistingResourceLocationOk() (*bool, bool)`

GetIsAssignedExistingResourceLocationOk returns a tuple with the IsAssignedExistingResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssignedExistingResourceLocation

`func (o *CatalogResourceLocationConfiguration) SetIsAssignedExistingResourceLocation(v bool)`

SetIsAssignedExistingResourceLocation sets IsAssignedExistingResourceLocation field to given value.

### HasIsAssignedExistingResourceLocation

`func (o *CatalogResourceLocationConfiguration) HasIsAssignedExistingResourceLocation() bool`

HasIsAssignedExistingResourceLocation returns a boolean if a field has been set.

### GetName

`func (o *CatalogResourceLocationConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogResourceLocationConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogResourceLocationConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetAzureResourceGroup

`func (o *CatalogResourceLocationConfiguration) GetAzureResourceGroup() string`

GetAzureResourceGroup returns the AzureResourceGroup field if non-nil, zero value otherwise.

### GetAzureResourceGroupOk

`func (o *CatalogResourceLocationConfiguration) GetAzureResourceGroupOk() (*string, bool)`

GetAzureResourceGroupOk returns a tuple with the AzureResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroup

`func (o *CatalogResourceLocationConfiguration) SetAzureResourceGroup(v string)`

SetAzureResourceGroup sets AzureResourceGroup field to given value.

### HasAzureResourceGroup

`func (o *CatalogResourceLocationConfiguration) HasAzureResourceGroup() bool`

HasAzureResourceGroup returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *CatalogResourceLocationConfiguration) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *CatalogResourceLocationConfiguration) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *CatalogResourceLocationConfiguration) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *CatalogResourceLocationConfiguration) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetConnectivityMethod

`func (o *CatalogResourceLocationConfiguration) GetConnectivityMethod() ConnectivityType`

GetConnectivityMethod returns the ConnectivityMethod field if non-nil, zero value otherwise.

### GetConnectivityMethodOk

`func (o *CatalogResourceLocationConfiguration) GetConnectivityMethodOk() (*ConnectivityType, bool)`

GetConnectivityMethodOk returns a tuple with the ConnectivityMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityMethod

`func (o *CatalogResourceLocationConfiguration) SetConnectivityMethod(v ConnectivityType)`

SetConnectivityMethod sets ConnectivityMethod field to given value.

### HasConnectivityMethod

`func (o *CatalogResourceLocationConfiguration) HasConnectivityMethod() bool`

HasConnectivityMethod returns a boolean if a field has been set.

### GetVmSize

`func (o *CatalogResourceLocationConfiguration) GetVmSize() string`

GetVmSize returns the VmSize field if non-nil, zero value otherwise.

### GetVmSizeOk

`func (o *CatalogResourceLocationConfiguration) GetVmSizeOk() (*string, bool)`

GetVmSizeOk returns a tuple with the VmSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSize

`func (o *CatalogResourceLocationConfiguration) SetVmSize(v string)`

SetVmSize sets VmSize field to given value.

### HasVmSize

`func (o *CatalogResourceLocationConfiguration) HasVmSize() bool`

HasVmSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


