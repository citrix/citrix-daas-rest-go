// Copyright © 2024. Citrix Systems, Inc.
package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/citrix/citrix-daas-rest-go/citrixstorefront/models"
)

type STFDeployment Service

// Create
type ApiCreateSTFDeploymentRequest struct {
	ctx                             context.Context
	ApiService                      *STFDeployment
	createSTFDeploymentRequestModel models.CreateSTFDeploymentRequestModel
}

func (r ApiCreateSTFDeploymentRequest) Execute() (models.STFDeploymentDetailModel, error) {
	bytes, err := r.ApiService.CreateSTFDeploymentExecute(r)
	if err != nil {
		return models.STFDeploymentDetailModel{}, err
	}
	var reponse = models.STFDeploymentDetailModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return models.STFDeploymentDetailModel{}, unMarshalErr
	}
	return reponse, nil
}

func (a *STFDeployment) CreateSTFDeploymentExecute(r ApiCreateSTFDeploymentRequest) ([]byte, error) {
	var param = StructToString(r.createSTFDeploymentRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFDeployment", param, "-Confirm:$false")
}

func (a *STFDeployment) STFDeploymentCreateSTFDeployment(ctx context.Context, createSTFDeploymentRequestModel models.CreateSTFDeploymentRequestModel) ApiCreateSTFDeploymentRequest {
	return ApiCreateSTFDeploymentRequest{
		ApiService:                      a,
		ctx:                             ctx,
		createSTFDeploymentRequestModel: createSTFDeploymentRequestModel,
	}
}

// Get
type ApiGetSTFDeploymentRequest struct {
	ctx                          context.Context
	ApiService                   *STFDeployment
	getSTFDeploymentRequestModel models.GetSTFDeploymentRequestModel
}

func (r ApiGetSTFDeploymentRequest) Execute() (models.STFDeploymentDetailModel, error) {
	bytes, err := r.ApiService.GetSTFDeploymentExecute(r)
	if strings.EqualFold(string(bytes), "WARNING: A StoreFront deployment does not exist.\n") {
		return models.STFDeploymentDetailModel{}, fmt.Errorf(NOT_EXIST)
	}
	if err != nil {
		return models.STFDeploymentDetailModel{}, err
	}
	var reponse = models.STFDeploymentDetailModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return models.STFDeploymentDetailModel{}, unMarshalErr
	}
	return reponse, nil
}

func (a *STFDeployment) GetSTFDeploymentExecute(r ApiGetSTFDeploymentRequest) ([]byte, error) {
	var param = StructToString(r.getSTFDeploymentRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFDeployment", param)
}

func (a *STFDeployment) STFDeploymentGetSTFDeployment(ctx context.Context, getSTFDeploymentRequestModel models.GetSTFDeploymentRequestModel) ApiGetSTFDeploymentRequest {
	return ApiGetSTFDeploymentRequest{
		ApiService:                   a,
		ctx:                          ctx,
		getSTFDeploymentRequestModel: getSTFDeploymentRequestModel,
	}
}

// Set
type ApiSetSTFDeploymentRequest struct {
	ctx                          context.Context
	ApiService                   *STFDeployment
	setSTFDeploymentRequestModel models.SetSTFDeploymentRequestModel
}

func (r ApiSetSTFDeploymentRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.SetSTFDeploymentExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFDeployment) SetSTFDeploymentExecute(r ApiSetSTFDeploymentRequest) ([]byte, error) {
	var param = StructToString(r.setSTFDeploymentRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFDeployment", param, "-Confirm:$false")
}

func (a *STFDeployment) STFDeploymentSetSTFDeployment(ctx context.Context, setSTFDeploymentRequestModel models.SetSTFDeploymentRequestModel) ApiSetSTFDeploymentRequest {
	return ApiSetSTFDeploymentRequest{
		ApiService:                   a,
		ctx:                          ctx,
		setSTFDeploymentRequestModel: setSTFDeploymentRequestModel,
	}
}

// Clear
type ApiClearSTFDeploymentRequest struct {
	ctx                            context.Context
	ApiService                     *STFDeployment
	clearSTFDeploymentRequestModel models.ClearSTFDeploymentRequestModel
}

func (r ApiClearSTFDeploymentRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.ClearSTFDeploymentExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFDeployment) ClearSTFDeploymentExecute(r ApiClearSTFDeploymentRequest) ([]byte, error) {
	var param = StructToString(r.clearSTFDeploymentRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Clear-STFDeployment", param, "-Confirm:$false")
}

func (a *STFDeployment) STFDeploymentClearSTFDeployment(ctx context.Context, clearSTFDeploymentRequestModel models.ClearSTFDeploymentRequestModel) ApiClearSTFDeploymentRequest {
	return ApiClearSTFDeploymentRequest{
		ApiService:                     a,
		ctx:                            ctx,
		clearSTFDeploymentRequestModel: clearSTFDeploymentRequestModel,
	}
}
