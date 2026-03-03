# EnableServiceEndpointsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceEndpoints** | **[]string** | List of service endpoint types to enable (e.g., Microsoft.Storage, Microsoft.Sql) | 

## Methods

### NewEnableServiceEndpointsRequest

`func NewEnableServiceEndpointsRequest(serviceEndpoints []string, ) *EnableServiceEndpointsRequest`

NewEnableServiceEndpointsRequest instantiates a new EnableServiceEndpointsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableServiceEndpointsRequestWithDefaults

`func NewEnableServiceEndpointsRequestWithDefaults() *EnableServiceEndpointsRequest`

NewEnableServiceEndpointsRequestWithDefaults instantiates a new EnableServiceEndpointsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceEndpoints

`func (o *EnableServiceEndpointsRequest) GetServiceEndpoints() []string`

GetServiceEndpoints returns the ServiceEndpoints field if non-nil, zero value otherwise.

### GetServiceEndpointsOk

`func (o *EnableServiceEndpointsRequest) GetServiceEndpointsOk() (*[]string, bool)`

GetServiceEndpointsOk returns a tuple with the ServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoints

`func (o *EnableServiceEndpointsRequest) SetServiceEndpoints(v []string)`

SetServiceEndpoints sets ServiceEndpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


