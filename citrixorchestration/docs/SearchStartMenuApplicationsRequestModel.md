# SearchStartMenuApplicationsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationGroups** | Pointer to **[]string** | The Application Groups which the start menu applications will be browsed from. | [optional] 
**DeliveryGroups** | Pointer to **[]string** | The Delivery Groups which the start menu applications will be browsed from | [optional] 
**TagRestrictions** | Pointer to **[]string** | The Tag Restrictions, which means only machines with these tags can be used to browse start menu applications. | [optional] 

## Methods

### NewSearchStartMenuApplicationsRequestModel

`func NewSearchStartMenuApplicationsRequestModel() *SearchStartMenuApplicationsRequestModel`

NewSearchStartMenuApplicationsRequestModel instantiates a new SearchStartMenuApplicationsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStartMenuApplicationsRequestModelWithDefaults

`func NewSearchStartMenuApplicationsRequestModelWithDefaults() *SearchStartMenuApplicationsRequestModel`

NewSearchStartMenuApplicationsRequestModelWithDefaults instantiates a new SearchStartMenuApplicationsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationGroups

`func (o *SearchStartMenuApplicationsRequestModel) GetApplicationGroups() []string`

GetApplicationGroups returns the ApplicationGroups field if non-nil, zero value otherwise.

### GetApplicationGroupsOk

`func (o *SearchStartMenuApplicationsRequestModel) GetApplicationGroupsOk() (*[]string, bool)`

GetApplicationGroupsOk returns a tuple with the ApplicationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationGroups

`func (o *SearchStartMenuApplicationsRequestModel) SetApplicationGroups(v []string)`

SetApplicationGroups sets ApplicationGroups field to given value.

### HasApplicationGroups

`func (o *SearchStartMenuApplicationsRequestModel) HasApplicationGroups() bool`

HasApplicationGroups returns a boolean if a field has been set.

### GetDeliveryGroups

`func (o *SearchStartMenuApplicationsRequestModel) GetDeliveryGroups() []string`

GetDeliveryGroups returns the DeliveryGroups field if non-nil, zero value otherwise.

### GetDeliveryGroupsOk

`func (o *SearchStartMenuApplicationsRequestModel) GetDeliveryGroupsOk() (*[]string, bool)`

GetDeliveryGroupsOk returns a tuple with the DeliveryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroups

`func (o *SearchStartMenuApplicationsRequestModel) SetDeliveryGroups(v []string)`

SetDeliveryGroups sets DeliveryGroups field to given value.

### HasDeliveryGroups

`func (o *SearchStartMenuApplicationsRequestModel) HasDeliveryGroups() bool`

HasDeliveryGroups returns a boolean if a field has been set.

### GetTagRestrictions

`func (o *SearchStartMenuApplicationsRequestModel) GetTagRestrictions() []string`

GetTagRestrictions returns the TagRestrictions field if non-nil, zero value otherwise.

### GetTagRestrictionsOk

`func (o *SearchStartMenuApplicationsRequestModel) GetTagRestrictionsOk() (*[]string, bool)`

GetTagRestrictionsOk returns a tuple with the TagRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagRestrictions

`func (o *SearchStartMenuApplicationsRequestModel) SetTagRestrictions(v []string)`

SetTagRestrictions sets TagRestrictions field to given value.

### HasTagRestrictions

`func (o *SearchStartMenuApplicationsRequestModel) HasTagRestrictions() bool`

HasTagRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


