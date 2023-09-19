# EditHypervisorAzureConnectionRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the hypervisor to create.  Optional; if not specified, will not be changed. | [optional] 
**InMaintenanceMode** | Pointer to **NullableBool** | Specifies whether the hypervisor is in maintenance mode, which disables all communication between XenApp &amp; XenDesktop and the Hypervisor.  Optional; if not specified, will not be changed. | [optional] 
**Scopes** | Pointer to **[]string** | Administrative scopes which the newly created hypervisor will be a part of. If not specified, will not be changed. The \&quot;All\&quot; scope, and any tenant scopes, are implicit and cannot be removed.  To remove from all non-implicit scopes, specify an empty array ([]). Specifying tenant scopes is equivalent to specifying the Tenants property and is subject to the same constraints. | [optional] 
**Tenants** | Pointer to **[]string** | Tenants to associate with the hypervisor. | [optional] 
**MaxAbsoluteActiveActions** | Pointer to **NullableInt32** | Maximum number of actions that can execute in parallel on the hypervisor.  Optional; if not specified, will not be changed. | [optional] 
**MaxAbsoluteNewActionsPerMinute** | Pointer to **NullableInt32** | Maximum number of actions that can be started on the hypervisor per-minute.  Optional; if not specified, will not be changed. | [optional] 
**MaxPowerActionsPercentageOfMachines** | Pointer to **NullableInt32** | Maximum percentage of machines on the hypervisor which can have their power state changed simultaneously.  Optional; if not specified, will not be changed. | [optional] 
**ConnectionOptions** | Pointer to **NullableString** | Connection options to use for the hypervisor.  Optional; if not specified, will not be changed.  May be removed by specifying an empty string (&#x60;\&quot;\&quot;&#x60;). | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata for hypervisor connection. When set the property value equal to null/empty means to remove this property. | [optional] 
**ApiKey** | Pointer to **NullableString** | The API key used to authenticate with the AWS APIs.  Optional.  If not specified, will not be changed.  If specified, the SecretKey must also be specified. | [optional] 
**SecretKey** | Pointer to **NullableString** | The secret key used to authenticate with the AWS APIs.  Optional. Must be specified in the format indicated by SecretKeyFormat. | [optional] 
**SecretKeyFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**MaximumConcurrentProvisioningOperations** | Pointer to **NullableInt32** | Maximum number of concurrent AWS provisioning operations. Optional.  If not specified, will not be changed. | [optional] 
**ApplicationId** | Pointer to **NullableString** | Application ID of the service principal used to access the Azure APIs.  Optional.  If not specified, will not be changed.  If specified, then ApplicationSecret must also be specified. | [optional] 
**ApplicationSecret** | Pointer to **NullableString** | The Application Secret of the service principal used to access the Azure APIs.  Optional.  If not specified, will not be changed.  If specified, must in the format indicated by ApplicationSecretFormat. | [optional] 
**ApplicationSecretFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**CustomProperties** | Pointer to **NullableString** | The properties of host connection that are specific to the target hosting infrastructure. | [optional] 
**ServiceAccountId** | Pointer to **NullableString** | The service account ID used to access the Google Cloud APIs. Optional.  If not specified, will not be changed.  If specified, ServiceAccountCredentials must also be specified. | [optional] 
**ServiceAccountCredentials** | Pointer to **NullableString** | the JSON-encoded service account credentials used to access the Google Cloud APIs.  Optional.  If not specified, will not be changed.  If specified, must be in the format indicated by ServiceAccountCredentialsFormat. | [optional] 
**ServiceAccountCredentialsFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**UserName** | Pointer to **NullableString** | Hypervisor user name.  Optional.  If not specified, will not be changed.  If specified, Password must also be specified. | [optional] 
**Password** | Pointer to **NullableString** | Hypervisor password.  Optional.  If specified, must be in the format indicated by PasswordFormat. | [optional] 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**Addresses** | Pointer to **[]string** | Hypervisor address(es).  Optional.  If not specified, will not be changed. | [optional] 
**SslThumbprints** | Pointer to **[]string** | SSL certificate thumbprints to consider acceptable for this connection.  Optional.  If not specified, will not be changed.  To remove previously specified values, specify an empty array (&#x60;[]&#x60;). | [optional] 
**SccmWakeUpProxy** | Pointer to **NullableBool** | Specifies whether to use Microsoft System Center Configuration Manager 2012 SP1 Wake-up Proxy for power management.  Optional.  If not specified, will not be changed. | [optional] 
**WakeOnLanPackets** | Pointer to [**WakeOnLanTransmission**](WakeOnLanTransmission.md) |  | [optional] 
**ServiceAccountCredential** | Pointer to **string** | The private key string to access the Oracle Cloud Infrastructure APIs. Required. Must be specified in the format indicated by ServiceAccountCredentialFormat. | [optional] 
**ServiceAccountCredentialFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**ServiceAccountFingerprint** | Pointer to **string** | The fingerprint of the public key associate with the ServiceAccountCredential. | [optional] 

## Methods

### NewEditHypervisorAzureConnectionRequestModel

`func NewEditHypervisorAzureConnectionRequestModel(connectionType HypervisorConnectionType, ) *EditHypervisorAzureConnectionRequestModel`

NewEditHypervisorAzureConnectionRequestModel instantiates a new EditHypervisorAzureConnectionRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHypervisorAzureConnectionRequestModelWithDefaults

`func NewEditHypervisorAzureConnectionRequestModelWithDefaults() *EditHypervisorAzureConnectionRequestModel`

NewEditHypervisorAzureConnectionRequestModelWithDefaults instantiates a new EditHypervisorAzureConnectionRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditHypervisorAzureConnectionRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditHypervisorAzureConnectionRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditHypervisorAzureConnectionRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetInMaintenanceMode

`func (o *EditHypervisorAzureConnectionRequestModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *EditHypervisorAzureConnectionRequestModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *EditHypervisorAzureConnectionRequestModel) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.

### SetInMaintenanceModeNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetInMaintenanceModeNil(b bool)`

 SetInMaintenanceModeNil sets the value for InMaintenanceMode to be an explicit nil

### UnsetInMaintenanceMode
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetInMaintenanceMode()`

UnsetInMaintenanceMode ensures that no value is present for InMaintenanceMode, not even an explicit nil
### GetScopes

`func (o *EditHypervisorAzureConnectionRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *EditHypervisorAzureConnectionRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *EditHypervisorAzureConnectionRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *EditHypervisorAzureConnectionRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *EditHypervisorAzureConnectionRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *EditHypervisorAzureConnectionRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetMaxAbsoluteActiveActions

`func (o *EditHypervisorAzureConnectionRequestModel) GetMaxAbsoluteActiveActions() int32`

GetMaxAbsoluteActiveActions returns the MaxAbsoluteActiveActions field if non-nil, zero value otherwise.

### GetMaxAbsoluteActiveActionsOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetMaxAbsoluteActiveActionsOk() (*int32, bool)`

GetMaxAbsoluteActiveActionsOk returns a tuple with the MaxAbsoluteActiveActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteActiveActions

`func (o *EditHypervisorAzureConnectionRequestModel) SetMaxAbsoluteActiveActions(v int32)`

SetMaxAbsoluteActiveActions sets MaxAbsoluteActiveActions field to given value.

### HasMaxAbsoluteActiveActions

`func (o *EditHypervisorAzureConnectionRequestModel) HasMaxAbsoluteActiveActions() bool`

HasMaxAbsoluteActiveActions returns a boolean if a field has been set.

### SetMaxAbsoluteActiveActionsNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetMaxAbsoluteActiveActionsNil(b bool)`

 SetMaxAbsoluteActiveActionsNil sets the value for MaxAbsoluteActiveActions to be an explicit nil

### UnsetMaxAbsoluteActiveActions
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetMaxAbsoluteActiveActions()`

UnsetMaxAbsoluteActiveActions ensures that no value is present for MaxAbsoluteActiveActions, not even an explicit nil
### GetMaxAbsoluteNewActionsPerMinute

`func (o *EditHypervisorAzureConnectionRequestModel) GetMaxAbsoluteNewActionsPerMinute() int32`

GetMaxAbsoluteNewActionsPerMinute returns the MaxAbsoluteNewActionsPerMinute field if non-nil, zero value otherwise.

### GetMaxAbsoluteNewActionsPerMinuteOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetMaxAbsoluteNewActionsPerMinuteOk() (*int32, bool)`

GetMaxAbsoluteNewActionsPerMinuteOk returns a tuple with the MaxAbsoluteNewActionsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAbsoluteNewActionsPerMinute

`func (o *EditHypervisorAzureConnectionRequestModel) SetMaxAbsoluteNewActionsPerMinute(v int32)`

SetMaxAbsoluteNewActionsPerMinute sets MaxAbsoluteNewActionsPerMinute field to given value.

### HasMaxAbsoluteNewActionsPerMinute

`func (o *EditHypervisorAzureConnectionRequestModel) HasMaxAbsoluteNewActionsPerMinute() bool`

HasMaxAbsoluteNewActionsPerMinute returns a boolean if a field has been set.

### SetMaxAbsoluteNewActionsPerMinuteNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetMaxAbsoluteNewActionsPerMinuteNil(b bool)`

 SetMaxAbsoluteNewActionsPerMinuteNil sets the value for MaxAbsoluteNewActionsPerMinute to be an explicit nil

### UnsetMaxAbsoluteNewActionsPerMinute
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetMaxAbsoluteNewActionsPerMinute()`

UnsetMaxAbsoluteNewActionsPerMinute ensures that no value is present for MaxAbsoluteNewActionsPerMinute, not even an explicit nil
### GetMaxPowerActionsPercentageOfMachines

`func (o *EditHypervisorAzureConnectionRequestModel) GetMaxPowerActionsPercentageOfMachines() int32`

GetMaxPowerActionsPercentageOfMachines returns the MaxPowerActionsPercentageOfMachines field if non-nil, zero value otherwise.

### GetMaxPowerActionsPercentageOfMachinesOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetMaxPowerActionsPercentageOfMachinesOk() (*int32, bool)`

GetMaxPowerActionsPercentageOfMachinesOk returns a tuple with the MaxPowerActionsPercentageOfMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPowerActionsPercentageOfMachines

`func (o *EditHypervisorAzureConnectionRequestModel) SetMaxPowerActionsPercentageOfMachines(v int32)`

SetMaxPowerActionsPercentageOfMachines sets MaxPowerActionsPercentageOfMachines field to given value.

### HasMaxPowerActionsPercentageOfMachines

`func (o *EditHypervisorAzureConnectionRequestModel) HasMaxPowerActionsPercentageOfMachines() bool`

HasMaxPowerActionsPercentageOfMachines returns a boolean if a field has been set.

### SetMaxPowerActionsPercentageOfMachinesNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetMaxPowerActionsPercentageOfMachinesNil(b bool)`

 SetMaxPowerActionsPercentageOfMachinesNil sets the value for MaxPowerActionsPercentageOfMachines to be an explicit nil

### UnsetMaxPowerActionsPercentageOfMachines
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetMaxPowerActionsPercentageOfMachines()`

UnsetMaxPowerActionsPercentageOfMachines ensures that no value is present for MaxPowerActionsPercentageOfMachines, not even an explicit nil
### GetConnectionOptions

`func (o *EditHypervisorAzureConnectionRequestModel) GetConnectionOptions() string`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetConnectionOptionsOk() (*string, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *EditHypervisorAzureConnectionRequestModel) SetConnectionOptions(v string)`

SetConnectionOptions sets ConnectionOptions field to given value.

### HasConnectionOptions

`func (o *EditHypervisorAzureConnectionRequestModel) HasConnectionOptions() bool`

HasConnectionOptions returns a boolean if a field has been set.

### SetConnectionOptionsNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetConnectionOptionsNil(b bool)`

 SetConnectionOptionsNil sets the value for ConnectionOptions to be an explicit nil

### UnsetConnectionOptions
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetConnectionOptions()`

UnsetConnectionOptions ensures that no value is present for ConnectionOptions, not even an explicit nil
### GetConnectionType

`func (o *EditHypervisorAzureConnectionRequestModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *EditHypervisorAzureConnectionRequestModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetMetadata

`func (o *EditHypervisorAzureConnectionRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EditHypervisorAzureConnectionRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EditHypervisorAzureConnectionRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetApiKey

`func (o *EditHypervisorAzureConnectionRequestModel) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *EditHypervisorAzureConnectionRequestModel) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *EditHypervisorAzureConnectionRequestModel) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetSecretKey

`func (o *EditHypervisorAzureConnectionRequestModel) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *EditHypervisorAzureConnectionRequestModel) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *EditHypervisorAzureConnectionRequestModel) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetSecretKeyFormat

`func (o *EditHypervisorAzureConnectionRequestModel) GetSecretKeyFormat() IdentityPasswordFormat`

GetSecretKeyFormat returns the SecretKeyFormat field if non-nil, zero value otherwise.

### GetSecretKeyFormatOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetSecretKeyFormatOk() (*IdentityPasswordFormat, bool)`

GetSecretKeyFormatOk returns a tuple with the SecretKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyFormat

`func (o *EditHypervisorAzureConnectionRequestModel) SetSecretKeyFormat(v IdentityPasswordFormat)`

SetSecretKeyFormat sets SecretKeyFormat field to given value.

### HasSecretKeyFormat

`func (o *EditHypervisorAzureConnectionRequestModel) HasSecretKeyFormat() bool`

HasSecretKeyFormat returns a boolean if a field has been set.

### GetMaximumConcurrentProvisioningOperations

`func (o *EditHypervisorAzureConnectionRequestModel) GetMaximumConcurrentProvisioningOperations() int32`

GetMaximumConcurrentProvisioningOperations returns the MaximumConcurrentProvisioningOperations field if non-nil, zero value otherwise.

### GetMaximumConcurrentProvisioningOperationsOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetMaximumConcurrentProvisioningOperationsOk() (*int32, bool)`

GetMaximumConcurrentProvisioningOperationsOk returns a tuple with the MaximumConcurrentProvisioningOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentProvisioningOperations

`func (o *EditHypervisorAzureConnectionRequestModel) SetMaximumConcurrentProvisioningOperations(v int32)`

SetMaximumConcurrentProvisioningOperations sets MaximumConcurrentProvisioningOperations field to given value.

### HasMaximumConcurrentProvisioningOperations

`func (o *EditHypervisorAzureConnectionRequestModel) HasMaximumConcurrentProvisioningOperations() bool`

HasMaximumConcurrentProvisioningOperations returns a boolean if a field has been set.

### SetMaximumConcurrentProvisioningOperationsNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetMaximumConcurrentProvisioningOperationsNil(b bool)`

 SetMaximumConcurrentProvisioningOperationsNil sets the value for MaximumConcurrentProvisioningOperations to be an explicit nil

### UnsetMaximumConcurrentProvisioningOperations
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetMaximumConcurrentProvisioningOperations()`

UnsetMaximumConcurrentProvisioningOperations ensures that no value is present for MaximumConcurrentProvisioningOperations, not even an explicit nil
### GetApplicationId

`func (o *EditHypervisorAzureConnectionRequestModel) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *EditHypervisorAzureConnectionRequestModel) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *EditHypervisorAzureConnectionRequestModel) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationIdNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetApplicationSecret

