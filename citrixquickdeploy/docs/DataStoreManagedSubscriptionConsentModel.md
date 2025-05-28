# DataStoreManagedSubscriptionConsentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentedBy** | Pointer to **string** | The name of the user that consented | [optional] 
**ConsentedAt** | Pointer to **time.Time** | The date time when the consent was given | [optional] 

## Methods

### NewDataStoreManagedSubscriptionConsentModel

`func NewDataStoreManagedSubscriptionConsentModel() *DataStoreManagedSubscriptionConsentModel`

NewDataStoreManagedSubscriptionConsentModel instantiates a new DataStoreManagedSubscriptionConsentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreManagedSubscriptionConsentModelWithDefaults

`func NewDataStoreManagedSubscriptionConsentModelWithDefaults() *DataStoreManagedSubscriptionConsentModel`

NewDataStoreManagedSubscriptionConsentModelWithDefaults instantiates a new DataStoreManagedSubscriptionConsentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentedBy

`func (o *DataStoreManagedSubscriptionConsentModel) GetConsentedBy() string`

GetConsentedBy returns the ConsentedBy field if non-nil, zero value otherwise.

### GetConsentedByOk

`func (o *DataStoreManagedSubscriptionConsentModel) GetConsentedByOk() (*string, bool)`

GetConsentedByOk returns a tuple with the ConsentedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentedBy

`func (o *DataStoreManagedSubscriptionConsentModel) SetConsentedBy(v string)`

SetConsentedBy sets ConsentedBy field to given value.

### HasConsentedBy

`func (o *DataStoreManagedSubscriptionConsentModel) HasConsentedBy() bool`

HasConsentedBy returns a boolean if a field has been set.

### GetConsentedAt

`func (o *DataStoreManagedSubscriptionConsentModel) GetConsentedAt() time.Time`

GetConsentedAt returns the ConsentedAt field if non-nil, zero value otherwise.

### GetConsentedAtOk

`func (o *DataStoreManagedSubscriptionConsentModel) GetConsentedAtOk() (*time.Time, bool)`

GetConsentedAtOk returns a tuple with the ConsentedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentedAt

`func (o *DataStoreManagedSubscriptionConsentModel) SetConsentedAt(v time.Time)`

SetConsentedAt sets ConsentedAt field to given value.

### HasConsentedAt

`func (o *DataStoreManagedSubscriptionConsentModel) HasConsentedAt() bool`

HasConsentedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


