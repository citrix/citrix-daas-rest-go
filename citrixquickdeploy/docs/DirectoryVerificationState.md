# DirectoryVerificationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryId** | Pointer to **string** | ID of the directory | [optional] 
**DirectoryName** | Pointer to **string** | Name of the directory | [optional] 
**DirectoryAssignment** | Pointer to [**DirectoryAssignmentType**](DirectoryAssignmentType.md) | Indicates if the directory is assigned to the pool or a customer | [optional] 
**AssignedCustomerId** | Pointer to **string** | ID of the assigned customer | [optional] [readonly] 
**AssignedCustomerSiteId** | Pointer to **string** | ID of the assigned customer&#39;s site | [optional] [readonly] 
**HasAssociatedSubscriptions** | Pointer to **bool** | Indicates if the Directory is invalid because it does not have any associated subscriptions | [optional] [readonly] 
**Subscriptions** | Pointer to [**[]SubscriptionVerificationState**](SubscriptionVerificationState.md) | List of subscriptions that have unexpected users | [optional] 
**UnexpectedAssignments** | Pointer to [**[]AzureAssignment**](AzureAssignment.md) | List of assignments in the directory that are not expected | [optional] 
**HasUnexpectedAssignments** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewDirectoryVerificationState

`func NewDirectoryVerificationState() *DirectoryVerificationState`

NewDirectoryVerificationState instantiates a new DirectoryVerificationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryVerificationStateWithDefaults

`func NewDirectoryVerificationStateWithDefaults() *DirectoryVerificationState`

NewDirectoryVerificationStateWithDefaults instantiates a new DirectoryVerificationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryId

`func (o *DirectoryVerificationState) GetDirectoryId() string`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *DirectoryVerificationState) GetDirectoryIdOk() (*string, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *DirectoryVerificationState) SetDirectoryId(v string)`

SetDirectoryId sets DirectoryId field to given value.

### HasDirectoryId

`func (o *DirectoryVerificationState) HasDirectoryId() bool`

HasDirectoryId returns a boolean if a field has been set.

### GetDirectoryName

`func (o *DirectoryVerificationState) GetDirectoryName() string`

GetDirectoryName returns the DirectoryName field if non-nil, zero value otherwise.

### GetDirectoryNameOk

`func (o *DirectoryVerificationState) GetDirectoryNameOk() (*string, bool)`

GetDirectoryNameOk returns a tuple with the DirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryName

`func (o *DirectoryVerificationState) SetDirectoryName(v string)`

SetDirectoryName sets DirectoryName field to given value.

### HasDirectoryName

`func (o *DirectoryVerificationState) HasDirectoryName() bool`

HasDirectoryName returns a boolean if a field has been set.

### GetDirectoryAssignment

`func (o *DirectoryVerificationState) GetDirectoryAssignment() DirectoryAssignmentType`

GetDirectoryAssignment returns the DirectoryAssignment field if non-nil, zero value otherwise.

### GetDirectoryAssignmentOk

`func (o *DirectoryVerificationState) GetDirectoryAssignmentOk() (*DirectoryAssignmentType, bool)`

GetDirectoryAssignmentOk returns a tuple with the DirectoryAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryAssignment

`func (o *DirectoryVerificationState) SetDirectoryAssignment(v DirectoryAssignmentType)`

SetDirectoryAssignment sets DirectoryAssignment field to given value.

### HasDirectoryAssignment

`func (o *DirectoryVerificationState) HasDirectoryAssignment() bool`

HasDirectoryAssignment returns a boolean if a field has been set.

### GetAssignedCustomerId

`func (o *DirectoryVerificationState) GetAssignedCustomerId() string`

GetAssignedCustomerId returns the AssignedCustomerId field if non-nil, zero value otherwise.

### GetAssignedCustomerIdOk

`func (o *DirectoryVerificationState) GetAssignedCustomerIdOk() (*string, bool)`

GetAssignedCustomerIdOk returns a tuple with the AssignedCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCustomerId

`func (o *DirectoryVerificationState) SetAssignedCustomerId(v string)`

SetAssignedCustomerId sets AssignedCustomerId field to given value.

### HasAssignedCustomerId

`func (o *DirectoryVerificationState) HasAssignedCustomerId() bool`

HasAssignedCustomerId returns a boolean if a field has been set.

### GetAssignedCustomerSiteId

`func (o *DirectoryVerificationState) GetAssignedCustomerSiteId() string`

GetAssignedCustomerSiteId returns the AssignedCustomerSiteId field if non-nil, zero value otherwise.

### GetAssignedCustomerSiteIdOk

`func (o *DirectoryVerificationState) GetAssignedCustomerSiteIdOk() (*string, bool)`

GetAssignedCustomerSiteIdOk returns a tuple with the AssignedCustomerSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCustomerSiteId

`func (o *DirectoryVerificationState) SetAssignedCustomerSiteId(v string)`

SetAssignedCustomerSiteId sets AssignedCustomerSiteId field to given value.

### HasAssignedCustomerSiteId

`func (o *DirectoryVerificationState) HasAssignedCustomerSiteId() bool`

HasAssignedCustomerSiteId returns a boolean if a field has been set.

### GetHasAssociatedSubscriptions

`func (o *DirectoryVerificationState) GetHasAssociatedSubscriptions() bool`

GetHasAssociatedSubscriptions returns the HasAssociatedSubscriptions field if non-nil, zero value otherwise.

### GetHasAssociatedSubscriptionsOk

`func (o *DirectoryVerificationState) GetHasAssociatedSubscriptionsOk() (*bool, bool)`

GetHasAssociatedSubscriptionsOk returns a tuple with the HasAssociatedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAssociatedSubscriptions

`func (o *DirectoryVerificationState) SetHasAssociatedSubscriptions(v bool)`

SetHasAssociatedSubscriptions sets HasAssociatedSubscriptions field to given value.

### HasHasAssociatedSubscriptions

`func (o *DirectoryVerificationState) HasHasAssociatedSubscriptions() bool`

HasHasAssociatedSubscriptions returns a boolean if a field has been set.

### GetSubscriptions

`func (o *DirectoryVerificationState) GetSubscriptions() []SubscriptionVerificationState`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *DirectoryVerificationState) GetSubscriptionsOk() (*[]SubscriptionVerificationState, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *DirectoryVerificationState) SetSubscriptions(v []SubscriptionVerificationState)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *DirectoryVerificationState) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetUnexpectedAssignments

`func (o *DirectoryVerificationState) GetUnexpectedAssignments() []AzureAssignment`

GetUnexpectedAssignments returns the UnexpectedAssignments field if non-nil, zero value otherwise.

### GetUnexpectedAssignmentsOk

`func (o *DirectoryVerificationState) GetUnexpectedAssignmentsOk() (*[]AzureAssignment, bool)`

GetUnexpectedAssignmentsOk returns a tuple with the UnexpectedAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexpectedAssignments

`func (o *DirectoryVerificationState) SetUnexpectedAssignments(v []AzureAssignment)`

SetUnexpectedAssignments sets UnexpectedAssignments field to given value.

### HasUnexpectedAssignments

`func (o *DirectoryVerificationState) HasUnexpectedAssignments() bool`

HasUnexpectedAssignments returns a boolean if a field has been set.

### GetHasUnexpectedAssignments

`func (o *DirectoryVerificationState) GetHasUnexpectedAssignments() bool`

GetHasUnexpectedAssignments returns the HasUnexpectedAssignments field if non-nil, zero value otherwise.

### GetHasUnexpectedAssignmentsOk

`func (o *DirectoryVerificationState) GetHasUnexpectedAssignmentsOk() (*bool, bool)`

GetHasUnexpectedAssignmentsOk returns a tuple with the HasUnexpectedAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnexpectedAssignments

`func (o *DirectoryVerificationState) SetHasUnexpectedAssignments(v bool)`

SetHasUnexpectedAssignments sets HasUnexpectedAssignments field to given value.

### HasHasUnexpectedAssignments

`func (o *DirectoryVerificationState) HasHasUnexpectedAssignments() bool`

HasHasUnexpectedAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


