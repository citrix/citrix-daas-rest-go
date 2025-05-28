# AddTemplateImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Friendly name of the template | 
**SubscriptionId** | **string** | ID of the Azure Subscription | 
**ResourceGroup** | **string** | Name of the Resource Group | 
**StorageAccount** | **string** | Name of the storage account | 
**VhdUrl** | Pointer to **string** | Url of the VHD image within the storage account | [optional] 
**VhdEncryptionUri** | Pointer to **string** | URI of the VHD guest disk file that will be imported | [optional] 
**OsPlatform** | Pointer to [**SupportedOperatingSystemType**](SupportedOperatingSystemType.md) | Type of operating system that will be imported | [optional] 
**Notes** | Pointer to **string** | Customer notes about template image | [optional] 
**CspCustomerId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**HyperVGen** | Pointer to **string** | The HyperVGeneration that should be set to either V1 or V2 | [optional] 
**VtpmEnabled** | Pointer to **bool** | The HyperVGeneration V2 supports vTPM TrustedLaunch | [optional] 
**SecureBootEnabled** | Pointer to **bool** | The Secure boot support enabled | [optional] 
**CmekEnabled** | Pointer to **bool** | The customer managed encryption key enabled | [optional] 
**CmekID** | Pointer to **string** | The customer managed encryption ID | [optional] 
**Validate** | Pointer to **bool** | Indicates if the template image should be validated upon creation | [optional] 

## Methods

### NewAddTemplateImageModel

`func NewAddTemplateImageModel(name string, subscriptionId string, resourceGroup string, storageAccount string, ) *AddTemplateImageModel`

NewAddTemplateImageModel instantiates a new AddTemplateImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTemplateImageModelWithDefaults

`func NewAddTemplateImageModelWithDefaults() *AddTemplateImageModel`

NewAddTemplateImageModelWithDefaults instantiates a new AddTemplateImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddTemplateImageModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddTemplateImageModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddTemplateImageModel) SetName(v string)`

SetName sets Name field to given value.


### GetSubscriptionId

`func (o *AddTemplateImageModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AddTemplateImageModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AddTemplateImageModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroup

`func (o *AddTemplateImageModel) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AddTemplateImageModel) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AddTemplateImageModel) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetStorageAccount

`func (o *AddTemplateImageModel) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AddTemplateImageModel) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AddTemplateImageModel) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.


### GetVhdUrl

`func (o *AddTemplateImageModel) GetVhdUrl() string`

GetVhdUrl returns the VhdUrl field if non-nil, zero value otherwise.

### GetVhdUrlOk

`func (o *AddTemplateImageModel) GetVhdUrlOk() (*string, bool)`

GetVhdUrlOk returns a tuple with the VhdUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhdUrl

`func (o *AddTemplateImageModel) SetVhdUrl(v string)`

SetVhdUrl sets VhdUrl field to given value.

### HasVhdUrl

`func (o *AddTemplateImageModel) HasVhdUrl() bool`

HasVhdUrl returns a boolean if a field has been set.

### GetVhdEncryptionUri

`func (o *AddTemplateImageModel) GetVhdEncryptionUri() string`

GetVhdEncryptionUri returns the VhdEncryptionUri field if non-nil, zero value otherwise.

### GetVhdEncryptionUriOk

`func (o *AddTemplateImageModel) GetVhdEncryptionUriOk() (*string, bool)`

GetVhdEncryptionUriOk returns a tuple with the VhdEncryptionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhdEncryptionUri

`func (o *AddTemplateImageModel) SetVhdEncryptionUri(v string)`

SetVhdEncryptionUri sets VhdEncryptionUri field to given value.

### HasVhdEncryptionUri

`func (o *AddTemplateImageModel) HasVhdEncryptionUri() bool`

HasVhdEncryptionUri returns a boolean if a field has been set.

### GetOsPlatform

`func (o *AddTemplateImageModel) GetOsPlatform() SupportedOperatingSystemType`

GetOsPlatform returns the OsPlatform field if non-nil, zero value otherwise.

### GetOsPlatformOk

`func (o *AddTemplateImageModel) GetOsPlatformOk() (*SupportedOperatingSystemType, bool)`

GetOsPlatformOk returns a tuple with the OsPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPlatform

`func (o *AddTemplateImageModel) SetOsPlatform(v SupportedOperatingSystemType)`

SetOsPlatform sets OsPlatform field to given value.

### HasOsPlatform