`func (o *EditHypervisorAzureConnectionRequestModel) GetApplicationSecret() string`

GetApplicationSecret returns the ApplicationSecret field if non-nil, zero value otherwise.

### GetApplicationSecretOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetApplicationSecretOk() (*string, bool)`

GetApplicationSecretOk returns a tuple with the ApplicationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecret

`func (o *EditHypervisorAzureConnectionRequestModel) SetApplicationSecret(v string)`

SetApplicationSecret sets ApplicationSecret field to given value.

### HasApplicationSecret

`func (o *EditHypervisorAzureConnectionRequestModel) HasApplicationSecret() bool`

HasApplicationSecret returns a boolean if a field has been set.

### SetApplicationSecretNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetApplicationSecretNil(b bool)`

 SetApplicationSecretNil sets the value for ApplicationSecret to be an explicit nil

### UnsetApplicationSecret
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetApplicationSecret()`

UnsetApplicationSecret ensures that no value is present for ApplicationSecret, not even an explicit nil
### GetApplicationSecretFormat

`func (o *EditHypervisorAzureConnectionRequestModel) GetApplicationSecretFormat() IdentityPasswordFormat`

GetApplicationSecretFormat returns the ApplicationSecretFormat field if non-nil, zero value otherwise.

### GetApplicationSecretFormatOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetApplicationSecretFormatOk() (*IdentityPasswordFormat, bool)`

