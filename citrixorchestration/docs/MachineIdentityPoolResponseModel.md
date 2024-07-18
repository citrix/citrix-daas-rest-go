# MachineIdentityPoolResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Machine identity pool id. | [optional] 
**Name** | Pointer to **NullableString** | Machine identity pool name. | [optional] 
**NamingSchemeType** | Pointer to [**AccountNamingSchemeType**](AccountNamingSchemeType.md) |  | [optional] 
**NamingScheme** | Pointer to **NullableString** | Machine identity pool naming scheme. | [optional] 
**StartCount** | Pointer to **int32** | Machine identity pool start count. | [optional] 
**CustomActiveDirectoryOU** | Pointer to **NullableString** | Machine identity pool custom active directory OU. | [optional] 
**UseDefaultOU** | Pointer to **bool** | Indicates whether use default OU | [optional] 
**DefaultOUDomain** | Pointer to [**ADDomainResponseModel**](ADDomainResponseModel.md) |  | [optional] 
**AvailableAccountsCount** | Pointer to **int32** | The number of available accounts in the machine identity pool | [optional] 
**InUseAccountsCount** | Pointer to **int32** | The number of accounts in use in the machine identity pool | [optional] 
**TaintedAccountsCount** | Pointer to **int32** | The number of tainted accounts in the machine identity pool | [optional] 
**ErrorAccountsCount** | Pointer to **int32** | The number of bad accounts in the machine identity pool | [optional] 
**TenantId** | Pointer to **NullableString** | Tenant id. | [optional] 
**WorkGroupMachines** | Pointer to **bool** | Work group machines. | [optional] 
**IdentityType** | Pointer to **NullableString** | Identity type. | [optional] 
**IdentityContent** | Pointer to **NullableString** | Identity content. | [optional] 
**AzureADSecurityGroupName** | Pointer to **NullableString** | Azure AD security group name. | [optional] 
**AzureADAccessToken** | Pointer to **NullableString** | Azure AD access token. | [optional] 
**DeviceManagementType** | Pointer to **NullableString** | Device management type. | [optional] 
**AzureADTenantId** | Pointer to **NullableString** | Azure AD tenant id. | [optional] 
**ServiceAccountUid** | Pointer to **[]string** | Service account Uids associated with this IdentityPool | [optional] 

## Methods

### NewMachineIdentityPoolResponseModel

`func NewMachineIdentityPoolResponseModel() *MachineIdentityPoolResponseModel`

NewMachineIdentityPoolResponseModel instantiates a new MachineIdentityPoolResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineIdentityPoolResponseModelWithDefaults

`func NewMachineIdentityPoolResponseModelWithDefaults() *MachineIdentityPoolResponseModel`

NewMachineIdentityPoolResponseModelWithDefaults instantiates a new MachineIdentityPoolResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MachineIdentityPoolResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineIdentityPoolResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineIdentityPoolResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MachineIdentityPoolResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MachineIdentityPoolResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineIdentityPoolResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineIdentityPoolResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineIdentityPoolResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MachineIdentityPoolResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MachineIdentityPoolResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamingSchemeType

`func (o *MachineIdentityPoolResponseModel) GetNamingSchemeType() AccountNamingSchemeType`

GetNamingSchemeType returns the NamingSchemeType field if non-nil, zero value otherwise.

### GetNamingSchemeTypeOk

`func (o *MachineIdentityPoolResponseModel) GetNamingSchemeTypeOk() (*AccountNamingSchemeType, bool)`

GetNamingSchemeTypeOk returns a tuple with the NamingSchemeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchemeType

`func (o *MachineIdentityPoolResponseModel) SetNamingSchemeType(v AccountNamingSchemeType)`

SetNamingSchemeType sets NamingSchemeType field to given value.

### HasNamingSchemeType

`func (o *MachineIdentityPoolResponseModel) HasNamingSchemeType() bool`

HasNamingSchemeType returns a boolean if a field has been set.

### GetNamingScheme

`func (o *MachineIdentityPoolResponseModel) GetNamingScheme() string`

GetNamingScheme returns the NamingScheme field if non-nil, zero value otherwise.

### GetNamingSchemeOk

`func (o *MachineIdentityPoolResponseModel) GetNamingSchemeOk() (*string, bool)`

