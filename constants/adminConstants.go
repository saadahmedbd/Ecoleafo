package constants

// Admin Roles
const (
	RoleSuperAdmin = "super_admin"
	RoleAdmin      = "admin"
	RoleManager    = "manager"
	RoleSupport    = "support"
)

// User Types
const (
	UserTypeBuyer  = "buyer"
	UserTypeSeller = "seller"
)

// User Status
const (
	StatusActive    = "active"
	StatusInactive  = "inactive"
	StatusSuspended = "suspended"
	StatusPending   = "pending"
)

// Order Status
const (
	OrderStatusPending    = "pending"
	OrderStatusConfirmed  = "confirmed"
	OrderStatusProcessing = "processing"
	OrderStatusShipped    = "shipped"
	OrderStatusDelivered  = "delivered"
	OrderStatusCancelled  = "cancelled"
	OrderStatusRefunded   = "refunded"
)

// Payment Status
const (
	PaymentStatusPending   = "pending"
	PaymentStatusCompleted = "completed"
	PaymentStatusFailed    = "failed"
	PaymentStatusRefunded  = "refunded"
)

// Product Status
const (
	ProductStatusDraft    = "draft"
	ProductStatusPending  = "pending"
	ProductStatusApproved = "approved"
	ProductStatusRejected = "rejected"
	ProductStatusDeleted  = "deleted"
)

// Activity Types
const (
	ActivityLogin      = "login"
	ActivityLogout     = "logout"
	ActivityCreate     = "create"
	ActivityUpdate     = "update"
	ActivityDelete     = "delete"
	ActivityApprove    = "approve"
	ActivityReject     = "reject"
	ActivitySuspend    = "suspend"
	ActivityActivate   = "activate"
	ActivityDeactivate = "deactivate"
)

// Permissions
const (
	PermissionManageUsers    = "can_manage_users"
	PermissionManageProducts = "can_manage_products"
	PermissionManageOrders   = "can_manage_orders"
	PermissionViewReports    = "can_view_reports"
	PermissionManageSettings = "can_manage_settings"
	PermissionManageSellers  = "can_manage_sellers"
	PermissionManagePayments = "can_manage_payments"
)

// Context Keys
const (
	ContextKeyUserID  = "user_id"
	ContextKeyAdminID = "admin_id"
	ContextKeyRole    = "role"
	ContextKeyEmail   = "email"
)

// Invitation
const (
	InvitationExpiryDays = 7
	TokenLength          = 64
)
