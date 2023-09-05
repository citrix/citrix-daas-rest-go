# SessionResponseModelUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The Domain the object belongs to. NOTE: When doing a forest level search, the domain property is populated only if the object contains a SID and the SID was requested as property to be fetched. | [optional] 
**Forest** | Pointer to **string** | The forest the object belongs to. | [optional] 
**Guid** | Pointer to **string** | The GUID of the object, the GUID is unique across the enterprise and anywhere else. | [optional] 
**DistinguishedName** | Pointer to **string** | The Distinguished name of the object. Distinguished names (DNs) are unique and they unambiguously identify objects in the directory. | [optional] 
**CanonicalName** | Pointer to **string** | The unique distinguished name of the object in canonical format. | [optional] 
**PossibleLookupFailure** | Pointer to **bool** | Set to True during a bulk identity lookup operation. For individual identity lookup, usually an exception will be thrown. | [optional] 
**DirectoryServer** | Pointer to **string** | Domain controller that performed the add/update operation on this object | [optional] 
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

### NewSessionResponseModelUser

`func NewSessionResponseModelUser(displayName string, propertiesFetched int32, ) *SessionResponseModelUser`

NewSessionResponseModelUser instantiates a new SessionResponseModelUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionResponseModelUserWithDefaults

`func NewSessionResponseModelUserWithDefaults() *SessionResponseModelUser`

NewSessionResponseModelUserWithDefaults instantiates a new SessionResponseModelUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SessionResponseModelUser) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SessionResponseModelUser) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SessionResponseModelUser) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SessionResponseModelUser) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetForest

`func (o *SessionResponseModelUser) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *SessionResponseModelUser) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *SessionResponseModelUser) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *SessionResponseModelUser) HasForest() bool`

HasForest returns a boolean if a field has been set.

### GetGuid

`func (o *SessionResponseModelUser) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SessionResponseModelUser) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SessionResponseModelUser) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *SessionResponseModelUser) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *SessionResponseModelUser) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *SessionResponseModelUser) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *SessionResponseModelUser) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *SessionResponseModelUser) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetCanonicalName

`func (o *SessionResponseModelUser) GetCanonicalName() string`

GetCanonicalName returns the CanonicalName field if non-nil, zero value otherwise.

### GetCanonicalNameOk

`func (o *SessionResponseModelUser) GetCanonicalNameOk() (*string, bool)`

GetCanonicalNameOk returns a tuple with the CanonicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalName

`func (o *SessionResponseModelUser) SetCanonicalName(v string)`

SetCanonicalName sets CanonicalName field to given value.

### HasCanonicalName

`func (o *SessionResponseModelUser) HasCanonicalName() bool`

HasCanonicalName returns a boolean if a field has been set.

### GetPossibleLookupFailure

`func (o *SessionResponseModelUser) GetPossibleLookupFailure() bool`

GetPossibleLookupFailure returns the PossibleLookupFailure field if non-nil, zero value otherwise.

### GetPossibleLookupFailureOk

`func (o *SessionResponseModelUser) GetPossibleLookupFailureOk() (*bool, bool)`

GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleLookupFailure

`func (o *SessionResponseModelUser) SetPossibleLookupFailure(v bool)`

SetPossibleLookupFailure sets PossibleLookupFailure field to given value.

### HasPossibleLookupFailure

`func (o *SessionResponseModelUser) HasPossibleLookupFailure() bool`

HasPossibleLookupFailure returns a boolean if a field has been set.

### GetDirectoryServer

`func (o *SessionResponseModelUser) GetDirectoryServer() string`

GetDirectoryServer returns the DirectoryServer field if non-nil, zero value otherwise.

### GetDirectoryServerOk

`func (o *SessionResponseModelUser) GetDirectoryServerOk() (*string, bool)`

GetDirectoryServerOk returns a tuple with the DirectoryServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServer

`func (o *SessionResponseModelUser) SetDirectoryServer(v string)`

SetDirectoryServer sets DirectoryServer field to given value.

### HasDirectoryServer

`func (o *SessionResponseModelUser) HasDirectoryServer() bool`

HasDirectoryServer returns a boolean if a field has been set.

### GetDisplayName

`func (o *SessionResponseModelUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SessionResponseModelUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SessionResponseModelUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSamName

`func (o *SessionResponseModelUser) GetSamName() string`

GetSamName returns the SamName field if non-nil, zero value otherwise.

### GetSamNameOk

`func (o *SessionResponseModelUser) GetSamNameOk() (*string, bool)`

GetSamNameOk returns a tuple with the SamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamName

`func (o *SessionResponseModelUser) SetSamName(v string)`

SetSamName sets SamName field to given value.

### HasSamName

`func (o *SessionResponseModelUser) HasSamName() bool`

HasSamName returns a boolean if a field has been set.

### GetSamAccountName

`func (o *SessionResponseModelUser) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *SessionResponseModelUser) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *SessionResponseModelUser) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.

### HasSamAccountName

`func (o *SessionResponseModelUser) HasSamAccountName() bool`

HasSamAccountName returns a boolean if a field has been set.

### GetCommonName

`func (o *SessionResponseModelUser) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *SessionResponseModelUser) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *SessionResponseModelUser) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *SessionResponseModelUser) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetPrincipalName

`func (o *SessionResponseModelUser) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *SessionResponseModelUser) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *SessionResponseModelUser) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *SessionResponseModelUser) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### GetOid

`func (o *SessionResponseModelUser) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *SessionResponseModelUser) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *SessionResponseModelUser) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *SessionResponseModelUser) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetSid

`func (o *SessionResponseModelUser) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SessionResponseModelUser) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SessionResponseModelUser) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SessionResponseModelUser) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetUserIdentity

`func (o *SessionResponseModelUser) GetUserIdentity() string`

GetUserIdentity returns the UserIdentity field if non-nil, zero value otherwise.

### GetUserIdentityOk

`func (o *SessionResponseModelUser) GetUserIdentityOk() (*string, bool)`

GetUserIdentityOk returns a tuple with the UserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentity

`func (o *SessionResponseModelUser) SetUserIdentity(v string)`

SetUserIdentity sets UserIdentity field to given value.

### HasUserIdentity

`func (o *SessionResponseModelUser) HasUserIdentity() bool`

HasUserIdentity returns a boolean if a field has been set.

### GetGroupSids

`func (o *SessionResponseModelUser) GetGroupSids() []string`

GetGroupSids returns the GroupSids field if non-nil, zero value otherwise.

### GetGroupSidsOk

`func (o *SessionResponseModelUser) GetGroupSidsOk() (*[]string, bool)`

GetGroupSidsOk returns a tuple with the GroupSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSids

`func (o *SessionResponseModelUser) SetGroupSids(v []string)`

SetGroupSids sets GroupSids field to given value.

### HasGroupSids

`func (o *SessionResponseModelUser) HasGroupSids() bool`

HasGroupSids returns a boolean if a field has been set.

### GetDenyOnlySids

`func (o *SessionResponseModelUser) GetDenyOnlySids() []string`

GetDenyOnlySids returns the DenyOnlySids field if non-nil, zero value otherwise.

### GetDenyOnlySidsOk

`func (o *SessionResponseModelUser) GetDenyOnlySidsOk() (*[]string, bool)`

GetDenyOnlySidsOk returns a tuple with the DenyOnlySids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyOnlySids

`func (o *SessionResponseModelUser) SetDenyOnlySids(v []string)`

SetDenyOnlySids sets DenyOnlySids field to given value.

### HasDenyOnlySids

`func (o *SessionResponseModelUser) HasDenyOnlySids() bool`

HasDenyOnlySids returns a boolean if a field has been set.

### GetPasswordCanExpire

`func (o *SessionResponseModelUser) GetPasswordCanExpire() bool`

GetPasswordCanExpire returns the PasswordCanExpire field if non-nil, zero value otherwise.

### GetPasswordCanExpireOk

`func (o *SessionResponseModelUser) GetPasswordCanExpireOk() (*bool, bool)`

GetPasswordCanExpireOk returns a tuple with the PasswordCanExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCanExpire

`func (o *SessionResponseModelUser) SetPasswordCanExpire(v bool)`

SetPasswordCanExpire sets PasswordCanExpire field to given value.

### HasPasswordCanExpire

`func (o *SessionResponseModelUser) HasPasswordCanExpire() bool`

HasPasswordCanExpire returns a boolean if a field has been set.

### GetDaysUntilPasswordExpiry

`func (o *SessionResponseModelUser) GetDaysUntilPasswordExpiry() int64`

GetDaysUntilPasswordExpiry returns the DaysUntilPasswordExpiry field if non-nil, zero value otherwise.

### GetDaysUntilPasswordExpiryOk

`func (o *SessionResponseModelUser) GetDaysUntilPasswordExpiryOk() (*int64, bool)`

GetDaysUntilPasswordExpiryOk returns a tuple with the DaysUntilPasswordExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilPasswordExpiry

`func (o *SessionResponseModelUser) SetDaysUntilPasswordExpiry(v int64)`

SetDaysUntilPasswordExpiry sets DaysUntilPasswordExpiry field to given value.

### HasDaysUntilPasswordExpiry

`func (o *SessionResponseModelUser) HasDaysUntilPasswordExpiry() bool`

HasDaysUntilPasswordExpiry returns a boolean if a field has been set.

### GetIsBuiltIn

`func (o *SessionResponseModelUser) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *SessionResponseModelUser) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *SessionResponseModelUser) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.

