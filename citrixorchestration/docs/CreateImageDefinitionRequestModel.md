# CreateImageDefinitionRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name associated with the image definition. | 
**Description** | Pointer to **NullableString** | The description associated with the image definition. | [optional] 
**OsType** | [**OsType**](OsType.md) |  | 
**VDASessionSupport** | [**SessionSupport**](SessionSupport.md) |  | 
**AssignedHypervisorConnection** | Pointer to [**AssignHypervisorConnectionToImageDefinitionRequestModel**](AssignHypervisorConnectionToImageDefinitionRequestModel.md) |  | [optional] 

## Methods

### NewCreateImageDefinitionRequestModel

`func NewCreateImageDefinitionRequestModel(name string, osType OsType, vDASessionSupport SessionSupport, ) *CreateImageDefinitionRequestModel`

NewCreateImageDefinitionRequestModel instantiates a new CreateImageDefinitionRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageDefinitionRequestModelWithDefaults

`func NewCreateImageDefinitionRequestModelWithDefaults() *CreateImageDefinitionRequestModel`

NewCreateImageDefinitionRequestModelWithDefaults instantiates a new CreateImageDefinitionRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateImageDefinitionRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateImageDefinitionRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateImageDefinitionRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateImageDefinitionRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateImageDefinitionRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateImageDefinitionRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateImageDefinitionRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateImageDefinitionRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateImageDefinitionRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOsType

`func (o *CreateImageDefinitionRequestModel) GetOsType() OsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CreateImageDefinitionRequestModel) GetOsTypeOk() (*OsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CreateImageDefinitionRequestModel) SetOsType(v OsType)`

SetOsType sets OsType field to given value.


### GetVDASessionSupport

`func (o *CreateImageDefinitionRequestModel) GetVDASessionSupport() SessionSupport`

GetVDASessionSupport returns the VDASessionSupport field if non-nil, zero value otherwise.

### GetVDASessionSupportOk

`func (o *CreateImageDefinitionRequestModel) GetVDASessionSupportOk() (*SessionSupport, bool)`

GetVDASessionSupportOk returns a tuple with the VDASessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDASessionSupport

`func (o *CreateImageDefinitionRequestModel) SetVDASessionSupport(v SessionSupport)`

SetVDASessionSupport sets VDASessionSupport field to given value.


### GetAssignedHypervisorConnection

`func (o *CreateImageDefinitionRequestModel) GetAssignedHypervisorConnection() AssignHypervisorConnectionToImageDefinitionRequestModel`

GetAssignedHypervisorConnection returns the AssignedHypervisorConnection field if non-nil, zero value otherwise.

### GetAssignedHypervisorConnectionOk

`func (o *CreateImageDefinitionRequestModel) GetAssignedHypervisorConnectionOk() (*AssignHypervisorConnectionToImageDefinitionRequestModel, bool)`

GetAssignedHypervisorConnectionOk returns a tuple with the AssignedHypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedHypervisorConnection

`func (o *CreateImageDefinitionRequestModel) SetAssignedHypervisorConnection(v AssignHypervisorConnectionToImageDefinitionRequestModel)`

SetAssignedHypervisorConnection sets AssignedHypervisorConnection field to given value.

### HasAssignedHypervisorConnection

`func (o *CreateImageDefinitionRequestModel) HasAssignedHypervisorConnection() bool`

HasAssignedHypervisorConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


