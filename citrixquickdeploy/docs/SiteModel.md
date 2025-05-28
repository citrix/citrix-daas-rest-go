# SiteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The virtual site id of the customer | 
**DisplayName** | Pointer to **string** | The display name of the virtual site | [optional] 

## Methods

### NewSiteModel

`func NewSiteModel(id string, ) *SiteModel`

NewSiteModel instantiates a new SiteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteModelWithDefaults

`func NewSiteModelWithDefaults() *SiteModel`

NewSiteModelWithDefaults instantiates a new SiteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteModel) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *SiteModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SiteModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SiteModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SiteModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


