# ImageVersionProvisioningSchemeRefResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the object. Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | [optional] 
**Uid** | Pointer to **NullableInt32** | DEPRECATED. Use Id. | [optional] 
**Name** | Pointer to **NullableString** | Name of the object. | [optional] 

## Methods

### NewImageVersionProvisioningSchemeRefResponseModel

`func NewImageVersionProvisioningSchemeRefResponseModel() *ImageVersionProvisioningSchemeRefResponseModel`

NewImageVersionProvisioningSchemeRefResponseModel instantiates a new ImageVersionProvisioningSchemeRefResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionProvisioningSchemeRefResponseModelWithDefaults

`func NewImageVersionProvisioningSchemeRefResponseModelWithDefaults() *ImageVersionProvisioningSchemeRefResponseModel`

NewImageVersionProvisioningSchemeRefResponseModelWithDefaults instantiates a new ImageVersionProvisioningSchemeRefResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageVersionProvisioningSchemeRefResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionProvisioningSchemeRefResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionProvisioningSchemeRefResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageVersionProvisioningSchemeRefResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ImageVersionProvisioningSchemeRefResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ImageVersionProvisioningSchemeRefResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *ImageVersionProvisioningSchemeRefResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ImageVersionProvisioningSchemeRefResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ImageVersionProvisioningSchemeRefResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ImageVersionProvisioningSchemeRefResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ImageVersionProvisioningSchemeRefResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ImageVersionProvisioningSchemeRefResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetName

`func (o *ImageVersionProvisioningSchemeRefResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageVersionProvisioningSchemeRefResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageVersionProvisioningSchemeRefResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageVersionProvisioningSchemeRefResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ImageVersionProvisioningSchemeRefResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImageVersionProvisioningSchemeRefResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


