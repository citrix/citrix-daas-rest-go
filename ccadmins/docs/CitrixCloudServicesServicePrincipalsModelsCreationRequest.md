# CitrixCloudServicesServicePrincipalsModelsCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AdminAccess** | [**CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel**](CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel.md) |  | 
**ExpirationDate** | **time.Time** |  | 

## Methods

### NewCitrixCloudServicesServicePrincipalsModelsCreationRequest

`func NewCitrixCloudServicesServicePrincipalsModelsCreationRequest(name string, adminAccess CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel, expirationDate time.Time, ) *CitrixCloudServicesServicePrincipalsModelsCreationRequest`

NewCitrixCloudServicesServicePrincipalsModelsCreationRequest instantiates a new CitrixCloudServicesServicePrincipalsModelsCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesServicePrincipalsModelsCreationRequestWithDefaults

`func NewCitrixCloudServicesServicePrincipalsModelsCreationRequestWithDefaults() *CitrixCloudServicesServicePrincipalsModelsCreationRequest`

NewCitrixCloudServicesServicePrincipalsModelsCreationRequestWithDefaults instantiates a new CitrixCloudServicesServicePrincipalsModelsCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CitrixCloudServicesServicePrincipalsModelsCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CitrixCloudServicesServicePrincipalsModelsCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CitrixCloudServicesServicePrincipalsModelsCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAdminAccess

`func (o *CitrixCloudServicesServicePrincipalsModelsCreationRequest) GetAdminAccess() CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel`

GetAdminAccess returns the AdminAccess field if non-nil, zero value otherwise.

### GetAdminAccessOk

`func (o *CitrixCloudServicesServicePrincipalsModelsCreationRequest) GetAdminAccessOk() (*CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel, bool)`

GetAdminAccessOk returns a tuple with the AdminAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAccess

`func (o *CitrixCloudServicesServicePrincipalsModelsCreationRequest) SetAdminAccess(v CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel)`

SetAdminAccess sets AdminAccess field to given value.


### GetExpirationDate

`func (o *CitrixCloudServicesServicePrincipalsModelsCreationRequest) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CitrixCloudServicesServicePrincipalsModelsCreationRequest) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CitrixCloudServicesServicePrincipalsModelsCreationRequest) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


