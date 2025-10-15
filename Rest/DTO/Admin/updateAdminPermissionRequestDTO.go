package admin

type UpdateAdminPermissionRequest struct {
	CanManageUsers    *bool `json:"can_manage_users"`
	CanManageProducts *bool `json:"can_manage_products"`
	CanManageOrders   *bool `json:"can_manage_orders"`
	CanViewReports    *bool `json:"can_view_reports"`
	CanManageSettings *bool `json:"can_manage_settings"`
}
