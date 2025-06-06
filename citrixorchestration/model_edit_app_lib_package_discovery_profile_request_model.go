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

// checks if the EditAppLibPackageDiscoveryProfileRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditAppLibPackageDiscoveryProfileRequestModel{}

// EditAppLibPackageDiscoveryProfileRequestModel Request object for editing applib package discovery profile.
type EditAppLibPackageDiscoveryProfileRequestModel struct {
	// The name of the package discovery profile
	Name NullableString `json:"Name,omitempty" validate:"regexp=(.*)*"`
	// The UID of the DesktopGroup from which the broker will select a VDA to run the discovery.
	DesktopGroupUid NullableInt32 `json:"DesktopGroupUid,omitempty"`
	// The root directory where the discovery will run
	Path NullableString `json:"Path,omitempty"`
	// A value indicating whether or not to search the child directory tree.
	Recurse NullableBool `json:"Recurse,omitempty"`
	// The url of the App-V Management server that packages will be discovered from.
	ManagementServer NullableString `json:"ManagementServer,omitempty"`
	// The url of the App-V Publishing server that packages will be discovered from.
	PublishingServer NullableString `json:"PublishingServer,omitempty"`
	// The user name of the App-V Server Administrator.
	Username NullableString `json:"Username,omitempty"`
	// The password of the App-V Server Administrator.
	Password       NullableString          `json:"Password,omitempty"`
	PasswordFormat *IdentityPasswordFormat `json:"PasswordFormat,omitempty"`
	// A value indicating whether or not to search should Discover App-V Packages.
	DiscoverAppV NullableBool `json:"DiscoverAppV,omitempty"`
	// A value indicating whether or not to search should Discover MSIX Packages.
	DiscoverMsix NullableBool `json:"DiscoverMsix,omitempty"`
	// A value indicating whether or not to search should discover AppAttach Packages.
	DiscoverAppAttach NullableBool `json:"DiscoverAppAttach,omitempty"`
	// A value indicating whether or not to search should discover FlexApp Packages.
	DiscoverFlexApp NullableBool `json:"DiscoverFlexApp,omitempty"`
	// A value indicating whether or not to search should discover ElasticApp Layers packages.
	DiscoverElasticAppLayers NullableBool `json:"DiscoverElasticAppLayers,omitempty"`
	// A value indicating whether or not the discovery should run automatically.
	AutomateDiscovery   NullableBool         `json:"AutomateDiscovery,omitempty"`
	AutoDiscoveryPeriod *AutoDiscoveryPeriod `json:"AutoDiscoveryPeriod,omitempty"`
	// A value indicating how often to trigger the discovery.
	AutoDiscoveryCadence NullableInt32 `json:"AutoDiscoveryCadence,omitempty"`
	// A value indicating whether to clean up absent packages.
	CleanupAbsentPackages NullableBool `json:"CleanupAbsentPackages,omitempty"`
}

// NewEditAppLibPackageDiscoveryProfileRequestModel instantiates a new EditAppLibPackageDiscoveryProfileRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditAppLibPackageDiscoveryProfileRequestModel() *EditAppLibPackageDiscoveryProfileRequestModel {
	this := EditAppLibPackageDiscoveryProfileRequestModel{}
	return &this
}

// NewEditAppLibPackageDiscoveryProfileRequestModelWithDefaults instantiates a new EditAppLibPackageDiscoveryProfileRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditAppLibPackageDiscoveryProfileRequestModelWithDefaults() *EditAppLibPackageDiscoveryProfileRequestModel {
	this := EditAppLibPackageDiscoveryProfileRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetName() {
	o.Name.Unset()
}

// GetDesktopGroupUid returns the DesktopGroupUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDesktopGroupUid() int32 {
	if o == nil || IsNil(o.DesktopGroupUid.Get()) {
		var ret int32
		return ret
	}
	return *o.DesktopGroupUid.Get()
}

// GetDesktopGroupUidOk returns a tuple with the DesktopGroupUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDesktopGroupUidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DesktopGroupUid.Get(), o.DesktopGroupUid.IsSet()
}

// HasDesktopGroupUid returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasDesktopGroupUid() bool {
	if o != nil && o.DesktopGroupUid.IsSet() {
		return true
	}

	return false
}

// SetDesktopGroupUid gets a reference to the given NullableInt32 and assigns it to the DesktopGroupUid field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDesktopGroupUid(v int32) {
	o.DesktopGroupUid.Set(&v)
}

// SetDesktopGroupUidNil sets the value for DesktopGroupUid to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDesktopGroupUidNil() {
	o.DesktopGroupUid.Set(nil)
}

