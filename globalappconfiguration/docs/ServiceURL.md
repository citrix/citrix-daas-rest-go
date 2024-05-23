# ServiceURL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigrationUrl** | Pointer to [**[]MigrationUrl**](MigrationUrl.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceURL

`func NewServiceURL() *ServiceURL`

NewServiceURL instantiates a new ServiceURL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceURLWithDefaults

`func NewServiceURLWithDefaults() *ServiceURL`

NewServiceURLWithDefaults instantiates a new ServiceURL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrationUrl

`func (o *ServiceURL) GetMigrationUrl() []MigrationUrl`

GetMigrationUrl returns the MigrationUrl field if non-nil, zero value otherwise.

### GetMigrationUrlOk

`func (o *ServiceURL) GetMigrationUrlOk() (*[]MigrationUrl, bool)`

GetMigrationUrlOk returns a tuple with the MigrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationUrl

`func (o *ServiceURL) SetMigrationUrl(v []MigrationUrl)`

SetMigrationUrl sets MigrationUrl field to given value.

### HasMigrationUrl

`func (o *ServiceURL) HasMigrationUrl() bool`

HasMigrationUrl returns a boolean if a field has been set.

### GetUrl

`func (o *ServiceURL) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceURL) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceURL) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServiceURL) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


