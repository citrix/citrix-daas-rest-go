# AzureTagAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**PoolName** | Pointer to **string** |  | [optional] 
**IssuesFound** | Pointer to **bool** |  | [optional] 
**AuditCompleted** | Pointer to **bool** |  | [optional] 
**AuditErrors** | Pointer to **string** |  | [optional] 
**IsActiveCustomerAssigned** | Pointer to **bool** |  | [optional] 
**IsActiveCspCustomer** | Pointer to **bool** |  | [optional] 
**IsAzurePoolAssigned** | Pointer to **bool** |  | [optional] 
**NumberOfPolicies** | Pointer to **int32** |  | [optional] 
**NumberOfPolicyDefinitions** | Pointer to **int32** |  | [optional] 
**NumberOfResourceGroups** | Pointer to **int32** |  | [optional] 
**NumberOfResources** | Pointer to **int32** |  | [optional] 
**Issues** | Pointer to [**[]AzureResourceTagAudit**](AzureResourceTagAudit.md) |  | [optional] 

## Methods

### NewAzureTagAudit

`func NewAzureTagAudit() *AzureTagAudit`

NewAzureTagAudit instantiates a new AzureTagAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureTagAuditWithDefaults

`func NewAzureTagAuditWithDefaults() *AzureTagAudit`

NewAzureTagAuditWithDefaults instantiates a new AzureTagAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryId

`func (o *AzureTagAudit) GetDirectoryId() string`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *AzureTagAudit) GetDirectoryIdOk() (*string, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *AzureTagAudit) SetDirectoryId(v string)`

SetDirectoryId sets DirectoryId field to given value.

### HasDirectoryId

`func (o *AzureTagAudit) HasDirectoryId() bool`

HasDirectoryId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AzureTagAudit) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureTagAudit) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureTagAudit) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureTagAudit) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AzureTagAudit) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AzureTagAudit) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AzureTagAudit) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AzureTagAudit) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetSiteId

`func (o *AzureTagAudit) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AzureTagAudit) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AzureTagAudit) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AzureTagAudit) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetPoolName

`func (o *AzureTagAudit) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AzureTagAudit) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AzureTagAudit) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AzureTagAudit) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetIssuesFound

`func (o *AzureTagAudit) GetIssuesFound() bool`

GetIssuesFound returns the IssuesFound field if non-nil, zero value otherwise.

### GetIssuesFoundOk

`func (o *AzureTagAudit) GetIssuesFoundOk() (*bool, bool)`

GetIssuesFoundOk returns a tuple with the IssuesFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesFound

`func (o *AzureTagAudit) SetIssuesFound(v bool)`

SetIssuesFound sets IssuesFound field to given value.

### HasIssuesFound

`func (o *AzureTagAudit) HasIssuesFound() bool`

HasIssuesFound returns a boolean if a field has been set.

### GetAuditCompleted

`func (o *AzureTagAudit) GetAuditCompleted() bool`

GetAuditCompleted returns the AuditCompleted field if non-nil, zero value otherwise.

### GetAuditCompletedOk

`func (o *AzureTagAudit) GetAuditCompletedOk() (*bool, bool)`

GetAuditCompletedOk returns a tuple with the AuditCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditCompleted

`func (o *AzureTagAudit) SetAuditCompleted(v bool)`

SetAuditCompleted sets AuditCompleted field to given value.

### HasAuditCompleted

`func (o *AzureTagAudit) HasAuditCompleted() bool`

HasAuditCompleted returns a boolean if a field has been set.

### GetAuditErrors

`func (o *AzureTagAudit) GetAuditErrors() string`

GetAuditErrors returns the AuditErrors field if non-nil, zero value otherwise.

### GetAuditErrorsOk

`func (o *AzureTagAudit) GetAuditErrorsOk() (*string, bool)`

GetAuditErrorsOk returns a tuple with the AuditErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditErrors

`func (o *AzureTagAudit) SetAuditErrors(v string)`

SetAuditErrors sets AuditErrors field to given value.

### HasAuditErrors

`func (o *AzureTagAudit) HasAuditErrors() bool`

HasAuditErrors returns a boolean if a field has been set.

### GetIsActiveCustomerAssigned

`func (o *AzureTagAudit) GetIsActiveCustomerAssigned() bool`

GetIsActiveCustomerAssigned returns the IsActiveCustomerAssigned field if non-nil, zero value otherwise.

### GetIsActiveCustomerAssignedOk

`func (o *AzureTagAudit) GetIsActiveCustomerAssignedOk() (*bool, bool)`

GetIsActiveCustomerAssignedOk returns a tuple with the IsActiveCustomerAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActiveCustomerAssigned

`func (o *AzureTagAudit) SetIsActiveCustomerAssigned(v bool)`

SetIsActiveCustomerAssigned sets IsActiveCustomerAssigned field to given value.

### HasIsActiveCustomerAssigned

`func (o *AzureTagAudit) HasIsActiveCustomerAssigned() bool`

HasIsActiveCustomerAssigned returns a boolean if a field has been set.

### GetIsActiveCspCustomer

`func (o *AzureTagAudit) GetIsActiveCspCustomer() bool`

GetIsActiveCspCustomer returns the IsActiveCspCustomer field if non-nil, zero value otherwise.

### GetIsActiveCspCustomerOk

`func (o *AzureTagAudit) GetIsActiveCspCustomerOk() (*bool, bool)`

GetIsActiveCspCustomerOk returns a tuple with the IsActiveCspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActiveCspCustomer

`func (o *AzureTagAudit) SetIsActiveCspCustomer(v bool)`

SetIsActiveCspCustomer sets IsActiveCspCustomer field to given value.

### HasIsActiveCspCustomer

`func (o *AzureTagAudit) HasIsActiveCspCustomer() bool`

HasIsActiveCspCustomer returns a boolean if a field has been set.

### GetIsAzurePoolAssigned

`func (o *AzureTagAudit) GetIsAzurePoolAssigned() bool`

GetIsAzurePoolAssigned returns the IsAzurePoolAssigned field if non-nil, zero value otherwise.

### GetIsAzurePoolAssignedOk

`func (o *AzureTagAudit) GetIsAzurePoolAssignedOk() (*bool, bool)`

GetIsAzurePoolAssignedOk returns a tuple with the IsAzurePoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAzurePoolAssigned

`func (o *AzureTagAudit) SetIsAzurePoolAssigned(v bool)`

SetIsAzurePoolAssigned sets IsAzurePoolAssigned field to given value.

### HasIsAzurePoolAssigned

`func (o *AzureTagAudit) HasIsAzurePoolAssigned() bool`

HasIsAzurePoolAssigned returns a boolean if a field has been set.

### GetNumberOfPolicies

`func (o *AzureTagAudit) GetNumberOfPolicies() int32`

GetNumberOfPolicies returns the NumberOfPolicies field if non-nil, zero value otherwise.

### GetNumberOfPoliciesOk

`func (o *AzureTagAudit) GetNumberOfPoliciesOk() (*int32, bool)`

GetNumberOfPoliciesOk returns a tuple with the NumberOfPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPolicies

`func (o *AzureTagAudit) SetNumberOfPolicies(v int32)`

SetNumberOfPolicies sets NumberOfPolicies field to given value.

### HasNumberOfPolicies

`func (o *AzureTagAudit) HasNumberOfPolicies() bool`

HasNumberOfPolicies returns a boolean if a field has been set.

### GetNumberOfPolicyDefinitions

`func (o *AzureTagAudit) GetNumberOfPolicyDefinitions() int32`

GetNumberOfPolicyDefinitions returns the NumberOfPolicyDefinitions field if non-nil, zero value otherwise.

### GetNumberOfPolicyDefinitionsOk

`func (o *AzureTagAudit) GetNumberOfPolicyDefinitionsOk() (*int32, bool)`

GetNumberOfPolicyDefinitionsOk returns a tuple with the NumberOfPolicyDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPolicyDefinitions

`func (o *AzureTagAudit) SetNumberOfPolicyDefinitions(v int32)`

SetNumberOfPolicyDefinitions sets NumberOfPolicyDefinitions field to given value.

### HasNumberOfPolicyDefinitions

`func (o *AzureTagAudit) HasNumberOfPolicyDefinitions() bool`

HasNumberOfPolicyDefinitions returns a boolean if a field has been set.

### GetNumberOfResourceGroups

`func (o *AzureTagAudit) GetNumberOfResourceGroups() int32`

GetNumberOfResourceGroups returns the NumberOfResourceGroups field if non-nil, zero value otherwise.

### GetNumberOfResourceGroupsOk

`func (o *AzureTagAudit) GetNumberOfResourceGroupsOk() (*int32, bool)`

GetNumberOfResourceGroupsOk returns a tuple with the NumberOfResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfResourceGroups

`func (o *AzureTagAudit) SetNumberOfResourceGroups(v int32)`

SetNumberOfResourceGroups sets NumberOfResourceGroups field to given value.

### HasNumberOfResourceGroups

`func (o *AzureTagAudit) HasNumberOfResourceGroups() bool`

HasNumberOfResourceGroups returns a boolean if a field has been set.

### GetNumberOfResources

`func (o *AzureTagAudit) GetNumberOfResources() int32`

GetNumberOfResources returns the NumberOfResources field if non-nil, zero value otherwise.

### GetNumberOfResourcesOk

`func (o *AzureTagAudit) GetNumberOfResourcesOk() (*int32, bool)`

GetNumberOfResourcesOk returns a tuple with the NumberOfResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfResources

`func (o *AzureTagAudit) SetNumberOfResources(v int32)`

SetNumberOfResources sets NumberOfResources field to given value.

### HasNumberOfResources

`func (o *AzureTagAudit) HasNumberOfResources() bool`

HasNumberOfResources returns a boolean if a field has been set.

### GetIssues

`func (o *AzureTagAudit) GetIssues() []AzureResourceTagAudit`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *AzureTagAudit) GetIssuesOk() (*[]AzureResourceTagAudit, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *AzureTagAudit) SetIssues(v []AzureResourceTagAudit)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *AzureTagAudit) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


