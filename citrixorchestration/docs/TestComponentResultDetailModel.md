# TestComponentResultDetailModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceSource** | Pointer to **string** | The service source of any error information present in this record. Typically, this will be the service that ran the test.  In some cases, particularly where an orphaned object target was detected and the test could not be run at all, the service source may be the framework itself, EnvTestService. | [optional] 
**Serverity** | Pointer to **string** | The severity of the problem. | [optional] 
**Explanation** | Pointer to **string** | The explanation of the problem. | [optional] 
**Action** | Pointer to **string** | The action to be take to fix this problem. | [optional] 

## Methods

### NewTestComponentResultDetailModel

`func NewTestComponentResultDetailModel() *TestComponentResultDetailModel`

NewTestComponentResultDetailModel instantiates a new TestComponentResultDetailModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestComponentResultDetailModelWithDefaults

`func NewTestComponentResultDetailModelWithDefaults() *TestComponentResultDetailModel`

NewTestComponentResultDetailModelWithDefaults instantiates a new TestComponentResultDetailModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceSource

`func (o *TestComponentResultDetailModel) GetServiceSource() string`

GetServiceSource returns the ServiceSource field if non-nil, zero value otherwise.

### GetServiceSourceOk

`func (o *TestComponentResultDetailModel) GetServiceSourceOk() (*string, bool)`

GetServiceSourceOk returns a tuple with the ServiceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSource

`func (o *TestComponentResultDetailModel) SetServiceSource(v string)`

SetServiceSource sets ServiceSource field to given value.

### HasServiceSource

`func (o *TestComponentResultDetailModel) HasServiceSource() bool`

HasServiceSource returns a boolean if a field has been set.

### GetServerity

`func (o *TestComponentResultDetailModel) GetServerity() string`

GetServerity returns the Serverity field if non-nil, zero value otherwise.

### GetServerityOk

`func (o *TestComponentResultDetailModel) GetServerityOk() (*string, bool)`

GetServerityOk returns a tuple with the Serverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerity

`func (o *TestComponentResultDetailModel) SetServerity(v string)`

SetServerity sets Serverity field to given value.

### HasServerity

`func (o *TestComponentResultDetailModel) HasServerity() bool`

HasServerity returns a boolean if a field has been set.

### GetExplanation

`func (o *TestComponentResultDetailModel) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *TestComponentResultDetailModel) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *TestComponentResultDetailModel) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *TestComponentResultDetailModel) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetAction

`func (o *TestComponentResultDetailModel) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TestComponentResultDetailModel) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TestComponentResultDetailModel) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TestComponentResultDetailModel) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


