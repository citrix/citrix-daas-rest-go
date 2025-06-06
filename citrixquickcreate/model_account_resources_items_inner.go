/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
	"fmt"
)

// AccountResourcesItemsInner - struct for AccountResourcesItemsInner
type AccountResourcesItemsInner struct {
	AwsEdcAccountResourceAmiImage         *AwsEdcAccountResourceAmiImage
	AwsEdcAccountResourceByolRegistration *AwsEdcAccountResourceByolRegistration
	AwsEdcAccountResourceCidrRange        *AwsEdcAccountResourceCidrRange
	AwsEdcAccountResourceDirectory        *AwsEdcAccountResourceDirectory
	AwsEdcAccountResourceEc2InstanceType  *AwsEdcAccountResourceEc2InstanceType
	AwsEdcAccountResourceFile             *AwsEdcAccountResourceFile
	AwsEdcAccountResourceKeyPair          *AwsEdcAccountResourceKeyPair
	AwsEdcAccountResourceRegion           *AwsEdcAccountResourceRegion
	AwsEdcAccountResourceSecurityGroup    *AwsEdcAccountResourceSecurityGroup
	AwsEdcAccountResourceSubnet           *AwsEdcAccountResourceSubnet
	AwsEdcAccountResourceVpc              *AwsEdcAccountResourceVpc
}

// AwsEdcAccountResourceAmiImageAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceAmiImage wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceAmiImageAsAccountResourcesItemsInner(v *AwsEdcAccountResourceAmiImage) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceAmiImage: v,
	}
}

// AwsEdcAccountResourceByolRegistrationAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceByolRegistration wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceByolRegistrationAsAccountResourcesItemsInner(v *AwsEdcAccountResourceByolRegistration) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceByolRegistration: v,
	}
}

// AwsEdcAccountResourceCidrRangeAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceCidrRange wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceCidrRangeAsAccountResourcesItemsInner(v *AwsEdcAccountResourceCidrRange) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceCidrRange: v,
	}
}

// AwsEdcAccountResourceDirectoryAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceDirectory wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceDirectoryAsAccountResourcesItemsInner(v *AwsEdcAccountResourceDirectory) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceDirectory: v,
	}
}

// AwsEdcAccountResourceEc2InstanceTypeAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceEc2InstanceType wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceEc2InstanceTypeAsAccountResourcesItemsInner(v *AwsEdcAccountResourceEc2InstanceType) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceEc2InstanceType: v,
	}
}

// AwsEdcAccountResourceFileAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceFile wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceFileAsAccountResourcesItemsInner(v *AwsEdcAccountResourceFile) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceFile: v,
	}
}

// AwsEdcAccountResourceKeyPairAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceKeyPair wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceKeyPairAsAccountResourcesItemsInner(v *AwsEdcAccountResourceKeyPair) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceKeyPair: v,
	}
}

// AwsEdcAccountResourceRegionAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceRegion wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceRegionAsAccountResourcesItemsInner(v *AwsEdcAccountResourceRegion) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceRegion: v,
	}
}

// AwsEdcAccountResourceSecurityGroupAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceSecurityGroup wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceSecurityGroupAsAccountResourcesItemsInner(v *AwsEdcAccountResourceSecurityGroup) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceSecurityGroup: v,
	}
}

// AwsEdcAccountResourceSubnetAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceSubnet wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceSubnetAsAccountResourcesItemsInner(v *AwsEdcAccountResourceSubnet) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceSubnet: v,
	}
}

