/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the MachineIdentityPoolResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineIdentityPoolResponseModel{}

// MachineIdentityPoolResponseModel Details about machine identity pool response model.
type MachineIdentityPoolResponseModel struct {
	// Machine identity pool id.
	Id *string `json:"Id,omitempty"`
	// Machine identity pool name.
	Name             NullableString           `json:"Name,omitempty"`
	NamingSchemeType *AccountNamingSchemeType `json:"NamingSchemeType,omitempty"`
	// Machine identity pool naming scheme.
	NamingScheme NullableString `json:"NamingScheme,omitempty"`
	// Machine identity pool start count.
	StartCount *int32 `json:"StartCount,omitempty"`
	// Machine identity pool custom active directory OU.
	CustomActiveDirectoryOU NullableString `json:"CustomActiveDirectoryOU,omitempty"`
	// Indicates whether use default OU
	UseDefaultOU    *bool                  `json:"UseDefaultOU,omitempty"`
	DefaultOUDomain *ADDomainResponseModel `json:"DefaultOUDomain,omitempty"`
	// The number of available accounts in the machine identity pool
	AvailableAccountsCount *int32 `json:"AvailableAccountsCount,omitempty"`
	// The number of accounts in use in the machine identity pool
	InUseAccountsCount *int32 `json:"InUseAccountsCount,omitempty"`
	// The number of tainted accounts in the machine identity pool
	TaintedAccountsCount *int32 `json:"TaintedAccountsCount,omitempty"`
	// The number of bad accounts in the machine identity pool
	ErrorAccountsCount *int32 `json:"ErrorAccountsCount,omitempty"`
	// Tenant id.
	TenantId NullableString `json:"TenantId,omitempty"`
	// Work group machines.
	WorkGroupMachines *bool `json:"WorkGroupMachines,omitempty"`
	// Identity type.
	IdentityType NullableString `json:"IdentityType,omitempty"`
	// Identity content.
	IdentityContent NullableString `json:"IdentityContent,omitempty"`
	// Azure AD security group name.
	AzureADSecurityGroupName NullableString `json:"AzureADSecurityGroupName,omitempty"`
	// Azure AD access token.
	AzureADAccessToken NullableString `json:"AzureADAccessToken,omitempty"`
	// Device management type.
	DeviceManagementType NullableString `json:"DeviceManagementType,omitempty"`
	// Azure AD tenant id.
	AzureADTenantId NullableString `json:"AzureADTenantId,omitempty"`
	// Service account Uids associated with this IdentityPool
	ServiceAccountUid []string `json:"ServiceAccountUid,omitempty"`
}

// NewMachineIdentityPoolResponseModel instantiates a new MachineIdentityPoolResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineIdentityPoolResponseModel() *MachineIdentityPoolResponseModel {
	this := MachineIdentityPoolResponseModel{}
	return &this
}

// NewMachineIdentityPoolResponseModelWithDefaults instantiates a new MachineIdentityPoolResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineIdentityPoolResponseModelWithDefaults() *MachineIdentityPoolResponseModel {
	this := MachineIdentityPoolResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MachineIdentityPoolResponseModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineIdentityPoolResponseModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MachineIdentityPoolResponseModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MachineIdentityPoolResponseModel) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *MachineIdentityPoolResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MachineIdentityPoolResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetNamingSchemeType returns the NamingSchemeType field value if set, zero value otherwise.
func (o *MachineIdentityPoolResponseModel) GetNamingSchemeType() AccountNamingSchemeType {
	if o == nil || IsNil(o.NamingSchemeType) {
		var ret AccountNamingSchemeType
		return ret
	}
	return *o.NamingSchemeType
}

// GetNamingSchemeTypeOk returns a tuple with the NamingSchemeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineIdentityPoolResponseModel) GetNamingSchemeTypeOk() (*AccountNamingSchemeType, bool) {
	if o == nil || IsNil(o.NamingSchemeType) {
		return nil, false
	}
	return o.NamingSchemeType, true
}

// HasNamingSchemeType returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasNamingSchemeType() bool {
	if o != nil && !IsNil(o.NamingSchemeType) {
		return true
	}

	return false
}

// SetNamingSchemeType gets a reference to the given AccountNamingSchemeType and assigns it to the NamingSchemeType field.
func (o *MachineIdentityPoolResponseModel) SetNamingSchemeType(v AccountNamingSchemeType) {
	o.NamingSchemeType = &v
}

