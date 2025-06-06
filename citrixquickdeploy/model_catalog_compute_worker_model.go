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

// checks if the CatalogComputeWorkerModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogComputeWorkerModel{}

// CatalogComputeWorkerModel struct for CatalogComputeWorkerModel
type CatalogComputeWorkerModel struct {
	// Indicates if Premium Storage will be used
	UsePremiumStorage *bool `json:"usePremiumStorage,omitempty"`
	// Indicates type of Storage that will be used
	StorageType *CatalogCapacityStorageType `json:"storageType,omitempty"`
	// Indicates if the catalog VMs should be deployed with Azure HUB license
	UseAzureHUB *bool `json:"useAzureHUB,omitempty"`
	// Number of concurrent users per VM
	MaxUsersPerVM *int32 `json:"maxUsersPerVM,omitempty"`
	// Type of VM to create
	InstanceTypeId *string `json:"instanceTypeId,omitempty"`
	// Name of the Azure VM Instance to use for the catalog
	InstanceName *string `json:"instanceName,omitempty"`
	// Use managed disks for VMs in the catalog
	UseManagedDisks *bool `json:"useManagedDisks,omitempty"`
}

// NewCatalogComputeWorkerModelWithDefaults instantiates a new CatalogComputeWorkerModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogComputeWorkerModelWithDefaults() *CatalogComputeWorkerModel {
	this := CatalogComputeWorkerModel{}
	return &this
}

// GetUsePremiumStorage returns the UsePremiumStorage field value if set, zero value otherwise.
func (o *CatalogComputeWorkerModel) GetUsePremiumStorage() bool {
	if o == nil || IsNil(o.UsePremiumStorage) {
		var ret bool
		return ret
	}
	return *o.UsePremiumStorage
}

// GetUsePremiumStorageOk returns a tuple with the UsePremiumStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogComputeWorkerModel) GetUsePremiumStorageOk() (*bool, bool) {
	if o == nil || IsNil(o.UsePremiumStorage) {
		return nil, false
	}
	return o.UsePremiumStorage, true
}

// SetUsePremiumStorage gets a reference to the given bool and assigns it to the UsePremiumStorage field.
func (o *CatalogComputeWorkerModel) SetUsePremiumStorage(v bool) {
	o.UsePremiumStorage = &v
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *CatalogComputeWorkerModel) GetStorageType() CatalogCapacityStorageType {
	if o == nil || IsNil(o.StorageType) {
		var ret CatalogCapacityStorageType
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogComputeWorkerModel) GetStorageTypeOk() (*CatalogCapacityStorageType, bool) {
	if o == nil || IsNil(o.StorageType) {
		return nil, false
	}
	return o.StorageType, true
}

// SetStorageType gets a reference to the given CatalogCapacityStorageType and assigns it to the StorageType field.
func (o *CatalogComputeWorkerModel) SetStorageType(v CatalogCapacityStorageType) {
	o.StorageType = &v
}

// GetUseAzureHUB returns the UseAzureHUB field value if set, zero value otherwise.
func (o *CatalogComputeWorkerModel) GetUseAzureHUB() bool {
	if o == nil || IsNil(o.UseAzureHUB) {
		var ret bool
		return ret
	}
	return *o.UseAzureHUB
}

// GetUseAzureHUBOk returns a tuple with the UseAzureHUB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogComputeWorkerModel) GetUseAzureHUBOk() (*bool, bool) {
	if o == nil || IsNil(o.UseAzureHUB) {
		return nil, false
	}
	return o.UseAzureHUB, true
}

// SetUseAzureHUB gets a reference to the given bool and assigns it to the UseAzureHUB field.
func (o *CatalogComputeWorkerModel) SetUseAzureHUB(v bool) {
	o.UseAzureHUB = &v
}

// GetMaxUsersPerVM returns the MaxUsersPerVM field value if set, zero value otherwise.
func (o *CatalogComputeWorkerModel) GetMaxUsersPerVM() int32 {
	if o == nil || IsNil(o.MaxUsersPerVM) {
		var ret int32
		return ret
	}
	return *o.MaxUsersPerVM
}

