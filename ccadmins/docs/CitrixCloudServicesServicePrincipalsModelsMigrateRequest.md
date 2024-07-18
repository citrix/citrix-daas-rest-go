# CitrixCloudServicesServicePrincipalsModelsMigrateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecureClients** | **[]string** |  | 
**ExpirationDate** | **time.Time** |  | 

## Methods

### NewCitrixCloudServicesServicePrincipalsModelsMigrateRequest

`func NewCitrixCloudServicesServicePrincipalsModelsMigrateRequest(secureClients []string, expirationDate time.Time, ) *CitrixCloudServicesServicePrincipalsModelsMigrateRequest`

NewCitrixCloudServicesServicePrincipalsModelsMigrateRequest instantiates a new CitrixCloudServicesServicePrincipalsModelsMigrateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesServicePrincipalsModelsMigrateRequestWithDefaults

`func NewCitrixCloudServicesServicePrincipalsModelsMigrateRequestWithDefaults() *CitrixCloudServicesServicePrincipalsModelsMigrateRequest`

NewCitrixCloudServicesServicePrincipalsModelsMigrateRequestWithDefaults instantiates a new CitrixCloudServicesServicePrincipalsModelsMigrateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecureClients

`func (o *CitrixCloudServicesServicePrincipalsModelsMigrateRequest) GetSecureClients() []string`

GetSecureClients returns the SecureClients field if non-nil, zero value otherwise.

### GetSecureClientsOk

`func (o *CitrixCloudServicesServicePrincipalsModelsMigrateRequest) GetSecureClientsOk() (*[]string, bool)`

GetSecureClientsOk returns a tuple with the SecureClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureClients

`func (o *CitrixCloudServicesServicePrincipalsModelsMigrateRequest) SetSecureClients(v []string)`

SetSecureClients sets SecureClients field to given value.


### GetExpirationDate

`func (o *CitrixCloudServicesServicePrincipalsModelsMigrateRequest) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CitrixCloudServicesServicePrincipalsModelsMigrateRequest) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CitrixCloudServicesServicePrincipalsModelsMigrateRequest) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


