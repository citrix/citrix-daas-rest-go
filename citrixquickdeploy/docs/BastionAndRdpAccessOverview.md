# BastionAndRdpAccessOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the bastion | [optional] 
**Name** | Pointer to **NullableString** | Name of the bastion | [optional] 
**State** | Pointer to **NullableString** | The current state of the bastion | [optional] 
**Error** | Pointer to **NullableString** | Error message if bastion is in failed state | [optional] 
**VnetName** | Pointer to **NullableString** | The name of the Vnet in which the bastion is created | [optional] 
**VnetAddressSpace** | Pointer to **NullableString** | The address space of the Vnet in which the bastion is created | [optional] 
**RestrictedIpAddressPrefix** | Pointer to **NullableString** | The Ip Address prefix, if any, to which the bastion rdp access is restricted | [optional] 
**TransactionId** | Pointer to **NullableString** | The transaction ID associated with the bastion | [optional] 
**HasRdpFileBeenRequested** | Pointer to **bool** | A boolean to check if the rdp file has been requested | [optional] 
**Type** | Pointer to **NullableString** | Specifies if it is a bastion or a rdp access | [optional] 
**CspCustomer** | Pointer to **NullableString** | Indicates that partner-tenant relationship exists if not null | [optional] 
**StartedAt** | Pointer to **NullableTime** | DateTime when the import image job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **NullableInt32** | Estimated total time in minute for the job to finish | [optional] 

## Methods

### NewBastionAndRdpAccessOverview

`func NewBastionAndRdpAccessOverview() *BastionAndRdpAccessOverview`

NewBastionAndRdpAccessOverview instantiates a new BastionAndRdpAccessOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBastionAndRdpAccessOverviewWithDefaults

`func NewBastionAndRdpAccessOverviewWithDefaults() *BastionAndRdpAccessOverview`

NewBastionAndRdpAccessOverviewWithDefaults instantiates a new BastionAndRdpAccessOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BastionAndRdpAccessOverview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BastionAndRdpAccessOverview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BastionAndRdpAccessOverview) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BastionAndRdpAccessOverview) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BastionAndRdpAccessOverview) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BastionAndRdpAccessOverview) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *BastionAndRdpAccessOverview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BastionAndRdpAccessOverview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BastionAndRdpAccessOverview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BastionAndRdpAccessOverview) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BastionAndRdpAccessOverview) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BastionAndRdpAccessOverview) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetState

`func (o *BastionAndRdpAccessOverview) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BastionAndRdpAccessOverview) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BastionAndRdpAccessOverview) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BastionAndRdpAccessOverview) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *BastionAndRdpAccessOverview) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *BastionAndRdpAccessOverview) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetError

`func (o *BastionAndRdpAccessOverview) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BastionAndRdpAccessOverview) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BastionAndRdpAccessOverview) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BastionAndRdpAccessOverview) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *BastionAndRdpAccessOverview) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BastionAndRdpAccessOverview) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetVnetName

`func (o *BastionAndRdpAccessOverview) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *BastionAndRdpAccessOverview) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *BastionAndRdpAccessOverview) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.

### HasVnetName

`func (o *BastionAndRdpAccessOverview) HasVnetName() bool`

HasVnetName returns a boolean if a field has been set.

### SetVnetNameNil

`func (o *BastionAndRdpAccessOverview) SetVnetNameNil(b bool)`

 SetVnetNameNil sets the value for VnetName to be an explicit nil

### UnsetVnetName
`func (o *BastionAndRdpAccessOverview) UnsetVnetName()`

UnsetVnetName ensures that no value is present for VnetName, not even an explicit nil
### GetVnetAddressSpace

`func (o *BastionAndRdpAccessOverview) GetVnetAddressSpace() string`

GetVnetAddressSpace returns the VnetAddressSpace field if non-nil, zero value otherwise.

### GetVnetAddressSpaceOk

`func (o *BastionAndRdpAccessOverview) GetVnetAddressSpaceOk() (*string, bool)`

GetVnetAddressSpaceOk returns a tuple with the VnetAddressSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetAddressSpace

`func (o *BastionAndRdpAccessOverview) SetVnetAddressSpace(v string)`

SetVnetAddressSpace sets VnetAddressSpace field to given value.

### HasVnetAddressSpace

`func (o *BastionAndRdpAccessOverview) HasVnetAddressSpace() bool`

HasVnetAddressSpace returns a boolean if a field has been set.

### SetVnetAddressSpaceNil

`func (o *BastionAndRdpAccessOverview) SetVnetAddressSpaceNil(b bool)`

 SetVnetAddressSpaceNil sets the value for VnetAddressSpace to be an explicit nil

### UnsetVnetAddressSpace
`func (o *BastionAndRdpAccessOverview) UnsetVnetAddressSpace()`

UnsetVnetAddressSpace ensures that no value is present for VnetAddressSpace, not even an explicit nil
### GetRestrictedIpAddressPrefix

`func (o *BastionAndRdpAccessOverview) GetRestrictedIpAddressPrefix() string`

GetRestrictedIpAddressPrefix returns the RestrictedIpAddressPrefix field if non-nil, zero value otherwise.

### GetRestrictedIpAddressPrefixOk

`func (o *BastionAndRdpAccessOverview) GetRestrictedIpAddressPrefixOk() (*string, bool)`

GetRestrictedIpAddressPrefixOk returns a tuple with the RestrictedIpAddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedIpAddressPrefix

`func (o *BastionAndRdpAccessOverview) SetRestrictedIpAddressPrefix(v string)`

SetRestrictedIpAddressPrefix sets RestrictedIpAddressPrefix field to given value.

### HasRestrictedIpAddressPrefix

`func (o *BastionAndRdpAccessOverview) HasRestrictedIpAddressPrefix() bool`

HasRestrictedIpAddressPrefix returns a boolean if a field has been set.

### SetRestrictedIpAddressPrefixNil

`func (o *BastionAndRdpAccessOverview) SetRestrictedIpAddressPrefixNil(b bool)`

 SetRestrictedIpAddressPrefixNil sets the value for RestrictedIpAddressPrefix to be an explicit nil

### UnsetRestrictedIpAddressPrefix
`func (o *BastionAndRdpAccessOverview) UnsetRestrictedIpAddressPrefix()`

UnsetRestrictedIpAddressPrefix ensures that no value is present for RestrictedIpAddressPrefix, not even an explicit nil
### GetTransactionId

`func (o *BastionAndRdpAccessOverview) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BastionAndRdpAccessOverview) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BastionAndRdpAccessOverview) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BastionAndRdpAccessOverview) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *BastionAndRdpAccessOverview) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *BastionAndRdpAccessOverview) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetHasRdpFileBeenRequested

