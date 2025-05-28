# TemplateImageAssociatedCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogId** | **string** | ID of the catalog | 
**CatalogName** | **string** | Name of the catalog | 
**State** | [**CatalogOverallState**](CatalogOverallState.md) | State of the catalog | 
**AzureSubscription** | **string** | Name of the Azure Subscription the catalog is associated with | 

## Methods

### NewTemplateImageAssociatedCatalog

`func NewTemplateImageAssociatedCatalog(catalogId string, catalogName string, state CatalogOverallState, azureSubscription string, ) *TemplateImageAssociatedCatalog`

NewTemplateImageAssociatedCatalog instantiates a new TemplateImageAssociatedCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateImageAssociatedCatalogWithDefaults

`func NewTemplateImageAssociatedCatalogWithDefaults() *TemplateImageAssociatedCatalog`

NewTemplateImageAssociatedCatalogWithDefaults instantiates a new TemplateImageAssociatedCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogId

`func (o *TemplateImageAssociatedCatalog) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *TemplateImageAssociatedCatalog) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *TemplateImageAssociatedCatalog) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.


### GetCatalogName

`func (o *TemplateImageAssociatedCatalog) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *TemplateImageAssociatedCatalog) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *TemplateImageAssociatedCatalog) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.


### GetState

`func (o *TemplateImageAssociatedCatalog) GetState() CatalogOverallState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TemplateImageAssociatedCatalog) GetStateOk() (*CatalogOverallState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TemplateImageAssociatedCatalog) SetState(v CatalogOverallState)`

SetState sets State field to given value.


### GetAzureSubscription

`func (o *TemplateImageAssociatedCatalog) GetAzureSubscription() string`

GetAzureSubscription returns the AzureSubscription field if non-nil, zero value otherwise.

### GetAzureSubscriptionOk

`func (o *TemplateImageAssociatedCatalog) GetAzureSubscriptionOk() (*string, bool)`

GetAzureSubscriptionOk returns a tuple with the AzureSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscription

`func (o *TemplateImageAssociatedCatalog) SetAzureSubscription(v string)`

SetAzureSubscription sets AzureSubscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


