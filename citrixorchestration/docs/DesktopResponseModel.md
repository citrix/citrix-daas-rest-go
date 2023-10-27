# DesktopResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the desktop. | [optional] 
**Uid** | Pointer to **int32** | &#x60;DEPRECATED&#x60; DEPRECATED. Use Id. | [optional] 
**ColorDepth** | Pointer to [**ColorDepth**](ColorDepth.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Optional description of the desktop. The text may be visible to the end user, for example, as a tooltip associated with the desktop within receiver. | [optional] 
**Enabled** | Pointer to **bool** | Whether the published desktop is enabled. A disabled desktop is ignored when determining which desktops are available for a user. | [optional] 
**ExcludedUserFilterEnabled** | Pointer to **bool** | Indicates whether the ExcludedUsers filter is enabled.  If the filter is disabled then any user entries in the filter are ignored when determining which desktops are available for a user. | [optional] 
**ExcludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The excluded users filter of the desktop; that is, the users and groups who are explicitly denied access to the published desktop. | [optional] 
**IconId** | Pointer to **NullableString** | Id of the icon used to display the published desktop to the user, and of assigned desktop(s) in the case where SharingKind is equal to Private. | [optional] 
**IncludedUserFilterEnabled** | Pointer to **bool** | Indicates whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access policy is implicitly granted an entitlement to the published desktop. | [optional] 
**IncludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | The included users filter of the desktop; that is, the users and groups who are explicitly granted access to the published desktop. | [optional] 
**LeasingBehavior** | Pointer to [**LeasingBehavior**](LeasingBehavior.md) |  | [optional] 
**MaxDesktops** | Pointer to **NullableInt32** | The number of machines from the delivery group which a user may privately allocate. | [optional] 
**Name** | Pointer to **NullableString** | The administrative name of the desktop. | [optional] 
**PublishedName** | Pointer to **NullableString** | The name of the published desktop as seen by the user, and of assigned desktop(s) in the case where SharingKind is equal to Private. | [optional] 
**RestrictToTag** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**SecureIcaRequired** | Pointer to **bool** | Indicates whether the desktop requires the SecureICA protocol for desktop sessions. | [optional] 
**SessionReconnection** | Pointer to [**SessionReconnection**](SessionReconnection.md) |  | [optional] 
**MachinesForAssignment** | Pointer to **NullableInt32** | Indicates the number of machines which are available for assignment based on this desktop configuration. | [optional] 
**Tenants** | Pointer to [**[]RefResponseModel**](RefResponseModel.md) | The tenant(s) that the desktop is assigned to.  If &#x60;null&#x60;, the desktop is not assigned to any tenants, and may be used by any tenant. | [optional] 

## Methods

### NewDesktopResponseModel

`func NewDesktopResponseModel() *DesktopResponseModel`

NewDesktopResponseModel instantiates a new DesktopResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopResponseModelWithDefaults

`func NewDesktopResponseModelWithDefaults() *DesktopResponseModel`

NewDesktopResponseModelWithDefaults instantiates a new DesktopResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DesktopResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DesktopResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DesktopResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DesktopResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DesktopResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DesktopResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUid

`func (o *DesktopResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DesktopResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DesktopResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DesktopResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetColorDepth

`func (o *DesktopResponseModel) GetColorDepth() ColorDepth`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *DesktopResponseModel) GetColorDepthOk() (*ColorDepth, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *DesktopResponseModel) SetColorDepth(v ColorDepth)`

SetColorDepth sets ColorDepth field to given value.

### HasColorDepth

`func (o *DesktopResponseModel) HasColorDepth() bool`

HasColorDepth returns a boolean if a field has been set.

### GetDescription

`func (o *DesktopResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DesktopResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DesktopResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DesktopResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DesktopResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DesktopResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *DesktopResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DesktopResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DesktopResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DesktopResponseModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExcludedUserFilterEnabled

`func (o *DesktopResponseModel) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *DesktopResponseModel) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *DesktopResponseModel) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.

