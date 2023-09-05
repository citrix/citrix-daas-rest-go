# AppVServerResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** | Uid of App-V server. | [optional] 
**Name** | **string** | Name of App-V server. | 
**ManagementServer** | **string** | Management server address.  (i.e. where packages/apps are enumerated from.) | 
**PublishingServer** | **string** | Publishing server address.  (i.e. where the package streams from.) | 
**IsStudioConfigured** | Pointer to **bool** | Is studio configured or not. | [optional] 

## Methods

### NewAppVServerResponseModel

`func NewAppVServerResponseModel(name string, managementServer string, publishingServer string, ) *AppVServerResponseModel`

NewAppVServerResponseModel instantiates a new AppVServerResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVServerResponseModelWithDefaults

`func NewAppVServerResponseModelWithDefaults() *AppVServerResponseModel`

NewAppVServerResponseModelWithDefaults instantiates a new AppVServerResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *AppVServerResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AppVServerResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AppVServerResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AppVServerResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *AppVServerResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppVServerResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppVServerResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetManagementServer

`func (o *AppVServerResponseModel) GetManagementServer() string`

GetManagementServer returns the ManagementServer field if non-nil, zero value otherwise.

### GetManagementServerOk

`func (o *AppVServerResponseModel) GetManagementServerOk() (*string, bool)`

GetManagementServerOk returns a tuple with the ManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServer

`func (o *AppVServerResponseModel) SetManagementServer(v string)`

SetManagementServer sets ManagementServer field to given value.


### GetPublishingServer

`func (o *AppVServerResponseModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *AppVServerResponseModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *AppVServerResponseModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.


### GetIsStudioConfigured

`func (o *AppVServerResponseModel) GetIsStudioConfigured() bool`

GetIsStudioConfigured returns the IsStudioConfigured field if non-nil, zero value otherwise.

### GetIsStudioConfiguredOk

`func (o *AppVServerResponseModel) GetIsStudioConfiguredOk() (*bool, bool)`

GetIsStudioConfiguredOk returns a tuple with the IsStudioConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStudioConfigured

`func (o *AppVServerResponseModel) SetIsStudioConfigured(v bool)`

SetIsStudioConfigured sets IsStudioConfigured field to given value.

### HasIsStudioConfigured

`func (o *AppVServerResponseModel) HasIsStudioConfigured() bool`

HasIsStudioConfigured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


