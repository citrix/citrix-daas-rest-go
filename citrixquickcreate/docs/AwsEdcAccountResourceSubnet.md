# AwsEdcAccountResourceSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetId** | Pointer to **NullableString** | Subnet Id | [optional] 
**Status** | Pointer to [**NullableAwsEdcSubnetStatus**](AwsEdcSubnetStatus.md) |  | [optional] 
**SubnetArn** | Pointer to **NullableString** | Subnet Status | [optional] 
**CidrBlock** | Pointer to **NullableString** | Subnet CIDR Block | [optional] 
**VpcId** | Pointer to **NullableString** | Subnet Vpc Id | [optional] 
**AwsRegion** | Pointer to **NullableString** | The AWS region where the subnet is located | [optional] 
**AvailabilityZoneId** | Pointer to **NullableString** | Subnet Availability Zone Id | [optional] 
**AvailabilityZone** | Pointer to **NullableString** | Subnet Availability Zone | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Subnet Tags | [optional] 
**Name** | Pointer to **NullableString** | The name of the subnet | [optional] [readonly] 
**IsSupportedByAwsManagedCapacity** | Pointer to **bool** | Indicates if the subnet can be used for Managed Capacity deployments | [optional] [readonly] 

## Methods

### NewAwsEdcAccountResourceSubnet

`func NewAwsEdcAccountResourceSubnet() *AwsEdcAccountResourceSubnet`

NewAwsEdcAccountResourceSubnet instantiates a new AwsEdcAccountResourceSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcAccountResourceSubnetWithDefaults

`func NewAwsEdcAccountResourceSubnetWithDefaults() *AwsEdcAccountResourceSubnet`

NewAwsEdcAccountResourceSubnetWithDefaults instantiates a new AwsEdcAccountResourceSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetId

`func (o *AwsEdcAccountResourceSubnet) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AwsEdcAccountResourceSubnet) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AwsEdcAccountResourceSubnet) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *AwsEdcAccountResourceSubnet) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *AwsEdcAccountResourceSubnet) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *AwsEdcAccountResourceSubnet) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetStatus

`func (o *AwsEdcAccountResourceSubnet) GetStatus() AwsEdcSubnetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsEdcAccountResourceSubnet) GetStatusOk() (*AwsEdcSubnetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsEdcAccountResourceSubnet) SetStatus(v AwsEdcSubnetStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsEdcAccountResourceSubnet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AwsEdcAccountResourceSubnet) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AwsEdcAccountResourceSubnet) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubnetArn

`func (o *AwsEdcAccountResourceSubnet) GetSubnetArn() string`

GetSubnetArn returns the SubnetArn field if non-nil, zero value otherwise.

### GetSubnetArnOk

`func (o *AwsEdcAccountResourceSubnet) GetSubnetArnOk() (*string, bool)`

GetSubnetArnOk returns a tuple with the SubnetArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetArn

`func (o *AwsEdcAccountResourceSubnet) SetSubnetArn(v string)`

SetSubnetArn sets SubnetArn field to given value.

### HasSubnetArn

`func (o *AwsEdcAccountResourceSubnet) HasSubnetArn() bool`

HasSubnetArn returns a boolean if a field has been set.

### SetSubnetArnNil

`func (o *AwsEdcAccountResourceSubnet) SetSubnetArnNil(b bool)`

 SetSubnetArnNil sets the value for SubnetArn to be an explicit nil

### UnsetSubnetArn
`func (o *AwsEdcAccountResourceSubnet) UnsetSubnetArn()`

UnsetSubnetArn ensures that no value is present for SubnetArn, not even an explicit nil
### GetCidrBlock

`func (o *AwsEdcAccountResourceSubnet) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AwsEdcAccountResourceSubnet) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AwsEdcAccountResourceSubnet) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *AwsEdcAccountResourceSubnet) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### SetCidrBlockNil

`func (o *AwsEdcAccountResourceSubnet) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *AwsEdcAccountResourceSubnet) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
### GetVpcId