// AwsEdcAccountResourceVpcAsAccountResourcesItemsInner is a convenience function that returns AwsEdcAccountResourceVpc wrapped in AccountResourcesItemsInner
func AwsEdcAccountResourceVpcAsAccountResourcesItemsInner(v *AwsEdcAccountResourceVpc) AccountResourcesItemsInner {
	return AccountResourcesItemsInner{
		AwsEdcAccountResourceVpc: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AccountResourcesItemsInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AWS_AMI'
	if jsonDict["resourceType"] == "AWS_AMI" {
		// try to unmarshal JSON data into AwsEdcAccountResourceAmiImage
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceAmiImage)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceAmiImage, return on the first match
		} else {
			dst.AwsEdcAccountResourceAmiImage = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceAmiImage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_BYOL_REGISTRATION'
	if jsonDict["resourceType"] == "AWS_BYOL_REGISTRATION" {
		// try to unmarshal JSON data into AwsEdcAccountResourceByolRegistration
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceByolRegistration)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceByolRegistration, return on the first match
		} else {
			dst.AwsEdcAccountResourceByolRegistration = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceByolRegistration: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_CIDRRANGE'
	if jsonDict["resourceType"] == "AWS_CIDRRANGE" {
		// try to unmarshal JSON data into AwsEdcAccountResourceCidrRange
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceCidrRange)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceCidrRange, return on the first match
		} else {
			dst.AwsEdcAccountResourceCidrRange = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceCidrRange: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_CLOUDFORMATION'
	if jsonDict["resourceType"] == "AWS_CLOUDFORMATION" {
		// try to unmarshal JSON data into AwsEdcAccountResourceFile
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceFile)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceFile, return on the first match
		} else {
			dst.AwsEdcAccountResourceFile = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceFile: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_DIRECTORY'
	if jsonDict["resourceType"] == "AWS_DIRECTORY" {
		// try to unmarshal JSON data into AwsEdcAccountResourceDirectory
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceDirectory)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceDirectory, return on the first match
		} else {
			dst.AwsEdcAccountResourceDirectory = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceDirectory: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_EC2_INSTANCE_TYPE'
	if jsonDict["resourceType"] == "AWS_EC2_INSTANCE_TYPE" {
		// try to unmarshal JSON data into AwsEdcAccountResourceEc2InstanceType
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceEc2InstanceType)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceEc2InstanceType, return on the first match
		} else {
			dst.AwsEdcAccountResourceEc2InstanceType = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceEc2InstanceType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_EC2_KEY_PAIR'
	if jsonDict["resourceType"] == "AWS_EC2_KEY_PAIR" {
		// try to unmarshal JSON data into AwsEdcAccountResourceKeyPair
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceKeyPair)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceKeyPair, return on the first match
		} else {
			dst.AwsEdcAccountResourceKeyPair = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceKeyPair: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_REGIONS'
	if jsonDict["resourceType"] == "AWS_REGIONS" {
		// try to unmarshal JSON data into AwsEdcAccountResourceRegion
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceRegion)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceRegion, return on the first match
		} else {
			dst.AwsEdcAccountResourceRegion = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceRegion: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_SECGROUP'
	if jsonDict["resourceType"] == "AWS_SECGROUP" {
		// try to unmarshal JSON data into AwsEdcAccountResourceSecurityGroup
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceSecurityGroup)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceSecurityGroup, return on the first match
		} else {
			dst.AwsEdcAccountResourceSecurityGroup = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceSecurityGroup: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_SUBNET'
	if jsonDict["resourceType"] == "AWS_SUBNET" {
		// try to unmarshal JSON data into AwsEdcAccountResourceSubnet
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceSubnet)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceSubnet, return on the first match
		} else {
			dst.AwsEdcAccountResourceSubnet = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceSubnet: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_VPC'
	if jsonDict["resourceType"] == "AWS_VPC" {
		// try to unmarshal JSON data into AwsEdcAccountResourceVpc
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceVpc)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceVpc, return on the first match
		} else {
			dst.AwsEdcAccountResourceVpc = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceVpc: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceAmiImage'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceAmiImage" {
		// try to unmarshal JSON data into AwsEdcAccountResourceAmiImage
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceAmiImage)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceAmiImage, return on the first match
		} else {
			dst.AwsEdcAccountResourceAmiImage = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceAmiImage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceByolRegistration'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceByolRegistration" {
		// try to unmarshal JSON data into AwsEdcAccountResourceByolRegistration
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceByolRegistration)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceByolRegistration, return on the first match
		} else {
			dst.AwsEdcAccountResourceByolRegistration = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceByolRegistration: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceCidrRange'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceCidrRange" {
		// try to unmarshal JSON data into AwsEdcAccountResourceCidrRange
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceCidrRange)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceCidrRange, return on the first match
		} else {
			dst.AwsEdcAccountResourceCidrRange = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceCidrRange: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceDirectory'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceDirectory" {
		// try to unmarshal JSON data into AwsEdcAccountResourceDirectory
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceDirectory)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceDirectory, return on the first match
		} else {
			dst.AwsEdcAccountResourceDirectory = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceDirectory: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceEc2InstanceType'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceEc2InstanceType" {
		// try to unmarshal JSON data into AwsEdcAccountResourceEc2InstanceType
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceEc2InstanceType)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceEc2InstanceType, return on the first match
		} else {
			dst.AwsEdcAccountResourceEc2InstanceType = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceEc2InstanceType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceFile'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceFile" {
		// try to unmarshal JSON data into AwsEdcAccountResourceFile
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceFile)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceFile, return on the first match
		} else {
			dst.AwsEdcAccountResourceFile = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceFile: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceKeyPair'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceKeyPair" {
		// try to unmarshal JSON data into AwsEdcAccountResourceKeyPair
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceKeyPair)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceKeyPair, return on the first match
		} else {
			dst.AwsEdcAccountResourceKeyPair = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceKeyPair: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceRegion'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceRegion" {
		// try to unmarshal JSON data into AwsEdcAccountResourceRegion
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceRegion)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceRegion, return on the first match
		} else {
			dst.AwsEdcAccountResourceRegion = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceRegion: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceSecurityGroup'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceSecurityGroup" {
		// try to unmarshal JSON data into AwsEdcAccountResourceSecurityGroup
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceSecurityGroup)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceSecurityGroup, return on the first match
		} else {
			dst.AwsEdcAccountResourceSecurityGroup = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceSecurityGroup: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceSubnet'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceSubnet" {
		// try to unmarshal JSON data into AwsEdcAccountResourceSubnet
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceSubnet)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceSubnet, return on the first match
		} else {
			dst.AwsEdcAccountResourceSubnet = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceSubnet: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsEdcAccountResourceVpc'
	if jsonDict["resourceType"] == "AwsEdcAccountResourceVpc" {
		// try to unmarshal JSON data into AwsEdcAccountResourceVpc
		err = json.Unmarshal(data, &dst.AwsEdcAccountResourceVpc)
		if err == nil {
			return nil // data stored in dst.AwsEdcAccountResourceVpc, return on the first match
		} else {
			dst.AwsEdcAccountResourceVpc = nil
			return fmt.Errorf("failed to unmarshal AccountResourcesItemsInner as AwsEdcAccountResourceVpc: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AccountResourcesItemsInner) MarshalJSON() ([]byte, error) {
	if src.AwsEdcAccountResourceAmiImage != nil {
		return json.Marshal(&src.AwsEdcAccountResourceAmiImage)
	}

	if src.AwsEdcAccountResourceByolRegistration != nil {
		return json.Marshal(&src.AwsEdcAccountResourceByolRegistration)
	}

	if src.AwsEdcAccountResourceCidrRange != nil {
		return json.Marshal(&src.AwsEdcAccountResourceCidrRange)
	}

	if src.AwsEdcAccountResourceDirectory != nil {
		return json.Marshal(&src.AwsEdcAccountResourceDirectory)
	}

	if src.AwsEdcAccountResourceEc2InstanceType != nil {
		return json.Marshal(&src.AwsEdcAccountResourceEc2InstanceType)
	}

	if src.AwsEdcAccountResourceFile != nil {
		return json.Marshal(&src.AwsEdcAccountResourceFile)
	}

	if src.AwsEdcAccountResourceKeyPair != nil {
		return json.Marshal(&src.AwsEdcAccountResourceKeyPair)
	}

	if src.AwsEdcAccountResourceRegion != nil {
		return json.Marshal(&src.AwsEdcAccountResourceRegion)
	}

	if src.AwsEdcAccountResourceSecurityGroup != nil {
		return json.Marshal(&src.AwsEdcAccountResourceSecurityGroup)
	}

	if src.AwsEdcAccountResourceSubnet != nil {
		return json.Marshal(&src.AwsEdcAccountResourceSubnet)
	}

	if src.AwsEdcAccountResourceVpc != nil {
		return json.Marshal(&src.AwsEdcAccountResourceVpc)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AccountResourcesItemsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AwsEdcAccountResourceAmiImage != nil {
		return obj.AwsEdcAccountResourceAmiImage
	}

	if obj.AwsEdcAccountResourceByolRegistration != nil {
		return obj.AwsEdcAccountResourceByolRegistration
	}

	if obj.AwsEdcAccountResourceCidrRange != nil {
		return obj.AwsEdcAccountResourceCidrRange
	}

	if obj.AwsEdcAccountResourceDirectory != nil {
		return obj.AwsEdcAccountResourceDirectory
	}

	if obj.AwsEdcAccountResourceEc2InstanceType != nil {
		return obj.AwsEdcAccountResourceEc2InstanceType
	}

	if obj.AwsEdcAccountResourceFile != nil {
		return obj.AwsEdcAccountResourceFile
	}

	if obj.AwsEdcAccountResourceKeyPair != nil {
		return obj.AwsEdcAccountResourceKeyPair
	}

	if obj.AwsEdcAccountResourceRegion != nil {
		return obj.AwsEdcAccountResourceRegion
	}

	if obj.AwsEdcAccountResourceSecurityGroup != nil {
		return obj.AwsEdcAccountResourceSecurityGroup
	}

	if obj.AwsEdcAccountResourceSubnet != nil {
		return obj.AwsEdcAccountResourceSubnet
	}

	if obj.AwsEdcAccountResourceVpc != nil {
		return obj.AwsEdcAccountResourceVpc
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AccountResourcesItemsInner) GetActualInstanceValue() interface{} {
	if obj.AwsEdcAccountResourceAmiImage != nil {
		return *obj.AwsEdcAccountResourceAmiImage
	}

	if obj.AwsEdcAccountResourceByolRegistration != nil {
		return *obj.AwsEdcAccountResourceByolRegistration
	}

	if obj.AwsEdcAccountResourceCidrRange != nil {
		return *obj.AwsEdcAccountResourceCidrRange
	}

	if obj.AwsEdcAccountResourceDirectory != nil {
		return *obj.AwsEdcAccountResourceDirectory
	}

	if obj.AwsEdcAccountResourceEc2InstanceType != nil {
		return *obj.AwsEdcAccountResourceEc2InstanceType
	}

	if obj.AwsEdcAccountResourceFile != nil {
		return *obj.AwsEdcAccountResourceFile
	}

	if obj.AwsEdcAccountResourceKeyPair != nil {
		return *obj.AwsEdcAccountResourceKeyPair
	}

	if obj.AwsEdcAccountResourceRegion != nil {
		return *obj.AwsEdcAccountResourceRegion
	}

	if obj.AwsEdcAccountResourceSecurityGroup != nil {
		return *obj.AwsEdcAccountResourceSecurityGroup
	}

	if obj.AwsEdcAccountResourceSubnet != nil {
		return *obj.AwsEdcAccountResourceSubnet
	}

	if obj.AwsEdcAccountResourceVpc != nil {
		return *obj.AwsEdcAccountResourceVpc
	}

	// all schemas are nil
	return nil
}

type NullableAccountResourcesItemsInner struct {
	value *AccountResourcesItemsInner
	isSet bool
}

func (v NullableAccountResourcesItemsInner) Get() *AccountResourcesItemsInner {
	return v.value
}

func (v *NullableAccountResourcesItemsInner) Set(val *AccountResourcesItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResourcesItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResourcesItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResourcesItemsInner(val *AccountResourcesItemsInner) *NullableAccountResourcesItemsInner {
	return &NullableAccountResourcesItemsInner{value: val, isSet: true}
}

func (v NullableAccountResourcesItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResourcesItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
