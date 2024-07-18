# CitrixCloudServicesApiModelsSegmentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]CitrixCloudServicesServicePrincipalsModelsServicePrincipal**](CitrixCloudServicesServicePrincipalsModelsServicePrincipal.md) |  | [optional] 
**ContinuationToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCitrixCloudServicesApiModelsSegmentResult

`func NewCitrixCloudServicesApiModelsSegmentResult() *CitrixCloudServicesApiModelsSegmentResult`

NewCitrixCloudServicesApiModelsSegmentResult instantiates a new CitrixCloudServicesApiModelsSegmentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesApiModelsSegmentResultWithDefaults

`func NewCitrixCloudServicesApiModelsSegmentResultWithDefaults() *CitrixCloudServicesApiModelsSegmentResult`

NewCitrixCloudServicesApiModelsSegmentResultWithDefaults instantiates a new CitrixCloudServicesApiModelsSegmentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CitrixCloudServicesApiModelsSegmentResult) GetItems() []CitrixCloudServicesServicePrincipalsModelsServicePrincipal`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CitrixCloudServicesApiModelsSegmentResult) GetItemsOk() (*[]CitrixCloudServicesServicePrincipalsModelsServicePrincipal, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CitrixCloudServicesApiModelsSegmentResult) SetItems(v []CitrixCloudServicesServicePrincipalsModelsServicePrincipal)`

SetItems sets Items field to given value.

### HasItems

`func (o *CitrixCloudServicesApiModelsSegmentResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *CitrixCloudServicesApiModelsSegmentResult) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *CitrixCloudServicesApiModelsSegmentResult) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetContinuationToken

`func (o *CitrixCloudServicesApiModelsSegmentResult) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *CitrixCloudServicesApiModelsSegmentResult) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *CitrixCloudServicesApiModelsSegmentResult) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *CitrixCloudServicesApiModelsSegmentResult) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *CitrixCloudServicesApiModelsSegmentResult) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *CitrixCloudServicesApiModelsSegmentResult) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


