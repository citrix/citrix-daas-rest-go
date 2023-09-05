# LogSiteResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggingEnabled** | Pointer to **bool** | Configuration logging settings enabled. | [optional] 
**AllowChangesWhileDatabaseDown** | Pointer to **bool** | Allow changes to the site configuration when the database is not available. | [optional] 
**Locale** | Pointer to **string** | Locale. | [optional] 

## Methods

### NewLogSiteResponseModel

`func NewLogSiteResponseModel() *LogSiteResponseModel`

NewLogSiteResponseModel instantiates a new LogSiteResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSiteResponseModelWithDefaults

`func NewLogSiteResponseModelWithDefaults() *LogSiteResponseModel`

NewLogSiteResponseModelWithDefaults instantiates a new LogSiteResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggingEnabled

`func (o *LogSiteResponseModel) GetLoggingEnabled() bool`

GetLoggingEnabled returns the LoggingEnabled field if non-nil, zero value otherwise.

### GetLoggingEnabledOk

`func (o *LogSiteResponseModel) GetLoggingEnabledOk() (*bool, bool)`

GetLoggingEnabledOk returns a tuple with the LoggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingEnabled

`func (o *LogSiteResponseModel) SetLoggingEnabled(v bool)`

SetLoggingEnabled sets LoggingEnabled field to given value.

### HasLoggingEnabled

`func (o *LogSiteResponseModel) HasLoggingEnabled() bool`

HasLoggingEnabled returns a boolean if a field has been set.

### GetAllowChangesWhileDatabaseDown

`func (o *LogSiteResponseModel) GetAllowChangesWhileDatabaseDown() bool`

GetAllowChangesWhileDatabaseDown returns the AllowChangesWhileDatabaseDown field if non-nil, zero value otherwise.

### GetAllowChangesWhileDatabaseDownOk

`func (o *LogSiteResponseModel) GetAllowChangesWhileDatabaseDownOk() (*bool, bool)`

GetAllowChangesWhileDatabaseDownOk returns a tuple with the AllowChangesWhileDatabaseDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChangesWhileDatabaseDown

`func (o *LogSiteResponseModel) SetAllowChangesWhileDatabaseDown(v bool)`

SetAllowChangesWhileDatabaseDown sets AllowChangesWhileDatabaseDown field to given value.

### HasAllowChangesWhileDatabaseDown

`func (o *LogSiteResponseModel) HasAllowChangesWhileDatabaseDown() bool`

HasAllowChangesWhileDatabaseDown returns a boolean if a field has been set.

### GetLocale

`func (o *LogSiteResponseModel) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *LogSiteResponseModel) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *LogSiteResponseModel) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *LogSiteResponseModel) HasLocale() bool`

HasLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