GetApplicationSecretFormatOk returns a tuple with the ApplicationSecretFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSecretFormat

`func (o *EditHypervisorAzureConnectionRequestModel) SetApplicationSecretFormat(v IdentityPasswordFormat)`

SetApplicationSecretFormat sets ApplicationSecretFormat field to given value.

### HasApplicationSecretFormat

`func (o *EditHypervisorAzureConnectionRequestModel) HasApplicationSecretFormat() bool`

HasApplicationSecretFormat returns a boolean if a field has been set.

### GetCustomProperties

`func (o *EditHypervisorAzureConnectionRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *EditHypervisorAzureConnectionRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *EditHypervisorAzureConnectionRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetServiceAccountId

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *EditHypervisorAzureConnectionRequestModel) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *EditHypervisorAzureConnectionRequestModel) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### SetServiceAccountIdNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetServiceAccountIdNil(b bool)`

 SetServiceAccountIdNil sets the value for ServiceAccountId to be an explicit nil

### UnsetServiceAccountId
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetServiceAccountId()`

UnsetServiceAccountId ensures that no value is present for ServiceAccountId, not even an explicit nil
### GetServiceAccountCredentials

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountCredentials() string`

GetServiceAccountCredentials returns the ServiceAccountCredentials field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountCredentialsOk() (*string, bool)`

GetServiceAccountCredentialsOk returns a tuple with the ServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentials

`func (o *EditHypervisorAzureConnectionRequestModel) SetServiceAccountCredentials(v string)`

SetServiceAccountCredentials sets ServiceAccountCredentials field to given value.

### HasServiceAccountCredentials

`func (o *EditHypervisorAzureConnectionRequestModel) HasServiceAccountCredentials() bool`

HasServiceAccountCredentials returns a boolean if a field has been set.

### SetServiceAccountCredentialsNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetServiceAccountCredentialsNil(b bool)`

 SetServiceAccountCredentialsNil sets the value for ServiceAccountCredentials to be an explicit nil

