# AddImageBuilderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Friendly name of the template | 
**ImageId** | **string** | ID of the image we will be building off of | 
**SubscriptionId** | **string** | ID of the Azure Subscription | 
**ResourceGroup** | **string** | Name of the Resource Group where the VNet is located | 
**VNet** | **string** | Name of the Virtual Network the Image Builder VM will connect to | 
**Subnet** | **string** | Name of the subnet the Image Builder VM will connect to | 
**DomainName** | Pointer to **string** | Name of the Domain the Image Builder will join | [optional] 
**OrganizationalUnit** | Pointer to **string** | The OU to associate the Image Builder VM with | [optional] 
**VmType** | **string** | The type of VM Instance type | 
**ServiceAccountName** | **string** | The service account used to join the Image Builder VM to the domain | 
**ServiceAccountPassword** | **string** | The service account password | 
**Notes** | Pointer to **string** | Customer notes about template image | [optional] 
**AllowedIPs** | Pointer to **[]string** | Ip Addresses allowed to RDP | [optional] 
**CspCustomerId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 

## Methods

### NewAddImageBuilderModel

`func NewAddImageBuilderModel(name string, imageId string, subscriptionId string, resourceGroup string, vNet string, subnet string, vmType string, serviceAccountName string, serviceAccountPassword string, ) *AddImageBuilderModel`

NewAddImageBuilderModel instantiates a new AddImageBuilderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddImageBuilderModelWithDefaults

`func NewAddImageBuilderModelWithDefaults() *AddImageBuilderModel`

NewAddImageBuilderModelWithDefaults instantiates a new AddImageBuilderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddImageBuilderModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddImageBuilderModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddImageBuilderModel) SetName(v string)`

SetName sets Name field to given value.


### GetImageId

`func (o *AddImageBuilderModel) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AddImageBuilderModel) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AddImageBuilderModel) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetSubscriptionId

`func (o *AddImageBuilderModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AddImageBuilderModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AddImageBuilderModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroup

`func (o *AddImageBuilderModel) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AddImageBuilderModel) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AddImageBuilderModel) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetVNet

`func (o *AddImageBuilderModel) GetVNet() string`

GetVNet returns the VNet field if non-nil, zero value otherwise.

### GetVNetOk

`func (o *AddImageBuilderModel) GetVNetOk() (*string, bool)`

GetVNetOk returns a tuple with the VNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNet

`func (o *AddImageBuilderModel) SetVNet(v string)`

SetVNet sets VNet field to given value.


### GetSubnet

`func (o *AddImageBuilderModel) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AddImageBuilderModel) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AddImageBuilderModel) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetDomainName

`func (o *AddImageBuilderModel) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AddImageBuilderModel) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *AddImageBuilderModel) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *AddImageBuilderModel) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *AddImageBuilderModel) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *AddImageBuilderModel) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *AddImageBuilderModel) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *AddImageBuilderModel) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetVmType

`func (o *AddImageBuilderModel) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *AddImageBuilderModel) GetVmTypeOk() (*string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmType

`func (o *AddImageBuilderModel) SetVmType(v string)`

SetVmType sets VmType field to given value.


### GetServiceAccountName

`func (o *AddImageBuilderModel) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *AddImageBuilderModel) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *AddImageBuilderModel) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.


### GetServiceAccountPassword

`func (o *AddImageBuilderModel) GetServiceAccountPassword() string`

GetServiceAccountPassword returns the ServiceAccountPassword field if non-nil, zero value otherwise.

### GetServiceAccountPasswordOk

`func (o *AddImageBuilderModel) GetServiceAccountPasswordOk() (*string, bool)`

GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountPassword

`func (o *AddImageBuilderModel) SetServiceAccountPassword(v string)`

SetServiceAccountPassword sets ServiceAccountPassword field to given value.


### GetNotes

`func (o *AddImageBuilderModel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddImageBuilderModel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddImageBuilderModel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddImageBuilderModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAllowedIPs

`func (o *AddImageBuilderModel) GetAllowedIPs() []string`

GetAllowedIPs returns the AllowedIPs field if non-nil, zero value otherwise.

### GetAllowedIPsOk

`func (o *AddImageBuilderModel) GetAllowedIPsOk() (*[]string, bool)`

GetAllowedIPsOk returns a tuple with the AllowedIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIPs

`func (o *AddImageBuilderModel) SetAllowedIPs(v []string)`

SetAllowedIPs sets AllowedIPs field to given value.

### HasAllowedIPs

`func (o *AddImageBuilderModel) HasAllowedIPs() bool`

HasAllowedIPs returns a boolean if a field has been set.

### GetCspCustomerId

`func (o *AddImageBuilderModel) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *AddImageBuilderModel) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *AddImageBuilderModel) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *AddImageBuilderModel) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### GetCspSiteId

`func (o *AddImageBuilderModel) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *AddImageBuilderModel) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *AddImageBuilderModel) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *AddImageBuilderModel) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


