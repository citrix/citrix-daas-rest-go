# EditHypervisorAzureConnectionRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** | Application ID of the service principal used to access the Azure APIs.  Optional.  If not specified, will not be changed.  If specified, then ApplicationSecret must also be specified. | [optional] 
**ApplicationSecret** | Pointer to **string** | The Application Secret of the service principal used to access the Azure APIs.  Optional.  If not specified, will not be changed.  If specified, must in the format indicated by ApplicationSecretFormat. | [optional] 
**ApplicationSecretFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**CustomProperties** | Pointer to **string** | The properties of host connection that are specific to the target hosting infrastructure. | [optional] 

## Methods

### NewEditHypervisorAzureConnectionRequestModelAllOf

`func NewEditHypervisorAzureConnectionRequestModelAllOf() *EditHypervisorAzureConnectionRequestModelAllOf`

NewEditHypervisorAzureConnectionRequestModelAllOf instantiates a new EditHypervisorAzureConnectionRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHypervisorAzureConnectionRequestModelAllOfWithDefaults

`func NewEditHypervisorAzureConnectionRequestModelAllOfWithDefaults() *EditHypervisorAzureConnectionRequestModelAllOf`

NewEditHypervisorAzureConnectionRequestModelAllOfWithDefaults instantiates a new EditHypervisorAzureConnectionRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationSecret

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationSecret() string`

GetApplicationSecret returns the ApplicationSecret field if non-nil, zero value otherwise.

### GetApplicationSecretOk

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationSecretOk() (*string, bool)`

GetApplicationSecretOk returns a tuple with the ApplicationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecret

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) SetApplicationSecret(v string)`

SetApplicationSecret sets ApplicationSecret field to given value.

### HasApplicationSecret

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) HasApplicationSecret() bool`

HasApplicationSecret returns a boolean if a field has been set.

### GetApplicationSecretFormat

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationSecretFormat() IdentityPasswordFormat`

GetApplicationSecretFormat returns the ApplicationSecretFormat field if non-nil, zero value otherwise.

### GetApplicationSecretFormatOk

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetApplicationSecretFormatOk() (*IdentityPasswordFormat, bool)`

GetApplicationSecretFormatOk returns a tuple with the ApplicationSecretFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecretFormat

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) SetApplicationSecretFormat(v IdentityPasswordFormat)`

SetApplicationSecretFormat sets ApplicationSecretFormat field to given value.

### HasApplicationSecretFormat

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) HasApplicationSecretFormat() bool`

HasApplicationSecretFormat returns a boolean if a field has been set.

### GetCustomProperties

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *EditHypervisorAzureConnectionRequestModelAllOf) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