// UnsetDesktopGroupUid ensures that no value is present for DesktopGroupUid, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetDesktopGroupUid() {
	o.DesktopGroupUid.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPath(v string) {
	o.Path.Set(&v)
}

// SetPathNil sets the value for Path to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetPath() {
	o.Path.Unset()
}

// GetRecurse returns the Recurse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetRecurse() bool {
	if o == nil || IsNil(o.Recurse.Get()) {
		var ret bool
		return ret
	}
	return *o.Recurse.Get()
}

// GetRecurseOk returns a tuple with the Recurse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetRecurseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recurse.Get(), o.Recurse.IsSet()
}

// HasRecurse returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasRecurse() bool {
	if o != nil && o.Recurse.IsSet() {
		return true
	}

	return false
}

// SetRecurse gets a reference to the given NullableBool and assigns it to the Recurse field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetRecurse(v bool) {
	o.Recurse.Set(&v)
}

// SetRecurseNil sets the value for Recurse to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetRecurseNil() {
	o.Recurse.Set(nil)
}

// UnsetRecurse ensures that no value is present for Recurse, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetRecurse() {
	o.Recurse.Unset()
}

// GetManagementServer returns the ManagementServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetManagementServer() string {
	if o == nil || IsNil(o.ManagementServer.Get()) {
		var ret string
		return ret
	}
	return *o.ManagementServer.Get()
}

// GetManagementServerOk returns a tuple with the ManagementServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetManagementServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagementServer.Get(), o.ManagementServer.IsSet()
}

// HasManagementServer returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasManagementServer() bool {
	if o != nil && o.ManagementServer.IsSet() {
		return true
	}

	return false
}

// SetManagementServer gets a reference to the given NullableString and assigns it to the ManagementServer field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetManagementServer(v string) {
	o.ManagementServer.Set(&v)
}

// SetManagementServerNil sets the value for ManagementServer to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetManagementServerNil() {
	o.ManagementServer.Set(nil)
}

// UnsetManagementServer ensures that no value is present for ManagementServer, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetManagementServer() {
	o.ManagementServer.Unset()
}

// GetPublishingServer returns the PublishingServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPublishingServer() string {
	if o == nil || IsNil(o.PublishingServer.Get()) {
		var ret string
		return ret
	}
	return *o.PublishingServer.Get()
}

// GetPublishingServerOk returns a tuple with the PublishingServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPublishingServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublishingServer.Get(), o.PublishingServer.IsSet()
}

// HasPublishingServer returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasPublishingServer() bool {
	if o != nil && o.PublishingServer.IsSet() {
		return true
	}

	return false
}

// SetPublishingServer gets a reference to the given NullableString and assigns it to the PublishingServer field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPublishingServer(v string) {
	o.PublishingServer.Set(&v)
}

// SetPublishingServerNil sets the value for PublishingServer to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPublishingServerNil() {
	o.PublishingServer.Set(nil)
}

// UnsetPublishingServer ensures that no value is present for PublishingServer, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetPublishingServer() {
	o.PublishingServer.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetUsername(v string) {
	o.Username.Set(&v)
}

// SetUsernameNil sets the value for Username to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordFormat returns the PasswordFormat field value if set, zero value otherwise.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPasswordFormat() IdentityPasswordFormat {
	if o == nil || IsNil(o.PasswordFormat) {
		var ret IdentityPasswordFormat
		return ret
	}
	return *o.PasswordFormat
}

// GetPasswordFormatOk returns a tuple with the PasswordFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool) {
	if o == nil || IsNil(o.PasswordFormat) {
		return nil, false
	}
	return o.PasswordFormat, true
}

// HasPasswordFormat returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasPasswordFormat() bool {
	if o != nil && !IsNil(o.PasswordFormat) {
		return true
	}

	return false
}

// SetPasswordFormat gets a reference to the given IdentityPasswordFormat and assigns it to the PasswordFormat field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetPasswordFormat(v IdentityPasswordFormat) {
	o.PasswordFormat = &v
}

// GetDiscoverAppV returns the DiscoverAppV field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppV() bool {
	if o == nil || IsNil(o.DiscoverAppV.Get()) {
		var ret bool
		return ret
	}
	return *o.DiscoverAppV.Get()
}

// GetDiscoverAppVOk returns a tuple with the DiscoverAppV field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppVOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscoverAppV.Get(), o.DiscoverAppV.IsSet()
}

// HasDiscoverAppV returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasDiscoverAppV() bool {
	if o != nil && o.DiscoverAppV.IsSet() {
		return true
	}

	return false
}

