# AzureAlertModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**Data**](Data.md) |  | [optional] 

## Methods

### NewAzureAlertModel

`func NewAzureAlertModel() *AzureAlertModel`

NewAzureAlertModel instantiates a new AzureAlertModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAlertModelWithDefaults

`func NewAzureAlertModelWithDefaults() *AzureAlertModel`

NewAzureAlertModelWithDefaults instantiates a new AzureAlertModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaId

`func (o *AzureAlertModel) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *AzureAlertModel) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *AzureAlertModel) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *AzureAlertModel) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetData

`func (o *AzureAlertModel) GetData() Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AzureAlertModel) GetDataOk() (*Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AzureAlertModel) SetData(v Data)`

SetData sets Data field to given value.

### HasData

`func (o *AzureAlertModel) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


