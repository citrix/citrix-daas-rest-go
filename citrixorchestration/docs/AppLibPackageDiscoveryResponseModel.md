# AppLibPackageDiscoveryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Get or set a unique persistent identifier for the PackageDiscoverySession. | [optional] 
**DiscoveryProfileUid** | Pointer to **int32** | Get or set a unique persistent identifier for the PackageDiscoveryProfile. | [optional] 
**DesktopGroupUid** | Pointer to **int32** | The UID of the DesktopGroup from which the broker will select a VDA to run the discovery. | [optional] 
**BrokerMachineUid** | Pointer to **int32** | Get or set a unique persistent identifier for the BrokderMachine. | [optional] 
**BrokerMachineSID** | Pointer to **NullableString** | Gets or sets the SID of the broker machine that was chosen to run the discovery. | [optional] 
**Path** | Pointer to **NullableString** | The path to the root directory where the discovery will run. | [optional] 
**ManagementServer** | Pointer to **NullableString** | The url of the App-V Management server that packages will be discovered from. | [optional] 
**PublishingServer** | Pointer to **NullableString** | The url of the App-V server that packages wil be discovered from. | [optional] 
**Status** | Pointer to [**AppLibPackageDiscoveryStatus**](AppLibPackageDiscoveryStatus.md) |  | [optional] 
**StatusMessage** | Pointer to **NullableString** | The status message of PackageDiscoverySession | [optional] 
**ImportedPackages** | Pointer to **int32** | the number of the imported packages | [optional] 
**ImportProgress** | Pointer to **int32** | the progress of package import | [optional] 
**DiscoveredPackages** | Pointer to **int32** | the number of discovered packages | [optional] 
**DiscoveryProgress** | Pointer to **int32** | the progress of the package discovery | [optional] 
**ExpectedPackages** | Pointer to **int32** | the number of the expected packages | [optional] 
**ReportingProgress** | Pointer to **int32** | Gets the percentage value of the reporting progress. | [optional] 
**ReportedPackages** | Pointer to **int32** | Gets or sets the number of packages that have reported so far.              | [optional] 
**DiscoverySessionTime** | Pointer to **NullableString** | Get the discovery session time | [optional] 

## Methods

### NewAppLibPackageDiscoveryResponseModel

`func NewAppLibPackageDiscoveryResponseModel() *AppLibPackageDiscoveryResponseModel`

NewAppLibPackageDiscoveryResponseModel instantiates a new AppLibPackageDiscoveryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLibPackageDiscoveryResponseModelWithDefaults

`func NewAppLibPackageDiscoveryResponseModelWithDefaults() *AppLibPackageDiscoveryResponseModel`

NewAppLibPackageDiscoveryResponseModelWithDefaults instantiates a new AppLibPackageDiscoveryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppLibPackageDiscoveryResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppLibPackageDiscoveryResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppLibPackageDiscoveryResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppLibPackageDiscoveryResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AppLibPackageDiscoveryResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AppLibPackageDiscoveryResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDiscoveryProfileUid

`func (o *AppLibPackageDiscoveryResponseModel) GetDiscoveryProfileUid() int32`

GetDiscoveryProfileUid returns the DiscoveryProfileUid field if non-nil, zero value otherwise.

### GetDiscoveryProfileUidOk

`func (o *AppLibPackageDiscoveryResponseModel) GetDiscoveryProfileUidOk() (*int32, bool)`

GetDiscoveryProfileUidOk returns a tuple with the DiscoveryProfileUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryProfileUid

`func (o *AppLibPackageDiscoveryResponseModel) SetDiscoveryProfileUid(v int32)`

SetDiscoveryProfileUid sets DiscoveryProfileUid field to given value.

### HasDiscoveryProfileUid

`func (o *AppLibPackageDiscoveryResponseModel) HasDiscoveryProfileUid() bool`

HasDiscoveryProfileUid returns a boolean if a field has been set.

### GetDesktopGroupUid

`func (o *AppLibPackageDiscoveryResponseModel) GetDesktopGroupUid() int32`

GetDesktopGroupUid returns the DesktopGroupUid field if non-nil, zero value otherwise.

### GetDesktopGroupUidOk

`func (o *AppLibPackageDiscoveryResponseModel) GetDesktopGroupUidOk() (*int32, bool)`

