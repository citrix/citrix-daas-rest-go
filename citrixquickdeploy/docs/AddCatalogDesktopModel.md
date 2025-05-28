# AddCatalogDesktopModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the desktop | 
**Description** | Pointer to **string** | Description | [optional] 
**IncludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | Users to be assigned as part of desktop publishing | [optional] 

## Methods

### NewAddCatalogDesktopModel

`func NewAddCatalogDesktopModel(name string, ) *AddCatalogDesktopModel`

NewAddCatalogDesktopModel instantiates a new AddCatalogDesktopModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogDesktopModelWithDefaults

`func NewAddCatalogDesktopModelWithDefaults() *AddCatalogDesktopModel`

NewAddCatalogDesktopModelWithDefaults instantiates a new AddCatalogDesktopModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddCatalogDesktopModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCatalogDesktopModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCatalogDesktopModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddCatalogDesktopModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCatalogDesktopModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCatalogDesktopModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCatalogDesktopModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludedUsers

`func (o *AddCatalogDesktopModel) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *AddCatalogDesktopModel) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *AddCatalogDesktopModel) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *AddCatalogDesktopModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


