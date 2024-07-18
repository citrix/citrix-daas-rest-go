# AwsEdcAccountResourceSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupId** | Pointer to **NullableString** | Security Group Id | [optional] 
**Name** | Pointer to **NullableString** | Security Group Name | [optional] 
**Description** | Pointer to **NullableString** | Security Group Description | [optional] 
**VpcId** | Pointer to **NullableString** | Security Group VpcId | [optional] 
**InboundPermissions** | Pointer to [**[]IpPermission**](IpPermission.md) | Security Group Inbound Permissions | [optional] 
**OutboundPermissions** | Pointer to [**[]IpPermission**](IpPermission.md) | Security Group Outbound Permissions | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Security Group Tags | [optional] 

## Methods

### NewAwsEdcAccountResourceSecurityGroup

`func NewAwsEdcAccountResourceSecurityGroup() *AwsEdcAccountResourceSecurityGroup`

NewAwsEdcAccountResourceSecurityGroup instantiates a new AwsEdcAccountResourceSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcAccountResourceSecurityGroupWithDefaults

`func NewAwsEdcAccountResourceSecurityGroupWithDefaults() *AwsEdcAccountResourceSecurityGroup`

NewAwsEdcAccountResourceSecurityGroupWithDefaults instantiates a new AwsEdcAccountResourceSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupId

`func (o *AwsEdcAccountResourceSecurityGroup) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *AwsEdcAccountResourceSecurityGroup) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *AwsEdcAccountResourceSecurityGroup) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *AwsEdcAccountResourceSecurityGroup) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupIdNil

`func (o *AwsEdcAccountResourceSecurityGroup) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *AwsEdcAccountResourceSecurityGroup) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil
### GetName

`func (o *AwsEdcAccountResourceSecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsEdcAccountResourceSecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsEdcAccountResourceSecurityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsEdcAccountResourceSecurityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsEdcAccountResourceSecurityGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsEdcAccountResourceSecurityGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AwsEdcAccountResourceSecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsEdcAccountResourceSecurityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsEdcAccountResourceSecurityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsEdcAccountResourceSecurityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AwsEdcAccountResourceSecurityGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AwsEdcAccountResourceSecurityGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVpcId

`func (o *AwsEdcAccountResourceSecurityGroup) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsEdcAccountResourceSecurityGroup) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsEdcAccountResourceSecurityGroup) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsEdcAccountResourceSecurityGroup) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *AwsEdcAccountResourceSecurityGroup) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AwsEdcAccountResourceSecurityGroup) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetInboundPermissions

`func (o *AwsEdcAccountResourceSecurityGroup) GetInboundPermissions() []IpPermission`

GetInboundPermissions returns the InboundPermissions field if non-nil, zero value otherwise.

### GetInboundPermissionsOk

`func (o *AwsEdcAccountResourceSecurityGroup) GetInboundPermissionsOk() (*[]IpPermission, bool)`

GetInboundPermissionsOk returns a tuple with the InboundPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundPermissions

`func (o *AwsEdcAccountResourceSecurityGroup) SetInboundPermissions(v []IpPermission)`

SetInboundPermissions sets InboundPermissions field to given value.

### HasInboundPermissions

`func (o *AwsEdcAccountResourceSecurityGroup) HasInboundPermissions() bool`

HasInboundPermissions returns a boolean if a field has been set.

### SetInboundPermissionsNil

`func (o *AwsEdcAccountResourceSecurityGroup) SetInboundPermissionsNil(b bool)`

 SetInboundPermissionsNil sets the value for InboundPermissions to be an explicit nil

### UnsetInboundPermissions
`func (o *AwsEdcAccountResourceSecurityGroup) UnsetInboundPermissions()`

UnsetInboundPermissions ensures that no value is present for InboundPermissions, not even an explicit nil
### GetOutboundPermissions

`func (o *AwsEdcAccountResourceSecurityGroup) GetOutboundPermissions() []IpPermission`

GetOutboundPermissions returns the OutboundPermissions field if non-nil, zero value otherwise.

### GetOutboundPermissionsOk

`func (o *AwsEdcAccountResourceSecurityGroup) GetOutboundPermissionsOk() (*[]IpPermission, bool)`

GetOutboundPermissionsOk returns a tuple with the OutboundPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundPermissions

`func (o *AwsEdcAccountResourceSecurityGroup) SetOutboundPermissions(v []IpPermission)`

SetOutboundPermissions sets OutboundPermissions field to given value.

### HasOutboundPermissions

`func (o *AwsEdcAccountResourceSecurityGroup) HasOutboundPermissions() bool`

HasOutboundPermissions returns a boolean if a field has been set.

### SetOutboundPermissionsNil

`func (o *AwsEdcAccountResourceSecurityGroup) SetOutboundPermissionsNil(b bool)`

 SetOutboundPermissionsNil sets the value for OutboundPermissions to be an explicit nil

### UnsetOutboundPermissions
`func (o *AwsEdcAccountResourceSecurityGroup) UnsetOutboundPermissions()`

UnsetOutboundPermissions ensures that no value is present for OutboundPermissions, not even an explicit nil
### GetTags

`func (o *AwsEdcAccountResourceSecurityGroup) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwsEdcAccountResourceSecurityGroup) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwsEdcAccountResourceSecurityGroup) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwsEdcAccountResourceSecurityGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AwsEdcAccountResourceSecurityGroup) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AwsEdcAccountResourceSecurityGroup) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


