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

// checks if the FtuConfigurationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FtuConfigurationModel{}

// FtuConfigurationModel struct for FtuConfigurationModel
type FtuConfigurationModel struct {
	// Indicates if FTU wizard has been completed by the customer
	IsFtuComplete *bool `json:"isFtuComplete,omitempty"`
	// The tye of on-prem connectivity
	Connectivity *OnPremConnectivity `json:"connectivity,omitempty"`
	// Indicates if at least one on-prem connection has been configured by the customer
	IsOnPremConnectionConfigured *bool `json:"isOnPremConnectionConfigured,omitempty"`
	// Indicates the prefered FtuAuthMethod method for Workspace users
	Authentication *AuthenticationMethod `json:"authentication,omitempty"`
	// Indicates whenver AzureQuickDeploy UI should be on/off
	CreatedWithAzureQuickDeploy *bool `json:"createdWithAzureQuickDeploy,omitempty"`
	StaleData                   *bool `json:"staleData,omitempty"`
}

// NewFtuConfigurationModelWithDefaults instantiates a new FtuConfigurationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFtuConfigurationModelWithDefaults() *FtuConfigurationModel {
	this := FtuConfigurationModel{}
	return &this
}

// GetIsFtuComplete returns the IsFtuComplete field value if set, zero value otherwise.
func (o *FtuConfigurationModel) GetIsFtuComplete() bool {
	if o == nil || IsNil(o.IsFtuComplete) {
		var ret bool
		return ret
	}
	return *o.IsFtuComplete
}

// GetIsFtuCompleteOk returns a tuple with the IsFtuComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtuConfigurationModel) GetIsFtuCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFtuComplete) {
		return nil, false
	}
	return o.IsFtuComplete, true
}

// SetIsFtuComplete gets a reference to the given bool and assigns it to the IsFtuComplete field.
func (o *FtuConfigurationModel) SetIsFtuComplete(v bool) {
	o.IsFtuComplete = &v
}

// GetConnectivity returns the Connectivity field value if set, zero value otherwise.
func (o *FtuConfigurationModel) GetConnectivity() OnPremConnectivity {
	if o == nil || IsNil(o.Connectivity) {
		var ret OnPremConnectivity
		return ret
	}
	return *o.Connectivity
}

// GetConnectivityOk returns a tuple with the Connectivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtuConfigurationModel) GetConnectivityOk() (*OnPremConnectivity, bool) {
	if o == nil || IsNil(o.Connectivity) {
		return nil, false
	}
	return o.Connectivity, true
}

// SetConnectivity gets a reference to the given OnPremConnectivity and assigns it to the Connectivity field.
func (o *FtuConfigurationModel) SetConnectivity(v OnPremConnectivity) {
	o.Connectivity = &v
}

// GetIsOnPremConnectionConfigured returns the IsOnPremConnectionConfigured field value if set, zero value otherwise.
func (o *FtuConfigurationModel) GetIsOnPremConnectionConfigured() bool {
	if o == nil || IsNil(o.IsOnPremConnectionConfigured) {
		var ret bool
		return ret
	}
	return *o.IsOnPremConnectionConfigured
}

// GetIsOnPremConnectionConfiguredOk returns a tuple with the IsOnPremConnectionConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtuConfigurationModel) GetIsOnPremConnectionConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOnPremConnectionConfigured) {
		return nil, false
	}
	return o.IsOnPremConnectionConfigured, true
}

// SetIsOnPremConnectionConfigured gets a reference to the given bool and assigns it to the IsOnPremConnectionConfigured field.
func (o *FtuConfigurationModel) SetIsOnPremConnectionConfigured(v bool) {
	o.IsOnPremConnectionConfigured = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *FtuConfigurationModel) GetAuthentication() AuthenticationMethod {
	if o == nil || IsNil(o.Authentication) {
		var ret AuthenticationMethod
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtuConfigurationModel) GetAuthenticationOk() (*AuthenticationMethod, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// SetAuthentication gets a reference to the given AuthenticationMethod and assigns it to the Authentication field.
func (o *FtuConfigurationModel) SetAuthentication(v AuthenticationMethod) {
	o.Authentication = &v
}

// GetCreatedWithAzureQuickDeploy returns the CreatedWithAzureQuickDeploy field value if set, zero value otherwise.
func (o *FtuConfigurationModel) GetCreatedWithAzureQuickDeploy() bool {
	if o == nil || IsNil(o.CreatedWithAzureQuickDeploy) {
		var ret bool
		return ret
	}
	return *o.CreatedWithAzureQuickDeploy
}

// GetCreatedWithAzureQuickDeployOk returns a tuple with the CreatedWithAzureQuickDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtuConfigurationModel) GetCreatedWithAzureQuickDeployOk() (*bool, bool) {
	if o == nil || IsNil(o.CreatedWithAzureQuickDeploy) {
		return nil, false
	}
	return o.CreatedWithAzureQuickDeploy, true
}

// SetCreatedWithAzureQuickDeploy gets a reference to the given bool and assigns it to the CreatedWithAzureQuickDeploy field.
func (o *FtuConfigurationModel) SetCreatedWithAzureQuickDeploy(v bool) {
	o.CreatedWithAzureQuickDeploy = &v
}

// GetStaleData returns the StaleData field value if set, zero value otherwise.
func (o *FtuConfigurationModel) GetStaleData() bool {
	if o == nil || IsNil(o.StaleData) {
		var ret bool
		return ret
	}
	return *o.StaleData
}

// GetStaleDataOk returns a tuple with the StaleData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtuConfigurationModel) GetStaleDataOk() (*bool, bool) {
	if o == nil || IsNil(o.StaleData) {
		return nil, false
	}
	return o.StaleData, true
}

// SetStaleData gets a reference to the given bool and assigns it to the StaleData field.
func (o *FtuConfigurationModel) SetStaleData(v bool) {
	o.StaleData = &v
}

func (o FtuConfigurationModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FtuConfigurationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsFtuComplete) {
		toSerialize["isFtuComplete"] = o.IsFtuComplete
	}
	if !IsNil(o.Connectivity) {
		toSerialize["connectivity"] = o.Connectivity
	}
	if !IsNil(o.IsOnPremConnectionConfigured) {
		toSerialize["isOnPremConnectionConfigured"] = o.IsOnPremConnectionConfigured
	}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if !IsNil(o.CreatedWithAzureQuickDeploy) {
		toSerialize["createdWithAzureQuickDeploy"] = o.CreatedWithAzureQuickDeploy
	}
	if !IsNil(o.StaleData) {
		toSerialize["staleData"] = o.StaleData
	}
	return toSerialize, nil
}

type NullableFtuConfigurationModel struct {
	value *FtuConfigurationModel
	isSet bool
}

func (v NullableFtuConfigurationModel) Get() *FtuConfigurationModel {
	return v.value
}

func (v *NullableFtuConfigurationModel) Set(val *FtuConfigurationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFtuConfigurationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFtuConfigurationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFtuConfigurationModel(val *FtuConfigurationModel) *NullableFtuConfigurationModel {
	return &NullableFtuConfigurationModel{value: val, isSet: true}
}

func (v NullableFtuConfigurationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFtuConfigurationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
