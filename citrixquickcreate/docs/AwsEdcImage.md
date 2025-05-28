# AwsEdcImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **NullableString** | Provides additional status on image like  error message | [optional] 
**AmazonImageId** | Pointer to **NullableString** | Id of the Image on AWS | [optional] 
**IngestionProcess** | Pointer to [**NullableAwsEdcWorkspaceImageIngestionProcess**](AwsEdcWorkspaceImageIngestionProcess.md) | Ingestion Process used during image import | [optional] 
**WorkspaceImageTenancy** | Pointer to [**NullableAwsEdcWorkspaceImageTenancy**](AwsEdcWorkspaceImageTenancy.md) | Tenancy of the image  Enum values DEDICATED, DEFAULT | [optional] 
**WorkspaceImageState** | Pointer to [**NullableAwsEdcWorkspaceImageState**](AwsEdcWorkspaceImageState.md) | State of the image  Enum values AVAILABLE, ERROR, PENDING | [optional] 
**ApplicationList** | Pointer to [**[]AwsEdcAmiImportApplications**](AwsEdcAmiImportApplications.md) | The list of installed applications | [optional] 

## Methods

### NewAwsEdcImage

`func NewAwsEdcImage() *AwsEdcImage`

NewAwsEdcImage instantiates a new AwsEdcImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcImageWithDefaults

`func NewAwsEdcImageWithDefaults() *AwsEdcImage`

NewAwsEdcImageWithDefaults instantiates a new AwsEdcImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsEdcImage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsEdcImage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsEdcImage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsEdcImage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AwsEdcImage) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AwsEdcImage) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAmazonImageId

`func (o *AwsEdcImage) GetAmazonImageId() string`

GetAmazonImageId returns the AmazonImageId field if non-nil, zero value otherwise.

### GetAmazonImageIdOk

`func (o *AwsEdcImage) GetAmazonImageIdOk() (*string, bool)`

GetAmazonImageIdOk returns a tuple with the AmazonImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonImageId

`func (o *AwsEdcImage) SetAmazonImageId(v string)`

SetAmazonImageId sets AmazonImageId field to given value.

### HasAmazonImageId

`func (o *AwsEdcImage) HasAmazonImageId() bool`

HasAmazonImageId returns a boolean if a field has been set.

### SetAmazonImageIdNil

`func (o *AwsEdcImage) SetAmazonImageIdNil(b bool)`

 SetAmazonImageIdNil sets the value for AmazonImageId to be an explicit nil

### UnsetAmazonImageId
`func (o *AwsEdcImage) UnsetAmazonImageId()`

UnsetAmazonImageId ensures that no value is present for AmazonImageId, not even an explicit nil
### GetIngestionProcess

`func (o *AwsEdcImage) GetIngestionProcess() AwsEdcWorkspaceImageIngestionProcess`

GetIngestionProcess returns the IngestionProcess field if non-nil, zero value otherwise.

### GetIngestionProcessOk

`func (o *AwsEdcImage) GetIngestionProcessOk() (*AwsEdcWorkspaceImageIngestionProcess, bool)`

GetIngestionProcessOk returns a tuple with the IngestionProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionProcess

`func (o *AwsEdcImage) SetIngestionProcess(v AwsEdcWorkspaceImageIngestionProcess)`

SetIngestionProcess sets IngestionProcess field to given value.

### HasIngestionProcess

`func (o *AwsEdcImage) HasIngestionProcess() bool`

HasIngestionProcess returns a boolean if a field has been set.

### SetIngestionProcessNil

`func (o *AwsEdcImage) SetIngestionProcessNil(b bool)`

 SetIngestionProcessNil sets the value for IngestionProcess to be an explicit nil

### UnsetIngestionProcess
`func (o *AwsEdcImage) UnsetIngestionProcess()`

UnsetIngestionProcess ensures that no value is present for IngestionProcess, not even an explicit nil
### GetWorkspaceImageTenancy

`func (o *AwsEdcImage) GetWorkspaceImageTenancy() AwsEdcWorkspaceImageTenancy`

GetWorkspaceImageTenancy returns the WorkspaceImageTenancy field if non-nil, zero value otherwise.

### GetWorkspaceImageTenancyOk

`func (o *AwsEdcImage) GetWorkspaceImageTenancyOk() (*AwsEdcWorkspaceImageTenancy, bool)`

GetWorkspaceImageTenancyOk returns a tuple with the WorkspaceImageTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceImageTenancy

`func (o *AwsEdcImage) SetWorkspaceImageTenancy(v AwsEdcWorkspaceImageTenancy)`

SetWorkspaceImageTenancy sets WorkspaceImageTenancy field to given value.

### HasWorkspaceImageTenancy

`func (o *AwsEdcImage) HasWorkspaceImageTenancy() bool`

HasWorkspaceImageTenancy returns a boolean if a field has been set.

### SetWorkspaceImageTenancyNil

`func (o *AwsEdcImage) SetWorkspaceImageTenancyNil(b bool)`

 SetWorkspaceImageTenancyNil sets the value for WorkspaceImageTenancy to be an explicit nil

### UnsetWorkspaceImageTenancy
`func (o *AwsEdcImage) UnsetWorkspaceImageTenancy()`

UnsetWorkspaceImageTenancy ensures that no value is present for WorkspaceImageTenancy, not even an explicit nil
### GetWorkspaceImageState

`func (o *AwsEdcImage) GetWorkspaceImageState() AwsEdcWorkspaceImageState`

GetWorkspaceImageState returns the WorkspaceImageState field if non-nil, zero value otherwise.

### GetWorkspaceImageStateOk

`func (o *AwsEdcImage) GetWorkspaceImageStateOk() (*AwsEdcWorkspaceImageState, bool)`

GetWorkspaceImageStateOk returns a tuple with the WorkspaceImageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceImageState

`func (o *AwsEdcImage) SetWorkspaceImageState(v AwsEdcWorkspaceImageState)`

SetWorkspaceImageState sets WorkspaceImageState field to given value.

### HasWorkspaceImageState

`func (o *AwsEdcImage) HasWorkspaceImageState() bool`

HasWorkspaceImageState returns a boolean if a field has been set.

### SetWorkspaceImageStateNil

`func (o *AwsEdcImage) SetWorkspaceImageStateNil(b bool)`

 SetWorkspaceImageStateNil sets the value for WorkspaceImageState to be an explicit nil

### UnsetWorkspaceImageState
`func (o *AwsEdcImage) UnsetWorkspaceImageState()`

UnsetWorkspaceImageState ensures that no value is present for WorkspaceImageState, not even an explicit nil
### GetApplicationList

`func (o *AwsEdcImage) GetApplicationList() []AwsEdcAmiImportApplications`

GetApplicationList returns the ApplicationList field if non-nil, zero value otherwise.

### GetApplicationListOk

`func (o *AwsEdcImage) GetApplicationListOk() (*[]AwsEdcAmiImportApplications, bool)`

GetApplicationListOk returns a tuple with the ApplicationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationList

`func (o *AwsEdcImage) SetApplicationList(v []AwsEdcAmiImportApplications)`

SetApplicationList sets ApplicationList field to given value.

### HasApplicationList

`func (o *AwsEdcImage) HasApplicationList() bool`

HasApplicationList returns a boolean if a field has been set.

### SetApplicationListNil

`func (o *AwsEdcImage) SetApplicationListNil(b bool)`

 SetApplicationListNil sets the value for ApplicationList to be an explicit nil

### UnsetApplicationList
`func (o *AwsEdcImage) UnsetApplicationList()`

UnsetApplicationList ensures that no value is present for ApplicationList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


