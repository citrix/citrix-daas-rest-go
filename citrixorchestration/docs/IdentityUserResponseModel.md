# IdentityUserResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | The Domain the object belongs to. NOTE: When doing a forest level search, the domain property is populated only if the object contains a SID and the SID was requested as property to be fetched. | [optional] 
**Forest** | Pointer to **NullableString** | The forest the object belongs to. | [optional] 
**Guid** | Pointer to **NullableString** | The GUID of the object, the GUID is unique across the enterprise and anywhere else. | [optional] 
**DistinguishedName** | Pointer to **NullableString** | The Distinguished name of the object. Distinguished names (DNs) are unique and they unambiguously identify objects in the directory. | [optional] 
**CanonicalName** | Pointer to **NullableString** | The unique distinguished name of the object in canonical format. | [optional] 
**PossibleLookupFailure** | Pointer to **bool** | Set to True during a bulk identity lookup operation. For individual identity lookup, usually an exception will be thrown. | [optional] 
**DirectoryServer** | Pointer to **NullableString** | Domain controller that performed the add/update operation on this object | [optional] 
**DisplayName** | **string** | Display name of the user or group. Usually something like \&quot;John Q. Doe\&quot;. | 
**SamName** | Pointer to **NullableString** | Fully qualified sAMAccountName of the User or Group object (including domain). | [optional] 
**SamAccountName** | Pointer to **NullableString** | Unqualified SAM name (Directory property) of the User or Group object.  Does not include the domain. The sAMAccountName is unique within a domain. | [optional] 
**CommonName** | Pointer to **NullableString** | The common name associated with the User or Group object. | [optional] 
**PrincipalName** | Pointer to **NullableString** | The user principal name of a User object. The UPN must be unique among all security principal objects within the directory forest. | [optional] 
**Oid** | Pointer to **NullableString** | The Object Identifier associated with the User or Group object. | [optional] 
**Sid** | Pointer to **NullableString** | The Security Identifier associated with the User or Group object. | [optional] 
**UserIdentity** | Pointer to **NullableString** | The User Identity. | [optional] 
**GroupSids** | Pointer to **[]string** | Security Identifiers of groups of which the object is a member. | [optional] 
**DenyOnlySids** | Pointer to **[]string** | Security Identifiers of resources to which the object is explicitly denied access. | [optional] 
**PasswordCanExpire** | Pointer to **NullableBool** | Indicates whether the password can expire on this User object. | [optional] 
**DaysUntilPasswordExpiry** | Pointer to **NullableInt64** | Indicates the number of days before the password can expire on this User object. | [optional] 
**IsBuiltIn** | Pointer to **NullableBool** | Indicates whether this a built-in object. | [optional] 
**IsGroup** | Pointer to **NullableBool** | Indicates whether this is a User or Group object. | [optional] 
**Name** | Pointer to **NullableString** | The name of the User or Group object. | [optional] 
**StreetAddress** | Pointer to **NullableString** | The street address associated with the User object. | [optional] 
**City** | Pointer to **NullableString** | The city associated with the User object. | [optional] 
**State** | Pointer to **NullableString** | The state associated with the User object. | [optional] 
**Country** | Pointer to **NullableString** | The country associated with the user object. | [optional] 
**Mail** | Pointer to **NullableString** | The e-mail address associated with the user object. | [optional] 
**TelephoneNumber** | Pointer to **NullableString** | The telephone number associated with the user object. | [optional] 
**HomePhone** | Pointer to **NullableString** | The home telephone number associated with the user object. | [optional] 
**Mobile** | Pointer to **NullableString** | The mobile telephone number associated with the user object. | [optional] 
**Enabled** | Pointer to **NullableBool** | Indicates whether the User object is disabled. NOTE that this is opposite of the low-level SDK, which has \&quot;IsDisabled\&quot;. By changing to \&quot;Enabled\&quot; it avoids a confusing double-negative.  It also matches all other SDK objects by using \&quot;Enabled\&quot; rather than \&quot;IsEnabled\&quot;. | [optional] 
**Locked** | Pointer to **NullableBool** | Indicates whether the User object is locked. | [optional] 
**PropertiesFetched** | **int32** | Properties fetched and populated in the User or Group object.  This is a bitfield indicating the fetched properties. | 
**Claims** | Pointer to [**IdentityClaimsResponseModel**](IdentityClaimsResponseModel.md) |  | [optional] 
**Directory** | Pointer to [**IdentityDirectoryResponseModel**](IdentityDirectoryResponseModel.md) |  | [optional] 