GetDesktopGroupUidOk returns a tuple with the DesktopGroupUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopGroupUid

`func (o *AppLibPackageDiscoveryResponseModel) SetDesktopGroupUid(v int32)`

SetDesktopGroupUid sets DesktopGroupUid field to given value.

### HasDesktopGroupUid

`func (o *AppLibPackageDiscoveryResponseModel) HasDesktopGroupUid() bool`

HasDesktopGroupUid returns a boolean if a field has been set.

### GetBrokerMachineUid

`func (o *AppLibPackageDiscoveryResponseModel) GetBrokerMachineUid() int32`

GetBrokerMachineUid returns the BrokerMachineUid field if non-nil, zero value otherwise.

### GetBrokerMachineUidOk

`func (o *AppLibPackageDiscoveryResponseModel) GetBrokerMachineUidOk() (*int32, bool)`

GetBrokerMachineUidOk returns a tuple with the BrokerMachineUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerMachineUid

`func (o *AppLibPackageDiscoveryResponseModel) SetBrokerMachineUid(v int32)`

SetBrokerMachineUid sets BrokerMachineUid field to given value.

### HasBrokerMachineUid

`func (o *AppLibPackageDiscoveryResponseModel) HasBrokerMachineUid() bool`

HasBrokerMachineUid returns a boolean if a field has been set.

### GetBrokerMachineSID

`func (o *AppLibPackageDiscoveryResponseModel) GetBrokerMachineSID() string`

GetBrokerMachineSID returns the BrokerMachineSID field if non-nil, zero value otherwise.

### GetBrokerMachineSIDOk

`func (o *AppLibPackageDiscoveryResponseModel) GetBrokerMachineSIDOk() (*string, bool)`

GetBrokerMachineSIDOk returns a tuple with the BrokerMachineSID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerMachineSID

`func (o *AppLibPackageDiscoveryResponseModel) SetBrokerMachineSID(v string)`

SetBrokerMachineSID sets BrokerMachineSID field to given value.

### HasBrokerMachineSID

`func (o *AppLibPackageDiscoveryResponseModel) HasBrokerMachineSID() bool`

HasBrokerMachineSID returns a boolean if a field has been set.

### SetBrokerMachineSIDNil

`func (o *AppLibPackageDiscoveryResponseModel) SetBrokerMachineSIDNil(b bool)`

 SetBrokerMachineSIDNil sets the value for BrokerMachineSID to be an explicit nil

### UnsetBrokerMachineSID
`func (o *AppLibPackageDiscoveryResponseModel) UnsetBrokerMachineSID()`

UnsetBrokerMachineSID ensures that no value is present for BrokerMachineSID, not even an explicit nil
### GetPath

