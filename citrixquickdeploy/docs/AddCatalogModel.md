# AddCatalogModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the catalog | 
**Type** | Pointer to [**AddCatalogType**](AddCatalogType.md) | Indicates if the catalog VDAs run a single session or multiple sessions | [optional] [default to ADDCATALOGTYPE_MULTI_SESSION]
**IsDomainJoined** | Pointer to **bool** | Indicates if the catalog will connected to the customers domain | [optional] 
**PersistStaticAllocatedVmDisks** | Pointer to **bool** | Indicates if catalogs that use statically allocated machines will have the disk contents persisted after shutdown | [optional] 
**MachineNamingScheme** | Pointer to [**NullableMachineNamingSchemeModel**](MachineNamingSchemeModel.md) |  | [optional] 
**IsAzureAdJoined** | Pointer to **bool** | Indicates if the machines in catalog will be Azure AD joined | [optional] 
**EnableIntuneEnroll** | Pointer to **bool** | Specifies if intune enroll is enabled | [optional] 
**MinUniqueUsers** | Pointer to **NullableInt32** | Specifies the minimum number of unique users that the catalog will support. This field (in conjunction with MaxUserPerVm in capacity settings) will determine the number of machines. MaxInstances in scale settings will be overridden if specified.  For Citrix Managed catalogs only | [optional] 

## Methods

### NewAddCatalogModel

`func NewAddCatalogModel(name string, ) *AddCatalogModel`

NewAddCatalogModel instantiates a new AddCatalogModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogModelWithDefaults

`func NewAddCatalogModelWithDefaults() *AddCatalogModel`

