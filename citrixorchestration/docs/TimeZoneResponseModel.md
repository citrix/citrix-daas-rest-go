# TimeZoneResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Human-readable name of the time zone, including the UTC offset of the standard (non-daylight-savings) time within the time zone.  For example, \&quot;(UTC-10:00) Aleutian Islands\&quot;.   This is localized into the language requested by the caller. | [optional] 
**FullName** | Pointer to **NullableString** | Human-readable name of the time zone, including the UTC offset of the standard (non-daylight-savings) time within the time zone.  For example, \&quot;(UTC-10:00) Aleutian Islands\&quot;.   This is localized into the language requested by the caller. | [optional] 
**Primary** | Pointer to **NullableString** | IANA identifier of the time zone within the primary territory covered by the time zone.   The  and  identifiers are usually, but not always, the same.  For example, the primary entry for \&quot;India Standard Time\&quot; is \&quot;Asia/Calcutta\&quot; while the canonical entry is \&quot;Asia/Kolkata\&quot;. | [optional] 
**Canonical** | Pointer to **NullableString** | Canonical IANA identifier of the time zone.   The  and  identifiers are usually, but not always, the same.  For example, the primary entry for \&quot;India Standard Time\&quot; is \&quot;Asia/Calcutta\&quot; while the canonical entry is \&quot;Asia/Kolkata\&quot;. | [optional] 
**WindowsId** | Pointer to **NullableString** | Windows identifier of the time zone.   Note that although this may appear to be human-readable, it is **NOT** localized and therefore, should not be displayed in user interfaces.  Instead, use  or  properties for this purpose. | [optional] 
**UtcOffsetSeconds** | Pointer to **int32** | UTC offset of the standard (non-daylight-savings) time within the time zone. | [optional] 
**UtcOffset** | Pointer to **NullableString** | String representation of UtcOffsetSeconds.  Always starts with &#x60;+&#x60; or &#x60;-&#x60;, followed by a two hours digits, &#x60;:&#x60;, and two minutes digits.  (e.g. &#x60;-05:00&#x60; / &#x60;+00:00&#x60;) | [optional] 
**TzdbIds** | Pointer to **[]string** | List of IANA identifiers which can be used to represent this time zone.  May be non-exhaustive. | [optional] 

## Methods

### NewTimeZoneResponseModel

`func NewTimeZoneResponseModel() *TimeZoneResponseModel`

NewTimeZoneResponseModel instantiates a new TimeZoneResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeZoneResponseModelWithDefaults

`func NewTimeZoneResponseModelWithDefaults() *TimeZoneResponseModel`

NewTimeZoneResponseModelWithDefaults instantiates a new TimeZoneResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TimeZoneResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeZoneResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeZoneResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimeZoneResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TimeZoneResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TimeZoneResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *TimeZoneResponseModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *TimeZoneResponseModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *TimeZoneResponseModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *TimeZoneResponseModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *TimeZoneResponseModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *TimeZoneResponseModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetPrimary

