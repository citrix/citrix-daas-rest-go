// Copyright © 2024. Citrix Systems, Inc.
package apis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/citrix/citrix-daas-rest-go/citrixstorefront/models"
)

type STFVersion Service

type ApiGetSTFVersionRequest struct {
	ctx        context.Context
	ApiService *STFVersion
}

func (r ApiGetSTFVersionRequest) Execute() (models.STFVersionModel, error) {
	bytes, err := r.ApiService.GetSTFVersionExecute(r)
	if err != nil {
		return models.STFVersionModel{}, err
	}
	var reponse = models.STFVersionModel{}
	unMarshalErr := json.Unmarshal(bytes, &reponse)
	if unMarshalErr != nil {
		fmt.Println("Error:", unMarshalErr)
		return models.STFVersionModel{}, fmt.Errorf("error unmarshal STFVersionModel: %v", unMarshalErr.Error())
	}
	return reponse, nil
}

func (a *STFVersion) GetSTFVersionExecute(r ApiGetSTFVersionRequest) ([]byte, error) {
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword(), a.client.GetDisableSSL()), "Get-STFVersion")
}

func (a *STFVersion) STFVersionGetVersion(ctx context.Context) ApiGetSTFVersionRequest {
	return ApiGetSTFVersionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}
