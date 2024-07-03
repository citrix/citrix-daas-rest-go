// Copyright © 2024. Citrix Systems, Inc.
package apis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/citrix/citrix-daas-rest-go/citrixstorefront/models"
)

type STFRoaming Service

// Get-STFRoamingService
type ApiGetSTFRoamingServiceRequest struct {
	ctx                           context.Context
	ApiService                    *STFRoaming
	STFRoamingServiceRequestModel models.STFRoamingServiceRequestModel
}

func (r ApiGetSTFRoamingServiceRequest) Execute() (models.STFRoamingServiceResponseModel, error) {
	bytes, err := r.ApiService.GetSTFRoamingServiceExecute(r)
	if err != nil {
		return models.STFRoamingServiceResponseModel{}, err
	}
	var reponse = models.STFRoamingServiceResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.STFRoamingServiceResponseModel{}, fmt.Errorf("error unmarshal STFRoamingServiceResponseModel: %v", unMarshalErr.Error())
	}
	return reponse, nil
}

func (a *STFRoaming) GetSTFRoamingServiceExecute(r ApiGetSTFRoamingServiceRequest) ([]byte, error) {
	params := StructToString(r.STFRoamingServiceRequestModel)
	return ExecuteCommandWithDepth(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), 5, "Get-STFRoamingService", params)
}

func (a *STFRoaming) STFRoamingServiceGet(ctx context.Context, stfRoamingServiceRequestModel models.STFRoamingServiceRequestModel) ApiGetSTFRoamingServiceRequest {
	return ApiGetSTFRoamingServiceRequest{
		ApiService:                    a,
		ctx:                           ctx,
		STFRoamingServiceRequestModel: stfRoamingServiceRequestModel,
	}
}

// Add-STFRoamingGateway
type ApiAddSTFRoamingGatewayRequest struct {
	ctx                              context.Context
	ApiService                       *STFRoaming
	AddSTFRoamingGatewayRequestModel models.AddSTFRoamingGatewayRequestModel
	GetSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel
	STFStaUrls                       []models.STFSTAUrlModel
	SessionReliability               bool
	RequestTicketTwoSTAs             bool
	StasUseLoadBalancings            bool
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
	var getRoamingServiceParam = StructToString(r.GetSTFRoamingServiceRequestModel)
	staUrlParam := ""
	sessionReliabilityParam := ""
	if r.SessionReliability {
		sessionReliabilityParam = "-SessionReliability"
	}

	requestTicketTwoSTAsParam := ""
	if r.RequestTicketTwoSTAs {
		requestTicketTwoSTAsParam = "-RequestTicketTwoSTAs"
	}

	stasUseLoadBalancingParam := ""
	if r.StasUseLoadBalancings {
		stasUseLoadBalancingParam = "-StasUseLoadBalancing"
	}
	if len(r.STFStaUrls) > 0 {
		staUrlParam = " -SecureTicketAuthorityObjs @("
		for _, staUrl := range r.STFStaUrls {
			staUrlParam += fmt.Sprintf("(New-STFSecureTicketAuthority -StaUrl '%s' -StaValidationEnabled $%t -StaValidationSecret '%s');", *staUrl.AuthorityId.Get(), *staUrl.StaValidationEnabled.Get(), *staUrl.StaValidationSecret.Get())
		}
		staUrlParam += ")"
	}

	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFRoamingGateway", fmt.Sprintf("-RoamingService (Get-STFRoamingService %s)", getRoamingServiceParam), param, staUrlParam, sessionReliabilityParam, requestTicketTwoSTAsParam, stasUseLoadBalancingParam)
}

func (a *STFRoaming) STFRoamingGatewayAdd(ctx context.Context, createSTFRoamingGatewayRequestModel models.AddSTFRoamingGatewayRequestModel, getSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel, stfStaUrls []models.STFSTAUrlModel, sessionReliability bool, requestTicketTwoSTAs bool, stasUseLoadBalancing bool) ApiAddSTFRoamingGatewayRequest {
	return ApiAddSTFRoamingGatewayRequest{
		ApiService:                       a,
		ctx:                              ctx,
		AddSTFRoamingGatewayRequestModel: createSTFRoamingGatewayRequestModel,
		GetSTFRoamingServiceRequestModel: getSTFRoamingServiceRequestModel,
		STFStaUrls:                       stfStaUrls,
	}
}

// Get-STFRoamingGateway
type ApiGetSTFRoamingGatewayRequest struct {
	ctx                              context.Context
	ApiService                       *STFRoaming
	GetSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel
	GetSTFRoamingGatewayRequestModel models.GetSTFRoamingGatewayRequestModel
}

func (r ApiGetSTFRoamingGatewayRequest) Execute() (models.STFRoamingGatewayResponseModel, error) {
	bytes, err := r.ApiService.GetSTFRoamingGatewayExecute(r)
	if err != nil {
		return models.STFRoamingGatewayResponseModel{}, err
	}
	var reponse = models.STFRoamingGatewayRawResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.STFRoamingGatewayResponseModel{}, fmt.Errorf("error unmarshal STFRoamingGatewayDetailModel: %v", unMarshalErr.Error())
	}
	reponse.SiteId = *r.GetSTFRoamingServiceRequestModel.SiteId.Get()
	return reponse.ConvertToResponseModel(), nil
}

