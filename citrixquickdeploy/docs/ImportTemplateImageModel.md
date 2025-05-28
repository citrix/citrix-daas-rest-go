# ImportTemplateImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the image | 
**VhdUri** | **string** | URI of the VHD file that will be imported | 
**VhdEncryptionUri** | Pointer to **string** | URI of the VHD guest disk file that will be imported | [optional] 
**Notes** | Pointer to **string** | Customer notes about template image | [optional] 
**OsPlatform** | Pointer to [**SupportedOperatingSystemType**](SupportedOperatingSystemType.md) | Type of operating system that will be imported | [optional] 
**CspCustomerId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**AzureSubscriptionId** | Pointer to **string** | The Id of the azure subscription where the image will be stored | [optional] 
**Region** | Pointer to **string** | The region where the storage account will be created for the image to be imported. | [optional] 
**HyperVGen** | Pointer to **string** | The HyperVGeneration that should be set to either V1 or V2 | [optional] 
**VtpmEnabled** | Pointer to **bool** | The HyperVGeneration V2 supports vTPM TrustedLaunch | [optional] 
**SecureBootEnabled** | Pointer to **bool** | The Secure boot support enabled | [optional] 
**CmekEnabled** | Pointer to **bool** | The customer managed encryption key enabled | [optional] 
**CmekID** | Pointer to **string** | The customer managed encryption ID | [optional] 

## Methods

### NewImportTemplateImageModel

`func NewImportTemplateImageModel(name string, vhdUri string, ) *ImportTemplateImageModel`

NewImportTemplateImageModel instantiates a new ImportTemplateImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportTemplateImageModelWithDefaults

`func NewImportTemplateImageModelWithDefaults() *ImportTemplateImageModel`

NewImportTemplateImageModelWithDefaults instantiates a new ImportTemplateImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImportTemplateImageModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportTemplateImageModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportTemplateImageModel) SetName(v string)`

SetName sets Name field to given value.


### GetVhdUri

`func (o *ImportTemplateImageModel) GetVhdUri() string`

GetVhdUri returns the VhdUri field if non-nil, zero value otherwise.

### GetVhdUriOk

`func (o *ImportTemplateImageModel) GetVhdUriOk() (*string, bool)`

GetVhdUriOk returns a tuple with the VhdUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhdUri

`func (o *ImportTemplateImageModel) SetVhdUri(v string)`

SetVhdUri sets VhdUri field to given value.


### GetVhdEncryptionUri

`func (o *ImportTemplateImageModel) GetVhdEncryptionUri() string`

GetVhdEncryptionUri returns the VhdEncryptionUri field if non-nil, zero value otherwise.

### GetVhdEncryptionUriOk

`func (o *ImportTemplateImageModel) GetVhdEncryptionUriOk() (*string, bool)`

GetVhdEncryptionUriOk returns a tuple with the VhdEncryptionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhdEncryptionUri

`func (o *ImportTemplateImageModel) SetVhdEncryptionUri(v string)`

SetVhdEncryptionUri sets VhdEncryptionUri field to given value.

### HasVhdEncryptionUri

`func (o *ImportTemplateImageModel) HasVhdEncryptionUri() bool`

HasVhdEncryptionUri returns a boolean if a field has been set.

### GetNotes

`func (o *ImportTemplateImageModel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ImportTemplateImageModel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ImportTemplateImageModel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ImportTemplateImageModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOsPlatform

`func (o *ImportTemplateImageModel) GetOsPlatform() SupportedOperatingSystemType`

GetOsPlatform returns the OsPlatform field if non-nil, zero value otherwise.

### GetOsPlatformOk

`func (o *ImportTemplateImageModel) GetOsPlatformOk() (*SupportedOperatingSystemType, bool)`

GetOsPlatformOk returns a tuple with the OsPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPlatform

`func (o *ImportTemplateImageModel) SetOsPlatform(v SupportedOperatingSystemType)`

SetOsPlatform sets OsPlatform field to given value.

### HasOsPlatform

`func (o *ImportTemplateImageModel) HasOsPlatform() bool`

HasOsPlatform returns a boolean if a field has been set.

### GetCspCustomerId

`func (o *ImportTemplateImageModel) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *ImportTemplateImageModel) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *ImportTemplateImageModel) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *ImportTemplateImageModel) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### GetCspSiteId

`func (o *ImportTemplateImageModel) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *ImportTemplateImageModel) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *ImportTemplateImageModel) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *ImportTemplateImageModel) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.

