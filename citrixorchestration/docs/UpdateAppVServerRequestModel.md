# UpdateAppVServerRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementServer** | Pointer to **NullableString** | The url of the App-V Management server that packages will be discovered from. | [optional] 
**PublishingServer** | Pointer to **NullableString** | The url of the App-V Publishing server that packages will be discovered from. | [optional] 

## Methods

### NewUpdateAppVServerRequestModel

`func NewUpdateAppVServerRequestModel() *UpdateAppVServerRequestModel`

NewUpdateAppVServerRequestModel instantiates a new UpdateAppVServerRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppVServerRequestModelWithDefaults

`func NewUpdateAppVServerRequestModelWithDefaults() *UpdateAppVServerRequestModel`

NewUpdateAppVServerRequestModelWithDefaults instantiates a new UpdateAppVServerRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementServer

`func (o *UpdateAppVServerRequestModel) GetManagementServer() string`

GetManagementServer returns the ManagementServer field if non-nil, zero value otherwise.

### GetManagementServerOk

`func (o *UpdateAppVServerRequestModel) GetManagementServerOk() (*string, bool)`

GetManagementServerOk returns a tuple with the ManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServer

`func (o *UpdateAppVServerRequestModel) SetManagementServer(v string)`

SetManagementServer sets ManagementServer field to given value.

### HasManagementServer

`func (o *UpdateAppVServerRequestModel) HasManagementServer() bool`

HasManagementServer returns a boolean if a field has been set.

### SetManagementServerNil

`func (o *UpdateAppVServerRequestModel) SetManagementServerNil(b bool)`

 SetManagementServerNil sets the value for ManagementServer to be an explicit nil

### UnsetManagementServer
`func (o *UpdateAppVServerRequestModel) UnsetManagementServer()`

UnsetManagementServer ensures that no value is present for ManagementServer, not even an explicit nil
### GetPublishingServer

`func (o *UpdateAppVServerRequestModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *UpdateAppVServerRequestModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *UpdateAppVServerRequestModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.

### HasPublishingServer

`func (o *UpdateAppVServerRequestModel) HasPublishingServer() bool`

HasPublishingServer returns a boolean if a field has been set.

### SetPublishingServerNil

`func (o *UpdateAppVServerRequestModel) SetPublishingServerNil(b bool)`

 SetPublishingServerNil sets the value for PublishingServer to be an explicit nil

### UnsetPublishingServer
`func (o *UpdateAppVServerRequestModel) UnsetPublishingServer()`

UnsetPublishingServer ensures that no value is present for PublishingServer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


