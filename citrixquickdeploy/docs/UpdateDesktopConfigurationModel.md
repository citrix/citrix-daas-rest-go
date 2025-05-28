# UpdateDesktopConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the desktop | 
**Description** | Pointer to **string** | Description | [optional] 
**State** | Pointer to [**AppsAndDesktopsProvisionState**](AppsAndDesktopsProvisionState.md) | The publish state of the desktop | [optional] 

## Methods

### NewUpdateDesktopConfigurationModel

`func NewUpdateDesktopConfigurationModel(name string, ) *UpdateDesktopConfigurationModel`

NewUpdateDesktopConfigurationModel instantiates a new UpdateDesktopConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDesktopConfigurationModelWithDefaults

`func NewUpdateDesktopConfigurationModelWithDefaults() *UpdateDesktopConfigurationModel`

NewUpdateDesktopConfigurationModelWithDefaults instantiates a new UpdateDesktopConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDesktopConfigurationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDesktopConfigurationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDesktopConfigurationModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateDesktopConfigurationModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDesktopConfigurationModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDesktopConfigurationModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDesktopConfigurationModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *UpdateDesktopConfigurationModel) GetState() AppsAndDesktopsProvisionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateDesktopConfigurationModel) GetStateOk() (*AppsAndDesktopsProvisionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateDesktopConfigurationModel) SetState(v AppsAndDesktopsProvisionState)`

SetState sets State field to given value.

### HasState

`func (o *UpdateDesktopConfigurationModel) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