### GetAzureSubscriptionId

`func (o *ImportTemplateImageModel) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *ImportTemplateImageModel) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *ImportTemplateImageModel) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *ImportTemplateImageModel) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.

### GetRegion

`func (o *ImportTemplateImageModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ImportTemplateImageModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ImportTemplateImageModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ImportTemplateImageModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetHyperVGen

`func (o *ImportTemplateImageModel) GetHyperVGen() string`

GetHyperVGen returns the HyperVGen field if non-nil, zero value otherwise.

### GetHyperVGenOk

`func (o *ImportTemplateImageModel) GetHyperVGenOk() (*string, bool)`

GetHyperVGenOk returns a tuple with the HyperVGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperVGen

`func (o *ImportTemplateImageModel) SetHyperVGen(v string)`

SetHyperVGen sets HyperVGen field to given value.

### HasHyperVGen

`func (o *ImportTemplateImageModel) HasHyperVGen() bool`

HasHyperVGen returns a boolean if a field has been set.

### GetVtpmEnabled

`func (o *ImportTemplateImageModel) GetVtpmEnabled() bool`

GetVtpmEnabled returns the VtpmEnabled field if non-nil, zero value otherwise.

### GetVtpmEnabledOk

`func (o *ImportTemplateImageModel) GetVtpmEnabledOk() (*bool, bool)`

GetVtpmEnabledOk returns a tuple with the VtpmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpmEnabled

`func (o *ImportTemplateImageModel) SetVtpmEnabled(v bool)`

SetVtpmEnabled sets VtpmEnabled field to given value.

### HasVtpmEnabled

`func (o *ImportTemplateImageModel) HasVtpmEnabled() bool`

HasVtpmEnabled returns a boolean if a field has been set.

### GetSecureBootEnabled

`func (o *ImportTemplateImageModel) GetSecureBootEnabled() bool`

GetSecureBootEnabled returns the SecureBootEnabled field if non-nil, zero value otherwise.

### GetSecureBootEnabledOk

`func (o *ImportTemplateImageModel) GetSecureBootEnabledOk() (*bool, bool)`

GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBootEnabled

`func (o *ImportTemplateImageModel) SetSecureBootEnabled(v bool)`

SetSecureBootEnabled sets SecureBootEnabled field to given value.

### HasSecureBootEnabled

`func (o *ImportTemplateImageModel) HasSecureBootEnabled() bool`

HasSecureBootEnabled returns a boolean if a field has been set.

### GetCmekEnabled

`func (o *ImportTemplateImageModel) GetCmekEnabled() bool`

GetCmekEnabled returns the CmekEnabled field if non-nil, zero value otherwise.

### GetCmekEnabledOk

`func (o *ImportTemplateImageModel) GetCmekEnabledOk() (*bool, bool)`

GetCmekEnabledOk returns a tuple with the CmekEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmekEnabled

`func (o *ImportTemplateImageModel) SetCmekEnabled(v bool)`

SetCmekEnabled sets CmekEnabled field to given value.

### HasCmekEnabled

`func (o *ImportTemplateImageModel) HasCmekEnabled() bool`

HasCmekEnabled returns a boolean if a field has been set.

### GetCmekID

`func (o *ImportTemplateImageModel) GetCmekID() string`

GetCmekID returns the CmekID field if non-nil, zero value otherwise.

### GetCmekIDOk

`func (o *ImportTemplateImageModel) GetCmekIDOk() (*string, bool)`

GetCmekIDOk returns a tuple with the CmekID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmekID

`func (o *ImportTemplateImageModel) SetCmekID(v string)`

SetCmekID sets CmekID field to given value.

### HasCmekID

`func (o *ImportTemplateImageModel) HasCmekID() bool`

HasCmekID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


