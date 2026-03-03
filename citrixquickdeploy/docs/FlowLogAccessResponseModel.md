# FlowLogAccessResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SasUrl** | Pointer to **NullableString** | SAS URL for accessing the flow log container | [optional] 
**ContainerName** | Pointer to **NullableString** | Name of the container containing flow logs | [optional] 
**StorageAccountName** | Pointer to **NullableString** | Storage account name for reference | [optional] 
**ExpiresAt** | Pointer to **time.Time** | When the SAS URL expires | [optional] 
**ValidForMinutes** | Pointer to **int32** | How many minutes the URL is valid for (remaining time until expiry) | [optional] 
**IsNewlyGenerated** | Pointer to **bool** | Whether this URL was newly generated (true) or returned from cache (false) | [optional] 

## Methods

### NewFlowLogAccessResponseModel

`func NewFlowLogAccessResponseModel() *FlowLogAccessResponseModel`

NewFlowLogAccessResponseModel instantiates a new FlowLogAccessResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowLogAccessResponseModelWithDefaults

`func NewFlowLogAccessResponseModelWithDefaults() *FlowLogAccessResponseModel`

NewFlowLogAccessResponseModelWithDefaults instantiates a new FlowLogAccessResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSasUrl

`func (o *FlowLogAccessResponseModel) GetSasUrl() string`

GetSasUrl returns the SasUrl field if non-nil, zero value otherwise.

### GetSasUrlOk

`func (o *FlowLogAccessResponseModel) GetSasUrlOk() (*string, bool)`

GetSasUrlOk returns a tuple with the SasUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasUrl

`func (o *FlowLogAccessResponseModel) SetSasUrl(v string)`

SetSasUrl sets SasUrl field to given value.

### HasSasUrl

`func (o *FlowLogAccessResponseModel) HasSasUrl() bool`

HasSasUrl returns a boolean if a field has been set.

### SetSasUrlNil

`func (o *FlowLogAccessResponseModel) SetSasUrlNil(b bool)`

 SetSasUrlNil sets the value for SasUrl to be an explicit nil

### UnsetSasUrl
`func (o *FlowLogAccessResponseModel) UnsetSasUrl()`

UnsetSasUrl ensures that no value is present for SasUrl, not even an explicit nil
### GetContainerName

`func (o *FlowLogAccessResponseModel) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *FlowLogAccessResponseModel) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *FlowLogAccessResponseModel) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *FlowLogAccessResponseModel) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### SetContainerNameNil

`func (o *FlowLogAccessResponseModel) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *FlowLogAccessResponseModel) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetStorageAccountName

`func (o *FlowLogAccessResponseModel) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *FlowLogAccessResponseModel) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *FlowLogAccessResponseModel) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.

### HasStorageAccountName

`func (o *FlowLogAccessResponseModel) HasStorageAccountName() bool`

HasStorageAccountName returns a boolean if a field has been set.

### SetStorageAccountNameNil

`func (o *FlowLogAccessResponseModel) SetStorageAccountNameNil(b bool)`

 SetStorageAccountNameNil sets the value for StorageAccountName to be an explicit nil

### UnsetStorageAccountName
`func (o *FlowLogAccessResponseModel) UnsetStorageAccountName()`

UnsetStorageAccountName ensures that no value is present for StorageAccountName, not even an explicit nil
### GetExpiresAt

`func (o *FlowLogAccessResponseModel) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FlowLogAccessResponseModel) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FlowLogAccessResponseModel) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FlowLogAccessResponseModel) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetValidForMinutes

`func (o *FlowLogAccessResponseModel) GetValidForMinutes() int32`

GetValidForMinutes returns the ValidForMinutes field if non-nil, zero value otherwise.

### GetValidForMinutesOk

`func (o *FlowLogAccessResponseModel) GetValidForMinutesOk() (*int32, bool)`

GetValidForMinutesOk returns a tuple with the ValidForMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidForMinutes

`func (o *FlowLogAccessResponseModel) SetValidForMinutes(v int32)`

SetValidForMinutes sets ValidForMinutes field to given value.

### HasValidForMinutes

`func (o *FlowLogAccessResponseModel) HasValidForMinutes() bool`

HasValidForMinutes returns a boolean if a field has been set.

### GetIsNewlyGenerated

`func (o *FlowLogAccessResponseModel) GetIsNewlyGenerated() bool`

GetIsNewlyGenerated returns the IsNewlyGenerated field if non-nil, zero value otherwise.

### GetIsNewlyGeneratedOk

`func (o *FlowLogAccessResponseModel) GetIsNewlyGeneratedOk() (*bool, bool)`

GetIsNewlyGeneratedOk returns a tuple with the IsNewlyGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewlyGenerated

`func (o *FlowLogAccessResponseModel) SetIsNewlyGenerated(v bool)`

SetIsNewlyGenerated sets IsNewlyGenerated field to given value.

### HasIsNewlyGenerated

`func (o *FlowLogAccessResponseModel) HasIsNewlyGenerated() bool`

HasIsNewlyGenerated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