// SetDiscoverAppV gets a reference to the given NullableBool and assigns it to the DiscoverAppV field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppV(v bool) {
	o.DiscoverAppV.Set(&v)
}

// SetDiscoverAppVNil sets the value for DiscoverAppV to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppVNil() {
	o.DiscoverAppV.Set(nil)
}

// UnsetDiscoverAppV ensures that no value is present for DiscoverAppV, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverAppV() {
	o.DiscoverAppV.Unset()
}

// GetDiscoverMsix returns the DiscoverMsix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverMsix() bool {
	if o == nil || IsNil(o.DiscoverMsix.Get()) {
		var ret bool
		return ret
	}
	return *o.DiscoverMsix.Get()
}

// GetDiscoverMsixOk returns a tuple with the DiscoverMsix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverMsixOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscoverMsix.Get(), o.DiscoverMsix.IsSet()
}

// HasDiscoverMsix returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasDiscoverMsix() bool {
	if o != nil && o.DiscoverMsix.IsSet() {
		return true
	}

	return false
}

// SetDiscoverMsix gets a reference to the given NullableBool and assigns it to the DiscoverMsix field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverMsix(v bool) {
	o.DiscoverMsix.Set(&v)
}

// SetDiscoverMsixNil sets the value for DiscoverMsix to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverMsixNil() {
	o.DiscoverMsix.Set(nil)
}

// UnsetDiscoverMsix ensures that no value is present for DiscoverMsix, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverMsix() {
	o.DiscoverMsix.Unset()
}

// GetDiscoverAppAttach returns the DiscoverAppAttach field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppAttach() bool {
	if o == nil || IsNil(o.DiscoverAppAttach.Get()) {
		var ret bool
		return ret
	}
	return *o.DiscoverAppAttach.Get()
}

// GetDiscoverAppAttachOk returns a tuple with the DiscoverAppAttach field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverAppAttachOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscoverAppAttach.Get(), o.DiscoverAppAttach.IsSet()
}

// HasDiscoverAppAttach returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasDiscoverAppAttach() bool {
	if o != nil && o.DiscoverAppAttach.IsSet() {
		return true
	}

	return false
}

// SetDiscoverAppAttach gets a reference to the given NullableBool and assigns it to the DiscoverAppAttach field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppAttach(v bool) {
	o.DiscoverAppAttach.Set(&v)
}

// SetDiscoverAppAttachNil sets the value for DiscoverAppAttach to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverAppAttachNil() {
	o.DiscoverAppAttach.Set(nil)
}

// UnsetDiscoverAppAttach ensures that no value is present for DiscoverAppAttach, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverAppAttach() {
	o.DiscoverAppAttach.Unset()
}

// GetDiscoverFlexApp returns the DiscoverFlexApp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverFlexApp() bool {
	if o == nil || IsNil(o.DiscoverFlexApp.Get()) {
		var ret bool
		return ret
	}
	return *o.DiscoverFlexApp.Get()
}

// GetDiscoverFlexAppOk returns a tuple with the DiscoverFlexApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverFlexAppOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscoverFlexApp.Get(), o.DiscoverFlexApp.IsSet()
}

// HasDiscoverFlexApp returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasDiscoverFlexApp() bool {
	if o != nil && o.DiscoverFlexApp.IsSet() {
		return true
	}

	return false
}

// SetDiscoverFlexApp gets a reference to the given NullableBool and assigns it to the DiscoverFlexApp field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverFlexApp(v bool) {
	o.DiscoverFlexApp.Set(&v)
}

// SetDiscoverFlexAppNil sets the value for DiscoverFlexApp to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverFlexAppNil() {
	o.DiscoverFlexApp.Set(nil)
}

// UnsetDiscoverFlexApp ensures that no value is present for DiscoverFlexApp, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverFlexApp() {
	o.DiscoverFlexApp.Unset()
}

// GetDiscoverElasticAppLayers returns the DiscoverElasticAppLayers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverElasticAppLayers() bool {
	if o == nil || IsNil(o.DiscoverElasticAppLayers.Get()) {
		var ret bool
		return ret
	}
	return *o.DiscoverElasticAppLayers.Get()
}

// GetDiscoverElasticAppLayersOk returns a tuple with the DiscoverElasticAppLayers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetDiscoverElasticAppLayersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscoverElasticAppLayers.Get(), o.DiscoverElasticAppLayers.IsSet()
}

// HasDiscoverElasticAppLayers returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasDiscoverElasticAppLayers() bool {
	if o != nil && o.DiscoverElasticAppLayers.IsSet() {
		return true
	}

	return false
}

