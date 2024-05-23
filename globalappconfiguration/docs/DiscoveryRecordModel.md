# DiscoveryRecordModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**Apps**](Apps.md) |  | [optional] 
**Domain** | Pointer to [**Domain**](Domain.md) |  | [optional] 

## Methods

### NewDiscoveryRecordModel

`func NewDiscoveryRecordModel() *DiscoveryRecordModel`

NewDiscoveryRecordModel instantiates a new DiscoveryRecordModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryRecordModelWithDefaults

`func NewDiscoveryRecordModelWithDefaults() *DiscoveryRecordModel`

NewDiscoveryRecordModelWithDefaults instantiates a new DiscoveryRecordModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *DiscoveryRecordModel) GetApp() Apps`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *DiscoveryRecordModel) GetAppOk() (*Apps, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *DiscoveryRecordModel) SetApp(v Apps)`

SetApp sets App field to given value.

### HasApp

`func (o *DiscoveryRecordModel) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetDomain

`func (o *DiscoveryRecordModel) GetDomain() Domain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DiscoveryRecordModel) GetDomainOk() (*Domain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DiscoveryRecordModel) SetDomain(v Domain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DiscoveryRecordModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


