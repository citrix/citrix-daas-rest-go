# ConnectionSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetAddress** | **string** | IP Address of the subnet | 
**SubnetMask** | **int32** | Subnet mask associated with the subnet | 

## Methods

### NewConnectionSubnet

`func NewConnectionSubnet(subnetAddress string, subnetMask int32, ) *ConnectionSubnet`

NewConnectionSubnet instantiates a new ConnectionSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionSubnetWithDefaults

`func NewConnectionSubnetWithDefaults() *ConnectionSubnet`

NewConnectionSubnetWithDefaults instantiates a new ConnectionSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetAddress

`func (o *ConnectionSubnet) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *ConnectionSubnet) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *ConnectionSubnet) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.


### GetSubnetMask

`func (o *ConnectionSubnet) GetSubnetMask() int32`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *ConnectionSubnet) GetSubnetMaskOk() (*int32, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *ConnectionSubnet) SetSubnetMask(v int32)`

SetSubnetMask sets SubnetMask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


