# BackupVmConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceOffering** | Pointer to **string** | The service offering for the backup VM.  Ex: \&quot;Standard_D2as_v4\&quot; | [optional] 
**Type** | Pointer to **string** | The service offering type for the backup VM.  Options are: \&quot;Regular\&quot; or \&quot;Spot\&quot;.  Default to \&quot;Regular\&quot; if unspecified. | [optional] 

## Methods

### NewBackupVmConfiguration

`func NewBackupVmConfiguration() *BackupVmConfiguration`

NewBackupVmConfiguration instantiates a new BackupVmConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupVmConfigurationWithDefaults

`func NewBackupVmConfigurationWithDefaults() *BackupVmConfiguration`

NewBackupVmConfigurationWithDefaults instantiates a new BackupVmConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceOffering

`func (o *BackupVmConfiguration) GetServiceOffering() string`

GetServiceOffering returns the ServiceOffering field if non-nil, zero value otherwise.

### GetServiceOfferingOk

`func (o *BackupVmConfiguration) GetServiceOfferingOk() (*string, bool)`

GetServiceOfferingOk returns a tuple with the ServiceOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOffering

`func (o *BackupVmConfiguration) SetServiceOffering(v string)`

SetServiceOffering sets ServiceOffering field to given value.

### HasServiceOffering

`func (o *BackupVmConfiguration) HasServiceOffering() bool`

HasServiceOffering returns a boolean if a field has been set.

### GetType

`func (o *BackupVmConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupVmConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupVmConfiguration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BackupVmConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