GetNamingSchemeOk returns a tuple with the NamingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingScheme

`func (o *MachineIdentityPoolResponseModel) SetNamingScheme(v string)`

SetNamingScheme sets NamingScheme field to given value.

### HasNamingScheme

`func (o *MachineIdentityPoolResponseModel) HasNamingScheme() bool`

HasNamingScheme returns a boolean if a field has been set.

### SetNamingSchemeNil

`func (o *MachineIdentityPoolResponseModel) SetNamingSchemeNil(b bool)`

 SetNamingSchemeNil sets the value for NamingScheme to be an explicit nil

### UnsetNamingScheme
`func (o *MachineIdentityPoolResponseModel) UnsetNamingScheme()`

UnsetNamingScheme ensures that no value is present for NamingScheme, not even an explicit nil
### GetStartCount

`func (o *MachineIdentityPoolResponseModel) GetStartCount() int32`

GetStartCount returns the StartCount field if non-nil, zero value otherwise.

### GetStartCountOk

`func (o *MachineIdentityPoolResponseModel) GetStartCountOk() (*int32, bool)`

GetStartCountOk returns a tuple with the StartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCount

`func (o *MachineIdentityPoolResponseModel) SetStartCount(v int32)`

SetStartCount sets StartCount field to given value.

### HasStartCount

`func (o *MachineIdentityPoolResponseModel) HasStartCount() bool`

HasStartCount returns a boolean if a field has been set.

### GetCustomActiveDirectoryOU

`func (o *MachineIdentityPoolResponseModel) GetCustomActiveDirectoryOU() string`

GetCustomActiveDirectoryOU returns the CustomActiveDirectoryOU field if non-nil, zero value otherwise.

### GetCustomActiveDirectoryOUOk

`func (o *MachineIdentityPoolResponseModel) GetCustomActiveDirectoryOUOk() (*string, bool)`

GetCustomActiveDirectoryOUOk returns a tuple with the CustomActiveDirectoryOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomActiveDirectoryOU

`func (o *MachineIdentityPoolResponseModel) SetCustomActiveDirectoryOU(v string)`

SetCustomActiveDirectoryOU sets CustomActiveDirectoryOU field to given value.

### HasCustomActiveDirectoryOU

`func (o *MachineIdentityPoolResponseModel) HasCustomActiveDirectoryOU() bool`

HasCustomActiveDirectoryOU returns a boolean if a field has been set.

### SetCustomActiveDirectoryOUNil

`func (o *MachineIdentityPoolResponseModel) SetCustomActiveDirectoryOUNil(b bool)`

 SetCustomActiveDirectoryOUNil sets the value for CustomActiveDirectoryOU to be an explicit nil

### UnsetCustomActiveDirectoryOU
`func (o *MachineIdentityPoolResponseModel) UnsetCustomActiveDirectoryOU()`

UnsetCustomActiveDirectoryOU ensures that no value is present for CustomActiveDirectoryOU, not even an explicit nil
### GetUseDefaultOU

`func (o *MachineIdentityPoolResponseModel) GetUseDefaultOU() bool`

GetUseDefaultOU returns the UseDefaultOU field if non-nil, zero value otherwise.

### GetUseDefaultOUOk

`func (o *MachineIdentityPoolResponseModel) GetUseDefaultOUOk() (*bool, bool)`

GetUseDefaultOUOk returns a tuple with the UseDefaultOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultOU

`func (o *MachineIdentityPoolResponseModel) SetUseDefaultOU(v bool)`

SetUseDefaultOU sets UseDefaultOU field to given value.

### HasUseDefaultOU

`func (o *MachineIdentityPoolResponseModel) HasUseDefaultOU() bool`

HasUseDefaultOU returns a boolean if a field has been set.

### GetDefaultOUDomain

`func (o *MachineIdentityPoolResponseModel) GetDefaultOUDomain() ADDomainResponseModel`

GetDefaultOUDomain returns the DefaultOUDomain field if non-nil, zero value otherwise.

### GetDefaultOUDomainOk

`func (o *MachineIdentityPoolResponseModel) GetDefaultOUDomainOk() (*ADDomainResponseModel, bool)`

GetDefaultOUDomainOk returns a tuple with the DefaultOUDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOUDomain

