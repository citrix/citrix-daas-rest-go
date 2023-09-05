# AddAppVServerRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementServer** | **string** | The url of the App-V Management server that packages will be discovered from. | 
**PublishingServer** | **string** | The url of the App-V Publishing server that packages will be discovered from. | 

## Methods

### NewAddAppVServerRequestModel

`func NewAddAppVServerRequestModel(managementServer string, publishingServer string, ) *AddAppVServerRequestModel`

NewAddAppVServerRequestModel instantiates a new AddAppVServerRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAppVServerRequestModelWithDefaults

`func NewAddAppVServerRequestModelWithDefaults() *AddAppVServerRequestModel`

NewAddAppVServerRequestModelWithDefaults instantiates a new AddAppVServerRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementServer

`func (o *AddAppVServerRequestModel) GetManagementServer() string`

GetManagementServer returns the ManagementServer field if non-nil, zero value otherwise.

### GetManagementServerOk

`func (o *AddAppVServerRequestModel) GetManagementServerOk() (*string, bool)`

GetManagementServerOk returns a tuple with the ManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServer

`func (o *AddAppVServerRequestModel) SetManagementServer(v string)`

SetManagementServer sets ManagementServer field to given value.


### GetPublishingServer

`func (o *AddAppVServerRequestModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *AddAppVServerRequestModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *AddAppVServerRequestModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


