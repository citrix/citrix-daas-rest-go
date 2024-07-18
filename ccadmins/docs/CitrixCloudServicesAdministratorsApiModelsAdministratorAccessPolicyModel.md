# CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ServiceName** | **string** |  | 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Checkable** | [**CitrixCloudServicesAdministratorsApiModelsPolicyProperty**](CitrixCloudServicesAdministratorsApiModelsPolicyProperty.md) |  | 
**ScopeChoices** | Pointer to [**CitrixCloudServicesAdministratorsApiModelsAdministratorAccessScopeChoices**](CitrixCloudServicesAdministratorsApiModelsAdministratorAccessScopeChoices.md) |  | [optional] 

## Methods

### NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel

`func NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel(name string, serviceName string, checkable CitrixCloudServicesAdministratorsApiModelsPolicyProperty, ) *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel`

NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel instantiates a new CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModelWithDefaults

`func NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModelWithDefaults() *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel`

NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModelWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) SetName(v string)`

SetName sets Name field to given value.


### GetServiceName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetCheckable

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) GetCheckable() CitrixCloudServicesAdministratorsApiModelsPolicyProperty`

GetCheckable returns the Checkable field if non-nil, zero value otherwise.

### GetCheckableOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) GetCheckableOk() (*CitrixCloudServicesAdministratorsApiModelsPolicyProperty, bool)`

GetCheckableOk returns a tuple with the Checkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckable

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) SetCheckable(v CitrixCloudServicesAdministratorsApiModelsPolicyProperty)`

SetCheckable sets Checkable field to given value.


### GetScopeChoices

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) GetScopeChoices() CitrixCloudServicesAdministratorsApiModelsAdministratorAccessScopeChoices`

GetScopeChoices returns the ScopeChoices field if non-nil, zero value otherwise.

### GetScopeChoicesOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) GetScopeChoicesOk() (*CitrixCloudServicesAdministratorsApiModelsAdministratorAccessScopeChoices, bool)`

GetScopeChoicesOk returns a tuple with the ScopeChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeChoices

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) SetScopeChoices(v CitrixCloudServicesAdministratorsApiModelsAdministratorAccessScopeChoices)`

SetScopeChoices sets ScopeChoices field to given value.

### HasScopeChoices

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessPolicyModel) HasScopeChoices() bool`

HasScopeChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


