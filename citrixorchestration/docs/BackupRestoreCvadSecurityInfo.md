# BackupRestoreCvadSecurityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityName** | Pointer to **NullableString** |  | [optional] 
**SecurityType** | Pointer to [**CvadSecurityTypes**](CvadSecurityTypes.md) |  | [optional] 
**Secret** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBackupRestoreCvadSecurityInfo

`func NewBackupRestoreCvadSecurityInfo() *BackupRestoreCvadSecurityInfo`

NewBackupRestoreCvadSecurityInfo instantiates a new BackupRestoreCvadSecurityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreCvadSecurityInfoWithDefaults

`func NewBackupRestoreCvadSecurityInfoWithDefaults() *BackupRestoreCvadSecurityInfo`

NewBackupRestoreCvadSecurityInfoWithDefaults instantiates a new BackupRestoreCvadSecurityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityName

`func (o *BackupRestoreCvadSecurityInfo) GetSecurityName() string`

GetSecurityName returns the SecurityName field if non-nil, zero value otherwise.

### GetSecurityNameOk

`func (o *BackupRestoreCvadSecurityInfo) GetSecurityNameOk() (*string, bool)`

GetSecurityNameOk returns a tuple with the SecurityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityName

`func (o *BackupRestoreCvadSecurityInfo) SetSecurityName(v string)`

SetSecurityName sets SecurityName field to given value.

### HasSecurityName

`func (o *BackupRestoreCvadSecurityInfo) HasSecurityName() bool`

HasSecurityName returns a boolean if a field has been set.

### SetSecurityNameNil

`func (o *BackupRestoreCvadSecurityInfo) SetSecurityNameNil(b bool)`

 SetSecurityNameNil sets the value for SecurityName to be an explicit nil

### UnsetSecurityName
`func (o *BackupRestoreCvadSecurityInfo) UnsetSecurityName()`

UnsetSecurityName ensures that no value is present for SecurityName, not even an explicit nil
### GetSecurityType

`func (o *BackupRestoreCvadSecurityInfo) GetSecurityType() CvadSecurityTypes`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *BackupRestoreCvadSecurityInfo) GetSecurityTypeOk() (*CvadSecurityTypes, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *BackupRestoreCvadSecurityInfo) SetSecurityType(v CvadSecurityTypes)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *BackupRestoreCvadSecurityInfo) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### GetSecret

`func (o *BackupRestoreCvadSecurityInfo) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BackupRestoreCvadSecurityInfo) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BackupRestoreCvadSecurityInfo) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BackupRestoreCvadSecurityInfo) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *BackupRestoreCvadSecurityInfo) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *BackupRestoreCvadSecurityInfo) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


