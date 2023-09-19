# DatabaseResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataStore** | Pointer to [**DataStore**](DataStore.md) |  | [optional] 
**DatabaseName** | Pointer to **NullableString** | Name of database | [optional] 
**DatabaseServerAddress** | Pointer to **NullableString** | Server address of database | [optional] 
**MirrorDatabaseServerAddress** | Pointer to **NullableString** | Mirror address of database | [optional] 
**CanChangeDatabase** | Pointer to **bool** | Indicates whether this database can be changed | [optional] 

## Methods

### NewDatabaseResponseModel

`func NewDatabaseResponseModel() *DatabaseResponseModel`

NewDatabaseResponseModel instantiates a new DatabaseResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseResponseModelWithDefaults

`func NewDatabaseResponseModelWithDefaults() *DatabaseResponseModel`

NewDatabaseResponseModelWithDefaults instantiates a new DatabaseResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataStore

`func (o *DatabaseResponseModel) GetDataStore() DataStore`

GetDataStore returns the DataStore field if non-nil, zero value otherwise.

### GetDataStoreOk

`func (o *DatabaseResponseModel) GetDataStoreOk() (*DataStore, bool)`

GetDataStoreOk returns a tuple with the DataStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStore

`func (o *DatabaseResponseModel) SetDataStore(v DataStore)`

SetDataStore sets DataStore field to given value.

### HasDataStore

`func (o *DatabaseResponseModel) HasDataStore() bool`

HasDataStore returns a boolean if a field has been set.

### GetDatabaseName

`func (o *DatabaseResponseModel) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *DatabaseResponseModel) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *DatabaseResponseModel) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *DatabaseResponseModel) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### SetDatabaseNameNil

`func (o *DatabaseResponseModel) SetDatabaseNameNil(b bool)`

 SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil

### UnsetDatabaseName
`func (o *DatabaseResponseModel) UnsetDatabaseName()`

UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
### GetDatabaseServerAddress

`func (o *DatabaseResponseModel) GetDatabaseServerAddress() string`

GetDatabaseServerAddress returns the DatabaseServerAddress field if non-nil, zero value otherwise.

### GetDatabaseServerAddressOk

`func (o *DatabaseResponseModel) GetDatabaseServerAddressOk() (*string, bool)`

GetDatabaseServerAddressOk returns a tuple with the DatabaseServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseServerAddress

`func (o *DatabaseResponseModel) SetDatabaseServerAddress(v string)`

SetDatabaseServerAddress sets DatabaseServerAddress field to given value.

### HasDatabaseServerAddress

`func (o *DatabaseResponseModel) HasDatabaseServerAddress() bool`

HasDatabaseServerAddress returns a boolean if a field has been set.

### SetDatabaseServerAddressNil

`func (o *DatabaseResponseModel) SetDatabaseServerAddressNil(b bool)`

 SetDatabaseServerAddressNil sets the value for DatabaseServerAddress to be an explicit nil

### UnsetDatabaseServerAddress
`func (o *DatabaseResponseModel) UnsetDatabaseServerAddress()`

UnsetDatabaseServerAddress ensures that no value is present for DatabaseServerAddress, not even an explicit nil
### GetMirrorDatabaseServerAddress

`func (o *DatabaseResponseModel) GetMirrorDatabaseServerAddress() string`

GetMirrorDatabaseServerAddress returns the MirrorDatabaseServerAddress field if non-nil, zero value otherwise.

### GetMirrorDatabaseServerAddressOk

`func (o *DatabaseResponseModel) GetMirrorDatabaseServerAddressOk() (*string, bool)`

GetMirrorDatabaseServerAddressOk returns a tuple with the MirrorDatabaseServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorDatabaseServerAddress

`func (o *DatabaseResponseModel) SetMirrorDatabaseServerAddress(v string)`

SetMirrorDatabaseServerAddress sets MirrorDatabaseServerAddress field to given value.

### HasMirrorDatabaseServerAddress

`func (o *DatabaseResponseModel) HasMirrorDatabaseServerAddress() bool`

HasMirrorDatabaseServerAddress returns a boolean if a field has been set.

### SetMirrorDatabaseServerAddressNil

`func (o *DatabaseResponseModel) SetMirrorDatabaseServerAddressNil(b bool)`

 SetMirrorDatabaseServerAddressNil sets the value for MirrorDatabaseServerAddress to be an explicit nil

### UnsetMirrorDatabaseServerAddress
`func (o *DatabaseResponseModel) UnsetMirrorDatabaseServerAddress()`

UnsetMirrorDatabaseServerAddress ensures that no value is present for MirrorDatabaseServerAddress, not even an explicit nil
### GetCanChangeDatabase

`func (o *DatabaseResponseModel) GetCanChangeDatabase() bool`

GetCanChangeDatabase returns the CanChangeDatabase field if non-nil, zero value otherwise.

### GetCanChangeDatabaseOk

`func (o *DatabaseResponseModel) GetCanChangeDatabaseOk() (*bool, bool)`

GetCanChangeDatabaseOk returns a tuple with the CanChangeDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChangeDatabase

`func (o *DatabaseResponseModel) SetCanChangeDatabase(v bool)`

SetCanChangeDatabase sets CanChangeDatabase field to given value.

### HasCanChangeDatabase

`func (o *DatabaseResponseModel) HasCanChangeDatabase() bool`

HasCanChangeDatabase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


