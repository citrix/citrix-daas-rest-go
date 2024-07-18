// Copyright © 2024. Citrix Systems, Inc.
package apis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/citrix/citrix-daas-rest-go/citrixstorefront/models"
)

type STFWebReceiver Service

// Add-STFWebReceiverService
type ApiCreateSTFWebReceiverRequest struct {
	ctx                              context.Context
	ApiService                       *STFWebReceiver
	CreateSTFWebReceiverRequestModel models.CreateSTFWebReceiverRequestModel
	GetSTFStoreServiceRequestModel   models.GetSTFStoreRequestModel
}

func (r ApiCreateSTFWebReceiverRequest) Execute() (models.STFWebReceiverDetailModel, error) {
	bytes, err := r.ApiService.CreateSTFWebReceiverExecute(r)
	if err != nil {
		return models.STFWebReceiverDetailModel{}, err
	}
	var reponse = models.STFWebReceiverDetailModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return models.STFWebReceiverDetailModel{}, unMarshalErr
	}
	return reponse, nil
}

func (a *STFWebReceiver) CreateSTFWebReceiverExecute(r ApiCreateSTFWebReceiverRequest) ([]byte, error) {
	var param = StructToString(r.CreateSTFWebReceiverRequestModel)
	var getSTFStoreParams = StructToString(r.GetSTFStoreServiceRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFWebReceiverService", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", getSTFStoreParams), param, "-Confirm:$false")
}

func (a *STFWebReceiver) STFWebReceiverCreateSTFWebReceiver(ctx context.Context, createSTFWebReceiverRequestModel models.CreateSTFWebReceiverRequestModel, getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel) ApiCreateSTFWebReceiverRequest {
	return ApiCreateSTFWebReceiverRequest{
		ApiService:                       a,
		ctx:                              ctx,
		CreateSTFWebReceiverRequestModel: createSTFWebReceiverRequestModel,
		GetSTFStoreServiceRequestModel:   getSTFStoreServiceRequestModel,
	}
}

// Get-STFWebReceiverService
type ApiGetSTFWebReceiverRequest struct {
	ctx                           context.Context
	ApiService                    *STFWebReceiver
	getSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiGetSTFWebReceiverRequest) Execute() (models.STFWebReceiverDetailModel, error) {
	bytes, err := r.ApiService.GetSTFWebReceiverExecute(r)
	if err != nil {
		return models.STFWebReceiverDetailModel{}, err
	}
	if len(bytes) == 0 {
		return models.STFWebReceiverDetailModel{}, fmt.Errorf(NOT_EXIST)
	}
	var reponse = models.STFWebReceiverDetailModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return models.STFWebReceiverDetailModel{}, unMarshalErr
	}
	return reponse, nil
}