`func (o *AppLibPackageDiscoveryResponseModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppLibPackageDiscoveryResponseModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppLibPackageDiscoveryResponseModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AppLibPackageDiscoveryResponseModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *AppLibPackageDiscoveryResponseModel) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *AppLibPackageDiscoveryResponseModel) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetManagementServer

`func (o *AppLibPackageDiscoveryResponseModel) GetManagementServer() string`

GetManagementServer returns the ManagementServer field if non-nil, zero value otherwise.

### GetManagementServerOk

`func (o *AppLibPackageDiscoveryResponseModel) GetManagementServerOk() (*string, bool)`

GetManagementServerOk returns a tuple with the ManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServer

`func (o *AppLibPackageDiscoveryResponseModel) SetManagementServer(v string)`

SetManagementServer sets ManagementServer field to given value.

### HasManagementServer

`func (o *AppLibPackageDiscoveryResponseModel) HasManagementServer() bool`

HasManagementServer returns a boolean if a field has been set.

### SetManagementServerNil

`func (o *AppLibPackageDiscoveryResponseModel) SetManagementServerNil(b bool)`

 SetManagementServerNil sets the value for ManagementServer to be an explicit nil

### UnsetManagementServer
`func (o *AppLibPackageDiscoveryResponseModel) UnsetManagementServer()`

UnsetManagementServer ensures that no value is present for ManagementServer, not even an explicit nil
### GetPublishingServer

`func (o *AppLibPackageDiscoveryResponseModel) GetPublishingServer() string`

GetPublishingServer returns the PublishingServer field if non-nil, zero value otherwise.

### GetPublishingServerOk

`func (o *AppLibPackageDiscoveryResponseModel) GetPublishingServerOk() (*string, bool)`

GetPublishingServerOk returns a tuple with the PublishingServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingServer

`func (o *AppLibPackageDiscoveryResponseModel) SetPublishingServer(v string)`

SetPublishingServer sets PublishingServer field to given value.

### HasPublishingServer

`func (o *AppLibPackageDiscoveryResponseModel) HasPublishingServer() bool`

HasPublishingServer returns a boolean if a field has been set.

### SetPublishingServerNil

`func (o *AppLibPackageDiscoveryResponseModel) SetPublishingServerNil(b bool)`

 SetPublishingServerNil sets the value for PublishingServer to be an explicit nil

### UnsetPublishingServer
`func (o *AppLibPackageDiscoveryResponseModel) UnsetPublishingServer()`

UnsetPublishingServer ensures that no value is present for PublishingServer, not even an explicit nil
### GetStatus

`func (o *AppLibPackageDiscoveryResponseModel) GetStatus() AppLibPackageDiscoveryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppLibPackageDiscoveryResponseModel) GetStatusOk() (*AppLibPackageDiscoveryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppLibPackageDiscoveryResponseModel) SetStatus(v AppLibPackageDiscoveryStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppLibPackageDiscoveryResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AppLibPackageDiscoveryResponseModel) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AppLibPackageDiscoveryResponseModel) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AppLibPackageDiscoveryResponseModel) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AppLibPackageDiscoveryResponseModel) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AppLibPackageDiscoveryResponseModel) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AppLibPackageDiscoveryResponseModel) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetImportedPackages

`func (o *AppLibPackageDiscoveryResponseModel) GetImportedPackages() int32`

GetImportedPackages returns the ImportedPackages field if non-nil, zero value otherwise.

### GetImportedPackagesOk

`func (o *AppLibPackageDiscoveryResponseModel) GetImportedPackagesOk() (*int32, bool)`

GetImportedPackagesOk returns a tuple with the ImportedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedPackages

`func (o *AppLibPackageDiscoveryResponseModel) SetImportedPackages(v int32)`

SetImportedPackages sets ImportedPackages field to given value.

### HasImportedPackages

`func (o *AppLibPackageDiscoveryResponseModel) HasImportedPackages() bool`

HasImportedPackages returns a boolean if a field has been set.

### GetImportProgress

`func (o *AppLibPackageDiscoveryResponseModel) GetImportProgress() int32`

GetImportProgress returns the ImportProgress field if non-nil, zero value otherwise.

### GetImportProgressOk

`func (o *AppLibPackageDiscoveryResponseModel) GetImportProgressOk() (*int32, bool)`

GetImportProgressOk returns a tuple with the ImportProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportProgress

`func (o *AppLibPackageDiscoveryResponseModel) SetImportProgress(v int32)`

SetImportProgress sets ImportProgress field to given value.

### HasImportProgress

`func (o *AppLibPackageDiscoveryResponseModel) HasImportProgress() bool`

HasImportProgress returns a boolean if a field has been set.

### GetDiscoveredPackages

`func (o *AppLibPackageDiscoveryResponseModel) GetDiscoveredPackages() int32`

GetDiscoveredPackages returns the DiscoveredPackages field if non-nil, zero value otherwise.

### GetDiscoveredPackagesOk

`func (o *AppLibPackageDiscoveryResponseModel) GetDiscoveredPackagesOk() (*int32, bool)`

GetDiscoveredPackagesOk returns a tuple with the DiscoveredPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredPackages

`func (o *AppLibPackageDiscoveryResponseModel) SetDiscoveredPackages(v int32)`

SetDiscoveredPackages sets DiscoveredPackages field to given value.

### HasDiscoveredPackages

`func (o *AppLibPackageDiscoveryResponseModel) HasDiscoveredPackages() bool`

HasDiscoveredPackages returns a boolean if a field has been set.

### GetDiscoveryProgress

`func (o *AppLibPackageDiscoveryResponseModel) GetDiscoveryProgress() int32`

GetDiscoveryProgress returns the DiscoveryProgress field if non-nil, zero value otherwise.

