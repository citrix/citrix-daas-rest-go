# DataStoreCustomerPowerResourcesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**PowerMode**](PowerMode.md) |  | [optional] 
**Status** | Pointer to [**PowerResourcesJobStatus**](PowerResourcesJobStatus.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDataStoreCustomerPowerResourcesModel

`func NewDataStoreCustomerPowerResourcesModel() *DataStoreCustomerPowerResourcesModel`

NewDataStoreCustomerPowerResourcesModel instantiates a new DataStoreCustomerPowerResourcesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreCustomerPowerResourcesModelWithDefaults

`func NewDataStoreCustomerPowerResourcesModelWithDefaults() *DataStoreCustomerPowerResourcesModel`

NewDataStoreCustomerPowerResourcesModelWithDefaults instantiates a new DataStoreCustomerPowerResourcesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataStoreCustomerPowerResourcesModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataStoreCustomerPowerResourcesModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataStoreCustomerPowerResourcesModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataStoreCustomerPowerResourcesModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *DataStoreCustomerPowerResourcesModel) GetMode() PowerMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DataStoreCustomerPowerResourcesModel) GetModeOk() (*PowerMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DataStoreCustomerPowerResourcesModel) SetMode(v PowerMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DataStoreCustomerPowerResourcesModel) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetStatus

`func (o *DataStoreCustomerPowerResourcesModel) GetStatus() PowerResourcesJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataStoreCustomerPowerResourcesModel) GetStatusOk() (*PowerResourcesJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataStoreCustomerPowerResourcesModel) SetStatus(v PowerResourcesJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataStoreCustomerPowerResourcesModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *DataStoreCustomerPowerResourcesModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DataStoreCustomerPowerResourcesModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DataStoreCustomerPowerResourcesModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DataStoreCustomerPowerResourcesModel) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DataStoreCustomerPowerResourcesModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataStoreCustomerPowerResourcesModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataStoreCustomerPowerResourcesModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataStoreCustomerPowerResourcesModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *DataStoreCustomerPowerResourcesModel) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DataStoreCustomerPowerResourcesModel) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DataStoreCustomerPowerResourcesModel) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DataStoreCustomerPowerResourcesModel) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *DataStoreCustomerPowerResourcesModel) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *DataStoreCustomerPowerResourcesModel) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *DataStoreCustomerPowerResourcesModel) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *DataStoreCustomerPowerResourcesModel) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


