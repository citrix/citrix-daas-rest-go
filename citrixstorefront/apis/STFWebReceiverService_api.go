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
	createSTFWebReceiverRequestModel models.CreateSTFWebReceiverRequestModel
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
	var param = StructToString(r.createSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFWebReceiverService", param, "-Confirm:$false")
}

func (a *STFWebReceiver) STFWebReceiverCreateSTFWebReceiver(ctx context.Context, createSTFWebReceiverRequestModel models.CreateSTFWebReceiverRequestModel) ApiCreateSTFWebReceiverRequest {
	return ApiCreateSTFWebReceiverRequest{
		ApiService:                       a,
		ctx:                              ctx,
		createSTFWebReceiverRequestModel: createSTFWebReceiverRequestModel,
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
	ctx                             context.Context
	ApiService                      *STFWebReceiver
	clearSTFWebReceiverRequestModel models.ClearSTFWebReceiverRequestModel
}

func (r ApiClearSTFWebReceiverRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.ClearSTFWebReceiverExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFWebReceiver) ClearSTFWebReceiverExecute(r ApiClearSTFWebReceiverRequest) ([]byte, error) {
	var param = StructToString(r.clearSTFWebReceiverRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Remove-STFWebReceiverService", param, "-Confirm:$false")
}

func (a *STFWebReceiver) STFWebReceiverClearSTFWebReceiver(ctx context.Context, clearSTFWebReceiverRequestModel models.ClearSTFWebReceiverRequestModel) ApiClearSTFWebReceiverRequest {
	return ApiClearSTFWebReceiverRequest{
		ApiService:                      a,
		ctx:                             ctx,
		clearSTFWebReceiverRequestModel: clearSTFWebReceiverRequestModel,
	}
}

//STFWebReceiverAuthenticationMethods

// Get
type ApiGetSTFWebReceiverAuthenticationMethodsRequest struct {
	ctx                                                context.Context
	ApiService                                         *STFWebReceiver
	getSTFWebReceiverAuthenticationMethodsRequestModel models.GetSTFWebReceiverAuthenticationMethodsRequestModel
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
	var param = StructToString(r.getSTFWebReceiverAuthenticationMethodsRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverAuthenticationMethods", param)
}

func (a *STFWebReceiver) STFWebReceiverGetSTFWebReceiverAuthenticationMethods(ctx context.Context, getSTFWebReceiverAuthenticationMethodsRequestModel models.GetSTFWebReceiverAuthenticationMethodsRequestModel) ApiGetSTFWebReceiverAuthenticationMethodsRequest {
	return ApiGetSTFWebReceiverAuthenticationMethodsRequest{
		ApiService: a,
		ctx:        ctx,
		getSTFWebReceiverAuthenticationMethodsRequestModel: getSTFWebReceiverAuthenticationMethodsRequestModel,
	}
}

// Set-STFWebReceiverAuthenticationMethods
type ApiSetSTFWebReceiverAuthenticationMethodsRequest struct {
	ctx                                                context.Context
	ApiService                                         *STFWebReceiver
	setSTFWebReceiverAuthenticationMethodsRequestModel models.UpdateSTFWebReceiverAuthenticationMethodsRequestModel
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
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverAuthenticationMethods", param, "-Confirm:$false")
}

func (a *STFWebReceiver) STFWebReceiverSetSTFWebReceiverAuthenticationMethods(ctx context.Context, setSTFWebReceiverAuthenticationMethodsRequestModel models.UpdateSTFWebReceiverAuthenticationMethodsRequestModel) ApiSetSTFWebReceiverAuthenticationMethodsRequest {
	return ApiSetSTFWebReceiverAuthenticationMethodsRequest{
		ApiService: a,
		ctx:        ctx,
		setSTFWebReceiverAuthenticationMethodsRequestModel: setSTFWebReceiverAuthenticationMethodsRequestModel,
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
	webReceiverPluginAssistantUpdateModel models.UpdateSTFWebReceiverPluginAssistantRequestModel
}

func (r ApiWebReceiverPluginAssistantUpdateRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.WebReceiverPluginAssistantUpdateExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFWebReceiver) WebReceiverPluginAssistantUpdateExecute(r ApiWebReceiverPluginAssistantUpdateRequest) ([]byte, error) {
	var param = StructToString(r.webReceiverPluginAssistantUpdateModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFWebReceiverPluginAssistant", param, "-Confirm:$false")
}

func (a *STFWebReceiver) STFWebReceiverPluginAssistantUpdate(ctx context.Context, webReceiverPluginAssistantUpdateModel models.UpdateSTFWebReceiverPluginAssistantRequestModel) ApiWebReceiverPluginAssistantUpdateRequest {
	return ApiWebReceiverPluginAssistantUpdateRequest{
		ApiService:                            a,
		ctx:                                   ctx,
		webReceiverPluginAssistantUpdateModel: webReceiverPluginAssistantUpdateModel,
	}
}

// Get-STFWebReceiverPluginAssistant
type ApiWebReceiverPluginAssistantGetRequest struct {
	ctx                                context.Context
	ApiService                         *STFWebReceiver
	webReceiverPluginAssistantGetModel models.GetSTFWebReceiverPluginAssistantRequestModel
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
	var param = StructToString(r.webReceiverPluginAssistantGetModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFWebReceiverPluginAssistant", param)
}

func (a *STFWebReceiver) STFWebReceiverPluginAssistantGet(ctx context.Context, webReceiverPluginAssistantGetModel models.GetSTFWebReceiverPluginAssistantRequestModel) ApiWebReceiverPluginAssistantGetRequest {
	return ApiWebReceiverPluginAssistantGetRequest{
		ApiService:                         a,
		ctx:                                ctx,
		webReceiverPluginAssistantGetModel: webReceiverPluginAssistantGetModel,
	}
}
