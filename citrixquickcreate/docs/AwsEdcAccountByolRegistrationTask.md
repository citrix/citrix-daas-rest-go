# AwsEdcAccountByolRegistrationTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedicatedTenancyManagementCidrRange** | Pointer to **NullableString** |  | [optional] 
**DedicatedTenancySupportEnabled** | Pointer to **NullableBool** |  | [optional] 
**ModificationState** | Pointer to [**NullableAwsEdcDedicatedTenancyState**](AwsEdcDedicatedTenancyState.md) |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAwsEdcAccountByolRegistrationTask

`func NewAwsEdcAccountByolRegistrationTask() *AwsEdcAccountByolRegistrationTask`

NewAwsEdcAccountByolRegistrationTask instantiates a new AwsEdcAccountByolRegistrationTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcAccountByolRegistrationTaskWithDefaults

`func NewAwsEdcAccountByolRegistrationTaskWithDefaults() *AwsEdcAccountByolRegistrationTask`

NewAwsEdcAccountByolRegistrationTaskWithDefaults instantiates a new AwsEdcAccountByolRegistrationTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedTenancyManagementCidrRange

`func (o *AwsEdcAccountByolRegistrationTask) GetDedicatedTenancyManagementCidrRange() string`

GetDedicatedTenancyManagementCidrRange returns the DedicatedTenancyManagementCidrRange field if non-nil, zero value otherwise.

### GetDedicatedTenancyManagementCidrRangeOk

`func (o *AwsEdcAccountByolRegistrationTask) GetDedicatedTenancyManagementCidrRangeOk() (*string, bool)`

GetDedicatedTenancyManagementCidrRangeOk returns a tuple with the DedicatedTenancyManagementCidrRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTenancyManagementCidrRange

`func (o *AwsEdcAccountByolRegistrationTask) SetDedicatedTenancyManagementCidrRange(v string)`

SetDedicatedTenancyManagementCidrRange sets DedicatedTenancyManagementCidrRange field to given value.

### HasDedicatedTenancyManagementCidrRange

`func (o *AwsEdcAccountByolRegistrationTask) HasDedicatedTenancyManagementCidrRange() bool`

HasDedicatedTenancyManagementCidrRange returns a boolean if a field has been set.

### SetDedicatedTenancyManagementCidrRangeNil

`func (o *AwsEdcAccountByolRegistrationTask) SetDedicatedTenancyManagementCidrRangeNil(b bool)`

 SetDedicatedTenancyManagementCidrRangeNil sets the value for DedicatedTenancyManagementCidrRange to be an explicit nil

### UnsetDedicatedTenancyManagementCidrRange
`func (o *AwsEdcAccountByolRegistrationTask) UnsetDedicatedTenancyManagementCidrRange()`

UnsetDedicatedTenancyManagementCidrRange ensures that no value is present for DedicatedTenancyManagementCidrRange, not even an explicit nil
### GetDedicatedTenancySupportEnabled

`func (o *AwsEdcAccountByolRegistrationTask) GetDedicatedTenancySupportEnabled() bool`

GetDedicatedTenancySupportEnabled returns the DedicatedTenancySupportEnabled field if non-nil, zero value otherwise.

### GetDedicatedTenancySupportEnabledOk

`func (o *AwsEdcAccountByolRegistrationTask) GetDedicatedTenancySupportEnabledOk() (*bool, bool)`

GetDedicatedTenancySupportEnabledOk returns a tuple with the DedicatedTenancySupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTenancySupportEnabled

`func (o *AwsEdcAccountByolRegistrationTask) SetDedicatedTenancySupportEnabled(v bool)`

SetDedicatedTenancySupportEnabled sets DedicatedTenancySupportEnabled field to given value.

### HasDedicatedTenancySupportEnabled

`func (o *AwsEdcAccountByolRegistrationTask) HasDedicatedTenancySupportEnabled() bool`

HasDedicatedTenancySupportEnabled returns a boolean if a field has been set.

### SetDedicatedTenancySupportEnabledNil

`func (o *AwsEdcAccountByolRegistrationTask) SetDedicatedTenancySupportEnabledNil(b bool)`

 SetDedicatedTenancySupportEnabledNil sets the value for DedicatedTenancySupportEnabled to be an explicit nil

### UnsetDedicatedTenancySupportEnabled
`func (o *AwsEdcAccountByolRegistrationTask) UnsetDedicatedTenancySupportEnabled()`

UnsetDedicatedTenancySupportEnabled ensures that no value is present for DedicatedTenancySupportEnabled, not even an explicit nil
### GetModificationState

`func (o *AwsEdcAccountByolRegistrationTask) GetModificationState() AwsEdcDedicatedTenancyState`

GetModificationState returns the ModificationState field if non-nil, zero value otherwise.

### GetModificationStateOk

`func (o *AwsEdcAccountByolRegistrationTask) GetModificationStateOk() (*AwsEdcDedicatedTenancyState, bool)`

GetModificationStateOk returns a tuple with the ModificationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationState

`func (o *AwsEdcAccountByolRegistrationTask) SetModificationState(v AwsEdcDedicatedTenancyState)`

SetModificationState sets ModificationState field to given value.

### HasModificationState

`func (o *AwsEdcAccountByolRegistrationTask) HasModificationState() bool`

HasModificationState returns a boolean if a field has been set.

### SetModificationStateNil

`func (o *AwsEdcAccountByolRegistrationTask) SetModificationStateNil(b bool)`

 SetModificationStateNil sets the value for ModificationState to be an explicit nil

### UnsetModificationState
`func (o *AwsEdcAccountByolRegistrationTask) UnsetModificationState()`

UnsetModificationState ensures that no value is present for ModificationState, not even an explicit nil
### GetErrorCode

`func (o *AwsEdcAccountByolRegistrationTask) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AwsEdcAccountByolRegistrationTask) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AwsEdcAccountByolRegistrationTask) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AwsEdcAccountByolRegistrationTask) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *AwsEdcAccountByolRegistrationTask) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AwsEdcAccountByolRegistrationTask) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *AwsEdcAccountByolRegistrationTask) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AwsEdcAccountByolRegistrationTask) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AwsEdcAccountByolRegistrationTask) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AwsEdcAccountByolRegistrationTask) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AwsEdcAccountByolRegistrationTask) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AwsEdcAccountByolRegistrationTask) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


