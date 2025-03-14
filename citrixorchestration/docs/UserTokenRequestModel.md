# UserTokenRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionKey** | **string** | The session key. | 
**Sid** | **string** | The users SID. | 
**VdaIPAddress** | **string** | The IP address of the VDA. | 
**BrokeringTime** | **time.Time** | The time at which the session is brokered. | 

## Methods

### NewUserTokenRequestModel

`func NewUserTokenRequestModel(sessionKey string, sid string, vdaIPAddress string, brokeringTime time.Time, ) *UserTokenRequestModel`

NewUserTokenRequestModel instantiates a new UserTokenRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTokenRequestModelWithDefaults

`func NewUserTokenRequestModelWithDefaults() *UserTokenRequestModel`

NewUserTokenRequestModelWithDefaults instantiates a new UserTokenRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionKey

`func (o *UserTokenRequestModel) GetSessionKey() string`

GetSessionKey returns the SessionKey field if non-nil, zero value otherwise.

### GetSessionKeyOk

`func (o *UserTokenRequestModel) GetSessionKeyOk() (*string, bool)`

GetSessionKeyOk returns a tuple with the SessionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionKey

`func (o *UserTokenRequestModel) SetSessionKey(v string)`

SetSessionKey sets SessionKey field to given value.


### GetSid

`func (o *UserTokenRequestModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *UserTokenRequestModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *UserTokenRequestModel) SetSid(v string)`

SetSid sets Sid field to given value.


### GetVdaIPAddress

`func (o *UserTokenRequestModel) GetVdaIPAddress() string`

GetVdaIPAddress returns the VdaIPAddress field if non-nil, zero value otherwise.

### GetVdaIPAddressOk

`func (o *UserTokenRequestModel) GetVdaIPAddressOk() (*string, bool)`

GetVdaIPAddressOk returns a tuple with the VdaIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaIPAddress

`func (o *UserTokenRequestModel) SetVdaIPAddress(v string)`

SetVdaIPAddress sets VdaIPAddress field to given value.


### GetBrokeringTime

`func (o *UserTokenRequestModel) GetBrokeringTime() time.Time`

GetBrokeringTime returns the BrokeringTime field if non-nil, zero value otherwise.

### GetBrokeringTimeOk

`func (o *UserTokenRequestModel) GetBrokeringTimeOk() (*time.Time, bool)`

GetBrokeringTimeOk returns a tuple with the BrokeringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokeringTime

`func (o *UserTokenRequestModel) SetBrokeringTime(v time.Time)`

SetBrokeringTime sets BrokeringTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


