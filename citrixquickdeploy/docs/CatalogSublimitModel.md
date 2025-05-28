# CatalogSublimitModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Limit on number of catalog creation. | [optional] 
**ExistingCatalogsCount** | Pointer to **int32** | Number of existing catalogs. | [optional] 
**HasLimitReached** | Pointer to **bool** | Indicates whether a customer&#39;s number of catalog has reached the limit. | [optional] [readonly] 
**MaxVdaCount** | Pointer to **int32** | The maximum number of VDAs per catalog. | [optional] 
**MaxVdaPerSubscription** | Pointer to **int32** | The max number of VDAs supported per Azure Subscription. | [optional] 

## Methods

### NewCatalogSublimitModel

`func NewCatalogSublimitModel() *CatalogSublimitModel`

NewCatalogSublimitModel instantiates a new CatalogSublimitModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogSublimitModelWithDefaults

`func NewCatalogSublimitModelWithDefaults() *CatalogSublimitModel`

NewCatalogSublimitModelWithDefaults instantiates a new CatalogSublimitModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *CatalogSublimitModel) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CatalogSublimitModel) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CatalogSublimitModel) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CatalogSublimitModel) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetExistingCatalogsCount

`func (o *CatalogSublimitModel) GetExistingCatalogsCount() int32`

GetExistingCatalogsCount returns the ExistingCatalogsCount field if non-nil, zero value otherwise.

### GetExistingCatalogsCountOk

`func (o *CatalogSublimitModel) GetExistingCatalogsCountOk() (*int32, bool)`

GetExistingCatalogsCountOk returns a tuple with the ExistingCatalogsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingCatalogsCount

`func (o *CatalogSublimitModel) SetExistingCatalogsCount(v int32)`

SetExistingCatalogsCount sets ExistingCatalogsCount field to given value.

### HasExistingCatalogsCount

`func (o *CatalogSublimitModel) HasExistingCatalogsCount() bool`

HasExistingCatalogsCount returns a boolean if a field has been set.

### GetHasLimitReached

`func (o *CatalogSublimitModel) GetHasLimitReached() bool`

GetHasLimitReached returns the HasLimitReached field if non-nil, zero value otherwise.

### GetHasLimitReachedOk

`func (o *CatalogSublimitModel) GetHasLimitReachedOk() (*bool, bool)`

GetHasLimitReachedOk returns a tuple with the HasLimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLimitReached

`func (o *CatalogSublimitModel) SetHasLimitReached(v bool)`

SetHasLimitReached sets HasLimitReached field to given value.

### HasHasLimitReached

`func (o *CatalogSublimitModel) HasHasLimitReached() bool`

HasHasLimitReached returns a boolean if a field has been set.

### GetMaxVdaCount

`func (o *CatalogSublimitModel) GetMaxVdaCount() int32`

GetMaxVdaCount returns the MaxVdaCount field if non-nil, zero value otherwise.

### GetMaxVdaCountOk

`func (o *CatalogSublimitModel) GetMaxVdaCountOk() (*int32, bool)`

GetMaxVdaCountOk returns a tuple with the MaxVdaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVdaCount

`func (o *CatalogSublimitModel) SetMaxVdaCount(v int32)`

SetMaxVdaCount sets MaxVdaCount field to given value.

### HasMaxVdaCount

`func (o *CatalogSublimitModel) HasMaxVdaCount() bool`

HasMaxVdaCount returns a boolean if a field has been set.

### GetMaxVdaPerSubscription

`func (o *CatalogSublimitModel) GetMaxVdaPerSubscription() int32`

GetMaxVdaPerSubscription returns the MaxVdaPerSubscription field if non-nil, zero value otherwise.

### GetMaxVdaPerSubscriptionOk

`func (o *CatalogSublimitModel) GetMaxVdaPerSubscriptionOk() (*int32, bool)`

GetMaxVdaPerSubscriptionOk returns a tuple with the MaxVdaPerSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVdaPerSubscription

`func (o *CatalogSublimitModel) SetMaxVdaPerSubscription(v int32)`

SetMaxVdaPerSubscription sets MaxVdaPerSubscription field to given value.

### HasMaxVdaPerSubscription

`func (o *CatalogSublimitModel) HasMaxVdaPerSubscription() bool`

HasMaxVdaPerSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


