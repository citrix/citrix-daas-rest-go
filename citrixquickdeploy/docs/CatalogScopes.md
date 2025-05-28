# CatalogScopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogId** | Pointer to **string** |  | [optional] 
**CatalogName** | Pointer to **string** |  | [optional] 
**Advanced** | Pointer to **bool** |  | [optional] 
**DeliveryGroupScopes** | Pointer to **[]string** |  | [optional] 
**DeliveryGroupMetadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) |  | [optional] 
**MachineCatalogScopes** | Pointer to **[]string** |  | [optional] 
**MachineCatalogMetadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) |  | [optional] 
**HostingConnectionScopes** | Pointer to **[]string** |  | [optional] 
**HostingConnectionMetadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) |  | [optional] 
**Error** | Pointer to **bool** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewCatalogScopes

`func NewCatalogScopes() *CatalogScopes`

NewCatalogScopes instantiates a new CatalogScopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogScopesWithDefaults

`func NewCatalogScopesWithDefaults() *CatalogScopes`

NewCatalogScopesWithDefaults instantiates a new CatalogScopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogId

`func (o *CatalogScopes) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *CatalogScopes) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *CatalogScopes) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *CatalogScopes) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetCatalogName

`func (o *CatalogScopes) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *CatalogScopes) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *CatalogScopes) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.

### HasCatalogName

`func (o *CatalogScopes) HasCatalogName() bool`

HasCatalogName returns a boolean if a field has been set.

### GetAdvanced

`func (o *CatalogScopes) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *CatalogScopes) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *CatalogScopes) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *CatalogScopes) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetDeliveryGroupScopes

`func (o *CatalogScopes) GetDeliveryGroupScopes() []string`

GetDeliveryGroupScopes returns the DeliveryGroupScopes field if non-nil, zero value otherwise.

### GetDeliveryGroupScopesOk

`func (o *CatalogScopes) GetDeliveryGroupScopesOk() (*[]string, bool)`

GetDeliveryGroupScopesOk returns a tuple with the DeliveryGroupScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroupScopes

`func (o *CatalogScopes) SetDeliveryGroupScopes(v []string)`

SetDeliveryGroupScopes sets DeliveryGroupScopes field to given value.

### HasDeliveryGroupScopes

`func (o *CatalogScopes) HasDeliveryGroupScopes() bool`

HasDeliveryGroupScopes returns a boolean if a field has been set.

### GetDeliveryGroupMetadata

`func (o *CatalogScopes) GetDeliveryGroupMetadata() []NameValueStringPairModel`

GetDeliveryGroupMetadata returns the DeliveryGroupMetadata field if non-nil, zero value otherwise.

### GetDeliveryGroupMetadataOk

`func (o *CatalogScopes) GetDeliveryGroupMetadataOk() (*[]NameValueStringPairModel, bool)`

GetDeliveryGroupMetadataOk returns a tuple with the DeliveryGroupMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroupMetadata

`func (o *CatalogScopes) SetDeliveryGroupMetadata(v []NameValueStringPairModel)`

SetDeliveryGroupMetadata sets DeliveryGroupMetadata field to given value.

### HasDeliveryGroupMetadata

`func (o *CatalogScopes) HasDeliveryGroupMetadata() bool`

HasDeliveryGroupMetadata returns a boolean if a field has been set.

### GetMachineCatalogScopes

`func (o *CatalogScopes) GetMachineCatalogScopes() []string`

GetMachineCatalogScopes returns the MachineCatalogScopes field if non-nil, zero value otherwise.

### GetMachineCatalogScopesOk

`func (o *CatalogScopes) GetMachineCatalogScopesOk() (*[]string, bool)`

GetMachineCatalogScopesOk returns a tuple with the MachineCatalogScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCatalogScopes

`func (o *CatalogScopes) SetMachineCatalogScopes(v []string)`

SetMachineCatalogScopes sets MachineCatalogScopes field to given value.

### HasMachineCatalogScopes

`func (o *CatalogScopes) HasMachineCatalogScopes() bool`

HasMachineCatalogScopes returns a boolean if a field has been set.

### GetMachineCatalogMetadata

`func (o *CatalogScopes) GetMachineCatalogMetadata() []NameValueStringPairModel`

GetMachineCatalogMetadata returns the MachineCatalogMetadata field if non-nil, zero value otherwise.

### GetMachineCatalogMetadataOk

`func (o *CatalogScopes) GetMachineCatalogMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMachineCatalogMetadataOk returns a tuple with the MachineCatalogMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCatalogMetadata

`func (o *CatalogScopes) SetMachineCatalogMetadata(v []NameValueStringPairModel)`

SetMachineCatalogMetadata sets MachineCatalogMetadata field to given value.

### HasMachineCatalogMetadata

`func (o *CatalogScopes) HasMachineCatalogMetadata() bool`

HasMachineCatalogMetadata returns a boolean if a field has been set.

### GetHostingConnectionScopes

`func (o *CatalogScopes) GetHostingConnectionScopes() []string`

GetHostingConnectionScopes returns the HostingConnectionScopes field if non-nil, zero value otherwise.

### GetHostingConnectionScopesOk

`func (o *CatalogScopes) GetHostingConnectionScopesOk() (*[]string, bool)`

GetHostingConnectionScopesOk returns a tuple with the HostingConnectionScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingConnectionScopes

`func (o *CatalogScopes) SetHostingConnectionScopes(v []string)`

SetHostingConnectionScopes sets HostingConnectionScopes field to given value.

### HasHostingConnectionScopes

`func (o *CatalogScopes) HasHostingConnectionScopes() bool`

HasHostingConnectionScopes returns a boolean if a field has been set.

### GetHostingConnectionMetadata

`func (o *CatalogScopes) GetHostingConnectionMetadata() []NameValueStringPairModel`

GetHostingConnectionMetadata returns the HostingConnectionMetadata field if non-nil, zero value otherwise.

### GetHostingConnectionMetadataOk

`func (o *CatalogScopes) GetHostingConnectionMetadataOk() (*[]NameValueStringPairModel, bool)`

GetHostingConnectionMetadataOk returns a tuple with the HostingConnectionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingConnectionMetadata

`func (o *CatalogScopes) SetHostingConnectionMetadata(v []NameValueStringPairModel)`

SetHostingConnectionMetadata sets HostingConnectionMetadata field to given value.

### HasHostingConnectionMetadata

`func (o *CatalogScopes) HasHostingConnectionMetadata() bool`

HasHostingConnectionMetadata returns a boolean if a field has been set.

### GetError

`func (o *CatalogScopes) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CatalogScopes) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CatalogScopes) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *CatalogScopes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CatalogScopes) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CatalogScopes) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CatalogScopes) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CatalogScopes) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


