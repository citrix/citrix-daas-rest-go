# InitiateDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**ConnectionId** | **string** | Hosting Connection Id | 
**DeploymentName** | **string** |  | 

## Methods

### NewInitiateDeployment

`func NewInitiateDeployment(accountType AccountType, connectionId string, deploymentName string, ) *InitiateDeployment`

NewInitiateDeployment instantiates a new InitiateDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateDeploymentWithDefaults

`func NewInitiateDeploymentWithDefaults() *InitiateDeployment`

NewInitiateDeploymentWithDefaults instantiates a new InitiateDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *InitiateDeployment) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *InitiateDeployment) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *InitiateDeployment) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetConnectionId

`func (o *InitiateDeployment) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *InitiateDeployment) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *InitiateDeployment) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetDeploymentName

`func (o *InitiateDeployment) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *InitiateDeployment) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *InitiateDeployment) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


