# SitesErrorWarningModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfErrors** | Pointer to **int32** |  | [optional] 
**NumberOfWarnings** | Pointer to **int32** |  | [optional] 
**ErrorWarning** | Pointer to [**[]SiteErrorWarningModel**](SiteErrorWarningModel.md) |  | [optional] 

## Methods

### NewSitesErrorWarningModel

`func NewSitesErrorWarningModel() *SitesErrorWarningModel`

NewSitesErrorWarningModel instantiates a new SitesErrorWarningModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSitesErrorWarningModelWithDefaults

`func NewSitesErrorWarningModelWithDefaults() *SitesErrorWarningModel`

NewSitesErrorWarningModelWithDefaults instantiates a new SitesErrorWarningModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfErrors

`func (o *SitesErrorWarningModel) GetNumberOfErrors() int32`

GetNumberOfErrors returns the NumberOfErrors field if non-nil, zero value otherwise.

### GetNumberOfErrorsOk

`func (o *SitesErrorWarningModel) GetNumberOfErrorsOk() (*int32, bool)`

GetNumberOfErrorsOk returns a tuple with the NumberOfErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfErrors

`func (o *SitesErrorWarningModel) SetNumberOfErrors(v int32)`

SetNumberOfErrors sets NumberOfErrors field to given value.

### HasNumberOfErrors

`func (o *SitesErrorWarningModel) HasNumberOfErrors() bool`

HasNumberOfErrors returns a boolean if a field has been set.

### GetNumberOfWarnings

`func (o *SitesErrorWarningModel) GetNumberOfWarnings() int32`

GetNumberOfWarnings returns the NumberOfWarnings field if non-nil, zero value otherwise.

### GetNumberOfWarningsOk

`func (o *SitesErrorWarningModel) GetNumberOfWarningsOk() (*int32, bool)`

GetNumberOfWarningsOk returns a tuple with the NumberOfWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfWarnings

`func (o *SitesErrorWarningModel) SetNumberOfWarnings(v int32)`

SetNumberOfWarnings sets NumberOfWarnings field to given value.

### HasNumberOfWarnings

`func (o *SitesErrorWarningModel) HasNumberOfWarnings() bool`

HasNumberOfWarnings returns a boolean if a field has been set.

### GetErrorWarning

`func (o *SitesErrorWarningModel) GetErrorWarning() []SiteErrorWarningModel`

GetErrorWarning returns the ErrorWarning field if non-nil, zero value otherwise.

### GetErrorWarningOk

`func (o *SitesErrorWarningModel) GetErrorWarningOk() (*[]SiteErrorWarningModel, bool)`

GetErrorWarningOk returns a tuple with the ErrorWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorWarning

`func (o *SitesErrorWarningModel) SetErrorWarning(v []SiteErrorWarningModel)`

SetErrorWarning sets ErrorWarning field to given value.

### HasErrorWarning

`func (o *SitesErrorWarningModel) HasErrorWarning() bool`

HasErrorWarning returns a boolean if a field has been set.

### SetErrorWarningNil

`func (o *SitesErrorWarningModel) SetErrorWarningNil(b bool)`

 SetErrorWarningNil sets the value for ErrorWarning to be an explicit nil

### UnsetErrorWarning
`func (o *SitesErrorWarningModel) UnsetErrorWarning()`

UnsetErrorWarning ensures that no value is present for ErrorWarning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


