# CitrixCloudServicesAgentHubApiEdgeServersGetResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**WindowsSid** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**CurrentVersion** | Pointer to **string** |  | [optional] 
**CurrentBootstrapperVersion** | Pointer to **string** |  | [optional] 
**ExpectedVersion** | Pointer to **string** |  | [optional] 
**ExpectedBootStrapperVersion** | Pointer to **string** |  | [optional] 
**VersionState** | Pointer to **string** |  | [optional] 
**InMaintenance** | Pointer to **bool** |  | [optional] 
**LeaseEndDateTime** | Pointer to **time.Time** |  | [optional] 
**UpgradeDisabled** | Pointer to **bool** |  | [optional] 
**ConnectorType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**LastContactDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCitrixCloudServicesAgentHubApiEdgeServersGetResultModel

`func NewCitrixCloudServicesAgentHubApiEdgeServersGetResultModel() *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel`

NewCitrixCloudServicesAgentHubApiEdgeServersGetResultModel instantiates a new CitrixCloudServicesAgentHubApiEdgeServersGetResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesAgentHubApiEdgeServersGetResultModelWithDefaults

`func NewCitrixCloudServicesAgentHubApiEdgeServersGetResultModelWithDefaults() *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel`

NewCitrixCloudServicesAgentHubApiEdgeServersGetResultModelWithDefaults instantiates a new CitrixCloudServicesAgentHubApiEdgeServersGetResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFqdn

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetRole

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetWindowsSid

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetWindowsSid() string`

GetWindowsSid returns the WindowsSid field if non-nil, zero value otherwise.

### GetWindowsSidOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetWindowsSidOk() (*string, bool)`

GetWindowsSidOk returns a tuple with the WindowsSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsSid

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetWindowsSid(v string)`

SetWindowsSid sets WindowsSid field to given value.

### HasWindowsSid

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasWindowsSid() bool`

HasWindowsSid returns a boolean if a field has been set.

### GetLocation

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetCurrentBootstrapperVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetCurrentBootstrapperVersion() string`

GetCurrentBootstrapperVersion returns the CurrentBootstrapperVersion field if non-nil, zero value otherwise.

### GetCurrentBootstrapperVersionOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetCurrentBootstrapperVersionOk() (*string, bool)`

GetCurrentBootstrapperVersionOk returns a tuple with the CurrentBootstrapperVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBootstrapperVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetCurrentBootstrapperVersion(v string)`

SetCurrentBootstrapperVersion sets CurrentBootstrapperVersion field to given value.

### HasCurrentBootstrapperVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasCurrentBootstrapperVersion() bool`

HasCurrentBootstrapperVersion returns a boolean if a field has been set.

### GetExpectedVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetExpectedVersion() string`

GetExpectedVersion returns the ExpectedVersion field if non-nil, zero value otherwise.

### GetExpectedVersionOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetExpectedVersionOk() (*string, bool)`

GetExpectedVersionOk returns a tuple with the ExpectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetExpectedVersion(v string)`

SetExpectedVersion sets ExpectedVersion field to given value.

### HasExpectedVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasExpectedVersion() bool`

HasExpectedVersion returns a boolean if a field has been set.

### GetExpectedBootStrapperVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetExpectedBootStrapperVersion() string`

GetExpectedBootStrapperVersion returns the ExpectedBootStrapperVersion field if non-nil, zero value otherwise.

### GetExpectedBootStrapperVersionOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetExpectedBootStrapperVersionOk() (*string, bool)`

GetExpectedBootStrapperVersionOk returns a tuple with the ExpectedBootStrapperVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedBootStrapperVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetExpectedBootStrapperVersion(v string)`

SetExpectedBootStrapperVersion sets ExpectedBootStrapperVersion field to given value.

### HasExpectedBootStrapperVersion

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasExpectedBootStrapperVersion() bool`

HasExpectedBootStrapperVersion returns a boolean if a field has been set.

### GetVersionState

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetVersionState() string`

GetVersionState returns the VersionState field if non-nil, zero value otherwise.

### GetVersionStateOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetVersionStateOk() (*string, bool)`

GetVersionStateOk returns a tuple with the VersionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionState

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetVersionState(v string)`

SetVersionState sets VersionState field to given value.

### HasVersionState

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasVersionState() bool`

HasVersionState returns a boolean if a field has been set.

### GetInMaintenance

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetInMaintenance() bool`

GetInMaintenance returns the InMaintenance field if non-nil, zero value otherwise.

### GetInMaintenanceOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetInMaintenanceOk() (*bool, bool)`

GetInMaintenanceOk returns a tuple with the InMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenance

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetInMaintenance(v bool)`

SetInMaintenance sets InMaintenance field to given value.

### HasInMaintenance

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasInMaintenance() bool`

HasInMaintenance returns a boolean if a field has been set.

### GetLeaseEndDateTime

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetLeaseEndDateTime() time.Time`

GetLeaseEndDateTime returns the LeaseEndDateTime field if non-nil, zero value otherwise.

### GetLeaseEndDateTimeOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetLeaseEndDateTimeOk() (*time.Time, bool)`

GetLeaseEndDateTimeOk returns a tuple with the LeaseEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseEndDateTime

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetLeaseEndDateTime(v time.Time)`

SetLeaseEndDateTime sets LeaseEndDateTime field to given value.

### HasLeaseEndDateTime

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasLeaseEndDateTime() bool`

HasLeaseEndDateTime returns a boolean if a field has been set.

### GetUpgradeDisabled

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetUpgradeDisabled() bool`

GetUpgradeDisabled returns the UpgradeDisabled field if non-nil, zero value otherwise.

### GetUpgradeDisabledOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetUpgradeDisabledOk() (*bool, bool)`

GetUpgradeDisabledOk returns a tuple with the UpgradeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDisabled

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetUpgradeDisabled(v bool)`

SetUpgradeDisabled sets UpgradeDisabled field to given value.

### HasUpgradeDisabled

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasUpgradeDisabled() bool`

HasUpgradeDisabled returns a boolean if a field has been set.

### GetConnectorType

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetStatus

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastContactDate

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetLastContactDate() time.Time`

GetLastContactDate returns the LastContactDate field if non-nil, zero value otherwise.

### GetLastContactDateOk

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) GetLastContactDateOk() (*time.Time, bool)`

GetLastContactDateOk returns a tuple with the LastContactDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContactDate

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) SetLastContactDate(v time.Time)`

SetLastContactDate sets LastContactDate field to given value.

### HasLastContactDate

`func (o *CitrixCloudServicesAgentHubApiEdgeServersGetResultModel) HasLastContactDate() bool`

HasLastContactDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


