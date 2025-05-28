# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) | The type of account the resource is associated with | 
**AccountId** | Pointer to **NullableString** | The ID of the account the resource is associated with | [optional] 
**ImageId** | Pointer to **NullableString** | The ID of the image | [optional] 
**Name** | Pointer to **NullableString** | The name of the image | [optional] 
**Description** | Pointer to **NullableString** | The description of the image | [optional] 
**Notes** | Pointer to **NullableString** | The notes of the image | [optional] 
**SessionSupport** | Pointer to [**NullableSessionSupport**](SessionSupport.md) | Session Type | [optional] 
**OperatingSystem** | Pointer to [**NullableOperatingSystemType**](OperatingSystemType.md) | The operating system of the image | [optional] 
**AssociatedDeployments** | Pointer to [**[]AssociatedDeployment**](AssociatedDeployment.md) |  | [optional] 

## Methods

### NewImage

`func NewImage(accountType AccountType, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *Image) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Image) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Image) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetAccountId

`func (o *Image) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Image) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Image) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Image) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *Image) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Image) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetImageId

`func (o *Image) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *Image) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *Image) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *Image) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *Image) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *Image) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetName

`func (o *Image) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Image) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Image) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Image) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Image) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Image) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Image) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Image) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Image) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Image) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Image) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Image) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNotes

`func (o *Image) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Image) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Image) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Image) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Image) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Image) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetSessionSupport

`func (o *Image) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *Image) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *Image) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *Image) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.

### SetSessionSupportNil

`func (o *Image) SetSessionSupportNil(b bool)`

 SetSessionSupportNil sets the value for SessionSupport to be an explicit nil

### UnsetSessionSupport
`func (o *Image) UnsetSessionSupport()`

UnsetSessionSupport ensures that no value is present for SessionSupport, not even an explicit nil
### GetOperatingSystem

`func (o *Image) GetOperatingSystem() OperatingSystemType`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Image) GetOperatingSystemOk() (*OperatingSystemType, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Image) SetOperatingSystem(v OperatingSystemType)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *Image) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystemNil

`func (o *Image) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *Image) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetAssociatedDeployments

`func (o *Image) GetAssociatedDeployments() []AssociatedDeployment`

GetAssociatedDeployments returns the AssociatedDeployments field if non-nil, zero value otherwise.

### GetAssociatedDeploymentsOk

`func (o *Image) GetAssociatedDeploymentsOk() (*[]AssociatedDeployment, bool)`

GetAssociatedDeploymentsOk returns a tuple with the AssociatedDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDeployments

`func (o *Image) SetAssociatedDeployments(v []AssociatedDeployment)`

SetAssociatedDeployments sets AssociatedDeployments field to given value.

### HasAssociatedDeployments

`func (o *Image) HasAssociatedDeployments() bool`

HasAssociatedDeployments returns a boolean if a field has been set.

### SetAssociatedDeploymentsNil

`func (o *Image) SetAssociatedDeploymentsNil(b bool)`

 SetAssociatedDeploymentsNil sets the value for AssociatedDeployments to be an explicit nil

### UnsetAssociatedDeployments
`func (o *Image) UnsetAssociatedDeployments()`

UnsetAssociatedDeployments ensures that no value is present for AssociatedDeployments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


