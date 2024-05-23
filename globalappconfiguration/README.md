# Go API client for globalappconfiguration

Describes API used by Global App Config Admin Service

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Generator version: 7.4.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import globalappconfiguration "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `globalappconfiguration.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), globalappconfiguration.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `globalappconfiguration.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), globalappconfiguration.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `globalappconfiguration.ContextOperationServerIndices` and `globalappconfiguration.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), globalappconfiguration.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), globalappconfiguration.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://wsaca.cloud.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DiscoveryControllerDAAS* | [**DeleteDiscoveryApiUsingDELETE**](docs/DiscoveryControllerDAAS.md#deletediscoveryapiusingdelete) | **Delete** /aca/discovery/app/{app}/domain/{domain} | Delete Discovery Record For App And Email Domain Suffix.
*DiscoveryControllerDAAS* | [**GetAllDiscoveryApiUsingGET**](docs/DiscoveryControllerDAAS.md#getalldiscoveryapiusingget) | **Get** /aca/discovery | Retrieve All Discovery Records For The Specified Customer Id
*DiscoveryControllerDAAS* | [**GetDiscoveryApiUsingGET**](docs/DiscoveryControllerDAAS.md#getdiscoveryapiusingget) | **Get** /aca/discovery/app/{app}/domain/{domain} | Retrieve Discovery Record For App And Email Domain Suffix.
*DiscoveryControllerDAAS* | [**PostDiscoveryApiUsingPOST**](docs/DiscoveryControllerDAAS.md#postdiscoveryapiusingpost) | **Post** /aca/discovery/app/{app}/domain | Create A New Discovery Record.
*DiscoveryControllerDAAS* | [**PutDiscoveryApiUsingPUT**](docs/DiscoveryControllerDAAS.md#putdiscoveryapiusingput) | **Put** /aca/discovery/app/{app}/domain/{domain} | Update Discovery Record For An App And Email Domain Suffix.
*SettingsControllerDAAS* | [**DeleteSettingsApiUsingDELETE**](docs/SettingsControllerDAAS.md#deletesettingsapiusingdelete) | **Delete** /aca/settings/app/{app}/url/{url} | Delete Record For App And Service-url.
*SettingsControllerDAAS* | [**DeleteSettingsForChannel**](docs/SettingsControllerDAAS.md#deletesettingsforchannel) | **Delete** /aca/settings/app/{app}/url/{url}/channel/{channelName} | Delete Settings Record For The Specified Customer Id, Product, Service And Channel.
*SettingsControllerDAAS* | [**GetAllSettingsApiUsingGET**](docs/SettingsControllerDAAS.md#getallsettingsapiusingget) | **Get** /aca/settings | Retrieve All Settings Records For The Specified Customer Id.
*SettingsControllerDAAS* | [**GetSettingsApiUsingGET**](docs/SettingsControllerDAAS.md#getsettingsapiusingget) | **Get** /aca/settings/app/{app}/url/{url} | Retrieve Settings Record For App And Service-URL.
*SettingsControllerDAAS* | [**PostSettingsApiUsingPOST**](docs/SettingsControllerDAAS.md#postsettingsapiusingpost) | **Post** /aca/settings/app/{app}/url | Create A New Settings Record.
*SettingsControllerDAAS* | [**PutSettingsApiUsingPUT**](docs/SettingsControllerDAAS.md#putsettingsapiusingput) | **Put** /aca/settings/app/{app}/url/{url} | Update Settings Record For An App And Service-URL.
*SettingsControllerDAAS* | [**RetrieveAllChannelSettings**](docs/SettingsControllerDAAS.md#retrieveallchannelsettings) | **Get** /aca/settings/app/{app}/url/{url}/channel | Retrieve All Channels Settings Records For The Specified Customer Id and url
*SettingsControllerDAAS* | [**RetrieveSettingsForChannel**](docs/SettingsControllerDAAS.md#retrievesettingsforchannel) | **Get** /aca/settings/app/{app}/url/{url}/channel/{channelName} | Retrieve Settings Record For The Specified Customer Id, Product, Service And Channel.


## Documentation For Models

 - [AllowedWebStoreURL](docs/AllowedWebStoreURL.md)
 - [AppSettings](docs/AppSettings.md)
 - [Apps](docs/Apps.md)
 - [CategorySettings](docs/CategorySettings.md)
 - [CitrixErrorModel](docs/CitrixErrorModel.md)
 - [DiscoveryRecordModel](docs/DiscoveryRecordModel.md)
 - [DiscoveryServiceURL](docs/DiscoveryServiceURL.md)
 - [Domain](docs/Domain.md)
 - [GetAllDiscoveryResponse](docs/GetAllDiscoveryResponse.md)
 - [GetAllSettingResponse](docs/GetAllSettingResponse.md)
 - [MigrationUrl](docs/MigrationUrl.md)
 - [Parameter](docs/Parameter.md)
 - [PlatformSettings](docs/PlatformSettings.md)
 - [ServiceURL](docs/ServiceURL.md)
 - [Settings](docs/Settings.md)
 - [SettingsChannel](docs/SettingsChannel.md)
 - [SettingsRecordModel](docs/SettingsRecordModel.md)
 - [Workspace](docs/Workspace.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerTokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		globalappconfiguration.ContextAPIKeys,
		map[string]globalappconfiguration.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


