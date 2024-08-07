# ServiceAccountCapabilityReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name of the capability. | [optional] 
**EffectiveScope** | Pointer to **NullableString** | Gets or sets the effective scope of the capability. | [optional] 
**IsHealthy** | Pointer to **bool** | Gets or sets the healthy status of the capability. | [optional] 
**FailureReason** | Pointer to **NullableString** | Gets or sets the failure reason of the capability. | [optional] 

## Methods

### NewServiceAccountCapabilityReference

`func NewServiceAccountCapabilityReference() *ServiceAccountCapabilityReference`

NewServiceAccountCapabilityReference instantiates a new ServiceAccountCapabilityReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountCapabilityReferenceWithDefaults

`func NewServiceAccountCapabilityReferenceWithDefaults() *ServiceAccountCapabilityReference`

NewServiceAccountCapabilityReferenceWithDefaults instantiates a new ServiceAccountCapabilityReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceAccountCapabilityReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountCapabilityReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountCapabilityReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceAccountCapabilityReference) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ServiceAccountCapabilityReference) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ServiceAccountCapabilityReference) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEffectiveScope

`func (o *ServiceAccountCapabilityReference) GetEffectiveScope() string`

GetEffectiveScope returns the EffectiveScope field if non-nil, zero value otherwise.

### GetEffectiveScopeOk

`func (o *ServiceAccountCapabilityReference) GetEffectiveScopeOk() (*string, bool)`

GetEffectiveScopeOk returns a tuple with the EffectiveScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveScope

`func (o *ServiceAccountCapabilityReference) SetEffectiveScope(v string)`

SetEffectiveScope sets EffectiveScope field to given value.

### HasEffectiveScope

`func (o *ServiceAccountCapabilityReference) HasEffectiveScope() bool`

HasEffectiveScope returns a boolean if a field has been set.

### SetEffectiveScopeNil

`func (o *ServiceAccountCapabilityReference) SetEffectiveScopeNil(b bool)`

 SetEffectiveScopeNil sets the value for EffectiveScope to be an explicit nil

### UnsetEffectiveScope
`func (o *ServiceAccountCapabilityReference) UnsetEffectiveScope()`

UnsetEffectiveScope ensures that no value is present for EffectiveScope, not even an explicit nil
### GetIsHealthy

`func (o *ServiceAccountCapabilityReference) GetIsHealthy() bool`

GetIsHealthy returns the IsHealthy field if non-nil, zero value otherwise.

### GetIsHealthyOk

`func (o *ServiceAccountCapabilityReference) GetIsHealthyOk() (*bool, bool)`

GetIsHealthyOk returns a tuple with the IsHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHealthy

`func (o *ServiceAccountCapabilityReference) SetIsHealthy(v bool)`

SetIsHealthy sets IsHealthy field to given value.

### HasIsHealthy

`func (o *ServiceAccountCapabilityReference) HasIsHealthy() bool`

HasIsHealthy returns a boolean if a field has been set.

### GetFailureReason

`func (o *ServiceAccountCapabilityReference) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ServiceAccountCapabilityReference) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ServiceAccountCapabilityReference) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ServiceAccountCapabilityReference) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *ServiceAccountCapabilityReference) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *ServiceAccountCapabilityReference) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