// GetNamingScheme returns the NamingScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetNamingScheme() string {
	if o == nil || IsNil(o.NamingScheme.Get()) {
		var ret string
		return ret
	}
	return *o.NamingScheme.Get()
}

// GetNamingSchemeOk returns a tuple with the NamingScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetNamingSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NamingScheme.Get(), o.NamingScheme.IsSet()
}

// HasNamingScheme returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasNamingScheme() bool {
	if o != nil && o.NamingScheme.IsSet() {
		return true
	}

	return false
}

// SetNamingScheme gets a reference to the given NullableString and assigns it to the NamingScheme field.
func (o *MachineIdentityPoolResponseModel) SetNamingScheme(v string) {
	o.NamingScheme.Set(&v)
}

// SetNamingSchemeNil sets the value for NamingScheme to be an explicit nil
func (o *MachineIdentityPoolResponseModel) SetNamingSchemeNil() {
	o.NamingScheme.Set(nil)
}

// UnsetNamingScheme ensures that no value is present for NamingScheme, not even an explicit nil
func (o *MachineIdentityPoolResponseModel) UnsetNamingScheme() {
	o.NamingScheme.Unset()
}

// GetStartCount returns the StartCount field value if set, zero value otherwise.
func (o *MachineIdentityPoolResponseModel) GetStartCount() int32 {
	if o == nil || IsNil(o.StartCount) {
		var ret int32
		return ret
	}
	return *o.StartCount
}

// GetStartCountOk returns a tuple with the StartCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineIdentityPoolResponseModel) GetStartCountOk() (*int32, bool) {
	if o == nil || IsNil(o.StartCount) {
		return nil, false
	}
	return o.StartCount, true
}

// HasStartCount returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasStartCount() bool {
	if o != nil && !IsNil(o.StartCount) {
		return true
	}

	return false
}

// SetStartCount gets a reference to the given int32 and assigns it to the StartCount field.
func (o *MachineIdentityPoolResponseModel) SetStartCount(v int32) {
	o.StartCount = &v
}

// GetCustomActiveDirectoryOU returns the CustomActiveDirectoryOU field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetCustomActiveDirectoryOU() string {
	if o == nil || IsNil(o.CustomActiveDirectoryOU.Get()) {
		var ret string
		return ret
	}
	return *o.CustomActiveDirectoryOU.Get()
}

// GetCustomActiveDirectoryOUOk returns a tuple with the CustomActiveDirectoryOU field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetCustomActiveDirectoryOUOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomActiveDirectoryOU.Get(), o.CustomActiveDirectoryOU.IsSet()
}

// HasCustomActiveDirectoryOU returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasCustomActiveDirectoryOU() bool {
	if o != nil && o.CustomActiveDirectoryOU.IsSet() {
		return true
	}

	return false
}

// SetCustomActiveDirectoryOU gets a reference to the given NullableString and assigns it to the CustomActiveDirectoryOU field.
func (o *MachineIdentityPoolResponseModel) SetCustomActiveDirectoryOU(v string) {
	o.CustomActiveDirectoryOU.Set(&v)
}

// SetCustomActiveDirectoryOUNil sets the value for CustomActiveDirectoryOU to be an explicit nil
func (o *MachineIdentityPoolResponseModel) SetCustomActiveDirectoryOUNil() {
	o.CustomActiveDirectoryOU.Set(nil)
}

// UnsetCustomActiveDirectoryOU ensures that no value is present for CustomActiveDirectoryOU, not even an explicit nil
func (o *MachineIdentityPoolResponseModel) UnsetCustomActiveDirectoryOU() {
	o.CustomActiveDirectoryOU.Unset()
}

// GetUseDefaultOU returns the UseDefaultOU field value if set, zero value otherwise.
func (o *MachineIdentityPoolResponseModel) GetUseDefaultOU() bool {
	if o == nil || IsNil(o.UseDefaultOU) {
		var ret bool
		return ret
	}
	return *o.UseDefaultOU
}

// GetUseDefaultOUOk returns a tuple with the UseDefaultOU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineIdentityPoolResponseModel) GetUseDefaultOUOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDefaultOU) {
		return nil, false
	}
	return o.UseDefaultOU, true
}

// HasUseDefaultOU returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasUseDefaultOU() bool {
	if o != nil && !IsNil(o.UseDefaultOU) {
		return true
	}

	return false
}

// SetUseDefaultOU gets a reference to the given bool and assigns it to the UseDefaultOU field.
func (o *MachineIdentityPoolResponseModel) SetUseDefaultOU(v bool) {
	o.UseDefaultOU = &v
}

