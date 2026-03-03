# AddRemoteBrowserIsolationWebAppModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | Pointer to **NullableString** | The friendly name displayed to the admin in the UI | [optional] 
**DestinationUrl** | Pointer to **NullableString** | The target URL of the published app | [optional] 
**WebBrowser** | Pointer to **NullableString** | Which web browser was used for the launch | [optional] [default to "ChromeLnx"]
**PasscodeBased** | Pointer to **bool** | Indicates that an app uses shared passcode | [optional] 
**Authenticated** | Pointer to **bool** | Indicates that an app requires user authentication | [optional] 
**Password** | Pointer to **NullableString** | The password of an app | [optional] 
**IconHash** | Pointer to **NullableString** | The hash of the Secure Browser app icon | [optional] 

## Methods

### NewAddRemoteBrowserIsolationWebAppModel

`func NewAddRemoteBrowserIsolationWebAppModel() *AddRemoteBrowserIsolationWebAppModel`

NewAddRemoteBrowserIsolationWebAppModel instantiates a new AddRemoteBrowserIsolationWebAppModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRemoteBrowserIsolationWebAppModelWithDefaults

`func NewAddRemoteBrowserIsolationWebAppModelWithDefaults() *AddRemoteBrowserIsolationWebAppModel`

NewAddRemoteBrowserIsolationWebAppModelWithDefaults instantiates a new AddRemoteBrowserIsolationWebAppModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFriendlyName

`func (o *AddRemoteBrowserIsolationWebAppModel) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AddRemoteBrowserIsolationWebAppModel) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AddRemoteBrowserIsolationWebAppModel) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AddRemoteBrowserIsolationWebAppModel) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *AddRemoteBrowserIsolationWebAppModel) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *AddRemoteBrowserIsolationWebAppModel) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetDestinationUrl

`func (o *AddRemoteBrowserIsolationWebAppModel) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *AddRemoteBrowserIsolationWebAppModel) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *AddRemoteBrowserIsolationWebAppModel) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *AddRemoteBrowserIsolationWebAppModel) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### SetDestinationUrlNil

`func (o *AddRemoteBrowserIsolationWebAppModel) SetDestinationUrlNil(b bool)`

 SetDestinationUrlNil sets the value for DestinationUrl to be an explicit nil

### UnsetDestinationUrl
`func (o *AddRemoteBrowserIsolationWebAppModel) UnsetDestinationUrl()`

UnsetDestinationUrl ensures that no value is present for DestinationUrl, not even an explicit nil
### GetWebBrowser

`func (o *AddRemoteBrowserIsolationWebAppModel) GetWebBrowser() string`

GetWebBrowser returns the WebBrowser field if non-nil, zero value otherwise.

### GetWebBrowserOk

`func (o *AddRemoteBrowserIsolationWebAppModel) GetWebBrowserOk() (*string, bool)`

GetWebBrowserOk returns a tuple with the WebBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebBrowser

`func (o *AddRemoteBrowserIsolationWebAppModel) SetWebBrowser(v string)`

SetWebBrowser sets WebBrowser field to given value.

### HasWebBrowser

`func (o *AddRemoteBrowserIsolationWebAppModel) HasWebBrowser() bool`

HasWebBrowser returns a boolean if a field has been set.

### SetWebBrowserNil

`func (o *AddRemoteBrowserIsolationWebAppModel) SetWebBrowserNil(b bool)`

 SetWebBrowserNil sets the value for WebBrowser to be an explicit nil

### UnsetWebBrowser
`func (o *AddRemoteBrowserIsolationWebAppModel) UnsetWebBrowser()`

UnsetWebBrowser ensures that no value is present for WebBrowser, not even an explicit nil
### GetPasscodeBased

`func (o *AddRemoteBrowserIsolationWebAppModel) GetPasscodeBased() bool`

GetPasscodeBased returns the PasscodeBased field if non-nil, zero value otherwise.

### GetPasscodeBasedOk

`func (o *AddRemoteBrowserIsolationWebAppModel) GetPasscodeBasedOk() (*bool, bool)`

GetPasscodeBasedOk returns a tuple with the PasscodeBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodeBased

`func (o *AddRemoteBrowserIsolationWebAppModel) SetPasscodeBased(v bool)`

SetPasscodeBased sets PasscodeBased field to given value.

### HasPasscodeBased

`func (o *AddRemoteBrowserIsolationWebAppModel) HasPasscodeBased() bool`

HasPasscodeBased returns a boolean if a field has been set.

### GetAuthenticated

`func (o *AddRemoteBrowserIsolationWebAppModel) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *AddRemoteBrowserIsolationWebAppModel) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *AddRemoteBrowserIsolationWebAppModel) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *AddRemoteBrowserIsolationWebAppModel) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### GetPassword

`func (o *AddRemoteBrowserIsolationWebAppModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddRemoteBrowserIsolationWebAppModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddRemoteBrowserIsolationWebAppModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddRemoteBrowserIsolationWebAppModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AddRemoteBrowserIsolationWebAppModel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AddRemoteBrowserIsolationWebAppModel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetIconHash

`func (o *AddRemoteBrowserIsolationWebAppModel) GetIconHash() string`

GetIconHash returns the IconHash field if non-nil, zero value otherwise.

### GetIconHashOk

`func (o *AddRemoteBrowserIsolationWebAppModel) GetIconHashOk() (*string, bool)`

GetIconHashOk returns a tuple with the IconHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconHash

`func (o *AddRemoteBrowserIsolationWebAppModel) SetIconHash(v string)`

SetIconHash sets IconHash field to given value.

### HasIconHash

`func (o *AddRemoteBrowserIsolationWebAppModel) HasIconHash() bool`

HasIconHash returns a boolean if a field has been set.

### SetIconHashNil

`func (o *AddRemoteBrowserIsolationWebAppModel) SetIconHashNil(b bool)`

 SetIconHashNil sets the value for IconHash to be an explicit nil

### UnsetIconHash
`func (o *AddRemoteBrowserIsolationWebAppModel) UnsetIconHash()`

UnsetIconHash ensures that no value is present for IconHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


