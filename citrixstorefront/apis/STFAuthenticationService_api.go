// Copyright © 2024. Citrix Systems, Inc.
package apis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/citrix/citrix-daas-rest-go/citrixstorefront/models"
)

type STFAuthentication Service

// Create
type ApiAddSTFAuthenticationServiceRequest struct {
	ctx                                     context.Context
	ApiService                              *STFAuthentication
	AddSTFAuthenticationServiceRequestModel models.AddSTFAuthenticationServiceRequestModel
}

func (r ApiAddSTFAuthenticationServiceRequest) Execute() (models.STFAuthenticationServiceResponseModel, error) {
	bytes, err := r.ApiService.ApiAddSTFAuthenticationServiceExecute(r)
	if err != nil {
		return models.STFAuthenticationServiceResponseModel{}, err
	}
	var reponse = models.STFAuthenticationServiceRawResponseModel{}
	unmarshalErr := json.Unmarshal(bytes, &reponse)
	if unmarshalErr != nil {
		fmt.Println("Error: ", unmarshalErr)
		return models.STFAuthenticationServiceResponseModel{}, fmt.Errorf("error unmarshal STFAuthenticationServiceResponseModel: %v", unmarshalErr.Error())
	}

	return reponse.ConvertToResponseModel()
}

func (a *STFAuthentication) ApiAddSTFAuthenticationServiceExecute(r ApiAddSTFAuthenticationServiceRequest) ([]byte, error) {
	var param = StructToString(r.AddSTFAuthenticationServiceRequestModel)
	return ExecuteCommandWithDepth(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), 5, "Add-STFAuthenticationService", param)
}

func (a *STFAuthentication) STFAuthenticationCreateSTFAuthenticationService(ctx context.Context, addSTFAuthenticationServiceRequestModel models.AddSTFAuthenticationServiceRequestModel) ApiAddSTFAuthenticationServiceRequest {
	return ApiAddSTFAuthenticationServiceRequest{
		ApiService:                              a,
		ctx:                                     ctx,
		AddSTFAuthenticationServiceRequestModel: addSTFAuthenticationServiceRequestModel,
	}
}

// Read
type ApiGetSTFAuthenticationServiceRequest struct {
	ctx                                     context.Context
	ApiService                              *STFAuthentication
	GetSTFAuthenticationServiceRequestModel models.GetSTFAuthenticationServiceRequestModel
}

func (r ApiGetSTFAuthenticationServiceRequest) Execute() (models.STFAuthenticationServiceResponseModel, error) {
	bytes, err := r.ApiService.GetSTFAuthenticationServiceExecute(r)
	if err != nil {
		return models.STFAuthenticationServiceResponseModel{}, err
	}
	if len(bytes) == 0 {
		return models.STFAuthenticationServiceResponseModel{}, fmt.Errorf(NOT_EXIST)
	}
	var reponse = models.STFAuthenticationServiceRawResponseModel{}
	unmarshalErr := json.Unmarshal(bytes, &reponse)
	if unmarshalErr != nil {
		fmt.Println("Error: ", unmarshalErr)
		return models.STFAuthenticationServiceResponseModel{}, fmt.Errorf("error unmarshal STFAuthenticationServiceResponseModel: %v", unmarshalErr.Error())
	}
	return reponse.ConvertToResponseModel()
}

func (a *STFAuthentication) GetSTFAuthenticationServiceExecute(r ApiGetSTFAuthenticationServiceRequest) ([]byte, error) {
	var param = StructToString(r.GetSTFAuthenticationServiceRequestModel)
	return ExecuteCommandWithDepth(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), 5, "Get-STFAuthenticationService", param)
}

func (a *STFAuthentication) STFAuthenticationGetSTFAuthenticationService(ctx context.Context, getSTFAuthenticationServiceRequestModel models.GetSTFAuthenticationServiceRequestModel) ApiGetSTFAuthenticationServiceRequest {
	return ApiGetSTFAuthenticationServiceRequest{
		ApiService:                              a,
		ctx:                                     ctx,
		GetSTFAuthenticationServiceRequestModel: getSTFAuthenticationServiceRequestModel,
	}
}

// Delete
type ApiRemoveSTFAuthenticationServiceRequest struct {
	ctx                                     context.Context
	ApiService                              *STFAuthentication
	GetSTFAuthenticationServiceRequestModel models.GetSTFAuthenticationServiceRequestModel
}

