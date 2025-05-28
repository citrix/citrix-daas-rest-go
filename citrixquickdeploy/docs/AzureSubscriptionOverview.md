# AzureSubscriptionOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedImages** | Pointer to **int32** | Number of Images that are stored within this subscription | [optional] 
**LinkedCatalogs** | Pointer to **int32** | Number of Catalogs that have their resources stored in this subscription | [optional] 
**LinkedAdvancedCatalogs** | Pointer to **int32** | Number of SimpleUI Advanced Catalogs that may have their resources stored in this subscription | [optional] [readonly] 
**Directory** | Pointer to **string** | The Azure Active Directory information for the directory tied to this subscription | [optional] [readonly] 
**DirectoryDomainName** | Pointer to **string** |  | [optional] [readonly] 
**LinkedCatalogsIds** | Pointer to **[]string** |  | [optional] [readonly] 
**LinkedImagesIds** | Pointer to **[]string** |  | [optional] [readonly] 
**Warnings** | Pointer to **[]string** |  | [optional] [readonly] 
**CatalogVMs** | Pointer to **int32** | Number of VMs that have been created in the catalogs that are associated with this subscription. | [optional] 
**SubscriptionId** | **string** | Id of the customers subscription | 
**Name** | **string** | Name of the Azure Subscription to display to the user | 
**IsLinked** | Pointer to **bool** | Indicates if this subscription is linked and we have the ability to manage resources within it | [optional] 
**IsValid** | Pointer to **bool** | Indicates whether catalog service can access the subscription or not | [optional] 
**CitrixManaged** | Pointer to **bool** | Indicates whether subscription is managed by Citrix or not | [optional] 
**CitrixManagedIndex** | Pointer to **int32** | Which Citrix Managed subscription is this? For the first Citrix Managedsubscription this is set to null | [optional] 
**Disabled** | Pointer to **bool** | Indicates if this subscription should not be used | [optional] 
**CspCustomer** | Pointer to **string** | Indicates that partner-tenant relationship exists if not null | [optional] 
**AzureAppId** | Pointer to **string** | ID of the Azure App User provided to manage the subscription | [optional] 
**ObfuscatedAzureAppSecret** | Pointer to **string** | An obfuscated form of the client secret that only shows the first few characters | [optional] 
**SecretExpirationDate** | Pointer to **time.Time** | The expiration date of the user provided secret, if one was used | [optional] 
**TenantId** | Pointer to **string** | Id of the directory that the user provided to manage the subscription | [optional] 
**UpdateJob** | Pointer to [**DataStoreAzureSubscriptionUpdateCredentialsJobModel**](DataStoreAzureSubscriptionUpdateCredentialsJobModel.md) | Details of job to update the azure app credentials | [optional] 
**UserConsentDetails** | Pointer to [**DataStoreManagedSubscriptionConsentModel**](DataStoreManagedSubscriptionConsentModel.md) | The details of consent obtained for use of managed subscription | [optional] 

## Methods

### NewAzureSubscriptionOverview

`func NewAzureSubscriptionOverview(subscriptionId string, name string, ) *AzureSubscriptionOverview`

NewAzureSubscriptionOverview instantiates a new AzureSubscriptionOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSubscriptionOverviewWithDefaults

`func NewAzureSubscriptionOverviewWithDefaults() *AzureSubscriptionOverview`

NewAzureSubscriptionOverviewWithDefaults instantiates a new AzureSubscriptionOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkedImages

`func (o *AzureSubscriptionOverview) GetLinkedImages() int32`

GetLinkedImages returns the LinkedImages field if non-nil, zero value otherwise.

### GetLinkedImagesOk

`func (o *AzureSubscriptionOverview) GetLinkedImagesOk() (*int32, bool)`

GetLinkedImagesOk returns a tuple with the LinkedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedImages

`func (o *AzureSubscriptionOverview) SetLinkedImages(v int32)`

SetLinkedImages sets LinkedImages field to given value.

### HasLinkedImages

`func (o *AzureSubscriptionOverview) HasLinkedImages() bool`

HasLinkedImages returns a boolean if a field has been set.

### GetLinkedCatalogs

`func (o *AzureSubscriptionOverview) GetLinkedCatalogs() int32`

