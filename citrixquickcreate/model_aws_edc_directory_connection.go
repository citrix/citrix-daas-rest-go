/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
)

// checks if the AwsEdcDirectoryConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsEdcDirectoryConnection{}

// AwsEdcDirectoryConnection struct for AwsEdcDirectoryConnection
type AwsEdcDirectoryConnection struct {
	ResourceConnection
	// Directory Id
	DirectoryId NullableString `json:"directoryId,omitempty"`
	// Registration status of the directory
	RegistrationStatus *AwsEdcDirectoryRegistrationStatus `json:"registrationStatus,omitempty"`
	// Directory Name
	DirectoryName NullableString `json:"directoryName,omitempty"`
	// Directory Vpc
	VpcId NullableString `json:"vpcId,omitempty"`
	// First Subnet Id
	Subnet1Id NullableString `json:"subnet1Id,omitempty"`
	// Second Subnet Id
	Subnet2Id NullableString `json:"subnet2Id,omitempty"`
	// Tenancy of directory  Enum values SHARED, DEDICATED
	Tenancy NullableAwsEdcDirectoryTenancy `json:"tenancy,omitempty"`
	// Enable Work Docs
	EnableWorkDocs NullableBool `json:"enableWorkDocs,omitempty"`
	// Enable Local Administrator
	UserEnabledAsLocalAdministrator NullableBool `json:"userEnabledAsLocalAdministrator,omitempty"`
	// Enable Internet Access
	EnableInternetAccess NullableBool `json:"enableInternetAccess,omitempty"`
	// Enable Maintanance Mode
	EnableMaintananceMode NullableBool `json:"enableMaintananceMode,omitempty"`
	// Custom Security Group Id
	SecurityGroupId NullableString `json:"securityGroupId,omitempty"`
	// Custom Organizational Unit
	DefaultOU NullableString `json:"defaultOU,omitempty"`
	// Any error message
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
}

type _AwsEdcDirectoryConnection AwsEdcDirectoryConnection

// NewAwsEdcDirectoryConnection instantiates a new AwsEdcDirectoryConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsEdcDirectoryConnection(accountType AccountType, connectionId string, name string, citrixManaged bool) *AwsEdcDirectoryConnection {
	this := AwsEdcDirectoryConnection{}
	this.AccountType = accountType
	this.ConnectionId = connectionId
	this.Name = name
	this.CitrixManaged = citrixManaged
	return &this
}

// NewAwsEdcDirectoryConnectionWithDefaults instantiates a new AwsEdcDirectoryConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsEdcDirectoryConnectionWithDefaults() *AwsEdcDirectoryConnection {
	this := AwsEdcDirectoryConnection{}
	return &this
}

// GetDirectoryId returns the DirectoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetDirectoryId() string {
	if o == nil || IsNil(o.DirectoryId.Get()) {
		var ret string
		return ret
	}
	return *o.DirectoryId.Get()
}

// GetDirectoryIdOk returns a tuple with the DirectoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetDirectoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectoryId.Get(), o.DirectoryId.IsSet()
}

// HasDirectoryId returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasDirectoryId() bool {
	if o != nil && o.DirectoryId.IsSet() {
		return true
	}

	return false
}

// SetDirectoryId gets a reference to the given NullableString and assigns it to the DirectoryId field.
func (o *AwsEdcDirectoryConnection) SetDirectoryId(v string) {
	o.DirectoryId.Set(&v)
}

// SetDirectoryIdNil sets the value for DirectoryId to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetDirectoryIdNil() {
	o.DirectoryId.Set(nil)
}

// UnsetDirectoryId ensures that no value is present for DirectoryId, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetDirectoryId() {
	o.DirectoryId.Unset()
}

// GetRegistrationStatus returns the RegistrationStatus field value if set, zero value otherwise.
func (o *AwsEdcDirectoryConnection) GetRegistrationStatus() AwsEdcDirectoryRegistrationStatus {
	if o == nil || IsNil(o.RegistrationStatus) {
		var ret AwsEdcDirectoryRegistrationStatus
		return ret
	}
	return *o.RegistrationStatus
}

