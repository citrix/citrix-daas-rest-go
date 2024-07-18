# LicenseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HdxLicenseState** | Pointer to [**LicenseState**](LicenseState.md) |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Windows365Licenses** | Pointer to **[]string** |  | [optional] 
**UserId** | Pointer to **NullableString** | AAD User ID for which we want to assign/revoke license | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | AAD UPN for provided User | [optional] 

## Methods

### NewLicenseUser

`func NewLicenseUser() *LicenseUser`

NewLicenseUser instantiates a new LicenseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseUserWithDefaults

`func NewLicenseUserWithDefaults() *LicenseUser`

NewLicenseUserWithDefaults instantiates a new LicenseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdxLicenseState

`func (o *LicenseUser) GetHdxLicenseState() LicenseState`

GetHdxLicenseState returns the HdxLicenseState field if non-nil, zero value otherwise.

### GetHdxLicenseStateOk

`func (o *LicenseUser) GetHdxLicenseStateOk() (*LicenseState, bool)`

GetHdxLicenseStateOk returns a tuple with the HdxLicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdxLicenseState

`func (o *LicenseUser) SetHdxLicenseState(v LicenseState)`

SetHdxLicenseState sets HdxLicenseState field to given value.

### HasHdxLicenseState

`func (o *LicenseUser) HasHdxLicenseState() bool`

HasHdxLicenseState returns a boolean if a field has been set.

### GetDisplayName

`func (o *LicenseUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LicenseUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LicenseUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *LicenseUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *LicenseUser) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *LicenseUser) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetWindows365Licenses

`func (o *LicenseUser) GetWindows365Licenses() []string`

GetWindows365Licenses returns the Windows365Licenses field if non-nil, zero value otherwise.

### GetWindows365LicensesOk

`func (o *LicenseUser) GetWindows365LicensesOk() (*[]string, bool)`

GetWindows365LicensesOk returns a tuple with the Windows365Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows365Licenses

`func (o *LicenseUser) SetWindows365Licenses(v []string)`

SetWindows365Licenses sets Windows365Licenses field to given value.

### HasWindows365Licenses

`func (o *LicenseUser) HasWindows365Licenses() bool`

HasWindows365Licenses returns a boolean if a field has been set.

### SetWindows365LicensesNil

`func (o *LicenseUser) SetWindows365LicensesNil(b bool)`

 SetWindows365LicensesNil sets the value for Windows365Licenses to be an explicit nil

### UnsetWindows365Licenses
`func (o *LicenseUser) UnsetWindows365Licenses()`

UnsetWindows365Licenses ensures that no value is present for Windows365Licenses, not even an explicit nil
### GetUserId

`func (o *LicenseUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LicenseUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LicenseUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LicenseUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *LicenseUser) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *LicenseUser) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserPrincipalName

`func (o *LicenseUser) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *LicenseUser) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *LicenseUser) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *LicenseUser) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *LicenseUser) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *LicenseUser) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


