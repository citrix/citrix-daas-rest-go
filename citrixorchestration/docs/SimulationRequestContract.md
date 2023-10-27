# SimulationRequestContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputerName** | Pointer to **NullableString** | Name of the selected computer. | [optional] 
**UserSid** | Pointer to **NullableString** | SID of the selected user. | [optional] 
**GroupSids** | Pointer to **[]string** | SIDs of the groups to which the user belongs. | [optional] 
**ClientIpAddress** | Pointer to **NullableString** | Client IP address. | [optional] 
**ClientName** | Pointer to **NullableString** | Client name | [optional] 
**DeliveryGroupGuid** | Pointer to **string** | Delivery group GUID. | [optional] 
**DeliveryGroupType** | Pointer to **NullableString** | Delivery group type, values are members of enum DesktopKind. | [optional] 
**OrganizationalUnit** | Pointer to **NullableString** | The OU the computer is in. | [optional] 
**Tags** | Pointer to **[]string** | Tags. | [optional] 
**IsUsingAccessGateway** | Pointer to **bool** | Using Access Gateway. | [optional] 
**AccessGatewayFarm** | Pointer to **NullableString** | Access Gateway farm. | [optional] 
**AccessGatewayTags** | Pointer to **[]string** | Access Gateway tags. | [optional] 
**IsUsingWanScaler** | Pointer to **bool** | Using WanScaler (NetScaler SD-WAN). | [optional] 

## Methods

### NewSimulationRequestContract

`func NewSimulationRequestContract() *SimulationRequestContract`

NewSimulationRequestContract instantiates a new SimulationRequestContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulationRequestContractWithDefaults

`func NewSimulationRequestContractWithDefaults() *SimulationRequestContract`

NewSimulationRequestContractWithDefaults instantiates a new SimulationRequestContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputerName

`func (o *SimulationRequestContract) GetComputerName() string`

GetComputerName returns the ComputerName field if non-nil, zero value otherwise.

### GetComputerNameOk

`func (o *SimulationRequestContract) GetComputerNameOk() (*string, bool)`

GetComputerNameOk returns a tuple with the ComputerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerName

`func (o *SimulationRequestContract) SetComputerName(v string)`

SetComputerName sets ComputerName field to given value.

### HasComputerName

`func (o *SimulationRequestContract) HasComputerName() bool`

HasComputerName returns a boolean if a field has been set.

### SetComputerNameNil

`func (o *SimulationRequestContract) SetComputerNameNil(b bool)`

 SetComputerNameNil sets the value for ComputerName to be an explicit nil

### UnsetComputerName
`func (o *SimulationRequestContract) UnsetComputerName()`

UnsetComputerName ensures that no value is present for ComputerName, not even an explicit nil
### GetUserSid

`func (o *SimulationRequestContract) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *SimulationRequestContract) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *SimulationRequestContract) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.

### HasUserSid

`func (o *SimulationRequestContract) HasUserSid() bool`

HasUserSid returns a boolean if a field has been set.

### SetUserSidNil

`func (o *SimulationRequestContract) SetUserSidNil(b bool)`

 SetUserSidNil sets the value for UserSid to be an explicit nil

### UnsetUserSid
`func (o *SimulationRequestContract) UnsetUserSid()`

UnsetUserSid ensures that no value is present for UserSid, not even an explicit nil
### GetGroupSids

`func (o *SimulationRequestContract) GetGroupSids() []string`

GetGroupSids returns the GroupSids field if non-nil, zero value otherwise.

### GetGroupSidsOk

`func (o *SimulationRequestContract) GetGroupSidsOk() (*[]string, bool)`

GetGroupSidsOk returns a tuple with the GroupSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSids

`func (o *SimulationRequestContract) SetGroupSids(v []string)`

SetGroupSids sets GroupSids field to given value.

### HasGroupSids

`func (o *SimulationRequestContract) HasGroupSids() bool`

HasGroupSids returns a boolean if a field has been set.

### SetGroupSidsNil

`func (o *SimulationRequestContract) SetGroupSidsNil(b bool)`

 SetGroupSidsNil sets the value for GroupSids to be an explicit nil

### UnsetGroupSids
`func (o *SimulationRequestContract) UnsetGroupSids()`

UnsetGroupSids ensures that no value is present for GroupSids, not even an explicit nil
### GetClientIpAddress

`func (o *SimulationRequestContract) GetClientIpAddress() string`

GetClientIpAddress returns the ClientIpAddress field if non-nil, zero value otherwise.

### GetClientIpAddressOk

`func (o *SimulationRequestContract) GetClientIpAddressOk() (*string, bool)`

GetClientIpAddressOk returns a tuple with the ClientIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAddress

`func (o *SimulationRequestContract) SetClientIpAddress(v string)`

SetClientIpAddress sets ClientIpAddress field to given value.

### HasClientIpAddress

`func (o *SimulationRequestContract) HasClientIpAddress() bool`

HasClientIpAddress returns a boolean if a field has been set.

### SetClientIpAddressNil

`func (o *SimulationRequestContract) SetClientIpAddressNil(b bool)`

 SetClientIpAddressNil sets the value for ClientIpAddress to be an explicit nil

### UnsetClientIpAddress
`func (o *SimulationRequestContract) UnsetClientIpAddress()`

UnsetClientIpAddress ensures that no value is present for ClientIpAddress, not even an explicit nil
### GetClientName

`func (o *SimulationRequestContract) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *SimulationRequestContract) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *SimulationRequestContract) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *SimulationRequestContract) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### SetClientNameNil

