# AddRemotePcCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the catalog. | 
**ResourceLocationId** | **string** | ID of the Resource Location to associate with this catalog. | 
**Type** | Pointer to [**AddRemotePcType**](AddRemotePcType.md) | Indicates if we are creating a catalog for Static or Random Assignment. | [optional] 
**EnrollmentScopes** | Pointer to [**[]RemotePcEnrollmentScope**](RemotePcEnrollmentScope.md) | List of OUs that will be associated with this catalog. | [optional] 
**MachineAssignments** | Pointer to [**[]AddCatalogMachineAssignment**](AddCatalogMachineAssignment.md) | List of Machines to add to the catalog and the users to assign them to. | [optional] 
**CspCustomerId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null. | [optional] 
**CspSiteId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null. | [optional] 

## Methods

### NewAddRemotePcCatalog

`func NewAddRemotePcCatalog(name string, resourceLocationId string, ) *AddRemotePcCatalog`

NewAddRemotePcCatalog instantiates a new AddRemotePcCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRemotePcCatalogWithDefaults

`func NewAddRemotePcCatalogWithDefaults() *AddRemotePcCatalog`

NewAddRemotePcCatalogWithDefaults instantiates a new AddRemotePcCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddRemotePcCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddRemotePcCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddRemotePcCatalog) SetName(v string)`

SetName sets Name field to given value.


### GetResourceLocationId

`func (o *AddRemotePcCatalog) GetResourceLocationId() string`

GetResourceLocationId returns the ResourceLocationId field if non-nil, zero value otherwise.

### GetResourceLocationIdOk

`func (o *AddRemotePcCatalog) GetResourceLocationIdOk() (*string, bool)`

GetResourceLocationIdOk returns a tuple with the ResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocationId

`func (o *AddRemotePcCatalog) SetResourceLocationId(v string)`

SetResourceLocationId sets ResourceLocationId field to given value.


### GetType

`func (o *AddRemotePcCatalog) GetType() AddRemotePcType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddRemotePcCatalog) GetTypeOk() (*AddRemotePcType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddRemotePcCatalog) SetType(v AddRemotePcType)`

SetType sets Type field to given value.

### HasType

`func (o *AddRemotePcCatalog) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnrollmentScopes

`func (o *AddRemotePcCatalog) GetEnrollmentScopes() []RemotePcEnrollmentScope`

GetEnrollmentScopes returns the EnrollmentScopes field if non-nil, zero value otherwise.

### GetEnrollmentScopesOk

`func (o *AddRemotePcCatalog) GetEnrollmentScopesOk() (*[]RemotePcEnrollmentScope, bool)`

GetEnrollmentScopesOk returns a tuple with the EnrollmentScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentScopes

`func (o *AddRemotePcCatalog) SetEnrollmentScopes(v []RemotePcEnrollmentScope)`

SetEnrollmentScopes sets EnrollmentScopes field to given value.

### HasEnrollmentScopes

`func (o *AddRemotePcCatalog) HasEnrollmentScopes() bool`

HasEnrollmentScopes returns a boolean if a field has been set.

### GetMachineAssignments

`func (o *AddRemotePcCatalog) GetMachineAssignments() []AddCatalogMachineAssignment`

GetMachineAssignments returns the MachineAssignments field if non-nil, zero value otherwise.

### GetMachineAssignmentsOk

`func (o *AddRemotePcCatalog) GetMachineAssignmentsOk() (*[]AddCatalogMachineAssignment, bool)`

GetMachineAssignmentsOk returns a tuple with the MachineAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAssignments

`func (o *AddRemotePcCatalog) SetMachineAssignments(v []AddCatalogMachineAssignment)`

SetMachineAssignments sets MachineAssignments field to given value.

### HasMachineAssignments

`func (o *AddRemotePcCatalog) HasMachineAssignments() bool`

HasMachineAssignments returns a boolean if a field has been set.

### GetCspCustomerId

`func (o *AddRemotePcCatalog) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *AddRemotePcCatalog) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *AddRemotePcCatalog) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *AddRemotePcCatalog) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### GetCspSiteId

`func (o *AddRemotePcCatalog) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *AddRemotePcCatalog) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *AddRemotePcCatalog) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *AddRemotePcCatalog) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


