# DesktopRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier of an existing desktop associated with the delivery group. | [optional] 
**ColorDepth** | Pointer to [**ColorDepth**](ColorDepth.md) |  | [optional] 
**Description** | Pointer to **string** | Specifies an optional description of the new desktop. The text may be visible to the end user, for example, as a tooltip associated with the desktop within receiver. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether the published desktop is enabled. A disabled desktop is ignored when determining which desktops are available for a user. | [optional] 
**ExcludedUserFilterEnabled** | Pointer to **bool** | Specifies whether the ExcludedUsers filter is enabled.  If the filter is disabled then any user entries in the filter are ignored when determining which desktops are available for a user. | [optional] 
**ExcludedUsers** | Pointer to **[]string** | Specifies the excluded users filter of the desktop; that is, the users and groups who are explicitly denied access to the published desktop. | [optional] 
**Icon** | Pointer to **string** | Specifies the image data of the icon used to display the published desktop to the user, and of assigned desktop(s) in the case where SharingKind is equal to Private. | [optional] 
**IncludedUserFilterEnabled** | Pointer to **bool** | Specifies whether the IncludedUsers filter is enabled.  If the filter is disabled then any user who satisfies the requirements of the delivery group&#39;s access policy is implicitly granted an entitlement to the published desktop. | [optional] 
**IncludedUsers** | Pointer to **[]string** | Specifies the included users filter of the desktop; that is, the users and groups who are explicitly granted access to the published desktop. | [optional] 
**LeasingBehavior** | Pointer to [**LeasingBehavior**](LeasingBehavior.md) |  | [optional] 
**MaxDesktops** | Pointer to **int32** | The number of machines from the delivery group which a user may privately allocate. | [optional] 
**Name** | Pointer to **string** | Specifies the administrative name of the desktop.  Each desktop in the site must have a unique name. | [optional] 
**PublishedName** | Pointer to **string** | The name of the published desktop as seen by the user, and of assigned desktop(s) in the case where SharingKind is equal to Private. | [optional] 
**RestrictToTag** | Pointer to **string** | Optional tag that may be used further to restrict which desktops may be made accessible to a user. | [optional] 
**SecureIcaRequired** | Pointer to **bool** | Specifies whether the desktop requires the SecureICA protocol for desktop sessions. | [optional] 
**SessionReconnection** | Pointer to [**SessionReconnection**](SessionReconnection.md) |  | [optional] 

## Methods

### NewDesktopRequestModel

`func NewDesktopRequestModel() *DesktopRequestModel`

NewDesktopRequestModel instantiates a new DesktopRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopRequestModelWithDefaults

`func NewDesktopRequestModelWithDefaults() *DesktopRequestModel`

NewDesktopRequestModelWithDefaults instantiates a new DesktopRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DesktopRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DesktopRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DesktopRequestModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DesktopRequestModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetColorDepth

`func (o *DesktopRequestModel) GetColorDepth() ColorDepth`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *DesktopRequestModel) GetColorDepthOk() (*ColorDepth, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *DesktopRequestModel) SetColorDepth(v ColorDepth)`

SetColorDepth sets ColorDepth field to given value.

### HasColorDepth

`func (o *DesktopRequestModel) HasColorDepth() bool`

HasColorDepth returns a boolean if a field has been set.

### GetDescription

`func (o *DesktopRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DesktopRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DesktopRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DesktopRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DesktopRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DesktopRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DesktopRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DesktopRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExcludedUserFilterEnabled

`func (o *DesktopRequestModel) GetExcludedUserFilterEnabled() bool`

GetExcludedUserFilterEnabled returns the ExcludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetExcludedUserFilterEnabledOk

`func (o *DesktopRequestModel) GetExcludedUserFilterEnabledOk() (*bool, bool)`

GetExcludedUserFilterEnabledOk returns a tuple with the ExcludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilterEnabled

`func (o *DesktopRequestModel) SetExcludedUserFilterEnabled(v bool)`

SetExcludedUserFilterEnabled sets ExcludedUserFilterEnabled field to given value.

### HasExcludedUserFilterEnabled

`func (o *DesktopRequestModel) HasExcludedUserFilterEnabled() bool`

HasExcludedUserFilterEnabled returns a boolean if a field has been set.

### GetExcludedUsers

`func (o *DesktopRequestModel) GetExcludedUsers() []string`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *DesktopRequestModel) GetExcludedUsersOk() (*[]string, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *DesktopRequestModel) SetExcludedUsers(v []string)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *DesktopRequestModel) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.

