# CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ServiceName** | **string** |  | 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Checkable** | [**CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty**](CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty.md) |  | 
**ScopeChoices** | Pointer to [**CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessScopeChoices**](CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessScopeChoices.md) |  | [optional] 

## Methods

### NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel

`func NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel(name string, serviceName string, checkable CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty, ) *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel`

NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel instantiates a new CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModelWithDefaults

`func NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModelWithDefaults() *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel`

NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModelWithDefaults instantiates a new CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) SetName(v string)`

SetName sets Name field to given value.


### GetServiceName

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetDisplayName

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetCheckable

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) GetCheckable() CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty`

GetCheckable returns the Checkable field if non-nil, zero value otherwise.

### GetCheckableOk

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) GetCheckableOk() (*CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty, bool)`

GetCheckableOk returns a tuple with the Checkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckable

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) SetCheckable(v CitrixCloudServicesServiceRuntimePolicySetsModelsPolicyProperty)`

SetCheckable sets Checkable field to given value.


### GetScopeChoices

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) GetScopeChoices() CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessScopeChoices`

GetScopeChoices returns the ScopeChoices field if non-nil, zero value otherwise.

### GetScopeChoicesOk

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) GetScopeChoicesOk() (*CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessScopeChoices, bool)`

GetScopeChoicesOk returns a tuple with the ScopeChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeChoices

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) SetScopeChoices(v CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessScopeChoices)`

SetScopeChoices sets ScopeChoices field to given value.

### HasScopeChoices

`func (o *CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessPolicyModel) HasScopeChoices() bool`

HasScopeChoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


