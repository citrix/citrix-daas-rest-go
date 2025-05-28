# RouteTableOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the route table | 
**State** | Pointer to [**RouteTableState**](RouteTableState.md) | The current state of the route table | [optional] 
**ErrorMessage** | Pointer to **string** | The error message if the route table failed | [optional] 
**ActiveRoutes** | Pointer to [**[]RouteOverview**](RouteOverview.md) | The routes in the route table that are currently active | [optional] 
**PendingRoutes** | Pointer to [**[]RouteOverview**](RouteOverview.md) | The routes that have to be applied to the route table | [optional] 
**TransactionId** | Pointer to **string** | The transaction id | [optional] 

## Methods

### NewRouteTableOverview

`func NewRouteTableOverview(id string, ) *RouteTableOverview`

NewRouteTableOverview instantiates a new RouteTableOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTableOverviewWithDefaults

`func NewRouteTableOverviewWithDefaults() *RouteTableOverview`

NewRouteTableOverviewWithDefaults instantiates a new RouteTableOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouteTableOverview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteTableOverview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteTableOverview) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *RouteTableOverview) GetState() RouteTableState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RouteTableOverview) GetStateOk() (*RouteTableState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RouteTableOverview) SetState(v RouteTableState)`

SetState sets State field to given value.

### HasState

`func (o *RouteTableOverview) HasState() bool`

HasState returns a boolean if a field has been set.

### GetErrorMessage

`func (o *RouteTableOverview) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *RouteTableOverview) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *RouteTableOverview) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *RouteTableOverview) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetActiveRoutes

`func (o *RouteTableOverview) GetActiveRoutes() []RouteOverview`

GetActiveRoutes returns the ActiveRoutes field if non-nil, zero value otherwise.

### GetActiveRoutesOk

`func (o *RouteTableOverview) GetActiveRoutesOk() (*[]RouteOverview, bool)`

GetActiveRoutesOk returns a tuple with the ActiveRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRoutes

`func (o *RouteTableOverview) SetActiveRoutes(v []RouteOverview)`

SetActiveRoutes sets ActiveRoutes field to given value.

### HasActiveRoutes

`func (o *RouteTableOverview) HasActiveRoutes() bool`

HasActiveRoutes returns a boolean if a field has been set.

### GetPendingRoutes

`func (o *RouteTableOverview) GetPendingRoutes() []RouteOverview`

GetPendingRoutes returns the PendingRoutes field if non-nil, zero value otherwise.

### GetPendingRoutesOk

`func (o *RouteTableOverview) GetPendingRoutesOk() (*[]RouteOverview, bool)`

GetPendingRoutesOk returns a tuple with the PendingRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRoutes

`func (o *RouteTableOverview) SetPendingRoutes(v []RouteOverview)`

SetPendingRoutes sets PendingRoutes field to given value.

### HasPendingRoutes

`func (o *RouteTableOverview) HasPendingRoutes() bool`

HasPendingRoutes returns a boolean if a field has been set.

### GetTransactionId

`func (o *RouteTableOverview) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *RouteTableOverview) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *RouteTableOverview) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *RouteTableOverview) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


