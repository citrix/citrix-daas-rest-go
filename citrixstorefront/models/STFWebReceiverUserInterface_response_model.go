package models

type UserInterfaceWorkspaceControlRawResponseModel struct {
	Enabled              NullableBool `json:"Enabled,omitempty"`
	AutoReconnectAtLogon NullableBool `json:"AutoReconnectAtLogon,omitempty"`
	LogoffAction         NullableInt  `json:"LogoffAction,omitempty"` // Enum {Disconnect,Terminate,None}
	ShowReconnectButton  NullableBool `json:"ShowReconnectButton,omitempty"`
	ShowDisconnectButton NullableBool `json:"ShowDisconnectButton,omitempty"`
}

func (rawResponse *UserInterfaceWorkspaceControlRawResponseModel) ConvertToResponseModel() UserInterfaceWorkspaceControlResponseModel {
	response := UserInterfaceWorkspaceControlResponseModel{}
	response.Enabled = rawResponse.Enabled
	response.AutoReconnectAtLogon = rawResponse.AutoReconnectAtLogon
	response.ShowReconnectButton = rawResponse.ShowReconnectButton
	response.ShowDisconnectButton = rawResponse.ShowDisconnectButton

	if rawResponse.LogoffAction.IsSet() {
		switch *rawResponse.LogoffAction.Get() {
		case 0:
			response.LogoffAction = "Disconnect"
		case 1:
			response.LogoffAction = "Terminate"
		case 2:
			response.LogoffAction = "None"
		}
	}
	return response
}

type UserInterfaceWorkspaceControlResponseModel struct {
	Enabled              NullableBool `json:"Enabled,omitempty"`
	AutoReconnectAtLogon NullableBool `json:"AutoReconnectAtLogon,omitempty"`
	LogoffAction         string       `json:"LogoffAction,omitempty"` // Enum {Disconnect,Terminate,None}
	ShowReconnectButton  NullableBool `json:"ShowReconnectButton,omitempty"`
	ShowDisconnectButton NullableBool `json:"ShowDisconnectButton,omitempty"`
}

type UserInterfaceReceiverConfigurationResponseModel struct {
	Enabled     NullableBool   `json:"Enabled,omitempty"`
	DownloadUrl NullableString `json:"DownloadUrl,omitempty"`
}

type UserInterfaceAppShortcutsResponseModel struct {
	Enabled               NullableBool `json:"Enabled,omitempty"`
	AllowSessionReconnect NullableBool `json:"AllowSessionReconnect,omitempty"`
}

type UserInterfaceUIViewsRawResponseModel struct {
	ShowAppsView     NullableBool `json:"ShowAppsView,omitempty"`
	ShowDesktopsView NullableBool `json:"ShowDesktopsView,omitempty"`
	DefaultView      NullableInt  `json:"DefaultView,omitempty"` // Enum {Auto,Desktops,Apps}
}

func (rawResponse *UserInterfaceUIViewsRawResponseModel) ConvertToResponseModel() UserInterfaceUIViewsResponseModel {
	response := UserInterfaceUIViewsResponseModel{}
	response.ShowAppsView = rawResponse.ShowAppsView
	response.ShowDesktopsView = rawResponse.ShowDesktopsView

	if rawResponse.DefaultView.IsSet() {
		switch *rawResponse.DefaultView.Get() {
		case 0:
			response.DefaultView = "Auto"
		case 1:
			response.DefaultView = "Desktops"
		case 2:
			response.DefaultView = "Apps"
		}
	}

	return response
}

type UserInterfaceUIViewsResponseModel struct {
	ShowAppsView     NullableBool `json:"ShowAppsView,omitempty"`
	ShowDesktopsView NullableBool `json:"ShowDesktopsView,omitempty"`
	DefaultView      string       `json:"DefaultView,omitempty"` // Enum {Auto,Desktops,Apps}
}

type UserInterfaceProgressiveWebAppResponseModel struct {
	Enabled           NullableBool `json:"Enabled,omitempty"`
	ShowInstallPrompt NullableBool `json:"ShowInstallPrompt,omitempty"`
}

type UserInterfaceBrandingResponseModel struct {
	BackgroundColor NullableString `json:"BackgroundColor,omitempty"`
	TextColor       NullableString `json:"TextColor,omitempty"`
	LinkColor       NullableString `json:"LinkColor,omitempty"`
	HeaderLogo      NullableString `json:"HeaderLogo,omitempty"`
	LogonLogo       NullableString `json:"LogonLogo,omitempty"`
}

