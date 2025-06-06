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

// checks if the BackupStorageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupStorageInfo{}

// BackupStorageInfo Backup/Restore cloud storage definition
type BackupStorageInfo struct {
	BlobStorageType *BackupRestoreBlobStorage `json:"BlobStorageType,omitempty"`
	// File Storage folder
	FileStorageFolder NullableString `json:"FileStorageFolder,omitempty"`
	// Info 1 for storage
	Info1 NullableString `json:"Info1,omitempty"`
	// Info 2 for storage
	Info2 NullableString `json:"Info2,omitempty"`
	// Info 3 for storage
	Info3 NullableString `json:"Info3,omitempty"`
	// Info 4 for storage
	Info4 NullableString `json:"Info4,omitempty"`
	// Storage initialized by admin
	Initialized *bool `json:"Initialized,omitempty"`
}

// NewBackupStorageInfo instantiates a new BackupStorageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupStorageInfo() *BackupStorageInfo {
	this := BackupStorageInfo{}
	return &this
}

// NewBackupStorageInfoWithDefaults instantiates a new BackupStorageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupStorageInfoWithDefaults() *BackupStorageInfo {
	this := BackupStorageInfo{}
	return &this
}

// GetBlobStorageType returns the BlobStorageType field value if set, zero value otherwise.
func (o *BackupStorageInfo) GetBlobStorageType() BackupRestoreBlobStorage {
	if o == nil || IsNil(o.BlobStorageType) {
		var ret BackupRestoreBlobStorage
		return ret
	}
	return *o.BlobStorageType
}

// GetBlobStorageTypeOk returns a tuple with the BlobStorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStorageInfo) GetBlobStorageTypeOk() (*BackupRestoreBlobStorage, bool) {
	if o == nil || IsNil(o.BlobStorageType) {
		return nil, false
	}
	return o.BlobStorageType, true
}

// HasBlobStorageType returns a boolean if a field has been set.
func (o *BackupStorageInfo) HasBlobStorageType() bool {
	if o != nil && !IsNil(o.BlobStorageType) {
		return true
	}

	return false
}

// SetBlobStorageType gets a reference to the given BackupRestoreBlobStorage and assigns it to the BlobStorageType field.
func (o *BackupStorageInfo) SetBlobStorageType(v BackupRestoreBlobStorage) {
	o.BlobStorageType = &v
}

// GetFileStorageFolder returns the FileStorageFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupStorageInfo) GetFileStorageFolder() string {
	if o == nil || IsNil(o.FileStorageFolder.Get()) {
		var ret string
		return ret
	}
	return *o.FileStorageFolder.Get()
}

// GetFileStorageFolderOk returns a tuple with the FileStorageFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupStorageInfo) GetFileStorageFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileStorageFolder.Get(), o.FileStorageFolder.IsSet()
}

// HasFileStorageFolder returns a boolean if a field has been set.
func (o *BackupStorageInfo) HasFileStorageFolder() bool {
	if o != nil && o.FileStorageFolder.IsSet() {
		return true
	}

	return false
}

// SetFileStorageFolder gets a reference to the given NullableString and assigns it to the FileStorageFolder field.
func (o *BackupStorageInfo) SetFileStorageFolder(v string) {
	o.FileStorageFolder.Set(&v)
}

// SetFileStorageFolderNil sets the value for FileStorageFolder to be an explicit nil
func (o *BackupStorageInfo) SetFileStorageFolderNil() {
	o.FileStorageFolder.Set(nil)
}

// UnsetFileStorageFolder ensures that no value is present for FileStorageFolder, not even an explicit nil
func (o *BackupStorageInfo) UnsetFileStorageFolder() {
	o.FileStorageFolder.Unset()
}

// GetInfo1 returns the Info1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupStorageInfo) GetInfo1() string {
	if o == nil || IsNil(o.Info1.Get()) {
		var ret string
		return ret
	}
	return *o.Info1.Get()
}

// GetInfo1Ok returns a tuple with the Info1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupStorageInfo) GetInfo1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info1.Get(), o.Info1.IsSet()
}

// HasInfo1 returns a boolean if a field has been set.
func (o *BackupStorageInfo) HasInfo1() bool {
	if o != nil && o.Info1.IsSet() {
		return true
	}

	return false
}

// SetInfo1 gets a reference to the given NullableString and assigns it to the Info1 field.
func (o *BackupStorageInfo) SetInfo1(v string) {
	o.Info1.Set(&v)
}

// SetInfo1Nil sets the value for Info1 to be an explicit nil
func (o *BackupStorageInfo) SetInfo1Nil() {
	o.Info1.Set(nil)
}

// UnsetInfo1 ensures that no value is present for Info1, not even an explicit nil
func (o *BackupStorageInfo) UnsetInfo1() {
	o.Info1.Unset()
}

// GetInfo2 returns the Info2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupStorageInfo) GetInfo2() string {
	if o == nil || IsNil(o.Info2.Get()) {
		var ret string
		return ret
	}
	return *o.Info2.Get()
}

// GetInfo2Ok returns a tuple with the Info2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupStorageInfo) GetInfo2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info2.Get(), o.Info2.IsSet()
}