`func (o *MachineIdentityPoolResponseModel) SetDefaultOUDomain(v ADDomainResponseModel)`

SetDefaultOUDomain sets DefaultOUDomain field to given value.

### HasDefaultOUDomain

`func (o *MachineIdentityPoolResponseModel) HasDefaultOUDomain() bool`

HasDefaultOUDomain returns a boolean if a field has been set.

### GetAvailableAccountsCount

`func (o *MachineIdentityPoolResponseModel) GetAvailableAccountsCount() int32`

GetAvailableAccountsCount returns the AvailableAccountsCount field if non-nil, zero value otherwise.

### GetAvailableAccountsCountOk

`func (o *MachineIdentityPoolResponseModel) GetAvailableAccountsCountOk() (*int32, bool)`

GetAvailableAccountsCountOk returns a tuple with the AvailableAccountsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAccountsCount

`func (o *MachineIdentityPoolResponseModel) SetAvailableAccountsCount(v int32)`

SetAvailableAccountsCount sets AvailableAccountsCount field to given value.

### HasAvailableAccountsCount

`func (o *MachineIdentityPoolResponseModel) HasAvailableAccountsCount() bool`

HasAvailableAccountsCount returns a boolean if a field has been set.

### GetInUseAccountsCount

`func (o *MachineIdentityPoolResponseModel) GetInUseAccountsCount() int32`

GetInUseAccountsCount returns the InUseAccountsCount field if non-nil, zero value otherwise.

### GetInUseAccountsCountOk

`func (o *MachineIdentityPoolResponseModel) GetInUseAccountsCountOk() (*int32, bool)`

GetInUseAccountsCountOk returns a tuple with the InUseAccountsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUseAccountsCount

`func (o *MachineIdentityPoolResponseModel) SetInUseAccountsCount(v int32)`

SetInUseAccountsCount sets InUseAccountsCount field to given value.

### HasInUseAccountsCount

`func (o *MachineIdentityPoolResponseModel) HasInUseAccountsCount() bool`

HasInUseAccountsCount returns a boolean if a field has been set.

### GetTaintedAccountsCount

`func (o *MachineIdentityPoolResponseModel) GetTaintedAccountsCount() int32`

GetTaintedAccountsCount returns the TaintedAccountsCount field if non-nil, zero value otherwise.

### GetTaintedAccountsCountOk

`func (o *MachineIdentityPoolResponseModel) GetTaintedAccountsCountOk() (*int32, bool)`

GetTaintedAccountsCountOk returns a tuple with the TaintedAccountsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaintedAccountsCount

`func (o *MachineIdentityPoolResponseModel) SetTaintedAccountsCount(v int32)`

SetTaintedAccountsCount sets TaintedAccountsCount field to given value.

### HasTaintedAccountsCount

`func (o *MachineIdentityPoolResponseModel) HasTaintedAccountsCount() bool`

HasTaintedAccountsCount returns a boolean if a field has been set.

### GetErrorAccountsCount

`func (o *MachineIdentityPoolResponseModel) GetErrorAccountsCount() int32`

GetErrorAccountsCount returns the ErrorAccountsCount field if non-nil, zero value otherwise.

### GetErrorAccountsCountOk

`func (o *MachineIdentityPoolResponseModel) GetErrorAccountsCountOk() (*int32, bool)`

GetErrorAccountsCountOk returns a tuple with the ErrorAccountsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorAccountsCount

`func (o *MachineIdentityPoolResponseModel) SetErrorAccountsCount(v int32)`

SetErrorAccountsCount sets ErrorAccountsCount field to given value.

### HasErrorAccountsCount

`func (o *MachineIdentityPoolResponseModel) HasErrorAccountsCount() bool`

HasErrorAccountsCount returns a boolean if a field has been set.

### GetTenantId

