# ImageVersionSpecResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Id of the image version specification. | 
**Context** | Pointer to [**ImageVersionSpecContextResponseModel**](ImageVersionSpecContextResponseModel.md) |  | [optional] 
**CreationTime** | Pointer to **NullableString** | The create time of the image version specification. | [optional] 
**Errors** | Pointer to **[]string** | The errors in this image version specification | [optional] 
**ImageRuntimeEnvironment** | Pointer to [**ImageRuntimeEnvironmentResponseModel**](ImageRuntimeEnvironmentResponseModel.md) |  | [optional] 
**ImageVersionSpecStatus** | Pointer to [**ImageVersionSpecStatus**](ImageVersionSpecStatus.md) |  | [optional] 
**MasterImage** | Pointer to [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | [optional] 
**PreparationType** | [**PreparationType**](PreparationType.md) |  | 
**ProvisioningSchemeCount** | Pointer to **int32** | Number of provisioning scheme created from this image version specification. | [optional] 
**ResourcePool** | [**HypervisorResourcePoolRefResponseModel**](HypervisorResourcePoolRefResponseModel.md) |  | 
**SourceImageVersionSpec** | Pointer to [**ImageVersionSpecRefResponseModel**](ImageVersionSpecRefResponseModel.md) |  | [optional] 
**DiskSize** | Pointer to **int32** | The disk size | [optional] 
**Warnings** | Pointer to [**[]ImageVersionSpecWarningResponseModel**](ImageVersionSpecWarningResponseModel.md) | List of warnings that are currently active on the image version specification, if any.  If there are no warnings this will not be specified. | [optional] 
**Zone** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 

## Methods

### NewImageVersionSpecResponseModel

`func NewImageVersionSpecResponseModel(id string, preparationType PreparationType, resourcePool HypervisorResourcePoolRefResponseModel, ) *ImageVersionSpecResponseModel`

NewImageVersionSpecResponseModel instantiates a new ImageVersionSpecResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionSpecResponseModelWithDefaults

`func NewImageVersionSpecResponseModelWithDefaults() *ImageVersionSpecResponseModel`

NewImageVersionSpecResponseModelWithDefaults instantiates a new ImageVersionSpecResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageVersionSpecResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionSpecResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionSpecResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetContext

`func (o *ImageVersionSpecResponseModel) GetContext() ImageVersionSpecContextResponseModel`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ImageVersionSpecResponseModel) GetContextOk() (*ImageVersionSpecContextResponseModel, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ImageVersionSpecResponseModel) SetContext(v ImageVersionSpecContextResponseModel)`

SetContext sets Context field to given value.

### HasContext

`func (o *ImageVersionSpecResponseModel) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreationTime

`func (o *ImageVersionSpecResponseModel) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ImageVersionSpecResponseModel) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ImageVersionSpecResponseModel) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ImageVersionSpecResponseModel) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### SetCreationTimeNil

`func (o *ImageVersionSpecResponseModel) SetCreationTimeNil(b bool)`

 SetCreationTimeNil sets the value for CreationTime to be an explicit nil

### UnsetCreationTime
`func (o *ImageVersionSpecResponseModel) UnsetCreationTime()`

UnsetCreationTime ensures that no value is present for CreationTime, not even an explicit nil
### GetErrors

`func (o *ImageVersionSpecResponseModel) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ImageVersionSpecResponseModel) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ImageVersionSpecResponseModel) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ImageVersionSpecResponseModel) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ImageVersionSpecResponseModel) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ImageVersionSpecResponseModel) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetImageRuntimeEnvironment

`func (o *ImageVersionSpecResponseModel) GetImageRuntimeEnvironment() ImageRuntimeEnvironmentResponseModel`

GetImageRuntimeEnvironment returns the ImageRuntimeEnvironment field if non-nil, zero value otherwise.

### GetImageRuntimeEnvironmentOk

`func (o *ImageVersionSpecResponseModel) GetImageRuntimeEnvironmentOk() (*ImageRuntimeEnvironmentResponseModel, bool)`

GetImageRuntimeEnvironmentOk returns a tuple with the ImageRuntimeEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRuntimeEnvironment

`func (o *ImageVersionSpecResponseModel) SetImageRuntimeEnvironment(v ImageRuntimeEnvironmentResponseModel)`

SetImageRuntimeEnvironment sets ImageRuntimeEnvironment field to given value.

### HasImageRuntimeEnvironment

`func (o *ImageVersionSpecResponseModel) HasImageRuntimeEnvironment() bool`

HasImageRuntimeEnvironment returns a boolean if a field has been set.

### GetImageVersionSpecStatus

`func (o *ImageVersionSpecResponseModel) GetImageVersionSpecStatus() ImageVersionSpecStatus`

GetImageVersionSpecStatus returns the ImageVersionSpecStatus field if non-nil, zero value otherwise.