`func (o *AwsEdcAccountResourceSubnet) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsEdcAccountResourceSubnet) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsEdcAccountResourceSubnet) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsEdcAccountResourceSubnet) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *AwsEdcAccountResourceSubnet) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AwsEdcAccountResourceSubnet) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetAwsRegion

`func (o *AwsEdcAccountResourceSubnet) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AwsEdcAccountResourceSubnet) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AwsEdcAccountResourceSubnet) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *AwsEdcAccountResourceSubnet) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *AwsEdcAccountResourceSubnet) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AwsEdcAccountResourceSubnet) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetAvailabilityZoneId

`func (o *AwsEdcAccountResourceSubnet) GetAvailabilityZoneId() string`

GetAvailabilityZoneId returns the AvailabilityZoneId field if non-nil, zero value otherwise.

### GetAvailabilityZoneIdOk

`func (o *AwsEdcAccountResourceSubnet) GetAvailabilityZoneIdOk() (*string, bool)`

GetAvailabilityZoneIdOk returns a tuple with the AvailabilityZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneId

`func (o *AwsEdcAccountResourceSubnet) SetAvailabilityZoneId(v string)`

SetAvailabilityZoneId sets AvailabilityZoneId field to given value.

### HasAvailabilityZoneId

`func (o *AwsEdcAccountResourceSubnet) HasAvailabilityZoneId() bool`

HasAvailabilityZoneId returns a boolean if a field has been set.

### SetAvailabilityZoneIdNil

`func (o *AwsEdcAccountResourceSubnet) SetAvailabilityZoneIdNil(b bool)`

 SetAvailabilityZoneIdNil sets the value for AvailabilityZoneId to be an explicit nil

### UnsetAvailabilityZoneId
`func (o *AwsEdcAccountResourceSubnet) UnsetAvailabilityZoneId()`

UnsetAvailabilityZoneId ensures that no value is present for AvailabilityZoneId, not even an explicit nil
### GetAvailabilityZone

`func (o *AwsEdcAccountResourceSubnet) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AwsEdcAccountResourceSubnet) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AwsEdcAccountResourceSubnet) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AwsEdcAccountResourceSubnet) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *AwsEdcAccountResourceSubnet) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *AwsEdcAccountResourceSubnet) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetTags

`func (o *AwsEdcAccountResourceSubnet) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwsEdcAccountResourceSubnet) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwsEdcAccountResourceSubnet) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwsEdcAccountResourceSubnet) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AwsEdcAccountResourceSubnet) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AwsEdcAccountResourceSubnet) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetName

`func (o *AwsEdcAccountResourceSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsEdcAccountResourceSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsEdcAccountResourceSubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsEdcAccountResourceSubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsEdcAccountResourceSubnet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsEdcAccountResourceSubnet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsSupportedByAwsManagedCapacity

`func (o *AwsEdcAccountResourceSubnet) GetIsSupportedByAwsManagedCapacity() bool`

GetIsSupportedByAwsManagedCapacity returns the IsSupportedByAwsManagedCapacity field if non-nil, zero value otherwise.

### GetIsSupportedByAwsManagedCapacityOk

`func (o *AwsEdcAccountResourceSubnet) GetIsSupportedByAwsManagedCapacityOk() (*bool, bool)`

GetIsSupportedByAwsManagedCapacityOk returns a tuple with the IsSupportedByAwsManagedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupportedByAwsManagedCapacity

`func (o *AwsEdcAccountResourceSubnet) SetIsSupportedByAwsManagedCapacity(v bool)`

SetIsSupportedByAwsManagedCapacity sets IsSupportedByAwsManagedCapacity field to given value.

### HasIsSupportedByAwsManagedCapacity

`func (o *AwsEdcAccountResourceSubnet) HasIsSupportedByAwsManagedCapacity() bool`

HasIsSupportedByAwsManagedCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


