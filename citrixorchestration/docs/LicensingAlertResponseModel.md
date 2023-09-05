# LicensingAlertResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | The priority of the licensing alert. | [optional] 
**AlertLevel** | Pointer to [**LicensingAlertLevel2**](LicensingAlertLevel2.md) |  | [optional] 
**TypeCode** | Pointer to **string** | The type code of the licensing alert (internal code used by licensing). | [optional] 
**Title** | Pointer to **string** | The title of the alert. | [optional] 
**Detail** | Pointer to **string** | More detailed information about the licensing alert. | [optional] 
**Action** | Pointer to **string** | The action the user should take to resolve the alert. | [optional] 

## Methods

### NewLicensingAlertResponseModel

`func NewLicensingAlertResponseModel() *LicensingAlertResponseModel`

NewLicensingAlertResponseModel instantiates a new LicensingAlertResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingAlertResponseModelWithDefaults

`func NewLicensingAlertResponseModelWithDefaults() *LicensingAlertResponseModel`

NewLicensingAlertResponseModelWithDefaults instantiates a new LicensingAlertResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *LicensingAlertResponseModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *LicensingAlertResponseModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *LicensingAlertResponseModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *LicensingAlertResponseModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAlertLevel

`func (o *LicensingAlertResponseModel) GetAlertLevel() LicensingAlertLevel2`

GetAlertLevel returns the AlertLevel field if non-nil, zero value otherwise.

### GetAlertLevelOk

`func (o *LicensingAlertResponseModel) GetAlertLevelOk() (*LicensingAlertLevel2, bool)`

GetAlertLevelOk returns a tuple with the AlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLevel

`func (o *LicensingAlertResponseModel) SetAlertLevel(v LicensingAlertLevel2)`

SetAlertLevel sets AlertLevel field to given value.

### HasAlertLevel

`func (o *LicensingAlertResponseModel) HasAlertLevel() bool`

HasAlertLevel returns a boolean if a field has been set.

### GetTypeCode

`func (o *LicensingAlertResponseModel) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *LicensingAlertResponseModel) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *LicensingAlertResponseModel) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *LicensingAlertResponseModel) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetTitle

`func (o *LicensingAlertResponseModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LicensingAlertResponseModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LicensingAlertResponseModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LicensingAlertResponseModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *LicensingAlertResponseModel) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *LicensingAlertResponseModel) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *LicensingAlertResponseModel) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *LicensingAlertResponseModel) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetAction

`func (o *LicensingAlertResponseModel) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LicensingAlertResponseModel) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LicensingAlertResponseModel) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *LicensingAlertResponseModel) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


