# ScheduledBackupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**VirtualSiteId** | **string** |  | 
**CatalogsForBackup** | [**[]CatalogsForBackup**](CatalogsForBackup.md) |  | 

## Methods

### NewScheduledBackupModel

`func NewScheduledBackupModel(customerId string, virtualSiteId string, catalogsForBackup []CatalogsForBackup, ) *ScheduledBackupModel`

NewScheduledBackupModel instantiates a new ScheduledBackupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledBackupModelWithDefaults

`func NewScheduledBackupModelWithDefaults() *ScheduledBackupModel`

NewScheduledBackupModelWithDefaults instantiates a new ScheduledBackupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *ScheduledBackupModel) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ScheduledBackupModel) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ScheduledBackupModel) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetVirtualSiteId

`func (o *ScheduledBackupModel) GetVirtualSiteId() string`

GetVirtualSiteId returns the VirtualSiteId field if non-nil, zero value otherwise.

### GetVirtualSiteIdOk

`func (o *ScheduledBackupModel) GetVirtualSiteIdOk() (*string, bool)`

GetVirtualSiteIdOk returns a tuple with the VirtualSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSiteId

`func (o *ScheduledBackupModel) SetVirtualSiteId(v string)`

SetVirtualSiteId sets VirtualSiteId field to given value.


### GetCatalogsForBackup

`func (o *ScheduledBackupModel) GetCatalogsForBackup() []CatalogsForBackup`

GetCatalogsForBackup returns the CatalogsForBackup field if non-nil, zero value otherwise.

### GetCatalogsForBackupOk

`func (o *ScheduledBackupModel) GetCatalogsForBackupOk() (*[]CatalogsForBackup, bool)`

GetCatalogsForBackupOk returns a tuple with the CatalogsForBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogsForBackup

`func (o *ScheduledBackupModel) SetCatalogsForBackup(v []CatalogsForBackup)`

SetCatalogsForBackup sets CatalogsForBackup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


