// Copyright © 2024. Citrix Systems, Inc.
package apis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/citrix/citrix-daas-rest-go/citrixstorefront/models"
)

type STFRoaming Service

// Add-STFRoamingGateway
type ApiAddSTFRoamingGatewayRequest struct {
	ctx                              context.Context
	ApiService                       *STFRoaming
	AddSTFRoamingGatewayRequestModel models.AddSTFRoamingGatewayRequestModel
}

func (r ApiAddSTFRoamingGatewayRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.AddSTFRoamingGatewayExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFRoaming) AddSTFRoamingGatewayExecute(r ApiAddSTFRoamingGatewayRequest) ([]byte, error) {
	var param = StructToString(r.AddSTFRoamingGatewayRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFRoamingGateway", param)
}

func (a *STFRoaming) STFRoamingAddGateway(ctx context.Context, createSTFRoamingGatewayRequestModel models.AddSTFRoamingGatewayRequestModel) ApiAddSTFRoamingGatewayRequest {
	return ApiAddSTFRoamingGatewayRequest{
		ApiService:                       a,
		ctx:                              ctx,
		AddSTFRoamingGatewayRequestModel: createSTFRoamingGatewayRequestModel,
	}
}

// New-STFRoamingGateway
type ApiNewSTFRoamingGatewayRequest struct {
	ctx                              context.Context
	ApiService                       *STFRoaming
	NewSTFRoamingGatewayRequestModel models.NewSTFRoamingGatewayRequestModel
}

func (r ApiNewSTFRoamingGatewayRequest) Execute() (models.STFRoamingDetailModel, error) {
	bytes, err := r.ApiService.NewSTFRoamingGatewayExecute(r)
	if err != nil {
		return models.STFRoamingDetailModel{}, err
	}
	var reponse = models.STFRoamingDetailModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.STFRoamingDetailModel{}, fmt.Errorf("error unmarshal STFRoamingDetailModel: %v", err)
	}
	return reponse, nil
}

func (a *STFRoaming) NewSTFRoamingGatewayExecute(r ApiNewSTFRoamingGatewayRequest) ([]byte, error) {
	var param = StructToString(r.NewSTFRoamingGatewayRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "New-STFRoamingGateway", param)
}

func (a *STFRoaming) STFRoamingNewGateway(ctx context.Context, newSTFRoamingGatewayRequestModel models.NewSTFRoamingGatewayRequestModel) ApiNewSTFRoamingGatewayRequest {
	return ApiNewSTFRoamingGatewayRequest{
		ApiService:                       a,
		ctx:                              ctx,
		NewSTFRoamingGatewayRequestModel: newSTFRoamingGatewayRequestModel,
	}
}

// Get-STFRoamingGateway
type ApiGetSTFRoamingGatewayRequest struct {
	ctx                              context.Context
	ApiService                       *STFRoaming
	GetSTFRoamingGatewayRequestModel models.GetSTFRoamingGatewayRequestModel
}

func (r ApiGetSTFRoamingGatewayRequest) Execute() (models.STFRoamingDetailModel, error) {
	bytes, err := r.ApiService.GetSTFRoamingGatewayExecute(r)
	if err != nil {
		return models.STFRoamingDetailModel{}, err
	}
	var reponse = models.STFRoamingDetailModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.STFRoamingDetailModel{}, fmt.Errorf("error unmarshal STFRoamingDetailModel: %v", err)
	}
	return reponse, nil
}

func (a *STFRoaming) GetSTFRoamingGatewayExecute(r ApiGetSTFRoamingGatewayRequest) ([]byte, error) {
	var param = StructToString(r.GetSTFRoamingGatewayRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFRoamingGateway", param)
}

func (a *STFRoaming) STFRoamingGetGateway(ctx context.Context, getSTFRoamingGatewayRequestModel models.GetSTFRoamingGatewayRequestModel) ApiGetSTFRoamingGatewayRequest {
	return ApiGetSTFRoamingGatewayRequest{
		ApiService:                       a,
		ctx:                              ctx,
		GetSTFRoamingGatewayRequestModel: getSTFRoamingGatewayRequestModel,
	}
}

// Remove-STFRoamingGateway
type ApiRemoveSTFoamingGatewayRequest struct {
	ctx                                 context.Context
	ApiService                          *STFRoaming
	removeSTFRoamingGatewayRequestModel models.RemoveSTFRoamingGatewayRequestModel
}

func (r ApiRemoveSTFoamingGatewayRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.RemoveSTFoamingGatewayExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFRoaming) RemoveSTFoamingGatewayExecute(r ApiRemoveSTFoamingGatewayRequest) ([]byte, error) {
	var param = StructToString(r.removeSTFRoamingGatewayRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Remove-STFRoamingGateway", param)
}

func (a *STFRoaming) STFStoreRemoveGateway(ctx context.Context, removeSTFoamingGatewayRequestModel models.RemoveSTFRoamingGatewayRequestModel) ApiRemoveSTFoamingGatewayRequest {
	return ApiRemoveSTFoamingGatewayRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		removeSTFRoamingGatewayRequestModel: removeSTFoamingGatewayRequestModel,
	}
}
