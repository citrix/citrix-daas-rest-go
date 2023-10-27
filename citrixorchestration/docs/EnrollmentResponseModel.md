# EnrollmentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Gets or sets the unique id of this VDA enrollment token meta data object. | [optional] 
**TokenName** | Pointer to **NullableString** | Gets or sets the name used to more easily identify this VDA enrollment token. | [optional] 
**IssuerId** | Pointer to **NullableString** | Gets or sets the unique id of the admin who created this VDA enrollment token meta data object. | [optional] 
**IssuedToUser** | Pointer to **NullableString** | Gets or sets the name of the user to whom this VDA enrollment token meta data object is issued. | [optional] 
**DateIssued** | Pointer to **NullableString** | Gets or sets the date this VDA enrollment token meta data object was created. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | Gets or sets the date this VDA enrollment token expires. | [optional] 
**NotValidBeforeDate** | Pointer to **NullableString** | Gets or sets the UTC date before which this VDA enrollment token will not be valid. | [optional] 
**NumMachinesAllowed** | Pointer to **int32** | Gets or sets the number of machines this VDA enrollment token allows to be added to the catalog. | [optional] 
**CatalogId** | Pointer to **NullableString** | Gets or sets the Catalog Id that this VDA enrollment token allows machines to be added to. | [optional] 
**IsRevoked** | Pointer to **bool** | Gets or sets the flag indicating that this VDA enrollment token has been revoked. | [optional] 
**LastUsed** | Pointer to **NullableString** | Gets or sets the date this VDA enrollment token was last used to add a machine to a catalog. | [optional] 
**NumTimesUsed** | Pointer to **int32** | Gets or sets the number of times this token was used to attempt to add machines to the specified catalog. | [optional] 
**Token** | Pointer to **NullableString** | The JWT VDA enrollment token created from the  data properties. | [optional] 
**HostConnectionId** | Pointer to **NullableString** | Gets or sets the HostConnectionId  which will be  used to determine how to handle power management. | [optional] 

## Methods

### NewEnrollmentResponseModel

`func NewEnrollmentResponseModel() *EnrollmentResponseModel`

NewEnrollmentResponseModel instantiates a new EnrollmentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentResponseModelWithDefaults

`func NewEnrollmentResponseModelWithDefaults() *EnrollmentResponseModel`

NewEnrollmentResponseModelWithDefaults instantiates a new EnrollmentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnrollmentResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnrollmentResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnrollmentResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnrollmentResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTokenName

`func (o *EnrollmentResponseModel) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *EnrollmentResponseModel) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *EnrollmentResponseModel) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *EnrollmentResponseModel) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### SetTokenNameNil

`func (o *EnrollmentResponseModel) SetTokenNameNil(b bool)`

 SetTokenNameNil sets the value for TokenName to be an explicit nil

### UnsetTokenName
`func (o *EnrollmentResponseModel) UnsetTokenName()`

UnsetTokenName ensures that no value is present for TokenName, not even an explicit nil
### GetIssuerId

`func (o *EnrollmentResponseModel) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *EnrollmentResponseModel) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *EnrollmentResponseModel) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.

### HasIssuerId

`func (o *EnrollmentResponseModel) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.

### SetIssuerIdNil

`func (o *EnrollmentResponseModel) SetIssuerIdNil(b bool)`

 SetIssuerIdNil sets the value for IssuerId to be an explicit nil

### UnsetIssuerId
`func (o *EnrollmentResponseModel) UnsetIssuerId()`

UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
### GetIssuedToUser

`func (o *EnrollmentResponseModel) GetIssuedToUser() string`

GetIssuedToUser returns the IssuedToUser field if non-nil, zero value otherwise.

### GetIssuedToUserOk

`func (o *EnrollmentResponseModel) GetIssuedToUserOk() (*string, bool)`

GetIssuedToUserOk returns a tuple with the IssuedToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedToUser

`func (o *EnrollmentResponseModel) SetIssuedToUser(v string)`

SetIssuedToUser sets IssuedToUser field to given value.

### HasIssuedToUser

`func (o *EnrollmentResponseModel) HasIssuedToUser() bool`

HasIssuedToUser returns a boolean if a field has been set.

### SetIssuedToUserNil

`func (o *EnrollmentResponseModel) SetIssuedToUserNil(b bool)`

 SetIssuedToUserNil sets the value for IssuedToUser to be an explicit nil

### UnsetIssuedToUser
`func (o *EnrollmentResponseModel) UnsetIssuedToUser()`

UnsetIssuedToUser ensures that no value is present for IssuedToUser, not even an explicit nil
### GetDateIssued

`func (o *EnrollmentResponseModel) GetDateIssued() string`

GetDateIssued returns the DateIssued field if non-nil, zero value otherwise.

### GetDateIssuedOk

`func (o *EnrollmentResponseModel) GetDateIssuedOk() (*string, bool)`

GetDateIssuedOk returns a tuple with the DateIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateIssued

`func (o *EnrollmentResponseModel) SetDateIssued(v string)`

SetDateIssued sets DateIssued field to given value.

### HasDateIssued

`func (o *EnrollmentResponseModel) HasDateIssued() bool`

HasDateIssued returns a boolean if a field has been set.

### SetDateIssuedNil

`func (o *EnrollmentResponseModel) SetDateIssuedNil(b bool)`

 SetDateIssuedNil sets the value for DateIssued to be an explicit nil

### UnsetDateIssued
`func (o *EnrollmentResponseModel) UnsetDateIssued()`

UnsetDateIssued ensures that no value is present for DateIssued, not even an explicit nil
### GetExpirationDate

`func (o *EnrollmentResponseModel) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *EnrollmentResponseModel) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *EnrollmentResponseModel) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *EnrollmentResponseModel) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *EnrollmentResponseModel) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *EnrollmentResponseModel) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetNotValidBeforeDate

