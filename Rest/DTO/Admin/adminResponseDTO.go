package admin

import "time"

type AdminResponse struct {
	ID                uint       `json:"id"`
	UserID            uint       `json:"user_id"`
	Email             string     `json:"email"`
	Role              string     `json:"role"`
	FullName          string     `json:"full_name"`
	Phone             string     `json:"phone"`
	Department        string     `json:"department"`
	CanManageUsers    bool       `json:"can_manage_users"`
	CanManageProducts bool       `json:"can_manage_products"`
	CanManageOrders   bool       `json:"can_manage_orders"`
	CanViewReports    bool       `json:"can_view_reports"`
	CanManageSettings bool       `json:"can_manage_settings"`
	IsActive          bool       `json:"is_active"`
	LastLogin         *time.Time `json:"last_login"`
	CreatedAt         time.Time  `json:"created_at"`
}
