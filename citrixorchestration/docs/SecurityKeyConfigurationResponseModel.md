# SecurityKeyConfigurationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key1** | **string** | The first security key. | 
**Key2** | **string** | The second security key. | 
**RequireKeyForSta** | **bool** | Indicates whether to require a key to authenticate communications over the STA port. | 
**RequireKeyForXml** | **bool** | Indicates whether to require a key to authenticate communications over the XML port. | 

## Methods

### NewSecurityKeyConfigurationResponseModel

`func NewSecurityKeyConfigurationResponseModel(key1 string, key2 string, requireKeyForSta bool, requireKeyForXml bool, ) *SecurityKeyConfigurationResponseModel`

NewSecurityKeyConfigurationResponseModel instantiates a new SecurityKeyConfigurationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityKeyConfigurationResponseModelWithDefaults

`func NewSecurityKeyConfigurationResponseModelWithDefaults() *SecurityKeyConfigurationResponseModel`

NewSecurityKeyConfigurationResponseModelWithDefaults instantiates a new SecurityKeyConfigurationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey1

`func (o *SecurityKeyConfigurationResponseModel) GetKey1() string`

GetKey1 returns the Key1 field if non-nil, zero value otherwise.

### GetKey1Ok

`func (o *SecurityKeyConfigurationResponseModel) GetKey1Ok() (*string, bool)`

GetKey1Ok returns a tuple with the Key1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey1

`func (o *SecurityKeyConfigurationResponseModel) SetKey1(v string)`

SetKey1 sets Key1 field to given value.


### GetKey2

`func (o *SecurityKeyConfigurationResponseModel) GetKey2() string`

GetKey2 returns the Key2 field if non-nil, zero value otherwise.

### GetKey2Ok

`func (o *SecurityKeyConfigurationResponseModel) GetKey2Ok() (*string, bool)`

GetKey2Ok returns a tuple with the Key2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey2

`func (o *SecurityKeyConfigurationResponseModel) SetKey2(v string)`

SetKey2 sets Key2 field to given value.


### GetRequireKeyForSta

`func (o *SecurityKeyConfigurationResponseModel) GetRequireKeyForSta() bool`

GetRequireKeyForSta returns the RequireKeyForSta field if non-nil, zero value otherwise.

### GetRequireKeyForStaOk

`func (o *SecurityKeyConfigurationResponseModel) GetRequireKeyForStaOk() (*bool, bool)`

GetRequireKeyForStaOk returns a tuple with the RequireKeyForSta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireKeyForSta

`func (o *SecurityKeyConfigurationResponseModel) SetRequireKeyForSta(v bool)`

SetRequireKeyForSta sets RequireKeyForSta field to given value.


### GetRequireKeyForXml

`func (o *SecurityKeyConfigurationResponseModel) GetRequireKeyForXml() bool`

GetRequireKeyForXml returns the RequireKeyForXml field if non-nil, zero value otherwise.

### GetRequireKeyForXmlOk

`func (o *SecurityKeyConfigurationResponseModel) GetRequireKeyForXmlOk() (*bool, bool)`

GetRequireKeyForXmlOk returns a tuple with the RequireKeyForXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireKeyForXml

`func (o *SecurityKeyConfigurationResponseModel) SetRequireKeyForXml(v bool)`

SetRequireKeyForXml sets RequireKeyForXml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


