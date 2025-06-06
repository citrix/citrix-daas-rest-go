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

// checks if the AwsEdcAccountResourceSubnet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsEdcAccountResourceSubnet{}

// AwsEdcAccountResourceSubnet struct for AwsEdcAccountResourceSubnet
type AwsEdcAccountResourceSubnet struct {
	AwsEdcAccountResource
	// Subnet Id
	SubnetId NullableString `json:"subnetId,omitempty"`
	// Subnet Status  Enum values AVAILABLE, PENDING
	Status NullableAwsEdcSubnetStatus `json:"status,omitempty"`
	// Subnet Status
	SubnetArn NullableString `json:"subnetArn,omitempty"`
	// Subnet CIDR Block
	CidrBlock NullableString `json:"cidrBlock,omitempty"`
	// Subnet Vpc Id
	VpcId NullableString `json:"vpcId,omitempty"`
	// The AWS region where the subnet is located
	AwsRegion NullableString `json:"awsRegion,omitempty"`
	// Subnet Availability Zone Id
	AvailabilityZoneId NullableString `json:"availabilityZoneId,omitempty"`
	// Subnet Availability Zone
	AvailabilityZone NullableString `json:"availabilityZone,omitempty"`
	// Subnet Tags
	Tags []Tag `json:"tags,omitempty"`
	// The name of the subnet
	Name NullableString `json:"name,omitempty"`
	// Indicates if the subnet can be used for Managed Capacity deployments
	IsSupportedByAwsManagedCapacity *bool `json:"isSupportedByAwsManagedCapacity,omitempty"`
}

type _AwsEdcAccountResourceSubnet AwsEdcAccountResourceSubnet

// NewAwsEdcAccountResourceSubnet instantiates a new AwsEdcAccountResourceSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsEdcAccountResourceSubnet(resourceType AwsAccountResourceType, accountType AccountType) *AwsEdcAccountResourceSubnet {
	this := AwsEdcAccountResourceSubnet{}
	this.ResourceType = resourceType
	this.AccountType = accountType
	return &this
}

// NewAwsEdcAccountResourceSubnetWithDefaults instantiates a new AwsEdcAccountResourceSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsEdcAccountResourceSubnetWithDefaults() *AwsEdcAccountResourceSubnet {
	this := AwsEdcAccountResourceSubnet{}
	return &this
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceSubnet) GetSubnetId() string {
	if o == nil || IsNil(o.SubnetId.Get()) {
		var ret string
		return ret
	}
	return *o.SubnetId.Get()
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceSubnet) GetSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubnetId.Get(), o.SubnetId.IsSet()
}

// HasSubnetId returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasSubnetId() bool {
	if o != nil && o.SubnetId.IsSet() {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given NullableString and assigns it to the SubnetId field.
func (o *AwsEdcAccountResourceSubnet) SetSubnetId(v string) {
	o.SubnetId.Set(&v)
}

// SetSubnetIdNil sets the value for SubnetId to be an explicit nil
func (o *AwsEdcAccountResourceSubnet) SetSubnetIdNil() {
	o.SubnetId.Set(nil)
}

// UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
func (o *AwsEdcAccountResourceSubnet) UnsetSubnetId() {
	o.SubnetId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceSubnet) GetStatus() AwsEdcSubnetStatus {
	if o == nil || IsNil(o.Status.Get()) {
		var ret AwsEdcSubnetStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceSubnet) GetStatusOk() (*AwsEdcSubnetStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableAwsEdcSubnetStatus and assigns it to the Status field.
func (o *AwsEdcAccountResourceSubnet) SetStatus(v AwsEdcSubnetStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *AwsEdcAccountResourceSubnet) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AwsEdcAccountResourceSubnet) UnsetStatus() {
	o.Status.Unset()
}

// GetSubnetArn returns the SubnetArn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceSubnet) GetSubnetArn() string {
	if o == nil || IsNil(o.SubnetArn.Get()) {
		var ret string
		return ret
	}
	return *o.SubnetArn.Get()
}

// GetSubnetArnOk returns a tuple with the SubnetArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceSubnet) GetSubnetArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubnetArn.Get(), o.SubnetArn.IsSet()
}

// HasSubnetArn returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasSubnetArn() bool {
	if o != nil && o.SubnetArn.IsSet() {
		return true
	}

	return false
}

// SetSubnetArn gets a reference to the given NullableString and assigns it to the SubnetArn field.
func (o *AwsEdcAccountResourceSubnet) SetSubnetArn(v string) {
	o.SubnetArn.Set(&v)
}

// SetSubnetArnNil sets the value for SubnetArn to be an explicit nil
func (o *AwsEdcAccountResourceSubnet) SetSubnetArnNil() {
	o.SubnetArn.Set(nil)
}

