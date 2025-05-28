# CustomerDetailsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | ID of the customer | [optional] 
**SiteId** | Pointer to **string** | ID of the customer&#39;s site | [optional] 
**Catalogs** | Pointer to [**[]CatalogConfiguration**](CatalogConfiguration.md) | List of catalogs configured by the user | [optional] 
**Images** | Pointer to **[]map[string]map[string]interface{}** | List of Images configured by the user | [optional] 
**Domains** | Pointer to **[]map[string]map[string]interface{}** | List of domains configured by the user | [optional] 
**ResourceLocations** | Pointer to **[]map[string]map[string]interface{}** | List of Resource Locations configured for the user | [optional] 
**Directories** | Pointer to **[]map[string]map[string]interface{}** | List of directories configured by the user | [optional] 
**OnpremConnections** | Pointer to **[]map[string]map[string]interface{}** | List of onprem connections configured by the user | [optional] 

## Methods

### NewCustomerDetailsModel

`func NewCustomerDetailsModel() *CustomerDetailsModel`

NewCustomerDetailsModel instantiates a new CustomerDetailsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDetailsModelWithDefaults

`func NewCustomerDetailsModelWithDefaults() *CustomerDetailsModel`

NewCustomerDetailsModelWithDefaults instantiates a new CustomerDetailsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CustomerDetailsModel) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerDetailsModel) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerDetailsModel) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CustomerDetailsModel) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetSiteId

`func (o *CustomerDetailsModel) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CustomerDetailsModel) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CustomerDetailsModel) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CustomerDetailsModel) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetCatalogs

`func (o *CustomerDetailsModel) GetCatalogs() []CatalogConfiguration`

GetCatalogs returns the Catalogs field if non-nil, zero value otherwise.

### GetCatalogsOk

`func (o *CustomerDetailsModel) GetCatalogsOk() (*[]CatalogConfiguration, bool)`

GetCatalogsOk returns a tuple with the Catalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogs

`func (o *CustomerDetailsModel) SetCatalogs(v []CatalogConfiguration)`

SetCatalogs sets Catalogs field to given value.

### HasCatalogs

`func (o *CustomerDetailsModel) HasCatalogs() bool`

HasCatalogs returns a boolean if a field has been set.

### GetImages

`func (o *CustomerDetailsModel) GetImages() []map[string]map[string]interface{}`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CustomerDetailsModel) GetImagesOk() (*[]map[string]map[string]interface{}, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CustomerDetailsModel) SetImages(v []map[string]map[string]interface{})`

SetImages sets Images field to given value.

### HasImages

`func (o *CustomerDetailsModel) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetDomains

`func (o *CustomerDetailsModel) GetDomains() []map[string]map[string]interface{}`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CustomerDetailsModel) GetDomainsOk() (*[]map[string]map[string]interface{}, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CustomerDetailsModel) SetDomains(v []map[string]map[string]interface{})`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CustomerDetailsModel) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetResourceLocations

`func (o *CustomerDetailsModel) GetResourceLocations() []map[string]map[string]interface{}`

GetResourceLocations returns the ResourceLocations field if non-nil, zero value otherwise.

### GetResourceLocationsOk

`func (o *CustomerDetailsModel) GetResourceLocationsOk() (*[]map[string]map[string]interface{}, bool)`

GetResourceLocationsOk returns a tuple with the ResourceLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocations

`func (o *CustomerDetailsModel) SetResourceLocations(v []map[string]map[string]interface{})`

SetResourceLocations sets ResourceLocations field to given value.

### HasResourceLocations

`func (o *CustomerDetailsModel) HasResourceLocations() bool`

HasResourceLocations returns a boolean if a field has been set.

### GetDirectories

`func (o *CustomerDetailsModel) GetDirectories() []map[string]map[string]interface{}`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *CustomerDetailsModel) GetDirectoriesOk() (*[]map[string]map[string]interface{}, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *CustomerDetailsModel) SetDirectories(v []map[string]map[string]interface{})`

SetDirectories sets Directories field to given value.

### HasDirectories

`func (o *CustomerDetailsModel) HasDirectories() bool`

HasDirectories returns a boolean if a field has been set.

### GetOnpremConnections

`func (o *CustomerDetailsModel) GetOnpremConnections() []map[string]map[string]interface{}`

GetOnpremConnections returns the OnpremConnections field if non-nil, zero value otherwise.

### GetOnpremConnectionsOk

`func (o *CustomerDetailsModel) GetOnpremConnectionsOk() (*[]map[string]map[string]interface{}, bool)`

GetOnpremConnectionsOk returns a tuple with the OnpremConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnpremConnections

`func (o *CustomerDetailsModel) SetOnpremConnections(v []map[string]map[string]interface{})`

SetOnpremConnections sets OnpremConnections field to given value.

### HasOnpremConnections

`func (o *CustomerDetailsModel) HasOnpremConnections() bool`

HasOnpremConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


