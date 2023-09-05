# IdentityCreateMachineRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forest** | **string** | Forest in which to create the machine identity. | 
**Domain** | Pointer to **string** | Domain in which to create the machine identity. | [optional] 
**Name** | **string** | The machine identity to create. | 
**ParentOU** | Pointer to **string** | The OU where the machine identity should be created. | [optional] 
**AccountPassword** | **string** | The machine account password, in the format specified by AccountPasswordFormat. | 
**AccountPasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**ResourceLocationId** | Pointer to **string** | Resource location within which the machine account should be created. | [optional] 

## Methods

### NewIdentityCreateMachineRequestModel

`func NewIdentityCreateMachineRequestModel(forest string, name string, accountPassword string, ) *IdentityCreateMachineRequestModel`

NewIdentityCreateMachineRequestModel instantiates a new IdentityCreateMachineRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityCreateMachineRequestModelWithDefaults

`func NewIdentityCreateMachineRequestModelWithDefaults() *IdentityCreateMachineRequestModel`

NewIdentityCreateMachineRequestModelWithDefaults instantiates a new IdentityCreateMachineRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForest

`func (o *IdentityCreateMachineRequestModel) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *IdentityCreateMachineRequestModel) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *IdentityCreateMachineRequestModel) SetForest(v string)`

SetForest sets Forest field to given value.


### GetDomain

`func (o *IdentityCreateMachineRequestModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentityCreateMachineRequestModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentityCreateMachineRequestModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdentityCreateMachineRequestModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *IdentityCreateMachineRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityCreateMachineRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityCreateMachineRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetParentOU

`func (o *IdentityCreateMachineRequestModel) GetParentOU() string`

GetParentOU returns the ParentOU field if non-nil, zero value otherwise.

### GetParentOUOk

`func (o *IdentityCreateMachineRequestModel) GetParentOUOk() (*string, bool)`

GetParentOUOk returns a tuple with the ParentOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOU

`func (o *IdentityCreateMachineRequestModel) SetParentOU(v string)`

SetParentOU sets ParentOU field to given value.

### HasParentOU

`func (o *IdentityCreateMachineRequestModel) HasParentOU() bool`

HasParentOU returns a boolean if a field has been set.

### GetAccountPassword

`func (o *IdentityCreateMachineRequestModel) GetAccountPassword() string`

GetAccountPassword returns the AccountPassword field if non-nil, zero value otherwise.

### GetAccountPasswordOk

`func (o *IdentityCreateMachineRequestModel) GetAccountPasswordOk() (*string, bool)`

GetAccountPasswordOk returns a tuple with the AccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPassword

`func (o *IdentityCreateMachineRequestModel) SetAccountPassword(v string)`

SetAccountPassword sets AccountPassword field to given value.


### GetAccountPasswordFormat

`func (o *IdentityCreateMachineRequestModel) GetAccountPasswordFormat() IdentityPasswordFormat`

GetAccountPasswordFormat returns the AccountPasswordFormat field if non-nil, zero value otherwise.

### GetAccountPasswordFormatOk

`func (o *IdentityCreateMachineRequestModel) GetAccountPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetAccountPasswordFormatOk returns a tuple with the AccountPasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPasswordFormat

`func (o *IdentityCreateMachineRequestModel) SetAccountPasswordFormat(v IdentityPasswordFormat)`

SetAccountPasswordFormat sets AccountPasswordFormat field to given value.

### HasAccountPasswordFormat

`func (o *IdentityCreateMachineRequestModel) HasAccountPasswordFormat() bool`

HasAccountPasswordFormat returns a boolean if a field has been set.

### GetResourceLocationId

`func (o *IdentityCreateMachineRequestModel) GetResourceLocationId() string`

GetResourceLocationId returns the ResourceLocationId field if non-nil, zero value otherwise.

### GetResourceLocationIdOk

`func (o *IdentityCreateMachineRequestModel) GetResourceLocationIdOk() (*string, bool)`

GetResourceLocationIdOk returns a tuple with the ResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocationId

`func (o *IdentityCreateMachineRequestModel) SetResourceLocationId(v string)`

SetResourceLocationId sets ResourceLocationId field to given value.

### HasResourceLocationId

`func (o *IdentityCreateMachineRequestModel) HasResourceLocationId() bool`

HasResourceLocationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