`func (o *EnrollmentResponseModel) GetNotValidBeforeDate() string`

GetNotValidBeforeDate returns the NotValidBeforeDate field if non-nil, zero value otherwise.

### GetNotValidBeforeDateOk

`func (o *EnrollmentResponseModel) GetNotValidBeforeDateOk() (*string, bool)`

GetNotValidBeforeDateOk returns a tuple with the NotValidBeforeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotValidBeforeDate

`func (o *EnrollmentResponseModel) SetNotValidBeforeDate(v string)`

SetNotValidBeforeDate sets NotValidBeforeDate field to given value.

### HasNotValidBeforeDate

`func (o *EnrollmentResponseModel) HasNotValidBeforeDate() bool`

HasNotValidBeforeDate returns a boolean if a field has been set.

### SetNotValidBeforeDateNil

`func (o *EnrollmentResponseModel) SetNotValidBeforeDateNil(b bool)`

 SetNotValidBeforeDateNil sets the value for NotValidBeforeDate to be an explicit nil

### UnsetNotValidBeforeDate
`func (o *EnrollmentResponseModel) UnsetNotValidBeforeDate()`

UnsetNotValidBeforeDate ensures that no value is present for NotValidBeforeDate, not even an explicit nil
### GetNumMachinesAllowed

`func (o *EnrollmentResponseModel) GetNumMachinesAllowed() int32`

GetNumMachinesAllowed returns the NumMachinesAllowed field if non-nil, zero value otherwise.

### GetNumMachinesAllowedOk

`func (o *EnrollmentResponseModel) GetNumMachinesAllowedOk() (*int32, bool)`

GetNumMachinesAllowedOk returns a tuple with the NumMachinesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachinesAllowed

`func (o *EnrollmentResponseModel) SetNumMachinesAllowed(v int32)`

SetNumMachinesAllowed sets NumMachinesAllowed field to given value.

### HasNumMachinesAllowed

`func (o *EnrollmentResponseModel) HasNumMachinesAllowed() bool`

HasNumMachinesAllowed returns a boolean if a field has been set.

### GetCatalogId

`func (o *EnrollmentResponseModel) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *EnrollmentResponseModel) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *EnrollmentResponseModel) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *EnrollmentResponseModel) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### SetCatalogIdNil

`func (o *EnrollmentResponseModel) SetCatalogIdNil(b bool)`

 SetCatalogIdNil sets the value for CatalogId to be an explicit nil

### UnsetCatalogId
`func (o *EnrollmentResponseModel) UnsetCatalogId()`

UnsetCatalogId ensures that no value is present for CatalogId, not even an explicit nil
### GetIsRevoked

`func (o *EnrollmentResponseModel) GetIsRevoked() bool`

GetIsRevoked returns the IsRevoked field if non-nil, zero value otherwise.

### GetIsRevokedOk

`func (o *EnrollmentResponseModel) GetIsRevokedOk() (*bool, bool)`

GetIsRevokedOk returns a tuple with the IsRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRevoked

`func (o *EnrollmentResponseModel) SetIsRevoked(v bool)`

SetIsRevoked sets IsRevoked field to given value.

### HasIsRevoked

`func (o *EnrollmentResponseModel) HasIsRevoked() bool`

HasIsRevoked returns a boolean if a field has been set.

### GetLastUsed

`func (o *EnrollmentResponseModel) GetLastUsed() string`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *EnrollmentResponseModel) GetLastUsedOk() (*string, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *EnrollmentResponseModel) SetLastUsed(v string)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *EnrollmentResponseModel) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *EnrollmentResponseModel) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *EnrollmentResponseModel) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetNumTimesUsed

`func (o *EnrollmentResponseModel) GetNumTimesUsed() int32`

GetNumTimesUsed returns the NumTimesUsed field if non-nil, zero value otherwise.

### GetNumTimesUsedOk

`func (o *EnrollmentResponseModel) GetNumTimesUsedOk() (*int32, bool)`

GetNumTimesUsedOk returns a tuple with the NumTimesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTimesUsed

`func (o *EnrollmentResponseModel) SetNumTimesUsed(v int32)`

SetNumTimesUsed sets NumTimesUsed field to given value.

### HasNumTimesUsed

`func (o *EnrollmentResponseModel) HasNumTimesUsed() bool`

HasNumTimesUsed returns a boolean if a field has been set.

### GetToken

`func (o *EnrollmentResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EnrollmentResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EnrollmentResponseModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EnrollmentResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *EnrollmentResponseModel) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *EnrollmentResponseModel) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetHostConnectionId

`func (o *EnrollmentResponseModel) GetHostConnectionId() string`

GetHostConnectionId returns the HostConnectionId field if non-nil, zero value otherwise.

### GetHostConnectionIdOk

`func (o *EnrollmentResponseModel) GetHostConnectionIdOk() (*string, bool)`

GetHostConnectionIdOk returns a tuple with the HostConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostConnectionId

`func (o *EnrollmentResponseModel) SetHostConnectionId(v string)`

SetHostConnectionId sets HostConnectionId field to given value.

### HasHostConnectionId

`func (o *EnrollmentResponseModel) HasHostConnectionId() bool`

HasHostConnectionId returns a boolean if a field has been set.

### SetHostConnectionIdNil

`func (o *EnrollmentResponseModel) SetHostConnectionIdNil(b bool)`

 SetHostConnectionIdNil sets the value for HostConnectionId to be an explicit nil

### UnsetHostConnectionId
`func (o *EnrollmentResponseModel) UnsetHostConnectionId()`

UnsetHostConnectionId ensures that no value is present for HostConnectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


