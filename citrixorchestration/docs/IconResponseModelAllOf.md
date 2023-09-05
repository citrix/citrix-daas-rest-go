# IconResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the icon. | 
**IsBuiltIn** | **bool** | Indicates if the icon is built-in (i.e. Citrix-provided).  The built-in icons can be used for applications or desktops that don&#39;t have their own icons. | 
**Uid** | Pointer to **int32** | DEPRECATED.  Use Id. | [optional] 

## Methods

### NewIconResponseModelAllOf

`func NewIconResponseModelAllOf(id string, isBuiltIn bool, ) *IconResponseModelAllOf`

NewIconResponseModelAllOf instantiates a new IconResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIconResponseModelAllOfWithDefaults

`func NewIconResponseModelAllOfWithDefaults() *IconResponseModelAllOf`

NewIconResponseModelAllOfWithDefaults instantiates a new IconResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IconResponseModelAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IconResponseModelAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IconResponseModelAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetIsBuiltIn

`func (o *IconResponseModelAllOf) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *IconResponseModelAllOf) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *IconResponseModelAllOf) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.


### GetUid

`func (o *IconResponseModelAllOf) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IconResponseModelAllOf) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IconResponseModelAllOf) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IconResponseModelAllOf) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