// SetDiscoverElasticAppLayers gets a reference to the given NullableBool and assigns it to the DiscoverElasticAppLayers field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverElasticAppLayers(v bool) {
	o.DiscoverElasticAppLayers.Set(&v)
}

// SetDiscoverElasticAppLayersNil sets the value for DiscoverElasticAppLayers to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetDiscoverElasticAppLayersNil() {
	o.DiscoverElasticAppLayers.Set(nil)
}

// UnsetDiscoverElasticAppLayers ensures that no value is present for DiscoverElasticAppLayers, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetDiscoverElasticAppLayers() {
	o.DiscoverElasticAppLayers.Unset()
}

// GetAutomateDiscovery returns the AutomateDiscovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutomateDiscovery() bool {
	if o == nil || IsNil(o.AutomateDiscovery.Get()) {
		var ret bool
		return ret
	}
	return *o.AutomateDiscovery.Get()
}

// GetAutomateDiscoveryOk returns a tuple with the AutomateDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutomateDiscoveryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutomateDiscovery.Get(), o.AutomateDiscovery.IsSet()
}

// HasAutomateDiscovery returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasAutomateDiscovery() bool {
	if o != nil && o.AutomateDiscovery.IsSet() {
		return true
	}

	return false
}

// SetAutomateDiscovery gets a reference to the given NullableBool and assigns it to the AutomateDiscovery field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetAutomateDiscovery(v bool) {
	o.AutomateDiscovery.Set(&v)
}

// SetAutomateDiscoveryNil sets the value for AutomateDiscovery to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetAutomateDiscoveryNil() {
	o.AutomateDiscovery.Set(nil)
}

// UnsetAutomateDiscovery ensures that no value is present for AutomateDiscovery, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetAutomateDiscovery() {
	o.AutomateDiscovery.Unset()
}

// GetAutoDiscoveryPeriod returns the AutoDiscoveryPeriod field value if set, zero value otherwise.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryPeriod() AutoDiscoveryPeriod {
	if o == nil || IsNil(o.AutoDiscoveryPeriod) {
		var ret AutoDiscoveryPeriod
		return ret
	}
	return *o.AutoDiscoveryPeriod
}

// GetAutoDiscoveryPeriodOk returns a tuple with the AutoDiscoveryPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryPeriodOk() (*AutoDiscoveryPeriod, bool) {
	if o == nil || IsNil(o.AutoDiscoveryPeriod) {
		return nil, false
	}
	return o.AutoDiscoveryPeriod, true
}

// HasAutoDiscoveryPeriod returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasAutoDiscoveryPeriod() bool {
	if o != nil && !IsNil(o.AutoDiscoveryPeriod) {
		return true
	}

	return false
}

// SetAutoDiscoveryPeriod gets a reference to the given AutoDiscoveryPeriod and assigns it to the AutoDiscoveryPeriod field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetAutoDiscoveryPeriod(v AutoDiscoveryPeriod) {
	o.AutoDiscoveryPeriod = &v
}

// GetAutoDiscoveryCadence returns the AutoDiscoveryCadence field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryCadence() int32 {
	if o == nil || IsNil(o.AutoDiscoveryCadence.Get()) {
		var ret int32
		return ret
	}
	return *o.AutoDiscoveryCadence.Get()
}

// GetAutoDiscoveryCadenceOk returns a tuple with the AutoDiscoveryCadence field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetAutoDiscoveryCadenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoDiscoveryCadence.Get(), o.AutoDiscoveryCadence.IsSet()
}

// HasAutoDiscoveryCadence returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasAutoDiscoveryCadence() bool {
	if o != nil && o.AutoDiscoveryCadence.IsSet() {
		return true
	}

	return false
}

// SetAutoDiscoveryCadence gets a reference to the given NullableInt32 and assigns it to the AutoDiscoveryCadence field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetAutoDiscoveryCadence(v int32) {
	o.AutoDiscoveryCadence.Set(&v)
}

// SetAutoDiscoveryCadenceNil sets the value for AutoDiscoveryCadence to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetAutoDiscoveryCadenceNil() {
	o.AutoDiscoveryCadence.Set(nil)
}

// UnsetAutoDiscoveryCadence ensures that no value is present for AutoDiscoveryCadence, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetAutoDiscoveryCadence() {
	o.AutoDiscoveryCadence.Unset()
}

// GetCleanupAbsentPackages returns the CleanupAbsentPackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetCleanupAbsentPackages() bool {
	if o == nil || IsNil(o.CleanupAbsentPackages.Get()) {
		var ret bool
		return ret
	}
	return *o.CleanupAbsentPackages.Get()
}

