# AdministratorNotificationsEmailsPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendNotificationEmails** | **bool** |  | 
**NotificationsSubscribed** | Pointer to [**[]AdministratorNotification**](AdministratorNotification.md) |  | [optional] 
**EnabledDate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAdministratorNotificationsEmailsPreferences

`func NewAdministratorNotificationsEmailsPreferences(sendNotificationEmails bool, ) *AdministratorNotificationsEmailsPreferences`

NewAdministratorNotificationsEmailsPreferences instantiates a new AdministratorNotificationsEmailsPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorNotificationsEmailsPreferencesWithDefaults

`func NewAdministratorNotificationsEmailsPreferencesWithDefaults() *AdministratorNotificationsEmailsPreferences`

NewAdministratorNotificationsEmailsPreferencesWithDefaults instantiates a new AdministratorNotificationsEmailsPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendNotificationEmails

`func (o *AdministratorNotificationsEmailsPreferences) GetSendNotificationEmails() bool`

GetSendNotificationEmails returns the SendNotificationEmails field if non-nil, zero value otherwise.

### GetSendNotificationEmailsOk

`func (o *AdministratorNotificationsEmailsPreferences) GetSendNotificationEmailsOk() (*bool, bool)`

GetSendNotificationEmailsOk returns a tuple with the SendNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotificationEmails

`func (o *AdministratorNotificationsEmailsPreferences) SetSendNotificationEmails(v bool)`

SetSendNotificationEmails sets SendNotificationEmails field to given value.


### GetNotificationsSubscribed

`func (o *AdministratorNotificationsEmailsPreferences) GetNotificationsSubscribed() []AdministratorNotification`

GetNotificationsSubscribed returns the NotificationsSubscribed field if non-nil, zero value otherwise.

### GetNotificationsSubscribedOk

`func (o *AdministratorNotificationsEmailsPreferences) GetNotificationsSubscribedOk() (*[]AdministratorNotification, bool)`

GetNotificationsSubscribedOk returns a tuple with the NotificationsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsSubscribed

`func (o *AdministratorNotificationsEmailsPreferences) SetNotificationsSubscribed(v []AdministratorNotification)`

SetNotificationsSubscribed sets NotificationsSubscribed field to given value.

### HasNotificationsSubscribed

`func (o *AdministratorNotificationsEmailsPreferences) HasNotificationsSubscribed() bool`

HasNotificationsSubscribed returns a boolean if a field has been set.

### SetNotificationsSubscribedNil

`func (o *AdministratorNotificationsEmailsPreferences) SetNotificationsSubscribedNil(b bool)`

 SetNotificationsSubscribedNil sets the value for NotificationsSubscribed to be an explicit nil

### UnsetNotificationsSubscribed
`func (o *AdministratorNotificationsEmailsPreferences) UnsetNotificationsSubscribed()`

UnsetNotificationsSubscribed ensures that no value is present for NotificationsSubscribed, not even an explicit nil
### GetEnabledDate

`func (o *AdministratorNotificationsEmailsPreferences) GetEnabledDate() string`

GetEnabledDate returns the EnabledDate field if non-nil, zero value otherwise.

### GetEnabledDateOk

`func (o *AdministratorNotificationsEmailsPreferences) GetEnabledDateOk() (*string, bool)`

GetEnabledDateOk returns a tuple with the EnabledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledDate

`func (o *AdministratorNotificationsEmailsPreferences) SetEnabledDate(v string)`

SetEnabledDate sets EnabledDate field to given value.

### HasEnabledDate

`func (o *AdministratorNotificationsEmailsPreferences) HasEnabledDate() bool`

HasEnabledDate returns a boolean if a field has been set.

### SetEnabledDateNil

`func (o *AdministratorNotificationsEmailsPreferences) SetEnabledDateNil(b bool)`

 SetEnabledDateNil sets the value for EnabledDate to be an explicit nil

### UnsetEnabledDate
`func (o *AdministratorNotificationsEmailsPreferences) UnsetEnabledDate()`

UnsetEnabledDate ensures that no value is present for EnabledDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


