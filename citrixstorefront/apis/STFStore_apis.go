// Copyright © 2024. Citrix Systems, Inc.
package apis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/citrix/citrix-daas-rest-go/citrixstorefront/models"
)

type STFStore Service

// Add-STFStoreService
type ApiCreateSTFStoreRequest struct {
	ctx                                  context.Context
	ApiService                           *STFStore
	CreateSTFStoreRequestModel           models.CreateSTFStoreRequestModel
	GetAuthenticationServiceRequestModel models.GetSTFAuthenticationServiceRequestModel
}

func (r ApiCreateSTFStoreRequest) Execute() (models.STFStoreDetailModel, error) {
	bytes, err := r.ApiService.CreateSTFStoreExecute(r)
	if err != nil {
		return models.STFStoreDetailModel{}, err
	}
	var response = models.STFStoreDetailModel{}
	unMarshalErr := json.Unmarshal(bytes, &response)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.STFStoreDetailModel{}, fmt.Errorf("error unmarshal STFStoreDetailModel: %v", unMarshalErr.Error())
	}
	return response, nil
}

func (a *STFStore) CreateSTFStoreExecute(r ApiCreateSTFStoreRequest) ([]byte, error) {
	var param = StructToString(r.CreateSTFStoreRequestModel)
	var getAuthServiceParams = StructToString(r.GetAuthenticationServiceRequestModel)
	if r.GetAuthenticationServiceRequestModel.VirtualPath.IsSet() && *r.GetAuthenticationServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFStoreService", fmt.Sprintf("-AuthenticationService (Get-STFAuthenticationService %s)", getAuthServiceParams), param)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFStoreService", param)
	}
}

func (a *STFStore) STFStoreCreateSTFStore(ctx context.Context, createSTFStoreRequestModel models.CreateSTFStoreRequestModel, getAuthenticationServiceRequestModel models.GetSTFAuthenticationServiceRequestModel) ApiCreateSTFStoreRequest {
	return ApiCreateSTFStoreRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		CreateSTFStoreRequestModel:           createSTFStoreRequestModel,
		GetAuthenticationServiceRequestModel: getAuthenticationServiceRequestModel,
	}
}

// Get-STFStoreService
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
	if len(bytes) == 0 {
		return models.STFStoreDetailModel{}, fmt.Errorf(NOT_EXIST)
	}
	var response = models.STFStoreDetailModel{}
	unMarshalErr := json.Unmarshal(bytes, &response)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.STFStoreDetailModel{}, fmt.Errorf("error occurred while unmarshalling STFStoreDetailModel: %v", unMarshalErr.Error())
	}
	return response, nil
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

// Set-STFStoreService
type ApiSetSTFStoreRequest struct {
	ctx                     context.Context
	ApiService              *STFStore
	GetSTFStoreRequestModel models.GetSTFStoreRequestModel
}

func (r ApiSetSTFStoreRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.SetSTFStoreExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFStore) SetSTFStoreExecute(r ApiSetSTFStoreRequest) ([]byte, error) {
	var getParams = StructToString(r.GetSTFStoreRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFStoreService", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", getParams), "-Confirm:$false")
}

func (a *STFStore) STFStoreSetSTFStore(ctx context.Context, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiSetSTFStoreRequest {
	return ApiSetSTFStoreRequest{
		ApiService:              a,
		ctx:                     ctx,
		GetSTFStoreRequestModel: getSTFStoreRequestModel,
	}
}

// Remove-STFStoreService
type ApiClearSTFStoreRequest struct {
	ctx                     context.Context
	ApiService              *STFStore
	GetSTFStoreRequestModel models.GetSTFStoreRequestModel
}

func (r ApiClearSTFStoreRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.ClearSTFStoreExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFStore) ClearSTFStoreExecute(r ApiClearSTFStoreRequest) ([]byte, error) {
	var getParams = StructToString(r.GetSTFStoreRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Remove-STFStoreService", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", getParams), "-Confirm:$false")
}

