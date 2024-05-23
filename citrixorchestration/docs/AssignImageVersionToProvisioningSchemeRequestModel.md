# AssignImageVersionToProvisioningSchemeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDefinition** | **string** | The image definition of the image version to assign. May be an Id or name. | 
**ImageVersion** | **string** | The image version to assign. May be an Id or number. | 

## Methods

### NewAssignImageVersionToProvisioningSchemeRequestModel

`func NewAssignImageVersionToProvisioningSchemeRequestModel(imageDefinition string, imageVersion string, ) *AssignImageVersionToProvisioningSchemeRequestModel`

NewAssignImageVersionToProvisioningSchemeRequestModel instantiates a new AssignImageVersionToProvisioningSchemeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignImageVersionToProvisioningSchemeRequestModelWithDefaults

`func NewAssignImageVersionToProvisioningSchemeRequestModelWithDefaults() *AssignImageVersionToProvisioningSchemeRequestModel`

NewAssignImageVersionToProvisioningSchemeRequestModelWithDefaults instantiates a new AssignImageVersionToProvisioningSchemeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDefinition

`func (o *AssignImageVersionToProvisioningSchemeRequestModel) GetImageDefinition() string`

GetImageDefinition returns the ImageDefinition field if non-nil, zero value otherwise.

### GetImageDefinitionOk

`func (o *AssignImageVersionToProvisioningSchemeRequestModel) GetImageDefinitionOk() (*string, bool)`

GetImageDefinitionOk returns a tuple with the ImageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinition

`func (o *AssignImageVersionToProvisioningSchemeRequestModel) SetImageDefinition(v string)`

SetImageDefinition sets ImageDefinition field to given value.


### GetImageVersion

`func (o *AssignImageVersionToProvisioningSchemeRequestModel) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *AssignImageVersionToProvisioningSchemeRequestModel) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *AssignImageVersionToProvisioningSchemeRequestModel) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


