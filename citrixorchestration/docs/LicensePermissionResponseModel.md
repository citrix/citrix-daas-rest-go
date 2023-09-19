# LicensePermissionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicensingAdministrator** | Pointer to **NullableString** | The administrator  for the license server | [optional] 
**LicensingPermissionLevel** | Pointer to [**LicensingPermissionLevel**](LicensingPermissionLevel.md) |  | [optional] 

## Methods

### NewLicensePermissionResponseModel

`func NewLicensePermissionResponseModel() *LicensePermissionResponseModel`

NewLicensePermissionResponseModel instantiates a new LicensePermissionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensePermissionResponseModelWithDefaults

`func NewLicensePermissionResponseModelWithDefaults() *LicensePermissionResponseModel`

NewLicensePermissionResponseModelWithDefaults instantiates a new LicensePermissionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicensingAdministrator

`func (o *LicensePermissionResponseModel) GetLicensingAdministrator() string`

GetLicensingAdministrator returns the LicensingAdministrator field if non-nil, zero value otherwise.

### GetLicensingAdministratorOk

`func (o *LicensePermissionResponseModel) GetLicensingAdministratorOk() (*string, bool)`

GetLicensingAdministratorOk returns a tuple with the LicensingAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingAdministrator

`func (o *LicensePermissionResponseModel) SetLicensingAdministrator(v string)`

SetLicensingAdministrator sets LicensingAdministrator field to given value.

### HasLicensingAdministrator

`func (o *LicensePermissionResponseModel) HasLicensingAdministrator() bool`

HasLicensingAdministrator returns a boolean if a field has been set.

### SetLicensingAdministratorNil

`func (o *LicensePermissionResponseModel) SetLicensingAdministratorNil(b bool)`

 SetLicensingAdministratorNil sets the value for LicensingAdministrator to be an explicit nil

### UnsetLicensingAdministrator
`func (o *LicensePermissionResponseModel) UnsetLicensingAdministrator()`

UnsetLicensingAdministrator ensures that no value is present for LicensingAdministrator, not even an explicit nil
### GetLicensingPermissionLevel

`func (o *LicensePermissionResponseModel) GetLicensingPermissionLevel() LicensingPermissionLevel`

GetLicensingPermissionLevel returns the LicensingPermissionLevel field if non-nil, zero value otherwise.

### GetLicensingPermissionLevelOk

`func (o *LicensePermissionResponseModel) GetLicensingPermissionLevelOk() (*LicensingPermissionLevel, bool)`

GetLicensingPermissionLevelOk returns a tuple with the LicensingPermissionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingPermissionLevel

`func (o *LicensePermissionResponseModel) SetLicensingPermissionLevel(v LicensingPermissionLevel)`

SetLicensingPermissionLevel sets LicensingPermissionLevel field to given value.

### HasLicensingPermissionLevel

`func (o *LicensePermissionResponseModel) HasLicensingPermissionLevel() bool`

HasLicensingPermissionLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


