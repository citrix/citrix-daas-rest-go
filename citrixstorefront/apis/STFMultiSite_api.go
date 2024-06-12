// Copyright © 2024. Citrix Systems, Inc.
package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/citrix/citrix-daas-rest-go/citrixstorefront/models"
)

type STFMultiSite Service

// Add-STFUserFarmMapping

type ApiAddSTFUserFarmMappingRequest struct {
	ctx                                  context.Context
	ApiService                           *STFMultiSite
	StoreVirtualPath                     models.NullableString
	UserFarmMappingName                  models.NullableString
	NewSTFEquivalentFarmSetRequestModels []models.STFEquivalentFarmSetRequestModel
	NewSTFUserFarmMappingGroups          []models.STFUserFarmMappingGroup
}

func (r ApiAddSTFUserFarmMappingRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.AddSTFUserFarmMappingExecute(r)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (a *STFMultiSite) AddSTFUserFarmMappingExecute(r ApiAddSTFUserFarmMappingRequest) ([]byte, error) {
	equivalentFarmSets := BuildEquivalentFarmSetsInput(r.NewSTFEquivalentFarmSetRequestModels)
	userFarmMappingGroupHashTable := BuildUserFarmMappingGroupHashTable(r.NewSTFUserFarmMappingGroups)

	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Add-STFUserFarmMapping", fmt.Sprintf("-StoreService (Get-STFStoreService -VirtualPath '%s')", *r.StoreVirtualPath.Get()), fmt.Sprintf("-Name '%s'", *r.UserFarmMappingName.Get()), fmt.Sprintf("-EquivalentFarmSet %s", equivalentFarmSets), fmt.Sprintf("-GroupMembers %s", userFarmMappingGroupHashTable))
}

func (a *STFMultiSite) STFMultiSiteAddUserFarmMapping(ctx context.Context, storeVirtualPath models.NullableString, userFarmMappingName models.NullableString, newSTFEquivalentFarmSetRequestModels []models.STFEquivalentFarmSetRequestModel, newSTFUserFarmMappingGroups []models.STFUserFarmMappingGroup) ApiAddSTFUserFarmMappingRequest {
	return ApiAddSTFUserFarmMappingRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		StoreVirtualPath:                     storeVirtualPath,
		UserFarmMappingName:                  userFarmMappingName,
		NewSTFEquivalentFarmSetRequestModels: newSTFEquivalentFarmSetRequestModels,
		NewSTFUserFarmMappingGroups:          newSTFUserFarmMappingGroups,
	}
}

// Get-STFUserFarmMapping

type ApiGetSTFUserFarmMappingRequest struct {
	ctx                    context.Context
	ApiService             *STFMultiSite
	STFUserFarmMappingName models.NullableString
	StoreVirtualPath       models.NullableString
}

