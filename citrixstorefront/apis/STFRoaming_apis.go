// Copyright © 2025. Citrix Systems, Inc.
package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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
	return ExecuteCommandWithDepth(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), 5, "Get-STFRoamingService", params)
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
	param = strings.Replace(param, "-SessionReliability ", "-SessionReliability:", -1)
	param = strings.Replace(param, "-RequestTicketTwoSTAs ", "-RequestTicketTwoSTAs:", -1)
	param = strings.Replace(param, "-StasUseLoadBalancing ", "-StasUseLoadBalancing:", -1)
	var getRoamingServiceParam = StructToString(r.GetSTFRoamingServiceRequestModel)
	staUrlParam := ""
	if len(r.STFStaUrls) > 0 {
		staUrlParam = " -SecureTicketAuthorityObjs @("
		for _, staUrl := range r.STFStaUrls {
			staUrlParam += fmt.Sprintf("(New-STFSecureTicketAuthority -StaUrl '%s' -StaValidationEnabled $%t -StaValidationSecret '%s');", *staUrl.StaUrl.Get(), *staUrl.StaValidationEnabled.Get(), *staUrl.StaValidationSecret.Get())
		}
		staUrlParam += ")"
	}
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Add-STFRoamingGateway", fmt.Sprintf("-RoamingService (Get-STFRoamingService %s)", getRoamingServiceParam), param, staUrlParam)
}

func (a *STFRoaming) STFRoamingGatewayAdd(ctx context.Context, createSTFRoamingGatewayRequestModel models.AddSTFRoamingGatewayRequestModel, getSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel, stfStaUrls []models.STFSTAUrlModel) ApiAddSTFRoamingGatewayRequest {
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
}

func (r ApiGetSTFRoamingGatewayRequest) Execute() ([]models.STFRoamingGatewayResponseModel, error) {
	bytes, err := r.ApiService.GetSTFRoamingGatewayExecute(r)

	if err != nil {
		return []models.STFRoamingGatewayResponseModel{}, err
	}
	if len(bytes) == 0 {
		return []models.STFRoamingGatewayResponseModel{}, nil
	}

	gateway := string(bytes)
	updatedResponses := []models.STFRoamingGatewayResponseModel{}

	if string(gateway[0]) == "{" {
		var reponses = models.STFRoamingGatewayRawResponseModel{}
		unMarshalErr := json.Unmarshal(bytes, &reponses)
		if unMarshalErr != nil {
			return []models.STFRoamingGatewayResponseModel{}, fmt.Errorf("error unmarshal STFRoamingGatewayDetailModel: %v", unMarshalErr)
		}

		reponses.SiteId = *r.GetSTFRoamingServiceRequestModel.SiteId.Get()
		updatedResponses = append(updatedResponses, reponses.ConvertToResponseModel())

	} else if string(gateway[0]) == "[" {
		var objectArraymap []interface{}
		unMarshalToInterfaceErr := json.Unmarshal([]byte(gateway), &objectArraymap)

		if unMarshalToInterfaceErr != nil {
			return []models.STFRoamingGatewayResponseModel{}, fmt.Errorf("error unmarshal STFRoamingGatewayDetailModel to Interface: %v", unMarshalToInterfaceErr)
		}

		for _, obj := range objectArraymap {
			response := models.STFRoamingGatewayRawResponseModel{}
			if myMap, ok := (obj.(map[string]interface{})); ok {
				sta, ok := myMap["SecureTicketAuthorityUrls"]
				if ok {
					if stastring, ok := sta.([]interface{}); ok {
						length := len(stastring)
						var updateStaString []interface{}
						for i := 0; i < length; i++ {
							if sta, ok := stastring[i].(string); ok {
								result := strings.Split(sta, ",")
								staMap := make(map[string]interface{})
								staMap["StaUrl"] = result[0]
								staMap["StaValidationEnabled"], err = strconv.ParseBool(result[1])
								if err != nil {
									return []models.STFRoamingGatewayResponseModel{}, fmt.Errorf("error converting STAValidationEnabled to bool: %v", err)
								}
								staMap["StaValidationSecret"] = result[2]
								updateStaString = append(updateStaString, staMap)
							}
						}
						myMap["SecureTicketAuthorityUrls"] = updateStaString
					}
				}
			}

			jsonData, marshallToJsonErr := json.Marshal(obj)
			if marshallToJsonErr != nil {
				return []models.STFRoamingGatewayResponseModel{}, fmt.Errorf("error marshalling STFRoamingGatewayDetailModel to JSON: %v", marshallToJsonErr)
			}

			unMarshalErr := json.Unmarshal(jsonData, &response)
			if unMarshalErr != nil {
				return []models.STFRoamingGatewayResponseModel{}, fmt.Errorf("error unmarshal STFRoamingGatewayDetailModel to ResponseModel: %v", unMarshalErr)
			}
			response.SiteId = *r.GetSTFRoamingServiceRequestModel.SiteId.Get()
			updatedResponses = append(updatedResponses, response.ConvertToResponseModel())
		}
	}
	return updatedResponses, nil
}