GetLinkedCatalogs returns the LinkedCatalogs field if non-nil, zero value otherwise.

### GetLinkedCatalogsOk

`func (o *AzureSubscriptionOverview) GetLinkedCatalogsOk() (*int32, bool)`

GetLinkedCatalogsOk returns a tuple with the LinkedCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCatalogs

`func (o *AzureSubscriptionOverview) SetLinkedCatalogs(v int32)`

SetLinkedCatalogs sets LinkedCatalogs field to given value.

### HasLinkedCatalogs

`func (o *AzureSubscriptionOverview) HasLinkedCatalogs() bool`

HasLinkedCatalogs returns a boolean if a field has been set.

### GetLinkedAdvancedCatalogs

`func (o *AzureSubscriptionOverview) GetLinkedAdvancedCatalogs() int32`

GetLinkedAdvancedCatalogs returns the LinkedAdvancedCatalogs field if non-nil, zero value otherwise.

### GetLinkedAdvancedCatalogsOk

`func (o *AzureSubscriptionOverview) GetLinkedAdvancedCatalogsOk() (*int32, bool)`

GetLinkedAdvancedCatalogsOk returns a tuple with the LinkedAdvancedCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAdvancedCatalogs

`func (o *AzureSubscriptionOverview) SetLinkedAdvancedCatalogs(v int32)`

SetLinkedAdvancedCatalogs sets LinkedAdvancedCatalogs field to given value.

### HasLinkedAdvancedCatalogs

`func (o *AzureSubscriptionOverview) HasLinkedAdvancedCatalogs() bool`

HasLinkedAdvancedCatalogs returns a boolean if a field has been set.

### GetDirectory

`func (o *AzureSubscriptionOverview) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *AzureSubscriptionOverview) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *AzureSubscriptionOverview) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *AzureSubscriptionOverview) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetDirectoryDomainName

`func (o *AzureSubscriptionOverview) GetDirectoryDomainName() string`

GetDirectoryDomainName returns the DirectoryDomainName field if non-nil, zero value otherwise.

### GetDirectoryDomainNameOk

`func (o *AzureSubscriptionOverview) GetDirectoryDomainNameOk() (*string, bool)`

GetDirectoryDomainNameOk returns a tuple with the DirectoryDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryDomainName

`func (o *AzureSubscriptionOverview) SetDirectoryDomainName(v string)`

SetDirectoryDomainName sets DirectoryDomainName field to given value.

### HasDirectoryDomainName

`func (o *AzureSubscriptionOverview) HasDirectoryDomainName() bool`

HasDirectoryDomainName returns a boolean if a field has been set.

### GetLinkedCatalogsIds

`func (o *AzureSubscriptionOverview) GetLinkedCatalogsIds() []string`

GetLinkedCatalogsIds returns the LinkedCatalogsIds field if non-nil, zero value otherwise.

### GetLinkedCatalogsIdsOk

`func (o *AzureSubscriptionOverview) GetLinkedCatalogsIdsOk() (*[]string, bool)`

GetLinkedCatalogsIdsOk returns a tuple with the LinkedCatalogsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCatalogsIds

`func (o *AzureSubscriptionOverview) SetLinkedCatalogsIds(v []string)`

SetLinkedCatalogsIds sets LinkedCatalogsIds field to given value.

### HasLinkedCatalogsIds

`func (o *AzureSubscriptionOverview) HasLinkedCatalogsIds() bool`

HasLinkedCatalogsIds returns a boolean if a field has been set.

### GetLinkedImagesIds

`func (o *AzureSubscriptionOverview) GetLinkedImagesIds() []string`

GetLinkedImagesIds returns the LinkedImagesIds field if non-nil, zero value otherwise.

### GetLinkedImagesIdsOk

`func (o *AzureSubscriptionOverview) GetLinkedImagesIdsOk() (*[]string, bool)`

GetLinkedImagesIdsOk returns a tuple with the LinkedImagesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedImagesIds

`func (o *AzureSubscriptionOverview) SetLinkedImagesIds(v []string)`

SetLinkedImagesIds sets LinkedImagesIds field to given value.

### HasLinkedImagesIds

`func (o *AzureSubscriptionOverview) HasLinkedImagesIds() bool`

HasLinkedImagesIds returns a boolean if a field has been set.

### GetWarnings

`func (o *AzureSubscriptionOverview) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AzureSubscriptionOverview) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AzureSubscriptionOverview) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AzureSubscriptionOverview) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetCatalogVMs

