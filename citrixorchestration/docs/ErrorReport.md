# ErrorReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CitrixErrorId** | **string** | Gets or sets CitrixErrorId | 
**ComponentName** | **string** | Gets or sets  the name of the component that threw the error | 
**ComponentVersion** | **string** | Gets or sets the version of the component that threw the error              | 
**ErrorData** | **map[string]string** | Gets Key/value pairs of error information to be added to the error report | 

## Methods

### NewErrorReport

`func NewErrorReport(citrixErrorId string, componentName string, componentVersion string, errorData map[string]string, ) *ErrorReport`

NewErrorReport instantiates a new ErrorReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorReportWithDefaults

`func NewErrorReportWithDefaults() *ErrorReport`

NewErrorReportWithDefaults instantiates a new ErrorReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitrixErrorId

`func (o *ErrorReport) GetCitrixErrorId() string`

GetCitrixErrorId returns the CitrixErrorId field if non-nil, zero value otherwise.

### GetCitrixErrorIdOk

`func (o *ErrorReport) GetCitrixErrorIdOk() (*string, bool)`

GetCitrixErrorIdOk returns a tuple with the CitrixErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixErrorId

`func (o *ErrorReport) SetCitrixErrorId(v string)`

SetCitrixErrorId sets CitrixErrorId field to given value.


### GetComponentName

`func (o *ErrorReport) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *ErrorReport) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *ErrorReport) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.


### GetComponentVersion

`func (o *ErrorReport) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *ErrorReport) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *ErrorReport) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.


### GetErrorData

`func (o *ErrorReport) GetErrorData() map[string]string`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *ErrorReport) GetErrorDataOk() (*map[string]string, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *ErrorReport) SetErrorData(v map[string]string)`

SetErrorData sets ErrorData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


