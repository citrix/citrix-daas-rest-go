# MachineCatalogValidationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumFailures** | **int32** | The number failures of validation. | 
**ValidationFailures** | [**[]MachineCatalogValidationFailureModel**](MachineCatalogValidationFailureModel.md) | The array of validation failures | 

## Methods

### NewMachineCatalogValidationResponseModel

`func NewMachineCatalogValidationResponseModel(numFailures int32, validationFailures []MachineCatalogValidationFailureModel, ) *MachineCatalogValidationResponseModel`

NewMachineCatalogValidationResponseModel instantiates a new MachineCatalogValidationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogValidationResponseModelWithDefaults

`func NewMachineCatalogValidationResponseModelWithDefaults() *MachineCatalogValidationResponseModel`

NewMachineCatalogValidationResponseModelWithDefaults instantiates a new MachineCatalogValidationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumFailures

`func (o *MachineCatalogValidationResponseModel) GetNumFailures() int32`

GetNumFailures returns the NumFailures field if non-nil, zero value otherwise.

### GetNumFailuresOk

`func (o *MachineCatalogValidationResponseModel) GetNumFailuresOk() (*int32, bool)`

GetNumFailuresOk returns a tuple with the NumFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailures

`func (o *MachineCatalogValidationResponseModel) SetNumFailures(v int32)`

SetNumFailures sets NumFailures field to given value.


### GetValidationFailures

`func (o *MachineCatalogValidationResponseModel) GetValidationFailures() []MachineCatalogValidationFailureModel`

GetValidationFailures returns the ValidationFailures field if non-nil, zero value otherwise.

### GetValidationFailuresOk

`func (o *MachineCatalogValidationResponseModel) GetValidationFailuresOk() (*[]MachineCatalogValidationFailureModel, bool)`

GetValidationFailuresOk returns a tuple with the ValidationFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFailures

`func (o *MachineCatalogValidationResponseModel) SetValidationFailures(v []MachineCatalogValidationFailureModel)`

SetValidationFailures sets ValidationFailures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