func (a *STFRoaming) GetSTFRoamingGatewayExecute(r ApiGetSTFRoamingGatewayRequest) ([]byte, error) {
	var getRoamingServiceParams = StructToString(r.GetSTFRoamingServiceRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Get-STFRoamingGateway", fmt.Sprintf("-RoamingService (Get-STFRoamingService %s)", getRoamingServiceParams))
}

func (a *STFRoaming) STFRoamingGatewayGet(ctx context.Context, getSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel) ApiGetSTFRoamingGatewayRequest {
	return ApiGetSTFRoamingGatewayRequest{
		ApiService:                       a,
		ctx:                              ctx,
		GetSTFRoamingServiceRequestModel: getSTFRoamingServiceRequestModel,
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
			staUrlParam += fmt.Sprintf("(New-STFSecureTicketAuthority -StaUrl '%s' -StaValidationEnabled $%t -StaValidationSecret '%s');", *staUrl.StaUrl.Get(), *staUrl.StaValidationEnabled.Get(), *staUrl.StaValidationSecret.Get())
		}
		staUrlParam += ")"
	}
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Set-STFRoamingGateway", fmt.Sprintf("-RoamingService (Get-STFRoamingService %s)", getRoamingServiceParam), param, staUrlParam)
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
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Remove-STFRoamingGateway", fmt.Sprintf("-RoamingService (Get-STFRoamingService %s)", getRoamingServiceParams), param, "-Confirm:$false")
}

func (a *STFRoaming) STFRoamingGatewayRemove(ctx context.Context, getSTFoamingGatewayRequestModel models.GetSTFRoamingGatewayRequestModel, getSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel) ApiRemoveSTFoamingGatewayRequest {
	return ApiRemoveSTFoamingGatewayRequest{
		ApiService:                       a,
		ctx:                              ctx,
		GetSTFRoamingGatewayRequestModel: getSTFoamingGatewayRequestModel,
		GetSTFRoamingServiceRequestModel: getSTFRoamingServiceRequestModel,
	}
}

// Set Roaming Beacon for Internal Param

type ApiSetRoamingInternalBeacon struct {
	ctx                                     context.Context
	ApiService                              *STFRoaming
	SetSTFRoamingInternalBeaconRequestModel models.SetSTFRoamingInternalBeaconRequestModel
}

