# UserZonePreferenceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | The full name of the user or user group. | [optional] 
**Name** | Pointer to **NullableString** | Name of the user or user group. | [optional] 
**Sid** | Pointer to **NullableString** | Sid of the user or user group. | [optional] 
**Upn** | Pointer to **NullableString** | UPN of the user or user group. | [optional] 
**Zone** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 

## Methods

### NewUserZonePreferenceResponseModel

`func NewUserZonePreferenceResponseModel() *UserZonePreferenceResponseModel`

NewUserZonePreferenceResponseModel instantiates a new UserZonePreferenceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserZonePreferenceResponseModelWithDefaults

`func NewUserZonePreferenceResponseModelWithDefaults() *UserZonePreferenceResponseModel`

NewUserZonePreferenceResponseModelWithDefaults instantiates a new UserZonePreferenceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *UserZonePreferenceResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UserZonePreferenceResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UserZonePreferenceResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UserZonePreferenceResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *UserZonePreferenceResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *UserZonePreferenceResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetName

`func (o *UserZonePreferenceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserZonePreferenceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserZonePreferenceResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserZonePreferenceResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserZonePreferenceResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserZonePreferenceResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSid

`func (o *UserZonePreferenceResponseModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *UserZonePreferenceResponseModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *UserZonePreferenceResponseModel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *UserZonePreferenceResponseModel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *UserZonePreferenceResponseModel) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *UserZonePreferenceResponseModel) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUpn

`func (o *UserZonePreferenceResponseModel) GetUpn() string`

GetUpn returns the Upn field if non-nil, zero value otherwise.

### GetUpnOk

`func (o *UserZonePreferenceResponseModel) GetUpnOk() (*string, bool)`

GetUpnOk returns a tuple with the Upn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpn

`func (o *UserZonePreferenceResponseModel) SetUpn(v string)`

SetUpn sets Upn field to given value.

### HasUpn

`func (o *UserZonePreferenceResponseModel) HasUpn() bool`

HasUpn returns a boolean if a field has been set.

### SetUpnNil

`func (o *UserZonePreferenceResponseModel) SetUpnNil(b bool)`

 SetUpnNil sets the value for Upn to be an explicit nil

### UnsetUpn
`func (o *UserZonePreferenceResponseModel) UnsetUpn()`

UnsetUpn ensures that no value is present for Upn, not even an explicit nil
### GetZone

`func (o *UserZonePreferenceResponseModel) GetZone() RefResponseModel`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UserZonePreferenceResponseModel) GetZoneOk() (*RefResponseModel, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UserZonePreferenceResponseModel) SetZone(v RefResponseModel)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UserZonePreferenceResponseModel) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


