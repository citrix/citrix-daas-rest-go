# SiteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Identity of site | [optional] [readonly] 
**Name** | Pointer to **string** | Name of site | [optional] 
**Description** | Pointer to **string** | Description of site | [optional] 

## Methods

### NewSiteModel

`func NewSiteModel() *SiteModel`

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

`func (o *SiteModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SiteModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SiteModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SiteModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SiteModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SiteModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SiteModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


