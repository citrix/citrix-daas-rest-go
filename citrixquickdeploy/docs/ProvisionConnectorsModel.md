# ProvisionConnectorsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** | Number of connectors to provision | [optional] 
**OrganizationalUnit** | Pointer to **string** | The OU the new connectors should be added to | [optional] 
**ServiceAccount** | Pointer to **string** | Service account to join the connector to the Resource Location&#39;s domain | [optional] 
**ServiceAccountPassword** | Pointer to **string** | Password for the service account | [optional] 
**AzureResourceGroup** | Pointer to **string** | Azure Resource Group where the connectors should be deployed | [optional] 
**UseAzureHub** | Pointer to **bool** | Indicates if the connector should be provisioned with Azure HUB enabled | [optional] 
**VmSize** | Pointer to **string** | The ID of the vm size | [optional] 
**DomainName** | Pointer to **string** | The domain the connectors will be joined to. Used when adding connectors to connectorless RL | [optional] 

## Methods

### NewProvisionConnectorsModel

`func NewProvisionConnectorsModel() *ProvisionConnectorsModel`

NewProvisionConnectorsModel instantiates a new ProvisionConnectorsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionConnectorsModelWithDefaults

`func NewProvisionConnectorsModelWithDefaults() *ProvisionConnectorsModel`

NewProvisionConnectorsModelWithDefaults instantiates a new ProvisionConnectorsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *ProvisionConnectorsModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProvisionConnectorsModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProvisionConnectorsModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProvisionConnectorsModel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *ProvisionConnectorsModel) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *ProvisionConnectorsModel) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *ProvisionConnectorsModel) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *ProvisionConnectorsModel) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetServiceAccount

`func (o *ProvisionConnectorsModel) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ProvisionConnectorsModel) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ProvisionConnectorsModel) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *ProvisionConnectorsModel) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetServiceAccountPassword

`func (o *ProvisionConnectorsModel) GetServiceAccountPassword() string`

GetServiceAccountPassword returns the ServiceAccountPassword field if non-nil, zero value otherwise.

### GetServiceAccountPasswordOk

`func (o *ProvisionConnectorsModel) GetServiceAccountPasswordOk() (*string, bool)`

GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountPassword

`func (o *ProvisionConnectorsModel) SetServiceAccountPassword(v string)`

SetServiceAccountPassword sets ServiceAccountPassword field to given value.

### HasServiceAccountPassword

`func (o *ProvisionConnectorsModel) HasServiceAccountPassword() bool`

HasServiceAccountPassword returns a boolean if a field has been set.

### GetAzureResourceGroup

`func (o *ProvisionConnectorsModel) GetAzureResourceGroup() string`

GetAzureResourceGroup returns the AzureResourceGroup field if non-nil, zero value otherwise.

### GetAzureResourceGroupOk

`func (o *ProvisionConnectorsModel) GetAzureResourceGroupOk() (*string, bool)`

GetAzureResourceGroupOk returns a tuple with the AzureResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroup

`func (o *ProvisionConnectorsModel) SetAzureResourceGroup(v string)`

SetAzureResourceGroup sets AzureResourceGroup field to given value.

### HasAzureResourceGroup

`func (o *ProvisionConnectorsModel) HasAzureResourceGroup() bool`

HasAzureResourceGroup returns a boolean if a field has been set.

### GetUseAzureHub

`func (o *ProvisionConnectorsModel) GetUseAzureHub() bool`

GetUseAzureHub returns the UseAzureHub field if non-nil, zero value otherwise.

### GetUseAzureHubOk

`func (o *ProvisionConnectorsModel) GetUseAzureHubOk() (*bool, bool)`

GetUseAzureHubOk returns a tuple with the UseAzureHub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAzureHub

`func (o *ProvisionConnectorsModel) SetUseAzureHub(v bool)`

SetUseAzureHub sets UseAzureHub field to given value.

### HasUseAzureHub

`func (o *ProvisionConnectorsModel) HasUseAzureHub() bool`

HasUseAzureHub returns a boolean if a field has been set.

### GetVmSize

`func (o *ProvisionConnectorsModel) GetVmSize() string`

GetVmSize returns the VmSize field if non-nil, zero value otherwise.

### GetVmSizeOk

`func (o *ProvisionConnectorsModel) GetVmSizeOk() (*string, bool)`

GetVmSizeOk returns a tuple with the VmSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSize

`func (o *ProvisionConnectorsModel) SetVmSize(v string)`

SetVmSize sets VmSize field to given value.

### HasVmSize

`func (o *ProvisionConnectorsModel) HasVmSize() bool`

HasVmSize returns a boolean if a field has been set.

### GetDomainName

`func (o *ProvisionConnectorsModel) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ProvisionConnectorsModel) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ProvisionConnectorsModel) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ProvisionConnectorsModel) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


