# HypervisorSecurityGroupRuleResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPort** | Pointer to **float32** | Start of the port number range, required for any protocol that uses ports. In AWS this can also be an ICMP port number, where -1 specifies all ICMP types. | [optional] 
**GroupIds** | Pointer to **[]string** | IDs of source or destination security groups. Empty when CIDR ranges are specified. | [optional] 
**IPRanges** | Pointer to **[]string** | Source or destination CIDR address ranges. Empty if group IDs are specified. | [optional] 
**Protocol** | **string** | IP protocol name or number. In AWS this is the IP protocol name or number, where -1 specifies all protocols. Protocol numbers: http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xml | 
**ToPort** | Pointer to **float32** | End of the port number range, required for any protocol that uses ports. In AWS this can also be an ICMP port number, where -1 specfies all ICMP types. | [optional] 

## Methods

### NewHypervisorSecurityGroupRuleResponseModel

`func NewHypervisorSecurityGroupRuleResponseModel(protocol string, ) *HypervisorSecurityGroupRuleResponseModel`

NewHypervisorSecurityGroupRuleResponseModel instantiates a new HypervisorSecurityGroupRuleResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorSecurityGroupRuleResponseModelWithDefaults

`func NewHypervisorSecurityGroupRuleResponseModelWithDefaults() *HypervisorSecurityGroupRuleResponseModel`

NewHypervisorSecurityGroupRuleResponseModelWithDefaults instantiates a new HypervisorSecurityGroupRuleResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPort

`func (o *HypervisorSecurityGroupRuleResponseModel) GetFromPort() float32`

GetFromPort returns the FromPort field if non-nil, zero value otherwise.

### GetFromPortOk

`func (o *HypervisorSecurityGroupRuleResponseModel) GetFromPortOk() (*float32, bool)`

GetFromPortOk returns a tuple with the FromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPort

`func (o *HypervisorSecurityGroupRuleResponseModel) SetFromPort(v float32)`

SetFromPort sets FromPort field to given value.

### HasFromPort

`func (o *HypervisorSecurityGroupRuleResponseModel) HasFromPort() bool`

HasFromPort returns a boolean if a field has been set.

### GetGroupIds

`func (o *HypervisorSecurityGroupRuleResponseModel) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *HypervisorSecurityGroupRuleResponseModel) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *HypervisorSecurityGroupRuleResponseModel) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *HypervisorSecurityGroupRuleResponseModel) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetIPRanges

`func (o *HypervisorSecurityGroupRuleResponseModel) GetIPRanges() []string`

GetIPRanges returns the IPRanges field if non-nil, zero value otherwise.

### GetIPRangesOk

`func (o *HypervisorSecurityGroupRuleResponseModel) GetIPRangesOk() (*[]string, bool)`

GetIPRangesOk returns a tuple with the IPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPRanges

`func (o *HypervisorSecurityGroupRuleResponseModel) SetIPRanges(v []string)`

SetIPRanges sets IPRanges field to given value.

### HasIPRanges

`func (o *HypervisorSecurityGroupRuleResponseModel) HasIPRanges() bool`

HasIPRanges returns a boolean if a field has been set.

### GetProtocol

`func (o *HypervisorSecurityGroupRuleResponseModel) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HypervisorSecurityGroupRuleResponseModel) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HypervisorSecurityGroupRuleResponseModel) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetToPort

`func (o *HypervisorSecurityGroupRuleResponseModel) GetToPort() float32`

GetToPort returns the ToPort field if non-nil, zero value otherwise.

### GetToPortOk

`func (o *HypervisorSecurityGroupRuleResponseModel) GetToPortOk() (*float32, bool)`

GetToPortOk returns a tuple with the ToPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPort

`func (o *HypervisorSecurityGroupRuleResponseModel) SetToPort(v float32)`

SetToPort sets ToPort field to given value.

### HasToPort

`func (o *HypervisorSecurityGroupRuleResponseModel) HasToPort() bool`

HasToPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


