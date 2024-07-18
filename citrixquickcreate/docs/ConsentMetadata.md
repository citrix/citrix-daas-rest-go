# ConsentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **NullableString** | Client/Application for MT AppRegistration for CloudPC Registration service  This is the app that we would get consent for | [optional] 
**Scope** | Pointer to **NullableString** | Scope for consent request  .default would get consent for all permissions listed in AppRegistration APIPermissions section | [optional] 

## Methods

### NewConsentMetadata

`func NewConsentMetadata() *ConsentMetadata`

NewConsentMetadata instantiates a new ConsentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentMetadataWithDefaults

`func NewConsentMetadataWithDefaults() *ConsentMetadata`

NewConsentMetadataWithDefaults instantiates a new ConsentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ConsentMetadata) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ConsentMetadata) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ConsentMetadata) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ConsentMetadata) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *ConsentMetadata) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *ConsentMetadata) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetScope

`func (o *ConsentMetadata) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ConsentMetadata) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ConsentMetadata) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ConsentMetadata) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *ConsentMetadata) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *ConsentMetadata) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


