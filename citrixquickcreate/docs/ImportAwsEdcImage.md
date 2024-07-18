# ImportAwsEdcImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ec2ImageId** | **string** | Source Image Id | 
**IngestionProcess** | Pointer to [**AwsEdcWorkspaceImageIngestionProcess**](AwsEdcWorkspaceImageIngestionProcess.md) |  | [optional] 
**ApplicationList** | Pointer to [**[]AwsEdcAmiImportApplications**](AwsEdcAmiImportApplications.md) | The list of installed applications | [optional] 

## Methods

### NewImportAwsEdcImage

`func NewImportAwsEdcImage(ec2ImageId string, ) *ImportAwsEdcImage`

NewImportAwsEdcImage instantiates a new ImportAwsEdcImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportAwsEdcImageWithDefaults

`func NewImportAwsEdcImageWithDefaults() *ImportAwsEdcImage`

NewImportAwsEdcImageWithDefaults instantiates a new ImportAwsEdcImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEc2ImageId

`func (o *ImportAwsEdcImage) GetEc2ImageId() string`

GetEc2ImageId returns the Ec2ImageId field if non-nil, zero value otherwise.

### GetEc2ImageIdOk

`func (o *ImportAwsEdcImage) GetEc2ImageIdOk() (*string, bool)`

GetEc2ImageIdOk returns a tuple with the Ec2ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEc2ImageId

`func (o *ImportAwsEdcImage) SetEc2ImageId(v string)`

SetEc2ImageId sets Ec2ImageId field to given value.


### GetIngestionProcess

`func (o *ImportAwsEdcImage) GetIngestionProcess() AwsEdcWorkspaceImageIngestionProcess`

GetIngestionProcess returns the IngestionProcess field if non-nil, zero value otherwise.

### GetIngestionProcessOk

`func (o *ImportAwsEdcImage) GetIngestionProcessOk() (*AwsEdcWorkspaceImageIngestionProcess, bool)`

GetIngestionProcessOk returns a tuple with the IngestionProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionProcess

`func (o *ImportAwsEdcImage) SetIngestionProcess(v AwsEdcWorkspaceImageIngestionProcess)`

SetIngestionProcess sets IngestionProcess field to given value.

### HasIngestionProcess

`func (o *ImportAwsEdcImage) HasIngestionProcess() bool`

HasIngestionProcess returns a boolean if a field has been set.

### GetApplicationList

`func (o *ImportAwsEdcImage) GetApplicationList() []AwsEdcAmiImportApplications`

GetApplicationList returns the ApplicationList field if non-nil, zero value otherwise.

### GetApplicationListOk

`func (o *ImportAwsEdcImage) GetApplicationListOk() (*[]AwsEdcAmiImportApplications, bool)`

GetApplicationListOk returns a tuple with the ApplicationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationList

`func (o *ImportAwsEdcImage) SetApplicationList(v []AwsEdcAmiImportApplications)`

SetApplicationList sets ApplicationList field to given value.

### HasApplicationList

`func (o *ImportAwsEdcImage) HasApplicationList() bool`

HasApplicationList returns a boolean if a field has been set.

### SetApplicationListNil

`func (o *ImportAwsEdcImage) SetApplicationListNil(b bool)`

 SetApplicationListNil sets the value for ApplicationList to be an explicit nil

### UnsetApplicationList
`func (o *ImportAwsEdcImage) UnsetApplicationList()`

UnsetApplicationList ensures that no value is present for ApplicationList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


