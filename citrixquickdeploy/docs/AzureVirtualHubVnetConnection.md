# AzureVirtualHubVnetConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name of the hub VNet connection. | [optional] 
**ResourceId** | Pointer to **NullableString** | Gets or sets the Azure resource ID of the connection. | [optional] 
**RemoteVnetId** | Pointer to **NullableString** | Gets or sets the resource ID of the remote VNet connected to the hub. | [optional] 
**RemoteVnetName** | Pointer to **NullableString** | Gets or sets the name of the remote VNet. | [optional] 
**ProvisioningState** | Pointer to **NullableString** | Gets or sets the provisioning state of the connection. | [optional] 
**EnableInternetSecurity** | Pointer to **NullableBool** | Gets or sets a value indicating whether internet security is enabled for this connection. | [optional] 

## Methods

### NewAzureVirtualHubVnetConnection

`func NewAzureVirtualHubVnetConnection() *AzureVirtualHubVnetConnection`

NewAzureVirtualHubVnetConnection instantiates a new AzureVirtualHubVnetConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVirtualHubVnetConnectionWithDefaults

`func NewAzureVirtualHubVnetConnectionWithDefaults() *AzureVirtualHubVnetConnection`

NewAzureVirtualHubVnetConnectionWithDefaults instantiates a new AzureVirtualHubVnetConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureVirtualHubVnetConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureVirtualHubVnetConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureVirtualHubVnetConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureVirtualHubVnetConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureVirtualHubVnetConnection) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureVirtualHubVnetConnection) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResourceId

`func (o *AzureVirtualHubVnetConnection) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AzureVirtualHubVnetConnection) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AzureVirtualHubVnetConnection) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AzureVirtualHubVnetConnection) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AzureVirtualHubVnetConnection) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AzureVirtualHubVnetConnection) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetRemoteVnetId

`func (o *AzureVirtualHubVnetConnection) GetRemoteVnetId() string`

GetRemoteVnetId returns the RemoteVnetId field if non-nil, zero value otherwise.

### GetRemoteVnetIdOk

`func (o *AzureVirtualHubVnetConnection) GetRemoteVnetIdOk() (*string, bool)`

GetRemoteVnetIdOk returns a tuple with the RemoteVnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVnetId

`func (o *AzureVirtualHubVnetConnection) SetRemoteVnetId(v string)`

SetRemoteVnetId sets RemoteVnetId field to given value.

### HasRemoteVnetId

`func (o *AzureVirtualHubVnetConnection) HasRemoteVnetId() bool`

HasRemoteVnetId returns a boolean if a field has been set.

### SetRemoteVnetIdNil

`func (o *AzureVirtualHubVnetConnection) SetRemoteVnetIdNil(b bool)`

 SetRemoteVnetIdNil sets the value for RemoteVnetId to be an explicit nil

### UnsetRemoteVnetId
`func (o *AzureVirtualHubVnetConnection) UnsetRemoteVnetId()`

UnsetRemoteVnetId ensures that no value is present for RemoteVnetId, not even an explicit nil
### GetRemoteVnetName

`func (o *AzureVirtualHubVnetConnection) GetRemoteVnetName() string`

GetRemoteVnetName returns the RemoteVnetName field if non-nil, zero value otherwise.

### GetRemoteVnetNameOk

`func (o *AzureVirtualHubVnetConnection) GetRemoteVnetNameOk() (*string, bool)`

GetRemoteVnetNameOk returns a tuple with the RemoteVnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVnetName

`func (o *AzureVirtualHubVnetConnection) SetRemoteVnetName(v string)`

SetRemoteVnetName sets RemoteVnetName field to given value.

### HasRemoteVnetName

`func (o *AzureVirtualHubVnetConnection) HasRemoteVnetName() bool`

HasRemoteVnetName returns a boolean if a field has been set.

### SetRemoteVnetNameNil

`func (o *AzureVirtualHubVnetConnection) SetRemoteVnetNameNil(b bool)`

 SetRemoteVnetNameNil sets the value for RemoteVnetName to be an explicit nil

### UnsetRemoteVnetName
`func (o *AzureVirtualHubVnetConnection) UnsetRemoteVnetName()`

UnsetRemoteVnetName ensures that no value is present for RemoteVnetName, not even an explicit nil
### GetProvisioningState

`func (o *AzureVirtualHubVnetConnection) GetProvisioningState() string`

GetProvisioningState returns the ProvisioningState field if non-nil, zero value otherwise.

### GetProvisioningStateOk

`func (o *AzureVirtualHubVnetConnection) GetProvisioningStateOk() (*string, bool)`

GetProvisioningStateOk returns a tuple with the ProvisioningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningState

`func (o *AzureVirtualHubVnetConnection) SetProvisioningState(v string)`

SetProvisioningState sets ProvisioningState field to given value.

### HasProvisioningState

`func (o *AzureVirtualHubVnetConnection) HasProvisioningState() bool`

HasProvisioningState returns a boolean if a field has been set.

### SetProvisioningStateNil

`func (o *AzureVirtualHubVnetConnection) SetProvisioningStateNil(b bool)`

 SetProvisioningStateNil sets the value for ProvisioningState to be an explicit nil

### UnsetProvisioningState
`func (o *AzureVirtualHubVnetConnection) UnsetProvisioningState()`

UnsetProvisioningState ensures that no value is present for ProvisioningState, not even an explicit nil
### GetEnableInternetSecurity

`func (o *AzureVirtualHubVnetConnection) GetEnableInternetSecurity() bool`

GetEnableInternetSecurity returns the EnableInternetSecurity field if non-nil, zero value otherwise.

### GetEnableInternetSecurityOk

`func (o *AzureVirtualHubVnetConnection) GetEnableInternetSecurityOk() (*bool, bool)`

GetEnableInternetSecurityOk returns a tuple with the EnableInternetSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternetSecurity

`func (o *AzureVirtualHubVnetConnection) SetEnableInternetSecurity(v bool)`

SetEnableInternetSecurity sets EnableInternetSecurity field to given value.

### HasEnableInternetSecurity

`func (o *AzureVirtualHubVnetConnection) HasEnableInternetSecurity() bool`

HasEnableInternetSecurity returns a boolean if a field has been set.

### SetEnableInternetSecurityNil

`func (o *AzureVirtualHubVnetConnection) SetEnableInternetSecurityNil(b bool)`

 SetEnableInternetSecurityNil sets the value for EnableInternetSecurity to be an explicit nil

### UnsetEnableInternetSecurity
`func (o *AzureVirtualHubVnetConnection) UnsetEnableInternetSecurity()`

UnsetEnableInternetSecurity ensures that no value is present for EnableInternetSecurity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