func (a *STFRoaming) GetSTFRoamingGatewayExecute(r ApiGetSTFRoamingGatewayRequest) ([]byte, error) {
	var param = StructToString(r.GetSTFRoamingGatewayRequestModel)
	var getRoamingServiceParams = StructToString(r.GetSTFRoamingServiceRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFRoamingGateway", fmt.Sprintf("-RoamingService (Get-STFRoamingService %s)", getRoamingServiceParams), param)
}

func (a *STFRoaming) STFRoamingGatewayGet(ctx context.Context, getSTFRoamingGatewayRequestModel models.GetSTFRoamingGatewayRequestModel, getSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel) ApiGetSTFRoamingGatewayRequest {
	return ApiGetSTFRoamingGatewayRequest{
		ApiService:                       a,
		ctx:                              ctx,
		GetSTFRoamingServiceRequestModel: getSTFRoamingServiceRequestModel,
		GetSTFRoamingGatewayRequestModel: getSTFRoamingGatewayRequestModel,
	}
}

// Set-STFRoamingGateway
type ApiSetSTFRoamingGatewayRequest struct {
	ctx                              context.Context
	ApiService                       *STFRoaming
	SetSTFRoamingGatewayRequestModel models.SetSTFRoamingGatewayRequestModel
	GetSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel
	STFStaUrls                       []models.STFSTAUrlModel
}

func (r ApiSetSTFRoamingGatewayRequest) Execute() error {
	_, err := r.ApiService.SetSTFRoamingGatewayExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFRoaming) SetSTFRoamingGatewayExecute(r ApiSetSTFRoamingGatewayRequest) ([]byte, error) {
	var param = StructToString(r.SetSTFRoamingGatewayRequestModel)
	var getRoamingServiceParam = StructToString(r.GetSTFRoamingServiceRequestModel)
	staUrlParam := ""
	if len(r.STFStaUrls) > 0 {
		staUrlParam = " -SecureTicketAuthorityObjs @("
		for _, staUrl := range r.STFStaUrls {
			staUrlParam += fmt.Sprintf("(New-STFSecureTicketAuthority -StaUrl '%s' -StaValidationEnabled $%t -StaValidationSecret '%s');", *staUrl.AuthorityId.Get(), *staUrl.StaValidationEnabled.Get(), *staUrl.StaValidationSecret.Get())
		}
		staUrlParam += ")"
	}

	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFRoamingGateway", fmt.Sprintf("-RoamingService (Get-STFRoamingService %s)", getRoamingServiceParam), param, staUrlParam)
}

func (a *STFRoaming) STFRoamingGatewaySet(ctx context.Context, setSTFRoamingGatewayRequestModel models.SetSTFRoamingGatewayRequestModel, getSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel, stfStaUrls []models.STFSTAUrlModel) ApiSetSTFRoamingGatewayRequest {
	return ApiSetSTFRoamingGatewayRequest{
		ApiService:                       a,
		ctx:                              ctx,
		SetSTFRoamingGatewayRequestModel: setSTFRoamingGatewayRequestModel,
		GetSTFRoamingServiceRequestModel: getSTFRoamingServiceRequestModel,
		STFStaUrls:                       stfStaUrls,
	}
}

// Remove-STFRoamingGateway
type ApiRemoveSTFoamingGatewayRequest struct {
	ctx                              context.Context
	ApiService                       *STFRoaming
	GetSTFRoamingGatewayRequestModel models.GetSTFRoamingGatewayRequestModel
	GetSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel
}

func (r ApiRemoveSTFoamingGatewayRequest) Execute() error {
	_, err := r.ApiService.RemoveSTFRoamingGatewayExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFRoaming) RemoveSTFRoamingGatewayExecute(r ApiRemoveSTFoamingGatewayRequest) ([]byte, error) {
	var getRoamingServiceParams = StructToString(r.GetSTFRoamingServiceRequestModel)
	var param = StructToString(r.GetSTFRoamingGatewayRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Remove-STFRoamingGateway", fmt.Sprintf("-RoamingService (Get-STFRoamingService %s)", getRoamingServiceParams), param, "-Confirm:$false")
}

func (a *STFRoaming) STFRoamingGatewayRemove(ctx context.Context, getSTFoamingGatewayRequestModel models.GetSTFRoamingGatewayRequestModel, getSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel) ApiRemoveSTFoamingGatewayRequest {
	return ApiRemoveSTFoamingGatewayRequest{
		ApiService:                       a,
		ctx:                              ctx,
		GetSTFRoamingGatewayRequestModel: getSTFoamingGatewayRequestModel,
		GetSTFRoamingServiceRequestModel: getSTFRoamingServiceRequestModel,
	}
}
