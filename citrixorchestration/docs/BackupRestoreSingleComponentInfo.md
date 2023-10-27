# BackupRestoreSingleComponentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalizedName** | Pointer to **NullableString** | Localized component name | [optional] 
**Component** | Pointer to [**BckRstrAutoConfigComponents**](BckRstrAutoConfigComponents.md) |  | [optional] 
**ControlPlane** | Pointer to [**BackupRestoreComponentControlPlane**](BackupRestoreComponentControlPlane.md) |  | [optional] 
**PreImportCheckMode** | Pointer to **bool** | Pre-Import check mode; prompt for secrets prior to restoring | [optional] 
**MembersViewable** | Pointer to **bool** | Compnent member names are viewable | [optional] 
**WithPrequisites** | Pointer to **bool** | Component supports restoring with prerequisites | [optional] 
**CanBeRestored** | Pointer to **bool** | Component can be restored | [optional] 
**HasDescription** | Pointer to **bool** | Component members have descriptions | [optional] 
**OnPremOnly** | Pointer to **bool** | Component valid On-Prem only | [optional] 

## Methods

### NewBackupRestoreSingleComponentInfo

`func NewBackupRestoreSingleComponentInfo() *BackupRestoreSingleComponentInfo`

NewBackupRestoreSingleComponentInfo instantiates a new BackupRestoreSingleComponentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreSingleComponentInfoWithDefaults

`func NewBackupRestoreSingleComponentInfoWithDefaults() *BackupRestoreSingleComponentInfo`

NewBackupRestoreSingleComponentInfoWithDefaults instantiates a new BackupRestoreSingleComponentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalizedName

`func (o *BackupRestoreSingleComponentInfo) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *BackupRestoreSingleComponentInfo) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *BackupRestoreSingleComponentInfo) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *BackupRestoreSingleComponentInfo) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### SetLocalizedNameNil

`func (o *BackupRestoreSingleComponentInfo) SetLocalizedNameNil(b bool)`

 SetLocalizedNameNil sets the value for LocalizedName to be an explicit nil

### UnsetLocalizedName
`func (o *BackupRestoreSingleComponentInfo) UnsetLocalizedName()`

UnsetLocalizedName ensures that no value is present for LocalizedName, not even an explicit nil
### GetComponent

`func (o *BackupRestoreSingleComponentInfo) GetComponent() BckRstrAutoConfigComponents`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *BackupRestoreSingleComponentInfo) GetComponentOk() (*BckRstrAutoConfigComponents, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *BackupRestoreSingleComponentInfo) SetComponent(v BckRstrAutoConfigComponents)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *BackupRestoreSingleComponentInfo) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetControlPlane

`func (o *BackupRestoreSingleComponentInfo) GetControlPlane() BackupRestoreComponentControlPlane`

GetControlPlane returns the ControlPlane field if non-nil, zero value otherwise.

### GetControlPlaneOk

`func (o *BackupRestoreSingleComponentInfo) GetControlPlaneOk() (*BackupRestoreComponentControlPlane, bool)`

GetControlPlaneOk returns a tuple with the ControlPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlane

`func (o *BackupRestoreSingleComponentInfo) SetControlPlane(v BackupRestoreComponentControlPlane)`

SetControlPlane sets ControlPlane field to given value.

### HasControlPlane

`func (o *BackupRestoreSingleComponentInfo) HasControlPlane() bool`

HasControlPlane returns a boolean if a field has been set.

### GetPreImportCheckMode

`func (o *BackupRestoreSingleComponentInfo) GetPreImportCheckMode() bool`

GetPreImportCheckMode returns the PreImportCheckMode field if non-nil, zero value otherwise.

### GetPreImportCheckModeOk

`func (o *BackupRestoreSingleComponentInfo) GetPreImportCheckModeOk() (*bool, bool)`

GetPreImportCheckModeOk returns a tuple with the PreImportCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreImportCheckMode

`func (o *BackupRestoreSingleComponentInfo) SetPreImportCheckMode(v bool)`

SetPreImportCheckMode sets PreImportCheckMode field to given value.

### HasPreImportCheckMode

`func (o *BackupRestoreSingleComponentInfo) HasPreImportCheckMode() bool`

HasPreImportCheckMode returns a boolean if a field has been set.

### GetMembersViewable

`func (o *BackupRestoreSingleComponentInfo) GetMembersViewable() bool`

GetMembersViewable returns the MembersViewable field if non-nil, zero value otherwise.

### GetMembersViewableOk

`func (o *BackupRestoreSingleComponentInfo) GetMembersViewableOk() (*bool, bool)`

GetMembersViewableOk returns a tuple with the MembersViewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersViewable

`func (o *BackupRestoreSingleComponentInfo) SetMembersViewable(v bool)`

SetMembersViewable sets MembersViewable field to given value.

### HasMembersViewable

`func (o *BackupRestoreSingleComponentInfo) HasMembersViewable() bool`

HasMembersViewable returns a boolean if a field has been set.

### GetWithPrequisites

`func (o *BackupRestoreSingleComponentInfo) GetWithPrequisites() bool`

GetWithPrequisites returns the WithPrequisites field if non-nil, zero value otherwise.

### GetWithPrequisitesOk

`func (o *BackupRestoreSingleComponentInfo) GetWithPrequisitesOk() (*bool, bool)`

GetWithPrequisitesOk returns a tuple with the WithPrequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithPrequisites

`func (o *BackupRestoreSingleComponentInfo) SetWithPrequisites(v bool)`

SetWithPrequisites sets WithPrequisites field to given value.

### HasWithPrequisites

`func (o *BackupRestoreSingleComponentInfo) HasWithPrequisites() bool`

HasWithPrequisites returns a boolean if a field has been set.

### GetCanBeRestored

`func (o *BackupRestoreSingleComponentInfo) GetCanBeRestored() bool`

GetCanBeRestored returns the CanBeRestored field if non-nil, zero value otherwise.

### GetCanBeRestoredOk

`func (o *BackupRestoreSingleComponentInfo) GetCanBeRestoredOk() (*bool, bool)`

GetCanBeRestoredOk returns a tuple with the CanBeRestored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeRestored

`func (o *BackupRestoreSingleComponentInfo) SetCanBeRestored(v bool)`

SetCanBeRestored sets CanBeRestored field to given value.

### HasCanBeRestored

`func (o *BackupRestoreSingleComponentInfo) HasCanBeRestored() bool`

HasCanBeRestored returns a boolean if a field has been set.

### GetHasDescription

`func (o *BackupRestoreSingleComponentInfo) GetHasDescription() bool`

GetHasDescription returns the HasDescription field if non-nil, zero value otherwise.

### GetHasDescriptionOk

`func (o *BackupRestoreSingleComponentInfo) GetHasDescriptionOk() (*bool, bool)`

GetHasDescriptionOk returns a tuple with the HasDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDescription

`func (o *BackupRestoreSingleComponentInfo) SetHasDescription(v bool)`

SetHasDescription sets HasDescription field to given value.

### HasHasDescription

`func (o *BackupRestoreSingleComponentInfo) HasHasDescription() bool`

HasHasDescription returns a boolean if a field has been set.

### GetOnPremOnly

`func (o *BackupRestoreSingleComponentInfo) GetOnPremOnly() bool`

GetOnPremOnly returns the OnPremOnly field if non-nil, zero value otherwise.

### GetOnPremOnlyOk

`func (o *BackupRestoreSingleComponentInfo) GetOnPremOnlyOk() (*bool, bool)`

GetOnPremOnlyOk returns a tuple with the OnPremOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremOnly

`func (o *BackupRestoreSingleComponentInfo) SetOnPremOnly(v bool)`

SetOnPremOnly sets OnPremOnly field to given value.

### HasOnPremOnly

`func (o *BackupRestoreSingleComponentInfo) HasOnPremOnly() bool`

HasOnPremOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


