# AotSettingResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to **NullableString** | The log server name. | [optional] 
**Port** | Pointer to **NullableInt32** | The log server port. | [optional] 
**EnabledOnDeliveryController** | Pointer to **NullableBool** | The AOT is enabled on delivery controller | [optional] 
**EnabledOnAllDeliveryGroups** | Pointer to **NullableBool** | The AOT is enabled on all delivery groups | [optional] 
**EnabledOnDeliveryGroups** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The delivery groups on which AOT is enabled | [optional] 
**LogServerResourceLocationId** | Pointer to **NullableString** | The resource location where the log server locate. | [optional] 
**LogServerApiKey** | Pointer to **NullableString** | The API key to access the log server. | [optional] 

## Methods

### NewAotSettingResponseModel

`func NewAotSettingResponseModel() *AotSettingResponseModel`

NewAotSettingResponseModel instantiates a new AotSettingResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAotSettingResponseModelWithDefaults

`func NewAotSettingResponseModelWithDefaults() *AotSettingResponseModel`

NewAotSettingResponseModelWithDefaults instantiates a new AotSettingResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *AotSettingResponseModel) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AotSettingResponseModel) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AotSettingResponseModel) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *AotSettingResponseModel) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *AotSettingResponseModel) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *AotSettingResponseModel) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetPort

