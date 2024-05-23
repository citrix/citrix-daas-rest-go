# ProvisioningSchemeImageVersionHistoryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageVersion** | [**ImageVersionRefResponseModel**](ImageVersionRefResponseModel.md) |  | 
**Date** | Pointer to **NullableString** | The date and time when the snapshot was used in the provisioning scheme. | [optional] 

## Methods

### NewProvisioningSchemeImageVersionHistoryResponseModel

`func NewProvisioningSchemeImageVersionHistoryResponseModel(imageVersion ImageVersionRefResponseModel, ) *ProvisioningSchemeImageVersionHistoryResponseModel`

NewProvisioningSchemeImageVersionHistoryResponseModel instantiates a new ProvisioningSchemeImageVersionHistoryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeImageVersionHistoryResponseModelWithDefaults

`func NewProvisioningSchemeImageVersionHistoryResponseModelWithDefaults() *ProvisioningSchemeImageVersionHistoryResponseModel`

NewProvisioningSchemeImageVersionHistoryResponseModelWithDefaults instantiates a new ProvisioningSchemeImageVersionHistoryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageVersion

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetImageVersion() ImageVersionRefResponseModel`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetImageVersionOk() (*ImageVersionRefResponseModel, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetImageVersion(v ImageVersionRefResponseModel)`

SetImageVersion sets ImageVersion field to given value.


### GetDate

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