// GetRegistrationStatusOk returns a tuple with the RegistrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsEdcDirectoryConnection) GetRegistrationStatusOk() (*AwsEdcDirectoryRegistrationStatus, bool) {
	if o == nil || IsNil(o.RegistrationStatus) {
		return nil, false
	}
	return o.RegistrationStatus, true
}

// HasRegistrationStatus returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasRegistrationStatus() bool {
	if o != nil && !IsNil(o.RegistrationStatus) {
		return true
	}

	return false
}

// SetRegistrationStatus gets a reference to the given AwsEdcDirectoryRegistrationStatus and assigns it to the RegistrationStatus field.
func (o *AwsEdcDirectoryConnection) SetRegistrationStatus(v AwsEdcDirectoryRegistrationStatus) {
	o.RegistrationStatus = &v
}

// GetDirectoryName returns the DirectoryName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetDirectoryName() string {
	if o == nil || IsNil(o.DirectoryName.Get()) {
		var ret string
		return ret
	}
	return *o.DirectoryName.Get()
}

// GetDirectoryNameOk returns a tuple with the DirectoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetDirectoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectoryName.Get(), o.DirectoryName.IsSet()
}

// HasDirectoryName returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasDirectoryName() bool {
	if o != nil && o.DirectoryName.IsSet() {
		return true
	}

	return false
}

// SetDirectoryName gets a reference to the given NullableString and assigns it to the DirectoryName field.
func (o *AwsEdcDirectoryConnection) SetDirectoryName(v string) {
	o.DirectoryName.Set(&v)
}

// SetDirectoryNameNil sets the value for DirectoryName to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetDirectoryNameNil() {
	o.DirectoryName.Set(nil)
}

// UnsetDirectoryName ensures that no value is present for DirectoryName, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetDirectoryName() {
	o.DirectoryName.Unset()
}

// GetVpcId returns the VpcId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetVpcId() string {
	if o == nil || IsNil(o.VpcId.Get()) {
		var ret string
		return ret
	}
	return *o.VpcId.Get()
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetVpcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VpcId.Get(), o.VpcId.IsSet()
}

// HasVpcId returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasVpcId() bool {
	if o != nil && o.VpcId.IsSet() {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given NullableString and assigns it to the VpcId field.
func (o *AwsEdcDirectoryConnection) SetVpcId(v string) {
	o.VpcId.Set(&v)
}

// SetVpcIdNil sets the value for VpcId to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetVpcIdNil() {
	o.VpcId.Set(nil)
}

// UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetVpcId() {
	o.VpcId.Unset()
}

// GetSubnet1Id returns the Subnet1Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetSubnet1Id() string {
	if o == nil || IsNil(o.Subnet1Id.Get()) {
		var ret string
		return ret
	}
	return *o.Subnet1Id.Get()
}

// GetSubnet1IdOk returns a tuple with the Subnet1Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetSubnet1IdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subnet1Id.Get(), o.Subnet1Id.IsSet()
}

// HasSubnet1Id returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasSubnet1Id() bool {
	if o != nil && o.Subnet1Id.IsSet() {
		return true
	}

	return false
}

// SetSubnet1Id gets a reference to the given NullableString and assigns it to the Subnet1Id field.
func (o *AwsEdcDirectoryConnection) SetSubnet1Id(v string) {
	o.Subnet1Id.Set(&v)
}

// SetSubnet1IdNil sets the value for Subnet1Id to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetSubnet1IdNil() {
	o.Subnet1Id.Set(nil)
}

// UnsetSubnet1Id ensures that no value is present for Subnet1Id, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetSubnet1Id() {
	o.Subnet1Id.Unset()
}

// GetSubnet2Id returns the Subnet2Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetSubnet2Id() string {
	if o == nil || IsNil(o.Subnet2Id.Get()) {
		var ret string
		return ret
	}
	return *o.Subnet2Id.Get()
}

