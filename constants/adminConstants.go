package constants

// ============================================================================
// ROLES & USER TYPES
// ============================================================================

// Admin Roles
const (
	RoleSuperAdmin = "super_admin"
	RoleAdmin      = "admin"
	RoleManager    = "manager"
	RoleSupport    = "support"
)

// User Types (also used as ActorType in AuditLog)
const (
	UserTypeBuyer  = "buyer"
	UserTypeSeller = "seller"
	UserTypeAdmin  = "admin"  // Added for audit logs
	UserTypeSystem = "system" // Added for audit logs
	UserTypeAPI    = "api"    // Added for audit logs
)

// ============================================================================
// STATUSES
// ============================================================================

// General Status
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

// Audit Log Status
const (
	AuditStatusSuccess = "success"
	AuditStatusFailed  = "failed"
	AuditStatusWarning = "warning"
)

// ============================================================================
// AUDIT LOG - ACTIONS (Merged with Activity Types)
// ============================================================================

// Authentication Actions
const (
	ActionLogin          = "login"
	ActionLogout         = "logout"
	ActionLoginFailed    = "login_failed"
	ActionPasswordChange = "password_change"
	ActionPasswordReset  = "password_reset"
	Action2FAEnabled     = "2fa_enabled"
	Action2FADisabled    = "2fa_disabled"
)

// CRUD Actions (Generic)
const (
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionView   = "view"
)

// Approval/Status Change Actions
const (
	ActionApprove    = "approve"
	ActionReject     = "reject"
	ActionSuspend    = "suspend"
	ActionActivate   = "activate"
	ActionDeactivate = "deactivate"
)

// User Management Actions
const (
	ActionUserCreate     = "user_create"
	ActionUserUpdate     = "user_update"
	ActionUserDelete     = "user_delete"
	ActionUserActivate   = "user_activate"
	ActionUserDeactivate = "user_deactivate"
	ActionUserSuspend    = "user_suspend"
)

// Seller Actions
const (
	ActionSellerApprove    = "seller_approve"
	ActionSellerReject     = "seller_reject"
	ActionSellerSuspend    = "seller_suspend"
	ActionSellerReactivate = "seller_reactivate"
)

// Product Actions
const (
	ActionProductCreate  = "product_create"
	ActionProductUpdate  = "product_update"
	ActionProductDelete  = "product_delete"
	ActionProductApprove = "product_approve"
	ActionProductReject  = "product_reject"
)

// Order Actions
const (
	ActionOrderCreate       = "order_create"
	ActionOrderUpdate       = "order_update"
	ActionOrderCancel       = "order_cancel"
	ActionOrderRefund       = "order_refund"
	ActionOrderStatusChange = "order_status_change"
)

// Review Actions
const (
	ActionReviewApprove = "review_approve"
	ActionReviewReject  = "review_reject"
	ActionReviewDelete  = "review_delete"
)

// Payment Actions
const (
	ActionPayoutRequest    = "payout_request"
	ActionPayoutApprove    = "payout_approve"
	ActionPayoutReject     = "payout_reject"
	ActionCommissionUpdate = "commission_update"
)

// Admin Actions
const (
	ActionAdminInvite      = "admin_invite"
	ActionAdminCreate      = "admin_create"
	ActionPermissionUpdate = "permission_update"
	ActionSettingsUpdate   = "settings_update"
)

// System Actions
const (
	ActionSystemBackup  = "system_backup"
	ActionSystemRestore = "system_restore"
	ActionDataExport    = "data_export"
	ActionBulkOperation = "bulk_operation"
)

// ============================================================================
// AUDIT LOG - ACTION GROUPS
// ============================================================================

const (
	GroupAuth     = "authentication"
	GroupUser     = "user_management"
	GroupSeller   = "seller_management"
	GroupProduct  = "product_management"
	GroupOrder    = "order_management"
	GroupReview   = "review_management"
	GroupPayment  = "payment_management"
	GroupAdmin    = "admin_management"
	GroupSystem   = "system"
	GroupSecurity = "security"
)

// ============================================================================
// AUDIT LOG - SEVERITY & CATEGORIES
// ============================================================================

// Severity Levels
const (
	SeverityInfo     = "info"
	SeverityWarning  = "warning"
	SeverityCritical = "critical"
)

// Audit Categories
const (
	CategorySecurity = "security"
	CategoryBusiness = "business"
	CategorySystem   = "system"
	CategoryAudit    = "audit"
)

// ============================================================================
// AUDIT LOG - ENTITY TYPES
// ============================================================================

const (
	EntityTypeUser     = "user"
	EntityTypeSeller   = "seller"
	EntityTypeBuyer    = "buyer"
	EntityTypeProduct  = "product"
	EntityTypeOrder    = "order"
	EntityTypeReview   = "review"
	EntityTypePayment  = "payment"
	EntityTypePayout   = "payout"
	EntityTypeAdmin    = "admin"
	EntityTypeSettings = "settings"
)

// ============================================================================
// PERMISSIONS
// ============================================================================

const (
	PermissionManageUsers    = "can_manage_users"
	PermissionManageProducts = "can_manage_products"
	PermissionManageOrders   = "can_manage_orders"
	PermissionViewReports    = "can_view_reports"
	PermissionManageSettings = "can_manage_settings"
	PermissionManageSellers  = "can_manage_sellers"
	PermissionManagePayments = "can_manage_payments"
)

// ============================================================================
// CONTEXT KEYS
// ============================================================================

const (
	ContextKeyUserID  = "user_id"
	ContextKeyAdminID = "admin_id"
	ContextKeyRole    = "role"
	ContextKeyEmail   = "email"
	ContextKeyRoleID  = "role_id"
)

// ============================================================================
// CONFIGURATION
// ============================================================================

// Invitation
const (
	InvitationExpiryDays = 7
	TokenLength          = 64
)
