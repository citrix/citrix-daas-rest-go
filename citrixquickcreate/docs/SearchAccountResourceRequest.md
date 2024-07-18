# SearchAccountResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**FilterProperty** | Pointer to **NullableString** | The filter property key. | [optional] 
**FilterValue** | Pointer to **NullableString** | The filter property value. | [optional] 

## Methods

### NewSearchAccountResourceRequest

`func NewSearchAccountResourceRequest(accountType AccountType, ) *SearchAccountResourceRequest`

NewSearchAccountResourceRequest instantiates a new SearchAccountResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAccountResourceRequestWithDefaults

`func NewSearchAccountResourceRequestWithDefaults() *SearchAccountResourceRequest`

NewSearchAccountResourceRequestWithDefaults instantiates a new SearchAccountResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *SearchAccountResourceRequest) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *SearchAccountResourceRequest) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *SearchAccountResourceRequest) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetFilterProperty

`func (o *SearchAccountResourceRequest) GetFilterProperty() string`

GetFilterProperty returns the FilterProperty field if non-nil, zero value otherwise.

### GetFilterPropertyOk

`func (o *SearchAccountResourceRequest) GetFilterPropertyOk() (*string, bool)`

GetFilterPropertyOk returns a tuple with the FilterProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterProperty

`func (o *SearchAccountResourceRequest) SetFilterProperty(v string)`

SetFilterProperty sets FilterProperty field to given value.

### HasFilterProperty

`func (o *SearchAccountResourceRequest) HasFilterProperty() bool`

HasFilterProperty returns a boolean if a field has been set.

### SetFilterPropertyNil

`func (o *SearchAccountResourceRequest) SetFilterPropertyNil(b bool)`

 SetFilterPropertyNil sets the value for FilterProperty to be an explicit nil

### UnsetFilterProperty
`func (o *SearchAccountResourceRequest) UnsetFilterProperty()`

UnsetFilterProperty ensures that no value is present for FilterProperty, not even an explicit nil
### GetFilterValue

`func (o *SearchAccountResourceRequest) GetFilterValue() string`

GetFilterValue returns the FilterValue field if non-nil, zero value otherwise.

### GetFilterValueOk

`func (o *SearchAccountResourceRequest) GetFilterValueOk() (*string, bool)`

GetFilterValueOk returns a tuple with the FilterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValue

`func (o *SearchAccountResourceRequest) SetFilterValue(v string)`

SetFilterValue sets FilterValue field to given value.

### HasFilterValue

`func (o *SearchAccountResourceRequest) HasFilterValue() bool`

HasFilterValue returns a boolean if a field has been set.

### SetFilterValueNil

`func (o *SearchAccountResourceRequest) SetFilterValueNil(b bool)`

 SetFilterValueNil sets the value for FilterValue to be an explicit nil

### UnsetFilterValue
`func (o *SearchAccountResourceRequest) UnsetFilterValue()`

UnsetFilterValue ensures that no value is present for FilterValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