// GetDefaultOUDomain returns the DefaultOUDomain field value if set, zero value otherwise.
func (o *MachineIdentityPoolResponseModel) GetDefaultOUDomain() ADDomainResponseModel {
	if o == nil || IsNil(o.DefaultOUDomain) {
		var ret ADDomainResponseModel
		return ret
	}
	return *o.DefaultOUDomain
}

// GetDefaultOUDomainOk returns a tuple with the DefaultOUDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineIdentityPoolResponseModel) GetDefaultOUDomainOk() (*ADDomainResponseModel, bool) {
	if o == nil || IsNil(o.DefaultOUDomain) {
		return nil, false
	}
	return o.DefaultOUDomain, true
}

// HasDefaultOUDomain returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasDefaultOUDomain() bool {
	if o != nil && !IsNil(o.DefaultOUDomain) {
		return true
	}

	return false
}

// SetDefaultOUDomain gets a reference to the given ADDomainResponseModel and assigns it to the DefaultOUDomain field.
func (o *MachineIdentityPoolResponseModel) SetDefaultOUDomain(v ADDomainResponseModel) {
	o.DefaultOUDomain = &v
}

// GetAvailableAccountsCount returns the AvailableAccountsCount field value if set, zero value otherwise.
func (o *MachineIdentityPoolResponseModel) GetAvailableAccountsCount() int32 {
	if o == nil || IsNil(o.AvailableAccountsCount) {
		var ret int32
		return ret
	}
	return *o.AvailableAccountsCount
}

// GetAvailableAccountsCountOk returns a tuple with the AvailableAccountsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineIdentityPoolResponseModel) GetAvailableAccountsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailableAccountsCount) {
		return nil, false
	}
	return o.AvailableAccountsCount, true
}

// HasAvailableAccountsCount returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasAvailableAccountsCount() bool {
	if o != nil && !IsNil(o.AvailableAccountsCount) {
		return true
	}

	return false
}

// SetAvailableAccountsCount gets a reference to the given int32 and assigns it to the AvailableAccountsCount field.
func (o *MachineIdentityPoolResponseModel) SetAvailableAccountsCount(v int32) {
	o.AvailableAccountsCount = &v
}

// GetInUseAccountsCount returns the InUseAccountsCount field value if set, zero value otherwise.
func (o *MachineIdentityPoolResponseModel) GetInUseAccountsCount() int32 {
	if o == nil || IsNil(o.InUseAccountsCount) {
		var ret int32
		return ret
	}
	return *o.InUseAccountsCount
}

// GetInUseAccountsCountOk returns a tuple with the InUseAccountsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineIdentityPoolResponseModel) GetInUseAccountsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InUseAccountsCount) {
		return nil, false
	}
	return o.InUseAccountsCount, true
}

// HasInUseAccountsCount returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasInUseAccountsCount() bool {
	if o != nil && !IsNil(o.InUseAccountsCount) {
		return true
	}

	return false
}

// SetInUseAccountsCount gets a reference to the given int32 and assigns it to the InUseAccountsCount field.
func (o *MachineIdentityPoolResponseModel) SetInUseAccountsCount(v int32) {
	o.InUseAccountsCount = &v
}

// GetTaintedAccountsCount returns the TaintedAccountsCount field value if set, zero value otherwise.
func (o *MachineIdentityPoolResponseModel) GetTaintedAccountsCount() int32 {
	if o == nil || IsNil(o.TaintedAccountsCount) {
		var ret int32
		return ret
	}
	return *o.TaintedAccountsCount
}

// GetTaintedAccountsCountOk returns a tuple with the TaintedAccountsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineIdentityPoolResponseModel) GetTaintedAccountsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TaintedAccountsCount) {
		return nil, false
	}
	return o.TaintedAccountsCount, true
}

// HasTaintedAccountsCount returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasTaintedAccountsCount() bool {
	if o != nil && !IsNil(o.TaintedAccountsCount) {
		return true
	}

	return false
}

// SetTaintedAccountsCount gets a reference to the given int32 and assigns it to the TaintedAccountsCount field.
func (o *MachineIdentityPoolResponseModel) SetTaintedAccountsCount(v int32) {
	o.TaintedAccountsCount = &v
}

// GetErrorAccountsCount returns the ErrorAccountsCount field value if set, zero value otherwise.
func (o *MachineIdentityPoolResponseModel) GetErrorAccountsCount() int32 {
	if o == nil || IsNil(o.ErrorAccountsCount) {
		var ret int32
		return ret
	}
	return *o.ErrorAccountsCount
}

