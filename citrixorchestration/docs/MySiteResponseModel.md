# MySiteResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the site. In the cloud this is the \&quot;Virtual site ID\&quot;.  On-prem this should be the XD farm guid. | 
**Name** | **string** | Name of the site. | 

## Methods

### NewMySiteResponseModel

`func NewMySiteResponseModel(id string, name string, ) *MySiteResponseModel`

NewMySiteResponseModel instantiates a new MySiteResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySiteResponseModelWithDefaults

`func NewMySiteResponseModelWithDefaults() *MySiteResponseModel`

NewMySiteResponseModelWithDefaults instantiates a new MySiteResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MySiteResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MySiteResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MySiteResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MySiteResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MySiteResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MySiteResponseModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


