# AutoscalePluginTemplateResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Autoscale plugin template name. | [optional] 
**Type** | Pointer to [**AutoscalePluginType**](AutoscalePluginType.md) |  | [optional] 
**Dates** | Pointer to **[]string** | Date range for the autoscale holiday plugin template. | [optional] 

## Methods

### NewAutoscalePluginTemplateResponseModel

`func NewAutoscalePluginTemplateResponseModel() *AutoscalePluginTemplateResponseModel`

NewAutoscalePluginTemplateResponseModel instantiates a new AutoscalePluginTemplateResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalePluginTemplateResponseModelWithDefaults

`func NewAutoscalePluginTemplateResponseModelWithDefaults() *AutoscalePluginTemplateResponseModel`

NewAutoscalePluginTemplateResponseModelWithDefaults instantiates a new AutoscalePluginTemplateResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutoscalePluginTemplateResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoscalePluginTemplateResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoscalePluginTemplateResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoscalePluginTemplateResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoscalePluginTemplateResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoscalePluginTemplateResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *AutoscalePluginTemplateResponseModel) GetType() AutoscalePluginType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoscalePluginTemplateResponseModel) GetTypeOk() (*AutoscalePluginType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoscalePluginTemplateResponseModel) SetType(v AutoscalePluginType)`

SetType sets Type field to given value.

### HasType

`func (o *AutoscalePluginTemplateResponseModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDates

`func (o *AutoscalePluginTemplateResponseModel) GetDates() []string`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *AutoscalePluginTemplateResponseModel) GetDatesOk() (*[]string, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *AutoscalePluginTemplateResponseModel) SetDates(v []string)`

SetDates sets Dates field to given value.

### HasDates

`func (o *AutoscalePluginTemplateResponseModel) HasDates() bool`

HasDates returns a boolean if a field has been set.

### SetDatesNil

`func (o *AutoscalePluginTemplateResponseModel) SetDatesNil(b bool)`

 SetDatesNil sets the value for Dates to be an explicit nil

### UnsetDates
`func (o *AutoscalePluginTemplateResponseModel) UnsetDates()`

UnsetDates ensures that no value is present for Dates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


