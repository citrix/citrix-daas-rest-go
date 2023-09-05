# HypervisorAzureEnvironmentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentDisplayName** | Pointer to **string** | Environment display name which used to show on the front end. e.g. Azure Global, Azure China and Azure US Government. | [optional] 
**EnvironmentName** | Pointer to [**AzureEnvironment**](AzureEnvironment.md) |  | [optional] 

## Methods

### NewHypervisorAzureEnvironmentResponseModel

`func NewHypervisorAzureEnvironmentResponseModel() *HypervisorAzureEnvironmentResponseModel`

NewHypervisorAzureEnvironmentResponseModel instantiates a new HypervisorAzureEnvironmentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorAzureEnvironmentResponseModelWithDefaults

`func NewHypervisorAzureEnvironmentResponseModelWithDefaults() *HypervisorAzureEnvironmentResponseModel`

NewHypervisorAzureEnvironmentResponseModelWithDefaults instantiates a new HypervisorAzureEnvironmentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentDisplayName

`func (o *HypervisorAzureEnvironmentResponseModel) GetEnvironmentDisplayName() string`

GetEnvironmentDisplayName returns the EnvironmentDisplayName field if non-nil, zero value otherwise.

### GetEnvironmentDisplayNameOk

`func (o *HypervisorAzureEnvironmentResponseModel) GetEnvironmentDisplayNameOk() (*string, bool)`

GetEnvironmentDisplayNameOk returns a tuple with the EnvironmentDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentDisplayName

`func (o *HypervisorAzureEnvironmentResponseModel) SetEnvironmentDisplayName(v string)`

SetEnvironmentDisplayName sets EnvironmentDisplayName field to given value.

### HasEnvironmentDisplayName

`func (o *HypervisorAzureEnvironmentResponseModel) HasEnvironmentDisplayName() bool`

HasEnvironmentDisplayName returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *HypervisorAzureEnvironmentResponseModel) GetEnvironmentName() AzureEnvironment`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *HypervisorAzureEnvironmentResponseModel) GetEnvironmentNameOk() (*AzureEnvironment, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *HypervisorAzureEnvironmentResponseModel) SetEnvironmentName(v AzureEnvironment)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *HypervisorAzureEnvironmentResponseModel) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


