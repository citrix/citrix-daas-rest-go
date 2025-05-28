# CatalogLimitModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingCatalogsCount** | Pointer to **int32** | Number of existing catalogs | [optional] 
**ServiceState** | Pointer to [**ServiceState**](ServiceState.md) | Current service state of customer | [optional] 
**CitrixManagedLimit** | Pointer to [**CatalogSublimitModel**](CatalogSublimitModel.md) | Limits on the Citrix Managed catalogs | [optional] 
**ByoaLimit** | Pointer to [**CatalogSublimitModel**](CatalogSublimitModel.md) | Limits on the BYOA catalogs | [optional] 

## Methods

### NewCatalogLimitModel

`func NewCatalogLimitModel() *CatalogLimitModel`

NewCatalogLimitModel instantiates a new CatalogLimitModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogLimitModelWithDefaults

`func NewCatalogLimitModelWithDefaults() *CatalogLimitModel`

NewCatalogLimitModelWithDefaults instantiates a new CatalogLimitModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingCatalogsCount

`func (o *CatalogLimitModel) GetExistingCatalogsCount() int32`

GetExistingCatalogsCount returns the ExistingCatalogsCount field if non-nil, zero value otherwise.

### GetExistingCatalogsCountOk

`func (o *CatalogLimitModel) GetExistingCatalogsCountOk() (*int32, bool)`

GetExistingCatalogsCountOk returns a tuple with the ExistingCatalogsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingCatalogsCount

`func (o *CatalogLimitModel) SetExistingCatalogsCount(v int32)`

SetExistingCatalogsCount sets ExistingCatalogsCount field to given value.

### HasExistingCatalogsCount

`func (o *CatalogLimitModel) HasExistingCatalogsCount() bool`

HasExistingCatalogsCount returns a boolean if a field has been set.

### GetServiceState

`func (o *CatalogLimitModel) GetServiceState() ServiceState`

GetServiceState returns the ServiceState field if non-nil, zero value otherwise.

### GetServiceStateOk

`func (o *CatalogLimitModel) GetServiceStateOk() (*ServiceState, bool)`

GetServiceStateOk returns a tuple with the ServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceState

`func (o *CatalogLimitModel) SetServiceState(v ServiceState)`

SetServiceState sets ServiceState field to given value.

### HasServiceState

`func (o *CatalogLimitModel) HasServiceState() bool`

HasServiceState returns a boolean if a field has been set.

### GetCitrixManagedLimit

`func (o *CatalogLimitModel) GetCitrixManagedLimit() CatalogSublimitModel`

GetCitrixManagedLimit returns the CitrixManagedLimit field if non-nil, zero value otherwise.

### GetCitrixManagedLimitOk

`func (o *CatalogLimitModel) GetCitrixManagedLimitOk() (*CatalogSublimitModel, bool)`

GetCitrixManagedLimitOk returns a tuple with the CitrixManagedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManagedLimit

`func (o *CatalogLimitModel) SetCitrixManagedLimit(v CatalogSublimitModel)`

SetCitrixManagedLimit sets CitrixManagedLimit field to given value.

### HasCitrixManagedLimit

`func (o *CatalogLimitModel) HasCitrixManagedLimit() bool`

HasCitrixManagedLimit returns a boolean if a field has been set.

### GetByoaLimit

`func (o *CatalogLimitModel) GetByoaLimit() CatalogSublimitModel`

GetByoaLimit returns the ByoaLimit field if non-nil, zero value otherwise.

### GetByoaLimitOk

`func (o *CatalogLimitModel) GetByoaLimitOk() (*CatalogSublimitModel, bool)`

GetByoaLimitOk returns a tuple with the ByoaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaLimit

`func (o *CatalogLimitModel) SetByoaLimit(v CatalogSublimitModel)`

SetByoaLimit sets ByoaLimit field to given value.

### HasByoaLimit

`func (o *CatalogLimitModel) HasByoaLimit() bool`

HasByoaLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