## Methods

### NewIdentityUserResponseModel

`func NewIdentityUserResponseModel(displayName string, propertiesFetched int32, ) *IdentityUserResponseModel`

NewIdentityUserResponseModel instantiates a new IdentityUserResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUserResponseModelWithDefaults

`func NewIdentityUserResponseModelWithDefaults() *IdentityUserResponseModel`

NewIdentityUserResponseModelWithDefaults instantiates a new IdentityUserResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *IdentityUserResponseModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentityUserResponseModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentityUserResponseModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdentityUserResponseModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *IdentityUserResponseModel) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *IdentityUserResponseModel) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetForest

`func (o *IdentityUserResponseModel) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *IdentityUserResponseModel) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *IdentityUserResponseModel) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *IdentityUserResponseModel) HasForest() bool`

HasForest returns a boolean if a field has been set.

### SetForestNil

`func (o *IdentityUserResponseModel) SetForestNil(b bool)`

 SetForestNil sets the value for Forest to be an explicit nil

### UnsetForest
`func (o *IdentityUserResponseModel) UnsetForest()`

UnsetForest ensures that no value is present for Forest, not even an explicit nil
### GetGuid

`func (o *IdentityUserResponseModel) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IdentityUserResponseModel) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IdentityUserResponseModel) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IdentityUserResponseModel) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *IdentityUserResponseModel) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *IdentityUserResponseModel) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetDistinguishedName

`func (o *IdentityUserResponseModel) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *IdentityUserResponseModel) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *IdentityUserResponseModel) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *IdentityUserResponseModel) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *IdentityUserResponseModel) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *IdentityUserResponseModel) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetCanonicalName

`func (o *IdentityUserResponseModel) GetCanonicalName() string`

GetCanonicalName returns the CanonicalName field if non-nil, zero value otherwise.

### GetCanonicalNameOk

`func (o *IdentityUserResponseModel) GetCanonicalNameOk() (*string, bool)`

GetCanonicalNameOk returns a tuple with the CanonicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalName

`func (o *IdentityUserResponseModel) SetCanonicalName(v string)`

SetCanonicalName sets CanonicalName field to given value.

### HasCanonicalName

`func (o *IdentityUserResponseModel) HasCanonicalName() bool`

HasCanonicalName returns a boolean if a field has been set.

### SetCanonicalNameNil

`func (o *IdentityUserResponseModel) SetCanonicalNameNil(b bool)`

 SetCanonicalNameNil sets the value for CanonicalName to be an explicit nil

### UnsetCanonicalName
`func (o *IdentityUserResponseModel) UnsetCanonicalName()`

UnsetCanonicalName ensures that no value is present for CanonicalName, not even an explicit nil
### GetPossibleLookupFailure

`func (o *IdentityUserResponseModel) GetPossibleLookupFailure() bool`

GetPossibleLookupFailure returns the PossibleLookupFailure field if non-nil, zero value otherwise.

### GetPossibleLookupFailureOk

`func (o *IdentityUserResponseModel) GetPossibleLookupFailureOk() (*bool, bool)`

GetPossibleLookupFailureOk returns a tuple with the PossibleLookupFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleLookupFailure

`func (o *IdentityUserResponseModel) SetPossibleLookupFailure(v bool)`

SetPossibleLookupFailure sets PossibleLookupFailure field to given value.

### HasPossibleLookupFailure

`func (o *IdentityUserResponseModel) HasPossibleLookupFailure() bool`

HasPossibleLookupFailure returns a boolean if a field has been set.

### GetDirectoryServer

`func (o *IdentityUserResponseModel) GetDirectoryServer() string`

GetDirectoryServer returns the DirectoryServer field if non-nil, zero value otherwise.