`func (o *AzureSubscriptionOverview) GetCatalogVMs() int32`

GetCatalogVMs returns the CatalogVMs field if non-nil, zero value otherwise.

### GetCatalogVMsOk

`func (o *AzureSubscriptionOverview) GetCatalogVMsOk() (*int32, bool)`

GetCatalogVMsOk returns a tuple with the CatalogVMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogVMs

`func (o *AzureSubscriptionOverview) SetCatalogVMs(v int32)`

SetCatalogVMs sets CatalogVMs field to given value.

### HasCatalogVMs

`func (o *AzureSubscriptionOverview) HasCatalogVMs() bool`

HasCatalogVMs returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AzureSubscriptionOverview) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureSubscriptionOverview) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureSubscriptionOverview) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetName

`func (o *AzureSubscriptionOverview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureSubscriptionOverview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureSubscriptionOverview) SetName(v string)`

SetName sets Name field to given value.


### GetIsLinked

`func (o *AzureSubscriptionOverview) GetIsLinked() bool`

GetIsLinked returns the IsLinked field if non-nil, zero value otherwise.

### GetIsLinkedOk

`func (o *AzureSubscriptionOverview) GetIsLinkedOk() (*bool, bool)`

GetIsLinkedOk returns a tuple with the IsLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinked

`func (o *AzureSubscriptionOverview) SetIsLinked(v bool)`

SetIsLinked sets IsLinked field to given value.

### HasIsLinked

`func (o *AzureSubscriptionOverview) HasIsLinked() bool`

HasIsLinked returns a boolean if a field has been set.

### GetIsValid

`func (o *AzureSubscriptionOverview) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *AzureSubscriptionOverview) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *AzureSubscriptionOverview) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *AzureSubscriptionOverview) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetCitrixManaged

`func (o *AzureSubscriptionOverview) GetCitrixManaged() bool`

GetCitrixManaged returns the CitrixManaged field if non-nil, zero value otherwise.

### GetCitrixManagedOk

`func (o *AzureSubscriptionOverview) GetCitrixManagedOk() (*bool, bool)`

GetCitrixManagedOk returns a tuple with the CitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManaged

`func (o *AzureSubscriptionOverview) SetCitrixManaged(v bool)`

SetCitrixManaged sets CitrixManaged field to given value.

### HasCitrixManaged

`func (o *AzureSubscriptionOverview) HasCitrixManaged() bool`

HasCitrixManaged returns a boolean if a field has been set.

### GetCitrixManagedIndex

`func (o *AzureSubscriptionOverview) GetCitrixManagedIndex() int32`

GetCitrixManagedIndex returns the CitrixManagedIndex field if non-nil, zero value otherwise.

### GetCitrixManagedIndexOk

`func (o *AzureSubscriptionOverview) GetCitrixManagedIndexOk() (*int32, bool)`

GetCitrixManagedIndexOk returns a tuple with the CitrixManagedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManagedIndex

`func (o *AzureSubscriptionOverview) SetCitrixManagedIndex(v int32)`

SetCitrixManagedIndex sets CitrixManagedIndex field to given value.

### HasCitrixManagedIndex

`func (o *AzureSubscriptionOverview) HasCitrixManagedIndex() bool`

HasCitrixManagedIndex returns a boolean if a field has been set.

### GetDisabled

`func (o *AzureSubscriptionOverview) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AzureSubscriptionOverview) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AzureSubscriptionOverview) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AzureSubscriptionOverview) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetCspCustomer

`func (o *AzureSubscriptionOverview) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *AzureSubscriptionOverview) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *AzureSubscriptionOverview) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *AzureSubscriptionOverview) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### GetAzureAppId

`func (o *AzureSubscriptionOverview) GetAzureAppId() string`

GetAzureAppId returns the AzureAppId field if non-nil, zero value otherwise.

### GetAzureAppIdOk

`func (o *AzureSubscriptionOverview) GetAzureAppIdOk() (*string, bool)`

GetAzureAppIdOk returns a tuple with the AzureAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppId

`func (o *AzureSubscriptionOverview) SetAzureAppId(v string)`

SetAzureAppId sets AzureAppId field to given value.

### HasAzureAppId

`func (o *AzureSubscriptionOverview) HasAzureAppId() bool`

HasAzureAppId returns a boolean if a field has been set.

### GetObfuscatedAzureAppSecret

`func (o *AzureSubscriptionOverview) GetObfuscatedAzureAppSecret() string`

GetObfuscatedAzureAppSecret returns the ObfuscatedAzureAppSecret field if non-nil, zero value otherwise.

### GetObfuscatedAzureAppSecretOk

`func (o *AzureSubscriptionOverview) GetObfuscatedAzureAppSecretOk() (*string, bool)`

GetObfuscatedAzureAppSecretOk returns a tuple with the ObfuscatedAzureAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscatedAzureAppSecret

`func (o *AzureSubscriptionOverview) SetObfuscatedAzureAppSecret(v string)`

SetObfuscatedAzureAppSecret sets ObfuscatedAzureAppSecret field to given value.

### HasObfuscatedAzureAppSecret

`func (o *AzureSubscriptionOverview) HasObfuscatedAzureAppSecret() bool`

HasObfuscatedAzureAppSecret returns a boolean if a field has been set.

### GetSecretExpirationDate

`func (o *AzureSubscriptionOverview) GetSecretExpirationDate() time.Time`

GetSecretExpirationDate returns the SecretExpirationDate field if non-nil, zero value otherwise.

### GetSecretExpirationDateOk

`func (o *AzureSubscriptionOverview) GetSecretExpirationDateOk() (*time.Time, bool)`

GetSecretExpirationDateOk returns a tuple with the SecretExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExpirationDate

`func (o *AzureSubscriptionOverview) SetSecretExpirationDate(v time.Time)`

SetSecretExpirationDate sets SecretExpirationDate field to given value.

### HasSecretExpirationDate

`func (o *AzureSubscriptionOverview) HasSecretExpirationDate() bool`

HasSecretExpirationDate returns a boolean if a field has been set.

### GetTenantId

`func (o *AzureSubscriptionOverview) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureSubscriptionOverview) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureSubscriptionOverview) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureSubscriptionOverview) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdateJob

`func (o *AzureSubscriptionOverview) GetUpdateJob() DataStoreAzureSubscriptionUpdateCredentialsJobModel`

GetUpdateJob returns the UpdateJob field if non-nil, zero value otherwise.

### GetUpdateJobOk

`func (o *AzureSubscriptionOverview) GetUpdateJobOk() (*DataStoreAzureSubscriptionUpdateCredentialsJobModel, bool)`

GetUpdateJobOk returns a tuple with the UpdateJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateJob

`func (o *AzureSubscriptionOverview) SetUpdateJob(v DataStoreAzureSubscriptionUpdateCredentialsJobModel)`

SetUpdateJob sets UpdateJob field to given value.

### HasUpdateJob

`func (o *AzureSubscriptionOverview) HasUpdateJob() bool`

HasUpdateJob returns a boolean if a field has been set.

### GetUserConsentDetails

`func (o *AzureSubscriptionOverview) GetUserConsentDetails() DataStoreManagedSubscriptionConsentModel`

GetUserConsentDetails returns the UserConsentDetails field if non-nil, zero value otherwise.

### GetUserConsentDetailsOk

`func (o *AzureSubscriptionOverview) GetUserConsentDetailsOk() (*DataStoreManagedSubscriptionConsentModel, bool)`

GetUserConsentDetailsOk returns a tuple with the UserConsentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserConsentDetails

`func (o *AzureSubscriptionOverview) SetUserConsentDetails(v DataStoreManagedSubscriptionConsentModel)`

SetUserConsentDetails sets UserConsentDetails field to given value.

### HasUserConsentDetails

`func (o *AzureSubscriptionOverview) HasUserConsentDetails() bool`

HasUserConsentDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


