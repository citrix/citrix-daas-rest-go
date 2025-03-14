# CreateAutoscalePluginTemplateRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Autoscale plugin template name. | 
**Dates** | Pointer to **[]string** | Date range for the autoscale holiday plugin template. | [optional] 

## Methods

### NewCreateAutoscalePluginTemplateRequestModel

`func NewCreateAutoscalePluginTemplateRequestModel(name string, ) *CreateAutoscalePluginTemplateRequestModel`

NewCreateAutoscalePluginTemplateRequestModel instantiates a new CreateAutoscalePluginTemplateRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoscalePluginTemplateRequestModelWithDefaults

`func NewCreateAutoscalePluginTemplateRequestModelWithDefaults() *CreateAutoscalePluginTemplateRequestModel`

NewCreateAutoscalePluginTemplateRequestModelWithDefaults instantiates a new CreateAutoscalePluginTemplateRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAutoscalePluginTemplateRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAutoscalePluginTemplateRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAutoscalePluginTemplateRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDates

`func (o *CreateAutoscalePluginTemplateRequestModel) GetDates() []string`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *CreateAutoscalePluginTemplateRequestModel) GetDatesOk() (*[]string, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *CreateAutoscalePluginTemplateRequestModel) SetDates(v []string)`

SetDates sets Dates field to given value.

### HasDates

`func (o *CreateAutoscalePluginTemplateRequestModel) HasDates() bool`

HasDates returns a boolean if a field has been set.

### SetDatesNil

`func (o *CreateAutoscalePluginTemplateRequestModel) SetDatesNil(b bool)`

 SetDatesNil sets the value for Dates to be an explicit nil

### UnsetDates
`func (o *CreateAutoscalePluginTemplateRequestModel) UnsetDates()`

UnsetDates ensures that no value is present for Dates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


