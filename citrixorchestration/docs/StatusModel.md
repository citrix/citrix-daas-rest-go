# StatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrchServiceVersion** | Pointer to **NullableString** | The Orchestration service version. | [optional] 
**ProductInternalVersion** | Pointer to **NullableString** | The Product internal version. | [optional] 
**ProductExternalVersion** | Pointer to **NullableString** | The product external version. | [optional] 
**OrchServiceStatus** | Pointer to [**ServiceStatus**](ServiceStatus.md) |  | [optional] 
**OrchRESTFulApiStatus** | Pointer to [**RestApiStatus**](RestApiStatus.md) |  | [optional] 
**SupportedAuthenticators** | Pointer to [**Authenticator**](Authenticator.md) |  | [optional] 

## Methods

### NewStatusModel

`func NewStatusModel() *StatusModel`

NewStatusModel instantiates a new StatusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusModelWithDefaults

`func NewStatusModelWithDefaults() *StatusModel`

NewStatusModelWithDefaults instantiates a new StatusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchServiceVersion

`func (o *StatusModel) GetOrchServiceVersion() string`

GetOrchServiceVersion returns the OrchServiceVersion field if non-nil, zero value otherwise.

### GetOrchServiceVersionOk

`func (o *StatusModel) GetOrchServiceVersionOk() (*string, bool)`

GetOrchServiceVersionOk returns a tuple with the OrchServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchServiceVersion

`func (o *StatusModel) SetOrchServiceVersion(v string)`

SetOrchServiceVersion sets OrchServiceVersion field to given value.

### HasOrchServiceVersion

`func (o *StatusModel) HasOrchServiceVersion() bool`

HasOrchServiceVersion returns a boolean if a field has been set.

### SetOrchServiceVersionNil

`func (o *StatusModel) SetOrchServiceVersionNil(b bool)`

 SetOrchServiceVersionNil sets the value for OrchServiceVersion to be an explicit nil

### UnsetOrchServiceVersion
`func (o *StatusModel) UnsetOrchServiceVersion()`

UnsetOrchServiceVersion ensures that no value is present for OrchServiceVersion, not even an explicit nil
### GetProductInternalVersion

`func (o *StatusModel) GetProductInternalVersion() string`

GetProductInternalVersion returns the ProductInternalVersion field if non-nil, zero value otherwise.

### GetProductInternalVersionOk

`func (o *StatusModel) GetProductInternalVersionOk() (*string, bool)`

GetProductInternalVersionOk returns a tuple with the ProductInternalVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInternalVersion

`func (o *StatusModel) SetProductInternalVersion(v string)`

SetProductInternalVersion sets ProductInternalVersion field to given value.

### HasProductInternalVersion

`func (o *StatusModel) HasProductInternalVersion() bool`

HasProductInternalVersion returns a boolean if a field has been set.

### SetProductInternalVersionNil

`func (o *StatusModel) SetProductInternalVersionNil(b bool)`

 SetProductInternalVersionNil sets the value for ProductInternalVersion to be an explicit nil

### UnsetProductInternalVersion
`func (o *StatusModel) UnsetProductInternalVersion()`

UnsetProductInternalVersion ensures that no value is present for ProductInternalVersion, not even an explicit nil
### GetProductExternalVersion

`func (o *StatusModel) GetProductExternalVersion() string`

GetProductExternalVersion returns the ProductExternalVersion field if non-nil, zero value otherwise.

### GetProductExternalVersionOk

`func (o *StatusModel) GetProductExternalVersionOk() (*string, bool)`

GetProductExternalVersionOk returns a tuple with the ProductExternalVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductExternalVersion

`func (o *StatusModel) SetProductExternalVersion(v string)`

SetProductExternalVersion sets ProductExternalVersion field to given value.

### HasProductExternalVersion

`func (o *StatusModel) HasProductExternalVersion() bool`

HasProductExternalVersion returns a boolean if a field has been set.

### SetProductExternalVersionNil

`func (o *StatusModel) SetProductExternalVersionNil(b bool)`

 SetProductExternalVersionNil sets the value for ProductExternalVersion to be an explicit nil

### UnsetProductExternalVersion
`func (o *StatusModel) UnsetProductExternalVersion()`

UnsetProductExternalVersion ensures that no value is present for ProductExternalVersion, not even an explicit nil
### GetOrchServiceStatus

`func (o *StatusModel) GetOrchServiceStatus() ServiceStatus`

GetOrchServiceStatus returns the OrchServiceStatus field if non-nil, zero value otherwise.

### GetOrchServiceStatusOk

`func (o *StatusModel) GetOrchServiceStatusOk() (*ServiceStatus, bool)`

GetOrchServiceStatusOk returns a tuple with the OrchServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchServiceStatus

`func (o *StatusModel) SetOrchServiceStatus(v ServiceStatus)`

SetOrchServiceStatus sets OrchServiceStatus field to given value.

### HasOrchServiceStatus

`func (o *StatusModel) HasOrchServiceStatus() bool`

HasOrchServiceStatus returns a boolean if a field has been set.

### GetOrchRESTFulApiStatus

`func (o *StatusModel) GetOrchRESTFulApiStatus() RestApiStatus`

GetOrchRESTFulApiStatus returns the OrchRESTFulApiStatus field if non-nil, zero value otherwise.

### GetOrchRESTFulApiStatusOk

`func (o *StatusModel) GetOrchRESTFulApiStatusOk() (*RestApiStatus, bool)`

GetOrchRESTFulApiStatusOk returns a tuple with the OrchRESTFulApiStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchRESTFulApiStatus

`func (o *StatusModel) SetOrchRESTFulApiStatus(v RestApiStatus)`

SetOrchRESTFulApiStatus sets OrchRESTFulApiStatus field to given value.

### HasOrchRESTFulApiStatus

`func (o *StatusModel) HasOrchRESTFulApiStatus() bool`

HasOrchRESTFulApiStatus returns a boolean if a field has been set.

### GetSupportedAuthenticators

`func (o *StatusModel) GetSupportedAuthenticators() Authenticator`

GetSupportedAuthenticators returns the SupportedAuthenticators field if non-nil, zero value otherwise.

### GetSupportedAuthenticatorsOk

`func (o *StatusModel) GetSupportedAuthenticatorsOk() (*Authenticator, bool)`

GetSupportedAuthenticatorsOk returns a tuple with the SupportedAuthenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthenticators

`func (o *StatusModel) SetSupportedAuthenticators(v Authenticator)`

SetSupportedAuthenticators sets SupportedAuthenticators field to given value.

### HasSupportedAuthenticators

`func (o *StatusModel) HasSupportedAuthenticators() bool`

HasSupportedAuthenticators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


