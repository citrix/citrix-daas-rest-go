# AboutModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to **string** | The commit hash when building.              | [optional] 
**IsCloud** | Pointer to **bool** | Indicates if the  current environment is cloud. | [optional] 

## Methods

### NewAboutModel

`func NewAboutModel() *AboutModel`

NewAboutModel instantiates a new AboutModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAboutModelWithDefaults

`func NewAboutModelWithDefaults() *AboutModel`

NewAboutModelWithDefaults instantiates a new AboutModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *AboutModel) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *AboutModel) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *AboutModel) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *AboutModel) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetIsCloud

`func (o *AboutModel) GetIsCloud() bool`

GetIsCloud returns the IsCloud field if non-nil, zero value otherwise.

### GetIsCloudOk

`func (o *AboutModel) GetIsCloudOk() (*bool, bool)`

GetIsCloudOk returns a tuple with the IsCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloud

`func (o *AboutModel) SetIsCloud(v bool)`

SetIsCloud sets IsCloud field to given value.

### HasIsCloud

`func (o *AboutModel) HasIsCloud() bool`

HasIsCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