// GetErrorAccountsCountOk returns a tuple with the ErrorAccountsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineIdentityPoolResponseModel) GetErrorAccountsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorAccountsCount) {
		return nil, false
	}
	return o.ErrorAccountsCount, true
}

// HasErrorAccountsCount returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasErrorAccountsCount() bool {
	if o != nil && !IsNil(o.ErrorAccountsCount) {
		return true
	}

	return false
}

// SetErrorAccountsCount gets a reference to the given int32 and assigns it to the ErrorAccountsCount field.
func (o *MachineIdentityPoolResponseModel) SetErrorAccountsCount(v int32) {
	o.ErrorAccountsCount = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *MachineIdentityPoolResponseModel) SetTenantId(v string) {
	o.TenantId.Set(&v)
}

// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *MachineIdentityPoolResponseModel) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *MachineIdentityPoolResponseModel) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetWorkGroupMachines returns the WorkGroupMachines field value if set, zero value otherwise.
func (o *MachineIdentityPoolResponseModel) GetWorkGroupMachines() bool {
	if o == nil || IsNil(o.WorkGroupMachines) {
		var ret bool
		return ret
	}
	return *o.WorkGroupMachines
}

// GetWorkGroupMachinesOk returns a tuple with the WorkGroupMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineIdentityPoolResponseModel) GetWorkGroupMachinesOk() (*bool, bool) {
	if o == nil || IsNil(o.WorkGroupMachines) {
		return nil, false
	}
	return o.WorkGroupMachines, true
}

// HasWorkGroupMachines returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasWorkGroupMachines() bool {
	if o != nil && !IsNil(o.WorkGroupMachines) {
		return true
	}

	return false
}

// SetWorkGroupMachines gets a reference to the given bool and assigns it to the WorkGroupMachines field.
func (o *MachineIdentityPoolResponseModel) SetWorkGroupMachines(v bool) {
	o.WorkGroupMachines = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetIdentityType() string {
	if o == nil || IsNil(o.IdentityType.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityType.Get()
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetIdentityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityType.Get(), o.IdentityType.IsSet()
}

// HasIdentityType returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasIdentityType() bool {
	if o != nil && o.IdentityType.IsSet() {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given NullableString and assigns it to the IdentityType field.
func (o *MachineIdentityPoolResponseModel) SetIdentityType(v string) {
	o.IdentityType.Set(&v)
}

// SetIdentityTypeNil sets the value for IdentityType to be an explicit nil
func (o *MachineIdentityPoolResponseModel) SetIdentityTypeNil() {
	o.IdentityType.Set(nil)
}

// UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
func (o *MachineIdentityPoolResponseModel) UnsetIdentityType() {
	o.IdentityType.Unset()
}

// GetIdentityContent returns the IdentityContent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetIdentityContent() string {
	if o == nil || IsNil(o.IdentityContent.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityContent.Get()
}

// GetIdentityContentOk returns a tuple with the IdentityContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetIdentityContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityContent.Get(), o.IdentityContent.IsSet()
}

// HasIdentityContent returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasIdentityContent() bool {
	if o != nil && o.IdentityContent.IsSet() {
		return true
	}

	return false
}

// SetIdentityContent gets a reference to the given NullableString and assigns it to the IdentityContent field.
func (o *MachineIdentityPoolResponseModel) SetIdentityContent(v string) {
	o.IdentityContent.Set(&v)
}

// SetIdentityContentNil sets the value for IdentityContent to be an explicit nil
func (o *MachineIdentityPoolResponseModel) SetIdentityContentNil() {
	o.IdentityContent.Set(nil)
}

// UnsetIdentityContent ensures that no value is present for IdentityContent, not even an explicit nil
func (o *MachineIdentityPoolResponseModel) UnsetIdentityContent() {
	o.IdentityContent.Unset()
}

// GetAzureADSecurityGroupName returns the AzureADSecurityGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetAzureADSecurityGroupName() string {
	if o == nil || IsNil(o.AzureADSecurityGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.AzureADSecurityGroupName.Get()
}

// GetAzureADSecurityGroupNameOk returns a tuple with the AzureADSecurityGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetAzureADSecurityGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureADSecurityGroupName.Get(), o.AzureADSecurityGroupName.IsSet()
}

// HasAzureADSecurityGroupName returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasAzureADSecurityGroupName() bool {
	if o != nil && o.AzureADSecurityGroupName.IsSet() {
		return true
	}

	return false
}