// GetSubnet2IdOk returns a tuple with the Subnet2Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetSubnet2IdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subnet2Id.Get(), o.Subnet2Id.IsSet()
}

// HasSubnet2Id returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasSubnet2Id() bool {
	if o != nil && o.Subnet2Id.IsSet() {
		return true
	}

	return false
}

// SetSubnet2Id gets a reference to the given NullableString and assigns it to the Subnet2Id field.
func (o *AwsEdcDirectoryConnection) SetSubnet2Id(v string) {
	o.Subnet2Id.Set(&v)
}

// SetSubnet2IdNil sets the value for Subnet2Id to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetSubnet2IdNil() {
	o.Subnet2Id.Set(nil)
}

// UnsetSubnet2Id ensures that no value is present for Subnet2Id, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetSubnet2Id() {
	o.Subnet2Id.Unset()
}

// GetTenancy returns the Tenancy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetTenancy() AwsEdcDirectoryTenancy {
	if o == nil || IsNil(o.Tenancy.Get()) {
		var ret AwsEdcDirectoryTenancy
		return ret
	}
	return *o.Tenancy.Get()
}

// GetTenancyOk returns a tuple with the Tenancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetTenancyOk() (*AwsEdcDirectoryTenancy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenancy.Get(), o.Tenancy.IsSet()
}

// HasTenancy returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasTenancy() bool {
	if o != nil && o.Tenancy.IsSet() {
		return true
	}

	return false
}

// SetTenancy gets a reference to the given NullableAwsEdcDirectoryTenancy and assigns it to the Tenancy field.
func (o *AwsEdcDirectoryConnection) SetTenancy(v AwsEdcDirectoryTenancy) {
	o.Tenancy.Set(&v)
}

// SetTenancyNil sets the value for Tenancy to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetTenancyNil() {
	o.Tenancy.Set(nil)
}

// UnsetTenancy ensures that no value is present for Tenancy, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetTenancy() {
	o.Tenancy.Unset()
}

// GetEnableWorkDocs returns the EnableWorkDocs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetEnableWorkDocs() bool {
	if o == nil || IsNil(o.EnableWorkDocs.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableWorkDocs.Get()
}

// GetEnableWorkDocsOk returns a tuple with the EnableWorkDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetEnableWorkDocsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableWorkDocs.Get(), o.EnableWorkDocs.IsSet()
}

// HasEnableWorkDocs returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasEnableWorkDocs() bool {
	if o != nil && o.EnableWorkDocs.IsSet() {
		return true
	}

	return false
}

// SetEnableWorkDocs gets a reference to the given NullableBool and assigns it to the EnableWorkDocs field.
func (o *AwsEdcDirectoryConnection) SetEnableWorkDocs(v bool) {
	o.EnableWorkDocs.Set(&v)
}

// SetEnableWorkDocsNil sets the value for EnableWorkDocs to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetEnableWorkDocsNil() {
	o.EnableWorkDocs.Set(nil)
}

// UnsetEnableWorkDocs ensures that no value is present for EnableWorkDocs, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetEnableWorkDocs() {
	o.EnableWorkDocs.Unset()
}

// GetUserEnabledAsLocalAdministrator returns the UserEnabledAsLocalAdministrator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetUserEnabledAsLocalAdministrator() bool {
	if o == nil || IsNil(o.UserEnabledAsLocalAdministrator.Get()) {
		var ret bool
		return ret
	}
	return *o.UserEnabledAsLocalAdministrator.Get()
}

// GetUserEnabledAsLocalAdministratorOk returns a tuple with the UserEnabledAsLocalAdministrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetUserEnabledAsLocalAdministratorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserEnabledAsLocalAdministrator.Get(), o.UserEnabledAsLocalAdministrator.IsSet()
}

// HasUserEnabledAsLocalAdministrator returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasUserEnabledAsLocalAdministrator() bool {
	if o != nil && o.UserEnabledAsLocalAdministrator.IsSet() {
		return true
	}

	return false
}

