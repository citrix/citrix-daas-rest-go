// Copyright © 2025. Citrix Systems, Inc.
package models

var _ MappedNullable = &CreateSTFDeploymentRequestModel{}

type STFDeploymentDetailModel struct {
	SiteId                  NullableInt64                  `json:"SiteId,omitempty"`                  // The IIS site id of the deployment
	HostBaseUrl             NullableString                 `json:"HostBaseUrl,omitempty"`             // Url used to access the StoreFront server group
	DeploymentExists        NullableBool                   `json:"DeploymentExists,omitempty"`        // Deployment exists
	InstalledFeatureClasses []InstalledFeatureClassesModel `json:"InstalledFeatureClasses,omitempty"` // InstalledFeatureClasses
	FeatureClassInstances   []FeatureClassInstancesModel   `json:"FeatureClassInstances,omitempty"`   // FeatureClassInstances
}

type InstalledFeatureClassesModel struct {
	FrameworkController         NullableString `json:"FrameworkController,omitempty"`         // Framework controller
	Type                        NullableString `json:"Type,omitempty"`                        // Type
	Name                        NullableString `json:"Name,omitempty"`                        // Name
	Version                     NullableString `json:"Version,omitempty"`                     // Version
	PreviousVersion             NullableString `json:"PreviousVersion,omitempty"`             // Previous version
	PackageFilename             NullableString `json:"PackageFilename,omitempty"`             // Package filename
	ParentType                  NullableString `json:"ParentType,omitempty"`                  // Parent type
	Multiplicity                NullableInt    `json:"Multiplicity,omitempty"`                // Multiplicity
	LocalisationPath            NullableString `json:"LocalisationPath,omitempty"`            // Localisation path
	TitleKey                    NullableString `json:"TitleKey,omitempty"`                    // Title key
	DescriptionKey              NullableString `json:"DescriptionKey,omitempty"`              // Description key
	DirectoryPath               NullableString `json:"DirectoryPath,omitempty"`               // Directory path
	InstanceFactoryAssemblyPath NullableString `json:"InstanceFactoryAssemblyPath,omitempty"` // Instance factory assembly path
	InstanceFactoryType         NullableString `json:"InstanceFactoryType,omitempty"`         // Instance factory type
	ClassDependencies           NullableString `json:"ClassDependencies,omitempty"`           // Class dependencies
	RegisterLibraries           NullableString `json:"RegisterLibraries,omitempty"`           // Register libraries
	InstallUtilLibraries        NullableString `json:"InstallUtilLibraries,omitempty"`        // Install util libraries
	GacLibraries                NullableString `json:"GacLibraries,omitempty"`                // Gac libraries
	PowerShellSdkFiles          NullableString `json:"PowerShellSdkFiles,omitempty"`          // PowerShell sdk files
	TemplateFiles               NullableString `json:"TemplateFiles,omitempty"`               // Template files
	ActiveDirectoryGroups       NullableString `json:"ActiveDirectoryGroups,omitempty"`       // Active directory groups
	TransformationDefinitions   NullableString `json:"TransformationDefinitions,omitempty"`   // Transformation definitions
	PowerShellSnapin            NullableString `json:"PowerShellSnapin,omitempty"`            // PowerShell snapin
	PowerShellAddCmdlet         NullableString `json:"PowerShellAddCmdlet,omitempty"`         // PowerShell add cmdlet
	Properties                  NullableString `json:"Properties,omitempty"`                  // Properties
	InstanceFactory             NullableString `json:"InstanceFactory,omitempty"`             // Instance factory
}

type FeatureClassInstancesModel struct {
	Path                           NullableString `json:"Path,omitempty"`                           // C:\Program Files\Citrix\Receiver StoreFront\Services\DefaultDomainServices
	ConfigLocation                 NullableString `json:"ConfigLocation,omitempty"`                 // C:\Program Files\Citrix\Receiver StoreFront\Services\DefaultDomainServices\Citrix.DeliveryServices.DomainServices.ServiceHost.exe.config
	ConfigTypeName                 NullableString `json:"ConfigTypeName,omitempty"`                 // Executable
	Id                             NullableString `json:"Id,omitempty"`                             // 9108053d-479c-4798-9799-d20c71fc2905
	ClassType                      NullableString `json:"ClassType,omitempty"`                      // 88e61a71-f70a-47b7-85c5-f6ca6f108410
	FrameworkController            NullableString `json:"FrameworkController,omitempty"`            // Citrix.DeliveryServices.Framework.FileBased.FrameworkController
	ParentInstance                 NullableString `json:"ParentInstance,omitempty"`                 // 00000000-0000-0000-0000-000000000000
	RootInstance                   NullableString `json:"RootInstance,omitempty"`                   // 9108053d-479c-4798-9799-d20c71fc2905
	TenantId                       NullableString `json:"TenantId,omitempty"`                       // 860e9401-39c8-4f2c-928d-34251102b840
	Data                           NullableString `json:"Data,omitempty"`                           // System.Collections.Generic.Dictionary`2[System.String,System.String]
	ReadOnlyData                   NullableString `json:"ReadOnlyData,omitempty"`                   // [WindowsServiceName, CitrixDefaultDomainService] [Name, DomainServices] [Cmdlet, Add-DSDomainService] [Snapin, Citrix.DeliveryServices.DomainService.Install] [Tenant, 860e9401-39c8-4f2c-928d-34251102b840] [IsService, true]
	ParameterData                  NullableString `json:"ParameterData,omitempty"`                  // [Folder, C:\Program Files\Citrix\Receiver StoreFront\Services\DefaultDomainServices] [StartService, ]
	AdditionalInstanceDependencies NullableString `json:"AdditionalInstanceDependencies,omitempty"` // ""
	IsDeployed                     NullableBool   `json:"IsDeployed,omitempty"`                     // true
	FeatureClass                   NullableString `json:"FeatureClass,omitempty"`                   // Citrix.DeliveryServices.Framework.Feature.FeatureClass
}
