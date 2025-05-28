# ManagedDomainConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** |  | [optional] 
**DomainUsers** | Pointer to [**[]ManagedDomainUser**](ManagedDomainUser.md) |  | [optional] [readonly] 

## Methods

### NewManagedDomainConfiguration

`func NewManagedDomainConfiguration() *ManagedDomainConfiguration`

NewManagedDomainConfiguration instantiates a new ManagedDomainConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDomainConfigurationWithDefaults

`func NewManagedDomainConfigurationWithDefaults() *ManagedDomainConfiguration`

NewManagedDomainConfigurationWithDefaults instantiates a new ManagedDomainConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *ManagedDomainConfiguration) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ManagedDomainConfiguration) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ManagedDomainConfiguration) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ManagedDomainConfiguration) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainUsers

`func (o *ManagedDomainConfiguration) GetDomainUsers() []ManagedDomainUser`

GetDomainUsers returns the DomainUsers field if non-nil, zero value otherwise.

### GetDomainUsersOk

`func (o *ManagedDomainConfiguration) GetDomainUsersOk() (*[]ManagedDomainUser, bool)`

GetDomainUsersOk returns a tuple with the DomainUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUsers

`func (o *ManagedDomainConfiguration) SetDomainUsers(v []ManagedDomainUser)`

SetDomainUsers sets DomainUsers field to given value.

### HasDomainUsers

`func (o *ManagedDomainConfiguration) HasDomainUsers() bool`

HasDomainUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