func (a *STFWebReceiver) GetSTFWebReceiverExecute(r ApiGetSTFWebReceiverRequest) ([]byte, error) {
	var param = StructToString(r.getSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverService", param)
}

func (a *STFWebReceiver) STFWebReceiverGetSTFWebReceiver(ctx context.Context, getSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel) ApiGetSTFWebReceiverRequest {
	return ApiGetSTFWebReceiverRequest{
		ApiService:                    a,
		ctx:                           ctx,
		getSTFWebReceiverRequestModel: getSTFWebReceiverRequestModel,
	}
}

// Set-STFWebReceiverService
type ApiSetSTFWebReceiverRequest struct {
	ctx                           context.Context
	ApiService                    *STFWebReceiver
	setSTFWebReceiverRequestModel models.SetSTFWebReceiverRequestModel
}

func (r ApiSetSTFWebReceiverRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.SetSTFWebReceiverExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFWebReceiver) SetSTFWebReceiverExecute(r ApiSetSTFWebReceiverRequest) ([]byte, error) {
	var param = StructToString(r.setSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverService", param, "-Confirm:$false")
}

func (a *STFWebReceiver) STFWebReceiverSetSTFWebReceiver(ctx context.Context, setSTFWebReceiverRequestModel models.SetSTFWebReceiverRequestModel) ApiSetSTFWebReceiverRequest {
	return ApiSetSTFWebReceiverRequest{
		ApiService:                    a,
		ctx:                           ctx,
		setSTFWebReceiverRequestModel: setSTFWebReceiverRequestModel,
	}
}

// Remove-STFWebReceiverService
type ApiClearSTFWebReceiverRequest struct {
	ctx                           context.Context
	ApiService                    *STFWebReceiver
	GetSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiClearSTFWebReceiverRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.ClearSTFWebReceiverExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFWebReceiver) ClearSTFWebReceiverExecute(r ApiClearSTFWebReceiverRequest) ([]byte, error) {
	var getParam = StructToString(r.GetSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Remove-STFWebReceiverService", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getParam), "-Confirm:$false")
}

func (a *STFWebReceiver) STFWebReceiverClearSTFWebReceiver(ctx context.Context, getSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel) ApiClearSTFWebReceiverRequest {
	return ApiClearSTFWebReceiverRequest{
		ApiService:                    a,
		ctx:                           ctx,
		GetSTFWebReceiverRequestModel: getSTFWebReceiverRequestModel,
	}
}

//STFWebReceiverAuthenticationMethods

// Get
type ApiGetSTFWebReceiverAuthenticationMethodsRequest struct {
	ctx                           context.Context
	ApiService                    *STFWebReceiver
	GetSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiGetSTFWebReceiverAuthenticationMethodsRequest) Execute() (models.STFWebReceiverAuthenticationMethodsResponse, error) {
	bytes, err := r.ApiService.GetSTFWebReceiverAuthenticationMethodsExecute(r)
	if err != nil {
		return models.STFWebReceiverAuthenticationMethodsResponse{}, err
	}
	var reponse = models.STFWebReceiverAuthenticationMethodsResponse{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.STFWebReceiverAuthenticationMethodsResponse{}, unMarshalErr
	}
	return reponse, nil
}

func (a *STFWebReceiver) GetSTFWebReceiverAuthenticationMethodsExecute(r ApiGetSTFWebReceiverAuthenticationMethodsRequest) ([]byte, error) {
	var param = StructToString(r.GetSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverAuthenticationMethods", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", param))
}

func (a *STFWebReceiver) STFWebReceiverGetSTFWebReceiverAuthenticationMethods(ctx context.Context, getSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel) ApiGetSTFWebReceiverAuthenticationMethodsRequest {
	return ApiGetSTFWebReceiverAuthenticationMethodsRequest{
		ApiService:                    a,
		ctx:                           ctx,
		GetSTFWebReceiverRequestModel: getSTFWebReceiverRequestModel,
	}
}

// Set-STFWebReceiverAuthenticationMethods
type ApiSetSTFWebReceiverAuthenticationMethodsRequest struct {
	ctx                                                context.Context
	ApiService                                         *STFWebReceiver
	setSTFWebReceiverAuthenticationMethodsRequestModel models.UpdateSTFWebReceiverAuthenticationMethodsRequestModel
	GetSTFWebReceiverRequestModel                      models.GetSTFWebReceiverRequestModel
}

func (r ApiSetSTFWebReceiverAuthenticationMethodsRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.UpdateSTFWebReceiverAuthenticationMethodsExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFWebReceiver) UpdateSTFWebReceiverAuthenticationMethodsExecute(r ApiSetSTFWebReceiverAuthenticationMethodsRequest) ([]byte, error) {
	var param = StructToString(r.setSTFWebReceiverAuthenticationMethodsRequestModel)
	var getWebReceiverParams = StructToString(r.GetSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverAuthenticationMethods", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams), param, "-Confirm:$false")
}

func (a *STFWebReceiver) STFWebReceiverSetSTFWebReceiverAuthenticationMethods(ctx context.Context, setSTFWebReceiverAuthenticationMethodsRequestModel models.UpdateSTFWebReceiverAuthenticationMethodsRequestModel, getSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel) ApiSetSTFWebReceiverAuthenticationMethodsRequest {
	return ApiSetSTFWebReceiverAuthenticationMethodsRequest{
		ApiService: a,
		ctx:        ctx,
		setSTFWebReceiverAuthenticationMethodsRequestModel: setSTFWebReceiverAuthenticationMethodsRequestModel,
		GetSTFWebReceiverRequestModel:                      getSTFWebReceiverRequestModel,
	}
}

// Get-STFWebReceiverAuthenticationMethodsAvailable
type ApiGetSTFWebReceiverAuthenticationMethodsAvailableRequest struct {
	ctx        context.Context
	ApiService *STFWebReceiver
}

func (r ApiGetSTFWebReceiverAuthenticationMethodsAvailableRequest) Execute() ([]string, error) {
	bytes, err := r.ApiService.GetSTFWebReceiverAuthenticationMethodsAvailableExecute(r)
	if err != nil {
		return nil, err
	}
	var reponse = []string{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return nil, unMarshalErr
	}
	return reponse, nil
}

func (a *STFWebReceiver) GetSTFWebReceiverAuthenticationMethodsAvailableExecute(r ApiGetSTFWebReceiverAuthenticationMethodsAvailableRequest) ([]byte, error) {
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverAuthenticationMethodsAvailable")
}

func (a *STFWebReceiver) STFWebReceiverGetSTFWebReceiverAuthenticationMethodsAvailable(ctx context.Context) ApiGetSTFWebReceiverAuthenticationMethodsAvailableRequest {
	return ApiGetSTFWebReceiverAuthenticationMethodsAvailableRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Set-STFWebReceiverPluginAssistant
type ApiWebReceiverPluginAssistantUpdateRequest struct {
	ctx                                   context.Context
	ApiService                            *STFWebReceiver
	WebReceiverPluginAssistantUpdateModel models.UpdateSTFWebReceiverPluginAssistantRequestModel
	GetSTFWebReceiverRequestModel         models.GetSTFWebReceiverRequestModel
}

func (r ApiWebReceiverPluginAssistantUpdateRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.WebReceiverPluginAssistantUpdateExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFWebReceiver) WebReceiverPluginAssistantUpdateExecute(r ApiWebReceiverPluginAssistantUpdateRequest) ([]byte, error) {
	var param = StructToString(r.WebReceiverPluginAssistantUpdateModel)
	var getWebReceiverParams = StructToString(r.GetSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverPluginAssistant", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams), param, "-Confirm:$false")
}

func (a *STFWebReceiver) STFWebReceiverPluginAssistantUpdate(ctx context.Context, webReceiverPluginAssistantUpdateModel models.UpdateSTFWebReceiverPluginAssistantRequestModel, getSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel) ApiWebReceiverPluginAssistantUpdateRequest {
	return ApiWebReceiverPluginAssistantUpdateRequest{
		ApiService:                            a,
		ctx:                                   ctx,
		WebReceiverPluginAssistantUpdateModel: webReceiverPluginAssistantUpdateModel,
		GetSTFWebReceiverRequestModel:         getSTFWebReceiverRequestModel,
	}
}

// Get-STFWebReceiverPluginAssistant
type ApiWebReceiverPluginAssistantGetRequest struct {
	ctx                           context.Context
	ApiService                    *STFWebReceiver
	GetSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiWebReceiverPluginAssistantGetRequest) Execute() (models.WebReceiverPluginAssistantModel, error) {
	bytes, err := r.ApiService.WebReceiverPluginAssistantGetExecute(r)
	if err != nil {
		return models.WebReceiverPluginAssistantModel{}, err
	}
	var reponse = models.WebReceiverPluginAssistantModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return models.WebReceiverPluginAssistantModel{}, unMarshalErr
	}
	return reponse, nil
}

func (a *STFWebReceiver) WebReceiverPluginAssistantGetExecute(r ApiWebReceiverPluginAssistantGetRequest) ([]byte, error) {
	var getWebReceiverParams = StructToString(r.GetSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverPluginAssistant", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams))
}

func (a *STFWebReceiver) STFWebReceiverPluginAssistantGet(ctx context.Context, getSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel) ApiWebReceiverPluginAssistantGetRequest {
	return ApiWebReceiverPluginAssistantGetRequest{
		ApiService:                    a,
		ctx:                           ctx,
		GetSTFWebReceiverRequestModel: getSTFWebReceiverRequestModel,
	}
}

// Set-STFWebReceiverApplicationShortcuts
type ApiWebReceiverApplicationShortcutsSetRequest struct {
	ctx                                            context.Context
	ApiService                                     *STFWebReceiver
	WebReceiverServiceGetRequestModel              models.GetSTFWebReceiverRequestModel
	WebReceiverApplicationShortcutsSetRequestModel models.SetWebReceiverApplicationShortcutsRequestModel
}

func (r ApiWebReceiverApplicationShortcutsSetRequest) Execute() error {
	_, err := r.ApiService.WebReceiverApplicationShortcutsSetExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFWebReceiver) WebReceiverApplicationShortcutsSetExecute(r ApiWebReceiverApplicationShortcutsSetRequest) ([]byte, error) {
	var getWebReceiverParams = StructToString(r.WebReceiverServiceGetRequestModel)
	var setParams = StructToString(r.WebReceiverApplicationShortcutsSetRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverApplicationShortcuts", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams), setParams)
}

func (a *STFWebReceiver) STFWebReceiverApplicationShortcutsSet(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel, webReceiverApplicationShortcutsSetRequestModel models.SetWebReceiverApplicationShortcutsRequestModel) ApiWebReceiverApplicationShortcutsSetRequest {
	return ApiWebReceiverApplicationShortcutsSetRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
		WebReceiverApplicationShortcutsSetRequestModel: webReceiverApplicationShortcutsSetRequestModel,
	}
}

// Get-STFWebReceiverApplicationShortcuts
type ApiWebReceiverApplicationShortcutsGetRequest struct {
	ctx                               context.Context
	ApiService                        *STFWebReceiver
	WebReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiWebReceiverApplicationShortcutsGetRequest) Execute() (models.GetWebReceiverApplicationShortcutsResponseModel, error) {
	bytes, err := r.ApiService.WebReceiverApplicationShortcutsGetExecute(r)
	if err != nil {
		return models.GetWebReceiverApplicationShortcutsResponseModel{}, err
	}
	var reponse = models.GetWebReceiverApplicationShortcutsResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return models.GetWebReceiverApplicationShortcutsResponseModel{}, unMarshalErr
	}
	return reponse, nil
}

func (a *STFWebReceiver) WebReceiverApplicationShortcutsGetExecute(r ApiWebReceiverApplicationShortcutsGetRequest) ([]byte, error) {
	var getWebReceiverParams = StructToString(r.WebReceiverServiceGetRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverApplicationShortcuts", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams))
}

func (a *STFWebReceiver) STFWebReceiverApplicationShortcutsGet(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel) ApiWebReceiverApplicationShortcutsGetRequest {
	return ApiWebReceiverApplicationShortcutsGetRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
	}
}

// Set-STFWebReceiverCommunication
type ApiWebReceiverCommunicationSetRequest struct {
	ctx                                     context.Context
	ApiService                              *STFWebReceiver
	WebReceiverServiceGetRequestModel       models.GetSTFWebReceiverRequestModel
	SetWebReceiverCommunicationRequestModel models.SetWebReceiverCommunicationRequestModel
}

func (r ApiWebReceiverCommunicationSetRequest) Execute() error {
	_, err := r.ApiService.WebReceiverCommunicationSetExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFWebReceiver) WebReceiverCommunicationSetExecute(r ApiWebReceiverCommunicationSetRequest) ([]byte, error) {
	getWebReceiverParams := StructToString(r.WebReceiverServiceGetRequestModel)
	setWebReceiverCommunicationParams := StructToString(r.SetWebReceiverCommunicationRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverCommunication", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams), setWebReceiverCommunicationParams)
}

func (a *STFWebReceiver) STFWebReceiverCommunicationSet(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel, setWebReceiverCommunicationRequestModel models.SetWebReceiverCommunicationRequestModel) ApiWebReceiverCommunicationSetRequest {
	return ApiWebReceiverCommunicationSetRequest{
		ApiService:                              a,
		ctx:                                     ctx,
		WebReceiverServiceGetRequestModel:       webReceiverServiceGetRequestModel,
		SetWebReceiverCommunicationRequestModel: setWebReceiverCommunicationRequestModel,
	}
}

// Get-STFWebReceiverCommunication
type ApiWebReceiverCommunicationGetRequest struct {
	ctx                               context.Context
	ApiService                        *STFWebReceiver
	WebReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiWebReceiverCommunicationGetRequest) Execute() (models.GetWebReceiverCommunicationResponseModel, error) {
	bytes, err := r.ApiService.WebReceiverCommunicationGetExecute(r)
	if err != nil {
		return models.GetWebReceiverCommunicationResponseModel{}, err
	}
	rawResponse := models.GetWebReceiverCommunicationRawResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &rawResponse)
	if unMarshalErr != nil {
		return models.GetWebReceiverCommunicationResponseModel{}, unMarshalErr
	}
	return rawResponse.ConvertToResponseModel(), nil
}

func (a *STFWebReceiver) WebReceiverCommunicationGetExecute(r ApiWebReceiverCommunicationGetRequest) ([]byte, error) {
	getWebReceiverParams := StructToString(r.WebReceiverServiceGetRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverCommunication", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams))
}

func (a *STFWebReceiver) STFWebReceiverCommunicationGet(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel) ApiWebReceiverCommunicationGetRequest {
	return ApiWebReceiverCommunicationGetRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
	}
}

// Set-STFWebReceiverStrictTransportSecurity
type ApiWebReceiverStrictTransportSecuritySetRequest struct {
	ctx                                               context.Context
	ApiService                                        *STFWebReceiver
	WebReceiverServiceGetRequestModel                 models.GetSTFWebReceiverRequestModel
	SetWebReceiverStrictTransportSecurityRequestModel models.SetWebReceiverStrictTransportSecurityRequestModel
}

func (r ApiWebReceiverStrictTransportSecuritySetRequest) Execute() error {
	_, err := r.ApiService.WebReceiverStrictTransportSecuritySetExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFWebReceiver) WebReceiverStrictTransportSecuritySetExecute(r ApiWebReceiverStrictTransportSecuritySetRequest) ([]byte, error) {
	getWebReceiverParams := StructToString(r.WebReceiverServiceGetRequestModel)
	setWebReceiverSTSParams := StructToString(r.SetWebReceiverStrictTransportSecurityRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverStrictTransportSecurity", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams), setWebReceiverSTSParams)
}

func (a *STFWebReceiver) STFWebReceiverStrictTransportSecuritySet(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel, setWebReceiverStrictTransportSecurityRequestModel models.SetWebReceiverStrictTransportSecurityRequestModel) ApiWebReceiverStrictTransportSecuritySetRequest {
	return ApiWebReceiverStrictTransportSecuritySetRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
		SetWebReceiverStrictTransportSecurityRequestModel: setWebReceiverStrictTransportSecurityRequestModel,
	}
}

// Get-STFWebReceiverStrictTransportSecurity
type ApiWebReceiverStrictTransportSecurityGetRequest struct {
	ctx                               context.Context
	ApiService                        *STFWebReceiver
	WebReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiWebReceiverStrictTransportSecurityGetRequest) Execute() (models.GetWebReceiverStrictTransportSecurityResponseModel, error) {
	bytes, err := r.ApiService.WebReceiverStrictTransportSecurityGetExecute(r)
	if err != nil {
		return models.GetWebReceiverStrictTransportSecurityResponseModel{}, err
	}
	rawResponse := models.GetWebReceiverStrictTransportSecurityRawResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &rawResponse)
	if unMarshalErr != nil {
		return models.GetWebReceiverStrictTransportSecurityResponseModel{}, unMarshalErr
	}
	return rawResponse.ConvertToResponseModel(), nil
}

func (a *STFWebReceiver) WebReceiverStrictTransportSecurityGetExecute(r ApiWebReceiverStrictTransportSecurityGetRequest) ([]byte, error) {
	getWebReceiverParams := StructToString(r.WebReceiverServiceGetRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverStrictTransportSecurity", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams))
}

func (a *STFWebReceiver) STFWebReceiverStrictTransportSecurityGet(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel) ApiWebReceiverStrictTransportSecurityGetRequest {
	return ApiWebReceiverStrictTransportSecurityGetRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
	}
}

// Set-STFWebReceiverAuthenticationManager
type ApiWebReceiverAuthenticationManagerSetRequest struct {
	ctx                                             context.Context
	ApiService                                      *STFWebReceiver
	WebReceiverServiceGetRequestModel               models.GetSTFWebReceiverRequestModel
	SetWebReceiverAuthenticationManagerRequestModel models.SetWebReceiverAuthenticationManagerRequestModel
}

func (r ApiWebReceiverAuthenticationManagerSetRequest) Execute() error {
	_, err := r.ApiService.WebReceiverAuthenticationManagerSetExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFWebReceiver) WebReceiverAuthenticationManagerSetExecute(r ApiWebReceiverAuthenticationManagerSetRequest) ([]byte, error) {
	getWebReceiverParams := StructToString(r.WebReceiverServiceGetRequestModel)
	setAuthManagerParams := StructToString(r.SetWebReceiverAuthenticationManagerRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverAuthenticationManager", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams), setAuthManagerParams)
}

func (a *STFWebReceiver) STFWebReceiverAuthenticationManagerSet(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel, setWebReceiverAuthenticationManagerRequestModel models.SetWebReceiverAuthenticationManagerRequestModel) ApiWebReceiverAuthenticationManagerSetRequest {
	return ApiWebReceiverAuthenticationManagerSetRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
		SetWebReceiverAuthenticationManagerRequestModel: setWebReceiverAuthenticationManagerRequestModel,
	}
}

// Get-STFWebReceiverAuthenticationManager
type ApiWebReceiverAuthenticationManagerGetRequest struct {
	ctx                               context.Context
	ApiService                        *STFWebReceiver
	WebReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiWebReceiverAuthenticationManagerGetRequest) Execute() (models.GetWebReceiverAuthenticationManagerResponseModel, error) {
	bytes, err := r.ApiService.WebReceiverAuthenticationManagerGetExecute(r)
	if err != nil {
		return models.GetWebReceiverAuthenticationManagerResponseModel{}, err
	}
	response := models.GetWebReceiverAuthenticationManagerResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &response)
	if unMarshalErr != nil {
		return models.GetWebReceiverAuthenticationManagerResponseModel{}, unMarshalErr
	}
	return response, nil
}

func (a *STFWebReceiver) WebReceiverAuthenticationManagerGetExecute(r ApiWebReceiverAuthenticationManagerGetRequest) ([]byte, error) {
	getWebReceiverParams := StructToString(r.WebReceiverServiceGetRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverAuthenticationManager", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams))
}

func (a *STFWebReceiver) STFWebReceiverAuthenticationManagerGet(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel) ApiWebReceiverAuthenticationManagerGetRequest {
	return ApiWebReceiverAuthenticationManagerGetRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
	}
}

// Set-STFWebReceiverUserInterface
type ApiWebReceiverUserInterfaceSetRequest struct {
	ctx                                        context.Context
	ApiService                                 *STFWebReceiver
	WebReceiverServiceGetRequestModel          models.GetSTFWebReceiverRequestModel
	SetSTFWebReceiverUserInterfaceRequestModel models.SetSTFWebReceiverUserInterfaceRequestModel
}

func (r ApiWebReceiverUserInterfaceSetRequest) Execute() error {
	_, err := r.ApiService.WebReceiverUserInterfaceSetExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFWebReceiver) WebReceiverUserInterfaceSetExecute(r ApiWebReceiverUserInterfaceSetRequest) ([]byte, error) {
	var getWebReceiverParams = StructToString(r.WebReceiverServiceGetRequestModel)
	var setParams = StructToString(r.SetSTFWebReceiverUserInterfaceRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverUserInterface", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams), setParams)
}

func (a *STFWebReceiver) STFWebReceiverUserInterfaceSet(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel, setSTFWebReceiverUserInterfaceRequestModel models.SetSTFWebReceiverUserInterfaceRequestModel) ApiWebReceiverUserInterfaceSetRequest {
	return ApiWebReceiverUserInterfaceSetRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
		SetSTFWebReceiverUserInterfaceRequestModel: setSTFWebReceiverUserInterfaceRequestModel,
	}
}

// Get-STFWebReceiverUserInterface
type ApiWebReceiverUserInterfaceGetRequest struct {
	ctx                               context.Context
	ApiService                        *STFWebReceiver
	WebReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiWebReceiverUserInterfaceGetRequest) Execute() (models.GetSTFWebReceiverUserInterfaceResponseModel, error) {
	bytes, err := r.ApiService.WebReceiverUserInterfaceGetExecute(r)
	if err != nil {
		return models.GetSTFWebReceiverUserInterfaceResponseModel{}, err
	}
	var reponse = models.GetSTFWebReceiverUserInterfaceRawResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return models.GetSTFWebReceiverUserInterfaceResponseModel{}, unMarshalErr
	}
	return reponse.ConvertToResponseModel(), nil
}

func (a *STFWebReceiver) WebReceiverUserInterfaceGetExecute(r ApiWebReceiverUserInterfaceGetRequest) ([]byte, error) {
	var getWebReceiverParams = StructToString(r.WebReceiverServiceGetRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverUserInterface", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams))
}

func (a *STFWebReceiver) STFWebReceiverUserInterfaceGet(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel) ApiWebReceiverUserInterfaceGetRequest {
	return ApiWebReceiverUserInterfaceGetRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
	}
}

// Set-STFWebReceiverResourcesService
type ApiWebReceiverResourcesServiceRequest struct {
	ctx                                           context.Context
	ApiService                                    *STFWebReceiver
	WebReceiverServiceGetRequestModel             models.GetSTFWebReceiverRequestModel
	SetSTFWebReceiverResourcesServiceRequestModel models.SetSTFWebReceiverResourcesServiceRequestModel
}

func (r ApiWebReceiverResourcesServiceRequest) Execute() error {
	_, err := r.ApiService.WebReceiverResourcesServiceExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFWebReceiver) WebReceiverResourcesServiceExecute(r ApiWebReceiverResourcesServiceRequest) ([]byte, error) {
	var setParams = StructToString(r.SetSTFWebReceiverResourcesServiceRequestModel)
	var getWebReceiverParams = StructToString(r.WebReceiverServiceGetRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverResourcesService", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams), setParams)
}

func (a *STFWebReceiver) SetSTFWebReceiverResourcesService(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel, setSTFWebReceiverResourcesServiceRequestModel models.SetSTFWebReceiverResourcesServiceRequestModel) ApiWebReceiverResourcesServiceRequest {
	return ApiWebReceiverResourcesServiceRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
		SetSTFWebReceiverResourcesServiceRequestModel: setSTFWebReceiverResourcesServiceRequestModel,
	}
}

// Get-STFWebReceiverResourcesService
type ApiWebReceiverResourcesServiceGetRequest struct {
	ctx                               context.Context
	ApiService                        *STFWebReceiver
	WebReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiWebReceiverResourcesServiceGetRequest) Execute() (models.GetSTFWebReceiverResourcesServiceResponseModel, error) {
	bytes, err := r.ApiService.WebReceiverResourcesServiceGetExecute(r)
	if err != nil {
		return models.GetSTFWebReceiverResourcesServiceResponseModel{}, err
	}
	var reponse = models.GetSTFWebReceiverResourcesServiceResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return models.GetSTFWebReceiverResourcesServiceResponseModel{}, unMarshalErr
	}
	return reponse, nil
}