`func (o *AotSettingResponseModel) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AotSettingResponseModel) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AotSettingResponseModel) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *AotSettingResponseModel) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *AotSettingResponseModel) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *AotSettingResponseModel) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetEnabledOnDeliveryController

`func (o *AotSettingResponseModel) GetEnabledOnDeliveryController() bool`

GetEnabledOnDeliveryController returns the EnabledOnDeliveryController field if non-nil, zero value otherwise.

### GetEnabledOnDeliveryControllerOk

`func (o *AotSettingResponseModel) GetEnabledOnDeliveryControllerOk() (*bool, bool)`

GetEnabledOnDeliveryControllerOk returns a tuple with the EnabledOnDeliveryController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnDeliveryController

`func (o *AotSettingResponseModel) SetEnabledOnDeliveryController(v bool)`

SetEnabledOnDeliveryController sets EnabledOnDeliveryController field to given value.

### HasEnabledOnDeliveryController

`func (o *AotSettingResponseModel) HasEnabledOnDeliveryController() bool`

HasEnabledOnDeliveryController returns a boolean if a field has been set.

### SetEnabledOnDeliveryControllerNil

`func (o *AotSettingResponseModel) SetEnabledOnDeliveryControllerNil(b bool)`

 SetEnabledOnDeliveryControllerNil sets the value for EnabledOnDeliveryController to be an explicit nil

### UnsetEnabledOnDeliveryController
`func (o *AotSettingResponseModel) UnsetEnabledOnDeliveryController()`

UnsetEnabledOnDeliveryController ensures that no value is present for EnabledOnDeliveryController, not even an explicit nil
### GetEnabledOnAllDeliveryGroups

`func (o *AotSettingResponseModel) GetEnabledOnAllDeliveryGroups() bool`

GetEnabledOnAllDeliveryGroups returns the EnabledOnAllDeliveryGroups field if non-nil, zero value otherwise.

### GetEnabledOnAllDeliveryGroupsOk

`func (o *AotSettingResponseModel) GetEnabledOnAllDeliveryGroupsOk() (*bool, bool)`

GetEnabledOnAllDeliveryGroupsOk returns a tuple with the EnabledOnAllDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnAllDeliveryGroups

`func (o *AotSettingResponseModel) SetEnabledOnAllDeliveryGroups(v bool)`

SetEnabledOnAllDeliveryGroups sets EnabledOnAllDeliveryGroups field to given value.

### HasEnabledOnAllDeliveryGroups

`func (o *AotSettingResponseModel) HasEnabledOnAllDeliveryGroups() bool`

HasEnabledOnAllDeliveryGroups returns a boolean if a field has been set.

### SetEnabledOnAllDeliveryGroupsNil

`func (o *AotSettingResponseModel) SetEnabledOnAllDeliveryGroupsNil(b bool)`

 SetEnabledOnAllDeliveryGroupsNil sets the value for EnabledOnAllDeliveryGroups to be an explicit nil

### UnsetEnabledOnAllDeliveryGroups
`func (o *AotSettingResponseModel) UnsetEnabledOnAllDeliveryGroups()`

UnsetEnabledOnAllDeliveryGroups ensures that no value is present for EnabledOnAllDeliveryGroups, not even an explicit nil
### GetEnabledOnDeliveryGroups

`func (o *AotSettingResponseModel) GetEnabledOnDeliveryGroups() []RefResponseModel`

GetEnabledOnDeliveryGroups returns the EnabledOnDeliveryGroups field if non-nil, zero value otherwise.

### GetEnabledOnDeliveryGroupsOk

`func (o *AotSettingResponseModel) GetEnabledOnDeliveryGroupsOk() (*[]RefResponseModel, bool)`

GetEnabledOnDeliveryGroupsOk returns a tuple with the EnabledOnDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnDeliveryGroups

`func (o *AotSettingResponseModel) SetEnabledOnDeliveryGroups(v []RefResponseModel)`

SetEnabledOnDeliveryGroups sets EnabledOnDeliveryGroups field to given value.

### HasEnabledOnDeliveryGroups

`func (o *AotSettingResponseModel) HasEnabledOnDeliveryGroups() bool`

HasEnabledOnDeliveryGroups returns a boolean if a field has been set.

### SetEnabledOnDeliveryGroupsNil

`func (o *AotSettingResponseModel) SetEnabledOnDeliveryGroupsNil(b bool)`

 SetEnabledOnDeliveryGroupsNil sets the value for EnabledOnDeliveryGroups to be an explicit nil

### UnsetEnabledOnDeliveryGroups
`func (o *AotSettingResponseModel) UnsetEnabledOnDeliveryGroups()`

UnsetEnabledOnDeliveryGroups ensures that no value is present for EnabledOnDeliveryGroups, not even an explicit nil
### GetLogServerResourceLocationId

`func (o *AotSettingResponseModel) GetLogServerResourceLocationId() string`

GetLogServerResourceLocationId returns the LogServerResourceLocationId field if non-nil, zero value otherwise.

### GetLogServerResourceLocationIdOk

`func (o *AotSettingResponseModel) GetLogServerResourceLocationIdOk() (*string, bool)`

GetLogServerResourceLocationIdOk returns a tuple with the LogServerResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogServerResourceLocationId

`func (o *AotSettingResponseModel) SetLogServerResourceLocationId(v string)`

SetLogServerResourceLocationId sets LogServerResourceLocationId field to given value.

### HasLogServerResourceLocationId

`func (o *AotSettingResponseModel) HasLogServerResourceLocationId() bool`

HasLogServerResourceLocationId returns a boolean if a field has been set.

### SetLogServerResourceLocationIdNil

`func (o *AotSettingResponseModel) SetLogServerResourceLocationIdNil(b bool)`

 SetLogServerResourceLocationIdNil sets the value for LogServerResourceLocationId to be an explicit nil

### UnsetLogServerResourceLocationId
`func (o *AotSettingResponseModel) UnsetLogServerResourceLocationId()`

UnsetLogServerResourceLocationId ensures that no value is present for LogServerResourceLocationId, not even an explicit nil
### GetLogServerApiKey

`func (o *AotSettingResponseModel) GetLogServerApiKey() string`

GetLogServerApiKey returns the LogServerApiKey field if non-nil, zero value otherwise.

### GetLogServerApiKeyOk

`func (o *AotSettingResponseModel) GetLogServerApiKeyOk() (*string, bool)`

GetLogServerApiKeyOk returns a tuple with the LogServerApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogServerApiKey

`func (o *AotSettingResponseModel) SetLogServerApiKey(v string)`

SetLogServerApiKey sets LogServerApiKey field to given value.

### HasLogServerApiKey

`func (o *AotSettingResponseModel) HasLogServerApiKey() bool`

HasLogServerApiKey returns a boolean if a field has been set.

### SetLogServerApiKeyNil

`func (o *AotSettingResponseModel) SetLogServerApiKeyNil(b bool)`

 SetLogServerApiKeyNil sets the value for LogServerApiKey to be an explicit nil

### UnsetLogServerApiKey
`func (o *AotSettingResponseModel) UnsetLogServerApiKey()`

UnsetLogServerApiKey ensures that no value is present for LogServerApiKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


