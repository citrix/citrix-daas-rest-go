# AssignHypervisorConnectionToImageDefinitionRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorConnection** | **string** | Specifies the hypervisor connection to assign to the image definition. | 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image definition that are specific to the target hypervisor connection. | [optional] 

## Methods

### NewAssignHypervisorConnectionToImageDefinitionRequestModel

`func NewAssignHypervisorConnectionToImageDefinitionRequestModel(hypervisorConnection string, ) *AssignHypervisorConnectionToImageDefinitionRequestModel`

NewAssignHypervisorConnectionToImageDefinitionRequestModel instantiates a new AssignHypervisorConnectionToImageDefinitionRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignHypervisorConnectionToImageDefinitionRequestModelWithDefaults

`func NewAssignHypervisorConnectionToImageDefinitionRequestModelWithDefaults() *AssignHypervisorConnectionToImageDefinitionRequestModel`

NewAssignHypervisorConnectionToImageDefinitionRequestModelWithDefaults instantiates a new AssignHypervisorConnectionToImageDefinitionRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHypervisorConnection

`func (o *AssignHypervisorConnectionToImageDefinitionRequestModel) GetHypervisorConnection() string`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *AssignHypervisorConnectionToImageDefinitionRequestModel) GetHypervisorConnectionOk() (*string, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *AssignHypervisorConnectionToImageDefinitionRequestModel) SetHypervisorConnection(v string)`

SetHypervisorConnection sets HypervisorConnection field to given value.


### GetCustomProperties

`func (o *AssignHypervisorConnectionToImageDefinitionRequestModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *AssignHypervisorConnectionToImageDefinitionRequestModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *AssignHypervisorConnectionToImageDefinitionRequestModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *AssignHypervisorConnectionToImageDefinitionRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *AssignHypervisorConnectionToImageDefinitionRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *AssignHypervisorConnectionToImageDefinitionRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


