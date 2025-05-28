# ObjectChangeEventModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ChangeType** | Pointer to [**ChangeEventsChangeType**](ChangeEventsChangeType.md) |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**ServiceProfile** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**Identity** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Geo** | Pointer to **string** |  | [optional] 
**ResourceProvider** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**BeforeChange** | Pointer to **string** |  | [optional] 
**AfterChange** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **time.Time** |  | [optional] 
**CallerServiceIdentity** | Pointer to **string** |  | [optional] 
**CallerPrincipal** | Pointer to **string** |  | [optional] 
**CallerServiceInstanceId** | Pointer to **string** |  | [optional] 
**SystemLogCorrelationId** | Pointer to **string** |  | [optional] 
**Oid** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Sid** | Pointer to **string** |  | [optional] 
**Upn** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