type GetSTFWebReceiverUserInterfaceRawResponseModel struct {
	AutoLaunchDesktop      NullableBool                                    `json:"AutoLaunchDesktop,omitempty"`
	MultiClickTimeout      NullableInt                                     `json:"MultiClickTimeout,omitempty"`
	CategoryViewCollapsed  NullableBool                                    `json:"CategoryViewCollapsed,omitempty"`
	MoveAppToUncategorized NullableBool                                    `json:"MoveAppToUncategorized,omitempty"`
	EnableAppsFolderView   NullableBool                                    `json:"EnableAppsFolderView,omitempty"`
	WorkspaceControl       UserInterfaceWorkspaceControlRawResponseModel   `json:"WorkspaceControl,omitempty"`
	ReceiverConfiguration  UserInterfaceReceiverConfigurationResponseModel `json:"ReceiverConfiguration,omitempty"`
	AppShortcuts           UserInterfaceAppShortcutsResponseModel          `json:"AppShortcuts,omitempty"`
	UIViews                UserInterfaceUIViewsRawResponseModel            `json:"UIViews,omitempty"`
	ProgressiveWebApp      UserInterfaceProgressiveWebAppResponseModel     `json:"ProgressiveWebApp,omitempty"`
	Branding               UserInterfaceBrandingResponseModel              `json:"Branding,omitempty"`
	ShowActivityManager    NullableBool                                    `json:"ShowActivityManager,omitempty"`
	ShowFirstTimeUse       NullableBool                                    `json:"ShowFirstTimeUse,omitempty"`
	PreventIcaDownloads    NullableBool                                    `json:"PreventIcaDownloads,omitempty"`
}

func (rawResponse *GetSTFWebReceiverUserInterfaceRawResponseModel) ConvertToResponseModel() GetSTFWebReceiverUserInterfaceResponseModel {
	return GetSTFWebReceiverUserInterfaceResponseModel{
		AutoLaunchDesktop:      rawResponse.AutoLaunchDesktop,
		MultiClickTimeout:      rawResponse.MultiClickTimeout,
		CategoryViewCollapsed:  rawResponse.CategoryViewCollapsed,
		MoveAppToUncategorized: rawResponse.MoveAppToUncategorized,
		EnableAppsFolderView:   rawResponse.EnableAppsFolderView,
		WorkspaceControl:       rawResponse.WorkspaceControl.ConvertToResponseModel(),
		ReceiverConfiguration:  rawResponse.ReceiverConfiguration,
		AppShortcuts:           rawResponse.AppShortcuts,
		UIViews:                rawResponse.UIViews.ConvertToResponseModel(),
		ProgressiveWebApp:      rawResponse.ProgressiveWebApp,
		Branding:               rawResponse.Branding,
		ShowActivityManager:    rawResponse.ShowActivityManager,
		ShowFirstTimeUse:       rawResponse.ShowFirstTimeUse,
		PreventIcaDownloads:    rawResponse.PreventIcaDownloads,
	}
}

type GetSTFWebReceiverUserInterfaceResponseModel struct {
	AutoLaunchDesktop      NullableBool                                    `json:"AutoLaunchDesktop,omitempty"`
	MultiClickTimeout      NullableInt                                     `json:"MultiClickTimeout,omitempty"`
	CategoryViewCollapsed  NullableBool                                    `json:"CategoryViewCollapsed,omitempty"`
	MoveAppToUncategorized NullableBool                                    `json:"MoveAppToUncategorized,omitempty"`
	EnableAppsFolderView   NullableBool                                    `json:"EnableAppsFolderView,omitempty"`
	WorkspaceControl       UserInterfaceWorkspaceControlResponseModel      `json:"WorkspaceControl,omitempty"`
	ReceiverConfiguration  UserInterfaceReceiverConfigurationResponseModel `json:"ReceiverConfiguration,omitempty"`
	AppShortcuts           UserInterfaceAppShortcutsResponseModel          `json:"AppShortcuts,omitempty"`
	UIViews                UserInterfaceUIViewsResponseModel               `json:"UIViews,omitempty"`
	ProgressiveWebApp      UserInterfaceProgressiveWebAppResponseModel     `json:"ProgressiveWebApp,omitempty"`
	Branding               UserInterfaceBrandingResponseModel              `json:"Branding,omitempty"`
	ShowActivityManager    NullableBool                                    `json:"ShowActivityManager,omitempty"`
	ShowFirstTimeUse       NullableBool                                    `json:"ShowFirstTimeUse,omitempty"`
	PreventIcaDownloads    NullableBool                                    `json:"PreventIcaDownloads,omitempty"`
}
