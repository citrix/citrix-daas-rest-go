# ClientDiscoveryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedWebStoreURLs** | Pointer to [**[]ClientDiscoveryURLModel**](ClientDiscoveryURLModel.md) |  | [optional] 
**Domain** | Pointer to [**Domain**](Domain.md) |  | [optional] 
**MatchedDomain** | Pointer to [**Domain**](Domain.md) |  | [optional] 
**PreferredURL** | Pointer to [**ClientPreferredURLModel**](ClientPreferredURLModel.md) |  | [optional] 
**ServiceURLs** | Pointer to [**[]ClientDiscoveryURLModel**](ClientDiscoveryURLModel.md) |  | [optional] 

## Methods

### NewClientDiscoveryModel

`func NewClientDiscoveryModel() *ClientDiscoveryModel`

NewClientDiscoveryModel instantiates a new ClientDiscoveryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDiscoveryModelWithDefaults

`func NewClientDiscoveryModelWithDefaults() *ClientDiscoveryModel`

NewClientDiscoveryModelWithDefaults instantiates a new ClientDiscoveryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedWebStoreURLs

`func (o *ClientDiscoveryModel) GetAllowedWebStoreURLs() []ClientDiscoveryURLModel`

GetAllowedWebStoreURLs returns the AllowedWebStoreURLs field if non-nil, zero value otherwise.

### GetAllowedWebStoreURLsOk

`func (o *ClientDiscoveryModel) GetAllowedWebStoreURLsOk() (*[]ClientDiscoveryURLModel, bool)`

GetAllowedWebStoreURLsOk returns a tuple with the AllowedWebStoreURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedWebStoreURLs

`func (o *ClientDiscoveryModel) SetAllowedWebStoreURLs(v []ClientDiscoveryURLModel)`

SetAllowedWebStoreURLs sets AllowedWebStoreURLs field to given value.

### HasAllowedWebStoreURLs

`func (o *ClientDiscoveryModel) HasAllowedWebStoreURLs() bool`

HasAllowedWebStoreURLs returns a boolean if a field has been set.

### GetDomain

`func (o *ClientDiscoveryModel) GetDomain() Domain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ClientDiscoveryModel) GetDomainOk() (*Domain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ClientDiscoveryModel) SetDomain(v Domain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ClientDiscoveryModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMatchedDomain

`func (o *ClientDiscoveryModel) GetMatchedDomain() Domain`

GetMatchedDomain returns the MatchedDomain field if non-nil, zero value otherwise.

### GetMatchedDomainOk

`func (o *ClientDiscoveryModel) GetMatchedDomainOk() (*Domain, bool)`

GetMatchedDomainOk returns a tuple with the MatchedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedDomain

`func (o *ClientDiscoveryModel) SetMatchedDomain(v Domain)`

SetMatchedDomain sets MatchedDomain field to given value.

### HasMatchedDomain

`func (o *ClientDiscoveryModel) HasMatchedDomain() bool`

HasMatchedDomain returns a boolean if a field has been set.

### GetPreferredURL

`func (o *ClientDiscoveryModel) GetPreferredURL() ClientPreferredURLModel`

GetPreferredURL returns the PreferredURL field if non-nil, zero value otherwise.

### GetPreferredURLOk

`func (o *ClientDiscoveryModel) GetPreferredURLOk() (*ClientPreferredURLModel, bool)`

GetPreferredURLOk returns a tuple with the PreferredURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredURL

`func (o *ClientDiscoveryModel) SetPreferredURL(v ClientPreferredURLModel)`

SetPreferredURL sets PreferredURL field to given value.

### HasPreferredURL

`func (o *ClientDiscoveryModel) HasPreferredURL() bool`

HasPreferredURL returns a boolean if a field has been set.

### GetServiceURLs

`func (o *ClientDiscoveryModel) GetServiceURLs() []ClientDiscoveryURLModel`

GetServiceURLs returns the ServiceURLs field if non-nil, zero value otherwise.

### GetServiceURLsOk

`func (o *ClientDiscoveryModel) GetServiceURLsOk() (*[]ClientDiscoveryURLModel, bool)`

GetServiceURLsOk returns a tuple with the ServiceURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceURLs

`func (o *ClientDiscoveryModel) SetServiceURLs(v []ClientDiscoveryURLModel)`

SetServiceURLs sets ServiceURLs field to given value.

### HasServiceURLs

`func (o *ClientDiscoveryModel) HasServiceURLs() bool`

HasServiceURLs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


