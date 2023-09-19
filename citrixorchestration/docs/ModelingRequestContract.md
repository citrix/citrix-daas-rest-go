# ModelingRequestContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainController** | Pointer to **NullableString** | The domain controller | [optional] 
**ComputerIdentity** | Pointer to **NullableString** | The computer | [optional] 
**UserIdentity** | Pointer to **NullableString** | The user | [optional] 
**SiteName** | Pointer to **NullableString** | The site name | [optional] 
**Computer** | Pointer to **NullableString** | The selected computer | [optional] 
**ComputerContainer** | Pointer to **NullableString** | Selected computer container | [optional] 
**User** | Pointer to **NullableString** | Selected user | [optional] 
**UserSid** | Pointer to **NullableString** | SID of the selected user. | [optional] 
**DistinguishedName** | Pointer to **NullableString** | User distinguished name. | [optional] 
**GroupSids** | Pointer to **[]string** | SIDs of the groups to which the user belongs. | [optional] 
**UserContainer** | Pointer to **NullableString** | Selected user container | [optional] 
**ClientIPAddress** | Pointer to **NullableString** | Client IP address | [optional] 
**ClientName** | Pointer to **NullableString** | Client name | [optional] 
**DeliveryGroup** | Pointer to **NullableString** | Delivery group | [optional] 
**DeliveryGroupType** | Pointer to **NullableString** | Delivery group type, values are members of enum DesktopKind. | [optional] 
**DeliveryGroupTags** | Pointer to **[]string** | Delivery group tags | [optional] 
**IsUsingAccessGateway** | Pointer to **bool** | Using Access Gateway | [optional] 
**AccessGatewayFarm** | Pointer to **NullableString** | Access Gateway farm | [optional] 
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

### SetDomainControllerNil

`func (o *ModelingRequestContract) SetDomainControllerNil(b bool)`

 SetDomainControllerNil sets the value for DomainController to be an explicit nil

### UnsetDomainController
`func (o *ModelingRequestContract) UnsetDomainController()`

UnsetDomainController ensures that no value is present for DomainController, not even an explicit nil
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

### SetComputerIdentityNil

`func (o *ModelingRequestContract) SetComputerIdentityNil(b bool)`

 SetComputerIdentityNil sets the value for ComputerIdentity to be an explicit nil

### UnsetComputerIdentity
`func (o *ModelingRequestContract) UnsetComputerIdentity()`

UnsetComputerIdentity ensures that no value is present for ComputerIdentity, not even an explicit nil
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

### SetUserIdentityNil

`func (o *ModelingRequestContract) SetUserIdentityNil(b bool)`

 SetUserIdentityNil sets the value for UserIdentity to be an explicit nil

### UnsetUserIdentity
`func (o *ModelingRequestContract) UnsetUserIdentity()`

UnsetUserIdentity ensures that no value is present for UserIdentity, not even an explicit nil
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

### SetSiteNameNil

`func (o *ModelingRequestContract) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *ModelingRequestContract) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
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

### SetComputerNil

`func (o *ModelingRequestContract) SetComputerNil(b bool)`

 SetComputerNil sets the value for Computer to be an explicit nil

### UnsetComputer
`func (o *ModelingRequestContract) UnsetComputer()`

UnsetComputer ensures that no value is present for Computer, not even an explicit nil
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

### SetComputerContainerNil

`func (o *ModelingRequestContract) SetComputerContainerNil(b bool)`

 SetComputerContainerNil sets the value for ComputerContainer to be an explicit nil

### UnsetComputerContainer
`func (o *ModelingRequestContract) UnsetComputerContainer()`

UnsetComputerContainer ensures that no value is present for ComputerContainer, not even an explicit nil
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

### SetUserNil

`func (o *ModelingRequestContract) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ModelingRequestContract) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
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

### SetUserSidNil

`func (o *ModelingRequestContract) SetUserSidNil(b bool)`

 SetUserSidNil sets the value for UserSid to be an explicit nil