`func (o *MachineIdentityPoolResponseModel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MachineIdentityPoolResponseModel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MachineIdentityPoolResponseModel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MachineIdentityPoolResponseModel) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MachineIdentityPoolResponseModel) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MachineIdentityPoolResponseModel) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetWorkGroupMachines

`func (o *MachineIdentityPoolResponseModel) GetWorkGroupMachines() bool`

GetWorkGroupMachines returns the WorkGroupMachines field if non-nil, zero value otherwise.

### GetWorkGroupMachinesOk

`func (o *MachineIdentityPoolResponseModel) GetWorkGroupMachinesOk() (*bool, bool)`

GetWorkGroupMachinesOk returns a tuple with the WorkGroupMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkGroupMachines

`func (o *MachineIdentityPoolResponseModel) SetWorkGroupMachines(v bool)`

SetWorkGroupMachines sets WorkGroupMachines field to given value.

### HasWorkGroupMachines

`func (o *MachineIdentityPoolResponseModel) HasWorkGroupMachines() bool`

HasWorkGroupMachines returns a boolean if a field has been set.

### GetIdentityType

`func (o *MachineIdentityPoolResponseModel) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *MachineIdentityPoolResponseModel) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *MachineIdentityPoolResponseModel) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *MachineIdentityPoolResponseModel) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *MachineIdentityPoolResponseModel) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *MachineIdentityPoolResponseModel) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
### GetIdentityContent

`func (o *MachineIdentityPoolResponseModel) GetIdentityContent() string`

GetIdentityContent returns the IdentityContent field if non-nil, zero value otherwise.

### GetIdentityContentOk

`func (o *MachineIdentityPoolResponseModel) GetIdentityContentOk() (*string, bool)`

GetIdentityContentOk returns a tuple with the IdentityContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityContent

`func (o *MachineIdentityPoolResponseModel) SetIdentityContent(v string)`

SetIdentityContent sets IdentityContent field to given value.

### HasIdentityContent

`func (o *MachineIdentityPoolResponseModel) HasIdentityContent() bool`

HasIdentityContent returns a boolean if a field has been set.

### SetIdentityContentNil

`func (o *MachineIdentityPoolResponseModel) SetIdentityContentNil(b bool)`

 SetIdentityContentNil sets the value for IdentityContent to be an explicit nil

### UnsetIdentityContent
`func (o *MachineIdentityPoolResponseModel) UnsetIdentityContent()`

UnsetIdentityContent ensures that no value is present for IdentityContent, not even an explicit nil
### GetAzureADSecurityGroupName

`func (o *MachineIdentityPoolResponseModel) GetAzureADSecurityGroupName() string`

GetAzureADSecurityGroupName returns the AzureADSecurityGroupName field if non-nil, zero value otherwise.

### GetAzureADSecurityGroupNameOk

`func (o *MachineIdentityPoolResponseModel) GetAzureADSecurityGroupNameOk() (*string, bool)`

GetAzureADSecurityGroupNameOk returns a tuple with the AzureADSecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADSecurityGroupName

`func (o *MachineIdentityPoolResponseModel) SetAzureADSecurityGroupName(v string)`

SetAzureADSecurityGroupName sets AzureADSecurityGroupName field to given value.

### HasAzureADSecurityGroupName

`func (o *MachineIdentityPoolResponseModel) HasAzureADSecurityGroupName() bool`

HasAzureADSecurityGroupName returns a boolean if a field has been set.

### SetAzureADSecurityGroupNameNil

`func (o *MachineIdentityPoolResponseModel) SetAzureADSecurityGroupNameNil(b bool)`

 SetAzureADSecurityGroupNameNil sets the value for AzureADSecurityGroupName to be an explicit nil

### UnsetAzureADSecurityGroupName
`func (o *MachineIdentityPoolResponseModel) UnsetAzureADSecurityGroupName()`

UnsetAzureADSecurityGroupName ensures that no value is present for AzureADSecurityGroupName, not even an explicit nil
### GetAzureADAccessToken

`func (o *MachineIdentityPoolResponseModel) GetAzureADAccessToken() string`

GetAzureADAccessToken returns the AzureADAccessToken field if non-nil, zero value otherwise.

### GetAzureADAccessTokenOk

`func (o *MachineIdentityPoolResponseModel) GetAzureADAccessTokenOk() (*string, bool)`

GetAzureADAccessTokenOk returns a tuple with the AzureADAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADAccessToken

`func (o *MachineIdentityPoolResponseModel) SetAzureADAccessToken(v string)`

SetAzureADAccessToken sets AzureADAccessToken field to given value.

### HasAzureADAccessToken

`func (o *MachineIdentityPoolResponseModel) HasAzureADAccessToken() bool`

