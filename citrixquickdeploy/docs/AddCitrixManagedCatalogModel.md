# AddCitrixManagedCatalogModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The region where the catalog should be deployed, if it does not have any on-prem connectivity | [optional] 
**Name** | **string** | Name of the catalog | 
**Type** | Pointer to [**AddCatalogType**](AddCatalogType.md) | Indicates if the catalog VDAs run a single session or multiple sessions | [optional] 
**IsDomainJoined** | Pointer to **bool** | Indicates if the catalog will connected to the customers domain | [optional] 
**PersistStaticAllocatedVmDisks** | Pointer to **bool** | Indicates if catalogs that use statically allocated machines will have the disk contents persisted after shutdown | [optional] 
**MachineNamingScheme** | Pointer to [**MachineNamingSchemeModel**](MachineNamingSchemeModel.md) |  | [optional] 
**IsAzureAdJoined** | Pointer to **bool** | Indicates if the machines in catalog will be Azure AD joined | [optional] 
**EnableIntuneEnroll** | Pointer to **bool** | Specifies if intune enroll is enabled | [optional] 

## Methods

### NewAddCitrixManagedCatalogModel

`func NewAddCitrixManagedCatalogModel(name string, ) *AddCitrixManagedCatalogModel`

NewAddCitrixManagedCatalogModel instantiates a new AddCitrixManagedCatalogModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCitrixManagedCatalogModelWithDefaults

`func NewAddCitrixManagedCatalogModelWithDefaults() *AddCitrixManagedCatalogModel`

NewAddCitrixManagedCatalogModelWithDefaults instantiates a new AddCitrixManagedCatalogModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AddCitrixManagedCatalogModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddCitrixManagedCatalogModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddCitrixManagedCatalogModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddCitrixManagedCatalogModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetName

`func (o *AddCitrixManagedCatalogModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCitrixManagedCatalogModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCitrixManagedCatalogModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddCitrixManagedCatalogModel) GetType() AddCatalogType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCitrixManagedCatalogModel) GetTypeOk() (*AddCatalogType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCitrixManagedCatalogModel) SetType(v AddCatalogType)`

SetType sets Type field to given value.

### HasType

`func (o *AddCitrixManagedCatalogModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsDomainJoined

`func (o *AddCitrixManagedCatalogModel) GetIsDomainJoined() bool`

GetIsDomainJoined returns the IsDomainJoined field if non-nil, zero value otherwise.

### GetIsDomainJoinedOk

`func (o *AddCitrixManagedCatalogModel) GetIsDomainJoinedOk() (*bool, bool)`

GetIsDomainJoinedOk returns a tuple with the IsDomainJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomainJoined

`func (o *AddCitrixManagedCatalogModel) SetIsDomainJoined(v bool)`

SetIsDomainJoined sets IsDomainJoined field to given value.

### HasIsDomainJoined

`func (o *AddCitrixManagedCatalogModel) HasIsDomainJoined() bool`

HasIsDomainJoined returns a boolean if a field has been set.

### GetPersistStaticAllocatedVmDisks

`func (o *AddCitrixManagedCatalogModel) GetPersistStaticAllocatedVmDisks() bool`

GetPersistStaticAllocatedVmDisks returns the PersistStaticAllocatedVmDisks field if non-nil, zero value otherwise.

### GetPersistStaticAllocatedVmDisksOk

`func (o *AddCitrixManagedCatalogModel) GetPersistStaticAllocatedVmDisksOk() (*bool, bool)`

GetPersistStaticAllocatedVmDisksOk returns a tuple with the PersistStaticAllocatedVmDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistStaticAllocatedVmDisks

`func (o *AddCitrixManagedCatalogModel) SetPersistStaticAllocatedVmDisks(v bool)`

SetPersistStaticAllocatedVmDisks sets PersistStaticAllocatedVmDisks field to given value.

### HasPersistStaticAllocatedVmDisks

`func (o *AddCitrixManagedCatalogModel) HasPersistStaticAllocatedVmDisks() bool`

HasPersistStaticAllocatedVmDisks returns a boolean if a field has been set.

### GetMachineNamingScheme

`func (o *AddCitrixManagedCatalogModel) GetMachineNamingScheme() MachineNamingSchemeModel`

GetMachineNamingScheme returns the MachineNamingScheme field if non-nil, zero value otherwise.

### GetMachineNamingSchemeOk

`func (o *AddCitrixManagedCatalogModel) GetMachineNamingSchemeOk() (*MachineNamingSchemeModel, bool)`

GetMachineNamingSchemeOk returns a tuple with the MachineNamingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineNamingScheme

`func (o *AddCitrixManagedCatalogModel) SetMachineNamingScheme(v MachineNamingSchemeModel)`

SetMachineNamingScheme sets MachineNamingScheme field to given value.

### HasMachineNamingScheme

`func (o *AddCitrixManagedCatalogModel) HasMachineNamingScheme() bool`

HasMachineNamingScheme returns a boolean if a field has been set.

### GetIsAzureAdJoined

`func (o *AddCitrixManagedCatalogModel) GetIsAzureAdJoined() bool`

GetIsAzureAdJoined returns the IsAzureAdJoined field if non-nil, zero value otherwise.

### GetIsAzureAdJoinedOk

`func (o *AddCitrixManagedCatalogModel) GetIsAzureAdJoinedOk() (*bool, bool)`

GetIsAzureAdJoinedOk returns a tuple with the IsAzureAdJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzureAdJoined

`func (o *AddCitrixManagedCatalogModel) SetIsAzureAdJoined(v bool)`

SetIsAzureAdJoined sets IsAzureAdJoined field to given value.

### HasIsAzureAdJoined

`func (o *AddCitrixManagedCatalogModel) HasIsAzureAdJoined() bool`

HasIsAzureAdJoined returns a boolean if a field has been set.

### GetEnableIntuneEnroll

`func (o *AddCitrixManagedCatalogModel) GetEnableIntuneEnroll() bool`

GetEnableIntuneEnroll returns the EnableIntuneEnroll field if non-nil, zero value otherwise.

### GetEnableIntuneEnrollOk

`func (o *AddCitrixManagedCatalogModel) GetEnableIntuneEnrollOk() (*bool, bool)`

GetEnableIntuneEnrollOk returns a tuple with the EnableIntuneEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntuneEnroll

`func (o *AddCitrixManagedCatalogModel) SetEnableIntuneEnroll(v bool)`

SetEnableIntuneEnroll sets EnableIntuneEnroll field to given value.

### HasEnableIntuneEnroll

`func (o *AddCitrixManagedCatalogModel) HasEnableIntuneEnroll() bool`

HasEnableIntuneEnroll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


