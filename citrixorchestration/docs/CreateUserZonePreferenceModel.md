# CreateUserZonePreferenceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | **string** | Zone Id. | 
**Name** | **string** | Name of the user or user group. Could be SamName, FullName, SID or UPN | 

## Methods

### NewCreateUserZonePreferenceModel

`func NewCreateUserZonePreferenceModel(zoneId string, name string, ) *CreateUserZonePreferenceModel`

NewCreateUserZonePreferenceModel instantiates a new CreateUserZonePreferenceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserZonePreferenceModelWithDefaults

`func NewCreateUserZonePreferenceModelWithDefaults() *CreateUserZonePreferenceModel`

NewCreateUserZonePreferenceModelWithDefaults instantiates a new CreateUserZonePreferenceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *CreateUserZonePreferenceModel) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *CreateUserZonePreferenceModel) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *CreateUserZonePreferenceModel) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetName

`func (o *CreateUserZonePreferenceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserZonePreferenceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserZonePreferenceModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


