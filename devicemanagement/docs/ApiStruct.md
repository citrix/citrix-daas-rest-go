# ApiStruct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Function** | Pointer to **string** |  | [optional] 
**ControllerName** | Pointer to **string** |  | [optional] 
**AuthLevel** | Pointer to **string** |  | [optional] 

## Methods

### NewApiStruct

`func NewApiStruct() *ApiStruct`

NewApiStruct instantiates a new ApiStruct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStructWithDefaults

`func NewApiStructWithDefaults() *ApiStruct`

NewApiStructWithDefaults instantiates a new ApiStruct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ApiStruct) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiStruct) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiStruct) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApiStruct) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMethod

`func (o *ApiStruct) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiStruct) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiStruct) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ApiStruct) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetFunction

`func (o *ApiStruct) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ApiStruct) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ApiStruct) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *ApiStruct) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetControllerName

`func (o *ApiStruct) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *ApiStruct) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *ApiStruct) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *ApiStruct) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetAuthLevel

`func (o *ApiStruct) GetAuthLevel() string`

GetAuthLevel returns the AuthLevel field if non-nil, zero value otherwise.

### GetAuthLevelOk

`func (o *ApiStruct) GetAuthLevelOk() (*string, bool)`

GetAuthLevelOk returns a tuple with the AuthLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthLevel

`func (o *ApiStruct) SetAuthLevel(v string)`

SetAuthLevel sets AuthLevel field to given value.

### HasAuthLevel

`func (o *ApiStruct) HasAuthLevel() bool`

HasAuthLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


