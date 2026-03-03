# ConnectivityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceLocationId** | **string** |  | 
**ResourceLocationName** | Pointer to **NullableString** |  | [optional] 
**Type** | [**ConnectivityType**](ConnectivityType.md) |  | 
**NetscalerOnPremisesUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConnectivityModel

`func NewConnectivityModel(resourceLocationId string, type_ ConnectivityType, ) *ConnectivityModel`

NewConnectivityModel instantiates a new ConnectivityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityModelWithDefaults

`func NewConnectivityModelWithDefaults() *ConnectivityModel`

NewConnectivityModelWithDefaults instantiates a new ConnectivityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceLocationId

`func (o *ConnectivityModel) GetResourceLocationId() string`

GetResourceLocationId returns the ResourceLocationId field if non-nil, zero value otherwise.

### GetResourceLocationIdOk

`func (o *ConnectivityModel) GetResourceLocationIdOk() (*string, bool)`

GetResourceLocationIdOk returns a tuple with the ResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocationId

`func (o *ConnectivityModel) SetResourceLocationId(v string)`

SetResourceLocationId sets ResourceLocationId field to given value.


### GetResourceLocationName

`func (o *ConnectivityModel) GetResourceLocationName() string`

GetResourceLocationName returns the ResourceLocationName field if non-nil, zero value otherwise.

### GetResourceLocationNameOk

`func (o *ConnectivityModel) GetResourceLocationNameOk() (*string, bool)`

GetResourceLocationNameOk returns a tuple with the ResourceLocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocationName

`func (o *ConnectivityModel) SetResourceLocationName(v string)`

SetResourceLocationName sets ResourceLocationName field to given value.

### HasResourceLocationName

`func (o *ConnectivityModel) HasResourceLocationName() bool`

HasResourceLocationName returns a boolean if a field has been set.

### SetResourceLocationNameNil

`func (o *ConnectivityModel) SetResourceLocationNameNil(b bool)`

 SetResourceLocationNameNil sets the value for ResourceLocationName to be an explicit nil

### UnsetResourceLocationName
`func (o *ConnectivityModel) UnsetResourceLocationName()`

UnsetResourceLocationName ensures that no value is present for ResourceLocationName, not even an explicit nil
### GetType

`func (o *ConnectivityModel) GetType() ConnectivityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectivityModel) GetTypeOk() (*ConnectivityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectivityModel) SetType(v ConnectivityType)`

SetType sets Type field to given value.


### GetNetscalerOnPremisesUrl

`func (o *ConnectivityModel) GetNetscalerOnPremisesUrl() string`

GetNetscalerOnPremisesUrl returns the NetscalerOnPremisesUrl field if non-nil, zero value otherwise.

### GetNetscalerOnPremisesUrlOk

`func (o *ConnectivityModel) GetNetscalerOnPremisesUrlOk() (*string, bool)`

GetNetscalerOnPremisesUrlOk returns a tuple with the NetscalerOnPremisesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetscalerOnPremisesUrl

`func (o *ConnectivityModel) SetNetscalerOnPremisesUrl(v string)`

SetNetscalerOnPremisesUrl sets NetscalerOnPremisesUrl field to given value.

### HasNetscalerOnPremisesUrl

`func (o *ConnectivityModel) HasNetscalerOnPremisesUrl() bool`

HasNetscalerOnPremisesUrl returns a boolean if a field has been set.

### SetNetscalerOnPremisesUrlNil

`func (o *ConnectivityModel) SetNetscalerOnPremisesUrlNil(b bool)`

 SetNetscalerOnPremisesUrlNil sets the value for NetscalerOnPremisesUrl to be an explicit nil

### UnsetNetscalerOnPremisesUrl
`func (o *ConnectivityModel) UnsetNetscalerOnPremisesUrl()`

UnsetNetscalerOnPremisesUrl ensures that no value is present for NetscalerOnPremisesUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


