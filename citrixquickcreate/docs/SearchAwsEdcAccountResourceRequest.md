# SearchAwsEdcAccountResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to [**AwsAccountResourceType**](AwsAccountResourceType.md) |  | [optional] 
**AwsRoleArn** | Pointer to **NullableString** | The ARN of the role to assume when searching for resources in an account that has not been added yet | [optional] 
**AwsRegion** | Pointer to **NullableString** | The AWS region to use | [optional] 

## Methods

### NewSearchAwsEdcAccountResourceRequest

`func NewSearchAwsEdcAccountResourceRequest() *SearchAwsEdcAccountResourceRequest`

NewSearchAwsEdcAccountResourceRequest instantiates a new SearchAwsEdcAccountResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAwsEdcAccountResourceRequestWithDefaults

`func NewSearchAwsEdcAccountResourceRequestWithDefaults() *SearchAwsEdcAccountResourceRequest`

NewSearchAwsEdcAccountResourceRequestWithDefaults instantiates a new SearchAwsEdcAccountResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *SearchAwsEdcAccountResourceRequest) GetResourceType() AwsAccountResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *SearchAwsEdcAccountResourceRequest) GetResourceTypeOk() (*AwsAccountResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *SearchAwsEdcAccountResourceRequest) SetResourceType(v AwsAccountResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *SearchAwsEdcAccountResourceRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetAwsRoleArn

`func (o *SearchAwsEdcAccountResourceRequest) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *SearchAwsEdcAccountResourceRequest) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *SearchAwsEdcAccountResourceRequest) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.

### HasAwsRoleArn

`func (o *SearchAwsEdcAccountResourceRequest) HasAwsRoleArn() bool`

HasAwsRoleArn returns a boolean if a field has been set.

### SetAwsRoleArnNil

`func (o *SearchAwsEdcAccountResourceRequest) SetAwsRoleArnNil(b bool)`

 SetAwsRoleArnNil sets the value for AwsRoleArn to be an explicit nil

### UnsetAwsRoleArn
`func (o *SearchAwsEdcAccountResourceRequest) UnsetAwsRoleArn()`

UnsetAwsRoleArn ensures that no value is present for AwsRoleArn, not even an explicit nil
### GetAwsRegion

`func (o *SearchAwsEdcAccountResourceRequest) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *SearchAwsEdcAccountResourceRequest) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *SearchAwsEdcAccountResourceRequest) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *SearchAwsEdcAccountResourceRequest) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *SearchAwsEdcAccountResourceRequest) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *SearchAwsEdcAccountResourceRequest) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