func (a *STFWebReceiver) WebReceiverResourcesServiceGetExecute(r ApiWebReceiverResourcesServiceGetRequest) ([]byte, error) {
	var getWebReceiverParams = StructToString(r.WebReceiverServiceGetRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverResourcesService", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getWebReceiverParams))
}

func (a *STFWebReceiver) GetSTFWebReceiverResourcesService(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel) ApiWebReceiverResourcesServiceGetRequest {
	return ApiWebReceiverResourcesServiceGetRequest{
		ApiService:                        a,
		ctx:                               ctx,
		WebReceiverServiceGetRequestModel: webReceiverServiceGetRequestModel,
	}
}

// Set-STFWebReceiverSiteStyle
type ApiSetSTFWebReceiverSiteStyleRequest struct {
	ctx                                    context.Context
	ApiService                             *STFWebReceiver
	setSTFWebReceiverSiteStyleRequestModel models.SetSTFWebReceiverSiteStyleRequestModel
	getSTFWebReceiverRequestModel          models.GetSTFWebReceiverRequestModel
}

func (r ApiSetSTFWebReceiverSiteStyleRequest) Execute() error {
	_, err := r.ApiService.SetSTFWebReceiverSiteStyleExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFWebReceiver) SetSTFWebReceiverSiteStyleExecute(r ApiSetSTFWebReceiverSiteStyleRequest) ([]byte, error) {
	var param = StructToString(r.setSTFWebReceiverSiteStyleRequestModel)
	var setSiteStyleParams = StructToString(r.getSTFWebReceiverRequestModel)
	if r.getSTFWebReceiverRequestModel.VirtualPath.IsSet() && *r.getSTFWebReceiverRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverSiteStyle", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", setSiteStyleParams), param)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverSiteStyle", param, "-Confirm:$false")
	}

}