`func (o *TimeZoneResponseModel) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *TimeZoneResponseModel) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *TimeZoneResponseModel) SetPrimary(v string)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *TimeZoneResponseModel) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *TimeZoneResponseModel) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *TimeZoneResponseModel) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
### GetCanonical

`func (o *TimeZoneResponseModel) GetCanonical() string`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *TimeZoneResponseModel) GetCanonicalOk() (*string, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *TimeZoneResponseModel) SetCanonical(v string)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *TimeZoneResponseModel) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### SetCanonicalNil

`func (o *TimeZoneResponseModel) SetCanonicalNil(b bool)`

 SetCanonicalNil sets the value for Canonical to be an explicit nil

### UnsetCanonical
`func (o *TimeZoneResponseModel) UnsetCanonical()`

UnsetCanonical ensures that no value is present for Canonical, not even an explicit nil
### GetWindowsId

`func (o *TimeZoneResponseModel) GetWindowsId() string`

GetWindowsId returns the WindowsId field if non-nil, zero value otherwise.

### GetWindowsIdOk

`func (o *TimeZoneResponseModel) GetWindowsIdOk() (*string, bool)`

GetWindowsIdOk returns a tuple with the WindowsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsId

`func (o *TimeZoneResponseModel) SetWindowsId(v string)`

SetWindowsId sets WindowsId field to given value.

### HasWindowsId

`func (o *TimeZoneResponseModel) HasWindowsId() bool`

HasWindowsId returns a boolean if a field has been set.

### SetWindowsIdNil

`func (o *TimeZoneResponseModel) SetWindowsIdNil(b bool)`

 SetWindowsIdNil sets the value for WindowsId to be an explicit nil

### UnsetWindowsId
`func (o *TimeZoneResponseModel) UnsetWindowsId()`

UnsetWindowsId ensures that no value is present for WindowsId, not even an explicit nil
### GetUtcOffsetSeconds

`func (o *TimeZoneResponseModel) GetUtcOffsetSeconds() int32`

GetUtcOffsetSeconds returns the UtcOffsetSeconds field if non-nil, zero value otherwise.

### GetUtcOffsetSecondsOk

`func (o *TimeZoneResponseModel) GetUtcOffsetSecondsOk() (*int32, bool)`

GetUtcOffsetSecondsOk returns a tuple with the UtcOffsetSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffsetSeconds

`func (o *TimeZoneResponseModel) SetUtcOffsetSeconds(v int32)`

SetUtcOffsetSeconds sets UtcOffsetSeconds field to given value.

### HasUtcOffsetSeconds

`func (o *TimeZoneResponseModel) HasUtcOffsetSeconds() bool`

HasUtcOffsetSeconds returns a boolean if a field has been set.

### GetUtcOffset

`func (o *TimeZoneResponseModel) GetUtcOffset() string`

GetUtcOffset returns the UtcOffset field if non-nil, zero value otherwise.

### GetUtcOffsetOk

`func (o *TimeZoneResponseModel) GetUtcOffsetOk() (*string, bool)`

GetUtcOffsetOk returns a tuple with the UtcOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffset

`func (o *TimeZoneResponseModel) SetUtcOffset(v string)`

SetUtcOffset sets UtcOffset field to given value.

### HasUtcOffset

`func (o *TimeZoneResponseModel) HasUtcOffset() bool`

HasUtcOffset returns a boolean if a field has been set.

### SetUtcOffsetNil

`func (o *TimeZoneResponseModel) SetUtcOffsetNil(b bool)`

 SetUtcOffsetNil sets the value for UtcOffset to be an explicit nil

### UnsetUtcOffset
`func (o *TimeZoneResponseModel) UnsetUtcOffset()`

UnsetUtcOffset ensures that no value is present for UtcOffset, not even an explicit nil
### GetTzdbIds

`func (o *TimeZoneResponseModel) GetTzdbIds() []string`

GetTzdbIds returns the TzdbIds field if non-nil, zero value otherwise.

### GetTzdbIdsOk

`func (o *TimeZoneResponseModel) GetTzdbIdsOk() (*[]string, bool)`

GetTzdbIdsOk returns a tuple with the TzdbIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTzdbIds

`func (o *TimeZoneResponseModel) SetTzdbIds(v []string)`

SetTzdbIds sets TzdbIds field to given value.

### HasTzdbIds

`func (o *TimeZoneResponseModel) HasTzdbIds() bool`

HasTzdbIds returns a boolean if a field has been set.

### SetTzdbIdsNil

`func (o *TimeZoneResponseModel) SetTzdbIdsNil(b bool)`

 SetTzdbIdsNil sets the value for TzdbIds to be an explicit nil

### UnsetTzdbIds
`func (o *TimeZoneResponseModel) UnsetTzdbIds()`

UnsetTzdbIds ensures that no value is present for TzdbIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