### UnsetServiceAccountCredentials
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetServiceAccountCredentials()`

UnsetServiceAccountCredentials ensures that no value is present for ServiceAccountCredentials, not even an explicit nil
### GetServiceAccountCredentialsFormat

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountCredentialsFormat() IdentityPasswordFormat`

GetServiceAccountCredentialsFormat returns the ServiceAccountCredentialsFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialsFormatOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountCredentialsFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialsFormatOk returns a tuple with the ServiceAccountCredentialsFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialsFormat

`func (o *EditHypervisorAzureConnectionRequestModel) SetServiceAccountCredentialsFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialsFormat sets ServiceAccountCredentialsFormat field to given value.

### HasServiceAccountCredentialsFormat

`func (o *EditHypervisorAzureConnectionRequestModel) HasServiceAccountCredentialsFormat() bool`

HasServiceAccountCredentialsFormat returns a boolean if a field has been set.

### GetUserName

`func (o *EditHypervisorAzureConnectionRequestModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *EditHypervisorAzureConnectionRequestModel) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *EditHypervisorAzureConnectionRequestModel) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetPassword

`func (o *EditHypervisorAzureConnectionRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EditHypervisorAzureConnectionRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EditHypervisorAzureConnectionRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordFormat

`func (o *EditHypervisorAzureConnectionRequestModel) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *EditHypervisorAzureConnectionRequestModel) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *EditHypervisorAzureConnectionRequestModel) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.

