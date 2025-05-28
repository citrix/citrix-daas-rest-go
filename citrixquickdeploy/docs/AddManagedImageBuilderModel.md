# AddManagedImageBuilderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Friendly name of the template | 
**ImageId** | **string** | ID of the image we will be building off of | 
**ManagedSubscriptionId** | Pointer to **string** | ID of the managed subscription that the image will be placed in. | [optional] 
**IsDomainJoined** | Pointer to **bool** | Indicates if the builder vm will be domain joined | [optional] 
**AzureRegion** | Pointer to **string** | Azure region to deploy image builder vm that does not have on-prem connectivity | [optional] 
**VnetPeeringId** | Pointer to **string** | ID of the vnet peering the image builder vm will be associated with | [optional] 
**AzureVpnId** | Pointer to **string** | ID of the Azure VPN the image builder will be associated with | [optional] 
**DomainName** | Pointer to **string** | Name of the Domain the Image Builder will join | [optional] 
**OrganizationalUnit** | Pointer to **string** | The OU to associate the Image Builder VM with | [optional] 
**VmType** | **string** | The type of VM Instance type | 
**ServiceAccountName** | **string** | The service account used to join the Image Builder VM to the domain or local account that will be created for non-domain joined images | 
**ServiceAccountPassword** | **string** | The service account password | 
**Notes** | Pointer to **string** | Customer notes about template image | [optional] 
**AllowedIPs** | Pointer to **[]string** | Ip Addresses allowed to RDP | [optional] 
**CspCustomerId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 

## Methods

### NewAddManagedImageBuilderModel

`func NewAddManagedImageBuilderModel(name string, imageId string, vmType string, serviceAccountName string, serviceAccountPassword string, ) *AddManagedImageBuilderModel`

NewAddManagedImageBuilderModel instantiates a new AddManagedImageBuilderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddManagedImageBuilderModelWithDefaults

`func NewAddManagedImageBuilderModelWithDefaults() *AddManagedImageBuilderModel`

NewAddManagedImageBuilderModelWithDefaults instantiates a new AddManagedImageBuilderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddManagedImageBuilderModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddManagedImageBuilderModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddManagedImageBuilderModel) SetName(v string)`

SetName sets Name field to given value.


### GetImageId

`func (o *AddManagedImageBuilderModel) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AddManagedImageBuilderModel) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AddManagedImageBuilderModel) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetManagedSubscriptionId

`func (o *AddManagedImageBuilderModel) GetManagedSubscriptionId() string`

GetManagedSubscriptionId returns the ManagedSubscriptionId field if non-nil, zero value otherwise.

### GetManagedSubscriptionIdOk

`func (o *AddManagedImageBuilderModel) GetManagedSubscriptionIdOk() (*string, bool)`

GetManagedSubscriptionIdOk returns a tuple with the ManagedSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSubscriptionId

`func (o *AddManagedImageBuilderModel) SetManagedSubscriptionId(v string)`

SetManagedSubscriptionId sets ManagedSubscriptionId field to given value.

### HasManagedSubscriptionId

`func (o *AddManagedImageBuilderModel) HasManagedSubscriptionId() bool`

HasManagedSubscriptionId returns a boolean if a field has been set.

### GetIsDomainJoined

`func (o *AddManagedImageBuilderModel) GetIsDomainJoined() bool`

GetIsDomainJoined returns the IsDomainJoined field if non-nil, zero value otherwise.

### GetIsDomainJoinedOk

`func (o *AddManagedImageBuilderModel) GetIsDomainJoinedOk() (*bool, bool)`

GetIsDomainJoinedOk returns a tuple with the IsDomainJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomainJoined

`func (o *AddManagedImageBuilderModel) SetIsDomainJoined(v bool)`

SetIsDomainJoined sets IsDomainJoined field to given value.

### HasIsDomainJoined

`func (o *AddManagedImageBuilderModel) HasIsDomainJoined() bool`

HasIsDomainJoined returns a boolean if a field has been set.

### GetAzureRegion

`func (o *AddManagedImageBuilderModel) GetAzureRegion() string`

GetAzureRegion returns the AzureRegion field if non-nil, zero value otherwise.

### GetAzureRegionOk

`func (o *AddManagedImageBuilderModel) GetAzureRegionOk() (*string, bool)`

GetAzureRegionOk returns a tuple with the AzureRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRegion

`func (o *AddManagedImageBuilderModel) SetAzureRegion(v string)`

SetAzureRegion sets AzureRegion field to given value.

### HasAzureRegion

`func (o *AddManagedImageBuilderModel) HasAzureRegion() bool`

HasAzureRegion returns a boolean if a field has been set.

### GetVnetPeeringId

`func (o *AddManagedImageBuilderModel) GetVnetPeeringId() string`

GetVnetPeeringId returns the VnetPeeringId field if non-nil, zero value otherwise.

### GetVnetPeeringIdOk

`func (o *AddManagedImageBuilderModel) GetVnetPeeringIdOk() (*string, bool)`

GetVnetPeeringIdOk returns a tuple with the VnetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetPeeringId

`func (o *AddManagedImageBuilderModel) SetVnetPeeringId(v string)`

SetVnetPeeringId sets VnetPeeringId field to given value.

### HasVnetPeeringId

`func (o *AddManagedImageBuilderModel) HasVnetPeeringId() bool`

HasVnetPeeringId returns a boolean if a field has been set.

### GetAzureVpnId

`func (o *AddManagedImageBuilderModel) GetAzureVpnId() string`

GetAzureVpnId returns the AzureVpnId field if non-nil, zero value otherwise.

### GetAzureVpnIdOk

`func (o *AddManagedImageBuilderModel) GetAzureVpnIdOk() (*string, bool)`

GetAzureVpnIdOk returns a tuple with the AzureVpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVpnId

`func (o *AddManagedImageBuilderModel) SetAzureVpnId(v string)`

SetAzureVpnId sets AzureVpnId field to given value.

### HasAzureVpnId

`func (o *AddManagedImageBuilderModel) HasAzureVpnId() bool`

HasAzureVpnId returns a boolean if a field has been set.

### GetDomainName

`func (o *AddManagedImageBuilderModel) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AddManagedImageBuilderModel) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *AddManagedImageBuilderModel) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *AddManagedImageBuilderModel) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *AddManagedImageBuilderModel) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *AddManagedImageBuilderModel) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *AddManagedImageBuilderModel) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *AddManagedImageBuilderModel) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetVmType

`func (o *AddManagedImageBuilderModel) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *AddManagedImageBuilderModel) GetVmTypeOk() (*string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmType

`func (o *AddManagedImageBuilderModel) SetVmType(v string)`

SetVmType sets VmType field to given value.


### GetServiceAccountName

`func (o *AddManagedImageBuilderModel) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *AddManagedImageBuilderModel) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *AddManagedImageBuilderModel) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.


### GetServiceAccountPassword

`func (o *AddManagedImageBuilderModel) GetServiceAccountPassword() string`

GetServiceAccountPassword returns the ServiceAccountPassword field if non-nil, zero value otherwise.

### GetServiceAccountPasswordOk

`func (o *AddManagedImageBuilderModel) GetServiceAccountPasswordOk() (*string, bool)`

GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountPassword

`func (o *AddManagedImageBuilderModel) SetServiceAccountPassword(v string)`

SetServiceAccountPassword sets ServiceAccountPassword field to given value.


### GetNotes

`func (o *AddManagedImageBuilderModel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddManagedImageBuilderModel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddManagedImageBuilderModel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddManagedImageBuilderModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAllowedIPs

`func (o *AddManagedImageBuilderModel) GetAllowedIPs() []string`

GetAllowedIPs returns the AllowedIPs field if non-nil, zero value otherwise.

### GetAllowedIPsOk

`func (o *AddManagedImageBuilderModel) GetAllowedIPsOk() (*[]string, bool)`

GetAllowedIPsOk returns a tuple with the AllowedIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIPs

`func (o *AddManagedImageBuilderModel) SetAllowedIPs(v []string)`

SetAllowedIPs sets AllowedIPs field to given value.

### HasAllowedIPs

`func (o *AddManagedImageBuilderModel) HasAllowedIPs() bool`

HasAllowedIPs returns a boolean if a field has been set.

### GetCspCustomerId

`func (o *AddManagedImageBuilderModel) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *AddManagedImageBuilderModel) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *AddManagedImageBuilderModel) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *AddManagedImageBuilderModel) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### GetCspSiteId

`func (o *AddManagedImageBuilderModel) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *AddManagedImageBuilderModel) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *AddManagedImageBuilderModel) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *AddManagedImageBuilderModel) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


