# MachineDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedClientName** | Pointer to **string** | The name of the endpoint client device that the machine has been assigned to. | [optional] 
**AssignedIPAddress** | Pointer to **string** | The IP address of the endpoint client device that the machine has been assigned to. | [optional] 
**BrowserName** | Pointer to **string** | Site-wide unique name identifying associated desktop to other components (for example StoreFront). This is typically non-null only for machines backing assigned private desktops. | [optional] 
**ColorDepth** | Pointer to [**ColorDepth**](ColorDepth.md) |  | [optional] 
**IconId** | Pointer to **string** | The machine&#39;s icon that is displayed in Receiver. | [optional] 
**IsReserved** | **bool** | Indicates if machine is reserved for special use, for example for AppDisk preparation. A reserved machine cannot be a member of a delivery group. | 
**LoadIndexes** | Pointer to **[]int32** | Gives the last reported individual load indexes that were used in the calculation of the LoadIndex value. Note that the LoadIndex value may have been subsequently adjusted due to session brokering operations. This value is only set when SessionSupport is equal to MultiSession. | [optional] 
**SecureIcaRequired** | Pointer to **bool** | Flag indicating whether SecureICA is required or not when starting a session on the machine. | [optional] 
**SessionsEstablished** | **int32** | Number of established sessions on this machine.  When SessionSupport is equal to MultiSession, this excludes established sessions which have not yet completed their logon processing. | 
**SessionsPending** | **int32** | Number of pending (brokered but not yet established) sessions on this machine.  When SessionSupport is equal to MultiSession, this also includes established sessions which have not yet completed their logon processing. | 
**StoreFrontServersForHostedReceiver** | Pointer to [**[]StoreFrontServerResponseModel**](StoreFrontServerResponseModel.md) | StoreFront servers to use for Receiver when it is hosted on the machine. | [optional] 
**VMToolsState** | Pointer to [**VMToolsState**](VMToolsState.md) |  | [optional] 
**FailureReason** | Pointer to **string** | Failure reason of power action. | [optional] 

## Methods

### NewMachineDetailResponseModelAllOf

`func NewMachineDetailResponseModelAllOf(isReserved bool, sessionsEstablished int32, sessionsPending int32, ) *MachineDetailResponseModelAllOf`

NewMachineDetailResponseModelAllOf instantiates a new MachineDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineDetailResponseModelAllOfWithDefaults

`func NewMachineDetailResponseModelAllOfWithDefaults() *MachineDetailResponseModelAllOf`

NewMachineDetailResponseModelAllOfWithDefaults instantiates a new MachineDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedClientName

`func (o *MachineDetailResponseModelAllOf) GetAssignedClientName() string`

GetAssignedClientName returns the AssignedClientName field if non-nil, zero value otherwise.

### GetAssignedClientNameOk

`func (o *MachineDetailResponseModelAllOf) GetAssignedClientNameOk() (*string, bool)`

GetAssignedClientNameOk returns a tuple with the AssignedClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedClientName

`func (o *MachineDetailResponseModelAllOf) SetAssignedClientName(v string)`

SetAssignedClientName sets AssignedClientName field to given value.

### HasAssignedClientName

`func (o *MachineDetailResponseModelAllOf) HasAssignedClientName() bool`

HasAssignedClientName returns a boolean if a field has been set.

### GetAssignedIPAddress

`func (o *MachineDetailResponseModelAllOf) GetAssignedIPAddress() string`

GetAssignedIPAddress returns the AssignedIPAddress field if non-nil, zero value otherwise.

### GetAssignedIPAddressOk

`func (o *MachineDetailResponseModelAllOf) GetAssignedIPAddressOk() (*string, bool)`

GetAssignedIPAddressOk returns a tuple with the AssignedIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedIPAddress

`func (o *MachineDetailResponseModelAllOf) SetAssignedIPAddress(v string)`

SetAssignedIPAddress sets AssignedIPAddress field to given value.

### HasAssignedIPAddress

`func (o *MachineDetailResponseModelAllOf) HasAssignedIPAddress() bool`

HasAssignedIPAddress returns a boolean if a field has been set.

### GetBrowserName

`func (o *MachineDetailResponseModelAllOf) GetBrowserName() string`

GetBrowserName returns the BrowserName field if non-nil, zero value otherwise.

### GetBrowserNameOk

`func (o *MachineDetailResponseModelAllOf) GetBrowserNameOk() (*string, bool)`

GetBrowserNameOk returns a tuple with the BrowserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserName

`func (o *MachineDetailResponseModelAllOf) SetBrowserName(v string)`

SetBrowserName sets BrowserName field to given value.

### HasBrowserName

`func (o *MachineDetailResponseModelAllOf) HasBrowserName() bool`

HasBrowserName returns a boolean if a field has been set.

### GetColorDepth

`func (o *MachineDetailResponseModelAllOf) GetColorDepth() ColorDepth`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *MachineDetailResponseModelAllOf) GetColorDepthOk() (*ColorDepth, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *MachineDetailResponseModelAllOf) SetColorDepth(v ColorDepth)`

SetColorDepth sets ColorDepth field to given value.

### HasColorDepth

`func (o *MachineDetailResponseModelAllOf) HasColorDepth() bool`

HasColorDepth returns a boolean if a field has been set.

### GetIconId

`func (o *MachineDetailResponseModelAllOf) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *MachineDetailResponseModelAllOf) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *MachineDetailResponseModelAllOf) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *MachineDetailResponseModelAllOf) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### GetIsReserved