### HasExcludedUserFilterEnabled

`func (o *DesktopResponseModel) HasExcludedUserFilterEnabled() bool`

HasExcludedUserFilterEnabled returns a boolean if a field has been set.

### GetExcludedUsers

`func (o *DesktopResponseModel) GetExcludedUsers() []IdentityUserResponseModel`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *DesktopResponseModel) GetExcludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *DesktopResponseModel) SetExcludedUsers(v []IdentityUserResponseModel)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *DesktopResponseModel) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.

### SetExcludedUsersNil

`func (o *DesktopResponseModel) SetExcludedUsersNil(b bool)`

 SetExcludedUsersNil sets the value for ExcludedUsers to be an explicit nil

### UnsetExcludedUsers
`func (o *DesktopResponseModel) UnsetExcludedUsers()`

UnsetExcludedUsers ensures that no value is present for ExcludedUsers, not even an explicit nil
### GetIconId

`func (o *DesktopResponseModel) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *DesktopResponseModel) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *DesktopResponseModel) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *DesktopResponseModel) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### SetIconIdNil

`func (o *DesktopResponseModel) SetIconIdNil(b bool)`

 SetIconIdNil sets the value for IconId to be an explicit nil

### UnsetIconId
`func (o *DesktopResponseModel) UnsetIconId()`

UnsetIconId ensures that no value is present for IconId, not even an explicit nil
### GetIncludedUserFilterEnabled

`func (o *DesktopResponseModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *DesktopResponseModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *DesktopResponseModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *DesktopResponseModel) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### GetIncludedUsers