`func (o *BastionAndRdpAccessOverview) GetHasRdpFileBeenRequested() bool`

GetHasRdpFileBeenRequested returns the HasRdpFileBeenRequested field if non-nil, zero value otherwise.

### GetHasRdpFileBeenRequestedOk

`func (o *BastionAndRdpAccessOverview) GetHasRdpFileBeenRequestedOk() (*bool, bool)`

GetHasRdpFileBeenRequestedOk returns a tuple with the HasRdpFileBeenRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRdpFileBeenRequested

`func (o *BastionAndRdpAccessOverview) SetHasRdpFileBeenRequested(v bool)`

SetHasRdpFileBeenRequested sets HasRdpFileBeenRequested field to given value.

### HasHasRdpFileBeenRequested

`func (o *BastionAndRdpAccessOverview) HasHasRdpFileBeenRequested() bool`

HasHasRdpFileBeenRequested returns a boolean if a field has been set.

### GetType

`func (o *BastionAndRdpAccessOverview) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BastionAndRdpAccessOverview) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BastionAndRdpAccessOverview) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BastionAndRdpAccessOverview) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BastionAndRdpAccessOverview) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BastionAndRdpAccessOverview) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCspCustomer

`func (o *BastionAndRdpAccessOverview) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *BastionAndRdpAccessOverview) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *BastionAndRdpAccessOverview) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *BastionAndRdpAccessOverview) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### SetCspCustomerNil

`func (o *BastionAndRdpAccessOverview) SetCspCustomerNil(b bool)`

 SetCspCustomerNil sets the value for CspCustomer to be an explicit nil

### UnsetCspCustomer
`func (o *BastionAndRdpAccessOverview) UnsetCspCustomer()`

UnsetCspCustomer ensures that no value is present for CspCustomer, not even an explicit nil
### GetStartedAt

`func (o *BastionAndRdpAccessOverview) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BastionAndRdpAccessOverview) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BastionAndRdpAccessOverview) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BastionAndRdpAccessOverview) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *BastionAndRdpAccessOverview) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *BastionAndRdpAccessOverview) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEstimatedTimeInMinute

`func (o *BastionAndRdpAccessOverview) GetEstimatedTimeInMinute() int32`

GetEstimatedTimeInMinute returns the EstimatedTimeInMinute field if non-nil, zero value otherwise.

### GetEstimatedTimeInMinuteOk

`func (o *BastionAndRdpAccessOverview) GetEstimatedTimeInMinuteOk() (*int32, bool)`

GetEstimatedTimeInMinuteOk returns a tuple with the EstimatedTimeInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeInMinute

`func (o *BastionAndRdpAccessOverview) SetEstimatedTimeInMinute(v int32)`

SetEstimatedTimeInMinute sets EstimatedTimeInMinute field to given value.

### HasEstimatedTimeInMinute

`func (o *BastionAndRdpAccessOverview) HasEstimatedTimeInMinute() bool`

HasEstimatedTimeInMinute returns a boolean if a field has been set.

### SetEstimatedTimeInMinuteNil

`func (o *BastionAndRdpAccessOverview) SetEstimatedTimeInMinuteNil(b bool)`

 SetEstimatedTimeInMinuteNil sets the value for EstimatedTimeInMinute to be an explicit nil

### UnsetEstimatedTimeInMinute
`func (o *BastionAndRdpAccessOverview) UnsetEstimatedTimeInMinute()`

UnsetEstimatedTimeInMinute ensures that no value is present for EstimatedTimeInMinute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


