# DeleteCatalogModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceDelete** | Pointer to **bool** | Indicates that we want to bypass processing checks on deploy and force delete the catalog | [optional] 
**DeleteVm** | Pointer to **bool** | Specifies whether to delete virtual machines that were provisioned in the machine catalog | [optional] 
**DeleteResourceLocationIfUnused** | Pointer to **bool** | Indicates if the Resource Location associated with the catalog should be removed (along with any Catalog Service deployed connectors)  if this is the last catalog to use the Resource Location | [optional] 
**ServiceAccount** | Pointer to **string** | Service account to perform delete with | [optional] 
**ServiceAccountPassword** | Pointer to **string** | Customer&#39;s domain password | [optional] 

## Methods

### NewDeleteCatalogModel

`func NewDeleteCatalogModel() *DeleteCatalogModel`

NewDeleteCatalogModel instantiates a new DeleteCatalogModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCatalogModelWithDefaults

`func NewDeleteCatalogModelWithDefaults() *DeleteCatalogModel`

NewDeleteCatalogModelWithDefaults instantiates a new DeleteCatalogModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceDelete

`func (o *DeleteCatalogModel) GetForceDelete() bool`

GetForceDelete returns the ForceDelete field if non-nil, zero value otherwise.

### GetForceDeleteOk

`func (o *DeleteCatalogModel) GetForceDeleteOk() (*bool, bool)`

GetForceDeleteOk returns a tuple with the ForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDelete

`func (o *DeleteCatalogModel) SetForceDelete(v bool)`

SetForceDelete sets ForceDelete field to given value.

### HasForceDelete

`func (o *DeleteCatalogModel) HasForceDelete() bool`

HasForceDelete returns a boolean if a field has been set.

### GetDeleteVm

`func (o *DeleteCatalogModel) GetDeleteVm() bool`

GetDeleteVm returns the DeleteVm field if non-nil, zero value otherwise.

### GetDeleteVmOk

`func (o *DeleteCatalogModel) GetDeleteVmOk() (*bool, bool)`

GetDeleteVmOk returns a tuple with the DeleteVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVm

`func (o *DeleteCatalogModel) SetDeleteVm(v bool)`

SetDeleteVm sets DeleteVm field to given value.

### HasDeleteVm

`func (o *DeleteCatalogModel) HasDeleteVm() bool`

HasDeleteVm returns a boolean if a field has been set.

### GetDeleteResourceLocationIfUnused

`func (o *DeleteCatalogModel) GetDeleteResourceLocationIfUnused() bool`

GetDeleteResourceLocationIfUnused returns the DeleteResourceLocationIfUnused field if non-nil, zero value otherwise.

### GetDeleteResourceLocationIfUnusedOk

`func (o *DeleteCatalogModel) GetDeleteResourceLocationIfUnusedOk() (*bool, bool)`

GetDeleteResourceLocationIfUnusedOk returns a tuple with the DeleteResourceLocationIfUnused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteResourceLocationIfUnused

`func (o *DeleteCatalogModel) SetDeleteResourceLocationIfUnused(v bool)`

SetDeleteResourceLocationIfUnused sets DeleteResourceLocationIfUnused field to given value.

### HasDeleteResourceLocationIfUnused

`func (o *DeleteCatalogModel) HasDeleteResourceLocationIfUnused() bool`

HasDeleteResourceLocationIfUnused returns a boolean if a field has been set.

### GetServiceAccount

`func (o *DeleteCatalogModel) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *DeleteCatalogModel) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *DeleteCatalogModel) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *DeleteCatalogModel) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetServiceAccountPassword

`func (o *DeleteCatalogModel) GetServiceAccountPassword() string`

GetServiceAccountPassword returns the ServiceAccountPassword field if non-nil, zero value otherwise.

### GetServiceAccountPasswordOk

`func (o *DeleteCatalogModel) GetServiceAccountPasswordOk() (*string, bool)`

GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountPassword

`func (o *DeleteCatalogModel) SetServiceAccountPassword(v string)`

SetServiceAccountPassword sets ServiceAccountPassword field to given value.

### HasServiceAccountPassword

`func (o *DeleteCatalogModel) HasServiceAccountPassword() bool`

HasServiceAccountPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