### GetImageVersionSpecStatusOk

`func (o *ImageVersionSpecResponseModel) GetImageVersionSpecStatusOk() (*ImageVersionSpecStatus, bool)`

GetImageVersionSpecStatusOk returns a tuple with the ImageVersionSpecStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersionSpecStatus

`func (o *ImageVersionSpecResponseModel) SetImageVersionSpecStatus(v ImageVersionSpecStatus)`

SetImageVersionSpecStatus sets ImageVersionSpecStatus field to given value.

### HasImageVersionSpecStatus

`func (o *ImageVersionSpecResponseModel) HasImageVersionSpecStatus() bool`

HasImageVersionSpecStatus returns a boolean if a field has been set.

### GetMasterImage

`func (o *ImageVersionSpecResponseModel) GetMasterImage() HypervisorResourceRefResponseModel`

GetMasterImage returns the MasterImage field if non-nil, zero value otherwise.

### GetMasterImageOk

`func (o *ImageVersionSpecResponseModel) GetMasterImageOk() (*HypervisorResourceRefResponseModel, bool)`

GetMasterImageOk returns a tuple with the MasterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImage

`func (o *ImageVersionSpecResponseModel) SetMasterImage(v HypervisorResourceRefResponseModel)`

SetMasterImage sets MasterImage field to given value.

### HasMasterImage

`func (o *ImageVersionSpecResponseModel) HasMasterImage() bool`

HasMasterImage returns a boolean if a field has been set.

### GetPreparationType

`func (o *ImageVersionSpecResponseModel) GetPreparationType() PreparationType`

GetPreparationType returns the PreparationType field if non-nil, zero value otherwise.

### GetPreparationTypeOk

`func (o *ImageVersionSpecResponseModel) GetPreparationTypeOk() (*PreparationType, bool)`

GetPreparationTypeOk returns a tuple with the PreparationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationType

`func (o *ImageVersionSpecResponseModel) SetPreparationType(v PreparationType)`

SetPreparationType sets PreparationType field to given value.


### GetProvisioningSchemeCount

`func (o *ImageVersionSpecResponseModel) GetProvisioningSchemeCount() int32`

GetProvisioningSchemeCount returns the ProvisioningSchemeCount field if non-nil, zero value otherwise.

### GetProvisioningSchemeCountOk

`func (o *ImageVersionSpecResponseModel) GetProvisioningSchemeCountOk() (*int32, bool)`

GetProvisioningSchemeCountOk returns a tuple with the ProvisioningSchemeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeCount

`func (o *ImageVersionSpecResponseModel) SetProvisioningSchemeCount(v int32)`

SetProvisioningSchemeCount sets ProvisioningSchemeCount field to given value.

### HasProvisioningSchemeCount

`func (o *ImageVersionSpecResponseModel) HasProvisioningSchemeCount() bool`

HasProvisioningSchemeCount returns a boolean if a field has been set.

### GetResourcePool

`func (o *ImageVersionSpecResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ImageVersionSpecResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ImageVersionSpecResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel)`

SetResourcePool sets ResourcePool field to given value.


### GetSourceImageVersionSpec

`func (o *ImageVersionSpecResponseModel) GetSourceImageVersionSpec() ImageVersionSpecRefResponseModel`

GetSourceImageVersionSpec returns the SourceImageVersionSpec field if non-nil, zero value otherwise.

### GetSourceImageVersionSpecOk

`func (o *ImageVersionSpecResponseModel) GetSourceImageVersionSpecOk() (*ImageVersionSpecRefResponseModel, bool)`

GetSourceImageVersionSpecOk returns a tuple with the SourceImageVersionSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImageVersionSpec

`func (o *ImageVersionSpecResponseModel) SetSourceImageVersionSpec(v ImageVersionSpecRefResponseModel)`

SetSourceImageVersionSpec sets SourceImageVersionSpec field to given value.

### HasSourceImageVersionSpec

`func (o *ImageVersionSpecResponseModel) HasSourceImageVersionSpec() bool`

HasSourceImageVersionSpec returns a boolean if a field has been set.

### GetDiskSize

`func (o *ImageVersionSpecResponseModel) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ImageVersionSpecResponseModel) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ImageVersionSpecResponseModel) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *ImageVersionSpecResponseModel) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetWarnings

`func (o *ImageVersionSpecResponseModel) GetWarnings() []ImageVersionSpecWarningResponseModel`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ImageVersionSpecResponseModel) GetWarningsOk() (*[]ImageVersionSpecWarningResponseModel, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ImageVersionSpecResponseModel) SetWarnings(v []ImageVersionSpecWarningResponseModel)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ImageVersionSpecResponseModel) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *ImageVersionSpecResponseModel) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *ImageVersionSpecResponseModel) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetZone

`func (o *ImageVersionSpecResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ImageVersionSpecResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ImageVersionSpecResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ImageVersionSpecResponseModel) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