func (r ApiRemoveSTFAuthenticationServiceRequest) Execute() error {
	_, err := r.ApiService.RemoveSTFAuthenticationServiceExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFAuthentication) RemoveSTFAuthenticationServiceExecute(r ApiRemoveSTFAuthenticationServiceRequest) ([]byte, error) {
	var param = StructToString(r.GetSTFAuthenticationServiceRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Remove-STFAuthenticationService", fmt.Sprintf("(Get-STFAuthenticationService %s)", param), "-Confirm:$false")
}

func (a *STFAuthentication) STFAuthenticationRemoveSTFAuthenticationService(ctx context.Context, getSTFAuthenticationServiceRequestModel models.GetSTFAuthenticationServiceRequestModel) ApiRemoveSTFAuthenticationServiceRequest {
	return ApiRemoveSTFAuthenticationServiceRequest{
		ApiService:                              a,
		ctx:                                     ctx,
		GetSTFAuthenticationServiceRequestModel: getSTFAuthenticationServiceRequestModel,
	}
}

// Get STFAuthenticationProtocolsAvailable
type ApiGetSTFAuthenticationProtocolsAvailableRequest struct {
	ctx        context.Context
	ApiService *STFAuthentication
}

func (r ApiGetSTFAuthenticationProtocolsAvailableRequest) Execute() ([]string, error) {
	bytes, err := r.ApiService.GetSTFAuthenticationProtocolsAvailableExecute(r)
	if err != nil {
		return nil, err
	}
	var reponse = []string{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return nil, unMarshalErr
	}
	return reponse, nil
}

func (a *STFAuthentication) GetSTFAuthenticationProtocolsAvailableExecute(r ApiGetSTFAuthenticationProtocolsAvailableRequest) ([]byte, error) {
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Get-STFAuthenticationProtocolsAvailable")
}

func (a *STFAuthentication) STFWebReceiverGetSTFAuthenticationProtocolsAvailable(ctx context.Context) ApiGetSTFAuthenticationProtocolsAvailableRequest {
	return ApiGetSTFAuthenticationProtocolsAvailableRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Get STFAuthenticationProtocols
type ApiGetSTFAuthenticationProtocolsRequest struct {
	ctx                                             context.Context
	ApiService                                      *STFAuthentication
	GetSTFAuthenticationServiceProtocolRequestModel models.GetSTFAuthenticationServiceProtocolRequestModel
}

func (r ApiGetSTFAuthenticationProtocolsRequest) Execute() ([]models.STFAuthenticationServiceProtocolResponseModel, error) {
	bytes, err := r.ApiService.GetSTFAuthenticationProtocolsExecute(r)
	if err != nil {
		return nil, err
	}
	var reponse = []models.STFAuthenticationServiceProtocolResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return nil, unMarshalErr
	}
	return reponse, nil
}

func (a *STFAuthentication) GetSTFAuthenticationProtocolsExecute(r ApiGetSTFAuthenticationProtocolsRequest) ([]byte, error) {
	var param = StructToString(r.GetSTFAuthenticationServiceProtocolRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Get-STFAuthenticationServiceProtocol", param)
}

func (a *STFAuthentication) STFWebReceiverGetSTFAuthenticationProtocols(ctx context.Context, getSTFAuthenticationServiceProtocolRequestModel models.GetSTFAuthenticationServiceProtocolRequestModel) ApiGetSTFAuthenticationProtocolsRequest {
	return ApiGetSTFAuthenticationProtocolsRequest{
		ApiService: a,
		ctx:        ctx,
		GetSTFAuthenticationServiceProtocolRequestModel: getSTFAuthenticationServiceProtocolRequestModel,
	}
}

// Add STFAuthenticationProtocols
type ApiAddSTFAuthenticationProtocolsRequest struct {
	ctx                                             context.Context
	ApiService                                      *STFAuthentication
	AddSTFAuthenticationServiceProtocolRequestModel models.AddUpdateSTFAuthenticationServiceProtocolRequestModel
}

func (r ApiAddSTFAuthenticationProtocolsRequest) Execute() error {
	_, err := r.ApiService.AddSTFAuthenticationProtocolsExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFAuthentication) AddSTFAuthenticationProtocolsExecute(r ApiAddSTFAuthenticationProtocolsRequest) ([]byte, error) {
	var param = StructToString(r.AddSTFAuthenticationServiceProtocolRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Add-STFAuthenticationServiceProtocol", param)
}

func (a *STFAuthentication) STFWebReceiverAddSTFAuthenticationProtocols(ctx context.Context, addSTFAuthenticationServiceProtocolRequestModel models.AddUpdateSTFAuthenticationServiceProtocolRequestModel) ApiAddSTFAuthenticationProtocolsRequest {
	return ApiAddSTFAuthenticationProtocolsRequest{
		ApiService: a,
		ctx:        ctx,
		AddSTFAuthenticationServiceProtocolRequestModel: addSTFAuthenticationServiceProtocolRequestModel,
	}
}

// Enable STFAuthenticationProtocols
type ApiEnableSTFAuthenticationProtocolsRequest struct {
	ctx                                             context.Context
	ApiService                                      *STFAuthentication
	EnableAuthenticationServiceProtocolRequestModel models.AddUpdateSTFAuthenticationServiceProtocolRequestModel
}

func (r ApiEnableSTFAuthenticationProtocolsRequest) Execute() error {
	_, err := r.ApiService.EnableSTFAuthenticationProtocolsExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFAuthentication) EnableSTFAuthenticationProtocolsExecute(r ApiEnableSTFAuthenticationProtocolsRequest) ([]byte, error) {
	var param = StructToString(r.EnableAuthenticationServiceProtocolRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Enable-STFAuthenticationServiceProtocol", param)
}

func (a *STFAuthentication) STFWebReceiverEnableSTFAuthenticationProtocols(ctx context.Context, EnableAuthenticationServiceProtocolRequestModel models.AddUpdateSTFAuthenticationServiceProtocolRequestModel) ApiEnableSTFAuthenticationProtocolsRequest {
	return ApiEnableSTFAuthenticationProtocolsRequest{
		ApiService: a,
		ctx:        ctx,
		EnableAuthenticationServiceProtocolRequestModel: EnableAuthenticationServiceProtocolRequestModel,
	}
}

// Disable STFAuthenticationProtocols
type ApiDisableSTFAuthenticationProtocolsRequest struct {
	ctx                                              context.Context
	ApiService                                       *STFAuthentication
	DisableAuthenticationServiceProtocolRequestModel models.AddUpdateSTFAuthenticationServiceProtocolRequestModel
}

func (r ApiDisableSTFAuthenticationProtocolsRequest) Execute() error {
	_, err := r.ApiService.DisableSTFAuthenticationProtocolsExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFAuthentication) DisableSTFAuthenticationProtocolsExecute(r ApiDisableSTFAuthenticationProtocolsRequest) ([]byte, error) {
	var param = StructToString(r.DisableAuthenticationServiceProtocolRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Disable-STFAuthenticationServiceProtocol", param)
}

func (a *STFAuthentication) STFWebReceiverDisableSTFAuthenticationProtocols(ctx context.Context, DisableAuthenticationServiceProtocolRequestModel models.AddUpdateSTFAuthenticationServiceProtocolRequestModel) ApiDisableSTFAuthenticationProtocolsRequest {
	return ApiDisableSTFAuthenticationProtocolsRequest{
		ApiService: a,
		ctx:        ctx,
		DisableAuthenticationServiceProtocolRequestModel: DisableAuthenticationServiceProtocolRequestModel,
	}
}

// Set STFClaimsFactoryNames

type ApiSetSTFClaimsFactoryNamesRequest struct {
	ctx                                  context.Context
	ApiService                           *STFAuthentication
	GetAuthenticationServiceRequestModel models.GetSTFAuthenticationServiceRequestModel
	SetSTFClaimsFactoryNamesRequestModel models.SetSTFClaimsFactoryNamesRequestModel
}

func (r ApiSetSTFClaimsFactoryNamesRequest) Execute() error {
	_, err := r.ApiService.SetSTFClaimsFactoryNamesExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFAuthentication) SetSTFClaimsFactoryNamesExecute(r ApiSetSTFClaimsFactoryNamesRequest) ([]byte, error) {
	var getAuthenticationServiceParams = StructToString(r.GetAuthenticationServiceRequestModel)
	var setClaimsFactoryNameParam = StructToString(r.SetSTFClaimsFactoryNamesRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Set-STFClaimsFactoryNames", fmt.Sprintf("-AuthenticationService (Get-STFAuthenticationService %s)", getAuthenticationServiceParams), setClaimsFactoryNameParam)
}

func (a *STFAuthentication) STFSetClaimsFactoryNames(ctx context.Context, getAuthenticationServiceRequestModel models.GetSTFAuthenticationServiceRequestModel, setSTFClaimsFactoryNamesRequestModel models.SetSTFClaimsFactoryNamesRequestModel) ApiSetSTFClaimsFactoryNamesRequest {
	return ApiSetSTFClaimsFactoryNamesRequest{
		ctx:                                  ctx,
		ApiService:                           a,
		GetAuthenticationServiceRequestModel: getAuthenticationServiceRequestModel,
		SetSTFClaimsFactoryNamesRequestModel: setSTFClaimsFactoryNamesRequestModel,
	}
}

// Set-STFCitrixAGBasicOptions
type ApiSetSTFCitrixAGBasicOptionsRequest struct {
	ctx                                    context.Context
	ApiService                             *STFAuthentication
	getSTFAuthenticationService            models.GetSTFAuthenticationServiceRequestModel
	SetSTFCitrixAGBasicOptionsRequestModel models.SetSTFCitrixAGBasicOptionsRequestModel
}

func (r ApiSetSTFCitrixAGBasicOptionsRequest) Execute() error {
	_, err := r.ApiService.SetSTFCitrixAGBasicOptionsExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFAuthentication) SetSTFCitrixAGBasicOptionsExecute(r ApiSetSTFCitrixAGBasicOptionsRequest) ([]byte, error) {
	var param = StructToString(r.SetSTFCitrixAGBasicOptionsRequestModel)
	var authParam = StructToString(r.getSTFAuthenticationService)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Set-STFCitrixAGBasicOptions", fmt.Sprintf("-AuthenticationService (Get-STFAuthenticationService %s)", authParam), param)
}

func (a *STFAuthentication) STFSetCitrixAGBasicOptions(ctx context.Context, getSTFAuthenticationService models.GetSTFAuthenticationServiceRequestModel, setSTFCitrixAGBasicOptionsRequestModel models.SetSTFCitrixAGBasicOptionsRequestModel) ApiSetSTFCitrixAGBasicOptionsRequest {
	return ApiSetSTFCitrixAGBasicOptionsRequest{
		ctx:                                    ctx,
		ApiService:                             a,
		getSTFAuthenticationService:            getSTFAuthenticationService,
		SetSTFCitrixAGBasicOptionsRequestModel: setSTFCitrixAGBasicOptionsRequestModel,
	}
}

// Get-STFCitrixAGBasicOptions
type ApiGetSTFCitrixAGBasicOptionsRequest struct {
	ctx                         context.Context
	ApiService                  *STFAuthentication
	getSTFAuthenticationService models.GetSTFAuthenticationServiceRequestModel
}

func (r ApiGetSTFCitrixAGBasicOptionsRequest) Execute() (models.STFCitrixAGBasicOptionsResponseModel, error) {

	bytes, err := r.ApiService.GetSTFCitrixAGBasicOptionsExecute(r)
	if err != nil {
		return models.STFCitrixAGBasicOptionsResponseModel{}, err
	}
	if len(bytes) == 0 {
		return models.STFCitrixAGBasicOptionsResponseModel{}, fmt.Errorf(NOT_EXIST)
	}
	var reponse = models.STFCitrixAGBasicOptionsResponseModel{}
	unmarshalErr := json.Unmarshal(bytes, &reponse)
	if unmarshalErr != nil {
		fmt.Println("Error: ", unmarshalErr)
		return models.STFCitrixAGBasicOptionsResponseModel{}, fmt.Errorf("error unmarshal STFCitrixAGBasicOptionsResponseModel: %v", unmarshalErr.Error())
	}
	return reponse, nil
}

func (a *STFAuthentication) GetSTFCitrixAGBasicOptionsExecute(r ApiGetSTFCitrixAGBasicOptionsRequest) ([]byte, error) {
	var param = StructToString(r.getSTFAuthenticationService)
	return ExecuteCommandWithDepth(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), 5, "Get-STFCitrixAGBasicOptions ", fmt.Sprintf("-AuthenticationService (Get-STFAuthenticationService %s)", param))
}

func (a *STFAuthentication) STFGetCitrixAGBasicOptions(ctx context.Context, getSTFAuthenticationService models.GetSTFAuthenticationServiceRequestModel) ApiGetSTFCitrixAGBasicOptionsRequest {
	return ApiGetSTFCitrixAGBasicOptionsRequest{
		ctx:                         ctx,
		ApiService:                  a,
		getSTFAuthenticationService: getSTFAuthenticationService,
	}
}