// UnsetSubnetArn ensures that no value is present for SubnetArn, not even an explicit nil
func (o *AwsEdcAccountResourceSubnet) UnsetSubnetArn() {
	o.SubnetArn.Unset()
}

// GetCidrBlock returns the CidrBlock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceSubnet) GetCidrBlock() string {
	if o == nil || IsNil(o.CidrBlock.Get()) {
		var ret string
		return ret
	}
	return *o.CidrBlock.Get()
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceSubnet) GetCidrBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CidrBlock.Get(), o.CidrBlock.IsSet()
}

// HasCidrBlock returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasCidrBlock() bool {
	if o != nil && o.CidrBlock.IsSet() {
		return true
	}

	return false
}

// SetCidrBlock gets a reference to the given NullableString and assigns it to the CidrBlock field.
func (o *AwsEdcAccountResourceSubnet) SetCidrBlock(v string) {
	o.CidrBlock.Set(&v)
}

// SetCidrBlockNil sets the value for CidrBlock to be an explicit nil
func (o *AwsEdcAccountResourceSubnet) SetCidrBlockNil() {
	o.CidrBlock.Set(nil)
}

// UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
func (o *AwsEdcAccountResourceSubnet) UnsetCidrBlock() {
	o.CidrBlock.Unset()
}

// GetVpcId returns the VpcId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceSubnet) GetVpcId() string {
	if o == nil || IsNil(o.VpcId.Get()) {
		var ret string
		return ret
	}
	return *o.VpcId.Get()
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceSubnet) GetVpcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VpcId.Get(), o.VpcId.IsSet()
}

// HasVpcId returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasVpcId() bool {
	if o != nil && o.VpcId.IsSet() {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given NullableString and assigns it to the VpcId field.
func (o *AwsEdcAccountResourceSubnet) SetVpcId(v string) {
	o.VpcId.Set(&v)
}

// SetVpcIdNil sets the value for VpcId to be an explicit nil
func (o *AwsEdcAccountResourceSubnet) SetVpcIdNil() {
	o.VpcId.Set(nil)
}

// UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
func (o *AwsEdcAccountResourceSubnet) UnsetVpcId() {
	o.VpcId.Unset()
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceSubnet) GetAwsRegion() string {
	if o == nil || IsNil(o.AwsRegion.Get()) {
		var ret string
		return ret
	}
	return *o.AwsRegion.Get()
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceSubnet) GetAwsRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsRegion.Get(), o.AwsRegion.IsSet()
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasAwsRegion() bool {
	if o != nil && o.AwsRegion.IsSet() {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given NullableString and assigns it to the AwsRegion field.
func (o *AwsEdcAccountResourceSubnet) SetAwsRegion(v string) {
	o.AwsRegion.Set(&v)
}

// SetAwsRegionNil sets the value for AwsRegion to be an explicit nil
func (o *AwsEdcAccountResourceSubnet) SetAwsRegionNil() {
	o.AwsRegion.Set(nil)
}

// UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
func (o *AwsEdcAccountResourceSubnet) UnsetAwsRegion() {
	o.AwsRegion.Unset()
}

// GetAvailabilityZoneId returns the AvailabilityZoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceSubnet) GetAvailabilityZoneId() string {
	if o == nil || IsNil(o.AvailabilityZoneId.Get()) {
		var ret string
		return ret
	}
	return *o.AvailabilityZoneId.Get()
}

// GetAvailabilityZoneIdOk returns a tuple with the AvailabilityZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceSubnet) GetAvailabilityZoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailabilityZoneId.Get(), o.AvailabilityZoneId.IsSet()
}

// HasAvailabilityZoneId returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasAvailabilityZoneId() bool {
	if o != nil && o.AvailabilityZoneId.IsSet() {
		return true
	}

	return false
}

// SetAvailabilityZoneId gets a reference to the given NullableString and assigns it to the AvailabilityZoneId field.
func (o *AwsEdcAccountResourceSubnet) SetAvailabilityZoneId(v string) {
	o.AvailabilityZoneId.Set(&v)
}

// SetAvailabilityZoneIdNil sets the value for AvailabilityZoneId to be an explicit nil
func (o *AwsEdcAccountResourceSubnet) SetAvailabilityZoneIdNil() {
	o.AvailabilityZoneId.Set(nil)
}

