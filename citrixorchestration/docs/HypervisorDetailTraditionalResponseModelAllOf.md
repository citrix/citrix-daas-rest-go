# HypervisorDetailTraditionalResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SslThumbprints** | Pointer to **[]string** | SSL thumbprints considered acceptable for the SSL certificate presented by the hypervisor. | [optional] 
**UserName** | **string** | The user name for the credentials used to communicate with the hypervisor. | 

## Methods

### NewHypervisorDetailTraditionalResponseModelAllOf

`func NewHypervisorDetailTraditionalResponseModelAllOf(userName string, ) *HypervisorDetailTraditionalResponseModelAllOf`

NewHypervisorDetailTraditionalResponseModelAllOf instantiates a new HypervisorDetailTraditionalResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorDetailTraditionalResponseModelAllOfWithDefaults

`func NewHypervisorDetailTraditionalResponseModelAllOfWithDefaults() *HypervisorDetailTraditionalResponseModelAllOf`

NewHypervisorDetailTraditionalResponseModelAllOfWithDefaults instantiates a new HypervisorDetailTraditionalResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSslThumbprints

`func (o *HypervisorDetailTraditionalResponseModelAllOf) GetSslThumbprints() []string`

GetSslThumbprints returns the SslThumbprints field if non-nil, zero value otherwise.

### GetSslThumbprintsOk

`func (o *HypervisorDetailTraditionalResponseModelAllOf) GetSslThumbprintsOk() (*[]string, bool)`

GetSslThumbprintsOk returns a tuple with the SslThumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprints

`func (o *HypervisorDetailTraditionalResponseModelAllOf) SetSslThumbprints(v []string)`

SetSslThumbprints sets SslThumbprints field to given value.

### HasSslThumbprints

`func (o *HypervisorDetailTraditionalResponseModelAllOf) HasSslThumbprints() bool`

HasSslThumbprints returns a boolean if a field has been set.

### GetUserName

`func (o *HypervisorDetailTraditionalResponseModelAllOf) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HypervisorDetailTraditionalResponseModelAllOf) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HypervisorDetailTraditionalResponseModelAllOf) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


