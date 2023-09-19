# LowLevelOperationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminMachineIP** | Pointer to **NullableString** | IP address of the admin machine from which the operation was performed.              | [optional] 
**Details** | Pointer to [**[]LowLevelOperationDetailResponseModel**](LowLevelOperationDetailResponseModel.md) | Details.              | [optional] 
**EndTime** | Pointer to **NullableString** | Time when the operation ended. If the operation is incomplete, will be null.              | [optional] 
**FormattedEndTime** | Pointer to **NullableString** | Formatted time when the operation ended. RFC 3339 compatible format. If the operation is incomplete, will be null.              | [optional] 
**Id** | **string** | ID of the logged operation.              | 
**IsSuccessful** | Pointer to **NullableBool** | Indicates whether the operation completed successfully.  If the operation is incomplete, will be null.              | [optional] 
**OperationId** | **string** | The id of the (high level) operation.              | 
**OperationType** | [**LogOperationType**](LogOperationType.md) |  | 
**Parameters** | [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Operation parameters.              | 
**Source** | **string** | Source of the operation.              | 
**SourceSdk** | Pointer to **NullableString** | The source SDK.              | [optional] 
**StartTime** | **string** | Time when the operation started.              | 
**FormattedStartTime** | **string** | Formatted time when the operation started. RFC 3339 compatible format.              | 
**TargetTypes** | **[]string** | The type(s) of object which were the target of the configuration change. For example, \&quot;Session\&quot; or \&quot;Machine\&quot;.              | 
**Text** | **string** | Human-readable description of the change.              | 
**User** | Pointer to **NullableString** | User who performed the change.              | [optional] 

## Methods

### NewLowLevelOperationResponseModel

`func NewLowLevelOperationResponseModel(id string, operationId string, operationType LogOperationType, parameters []NameValueStringPairModel, source string, startTime string, formattedStartTime string, targetTypes []string, text string, ) *LowLevelOperationResponseModel`

NewLowLevelOperationResponseModel instantiates a new LowLevelOperationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLowLevelOperationResponseModelWithDefaults

`func NewLowLevelOperationResponseModelWithDefaults() *LowLevelOperationResponseModel`

NewLowLevelOperationResponseModelWithDefaults instantiates a new LowLevelOperationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminMachineIP

`func (o *LowLevelOperationResponseModel) GetAdminMachineIP() string`

GetAdminMachineIP returns the AdminMachineIP field if non-nil, zero value otherwise.

### GetAdminMachineIPOk

`func (o *LowLevelOperationResponseModel) GetAdminMachineIPOk() (*string, bool)`

GetAdminMachineIPOk returns a tuple with the AdminMachineIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminMachineIP

`func (o *LowLevelOperationResponseModel) SetAdminMachineIP(v string)`

SetAdminMachineIP sets AdminMachineIP field to given value.

### HasAdminMachineIP

`func (o *LowLevelOperationResponseModel) HasAdminMachineIP() bool`

HasAdminMachineIP returns a boolean if a field has been set.

### SetAdminMachineIPNil

`func (o *LowLevelOperationResponseModel) SetAdminMachineIPNil(b bool)`

 SetAdminMachineIPNil sets the value for AdminMachineIP to be an explicit nil

### UnsetAdminMachineIP
`func (o *LowLevelOperationResponseModel) UnsetAdminMachineIP()`

UnsetAdminMachineIP ensures that no value is present for AdminMachineIP, not even an explicit nil
### GetDetails

`func (o *LowLevelOperationResponseModel) GetDetails() []LowLevelOperationDetailResponseModel`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *LowLevelOperationResponseModel) GetDetailsOk() (*[]LowLevelOperationDetailResponseModel, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *LowLevelOperationResponseModel) SetDetails(v []LowLevelOperationDetailResponseModel)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *LowLevelOperationResponseModel) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *LowLevelOperationResponseModel) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *LowLevelOperationResponseModel) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetEndTime

`func (o *LowLevelOperationResponseModel) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LowLevelOperationResponseModel) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LowLevelOperationResponseModel) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *LowLevelOperationResponseModel) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *LowLevelOperationResponseModel) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *LowLevelOperationResponseModel) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetFormattedEndTime

`func (o *LowLevelOperationResponseModel) GetFormattedEndTime() string`

GetFormattedEndTime returns the FormattedEndTime field if non-nil, zero value otherwise.

### GetFormattedEndTimeOk

`func (o *LowLevelOperationResponseModel) GetFormattedEndTimeOk() (*string, bool)`

GetFormattedEndTimeOk returns a tuple with the FormattedEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedEndTime

`func (o *LowLevelOperationResponseModel) SetFormattedEndTime(v string)`

SetFormattedEndTime sets FormattedEndTime field to given value.

### HasFormattedEndTime

`func (o *LowLevelOperationResponseModel) HasFormattedEndTime() bool`

HasFormattedEndTime returns a boolean if a field has been set.

### SetFormattedEndTimeNil

`func (o *LowLevelOperationResponseModel) SetFormattedEndTimeNil(b bool)`

 SetFormattedEndTimeNil sets the value for FormattedEndTime to be an explicit nil

### UnsetFormattedEndTime
`func (o *LowLevelOperationResponseModel) UnsetFormattedEndTime()`

UnsetFormattedEndTime ensures that no value is present for FormattedEndTime, not even an explicit nil
### GetId

`func (o *LowLevelOperationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LowLevelOperationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LowLevelOperationResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetIsSuccessful

`func (o *LowLevelOperationResponseModel) GetIsSuccessful() bool`

GetIsSuccessful returns the IsSuccessful field if non-nil, zero value otherwise.

### GetIsSuccessfulOk

`func (o *LowLevelOperationResponseModel) GetIsSuccessfulOk() (*bool, bool)`

GetIsSuccessfulOk returns a tuple with the IsSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccessful

`func (o *LowLevelOperationResponseModel) SetIsSuccessful(v bool)`

SetIsSuccessful sets IsSuccessful field to given value.

### HasIsSuccessful

`func (o *LowLevelOperationResponseModel) HasIsSuccessful() bool`

HasIsSuccessful returns a boolean if a field has been set.

### SetIsSuccessfulNil

`func (o *LowLevelOperationResponseModel) SetIsSuccessfulNil(b bool)`

 SetIsSuccessfulNil sets the value for IsSuccessful to be an explicit nil

### UnsetIsSuccessful
`func (o *LowLevelOperationResponseModel) UnsetIsSuccessful()`

UnsetIsSuccessful ensures that no value is present for IsSuccessful, not even an explicit nil
### GetOperationId

`func (o *LowLevelOperationResponseModel) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *LowLevelOperationResponseModel) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *LowLevelOperationResponseModel) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetOperationType

`func (o *LowLevelOperationResponseModel) GetOperationType() LogOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *LowLevelOperationResponseModel) GetOperationTypeOk() (*LogOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *LowLevelOperationResponseModel) SetOperationType(v LogOperationType)`

SetOperationType sets OperationType field to given value.


### GetParameters

`func (o *LowLevelOperationResponseModel) GetParameters() []NameValueStringPairModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *LowLevelOperationResponseModel) GetParametersOk() (*[]NameValueStringPairModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *LowLevelOperationResponseModel) SetParameters(v []NameValueStringPairModel)`

SetParameters sets Parameters field to given value.


### GetSource

`func (o *LowLevelOperationResponseModel) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LowLevelOperationResponseModel) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LowLevelOperationResponseModel) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceSdk

`func (o *LowLevelOperationResponseModel) GetSourceSdk() string`

GetSourceSdk returns the SourceSdk field if non-nil, zero value otherwise.

### GetSourceSdkOk

`func (o *LowLevelOperationResponseModel) GetSourceSdkOk() (*string, bool)`

GetSourceSdkOk returns a tuple with the SourceSdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSdk

`func (o *LowLevelOperationResponseModel) SetSourceSdk(v string)`

SetSourceSdk sets SourceSdk field to given value.

### HasSourceSdk

`func (o *LowLevelOperationResponseModel) HasSourceSdk() bool`

HasSourceSdk returns a boolean if a field has been set.

### SetSourceSdkNil

`func (o *LowLevelOperationResponseModel) SetSourceSdkNil(b bool)`

 SetSourceSdkNil sets the value for SourceSdk to be an explicit nil

### UnsetSourceSdk
`func (o *LowLevelOperationResponseModel) UnsetSourceSdk()`

UnsetSourceSdk ensures that no value is present for SourceSdk, not even an explicit nil
### GetStartTime

`func (o *LowLevelOperationResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LowLevelOperationResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LowLevelOperationResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetFormattedStartTime

`func (o *LowLevelOperationResponseModel) GetFormattedStartTime() string`

GetFormattedStartTime returns the FormattedStartTime field if non-nil, zero value otherwise.

### GetFormattedStartTimeOk

`func (o *LowLevelOperationResponseModel) GetFormattedStartTimeOk() (*string, bool)`

GetFormattedStartTimeOk returns a tuple with the FormattedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedStartTime

`func (o *LowLevelOperationResponseModel) SetFormattedStartTime(v string)`

SetFormattedStartTime sets FormattedStartTime field to given value.


### GetTargetTypes

`func (o *LowLevelOperationResponseModel) GetTargetTypes() []string`

GetTargetTypes returns the TargetTypes field if non-nil, zero value otherwise.

### GetTargetTypesOk

`func (o *LowLevelOperationResponseModel) GetTargetTypesOk() (*[]string, bool)`

GetTargetTypesOk returns a tuple with the TargetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTypes

`func (o *LowLevelOperationResponseModel) SetTargetTypes(v []string)`

SetTargetTypes sets TargetTypes field to given value.


### GetText

`func (o *LowLevelOperationResponseModel) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *LowLevelOperationResponseModel) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *LowLevelOperationResponseModel) SetText(v string)`

SetText sets Text field to given value.


### GetUser

`func (o *LowLevelOperationResponseModel) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LowLevelOperationResponseModel) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LowLevelOperationResponseModel) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LowLevelOperationResponseModel) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *LowLevelOperationResponseModel) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *LowLevelOperationResponseModel) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


