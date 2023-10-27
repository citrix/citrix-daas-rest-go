# SessionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the session. | [optional] 
**Uid** | Pointer to **NullableInt32** | &#x60;Deprecated, use Id.&#x60; DEPRECATED. Use Id. | [optional] 
**ApplicationsInUse** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | List of applications in use in the session. | [optional] 
**AppState** | Pointer to [**AppState**](AppState.md) |  | [optional] 
**AppStateLastChangeTime** | Pointer to **NullableString** | The time when the session entered the current app state. | [optional] 
**FormattedAppStateLastChangeTime** | Pointer to **NullableString** | The formatted time when the session entered the current app state. RFC 3339 compatible format. | [optional] 
**Brokering** | Pointer to [**SessionBrokeringResponseModel**](SessionBrokeringResponseModel.md) |  | [optional] 
**Client** | Pointer to [**SessionClientResponseModel**](SessionClientResponseModel.md) |  | [optional] 
**Connection** | Pointer to [**SessionConnectionResponseModel**](SessionConnectionResponseModel.md) |  | [optional] 
**ContainerScopes** | Pointer to [**[]ContainerScopeResponseModel**](ContainerScopeResponseModel.md) | Delegated admin scopes in which the containers of the session reside. | [optional] 
**EstablishmentDurationMilliseconds** | Pointer to **NullableInt32** | Duration that it took to establish the session. | [optional] 
**EstablishmentTime** | Pointer to **NullableString** | Time at which the session was established. | [optional] 
**FormattedEstablishmentTime** | Pointer to **NullableString** | Formatted time at which the session was established. RFC 3339 compatible format. | [optional] 
**IsAnonymousUser** | Pointer to **bool** | Indicates whether the session was established anonymously (without user credentials), in this case a temporary local user account on the machine is used. | [optional] 
**IsHidden** | Pointer to **bool** | Flag to indicate if the session is currently hidden from the user and not to be reconnected to. | [optional] 
**LogoffInProgress** | Pointer to **bool** | Indicates whether the session is in the process of being logged off. | [optional] 
**LogonInProgress** | Pointer to **bool** | Indicates whether the session is still executing user logon processing or not. | [optional] 
**Machine** | Pointer to [**MachineResponseModel**](MachineResponseModel.md) |  | [optional] 
**SessionType** | Pointer to [**SessionType**](SessionType.md) |  | [optional] 
**StartTime** | Pointer to **NullableString** | The time indicates when the session was started. | [optional] 
**FormattedStartTime** | Pointer to **NullableString** | The formatted time indicates when the session was started. RFC 3339 compatible format. | [optional] 
**State** | Pointer to [**SessionState**](SessionState.md) |  | [optional] 
**StateChangeTime** | Pointer to **NullableString** | The time of the most recent state change for the session. | [optional] 
**FormattedStateChangeTime** | Pointer to **NullableString** | The formatted time of the most recent state change for the session. RFC 3339 compatible format. | [optional] 
**UntrustedUserName** | Pointer to **NullableString** | The name of the logged-on user reported directly from the machine (in the form DOMAIN\\user). This may be useful where the user is logged in to a non-domain account, however the name cannot be verified and must therefore be considered untrusted. | [optional] 
**User** | Pointer to [**IdentityUserResponseModel**](IdentityUserResponseModel.md) |  | [optional] 

## Methods

### NewSessionResponseModel

`func NewSessionResponseModel() *SessionResponseModel`

NewSessionResponseModel instantiates a new SessionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionResponseModelWithDefaults

`func NewSessionResponseModelWithDefaults() *SessionResponseModel`

NewSessionResponseModelWithDefaults instantiates a new SessionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SessionResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SessionResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *SessionResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SessionResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SessionResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SessionResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *SessionResponseModel) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *SessionResponseModel) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetApplicationsInUse

`func (o *SessionResponseModel) GetApplicationsInUse() []RefResponseModel`

GetApplicationsInUse returns the ApplicationsInUse field if non-nil, zero value otherwise.

### GetApplicationsInUseOk

`func (o *SessionResponseModel) GetApplicationsInUseOk() (*[]RefResponseModel, bool)`

GetApplicationsInUseOk returns a tuple with the ApplicationsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsInUse

`func (o *SessionResponseModel) SetApplicationsInUse(v []RefResponseModel)`

SetApplicationsInUse sets ApplicationsInUse field to given value.

### HasApplicationsInUse

`func (o *SessionResponseModel) HasApplicationsInUse() bool`