`func (o *MachineDetailResponseModelAllOf) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *MachineDetailResponseModelAllOf) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *MachineDetailResponseModelAllOf) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.


### GetLoadIndexes

`func (o *MachineDetailResponseModelAllOf) GetLoadIndexes() []int32`

GetLoadIndexes returns the LoadIndexes field if non-nil, zero value otherwise.

### GetLoadIndexesOk

`func (o *MachineDetailResponseModelAllOf) GetLoadIndexesOk() (*[]int32, bool)`

GetLoadIndexesOk returns a tuple with the LoadIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadIndexes

`func (o *MachineDetailResponseModelAllOf) SetLoadIndexes(v []int32)`

SetLoadIndexes sets LoadIndexes field to given value.

### HasLoadIndexes

`func (o *MachineDetailResponseModelAllOf) HasLoadIndexes() bool`

HasLoadIndexes returns a boolean if a field has been set.

### GetSecureIcaRequired

`func (o *MachineDetailResponseModelAllOf) GetSecureIcaRequired() bool`

GetSecureIcaRequired returns the SecureIcaRequired field if non-nil, zero value otherwise.

### GetSecureIcaRequiredOk

`func (o *MachineDetailResponseModelAllOf) GetSecureIcaRequiredOk() (*bool, bool)`

GetSecureIcaRequiredOk returns a tuple with the SecureIcaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureIcaRequired

`func (o *MachineDetailResponseModelAllOf) SetSecureIcaRequired(v bool)`

SetSecureIcaRequired sets SecureIcaRequired field to given value.

### HasSecureIcaRequired

`func (o *MachineDetailResponseModelAllOf) HasSecureIcaRequired() bool`

HasSecureIcaRequired returns a boolean if a field has been set.

### GetSessionsEstablished

`func (o *MachineDetailResponseModelAllOf) GetSessionsEstablished() int32`

GetSessionsEstablished returns the SessionsEstablished field if non-nil, zero value otherwise.

### GetSessionsEstablishedOk

`func (o *MachineDetailResponseModelAllOf) GetSessionsEstablishedOk() (*int32, bool)`

GetSessionsEstablishedOk returns a tuple with the SessionsEstablished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsEstablished

`func (o *MachineDetailResponseModelAllOf) SetSessionsEstablished(v int32)`

SetSessionsEstablished sets SessionsEstablished field to given value.


### GetSessionsPending

`func (o *MachineDetailResponseModelAllOf) GetSessionsPending() int32`

GetSessionsPending returns the SessionsPending field if non-nil, zero value otherwise.

### GetSessionsPendingOk

`func (o *MachineDetailResponseModelAllOf) GetSessionsPendingOk() (*int32, bool)`

GetSessionsPendingOk returns a tuple with the SessionsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsPending

`func (o *MachineDetailResponseModelAllOf) SetSessionsPending(v int32)`

SetSessionsPending sets SessionsPending field to given value.


### GetStoreFrontServersForHostedReceiver

`func (o *MachineDetailResponseModelAllOf) GetStoreFrontServersForHostedReceiver() []StoreFrontServerResponseModel`

GetStoreFrontServersForHostedReceiver returns the StoreFrontServersForHostedReceiver field if non-nil, zero value otherwise.

### GetStoreFrontServersForHostedReceiverOk

`func (o *MachineDetailResponseModelAllOf) GetStoreFrontServersForHostedReceiverOk() (*[]StoreFrontServerResponseModel, bool)`

GetStoreFrontServersForHostedReceiverOk returns a tuple with the StoreFrontServersForHostedReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFrontServersForHostedReceiver

`func (o *MachineDetailResponseModelAllOf) SetStoreFrontServersForHostedReceiver(v []StoreFrontServerResponseModel)`

SetStoreFrontServersForHostedReceiver sets StoreFrontServersForHostedReceiver field to given value.

### HasStoreFrontServersForHostedReceiver

`func (o *MachineDetailResponseModelAllOf) HasStoreFrontServersForHostedReceiver() bool`

HasStoreFrontServersForHostedReceiver returns a boolean if a field has been set.

### GetVMToolsState

`func (o *MachineDetailResponseModelAllOf) GetVMToolsState() VMToolsState`

GetVMToolsState returns the VMToolsState field if non-nil, zero value otherwise.

### GetVMToolsStateOk

`func (o *MachineDetailResponseModelAllOf) GetVMToolsStateOk() (*VMToolsState, bool)`

GetVMToolsStateOk returns a tuple with the VMToolsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMToolsState

`func (o *MachineDetailResponseModelAllOf) SetVMToolsState(v VMToolsState)`

SetVMToolsState sets VMToolsState field to given value.

### HasVMToolsState

`func (o *MachineDetailResponseModelAllOf) HasVMToolsState() bool`

HasVMToolsState returns a boolean if a field has been set.

### GetFailureReason

`func (o *MachineDetailResponseModelAllOf) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MachineDetailResponseModelAllOf) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *MachineDetailResponseModelAllOf) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *MachineDetailResponseModelAllOf) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


