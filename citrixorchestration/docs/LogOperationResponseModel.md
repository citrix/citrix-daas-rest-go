# LogOperationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the logged operation. | 
**AdminMachineIP** | Pointer to **NullableString** | IP address of the admin machine from which the operation was performed. | [optional] 
**EndTime** | Pointer to **NullableString** | Time when the operation ended. If the operation is incomplete, will be null. | [optional] 
**FormattedEndTime** | Pointer to **NullableString** | Formatted time when the operation ended. If the operation is incomplete, will be null. RFC 3339 compatible format. | [optional] 
**IsSuccessful** | Pointer to **NullableBool** | Indicates whether the operation completed successfully.  If the operation is incomplete, will be null. | [optional] 
**OperationType** | [**LogOperationType**](LogOperationType.md) |  | 
**Parameters** | [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Operation parameters. It is a list of parameter names. | 
**Source** | **string** | Source of the operation. | 
**StartTime** | **string** | Time when the operation started. | 
**FormattedStartTime** | **string** | Formatted time when the operation started. RFC 3339 compatible format. | 
**TargetTypes** | **[]string** | The type(s) of object which were the target of the configuration change. For example, \&quot;Session\&quot; or \&quot;Machine\&quot;. | 
**Text** | **string** | Human-readable description of the change. | 
**User** | Pointer to **NullableString** | User who performed the change. | [optional] 

## Methods

### NewLogOperationResponseModel

`func NewLogOperationResponseModel(id string, operationType LogOperationType, parameters []NameValueStringPairModel, source string, startTime string, formattedStartTime string, targetTypes []string, text string, ) *LogOperationResponseModel`

NewLogOperationResponseModel instantiates a new LogOperationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogOperationResponseModelWithDefaults

`func NewLogOperationResponseModelWithDefaults() *LogOperationResponseModel`

NewLogOperationResponseModelWithDefaults instantiates a new LogOperationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogOperationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogOperationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogOperationResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetAdminMachineIP

`func (o *LogOperationResponseModel) GetAdminMachineIP() string`

GetAdminMachineIP returns the AdminMachineIP field if non-nil, zero value otherwise.

### GetAdminMachineIPOk

`func (o *LogOperationResponseModel) GetAdminMachineIPOk() (*string, bool)`

GetAdminMachineIPOk returns a tuple with the AdminMachineIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminMachineIP

`func (o *LogOperationResponseModel) SetAdminMachineIP(v string)`

SetAdminMachineIP sets AdminMachineIP field to given value.

### HasAdminMachineIP

`func (o *LogOperationResponseModel) HasAdminMachineIP() bool`

HasAdminMachineIP returns a boolean if a field has been set.

### SetAdminMachineIPNil

`func (o *LogOperationResponseModel) SetAdminMachineIPNil(b bool)`

 SetAdminMachineIPNil sets the value for AdminMachineIP to be an explicit nil

### UnsetAdminMachineIP
`func (o *LogOperationResponseModel) UnsetAdminMachineIP()`

UnsetAdminMachineIP ensures that no value is present for AdminMachineIP, not even an explicit nil
### GetEndTime

`func (o *LogOperationResponseModel) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LogOperationResponseModel) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LogOperationResponseModel) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *LogOperationResponseModel) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *LogOperationResponseModel) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *LogOperationResponseModel) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetFormattedEndTime

`func (o *LogOperationResponseModel) GetFormattedEndTime() string`

GetFormattedEndTime returns the FormattedEndTime field if non-nil, zero value otherwise.

### GetFormattedEndTimeOk

`func (o *LogOperationResponseModel) GetFormattedEndTimeOk() (*string, bool)`

GetFormattedEndTimeOk returns a tuple with the FormattedEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedEndTime

`func (o *LogOperationResponseModel) SetFormattedEndTime(v string)`

SetFormattedEndTime sets FormattedEndTime field to given value.

### HasFormattedEndTime

`func (o *LogOperationResponseModel) HasFormattedEndTime() bool`

HasFormattedEndTime returns a boolean if a field has been set.

### SetFormattedEndTimeNil

`func (o *LogOperationResponseModel) SetFormattedEndTimeNil(b bool)`

 SetFormattedEndTimeNil sets the value for FormattedEndTime to be an explicit nil

### UnsetFormattedEndTime
`func (o *LogOperationResponseModel) UnsetFormattedEndTime()`

UnsetFormattedEndTime ensures that no value is present for FormattedEndTime, not even an explicit nil
### GetIsSuccessful

`func (o *LogOperationResponseModel) GetIsSuccessful() bool`

GetIsSuccessful returns the IsSuccessful field if non-nil, zero value otherwise.

### GetIsSuccessfulOk

`func (o *LogOperationResponseModel) GetIsSuccessfulOk() (*bool, bool)`

GetIsSuccessfulOk returns a tuple with the IsSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccessful

`func (o *LogOperationResponseModel) SetIsSuccessful(v bool)`

SetIsSuccessful sets IsSuccessful field to given value.

### HasIsSuccessful

`func (o *LogOperationResponseModel) HasIsSuccessful() bool`

HasIsSuccessful returns a boolean if a field has been set.

### SetIsSuccessfulNil

`func (o *LogOperationResponseModel) SetIsSuccessfulNil(b bool)`

 SetIsSuccessfulNil sets the value for IsSuccessful to be an explicit nil

### UnsetIsSuccessful
`func (o *LogOperationResponseModel) UnsetIsSuccessful()`

UnsetIsSuccessful ensures that no value is present for IsSuccessful, not even an explicit nil
### GetOperationType

`func (o *LogOperationResponseModel) GetOperationType() LogOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *LogOperationResponseModel) GetOperationTypeOk() (*LogOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *LogOperationResponseModel) SetOperationType(v LogOperationType)`

SetOperationType sets OperationType field to given value.


### GetParameters

`func (o *LogOperationResponseModel) GetParameters() []NameValueStringPairModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *LogOperationResponseModel) GetParametersOk() (*[]NameValueStringPairModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *LogOperationResponseModel) SetParameters(v []NameValueStringPairModel)`

SetParameters sets Parameters field to given value.


### GetSource

`func (o *LogOperationResponseModel) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LogOperationResponseModel) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LogOperationResponseModel) SetSource(v string)`

SetSource sets Source field to given value.


### GetStartTime

`func (o *LogOperationResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LogOperationResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LogOperationResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetFormattedStartTime

`func (o *LogOperationResponseModel) GetFormattedStartTime() string`

GetFormattedStartTime returns the FormattedStartTime field if non-nil, zero value otherwise.

### GetFormattedStartTimeOk

`func (o *LogOperationResponseModel) GetFormattedStartTimeOk() (*string, bool)`

GetFormattedStartTimeOk returns a tuple with the FormattedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedStartTime

`func (o *LogOperationResponseModel) SetFormattedStartTime(v string)`

SetFormattedStartTime sets FormattedStartTime field to given value.


### GetTargetTypes

`func (o *LogOperationResponseModel) GetTargetTypes() []string`

GetTargetTypes returns the TargetTypes field if non-nil, zero value otherwise.

### GetTargetTypesOk

`func (o *LogOperationResponseModel) GetTargetTypesOk() (*[]string, bool)`

GetTargetTypesOk returns a tuple with the TargetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTypes

`func (o *LogOperationResponseModel) SetTargetTypes(v []string)`

SetTargetTypes sets TargetTypes field to given value.


### GetText

`func (o *LogOperationResponseModel) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *LogOperationResponseModel) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *LogOperationResponseModel) SetText(v string)`

SetText sets Text field to given value.


### GetUser

`func (o *LogOperationResponseModel) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LogOperationResponseModel) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LogOperationResponseModel) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LogOperationResponseModel) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *LogOperationResponseModel) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *LogOperationResponseModel) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