// SetUserEnabledAsLocalAdministrator gets a reference to the given NullableBool and assigns it to the UserEnabledAsLocalAdministrator field.
func (o *AwsEdcDirectoryConnection) SetUserEnabledAsLocalAdministrator(v bool) {
	o.UserEnabledAsLocalAdministrator.Set(&v)
}

// SetUserEnabledAsLocalAdministratorNil sets the value for UserEnabledAsLocalAdministrator to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetUserEnabledAsLocalAdministratorNil() {
	o.UserEnabledAsLocalAdministrator.Set(nil)
}

// UnsetUserEnabledAsLocalAdministrator ensures that no value is present for UserEnabledAsLocalAdministrator, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetUserEnabledAsLocalAdministrator() {
	o.UserEnabledAsLocalAdministrator.Unset()
}

// GetEnableInternetAccess returns the EnableInternetAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetEnableInternetAccess() bool {
	if o == nil || IsNil(o.EnableInternetAccess.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableInternetAccess.Get()
}

// GetEnableInternetAccessOk returns a tuple with the EnableInternetAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetEnableInternetAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableInternetAccess.Get(), o.EnableInternetAccess.IsSet()
}

// HasEnableInternetAccess returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasEnableInternetAccess() bool {
	if o != nil && o.EnableInternetAccess.IsSet() {
		return true
	}

	return false
}

// SetEnableInternetAccess gets a reference to the given NullableBool and assigns it to the EnableInternetAccess field.
func (o *AwsEdcDirectoryConnection) SetEnableInternetAccess(v bool) {
	o.EnableInternetAccess.Set(&v)
}

// SetEnableInternetAccessNil sets the value for EnableInternetAccess to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetEnableInternetAccessNil() {
	o.EnableInternetAccess.Set(nil)
}

// UnsetEnableInternetAccess ensures that no value is present for EnableInternetAccess, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetEnableInternetAccess() {
	o.EnableInternetAccess.Unset()
}

// GetEnableMaintananceMode returns the EnableMaintananceMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetEnableMaintananceMode() bool {
	if o == nil || IsNil(o.EnableMaintananceMode.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableMaintananceMode.Get()
}

// GetEnableMaintananceModeOk returns a tuple with the EnableMaintananceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetEnableMaintananceModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableMaintananceMode.Get(), o.EnableMaintananceMode.IsSet()
}

// HasEnableMaintananceMode returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasEnableMaintananceMode() bool {
	if o != nil && o.EnableMaintananceMode.IsSet() {
		return true
	}

	return false
}

// SetEnableMaintananceMode gets a reference to the given NullableBool and assigns it to the EnableMaintananceMode field.
func (o *AwsEdcDirectoryConnection) SetEnableMaintananceMode(v bool) {
	o.EnableMaintananceMode.Set(&v)
}

// SetEnableMaintananceModeNil sets the value for EnableMaintananceMode to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetEnableMaintananceModeNil() {
	o.EnableMaintananceMode.Set(nil)
}

// UnsetEnableMaintananceMode ensures that no value is present for EnableMaintananceMode, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetEnableMaintananceMode() {
	o.EnableMaintananceMode.Unset()
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetSecurityGroupId() string {
	if o == nil || IsNil(o.SecurityGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.SecurityGroupId.Get()
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecurityGroupId.Get(), o.SecurityGroupId.IsSet()
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasSecurityGroupId() bool {
	if o != nil && o.SecurityGroupId.IsSet() {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given NullableString and assigns it to the SecurityGroupId field.
func (o *AwsEdcDirectoryConnection) SetSecurityGroupId(v string) {
	o.SecurityGroupId.Set(&v)
}

// SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetSecurityGroupIdNil() {
	o.SecurityGroupId.Set(nil)
}

// UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetSecurityGroupId() {
	o.SecurityGroupId.Unset()
}

// GetDefaultOU returns the DefaultOU field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetDefaultOU() string {
	if o == nil || IsNil(o.DefaultOU.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultOU.Get()
}

// GetDefaultOUOk returns a tuple with the DefaultOU field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetDefaultOUOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultOU.Get(), o.DefaultOU.IsSet()
}

// HasDefaultOU returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasDefaultOU() bool {
	if o != nil && o.DefaultOU.IsSet() {
		return true
	}

	return false
}