func (a *STFStore) STFStoreClearSTFStore(ctx context.Context, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiClearSTFStoreRequest {
	return ApiClearSTFStoreRequest{
		ApiService:              a,
		ctx:                     ctx,
		GetSTFStoreRequestModel: getSTFStoreRequestModel,
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

// Add-STFStoreFarm
type ApiAddSTFStoreFarmRequest struct {
	ctx                            context.Context
	ApiService                     *STFStore
	addSTFStoreFarmRequestModel    models.AddSTFStoreFarmRequestModel
	getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel //This is used to set the StoreService for the SetSTFStoreFarmRequest
}

func (r ApiAddSTFStoreFarmRequest) Execute() (models.StoreFarmModel, error) {
	_, err := r.ApiService.AddSTFStoreFarmExecute(r)
	if err != nil {
		return models.StoreFarmModel{}, err
	}
	return models.StoreFarmModel{}, nil
}

func (a *STFStore) AddSTFStoreFarmExecute(r ApiAddSTFStoreFarmRequest) ([]byte, error) {
	var param = StructToString(r.addSTFStoreFarmRequestModel.BaseSTFStoreFarmRequestModel)
	var addStoreServiceParams = StructToString(r.getSTFStoreServiceRequestModel)

	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFStoreFarm", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", addStoreServiceParams), param)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFStoreFarm", param)
	}
}

func (a *STFStore) STFStoreNewStoreFarm(ctx context.Context, addSTFStoreFarmRequestModel models.AddSTFStoreFarmRequestModel, getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel) ApiAddSTFStoreFarmRequest {
	return ApiAddSTFStoreFarmRequest{
		ApiService:                     a,
		ctx:                            ctx,
		addSTFStoreFarmRequestModel:    addSTFStoreFarmRequestModel,
		getSTFStoreServiceRequestModel: getSTFStoreServiceRequestModel,
	}
}

// Set-STFStoreFarm
type ApiSetSTFStoreFarmRequest struct {
	ctx                            context.Context
	ApiService                     *STFStore
	setSTFStoreFarmRequestModel    models.AddSTFStoreFarmRequestModel //Add and Set are same in this case
	getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel     //This is used to set the StoreService for the SetSTFStoreFarmRequest
}

func (r ApiSetSTFStoreFarmRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.SetSTFStoreFarmExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFStore) SetSTFStoreFarmExecute(r ApiSetSTFStoreFarmRequest) ([]byte, error) {
	var param = StructToString(r.setSTFStoreFarmRequestModel.BaseSTFStoreFarmRequestModel)
	var setStoreServiceParams = StructToString(r.getSTFStoreServiceRequestModel)

	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFStoreFarm", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", setStoreServiceParams), param)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFStoreFarm", param)
	}
}

func (a *STFStore) STFStoreSetStoreFarm(ctx context.Context, setSTFStoreFarmRequestModel models.AddSTFStoreFarmRequestModel, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiSetSTFStoreFarmRequest {
	return ApiSetSTFStoreFarmRequest{
		ApiService:                     a,
		ctx:                            ctx,
		setSTFStoreFarmRequestModel:    setSTFStoreFarmRequestModel,
		getSTFStoreServiceRequestModel: getSTFStoreRequestModel,
	}
}

// Get-STFStoreFarm
type ApiGetSTFStoreFarmRequest struct {
	ctx                            context.Context
	ApiService                     *STFStore
	getSTFStoreFarmRequestModel    models.GetSTFStoreFarmRequestModel
	getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel //This is used to set the StoreService for the GetSTFStoreFarmRequest
}

func (r ApiGetSTFStoreFarmRequest) Execute() (models.StoreFarmModel, error) {
	bytes, err := r.ApiService.GetSTFStoreFarmExecute(r)
	if err != nil {
		return models.StoreFarmModel{}, err
	}
	if len(bytes) == 0 {
		return models.StoreFarmModel{}, fmt.Errorf(NOT_EXIST)
	}

	var reponse = models.StoreFarmModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.StoreFarmModel{}, fmt.Errorf("error unmarshal STFStoreFarmModel: %v", err)
	}
	return reponse, nil
}

