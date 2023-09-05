# IdentityUserResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name of the user or group. Usually something like \&quot;John Q. Doe\&quot;. | 
**SamName** | Pointer to **string** | Fully qualified sAMAccountName of the User or Group object (including domain). | [optional] 
**SamAccountName** | Pointer to **string** | Unqualified SAM name (Directory property) of the User or Group object.  Does not include the domain. The sAMAccountName is unique within a domain. | [optional] 
**CommonName** | Pointer to **string** | The common name associated with the User or Group object. | [optional] 
**PrincipalName** | Pointer to **string** | The user principal name of a User object. The UPN must be unique among all security principal objects within the directory forest. | [optional] 
**Oid** | Pointer to **string** | The Object Identifier associated with the User or Group object. | [optional] 
**Sid** | Pointer to **string** | The Security Identifier associated with the User or Group object. | [optional] 
**UserIdentity** | Pointer to **string** | The User Identity. | [optional] 
**GroupSids** | Pointer to **[]string** | Security Identifiers of groups of which the object is a member. | [optional] 
**DenyOnlySids** | Pointer to **[]string** | Security Identifiers of resources to which the object is explicitly denied access. | [optional] 
**PasswordCanExpire** | Pointer to **bool** | Indicates whether the password can expire on this User object. | [optional] 
**DaysUntilPasswordExpiry** | Pointer to **int64** | Indicates the number of days before the password can expire on this User object. | [optional] 
**IsBuiltIn** | Pointer to **bool** | Indicates whether this a built-in object. | [optional] 
**IsGroup** | Pointer to **bool** | Indicates whether this is a User or Group object. | [optional] 
**Name** | Pointer to **string** | The name of the User or Group object. | [optional] 
**StreetAddress** | Pointer to **string** | The street address associated with the User object. | [optional] 
**City** | Pointer to **string** | The city associated with the User object. | [optional] 
**State** | Pointer to **string** | The state associated with the User object. | [optional] 
**Country** | Pointer to **string** | The country associated with the user object. | [optional] 
**Mail** | Pointer to **string** | The e-mail address associated with the user object. | [optional] 
**TelephoneNumber** | Pointer to **string** | The telephone number associated with the user object. | [optional] 
**HomePhone** | Pointer to **string** | The home telephone number associated with the user object. | [optional] 
**Mobile** | Pointer to **string** | The mobile telephone number associated with the user object. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the User object is disabled. NOTE that this is opposite of the low-level SDK, which has \&quot;IsDisabled\&quot;. By changing to \&quot;Enabled\&quot; it avoids a confusing double-negative.  It also matches all other SDK objects by using \&quot;Enabled\&quot; rather than \&quot;IsEnabled\&quot;. | [optional] 
**Locked** | Pointer to **bool** | Indicates whether the User object is locked. | [optional] 
**PropertiesFetched** | **int32** | Properties fetched and populated in the User or Group object.  This is a bitfield indicating the fetched properties. | 
**Claims** | Pointer to [**IdentityUserResponseModelAllOfClaims**](IdentityUserResponseModelAllOfClaims.md) |  | [optional] 
**Directory** | Pointer to [**IdentityUserResponseModelAllOfDirectory**](IdentityUserResponseModelAllOfDirectory.md) |  | [optional] 

## Methods

### NewIdentityUserResponseModelAllOf

`func NewIdentityUserResponseModelAllOf(displayName string, propertiesFetched int32, ) *IdentityUserResponseModelAllOf`

NewIdentityUserResponseModelAllOf instantiates a new IdentityUserResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUserResponseModelAllOfWithDefaults

`func NewIdentityUserResponseModelAllOfWithDefaults() *IdentityUserResponseModelAllOf`

NewIdentityUserResponseModelAllOfWithDefaults instantiates a new IdentityUserResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *IdentityUserResponseModelAllOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityUserResponseModelAllOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityUserResponseModelAllOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSamName

`func (o *IdentityUserResponseModelAllOf) GetSamName() string`

GetSamName returns the SamName field if non-nil, zero value otherwise.

### GetSamNameOk

`func (o *IdentityUserResponseModelAllOf) GetSamNameOk() (*string, bool)`

GetSamNameOk returns a tuple with the SamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamName

`func (o *IdentityUserResponseModelAllOf) SetSamName(v string)`

SetSamName sets SamName field to given value.

### HasSamName

`func (o *IdentityUserResponseModelAllOf) HasSamName() bool`

HasSamName returns a boolean if a field has been set.

### GetSamAccountName

