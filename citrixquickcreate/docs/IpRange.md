# IpRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrIp** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIpRange

`func NewIpRange() *IpRange`

NewIpRange instantiates a new IpRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpRangeWithDefaults

`func NewIpRangeWithDefaults() *IpRange`

NewIpRangeWithDefaults instantiates a new IpRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrIp

`func (o *IpRange) GetCidrIp() string`

GetCidrIp returns the CidrIp field if non-nil, zero value otherwise.

### GetCidrIpOk

`func (o *IpRange) GetCidrIpOk() (*string, bool)`

GetCidrIpOk returns a tuple with the CidrIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIp

`func (o *IpRange) SetCidrIp(v string)`

SetCidrIp sets CidrIp field to given value.

### HasCidrIp

`func (o *IpRange) HasCidrIp() bool`

HasCidrIp returns a boolean if a field has been set.

### SetCidrIpNil

`func (o *IpRange) SetCidrIpNil(b bool)`

 SetCidrIpNil sets the value for CidrIp to be an explicit nil

### UnsetCidrIp
`func (o *IpRange) UnsetCidrIp()`

UnsetCidrIp ensures that no value is present for CidrIp, not even an explicit nil
### GetDescription

`func (o *IpRange) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpRange) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpRange) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpRange) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IpRange) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IpRange) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


