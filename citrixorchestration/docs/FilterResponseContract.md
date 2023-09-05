# FilterResponseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** | The filter type. | [optional] 
**Identifier** | Pointer to **string** | The filter name. | [optional] 
**DisplayName** | Pointer to **string** | Filter display name, this is translated and unique. | [optional] 
**IsSingleton** | Pointer to **bool** | Is filter a singleton, only the NetScaler SD-Wan filter is singleton. | [optional] 
**Explanation** | Pointer to **string** | Translated filter explanation. | [optional] 
**Synopsis** | Pointer to **string** | A summary of filter value. | [optional] 
**IsUserFilter** | Pointer to **bool** | True &#x3D; user filter, False &#x3D; machine filter | [optional] 
**IsAllowed** | Pointer to **bool** | Allow or deny | [optional] 
**IsEnabled** | Pointer to **bool** | Enabled or disabled | [optional] 
**FilterValue** | Pointer to **string** | Serialized JSON string. | [optional] 

## Methods

### NewFilterResponseContract

`func NewFilterResponseContract() *FilterResponseContract`

NewFilterResponseContract instantiates a new FilterResponseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterResponseContractWithDefaults

`func NewFilterResponseContractWithDefaults() *FilterResponseContract`

NewFilterResponseContractWithDefaults instantiates a new FilterResponseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *FilterResponseContract) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *FilterResponseContract) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *FilterResponseContract) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *FilterResponseContract) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetIdentifier

`func (o *FilterResponseContract) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *FilterResponseContract) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *FilterResponseContract) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *FilterResponseContract) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDisplayName

`func (o *FilterResponseContract) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FilterResponseContract) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FilterResponseContract) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FilterResponseContract) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsSingleton

`func (o *FilterResponseContract) GetIsSingleton() bool`

GetIsSingleton returns the IsSingleton field if non-nil, zero value otherwise.

### GetIsSingletonOk

`func (o *FilterResponseContract) GetIsSingletonOk() (*bool, bool)`

GetIsSingletonOk returns a tuple with the IsSingleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleton

`func (o *FilterResponseContract) SetIsSingleton(v bool)`

SetIsSingleton sets IsSingleton field to given value.

### HasIsSingleton

`func (o *FilterResponseContract) HasIsSingleton() bool`

HasIsSingleton returns a boolean if a field has been set.

### GetExplanation

`func (o *FilterResponseContract) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *FilterResponseContract) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *FilterResponseContract) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *FilterResponseContract) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetSynopsis

`func (o *FilterResponseContract) GetSynopsis() string`

GetSynopsis returns the Synopsis field if non-nil, zero value otherwise.

### GetSynopsisOk

`func (o *FilterResponseContract) GetSynopsisOk() (*string, bool)`

GetSynopsisOk returns a tuple with the Synopsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynopsis

`func (o *FilterResponseContract) SetSynopsis(v string)`

SetSynopsis sets Synopsis field to given value.

### HasSynopsis

`func (o *FilterResponseContract) HasSynopsis() bool`

HasSynopsis returns a boolean if a field has been set.

### GetIsUserFilter

`func (o *FilterResponseContract) GetIsUserFilter() bool`

GetIsUserFilter returns the IsUserFilter field if non-nil, zero value otherwise.

### GetIsUserFilterOk

`func (o *FilterResponseContract) GetIsUserFilterOk() (*bool, bool)`

GetIsUserFilterOk returns a tuple with the IsUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserFilter

`func (o *FilterResponseContract) SetIsUserFilter(v bool)`

SetIsUserFilter sets IsUserFilter field to given value.

### HasIsUserFilter

`func (o *FilterResponseContract) HasIsUserFilter() bool`

HasIsUserFilter returns a boolean if a field has been set.

### GetIsAllowed

`func (o *FilterResponseContract) GetIsAllowed() bool`

GetIsAllowed returns the IsAllowed field if non-nil, zero value otherwise.

### GetIsAllowedOk

`func (o *FilterResponseContract) GetIsAllowedOk() (*bool, bool)`

GetIsAllowedOk returns a tuple with the IsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowed

`func (o *FilterResponseContract) SetIsAllowed(v bool)`

SetIsAllowed sets IsAllowed field to given value.

### HasIsAllowed

`func (o *FilterResponseContract) HasIsAllowed() bool`

HasIsAllowed returns a boolean if a field has been set.

### GetIsEnabled

`func (o *FilterResponseContract) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FilterResponseContract) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FilterResponseContract) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FilterResponseContract) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetFilterValue

`func (o *FilterResponseContract) GetFilterValue() string`

GetFilterValue returns the FilterValue field if non-nil, zero value otherwise.

### GetFilterValueOk

`func (o *FilterResponseContract) GetFilterValueOk() (*string, bool)`

GetFilterValueOk returns a tuple with the FilterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValue

`func (o *FilterResponseContract) SetFilterValue(v string)`

SetFilterValue sets FilterValue field to given value.

### HasFilterValue

`func (o *FilterResponseContract) HasFilterValue() bool`

HasFilterValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