`func (o *IdentityUserResponseModelAllOf) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *IdentityUserResponseModelAllOf) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *IdentityUserResponseModelAllOf) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.

### HasSamAccountName

`func (o *IdentityUserResponseModelAllOf) HasSamAccountName() bool`

HasSamAccountName returns a boolean if a field has been set.

### GetCommonName

`func (o *IdentityUserResponseModelAllOf) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *IdentityUserResponseModelAllOf) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *IdentityUserResponseModelAllOf) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *IdentityUserResponseModelAllOf) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetPrincipalName

`func (o *IdentityUserResponseModelAllOf) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *IdentityUserResponseModelAllOf) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *IdentityUserResponseModelAllOf) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *IdentityUserResponseModelAllOf) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### GetOid

`func (o *IdentityUserResponseModelAllOf) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *IdentityUserResponseModelAllOf) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *IdentityUserResponseModelAllOf) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *IdentityUserResponseModelAllOf) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetSid

`func (o *IdentityUserResponseModelAllOf) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IdentityUserResponseModelAllOf) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IdentityUserResponseModelAllOf) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IdentityUserResponseModelAllOf) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetUserIdentity

`func (o *IdentityUserResponseModelAllOf) GetUserIdentity() string`

GetUserIdentity returns the UserIdentity field if non-nil, zero value otherwise.

### GetUserIdentityOk

`func (o *IdentityUserResponseModelAllOf) GetUserIdentityOk() (*string, bool)`

GetUserIdentityOk returns a tuple with the UserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentity

`func (o *IdentityUserResponseModelAllOf) SetUserIdentity(v string)`

SetUserIdentity sets UserIdentity field to given value.

### HasUserIdentity

`func (o *IdentityUserResponseModelAllOf) HasUserIdentity() bool`

HasUserIdentity returns a boolean if a field has been set.

### GetGroupSids

`func (o *IdentityUserResponseModelAllOf) GetGroupSids() []string`

GetGroupSids returns the GroupSids field if non-nil, zero value otherwise.

### GetGroupSidsOk

`func (o *IdentityUserResponseModelAllOf) GetGroupSidsOk() (*[]string, bool)`

GetGroupSidsOk returns a tuple with the GroupSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSids

`func (o *IdentityUserResponseModelAllOf) SetGroupSids(v []string)`

SetGroupSids sets GroupSids field to given value.

### HasGroupSids

`func (o *IdentityUserResponseModelAllOf) HasGroupSids() bool`

HasGroupSids returns a boolean if a field has been set.

### GetDenyOnlySids

`func (o *IdentityUserResponseModelAllOf) GetDenyOnlySids() []string`

GetDenyOnlySids returns the DenyOnlySids field if non-nil, zero value otherwise.

### GetDenyOnlySidsOk

`func (o *IdentityUserResponseModelAllOf) GetDenyOnlySidsOk() (*[]string, bool)`

GetDenyOnlySidsOk returns a tuple with the DenyOnlySids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyOnlySids

`func (o *IdentityUserResponseModelAllOf) SetDenyOnlySids(v []string)`

SetDenyOnlySids sets DenyOnlySids field to given value.

### HasDenyOnlySids

`func (o *IdentityUserResponseModelAllOf) HasDenyOnlySids() bool`

HasDenyOnlySids returns a boolean if a field has been set.

### GetPasswordCanExpire

`func (o *IdentityUserResponseModelAllOf) GetPasswordCanExpire() bool`

GetPasswordCanExpire returns the PasswordCanExpire field if non-nil, zero value otherwise.

### GetPasswordCanExpireOk

`func (o *IdentityUserResponseModelAllOf) GetPasswordCanExpireOk() (*bool, bool)`

GetPasswordCanExpireOk returns a tuple with the PasswordCanExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCanExpire

`func (o *IdentityUserResponseModelAllOf) SetPasswordCanExpire(v bool)`

SetPasswordCanExpire sets PasswordCanExpire field to given value.

### HasPasswordCanExpire

`func (o *IdentityUserResponseModelAllOf) HasPasswordCanExpire() bool`

HasPasswordCanExpire returns a boolean if a field has been set.

### GetDaysUntilPasswordExpiry

`func (o *IdentityUserResponseModelAllOf) GetDaysUntilPasswordExpiry() int64`

GetDaysUntilPasswordExpiry returns the DaysUntilPasswordExpiry field if non-nil, zero value otherwise.

### GetDaysUntilPasswordExpiryOk

`func (o *IdentityUserResponseModelAllOf) GetDaysUntilPasswordExpiryOk() (*int64, bool)`

GetDaysUntilPasswordExpiryOk returns a tuple with the DaysUntilPasswordExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilPasswordExpiry

`func (o *IdentityUserResponseModelAllOf) SetDaysUntilPasswordExpiry(v int64)`

SetDaysUntilPasswordExpiry sets DaysUntilPasswordExpiry field to given value.

### HasDaysUntilPasswordExpiry

`func (o *IdentityUserResponseModelAllOf) HasDaysUntilPasswordExpiry() bool`

HasDaysUntilPasswordExpiry returns a boolean if a field has been set.

### GetIsBuiltIn

`func (o *IdentityUserResponseModelAllOf) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *IdentityUserResponseModelAllOf) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *IdentityUserResponseModelAllOf) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.

