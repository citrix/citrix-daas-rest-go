# AzureADJoinStatusResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureADJoinedMode** | Pointer to [**AzureAdJoinedMode**](AzureAdJoinedMode.md) |  | [optional] 
**DomainJoinStatus** | Pointer to [**DomainJoinStatus**](DomainJoinStatus.md) |  | [optional] 

## Methods

### NewAzureADJoinStatusResponseModel

`func NewAzureADJoinStatusResponseModel() *AzureADJoinStatusResponseModel`

NewAzureADJoinStatusResponseModel instantiates a new AzureADJoinStatusResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureADJoinStatusResponseModelWithDefaults

`func NewAzureADJoinStatusResponseModelWithDefaults() *AzureADJoinStatusResponseModel`

NewAzureADJoinStatusResponseModelWithDefaults instantiates a new AzureADJoinStatusResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureADJoinedMode

`func (o *AzureADJoinStatusResponseModel) GetAzureADJoinedMode() AzureAdJoinedMode`

GetAzureADJoinedMode returns the AzureADJoinedMode field if non-nil, zero value otherwise.

### GetAzureADJoinedModeOk

`func (o *AzureADJoinStatusResponseModel) GetAzureADJoinedModeOk() (*AzureAdJoinedMode, bool)`

GetAzureADJoinedModeOk returns a tuple with the AzureADJoinedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADJoinedMode

`func (o *AzureADJoinStatusResponseModel) SetAzureADJoinedMode(v AzureAdJoinedMode)`

SetAzureADJoinedMode sets AzureADJoinedMode field to given value.

### HasAzureADJoinedMode

`func (o *AzureADJoinStatusResponseModel) HasAzureADJoinedMode() bool`

HasAzureADJoinedMode returns a boolean if a field has been set.

### GetDomainJoinStatus

`func (o *AzureADJoinStatusResponseModel) GetDomainJoinStatus() DomainJoinStatus`

GetDomainJoinStatus returns the DomainJoinStatus field if non-nil, zero value otherwise.

### GetDomainJoinStatusOk

`func (o *AzureADJoinStatusResponseModel) GetDomainJoinStatusOk() (*DomainJoinStatus, bool)`

GetDomainJoinStatusOk returns a tuple with the DomainJoinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainJoinStatus

`func (o *AzureADJoinStatusResponseModel) SetDomainJoinStatus(v DomainJoinStatus)`

SetDomainJoinStatus sets DomainJoinStatus field to given value.

### HasDomainJoinStatus

`func (o *AzureADJoinStatusResponseModel) HasDomainJoinStatus() bool`

HasDomainJoinStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


