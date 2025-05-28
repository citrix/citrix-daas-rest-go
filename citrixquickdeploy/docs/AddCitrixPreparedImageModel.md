# AddCitrixPreparedImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the citrix prepared image | 
**Description** | Pointer to **string** | Description of the image | [optional] 
**AzureSubscriptionId** | **string** | ID of the Azure Subscription containing the image | 
**AzureResourceGroup** | **string** | Name of the resource group where the vhd&#39;s storage account is located | 
**SessionSupport** | [**TemplateImageSessionSupport**](TemplateImageSessionSupport.md) | The type of sessions that are supported by the image&#39;s OS | 
**StartMenuAppsXmlUri** | **string** | URI to the xml file containing the start menu&#39;s applications | 
**VhdLocations** | Pointer to [**[]CitrixPreparedImageVhdLocation**](CitrixPreparedImageVhdLocation.md) | List of locations where the VHDs are stored | [optional] 
**OsPlatform** | Pointer to [**SupportedOperatingSystemType**](SupportedOperatingSystemType.md) | Type of operating system that will be imported | [optional] 
**HyperVGen** | Pointer to **string** | Type of operating system that will be imported | [optional] 
**VtpmEnabled** | Pointer to **bool** | Is trusted launch enabled for V2 gen | [optional] 

## Methods

### NewAddCitrixPreparedImageModel

`func NewAddCitrixPreparedImageModel(name string, azureSubscriptionId string, azureResourceGroup string, sessionSupport TemplateImageSessionSupport, startMenuAppsXmlUri string, ) *AddCitrixPreparedImageModel`

NewAddCitrixPreparedImageModel instantiates a new AddCitrixPreparedImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCitrixPreparedImageModelWithDefaults

`func NewAddCitrixPreparedImageModelWithDefaults() *AddCitrixPreparedImageModel`

NewAddCitrixPreparedImageModelWithDefaults instantiates a new AddCitrixPreparedImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddCitrixPreparedImageModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCitrixPreparedImageModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCitrixPreparedImageModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddCitrixPreparedImageModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCitrixPreparedImageModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCitrixPreparedImageModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCitrixPreparedImageModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAzureSubscriptionId

`func (o *AddCitrixPreparedImageModel) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *AddCitrixPreparedImageModel) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *AddCitrixPreparedImageModel) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### GetAzureResourceGroup

`func (o *AddCitrixPreparedImageModel) GetAzureResourceGroup() string`

GetAzureResourceGroup returns the AzureResourceGroup field if non-nil, zero value otherwise.

### GetAzureResourceGroupOk

`func (o *AddCitrixPreparedImageModel) GetAzureResourceGroupOk() (*string, bool)`

GetAzureResourceGroupOk returns a tuple with the AzureResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroup

`func (o *AddCitrixPreparedImageModel) SetAzureResourceGroup(v string)`

SetAzureResourceGroup sets AzureResourceGroup field to given value.


### GetSessionSupport

`func (o *AddCitrixPreparedImageModel) GetSessionSupport() TemplateImageSessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *AddCitrixPreparedImageModel) GetSessionSupportOk() (*TemplateImageSessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *AddCitrixPreparedImageModel) SetSessionSupport(v TemplateImageSessionSupport)`

SetSessionSupport sets SessionSupport field to given value.


### GetStartMenuAppsXmlUri

`func (o *AddCitrixPreparedImageModel) GetStartMenuAppsXmlUri() string`

GetStartMenuAppsXmlUri returns the StartMenuAppsXmlUri field if non-nil, zero value otherwise.

### GetStartMenuAppsXmlUriOk

`func (o *AddCitrixPreparedImageModel) GetStartMenuAppsXmlUriOk() (*string, bool)`

GetStartMenuAppsXmlUriOk returns a tuple with the StartMenuAppsXmlUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMenuAppsXmlUri

`func (o *AddCitrixPreparedImageModel) SetStartMenuAppsXmlUri(v string)`

SetStartMenuAppsXmlUri sets StartMenuAppsXmlUri field to given value.


### GetVhdLocations

`func (o *AddCitrixPreparedImageModel) GetVhdLocations() []CitrixPreparedImageVhdLocation`

GetVhdLocations returns the VhdLocations field if non-nil, zero value otherwise.

### GetVhdLocationsOk

`func (o *AddCitrixPreparedImageModel) GetVhdLocationsOk() (*[]CitrixPreparedImageVhdLocation, bool)`

GetVhdLocationsOk returns a tuple with the VhdLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhdLocations

`func (o *AddCitrixPreparedImageModel) SetVhdLocations(v []CitrixPreparedImageVhdLocation)`

SetVhdLocations sets VhdLocations field to given value.

### HasVhdLocations

`func (o *AddCitrixPreparedImageModel) HasVhdLocations() bool`

HasVhdLocations returns a boolean if a field has been set.

### GetOsPlatform

`func (o *AddCitrixPreparedImageModel) GetOsPlatform() SupportedOperatingSystemType`

GetOsPlatform returns the OsPlatform field if non-nil, zero value otherwise.

### GetOsPlatformOk

`func (o *AddCitrixPreparedImageModel) GetOsPlatformOk() (*SupportedOperatingSystemType, bool)`

GetOsPlatformOk returns a tuple with the OsPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPlatform

`func (o *AddCitrixPreparedImageModel) SetOsPlatform(v SupportedOperatingSystemType)`

SetOsPlatform sets OsPlatform field to given value.

### HasOsPlatform

`func (o *AddCitrixPreparedImageModel) HasOsPlatform() bool`

HasOsPlatform returns a boolean if a field has been set.

### GetHyperVGen

`func (o *AddCitrixPreparedImageModel) GetHyperVGen() string`

GetHyperVGen returns the HyperVGen field if non-nil, zero value otherwise.

### GetHyperVGenOk

`func (o *AddCitrixPreparedImageModel) GetHyperVGenOk() (*string, bool)`

GetHyperVGenOk returns a tuple with the HyperVGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperVGen

`func (o *AddCitrixPreparedImageModel) SetHyperVGen(v string)`

SetHyperVGen sets HyperVGen field to given value.

### HasHyperVGen

`func (o *AddCitrixPreparedImageModel) HasHyperVGen() bool`

HasHyperVGen returns a boolean if a field has been set.

### GetVtpmEnabled

`func (o *AddCitrixPreparedImageModel) GetVtpmEnabled() bool`

GetVtpmEnabled returns the VtpmEnabled field if non-nil, zero value otherwise.

### GetVtpmEnabledOk

`func (o *AddCitrixPreparedImageModel) GetVtpmEnabledOk() (*bool, bool)`

GetVtpmEnabledOk returns a tuple with the VtpmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpmEnabled

`func (o *AddCitrixPreparedImageModel) SetVtpmEnabled(v bool)`

SetVtpmEnabled sets VtpmEnabled field to given value.

### HasVtpmEnabled

`func (o *AddCitrixPreparedImageModel) HasVtpmEnabled() bool`

HasVtpmEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