### HasIsBuiltIn

`func (o *IdentityUserResponseModelAllOf) HasIsBuiltIn() bool`

HasIsBuiltIn returns a boolean if a field has been set.

### GetIsGroup

`func (o *IdentityUserResponseModelAllOf) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *IdentityUserResponseModelAllOf) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *IdentityUserResponseModelAllOf) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *IdentityUserResponseModelAllOf) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.

### GetName

`func (o *IdentityUserResponseModelAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityUserResponseModelAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityUserResponseModelAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityUserResponseModelAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStreetAddress

`func (o *IdentityUserResponseModelAllOf) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *IdentityUserResponseModelAllOf) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *IdentityUserResponseModelAllOf) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *IdentityUserResponseModelAllOf) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetCity

`func (o *IdentityUserResponseModelAllOf) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *IdentityUserResponseModelAllOf) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *IdentityUserResponseModelAllOf) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *IdentityUserResponseModelAllOf) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *IdentityUserResponseModelAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IdentityUserResponseModelAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IdentityUserResponseModelAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IdentityUserResponseModelAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *IdentityUserResponseModelAllOf) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IdentityUserResponseModelAllOf) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IdentityUserResponseModelAllOf) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *IdentityUserResponseModelAllOf) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetMail

`func (o *IdentityUserResponseModelAllOf) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *IdentityUserResponseModelAllOf) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *IdentityUserResponseModelAllOf) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *IdentityUserResponseModelAllOf) HasMail() bool`

HasMail returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *IdentityUserResponseModelAllOf) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *IdentityUserResponseModelAllOf) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *IdentityUserResponseModelAllOf) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *IdentityUserResponseModelAllOf) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetHomePhone

`func (o *IdentityUserResponseModelAllOf) GetHomePhone() string`

GetHomePhone returns the HomePhone field if non-nil, zero value otherwise.

### GetHomePhoneOk

`func (o *IdentityUserResponseModelAllOf) GetHomePhoneOk() (*string, bool)`

GetHomePhoneOk returns a tuple with the HomePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePhone

`func (o *IdentityUserResponseModelAllOf) SetHomePhone(v string)`

SetHomePhone sets HomePhone field to given value.

### HasHomePhone

`func (o *IdentityUserResponseModelAllOf) HasHomePhone() bool`

HasHomePhone returns a boolean if a field has been set.

### GetMobile

`func (o *IdentityUserResponseModelAllOf) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *IdentityUserResponseModelAllOf) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *IdentityUserResponseModelAllOf) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *IdentityUserResponseModelAllOf) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentityUserResponseModelAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityUserResponseModelAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityUserResponseModelAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdentityUserResponseModelAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocked

`func (o *IdentityUserResponseModelAllOf) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *IdentityUserResponseModelAllOf) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *IdentityUserResponseModelAllOf) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *IdentityUserResponseModelAllOf) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPropertiesFetched

`func (o *IdentityUserResponseModelAllOf) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *IdentityUserResponseModelAllOf) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *IdentityUserResponseModelAllOf) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.


### GetClaims

`func (o *IdentityUserResponseModelAllOf) GetClaims() IdentityUserResponseModelAllOfClaims`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *IdentityUserResponseModelAllOf) GetClaimsOk() (*IdentityUserResponseModelAllOfClaims, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *IdentityUserResponseModelAllOf) SetClaims(v IdentityUserResponseModelAllOfClaims)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *IdentityUserResponseModelAllOf) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetDirectory

`func (o *IdentityUserResponseModelAllOf) GetDirectory() IdentityUserResponseModelAllOfDirectory`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *IdentityUserResponseModelAllOf) GetDirectoryOk() (*IdentityUserResponseModelAllOfDirectory, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *IdentityUserResponseModelAllOf) SetDirectory(v IdentityUserResponseModelAllOfDirectory)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *IdentityUserResponseModelAllOf) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