`func (o *SimulationRequestContract) SetClientNameNil(b bool)`

 SetClientNameNil sets the value for ClientName to be an explicit nil

### UnsetClientName
`func (o *SimulationRequestContract) UnsetClientName()`

UnsetClientName ensures that no value is present for ClientName, not even an explicit nil
### GetDeliveryGroupGuid

`func (o *SimulationRequestContract) GetDeliveryGroupGuid() string`

GetDeliveryGroupGuid returns the DeliveryGroupGuid field if non-nil, zero value otherwise.

### GetDeliveryGroupGuidOk

`func (o *SimulationRequestContract) GetDeliveryGroupGuidOk() (*string, bool)`

GetDeliveryGroupGuidOk returns a tuple with the DeliveryGroupGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroupGuid

`func (o *SimulationRequestContract) SetDeliveryGroupGuid(v string)`

SetDeliveryGroupGuid sets DeliveryGroupGuid field to given value.

### HasDeliveryGroupGuid

`func (o *SimulationRequestContract) HasDeliveryGroupGuid() bool`

HasDeliveryGroupGuid returns a boolean if a field has been set.

### GetDeliveryGroupType

`func (o *SimulationRequestContract) GetDeliveryGroupType() string`

GetDeliveryGroupType returns the DeliveryGroupType field if non-nil, zero value otherwise.

### GetDeliveryGroupTypeOk

`func (o *SimulationRequestContract) GetDeliveryGroupTypeOk() (*string, bool)`

GetDeliveryGroupTypeOk returns a tuple with the DeliveryGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroupType

`func (o *SimulationRequestContract) SetDeliveryGroupType(v string)`

SetDeliveryGroupType sets DeliveryGroupType field to given value.

### HasDeliveryGroupType

`func (o *SimulationRequestContract) HasDeliveryGroupType() bool`

HasDeliveryGroupType returns a boolean if a field has been set.

### SetDeliveryGroupTypeNil

`func (o *SimulationRequestContract) SetDeliveryGroupTypeNil(b bool)`

 SetDeliveryGroupTypeNil sets the value for DeliveryGroupType to be an explicit nil

### UnsetDeliveryGroupType
`func (o *SimulationRequestContract) UnsetDeliveryGroupType()`

UnsetDeliveryGroupType ensures that no value is present for DeliveryGroupType, not even an explicit nil
### GetOrganizationalUnit

`func (o *SimulationRequestContract) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *SimulationRequestContract) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *SimulationRequestContract) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *SimulationRequestContract) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### SetOrganizationalUnitNil

`func (o *SimulationRequestContract) SetOrganizationalUnitNil(b bool)`

 SetOrganizationalUnitNil sets the value for OrganizationalUnit to be an explicit nil

### UnsetOrganizationalUnit
`func (o *SimulationRequestContract) UnsetOrganizationalUnit()`

UnsetOrganizationalUnit ensures that no value is present for OrganizationalUnit, not even an explicit nil
### GetTags

`func (o *SimulationRequestContract) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SimulationRequestContract) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SimulationRequestContract) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SimulationRequestContract) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SimulationRequestContract) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SimulationRequestContract) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetIsUsingAccessGateway

`func (o *SimulationRequestContract) GetIsUsingAccessGateway() bool`

GetIsUsingAccessGateway returns the IsUsingAccessGateway field if non-nil, zero value otherwise.

### GetIsUsingAccessGatewayOk

`func (o *SimulationRequestContract) GetIsUsingAccessGatewayOk() (*bool, bool)`

GetIsUsingAccessGatewayOk returns a tuple with the IsUsingAccessGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingAccessGateway

`func (o *SimulationRequestContract) SetIsUsingAccessGateway(v bool)`

SetIsUsingAccessGateway sets IsUsingAccessGateway field to given value.

### HasIsUsingAccessGateway

`func (o *SimulationRequestContract) HasIsUsingAccessGateway() bool`

HasIsUsingAccessGateway returns a boolean if a field has been set.

### GetAccessGatewayFarm

`func (o *SimulationRequestContract) GetAccessGatewayFarm() string`

GetAccessGatewayFarm returns the AccessGatewayFarm field if non-nil, zero value otherwise.

### GetAccessGatewayFarmOk

`func (o *SimulationRequestContract) GetAccessGatewayFarmOk() (*string, bool)`

GetAccessGatewayFarmOk returns a tuple with the AccessGatewayFarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGatewayFarm

`func (o *SimulationRequestContract) SetAccessGatewayFarm(v string)`

SetAccessGatewayFarm sets AccessGatewayFarm field to given value.

### HasAccessGatewayFarm

`func (o *SimulationRequestContract) HasAccessGatewayFarm() bool`

HasAccessGatewayFarm returns a boolean if a field has been set.

### SetAccessGatewayFarmNil

`func (o *SimulationRequestContract) SetAccessGatewayFarmNil(b bool)`

 SetAccessGatewayFarmNil sets the value for AccessGatewayFarm to be an explicit nil

### UnsetAccessGatewayFarm
`func (o *SimulationRequestContract) UnsetAccessGatewayFarm()`

UnsetAccessGatewayFarm ensures that no value is present for AccessGatewayFarm, not even an explicit nil
### GetAccessGatewayTags

`func (o *SimulationRequestContract) GetAccessGatewayTags() []string`

GetAccessGatewayTags returns the AccessGatewayTags field if non-nil, zero value otherwise.

### GetAccessGatewayTagsOk

`func (o *SimulationRequestContract) GetAccessGatewayTagsOk() (*[]string, bool)`

GetAccessGatewayTagsOk returns a tuple with the AccessGatewayTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGatewayTags

`func (o *SimulationRequestContract) SetAccessGatewayTags(v []string)`

SetAccessGatewayTags sets AccessGatewayTags field to given value.

### HasAccessGatewayTags

`func (o *SimulationRequestContract) HasAccessGatewayTags() bool`

HasAccessGatewayTags returns a boolean if a field has been set.

### SetAccessGatewayTagsNil

`func (o *SimulationRequestContract) SetAccessGatewayTagsNil(b bool)`

 SetAccessGatewayTagsNil sets the value for AccessGatewayTags to be an explicit nil

### UnsetAccessGatewayTags
`func (o *SimulationRequestContract) UnsetAccessGatewayTags()`

UnsetAccessGatewayTags ensures that no value is present for AccessGatewayTags, not even an explicit nil
### GetIsUsingWanScaler

`func (o *SimulationRequestContract) GetIsUsingWanScaler() bool`

GetIsUsingWanScaler returns the IsUsingWanScaler field if non-nil, zero value otherwise.

### GetIsUsingWanScalerOk

`func (o *SimulationRequestContract) GetIsUsingWanScalerOk() (*bool, bool)`

GetIsUsingWanScalerOk returns a tuple with the IsUsingWanScaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingWanScaler

`func (o *SimulationRequestContract) SetIsUsingWanScaler(v bool)`

SetIsUsingWanScaler sets IsUsingWanScaler field to given value.

### HasIsUsingWanScaler

`func (o *SimulationRequestContract) HasIsUsingWanScaler() bool`

HasIsUsingWanScaler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