func (r ApiGetSTFUserFarmMappingRequest) Execute() (models.STFUserFarmMappingResponseModel, error) {
	bytes, err := r.ApiService.GetSTFUserFarmMappingExecute(r)
	if err != nil {
		return models.STFUserFarmMappingResponseModel{}, err
	}
	if len(bytes) == 0 {
		return models.STFUserFarmMappingResponseModel{}, fmt.Errorf(NOT_EXIST)
	}
	var rawResponse = models.STFUserFarmMappingRawResponseModel{}
	unMarshalErr := json.Unmarshal(bytes, &rawResponse)
	if unMarshalErr != nil {
		return models.STFUserFarmMappingResponseModel{}, unMarshalErr
	}

	responseModel := models.STFUserFarmMappingResponseModel{}
	responseModel.Name.Set(rawResponse.Name.Get())
	responseModel.VirtualPath.Set(r.StoreVirtualPath.Get())
	farmSets := []models.STFFarmSetResponseModel{}
	for _, rawFarmSet := range rawResponse.FarmSets {
		farmSet := models.STFFarmSetResponseModel{}
		primaryFarms := []string{}
		backupFarms := []string{}

		farmSet.SetName(*rawFarmSet.Name.Get())
		farmSet.SetAggregationGroupName(*rawFarmSet.AggregationGroupName.Get())
		farmSet.SetFarmsAreIdentical(*rawFarmSet.FarmsAreIdentical.Get())
		if *rawFarmSet.LoadBalanceMode.Get() == 0 {
			farmSet.SetLoadBalanceMode("Failover")
		} else if *rawFarmSet.LoadBalanceMode.Get() == 1 {
			farmSet.SetLoadBalanceMode("LoadBalanced")
		}

		for _, rawPrimaryFarm := range strings.Split(*rawFarmSet.PrimaryFarms.Get(), " ") {
			if !strings.EqualFold(rawPrimaryFarm, "") {
				primaryFarms = append(primaryFarms, rawPrimaryFarm)
			}
		}
		farmSet.SetPrimaryFarms(primaryFarms)

		for _, rawBackupFarm := range strings.Split(*rawFarmSet.BackupFarms.Get(), " ") {
			if !strings.EqualFold(rawBackupFarm, "") {
				backupFarms = append(backupFarms, rawBackupFarm)
			}
		}
		farmSet.SetBackupFarms(backupFarms)

		farmSets = append(farmSets, farmSet)
	}

	responseModel.FarmSets = farmSets

	groupMembers := []models.STFGroupMemberResponseModel{}
	for groupName, accountSid := range rawResponse.GroupMembers {
		groupMember := models.STFGroupMemberResponseModel{}
		groupMember.SetGroupName(groupName)
		groupMember.SetAccountSid(accountSid)
		groupMembers = append(groupMembers, groupMember)
	}
	responseModel.GroupMembers = groupMembers
	return responseModel, nil
}

func (a *STFMultiSite) GetSTFUserFarmMappingExecute(r ApiGetSTFUserFarmMappingRequest) ([]byte, error) {
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Get-STFUserFarmMapping", fmt.Sprintf("-StoreService (Get-STFStoreService -VirtualPath '%s')", *r.StoreVirtualPath.Get()), fmt.Sprintf("-Name '%s'", *r.STFUserFarmMappingName.Get()))
}

func (a *STFMultiSite) STFMultiSiteGetUserFarmMapping(ctx context.Context, storeVirtualPath models.NullableString, getSTFUserFarmMappingName models.NullableString) ApiGetSTFUserFarmMappingRequest {
	return ApiGetSTFUserFarmMappingRequest{
		ApiService:             a,
		ctx:                    ctx,
		STFUserFarmMappingName: getSTFUserFarmMappingName,
		StoreVirtualPath:       storeVirtualPath,
	}
}

// Set-STFUserFarmMapping
type ApiSetSTFUserFarmMappingRequest struct {
	ctx                                  context.Context
	ApiService                           *STFMultiSite
	StoreVirtualPath                     models.NullableString
	UserFarmMappingName                  models.NullableString
	SetSTFEquivalentFarmSetRequestModels []models.STFEquivalentFarmSetRequestModel
	SetSTFUserFarmMappingGroups          []models.STFUserFarmMappingGroup
}

func (r ApiSetSTFUserFarmMappingRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.SetSTFUserFarmMappingExecute(r)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (a *STFMultiSite) SetSTFUserFarmMappingExecute(r ApiSetSTFUserFarmMappingRequest) ([]byte, error) {
	equivalentFarmSets := BuildEquivalentFarmSetsInput(r.SetSTFEquivalentFarmSetRequestModels)
	userFarmMappingGroups := BuildUserFarmMappingGroupHashTable(r.SetSTFUserFarmMappingGroups)

	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Set-STFUserFarmMapping", fmt.Sprintf("-StoreService (Get-STFStoreService -VirtualPath '%s')", *r.StoreVirtualPath.Get()), fmt.Sprintf("-UserFarmMapping (Get-STFUserFarmMapping -StoreService (Get-STFStoreService -VirtualPath '%s') -Name '%s')", *r.StoreVirtualPath.Get(), *r.UserFarmMappingName.Get()), fmt.Sprintf("-EquivalentFarmSet %s", equivalentFarmSets), fmt.Sprintf("-GroupMembers %s", userFarmMappingGroups))
}

func (a *STFMultiSite) STFMultiSiteSetUserFarmMapping(ctx context.Context, storeVirtualPath models.NullableString, userFarmMappingName models.NullableString, setSTFEquivalentFarmSetRequestModels []models.STFEquivalentFarmSetRequestModel, setSTFUserFarmMappingGroups []models.STFUserFarmMappingGroup) ApiSetSTFUserFarmMappingRequest {
	return ApiSetSTFUserFarmMappingRequest{
		ApiService:                           a,
		ctx:                                  ctx,
		StoreVirtualPath:                     storeVirtualPath,
		UserFarmMappingName:                  userFarmMappingName,
		SetSTFEquivalentFarmSetRequestModels: setSTFEquivalentFarmSetRequestModels,
		SetSTFUserFarmMappingGroups:          setSTFUserFarmMappingGroups,
	}
}

// Remove-STFUserFarmMapping
type ApiRemoveSTFUserFarmMappingRequest struct {
	ctx                 context.Context
	ApiService          *STFMultiSite
	UserFarmMappingName models.NullableString
	StoreVirtualPath    models.NullableString
}

func (r ApiRemoveSTFUserFarmMappingRequest) Execute() ([]byte, error) {
	bytes, err := r.ApiService.RemoveSTFUserFarmMappingExecute(r)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (a *STFMultiSite) RemoveSTFUserFarmMappingExecute(r ApiRemoveSTFUserFarmMappingRequest) ([]byte, error) {
	return ExecuteCommand(BuildAuth(a.client.GetComputerName(), a.client.GetAdUserName(), a.client.GetAdPassword()), "Remove-STFUserFarmMapping", fmt.Sprintf("-StoreService (Get-STFStoreService -VirtualPath '%s')", *r.StoreVirtualPath.Get()), fmt.Sprintf("-Name '%s'", *r.UserFarmMappingName.Get()), "-Confirm:$false")
}

func (a *STFMultiSite) STFMultiSiteRemoveUserFarmMapping(ctx context.Context, storeVirtualPath models.NullableString, userFarmMappingName models.NullableString) ApiRemoveSTFUserFarmMappingRequest {
	return ApiRemoveSTFUserFarmMappingRequest{
		ApiService:          a,
		ctx:                 ctx,
		UserFarmMappingName: userFarmMappingName,
		StoreVirtualPath:    storeVirtualPath,
	}
}

func BuildEquivalentFarmSetsInput(equivalentFarmSetModels []models.STFEquivalentFarmSetRequestModel) string {
	equivalentFarmSets := "@("
	for _, equivalentFarmSet := range equivalentFarmSetModels {
		primaryFarms := "@("
		for _, primaryFarm := range equivalentFarmSet.PrimaryFarms {
			primaryFarms += fmt.Sprintf("'%s';", primaryFarm)
		}
		primaryFarms += ")"

		backupFarms := ""
		if len(equivalentFarmSet.BackupFarms) > 0 {
			backupFarms += "-BackupFarms @(@("
			for _, backupFarm := range equivalentFarmSet.BackupFarms {
				backupFarms += fmt.Sprintf("'%s';", backupFarm)
			}
			backupFarms += "))"
		}

		equivalentFarmSets += fmt.Sprintf("(New-STFEquivalentFarmset -Name '%s' -LoadBalanceMode %s -FarmsAreIdentical $%t -AggregationGroupName '%s' -PrimaryFarms @(%s) %s);", *equivalentFarmSet.Name.Get(), *equivalentFarmSet.LoadBalanceMode.Get(), *equivalentFarmSet.FarmsAreIdentical.Get(), *equivalentFarmSet.AggregationGroupName.Get(), primaryFarms, backupFarms)
	}
	equivalentFarmSets += ")"
	return equivalentFarmSets
}

func BuildUserFarmMappingGroupHashTable(userFarmMappingGroups []models.STFUserFarmMappingGroup) string {
	userFarmMappingGroupHashTable := "@{"
	if len(userFarmMappingGroups) == 0 {
		userFarmMappingGroupHashTable += "'Everyone' = 'Everyone'"
	} else {
		for _, userFarmMappingGroup := range userFarmMappingGroups {
			userFarmMappingGroupHashTable += fmt.Sprintf("'%s' = '%s';", *userFarmMappingGroup.GroupName.Get(), *userFarmMappingGroup.AccountSid.Get())
		}
	}
	userFarmMappingGroupHashTable += "}"
	return userFarmMappingGroupHashTable
}