### GetAddresses

`func (o *EditHypervisorAzureConnectionRequestModel) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *EditHypervisorAzureConnectionRequestModel) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *EditHypervisorAzureConnectionRequestModel) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetSslThumbprints

`func (o *EditHypervisorAzureConnectionRequestModel) GetSslThumbprints() []string`

GetSslThumbprints returns the SslThumbprints field if non-nil, zero value otherwise.

### GetSslThumbprintsOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetSslThumbprintsOk() (*[]string, bool)`

GetSslThumbprintsOk returns a tuple with the SslThumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprints

`func (o *EditHypervisorAzureConnectionRequestModel) SetSslThumbprints(v []string)`

SetSslThumbprints sets SslThumbprints field to given value.

### HasSslThumbprints

`func (o *EditHypervisorAzureConnectionRequestModel) HasSslThumbprints() bool`

HasSslThumbprints returns a boolean if a field has been set.

### SetSslThumbprintsNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetSslThumbprintsNil(b bool)`

 SetSslThumbprintsNil sets the value for SslThumbprints to be an explicit nil

### UnsetSslThumbprints
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetSslThumbprints()`

UnsetSslThumbprints ensures that no value is present for SslThumbprints, not even an explicit nil
### GetSccmWakeUpProxy

`func (o *EditHypervisorAzureConnectionRequestModel) GetSccmWakeUpProxy() bool`

GetSccmWakeUpProxy returns the SccmWakeUpProxy field if non-nil, zero value otherwise.

### GetSccmWakeUpProxyOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetSccmWakeUpProxyOk() (*bool, bool)`

