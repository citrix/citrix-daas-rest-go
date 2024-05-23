package apis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/citrix/citrix-daas-rest-go/citrixstorefront/models"
)

type STFStore Service

// Create
type ApiCreateSTFStoreRequest struct {
	ctx                        context.Context
	ApiService                 *STFStore
	createSTFStoreRequestModel models.CreateSTFStoreRequestModel
}

func (r ApiCreateSTFStoreRequest) Execute() (models.STFStoreDetailModel, error) {
	bytes, err := r.ApiService.CreateSTFStoreExecute(r)
	if err != nil {
		return models.STFStoreDetailModel{}, err
	}
	var reponse = models.STFStoreDetailModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.STFStoreDetailModel{}, fmt.Errorf("error unmarshal STFStoreDetailModel: %v", err)
	}
	return reponse, nil
}

func (a *STFStore) CreateSTFStoreExecute(r ApiCreateSTFStoreRequest) ([]byte, error) {
	var param = StructToString(r.createSTFStoreRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFStoreService", param)
}

func (a *STFStore) STFStoreCreateSTFStore(ctx context.Context, createSTFStoreRequestModel models.CreateSTFStoreRequestModel) ApiCreateSTFStoreRequest {
	return ApiCreateSTFStoreRequest{
		ApiService:                 a,
		ctx:                        ctx,
		createSTFStoreRequestModel: createSTFStoreRequestModel,
	}
}

// Get
type ApiGetSTFStoreRequest struct {
	ctx                     context.Context
	ApiService              *STFStore
	getSTFStoreRequestModel models.GetSTFStoreRequestModel
}

func (r ApiGetSTFStoreRequest) Execute() (models.STFStoreDetailModel, error) {
	bytes, err := r.ApiService.GetSTFStoreExecute(r)
	if err != nil {
		return models.STFStoreDetailModel{}, err
	}
	var reponse = models.STFStoreDetailModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.STFStoreDetailModel{}, fmt.Errorf("error unmarshal STFStoreDetailModel: %v", err)
	}
	return reponse, nil
}

func (a *STFStore) GetSTFStoreExecute(r ApiGetSTFStoreRequest) ([]byte, error) {
	var param = StructToString(r.getSTFStoreRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFStoreService", param)
}

func (a *STFStore) STFStoreGetSTFStore(ctx context.Context, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiGetSTFStoreRequest {
	return ApiGetSTFStoreRequest{
		ApiService:              a,
		ctx:                     ctx,
		getSTFStoreRequestModel: getSTFStoreRequestModel,
	}
}

// Set
type ApiSetSTFStoreRequest struct {
	ctx                     context.Context
	ApiService              *STFStore
	setSTFStoreRequestModel models.SetSTFStoreRequestModel
}

func (r ApiSetSTFStoreRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.SetSTFStoreExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFStore) SetSTFStoreExecute(r ApiSetSTFStoreRequest) ([]byte, error) {
	var param = StructToString(r.setSTFStoreRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFStoreService", param, "-Confirm:$false")
}

func (a *STFStore) STFStoreSetSTFStore(ctx context.Context, setSTFStoreRequestModel models.SetSTFStoreRequestModel) ApiSetSTFStoreRequest {
	return ApiSetSTFStoreRequest{
		ApiService:              a,
		ctx:                     ctx,
		setSTFStoreRequestModel: setSTFStoreRequestModel,
	}
}

// Clear
type ApiClearSTFStoreRequest struct {
	ctx                       context.Context
	ApiService                *STFStore
	clearSTFStoreRequestModel models.ClearSTFStoreRequestModel
}

func (r ApiClearSTFStoreRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.ClearSTFStoreExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFStore) ClearSTFStoreExecute(r ApiClearSTFStoreRequest) ([]byte, error) {
	var param = StructToString(r.clearSTFStoreRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Remove-STFStoreService", param, "-Confirm:$false")
}

func (a *STFStore) STFStoreClearSTFStore(ctx context.Context, clearSTFStoreRequestModel models.ClearSTFStoreRequestModel) ApiClearSTFStoreRequest {
	return ApiClearSTFStoreRequest{
		ApiService:                a,
		ctx:                       ctx,
		clearSTFStoreRequestModel: clearSTFStoreRequestModel,
	}
}

// Register-STFStoreGateway
type ApiRegisterSTFStoreGatewayRequest struct {
	ctx                                 context.Context
	ApiService                          *STFStore
	registerSTFStoreGatewayRequestModel models.RegisterSTFStoreGatewayRequestModel
}

func (r ApiRegisterSTFStoreGatewayRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.RegisterSTFStoreGatewayExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFStore) RegisterSTFStoreGatewayExecute(r ApiRegisterSTFStoreGatewayRequest) ([]byte, error) {
	var param = StructToString(r.registerSTFStoreGatewayRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Register-STFStoreGateway", param)
}

func (a *STFStore) STFStoreRegisterSTFStoreGateway(ctx context.Context, registerSTFStoreGatewayRequestModel models.RegisterSTFStoreGatewayRequestModel) ApiRegisterSTFStoreGatewayRequest {
	return ApiRegisterSTFStoreGatewayRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		registerSTFStoreGatewayRequestModel: registerSTFStoreGatewayRequestModel,
	}
}

// Unregister-STFStoreGateway
type ApiUnregisterSTFStoreGatewayRequest struct {
	ctx                                   context.Context
	ApiService                            *STFStore
	unregisterSTFStoreGatewayRequestModel models.UnregisterSTFStoreGatewayRequestModel
}

func (r ApiUnregisterSTFStoreGatewayRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.UnregisterSTFStoreGatewayExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFStore) UnregisterSTFStoreGatewayExecute(r ApiUnregisterSTFStoreGatewayRequest) ([]byte, error) {
	var param = StructToString(r.unregisterSTFStoreGatewayRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Unregister-STFStoreGateway", param)
}

func (a *STFStore) STFStoreUnregisterSTFStoreGateway(ctx context.Context, unregisterSTFStoreGatewayRequestModel models.UnregisterSTFStoreGatewayRequestModel) ApiUnregisterSTFStoreGatewayRequest {
	return ApiUnregisterSTFStoreGatewayRequest{
		ApiService:                            a,
		ctx:                                   ctx,
		unregisterSTFStoreGatewayRequestModel: unregisterSTFStoreGatewayRequestModel,
	}
}
