# AzureAdConnectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppPermission** | **string** |  | 
**TenantId** | **string** |  | 

## Methods

### NewAzureAdConnectionSettings

`func NewAzureAdConnectionSettings(appPermission string, tenantId string, ) *AzureAdConnectionSettings`

NewAzureAdConnectionSettings instantiates a new AzureAdConnectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAdConnectionSettingsWithDefaults

`func NewAzureAdConnectionSettingsWithDefaults() *AzureAdConnectionSettings`

NewAzureAdConnectionSettingsWithDefaults instantiates a new AzureAdConnectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPermission

`func (o *AzureAdConnectionSettings) GetAppPermission() string`

GetAppPermission returns the AppPermission field if non-nil, zero value otherwise.

### GetAppPermissionOk

`func (o *AzureAdConnectionSettings) GetAppPermissionOk() (*string, bool)`

GetAppPermissionOk returns a tuple with the AppPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPermission

`func (o *AzureAdConnectionSettings) SetAppPermission(v string)`

SetAppPermission sets AppPermission field to given value.


### GetTenantId

`func (o *AzureAdConnectionSettings) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureAdConnectionSettings) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureAdConnectionSettings) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


