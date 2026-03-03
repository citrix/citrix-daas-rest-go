# ObjectChangeEventModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ChangeType** | Pointer to [**ChangeEventsChangeType**](ChangeEventsChangeType.md) |  | [optional] 
**Identifier** | Pointer to **NullableString** |  | [optional] 
**ServiceProfile** | Pointer to **NullableString** |  | [optional] 
**TransactionId** | Pointer to **NullableString** |  | [optional] 
**Identity** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**CustomerGeo** | Pointer to **NullableString** |  | [optional] 
**Geo** | Pointer to **NullableString** |  | [optional] 
**ResourceProvider** | Pointer to **NullableString** |  | [optional] 
**Service** | Pointer to **NullableString** |  | [optional] 
**BeforeChange** | Pointer to **NullableString** |  | [optional] 
**AfterChange** | Pointer to **NullableString** |  | [optional] 
**TimeStamp** | Pointer to **NullableTime** |  | [optional] 
**CallerServiceIdentity** | Pointer to **NullableString** |  | [optional] 
**CallerPrincipal** | Pointer to **NullableString** |  | [optional] 
**CallerServiceInstanceId** | Pointer to **NullableString** |  | [optional] 
**SystemLogCorrelationId** | Pointer to **NullableString** |  | [optional] 
**Oid** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Sid** | Pointer to **NullableString** |  | [optional] 
**Upn** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewObjectChangeEventModel

`func NewObjectChangeEventModel() *ObjectChangeEventModel`

NewObjectChangeEventModel instantiates a new ObjectChangeEventModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectChangeEventModelWithDefaults

`func NewObjectChangeEventModelWithDefaults() *ObjectChangeEventModel`

NewObjectChangeEventModelWithDefaults instantiates a new ObjectChangeEventModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *ObjectChangeEventModel) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ObjectChangeEventModel) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ObjectChangeEventModel) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ObjectChangeEventModel) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *ObjectChangeEventModel) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *ObjectChangeEventModel) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetType

`func (o *ObjectChangeEventModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectChangeEventModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectChangeEventModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ObjectChangeEventModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ObjectChangeEventModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ObjectChangeEventModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetChangeType

`func (o *ObjectChangeEventModel) GetChangeType() ChangeEventsChangeType`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *ObjectChangeEventModel) GetChangeTypeOk() (*ChangeEventsChangeType, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *ObjectChangeEventModel) SetChangeType(v ChangeEventsChangeType)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *ObjectChangeEventModel) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetIdentifier

`func (o *ObjectChangeEventModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ObjectChangeEventModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ObjectChangeEventModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ObjectChangeEventModel) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *ObjectChangeEventModel) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *ObjectChangeEventModel) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetServiceProfile

`func (o *ObjectChangeEventModel) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *ObjectChangeEventModel) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *ObjectChangeEventModel) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *ObjectChangeEventModel) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### SetServiceProfileNil

`func (o *ObjectChangeEventModel) SetServiceProfileNil(b bool)`

 SetServiceProfileNil sets the value for ServiceProfile to be an explicit nil

### UnsetServiceProfile
`func (o *ObjectChangeEventModel) UnsetServiceProfile()`

UnsetServiceProfile ensures that no value is present for ServiceProfile, not even an explicit nil
### GetTransactionId

`func (o *ObjectChangeEventModel) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ObjectChangeEventModel) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ObjectChangeEventModel) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ObjectChangeEventModel) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *ObjectChangeEventModel) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *ObjectChangeEventModel) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetIdentity

`func (o *ObjectChangeEventModel) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ObjectChangeEventModel) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ObjectChangeEventModel) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ObjectChangeEventModel) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *ObjectChangeEventModel) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *ObjectChangeEventModel) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetUserId

`func (o *ObjectChangeEventModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ObjectChangeEventModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ObjectChangeEventModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ObjectChangeEventModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ObjectChangeEventModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ObjectChangeEventModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetCustomerGeo

`func (o *ObjectChangeEventModel) GetCustomerGeo() string`

GetCustomerGeo returns the CustomerGeo field if non-nil, zero value otherwise.

### GetCustomerGeoOk

`func (o *ObjectChangeEventModel) GetCustomerGeoOk() (*string, bool)`

GetCustomerGeoOk returns a tuple with the CustomerGeo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGeo

`func (o *ObjectChangeEventModel) SetCustomerGeo(v string)`

SetCustomerGeo sets CustomerGeo field to given value.

### HasCustomerGeo

`func (o *ObjectChangeEventModel) HasCustomerGeo() bool`

HasCustomerGeo returns a boolean if a field has been set.

### SetCustomerGeoNil

`func (o *ObjectChangeEventModel) SetCustomerGeoNil(b bool)`

 SetCustomerGeoNil sets the value for CustomerGeo to be an explicit nil

### UnsetCustomerGeo
`func (o *ObjectChangeEventModel) UnsetCustomerGeo()`

UnsetCustomerGeo ensures that no value is present for CustomerGeo, not even an explicit nil
### GetGeo

`func (o *ObjectChangeEventModel) GetGeo() string`

GetGeo returns the Geo field if non-nil, zero value otherwise.

### GetGeoOk

`func (o *ObjectChangeEventModel) GetGeoOk() (*string, bool)`

GetGeoOk returns a tuple with the Geo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeo

`func (o *ObjectChangeEventModel) SetGeo(v string)`

SetGeo sets Geo field to given value.

### HasGeo

`func (o *ObjectChangeEventModel) HasGeo() bool`

HasGeo returns a boolean if a field has been set.

### SetGeoNil

`func (o *ObjectChangeEventModel) SetGeoNil(b bool)`

 SetGeoNil sets the value for Geo to be an explicit nil

### UnsetGeo
`func (o *ObjectChangeEventModel) UnsetGeo()`

UnsetGeo ensures that no value is present for Geo, not even an explicit nil
### GetResourceProvider

`func (o *ObjectChangeEventModel) GetResourceProvider() string`

GetResourceProvider returns the ResourceProvider field if non-nil, zero value otherwise.

### GetResourceProviderOk

`func (o *ObjectChangeEventModel) GetResourceProviderOk() (*string, bool)`

GetResourceProviderOk returns a tuple with the ResourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProvider

`func (o *ObjectChangeEventModel) SetResourceProvider(v string)`

SetResourceProvider sets ResourceProvider field to given value.

### HasResourceProvider

`func (o *ObjectChangeEventModel) HasResourceProvider() bool`

HasResourceProvider returns a boolean if a field has been set.

### SetResourceProviderNil

`func (o *ObjectChangeEventModel) SetResourceProviderNil(b bool)`

 SetResourceProviderNil sets the value for ResourceProvider to be an explicit nil

### UnsetResourceProvider
`func (o *ObjectChangeEventModel) UnsetResourceProvider()`

UnsetResourceProvider ensures that no value is present for ResourceProvider, not even an explicit nil
### GetService

`func (o *ObjectChangeEventModel) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ObjectChangeEventModel) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ObjectChangeEventModel) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ObjectChangeEventModel) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *ObjectChangeEventModel) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *ObjectChangeEventModel) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetBeforeChange

`func (o *ObjectChangeEventModel) GetBeforeChange() string`

GetBeforeChange returns the BeforeChange field if non-nil, zero value otherwise.

### GetBeforeChangeOk

`func (o *ObjectChangeEventModel) GetBeforeChangeOk() (*string, bool)`

GetBeforeChangeOk returns a tuple with the BeforeChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeChange

`func (o *ObjectChangeEventModel) SetBeforeChange(v string)`

SetBeforeChange sets BeforeChange field to given value.

### HasBeforeChange

`func (o *ObjectChangeEventModel) HasBeforeChange() bool`

HasBeforeChange returns a boolean if a field has been set.

### SetBeforeChangeNil

`func (o *ObjectChangeEventModel) SetBeforeChangeNil(b bool)`

 SetBeforeChangeNil sets the value for BeforeChange to be an explicit nil

### UnsetBeforeChange
`func (o *ObjectChangeEventModel) UnsetBeforeChange()`

UnsetBeforeChange ensures that no value is present for BeforeChange, not even an explicit nil
### GetAfterChange

`func (o *ObjectChangeEventModel) GetAfterChange() string`

GetAfterChange returns the AfterChange field if non-nil, zero value otherwise.

### GetAfterChangeOk

`func (o *ObjectChangeEventModel) GetAfterChangeOk() (*string, bool)`

GetAfterChangeOk returns a tuple with the AfterChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterChange

`func (o *ObjectChangeEventModel) SetAfterChange(v string)`

SetAfterChange sets AfterChange field to given value.

### HasAfterChange

`func (o *ObjectChangeEventModel) HasAfterChange() bool`

HasAfterChange returns a boolean if a field has been set.

### SetAfterChangeNil

`func (o *ObjectChangeEventModel) SetAfterChangeNil(b bool)`

 SetAfterChangeNil sets the value for AfterChange to be an explicit nil

### UnsetAfterChange
`func (o *ObjectChangeEventModel) UnsetAfterChange()`

UnsetAfterChange ensures that no value is present for AfterChange, not even an explicit nil
### GetTimeStamp

`func (o *ObjectChangeEventModel) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ObjectChangeEventModel) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ObjectChangeEventModel) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *ObjectChangeEventModel) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### SetTimeStampNil

