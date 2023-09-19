# LowLevelOperationDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **NullableString** | The date and time that the operation completed. This will be null if the operation is still in progress, or if the operation never completed.              | [optional] 
**FormattedEndTime** | Pointer to **NullableString** | The formatted date and time that the operation completed. RFC 3339 compatible format. This will be null if the operation is still in progress, or if the operation never completed.              | [optional] 
**IsSuccessful** | Pointer to **NullableBool** | Indicates whether the operation completed successfully or not. This will be null if the operation is still in progress, or if the operation didn&#39;t complete.              | [optional] 
**NewValue** | Pointer to **NullableString** | The new property value.              | [optional] 
**PreviousValue** | Pointer to **NullableString** | The previous property value.              | [optional] 
**PropertyName** | Pointer to **NullableString** | The name of the changed property.              | [optional] 
**AddValue** | Pointer to **NullableString** | If the object property contains a set of values, this specifies the new value which was added to the set.              | [optional] 
**RemoveValue** | Pointer to **NullableString** | If the object property contains a set of values, this specifies the value which was removed from the set.              | [optional] 
**StartTime** | **string** | The date and time that the operation started.              | 
**FormattedStartTime** | **string** | The formatted date and time that the operation started. RFC 3339 compatible format.              | 
**TargetName** | **string** | The name of the target object affected by the operation.              | 
**TargetUid** | **string** | The unique identifier of the target object affected by the operation.              | 
**TargetType** | Pointer to **NullableString** | The type of the target object.              | [optional] 
**Text** | **string** | The description of operation performed on the target object.              | 

## Methods

### NewLowLevelOperationDetailResponseModel

`func NewLowLevelOperationDetailResponseModel(startTime string, formattedStartTime string, targetName string, targetUid string, text string, ) *LowLevelOperationDetailResponseModel`

NewLowLevelOperationDetailResponseModel instantiates a new LowLevelOperationDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLowLevelOperationDetailResponseModelWithDefaults

`func NewLowLevelOperationDetailResponseModelWithDefaults() *LowLevelOperationDetailResponseModel`

NewLowLevelOperationDetailResponseModelWithDefaults instantiates a new LowLevelOperationDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *LowLevelOperationDetailResponseModel) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LowLevelOperationDetailResponseModel) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LowLevelOperationDetailResponseModel) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *LowLevelOperationDetailResponseModel) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *LowLevelOperationDetailResponseModel) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *LowLevelOperationDetailResponseModel) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetFormattedEndTime

`func (o *LowLevelOperationDetailResponseModel) GetFormattedEndTime() string`

GetFormattedEndTime returns the FormattedEndTime field if non-nil, zero value otherwise.

### GetFormattedEndTimeOk

`func (o *LowLevelOperationDetailResponseModel) GetFormattedEndTimeOk() (*string, bool)`

GetFormattedEndTimeOk returns a tuple with the FormattedEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedEndTime

`func (o *LowLevelOperationDetailResponseModel) SetFormattedEndTime(v string)`

SetFormattedEndTime sets FormattedEndTime field to given value.

### HasFormattedEndTime

`func (o *LowLevelOperationDetailResponseModel) HasFormattedEndTime() bool`

HasFormattedEndTime returns a boolean if a field has been set.

### SetFormattedEndTimeNil

`func (o *LowLevelOperationDetailResponseModel) SetFormattedEndTimeNil(b bool)`

 SetFormattedEndTimeNil sets the value for FormattedEndTime to be an explicit nil

### UnsetFormattedEndTime
`func (o *LowLevelOperationDetailResponseModel) UnsetFormattedEndTime()`

UnsetFormattedEndTime ensures that no value is present for FormattedEndTime, not even an explicit nil
### GetIsSuccessful

`func (o *LowLevelOperationDetailResponseModel) GetIsSuccessful() bool`

GetIsSuccessful returns the IsSuccessful field if non-nil, zero value otherwise.

### GetIsSuccessfulOk

`func (o *LowLevelOperationDetailResponseModel) GetIsSuccessfulOk() (*bool, bool)`

GetIsSuccessfulOk returns a tuple with the IsSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccessful

`func (o *LowLevelOperationDetailResponseModel) SetIsSuccessful(v bool)`

SetIsSuccessful sets IsSuccessful field to given value.

### HasIsSuccessful

`func (o *LowLevelOperationDetailResponseModel) HasIsSuccessful() bool`

HasIsSuccessful returns a boolean if a field has been set.

### SetIsSuccessfulNil

`func (o *LowLevelOperationDetailResponseModel) SetIsSuccessfulNil(b bool)`

 SetIsSuccessfulNil sets the value for IsSuccessful to be an explicit nil

### UnsetIsSuccessful
`func (o *LowLevelOperationDetailResponseModel) UnsetIsSuccessful()`

UnsetIsSuccessful ensures that no value is present for IsSuccessful, not even an explicit nil
### GetNewValue

