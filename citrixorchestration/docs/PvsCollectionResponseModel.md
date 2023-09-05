# PvsCollectionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the PVS collection. | 
**Name** | **string** | Name of the PVS collection. | 
**Domain** | Pointer to **string** | Active directory domain of machines in the collection. internally: when querying a collection, get a single device from the collection and get its AD domain. | [optional] 
**PvsSite** | [**PvsCollectionResponseModelPvsSite**](PvsCollectionResponseModelPvsSite.md) |  | 

## Methods

### NewPvsCollectionResponseModel

`func NewPvsCollectionResponseModel(id string, name string, pvsSite PvsCollectionResponseModelPvsSite, ) *PvsCollectionResponseModel`

NewPvsCollectionResponseModel instantiates a new PvsCollectionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvsCollectionResponseModelWithDefaults

`func NewPvsCollectionResponseModelWithDefaults() *PvsCollectionResponseModel`

NewPvsCollectionResponseModelWithDefaults instantiates a new PvsCollectionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PvsCollectionResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PvsCollectionResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PvsCollectionResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PvsCollectionResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PvsCollectionResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PvsCollectionResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDomain

`func (o *PvsCollectionResponseModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PvsCollectionResponseModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PvsCollectionResponseModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PvsCollectionResponseModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPvsSite

`func (o *PvsCollectionResponseModel) GetPvsSite() PvsCollectionResponseModelPvsSite`

GetPvsSite returns the PvsSite field if non-nil, zero value otherwise.

### GetPvsSiteOk

`func (o *PvsCollectionResponseModel) GetPvsSiteOk() (*PvsCollectionResponseModelPvsSite, bool)`

GetPvsSiteOk returns a tuple with the PvsSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsSite

`func (o *PvsCollectionResponseModel) SetPvsSite(v PvsCollectionResponseModelPvsSite)`

SetPvsSite sets PvsSite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