func (a *STFWebReceiver) STFWebReceiverSetSTFWebReceiverSiteStyle(ctx context.Context, webReceiverServiceGetRequestModel models.GetSTFWebReceiverRequestModel, setSTFWebReceiverSiteStyleRequestModel models.SetSTFWebReceiverSiteStyleRequestModel) ApiSetSTFWebReceiverSiteStyleRequest {
	return ApiSetSTFWebReceiverSiteStyleRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		setSTFWebReceiverSiteStyleRequestModel: setSTFWebReceiverSiteStyleRequestModel,
		getSTFWebReceiverRequestModel:          webReceiverServiceGetRequestModel,
	}
}

// Get-STFWebReceiverSiteStyle
type ApiGetSTFWebReceiverSiteStyleRequest struct {
	ctx                           context.Context
	ApiService                    *STFWebReceiver
	getSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiGetSTFWebReceiverSiteStyleRequest) Execute() (models.STFWebReceiverSiteStyleResponseModel, error) {
	bytes, err := r.ApiService.GetSTFWebReceiverSiteStyleExecute(r)
	if err != nil {
		return models.STFWebReceiverSiteStyleResponseModel{}, err
	}
	var response = models.STFWebReceiverSiteStyleResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &response)
	if unMarshalErr != nil {
		return models.STFWebReceiverSiteStyleResponseModel{}, unMarshalErr
	}
	return response, nil
}

