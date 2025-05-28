# VNetResourceLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | The id of the subscription | [optional] 
**ResourceLocations** | Pointer to [**[]VNetResourceLocation**](VNetResourceLocation.md) | The list of resource locations associated with the subscription | [optional] 

## Methods

### NewVNetResourceLocations

`func NewVNetResourceLocations() *VNetResourceLocations`

NewVNetResourceLocations instantiates a new VNetResourceLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVNetResourceLocationsWithDefaults

`func NewVNetResourceLocationsWithDefaults() *VNetResourceLocations`

NewVNetResourceLocationsWithDefaults instantiates a new VNetResourceLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *VNetResourceLocations) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *VNetResourceLocations) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *VNetResourceLocations) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *VNetResourceLocations) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetResourceLocations

`func (o *VNetResourceLocations) GetResourceLocations() []VNetResourceLocation`

GetResourceLocations returns the ResourceLocations field if non-nil, zero value otherwise.

### GetResourceLocationsOk

`func (o *VNetResourceLocations) GetResourceLocationsOk() (*[]VNetResourceLocation, bool)`

GetResourceLocationsOk returns a tuple with the ResourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocations

`func (o *VNetResourceLocations) SetResourceLocations(v []VNetResourceLocation)`

SetResourceLocations sets ResourceLocations field to given value.

### HasResourceLocations

`func (o *VNetResourceLocations) HasResourceLocations() bool`

HasResourceLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


