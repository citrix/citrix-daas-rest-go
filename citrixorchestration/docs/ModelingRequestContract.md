# ModelingRequestContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainController** | Pointer to **string** | The domain controller | [optional] 
**ComputerIdentity** | Pointer to **string** | The computer | [optional] 
**UserIdentity** | Pointer to **string** | The user | [optional] 
**SiteName** | Pointer to **string** | The site name | [optional] 
**Computer** | Pointer to **string** | The selected computer | [optional] 
**ComputerContainer** | Pointer to **string** | Selected computer container | [optional] 
**User** | Pointer to **string** | Selected user | [optional] 
**UserSid** | Pointer to **string** | SID of the selected user. | [optional] 
**DistinguishedName** | Pointer to **string** | User distinguished name. | [optional] 
**GroupSids** | Pointer to **[]string** | SIDs of the groups to which the user belongs. | [optional] 
**UserContainer** | Pointer to **string** | Selected user container | [optional] 
**ClientIPAddress** | Pointer to **string** | Client IP address | [optional] 
**ClientName** | Pointer to **string** | Client name | [optional] 
**DeliveryGroup** | Pointer to **string** | Delivery group | [optional] 
**DeliveryGroupType** | Pointer to **string** | Delivery group type, values are members of enum DesktopKind. | [optional] 
**DeliveryGroupTags** | Pointer to **[]string** | Delivery group tags | [optional] 
**IsUsingAccessGateway** | Pointer to **bool** | Using Access Gateway | [optional] 
**AccessGatewayFarm** | Pointer to **string** | Access Gateway farm | [optional] 
**AccessGatewayTags** | Pointer to **[]string** | Access Gateway tags | [optional] 
**IsUsingWanScaler** | Pointer to **bool** | Using WanScaler (NetScaler SD-WAN) | [optional] 

## Methods

### NewModelingRequestContract

`func NewModelingRequestContract() *ModelingRequestContract`

NewModelingRequestContract instantiates a new ModelingRequestContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelingRequestContractWithDefaults

`func NewModelingRequestContractWithDefaults() *ModelingRequestContract`

NewModelingRequestContractWithDefaults instantiates a new ModelingRequestContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainController

`func (o *ModelingRequestContract) GetDomainController() string`

GetDomainController returns the DomainController field if non-nil, zero value otherwise.

### GetDomainControllerOk

`func (o *ModelingRequestContract) GetDomainControllerOk() (*string, bool)`

GetDomainControllerOk returns a tuple with the DomainController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainController

`func (o *ModelingRequestContract) SetDomainController(v string)`

SetDomainController sets DomainController field to given value.

### HasDomainController

`func (o *ModelingRequestContract) HasDomainController() bool`

HasDomainController returns a boolean if a field has been set.

### GetComputerIdentity

`func (o *ModelingRequestContract) GetComputerIdentity() string`

GetComputerIdentity returns the ComputerIdentity field if non-nil, zero value otherwise.

### GetComputerIdentityOk

`func (o *ModelingRequestContract) GetComputerIdentityOk() (*string, bool)`

GetComputerIdentityOk returns a tuple with the ComputerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerIdentity

`func (o *ModelingRequestContract) SetComputerIdentity(v string)`

SetComputerIdentity sets ComputerIdentity field to given value.

### HasComputerIdentity

`func (o *ModelingRequestContract) HasComputerIdentity() bool`

HasComputerIdentity returns a boolean if a field has been set.

### GetUserIdentity

`func (o *ModelingRequestContract) GetUserIdentity() string`

GetUserIdentity returns the UserIdentity field if non-nil, zero value otherwise.

### GetUserIdentityOk

`func (o *ModelingRequestContract) GetUserIdentityOk() (*string, bool)`

GetUserIdentityOk returns a tuple with the UserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentity

`func (o *ModelingRequestContract) SetUserIdentity(v string)`

SetUserIdentity sets UserIdentity field to given value.

### HasUserIdentity

`func (o *ModelingRequestContract) HasUserIdentity() bool`

HasUserIdentity returns a boolean if a field has been set.

### GetSiteName

