# AzureCatalogVdaCostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmName** | Pointer to **string** | Name of the VDA VM | [optional] 
**AzureResourceGroup** | Pointer to **string** | Name of the Azure resource group which the VDA VM is located | [optional] 
**TotalCost** | Pointer to **float64** | Total cost on Azure for the VDA | [optional] 
**DiskCost** | Pointer to **float64** | Cost of the VDA Disk | [optional] 
**VmCost** | Pointer to **float64** | Cost of the VDA VM | [optional] 
**VmUsage** | Pointer to **float64** | Unit of usage of the VDA VM | [optional] 
**AssignedUsers** | Pointer to **string** | A list of users assigned to the VDA, concatenated into a string with comma for csv output | [optional] 

## Methods

### NewAzureCatalogVdaCostModel

`func NewAzureCatalogVdaCostModel() *AzureCatalogVdaCostModel`

NewAzureCatalogVdaCostModel instantiates a new AzureCatalogVdaCostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCatalogVdaCostModelWithDefaults

`func NewAzureCatalogVdaCostModelWithDefaults() *AzureCatalogVdaCostModel`

NewAzureCatalogVdaCostModelWithDefaults instantiates a new AzureCatalogVdaCostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmName

`func (o *AzureCatalogVdaCostModel) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *AzureCatalogVdaCostModel) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *AzureCatalogVdaCostModel) SetVmName(v string)`

SetVmName sets VmName field to given value.

### HasVmName

`func (o *AzureCatalogVdaCostModel) HasVmName() bool`

HasVmName returns a boolean if a field has been set.

### GetAzureResourceGroup

`func (o *AzureCatalogVdaCostModel) GetAzureResourceGroup() string`

GetAzureResourceGroup returns the AzureResourceGroup field if non-nil, zero value otherwise.

### GetAzureResourceGroupOk

`func (o *AzureCatalogVdaCostModel) GetAzureResourceGroupOk() (*string, bool)`

GetAzureResourceGroupOk returns a tuple with the AzureResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroup

`func (o *AzureCatalogVdaCostModel) SetAzureResourceGroup(v string)`

SetAzureResourceGroup sets AzureResourceGroup field to given value.

### HasAzureResourceGroup

`func (o *AzureCatalogVdaCostModel) HasAzureResourceGroup() bool`

HasAzureResourceGroup returns a boolean if a field has been set.

### GetTotalCost

`func (o *AzureCatalogVdaCostModel) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *AzureCatalogVdaCostModel) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *AzureCatalogVdaCostModel) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *AzureCatalogVdaCostModel) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetDiskCost

`func (o *AzureCatalogVdaCostModel) GetDiskCost() float64`

GetDiskCost returns the DiskCost field if non-nil, zero value otherwise.

### GetDiskCostOk

`func (o *AzureCatalogVdaCostModel) GetDiskCostOk() (*float64, bool)`

GetDiskCostOk returns a tuple with the DiskCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCost

`func (o *AzureCatalogVdaCostModel) SetDiskCost(v float64)`

SetDiskCost sets DiskCost field to given value.

### HasDiskCost

`func (o *AzureCatalogVdaCostModel) HasDiskCost() bool`

HasDiskCost returns a boolean if a field has been set.

### GetVmCost

`func (o *AzureCatalogVdaCostModel) GetVmCost() float64`

GetVmCost returns the VmCost field if non-nil, zero value otherwise.

### GetVmCostOk

`func (o *AzureCatalogVdaCostModel) GetVmCostOk() (*float64, bool)`

GetVmCostOk returns a tuple with the VmCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCost

`func (o *AzureCatalogVdaCostModel) SetVmCost(v float64)`

SetVmCost sets VmCost field to given value.

### HasVmCost

`func (o *AzureCatalogVdaCostModel) HasVmCost() bool`

HasVmCost returns a boolean if a field has been set.

### GetVmUsage

`func (o *AzureCatalogVdaCostModel) GetVmUsage() float64`

GetVmUsage returns the VmUsage field if non-nil, zero value otherwise.

### GetVmUsageOk

`func (o *AzureCatalogVdaCostModel) GetVmUsageOk() (*float64, bool)`

GetVmUsageOk returns a tuple with the VmUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUsage

`func (o *AzureCatalogVdaCostModel) SetVmUsage(v float64)`

SetVmUsage sets VmUsage field to given value.

### HasVmUsage

`func (o *AzureCatalogVdaCostModel) HasVmUsage() bool`

HasVmUsage returns a boolean if a field has been set.

### GetAssignedUsers

`func (o *AzureCatalogVdaCostModel) GetAssignedUsers() string`

GetAssignedUsers returns the AssignedUsers field if non-nil, zero value otherwise.

### GetAssignedUsersOk

`func (o *AzureCatalogVdaCostModel) GetAssignedUsersOk() (*string, bool)`

GetAssignedUsersOk returns a tuple with the AssignedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUsers

`func (o *AzureCatalogVdaCostModel) SetAssignedUsers(v string)`

SetAssignedUsers sets AssignedUsers field to given value.

### HasAssignedUsers

`func (o *AzureCatalogVdaCostModel) HasAssignedUsers() bool`

HasAssignedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


