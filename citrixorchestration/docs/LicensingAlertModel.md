# LicensingAlertModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | The priority of the licensing alert. | [optional] 
**LicensingAlertLevel** | Pointer to [**LicensingAlertLevel**](LicensingAlertLevel.md) |  | [optional] 
**TypeCode** | Pointer to **string** | The type code of the licensing alert (internal code used by licensing). | [optional] 
**Title** | Pointer to **string** | The title of the alert. | [optional] 
**Detail** | Pointer to **string** | More detailed information about the licensing alert. | [optional] 
**Action** | Pointer to **string** | The action the user should take to resolve the alert. | [optional] 

## Methods

### NewLicensingAlertModel

`func NewLicensingAlertModel() *LicensingAlertModel`

NewLicensingAlertModel instantiates a new LicensingAlertModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingAlertModelWithDefaults

`func NewLicensingAlertModelWithDefaults() *LicensingAlertModel`

NewLicensingAlertModelWithDefaults instantiates a new LicensingAlertModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *LicensingAlertModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *LicensingAlertModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *LicensingAlertModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *LicensingAlertModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetLicensingAlertLevel

`func (o *LicensingAlertModel) GetLicensingAlertLevel() LicensingAlertLevel`

GetLicensingAlertLevel returns the LicensingAlertLevel field if non-nil, zero value otherwise.

### GetLicensingAlertLevelOk

`func (o *LicensingAlertModel) GetLicensingAlertLevelOk() (*LicensingAlertLevel, bool)`

GetLicensingAlertLevelOk returns a tuple with the LicensingAlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingAlertLevel

`func (o *LicensingAlertModel) SetLicensingAlertLevel(v LicensingAlertLevel)`

SetLicensingAlertLevel sets LicensingAlertLevel field to given value.

### HasLicensingAlertLevel

`func (o *LicensingAlertModel) HasLicensingAlertLevel() bool`

HasLicensingAlertLevel returns a boolean if a field has been set.

### GetTypeCode

`func (o *LicensingAlertModel) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *LicensingAlertModel) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *LicensingAlertModel) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *LicensingAlertModel) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetTitle

`func (o *LicensingAlertModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LicensingAlertModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LicensingAlertModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LicensingAlertModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *LicensingAlertModel) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *LicensingAlertModel) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *LicensingAlertModel) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *LicensingAlertModel) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetAction

`func (o *LicensingAlertModel) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LicensingAlertModel) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LicensingAlertModel) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *LicensingAlertModel) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