HasAzureADAccessToken returns a boolean if a field has been set.

### SetAzureADAccessTokenNil

`func (o *MachineIdentityPoolResponseModel) SetAzureADAccessTokenNil(b bool)`

 SetAzureADAccessTokenNil sets the value for AzureADAccessToken to be an explicit nil

### UnsetAzureADAccessToken
`func (o *MachineIdentityPoolResponseModel) UnsetAzureADAccessToken()`

UnsetAzureADAccessToken ensures that no value is present for AzureADAccessToken, not even an explicit nil
### GetDeviceManagementType

`func (o *MachineIdentityPoolResponseModel) GetDeviceManagementType() string`

GetDeviceManagementType returns the DeviceManagementType field if non-nil, zero value otherwise.

### GetDeviceManagementTypeOk

`func (o *MachineIdentityPoolResponseModel) GetDeviceManagementTypeOk() (*string, bool)`

GetDeviceManagementTypeOk returns a tuple with the DeviceManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManagementType

`func (o *MachineIdentityPoolResponseModel) SetDeviceManagementType(v string)`

SetDeviceManagementType sets DeviceManagementType field to given value.

### HasDeviceManagementType

`func (o *MachineIdentityPoolResponseModel) HasDeviceManagementType() bool`

HasDeviceManagementType returns a boolean if a field has been set.

### SetDeviceManagementTypeNil

`func (o *MachineIdentityPoolResponseModel) SetDeviceManagementTypeNil(b bool)`

 SetDeviceManagementTypeNil sets the value for DeviceManagementType to be an explicit nil

### UnsetDeviceManagementType
`func (o *MachineIdentityPoolResponseModel) UnsetDeviceManagementType()`

UnsetDeviceManagementType ensures that no value is present for DeviceManagementType, not even an explicit nil
### GetAzureADTenantId

`func (o *MachineIdentityPoolResponseModel) GetAzureADTenantId() string`

GetAzureADTenantId returns the AzureADTenantId field if non-nil, zero value otherwise.

### GetAzureADTenantIdOk

`func (o *MachineIdentityPoolResponseModel) GetAzureADTenantIdOk() (*string, bool)`

GetAzureADTenantIdOk returns a tuple with the AzureADTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADTenantId

`func (o *MachineIdentityPoolResponseModel) SetAzureADTenantId(v string)`

SetAzureADTenantId sets AzureADTenantId field to given value.

### HasAzureADTenantId

`func (o *MachineIdentityPoolResponseModel) HasAzureADTenantId() bool`

HasAzureADTenantId returns a boolean if a field has been set.

### SetAzureADTenantIdNil

`func (o *MachineIdentityPoolResponseModel) SetAzureADTenantIdNil(b bool)`

 SetAzureADTenantIdNil sets the value for AzureADTenantId to be an explicit nil

### UnsetAzureADTenantId
`func (o *MachineIdentityPoolResponseModel) UnsetAzureADTenantId()`

UnsetAzureADTenantId ensures that no value is present for AzureADTenantId, not even an explicit nil
### GetServiceAccountUid

`func (o *MachineIdentityPoolResponseModel) GetServiceAccountUid() []string`

GetServiceAccountUid returns the ServiceAccountUid field if non-nil, zero value otherwise.

### GetServiceAccountUidOk

`func (o *MachineIdentityPoolResponseModel) GetServiceAccountUidOk() (*[]string, bool)`

GetServiceAccountUidOk returns a tuple with the ServiceAccountUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountUid

`func (o *MachineIdentityPoolResponseModel) SetServiceAccountUid(v []string)`

SetServiceAccountUid sets ServiceAccountUid field to given value.

### HasServiceAccountUid

`func (o *MachineIdentityPoolResponseModel) HasServiceAccountUid() bool`

HasServiceAccountUid returns a boolean if a field has been set.

### SetServiceAccountUidNil

`func (o *MachineIdentityPoolResponseModel) SetServiceAccountUidNil(b bool)`

 SetServiceAccountUidNil sets the value for ServiceAccountUid to be an explicit nil

### UnsetServiceAccountUid
`func (o *MachineIdentityPoolResponseModel) UnsetServiceAccountUid()`

UnsetServiceAccountUid ensures that no value is present for ServiceAccountUid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