HasApplicationsInUse returns a boolean if a field has been set.

### SetApplicationsInUseNil

`func (o *SessionResponseModel) SetApplicationsInUseNil(b bool)`

 SetApplicationsInUseNil sets the value for ApplicationsInUse to be an explicit nil

### UnsetApplicationsInUse
`func (o *SessionResponseModel) UnsetApplicationsInUse()`

UnsetApplicationsInUse ensures that no value is present for ApplicationsInUse, not even an explicit nil
### GetAppState

`func (o *SessionResponseModel) GetAppState() AppState`

GetAppState returns the AppState field if non-nil, zero value otherwise.

### GetAppStateOk

`func (o *SessionResponseModel) GetAppStateOk() (*AppState, bool)`

GetAppStateOk returns a tuple with the AppState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppState

`func (o *SessionResponseModel) SetAppState(v AppState)`

SetAppState sets AppState field to given value.

### HasAppState

`func (o *SessionResponseModel) HasAppState() bool`

HasAppState returns a boolean if a field has been set.

### GetAppStateLastChangeTime

`func (o *SessionResponseModel) GetAppStateLastChangeTime() string`

GetAppStateLastChangeTime returns the AppStateLastChangeTime field if non-nil, zero value otherwise.

### GetAppStateLastChangeTimeOk

`func (o *SessionResponseModel) GetAppStateLastChangeTimeOk() (*string, bool)`

GetAppStateLastChangeTimeOk returns a tuple with the AppStateLastChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStateLastChangeTime

`func (o *SessionResponseModel) SetAppStateLastChangeTime(v string)`

SetAppStateLastChangeTime sets AppStateLastChangeTime field to given value.

### HasAppStateLastChangeTime

`func (o *SessionResponseModel) HasAppStateLastChangeTime() bool`

HasAppStateLastChangeTime returns a boolean if a field has been set.

### SetAppStateLastChangeTimeNil

`func (o *SessionResponseModel) SetAppStateLastChangeTimeNil(b bool)`

 SetAppStateLastChangeTimeNil sets the value for AppStateLastChangeTime to be an explicit nil

### UnsetAppStateLastChangeTime
`func (o *SessionResponseModel) UnsetAppStateLastChangeTime()`

UnsetAppStateLastChangeTime ensures that no value is present for AppStateLastChangeTime, not even an explicit nil
### GetFormattedAppStateLastChangeTime

`func (o *SessionResponseModel) GetFormattedAppStateLastChangeTime() string`

GetFormattedAppStateLastChangeTime returns the FormattedAppStateLastChangeTime field if non-nil, zero value otherwise.

### GetFormattedAppStateLastChangeTimeOk

`func (o *SessionResponseModel) GetFormattedAppStateLastChangeTimeOk() (*string, bool)`

GetFormattedAppStateLastChangeTimeOk returns a tuple with the FormattedAppStateLastChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAppStateLastChangeTime

`func (o *SessionResponseModel) SetFormattedAppStateLastChangeTime(v string)`

SetFormattedAppStateLastChangeTime sets FormattedAppStateLastChangeTime field to given value.

### HasFormattedAppStateLastChangeTime

`func (o *SessionResponseModel) HasFormattedAppStateLastChangeTime() bool`

HasFormattedAppStateLastChangeTime returns a boolean if a field has been set.

### SetFormattedAppStateLastChangeTimeNil

`func (o *SessionResponseModel) SetFormattedAppStateLastChangeTimeNil(b bool)`

 SetFormattedAppStateLastChangeTimeNil sets the value for FormattedAppStateLastChangeTime to be an explicit nil

### UnsetFormattedAppStateLastChangeTime
`func (o *SessionResponseModel) UnsetFormattedAppStateLastChangeTime()`

UnsetFormattedAppStateLastChangeTime ensures that no value is present for FormattedAppStateLastChangeTime, not even an explicit nil
### GetBrokering

`func (o *SessionResponseModel) GetBrokering() SessionBrokeringResponseModel`

GetBrokering returns the Brokering field if non-nil, zero value otherwise.

### GetBrokeringOk

`func (o *SessionResponseModel) GetBrokeringOk() (*SessionBrokeringResponseModel, bool)`

GetBrokeringOk returns a tuple with the Brokering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokering

`func (o *SessionResponseModel) SetBrokering(v SessionBrokeringResponseModel)`

SetBrokering sets Brokering field to given value.

### HasBrokering

`func (o *SessionResponseModel) HasBrokering() bool`

