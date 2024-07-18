# UserIdGroupPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**GroupId** | Pointer to **NullableString** |  | [optional] 
**GroupName** | Pointer to **NullableString** |  | [optional] 
**PeeringStatus** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**VpcPeeringConnectionId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserIdGroupPair

`func NewUserIdGroupPair() *UserIdGroupPair`

NewUserIdGroupPair instantiates a new UserIdGroupPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdGroupPairWithDefaults

`func NewUserIdGroupPairWithDefaults() *UserIdGroupPair`

NewUserIdGroupPairWithDefaults instantiates a new UserIdGroupPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UserIdGroupPair) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserIdGroupPair) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserIdGroupPair) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserIdGroupPair) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UserIdGroupPair) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UserIdGroupPair) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGroupId

`func (o *UserIdGroupPair) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UserIdGroupPair) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UserIdGroupPair) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UserIdGroupPair) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *UserIdGroupPair) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *UserIdGroupPair) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *UserIdGroupPair) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *UserIdGroupPair) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *UserIdGroupPair) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *UserIdGroupPair) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *UserIdGroupPair) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *UserIdGroupPair) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetPeeringStatus

`func (o *UserIdGroupPair) GetPeeringStatus() string`

GetPeeringStatus returns the PeeringStatus field if non-nil, zero value otherwise.

### GetPeeringStatusOk

`func (o *UserIdGroupPair) GetPeeringStatusOk() (*string, bool)`

GetPeeringStatusOk returns a tuple with the PeeringStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringStatus

`func (o *UserIdGroupPair) SetPeeringStatus(v string)`

SetPeeringStatus sets PeeringStatus field to given value.

### HasPeeringStatus

`func (o *UserIdGroupPair) HasPeeringStatus() bool`

HasPeeringStatus returns a boolean if a field has been set.

### SetPeeringStatusNil

`func (o *UserIdGroupPair) SetPeeringStatusNil(b bool)`

 SetPeeringStatusNil sets the value for PeeringStatus to be an explicit nil

### UnsetPeeringStatus
`func (o *UserIdGroupPair) UnsetPeeringStatus()`

UnsetPeeringStatus ensures that no value is present for PeeringStatus, not even an explicit nil
### GetUserId

`func (o *UserIdGroupPair) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserIdGroupPair) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserIdGroupPair) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserIdGroupPair) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UserIdGroupPair) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UserIdGroupPair) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetVpcId

`func (o *UserIdGroupPair) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *UserIdGroupPair) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *UserIdGroupPair) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *UserIdGroupPair) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *UserIdGroupPair) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *UserIdGroupPair) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcPeeringConnectionId

`func (o *UserIdGroupPair) GetVpcPeeringConnectionId() string`

GetVpcPeeringConnectionId returns the VpcPeeringConnectionId field if non-nil, zero value otherwise.

### GetVpcPeeringConnectionIdOk

`func (o *UserIdGroupPair) GetVpcPeeringConnectionIdOk() (*string, bool)`

GetVpcPeeringConnectionIdOk returns a tuple with the VpcPeeringConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPeeringConnectionId

`func (o *UserIdGroupPair) SetVpcPeeringConnectionId(v string)`

SetVpcPeeringConnectionId sets VpcPeeringConnectionId field to given value.

### HasVpcPeeringConnectionId

`func (o *UserIdGroupPair) HasVpcPeeringConnectionId() bool`

HasVpcPeeringConnectionId returns a boolean if a field has been set.

### SetVpcPeeringConnectionIdNil

`func (o *UserIdGroupPair) SetVpcPeeringConnectionIdNil(b bool)`

 SetVpcPeeringConnectionIdNil sets the value for VpcPeeringConnectionId to be an explicit nil

### UnsetVpcPeeringConnectionId
`func (o *UserIdGroupPair) UnsetVpcPeeringConnectionId()`

UnsetVpcPeeringConnectionId ensures that no value is present for VpcPeeringConnectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


