# SecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPort** | **int32** | Gets or sets the FromPort. | 
**SecurityGroupIds** | **[]string** | Gets the SecurityGroupIds. | 
**IPRanges** | **[]string** | Gets the IPRanges. | 
**Protocol** | **string** | Gets or sets the Protocol. | 
**ToPort** | **int32** | Gets or sets the ToPort. | 

## Methods

### NewSecurityGroupRule

`func NewSecurityGroupRule(fromPort int32, securityGroupIds []string, iPRanges []string, protocol string, toPort int32, ) *SecurityGroupRule`

NewSecurityGroupRule instantiates a new SecurityGroupRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRuleWithDefaults

`func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule`

NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPort

`func (o *SecurityGroupRule) GetFromPort() int32`

GetFromPort returns the FromPort field if non-nil, zero value otherwise.

### GetFromPortOk

`func (o *SecurityGroupRule) GetFromPortOk() (*int32, bool)`

GetFromPortOk returns a tuple with the FromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPort

`func (o *SecurityGroupRule) SetFromPort(v int32)`

SetFromPort sets FromPort field to given value.


### GetSecurityGroupIds

`func (o *SecurityGroupRule) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *SecurityGroupRule) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *SecurityGroupRule) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.


### GetIPRanges

`func (o *SecurityGroupRule) GetIPRanges() []string`

GetIPRanges returns the IPRanges field if non-nil, zero value otherwise.

### GetIPRangesOk

`func (o *SecurityGroupRule) GetIPRangesOk() (*[]string, bool)`

GetIPRangesOk returns a tuple with the IPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPRanges

`func (o *SecurityGroupRule) SetIPRanges(v []string)`

SetIPRanges sets IPRanges field to given value.


### GetProtocol

`func (o *SecurityGroupRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetToPort

`func (o *SecurityGroupRule) GetToPort() int32`

GetToPort returns the ToPort field if non-nil, zero value otherwise.

### GetToPortOk

`func (o *SecurityGroupRule) GetToPortOk() (*int32, bool)`

GetToPortOk returns a tuple with the ToPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPort

`func (o *SecurityGroupRule) SetToPort(v int32)`

SetToPort sets ToPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


