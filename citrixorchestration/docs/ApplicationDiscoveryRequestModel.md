# ApplicationDiscoveryRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationGroups** | Pointer to **[]string** | The machines associated with Application Groups will be browsed. | [optional] 
**DeliveryGroups** | Pointer to **[]string** | The machines associated with Delivery Groups will be browsed. | [optional] 
**MachineCatalogs** | Pointer to **[]string** | The machines in the catalogs will be browsed. | [optional] 
**ServerOrPath** | Pointer to **NullableString** | Specify machine IP or DNS to get machine shares; Specify local path to get directory/files. | [optional] 

## Methods

### NewApplicationDiscoveryRequestModel

`func NewApplicationDiscoveryRequestModel() *ApplicationDiscoveryRequestModel`

NewApplicationDiscoveryRequestModel instantiates a new ApplicationDiscoveryRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDiscoveryRequestModelWithDefaults

`func NewApplicationDiscoveryRequestModelWithDefaults() *ApplicationDiscoveryRequestModel`

NewApplicationDiscoveryRequestModelWithDefaults instantiates a new ApplicationDiscoveryRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationGroups

`func (o *ApplicationDiscoveryRequestModel) GetApplicationGroups() []string`

GetApplicationGroups returns the ApplicationGroups field if non-nil, zero value otherwise.

### GetApplicationGroupsOk

`func (o *ApplicationDiscoveryRequestModel) GetApplicationGroupsOk() (*[]string, bool)`

GetApplicationGroupsOk returns a tuple with the ApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroups

`func (o *ApplicationDiscoveryRequestModel) SetApplicationGroups(v []string)`

SetApplicationGroups sets ApplicationGroups field to given value.

### HasApplicationGroups

`func (o *ApplicationDiscoveryRequestModel) HasApplicationGroups() bool`

HasApplicationGroups returns a boolean if a field has been set.

### SetApplicationGroupsNil

`func (o *ApplicationDiscoveryRequestModel) SetApplicationGroupsNil(b bool)`

 SetApplicationGroupsNil sets the value for ApplicationGroups to be an explicit nil

### UnsetApplicationGroups
`func (o *ApplicationDiscoveryRequestModel) UnsetApplicationGroups()`

UnsetApplicationGroups ensures that no value is present for ApplicationGroups, not even an explicit nil
### GetDeliveryGroups

`func (o *ApplicationDiscoveryRequestModel) GetDeliveryGroups() []string`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *ApplicationDiscoveryRequestModel) GetDeliveryGroupsOk() (*[]string, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *ApplicationDiscoveryRequestModel) SetDeliveryGroups(v []string)`

SetDeliveryGroups sets DeliveryGroups field to given value.

### HasDeliveryGroups

`func (o *ApplicationDiscoveryRequestModel) HasDeliveryGroups() bool`

HasDeliveryGroups returns a boolean if a field has been set.

### SetDeliveryGroupsNil

`func (o *ApplicationDiscoveryRequestModel) SetDeliveryGroupsNil(b bool)`

 SetDeliveryGroupsNil sets the value for DeliveryGroups to be an explicit nil

### UnsetDeliveryGroups
`func (o *ApplicationDiscoveryRequestModel) UnsetDeliveryGroups()`

UnsetDeliveryGroups ensures that no value is present for DeliveryGroups, not even an explicit nil
### GetMachineCatalogs

`func (o *ApplicationDiscoveryRequestModel) GetMachineCatalogs() []string`

GetMachineCatalogs returns the MachineCatalogs field if non-nil, zero value otherwise.

### GetMachineCatalogsOk

`func (o *ApplicationDiscoveryRequestModel) GetMachineCatalogsOk() (*[]string, bool)`

GetMachineCatalogsOk returns a tuple with the MachineCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCatalogs

`func (o *ApplicationDiscoveryRequestModel) SetMachineCatalogs(v []string)`

SetMachineCatalogs sets MachineCatalogs field to given value.

### HasMachineCatalogs

`func (o *ApplicationDiscoveryRequestModel) HasMachineCatalogs() bool`

HasMachineCatalogs returns a boolean if a field has been set.

### SetMachineCatalogsNil

`func (o *ApplicationDiscoveryRequestModel) SetMachineCatalogsNil(b bool)`

 SetMachineCatalogsNil sets the value for MachineCatalogs to be an explicit nil

### UnsetMachineCatalogs
`func (o *ApplicationDiscoveryRequestModel) UnsetMachineCatalogs()`

UnsetMachineCatalogs ensures that no value is present for MachineCatalogs, not even an explicit nil
### GetServerOrPath

`func (o *ApplicationDiscoveryRequestModel) GetServerOrPath() string`

GetServerOrPath returns the ServerOrPath field if non-nil, zero value otherwise.

### GetServerOrPathOk

`func (o *ApplicationDiscoveryRequestModel) GetServerOrPathOk() (*string, bool)`

GetServerOrPathOk returns a tuple with the ServerOrPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOrPath

`func (o *ApplicationDiscoveryRequestModel) SetServerOrPath(v string)`

SetServerOrPath sets ServerOrPath field to given value.

### HasServerOrPath

`func (o *ApplicationDiscoveryRequestModel) HasServerOrPath() bool`

HasServerOrPath returns a boolean if a field has been set.

### SetServerOrPathNil

`func (o *ApplicationDiscoveryRequestModel) SetServerOrPathNil(b bool)`

 SetServerOrPathNil sets the value for ServerOrPath to be an explicit nil

### UnsetServerOrPath
`func (o *ApplicationDiscoveryRequestModel) UnsetServerOrPath()`

UnsetServerOrPath ensures that no value is present for ServerOrPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