func (a *STFStore) GetSTFStoreFarmExecute(r ApiGetSTFStoreFarmRequest) ([]byte, error) {
	var param = StructToString(r.getSTFStoreFarmRequestModel)
	var getStoreServiceParams = StructToString(r.getSTFStoreServiceRequestModel)

	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFStoreFarm", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", getStoreServiceParams), param)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFStoreFarm", param)
	}
}

func (a *STFStore) STFStoreGetStoreFarm(ctx context.Context, getSTFStoreFarmRequestModel models.GetSTFStoreFarmRequestModel, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiGetSTFStoreFarmRequest {
	return ApiGetSTFStoreFarmRequest{
		ApiService:                     a,
		ctx:                            ctx,
		getSTFStoreFarmRequestModel:    getSTFStoreFarmRequestModel,
		getSTFStoreServiceRequestModel: getSTFStoreRequestModel,
	}
}

// Remove-STFStoreFarm
type ApiRemoveSTFStoreFarmRequest struct {
	ctx                            context.Context
	ApiService                     *STFStore
	removeSTFStoreFarmRequestModel models.GetSTFStoreFarmRequestModel
	getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel
}

func (r ApiRemoveSTFStoreFarmRequest) Execute() error {
	_, err := r.ApiService.RemoveSTFStoreFarmExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFStore) RemoveSTFStoreFarmExecute(r ApiRemoveSTFStoreFarmRequest) ([]byte, error) {
	var param = StructToString(r.removeSTFStoreFarmRequestModel)
	var removeStoreServiceParams = StructToString(r.getSTFStoreServiceRequestModel)

	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Remove-STFStoreFarm", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", removeStoreServiceParams), param)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Remove-STFStoreFarm", param)
	}
}

func (a *STFStore) STFStoreRemoveStoreFarm(ctx context.Context, removeModel models.GetSTFStoreFarmRequestModel, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiRemoveSTFStoreFarmRequest {
	return ApiRemoveSTFStoreFarmRequest{
		ApiService:                     a,
		ctx:                            ctx,
		removeSTFStoreFarmRequestModel: removeModel,
		getSTFStoreServiceRequestModel: getSTFStoreRequestModel,
	}
}

// Get-STFStoreFarmConfiguration
type ApiGetSTFStoreFarmConfigurationRequest struct {
	ctx                            context.Context
	ApiService                     *STFStore
	getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel
}

func (r ApiGetSTFStoreFarmConfigurationRequest) Execute() (models.StoreFarmConfigurationResponseModel, error) {
	bytes, err := r.ApiService.GetSTFStoreFarmConfigurationExecute(r)
	if err != nil {
		return models.StoreFarmConfigurationResponseModel{}, err
	}
	var rawResponse = models.StoreFarmConfigurationRawResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &rawResponse)
	if unMarshalErr != nil {
		return models.StoreFarmConfigurationResponseModel{}, fmt.Errorf("error unmarshal StoreFarmConfigurationResponseModel: %v", unMarshalErr.Error())
	}
	return rawResponse.ConvertToResponseModel(), nil
}

func (a *STFStore) GetSTFStoreFarmConfigurationExecute(r ApiGetSTFStoreFarmConfigurationRequest) ([]byte, error) {
	var param = StructToString(r.getSTFStoreServiceRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFStoreFarmConfiguration", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", param))
}

