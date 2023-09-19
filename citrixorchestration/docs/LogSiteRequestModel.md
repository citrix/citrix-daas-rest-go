# LogSiteRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggingEnabled** | Pointer to **bool** | Configuration logging settings enabled. | [optional] 
**AllowChangesWhileDatabaseDown** | Pointer to **bool** | Allow changes to the site configuration when the database is not available. | [optional] 
**Locale** | Pointer to **NullableString** | Locale. | [optional] 

## Methods

### NewLogSiteRequestModel

`func NewLogSiteRequestModel() *LogSiteRequestModel`

NewLogSiteRequestModel instantiates a new LogSiteRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSiteRequestModelWithDefaults

`func NewLogSiteRequestModelWithDefaults() *LogSiteRequestModel`

NewLogSiteRequestModelWithDefaults instantiates a new LogSiteRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggingEnabled

`func (o *LogSiteRequestModel) GetLoggingEnabled() bool`

GetLoggingEnabled returns the LoggingEnabled field if non-nil, zero value otherwise.

### GetLoggingEnabledOk

`func (o *LogSiteRequestModel) GetLoggingEnabledOk() (*bool, bool)`

GetLoggingEnabledOk returns a tuple with the LoggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingEnabled

`func (o *LogSiteRequestModel) SetLoggingEnabled(v bool)`

SetLoggingEnabled sets LoggingEnabled field to given value.

### HasLoggingEnabled

`func (o *LogSiteRequestModel) HasLoggingEnabled() bool`

HasLoggingEnabled returns a boolean if a field has been set.

### GetAllowChangesWhileDatabaseDown

`func (o *LogSiteRequestModel) GetAllowChangesWhileDatabaseDown() bool`

GetAllowChangesWhileDatabaseDown returns the AllowChangesWhileDatabaseDown field if non-nil, zero value otherwise.

### GetAllowChangesWhileDatabaseDownOk

`func (o *LogSiteRequestModel) GetAllowChangesWhileDatabaseDownOk() (*bool, bool)`

GetAllowChangesWhileDatabaseDownOk returns a tuple with the AllowChangesWhileDatabaseDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChangesWhileDatabaseDown

`func (o *LogSiteRequestModel) SetAllowChangesWhileDatabaseDown(v bool)`

SetAllowChangesWhileDatabaseDown sets AllowChangesWhileDatabaseDown field to given value.

### HasAllowChangesWhileDatabaseDown

`func (o *LogSiteRequestModel) HasAllowChangesWhileDatabaseDown() bool`

HasAllowChangesWhileDatabaseDown returns a boolean if a field has been set.

### GetLocale

`func (o *LogSiteRequestModel) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *LogSiteRequestModel) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *LogSiteRequestModel) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *LogSiteRequestModel) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *LogSiteRequestModel) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *LogSiteRequestModel) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


