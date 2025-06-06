/*
Citrix Virtual App & Desktop Catalog Service 147.0.26651.57932

Catalog Service

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickdeploy

import (
	"encoding/json"
)

// checks if the AddCitrixManagedCatalogModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCitrixManagedCatalogModel{}

// AddCitrixManagedCatalogModel struct for AddCitrixManagedCatalogModel
type AddCitrixManagedCatalogModel struct {
	// The region where the catalog should be deployed, if it does not have any on-prem connectivity
	Region *string `json:"region,omitempty"`
	// Name of the catalog
	Name string `json:"name"`
	// Indicates if the catalog VDAs run a single session or multiple sessions
	Type *AddCatalogType `json:"type,omitempty"`
	// Indicates if the catalog will connected to the customers domain
	IsDomainJoined *bool `json:"isDomainJoined,omitempty"`
	// Indicates if catalogs that use statically allocated machines will have the disk contents persisted after shutdown
	PersistStaticAllocatedVmDisks *bool                     `json:"persistStaticAllocatedVmDisks,omitempty"`
	MachineNamingScheme           *MachineNamingSchemeModel `json:"machineNamingScheme,omitempty"`
	// Indicates if the machines in catalog will be Azure AD joined
	IsAzureAdJoined *bool `json:"isAzureAdJoined,omitempty"`
	// Specifies if intune enroll is enabled
	EnableIntuneEnroll *bool `json:"enableIntuneEnroll,omitempty"`
}

// NewAddCitrixManagedCatalogModelWithDefaults instantiates a new AddCitrixManagedCatalogModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCitrixManagedCatalogModelWithDefaults() *AddCitrixManagedCatalogModel {
	this := AddCitrixManagedCatalogModel{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AddCitrixManagedCatalogModel) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCitrixManagedCatalogModel) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AddCitrixManagedCatalogModel) SetRegion(v string) {
	o.Region = &v
}

// GetName returns the Name field value
func (o *AddCitrixManagedCatalogModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddCitrixManagedCatalogModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddCitrixManagedCatalogModel) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AddCitrixManagedCatalogModel) GetType() AddCatalogType {
	if o == nil || IsNil(o.Type) {
		var ret AddCatalogType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCitrixManagedCatalogModel) GetTypeOk() (*AddCatalogType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// SetType gets a reference to the given AddCatalogType and assigns it to the Type field.
func (o *AddCitrixManagedCatalogModel) SetType(v AddCatalogType) {
	o.Type = &v
}

// GetIsDomainJoined returns the IsDomainJoined field value if set, zero value otherwise.
func (o *AddCitrixManagedCatalogModel) GetIsDomainJoined() bool {
	if o == nil || IsNil(o.IsDomainJoined) {
		var ret bool
		return ret
	}
	return *o.IsDomainJoined
}

// GetIsDomainJoinedOk returns a tuple with the IsDomainJoined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCitrixManagedCatalogModel) GetIsDomainJoinedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDomainJoined) {
		return nil, false
	}
	return o.IsDomainJoined, true
}

// SetIsDomainJoined gets a reference to the given bool and assigns it to the IsDomainJoined field.
func (o *AddCitrixManagedCatalogModel) SetIsDomainJoined(v bool) {
	o.IsDomainJoined = &v
}

// GetPersistStaticAllocatedVmDisks returns the PersistStaticAllocatedVmDisks field value if set, zero value otherwise.
func (o *AddCitrixManagedCatalogModel) GetPersistStaticAllocatedVmDisks() bool {
	if o == nil || IsNil(o.PersistStaticAllocatedVmDisks) {
		var ret bool
		return ret
	}
	return *o.PersistStaticAllocatedVmDisks
}

// GetPersistStaticAllocatedVmDisksOk returns a tuple with the PersistStaticAllocatedVmDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCitrixManagedCatalogModel) GetPersistStaticAllocatedVmDisksOk() (*bool, bool) {
	if o == nil || IsNil(o.PersistStaticAllocatedVmDisks) {
		return nil, false
	}
	return o.PersistStaticAllocatedVmDisks, true
}

// SetPersistStaticAllocatedVmDisks gets a reference to the given bool and assigns it to the PersistStaticAllocatedVmDisks field.
func (o *AddCitrixManagedCatalogModel) SetPersistStaticAllocatedVmDisks(v bool) {
	o.PersistStaticAllocatedVmDisks = &v
}

// GetMachineNamingScheme returns the MachineNamingScheme field value if set, zero value otherwise.
func (o *AddCitrixManagedCatalogModel) GetMachineNamingScheme() MachineNamingSchemeModel {
	if o == nil || IsNil(o.MachineNamingScheme) {
		var ret MachineNamingSchemeModel
		return ret
	}
	return *o.MachineNamingScheme
}

// GetMachineNamingSchemeOk returns a tuple with the MachineNamingScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCitrixManagedCatalogModel) GetMachineNamingSchemeOk() (*MachineNamingSchemeModel, bool) {
	if o == nil || IsNil(o.MachineNamingScheme) {
		return nil, false
	}
	return o.MachineNamingScheme, true
}

// SetMachineNamingScheme gets a reference to the given MachineNamingSchemeModel and assigns it to the MachineNamingScheme field.
func (o *AddCitrixManagedCatalogModel) SetMachineNamingScheme(v MachineNamingSchemeModel) {
	o.MachineNamingScheme = &v
}

// GetIsAzureAdJoined returns the IsAzureAdJoined field value if set, zero value otherwise.
func (o *AddCitrixManagedCatalogModel) GetIsAzureAdJoined() bool {
	if o == nil || IsNil(o.IsAzureAdJoined) {
		var ret bool
		return ret
	}
	return *o.IsAzureAdJoined
}

// GetIsAzureAdJoinedOk returns a tuple with the IsAzureAdJoined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCitrixManagedCatalogModel) GetIsAzureAdJoinedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAzureAdJoined) {
		return nil, false
	}
	return o.IsAzureAdJoined, true
}

// SetIsAzureAdJoined gets a reference to the given bool and assigns it to the IsAzureAdJoined field.
func (o *AddCitrixManagedCatalogModel) SetIsAzureAdJoined(v bool) {
	o.IsAzureAdJoined = &v
}

// GetEnableIntuneEnroll returns the EnableIntuneEnroll field value if set, zero value otherwise.
func (o *AddCitrixManagedCatalogModel) GetEnableIntuneEnroll() bool {
	if o == nil || IsNil(o.EnableIntuneEnroll) {
		var ret bool
		return ret
	}
	return *o.EnableIntuneEnroll
}

// GetEnableIntuneEnrollOk returns a tuple with the EnableIntuneEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCitrixManagedCatalogModel) GetEnableIntuneEnrollOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableIntuneEnroll) {
		return nil, false
	}
	return o.EnableIntuneEnroll, true
}

// SetEnableIntuneEnroll gets a reference to the given bool and assigns it to the EnableIntuneEnroll field.
func (o *AddCitrixManagedCatalogModel) SetEnableIntuneEnroll(v bool) {
	o.EnableIntuneEnroll = &v
}

func (o AddCitrixManagedCatalogModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCitrixManagedCatalogModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsDomainJoined) {
		toSerialize["isDomainJoined"] = o.IsDomainJoined
	}
	if !IsNil(o.PersistStaticAllocatedVmDisks) {
		toSerialize["persistStaticAllocatedVmDisks"] = o.PersistStaticAllocatedVmDisks
	}
	if !IsNil(o.MachineNamingScheme) {
		toSerialize["machineNamingScheme"] = o.MachineNamingScheme
	}
	if !IsNil(o.IsAzureAdJoined) {
		toSerialize["isAzureAdJoined"] = o.IsAzureAdJoined
	}
	if !IsNil(o.EnableIntuneEnroll) {
		toSerialize["enableIntuneEnroll"] = o.EnableIntuneEnroll
	}
	return toSerialize, nil
}

type NullableAddCitrixManagedCatalogModel struct {
	value *AddCitrixManagedCatalogModel
	isSet bool
}

func (v NullableAddCitrixManagedCatalogModel) Get() *AddCitrixManagedCatalogModel {
	return v.value
}

func (v *NullableAddCitrixManagedCatalogModel) Set(val *AddCitrixManagedCatalogModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCitrixManagedCatalogModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCitrixManagedCatalogModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCitrixManagedCatalogModel(val *AddCitrixManagedCatalogModel) *NullableAddCitrixManagedCatalogModel {
	return &NullableAddCitrixManagedCatalogModel{value: val, isSet: true}
}

func (v NullableAddCitrixManagedCatalogModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCitrixManagedCatalogModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
