# SiteServiceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** | Name of the service. | 
**ServiceType** | **string** | Type of the service, which will not be changed across languages. | 
**CurrentSchemaVersion** | Pointer to **string** | Current schema version of the service. Will be &#x60;null&#x60; for XenApp &amp; XenDesktop service. | [optional] 
**DesiredSchemaVersion** | Pointer to **string** | Desired schema version of the service. Will be &#x60;null&#x60; for XenApp &amp; XenDesktop service. | [optional] 
**Capabilities** | **[]string** | List of capabilities exposed by the service. | 

## Methods

### NewSiteServiceResponseModel

`func NewSiteServiceResponseModel(serviceName string, serviceType string, capabilities []string, ) *SiteServiceResponseModel`

NewSiteServiceResponseModel instantiates a new SiteServiceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteServiceResponseModelWithDefaults

`func NewSiteServiceResponseModelWithDefaults() *SiteServiceResponseModel`

NewSiteServiceResponseModelWithDefaults instantiates a new SiteServiceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *SiteServiceResponseModel) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SiteServiceResponseModel) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SiteServiceResponseModel) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServiceType

`func (o *SiteServiceResponseModel) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *SiteServiceResponseModel) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *SiteServiceResponseModel) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetCurrentSchemaVersion

`func (o *SiteServiceResponseModel) GetCurrentSchemaVersion() string`

GetCurrentSchemaVersion returns the CurrentSchemaVersion field if non-nil, zero value otherwise.

### GetCurrentSchemaVersionOk

`func (o *SiteServiceResponseModel) GetCurrentSchemaVersionOk() (*string, bool)`

GetCurrentSchemaVersionOk returns a tuple with the CurrentSchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSchemaVersion

`func (o *SiteServiceResponseModel) SetCurrentSchemaVersion(v string)`

SetCurrentSchemaVersion sets CurrentSchemaVersion field to given value.

### HasCurrentSchemaVersion

`func (o *SiteServiceResponseModel) HasCurrentSchemaVersion() bool`

HasCurrentSchemaVersion returns a boolean if a field has been set.

### GetDesiredSchemaVersion

`func (o *SiteServiceResponseModel) GetDesiredSchemaVersion() string`

GetDesiredSchemaVersion returns the DesiredSchemaVersion field if non-nil, zero value otherwise.

### GetDesiredSchemaVersionOk

`func (o *SiteServiceResponseModel) GetDesiredSchemaVersionOk() (*string, bool)`

GetDesiredSchemaVersionOk returns a tuple with the DesiredSchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredSchemaVersion

`func (o *SiteServiceResponseModel) SetDesiredSchemaVersion(v string)`

SetDesiredSchemaVersion sets DesiredSchemaVersion field to given value.

### HasDesiredSchemaVersion

`func (o *SiteServiceResponseModel) HasDesiredSchemaVersion() bool`

HasDesiredSchemaVersion returns a boolean if a field has been set.

### GetCapabilities

`func (o *SiteServiceResponseModel) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *SiteServiceResponseModel) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *SiteServiceResponseModel) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


