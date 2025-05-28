# Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Essentials** | Pointer to [**Essentials**](Essentials.md) |  | [optional] 
**AlertContext** | Pointer to [**AlertContext**](AlertContext.md) |  | [optional] 

## Methods

### NewData

`func NewData() *Data`

NewData instantiates a new Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWithDefaults

`func NewDataWithDefaults() *Data`

NewDataWithDefaults instantiates a new Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEssentials

`func (o *Data) GetEssentials() Essentials`

GetEssentials returns the Essentials field if non-nil, zero value otherwise.

### GetEssentialsOk

`func (o *Data) GetEssentialsOk() (*Essentials, bool)`

GetEssentialsOk returns a tuple with the Essentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssentials

`func (o *Data) SetEssentials(v Essentials)`

SetEssentials sets Essentials field to given value.

### HasEssentials

`func (o *Data) HasEssentials() bool`

HasEssentials returns a boolean if a field has been set.

### GetAlertContext

`func (o *Data) GetAlertContext() AlertContext`

GetAlertContext returns the AlertContext field if non-nil, zero value otherwise.

### GetAlertContextOk

`func (o *Data) GetAlertContextOk() (*AlertContext, bool)`

GetAlertContextOk returns a tuple with the AlertContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertContext

`func (o *Data) SetAlertContext(v AlertContext)`

SetAlertContext sets AlertContext field to given value.

### HasAlertContext

`func (o *Data) HasAlertContext() bool`

HasAlertContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


