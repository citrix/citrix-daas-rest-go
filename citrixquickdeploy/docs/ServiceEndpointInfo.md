# ServiceEndpointInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **NullableString** | The service endpoint type (e.g., Microsoft.Storage, Microsoft.Sql) | [optional] 
**IsEnabled** | Pointer to **bool** | Whether this service endpoint is currently enabled on the subnet | [optional] 
**Locations** | Pointer to **[]string** | The locations/regions where this service endpoint is available | [optional] 

## Methods

### NewServiceEndpointInfo

`func NewServiceEndpointInfo() *ServiceEndpointInfo`

NewServiceEndpointInfo instantiates a new ServiceEndpointInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceEndpointInfoWithDefaults

`func NewServiceEndpointInfoWithDefaults() *ServiceEndpointInfo`

NewServiceEndpointInfoWithDefaults instantiates a new ServiceEndpointInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServiceEndpointInfo) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceEndpointInfo) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceEndpointInfo) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceEndpointInfo) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *ServiceEndpointInfo) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *ServiceEndpointInfo) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetIsEnabled

`func (o *ServiceEndpointInfo) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ServiceEndpointInfo) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ServiceEndpointInfo) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ServiceEndpointInfo) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetLocations

`func (o *ServiceEndpointInfo) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ServiceEndpointInfo) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ServiceEndpointInfo) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ServiceEndpointInfo) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *ServiceEndpointInfo) SetLocationsNil(b bool)`

 SetLocationsNil sets the value for Locations to be an explicit nil

### UnsetLocations
`func (o *ServiceEndpointInfo) UnsetLocations()`

UnsetLocations ensures that no value is present for Locations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


