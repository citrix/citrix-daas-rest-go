# ImageVersionSpecResourcePoolResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the resource. | [optional] 
**Name** | Pointer to **NullableString** | Name of the resource. | [optional] 
**XDPath** | Pointer to **NullableString** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**FullRelativePath** | **string** | Full path to the resources within the resource pool, including the hypervisor, relative to the root of the API. Example: &#x60;Hypervisors/{{hypervisor id}}/ResourcePools/{{resource pool id}}/Resources&#x60; | 
**Hypervisor** | [**HypervisorRefResponseModel**](HypervisorRefResponseModel.md) |  | 
**IsPrimary** | Pointer to **bool** | Indicates whether the resource pool is the primary resource pool for the image version. Resource pool used to prepare the image is the default primary resource pool. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image that are specific to the target hosting infrastructure. | [optional] 
**CustomPropertiesInString** | Pointer to **NullableString** | The properties of the image that are specific to the target hosting infrastructure in string format. | [optional] 
**Errors** | Pointer to **[]string** | The errors in this image version specification within this resource pool. | [optional] 
**ImageVersionSpecResourcePoolStatus** | Pointer to [**ImageVersionSpecResourcePoolStatus**](ImageVersionSpecResourcePoolStatus.md) |  | [optional] 

## Methods

### NewImageVersionSpecResourcePoolResponseModel

`func NewImageVersionSpecResourcePoolResponseModel(fullRelativePath string, hypervisor HypervisorRefResponseModel, ) *ImageVersionSpecResourcePoolResponseModel`

NewImageVersionSpecResourcePoolResponseModel instantiates a new ImageVersionSpecResourcePoolResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionSpecResourcePoolResponseModelWithDefaults

`func NewImageVersionSpecResourcePoolResponseModelWithDefaults() *ImageVersionSpecResourcePoolResponseModel`

NewImageVersionSpecResourcePoolResponseModelWithDefaults instantiates a new ImageVersionSpecResourcePoolResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageVersionSpecResourcePoolResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionSpecResourcePoolResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionSpecResourcePoolResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageVersionSpecResourcePoolResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ImageVersionSpecResourcePoolResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ImageVersionSpecResourcePoolResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ImageVersionSpecResourcePoolResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageVersionSpecResourcePoolResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageVersionSpecResourcePoolResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageVersionSpecResourcePoolResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ImageVersionSpecResourcePoolResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImageVersionSpecResourcePoolResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *ImageVersionSpecResourcePoolResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *ImageVersionSpecResourcePoolResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *ImageVersionSpecResourcePoolResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *ImageVersionSpecResourcePoolResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *ImageVersionSpecResourcePoolResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *ImageVersionSpecResourcePoolResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetFullRelativePath

`func (o *ImageVersionSpecResourcePoolResponseModel) GetFullRelativePath() string`

GetFullRelativePath returns the FullRelativePath field if non-nil, zero value otherwise.

### GetFullRelativePathOk

`func (o *ImageVersionSpecResourcePoolResponseModel) GetFullRelativePathOk() (*string, bool)`

GetFullRelativePathOk returns a tuple with the FullRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRelativePath

`func (o *ImageVersionSpecResourcePoolResponseModel) SetFullRelativePath(v string)`

SetFullRelativePath sets FullRelativePath field to given value.


### GetHypervisor

`func (o *ImageVersionSpecResourcePoolResponseModel) GetHypervisor() HypervisorRefResponseModel`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *ImageVersionSpecResourcePoolResponseModel) GetHypervisorOk() (*HypervisorRefResponseModel, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *ImageVersionSpecResourcePoolResponseModel) SetHypervisor(v HypervisorRefResponseModel)`

SetHypervisor sets Hypervisor field to given value.


### GetIsPrimary

`func (o *ImageVersionSpecResourcePoolResponseModel) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *ImageVersionSpecResourcePoolResponseModel) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *ImageVersionSpecResourcePoolResponseModel) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *ImageVersionSpecResourcePoolResponseModel) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ImageVersionSpecResourcePoolResponseModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ImageVersionSpecResourcePoolResponseModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ImageVersionSpecResourcePoolResponseModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ImageVersionSpecResourcePoolResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *ImageVersionSpecResourcePoolResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *ImageVersionSpecResourcePoolResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetCustomPropertiesInString

`func (o *ImageVersionSpecResourcePoolResponseModel) GetCustomPropertiesInString() string`

GetCustomPropertiesInString returns the CustomPropertiesInString field if non-nil, zero value otherwise.

### GetCustomPropertiesInStringOk

`func (o *ImageVersionSpecResourcePoolResponseModel) GetCustomPropertiesInStringOk() (*string, bool)`

GetCustomPropertiesInStringOk returns a tuple with the CustomPropertiesInString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPropertiesInString

`func (o *ImageVersionSpecResourcePoolResponseModel) SetCustomPropertiesInString(v string)`

SetCustomPropertiesInString sets CustomPropertiesInString field to given value.

### HasCustomPropertiesInString

`func (o *ImageVersionSpecResourcePoolResponseModel) HasCustomPropertiesInString() bool`

HasCustomPropertiesInString returns a boolean if a field has been set.

### SetCustomPropertiesInStringNil

`func (o *ImageVersionSpecResourcePoolResponseModel) SetCustomPropertiesInStringNil(b bool)`

 SetCustomPropertiesInStringNil sets the value for CustomPropertiesInString to be an explicit nil

### UnsetCustomPropertiesInString
`func (o *ImageVersionSpecResourcePoolResponseModel) UnsetCustomPropertiesInString()`

UnsetCustomPropertiesInString ensures that no value is present for CustomPropertiesInString, not even an explicit nil
### GetErrors

`func (o *ImageVersionSpecResourcePoolResponseModel) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ImageVersionSpecResourcePoolResponseModel) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ImageVersionSpecResourcePoolResponseModel) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ImageVersionSpecResourcePoolResponseModel) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ImageVersionSpecResourcePoolResponseModel) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ImageVersionSpecResourcePoolResponseModel) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetImageVersionSpecResourcePoolStatus

`func (o *ImageVersionSpecResourcePoolResponseModel) GetImageVersionSpecResourcePoolStatus() ImageVersionSpecResourcePoolStatus`

GetImageVersionSpecResourcePoolStatus returns the ImageVersionSpecResourcePoolStatus field if non-nil, zero value otherwise.

### GetImageVersionSpecResourcePoolStatusOk

`func (o *ImageVersionSpecResourcePoolResponseModel) GetImageVersionSpecResourcePoolStatusOk() (*ImageVersionSpecResourcePoolStatus, bool)`

GetImageVersionSpecResourcePoolStatusOk returns a tuple with the ImageVersionSpecResourcePoolStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersionSpecResourcePoolStatus

`func (o *ImageVersionSpecResourcePoolResponseModel) SetImageVersionSpecResourcePoolStatus(v ImageVersionSpecResourcePoolStatus)`

SetImageVersionSpecResourcePoolStatus sets ImageVersionSpecResourcePoolStatus field to given value.

### HasImageVersionSpecResourcePoolStatus

`func (o *ImageVersionSpecResourcePoolResponseModel) HasImageVersionSpecResourcePoolStatus() bool`

HasImageVersionSpecResourcePoolStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