// UnsetAvailabilityZoneId ensures that no value is present for AvailabilityZoneId, not even an explicit nil
func (o *AwsEdcAccountResourceSubnet) UnsetAvailabilityZoneId() {
	o.AvailabilityZoneId.Unset()
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceSubnet) GetAvailabilityZone() string {
	if o == nil || IsNil(o.AvailabilityZone.Get()) {
		var ret string
		return ret
	}
	return *o.AvailabilityZone.Get()
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceSubnet) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailabilityZone.Get(), o.AvailabilityZone.IsSet()
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasAvailabilityZone() bool {
	if o != nil && o.AvailabilityZone.IsSet() {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given NullableString and assigns it to the AvailabilityZone field.
func (o *AwsEdcAccountResourceSubnet) SetAvailabilityZone(v string) {
	o.AvailabilityZone.Set(&v)
}

// SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil
func (o *AwsEdcAccountResourceSubnet) SetAvailabilityZoneNil() {
	o.AvailabilityZone.Set(nil)
}

// UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
func (o *AwsEdcAccountResourceSubnet) UnsetAvailabilityZone() {
	o.AvailabilityZone.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceSubnet) GetTags() []Tag {
	if o == nil {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceSubnet) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *AwsEdcAccountResourceSubnet) SetTags(v []Tag) {
	o.Tags = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceSubnet) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceSubnet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AwsEdcAccountResourceSubnet) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *AwsEdcAccountResourceSubnet) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AwsEdcAccountResourceSubnet) UnsetName() {
	o.Name.Unset()
}

// GetIsSupportedByAwsManagedCapacity returns the IsSupportedByAwsManagedCapacity field value if set, zero value otherwise.
func (o *AwsEdcAccountResourceSubnet) GetIsSupportedByAwsManagedCapacity() bool {
	if o == nil || IsNil(o.IsSupportedByAwsManagedCapacity) {
		var ret bool
		return ret
	}
	return *o.IsSupportedByAwsManagedCapacity
}

// GetIsSupportedByAwsManagedCapacityOk returns a tuple with the IsSupportedByAwsManagedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsEdcAccountResourceSubnet) GetIsSupportedByAwsManagedCapacityOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSupportedByAwsManagedCapacity) {
		return nil, false
	}
	return o.IsSupportedByAwsManagedCapacity, true
}

// HasIsSupportedByAwsManagedCapacity returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceSubnet) HasIsSupportedByAwsManagedCapacity() bool {
	if o != nil && !IsNil(o.IsSupportedByAwsManagedCapacity) {
		return true
	}

	return false
}

// SetIsSupportedByAwsManagedCapacity gets a reference to the given bool and assigns it to the IsSupportedByAwsManagedCapacity field.
func (o *AwsEdcAccountResourceSubnet) SetIsSupportedByAwsManagedCapacity(v bool) {
	o.IsSupportedByAwsManagedCapacity = &v
}

func (o AwsEdcAccountResourceSubnet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsEdcAccountResourceSubnet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAwsEdcAccountResource, errAwsEdcAccountResource := json.Marshal(o.AwsEdcAccountResource)
	if errAwsEdcAccountResource != nil {
		return map[string]interface{}{}, errAwsEdcAccountResource
	}
	errAwsEdcAccountResource = json.Unmarshal([]byte(serializedAwsEdcAccountResource), &toSerialize)
	if errAwsEdcAccountResource != nil {
		return map[string]interface{}{}, errAwsEdcAccountResource
	}
	if o.SubnetId.IsSet() {
		toSerialize["subnetId"] = o.SubnetId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.SubnetArn.IsSet() {
		toSerialize["subnetArn"] = o.SubnetArn.Get()
	}
	if o.CidrBlock.IsSet() {
		toSerialize["cidrBlock"] = o.CidrBlock.Get()
	}
	if o.VpcId.IsSet() {
		toSerialize["vpcId"] = o.VpcId.Get()
	}
	if o.AwsRegion.IsSet() {
		toSerialize["awsRegion"] = o.AwsRegion.Get()
	}
	if o.AvailabilityZoneId.IsSet() {
		toSerialize["availabilityZoneId"] = o.AvailabilityZoneId.Get()
	}
	if o.AvailabilityZone.IsSet() {
		toSerialize["availabilityZone"] = o.AvailabilityZone.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.IsSupportedByAwsManagedCapacity) {
		toSerialize["isSupportedByAwsManagedCapacity"] = o.IsSupportedByAwsManagedCapacity
	}
	return toSerialize, nil
}

type NullableAwsEdcAccountResourceSubnet struct {
	value *AwsEdcAccountResourceSubnet
	isSet bool
}

func (v NullableAwsEdcAccountResourceSubnet) Get() *AwsEdcAccountResourceSubnet {
	return v.value
}

func (v *NullableAwsEdcAccountResourceSubnet) Set(val *AwsEdcAccountResourceSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcAccountResourceSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcAccountResourceSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcAccountResourceSubnet(val *AwsEdcAccountResourceSubnet) *NullableAwsEdcAccountResourceSubnet {
	return &NullableAwsEdcAccountResourceSubnet{value: val, isSet: true}
}

func (v NullableAwsEdcAccountResourceSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcAccountResourceSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