func (a *STFWebReceiver) GetSTFWebReceiverSiteStyleExecute(r ApiGetSTFWebReceiverSiteStyleRequest) ([]byte, error) {
	var param = StructToString(r.getSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverSiteStyle", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", param))
}

func (a *STFWebReceiver) STFWebReceiverGetSTFWebReceiverSiteStyle(ctx context.Context, getSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel) ApiGetSTFWebReceiverSiteStyleRequest {
	return ApiGetSTFWebReceiverSiteStyleRequest{
		ApiService:                    a,
		ctx:                           ctx,
		getSTFWebReceiverRequestModel: getSTFWebReceiverRequestModel,
	}
}

// Remove-STFWebReceiverService
type ApiClearSTFWebReceiverSiteStyleRequest struct {
	ctx                           context.Context
	ApiService                    *STFWebReceiver
	GetSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel
}

func (r ApiClearSTFWebReceiverSiteStyleRequest) Execute() ([]byte, error) {
	byte, err := r.ApiService.ClearSTFWebReceiverSiteStyleExecute(r)
	if err != nil {
		return byte, err
	}
	return nil, nil
}

func (a *STFWebReceiver) ClearSTFWebReceiverSiteStyleExecute(r ApiClearSTFWebReceiverSiteStyleRequest) ([]byte, error) {
	var getParam = StructToString(r.GetSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Clear-STFWebReceiverSiteStyle", fmt.Sprintf("-WebReceiverService (Get-STFWebReceiverService %s)", getParam), "-Confirm:$false")
}

func (a *STFWebReceiver) STFWebReceiverClearSTFWebReceiverSiteStyle(ctx context.Context, getSTFWebReceiverRequestModel models.GetSTFWebReceiverRequestModel) ApiClearSTFWebReceiverSiteStyleRequest {
	return ApiClearSTFWebReceiverSiteStyleRequest{
		ApiService:                    a,
		ctx:                           ctx,
		GetSTFWebReceiverRequestModel: getSTFWebReceiverRequestModel,
	}
}
