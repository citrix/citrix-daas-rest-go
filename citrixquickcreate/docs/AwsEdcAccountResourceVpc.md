# AwsEdcAccountResourceVpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcId** | Pointer to **NullableString** | Vpc Id | [optional] 
**Status** | Pointer to [**NullableAwsEdcVpcState**](AwsEdcVpcState.md) |  | [optional] 
**VpcTenancy** | Pointer to [**NullableAwsEdcVpcTenancy**](AwsEdcVpcTenancy.md) |  | [optional] 
**CidrBlock** | Pointer to **NullableString** | Vpc CIDR Block | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Vpc Tags | [optional] 

## Methods

### NewAwsEdcAccountResourceVpc

`func NewAwsEdcAccountResourceVpc() *AwsEdcAccountResourceVpc`

NewAwsEdcAccountResourceVpc instantiates a new AwsEdcAccountResourceVpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcAccountResourceVpcWithDefaults

`func NewAwsEdcAccountResourceVpcWithDefaults() *AwsEdcAccountResourceVpc`

NewAwsEdcAccountResourceVpcWithDefaults instantiates a new AwsEdcAccountResourceVpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcId

`func (o *AwsEdcAccountResourceVpc) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsEdcAccountResourceVpc) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsEdcAccountResourceVpc) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsEdcAccountResourceVpc) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *AwsEdcAccountResourceVpc) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AwsEdcAccountResourceVpc) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetStatus

`func (o *AwsEdcAccountResourceVpc) GetStatus() AwsEdcVpcState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsEdcAccountResourceVpc) GetStatusOk() (*AwsEdcVpcState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsEdcAccountResourceVpc) SetStatus(v AwsEdcVpcState)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsEdcAccountResourceVpc) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AwsEdcAccountResourceVpc) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AwsEdcAccountResourceVpc) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetVpcTenancy

`func (o *AwsEdcAccountResourceVpc) GetVpcTenancy() AwsEdcVpcTenancy`

GetVpcTenancy returns the VpcTenancy field if non-nil, zero value otherwise.

### GetVpcTenancyOk

`func (o *AwsEdcAccountResourceVpc) GetVpcTenancyOk() (*AwsEdcVpcTenancy, bool)`

GetVpcTenancyOk returns a tuple with the VpcTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcTenancy

`func (o *AwsEdcAccountResourceVpc) SetVpcTenancy(v AwsEdcVpcTenancy)`

SetVpcTenancy sets VpcTenancy field to given value.

### HasVpcTenancy

`func (o *AwsEdcAccountResourceVpc) HasVpcTenancy() bool`

HasVpcTenancy returns a boolean if a field has been set.

### SetVpcTenancyNil

`func (o *AwsEdcAccountResourceVpc) SetVpcTenancyNil(b bool)`

 SetVpcTenancyNil sets the value for VpcTenancy to be an explicit nil

### UnsetVpcTenancy
`func (o *AwsEdcAccountResourceVpc) UnsetVpcTenancy()`

UnsetVpcTenancy ensures that no value is present for VpcTenancy, not even an explicit nil
### GetCidrBlock

`func (o *AwsEdcAccountResourceVpc) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AwsEdcAccountResourceVpc) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AwsEdcAccountResourceVpc) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *AwsEdcAccountResourceVpc) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### SetCidrBlockNil

`func (o *AwsEdcAccountResourceVpc) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *AwsEdcAccountResourceVpc) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
### GetTags

`func (o *AwsEdcAccountResourceVpc) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwsEdcAccountResourceVpc) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwsEdcAccountResourceVpc) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwsEdcAccountResourceVpc) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AwsEdcAccountResourceVpc) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AwsEdcAccountResourceVpc) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


