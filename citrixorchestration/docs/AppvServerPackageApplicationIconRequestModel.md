# AppvServerPackageApplicationIconRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | **string** | ManagementServeraddress of the App-V server | 
**PackageId** | **string** | Id of the App-V package | 
**ApplicationId** | **string** | Id of the App-V application within the package. | 
**IconFormat** | Pointer to **NullableString** | Icon format.  Must be: &#x60;{mime-type};{width}x{height}x{colordepth}&#x60;  where:   * _mime-type_ must be &#x60;image/png&#x60;.  (Other formats may be supported in future.) * _width_ and _height_ are specified in pixels. * _colordepth_ (optional) is either &#x60;8&#x60; or &#x60;24&#x60;.  Optional. If not specified, only the raw icon data will be returned. Note that this is typically in ICO format, which some clients cannot display properly. | [optional] 

## Methods

### NewAppvServerPackageApplicationIconRequestModel

`func NewAppvServerPackageApplicationIconRequestModel(server string, packageId string, applicationId string, ) *AppvServerPackageApplicationIconRequestModel`

NewAppvServerPackageApplicationIconRequestModel instantiates a new AppvServerPackageApplicationIconRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppvServerPackageApplicationIconRequestModelWithDefaults

`func NewAppvServerPackageApplicationIconRequestModelWithDefaults() *AppvServerPackageApplicationIconRequestModel`

NewAppvServerPackageApplicationIconRequestModelWithDefaults instantiates a new AppvServerPackageApplicationIconRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *AppvServerPackageApplicationIconRequestModel) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AppvServerPackageApplicationIconRequestModel) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AppvServerPackageApplicationIconRequestModel) SetServer(v string)`

SetServer sets Server field to given value.


### GetPackageId

`func (o *AppvServerPackageApplicationIconRequestModel) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *AppvServerPackageApplicationIconRequestModel) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *AppvServerPackageApplicationIconRequestModel) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetApplicationId

`func (o *AppvServerPackageApplicationIconRequestModel) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AppvServerPackageApplicationIconRequestModel) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AppvServerPackageApplicationIconRequestModel) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetIconFormat

`func (o *AppvServerPackageApplicationIconRequestModel) GetIconFormat() string`

GetIconFormat returns the IconFormat field if non-nil, zero value otherwise.

### GetIconFormatOk

`func (o *AppvServerPackageApplicationIconRequestModel) GetIconFormatOk() (*string, bool)`

GetIconFormatOk returns a tuple with the IconFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFormat

`func (o *AppvServerPackageApplicationIconRequestModel) SetIconFormat(v string)`

SetIconFormat sets IconFormat field to given value.

### HasIconFormat

`func (o *AppvServerPackageApplicationIconRequestModel) HasIconFormat() bool`

HasIconFormat returns a boolean if a field has been set.

### SetIconFormatNil

`func (o *AppvServerPackageApplicationIconRequestModel) SetIconFormatNil(b bool)`

 SetIconFormatNil sets the value for IconFormat to be an explicit nil

### UnsetIconFormat
`func (o *AppvServerPackageApplicationIconRequestModel) UnsetIconFormat()`

UnsetIconFormat ensures that no value is present for IconFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


