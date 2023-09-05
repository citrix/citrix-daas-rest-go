# MachineCatalogDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentVersion** | Pointer to **string** | Version of the Citrix Virtual Delivery Agent (VDA) installed on the machine. | [optional] 
**HypervisorConnection** | Pointer to [**MachineCatalogDetailResponseModelAllOfHypervisorConnection**](MachineCatalogDetailResponseModelAllOfHypervisorConnection.md) |  | [optional] 
**OSType** | Pointer to **string** | A string that can be used to identify the operating system that is running on machines in the catalog. | [optional] 
**OSVersion** | Pointer to **string** | A string that can be used to identify the version of the operating system running on the machine, if known. | [optional] 
**DeliveryGroups** | [**[]MachineCatalogDeliveryGroupRefResponseModel**](MachineCatalogDeliveryGroupRefResponseModel.md) | Delivery groups associated with this catalog. | 
**UpgradeDetail** | Pointer to [**MachineCatalogDetailResponseModelAllOfUpgradeDetail**](MachineCatalogDetailResponseModelAllOfUpgradeDetail.md) |  | [optional] 

## Methods

### NewMachineCatalogDetailResponseModelAllOf

`func NewMachineCatalogDetailResponseModelAllOf(deliveryGroups []MachineCatalogDeliveryGroupRefResponseModel, ) *MachineCatalogDetailResponseModelAllOf`

NewMachineCatalogDetailResponseModelAllOf instantiates a new MachineCatalogDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogDetailResponseModelAllOfWithDefaults

`func NewMachineCatalogDetailResponseModelAllOfWithDefaults() *MachineCatalogDetailResponseModelAllOf`

NewMachineCatalogDetailResponseModelAllOfWithDefaults instantiates a new MachineCatalogDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentVersion

`func (o *MachineCatalogDetailResponseModelAllOf) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *MachineCatalogDetailResponseModelAllOf) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *MachineCatalogDetailResponseModelAllOf) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *MachineCatalogDetailResponseModelAllOf) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetHypervisorConnection

`func (o *MachineCatalogDetailResponseModelAllOf) GetHypervisorConnection() MachineCatalogDetailResponseModelAllOfHypervisorConnection`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *MachineCatalogDetailResponseModelAllOf) GetHypervisorConnectionOk() (*MachineCatalogDetailResponseModelAllOfHypervisorConnection, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *MachineCatalogDetailResponseModelAllOf) SetHypervisorConnection(v MachineCatalogDetailResponseModelAllOfHypervisorConnection)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *MachineCatalogDetailResponseModelAllOf) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### GetOSType

`func (o *MachineCatalogDetailResponseModelAllOf) GetOSType() string`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *MachineCatalogDetailResponseModelAllOf) GetOSTypeOk() (*string, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *MachineCatalogDetailResponseModelAllOf) SetOSType(v string)`

SetOSType sets OSType field to given value.

### HasOSType

`func (o *MachineCatalogDetailResponseModelAllOf) HasOSType() bool`

HasOSType returns a boolean if a field has been set.

### GetOSVersion

`func (o *MachineCatalogDetailResponseModelAllOf) GetOSVersion() string`

GetOSVersion returns the OSVersion field if non-nil, zero value otherwise.

### GetOSVersionOk

`func (o *MachineCatalogDetailResponseModelAllOf) GetOSVersionOk() (*string, bool)`

GetOSVersionOk returns a tuple with the OSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSVersion

`func (o *MachineCatalogDetailResponseModelAllOf) SetOSVersion(v string)`

SetOSVersion sets OSVersion field to given value.

### HasOSVersion

`func (o *MachineCatalogDetailResponseModelAllOf) HasOSVersion() bool`

HasOSVersion returns a boolean if a field has been set.

### GetDeliveryGroups

`func (o *MachineCatalogDetailResponseModelAllOf) GetDeliveryGroups() []MachineCatalogDeliveryGroupRefResponseModel`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *MachineCatalogDetailResponseModelAllOf) GetDeliveryGroupsOk() (*[]MachineCatalogDeliveryGroupRefResponseModel, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *MachineCatalogDetailResponseModelAllOf) SetDeliveryGroups(v []MachineCatalogDeliveryGroupRefResponseModel)`

SetDeliveryGroups sets DeliveryGroups field to given value.


### GetUpgradeDetail

`func (o *MachineCatalogDetailResponseModelAllOf) GetUpgradeDetail() MachineCatalogDetailResponseModelAllOfUpgradeDetail`

GetUpgradeDetail returns the UpgradeDetail field if non-nil, zero value otherwise.

### GetUpgradeDetailOk

`func (o *MachineCatalogDetailResponseModelAllOf) GetUpgradeDetailOk() (*MachineCatalogDetailResponseModelAllOfUpgradeDetail, bool)`

GetUpgradeDetailOk returns a tuple with the UpgradeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDetail

`func (o *MachineCatalogDetailResponseModelAllOf) SetUpgradeDetail(v MachineCatalogDetailResponseModelAllOfUpgradeDetail)`

SetUpgradeDetail sets UpgradeDetail field to given value.

### HasUpgradeDetail

`func (o *MachineCatalogDetailResponseModelAllOf) HasUpgradeDetail() bool`

HasUpgradeDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


