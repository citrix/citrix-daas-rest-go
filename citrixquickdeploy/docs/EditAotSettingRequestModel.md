# EditAotSettingRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to **NullableString** | The log server name. | [optional] 
**Port** | Pointer to **NullableInt32** | The log server port. | [optional] 
**EnabledOnDeliveryController** | Pointer to **NullableBool** | Enable AOT on delivery controller. | [optional] 
**EnabledOnAllDeliveryGroups** | Pointer to **NullableBool** | Enable AOT on all delivery groups. | [optional] 
**EnabledOnDeliveryGroups** | Pointer to **[]string** | Enable AOT on the specified delivery groups. | [optional] 
**LogServerResourceLocationId** | Pointer to **NullableString** | The resource location where the log server locate. | [optional] 
**LogServerApiKey** | Pointer to **NullableString** | The API key to access the log server. | [optional] 

## Methods

### NewEditAotSettingRequestModel

`func NewEditAotSettingRequestModel() *EditAotSettingRequestModel`

NewEditAotSettingRequestModel instantiates a new EditAotSettingRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAotSettingRequestModelWithDefaults

`func NewEditAotSettingRequestModelWithDefaults() *EditAotSettingRequestModel`

NewEditAotSettingRequestModelWithDefaults instantiates a new EditAotSettingRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *EditAotSettingRequestModel) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *EditAotSettingRequestModel) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *EditAotSettingRequestModel) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *EditAotSettingRequestModel) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *EditAotSettingRequestModel) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *EditAotSettingRequestModel) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetPort

`func (o *EditAotSettingRequestModel) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EditAotSettingRequestModel) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EditAotSettingRequestModel) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *EditAotSettingRequestModel) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *EditAotSettingRequestModel) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *EditAotSettingRequestModel) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetEnabledOnDeliveryController

`func (o *EditAotSettingRequestModel) GetEnabledOnDeliveryController() bool`

GetEnabledOnDeliveryController returns the EnabledOnDeliveryController field if non-nil, zero value otherwise.

### GetEnabledOnDeliveryControllerOk

`func (o *EditAotSettingRequestModel) GetEnabledOnDeliveryControllerOk() (*bool, bool)`

GetEnabledOnDeliveryControllerOk returns a tuple with the EnabledOnDeliveryController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnDeliveryController

`func (o *EditAotSettingRequestModel) SetEnabledOnDeliveryController(v bool)`

SetEnabledOnDeliveryController sets EnabledOnDeliveryController field to given value.

### HasEnabledOnDeliveryController

`func (o *EditAotSettingRequestModel) HasEnabledOnDeliveryController() bool`

HasEnabledOnDeliveryController returns a boolean if a field has been set.

### SetEnabledOnDeliveryControllerNil

`func (o *EditAotSettingRequestModel) SetEnabledOnDeliveryControllerNil(b bool)`

 SetEnabledOnDeliveryControllerNil sets the value for EnabledOnDeliveryController to be an explicit nil

### UnsetEnabledOnDeliveryController
`func (o *EditAotSettingRequestModel) UnsetEnabledOnDeliveryController()`

UnsetEnabledOnDeliveryController ensures that no value is present for EnabledOnDeliveryController, not even an explicit nil
### GetEnabledOnAllDeliveryGroups

`func (o *EditAotSettingRequestModel) GetEnabledOnAllDeliveryGroups() bool`

GetEnabledOnAllDeliveryGroups returns the EnabledOnAllDeliveryGroups field if non-nil, zero value otherwise.

### GetEnabledOnAllDeliveryGroupsOk

`func (o *EditAotSettingRequestModel) GetEnabledOnAllDeliveryGroupsOk() (*bool, bool)`

GetEnabledOnAllDeliveryGroupsOk returns a tuple with the EnabledOnAllDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnAllDeliveryGroups

`func (o *EditAotSettingRequestModel) SetEnabledOnAllDeliveryGroups(v bool)`

SetEnabledOnAllDeliveryGroups sets EnabledOnAllDeliveryGroups field to given value.

### HasEnabledOnAllDeliveryGroups

`func (o *EditAotSettingRequestModel) HasEnabledOnAllDeliveryGroups() bool`

HasEnabledOnAllDeliveryGroups returns a boolean if a field has been set.

### SetEnabledOnAllDeliveryGroupsNil

`func (o *EditAotSettingRequestModel) SetEnabledOnAllDeliveryGroupsNil(b bool)`

 SetEnabledOnAllDeliveryGroupsNil sets the value for EnabledOnAllDeliveryGroups to be an explicit nil

### UnsetEnabledOnAllDeliveryGroups
`func (o *EditAotSettingRequestModel) UnsetEnabledOnAllDeliveryGroups()`

UnsetEnabledOnAllDeliveryGroups ensures that no value is present for EnabledOnAllDeliveryGroups, not even an explicit nil
### GetEnabledOnDeliveryGroups

`func (o *EditAotSettingRequestModel) GetEnabledOnDeliveryGroups() []string`

GetEnabledOnDeliveryGroups returns the EnabledOnDeliveryGroups field if non-nil, zero value otherwise.

### GetEnabledOnDeliveryGroupsOk

`func (o *EditAotSettingRequestModel) GetEnabledOnDeliveryGroupsOk() (*[]string, bool)`

GetEnabledOnDeliveryGroupsOk returns a tuple with the EnabledOnDeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnDeliveryGroups

`func (o *EditAotSettingRequestModel) SetEnabledOnDeliveryGroups(v []string)`

SetEnabledOnDeliveryGroups sets EnabledOnDeliveryGroups field to given value.

### HasEnabledOnDeliveryGroups

`func (o *EditAotSettingRequestModel) HasEnabledOnDeliveryGroups() bool`

HasEnabledOnDeliveryGroups returns a boolean if a field has been set.

### SetEnabledOnDeliveryGroupsNil

`func (o *EditAotSettingRequestModel) SetEnabledOnDeliveryGroupsNil(b bool)`

 SetEnabledOnDeliveryGroupsNil sets the value for EnabledOnDeliveryGroups to be an explicit nil

### UnsetEnabledOnDeliveryGroups
`func (o *EditAotSettingRequestModel) UnsetEnabledOnDeliveryGroups()`

UnsetEnabledOnDeliveryGroups ensures that no value is present for EnabledOnDeliveryGroups, not even an explicit nil
### GetLogServerResourceLocationId

`func (o *EditAotSettingRequestModel) GetLogServerResourceLocationId() string`

GetLogServerResourceLocationId returns the LogServerResourceLocationId field if non-nil, zero value otherwise.

### GetLogServerResourceLocationIdOk

`func (o *EditAotSettingRequestModel) GetLogServerResourceLocationIdOk() (*string, bool)`

GetLogServerResourceLocationIdOk returns a tuple with the LogServerResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogServerResourceLocationId

`func (o *EditAotSettingRequestModel) SetLogServerResourceLocationId(v string)`

SetLogServerResourceLocationId sets LogServerResourceLocationId field to given value.

### HasLogServerResourceLocationId

`func (o *EditAotSettingRequestModel) HasLogServerResourceLocationId() bool`

HasLogServerResourceLocationId returns a boolean if a field has been set.

### SetLogServerResourceLocationIdNil

`func (o *EditAotSettingRequestModel) SetLogServerResourceLocationIdNil(b bool)`

 SetLogServerResourceLocationIdNil sets the value for LogServerResourceLocationId to be an explicit nil

### UnsetLogServerResourceLocationId
`func (o *EditAotSettingRequestModel) UnsetLogServerResourceLocationId()`

UnsetLogServerResourceLocationId ensures that no value is present for LogServerResourceLocationId, not even an explicit nil
### GetLogServerApiKey

`func (o *EditAotSettingRequestModel) GetLogServerApiKey() string`

GetLogServerApiKey returns the LogServerApiKey field if non-nil, zero value otherwise.

### GetLogServerApiKeyOk

`func (o *EditAotSettingRequestModel) GetLogServerApiKeyOk() (*string, bool)`

GetLogServerApiKeyOk returns a tuple with the LogServerApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogServerApiKey

`func (o *EditAotSettingRequestModel) SetLogServerApiKey(v string)`

SetLogServerApiKey sets LogServerApiKey field to given value.

### HasLogServerApiKey

`func (o *EditAotSettingRequestModel) HasLogServerApiKey() bool`

HasLogServerApiKey returns a boolean if a field has been set.

### SetLogServerApiKeyNil

`func (o *EditAotSettingRequestModel) SetLogServerApiKeyNil(b bool)`

 SetLogServerApiKeyNil sets the value for LogServerApiKey to be an explicit nil

### UnsetLogServerApiKey
`func (o *EditAotSettingRequestModel) UnsetLogServerApiKey()`

UnsetLogServerApiKey ensures that no value is present for LogServerApiKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


