# AzureSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the subnet | 
**AddressPrefix** | **string** | The address range specified by this subnet | 

## Methods

### NewAzureSubnet

`func NewAzureSubnet(name string, addressPrefix string, ) *AzureSubnet`

NewAzureSubnet instantiates a new AzureSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSubnetWithDefaults

`func NewAzureSubnetWithDefaults() *AzureSubnet`

NewAzureSubnetWithDefaults instantiates a new AzureSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureSubnet) SetName(v string)`

SetName sets Name field to given value.


### GetAddressPrefix

`func (o *AzureSubnet) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *AzureSubnet) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *AzureSubnet) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


