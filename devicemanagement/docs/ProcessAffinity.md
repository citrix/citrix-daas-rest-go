# ProcessAffinity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessName** | Pointer to **string** |  | [optional] 
**Affinity** | Pointer to **int32** |  | [optional] 

## Methods

### NewProcessAffinity

`func NewProcessAffinity() *ProcessAffinity`

NewProcessAffinity instantiates a new ProcessAffinity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessAffinityWithDefaults

`func NewProcessAffinityWithDefaults() *ProcessAffinity`

NewProcessAffinityWithDefaults instantiates a new ProcessAffinity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessName

`func (o *ProcessAffinity) GetProcessName() string`

GetProcessName returns the ProcessName field if non-nil, zero value otherwise.

### GetProcessNameOk

`func (o *ProcessAffinity) GetProcessNameOk() (*string, bool)`

GetProcessNameOk returns a tuple with the ProcessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessName

`func (o *ProcessAffinity) SetProcessName(v string)`

SetProcessName sets ProcessName field to given value.

### HasProcessName

`func (o *ProcessAffinity) HasProcessName() bool`

HasProcessName returns a boolean if a field has been set.

### GetAffinity

`func (o *ProcessAffinity) GetAffinity() int32`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *ProcessAffinity) GetAffinityOk() (*int32, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *ProcessAffinity) SetAffinity(v int32)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *ProcessAffinity) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