// SetAzureADSecurityGroupName gets a reference to the given NullableString and assigns it to the AzureADSecurityGroupName field.
func (o *MachineIdentityPoolResponseModel) SetAzureADSecurityGroupName(v string) {
	o.AzureADSecurityGroupName.Set(&v)
}

// SetAzureADSecurityGroupNameNil sets the value for AzureADSecurityGroupName to be an explicit nil
func (o *MachineIdentityPoolResponseModel) SetAzureADSecurityGroupNameNil() {
	o.AzureADSecurityGroupName.Set(nil)
}

// UnsetAzureADSecurityGroupName ensures that no value is present for AzureADSecurityGroupName, not even an explicit nil
func (o *MachineIdentityPoolResponseModel) UnsetAzureADSecurityGroupName() {
	o.AzureADSecurityGroupName.Unset()
}

// GetAzureADAccessToken returns the AzureADAccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetAzureADAccessToken() string {
	if o == nil || IsNil(o.AzureADAccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AzureADAccessToken.Get()
}

// GetAzureADAccessTokenOk returns a tuple with the AzureADAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetAzureADAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureADAccessToken.Get(), o.AzureADAccessToken.IsSet()
}

// HasAzureADAccessToken returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasAzureADAccessToken() bool {
	if o != nil && o.AzureADAccessToken.IsSet() {
		return true
	}

	return false
}

// SetAzureADAccessToken gets a reference to the given NullableString and assigns it to the AzureADAccessToken field.
func (o *MachineIdentityPoolResponseModel) SetAzureADAccessToken(v string) {
	o.AzureADAccessToken.Set(&v)
}

// SetAzureADAccessTokenNil sets the value for AzureADAccessToken to be an explicit nil
func (o *MachineIdentityPoolResponseModel) SetAzureADAccessTokenNil() {
	o.AzureADAccessToken.Set(nil)
}

// UnsetAzureADAccessToken ensures that no value is present for AzureADAccessToken, not even an explicit nil
func (o *MachineIdentityPoolResponseModel) UnsetAzureADAccessToken() {
	o.AzureADAccessToken.Unset()
}

// GetDeviceManagementType returns the DeviceManagementType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetDeviceManagementType() string {
	if o == nil || IsNil(o.DeviceManagementType.Get()) {
		var ret string
		return ret
	}
	return *o.DeviceManagementType.Get()
}

// GetDeviceManagementTypeOk returns a tuple with the DeviceManagementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetDeviceManagementTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceManagementType.Get(), o.DeviceManagementType.IsSet()
}

// HasDeviceManagementType returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasDeviceManagementType() bool {
	if o != nil && o.DeviceManagementType.IsSet() {
		return true
	}

	return false
}

// SetDeviceManagementType gets a reference to the given NullableString and assigns it to the DeviceManagementType field.
func (o *MachineIdentityPoolResponseModel) SetDeviceManagementType(v string) {
	o.DeviceManagementType.Set(&v)
}

// SetDeviceManagementTypeNil sets the value for DeviceManagementType to be an explicit nil
func (o *MachineIdentityPoolResponseModel) SetDeviceManagementTypeNil() {
	o.DeviceManagementType.Set(nil)
}

// UnsetDeviceManagementType ensures that no value is present for DeviceManagementType, not even an explicit nil
func (o *MachineIdentityPoolResponseModel) UnsetDeviceManagementType() {
	o.DeviceManagementType.Unset()
}

// GetAzureADTenantId returns the AzureADTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetAzureADTenantId() string {
	if o == nil || IsNil(o.AzureADTenantId.Get()) {
		var ret string
		return ret
	}
	return *o.AzureADTenantId.Get()
}

// GetAzureADTenantIdOk returns a tuple with the AzureADTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetAzureADTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureADTenantId.Get(), o.AzureADTenantId.IsSet()
}

// HasAzureADTenantId returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasAzureADTenantId() bool {
	if o != nil && o.AzureADTenantId.IsSet() {
		return true
	}

	return false
}

// SetAzureADTenantId gets a reference to the given NullableString and assigns it to the AzureADTenantId field.
func (o *MachineIdentityPoolResponseModel) SetAzureADTenantId(v string) {
	o.AzureADTenantId.Set(&v)
}

// SetAzureADTenantIdNil sets the value for AzureADTenantId to be an explicit nil
func (o *MachineIdentityPoolResponseModel) SetAzureADTenantIdNil() {
	o.AzureADTenantId.Set(nil)
}