NewAddCatalogModelWithDefaults instantiates a new AddCatalogModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddCatalogModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCatalogModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCatalogModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddCatalogModel) GetType() AddCatalogType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCatalogModel) GetTypeOk() (*AddCatalogType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCatalogModel) SetType(v AddCatalogType)`

SetType sets Type field to given value.

### HasType

`func (o *AddCatalogModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsDomainJoined

`func (o *AddCatalogModel) GetIsDomainJoined() bool`

GetIsDomainJoined returns the IsDomainJoined field if non-nil, zero value otherwise.

### GetIsDomainJoinedOk

`func (o *AddCatalogModel) GetIsDomainJoinedOk() (*bool, bool)`

GetIsDomainJoinedOk returns a tuple with the IsDomainJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomainJoined

`func (o *AddCatalogModel) SetIsDomainJoined(v bool)`

SetIsDomainJoined sets IsDomainJoined field to given value.

### HasIsDomainJoined

`func (o *AddCatalogModel) HasIsDomainJoined() bool`

HasIsDomainJoined returns a boolean if a field has been set.

### GetPersistStaticAllocatedVmDisks

`func (o *AddCatalogModel) GetPersistStaticAllocatedVmDisks() bool`

GetPersistStaticAllocatedVmDisks returns the PersistStaticAllocatedVmDisks field if non-nil, zero value otherwise.

### GetPersistStaticAllocatedVmDisksOk

`func (o *AddCatalogModel) GetPersistStaticAllocatedVmDisksOk() (*bool, bool)`

GetPersistStaticAllocatedVmDisksOk returns a tuple with the PersistStaticAllocatedVmDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistStaticAllocatedVmDisks

`func (o *AddCatalogModel) SetPersistStaticAllocatedVmDisks(v bool)`

SetPersistStaticAllocatedVmDisks sets PersistStaticAllocatedVmDisks field to given value.

### HasPersistStaticAllocatedVmDisks

`func (o *AddCatalogModel) HasPersistStaticAllocatedVmDisks() bool`

HasPersistStaticAllocatedVmDisks returns a boolean if a field has been set.

### GetMachineNamingScheme

`func (o *AddCatalogModel) GetMachineNamingScheme() MachineNamingSchemeModel`

GetMachineNamingScheme returns the MachineNamingScheme field if non-nil, zero value otherwise.

### GetMachineNamingSchemeOk

`func (o *AddCatalogModel) GetMachineNamingSchemeOk() (*MachineNamingSchemeModel, bool)`

GetMachineNamingSchemeOk returns a tuple with the MachineNamingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineNamingScheme

`func (o *AddCatalogModel) SetMachineNamingScheme(v MachineNamingSchemeModel)`

SetMachineNamingScheme sets MachineNamingScheme field to given value.

### HasMachineNamingScheme

`func (o *AddCatalogModel) HasMachineNamingScheme() bool`

HasMachineNamingScheme returns a boolean if a field has been set.

### SetMachineNamingSchemeNil

`func (o *AddCatalogModel) SetMachineNamingSchemeNil(b bool)`

 SetMachineNamingSchemeNil sets the value for MachineNamingScheme to be an explicit nil

### UnsetMachineNamingScheme
`func (o *AddCatalogModel) UnsetMachineNamingScheme()`

UnsetMachineNamingScheme ensures that no value is present for MachineNamingScheme, not even an explicit nil
### GetIsAzureAdJoined

`func (o *AddCatalogModel) GetIsAzureAdJoined() bool`

GetIsAzureAdJoined returns the IsAzureAdJoined field if non-nil, zero value otherwise.

### GetIsAzureAdJoinedOk

`func (o *AddCatalogModel) GetIsAzureAdJoinedOk() (*bool, bool)`

GetIsAzureAdJoinedOk returns a tuple with the IsAzureAdJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzureAdJoined

`func (o *AddCatalogModel) SetIsAzureAdJoined(v bool)`

SetIsAzureAdJoined sets IsAzureAdJoined field to given value.

### HasIsAzureAdJoined

`func (o *AddCatalogModel) HasIsAzureAdJoined() bool`

HasIsAzureAdJoined returns a boolean if a field has been set.

### GetEnableIntuneEnroll

`func (o *AddCatalogModel) GetEnableIntuneEnroll() bool`

GetEnableIntuneEnroll returns the EnableIntuneEnroll field if non-nil, zero value otherwise.

### GetEnableIntuneEnrollOk

`func (o *AddCatalogModel) GetEnableIntuneEnrollOk() (*bool, bool)`

GetEnableIntuneEnrollOk returns a tuple with the EnableIntuneEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntuneEnroll

`func (o *AddCatalogModel) SetEnableIntuneEnroll(v bool)`

SetEnableIntuneEnroll sets EnableIntuneEnroll field to given value.

### HasEnableIntuneEnroll

`func (o *AddCatalogModel) HasEnableIntuneEnroll() bool`

HasEnableIntuneEnroll returns a boolean if a field has been set.

### GetMinUniqueUsers

`func (o *AddCatalogModel) GetMinUniqueUsers() int32`

GetMinUniqueUsers returns the MinUniqueUsers field if non-nil, zero value otherwise.

### GetMinUniqueUsersOk

`func (o *AddCatalogModel) GetMinUniqueUsersOk() (*int32, bool)`

GetMinUniqueUsersOk returns a tuple with the MinUniqueUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUniqueUsers

`func (o *AddCatalogModel) SetMinUniqueUsers(v int32)`

SetMinUniqueUsers sets MinUniqueUsers field to given value.

### HasMinUniqueUsers

`func (o *AddCatalogModel) HasMinUniqueUsers() bool`

HasMinUniqueUsers returns a boolean if a field has been set.

### SetMinUniqueUsersNil

`func (o *AddCatalogModel) SetMinUniqueUsersNil(b bool)`

 SetMinUniqueUsersNil sets the value for MinUniqueUsers to be an explicit nil

### UnsetMinUniqueUsers
`func (o *AddCatalogModel) UnsetMinUniqueUsers()`

UnsetMinUniqueUsers ensures that no value is present for MinUniqueUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


