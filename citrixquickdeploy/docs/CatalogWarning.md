# CatalogWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CatalogWarningType**](CatalogWarningType.md) | Type of warning | [optional] 
**IsDismissible** | Pointer to **bool** | Indicates if the warning message can be cleared out by the user | [optional] 
**WarningMessage** | Pointer to **string** | Warning message to display to the user | [optional] 
**IsError** | Pointer to **bool** |  | [optional] 

## Methods

### NewCatalogWarning

`func NewCatalogWarning() *CatalogWarning`

NewCatalogWarning instantiates a new CatalogWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogWarningWithDefaults

`func NewCatalogWarningWithDefaults() *CatalogWarning`

NewCatalogWarningWithDefaults instantiates a new CatalogWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogWarning) GetType() CatalogWarningType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogWarning) GetTypeOk() (*CatalogWarningType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogWarning) SetType(v CatalogWarningType)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogWarning) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsDismissible

`func (o *CatalogWarning) GetIsDismissible() bool`

GetIsDismissible returns the IsDismissible field if non-nil, zero value otherwise.

### GetIsDismissibleOk

`func (o *CatalogWarning) GetIsDismissibleOk() (*bool, bool)`

GetIsDismissibleOk returns a tuple with the IsDismissible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDismissible

`func (o *CatalogWarning) SetIsDismissible(v bool)`

SetIsDismissible sets IsDismissible field to given value.

### HasIsDismissible

`func (o *CatalogWarning) HasIsDismissible() bool`

HasIsDismissible returns a boolean if a field has been set.

### GetWarningMessage

`func (o *CatalogWarning) GetWarningMessage() string`

GetWarningMessage returns the WarningMessage field if non-nil, zero value otherwise.

### GetWarningMessageOk

`func (o *CatalogWarning) GetWarningMessageOk() (*string, bool)`

GetWarningMessageOk returns a tuple with the WarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessage

`func (o *CatalogWarning) SetWarningMessage(v string)`

SetWarningMessage sets WarningMessage field to given value.

### HasWarningMessage

`func (o *CatalogWarning) HasWarningMessage() bool`

HasWarningMessage returns a boolean if a field has been set.

### GetIsError

`func (o *CatalogWarning) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *CatalogWarning) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *CatalogWarning) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *CatalogWarning) HasIsError() bool`

HasIsError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


