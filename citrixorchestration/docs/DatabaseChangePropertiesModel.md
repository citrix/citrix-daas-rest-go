# DatabaseChangePropertiesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | **string** | Name of database | 
**DatabaseServerAddress** | **string** | Server address of database | 
**ManualUpdateDatabase** | **bool** | Indicates whether the database should be created manually by scripts | 
**SkipWarning** | **bool** | Indicates whether ignore warnings where changing database | 

## Methods

### NewDatabaseChangePropertiesModel

`func NewDatabaseChangePropertiesModel(databaseName string, databaseServerAddress string, manualUpdateDatabase bool, skipWarning bool, ) *DatabaseChangePropertiesModel`

NewDatabaseChangePropertiesModel instantiates a new DatabaseChangePropertiesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseChangePropertiesModelWithDefaults

`func NewDatabaseChangePropertiesModelWithDefaults() *DatabaseChangePropertiesModel`

NewDatabaseChangePropertiesModelWithDefaults instantiates a new DatabaseChangePropertiesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *DatabaseChangePropertiesModel) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *DatabaseChangePropertiesModel) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *DatabaseChangePropertiesModel) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDatabaseServerAddress

`func (o *DatabaseChangePropertiesModel) GetDatabaseServerAddress() string`

GetDatabaseServerAddress returns the DatabaseServerAddress field if non-nil, zero value otherwise.

### GetDatabaseServerAddressOk

`func (o *DatabaseChangePropertiesModel) GetDatabaseServerAddressOk() (*string, bool)`

GetDatabaseServerAddressOk returns a tuple with the DatabaseServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseServerAddress

`func (o *DatabaseChangePropertiesModel) SetDatabaseServerAddress(v string)`

SetDatabaseServerAddress sets DatabaseServerAddress field to given value.


### GetManualUpdateDatabase

`func (o *DatabaseChangePropertiesModel) GetManualUpdateDatabase() bool`

GetManualUpdateDatabase returns the ManualUpdateDatabase field if non-nil, zero value otherwise.

### GetManualUpdateDatabaseOk

`func (o *DatabaseChangePropertiesModel) GetManualUpdateDatabaseOk() (*bool, bool)`

GetManualUpdateDatabaseOk returns a tuple with the ManualUpdateDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualUpdateDatabase

`func (o *DatabaseChangePropertiesModel) SetManualUpdateDatabase(v bool)`

SetManualUpdateDatabase sets ManualUpdateDatabase field to given value.


### GetSkipWarning

`func (o *DatabaseChangePropertiesModel) GetSkipWarning() bool`

GetSkipWarning returns the SkipWarning field if non-nil, zero value otherwise.

### GetSkipWarningOk

`func (o *DatabaseChangePropertiesModel) GetSkipWarningOk() (*bool, bool)`

GetSkipWarningOk returns a tuple with the SkipWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWarning

`func (o *DatabaseChangePropertiesModel) SetSkipWarning(v bool)`

SetSkipWarning sets SkipWarning field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


