# EnableServiceEndpointsForOnPremConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnPremConnectionType** | [**OnPremConnectionType**](OnPremConnectionType.md) | Type of on-premises connection for which to enable service endpoints | 
**ServiceEndpoints** | **[]string** | List of service endpoint types to enable (e.g., Microsoft.Storage, Microsoft.Sql) | 

## Methods

### NewEnableServiceEndpointsForOnPremConnectionRequest

`func NewEnableServiceEndpointsForOnPremConnectionRequest(onPremConnectionType OnPremConnectionType, serviceEndpoints []string, ) *EnableServiceEndpointsForOnPremConnectionRequest`

NewEnableServiceEndpointsForOnPremConnectionRequest instantiates a new EnableServiceEndpointsForOnPremConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableServiceEndpointsForOnPremConnectionRequestWithDefaults

`func NewEnableServiceEndpointsForOnPremConnectionRequestWithDefaults() *EnableServiceEndpointsForOnPremConnectionRequest`

NewEnableServiceEndpointsForOnPremConnectionRequestWithDefaults instantiates a new EnableServiceEndpointsForOnPremConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnPremConnectionType

`func (o *EnableServiceEndpointsForOnPremConnectionRequest) GetOnPremConnectionType() OnPremConnectionType`

GetOnPremConnectionType returns the OnPremConnectionType field if non-nil, zero value otherwise.

### GetOnPremConnectionTypeOk

`func (o *EnableServiceEndpointsForOnPremConnectionRequest) GetOnPremConnectionTypeOk() (*OnPremConnectionType, bool)`

GetOnPremConnectionTypeOk returns a tuple with the OnPremConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremConnectionType

`func (o *EnableServiceEndpointsForOnPremConnectionRequest) SetOnPremConnectionType(v OnPremConnectionType)`

SetOnPremConnectionType sets OnPremConnectionType field to given value.


### GetServiceEndpoints

`func (o *EnableServiceEndpointsForOnPremConnectionRequest) GetServiceEndpoints() []string`

GetServiceEndpoints returns the ServiceEndpoints field if non-nil, zero value otherwise.

### GetServiceEndpointsOk

`func (o *EnableServiceEndpointsForOnPremConnectionRequest) GetServiceEndpointsOk() (*[]string, bool)`

GetServiceEndpointsOk returns a tuple with the ServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoints

`func (o *EnableServiceEndpointsForOnPremConnectionRequest) SetServiceEndpoints(v []string)`

SetServiceEndpoints sets ServiceEndpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