func (a *STFStore) STFStoreFarmGetStoreConfiguration(ctx context.Context, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiGetSTFStoreFarmConfigurationRequest {
	return ApiGetSTFStoreFarmConfigurationRequest{
		ApiService:                     a,
		ctx:                            ctx,
		getSTFStoreServiceRequestModel: getSTFStoreRequestModel,
	}
}

// Set-STFStoreFarmConfiguration
type ApiSetSTFStoreFarmConfigurationRequest struct {
	ctx                                   context.Context
	ApiService                            *STFStore
	setStoreFarmConfigurationRequestModel models.SetStoreFarmConfigurationRequestModel
	getSTFStoreServiceRequestModel        models.GetSTFStoreRequestModel
}

func (r ApiSetSTFStoreFarmConfigurationRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.SetSTFStoreFarmConfigurationExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFStore) SetSTFStoreFarmConfigurationExecute(r ApiSetSTFStoreFarmConfigurationRequest) ([]byte, error) {
	var param = StructToString(r.setStoreFarmConfigurationRequestModel)
	var setStoreFarmSettingsParams = StructToString(r.getSTFStoreServiceRequestModel)
	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFStoreFarmConfiguration", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", setStoreFarmSettingsParams), param)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFStoreFarmConfiguration", param, "-Confirm:$false")
	}
}

func (a *STFStore) STFStoreFarmSetSTFStoreConfiguration(ctx context.Context, setStoreFarmConfigurationRequestModel models.SetStoreFarmConfigurationRequestModel, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiSetSTFStoreFarmConfigurationRequest {
	return ApiSetSTFStoreFarmConfigurationRequest{
		ApiService:                            a,
		ctx:                                   ctx,
		setStoreFarmConfigurationRequestModel: setStoreFarmConfigurationRequestModel,
		getSTFStoreServiceRequestModel:        getSTFStoreRequestModel,
	}
}

// Get-STFStoreEnumerationOptions
type ApiGetSTFStoreEnumerationOptionsRequest struct {
	ctx                            context.Context
	ApiService                     *STFStore
	getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel
}

func (r ApiGetSTFStoreEnumerationOptionsRequest) Execute() (models.GetSTFStoreEnumerationOptionsResponseModel, error) {
	bytes, err := r.ApiService.GetSTFStoreEnumerationOptionsExecute(r)
	if err != nil {
		return models.GetSTFStoreEnumerationOptionsResponseModel{}, err
	}
	if len(bytes) == 0 {
		return models.GetSTFStoreEnumerationOptionsResponseModel{}, fmt.Errorf(NOT_EXIST)
	}
	var rawResponse = models.GetSTFStoreEnumerationOptionsRawResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &rawResponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.GetSTFStoreEnumerationOptionsResponseModel{}, fmt.Errorf("error occurred while unmarshalling EnumerationOptions response: %v", unMarshalErr.Error())
	}
	return rawResponse.ConvertToResponseModel(), nil
}

func (a *STFStore) GetSTFStoreEnumerationOptionsExecute(r ApiGetSTFStoreEnumerationOptionsRequest) ([]byte, error) {
	var setStoreEnumerationOptionsParams = StructToString(r.getSTFStoreServiceRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFStoreEnumerationOptions", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", setStoreEnumerationOptionsParams))
}

func (a *STFStore) STFStoreGetSTFStoreEnumerationOptions(ctx context.Context, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiGetSTFStoreEnumerationOptionsRequest {
	return ApiGetSTFStoreEnumerationOptionsRequest{
		ApiService:                     a,
		ctx:                            ctx,
		getSTFStoreServiceRequestModel: getSTFStoreRequestModel,
	}
}

// Enable-STFStorePna
type ApiEnableSTFStorePnaRequest struct {
	ctx                            context.Context
	ApiService                     *STFStore
	enableSTFStorePnaModel         models.STFStorePnaSetRequestModel
	getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel //This is used to set the StoreService for the SetSTFStoreFarmRequest
}

func (r ApiEnableSTFStorePnaRequest) Execute() error {
	_, err := r.ApiService.EnableSTFStorePnaExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFStore) EnableSTFStorePnaExecute(r ApiEnableSTFStorePnaRequest) ([]byte, error) {
	var param = StructToString(r.enableSTFStorePnaModel)
	var removeStoreServiceParams = StructToString(r.getSTFStoreServiceRequestModel)

	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Enable-STFStorePna", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", removeStoreServiceParams), param)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Enable-STFStorePna", param)
	}
}

