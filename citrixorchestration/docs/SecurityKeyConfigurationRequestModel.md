# SecurityKeyConfigurationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key1** | Pointer to **string** | The first security key. | [optional] 
**Key2** | Pointer to **string** | The second security key. | [optional] 
**RequireKeyForSta** | Pointer to **bool** | Indicates whether to require a key to authenticate communications over the STA port. | [optional] 
**RequireKeyForXml** | Pointer to **bool** | Indicates whether to require a key to authenticate communications over the XML port. | [optional] 

## Methods

### NewSecurityKeyConfigurationRequestModel

`func NewSecurityKeyConfigurationRequestModel() *SecurityKeyConfigurationRequestModel`

NewSecurityKeyConfigurationRequestModel instantiates a new SecurityKeyConfigurationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityKeyConfigurationRequestModelWithDefaults

`func NewSecurityKeyConfigurationRequestModelWithDefaults() *SecurityKeyConfigurationRequestModel`

NewSecurityKeyConfigurationRequestModelWithDefaults instantiates a new SecurityKeyConfigurationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey1

`func (o *SecurityKeyConfigurationRequestModel) GetKey1() string`

GetKey1 returns the Key1 field if non-nil, zero value otherwise.

### GetKey1Ok

`func (o *SecurityKeyConfigurationRequestModel) GetKey1Ok() (*string, bool)`

GetKey1Ok returns a tuple with the Key1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey1

`func (o *SecurityKeyConfigurationRequestModel) SetKey1(v string)`

SetKey1 sets Key1 field to given value.

### HasKey1

`func (o *SecurityKeyConfigurationRequestModel) HasKey1() bool`

HasKey1 returns a boolean if a field has been set.

### GetKey2

`func (o *SecurityKeyConfigurationRequestModel) GetKey2() string`

GetKey2 returns the Key2 field if non-nil, zero value otherwise.

### GetKey2Ok

`func (o *SecurityKeyConfigurationRequestModel) GetKey2Ok() (*string, bool)`

GetKey2Ok returns a tuple with the Key2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey2

`func (o *SecurityKeyConfigurationRequestModel) SetKey2(v string)`

SetKey2 sets Key2 field to given value.

### HasKey2

`func (o *SecurityKeyConfigurationRequestModel) HasKey2() bool`

HasKey2 returns a boolean if a field has been set.

### GetRequireKeyForSta

`func (o *SecurityKeyConfigurationRequestModel) GetRequireKeyForSta() bool`

GetRequireKeyForSta returns the RequireKeyForSta field if non-nil, zero value otherwise.

### GetRequireKeyForStaOk

`func (o *SecurityKeyConfigurationRequestModel) GetRequireKeyForStaOk() (*bool, bool)`

GetRequireKeyForStaOk returns a tuple with the RequireKeyForSta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireKeyForSta

`func (o *SecurityKeyConfigurationRequestModel) SetRequireKeyForSta(v bool)`

SetRequireKeyForSta sets RequireKeyForSta field to given value.

### HasRequireKeyForSta

`func (o *SecurityKeyConfigurationRequestModel) HasRequireKeyForSta() bool`

HasRequireKeyForSta returns a boolean if a field has been set.

### GetRequireKeyForXml

`func (o *SecurityKeyConfigurationRequestModel) GetRequireKeyForXml() bool`

GetRequireKeyForXml returns the RequireKeyForXml field if non-nil, zero value otherwise.

### GetRequireKeyForXmlOk

`func (o *SecurityKeyConfigurationRequestModel) GetRequireKeyForXmlOk() (*bool, bool)`

GetRequireKeyForXmlOk returns a tuple with the RequireKeyForXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireKeyForXml

`func (o *SecurityKeyConfigurationRequestModel) SetRequireKeyForXml(v bool)`

SetRequireKeyForXml sets RequireKeyForXml field to given value.

### HasRequireKeyForXml

`func (o *SecurityKeyConfigurationRequestModel) HasRequireKeyForXml() bool`

HasRequireKeyForXml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


