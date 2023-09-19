# MyCustomerResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the customer. internally: the CC customer id / short-name, or for on-prem this could be any generated guid. | 
**Name** | Pointer to **NullableString** | Human-readable name of the customer.  Not necessarily unique. internally: the CC customer full name, or for on-prem this can be something hardcoded. | [optional] 
**Sites** | [**[]MySiteResponseModel**](MySiteResponseModel.md) | Sites which the admin has access to within the customer. | 

## Methods

### NewMyCustomerResponseModel

`func NewMyCustomerResponseModel(id string, sites []MySiteResponseModel, ) *MyCustomerResponseModel`

NewMyCustomerResponseModel instantiates a new MyCustomerResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyCustomerResponseModelWithDefaults

`func NewMyCustomerResponseModelWithDefaults() *MyCustomerResponseModel`

NewMyCustomerResponseModelWithDefaults instantiates a new MyCustomerResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MyCustomerResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyCustomerResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyCustomerResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MyCustomerResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MyCustomerResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MyCustomerResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MyCustomerResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MyCustomerResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MyCustomerResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSites

`func (o *MyCustomerResponseModel) GetSites() []MySiteResponseModel`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *MyCustomerResponseModel) GetSitesOk() (*[]MySiteResponseModel, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *MyCustomerResponseModel) SetSites(v []MySiteResponseModel)`

SetSites sets Sites field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


