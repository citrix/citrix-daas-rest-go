# PublishedDesktop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**AppsAndDesktopsProvisionState**](AppsAndDesktopsProvisionState.md) | Current state in publishing the desktop | 
**Status** | Pointer to **string** | Status message related to desktop state | [optional] 
**Id** | **string** | Unique identifier of desktop, generally a GUID | 
**Name** | **string** | Display name of desktop | 
**Description** | Pointer to **string** | Description | [optional] 

## Methods

### NewPublishedDesktop

`func NewPublishedDesktop(state AppsAndDesktopsProvisionState, id string, name string, ) *PublishedDesktop`

NewPublishedDesktop instantiates a new PublishedDesktop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedDesktopWithDefaults

`func NewPublishedDesktopWithDefaults() *PublishedDesktop`

NewPublishedDesktopWithDefaults instantiates a new PublishedDesktop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PublishedDesktop) GetState() AppsAndDesktopsProvisionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PublishedDesktop) GetStateOk() (*AppsAndDesktopsProvisionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PublishedDesktop) SetState(v AppsAndDesktopsProvisionState)`

SetState sets State field to given value.


### GetStatus

`func (o *PublishedDesktop) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublishedDesktop) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublishedDesktop) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublishedDesktop) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *PublishedDesktop) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublishedDesktop) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublishedDesktop) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PublishedDesktop) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublishedDesktop) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublishedDesktop) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PublishedDesktop) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublishedDesktop) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublishedDesktop) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublishedDesktop) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


