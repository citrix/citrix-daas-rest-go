# HypervisorSecurityGroupResourceResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the security group. | [optional] 
**EgressRules** | [**[]HypervisorSecurityGroupRuleResponseModel**](HypervisorSecurityGroupRuleResponseModel.md) | Egress rules governing outbound network traffic. | 
**IngressRules** | [**[]HypervisorSecurityGroupRuleResponseModel**](HypervisorSecurityGroupRuleResponseModel.md) | Ingress rules governing inbound network traffic. | 
**VirtualPrivateCloudId** | Pointer to **string** | The identifier of the associated VPC. | [optional] 

## Methods

### NewHypervisorSecurityGroupResourceResponseModelAllOf

`func NewHypervisorSecurityGroupResourceResponseModelAllOf(egressRules []HypervisorSecurityGroupRuleResponseModel, ingressRules []HypervisorSecurityGroupRuleResponseModel, ) *HypervisorSecurityGroupResourceResponseModelAllOf`

NewHypervisorSecurityGroupResourceResponseModelAllOf instantiates a new HypervisorSecurityGroupResourceResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorSecurityGroupResourceResponseModelAllOfWithDefaults

`func NewHypervisorSecurityGroupResourceResponseModelAllOfWithDefaults() *HypervisorSecurityGroupResourceResponseModelAllOf`

NewHypervisorSecurityGroupResourceResponseModelAllOfWithDefaults instantiates a new HypervisorSecurityGroupResourceResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEgressRules

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) GetEgressRules() []HypervisorSecurityGroupRuleResponseModel`

GetEgressRules returns the EgressRules field if non-nil, zero value otherwise.

### GetEgressRulesOk

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) GetEgressRulesOk() (*[]HypervisorSecurityGroupRuleResponseModel, bool)`

GetEgressRulesOk returns a tuple with the EgressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressRules

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) SetEgressRules(v []HypervisorSecurityGroupRuleResponseModel)`

SetEgressRules sets EgressRules field to given value.


### GetIngressRules

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) GetIngressRules() []HypervisorSecurityGroupRuleResponseModel`

GetIngressRules returns the IngressRules field if non-nil, zero value otherwise.

### GetIngressRulesOk

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) GetIngressRulesOk() (*[]HypervisorSecurityGroupRuleResponseModel, bool)`

GetIngressRulesOk returns a tuple with the IngressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRules

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) SetIngressRules(v []HypervisorSecurityGroupRuleResponseModel)`

SetIngressRules sets IngressRules field to given value.


### GetVirtualPrivateCloudId

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) GetVirtualPrivateCloudId() string`

GetVirtualPrivateCloudId returns the VirtualPrivateCloudId field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudIdOk

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) GetVirtualPrivateCloudIdOk() (*string, bool)`

GetVirtualPrivateCloudIdOk returns a tuple with the VirtualPrivateCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloudId

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) SetVirtualPrivateCloudId(v string)`

SetVirtualPrivateCloudId sets VirtualPrivateCloudId field to given value.

### HasVirtualPrivateCloudId

`func (o *HypervisorSecurityGroupResourceResponseModelAllOf) HasVirtualPrivateCloudId() bool`

HasVirtualPrivateCloudId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


