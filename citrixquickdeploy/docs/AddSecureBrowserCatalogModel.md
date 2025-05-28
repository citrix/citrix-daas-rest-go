# AddSecureBrowserCatalogModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | The region where the catalog should be deployed, if it does not have any on-prem connectivity | 
**SubscriptionId** | **string** | The Id of the BYOA subscription | 
**Name** | **string** | Name of the secure browser catalog | 
**Type** | Pointer to [**AddCatalogType**](AddCatalogType.md) | Indicates if the catalog VDAs run a single session or multiple sessions | [optional] 
**IsDomainJoined** | Pointer to **bool** | Indicates if the catalog will connected to the customers domain | [optional] 
**PersistStaticAllocatedVmDisks** | Pointer to **bool** | Indicates if catalogs that use statically allocated machines will have the disk contents persisted after shutdown | [optional] 
**MachineNamingScheme** | Pointer to [**MachineNamingSchemeModel**](MachineNamingSchemeModel.md) |  | [optional] 
**IsAzureAdJoined** | Pointer to **bool** | Indicates if the machines in catalog will be Azure AD joined | [optional] 
**EnableIntuneEnroll** | Pointer to **bool** | Specifies if intune enroll is enabled | [optional] 

## Methods

### NewAddSecureBrowserCatalogModel

`func NewAddSecureBrowserCatalogModel(region string, subscriptionId string, name string, ) *AddSecureBrowserCatalogModel`

NewAddSecureBrowserCatalogModel instantiates a new AddSecureBrowserCatalogModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecureBrowserCatalogModelWithDefaults

`func NewAddSecureBrowserCatalogModelWithDefaults() *AddSecureBrowserCatalogModel`

NewAddSecureBrowserCatalogModelWithDefaults instantiates a new AddSecureBrowserCatalogModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AddSecureBrowserCatalogModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddSecureBrowserCatalogModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddSecureBrowserCatalogModel) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSubscriptionId

`func (o *AddSecureBrowserCatalogModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AddSecureBrowserCatalogModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AddSecureBrowserCatalogModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetName

`func (o *AddSecureBrowserCatalogModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSecureBrowserCatalogModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSecureBrowserCatalogModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddSecureBrowserCatalogModel) GetType() AddCatalogType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddSecureBrowserCatalogModel) GetTypeOk() (*AddCatalogType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddSecureBrowserCatalogModel) SetType(v AddCatalogType)`

SetType sets Type field to given value.

### HasType

`func (o *AddSecureBrowserCatalogModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsDomainJoined

`func (o *AddSecureBrowserCatalogModel) GetIsDomainJoined() bool`

GetIsDomainJoined returns the IsDomainJoined field if non-nil, zero value otherwise.

### GetIsDomainJoinedOk

`func (o *AddSecureBrowserCatalogModel) GetIsDomainJoinedOk() (*bool, bool)`

GetIsDomainJoinedOk returns a tuple with the IsDomainJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomainJoined

`func (o *AddSecureBrowserCatalogModel) SetIsDomainJoined(v bool)`

SetIsDomainJoined sets IsDomainJoined field to given value.

### HasIsDomainJoined

`func (o *AddSecureBrowserCatalogModel) HasIsDomainJoined() bool`

HasIsDomainJoined returns a boolean if a field has been set.

### GetPersistStaticAllocatedVmDisks

`func (o *AddSecureBrowserCatalogModel) GetPersistStaticAllocatedVmDisks() bool`

GetPersistStaticAllocatedVmDisks returns the PersistStaticAllocatedVmDisks field if non-nil, zero value otherwise.

### GetPersistStaticAllocatedVmDisksOk

`func (o *AddSecureBrowserCatalogModel) GetPersistStaticAllocatedVmDisksOk() (*bool, bool)`

GetPersistStaticAllocatedVmDisksOk returns a tuple with the PersistStaticAllocatedVmDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistStaticAllocatedVmDisks

`func (o *AddSecureBrowserCatalogModel) SetPersistStaticAllocatedVmDisks(v bool)`

SetPersistStaticAllocatedVmDisks sets PersistStaticAllocatedVmDisks field to given value.

### HasPersistStaticAllocatedVmDisks

`func (o *AddSecureBrowserCatalogModel) HasPersistStaticAllocatedVmDisks() bool`

HasPersistStaticAllocatedVmDisks returns a boolean if a field has been set.

### GetMachineNamingScheme

`func (o *AddSecureBrowserCatalogModel) GetMachineNamingScheme() MachineNamingSchemeModel`

GetMachineNamingScheme returns the MachineNamingScheme field if non-nil, zero value otherwise.

### GetMachineNamingSchemeOk

`func (o *AddSecureBrowserCatalogModel) GetMachineNamingSchemeOk() (*MachineNamingSchemeModel, bool)`

GetMachineNamingSchemeOk returns a tuple with the MachineNamingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineNamingScheme

`func (o *AddSecureBrowserCatalogModel) SetMachineNamingScheme(v MachineNamingSchemeModel)`

SetMachineNamingScheme sets MachineNamingScheme field to given value.

### HasMachineNamingScheme

`func (o *AddSecureBrowserCatalogModel) HasMachineNamingScheme() bool`

HasMachineNamingScheme returns a boolean if a field has been set.

### GetIsAzureAdJoined

`func (o *AddSecureBrowserCatalogModel) GetIsAzureAdJoined() bool`

GetIsAzureAdJoined returns the IsAzureAdJoined field if non-nil, zero value otherwise.

### GetIsAzureAdJoinedOk

`func (o *AddSecureBrowserCatalogModel) GetIsAzureAdJoinedOk() (*bool, bool)`

GetIsAzureAdJoinedOk returns a tuple with the IsAzureAdJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzureAdJoined

`func (o *AddSecureBrowserCatalogModel) SetIsAzureAdJoined(v bool)`

SetIsAzureAdJoined sets IsAzureAdJoined field to given value.

### HasIsAzureAdJoined

`func (o *AddSecureBrowserCatalogModel) HasIsAzureAdJoined() bool`

HasIsAzureAdJoined returns a boolean if a field has been set.

### GetEnableIntuneEnroll

`func (o *AddSecureBrowserCatalogModel) GetEnableIntuneEnroll() bool`

GetEnableIntuneEnroll returns the EnableIntuneEnroll field if non-nil, zero value otherwise.

### GetEnableIntuneEnrollOk

`func (o *AddSecureBrowserCatalogModel) GetEnableIntuneEnrollOk() (*bool, bool)`

GetEnableIntuneEnrollOk returns a tuple with the EnableIntuneEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntuneEnroll

`func (o *AddSecureBrowserCatalogModel) SetEnableIntuneEnroll(v bool)`

SetEnableIntuneEnroll sets EnableIntuneEnroll field to given value.

### HasEnableIntuneEnroll

`func (o *AddSecureBrowserCatalogModel) HasEnableIntuneEnroll() bool`

HasEnableIntuneEnroll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


