# CreateOrUpdateRouteTableModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureRoutes** | [**[]AzureRoute**](AzureRoute.md) | The list of routes. | 
**DisableRoutePropagation** | Pointer to **bool** | Indicates whether disable the border gateway protocol route propagation. | [optional] 

## Methods

### NewCreateOrUpdateRouteTableModel

`func NewCreateOrUpdateRouteTableModel(azureRoutes []AzureRoute, ) *CreateOrUpdateRouteTableModel`

NewCreateOrUpdateRouteTableModel instantiates a new CreateOrUpdateRouteTableModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateRouteTableModelWithDefaults

`func NewCreateOrUpdateRouteTableModelWithDefaults() *CreateOrUpdateRouteTableModel`

NewCreateOrUpdateRouteTableModelWithDefaults instantiates a new CreateOrUpdateRouteTableModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureRoutes

`func (o *CreateOrUpdateRouteTableModel) GetAzureRoutes() []AzureRoute`

GetAzureRoutes returns the AzureRoutes field if non-nil, zero value otherwise.

### GetAzureRoutesOk

`func (o *CreateOrUpdateRouteTableModel) GetAzureRoutesOk() (*[]AzureRoute, bool)`

GetAzureRoutesOk returns a tuple with the AzureRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRoutes

`func (o *CreateOrUpdateRouteTableModel) SetAzureRoutes(v []AzureRoute)`

SetAzureRoutes sets AzureRoutes field to given value.


### GetDisableRoutePropagation

`func (o *CreateOrUpdateRouteTableModel) GetDisableRoutePropagation() bool`

GetDisableRoutePropagation returns the DisableRoutePropagation field if non-nil, zero value otherwise.

### GetDisableRoutePropagationOk

`func (o *CreateOrUpdateRouteTableModel) GetDisableRoutePropagationOk() (*bool, bool)`

GetDisableRoutePropagationOk returns a tuple with the DisableRoutePropagation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRoutePropagation

`func (o *CreateOrUpdateRouteTableModel) SetDisableRoutePropagation(v bool)`

SetDisableRoutePropagation sets DisableRoutePropagation field to given value.

### HasDisableRoutePropagation

`func (o *CreateOrUpdateRouteTableModel) HasDisableRoutePropagation() bool`

HasDisableRoutePropagation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