func (r ApiSetRoamingInternalBeacon) Execute() error {
	_, err := r.ApiService.SetRoamingInternalBeaconExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFRoaming) SetRoamingInternalBeaconExecute(r ApiSetRoamingInternalBeacon) ([]byte, error) {
	var int_param = StructToString(r.SetSTFRoamingInternalBeaconRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Set-STFRoamingBeacon", int_param, "-Confirm:$false")

}

func (a *STFRoaming) SetRoamingInternalBeacon(ctx context.Context, setSTFRoamingInternalBeaconRequestModel models.SetSTFRoamingInternalBeaconRequestModel) ApiSetRoamingInternalBeacon {
	return ApiSetRoamingInternalBeacon{
		ApiService:                              a,
		ctx:                                     ctx,
		SetSTFRoamingInternalBeaconRequestModel: setSTFRoamingInternalBeaconRequestModel,
	}
}

// Get Roaming Beacon for Internal Param

type ApiGetRoamingInternalBeacon struct {
	ctx                              context.Context
	ApiService                       *STFRoaming
	GetSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel
}

func (r ApiGetRoamingInternalBeacon) Execute() (models.GetSTFRoamingInternalBeaconResponseModel, error) {
	bytes, err := r.ApiService.GetRoamingInternalBeaconExecute(r)
	if err != nil {
		return models.GetSTFRoamingInternalBeaconResponseModel{}, err
	}
	var reponse = models.GetSTFRoamingInternalBeaconResponseModel{}
	internal_string := string(bytes)
	internal_string = strings.Replace(internal_string, "\r", "", -1)
	internal_string = strings.Replace(internal_string, "\n", "", -1)
	internal_string = strings.Replace(internal_string, "\"", "", -1)
	internal_string = strings.Replace(internal_string, " ", "", -1)
	reponse.Internal = internal_string

	return reponse, nil
}

func (a *STFRoaming) GetRoamingInternalBeaconExecute(r ApiGetRoamingInternalBeacon) ([]byte, error) {
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Get-STFRoamingBeacon -Internal")

}

func (a *STFRoaming) GetRoamingInternalBeacon(ctx context.Context) ApiGetRoamingInternalBeacon {
	return ApiGetRoamingInternalBeacon{
		ApiService: a,
		ctx:        ctx,
	}
}

// Clear Roaming Beacon for Internal Param

type ApiRemoveSTFRoamingInternalBeaconRequest struct {
	ctx                              context.Context
	ApiService                       *STFRoaming
	GetSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel
}

func (r ApiRemoveSTFRoamingInternalBeaconRequest) Execute() error {
	_, err := r.ApiService.RemoveSTFRoamingBeaconExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFRoaming) RemoveSTFRoamingBeaconExecute(r ApiRemoveSTFRoamingInternalBeaconRequest) ([]byte, error) {
	var getRoamingServiceParams = StructToString(r.GetSTFRoamingServiceRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Clear-STFRoamingBeacon", fmt.Sprintf("-RoamingService (Get-STFRoamingService %s)", getRoamingServiceParams), "-Confirm:$false")
}

func (a *STFRoaming) STFRoamingBeaconInternalRemove(ctx context.Context, getSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel) ApiRemoveSTFRoamingInternalBeaconRequest {
	return ApiRemoveSTFRoamingInternalBeaconRequest{
		ApiService:                       a,
		ctx:                              ctx,
		GetSTFRoamingServiceRequestModel: getSTFRoamingServiceRequestModel,
	}
}

// Set Roaming Beacon for External Param

type ApiSetRoamingExternalBeacon struct {
	ctx                                     context.Context
	ApiService                              *STFRoaming
	SetSTFRoamingExternalBeaconRequestModel models.SetSTFRoamingExternalBeaconRequestModel
	SetSTFRoamingInternalBeaconRequestModel models.SetSTFRoamingInternalBeaconRequestModel
}

func (r ApiSetRoamingExternalBeacon) Execute() error {
	_, err := r.ApiService.SetRoamingExternalBeaconExecute(r)
	if err != nil {
		return err
	}
	return nil
}

func (a *STFRoaming) SetRoamingExternalBeaconExecute(r ApiSetRoamingExternalBeacon) ([]byte, error) {
	var int_param = StructToString(r.SetSTFRoamingInternalBeaconRequestModel)
	var ext_param = StructToString(r.SetSTFRoamingExternalBeaconRequestModel)
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Set-STFRoamingBeacon", ext_param, int_param, "-Confirm:$false")

}

func (a *STFRoaming) SetRoamingExternalBeacon(ctx context.Context, setSTFRoamingExternalBeaconRequestModel models.SetSTFRoamingExternalBeaconRequestModel, setSTFRoamingInternalBeaconRequestModel models.SetSTFRoamingInternalBeaconRequestModel) ApiSetRoamingExternalBeacon {
	return ApiSetRoamingExternalBeacon{
		ApiService:                              a,
		ctx:                                     ctx,
		SetSTFRoamingExternalBeaconRequestModel: setSTFRoamingExternalBeaconRequestModel,
		SetSTFRoamingInternalBeaconRequestModel: setSTFRoamingInternalBeaconRequestModel,
	}
}

// Get Roaming Beacon for External Param

type ApiGetRoamingExternalBeacon struct {
	ctx                              context.Context
	ApiService                       *STFRoaming
	GetSTFRoamingServiceRequestModel models.STFRoamingServiceRequestModel
}

func (r ApiGetRoamingExternalBeacon) Execute() (models.GetSTFRoamingExternalBeaconResponseModel, error) {
	bytes, err := r.ApiService.GetRoamingExternalBeaconExecute(r)
	if err != nil {
		return models.GetSTFRoamingExternalBeaconResponseModel{}, err
	}
	var reponse = models.GetSTFRoamingExternalBeaconResponseModel{}
	ext_string := string(bytes)
	ext_str_arr := strings.Split(ext_string, ",")
	for i := range ext_str_arr {
		ext_str_arr[i] = strings.Replace(ext_str_arr[i], "\r", "", -1)
		ext_str_arr[i] = strings.Replace(ext_str_arr[i], "\n", "", -1)
		ext_str_arr[i] = strings.Replace(ext_str_arr[i], "\\", "", -1)
		ext_str_arr[i] = strings.Replace(ext_str_arr[i], "\"", "", -1)
		ext_str_arr[i] = strings.Replace(ext_str_arr[i], " ", "", -1)
		ext_str_arr[i] = strings.Replace(ext_str_arr[i], "[", "", -1)
		ext_str_arr[i] = strings.Replace(ext_str_arr[i], "]", "", -1)
	}
	reponse.External = ext_str_arr
	return reponse, nil
}

func (a *STFRoaming) GetRoamingExternalBeaconExecute(r ApiGetRoamingExternalBeacon) ([]byte, error) {
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Get-STFRoamingBeacon -External")

}

func (a *STFRoaming) GetRoamingExternalBeacon(ctx context.Context) ApiGetRoamingExternalBeacon {
	return ApiGetRoamingExternalBeacon{
		ApiService: a,
		ctx:        ctx,
	}
}
