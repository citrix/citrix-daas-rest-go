# ManagedCapacityAssignedCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**CustomerName** | **string** |  | 
**OrgId** | **string** |  | 
**VirtualSiteId** | **string** |  | 
**Edition** | **string** |  | 
**State** | [**CommissionState**](CommissionState.md) |  | 
**TenantCustomerId** | Pointer to **string** |  | [optional] 
**TenantVirtualSiteId** | Pointer to **string** |  | [optional] 
**TenantOrgId** | Pointer to **string** |  | [optional] 
**TenantName** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedCapacityAssignedCustomer

`func NewManagedCapacityAssignedCustomer(customerId string, customerName string, orgId string, virtualSiteId string, edition string, state CommissionState, ) *ManagedCapacityAssignedCustomer`

NewManagedCapacityAssignedCustomer instantiates a new ManagedCapacityAssignedCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedCapacityAssignedCustomerWithDefaults

`func NewManagedCapacityAssignedCustomerWithDefaults() *ManagedCapacityAssignedCustomer`

NewManagedCapacityAssignedCustomerWithDefaults instantiates a new ManagedCapacityAssignedCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *ManagedCapacityAssignedCustomer) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ManagedCapacityAssignedCustomer) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ManagedCapacityAssignedCustomer) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCustomerName

`func (o *ManagedCapacityAssignedCustomer) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *ManagedCapacityAssignedCustomer) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *ManagedCapacityAssignedCustomer) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### GetOrgId

`func (o *ManagedCapacityAssignedCustomer) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ManagedCapacityAssignedCustomer) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ManagedCapacityAssignedCustomer) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetVirtualSiteId

`func (o *ManagedCapacityAssignedCustomer) GetVirtualSiteId() string`

GetVirtualSiteId returns the VirtualSiteId field if non-nil, zero value otherwise.

### GetVirtualSiteIdOk

`func (o *ManagedCapacityAssignedCustomer) GetVirtualSiteIdOk() (*string, bool)`

GetVirtualSiteIdOk returns a tuple with the VirtualSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSiteId

`func (o *ManagedCapacityAssignedCustomer) SetVirtualSiteId(v string)`

SetVirtualSiteId sets VirtualSiteId field to given value.


### GetEdition

`func (o *ManagedCapacityAssignedCustomer) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *ManagedCapacityAssignedCustomer) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *ManagedCapacityAssignedCustomer) SetEdition(v string)`

SetEdition sets Edition field to given value.


### GetState

`func (o *ManagedCapacityAssignedCustomer) GetState() CommissionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagedCapacityAssignedCustomer) GetStateOk() (*CommissionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagedCapacityAssignedCustomer) SetState(v CommissionState)`

SetState sets State field to given value.


### GetTenantCustomerId

`func (o *ManagedCapacityAssignedCustomer) GetTenantCustomerId() string`

GetTenantCustomerId returns the TenantCustomerId field if non-nil, zero value otherwise.

### GetTenantCustomerIdOk

`func (o *ManagedCapacityAssignedCustomer) GetTenantCustomerIdOk() (*string, bool)`

GetTenantCustomerIdOk returns a tuple with the TenantCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCustomerId

`func (o *ManagedCapacityAssignedCustomer) SetTenantCustomerId(v string)`

SetTenantCustomerId sets TenantCustomerId field to given value.

### HasTenantCustomerId

`func (o *ManagedCapacityAssignedCustomer) HasTenantCustomerId() bool`

HasTenantCustomerId returns a boolean if a field has been set.

### GetTenantVirtualSiteId

`func (o *ManagedCapacityAssignedCustomer) GetTenantVirtualSiteId() string`

GetTenantVirtualSiteId returns the TenantVirtualSiteId field if non-nil, zero value otherwise.

### GetTenantVirtualSiteIdOk

`func (o *ManagedCapacityAssignedCustomer) GetTenantVirtualSiteIdOk() (*string, bool)`

GetTenantVirtualSiteIdOk returns a tuple with the TenantVirtualSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantVirtualSiteId

`func (o *ManagedCapacityAssignedCustomer) SetTenantVirtualSiteId(v string)`

SetTenantVirtualSiteId sets TenantVirtualSiteId field to given value.

### HasTenantVirtualSiteId

`func (o *ManagedCapacityAssignedCustomer) HasTenantVirtualSiteId() bool`

HasTenantVirtualSiteId returns a boolean if a field has been set.

### GetTenantOrgId

`func (o *ManagedCapacityAssignedCustomer) GetTenantOrgId() string`

GetTenantOrgId returns the TenantOrgId field if non-nil, zero value otherwise.

### GetTenantOrgIdOk

`func (o *ManagedCapacityAssignedCustomer) GetTenantOrgIdOk() (*string, bool)`

GetTenantOrgIdOk returns a tuple with the TenantOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantOrgId

`func (o *ManagedCapacityAssignedCustomer) SetTenantOrgId(v string)`

SetTenantOrgId sets TenantOrgId field to given value.

### HasTenantOrgId

`func (o *ManagedCapacityAssignedCustomer) HasTenantOrgId() bool`

HasTenantOrgId returns a boolean if a field has been set.

### GetTenantName

`func (o *ManagedCapacityAssignedCustomer) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *ManagedCapacityAssignedCustomer) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *ManagedCapacityAssignedCustomer) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *ManagedCapacityAssignedCustomer) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


