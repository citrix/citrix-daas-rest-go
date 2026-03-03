# AzureTagAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**SiteId** | Pointer to **NullableString** |  | [optional] 
**PoolName** | Pointer to **NullableString** |  | [optional] 
**IssuesFound** | Pointer to **NullableBool** |  | [optional] 
**AuditCompleted** | Pointer to **NullableBool** |  | [optional] 
**AuditErrors** | Pointer to **NullableString** |  | [optional] 
**IsActiveCustomerAssigned** | Pointer to **NullableBool** |  | [optional] 
**IsActiveCspCustomer** | Pointer to **NullableBool** |  | [optional] 
**IsAzurePoolAssigned** | Pointer to **NullableBool** |  | [optional] 
**NumberOfPolicies** | Pointer to **NullableInt32** |  | [optional] 
**NumberOfPolicyDefinitions** | Pointer to **NullableInt32** |  | [optional] 
**NumberOfResourceGroups** | Pointer to **NullableInt32** |  | [optional] 
**NumberOfResources** | Pointer to **NullableInt32** |  | [optional] 
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

### SetDirectoryIdNil

`func (o *AzureTagAudit) SetDirectoryIdNil(b bool)`

 SetDirectoryIdNil sets the value for DirectoryId to be an explicit nil

### UnsetDirectoryId
`func (o *AzureTagAudit) UnsetDirectoryId()`

UnsetDirectoryId ensures that no value is present for DirectoryId, not even an explicit nil
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

### SetSubscriptionIdNil

`func (o *AzureTagAudit) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AzureTagAudit) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
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

### SetCustomerIdNil

`func (o *AzureTagAudit) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *AzureTagAudit) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
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

### SetSiteIdNil

`func (o *AzureTagAudit) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *AzureTagAudit) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
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

### SetPoolNameNil

`func (o *AzureTagAudit) SetPoolNameNil(b bool)`

 SetPoolNameNil sets the value for PoolName to be an explicit nil

### UnsetPoolName
`func (o *AzureTagAudit) UnsetPoolName()`

UnsetPoolName ensures that no value is present for PoolName, not even an explicit nil
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

### SetIssuesFoundNil

`func (o *AzureTagAudit) SetIssuesFoundNil(b bool)`

 SetIssuesFoundNil sets the value for IssuesFound to be an explicit nil

### UnsetIssuesFound
`func (o *AzureTagAudit) UnsetIssuesFound()`

UnsetIssuesFound ensures that no value is present for IssuesFound, not even an explicit nil
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

### SetAuditCompletedNil

`func (o *AzureTagAudit) SetAuditCompletedNil(b bool)`

 SetAuditCompletedNil sets the value for AuditCompleted to be an explicit nil

### UnsetAuditCompleted
`func (o *AzureTagAudit) UnsetAuditCompleted()`

UnsetAuditCompleted ensures that no value is present for AuditCompleted, not even an explicit nil
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

### SetAuditErrorsNil

`func (o *AzureTagAudit) SetAuditErrorsNil(b bool)`

 SetAuditErrorsNil sets the value for AuditErrors to be an explicit nil

### UnsetAuditErrors
`func (o *AzureTagAudit) UnsetAuditErrors()`

UnsetAuditErrors ensures that no value is present for AuditErrors, not even an explicit nil
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

### SetIsActiveCustomerAssignedNil

`func (o *AzureTagAudit) SetIsActiveCustomerAssignedNil(b bool)`

 SetIsActiveCustomerAssignedNil sets the value for IsActiveCustomerAssigned to be an explicit nil

### UnsetIsActiveCustomerAssigned
`func (o *AzureTagAudit) UnsetIsActiveCustomerAssigned()`

UnsetIsActiveCustomerAssigned ensures that no value is present for IsActiveCustomerAssigned, not even an explicit nil
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

### SetIsActiveCspCustomerNil

`func (o *AzureTagAudit) SetIsActiveCspCustomerNil(b bool)`

 SetIsActiveCspCustomerNil sets the value for IsActiveCspCustomer to be an explicit nil

### UnsetIsActiveCspCustomer
`func (o *AzureTagAudit) UnsetIsActiveCspCustomer()`

UnsetIsActiveCspCustomer ensures that no value is present for IsActiveCspCustomer, not even an explicit nil
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

### SetIsAzurePoolAssignedNil

`func (o *AzureTagAudit) SetIsAzurePoolAssignedNil(b bool)`

 SetIsAzurePoolAssignedNil sets the value for IsAzurePoolAssigned to be an explicit nil

### UnsetIsAzurePoolAssigned
`func (o *AzureTagAudit) UnsetIsAzurePoolAssigned()`

UnsetIsAzurePoolAssigned ensures that no value is present for IsAzurePoolAssigned, not even an explicit nil
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

### SetNumberOfPoliciesNil

`func (o *AzureTagAudit) SetNumberOfPoliciesNil(b bool)`

 SetNumberOfPoliciesNil sets the value for NumberOfPolicies to be an explicit nil

### UnsetNumberOfPolicies
`func (o *AzureTagAudit) UnsetNumberOfPolicies()`

UnsetNumberOfPolicies ensures that no value is present for NumberOfPolicies, not even an explicit nil
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

### SetNumberOfPolicyDefinitionsNil

`func (o *AzureTagAudit) SetNumberOfPolicyDefinitionsNil(b bool)`

 SetNumberOfPolicyDefinitionsNil sets the value for NumberOfPolicyDefinitions to be an explicit nil

### UnsetNumberOfPolicyDefinitions
`func (o *AzureTagAudit) UnsetNumberOfPolicyDefinitions()`

UnsetNumberOfPolicyDefinitions ensures that no value is present for NumberOfPolicyDefinitions, not even an explicit nil
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

### SetNumberOfResourceGroupsNil

`func (o *AzureTagAudit) SetNumberOfResourceGroupsNil(b bool)`

 SetNumberOfResourceGroupsNil sets the value for NumberOfResourceGroups to be an explicit nil

### UnsetNumberOfResourceGroups
`func (o *AzureTagAudit) UnsetNumberOfResourceGroups()`

UnsetNumberOfResourceGroups ensures that no value is present for NumberOfResourceGroups, not even an explicit nil
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

### SetNumberOfResourcesNil

`func (o *AzureTagAudit) SetNumberOfResourcesNil(b bool)`

 SetNumberOfResourcesNil sets the value for NumberOfResources to be an explicit nil

### UnsetNumberOfResources
`func (o *AzureTagAudit) UnsetNumberOfResources()`

UnsetNumberOfResources ensures that no value is present for NumberOfResources, not even an explicit nil
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

### SetIssuesNil

`func (o *AzureTagAudit) SetIssuesNil(b bool)`

 SetIssuesNil sets the value for Issues to be an explicit nil

### UnsetIssues
`func (o *AzureTagAudit) UnsetIssues()`

UnsetIssues ensures that no value is present for Issues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