`func (o *ObjectChangeEventModel) SetTimeStampNil(b bool)`

 SetTimeStampNil sets the value for TimeStamp to be an explicit nil

### UnsetTimeStamp
`func (o *ObjectChangeEventModel) UnsetTimeStamp()`

UnsetTimeStamp ensures that no value is present for TimeStamp, not even an explicit nil
### GetCallerServiceIdentity

`func (o *ObjectChangeEventModel) GetCallerServiceIdentity() string`

GetCallerServiceIdentity returns the CallerServiceIdentity field if non-nil, zero value otherwise.

### GetCallerServiceIdentityOk

`func (o *ObjectChangeEventModel) GetCallerServiceIdentityOk() (*string, bool)`

GetCallerServiceIdentityOk returns a tuple with the CallerServiceIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerServiceIdentity

`func (o *ObjectChangeEventModel) SetCallerServiceIdentity(v string)`

SetCallerServiceIdentity sets CallerServiceIdentity field to given value.

### HasCallerServiceIdentity

`func (o *ObjectChangeEventModel) HasCallerServiceIdentity() bool`

HasCallerServiceIdentity returns a boolean if a field has been set.

### SetCallerServiceIdentityNil

`func (o *ObjectChangeEventModel) SetCallerServiceIdentityNil(b bool)`

 SetCallerServiceIdentityNil sets the value for CallerServiceIdentity to be an explicit nil

### UnsetCallerServiceIdentity
`func (o *ObjectChangeEventModel) UnsetCallerServiceIdentity()`

UnsetCallerServiceIdentity ensures that no value is present for CallerServiceIdentity, not even an explicit nil
### GetCallerPrincipal

`func (o *ObjectChangeEventModel) GetCallerPrincipal() string`

GetCallerPrincipal returns the CallerPrincipal field if non-nil, zero value otherwise.

### GetCallerPrincipalOk

`func (o *ObjectChangeEventModel) GetCallerPrincipalOk() (*string, bool)`

GetCallerPrincipalOk returns a tuple with the CallerPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerPrincipal

`func (o *ObjectChangeEventModel) SetCallerPrincipal(v string)`

SetCallerPrincipal sets CallerPrincipal field to given value.

### HasCallerPrincipal

`func (o *ObjectChangeEventModel) HasCallerPrincipal() bool`

HasCallerPrincipal returns a boolean if a field has been set.

### SetCallerPrincipalNil

`func (o *ObjectChangeEventModel) SetCallerPrincipalNil(b bool)`

 SetCallerPrincipalNil sets the value for CallerPrincipal to be an explicit nil

### UnsetCallerPrincipal
`func (o *ObjectChangeEventModel) UnsetCallerPrincipal()`

UnsetCallerPrincipal ensures that no value is present for CallerPrincipal, not even an explicit nil
### GetCallerServiceInstanceId

`func (o *ObjectChangeEventModel) GetCallerServiceInstanceId() string`

GetCallerServiceInstanceId returns the CallerServiceInstanceId field if non-nil, zero value otherwise.

### GetCallerServiceInstanceIdOk

`func (o *ObjectChangeEventModel) GetCallerServiceInstanceIdOk() (*string, bool)`

GetCallerServiceInstanceIdOk returns a tuple with the CallerServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerServiceInstanceId

`func (o *ObjectChangeEventModel) SetCallerServiceInstanceId(v string)`

SetCallerServiceInstanceId sets CallerServiceInstanceId field to given value.

### HasCallerServiceInstanceId

`func (o *ObjectChangeEventModel) HasCallerServiceInstanceId() bool`

HasCallerServiceInstanceId returns a boolean if a field has been set.

### SetCallerServiceInstanceIdNil

`func (o *ObjectChangeEventModel) SetCallerServiceInstanceIdNil(b bool)`

 SetCallerServiceInstanceIdNil sets the value for CallerServiceInstanceId to be an explicit nil

### UnsetCallerServiceInstanceId
`func (o *ObjectChangeEventModel) UnsetCallerServiceInstanceId()`

UnsetCallerServiceInstanceId ensures that no value is present for CallerServiceInstanceId, not even an explicit nil
### GetSystemLogCorrelationId

`func (o *ObjectChangeEventModel) GetSystemLogCorrelationId() string`

GetSystemLogCorrelationId returns the SystemLogCorrelationId field if non-nil, zero value otherwise.

### GetSystemLogCorrelationIdOk

`func (o *ObjectChangeEventModel) GetSystemLogCorrelationIdOk() (*string, bool)`

GetSystemLogCorrelationIdOk returns a tuple with the SystemLogCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLogCorrelationId

`func (o *ObjectChangeEventModel) SetSystemLogCorrelationId(v string)`

SetSystemLogCorrelationId sets SystemLogCorrelationId field to given value.

### HasSystemLogCorrelationId

`func (o *ObjectChangeEventModel) HasSystemLogCorrelationId() bool`

HasSystemLogCorrelationId returns a boolean if a field has been set.

### SetSystemLogCorrelationIdNil

`func (o *ObjectChangeEventModel) SetSystemLogCorrelationIdNil(b bool)`

 SetSystemLogCorrelationIdNil sets the value for SystemLogCorrelationId to be an explicit nil

### UnsetSystemLogCorrelationId
`func (o *ObjectChangeEventModel) UnsetSystemLogCorrelationId()`

UnsetSystemLogCorrelationId ensures that no value is present for SystemLogCorrelationId, not even an explicit nil
### GetOid

`func (o *ObjectChangeEventModel) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *ObjectChangeEventModel) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *ObjectChangeEventModel) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *ObjectChangeEventModel) HasOid() bool`

HasOid returns a boolean if a field has been set.

### SetOidNil

`func (o *ObjectChangeEventModel) SetOidNil(b bool)`

 SetOidNil sets the value for Oid to be an explicit nil

### UnsetOid
`func (o *ObjectChangeEventModel) UnsetOid()`

UnsetOid ensures that no value is present for Oid, not even an explicit nil
### GetEmail

`func (o *ObjectChangeEventModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ObjectChangeEventModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ObjectChangeEventModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ObjectChangeEventModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ObjectChangeEventModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ObjectChangeEventModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDisplayName

`func (o *ObjectChangeEventModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ObjectChangeEventModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ObjectChangeEventModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ObjectChangeEventModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ObjectChangeEventModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ObjectChangeEventModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSid

`func (o *ObjectChangeEventModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ObjectChangeEventModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ObjectChangeEventModel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ObjectChangeEventModel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ObjectChangeEventModel) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ObjectChangeEventModel) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUpn

`func (o *ObjectChangeEventModel) GetUpn() string`

GetUpn returns the Upn field if non-nil, zero value otherwise.

### GetUpnOk

`func (o *ObjectChangeEventModel) GetUpnOk() (*string, bool)`

GetUpnOk returns a tuple with the Upn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpn

`func (o *ObjectChangeEventModel) SetUpn(v string)`

SetUpn sets Upn field to given value.

### HasUpn

`func (o *ObjectChangeEventModel) HasUpn() bool`

HasUpn returns a boolean if a field has been set.

### SetUpnNil

`func (o *ObjectChangeEventModel) SetUpnNil(b bool)`

 SetUpnNil sets the value for Upn to be an explicit nil

### UnsetUpn
`func (o *ObjectChangeEventModel) UnsetUpn()`

UnsetUpn ensures that no value is present for Upn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


