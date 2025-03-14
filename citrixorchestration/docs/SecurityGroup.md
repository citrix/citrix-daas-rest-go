# SecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Gets or sets the Description. | 
**EgressRules** | [**[]SecurityGroupRule**](SecurityGroupRule.md) | Gets the EgressRules. | 
**IngressRules** | [**[]SecurityGroupRule**](SecurityGroupRule.md) | Gets the IngressRules. | 
**VirtualPrivateCloudId** | **string** | Gets or sets the VirtualPrivateCloudId. | 

## Methods

### NewSecurityGroup

`func NewSecurityGroup(description string, egressRules []SecurityGroupRule, ingressRules []SecurityGroupRule, virtualPrivateCloudId string, ) *SecurityGroup`

NewSecurityGroup instantiates a new SecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupWithDefaults

`func NewSecurityGroupWithDefaults() *SecurityGroup`

NewSecurityGroupWithDefaults instantiates a new SecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEgressRules

`func (o *SecurityGroup) GetEgressRules() []SecurityGroupRule`

GetEgressRules returns the EgressRules field if non-nil, zero value otherwise.

### GetEgressRulesOk

`func (o *SecurityGroup) GetEgressRulesOk() (*[]SecurityGroupRule, bool)`

GetEgressRulesOk returns a tuple with the EgressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressRules

`func (o *SecurityGroup) SetEgressRules(v []SecurityGroupRule)`

SetEgressRules sets EgressRules field to given value.


### GetIngressRules

`func (o *SecurityGroup) GetIngressRules() []SecurityGroupRule`

GetIngressRules returns the IngressRules field if non-nil, zero value otherwise.

### GetIngressRulesOk

`func (o *SecurityGroup) GetIngressRulesOk() (*[]SecurityGroupRule, bool)`

GetIngressRulesOk returns a tuple with the IngressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRules

`func (o *SecurityGroup) SetIngressRules(v []SecurityGroupRule)`

SetIngressRules sets IngressRules field to given value.


### GetVirtualPrivateCloudId

`func (o *SecurityGroup) GetVirtualPrivateCloudId() string`

GetVirtualPrivateCloudId returns the VirtualPrivateCloudId field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudIdOk

`func (o *SecurityGroup) GetVirtualPrivateCloudIdOk() (*string, bool)`

GetVirtualPrivateCloudIdOk returns a tuple with the VirtualPrivateCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloudId

`func (o *SecurityGroup) SetVirtualPrivateCloudId(v string)`

SetVirtualPrivateCloudId sets VirtualPrivateCloudId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