// GetMaxUsersPerVMOk returns a tuple with the MaxUsersPerVM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogComputeWorkerModel) GetMaxUsersPerVMOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUsersPerVM) {
		return nil, false
	}
	return o.MaxUsersPerVM, true
}

// SetMaxUsersPerVM gets a reference to the given int32 and assigns it to the MaxUsersPerVM field.
func (o *CatalogComputeWorkerModel) SetMaxUsersPerVM(v int32) {
	o.MaxUsersPerVM = &v
}

// GetInstanceTypeId returns the InstanceTypeId field value if set, zero value otherwise.
func (o *CatalogComputeWorkerModel) GetInstanceTypeId() string {
	if o == nil || IsNil(o.InstanceTypeId) {
		var ret string
		return ret
	}
	return *o.InstanceTypeId
}

// GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogComputeWorkerModel) GetInstanceTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceTypeId) {
		return nil, false
	}
	return o.InstanceTypeId, true
}

// SetInstanceTypeId gets a reference to the given string and assigns it to the InstanceTypeId field.
func (o *CatalogComputeWorkerModel) SetInstanceTypeId(v string) {
	o.InstanceTypeId = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *CatalogComputeWorkerModel) GetInstanceName() string {
	if o == nil || IsNil(o.InstanceName) {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogComputeWorkerModel) GetInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceName) {
		return nil, false
	}
	return o.InstanceName, true
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *CatalogComputeWorkerModel) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetUseManagedDisks returns the UseManagedDisks field value if set, zero value otherwise.
func (o *CatalogComputeWorkerModel) GetUseManagedDisks() bool {
	if o == nil || IsNil(o.UseManagedDisks) {
		var ret bool
		return ret
	}
	return *o.UseManagedDisks
}

// GetUseManagedDisksOk returns a tuple with the UseManagedDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogComputeWorkerModel) GetUseManagedDisksOk() (*bool, bool) {
	if o == nil || IsNil(o.UseManagedDisks) {
		return nil, false
	}
	return o.UseManagedDisks, true
}

// SetUseManagedDisks gets a reference to the given bool and assigns it to the UseManagedDisks field.
func (o *CatalogComputeWorkerModel) SetUseManagedDisks(v bool) {
	o.UseManagedDisks = &v
}

func (o CatalogComputeWorkerModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogComputeWorkerModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UsePremiumStorage) {
		toSerialize["usePremiumStorage"] = o.UsePremiumStorage
	}
	if !IsNil(o.StorageType) {
		toSerialize["storageType"] = o.StorageType
	}
	if !IsNil(o.UseAzureHUB) {
		toSerialize["useAzureHUB"] = o.UseAzureHUB
	}
	if !IsNil(o.MaxUsersPerVM) {
		toSerialize["maxUsersPerVM"] = o.MaxUsersPerVM
	}
	if !IsNil(o.InstanceTypeId) {
		toSerialize["instanceTypeId"] = o.InstanceTypeId
	}
	if !IsNil(o.InstanceName) {
		toSerialize["instanceName"] = o.InstanceName
	}
	if !IsNil(o.UseManagedDisks) {
		toSerialize["useManagedDisks"] = o.UseManagedDisks
	}
	return toSerialize, nil
}

type NullableCatalogComputeWorkerModel struct {
	value *CatalogComputeWorkerModel
	isSet bool
}

func (v NullableCatalogComputeWorkerModel) Get() *CatalogComputeWorkerModel {
	return v.value
}

func (v *NullableCatalogComputeWorkerModel) Set(val *CatalogComputeWorkerModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogComputeWorkerModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogComputeWorkerModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogComputeWorkerModel(val *CatalogComputeWorkerModel) *NullableCatalogComputeWorkerModel {
	return &NullableCatalogComputeWorkerModel{value: val, isSet: true}
}

func (v NullableCatalogComputeWorkerModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogComputeWorkerModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
