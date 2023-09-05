# VersionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | XA or XD, no other values. | [optional] 
**Version** | Pointer to **string** | One of the values in XaVersions or XdVersions. | [optional] 
**Edition** | Pointer to **string** | MS or SS, no other values. | [optional] 
**DisplayName** | Pointer to **string** | Display name. | [optional] 

## Methods

### NewVersionDetail

`func NewVersionDetail() *VersionDetail`

NewVersionDetail instantiates a new VersionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionDetailWithDefaults

`func NewVersionDetailWithDefaults() *VersionDetail`

NewVersionDetailWithDefaults instantiates a new VersionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *VersionDetail) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *VersionDetail) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *VersionDetail) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *VersionDetail) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetVersion

`func (o *VersionDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionDetail) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VersionDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEdition

`func (o *VersionDetail) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *VersionDetail) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *VersionDetail) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *VersionDetail) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetDisplayName

`func (o *VersionDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VersionDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VersionDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VersionDetail) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