func (a *STFStore) STFStoreEnableStorePna(ctx context.Context, enablePnaModel models.STFStorePnaSetRequestModel, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiEnableSTFStorePnaRequest {
	return ApiEnableSTFStorePnaRequest{
		ApiService:                     a,
		ctx:                            ctx,
		enableSTFStorePnaModel:         enablePnaModel,
		getSTFStoreServiceRequestModel: getSTFStoreRequestModel,
	}
}

// Disable-STFStorePna
type ApiDisableSTFStorePnaRequest struct {
	ctx                            context.Context
	ApiService                     *STFStore
	getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel //This is used to set the StoreService for the SetSTFStoreFarmRequest
}

func (r ApiDisableSTFStorePnaRequest) Execute() error {
	_, err := r.ApiService.DisableSTFStorePnaExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFStore) DisableSTFStorePnaExecute(r ApiDisableSTFStorePnaRequest) ([]byte, error) {
	var removeStoreServiceParams = StructToString(r.getSTFStoreServiceRequestModel)

	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Disable-STFStorePna", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", removeStoreServiceParams), "-Confirm:$false")
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Disable-STFStorePna", "-Confirm:$false")
	}
}

func (a *STFStore) STFStoreDisableStorePna(ctx context.Context, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiDisableSTFStorePnaRequest {
	return ApiDisableSTFStorePnaRequest{
		ApiService:                     a,
		ctx:                            ctx,
		getSTFStoreServiceRequestModel: getSTFStoreRequestModel,
	}
}

// Set-STFStoreEnumerationOptions

type ApiSetSTFStoreEnumerationOptionsRequest struct {
	ctx                                       context.Context
	ApiService                                *STFStore
	setSTFStoreEnumerationOptionsRequestModel models.SetSTFStoreEnumerationOptionsRequestModel
	getSTFStoreServiceRequestModel            models.GetSTFStoreRequestModel
}

func (r ApiSetSTFStoreEnumerationOptionsRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.SetSTFStoreEnumerationOptionsExecute(r)
	if err != nil {
		return bytes, err
	}
	return nil, nil
}

func (a *STFStore) SetSTFStoreEnumerationOptionsExecute(r ApiSetSTFStoreEnumerationOptionsRequest) ([]byte, error) {
	var param = StructToString(r.setSTFStoreEnumerationOptionsRequestModel)
	var setStoreEnumerationOptionsParams = StructToString(r.getSTFStoreServiceRequestModel)

	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFStoreEnumerationOptions", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", setStoreEnumerationOptionsParams), param)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFStoreEnumerationOptions", param)
	}
}

func (a *STFStore) STFStoreSetSTFStoreEnumerationOptions(ctx context.Context, setSTFStoreEnumerationOptionsRequestModel models.SetSTFStoreEnumerationOptionsRequestModel, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiSetSTFStoreEnumerationOptionsRequest {
	return ApiSetSTFStoreEnumerationOptionsRequest{
		ApiService: a,
		ctx:        ctx,
		setSTFStoreEnumerationOptionsRequestModel: setSTFStoreEnumerationOptionsRequestModel,
		getSTFStoreServiceRequestModel:            getSTFStoreRequestModel,
	}
}

// Get-STFStorePna
type ApiGETSTFStorePnaRequest struct {
	ctx                            context.Context
	ApiService                     *STFStore
	getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel //This is used to set the StoreService for the SetSTFStoreFarmRequest
}

type tempStoreService struct {
	FeatureData json.RawMessage `json:"FeatureData"`
}

func (r ApiGETSTFStorePnaRequest) Execute() (models.STFPna, error) {
	bytes, err := r.ApiService.GetSTFStorePnaExecute(r)
	if err != nil {
		return models.STFPna{}, err
	}
	if len(bytes) == 0 {
		return models.STFPna{}, fmt.Errorf(NOT_EXIST)
	}
	var reponse = models.STFPna{}
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(bytes, &objmap)
	if err != nil {
		fmt.Println("Error:", err)
		return models.STFPna{}, fmt.Errorf("error unmarshal STFStorePna: %v", err)
	}
	err = json.Unmarshal(*objmap["DefaultPnaService"], &reponse.DefaultPnaService)
	if err != nil {
		fmt.Println("Error:", err)
		return models.STFPna{}, fmt.Errorf("error unmarshal STFStorePna: %v", err)
	}
	var storeService tempStoreService
	err = json.Unmarshal(*objmap["StoreService"], &storeService)
	if err != nil {
		fmt.Println("Error:", err)
		return models.STFPna{}, fmt.Errorf("error unmarshal StoreService: %v", err)
	}
	err = json.Unmarshal(storeService.FeatureData, &reponse.FeatureData)
	if err != nil {
		fmt.Println("Error:", err)
		return models.STFPna{}, fmt.Errorf("error unmarshal STFStorePna: %v", err)
	}
	err = json.Unmarshal(*objmap["PnaEnabled"], &reponse.PnaEnabled)
	if err != nil {
		fmt.Println("Error:", err)
		return models.STFPna{}, fmt.Errorf("error unmarshal STFStorePna: %v", err)
	}
	return reponse, nil
}

func (a *STFStore) GetSTFStorePnaExecute(r ApiGETSTFStorePnaRequest) ([]byte, error) {
	var removeStoreServiceParams = StructToString(r.getSTFStoreServiceRequestModel)

	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFStorePna", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", removeStoreServiceParams))
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFStorePna", "-Confirm:$false")
	}
}

