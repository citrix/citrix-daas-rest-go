# Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedWebStoreURLs** | Pointer to [**[]AdminDomainURL**](AdminDomainURL.md) |  | [optional] 
**ServiceURLs** | Pointer to [**[]AdminDomainURL**](AdminDomainURL.md) |  | [optional] 

## Methods

### NewWorkspace

`func NewWorkspace() *Workspace`

NewWorkspace instantiates a new Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWithDefaults

`func NewWorkspaceWithDefaults() *Workspace`

NewWorkspaceWithDefaults instantiates a new Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedWebStoreURLs

`func (o *Workspace) GetAllowedWebStoreURLs() []AdminDomainURL`

GetAllowedWebStoreURLs returns the AllowedWebStoreURLs field if non-nil, zero value otherwise.

### GetAllowedWebStoreURLsOk

`func (o *Workspace) GetAllowedWebStoreURLsOk() (*[]AdminDomainURL, bool)`

GetAllowedWebStoreURLsOk returns a tuple with the AllowedWebStoreURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedWebStoreURLs

`func (o *Workspace) SetAllowedWebStoreURLs(v []AdminDomainURL)`

SetAllowedWebStoreURLs sets AllowedWebStoreURLs field to given value.

### HasAllowedWebStoreURLs

`func (o *Workspace) HasAllowedWebStoreURLs() bool`

HasAllowedWebStoreURLs returns a boolean if a field has been set.

### GetServiceURLs

`func (o *Workspace) GetServiceURLs() []AdminDomainURL`

GetServiceURLs returns the ServiceURLs field if non-nil, zero value otherwise.

### GetServiceURLsOk

`func (o *Workspace) GetServiceURLsOk() (*[]AdminDomainURL, bool)`

GetServiceURLsOk returns a tuple with the ServiceURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceURLs

`func (o *Workspace) SetServiceURLs(v []AdminDomainURL)`

SetServiceURLs sets ServiceURLs field to given value.

### HasServiceURLs

`func (o *Workspace) HasServiceURLs() bool`

HasServiceURLs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


