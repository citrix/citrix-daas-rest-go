# AppVApplicationDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceLocation** | **string** | App sequence location. | 
**TargetInPackage** | **bool** | Whether the app target is in a package. | 
**Users** | [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | Users associated with the App-V application on the management server. | 
**WorkingDirectory** | **string** | Working directory for the App-V application as defined on the management server. | 

## Methods

### NewAppVApplicationDetailResponseModelAllOf

`func NewAppVApplicationDetailResponseModelAllOf(sequenceLocation string, targetInPackage bool, users []IdentityUserResponseModel, workingDirectory string, ) *AppVApplicationDetailResponseModelAllOf`

NewAppVApplicationDetailResponseModelAllOf instantiates a new AppVApplicationDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVApplicationDetailResponseModelAllOfWithDefaults

`func NewAppVApplicationDetailResponseModelAllOfWithDefaults() *AppVApplicationDetailResponseModelAllOf`

NewAppVApplicationDetailResponseModelAllOfWithDefaults instantiates a new AppVApplicationDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequenceLocation

`func (o *AppVApplicationDetailResponseModelAllOf) GetSequenceLocation() string`

GetSequenceLocation returns the SequenceLocation field if non-nil, zero value otherwise.

### GetSequenceLocationOk

`func (o *AppVApplicationDetailResponseModelAllOf) GetSequenceLocationOk() (*string, bool)`

GetSequenceLocationOk returns a tuple with the SequenceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceLocation

`func (o *AppVApplicationDetailResponseModelAllOf) SetSequenceLocation(v string)`

SetSequenceLocation sets SequenceLocation field to given value.


### GetTargetInPackage

`func (o *AppVApplicationDetailResponseModelAllOf) GetTargetInPackage() bool`

GetTargetInPackage returns the TargetInPackage field if non-nil, zero value otherwise.

### GetTargetInPackageOk

`func (o *AppVApplicationDetailResponseModelAllOf) GetTargetInPackageOk() (*bool, bool)`

GetTargetInPackageOk returns a tuple with the TargetInPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetInPackage

`func (o *AppVApplicationDetailResponseModelAllOf) SetTargetInPackage(v bool)`

SetTargetInPackage sets TargetInPackage field to given value.


### GetUsers

`func (o *AppVApplicationDetailResponseModelAllOf) GetUsers() []IdentityUserResponseModel`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AppVApplicationDetailResponseModelAllOf) GetUsersOk() (*[]IdentityUserResponseModel, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AppVApplicationDetailResponseModelAllOf) SetUsers(v []IdentityUserResponseModel)`

SetUsers sets Users field to given value.


### GetWorkingDirectory

`func (o *AppVApplicationDetailResponseModelAllOf) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *AppVApplicationDetailResponseModelAllOf) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *AppVApplicationDetailResponseModelAllOf) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