GetSccmWakeUpProxyOk returns a tuple with the SccmWakeUpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSccmWakeUpProxy

`func (o *EditHypervisorAzureConnectionRequestModel) SetSccmWakeUpProxy(v bool)`

SetSccmWakeUpProxy sets SccmWakeUpProxy field to given value.

### HasSccmWakeUpProxy

`func (o *EditHypervisorAzureConnectionRequestModel) HasSccmWakeUpProxy() bool`

HasSccmWakeUpProxy returns a boolean if a field has been set.

### SetSccmWakeUpProxyNil

`func (o *EditHypervisorAzureConnectionRequestModel) SetSccmWakeUpProxyNil(b bool)`

 SetSccmWakeUpProxyNil sets the value for SccmWakeUpProxy to be an explicit nil

### UnsetSccmWakeUpProxy
`func (o *EditHypervisorAzureConnectionRequestModel) UnsetSccmWakeUpProxy()`

UnsetSccmWakeUpProxy ensures that no value is present for SccmWakeUpProxy, not even an explicit nil
### GetWakeOnLanPackets

`func (o *EditHypervisorAzureConnectionRequestModel) GetWakeOnLanPackets() WakeOnLanTransmission`

GetWakeOnLanPackets returns the WakeOnLanPackets field if non-nil, zero value otherwise.

### GetWakeOnLanPacketsOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetWakeOnLanPacketsOk() (*WakeOnLanTransmission, bool)`

GetWakeOnLanPacketsOk returns a tuple with the WakeOnLanPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanPackets

`func (o *EditHypervisorAzureConnectionRequestModel) SetWakeOnLanPackets(v WakeOnLanTransmission)`

SetWakeOnLanPackets sets WakeOnLanPackets field to given value.

### HasWakeOnLanPackets

`func (o *EditHypervisorAzureConnectionRequestModel) HasWakeOnLanPackets() bool`

HasWakeOnLanPackets returns a boolean if a field has been set.

### GetServiceAccountCredential

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountCredential() string`

GetServiceAccountCredential returns the ServiceAccountCredential field if non-nil, zero value otherwise.

### GetServiceAccountCredentialOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountCredentialOk() (*string, bool)`

GetServiceAccountCredentialOk returns a tuple with the ServiceAccountCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredential

`func (o *EditHypervisorAzureConnectionRequestModel) SetServiceAccountCredential(v string)`

SetServiceAccountCredential sets ServiceAccountCredential field to given value.

### HasServiceAccountCredential

`func (o *EditHypervisorAzureConnectionRequestModel) HasServiceAccountCredential() bool`

HasServiceAccountCredential returns a boolean if a field has been set.

### GetServiceAccountCredentialFormat

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountCredentialFormat() IdentityPasswordFormat`

GetServiceAccountCredentialFormat returns the ServiceAccountCredentialFormat field if non-nil, zero value otherwise.

### GetServiceAccountCredentialFormatOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountCredentialFormatOk() (*IdentityPasswordFormat, bool)`

GetServiceAccountCredentialFormatOk returns a tuple with the ServiceAccountCredentialFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredentialFormat

`func (o *EditHypervisorAzureConnectionRequestModel) SetServiceAccountCredentialFormat(v IdentityPasswordFormat)`

SetServiceAccountCredentialFormat sets ServiceAccountCredentialFormat field to given value.

### HasServiceAccountCredentialFormat

`func (o *EditHypervisorAzureConnectionRequestModel) HasServiceAccountCredentialFormat() bool`

HasServiceAccountCredentialFormat returns a boolean if a field has been set.

### GetServiceAccountFingerprint

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountFingerprint() string`

GetServiceAccountFingerprint returns the ServiceAccountFingerprint field if non-nil, zero value otherwise.

### GetServiceAccountFingerprintOk

`func (o *EditHypervisorAzureConnectionRequestModel) GetServiceAccountFingerprintOk() (*string, bool)`

GetServiceAccountFingerprintOk returns a tuple with the ServiceAccountFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountFingerprint

`func (o *EditHypervisorAzureConnectionRequestModel) SetServiceAccountFingerprint(v string)`

SetServiceAccountFingerprint sets ServiceAccountFingerprint field to given value.

### HasServiceAccountFingerprint

`func (o *EditHypervisorAzureConnectionRequestModel) HasServiceAccountFingerprint() bool`

HasServiceAccountFingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


