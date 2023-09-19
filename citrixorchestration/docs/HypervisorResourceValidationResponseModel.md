# HypervisorResourceValidationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceRef** | Pointer to [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | [optional] 
**Reports** | Pointer to [**[]ResourceValidationReportModel**](ResourceValidationReportModel.md) | The validation reports. | [optional] 

## Methods

### NewHypervisorResourceValidationResponseModel

`func NewHypervisorResourceValidationResponseModel() *HypervisorResourceValidationResponseModel`

NewHypervisorResourceValidationResponseModel instantiates a new HypervisorResourceValidationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourceValidationResponseModelWithDefaults

`func NewHypervisorResourceValidationResponseModelWithDefaults() *HypervisorResourceValidationResponseModel`

NewHypervisorResourceValidationResponseModelWithDefaults instantiates a new HypervisorResourceValidationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceRef

`func (o *HypervisorResourceValidationResponseModel) GetResourceRef() HypervisorResourceRefResponseModel`

GetResourceRef returns the ResourceRef field if non-nil, zero value otherwise.

### GetResourceRefOk

`func (o *HypervisorResourceValidationResponseModel) GetResourceRefOk() (*HypervisorResourceRefResponseModel, bool)`

GetResourceRefOk returns a tuple with the ResourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRef

`func (o *HypervisorResourceValidationResponseModel) SetResourceRef(v HypervisorResourceRefResponseModel)`

SetResourceRef sets ResourceRef field to given value.

### HasResourceRef

`func (o *HypervisorResourceValidationResponseModel) HasResourceRef() bool`

HasResourceRef returns a boolean if a field has been set.

### GetReports

`func (o *HypervisorResourceValidationResponseModel) GetReports() []ResourceValidationReportModel`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *HypervisorResourceValidationResponseModel) GetReportsOk() (*[]ResourceValidationReportModel, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *HypervisorResourceValidationResponseModel) SetReports(v []ResourceValidationReportModel)`

SetReports sets Reports field to given value.

### HasReports

`func (o *HypervisorResourceValidationResponseModel) HasReports() bool`

HasReports returns a boolean if a field has been set.

### SetReportsNil

`func (o *HypervisorResourceValidationResponseModel) SetReportsNil(b bool)`

 SetReportsNil sets the value for Reports to be an explicit nil

### UnsetReports
`func (o *HypervisorResourceValidationResponseModel) UnsetReports()`

UnsetReports ensures that no value is present for Reports, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