// SetDefaultOU gets a reference to the given NullableString and assigns it to the DefaultOU field.
func (o *AwsEdcDirectoryConnection) SetDefaultOU(v string) {
	o.DefaultOU.Set(&v)
}

// SetDefaultOUNil sets the value for DefaultOU to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetDefaultOUNil() {
	o.DefaultOU.Set(nil)
}

// UnsetDefaultOU ensures that no value is present for DefaultOU, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetDefaultOU() {
	o.DefaultOU.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcDirectoryConnection) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcDirectoryConnection) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AwsEdcDirectoryConnection) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *AwsEdcDirectoryConnection) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *AwsEdcDirectoryConnection) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *AwsEdcDirectoryConnection) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

func (o AwsEdcDirectoryConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsEdcDirectoryConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedResourceConnection, errResourceConnection := json.Marshal(o.ResourceConnection)
	if errResourceConnection != nil {
		return map[string]interface{}{}, errResourceConnection
	}
	errResourceConnection = json.Unmarshal([]byte(serializedResourceConnection), &toSerialize)
	if errResourceConnection != nil {
		return map[string]interface{}{}, errResourceConnection
	}
	if o.DirectoryId.IsSet() {
		toSerialize["directoryId"] = o.DirectoryId.Get()
	}
	if !IsNil(o.RegistrationStatus) {
		toSerialize["registrationStatus"] = o.RegistrationStatus
	}
	if o.DirectoryName.IsSet() {
		toSerialize["directoryName"] = o.DirectoryName.Get()
	}
	if o.VpcId.IsSet() {
		toSerialize["vpcId"] = o.VpcId.Get()
	}
	if o.Subnet1Id.IsSet() {
		toSerialize["subnet1Id"] = o.Subnet1Id.Get()
	}
	if o.Subnet2Id.IsSet() {
		toSerialize["subnet2Id"] = o.Subnet2Id.Get()
	}
	if o.Tenancy.IsSet() {
		toSerialize["tenancy"] = o.Tenancy.Get()
	}
	if o.EnableWorkDocs.IsSet() {
		toSerialize["enableWorkDocs"] = o.EnableWorkDocs.Get()
	}
	if o.UserEnabledAsLocalAdministrator.IsSet() {
		toSerialize["userEnabledAsLocalAdministrator"] = o.UserEnabledAsLocalAdministrator.Get()
	}
	if o.EnableInternetAccess.IsSet() {
		toSerialize["enableInternetAccess"] = o.EnableInternetAccess.Get()
	}
	if o.EnableMaintananceMode.IsSet() {
		toSerialize["enableMaintananceMode"] = o.EnableMaintananceMode.Get()
	}
	if o.SecurityGroupId.IsSet() {
		toSerialize["securityGroupId"] = o.SecurityGroupId.Get()
	}
	if o.DefaultOU.IsSet() {
		toSerialize["defaultOU"] = o.DefaultOU.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	return toSerialize, nil
}

type NullableAwsEdcDirectoryConnection struct {
	value *AwsEdcDirectoryConnection
	isSet bool
}

func (v NullableAwsEdcDirectoryConnection) Get() *AwsEdcDirectoryConnection {
	return v.value
}

func (v *NullableAwsEdcDirectoryConnection) Set(val *AwsEdcDirectoryConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcDirectoryConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcDirectoryConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcDirectoryConnection(val *AwsEdcDirectoryConnection) *NullableAwsEdcDirectoryConnection {
	return &NullableAwsEdcDirectoryConnection{value: val, isSet: true}
}

func (v NullableAwsEdcDirectoryConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcDirectoryConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
