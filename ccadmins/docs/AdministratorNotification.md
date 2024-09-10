# AdministratorNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AdministratorNotificationType**](AdministratorNotificationType.md) |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewAdministratorNotification

`func NewAdministratorNotification() *AdministratorNotification`

NewAdministratorNotification instantiates a new AdministratorNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorNotificationWithDefaults

`func NewAdministratorNotificationWithDefaults() *AdministratorNotification`

NewAdministratorNotificationWithDefaults instantiates a new AdministratorNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdministratorNotification) GetType() AdministratorNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdministratorNotification) GetTypeOk() (*AdministratorNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdministratorNotification) SetType(v AdministratorNotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *AdministratorNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *AdministratorNotification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdministratorNotification) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdministratorNotification) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AdministratorNotification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *AdministratorNotification) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *AdministratorNotification) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