### HasIsBuiltIn

`func (o *SessionResponseModelUser) HasIsBuiltIn() bool`

HasIsBuiltIn returns a boolean if a field has been set.

### GetIsGroup

`func (o *SessionResponseModelUser) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *SessionResponseModelUser) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *SessionResponseModelUser) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *SessionResponseModelUser) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.

### GetName

`func (o *SessionResponseModelUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionResponseModelUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionResponseModelUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SessionResponseModelUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStreetAddress

`func (o *SessionResponseModelUser) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *SessionResponseModelUser) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *SessionResponseModelUser) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *SessionResponseModelUser) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetCity

`func (o *SessionResponseModelUser) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SessionResponseModelUser) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SessionResponseModelUser) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SessionResponseModelUser) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *SessionResponseModelUser) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SessionResponseModelUser) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SessionResponseModelUser) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SessionResponseModelUser) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *SessionResponseModelUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SessionResponseModelUser) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SessionResponseModelUser) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SessionResponseModelUser) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetMail

`func (o *SessionResponseModelUser) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *SessionResponseModelUser) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *SessionResponseModelUser) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *SessionResponseModelUser) HasMail() bool`

HasMail returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *SessionResponseModelUser) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *SessionResponseModelUser) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *SessionResponseModelUser) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *SessionResponseModelUser) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetHomePhone

`func (o *SessionResponseModelUser) GetHomePhone() string`

GetHomePhone returns the HomePhone field if non-nil, zero value otherwise.

### GetHomePhoneOk

`func (o *SessionResponseModelUser) GetHomePhoneOk() (*string, bool)`

GetHomePhoneOk returns a tuple with the HomePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePhone

`func (o *SessionResponseModelUser) SetHomePhone(v string)`

SetHomePhone sets HomePhone field to given value.

### HasHomePhone

`func (o *SessionResponseModelUser) HasHomePhone() bool`

HasHomePhone returns a boolean if a field has been set.

### GetMobile

`func (o *SessionResponseModelUser) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *SessionResponseModelUser) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *SessionResponseModelUser) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *SessionResponseModelUser) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetEnabled

`func (o *SessionResponseModelUser) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SessionResponseModelUser) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SessionResponseModelUser) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SessionResponseModelUser) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocked

`func (o *SessionResponseModelUser) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SessionResponseModelUser) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SessionResponseModelUser) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SessionResponseModelUser) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPropertiesFetched

`func (o *SessionResponseModelUser) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *SessionResponseModelUser) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *SessionResponseModelUser) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.


### GetClaims

`func (o *SessionResponseModelUser) GetClaims() IdentityUserResponseModelAllOfClaims`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *SessionResponseModelUser) GetClaimsOk() (*IdentityUserResponseModelAllOfClaims, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *SessionResponseModelUser) SetClaims(v IdentityUserResponseModelAllOfClaims)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *SessionResponseModelUser) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetDirectory

`func (o *SessionResponseModelUser) GetDirectory() IdentityUserResponseModelAllOfDirectory`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *SessionResponseModelUser) GetDirectoryOk() (*IdentityUserResponseModelAllOfDirectory, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *SessionResponseModelUser) SetDirectory(v IdentityUserResponseModelAllOfDirectory)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *SessionResponseModelUser) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