`func (o *ModelingRequestContract) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *ModelingRequestContract) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *ModelingRequestContract) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *ModelingRequestContract) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetComputer

`func (o *ModelingRequestContract) GetComputer() string`

GetComputer returns the Computer field if non-nil, zero value otherwise.

### GetComputerOk

`func (o *ModelingRequestContract) GetComputerOk() (*string, bool)`

GetComputerOk returns a tuple with the Computer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputer

`func (o *ModelingRequestContract) SetComputer(v string)`

SetComputer sets Computer field to given value.

### HasComputer

`func (o *ModelingRequestContract) HasComputer() bool`

HasComputer returns a boolean if a field has been set.

### GetComputerContainer

`func (o *ModelingRequestContract) GetComputerContainer() string`

GetComputerContainer returns the ComputerContainer field if non-nil, zero value otherwise.

### GetComputerContainerOk

`func (o *ModelingRequestContract) GetComputerContainerOk() (*string, bool)`

GetComputerContainerOk returns a tuple with the ComputerContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerContainer

`func (o *ModelingRequestContract) SetComputerContainer(v string)`

SetComputerContainer sets ComputerContainer field to given value.

### HasComputerContainer

`func (o *ModelingRequestContract) HasComputerContainer() bool`

HasComputerContainer returns a boolean if a field has been set.

### GetUser

`func (o *ModelingRequestContract) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelingRequestContract) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelingRequestContract) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelingRequestContract) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserSid

`func (o *ModelingRequestContract) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *ModelingRequestContract) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *ModelingRequestContract) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.

### HasUserSid

`func (o *ModelingRequestContract) HasUserSid() bool`

HasUserSid returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *ModelingRequestContract) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *ModelingRequestContract) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *ModelingRequestContract) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *ModelingRequestContract) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetGroupSids

`func (o *ModelingRequestContract) GetGroupSids() []string`

GetGroupSids returns the GroupSids field if non-nil, zero value otherwise.

### GetGroupSidsOk

`func (o *ModelingRequestContract) GetGroupSidsOk() (*[]string, bool)`

GetGroupSidsOk returns a tuple with the GroupSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSids

`func (o *ModelingRequestContract) SetGroupSids(v []string)`

SetGroupSids sets GroupSids field to given value.

### HasGroupSids

`func (o *ModelingRequestContract) HasGroupSids() bool`

HasGroupSids returns a boolean if a field has been set.

### GetUserContainer

`func (o *ModelingRequestContract) GetUserContainer() string`

GetUserContainer returns the UserContainer field if non-nil, zero value otherwise.

### GetUserContainerOk

`func (o *ModelingRequestContract) GetUserContainerOk() (*string, bool)`

GetUserContainerOk returns a tuple with the UserContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserContainer

`func (o *ModelingRequestContract) SetUserContainer(v string)`

SetUserContainer sets UserContainer field to given value.

### HasUserContainer

`func (o *ModelingRequestContract) HasUserContainer() bool`

HasUserContainer returns a boolean if a field has been set.

### GetClientIPAddress

`func (o *ModelingRequestContract) GetClientIPAddress() string`

GetClientIPAddress returns the ClientIPAddress field if non-nil, zero value otherwise.

### GetClientIPAddressOk

`func (o *ModelingRequestContract) GetClientIPAddressOk() (*string, bool)`

GetClientIPAddressOk returns a tuple with the ClientIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIPAddress

`func (o *ModelingRequestContract) SetClientIPAddress(v string)`

SetClientIPAddress sets ClientIPAddress field to given value.

### HasClientIPAddress

`func (o *ModelingRequestContract) HasClientIPAddress() bool`

HasClientIPAddress returns a boolean if a field has been set.

### GetClientName

`func (o *ModelingRequestContract) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ModelingRequestContract) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ModelingRequestContract) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ModelingRequestContract) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDeliveryGroup

`func (o *ModelingRequestContract) GetDeliveryGroup() string`

GetDeliveryGroup returns the DeliveryGroup field if non-nil, zero value otherwise.

### GetDeliveryGroupOk

`func (o *ModelingRequestContract) GetDeliveryGroupOk() (*string, bool)`

GetDeliveryGroupOk returns a tuple with the DeliveryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroup

`func (o *ModelingRequestContract) SetDeliveryGroup(v string)`

SetDeliveryGroup sets DeliveryGroup field to given value.

### HasDeliveryGroup

`func (o *ModelingRequestContract) HasDeliveryGroup() bool`

HasDeliveryGroup returns a boolean if a field has been set.

### GetDeliveryGroupType

`func (o *ModelingRequestContract) GetDeliveryGroupType() string`

GetDeliveryGroupType returns the DeliveryGroupType field if non-nil, zero value otherwise.

### GetDeliveryGroupTypeOk

