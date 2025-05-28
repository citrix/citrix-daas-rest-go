# ConnectorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the connector. | [optional] 
**State** | Pointer to [**ConnectorDetailsState**](ConnectorDetailsState.md) | State of the connector. | [optional] 
**Fqdn** | Pointer to **string** | Hostname of the connector. | [optional] 
**OrganizationalUnit** | Pointer to **string** |  | [optional] 
**LastContactDate** | Pointer to **time.Time** | Time at which the connector made last contact. | [optional] 
**Status** | Pointer to **string** | Connection status of the connector | [optional] 
**StatusMessage** | Pointer to **string** | Details about the connector status. | [optional] 
**TransactionId** | Pointer to **string** | ID of the transaction in which the connector was provisioned. | [optional] 
**IpAddress** | Pointer to **string** | The IP Address of the connector | [optional] 
**AzureResourceGroup** | Pointer to **string** | The resource group of the connector | [optional] 
**VmSize** | Pointer to **string** | The vm size of the resource group | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