### UnsetUserSid
`func (o *ModelingRequestContract) UnsetUserSid()`

UnsetUserSid ensures that no value is present for UserSid, not even an explicit nil
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

### SetDistinguishedNameNil

`func (o *ModelingRequestContract) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *ModelingRequestContract) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
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

### SetGroupSidsNil

`func (o *ModelingRequestContract) SetGroupSidsNil(b bool)`

 SetGroupSidsNil sets the value for GroupSids to be an explicit nil

### UnsetGroupSids
`func (o *ModelingRequestContract) UnsetGroupSids()`

UnsetGroupSids ensures that no value is present for GroupSids, not even an explicit nil
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

### SetUserContainerNil

`func (o *ModelingRequestContract) SetUserContainerNil(b bool)`

 SetUserContainerNil sets the value for UserContainer to be an explicit nil

### UnsetUserContainer
`func (o *ModelingRequestContract) UnsetUserContainer()`

UnsetUserContainer ensures that no value is present for UserContainer, not even an explicit nil
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

### SetClientIPAddressNil

`func (o *ModelingRequestContract) SetClientIPAddressNil(b bool)`

 SetClientIPAddressNil sets the value for ClientIPAddress to be an explicit nil

### UnsetClientIPAddress
`func (o *ModelingRequestContract) UnsetClientIPAddress()`

UnsetClientIPAddress ensures that no value is present for ClientIPAddress, not even an explicit nil
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

### SetClientNameNil

`func (o *ModelingRequestContract) SetClientNameNil(b bool)`

 SetClientNameNil sets the value for ClientName to be an explicit nil

### UnsetClientName
`func (o *ModelingRequestContract) UnsetClientName()`

UnsetClientName ensures that no value is present for ClientName, not even an explicit nil
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

### SetDeliveryGroupNil

`func (o *ModelingRequestContract) SetDeliveryGroupNil(b bool)`

 SetDeliveryGroupNil sets the value for DeliveryGroup to be an explicit nil

### UnsetDeliveryGroup
`func (o *ModelingRequestContract) UnsetDeliveryGroup()`

UnsetDeliveryGroup ensures that no value is present for DeliveryGroup, not even an explicit nil
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

### SetDeliveryGroupTypeNil

`func (o *ModelingRequestContract) SetDeliveryGroupTypeNil(b bool)`

 SetDeliveryGroupTypeNil sets the value for DeliveryGroupType to be an explicit nil

### UnsetDeliveryGroupType
`func (o *ModelingRequestContract) UnsetDeliveryGroupType()`

UnsetDeliveryGroupType ensures that no value is present for DeliveryGroupType, not even an explicit nil
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

### SetDeliveryGroupTagsNil

`func (o *ModelingRequestContract) SetDeliveryGroupTagsNil(b bool)`

 SetDeliveryGroupTagsNil sets the value for DeliveryGroupTags to be an explicit nil

### UnsetDeliveryGroupTags
`func (o *ModelingRequestContract) UnsetDeliveryGroupTags()`

UnsetDeliveryGroupTags ensures that no value is present for DeliveryGroupTags, not even an explicit nil
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

### SetAccessGatewayFarmNil

`func (o *ModelingRequestContract) SetAccessGatewayFarmNil(b bool)`

 SetAccessGatewayFarmNil sets the value for AccessGatewayFarm to be an explicit nil

### UnsetAccessGatewayFarm
`func (o *ModelingRequestContract) UnsetAccessGatewayFarm()`

UnsetAccessGatewayFarm ensures that no value is present for AccessGatewayFarm, not even an explicit nil
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

### SetAccessGatewayTagsNil

`func (o *ModelingRequestContract) SetAccessGatewayTagsNil(b bool)`

 SetAccessGatewayTagsNil sets the value for AccessGatewayTags to be an explicit nil

### UnsetAccessGatewayTags
`func (o *ModelingRequestContract) UnsetAccessGatewayTags()`

UnsetAccessGatewayTags ensures that no value is present for AccessGatewayTags, not even an explicit nil
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


