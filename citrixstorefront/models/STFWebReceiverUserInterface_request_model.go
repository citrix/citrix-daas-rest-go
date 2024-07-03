package models

var _ MappedNullable = &SetSTFWebReceiverUserInterfaceRequestModel{}

type SetSTFWebReceiverUserInterfaceRequestModel struct {
	AutoLaunchDesktop                    NullableBool   `json:"AutoLaunchDesktop,omitempty"`
	MultiClickTimeout                    NullableInt    `json:"MultiClickTimeout,omitempty"`
	EnableAppsFolderView                 NullableBool   `json:"EnableAppsFolderView,omitempty"`
	ShowAppsView                         NullableBool   `json:"ShowAppsView,omitempty"`
	ShowDesktopsView                     NullableBool   `json:"ShowDesktopsView,omitempty"`
	DefaultView                          NullableString `json:"DefaultView,omitempty"` // Enum {Auto,Desktops,Apps}
	WorkspaceControlEnabled              NullableBool   `json:"WorkspaceControlEnabled,omitempty"`
	WorkspaceControlAutoReconnectAtLogon NullableBool   `json:"WorkspaceControlAutoReconnectAtLogon,omitempty"`
	WorkspaceControlLogoffAction         NullableString `json:"WorkspaceControlLogoffAction,omitempty"` // Enum {Disconnect,Terminate,None}
	WorkspaceControlShowReconnectButton  NullableBool   `json:"WorkspaceControlShowReconnectButton,omitempty"`
	WorkspaceControlShowDisconnectButton NullableBool   `json:"WorkspaceControlShowDisconnectButton,omitempty"`
	ReceiverConfigurationEnabled         NullableBool   `json:"ReceiverConfigurationEnabled,omitempty"`
	AppShortcutsEnabled                  NullableBool   `json:"AppShortcutsEnabled,omitempty"`
	AppShortcutsAllowSessionReconnect    NullableBool   `json:"AppShortcutsAllowSessionReconnect,omitempty"`
	CategoryViewCollapsed                NullableBool   `json:"CategoryViewCollapsed,omitempty"`
	MoveAppToUncategorized               NullableBool   `json:"MoveAppToUncategorized,omitempty"`
	ProgressiveWebAppEnabled             NullableBool   `json:"ProgressiveWebAppEnabled,omitempty"`
	ShowProgressiveWebAppInstallPrompt   NullableBool   `json:"ShowProgressiveWebAppInstallPrompt,omitempty"`
	ShowActivityManager                  NullableBool   `json:"ShowActivityManager,omitempty"`
	ShowFirstTimeUse                     NullableBool   `json:"ShowFirstTimeUse,omitempty"`
	PreventIcaDownloads                  NullableBool   `json:"PreventIcaDownloads,omitempty"`
}

