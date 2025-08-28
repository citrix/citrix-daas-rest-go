# \BackupRestoreAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupRestoreCreateBackupSchedule**](BackupRestoreAPIsDAAS.md#BackupRestoreCreateBackupSchedule) | **Post** /BackupRestore/Schedules | Create backup schedule
[**BackupRestoreDeleteBackup**](BackupRestoreAPIsDAAS.md#BackupRestoreDeleteBackup) | **Delete** /BackupRestore/{backupName} | Delete a backup
[**BackupRestoreDeleteBackupSchedule**](BackupRestoreAPIsDAAS.md#BackupRestoreDeleteBackupSchedule) | **Delete** /BackupRestore/Schedule/{uid} | Delete backup schedule
[**BackupRestoreDisableSingleBackupSchedule**](BackupRestoreAPIsDAAS.md#BackupRestoreDisableSingleBackupSchedule) | **Post** /BackupRestore/Schedules/{uid}/Disable | Disable single backup schedule
[**BackupRestoreDoBackup**](BackupRestoreAPIsDAAS.md#BackupRestoreDoBackup) | **Post** /BackupRestore/DoBackup | Do backup
[**BackupRestoreDownloadSingleBackup**](BackupRestoreAPIsDAAS.md#BackupRestoreDownloadSingleBackup) | **Get** /BackupRestore/{backupName}/Download | Download single backup
[**BackupRestoreEnableSingleBackupSchedule**](BackupRestoreAPIsDAAS.md#BackupRestoreEnableSingleBackupSchedule) | **Post** /BackupRestore/Schedules/{uid}/Enable | Enable single backup schedule
[**BackupRestoreGetAllBackupHistory**](BackupRestoreAPIsDAAS.md#BackupRestoreGetAllBackupHistory) | **Get** /BackupRestore/History | Get all backup history
[**BackupRestoreGetBackupHistoryForSpecificHistory**](BackupRestoreAPIsDAAS.md#BackupRestoreGetBackupHistoryForSpecificHistory) | **Get** /BackupRestore/History/{uid} | Get backup history for a single specific backup
[**BackupRestoreGetBackupRestoreOptions**](BackupRestoreAPIsDAAS.md#BackupRestoreGetBackupRestoreOptions) | **Get** /BackupRestore/Options | Get backup / restore options
[**BackupRestoreGetBackupSchedule**](BackupRestoreAPIsDAAS.md#BackupRestoreGetBackupSchedule) | **Get** /BackupRestore/Schedules/{Uid} | Get single backup schedule
[**BackupRestoreGetBackupSchedules**](BackupRestoreAPIsDAAS.md#BackupRestoreGetBackupSchedules) | **Get** /BackupRestore/Schedules | Get backup schedules
[**BackupRestoreGetBackups**](BackupRestoreAPIsDAAS.md#BackupRestoreGetBackups) | **Get** /BackupRestore | Get backups
[**BackupRestoreGetBlobStorage**](BackupRestoreAPIsDAAS.md#BackupRestoreGetBlobStorage) | **Get** /BackupRestore/Storage | Get Blob Storage Information
[**BackupRestoreGetComponentMembers**](BackupRestoreAPIsDAAS.md#BackupRestoreGetComponentMembers) | **Post** /BackupRestore/{backupName}/Members | Get backup Members
[**BackupRestoreGetCurrentActionStatus**](BackupRestoreAPIsDAAS.md#BackupRestoreGetCurrentActionStatus) | **Get** /BackupRestore/Status | Get backup / restore status
[**BackupRestoreGetHistoryForSpecificBackup**](BackupRestoreAPIsDAAS.md#BackupRestoreGetHistoryForSpecificBackup) | **Get** /BackupRestore/{backupName}/History/All | Get all backup history for specific backup
[**BackupRestoreGetInformation**](BackupRestoreAPIsDAAS.md#BackupRestoreGetInformation) | **Get** /BackupRestore/Information | Get backup / restore information
[**BackupRestoreGetResultLog**](BackupRestoreAPIsDAAS.md#BackupRestoreGetResultLog) | **Post** /BackupRestore/{backupName}/ResultLog | Get result log
[**BackupRestoreModifyBackupSchedule**](BackupRestoreAPIsDAAS.md#BackupRestoreModifyBackupSchedule) | **Patch** /BackupRestore/Schedules/{uid} | Modify backup schedules
[**BackupRestorePinSingleBackup**](BackupRestoreAPIsDAAS.md#BackupRestorePinSingleBackup) | **Post** /BackupRestore/{backupName}/PinBackup | Pin a backup
[**BackupRestoreRestoreBackupToSite**](BackupRestoreAPIsDAAS.md#BackupRestoreRestoreBackupToSite) | **Post** /BackupRestore/{backupName}/Restore | Restore a backup
[**BackupRestoreSetBackupRestoreOptions**](BackupRestoreAPIsDAAS.md#BackupRestoreSetBackupRestoreOptions) | **Post** /BackupRestore/Options | Set backup / restore options
[**BackupRestoreSetBlobStorage**](BackupRestoreAPIsDAAS.md#BackupRestoreSetBlobStorage) | **Post** /BackupRestore/Storage | Set Blob Storage Information
[**BackupRestoreSetCurrentActionStatus**](BackupRestoreAPIsDAAS.md#BackupRestoreSetCurrentActionStatus) | **Post** /BackupRestore/Status | Set backup / restore status
[**BackupRestoreUnpinSingleBackup**](BackupRestoreAPIsDAAS.md#BackupRestoreUnpinSingleBackup) | **Post** /BackupRestore/{backupName}/UnpinBackup | Unpin a backup
[**BackupRestoreUpdateBckRstrHistoryNotes**](BackupRestoreAPIsDAAS.md#BackupRestoreUpdateBckRstrHistoryNotes) | **Post** /BackupRestore/History/{uid} | Update notes in history entry
[**BackupRestoreValidateStorage**](BackupRestoreAPIsDAAS.md#BackupRestoreValidateStorage) | **Get** /BackupRestore/Storage/Validate | Validate Storage



## BackupRestoreCreateBackupSchedule

> BackupRestoreScheduleResponseModel BackupRestoreCreateBackupSchedule(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreScheduleRequestModel(backupRestoreScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create backup schedule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupRestoreScheduleRequestModel := *openapiclient.NewBackupRestoreScheduleRequestModel("Name_example", "StartDate_example", int32(123), false, openapiclient.BackupRestoreScheduleFrequency("Daily"), "StartTime_example", "TimeZoneId_example") // BackupRestoreScheduleRequestModel | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreCreateBackupSchedule(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreScheduleRequestModel(backupRestoreScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreCreateBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreCreateBackupSchedule`: BackupRestoreScheduleResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreCreateBackupSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreCreateBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **backupRestoreScheduleRequestModel** | [**BackupRestoreScheduleRequestModel**](BackupRestoreScheduleRequestModel.md) |  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreScheduleResponseModel**](BackupRestoreScheduleResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreDeleteBackup

> BackupRestoreDeleteBackup(ctx, backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete a backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupName := "backupName_example" // string | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreDeleteBackup(context.Background(), backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreDeleteBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreDeleteBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreDeleteBackupSchedule

> BackupRestoreDeleteBackupSchedule(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete backup schedule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := int32(56) // int32 | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreDeleteBackupSchedule(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreDeleteBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreDeleteBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreDisableSingleBackupSchedule

> BackupRestoreDisableSingleBackupSchedule(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Disable single backup schedule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := int32(56) // int32 | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreDisableSingleBackupSchedule(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreDisableSingleBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreDisableSingleBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreDoBackup

> BackupDoBackupResponseModel BackupRestoreDoBackup(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreDoBackupRequestModel(backupRestoreDoBackupRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Do backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupRestoreDoBackupRequestModel := *openapiclient.NewBackupRestoreDoBackupRequestModel() // BackupRestoreDoBackupRequestModel | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreDoBackup(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreDoBackupRequestModel(backupRestoreDoBackupRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreDoBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreDoBackup`: BackupDoBackupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreDoBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreDoBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **backupRestoreDoBackupRequestModel** | [**BackupRestoreDoBackupRequestModel**](BackupRestoreDoBackupRequestModel.md) |  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupDoBackupResponseModel**](BackupDoBackupResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreDownloadSingleBackup

> BackupRestoreDownloadResponseModel BackupRestoreDownloadSingleBackup(ctx, backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).ContinuationToken(continuationToken).Execute()

Download single backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupName := "backupName_example" // string | Name of the backp to download
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	continuationToken := "continuationToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreDownloadSingleBackup(context.Background(), backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreDownloadSingleBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreDownloadSingleBackup`: BackupRestoreDownloadResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreDownloadSingleBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupName** | **string** | Name of the backp to download | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreDownloadSingleBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **continuationToken** | **string** |  | 

### Return type

[**BackupRestoreDownloadResponseModel**](BackupRestoreDownloadResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreEnableSingleBackupSchedule

> BackupRestoreEnableSingleBackupSchedule(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Enable single backup schedule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := int32(56) // int32 | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreEnableSingleBackupSchedule(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreEnableSingleBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreEnableSingleBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetAllBackupHistory

> BackupRestoreHistoryResponseModelCollection BackupRestoreGetAllBackupHistory(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get all backup history

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetAllBackupHistory(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetAllBackupHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetAllBackupHistory`: BackupRestoreHistoryResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetAllBackupHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetAllBackupHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]

### Return type

[**BackupRestoreHistoryResponseModelCollection**](BackupRestoreHistoryResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetBackupHistoryForSpecificHistory

> BackupRestoreHistoryInformation BackupRestoreGetBackupHistoryForSpecificHistory(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get backup history for a single specific backup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := "uid_example" // string | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetBackupHistoryForSpecificHistory(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetBackupHistoryForSpecificHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetBackupHistoryForSpecificHistory`: BackupRestoreHistoryInformation
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetBackupHistoryForSpecificHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetBackupHistoryForSpecificHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]

### Return type

[**BackupRestoreHistoryInformation**](BackupRestoreHistoryInformation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetBackupRestoreOptions

> BackupRestoreOptionsResponseModel BackupRestoreGetBackupRestoreOptions(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get backup / restore options



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetBackupRestoreOptions(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetBackupRestoreOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetBackupRestoreOptions`: BackupRestoreOptionsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetBackupRestoreOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetBackupRestoreOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreOptionsResponseModel**](BackupRestoreOptionsResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetBackupSchedule

> BackupRestoreScheduleModel BackupRestoreGetBackupSchedule(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get single backup schedule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := int32(56) // int32 | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetBackupSchedule(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetBackupSchedule`: BackupRestoreScheduleModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetBackupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreScheduleModel**](BackupRestoreScheduleModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetBackupSchedules

> BackupRestoreScheduleModelCollection BackupRestoreGetBackupSchedules(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get backup schedules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetBackupSchedules(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetBackupSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetBackupSchedules`: BackupRestoreScheduleModelCollection
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetBackupSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetBackupSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreScheduleModelCollection**](BackupRestoreScheduleModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetBackups

> BackupResponseModelCollection BackupRestoreGetBackups(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).AllFiles(allFiles).Execute()

Get backups

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
	allFiles := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetBackups(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).AllFiles(allFiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetBackups`: BackupResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **allFiles** | **bool** |  | [default to false]

### Return type

[**BackupResponseModelCollection**](BackupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetBlobStorage

> BackupRestoreStorageResponseModel BackupRestoreGetBlobStorage(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get Blob Storage Information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetBlobStorage(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetBlobStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetBlobStorage`: BackupRestoreStorageResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetBlobStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetBlobStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreStorageResponseModel**](BackupRestoreStorageResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetComponentMembers

> BackupMembersResponseModelCollection BackupRestoreGetComponentMembers(ctx, backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreGetComponentMemberRequestModel(backupRestoreGetComponentMemberRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get backup Members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupName := "backupName_example" // string | Name of the backp to restore
	backupRestoreGetComponentMemberRequestModel := *openapiclient.NewBackupRestoreGetComponentMemberRequestModel(openapiclient.BckRstrAutoConfigComponents("Tags")) // BackupRestoreGetComponentMemberRequestModel | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetComponentMembers(context.Background(), backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreGetComponentMemberRequestModel(backupRestoreGetComponentMemberRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetComponentMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetComponentMembers`: BackupMembersResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetComponentMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupName** | **string** | Name of the backp to restore | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetComponentMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **backupRestoreGetComponentMemberRequestModel** | [**BackupRestoreGetComponentMemberRequestModel**](BackupRestoreGetComponentMemberRequestModel.md) |  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** |  | [default to false]

### Return type

[**BackupMembersResponseModelCollection**](BackupMembersResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetCurrentActionStatus

> BackupRestoreStatusResponseModel BackupRestoreGetCurrentActionStatus(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get backup / restore status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetCurrentActionStatus(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetCurrentActionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetCurrentActionStatus`: BackupRestoreStatusResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetCurrentActionStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetCurrentActionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreStatusResponseModel**](BackupRestoreStatusResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetHistoryForSpecificBackup

> BackupRestoreHistoryResponseModelCollection BackupRestoreGetHistoryForSpecificBackup(ctx, backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get all backup history for specific backup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupName := "backupName_example" // string | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetHistoryForSpecificBackup(context.Background(), backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetHistoryForSpecificBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetHistoryForSpecificBackup`: BackupRestoreHistoryResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetHistoryForSpecificBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetHistoryForSpecificBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]

### Return type

[**BackupRestoreHistoryResponseModelCollection**](BackupRestoreHistoryResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetInformation

> BackupRestoreInformationResponseModel BackupRestoreGetInformation(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get backup / restore information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetInformation(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetInformation`: BackupRestoreInformationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreInformationResponseModel**](BackupRestoreInformationResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreGetResultLog

> BackupRestoreResultLogResponseModel BackupRestoreGetResultLog(ctx, backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get result log



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupName := "backupName_example" // string | Name of the backp to restore
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreGetResultLog(context.Background(), backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreGetResultLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreGetResultLog`: BackupRestoreResultLogResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreGetResultLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupName** | **string** | Name of the backp to restore | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreGetResultLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** |  | [default to false]

### Return type

[**BackupRestoreResultLogResponseModel**](BackupRestoreResultLogResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreModifyBackupSchedule

> BackupRestoreScheduleResponseModel BackupRestoreModifyBackupSchedule(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreScheduleRequestModel(backupRestoreScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Modify backup schedules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := int32(56) // int32 | 
	backupRestoreScheduleRequestModel := *openapiclient.NewBackupRestoreScheduleRequestModel("Name_example", "StartDate_example", int32(123), false, openapiclient.BackupRestoreScheduleFrequency("Daily"), "StartTime_example", "TimeZoneId_example") // BackupRestoreScheduleRequestModel | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreModifyBackupSchedule(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreScheduleRequestModel(backupRestoreScheduleRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreModifyBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreModifyBackupSchedule`: BackupRestoreScheduleResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreModifyBackupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreModifyBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **backupRestoreScheduleRequestModel** | [**BackupRestoreScheduleRequestModel**](BackupRestoreScheduleRequestModel.md) |  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreScheduleResponseModel**](BackupRestoreScheduleResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestorePinSingleBackup

> bool BackupRestorePinSingleBackup(ctx, backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Pin a backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupName := "backupName_example" // string | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestorePinSingleBackup(context.Background(), backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestorePinSingleBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestorePinSingleBackup`: bool
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestorePinSingleBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestorePinSingleBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** |  | [default to false]

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreRestoreBackupToSite

> BackupRestoreRestoreBackupResponseModel BackupRestoreRestoreBackupToSite(ctx, backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreRestoreBackupRequestModel(backupRestoreRestoreBackupRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Restore a backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupName := "backupName_example" // string | Name of the backp to restore
	backupRestoreRestoreBackupRequestModel := *openapiclient.NewBackupRestoreRestoreBackupRequestModel("BackupName_example", int32(123), openapiclient.BckRstrAutoConfigComponents("Tags")) // BackupRestoreRestoreBackupRequestModel | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreRestoreBackupToSite(context.Background(), backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreRestoreBackupRequestModel(backupRestoreRestoreBackupRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreRestoreBackupToSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreRestoreBackupToSite`: BackupRestoreRestoreBackupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreRestoreBackupToSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupName** | **string** | Name of the backp to restore | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreRestoreBackupToSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **backupRestoreRestoreBackupRequestModel** | [**BackupRestoreRestoreBackupRequestModel**](BackupRestoreRestoreBackupRequestModel.md) |  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreRestoreBackupResponseModel**](BackupRestoreRestoreBackupResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreSetBackupRestoreOptions

> BackupRestoreOptionsResponseModel BackupRestoreSetBackupRestoreOptions(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreOptionsRequestModel(backupRestoreOptionsRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Set backup / restore options



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupRestoreOptionsRequestModel := *openapiclient.NewBackupRestoreOptionsRequestModel() // BackupRestoreOptionsRequestModel | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreSetBackupRestoreOptions(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreOptionsRequestModel(backupRestoreOptionsRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreSetBackupRestoreOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreSetBackupRestoreOptions`: BackupRestoreOptionsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreSetBackupRestoreOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreSetBackupRestoreOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **backupRestoreOptionsRequestModel** | [**BackupRestoreOptionsRequestModel**](BackupRestoreOptionsRequestModel.md) |  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreOptionsResponseModel**](BackupRestoreOptionsResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreSetBlobStorage

> BackupRestoreValidateStorageResponseModel BackupRestoreSetBlobStorage(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreStorageRequestModel(backupRestoreStorageRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Set Blob Storage Information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupRestoreStorageRequestModel := *openapiclient.NewBackupRestoreStorageRequestModel() // BackupRestoreStorageRequestModel | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreSetBlobStorage(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreStorageRequestModel(backupRestoreStorageRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreSetBlobStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreSetBlobStorage`: BackupRestoreValidateStorageResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreSetBlobStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreSetBlobStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **backupRestoreStorageRequestModel** | [**BackupRestoreStorageRequestModel**](BackupRestoreStorageRequestModel.md) |  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreValidateStorageResponseModel**](BackupRestoreValidateStorageResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreSetCurrentActionStatus

> BackupRestoreStatusResponseModel BackupRestoreSetCurrentActionStatus(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreStatusRequestModel(backupRestoreStatusRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Set backup / restore status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupRestoreStatusRequestModel := *openapiclient.NewBackupRestoreStatusRequestModel(openapiclient.BackupRestoreActions("Backup"), "ExecutionId_example", int32(123), openapiclient.BackupRestoreStatus("Unknown"), time.Now(), time.Now()) // BackupRestoreStatusRequestModel | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreSetCurrentActionStatus(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreStatusRequestModel(backupRestoreStatusRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreSetCurrentActionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreSetCurrentActionStatus`: BackupRestoreStatusResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreSetCurrentActionStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreSetCurrentActionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **backupRestoreStatusRequestModel** | [**BackupRestoreStatusRequestModel**](BackupRestoreStatusRequestModel.md) |  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreStatusResponseModel**](BackupRestoreStatusResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreUnpinSingleBackup

> bool BackupRestoreUnpinSingleBackup(ctx, backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Unpin a backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	backupName := "backupName_example" // string | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreUnpinSingleBackup(context.Background(), backupName).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreUnpinSingleBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreUnpinSingleBackup`: bool
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreUnpinSingleBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreUnpinSingleBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreUpdateBckRstrHistoryNotes

> bool BackupRestoreUpdateBckRstrHistoryNotes(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreUpdateNotesRequestModel(backupRestoreUpdateNotesRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update notes in history entry

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := int32(56) // int32 | History entry unique identifier
	backupRestoreUpdateNotesRequestModel := *openapiclient.NewBackupRestoreUpdateNotesRequestModel() // BackupRestoreUpdateNotesRequestModel | new notes
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreUpdateBckRstrHistoryNotes(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BackupRestoreUpdateNotesRequestModel(backupRestoreUpdateNotesRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreUpdateBckRstrHistoryNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreUpdateBckRstrHistoryNotes`: bool
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreUpdateBckRstrHistoryNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** | History entry unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreUpdateBckRstrHistoryNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **backupRestoreUpdateNotesRequestModel** | [**BackupRestoreUpdateNotesRequestModel**](BackupRestoreUpdateNotesRequestModel.md) | new notes | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupRestoreValidateStorage

> BackupRestoreValidateStorageResponseModel BackupRestoreValidateStorage(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Validate Storage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupRestoreAPIsDAAS.BackupRestoreValidateStorage(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreAPIsDAAS.BackupRestoreValidateStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackupRestoreValidateStorage`: BackupRestoreValidateStorageResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupRestoreAPIsDAAS.BackupRestoreValidateStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupRestoreValidateStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**BackupRestoreValidateStorageResponseModel**](BackupRestoreValidateStorageResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

