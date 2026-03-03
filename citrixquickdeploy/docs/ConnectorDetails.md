# ConnectorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | ID of the connector. | [optional] 
**State** | Pointer to [**ConnectorDetailsState**](ConnectorDetailsState.md) | State of the connector. | [optional] 
**Fqdn** | Pointer to **NullableString** | Hostname of the connector. | [optional] 
**OrganizationalUnit** | Pointer to **NullableString** |  | [optional] 
**LastContactDate** | Pointer to **NullableTime** | Time at which the connector made last contact. | [optional] 
**Status** | Pointer to **NullableString** | Connection status of the connector | [optional] 
**StatusMessage** | Pointer to **NullableString** | Details about the connector status. | [optional] 
**TransactionId** | Pointer to **NullableString** | ID of the transaction in which the connector was provisioned. | [optional] 
**IpAddress** | Pointer to **NullableString** | The IP Address of the connector | [optional] 
**AzureResourceGroup** | Pointer to **NullableString** | The resource group of the connector | [optional] 
**VmSize** | Pointer to **NullableString** | The vm size of the resource group | [optional] 

## Methods

### NewConnectorDetails

`func NewConnectorDetails() *ConnectorDetails`

NewConnectorDetails instantiates a new ConnectorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorDetailsWithDefaults

`func NewConnectorDetailsWithDefaults() *ConnectorDetails`

NewConnectorDetailsWithDefaults instantiates a new ConnectorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectorDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectorDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ConnectorDetails) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ConnectorDetails) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetState

`func (o *ConnectorDetails) GetState() ConnectorDetailsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectorDetails) GetStateOk() (*ConnectorDetailsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectorDetails) SetState(v ConnectorDetailsState)`

SetState sets State field to given value.

### HasState

`func (o *ConnectorDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetFqdn

`func (o *ConnectorDetails) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ConnectorDetails) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ConnectorDetails) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ConnectorDetails) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *ConnectorDetails) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *ConnectorDetails) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetOrganizationalUnit

`func (o *ConnectorDetails) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *ConnectorDetails) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *ConnectorDetails) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *ConnectorDetails) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### SetOrganizationalUnitNil

`func (o *ConnectorDetails) SetOrganizationalUnitNil(b bool)`

 SetOrganizationalUnitNil sets the value for OrganizationalUnit to be an explicit nil

### UnsetOrganizationalUnit
`func (o *ConnectorDetails) UnsetOrganizationalUnit()`

UnsetOrganizationalUnit ensures that no value is present for OrganizationalUnit, not even an explicit nil
### GetLastContactDate

`func (o *ConnectorDetails) GetLastContactDate() time.Time`

GetLastContactDate returns the LastContactDate field if non-nil, zero value otherwise.

### GetLastContactDateOk

`func (o *ConnectorDetails) GetLastContactDateOk() (*time.Time, bool)`

GetLastContactDateOk returns a tuple with the LastContactDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContactDate

`func (o *ConnectorDetails) SetLastContactDate(v time.Time)`

SetLastContactDate sets LastContactDate field to given value.

### HasLastContactDate

`func (o *ConnectorDetails) HasLastContactDate() bool`

HasLastContactDate returns a boolean if a field has been set.

### SetLastContactDateNil

`func (o *ConnectorDetails) SetLastContactDateNil(b bool)`

 SetLastContactDateNil sets the value for LastContactDate to be an explicit nil

### UnsetLastContactDate
`func (o *ConnectorDetails) UnsetLastContactDate()`

UnsetLastContactDate ensures that no value is present for LastContactDate, not even an explicit nil
### GetStatus

`func (o *ConnectorDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectorDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ConnectorDetails) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ConnectorDetails) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusMessage

`func (o *ConnectorDetails) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ConnectorDetails) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ConnectorDetails) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ConnectorDetails) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ConnectorDetails) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ConnectorDetails) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetTransactionId

`func (o *ConnectorDetails) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ConnectorDetails) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ConnectorDetails) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ConnectorDetails) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *ConnectorDetails) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *ConnectorDetails) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetIpAddress

`func (o *ConnectorDetails) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ConnectorDetails) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ConnectorDetails) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ConnectorDetails) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *ConnectorDetails) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *ConnectorDetails) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetAzureResourceGroup

`func (o *ConnectorDetails) GetAzureResourceGroup() string`

GetAzureResourceGroup returns the AzureResourceGroup field if non-nil, zero value otherwise.

### GetAzureResourceGroupOk

`func (o *ConnectorDetails) GetAzureResourceGroupOk() (*string, bool)`

GetAzureResourceGroupOk returns a tuple with the AzureResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroup

`func (o *ConnectorDetails) SetAzureResourceGroup(v string)`

SetAzureResourceGroup sets AzureResourceGroup field to given value.

### HasAzureResourceGroup

`func (o *ConnectorDetails) HasAzureResourceGroup() bool`

HasAzureResourceGroup returns a boolean if a field has been set.

### SetAzureResourceGroupNil

`func (o *ConnectorDetails) SetAzureResourceGroupNil(b bool)`

 SetAzureResourceGroupNil sets the value for AzureResourceGroup to be an explicit nil

### UnsetAzureResourceGroup
`func (o *ConnectorDetails) UnsetAzureResourceGroup()`

UnsetAzureResourceGroup ensures that no value is present for AzureResourceGroup, not even an explicit nil
### GetVmSize

`func (o *ConnectorDetails) GetVmSize() string`

GetVmSize returns the VmSize field if non-nil, zero value otherwise.

### GetVmSizeOk

`func (o *ConnectorDetails) GetVmSizeOk() (*string, bool)`

GetVmSizeOk returns a tuple with the VmSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSize

`func (o *ConnectorDetails) SetVmSize(v string)`

SetVmSize sets VmSize field to given value.

### HasVmSize

`func (o *ConnectorDetails) HasVmSize() bool`

HasVmSize returns a boolean if a field has been set.

### SetVmSizeNil

`func (o *ConnectorDetails) SetVmSizeNil(b bool)`

 SetVmSizeNil sets the value for VmSize to be an explicit nil

### UnsetVmSize
`func (o *ConnectorDetails) UnsetVmSize()`

UnsetVmSize ensures that no value is present for VmSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


