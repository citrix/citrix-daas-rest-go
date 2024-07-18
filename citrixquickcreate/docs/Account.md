# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**AccountId** | Pointer to **NullableString** | ID of the account | [optional] 
**Name** | Pointer to **NullableString** | Name of the account | [optional] 
**AccountState** | Pointer to [**AccountState**](AccountState.md) |  | [optional] 
**AssociatedDeployments** | Pointer to [**[]AssociatedDeployment**](AssociatedDeployment.md) |  | [optional] 
**AssociatedResourceConnections** | Pointer to [**[]AssociatedResourceConnection**](AssociatedResourceConnection.md) |  | [optional] 
**AssociatedImages** | Pointer to [**[]AssociatedImage**](AssociatedImage.md) |  | [optional] 
**CitrixManaged** | Pointer to **bool** | Indicates whether the account is Citrix managed | [optional] 
**ErrorMessage** | Pointer to **NullableString** | The error message for account | [optional] 
**Warnings** | Pointer to [**[]AccountWarning**](AccountWarning.md) | Warnings and errors associated with the account | [optional] 

## Methods

### NewAccount

`func NewAccount(accountType AccountType, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *Account) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Account) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Account) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetAccountId

`func (o *Account) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Account) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Account) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Account) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *Account) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Account) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Account) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Account) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Account) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccountState

`func (o *Account) GetAccountState() AccountState`

GetAccountState returns the AccountState field if non-nil, zero value otherwise.

### GetAccountStateOk

`func (o *Account) GetAccountStateOk() (*AccountState, bool)`

GetAccountStateOk returns a tuple with the AccountState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountState

`func (o *Account) SetAccountState(v AccountState)`

SetAccountState sets AccountState field to given value.

### HasAccountState

`func (o *Account) HasAccountState() bool`

HasAccountState returns a boolean if a field has been set.

### GetAssociatedDeployments

`func (o *Account) GetAssociatedDeployments() []AssociatedDeployment`

GetAssociatedDeployments returns the AssociatedDeployments field if non-nil, zero value otherwise.

### GetAssociatedDeploymentsOk

`func (o *Account) GetAssociatedDeploymentsOk() (*[]AssociatedDeployment, bool)`

GetAssociatedDeploymentsOk returns a tuple with the AssociatedDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDeployments

`func (o *Account) SetAssociatedDeployments(v []AssociatedDeployment)`

SetAssociatedDeployments sets AssociatedDeployments field to given value.

### HasAssociatedDeployments

`func (o *Account) HasAssociatedDeployments() bool`

HasAssociatedDeployments returns a boolean if a field has been set.

### SetAssociatedDeploymentsNil

`func (o *Account) SetAssociatedDeploymentsNil(b bool)`

 SetAssociatedDeploymentsNil sets the value for AssociatedDeployments to be an explicit nil

### UnsetAssociatedDeployments
`func (o *Account) UnsetAssociatedDeployments()`

UnsetAssociatedDeployments ensures that no value is present for AssociatedDeployments, not even an explicit nil
### GetAssociatedResourceConnections

`func (o *Account) GetAssociatedResourceConnections() []AssociatedResourceConnection`

GetAssociatedResourceConnections returns the AssociatedResourceConnections field if non-nil, zero value otherwise.

### GetAssociatedResourceConnectionsOk

`func (o *Account) GetAssociatedResourceConnectionsOk() (*[]AssociatedResourceConnection, bool)`

GetAssociatedResourceConnectionsOk returns a tuple with the AssociatedResourceConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedResourceConnections

`func (o *Account) SetAssociatedResourceConnections(v []AssociatedResourceConnection)`

SetAssociatedResourceConnections sets AssociatedResourceConnections field to given value.

### HasAssociatedResourceConnections

`func (o *Account) HasAssociatedResourceConnections() bool`

HasAssociatedResourceConnections returns a boolean if a field has been set.

### SetAssociatedResourceConnectionsNil

`func (o *Account) SetAssociatedResourceConnectionsNil(b bool)`

 SetAssociatedResourceConnectionsNil sets the value for AssociatedResourceConnections to be an explicit nil

### UnsetAssociatedResourceConnections
`func (o *Account) UnsetAssociatedResourceConnections()`

UnsetAssociatedResourceConnections ensures that no value is present for AssociatedResourceConnections, not even an explicit nil
### GetAssociatedImages

`func (o *Account) GetAssociatedImages() []AssociatedImage`

GetAssociatedImages returns the AssociatedImages field if non-nil, zero value otherwise.

### GetAssociatedImagesOk

`func (o *Account) GetAssociatedImagesOk() (*[]AssociatedImage, bool)`

GetAssociatedImagesOk returns a tuple with the AssociatedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedImages

`func (o *Account) SetAssociatedImages(v []AssociatedImage)`

SetAssociatedImages sets AssociatedImages field to given value.

### HasAssociatedImages

`func (o *Account) HasAssociatedImages() bool`

HasAssociatedImages returns a boolean if a field has been set.

### SetAssociatedImagesNil

`func (o *Account) SetAssociatedImagesNil(b bool)`

 SetAssociatedImagesNil sets the value for AssociatedImages to be an explicit nil

### UnsetAssociatedImages
`func (o *Account) UnsetAssociatedImages()`

UnsetAssociatedImages ensures that no value is present for AssociatedImages, not even an explicit nil
### GetCitrixManaged

`func (o *Account) GetCitrixManaged() bool`

GetCitrixManaged returns the CitrixManaged field if non-nil, zero value otherwise.

### GetCitrixManagedOk

`func (o *Account) GetCitrixManagedOk() (*bool, bool)`

GetCitrixManagedOk returns a tuple with the CitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManaged

`func (o *Account) SetCitrixManaged(v bool)`

SetCitrixManaged sets CitrixManaged field to given value.

### HasCitrixManaged

`func (o *Account) HasCitrixManaged() bool`

HasCitrixManaged returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Account) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Account) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Account) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Account) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Account) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Account) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetWarnings

`func (o *Account) GetWarnings() []AccountWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Account) GetWarningsOk() (*[]AccountWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Account) SetWarnings(v []AccountWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Account) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *Account) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *Account) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


