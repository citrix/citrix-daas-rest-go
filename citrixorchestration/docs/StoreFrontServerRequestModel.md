# StoreFrontServerRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of an existing StoreFront server.  If specified, this must be the only property specified in the request model. | [optional] 
**Name** | Pointer to **string** | Name of the StoreFront server. | [optional] 
**Description** | Pointer to **string** | Description of the StoreFront server. | [optional] 
**Url** | Pointer to **string** | Url of the StoreFront server. | [optional] 
**Enabled** | Pointer to **bool** | Whether the StoreFront server is enabled.  Disabled StoreFront servers will not have thier URLs added to hosted receiver. | [optional] [default to true]

## Methods

### NewStoreFrontServerRequestModel

`func NewStoreFrontServerRequestModel() *StoreFrontServerRequestModel`

NewStoreFrontServerRequestModel instantiates a new StoreFrontServerRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreFrontServerRequestModelWithDefaults

`func NewStoreFrontServerRequestModelWithDefaults() *StoreFrontServerRequestModel`

NewStoreFrontServerRequestModelWithDefaults instantiates a new StoreFrontServerRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StoreFrontServerRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoreFrontServerRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoreFrontServerRequestModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StoreFrontServerRequestModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StoreFrontServerRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreFrontServerRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreFrontServerRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoreFrontServerRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StoreFrontServerRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoreFrontServerRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoreFrontServerRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoreFrontServerRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *StoreFrontServerRequestModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StoreFrontServerRequestModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StoreFrontServerRequestModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StoreFrontServerRequestModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *StoreFrontServerRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StoreFrontServerRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StoreFrontServerRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StoreFrontServerRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


