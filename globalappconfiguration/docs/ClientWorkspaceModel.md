# ClientWorkspaceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GacsDiscovery** | Pointer to [**ClientDiscoveryModel**](ClientDiscoveryModel.md) |  | [optional] 
**GacsSettingsEndpoints** | Pointer to [**[]ClientSettingsEndpointsModel**](ClientSettingsEndpointsModel.md) |  | [optional] 

## Methods

### NewClientWorkspaceModel

`func NewClientWorkspaceModel() *ClientWorkspaceModel`

NewClientWorkspaceModel instantiates a new ClientWorkspaceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWorkspaceModelWithDefaults

`func NewClientWorkspaceModelWithDefaults() *ClientWorkspaceModel`

NewClientWorkspaceModelWithDefaults instantiates a new ClientWorkspaceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGacsDiscovery

`func (o *ClientWorkspaceModel) GetGacsDiscovery() ClientDiscoveryModel`

GetGacsDiscovery returns the GacsDiscovery field if non-nil, zero value otherwise.

### GetGacsDiscoveryOk

`func (o *ClientWorkspaceModel) GetGacsDiscoveryOk() (*ClientDiscoveryModel, bool)`

GetGacsDiscoveryOk returns a tuple with the GacsDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGacsDiscovery

`func (o *ClientWorkspaceModel) SetGacsDiscovery(v ClientDiscoveryModel)`

SetGacsDiscovery sets GacsDiscovery field to given value.

### HasGacsDiscovery

`func (o *ClientWorkspaceModel) HasGacsDiscovery() bool`

HasGacsDiscovery returns a boolean if a field has been set.

### GetGacsSettingsEndpoints

`func (o *ClientWorkspaceModel) GetGacsSettingsEndpoints() []ClientSettingsEndpointsModel`

GetGacsSettingsEndpoints returns the GacsSettingsEndpoints field if non-nil, zero value otherwise.

### GetGacsSettingsEndpointsOk

`func (o *ClientWorkspaceModel) GetGacsSettingsEndpointsOk() (*[]ClientSettingsEndpointsModel, bool)`

GetGacsSettingsEndpointsOk returns a tuple with the GacsSettingsEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGacsSettingsEndpoints

`func (o *ClientWorkspaceModel) SetGacsSettingsEndpoints(v []ClientSettingsEndpointsModel)`

SetGacsSettingsEndpoints sets GacsSettingsEndpoints field to given value.

### HasGacsSettingsEndpoints

`func (o *ClientWorkspaceModel) HasGacsSettingsEndpoints() bool`

HasGacsSettingsEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