### GetDiscoveryProgressOk

`func (o *AppLibPackageDiscoveryResponseModel) GetDiscoveryProgressOk() (*int32, bool)`

GetDiscoveryProgressOk returns a tuple with the DiscoveryProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryProgress

`func (o *AppLibPackageDiscoveryResponseModel) SetDiscoveryProgress(v int32)`

SetDiscoveryProgress sets DiscoveryProgress field to given value.

### HasDiscoveryProgress

`func (o *AppLibPackageDiscoveryResponseModel) HasDiscoveryProgress() bool`

HasDiscoveryProgress returns a boolean if a field has been set.

### GetExpectedPackages

`func (o *AppLibPackageDiscoveryResponseModel) GetExpectedPackages() int32`

GetExpectedPackages returns the ExpectedPackages field if non-nil, zero value otherwise.

### GetExpectedPackagesOk

`func (o *AppLibPackageDiscoveryResponseModel) GetExpectedPackagesOk() (*int32, bool)`

GetExpectedPackagesOk returns a tuple with the ExpectedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedPackages

`func (o *AppLibPackageDiscoveryResponseModel) SetExpectedPackages(v int32)`

SetExpectedPackages sets ExpectedPackages field to given value.

### HasExpectedPackages

`func (o *AppLibPackageDiscoveryResponseModel) HasExpectedPackages() bool`

HasExpectedPackages returns a boolean if a field has been set.

### GetReportingProgress

`func (o *AppLibPackageDiscoveryResponseModel) GetReportingProgress() int32`

GetReportingProgress returns the ReportingProgress field if non-nil, zero value otherwise.

### GetReportingProgressOk

`func (o *AppLibPackageDiscoveryResponseModel) GetReportingProgressOk() (*int32, bool)`

GetReportingProgressOk returns a tuple with the ReportingProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingProgress

`func (o *AppLibPackageDiscoveryResponseModel) SetReportingProgress(v int32)`

SetReportingProgress sets ReportingProgress field to given value.

### HasReportingProgress

`func (o *AppLibPackageDiscoveryResponseModel) HasReportingProgress() bool`

HasReportingProgress returns a boolean if a field has been set.

### GetReportedPackages

`func (o *AppLibPackageDiscoveryResponseModel) GetReportedPackages() int32`

GetReportedPackages returns the ReportedPackages field if non-nil, zero value otherwise.

### GetReportedPackagesOk

`func (o *AppLibPackageDiscoveryResponseModel) GetReportedPackagesOk() (*int32, bool)`

GetReportedPackagesOk returns a tuple with the ReportedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedPackages

`func (o *AppLibPackageDiscoveryResponseModel) SetReportedPackages(v int32)`

SetReportedPackages sets ReportedPackages field to given value.

### HasReportedPackages

`func (o *AppLibPackageDiscoveryResponseModel) HasReportedPackages() bool`

HasReportedPackages returns a boolean if a field has been set.

### GetDiscoverySessionTime

`func (o *AppLibPackageDiscoveryResponseModel) GetDiscoverySessionTime() string`

GetDiscoverySessionTime returns the DiscoverySessionTime field if non-nil, zero value otherwise.

### GetDiscoverySessionTimeOk

`func (o *AppLibPackageDiscoveryResponseModel) GetDiscoverySessionTimeOk() (*string, bool)`

GetDiscoverySessionTimeOk returns a tuple with the DiscoverySessionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverySessionTime

`func (o *AppLibPackageDiscoveryResponseModel) SetDiscoverySessionTime(v string)`

SetDiscoverySessionTime sets DiscoverySessionTime field to given value.

### HasDiscoverySessionTime

`func (o *AppLibPackageDiscoveryResponseModel) HasDiscoverySessionTime() bool`

HasDiscoverySessionTime returns a boolean if a field has been set.

### SetDiscoverySessionTimeNil

`func (o *AppLibPackageDiscoveryResponseModel) SetDiscoverySessionTimeNil(b bool)`

 SetDiscoverySessionTimeNil sets the value for DiscoverySessionTime to be an explicit nil

### UnsetDiscoverySessionTime
`func (o *AppLibPackageDiscoveryResponseModel) UnsetDiscoverySessionTime()`

UnsetDiscoverySessionTime ensures that no value is present for DiscoverySessionTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


