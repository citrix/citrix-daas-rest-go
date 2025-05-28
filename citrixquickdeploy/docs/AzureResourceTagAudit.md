# AzureResourceTagAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AzureResourceTagAuditType**](AzureResourceTagAuditType.md) |  | [optional] 
**Violation** | Pointer to [**AzureResourceTagAuditViolation**](AzureResourceTagAuditViolation.md) |  | [optional] 
**ResourceIdTag** | Pointer to **string** |  | [optional] 
**TagName** | Pointer to **string** |  | [optional] 
**TagValue** | Pointer to **string** |  | [optional] 
**CorrectTagValue** | Pointer to **string** |  | [optional] 
**RepairCompleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewAzureResourceTagAudit

`func NewAzureResourceTagAudit() *AzureResourceTagAudit`

NewAzureResourceTagAudit instantiates a new AzureResourceTagAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureResourceTagAuditWithDefaults

`func NewAzureResourceTagAuditWithDefaults() *AzureResourceTagAudit`

NewAzureResourceTagAuditWithDefaults instantiates a new AzureResourceTagAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AzureResourceTagAudit) GetType() AzureResourceTagAuditType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureResourceTagAudit) GetTypeOk() (*AzureResourceTagAuditType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureResourceTagAudit) SetType(v AzureResourceTagAuditType)`

SetType sets Type field to given value.

### HasType

`func (o *AzureResourceTagAudit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetViolation

`func (o *AzureResourceTagAudit) GetViolation() AzureResourceTagAuditViolation`

GetViolation returns the Violation field if non-nil, zero value otherwise.

### GetViolationOk

`func (o *AzureResourceTagAudit) GetViolationOk() (*AzureResourceTagAuditViolation, bool)`

GetViolationOk returns a tuple with the Violation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolation

`func (o *AzureResourceTagAudit) SetViolation(v AzureResourceTagAuditViolation)`

SetViolation sets Violation field to given value.

### HasViolation

`func (o *AzureResourceTagAudit) HasViolation() bool`

HasViolation returns a boolean if a field has been set.

### GetResourceIdTag

`func (o *AzureResourceTagAudit) GetResourceIdTag() string`

GetResourceIdTag returns the ResourceIdTag field if non-nil, zero value otherwise.

### GetResourceIdTagOk

`func (o *AzureResourceTagAudit) GetResourceIdTagOk() (*string, bool)`

GetResourceIdTagOk returns a tuple with the ResourceIdTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIdTag

`func (o *AzureResourceTagAudit) SetResourceIdTag(v string)`

SetResourceIdTag sets ResourceIdTag field to given value.

### HasResourceIdTag

`func (o *AzureResourceTagAudit) HasResourceIdTag() bool`

HasResourceIdTag returns a boolean if a field has been set.

### GetTagName

`func (o *AzureResourceTagAudit) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *AzureResourceTagAudit) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *AzureResourceTagAudit) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *AzureResourceTagAudit) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetTagValue

`func (o *AzureResourceTagAudit) GetTagValue() string`

GetTagValue returns the TagValue field if non-nil, zero value otherwise.

### GetTagValueOk

`func (o *AzureResourceTagAudit) GetTagValueOk() (*string, bool)`

GetTagValueOk returns a tuple with the TagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValue

`func (o *AzureResourceTagAudit) SetTagValue(v string)`

SetTagValue sets TagValue field to given value.

### HasTagValue

`func (o *AzureResourceTagAudit) HasTagValue() bool`

HasTagValue returns a boolean if a field has been set.

### GetCorrectTagValue

`func (o *AzureResourceTagAudit) GetCorrectTagValue() string`

GetCorrectTagValue returns the CorrectTagValue field if non-nil, zero value otherwise.

### GetCorrectTagValueOk

`func (o *AzureResourceTagAudit) GetCorrectTagValueOk() (*string, bool)`

GetCorrectTagValueOk returns a tuple with the CorrectTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectTagValue

`func (o *AzureResourceTagAudit) SetCorrectTagValue(v string)`

SetCorrectTagValue sets CorrectTagValue field to given value.

### HasCorrectTagValue

`func (o *AzureResourceTagAudit) HasCorrectTagValue() bool`

HasCorrectTagValue returns a boolean if a field has been set.

### GetRepairCompleted

`func (o *AzureResourceTagAudit) GetRepairCompleted() bool`

GetRepairCompleted returns the RepairCompleted field if non-nil, zero value otherwise.

### GetRepairCompletedOk

`func (o *AzureResourceTagAudit) GetRepairCompletedOk() (*bool, bool)`

GetRepairCompletedOk returns a tuple with the RepairCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairCompleted

`func (o *AzureResourceTagAudit) SetRepairCompleted(v bool)`

SetRepairCompleted sets RepairCompleted field to given value.

### HasRepairCompleted

`func (o *AzureResourceTagAudit) HasRepairCompleted() bool`

HasRepairCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