### GetIcon

`func (o *DesktopRequestModel) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *DesktopRequestModel) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *DesktopRequestModel) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *DesktopRequestModel) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIncludedUserFilterEnabled

`func (o *DesktopRequestModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *DesktopRequestModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *DesktopRequestModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *DesktopRequestModel) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### GetIncludedUsers

`func (o *DesktopRequestModel) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *DesktopRequestModel) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *DesktopRequestModel) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *DesktopRequestModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### GetLeasingBehavior

`func (o *DesktopRequestModel) GetLeasingBehavior() LeasingBehavior`

GetLeasingBehavior returns the LeasingBehavior field if non-nil, zero value otherwise.

### GetLeasingBehaviorOk

`func (o *DesktopRequestModel) GetLeasingBehaviorOk() (*LeasingBehavior, bool)`

GetLeasingBehaviorOk returns a tuple with the LeasingBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasingBehavior

`func (o *DesktopRequestModel) SetLeasingBehavior(v LeasingBehavior)`

SetLeasingBehavior sets LeasingBehavior field to given value.

### HasLeasingBehavior

`func (o *DesktopRequestModel) HasLeasingBehavior() bool`

HasLeasingBehavior returns a boolean if a field has been set.

### GetMaxDesktops

`func (o *DesktopRequestModel) GetMaxDesktops() int32`

GetMaxDesktops returns the MaxDesktops field if non-nil, zero value otherwise.

### GetMaxDesktopsOk

`func (o *DesktopRequestModel) GetMaxDesktopsOk() (*int32, bool)`

GetMaxDesktopsOk returns a tuple with the MaxDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDesktops

`func (o *DesktopRequestModel) SetMaxDesktops(v int32)`

SetMaxDesktops sets MaxDesktops field to given value.

### HasMaxDesktops

`func (o *DesktopRequestModel) HasMaxDesktops() bool`

HasMaxDesktops returns a boolean if a field has been set.

### GetName

`func (o *DesktopRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DesktopRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DesktopRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DesktopRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublishedName

`func (o *DesktopRequestModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *DesktopRequestModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *DesktopRequestModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *DesktopRequestModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### GetRestrictToTag

`func (o *DesktopRequestModel) GetRestrictToTag() string`

GetRestrictToTag returns the RestrictToTag field if non-nil, zero value otherwise.

### GetRestrictToTagOk

`func (o *DesktopRequestModel) GetRestrictToTagOk() (*string, bool)`

GetRestrictToTagOk returns a tuple with the RestrictToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToTag

`func (o *DesktopRequestModel) SetRestrictToTag(v string)`

SetRestrictToTag sets RestrictToTag field to given value.

### HasRestrictToTag

`func (o *DesktopRequestModel) HasRestrictToTag() bool`

HasRestrictToTag returns a boolean if a field has been set.

### GetSecureIcaRequired

`func (o *DesktopRequestModel) GetSecureIcaRequired() bool`

GetSecureIcaRequired returns the SecureIcaRequired field if non-nil, zero value otherwise.

### GetSecureIcaRequiredOk

`func (o *DesktopRequestModel) GetSecureIcaRequiredOk() (*bool, bool)`

GetSecureIcaRequiredOk returns a tuple with the SecureIcaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureIcaRequired

`func (o *DesktopRequestModel) SetSecureIcaRequired(v bool)`

SetSecureIcaRequired sets SecureIcaRequired field to given value.

### HasSecureIcaRequired

`func (o *DesktopRequestModel) HasSecureIcaRequired() bool`

HasSecureIcaRequired returns a boolean if a field has been set.

### GetSessionReconnection

`func (o *DesktopRequestModel) GetSessionReconnection() SessionReconnection`

GetSessionReconnection returns the SessionReconnection field if non-nil, zero value otherwise.

### GetSessionReconnectionOk

`func (o *DesktopRequestModel) GetSessionReconnectionOk() (*SessionReconnection, bool)`

GetSessionReconnectionOk returns a tuple with the SessionReconnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReconnection

`func (o *DesktopRequestModel) SetSessionReconnection(v SessionReconnection)`

SetSessionReconnection sets SessionReconnection field to given value.

### HasSessionReconnection

`func (o *DesktopRequestModel) HasSessionReconnection() bool`

HasSessionReconnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


