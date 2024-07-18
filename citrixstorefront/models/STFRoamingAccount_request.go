// Copyright © 2024. Citrix Systems, Inc.

package models

// Get-STFRoamingAccount

var _ MappedNullable = &GetSTFRoamingAccountRequestModel{}

type GetSTFRoamingAccountRequestModel struct {
	StoreService NullableString //Configuration of a StoreFront Store Service

}

// ToMap implements MappedNullable.

func (o GetSTFRoamingAccountRequestModel) ToMap() (map[string]interface{}, error) {

	toSerialize := map[string]interface{}{}

	if o.StoreService.IsSet() {
		toSerialize["StoreService"] = o.StoreService.Get()
	}

	return toSerialize, nil
}

func (o *GetSTFRoamingAccountRequestModel) SetStoreService(v string) {
	o.StoreService.Set(&v)
}

// Set-STFRoamingAccount

var _ MappedNullable = &SetSTFRoamingAccountRequestModel{}

type SetSTFRoamingAccountRequestModel struct {
	StoreService NullableString //Configuration of a StoreFront Store Service
	Published    NullableBool   //Whether the account is published
}

// ToMap implements MappedNullable.

func (o SetSTFRoamingAccountRequestModel) ToMap() (map[string]interface{}, error) {

	toSerialize := map[string]interface{}{}

	if o.StoreService.IsSet() {
		toSerialize["StoreService"] = o.StoreService.Get()
	}

	if o.Published.IsSet() {
		toSerialize["Published"] = o.Published.Get()
	}

	return toSerialize, nil
}

func (o *SetSTFRoamingAccountRequestModel) SetStoreService(v string) {
	o.StoreService.Set(&v)
}

func (o *SetSTFRoamingAccountRequestModel) SetPublished(v bool) {
	o.Published.Set(&v)
}