`func (o *ModelingRequestContract) GetDeliveryGroupTypeOk() (*string, bool)`

GetDeliveryGroupTypeOk returns a tuple with the DeliveryGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroupType

`func (o *ModelingRequestContract) SetDeliveryGroupType(v string)`

SetDeliveryGroupType sets DeliveryGroupType field to given value.

### HasDeliveryGroupType

`func (o *ModelingRequestContract) HasDeliveryGroupType() bool`

HasDeliveryGroupType returns a boolean if a field has been set.

### GetDeliveryGroupTags

`func (o *ModelingRequestContract) GetDeliveryGroupTags() []string`

GetDeliveryGroupTags returns the DeliveryGroupTags field if non-nil, zero value otherwise.

### GetDeliveryGroupTagsOk

`func (o *ModelingRequestContract) GetDeliveryGroupTagsOk() (*[]string, bool)`

GetDeliveryGroupTagsOk returns a tuple with the DeliveryGroupTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroupTags

`func (o *ModelingRequestContract) SetDeliveryGroupTags(v []string)`

SetDeliveryGroupTags sets DeliveryGroupTags field to given value.

### HasDeliveryGroupTags

`func (o *ModelingRequestContract) HasDeliveryGroupTags() bool`

HasDeliveryGroupTags returns a boolean if a field has been set.

### GetIsUsingAccessGateway

`func (o *ModelingRequestContract) GetIsUsingAccessGateway() bool`

GetIsUsingAccessGateway returns the IsUsingAccessGateway field if non-nil, zero value otherwise.

### GetIsUsingAccessGatewayOk

`func (o *ModelingRequestContract) GetIsUsingAccessGatewayOk() (*bool, bool)`

GetIsUsingAccessGatewayOk returns a tuple with the IsUsingAccessGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingAccessGateway

`func (o *ModelingRequestContract) SetIsUsingAccessGateway(v bool)`

SetIsUsingAccessGateway sets IsUsingAccessGateway field to given value.

### HasIsUsingAccessGateway

`func (o *ModelingRequestContract) HasIsUsingAccessGateway() bool`

HasIsUsingAccessGateway returns a boolean if a field has been set.

### GetAccessGatewayFarm

`func (o *ModelingRequestContract) GetAccessGatewayFarm() string`

GetAccessGatewayFarm returns the AccessGatewayFarm field if non-nil, zero value otherwise.

### GetAccessGatewayFarmOk

`func (o *ModelingRequestContract) GetAccessGatewayFarmOk() (*string, bool)`

GetAccessGatewayFarmOk returns a tuple with the AccessGatewayFarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGatewayFarm

`func (o *ModelingRequestContract) SetAccessGatewayFarm(v string)`

SetAccessGatewayFarm sets AccessGatewayFarm field to given value.

### HasAccessGatewayFarm

`func (o *ModelingRequestContract) HasAccessGatewayFarm() bool`

HasAccessGatewayFarm returns a boolean if a field has been set.

### GetAccessGatewayTags

`func (o *ModelingRequestContract) GetAccessGatewayTags() []string`

GetAccessGatewayTags returns the AccessGatewayTags field if non-nil, zero value otherwise.

### GetAccessGatewayTagsOk

`func (o *ModelingRequestContract) GetAccessGatewayTagsOk() (*[]string, bool)`

GetAccessGatewayTagsOk returns a tuple with the AccessGatewayTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGatewayTags

`func (o *ModelingRequestContract) SetAccessGatewayTags(v []string)`

SetAccessGatewayTags sets AccessGatewayTags field to given value.

### HasAccessGatewayTags

`func (o *ModelingRequestContract) HasAccessGatewayTags() bool`

HasAccessGatewayTags returns a boolean if a field has been set.

### GetIsUsingWanScaler

`func (o *ModelingRequestContract) GetIsUsingWanScaler() bool`

GetIsUsingWanScaler returns the IsUsingWanScaler field if non-nil, zero value otherwise.

### GetIsUsingWanScalerOk

`func (o *ModelingRequestContract) GetIsUsingWanScalerOk() (*bool, bool)`

GetIsUsingWanScalerOk returns a tuple with the IsUsingWanScaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingWanScaler

`func (o *ModelingRequestContract) SetIsUsingWanScaler(v bool)`

SetIsUsingWanScaler sets IsUsingWanScaler field to given value.

### HasIsUsingWanScaler

`func (o *ModelingRequestContract) HasIsUsingWanScaler() bool`

HasIsUsingWanScaler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