`func (o *DesktopResponseModel) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *DesktopResponseModel) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *DesktopResponseModel) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *DesktopResponseModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *DesktopResponseModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *DesktopResponseModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil
### GetLeasingBehavior

`func (o *DesktopResponseModel) GetLeasingBehavior() LeasingBehavior`

GetLeasingBehavior returns the LeasingBehavior field if non-nil, zero value otherwise.

### GetLeasingBehaviorOk

`func (o *DesktopResponseModel) GetLeasingBehaviorOk() (*LeasingBehavior, bool)`

GetLeasingBehaviorOk returns a tuple with the LeasingBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasingBehavior

`func (o *DesktopResponseModel) SetLeasingBehavior(v LeasingBehavior)`

SetLeasingBehavior sets LeasingBehavior field to given value.

### HasLeasingBehavior

`func (o *DesktopResponseModel) HasLeasingBehavior() bool`

HasLeasingBehavior returns a boolean if a field has been set.

### GetMaxDesktops

`func (o *DesktopResponseModel) GetMaxDesktops() int32`

GetMaxDesktops returns the MaxDesktops field if non-nil, zero value otherwise.

### GetMaxDesktopsOk

`func (o *DesktopResponseModel) GetMaxDesktopsOk() (*int32, bool)`

GetMaxDesktopsOk returns a tuple with the MaxDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDesktops

`func (o *DesktopResponseModel) SetMaxDesktops(v int32)`

SetMaxDesktops sets MaxDesktops field to given value.

### HasMaxDesktops

`func (o *DesktopResponseModel) HasMaxDesktops() bool`

HasMaxDesktops returns a boolean if a field has been set.

### SetMaxDesktopsNil

`func (o *DesktopResponseModel) SetMaxDesktopsNil(b bool)`

 SetMaxDesktopsNil sets the value for MaxDesktops to be an explicit nil

### UnsetMaxDesktops
`func (o *DesktopResponseModel) UnsetMaxDesktops()`

UnsetMaxDesktops ensures that no value is present for MaxDesktops, not even an explicit nil
### GetName

`func (o *DesktopResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DesktopResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DesktopResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DesktopResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DesktopResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DesktopResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPublishedName

`func (o *DesktopResponseModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *DesktopResponseModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *DesktopResponseModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *DesktopResponseModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *DesktopResponseModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *DesktopResponseModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetRestrictToTag

`func (o *DesktopResponseModel) GetRestrictToTag() RefResponseModel`

GetRestrictToTag returns the RestrictToTag field if non-nil, zero value otherwise.

### GetRestrictToTagOk

`func (o *DesktopResponseModel) GetRestrictToTagOk() (*RefResponseModel, bool)`

GetRestrictToTagOk returns a tuple with the RestrictToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToTag

`func (o *DesktopResponseModel) SetRestrictToTag(v RefResponseModel)`

SetRestrictToTag sets RestrictToTag field to given value.

### HasRestrictToTag

`func (o *DesktopResponseModel) HasRestrictToTag() bool`

HasRestrictToTag returns a boolean if a field has been set.

### GetSecureIcaRequired

`func (o *DesktopResponseModel) GetSecureIcaRequired() bool`

GetSecureIcaRequired returns the SecureIcaRequired field if non-nil, zero value otherwise.

### GetSecureIcaRequiredOk

`func (o *DesktopResponseModel) GetSecureIcaRequiredOk() (*bool, bool)`

GetSecureIcaRequiredOk returns a tuple with the SecureIcaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureIcaRequired

`func (o *DesktopResponseModel) SetSecureIcaRequired(v bool)`

SetSecureIcaRequired sets SecureIcaRequired field to given value.

### HasSecureIcaRequired

`func (o *DesktopResponseModel) HasSecureIcaRequired() bool`

HasSecureIcaRequired returns a boolean if a field has been set.

### GetSessionReconnection

`func (o *DesktopResponseModel) GetSessionReconnection() SessionReconnection`

GetSessionReconnection returns the SessionReconnection field if non-nil, zero value otherwise.

### GetSessionReconnectionOk

`func (o *DesktopResponseModel) GetSessionReconnectionOk() (*SessionReconnection, bool)`

GetSessionReconnectionOk returns a tuple with the SessionReconnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReconnection

`func (o *DesktopResponseModel) SetSessionReconnection(v SessionReconnection)`

SetSessionReconnection sets SessionReconnection field to given value.

### HasSessionReconnection

`func (o *DesktopResponseModel) HasSessionReconnection() bool`

HasSessionReconnection returns a boolean if a field has been set.

### GetMachinesForAssignment

`func (o *DesktopResponseModel) GetMachinesForAssignment() int32`

GetMachinesForAssignment returns the MachinesForAssignment field if non-nil, zero value otherwise.

### GetMachinesForAssignmentOk

`func (o *DesktopResponseModel) GetMachinesForAssignmentOk() (*int32, bool)`

GetMachinesForAssignmentOk returns a tuple with the MachinesForAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinesForAssignment

`func (o *DesktopResponseModel) SetMachinesForAssignment(v int32)`

SetMachinesForAssignment sets MachinesForAssignment field to given value.

### HasMachinesForAssignment

`func (o *DesktopResponseModel) HasMachinesForAssignment() bool`

HasMachinesForAssignment returns a boolean if a field has been set.

### SetMachinesForAssignmentNil

`func (o *DesktopResponseModel) SetMachinesForAssignmentNil(b bool)`

 SetMachinesForAssignmentNil sets the value for MachinesForAssignment to be an explicit nil

### UnsetMachinesForAssignment
`func (o *DesktopResponseModel) UnsetMachinesForAssignment()`

UnsetMachinesForAssignment ensures that no value is present for MachinesForAssignment, not even an explicit nil
### GetTenants

`func (o *DesktopResponseModel) GetTenants() []RefResponseModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *DesktopResponseModel) GetTenantsOk() (*[]RefResponseModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *DesktopResponseModel) SetTenants(v []RefResponseModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *DesktopResponseModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *DesktopResponseModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *DesktopResponseModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