`func (o *AddTemplateImageModel) HasOsPlatform() bool`

HasOsPlatform returns a boolean if a field has been set.

### GetNotes

`func (o *AddTemplateImageModel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddTemplateImageModel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddTemplateImageModel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddTemplateImageModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCspCustomerId

`func (o *AddTemplateImageModel) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *AddTemplateImageModel) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *AddTemplateImageModel) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *AddTemplateImageModel) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### GetCspSiteId

`func (o *AddTemplateImageModel) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *AddTemplateImageModel) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *AddTemplateImageModel) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *AddTemplateImageModel) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.

### GetHyperVGen

`func (o *AddTemplateImageModel) GetHyperVGen() string`

GetHyperVGen returns the HyperVGen field if non-nil, zero value otherwise.

### GetHyperVGenOk

`func (o *AddTemplateImageModel) GetHyperVGenOk() (*string, bool)`

GetHyperVGenOk returns a tuple with the HyperVGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperVGen

`func (o *AddTemplateImageModel) SetHyperVGen(v string)`

SetHyperVGen sets HyperVGen field to given value.

### HasHyperVGen

`func (o *AddTemplateImageModel) HasHyperVGen() bool`

HasHyperVGen returns a boolean if a field has been set.

### GetVtpmEnabled

`func (o *AddTemplateImageModel) GetVtpmEnabled() bool`

GetVtpmEnabled returns the VtpmEnabled field if non-nil, zero value otherwise.

### GetVtpmEnabledOk

`func (o *AddTemplateImageModel) GetVtpmEnabledOk() (*bool, bool)`

GetVtpmEnabledOk returns a tuple with the VtpmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpmEnabled

`func (o *AddTemplateImageModel) SetVtpmEnabled(v bool)`

SetVtpmEnabled sets VtpmEnabled field to given value.

### HasVtpmEnabled

`func (o *AddTemplateImageModel) HasVtpmEnabled() bool`

HasVtpmEnabled returns a boolean if a field has been set.

### GetSecureBootEnabled

`func (o *AddTemplateImageModel) GetSecureBootEnabled() bool`

GetSecureBootEnabled returns the SecureBootEnabled field if non-nil, zero value otherwise.

### GetSecureBootEnabledOk

`func (o *AddTemplateImageModel) GetSecureBootEnabledOk() (*bool, bool)`

GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBootEnabled

`func (o *AddTemplateImageModel) SetSecureBootEnabled(v bool)`

SetSecureBootEnabled sets SecureBootEnabled field to given value.

### HasSecureBootEnabled

`func (o *AddTemplateImageModel) HasSecureBootEnabled() bool`

HasSecureBootEnabled returns a boolean if a field has been set.

### GetCmekEnabled

`func (o *AddTemplateImageModel) GetCmekEnabled() bool`

GetCmekEnabled returns the CmekEnabled field if non-nil, zero value otherwise.

### GetCmekEnabledOk

`func (o *AddTemplateImageModel) GetCmekEnabledOk() (*bool, bool)`

GetCmekEnabledOk returns a tuple with the CmekEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmekEnabled

`func (o *AddTemplateImageModel) SetCmekEnabled(v bool)`

SetCmekEnabled sets CmekEnabled field to given value.

### HasCmekEnabled

`func (o *AddTemplateImageModel) HasCmekEnabled() bool`

HasCmekEnabled returns a boolean if a field has been set.

### GetCmekID

`func (o *AddTemplateImageModel) GetCmekID() string`

GetCmekID returns the CmekID field if non-nil, zero value otherwise.

### GetCmekIDOk

`func (o *AddTemplateImageModel) GetCmekIDOk() (*string, bool)`

GetCmekIDOk returns a tuple with the CmekID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmekID

`func (o *AddTemplateImageModel) SetCmekID(v string)`

SetCmekID sets CmekID field to given value.

### HasCmekID

`func (o *AddTemplateImageModel) HasCmekID() bool`

HasCmekID returns a boolean if a field has been set.

### GetValidate

`func (o *AddTemplateImageModel) GetValidate() bool`

GetValidate returns the Validate field if non-nil, zero value otherwise.

### GetValidateOk

`func (o *AddTemplateImageModel) GetValidateOk() (*bool, bool)`

GetValidateOk returns a tuple with the Validate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidate

`func (o *AddTemplateImageModel) SetValidate(v bool)`

SetValidate sets Validate field to given value.

### HasValidate

`func (o *AddTemplateImageModel) HasValidate() bool`

HasValidate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


