# IpPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRanges** | Pointer to **[]string** |  | [optional] 
**FromPort** | Pointer to **int32** |  | [optional] 
**IpProtocol** | Pointer to **NullableString** |  | [optional] 
**Ipv4Ranges** | Pointer to [**[]IpRange**](IpRange.md) |  | [optional] 
**Ipv6Ranges** | Pointer to [**[]Ipv6Range**](Ipv6Range.md) |  | [optional] 
**PrefixListIds** | Pointer to [**[]PrefixListId**](PrefixListId.md) |  | [optional] 
**ToPort** | Pointer to **int32** |  | [optional] 
**UserIdGroupPairs** | Pointer to [**[]UserIdGroupPair**](UserIdGroupPair.md) |  | [optional] 

## Methods

### NewIpPermission

`func NewIpPermission() *IpPermission`

NewIpPermission instantiates a new IpPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpPermissionWithDefaults

`func NewIpPermissionWithDefaults() *IpPermission`

NewIpPermissionWithDefaults instantiates a new IpPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpRanges

`func (o *IpPermission) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *IpPermission) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *IpPermission) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *IpPermission) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### SetIpRangesNil

`func (o *IpPermission) SetIpRangesNil(b bool)`

 SetIpRangesNil sets the value for IpRanges to be an explicit nil

### UnsetIpRanges
`func (o *IpPermission) UnsetIpRanges()`

UnsetIpRanges ensures that no value is present for IpRanges, not even an explicit nil
### GetFromPort

`func (o *IpPermission) GetFromPort() int32`

GetFromPort returns the FromPort field if non-nil, zero value otherwise.

### GetFromPortOk

`func (o *IpPermission) GetFromPortOk() (*int32, bool)`

GetFromPortOk returns a tuple with the FromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPort

`func (o *IpPermission) SetFromPort(v int32)`

SetFromPort sets FromPort field to given value.

### HasFromPort

`func (o *IpPermission) HasFromPort() bool`

HasFromPort returns a boolean if a field has been set.

### GetIpProtocol

`func (o *IpPermission) GetIpProtocol() string`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *IpPermission) GetIpProtocolOk() (*string, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpProtocol

`func (o *IpPermission) SetIpProtocol(v string)`

SetIpProtocol sets IpProtocol field to given value.

### HasIpProtocol

`func (o *IpPermission) HasIpProtocol() bool`

HasIpProtocol returns a boolean if a field has been set.

### SetIpProtocolNil

`func (o *IpPermission) SetIpProtocolNil(b bool)`

 SetIpProtocolNil sets the value for IpProtocol to be an explicit nil

### UnsetIpProtocol
`func (o *IpPermission) UnsetIpProtocol()`

UnsetIpProtocol ensures that no value is present for IpProtocol, not even an explicit nil
### GetIpv4Ranges

`func (o *IpPermission) GetIpv4Ranges() []IpRange`

GetIpv4Ranges returns the Ipv4Ranges field if non-nil, zero value otherwise.

### GetIpv4RangesOk

`func (o *IpPermission) GetIpv4RangesOk() (*[]IpRange, bool)`

GetIpv4RangesOk returns a tuple with the Ipv4Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Ranges

`func (o *IpPermission) SetIpv4Ranges(v []IpRange)`

SetIpv4Ranges sets Ipv4Ranges field to given value.

### HasIpv4Ranges

`func (o *IpPermission) HasIpv4Ranges() bool`

HasIpv4Ranges returns a boolean if a field has been set.

### SetIpv4RangesNil

`func (o *IpPermission) SetIpv4RangesNil(b bool)`

 SetIpv4RangesNil sets the value for Ipv4Ranges to be an explicit nil

### UnsetIpv4Ranges
`func (o *IpPermission) UnsetIpv4Ranges()`

UnsetIpv4Ranges ensures that no value is present for Ipv4Ranges, not even an explicit nil
### GetIpv6Ranges

`func (o *IpPermission) GetIpv6Ranges() []Ipv6Range`

GetIpv6Ranges returns the Ipv6Ranges field if non-nil, zero value otherwise.

### GetIpv6RangesOk

`func (o *IpPermission) GetIpv6RangesOk() (*[]Ipv6Range, bool)`

GetIpv6RangesOk returns a tuple with the Ipv6Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Ranges

`func (o *IpPermission) SetIpv6Ranges(v []Ipv6Range)`

SetIpv6Ranges sets Ipv6Ranges field to given value.

### HasIpv6Ranges

`func (o *IpPermission) HasIpv6Ranges() bool`

HasIpv6Ranges returns a boolean if a field has been set.

### SetIpv6RangesNil

`func (o *IpPermission) SetIpv6RangesNil(b bool)`

 SetIpv6RangesNil sets the value for Ipv6Ranges to be an explicit nil

### UnsetIpv6Ranges
`func (o *IpPermission) UnsetIpv6Ranges()`

UnsetIpv6Ranges ensures that no value is present for Ipv6Ranges, not even an explicit nil
### GetPrefixListIds

`func (o *IpPermission) GetPrefixListIds() []PrefixListId`

GetPrefixListIds returns the PrefixListIds field if non-nil, zero value otherwise.

### GetPrefixListIdsOk

`func (o *IpPermission) GetPrefixListIdsOk() (*[]PrefixListId, bool)`

GetPrefixListIdsOk returns a tuple with the PrefixListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixListIds

`func (o *IpPermission) SetPrefixListIds(v []PrefixListId)`

SetPrefixListIds sets PrefixListIds field to given value.

### HasPrefixListIds

`func (o *IpPermission) HasPrefixListIds() bool`

HasPrefixListIds returns a boolean if a field has been set.

### SetPrefixListIdsNil

`func (o *IpPermission) SetPrefixListIdsNil(b bool)`

 SetPrefixListIdsNil sets the value for PrefixListIds to be an explicit nil

### UnsetPrefixListIds
`func (o *IpPermission) UnsetPrefixListIds()`

UnsetPrefixListIds ensures that no value is present for PrefixListIds, not even an explicit nil
### GetToPort

`func (o *IpPermission) GetToPort() int32`

GetToPort returns the ToPort field if non-nil, zero value otherwise.

### GetToPortOk

`func (o *IpPermission) GetToPortOk() (*int32, bool)`

GetToPortOk returns a tuple with the ToPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPort

`func (o *IpPermission) SetToPort(v int32)`

SetToPort sets ToPort field to given value.

### HasToPort

`func (o *IpPermission) HasToPort() bool`

HasToPort returns a boolean if a field has been set.

### GetUserIdGroupPairs

`func (o *IpPermission) GetUserIdGroupPairs() []UserIdGroupPair`

GetUserIdGroupPairs returns the UserIdGroupPairs field if non-nil, zero value otherwise.

### GetUserIdGroupPairsOk

`func (o *IpPermission) GetUserIdGroupPairsOk() (*[]UserIdGroupPair, bool)`

GetUserIdGroupPairsOk returns a tuple with the UserIdGroupPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdGroupPairs

`func (o *IpPermission) SetUserIdGroupPairs(v []UserIdGroupPair)`

SetUserIdGroupPairs sets UserIdGroupPairs field to given value.

### HasUserIdGroupPairs

`func (o *IpPermission) HasUserIdGroupPairs() bool`

HasUserIdGroupPairs returns a boolean if a field has been set.

### SetUserIdGroupPairsNil

`func (o *IpPermission) SetUserIdGroupPairsNil(b bool)`

 SetUserIdGroupPairsNil sets the value for UserIdGroupPairs to be an explicit nil

### UnsetUserIdGroupPairs
`func (o *IpPermission) UnsetUserIdGroupPairs()`

UnsetUserIdGroupPairs ensures that no value is present for UserIdGroupPairs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


