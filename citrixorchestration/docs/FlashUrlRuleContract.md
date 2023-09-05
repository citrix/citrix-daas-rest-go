# FlashUrlRuleContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RenderAction** | Pointer to [**FlashUrlRenderAction**](FlashUrlRenderAction.md) |  | [optional] 
**UrlPattern** | Pointer to **string** |  | [optional] 
**FlashInstance** | Pointer to **string** |  | [optional] 

## Methods

### NewFlashUrlRuleContract

`func NewFlashUrlRuleContract() *FlashUrlRuleContract`

NewFlashUrlRuleContract instantiates a new FlashUrlRuleContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlashUrlRuleContractWithDefaults

`func NewFlashUrlRuleContractWithDefaults() *FlashUrlRuleContract`

NewFlashUrlRuleContractWithDefaults instantiates a new FlashUrlRuleContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenderAction

`func (o *FlashUrlRuleContract) GetRenderAction() FlashUrlRenderAction`

GetRenderAction returns the RenderAction field if non-nil, zero value otherwise.

### GetRenderActionOk

`func (o *FlashUrlRuleContract) GetRenderActionOk() (*FlashUrlRenderAction, bool)`

GetRenderActionOk returns a tuple with the RenderAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderAction

`func (o *FlashUrlRuleContract) SetRenderAction(v FlashUrlRenderAction)`

SetRenderAction sets RenderAction field to given value.

### HasRenderAction

`func (o *FlashUrlRuleContract) HasRenderAction() bool`

HasRenderAction returns a boolean if a field has been set.

### GetUrlPattern

`func (o *FlashUrlRuleContract) GetUrlPattern() string`

GetUrlPattern returns the UrlPattern field if non-nil, zero value otherwise.

### GetUrlPatternOk

`func (o *FlashUrlRuleContract) GetUrlPatternOk() (*string, bool)`

GetUrlPatternOk returns a tuple with the UrlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPattern

`func (o *FlashUrlRuleContract) SetUrlPattern(v string)`

SetUrlPattern sets UrlPattern field to given value.

### HasUrlPattern

`func (o *FlashUrlRuleContract) HasUrlPattern() bool`

HasUrlPattern returns a boolean if a field has been set.

### GetFlashInstance

`func (o *FlashUrlRuleContract) GetFlashInstance() string`

GetFlashInstance returns the FlashInstance field if non-nil, zero value otherwise.

### GetFlashInstanceOk

`func (o *FlashUrlRuleContract) GetFlashInstanceOk() (*string, bool)`

GetFlashInstanceOk returns a tuple with the FlashInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashInstance

`func (o *FlashUrlRuleContract) SetFlashInstance(v string)`

SetFlashInstance sets FlashInstance field to given value.

### HasFlashInstance

`func (o *FlashUrlRuleContract) HasFlashInstance() bool`

HasFlashInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