### GetDirectoryServerOk

`func (o *IdentityUserResponseModel) GetDirectoryServerOk() (*string, bool)`

GetDirectoryServerOk returns a tuple with the DirectoryServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServer

`func (o *IdentityUserResponseModel) SetDirectoryServer(v string)`

SetDirectoryServer sets DirectoryServer field to given value.

### HasDirectoryServer

`func (o *IdentityUserResponseModel) HasDirectoryServer() bool`

HasDirectoryServer returns a boolean if a field has been set.

### SetDirectoryServerNil

`func (o *IdentityUserResponseModel) SetDirectoryServerNil(b bool)`

 SetDirectoryServerNil sets the value for DirectoryServer to be an explicit nil

### UnsetDirectoryServer
`func (o *IdentityUserResponseModel) UnsetDirectoryServer()`

UnsetDirectoryServer ensures that no value is present for DirectoryServer, not even an explicit nil
### GetDisplayName

`func (o *IdentityUserResponseModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityUserResponseModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityUserResponseModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSamName

`func (o *IdentityUserResponseModel) GetSamName() string`

GetSamName returns the SamName field if non-nil, zero value otherwise.

### GetSamNameOk

`func (o *IdentityUserResponseModel) GetSamNameOk() (*string, bool)`

GetSamNameOk returns a tuple with the SamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamName

`func (o *IdentityUserResponseModel) SetSamName(v string)`

SetSamName sets SamName field to given value.

### HasSamName

`func (o *IdentityUserResponseModel) HasSamName() bool`

HasSamName returns a boolean if a field has been set.

### SetSamNameNil

`func (o *IdentityUserResponseModel) SetSamNameNil(b bool)`

 SetSamNameNil sets the value for SamName to be an explicit nil

### UnsetSamName
`func (o *IdentityUserResponseModel) UnsetSamName()`

UnsetSamName ensures that no value is present for SamName, not even an explicit nil
### GetSamAccountName

`func (o *IdentityUserResponseModel) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *IdentityUserResponseModel) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *IdentityUserResponseModel) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.

### HasSamAccountName

`func (o *IdentityUserResponseModel) HasSamAccountName() bool`

HasSamAccountName returns a boolean if a field has been set.

### SetSamAccountNameNil

`func (o *IdentityUserResponseModel) SetSamAccountNameNil(b bool)`

 SetSamAccountNameNil sets the value for SamAccountName to be an explicit nil

### UnsetSamAccountName
`func (o *IdentityUserResponseModel) UnsetSamAccountName()`

UnsetSamAccountName ensures that no value is present for SamAccountName, not even an explicit nil
### GetCommonName

`func (o *IdentityUserResponseModel) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *IdentityUserResponseModel) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *IdentityUserResponseModel) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *IdentityUserResponseModel) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *IdentityUserResponseModel) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *IdentityUserResponseModel) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetPrincipalName

`func (o *IdentityUserResponseModel) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *IdentityUserResponseModel) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *IdentityUserResponseModel) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *IdentityUserResponseModel) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *IdentityUserResponseModel) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *IdentityUserResponseModel) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
### GetOid

`func (o *IdentityUserResponseModel) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *IdentityUserResponseModel) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *IdentityUserResponseModel) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *IdentityUserResponseModel) HasOid() bool`

HasOid returns a boolean if a field has been set.

### SetOidNil

`func (o *IdentityUserResponseModel) SetOidNil(b bool)`

 SetOidNil sets the value for Oid to be an explicit nil

### UnsetOid
`func (o *IdentityUserResponseModel) UnsetOid()`

UnsetOid ensures that no value is present for Oid, not even an explicit nil
### GetSid

