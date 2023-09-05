# VdaVersionContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | XA or XD, no other values. | [optional] 
**Version** | Pointer to **string** | One of the values in XaVersions or XdVersions. | [optional] 
**Edition** | Pointer to **string** | MS or SS, no other values. | [optional] 
**DisplayName** | Pointer to **string** | Display name. | [optional] 

## Methods

### NewVdaVersionContract

`func NewVdaVersionContract() *VdaVersionContract`

NewVdaVersionContract instantiates a new VdaVersionContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdaVersionContractWithDefaults

`func NewVdaVersionContractWithDefaults() *VdaVersionContract`

NewVdaVersionContractWithDefaults instantiates a new VdaVersionContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *VdaVersionContract) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *VdaVersionContract) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *VdaVersionContract) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *VdaVersionContract) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetVersion

`func (o *VdaVersionContract) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VdaVersionContract) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VdaVersionContract) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VdaVersionContract) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEdition

`func (o *VdaVersionContract) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *VdaVersionContract) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *VdaVersionContract) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *VdaVersionContract) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetDisplayName

`func (o *VdaVersionContract) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VdaVersionContract) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VdaVersionContract) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VdaVersionContract) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


