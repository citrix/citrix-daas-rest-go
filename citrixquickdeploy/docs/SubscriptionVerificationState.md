# SubscriptionVerificationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | ID of the subscription | [optional] 
**UnexpectedAssignments** | Pointer to [**[]AzureAssignment**](AzureAssignment.md) | List of assignments in the subscription that are not expected | [optional] 
**HasUnexpectedAssignments** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewSubscriptionVerificationState

`func NewSubscriptionVerificationState() *SubscriptionVerificationState`

NewSubscriptionVerificationState instantiates a new SubscriptionVerificationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionVerificationStateWithDefaults

`func NewSubscriptionVerificationStateWithDefaults() *SubscriptionVerificationState`

NewSubscriptionVerificationStateWithDefaults instantiates a new SubscriptionVerificationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *SubscriptionVerificationState) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionVerificationState) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionVerificationState) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SubscriptionVerificationState) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUnexpectedAssignments

`func (o *SubscriptionVerificationState) GetUnexpectedAssignments() []AzureAssignment`

GetUnexpectedAssignments returns the UnexpectedAssignments field if non-nil, zero value otherwise.

### GetUnexpectedAssignmentsOk

`func (o *SubscriptionVerificationState) GetUnexpectedAssignmentsOk() (*[]AzureAssignment, bool)`

GetUnexpectedAssignmentsOk returns a tuple with the UnexpectedAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexpectedAssignments

`func (o *SubscriptionVerificationState) SetUnexpectedAssignments(v []AzureAssignment)`

SetUnexpectedAssignments sets UnexpectedAssignments field to given value.

### HasUnexpectedAssignments

`func (o *SubscriptionVerificationState) HasUnexpectedAssignments() bool`

HasUnexpectedAssignments returns a boolean if a field has been set.

### GetHasUnexpectedAssignments

`func (o *SubscriptionVerificationState) GetHasUnexpectedAssignments() bool`

GetHasUnexpectedAssignments returns the HasUnexpectedAssignments field if non-nil, zero value otherwise.

### GetHasUnexpectedAssignmentsOk

`func (o *SubscriptionVerificationState) GetHasUnexpectedAssignmentsOk() (*bool, bool)`

GetHasUnexpectedAssignmentsOk returns a tuple with the HasUnexpectedAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnexpectedAssignments

`func (o *SubscriptionVerificationState) SetHasUnexpectedAssignments(v bool)`

SetHasUnexpectedAssignments sets HasUnexpectedAssignments field to given value.

### HasHasUnexpectedAssignments

`func (o *SubscriptionVerificationState) HasHasUnexpectedAssignments() bool`

HasHasUnexpectedAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


