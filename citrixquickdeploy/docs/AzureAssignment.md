# AzureAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | Pointer to **string** | ID of the assignment | [optional] 
**RoleId** | Pointer to **string** | ID of the Role that has been assigned | [optional] 
**RoleName** | Pointer to **string** | Name of the assigned Role | [optional] 
**ObjectType** | Pointer to **string** | Type of object that has been assigned the role | [optional] 
**DisplayName** | Pointer to **string** | Display name of the app / user assigned the role | [optional] 
**AppId** | Pointer to **string** | ID of the app that is associated with the Service Principal that hasd the assignment | [optional] 
**Upn** | Pointer to **string** | UPN of the user assigned the role | [optional] 

## Methods

### NewAzureAssignment

`func NewAzureAssignment() *AzureAssignment`

NewAzureAssignment instantiates a new AzureAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAssignmentWithDefaults

`func NewAzureAssignmentWithDefaults() *AzureAssignment`

NewAzureAssignmentWithDefaults instantiates a new AzureAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *AzureAssignment) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *AzureAssignment) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *AzureAssignment) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *AzureAssignment) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetRoleId

`func (o *AzureAssignment) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AzureAssignment) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AzureAssignment) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *AzureAssignment) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *AzureAssignment) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AzureAssignment) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AzureAssignment) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AzureAssignment) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetObjectType

`func (o *AzureAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AzureAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AzureAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AzureAssignment) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetDisplayName

`func (o *AzureAssignment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AzureAssignment) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AzureAssignment) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AzureAssignment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAppId

`func (o *AzureAssignment) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AzureAssignment) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AzureAssignment) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AzureAssignment) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetUpn

`func (o *AzureAssignment) GetUpn() string`

GetUpn returns the Upn field if non-nil, zero value otherwise.

### GetUpnOk

`func (o *AzureAssignment) GetUpnOk() (*string, bool)`

GetUpnOk returns a tuple with the Upn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpn

`func (o *AzureAssignment) SetUpn(v string)`

SetUpn sets Upn field to given value.

### HasUpn

`func (o *AzureAssignment) HasUpn() bool`

HasUpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


