# EnrollmentRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenName** | **string** | Gets or sets the name used to more easily identify this VDA enrollment token. | 
**IssuedToUser** | Pointer to **NullableString** | Gets or sets the name of the user to whom this VDA enrollment token meta data object is issued. | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | Gets or sets the date this VDA enrollment token expires. | [optional] 
**NotValidBeforeDate** | Pointer to **NullableTime** | Gets or sets the UTC date before which this VDA enrollment token will not be valid. | [optional] 
**NumMachinesAllowed** | Pointer to **int32** | Gets or sets the number of machines this VDA enrollment token allows to be added to the catalog. | [optional] 
**CatalogId** | **string** | Gets or sets the Catalog Id that this VDA enrollment token allows machines to be added to. | 
**HostConnectionId** | Pointer to **NullableString** | Gets or sets the HostConnectionId  which will be  used to determine how to handle power management. | [optional] 

## Methods

### NewEnrollmentRequestModel

`func NewEnrollmentRequestModel(tokenName string, catalogId string, ) *EnrollmentRequestModel`

NewEnrollmentRequestModel instantiates a new EnrollmentRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentRequestModelWithDefaults

`func NewEnrollmentRequestModelWithDefaults() *EnrollmentRequestModel`

NewEnrollmentRequestModelWithDefaults instantiates a new EnrollmentRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenName

`func (o *EnrollmentRequestModel) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *EnrollmentRequestModel) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *EnrollmentRequestModel) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetIssuedToUser

`func (o *EnrollmentRequestModel) GetIssuedToUser() string`

GetIssuedToUser returns the IssuedToUser field if non-nil, zero value otherwise.

### GetIssuedToUserOk

`func (o *EnrollmentRequestModel) GetIssuedToUserOk() (*string, bool)`

GetIssuedToUserOk returns a tuple with the IssuedToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedToUser

`func (o *EnrollmentRequestModel) SetIssuedToUser(v string)`

SetIssuedToUser sets IssuedToUser field to given value.

### HasIssuedToUser

`func (o *EnrollmentRequestModel) HasIssuedToUser() bool`

HasIssuedToUser returns a boolean if a field has been set.

### SetIssuedToUserNil

`func (o *EnrollmentRequestModel) SetIssuedToUserNil(b bool)`

 SetIssuedToUserNil sets the value for IssuedToUser to be an explicit nil

### UnsetIssuedToUser
`func (o *EnrollmentRequestModel) UnsetIssuedToUser()`

UnsetIssuedToUser ensures that no value is present for IssuedToUser, not even an explicit nil
### GetExpirationDate

`func (o *EnrollmentRequestModel) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *EnrollmentRequestModel) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *EnrollmentRequestModel) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *EnrollmentRequestModel) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *EnrollmentRequestModel) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *EnrollmentRequestModel) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetNotValidBeforeDate

`func (o *EnrollmentRequestModel) GetNotValidBeforeDate() time.Time`

GetNotValidBeforeDate returns the NotValidBeforeDate field if non-nil, zero value otherwise.

### GetNotValidBeforeDateOk

`func (o *EnrollmentRequestModel) GetNotValidBeforeDateOk() (*time.Time, bool)`

GetNotValidBeforeDateOk returns a tuple with the NotValidBeforeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotValidBeforeDate

`func (o *EnrollmentRequestModel) SetNotValidBeforeDate(v time.Time)`

SetNotValidBeforeDate sets NotValidBeforeDate field to given value.

### HasNotValidBeforeDate

`func (o *EnrollmentRequestModel) HasNotValidBeforeDate() bool`

HasNotValidBeforeDate returns a boolean if a field has been set.

### SetNotValidBeforeDateNil

`func (o *EnrollmentRequestModel) SetNotValidBeforeDateNil(b bool)`

 SetNotValidBeforeDateNil sets the value for NotValidBeforeDate to be an explicit nil

### UnsetNotValidBeforeDate
`func (o *EnrollmentRequestModel) UnsetNotValidBeforeDate()`

UnsetNotValidBeforeDate ensures that no value is present for NotValidBeforeDate, not even an explicit nil
### GetNumMachinesAllowed

`func (o *EnrollmentRequestModel) GetNumMachinesAllowed() int32`

GetNumMachinesAllowed returns the NumMachinesAllowed field if non-nil, zero value otherwise.

### GetNumMachinesAllowedOk

`func (o *EnrollmentRequestModel) GetNumMachinesAllowedOk() (*int32, bool)`

GetNumMachinesAllowedOk returns a tuple with the NumMachinesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachinesAllowed

`func (o *EnrollmentRequestModel) SetNumMachinesAllowed(v int32)`

SetNumMachinesAllowed sets NumMachinesAllowed field to given value.

### HasNumMachinesAllowed

`func (o *EnrollmentRequestModel) HasNumMachinesAllowed() bool`

HasNumMachinesAllowed returns a boolean if a field has been set.

### GetCatalogId

`func (o *EnrollmentRequestModel) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *EnrollmentRequestModel) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *EnrollmentRequestModel) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.


### GetHostConnectionId

`func (o *EnrollmentRequestModel) GetHostConnectionId() string`

GetHostConnectionId returns the HostConnectionId field if non-nil, zero value otherwise.

### GetHostConnectionIdOk

`func (o *EnrollmentRequestModel) GetHostConnectionIdOk() (*string, bool)`

GetHostConnectionIdOk returns a tuple with the HostConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostConnectionId

`func (o *EnrollmentRequestModel) SetHostConnectionId(v string)`

SetHostConnectionId sets HostConnectionId field to given value.

### HasHostConnectionId

`func (o *EnrollmentRequestModel) HasHostConnectionId() bool`

HasHostConnectionId returns a boolean if a field has been set.

### SetHostConnectionIdNil

`func (o *EnrollmentRequestModel) SetHostConnectionIdNil(b bool)`

 SetHostConnectionIdNil sets the value for HostConnectionId to be an explicit nil

### UnsetHostConnectionId
`func (o *EnrollmentRequestModel) UnsetHostConnectionId()`

UnsetHostConnectionId ensures that no value is present for HostConnectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


