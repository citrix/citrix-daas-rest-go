# RemotePcEnrollmentScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ou** | Pointer to **string** | The OU to include in the catalog. | [optional] 
**IncludeSubfolders** | Pointer to **bool** | Indicates if the subfolders of this OU should also be included. | [optional] 

## Methods

### NewRemotePcEnrollmentScope

`func NewRemotePcEnrollmentScope() *RemotePcEnrollmentScope`

NewRemotePcEnrollmentScope instantiates a new RemotePcEnrollmentScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemotePcEnrollmentScopeWithDefaults

`func NewRemotePcEnrollmentScopeWithDefaults() *RemotePcEnrollmentScope`

NewRemotePcEnrollmentScopeWithDefaults instantiates a new RemotePcEnrollmentScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOu

`func (o *RemotePcEnrollmentScope) GetOu() string`

GetOu returns the Ou field if non-nil, zero value otherwise.

### GetOuOk

`func (o *RemotePcEnrollmentScope) GetOuOk() (*string, bool)`

GetOuOk returns a tuple with the Ou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOu

`func (o *RemotePcEnrollmentScope) SetOu(v string)`

SetOu sets Ou field to given value.

### HasOu

`func (o *RemotePcEnrollmentScope) HasOu() bool`

HasOu returns a boolean if a field has been set.

### GetIncludeSubfolders

`func (o *RemotePcEnrollmentScope) GetIncludeSubfolders() bool`

GetIncludeSubfolders returns the IncludeSubfolders field if non-nil, zero value otherwise.

### GetIncludeSubfoldersOk

`func (o *RemotePcEnrollmentScope) GetIncludeSubfoldersOk() (*bool, bool)`

GetIncludeSubfoldersOk returns a tuple with the IncludeSubfolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubfolders

`func (o *RemotePcEnrollmentScope) SetIncludeSubfolders(v bool)`

SetIncludeSubfolders sets IncludeSubfolders field to given value.

### HasIncludeSubfolders

`func (o *RemotePcEnrollmentScope) HasIncludeSubfolders() bool`

HasIncludeSubfolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