func (a *STFStore) STFStoreGetStorePna(ctx context.Context, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiGETSTFStorePnaRequest {
	return ApiGETSTFStorePnaRequest{
		ApiService:                     a,
		ctx:                            ctx,
		getSTFStoreServiceRequestModel: getSTFStoreRequestModel,
	}
}

// Get-STFStoreLaunchOptions
type ApiGetSTFStoreLaunchOptionsRequest struct {
	ctx                            context.Context
	ApiService                     *STFStore
	getSTFStoreServiceRequestModel models.GetSTFStoreRequestModel
}

func (r ApiGetSTFStoreLaunchOptionsRequest) Execute() (models.GetSTFStoreLaunchOptionsResponseModel, error) {
	bytes, err := r.ApiService.GetSTFStoreLaunchOptionsExecute(r)
	if err != nil {
		return models.GetSTFStoreLaunchOptionsResponseModel{}, err
	}

	var rawResponse = models.GetSTFStoreLaunchOptionsRawResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &rawResponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.GetSTFStoreLaunchOptionsResponseModel{}, fmt.Errorf("error occurred while unmarshalling LaunchOptions response: %v", unMarshalErr.Error())
	}
	return rawResponse.ConvertToResponseModel(), nil
}

func (a *STFStore) GetSTFStoreLaunchOptionsExecute(r ApiGetSTFStoreLaunchOptionsRequest) ([]byte, error) {
	var setStoreLaunchOptionsParams = StructToString(r.getSTFStoreServiceRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFStoreLaunchOptions", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", setStoreLaunchOptionsParams))
}

func (a *STFStore) STFStoreGetSTFStoreLaunchOptions(ctx context.Context, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiGetSTFStoreLaunchOptionsRequest {
	return ApiGetSTFStoreLaunchOptionsRequest{
		ApiService:                     a,
		ctx:                            ctx,
		getSTFStoreServiceRequestModel: getSTFStoreRequestModel,
	}
}

// Set-STFStoreLaunchOptions
type ApiSetSTFStoreLaunchOptionsRequest struct {
	ctx                             context.Context
	ApiService                      *STFStore
	setSTFStoreLaunchOptionsRequest models.SetSTFStoreLaunchOptionsRequestModel
	getSTFStoreServiceRequestModel  models.GetSTFStoreRequestModel
}

func (r ApiSetSTFStoreLaunchOptionsRequest) Execute() error {
	_, err := r.ApiService.SetSTFStoreLaunchOptionsExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFStore) SetSTFStoreLaunchOptionsExecute(r ApiSetSTFStoreLaunchOptionsRequest) ([]byte, error) {
	var setStoreLaunchOptionsParams = StructToString(r.setSTFStoreLaunchOptionsRequest)
	var getStoreLaunchOptionsParams = StructToString(r.getSTFStoreServiceRequestModel)

	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFStoreLaunchOptions", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", getStoreLaunchOptionsParams), setStoreLaunchOptionsParams)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFStoreLaunchOptions", setStoreLaunchOptionsParams)
	}
}