`func (o *IdentityUserResponseModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IdentityUserResponseModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IdentityUserResponseModel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IdentityUserResponseModel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *IdentityUserResponseModel) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *IdentityUserResponseModel) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUserIdentity

`func (o *IdentityUserResponseModel) GetUserIdentity() string`

GetUserIdentity returns the UserIdentity field if non-nil, zero value otherwise.

### GetUserIdentityOk

`func (o *IdentityUserResponseModel) GetUserIdentityOk() (*string, bool)`

GetUserIdentityOk returns a tuple with the UserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentity

`func (o *IdentityUserResponseModel) SetUserIdentity(v string)`

SetUserIdentity sets UserIdentity field to given value.

### HasUserIdentity

`func (o *IdentityUserResponseModel) HasUserIdentity() bool`

HasUserIdentity returns a boolean if a field has been set.

### SetUserIdentityNil

`func (o *IdentityUserResponseModel) SetUserIdentityNil(b bool)`

 SetUserIdentityNil sets the value for UserIdentity to be an explicit nil

### UnsetUserIdentity
`func (o *IdentityUserResponseModel) UnsetUserIdentity()`

UnsetUserIdentity ensures that no value is present for UserIdentity, not even an explicit nil
### GetGroupSids

`func (o *IdentityUserResponseModel) GetGroupSids() []string`

GetGroupSids returns the GroupSids field if non-nil, zero value otherwise.

### GetGroupSidsOk

`func (o *IdentityUserResponseModel) GetGroupSidsOk() (*[]string, bool)`

GetGroupSidsOk returns a tuple with the GroupSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSids

`func (o *IdentityUserResponseModel) SetGroupSids(v []string)`

SetGroupSids sets GroupSids field to given value.

### HasGroupSids

`func (o *IdentityUserResponseModel) HasGroupSids() bool`

HasGroupSids returns a boolean if a field has been set.

### SetGroupSidsNil

`func (o *IdentityUserResponseModel) SetGroupSidsNil(b bool)`

 SetGroupSidsNil sets the value for GroupSids to be an explicit nil

### UnsetGroupSids
`func (o *IdentityUserResponseModel) UnsetGroupSids()`

UnsetGroupSids ensures that no value is present for GroupSids, not even an explicit nil
### GetDenyOnlySids

`func (o *IdentityUserResponseModel) GetDenyOnlySids() []string`

GetDenyOnlySids returns the DenyOnlySids field if non-nil, zero value otherwise.

### GetDenyOnlySidsOk

`func (o *IdentityUserResponseModel) GetDenyOnlySidsOk() (*[]string, bool)`

GetDenyOnlySidsOk returns a tuple with the DenyOnlySids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyOnlySids

`func (o *IdentityUserResponseModel) SetDenyOnlySids(v []string)`

SetDenyOnlySids sets DenyOnlySids field to given value.

### HasDenyOnlySids

`func (o *IdentityUserResponseModel) HasDenyOnlySids() bool`

HasDenyOnlySids returns a boolean if a field has been set.

### SetDenyOnlySidsNil

`func (o *IdentityUserResponseModel) SetDenyOnlySidsNil(b bool)`

 SetDenyOnlySidsNil sets the value for DenyOnlySids to be an explicit nil

### UnsetDenyOnlySids
`func (o *IdentityUserResponseModel) UnsetDenyOnlySids()`

UnsetDenyOnlySids ensures that no value is present for DenyOnlySids, not even an explicit nil
### GetPasswordCanExpire

`func (o *IdentityUserResponseModel) GetPasswordCanExpire() bool`

GetPasswordCanExpire returns the PasswordCanExpire field if non-nil, zero value otherwise.

### GetPasswordCanExpireOk

`func (o *IdentityUserResponseModel) GetPasswordCanExpireOk() (*bool, bool)`

GetPasswordCanExpireOk returns a tuple with the PasswordCanExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCanExpire

`func (o *IdentityUserResponseModel) SetPasswordCanExpire(v bool)`

SetPasswordCanExpire sets PasswordCanExpire field to given value.

### HasPasswordCanExpire

`func (o *IdentityUserResponseModel) HasPasswordCanExpire() bool`

HasPasswordCanExpire returns a boolean if a field has been set.

### SetPasswordCanExpireNil

`func (o *IdentityUserResponseModel) SetPasswordCanExpireNil(b bool)`

 SetPasswordCanExpireNil sets the value for PasswordCanExpire to be an explicit nil

### UnsetPasswordCanExpire
`func (o *IdentityUserResponseModel) UnsetPasswordCanExpire()`

UnsetPasswordCanExpire ensures that no value is present for PasswordCanExpire, not even an explicit nil
### GetDaysUntilPasswordExpiry

`func (o *IdentityUserResponseModel) GetDaysUntilPasswordExpiry() int64`

GetDaysUntilPasswordExpiry returns the DaysUntilPasswordExpiry field if non-nil, zero value otherwise.

### GetDaysUntilPasswordExpiryOk

`func (o *IdentityUserResponseModel) GetDaysUntilPasswordExpiryOk() (*int64, bool)`

GetDaysUntilPasswordExpiryOk returns a tuple with the DaysUntilPasswordExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilPasswordExpiry

`func (o *IdentityUserResponseModel) SetDaysUntilPasswordExpiry(v int64)`

SetDaysUntilPasswordExpiry sets DaysUntilPasswordExpiry field to given value.

### HasDaysUntilPasswordExpiry

`func (o *IdentityUserResponseModel) HasDaysUntilPasswordExpiry() bool`

HasDaysUntilPasswordExpiry returns a boolean if a field has been set.

### SetDaysUntilPasswordExpiryNil

`func (o *IdentityUserResponseModel) SetDaysUntilPasswordExpiryNil(b bool)`

 SetDaysUntilPasswordExpiryNil sets the value for DaysUntilPasswordExpiry to be an explicit nil

### UnsetDaysUntilPasswordExpiry
`func (o *IdentityUserResponseModel) UnsetDaysUntilPasswordExpiry()`

UnsetDaysUntilPasswordExpiry ensures that no value is present for DaysUntilPasswordExpiry, not even an explicit nil
### GetIsBuiltIn

`func (o *IdentityUserResponseModel) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *IdentityUserResponseModel) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *IdentityUserResponseModel) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.