`func (o *LowLevelOperationDetailResponseModel) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *LowLevelOperationDetailResponseModel) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *LowLevelOperationDetailResponseModel) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *LowLevelOperationDetailResponseModel) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### SetNewValueNil

`func (o *LowLevelOperationDetailResponseModel) SetNewValueNil(b bool)`

 SetNewValueNil sets the value for NewValue to be an explicit nil

### UnsetNewValue
`func (o *LowLevelOperationDetailResponseModel) UnsetNewValue()`

UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
### GetPreviousValue

`func (o *LowLevelOperationDetailResponseModel) GetPreviousValue() string`

GetPreviousValue returns the PreviousValue field if non-nil, zero value otherwise.

### GetPreviousValueOk

`func (o *LowLevelOperationDetailResponseModel) GetPreviousValueOk() (*string, bool)`

GetPreviousValueOk returns a tuple with the PreviousValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValue

`func (o *LowLevelOperationDetailResponseModel) SetPreviousValue(v string)`

SetPreviousValue sets PreviousValue field to given value.

### HasPreviousValue

`func (o *LowLevelOperationDetailResponseModel) HasPreviousValue() bool`

HasPreviousValue returns a boolean if a field has been set.

### SetPreviousValueNil

`func (o *LowLevelOperationDetailResponseModel) SetPreviousValueNil(b bool)`

 SetPreviousValueNil sets the value for PreviousValue to be an explicit nil

### UnsetPreviousValue
`func (o *LowLevelOperationDetailResponseModel) UnsetPreviousValue()`

UnsetPreviousValue ensures that no value is present for PreviousValue, not even an explicit nil
### GetPropertyName

`func (o *LowLevelOperationDetailResponseModel) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *LowLevelOperationDetailResponseModel) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *LowLevelOperationDetailResponseModel) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *LowLevelOperationDetailResponseModel) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### SetPropertyNameNil

`func (o *LowLevelOperationDetailResponseModel) SetPropertyNameNil(b bool)`

 SetPropertyNameNil sets the value for PropertyName to be an explicit nil

### UnsetPropertyName
`func (o *LowLevelOperationDetailResponseModel) UnsetPropertyName()`

UnsetPropertyName ensures that no value is present for PropertyName, not even an explicit nil
### GetAddValue

`func (o *LowLevelOperationDetailResponseModel) GetAddValue() string`

GetAddValue returns the AddValue field if non-nil, zero value otherwise.

### GetAddValueOk

`func (o *LowLevelOperationDetailResponseModel) GetAddValueOk() (*string, bool)`

GetAddValueOk returns a tuple with the AddValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddValue

`func (o *LowLevelOperationDetailResponseModel) SetAddValue(v string)`

SetAddValue sets AddValue field to given value.

### HasAddValue

`func (o *LowLevelOperationDetailResponseModel) HasAddValue() bool`

HasAddValue returns a boolean if a field has been set.

### SetAddValueNil

`func (o *LowLevelOperationDetailResponseModel) SetAddValueNil(b bool)`

 SetAddValueNil sets the value for AddValue to be an explicit nil

### UnsetAddValue
`func (o *LowLevelOperationDetailResponseModel) UnsetAddValue()`

UnsetAddValue ensures that no value is present for AddValue, not even an explicit nil
### GetRemoveValue

`func (o *LowLevelOperationDetailResponseModel) GetRemoveValue() string`

GetRemoveValue returns the RemoveValue field if non-nil, zero value otherwise.

### GetRemoveValueOk

`func (o *LowLevelOperationDetailResponseModel) GetRemoveValueOk() (*string, bool)`

GetRemoveValueOk returns a tuple with the RemoveValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveValue

`func (o *LowLevelOperationDetailResponseModel) SetRemoveValue(v string)`

SetRemoveValue sets RemoveValue field to given value.

### HasRemoveValue

`func (o *LowLevelOperationDetailResponseModel) HasRemoveValue() bool`

HasRemoveValue returns a boolean if a field has been set.

### SetRemoveValueNil

`func (o *LowLevelOperationDetailResponseModel) SetRemoveValueNil(b bool)`

 SetRemoveValueNil sets the value for RemoveValue to be an explicit nil

### UnsetRemoveValue
`func (o *LowLevelOperationDetailResponseModel) UnsetRemoveValue()`

UnsetRemoveValue ensures that no value is present for RemoveValue, not even an explicit nil
### GetStartTime

`func (o *LowLevelOperationDetailResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LowLevelOperationDetailResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LowLevelOperationDetailResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetFormattedStartTime

`func (o *LowLevelOperationDetailResponseModel) GetFormattedStartTime() string`

GetFormattedStartTime returns the FormattedStartTime field if non-nil, zero value otherwise.

### GetFormattedStartTimeOk

`func (o *LowLevelOperationDetailResponseModel) GetFormattedStartTimeOk() (*string, bool)`

GetFormattedStartTimeOk returns a tuple with the FormattedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedStartTime

`func (o *LowLevelOperationDetailResponseModel) SetFormattedStartTime(v string)`

SetFormattedStartTime sets FormattedStartTime field to given value.


### GetTargetName

`func (o *LowLevelOperationDetailResponseModel) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *LowLevelOperationDetailResponseModel) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *LowLevelOperationDetailResponseModel) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetTargetUid

`func (o *LowLevelOperationDetailResponseModel) GetTargetUid() string`

GetTargetUid returns the TargetUid field if non-nil, zero value otherwise.

### GetTargetUidOk

`func (o *LowLevelOperationDetailResponseModel) GetTargetUidOk() (*string, bool)`

GetTargetUidOk returns a tuple with the TargetUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUid

`func (o *LowLevelOperationDetailResponseModel) SetTargetUid(v string)`

SetTargetUid sets TargetUid field to given value.


### GetTargetType

`func (o *LowLevelOperationDetailResponseModel) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *LowLevelOperationDetailResponseModel) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *LowLevelOperationDetailResponseModel) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *LowLevelOperationDetailResponseModel) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *LowLevelOperationDetailResponseModel) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *LowLevelOperationDetailResponseModel) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetText

`func (o *LowLevelOperationDetailResponseModel) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *LowLevelOperationDetailResponseModel) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *LowLevelOperationDetailResponseModel) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


