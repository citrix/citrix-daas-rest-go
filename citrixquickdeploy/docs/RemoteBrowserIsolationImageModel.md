# RemoteBrowserIsolationImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the image | [optional] 
**SbSessionVdaVersion** | Pointer to **string** | Shows version of sbsession | [optional] 
**CommandLineExecutable** | Pointer to **string** | The init session path | [optional] 
**CommandLineArguments** | Pointer to **string** | The init session arguments | [optional] 
**WorkingDirectory** | Pointer to **string** | The init session working directory | [optional] 
**LinuxVdaVersion** | Pointer to **string** | Shows version of image Virtual Desktop Agent | [optional] 
**OsVersion** | Pointer to **string** | Shows version of image OS | [optional] 
**VhdLocation** | Pointer to **string** | Url of the image which was provided by secure browser | [optional] 
**VhdSasLink** | Pointer to **string** | Url of the image with Sas token which was provided by secure browser | [optional] 
**TicketRequestAddress** | Pointer to **string** | Address of secure browser ticket request | [optional] 
**PolicyContent** | Pointer to **string** | The Group Policy content in base64 format | [optional] 
**VdaUpdateDelay** | Pointer to **int32** | Approximate maximum duration over which the reboot cycle runs, in minutes.  Value of 0 causes all machines to reboot immediately.  Defaults to -1: a \&quot;natural reboot\&quot; cycle allowing machines that are in-use to  continue working and be restarted only after they become idle. | [optional] 

## Methods

### NewRemoteBrowserIsolationImageModel

`func NewRemoteBrowserIsolationImageModel() *RemoteBrowserIsolationImageModel`

NewRemoteBrowserIsolationImageModel instantiates a new RemoteBrowserIsolationImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteBrowserIsolationImageModelWithDefaults

`func NewRemoteBrowserIsolationImageModelWithDefaults() *RemoteBrowserIsolationImageModel`

NewRemoteBrowserIsolationImageModelWithDefaults instantiates a new RemoteBrowserIsolationImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteBrowserIsolationImageModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteBrowserIsolationImageModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteBrowserIsolationImageModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteBrowserIsolationImageModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSbSessionVdaVersion

`func (o *RemoteBrowserIsolationImageModel) GetSbSessionVdaVersion() string`

GetSbSessionVdaVersion returns the SbSessionVdaVersion field if non-nil, zero value otherwise.

### GetSbSessionVdaVersionOk

`func (o *RemoteBrowserIsolationImageModel) GetSbSessionVdaVersionOk() (*string, bool)`

GetSbSessionVdaVersionOk returns a tuple with the SbSessionVdaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbSessionVdaVersion

`func (o *RemoteBrowserIsolationImageModel) SetSbSessionVdaVersion(v string)`

SetSbSessionVdaVersion sets SbSessionVdaVersion field to given value.

### HasSbSessionVdaVersion

`func (o *RemoteBrowserIsolationImageModel) HasSbSessionVdaVersion() bool`

HasSbSessionVdaVersion returns a boolean if a field has been set.

### GetCommandLineExecutable

`func (o *RemoteBrowserIsolationImageModel) GetCommandLineExecutable() string`

GetCommandLineExecutable returns the CommandLineExecutable field if non-nil, zero value otherwise.

### GetCommandLineExecutableOk

`func (o *RemoteBrowserIsolationImageModel) GetCommandLineExecutableOk() (*string, bool)`

GetCommandLineExecutableOk returns a tuple with the CommandLineExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineExecutable

`func (o *RemoteBrowserIsolationImageModel) SetCommandLineExecutable(v string)`

SetCommandLineExecutable sets CommandLineExecutable field to given value.

### HasCommandLineExecutable

`func (o *RemoteBrowserIsolationImageModel) HasCommandLineExecutable() bool`

HasCommandLineExecutable returns a boolean if a field has been set.

### GetCommandLineArguments

`func (o *RemoteBrowserIsolationImageModel) GetCommandLineArguments() string`

GetCommandLineArguments returns the CommandLineArguments field if non-nil, zero value otherwise.

### GetCommandLineArgumentsOk

`func (o *RemoteBrowserIsolationImageModel) GetCommandLineArgumentsOk() (*string, bool)`

GetCommandLineArgumentsOk returns a tuple with the CommandLineArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArguments

`func (o *RemoteBrowserIsolationImageModel) SetCommandLineArguments(v string)`

SetCommandLineArguments sets CommandLineArguments field to given value.

### HasCommandLineArguments

`func (o *RemoteBrowserIsolationImageModel) HasCommandLineArguments() bool`

HasCommandLineArguments returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *RemoteBrowserIsolationImageModel) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *RemoteBrowserIsolationImageModel) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *RemoteBrowserIsolationImageModel) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *RemoteBrowserIsolationImageModel) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetLinuxVdaVersion

`func (o *RemoteBrowserIsolationImageModel) GetLinuxVdaVersion() string`

GetLinuxVdaVersion returns the LinuxVdaVersion field if non-nil, zero value otherwise.

### GetLinuxVdaVersionOk

`func (o *RemoteBrowserIsolationImageModel) GetLinuxVdaVersionOk() (*string, bool)`

GetLinuxVdaVersionOk returns a tuple with the LinuxVdaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxVdaVersion

`func (o *RemoteBrowserIsolationImageModel) SetLinuxVdaVersion(v string)`

SetLinuxVdaVersion sets LinuxVdaVersion field to given value.

### HasLinuxVdaVersion

`func (o *RemoteBrowserIsolationImageModel) HasLinuxVdaVersion() bool`

HasLinuxVdaVersion returns a boolean if a field has been set.

### GetOsVersion

`func (o *RemoteBrowserIsolationImageModel) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *RemoteBrowserIsolationImageModel) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *RemoteBrowserIsolationImageModel) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *RemoteBrowserIsolationImageModel) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetVhdLocation

`func (o *RemoteBrowserIsolationImageModel) GetVhdLocation() string`

GetVhdLocation returns the VhdLocation field if non-nil, zero value otherwise.

### GetVhdLocationOk

`func (o *RemoteBrowserIsolationImageModel) GetVhdLocationOk() (*string, bool)`

GetVhdLocationOk returns a tuple with the VhdLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhdLocation

`func (o *RemoteBrowserIsolationImageModel) SetVhdLocation(v string)`

SetVhdLocation sets VhdLocation field to given value.

### HasVhdLocation

`func (o *RemoteBrowserIsolationImageModel) HasVhdLocation() bool`

HasVhdLocation returns a boolean if a field has been set.

### GetVhdSasLink

`func (o *RemoteBrowserIsolationImageModel) GetVhdSasLink() string`

GetVhdSasLink returns the VhdSasLink field if non-nil, zero value otherwise.

### GetVhdSasLinkOk

`func (o *RemoteBrowserIsolationImageModel) GetVhdSasLinkOk() (*string, bool)`

GetVhdSasLinkOk returns a tuple with the VhdSasLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhdSasLink

`func (o *RemoteBrowserIsolationImageModel) SetVhdSasLink(v string)`

SetVhdSasLink sets VhdSasLink field to given value.

### HasVhdSasLink

`func (o *RemoteBrowserIsolationImageModel) HasVhdSasLink() bool`

HasVhdSasLink returns a boolean if a field has been set.

### GetTicketRequestAddress

`func (o *RemoteBrowserIsolationImageModel) GetTicketRequestAddress() string`

GetTicketRequestAddress returns the TicketRequestAddress field if non-nil, zero value otherwise.

### GetTicketRequestAddressOk

`func (o *RemoteBrowserIsolationImageModel) GetTicketRequestAddressOk() (*string, bool)`

GetTicketRequestAddressOk returns a tuple with the TicketRequestAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketRequestAddress

`func (o *RemoteBrowserIsolationImageModel) SetTicketRequestAddress(v string)`

SetTicketRequestAddress sets TicketRequestAddress field to given value.

### HasTicketRequestAddress

`func (o *RemoteBrowserIsolationImageModel) HasTicketRequestAddress() bool`

HasTicketRequestAddress returns a boolean if a field has been set.

### GetPolicyContent

`func (o *RemoteBrowserIsolationImageModel) GetPolicyContent() string`

GetPolicyContent returns the PolicyContent field if non-nil, zero value otherwise.

### GetPolicyContentOk

`func (o *RemoteBrowserIsolationImageModel) GetPolicyContentOk() (*string, bool)`

GetPolicyContentOk returns a tuple with the PolicyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyContent

`func (o *RemoteBrowserIsolationImageModel) SetPolicyContent(v string)`

SetPolicyContent sets PolicyContent field to given value.

### HasPolicyContent

`func (o *RemoteBrowserIsolationImageModel) HasPolicyContent() bool`

HasPolicyContent returns a boolean if a field has been set.

### GetVdaUpdateDelay

`func (o *RemoteBrowserIsolationImageModel) GetVdaUpdateDelay() int32`

GetVdaUpdateDelay returns the VdaUpdateDelay field if non-nil, zero value otherwise.

### GetVdaUpdateDelayOk

`func (o *RemoteBrowserIsolationImageModel) GetVdaUpdateDelayOk() (*int32, bool)`

GetVdaUpdateDelayOk returns a tuple with the VdaUpdateDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaUpdateDelay

`func (o *RemoteBrowserIsolationImageModel) SetVdaUpdateDelay(v int32)`

SetVdaUpdateDelay sets VdaUpdateDelay field to given value.

### HasVdaUpdateDelay

`func (o *RemoteBrowserIsolationImageModel) HasVdaUpdateDelay() bool`

HasVdaUpdateDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