### HasIsBuiltIn

`func (o *IdentityUserResponseModel) HasIsBuiltIn() bool`

HasIsBuiltIn returns a boolean if a field has been set.

### SetIsBuiltInNil

`func (o *IdentityUserResponseModel) SetIsBuiltInNil(b bool)`

 SetIsBuiltInNil sets the value for IsBuiltIn to be an explicit nil

### UnsetIsBuiltIn
`func (o *IdentityUserResponseModel) UnsetIsBuiltIn()`

UnsetIsBuiltIn ensures that no value is present for IsBuiltIn, not even an explicit nil
### GetIsGroup

`func (o *IdentityUserResponseModel) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *IdentityUserResponseModel) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *IdentityUserResponseModel) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *IdentityUserResponseModel) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.

### SetIsGroupNil

`func (o *IdentityUserResponseModel) SetIsGroupNil(b bool)`

 SetIsGroupNil sets the value for IsGroup to be an explicit nil

### UnsetIsGroup
`func (o *IdentityUserResponseModel) UnsetIsGroup()`

UnsetIsGroup ensures that no value is present for IsGroup, not even an explicit nil
### GetName

`func (o *IdentityUserResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityUserResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityUserResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityUserResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IdentityUserResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdentityUserResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStreetAddress

`func (o *IdentityUserResponseModel) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *IdentityUserResponseModel) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *IdentityUserResponseModel) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *IdentityUserResponseModel) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *IdentityUserResponseModel) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *IdentityUserResponseModel) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetCity

`func (o *IdentityUserResponseModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *IdentityUserResponseModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *IdentityUserResponseModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *IdentityUserResponseModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *IdentityUserResponseModel) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *IdentityUserResponseModel) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *IdentityUserResponseModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IdentityUserResponseModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IdentityUserResponseModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IdentityUserResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *IdentityUserResponseModel) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *IdentityUserResponseModel) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *IdentityUserResponseModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IdentityUserResponseModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IdentityUserResponseModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *IdentityUserResponseModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *IdentityUserResponseModel) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *IdentityUserResponseModel) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetMail

`func (o *IdentityUserResponseModel) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *IdentityUserResponseModel) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *IdentityUserResponseModel) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *IdentityUserResponseModel) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMailNil

`func (o *IdentityUserResponseModel) SetMailNil(b bool)`

 SetMailNil sets the value for Mail to be an explicit nil

### UnsetMail
`func (o *IdentityUserResponseModel) UnsetMail()`

UnsetMail ensures that no value is present for Mail, not even an explicit nil
### GetTelephoneNumber

`func (o *IdentityUserResponseModel) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *IdentityUserResponseModel) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *IdentityUserResponseModel) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *IdentityUserResponseModel) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### SetTelephoneNumberNil

`func (o *IdentityUserResponseModel) SetTelephoneNumberNil(b bool)`

 SetTelephoneNumberNil sets the value for TelephoneNumber to be an explicit nil

### UnsetTelephoneNumber
`func (o *IdentityUserResponseModel) UnsetTelephoneNumber()`

UnsetTelephoneNumber ensures that no value is present for TelephoneNumber, not even an explicit nil
### GetHomePhone

`func (o *IdentityUserResponseModel) GetHomePhone() string`

GetHomePhone returns the HomePhone field if non-nil, zero value otherwise.

### GetHomePhoneOk

`func (o *IdentityUserResponseModel) GetHomePhoneOk() (*string, bool)`

GetHomePhoneOk returns a tuple with the HomePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePhone

`func (o *IdentityUserResponseModel) SetHomePhone(v string)`

SetHomePhone sets HomePhone field to given value.

### HasHomePhone

`func (o *IdentityUserResponseModel) HasHomePhone() bool`

HasHomePhone returns a boolean if a field has been set.

### SetHomePhoneNil

`func (o *IdentityUserResponseModel) SetHomePhoneNil(b bool)`

 SetHomePhoneNil sets the value for HomePhone to be an explicit nil

### UnsetHomePhone
`func (o *IdentityUserResponseModel) UnsetHomePhone()`

UnsetHomePhone ensures that no value is present for HomePhone, not even an explicit nil
### GetMobile

`func (o *IdentityUserResponseModel) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *IdentityUserResponseModel) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *IdentityUserResponseModel) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *IdentityUserResponseModel) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### SetMobileNil

`func (o *IdentityUserResponseModel) SetMobileNil(b bool)`

 SetMobileNil sets the value for Mobile to be an explicit nil

### UnsetMobile
`func (o *IdentityUserResponseModel) UnsetMobile()`

UnsetMobile ensures that no value is present for Mobile, not even an explicit nil
### GetEnabled

`func (o *IdentityUserResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityUserResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityUserResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdentityUserResponseModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *IdentityUserResponseModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *IdentityUserResponseModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetLocked

`func (o *IdentityUserResponseModel) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *IdentityUserResponseModel) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *IdentityUserResponseModel) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *IdentityUserResponseModel) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *IdentityUserResponseModel) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *IdentityUserResponseModel) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetPropertiesFetched

`func (o *IdentityUserResponseModel) GetPropertiesFetched() int32`

GetPropertiesFetched returns the PropertiesFetched field if non-nil, zero value otherwise.

### GetPropertiesFetchedOk

`func (o *IdentityUserResponseModel) GetPropertiesFetchedOk() (*int32, bool)`

GetPropertiesFetchedOk returns a tuple with the PropertiesFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesFetched

`func (o *IdentityUserResponseModel) SetPropertiesFetched(v int32)`

SetPropertiesFetched sets PropertiesFetched field to given value.


### GetClaims

`func (o *IdentityUserResponseModel) GetClaims() IdentityClaimsResponseModel`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *IdentityUserResponseModel) GetClaimsOk() (*IdentityClaimsResponseModel, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *IdentityUserResponseModel) SetClaims(v IdentityClaimsResponseModel)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *IdentityUserResponseModel) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetDirectory

`func (o *IdentityUserResponseModel) GetDirectory() IdentityDirectoryResponseModel`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *IdentityUserResponseModel) GetDirectoryOk() (*IdentityDirectoryResponseModel, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *IdentityUserResponseModel) SetDirectory(v IdentityDirectoryResponseModel)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *IdentityUserResponseModel) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