// ToMap implements MappedNullable.
func (o SetSTFWebReceiverUserInterfaceRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoLaunchDesktop.IsSet() {
		toSerialize["AutoLaunchDesktop"] = o.AutoLaunchDesktop.Get()
	}
	if o.MultiClickTimeout.IsSet() {
		toSerialize["MultiClickTimeout"] = o.MultiClickTimeout.Get()
	}
	if o.EnableAppsFolderView.IsSet() {
		toSerialize["EnableAppsFolderView"] = o.EnableAppsFolderView.Get()
	}
	if o.ShowAppsView.IsSet() {
		toSerialize["ShowAppsView"] = o.ShowAppsView.Get()
	}
	if o.ShowDesktopsView.IsSet() {
		toSerialize["ShowDesktopsView"] = o.ShowDesktopsView.Get()
	}
	if o.DefaultView.IsSet() {
		toSerialize["DefaultView"] = o.DefaultView.Get()
	}
	if o.WorkspaceControlEnabled.IsSet() {
		toSerialize["WorkspaceControlEnabled"] = o.WorkspaceControlEnabled.Get()
	}
	if o.WorkspaceControlAutoReconnectAtLogon.IsSet() {
		toSerialize["WorkspaceControlAutoReconnectAtLogon"] = o.WorkspaceControlAutoReconnectAtLogon.Get()
	}
	if o.WorkspaceControlLogoffAction.IsSet() {
		toSerialize["WorkspaceControlLogoffAction"] = o.WorkspaceControlLogoffAction.Get()
	}
	if o.WorkspaceControlShowReconnectButton.IsSet() {
		toSerialize["WorkspaceControlShowReconnectButton"] = o.WorkspaceControlShowReconnectButton.Get()
	}
	if o.WorkspaceControlShowDisconnectButton.IsSet() {
		toSerialize["WorkspaceControlShowDisconnectButton"] = o.WorkspaceControlShowDisconnectButton.Get()
	}
	if o.ReceiverConfigurationEnabled.IsSet() {
		toSerialize["ReceiverConfigurationEnabled"] = o.ReceiverConfigurationEnabled.Get()
	}
	if o.AppShortcutsEnabled.IsSet() {
		toSerialize["AppShortcutsEnabled"] = o.AppShortcutsEnabled.Get()
	}
	if o.AppShortcutsAllowSessionReconnect.IsSet() {
		toSerialize["AppShortcutsAllowSessionReconnect"] = o.AppShortcutsAllowSessionReconnect.Get()
	}
	if o.CategoryViewCollapsed.IsSet() {
		toSerialize["CategoryViewCollapsed"] = o.CategoryViewCollapsed.Get()
	}
	if o.MoveAppToUncategorized.IsSet() {
		toSerialize["MoveAppToUncategorized"] = o.MoveAppToUncategorized.Get()
	}
	if o.ProgressiveWebAppEnabled.IsSet() {
		toSerialize["ProgressiveWebAppEnabled"] = o.ProgressiveWebAppEnabled.Get()
	}
	if o.ShowProgressiveWebAppInstallPrompt.IsSet() {
		toSerialize["ShowProgressiveWebAppInstallPrompt"] = o.ShowProgressiveWebAppInstallPrompt.Get()
	}
	if o.ShowActivityManager.IsSet() {
		toSerialize["ShowActivityManager"] = o.ShowActivityManager.Get()
	}
	if o.ShowFirstTimeUse.IsSet() {
		toSerialize["ShowFirstTimeUse"] = o.ShowFirstTimeUse.Get()
	}
	if o.PreventIcaDownloads.IsSet() {
		toSerialize["PreventIcaDownloads"] = o.PreventIcaDownloads.Get()
	}
	return toSerialize, nil
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetAutoLaunchDesktop(v bool) {
	o.AutoLaunchDesktop.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetMultiClickTimeout(v int) {
	o.MultiClickTimeout.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetEnableAppsFolderView(v bool) {
	o.EnableAppsFolderView.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetShowAppsView(v bool) {
	o.ShowAppsView.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetShowDesktopsView(v bool) {
	o.ShowDesktopsView.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetDefaultView(v string) {
	o.DefaultView.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetWorkspaceControlEnabled(v bool) {
	o.WorkspaceControlEnabled.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetWorkspaceControlAutoReconnectAtLogon(v bool) {
	o.WorkspaceControlAutoReconnectAtLogon.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetWorkspaceControlLogoffAction(v string) {
	o.WorkspaceControlLogoffAction.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetWorkspaceControlShowReconnectButton(v bool) {
	o.WorkspaceControlShowReconnectButton.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetWorkspaceControlShowDisconnectButton(v bool) {
	o.WorkspaceControlShowDisconnectButton.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetReceiverConfigurationEnabled(v bool) {
	o.ReceiverConfigurationEnabled.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetAppShortcutsEnabled(v bool) {
	o.AppShortcutsEnabled.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetAppShortcutsAllowSessionReconnect(v bool) {
	o.AppShortcutsAllowSessionReconnect.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetCategoryViewCollapsed(v bool) {
	o.CategoryViewCollapsed.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetMoveAppToUncategorized(v bool) {
	o.MoveAppToUncategorized.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetProgressiveWebAppEnabled(v bool) {
	o.ProgressiveWebAppEnabled.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetShowProgressiveWebAppInstallPrompt(v bool) {
	o.ShowProgressiveWebAppInstallPrompt.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetShowActivityManager(v bool) {
	o.ShowActivityManager.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetShowFirstTimeUse(v bool) {
	o.ShowFirstTimeUse.Set(&v)
}

func (o *SetSTFWebReceiverUserInterfaceRequestModel) SetPreventIcaDownloads(v bool) {
	o.PreventIcaDownloads.Set(&v)
}
