# MultiSiteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Site Name | [optional] 
**DisplayName** | **string** | The site&#39;s friendly name | 
**DeliveryControllers** | Pointer to **[]string** | The Orchestration servers belongs to the site | [optional] 
**Default** | Pointer to **bool** | Indicate if it is the default site. | [optional] 

## Methods

### NewMultiSiteModel

`func NewMultiSiteModel(displayName string, ) *MultiSiteModel`

NewMultiSiteModel instantiates a new MultiSiteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiSiteModelWithDefaults

`func NewMultiSiteModelWithDefaults() *MultiSiteModel`

NewMultiSiteModelWithDefaults instantiates a new MultiSiteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MultiSiteModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MultiSiteModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MultiSiteModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MultiSiteModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *MultiSiteModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MultiSiteModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MultiSiteModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDeliveryControllers

`func (o *MultiSiteModel) GetDeliveryControllers() []string`

GetDeliveryControllers returns the DeliveryControllers field if non-nil, zero value otherwise.

### GetDeliveryControllersOk

`func (o *MultiSiteModel) GetDeliveryControllersOk() (*[]string, bool)`

GetDeliveryControllersOk returns a tuple with the DeliveryControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryControllers

`func (o *MultiSiteModel) SetDeliveryControllers(v []string)`

SetDeliveryControllers sets DeliveryControllers field to given value.

### HasDeliveryControllers

`func (o *MultiSiteModel) HasDeliveryControllers() bool`

HasDeliveryControllers returns a boolean if a field has been set.

### GetDefault

`func (o *MultiSiteModel) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *MultiSiteModel) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *MultiSiteModel) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *MultiSiteModel) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