// UnsetAzureADTenantId ensures that no value is present for AzureADTenantId, not even an explicit nil
func (o *MachineIdentityPoolResponseModel) UnsetAzureADTenantId() {
	o.AzureADTenantId.Unset()
}

// GetServiceAccountUid returns the ServiceAccountUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdentityPoolResponseModel) GetServiceAccountUid() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ServiceAccountUid
}

// GetServiceAccountUidOk returns a tuple with the ServiceAccountUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdentityPoolResponseModel) GetServiceAccountUidOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceAccountUid) {
		return nil, false
	}
	return o.ServiceAccountUid, true
}

// HasServiceAccountUid returns a boolean if a field has been set.
func (o *MachineIdentityPoolResponseModel) HasServiceAccountUid() bool {
	if o != nil && IsNil(o.ServiceAccountUid) {
		return true
	}

	return false
}

// SetServiceAccountUid gets a reference to the given []string and assigns it to the ServiceAccountUid field.
func (o *MachineIdentityPoolResponseModel) SetServiceAccountUid(v []string) {
	o.ServiceAccountUid = v
}

func (o MachineIdentityPoolResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineIdentityPoolResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.NamingSchemeType) {
		toSerialize["NamingSchemeType"] = o.NamingSchemeType
	}
	if o.NamingScheme.IsSet() {
		toSerialize["NamingScheme"] = o.NamingScheme.Get()
	}
	if !IsNil(o.StartCount) {
		toSerialize["StartCount"] = o.StartCount
	}
	if o.CustomActiveDirectoryOU.IsSet() {
		toSerialize["CustomActiveDirectoryOU"] = o.CustomActiveDirectoryOU.Get()
	}
	if !IsNil(o.UseDefaultOU) {
		toSerialize["UseDefaultOU"] = o.UseDefaultOU
	}
	if !IsNil(o.DefaultOUDomain) {
		toSerialize["DefaultOUDomain"] = o.DefaultOUDomain
	}
	if !IsNil(o.AvailableAccountsCount) {
		toSerialize["AvailableAccountsCount"] = o.AvailableAccountsCount
	}
	if !IsNil(o.InUseAccountsCount) {
		toSerialize["InUseAccountsCount"] = o.InUseAccountsCount
	}
	if !IsNil(o.TaintedAccountsCount) {
		toSerialize["TaintedAccountsCount"] = o.TaintedAccountsCount
	}
	if !IsNil(o.ErrorAccountsCount) {
		toSerialize["ErrorAccountsCount"] = o.ErrorAccountsCount
	}
	if o.TenantId.IsSet() {
		toSerialize["TenantId"] = o.TenantId.Get()
	}
	if !IsNil(o.WorkGroupMachines) {
		toSerialize["WorkGroupMachines"] = o.WorkGroupMachines
	}
	if o.IdentityType.IsSet() {
		toSerialize["IdentityType"] = o.IdentityType.Get()
	}
	if o.IdentityContent.IsSet() {
		toSerialize["IdentityContent"] = o.IdentityContent.Get()
	}
	if o.AzureADSecurityGroupName.IsSet() {
		toSerialize["AzureADSecurityGroupName"] = o.AzureADSecurityGroupName.Get()
	}
	if o.AzureADAccessToken.IsSet() {
		toSerialize["AzureADAccessToken"] = o.AzureADAccessToken.Get()
	}
	if o.DeviceManagementType.IsSet() {
		toSerialize["DeviceManagementType"] = o.DeviceManagementType.Get()
	}
	if o.AzureADTenantId.IsSet() {
		toSerialize["AzureADTenantId"] = o.AzureADTenantId.Get()
	}
	if o.ServiceAccountUid != nil {
		toSerialize["ServiceAccountUid"] = o.ServiceAccountUid
	}
	return toSerialize, nil
}

type NullableMachineIdentityPoolResponseModel struct {
	value *MachineIdentityPoolResponseModel
	isSet bool
}

func (v NullableMachineIdentityPoolResponseModel) Get() *MachineIdentityPoolResponseModel {
	return v.value
}

func (v *NullableMachineIdentityPoolResponseModel) Set(val *MachineIdentityPoolResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineIdentityPoolResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineIdentityPoolResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineIdentityPoolResponseModel(val *MachineIdentityPoolResponseModel) *NullableMachineIdentityPoolResponseModel {
	return &NullableMachineIdentityPoolResponseModel{value: val, isSet: true}
}

func (v NullableMachineIdentityPoolResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineIdentityPoolResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