HasBrokering returns a boolean if a field has been set.

### GetClient

`func (o *SessionResponseModel) GetClient() SessionClientResponseModel`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *SessionResponseModel) GetClientOk() (*SessionClientResponseModel, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *SessionResponseModel) SetClient(v SessionClientResponseModel)`

SetClient sets Client field to given value.

### HasClient

`func (o *SessionResponseModel) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetConnection

`func (o *SessionResponseModel) GetConnection() SessionConnectionResponseModel`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *SessionResponseModel) GetConnectionOk() (*SessionConnectionResponseModel, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *SessionResponseModel) SetConnection(v SessionConnectionResponseModel)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *SessionResponseModel) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetContainerScopes

`func (o *SessionResponseModel) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *SessionResponseModel) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *SessionResponseModel) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.

### HasContainerScopes

`func (o *SessionResponseModel) HasContainerScopes() bool`

HasContainerScopes returns a boolean if a field has been set.

### SetContainerScopesNil

`func (o *SessionResponseModel) SetContainerScopesNil(b bool)`

 SetContainerScopesNil sets the value for ContainerScopes to be an explicit nil

### UnsetContainerScopes
`func (o *SessionResponseModel) UnsetContainerScopes()`

UnsetContainerScopes ensures that no value is present for ContainerScopes, not even an explicit nil
### GetEstablishmentDurationMilliseconds

`func (o *SessionResponseModel) GetEstablishmentDurationMilliseconds() int32`

GetEstablishmentDurationMilliseconds returns the EstablishmentDurationMilliseconds field if non-nil, zero value otherwise.

### GetEstablishmentDurationMillisecondsOk

`func (o *SessionResponseModel) GetEstablishmentDurationMillisecondsOk() (*int32, bool)`

GetEstablishmentDurationMillisecondsOk returns a tuple with the EstablishmentDurationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablishmentDurationMilliseconds

`func (o *SessionResponseModel) SetEstablishmentDurationMilliseconds(v int32)`

SetEstablishmentDurationMilliseconds sets EstablishmentDurationMilliseconds field to given value.

### HasEstablishmentDurationMilliseconds

`func (o *SessionResponseModel) HasEstablishmentDurationMilliseconds() bool`

HasEstablishmentDurationMilliseconds returns a boolean if a field has been set.

### SetEstablishmentDurationMillisecondsNil

`func (o *SessionResponseModel) SetEstablishmentDurationMillisecondsNil(b bool)`

 SetEstablishmentDurationMillisecondsNil sets the value for EstablishmentDurationMilliseconds to be an explicit nil

### UnsetEstablishmentDurationMilliseconds
`func (o *SessionResponseModel) UnsetEstablishmentDurationMilliseconds()`

UnsetEstablishmentDurationMilliseconds ensures that no value is present for EstablishmentDurationMilliseconds, not even an explicit nil
### GetEstablishmentTime

`func (o *SessionResponseModel) GetEstablishmentTime() string`

GetEstablishmentTime returns the EstablishmentTime field if non-nil, zero value otherwise.

### GetEstablishmentTimeOk

`func (o *SessionResponseModel) GetEstablishmentTimeOk() (*string, bool)`

GetEstablishmentTimeOk returns a tuple with the EstablishmentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablishmentTime

`func (o *SessionResponseModel) SetEstablishmentTime(v string)`

SetEstablishmentTime sets EstablishmentTime field to given value.

### HasEstablishmentTime

`func (o *SessionResponseModel) HasEstablishmentTime() bool`

HasEstablishmentTime returns a boolean if a field has been set.

### SetEstablishmentTimeNil

`func (o *SessionResponseModel) SetEstablishmentTimeNil(b bool)`

 SetEstablishmentTimeNil sets the value for EstablishmentTime to be an explicit nil

### UnsetEstablishmentTime
`func (o *SessionResponseModel) UnsetEstablishmentTime()`

UnsetEstablishmentTime ensures that no value is present for EstablishmentTime, not even an explicit nil
### GetFormattedEstablishmentTime

`func (o *SessionResponseModel) GetFormattedEstablishmentTime() string`

GetFormattedEstablishmentTime returns the FormattedEstablishmentTime field if non-nil, zero value otherwise.

### GetFormattedEstablishmentTimeOk

`func (o *SessionResponseModel) GetFormattedEstablishmentTimeOk() (*string, bool)`

GetFormattedEstablishmentTimeOk returns a tuple with the FormattedEstablishmentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedEstablishmentTime

`func (o *SessionResponseModel) SetFormattedEstablishmentTime(v string)`

SetFormattedEstablishmentTime sets FormattedEstablishmentTime field to given value.

### HasFormattedEstablishmentTime

`func (o *SessionResponseModel) HasFormattedEstablishmentTime() bool`

HasFormattedEstablishmentTime returns a boolean if a field has been set.

### SetFormattedEstablishmentTimeNil

`func (o *SessionResponseModel) SetFormattedEstablishmentTimeNil(b bool)`

 SetFormattedEstablishmentTimeNil sets the value for FormattedEstablishmentTime to be an explicit nil

### UnsetFormattedEstablishmentTime
`func (o *SessionResponseModel) UnsetFormattedEstablishmentTime()`

UnsetFormattedEstablishmentTime ensures that no value is present for FormattedEstablishmentTime, not even an explicit nil
### GetIsAnonymousUser

`func (o *SessionResponseModel) GetIsAnonymousUser() bool`

GetIsAnonymousUser returns the IsAnonymousUser field if non-nil, zero value otherwise.

### GetIsAnonymousUserOk

`func (o *SessionResponseModel) GetIsAnonymousUserOk() (*bool, bool)`

GetIsAnonymousUserOk returns a tuple with the IsAnonymousUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymousUser

`func (o *SessionResponseModel) SetIsAnonymousUser(v bool)`

SetIsAnonymousUser sets IsAnonymousUser field to given value.

### HasIsAnonymousUser

`func (o *SessionResponseModel) HasIsAnonymousUser() bool`

HasIsAnonymousUser returns a boolean if a field has been set.

### GetIsHidden

`func (o *SessionResponseModel) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *SessionResponseModel) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *SessionResponseModel) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *SessionResponseModel) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetLogoffInProgress

`func (o *SessionResponseModel) GetLogoffInProgress() bool`

GetLogoffInProgress returns the LogoffInProgress field if non-nil, zero value otherwise.

### GetLogoffInProgressOk

`func (o *SessionResponseModel) GetLogoffInProgressOk() (*bool, bool)`

GetLogoffInProgressOk returns a tuple with the LogoffInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffInProgress

`func (o *SessionResponseModel) SetLogoffInProgress(v bool)`

SetLogoffInProgress sets LogoffInProgress field to given value.

### HasLogoffInProgress

`func (o *SessionResponseModel) HasLogoffInProgress() bool`

HasLogoffInProgress returns a boolean if a field has been set.

### GetLogonInProgress

`func (o *SessionResponseModel) GetLogonInProgress() bool`

GetLogonInProgress returns the LogonInProgress field if non-nil, zero value otherwise.

### GetLogonInProgressOk

`func (o *SessionResponseModel) GetLogonInProgressOk() (*bool, bool)`

GetLogonInProgressOk returns a tuple with the LogonInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonInProgress

`func (o *SessionResponseModel) SetLogonInProgress(v bool)`

SetLogonInProgress sets LogonInProgress field to given value.

### HasLogonInProgress

`func (o *SessionResponseModel) HasLogonInProgress() bool`

HasLogonInProgress returns a boolean if a field has been set.

### GetMachine

`func (o *SessionResponseModel) GetMachine() MachineResponseModel`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *SessionResponseModel) GetMachineOk() (*MachineResponseModel, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *SessionResponseModel) SetMachine(v MachineResponseModel)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *SessionResponseModel) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetSessionType

`func (o *SessionResponseModel) GetSessionType() SessionType`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *SessionResponseModel) GetSessionTypeOk() (*SessionType, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *SessionResponseModel) SetSessionType(v SessionType)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *SessionResponseModel) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetStartTime

`func (o *SessionResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SessionResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SessionResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SessionResponseModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *SessionResponseModel) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *SessionResponseModel) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetFormattedStartTime

`func (o *SessionResponseModel) GetFormattedStartTime() string`

GetFormattedStartTime returns the FormattedStartTime field if non-nil, zero value otherwise.

### GetFormattedStartTimeOk

`func (o *SessionResponseModel) GetFormattedStartTimeOk() (*string, bool)`

GetFormattedStartTimeOk returns a tuple with the FormattedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedStartTime

`func (o *SessionResponseModel) SetFormattedStartTime(v string)`

SetFormattedStartTime sets FormattedStartTime field to given value.

### HasFormattedStartTime

`func (o *SessionResponseModel) HasFormattedStartTime() bool`

HasFormattedStartTime returns a boolean if a field has been set.

### SetFormattedStartTimeNil

`func (o *SessionResponseModel) SetFormattedStartTimeNil(b bool)`

 SetFormattedStartTimeNil sets the value for FormattedStartTime to be an explicit nil

### UnsetFormattedStartTime
`func (o *SessionResponseModel) UnsetFormattedStartTime()`

UnsetFormattedStartTime ensures that no value is present for FormattedStartTime, not even an explicit nil
### GetState

`func (o *SessionResponseModel) GetState() SessionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SessionResponseModel) GetStateOk() (*SessionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SessionResponseModel) SetState(v SessionState)`

SetState sets State field to given value.

### HasState

`func (o *SessionResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateChangeTime

`func (o *SessionResponseModel) GetStateChangeTime() string`

GetStateChangeTime returns the StateChangeTime field if non-nil, zero value otherwise.

### GetStateChangeTimeOk

`func (o *SessionResponseModel) GetStateChangeTimeOk() (*string, bool)`

GetStateChangeTimeOk returns a tuple with the StateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeTime

`func (o *SessionResponseModel) SetStateChangeTime(v string)`

SetStateChangeTime sets StateChangeTime field to given value.

### HasStateChangeTime

`func (o *SessionResponseModel) HasStateChangeTime() bool`

HasStateChangeTime returns a boolean if a field has been set.

### SetStateChangeTimeNil

`func (o *SessionResponseModel) SetStateChangeTimeNil(b bool)`

 SetStateChangeTimeNil sets the value for StateChangeTime to be an explicit nil

### UnsetStateChangeTime
`func (o *SessionResponseModel) UnsetStateChangeTime()`

UnsetStateChangeTime ensures that no value is present for StateChangeTime, not even an explicit nil
### GetFormattedStateChangeTime

`func (o *SessionResponseModel) GetFormattedStateChangeTime() string`

GetFormattedStateChangeTime returns the FormattedStateChangeTime field if non-nil, zero value otherwise.

### GetFormattedStateChangeTimeOk

`func (o *SessionResponseModel) GetFormattedStateChangeTimeOk() (*string, bool)`

GetFormattedStateChangeTimeOk returns a tuple with the FormattedStateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedStateChangeTime

`func (o *SessionResponseModel) SetFormattedStateChangeTime(v string)`

SetFormattedStateChangeTime sets FormattedStateChangeTime field to given value.

### HasFormattedStateChangeTime

`func (o *SessionResponseModel) HasFormattedStateChangeTime() bool`

HasFormattedStateChangeTime returns a boolean if a field has been set.

### SetFormattedStateChangeTimeNil

`func (o *SessionResponseModel) SetFormattedStateChangeTimeNil(b bool)`

 SetFormattedStateChangeTimeNil sets the value for FormattedStateChangeTime to be an explicit nil

### UnsetFormattedStateChangeTime
`func (o *SessionResponseModel) UnsetFormattedStateChangeTime()`

UnsetFormattedStateChangeTime ensures that no value is present for FormattedStateChangeTime, not even an explicit nil
### GetUntrustedUserName

`func (o *SessionResponseModel) GetUntrustedUserName() string`

GetUntrustedUserName returns the UntrustedUserName field if non-nil, zero value otherwise.

### GetUntrustedUserNameOk

`func (o *SessionResponseModel) GetUntrustedUserNameOk() (*string, bool)`

GetUntrustedUserNameOk returns a tuple with the UntrustedUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntrustedUserName

`func (o *SessionResponseModel) SetUntrustedUserName(v string)`

SetUntrustedUserName sets UntrustedUserName field to given value.

### HasUntrustedUserName

`func (o *SessionResponseModel) HasUntrustedUserName() bool`

HasUntrustedUserName returns a boolean if a field has been set.

### SetUntrustedUserNameNil

`func (o *SessionResponseModel) SetUntrustedUserNameNil(b bool)`

 SetUntrustedUserNameNil sets the value for UntrustedUserName to be an explicit nil

### UnsetUntrustedUserName
`func (o *SessionResponseModel) UnsetUntrustedUserName()`

UnsetUntrustedUserName ensures that no value is present for UntrustedUserName, not even an explicit nil
### GetUser

`func (o *SessionResponseModel) GetUser() IdentityUserResponseModel`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SessionResponseModel) GetUserOk() (*IdentityUserResponseModel, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SessionResponseModel) SetUser(v IdentityUserResponseModel)`

SetUser sets User field to given value.

### HasUser

`func (o *SessionResponseModel) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


