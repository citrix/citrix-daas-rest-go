# ContainerScopeResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | [**[]ScopeResponseModel**](ScopeResponseModel.md) | List of delegated admin scopes. | 
**ScopeType** | [**ContainerScopeType**](ContainerScopeType.md) |  | 

## Methods

### NewContainerScopeResponseModel

`func NewContainerScopeResponseModel(scopes []ScopeResponseModel, scopeType ContainerScopeType, ) *ContainerScopeResponseModel`

NewContainerScopeResponseModel instantiates a new ContainerScopeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerScopeResponseModelWithDefaults

`func NewContainerScopeResponseModelWithDefaults() *ContainerScopeResponseModel`

NewContainerScopeResponseModelWithDefaults instantiates a new ContainerScopeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *ContainerScopeResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ContainerScopeResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ContainerScopeResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.


### GetScopeType

`func (o *ContainerScopeResponseModel) GetScopeType() ContainerScopeType`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *ContainerScopeResponseModel) GetScopeTypeOk() (*ContainerScopeType, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *ContainerScopeResponseModel) SetScopeType(v ContainerScopeType)`

SetScopeType sets ScopeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