// GetCleanupAbsentPackagesOk returns a tuple with the CleanupAbsentPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditAppLibPackageDiscoveryProfileRequestModel) GetCleanupAbsentPackagesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CleanupAbsentPackages.Get(), o.CleanupAbsentPackages.IsSet()
}

// HasCleanupAbsentPackages returns a boolean if a field has been set.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) HasCleanupAbsentPackages() bool {
	if o != nil && o.CleanupAbsentPackages.IsSet() {
		return true
	}

	return false
}

// SetCleanupAbsentPackages gets a reference to the given NullableBool and assigns it to the CleanupAbsentPackages field.
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetCleanupAbsentPackages(v bool) {
	o.CleanupAbsentPackages.Set(&v)
}

// SetCleanupAbsentPackagesNil sets the value for CleanupAbsentPackages to be an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) SetCleanupAbsentPackagesNil() {
	o.CleanupAbsentPackages.Set(nil)
}

// UnsetCleanupAbsentPackages ensures that no value is present for CleanupAbsentPackages, not even an explicit nil
func (o *EditAppLibPackageDiscoveryProfileRequestModel) UnsetCleanupAbsentPackages() {
	o.CleanupAbsentPackages.Unset()
}

func (o EditAppLibPackageDiscoveryProfileRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditAppLibPackageDiscoveryProfileRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.DesktopGroupUid.IsSet() {
		toSerialize["DesktopGroupUid"] = o.DesktopGroupUid.Get()
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.Recurse.IsSet() {
		toSerialize["Recurse"] = o.Recurse.Get()
	}
	if o.ManagementServer.IsSet() {
		toSerialize["ManagementServer"] = o.ManagementServer.Get()
	}
	if o.PublishingServer.IsSet() {
		toSerialize["PublishingServer"] = o.PublishingServer.Get()
	}
	if o.Username.IsSet() {
		toSerialize["Username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["Password"] = o.Password.Get()
	}
	if !IsNil(o.PasswordFormat) {
		toSerialize["PasswordFormat"] = o.PasswordFormat
	}
	if o.DiscoverAppV.IsSet() {
		toSerialize["DiscoverAppV"] = o.DiscoverAppV.Get()
	}
	if o.DiscoverMsix.IsSet() {
		toSerialize["DiscoverMsix"] = o.DiscoverMsix.Get()
	}
	if o.DiscoverAppAttach.IsSet() {
		toSerialize["DiscoverAppAttach"] = o.DiscoverAppAttach.Get()
	}
	if o.DiscoverFlexApp.IsSet() {
		toSerialize["DiscoverFlexApp"] = o.DiscoverFlexApp.Get()
	}
	if o.DiscoverElasticAppLayers.IsSet() {
		toSerialize["DiscoverElasticAppLayers"] = o.DiscoverElasticAppLayers.Get()
	}
	if o.AutomateDiscovery.IsSet() {
		toSerialize["AutomateDiscovery"] = o.AutomateDiscovery.Get()
	}
	if !IsNil(o.AutoDiscoveryPeriod) {
		toSerialize["AutoDiscoveryPeriod"] = o.AutoDiscoveryPeriod
	}
	if o.AutoDiscoveryCadence.IsSet() {
		toSerialize["AutoDiscoveryCadence"] = o.AutoDiscoveryCadence.Get()
	}
	if o.CleanupAbsentPackages.IsSet() {
		toSerialize["CleanupAbsentPackages"] = o.CleanupAbsentPackages.Get()
	}
	return toSerialize, nil
}

type NullableEditAppLibPackageDiscoveryProfileRequestModel struct {
	value *EditAppLibPackageDiscoveryProfileRequestModel
	isSet bool
}

func (v NullableEditAppLibPackageDiscoveryProfileRequestModel) Get() *EditAppLibPackageDiscoveryProfileRequestModel {
	return v.value
}

func (v *NullableEditAppLibPackageDiscoveryProfileRequestModel) Set(val *EditAppLibPackageDiscoveryProfileRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEditAppLibPackageDiscoveryProfileRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEditAppLibPackageDiscoveryProfileRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditAppLibPackageDiscoveryProfileRequestModel(val *EditAppLibPackageDiscoveryProfileRequestModel) *NullableEditAppLibPackageDiscoveryProfileRequestModel {
	return &NullableEditAppLibPackageDiscoveryProfileRequestModel{value: val, isSet: true}
}

func (v NullableEditAppLibPackageDiscoveryProfileRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditAppLibPackageDiscoveryProfileRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