// HasInfo2 returns a boolean if a field has been set.
func (o *BackupStorageInfo) HasInfo2() bool {
	if o != nil && o.Info2.IsSet() {
		return true
	}

	return false
}

// SetInfo2 gets a reference to the given NullableString and assigns it to the Info2 field.
func (o *BackupStorageInfo) SetInfo2(v string) {
	o.Info2.Set(&v)
}

// SetInfo2Nil sets the value for Info2 to be an explicit nil
func (o *BackupStorageInfo) SetInfo2Nil() {
	o.Info2.Set(nil)
}

// UnsetInfo2 ensures that no value is present for Info2, not even an explicit nil
func (o *BackupStorageInfo) UnsetInfo2() {
	o.Info2.Unset()
}

// GetInfo3 returns the Info3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupStorageInfo) GetInfo3() string {
	if o == nil || IsNil(o.Info3.Get()) {
		var ret string
		return ret
	}
	return *o.Info3.Get()
}

// GetInfo3Ok returns a tuple with the Info3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupStorageInfo) GetInfo3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info3.Get(), o.Info3.IsSet()
}

// HasInfo3 returns a boolean if a field has been set.
func (o *BackupStorageInfo) HasInfo3() bool {
	if o != nil && o.Info3.IsSet() {
		return true
	}

	return false
}

// SetInfo3 gets a reference to the given NullableString and assigns it to the Info3 field.
func (o *BackupStorageInfo) SetInfo3(v string) {
	o.Info3.Set(&v)
}

// SetInfo3Nil sets the value for Info3 to be an explicit nil
func (o *BackupStorageInfo) SetInfo3Nil() {
	o.Info3.Set(nil)
}

// UnsetInfo3 ensures that no value is present for Info3, not even an explicit nil
func (o *BackupStorageInfo) UnsetInfo3() {
	o.Info3.Unset()
}

// GetInfo4 returns the Info4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupStorageInfo) GetInfo4() string {
	if o == nil || IsNil(o.Info4.Get()) {
		var ret string
		return ret
	}
	return *o.Info4.Get()
}

// GetInfo4Ok returns a tuple with the Info4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupStorageInfo) GetInfo4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info4.Get(), o.Info4.IsSet()
}

// HasInfo4 returns a boolean if a field has been set.
func (o *BackupStorageInfo) HasInfo4() bool {
	if o != nil && o.Info4.IsSet() {
		return true
	}

	return false
}

// SetInfo4 gets a reference to the given NullableString and assigns it to the Info4 field.
func (o *BackupStorageInfo) SetInfo4(v string) {
	o.Info4.Set(&v)
}

// SetInfo4Nil sets the value for Info4 to be an explicit nil
func (o *BackupStorageInfo) SetInfo4Nil() {
	o.Info4.Set(nil)
}

// UnsetInfo4 ensures that no value is present for Info4, not even an explicit nil
func (o *BackupStorageInfo) UnsetInfo4() {
	o.Info4.Unset()
}

// GetInitialized returns the Initialized field value if set, zero value otherwise.
func (o *BackupStorageInfo) GetInitialized() bool {
	if o == nil || IsNil(o.Initialized) {
		var ret bool
		return ret
	}
	return *o.Initialized
}

// GetInitializedOk returns a tuple with the Initialized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStorageInfo) GetInitializedOk() (*bool, bool) {
	if o == nil || IsNil(o.Initialized) {
		return nil, false
	}
	return o.Initialized, true
}

// HasInitialized returns a boolean if a field has been set.
func (o *BackupStorageInfo) HasInitialized() bool {
	if o != nil && !IsNil(o.Initialized) {
		return true
	}

	return false
}

// SetInitialized gets a reference to the given bool and assigns it to the Initialized field.
func (o *BackupStorageInfo) SetInitialized(v bool) {
	o.Initialized = &v
}

func (o BackupStorageInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupStorageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlobStorageType) {
		toSerialize["BlobStorageType"] = o.BlobStorageType
	}
	if o.FileStorageFolder.IsSet() {
		toSerialize["FileStorageFolder"] = o.FileStorageFolder.Get()
	}
	if o.Info1.IsSet() {
		toSerialize["Info1"] = o.Info1.Get()
	}
	if o.Info2.IsSet() {
		toSerialize["Info2"] = o.Info2.Get()
	}
	if o.Info3.IsSet() {
		toSerialize["Info3"] = o.Info3.Get()
	}
	if o.Info4.IsSet() {
		toSerialize["Info4"] = o.Info4.Get()
	}
	if !IsNil(o.Initialized) {
		toSerialize["Initialized"] = o.Initialized
	}
	return toSerialize, nil
}

type NullableBackupStorageInfo struct {
	value *BackupStorageInfo
	isSet bool
}

func (v NullableBackupStorageInfo) Get() *BackupStorageInfo {
	return v.value
}

func (v *NullableBackupStorageInfo) Set(val *BackupStorageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupStorageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupStorageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupStorageInfo(val *BackupStorageInfo) *NullableBackupStorageInfo {
	return &NullableBackupStorageInfo{value: val, isSet: true}
}

func (v NullableBackupStorageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupStorageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