func (a *STFStore) STFStoreSetSTFStoreLaunchOptions(ctx context.Context, setSTFStoreLaunchOptionsRequest models.SetSTFStoreLaunchOptionsRequestModel, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiSetSTFStoreLaunchOptionsRequest {
	return ApiSetSTFStoreLaunchOptionsRequest{
		ApiService:                      a,
		ctx:                             ctx,
		setSTFStoreLaunchOptionsRequest: setSTFStoreLaunchOptionsRequest,
		getSTFStoreServiceRequestModel:  getSTFStoreRequestModel,
	}
}

// Get-STFRoamingAccount

type ApiGetSTFRoamingAccountRequest struct {
	ctx                     context.Context
	ApiService              *STFStore
	getSTFStoreRequestModel models.GetSTFStoreRequestModel
}

func (r ApiGetSTFRoamingAccountRequest) Execute() (models.GetSTFRoamingAccountResponseModel, error) {
	bytes, err := r.ApiService.GetSTFRoamingAccountExecute(r)
	if err != nil {
		return models.GetSTFRoamingAccountResponseModel{}, err
	}
	var reponse = models.GetSTFRoamingAccountResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		return models.GetSTFRoamingAccountResponseModel{}, fmt.Errorf("error unmarshal GetSTFRoamingAccountResponseModel: %v", unMarshalErr.Error())
	}
	return reponse, nil
}

func (a *STFStore) GetSTFRoamingAccountExecute(r ApiGetSTFRoamingAccountRequest) ([]byte, error) {
	var param = StructToString(r.getSTFStoreRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFRoamingAccount", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", param))
}

func (a *STFStore) STFRoamingAccountGet(ctx context.Context, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiGetSTFRoamingAccountRequest {
	return ApiGetSTFRoamingAccountRequest{
		ApiService:              a,
		ctx:                     ctx,
		getSTFStoreRequestModel: getSTFStoreRequestModel,
	}
}

// Set-STFRoamingAccount

type ApiSetSTFRoamingAccountRequest struct {
	ctx                              context.Context
	ApiService                       *STFStore
	SetSTFRoamingAccountRequestModel models.SetSTFRoamingAccountRequestModel
	getSTFStoreServiceRequestModel   models.GetSTFStoreRequestModel
}

func (r ApiSetSTFRoamingAccountRequest) Execute() error {
	_, err := r.ApiService.SetSTFRoamingAccountExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFStore) SetSTFRoamingAccountExecute(r ApiSetSTFRoamingAccountRequest) ([]byte, error) {
	var getStoreParam = StructToString(r.getSTFStoreServiceRequestModel)
	var getRoamAccParam = StructToString(r.SetSTFRoamingAccountRequestModel)
	if r.getSTFStoreServiceRequestModel.VirtualPath.IsSet() && *r.getSTFStoreServiceRequestModel.VirtualPath.Get() != "" {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFRoamingAccount", fmt.Sprintf("-StoreService (Get-STFStoreService %s)", getStoreParam), getRoamAccParam)
	} else {
		return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFRoamingAccount", getRoamAccParam, "-Confirm:$false")
	}
}

func (a *STFStore) STFRoamingAccountSet(ctx context.Context, setSTFRoamingAccountRequestModel models.SetSTFRoamingAccountRequestModel, getSTFStoreRequestModel models.GetSTFStoreRequestModel) ApiSetSTFRoamingAccountRequest {
	return ApiSetSTFRoamingAccountRequest{
		ApiService:                       a,
		ctx:                              ctx,
		SetSTFRoamingAccountRequestModel: setSTFRoamingAccountRequestModel,
		getSTFStoreServiceRequestModel:   getSTFStoreRequestModel,
	}
}
