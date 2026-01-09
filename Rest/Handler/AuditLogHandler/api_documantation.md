
-- Add indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_audit_logs_actor ON audit_logs(actor_id, actor_type);
CREATE INDEX IF NOT EXISTS idx_audit_logs_action ON audit_logs(action);
CREATE INDEX IF NOT EXISTS idx_audit_logs_action_group ON audit_logs(action_group);
CREATE INDEX IF NOT EXISTS idx_audit_logs_entity ON audit_logs(entity_type, entity_id);
CREATE INDEX IF NOT EXISTS idx_audit_logs_created ON audit_logs(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_audit_logs_severity ON audit_logs(severity);
CREATE INDEX IF NOT EXISTS idx_audit_logs_category ON audit_logs(category);
CREATE INDEX IF NOT EXISTS idx_audit_logs_status ON audit_logs(status);
CREATE INDEX IF NOT EXISTS idx_audit_logs_ip ON audit_logs(ip_address);

-- Composite indexes for common queries
CREATE INDEX IF NOT EXISTS idx_audit_logs_actor_created 
    ON audit_logs(actor_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_audit_logs_entity_created 
    ON audit_logs(entity_type, entity_id, created_at DESC);
*/

// ============================================================================
// STEP 10: API DOCUMENTATION
// ============================================================================

/*
=============================================================================
ACTIVITY LOG ENHANCEMENT API DOCUMENTATION
=============================================================================

BASE URL: http://localhost:3000
AUTHENTICATION: JWT Bearer token required (Admin/Super Admin)
HEADERS: Authorization: Bearer {admin_token}

=============================================================================
ENDPOINTS
=============================================================================

1. GET ALL AUDIT LOGS
   GET /api/admin/logs?page=1&limit=20&action=login&actor_type=admin
   
   Query Parameters:
   - page: Page number (default: 1)
   - limit: Items per page (default: 20, max: 100)
   - actor_id: Filter by actor ID
   - actor_type: admin, seller, buyer, system
   - action: login, logout, product_approve, etc.
   - action_group: authentication, user_management, order_management, etc.
   - entity_type: user, seller, product, order, review, etc.
   - entity_id: Filter by entity ID
   - status: success, failed, warning
   - severity: info, warning, critical
   - category: security, business, system, audit
   - start_date: YYYY-MM-DD
   - end_date: YYYY-MM-DD
   - search: Search in description, actor_name, entity_name
   - ip: Filter by IP address
   
   Response:
   {
     "status": "success",
     "data": [
       {
         "id": 1,
         "actor_id": 5,
         "actor_type": "admin",
         "actor_name": "John Admin",
         "action": "seller_approve",
         "action_group": "seller_management",
         "action_label": "Seller Approved",
         "description": "Approved seller: Tech Store",
         "entity_type": "seller",
         "entity_id": 10,
         "entity_name": "Tech Store",
         "ip_address": "192.168.1.1",
         "status": "success",
         "severity": "info",
         "category": "business",
         "created_at": "2025-11-19T10:30:00Z",
         "time_ago": "2 hours ago"
       }
     ],
     "pagination": {
       "page": 1,
       "limit": 20,
       "total": 500
     }
   }

2. GET LOG BY ID
   GET /api/admin/logs/detail?id=5
   
   Response: Single log with full details including old/new values and changes

3. GET ACTOR LOGS
   GET /api/admin/logs/actor?actor_id=5&actor_type=admin&page=1&limit=20
   
   Returns: All actions performed by specific actor

4. GET ENTITY HISTORY
   GET /api/admin/logs/entity?entity_type=product&entity_id=10&page=1&limit=20
   
   Returns: Complete audit trail for specific entity

5. GET SECURITY LOGS
   GET /api/admin/logs/security?page=1&limit=20
   
   Returns: Login failures, password changes, critical severity events

6. GET RECENT ACTIVITY
   GET /api/admin/logs/recent?limit=20
   
   Returns: Latest activity across the platform

7. GET STATISTICS
   GET /api/admin/logs/stats
   
   Response:
   {
     "status": "success",
     "data": {
       "total_logs": 5000,
       "today_logs": 150,
       "this_week_logs": 800,
       "success_count": 4800,
       "failed_count": 200,
       "by_action_group": {
         "authentication": 500,
         "seller_management": 300,
         "product_management": 1200
       },
       "by_severity": {
         "info": 4500,
         "warning": 400,
         "critical": 100
       },
       "by_actor_type": {
         "admin": 2000,
         "seller": 2500,
         "buyer": 500
       },
       "recent_activity": [...],
       "top_actors": [
         {
           "actor_id": 5,
           "actor_name": "John Admin",
           "actor_type": "admin",
           "action_count": 250
         }
       ],
       "security_alerts": 15
     }
   }

8. GET ACTIVITY TIMELINE
   GET /api/admin/logs/timeline?start_date=2025-11-01&end_date=2025-11-30&group_by=day
   
   Returns: Activity counts grouped by time period

9. EXPORT LOGS
   POST /api/admin/logs/export
   
   Request:
   {
     "format": "csv",
     "start_date": "2025-11-01",
     "end_date": "2025-11-30",
     "filters": {
       "actor_type": "admin",
       "action_group": "seller_management"
     }
   }
   
   Response: File download (CSV or JSON)

10. CLEANUP OLD LOGS (Super Admin)
    POST /api/admin/logs/cleanup?days=90
    
    Returns: Confirmation message

=============================================================================
ACTION TYPES REFERENCE
=============================================================================

Authentication:
- login, logout, login_failed, password_change, password_reset

User Management:
- user_create, user_update, user_delete, user_activate, user_deactivate, user_suspend

Seller Management:
- seller_approve, seller_reject, seller_suspend, seller_reactivate

Product Management:
- product_create, product_update, product_delete, product_approve, product_reject

Order Management:
- order_create, order_update, order_cancel, order_refund, order_status_change

Review Management:
- review_approve, review_reject, review_delete

Payment Management:
- payout_request, payout_approve, payout_reject, commission_update

Admin Management:
- admin_invite, admin_create, permission_update

System:
- settings_update, data_export, system_backup, bulk_operation

=============================================================================
INTEGRATION EXAMPLES
=============================================================================

1. In SellerService - Log seller approval:

func (s *SellerService) ApproveSeller(sellerID uint, adminID uint, adminName string) error {
    seller, err := s.sellerRepo.GetByID(sellerID)
    if err != nil {
        return err
    }
    
    // Approve seller
    seller.Status = "approved"
    s.sellerRepo.Update(seller)
    
    // Log action
    s.auditHelper.LogSellerApproval(adminID, adminName, seller)
    
    return nil
}

2. In AuthService - Log login attempts:

func (s *AuthService) Login(email, password, ip, userAgent string) (*Token, error) {
    user, err := s.userRepo.GetByEmail(email)
    
    if err != nil || !user.CheckPassword(password) {
        // Log failed attempt
        s.auditHelper.LogLoginAttempt(nil, email, ip, userAgent, false)
        return nil, errors.New("invalid credentials")
    }
    
    // Log successful login
    s.auditHelper.LogLoginAttempt(&user.ID, email, ip, userAgent, true)
    
    return s.generateToken(user)
}

3. In ProductService - Log with data changes:

func (s *ProductService) UpdateProduct(productID uint, updates *UpdateProductDTO, adminID uint) error {
    oldProduct, _ := s.productRepo.GetByID(productID)
    
    // Apply updates
    s.productRepo.Update(productID, updates)
    
    newProduct, _ := s.productRepo.GetByID(productID)
    
    // Log with change tracking
    s.auditService.LogDataChange(
        adminID,
        "admin",
        ActionProductUpdate,
        "product",
        productID,
        oldProduct,
        newProduct,
    )
    
    return nil
}

=============================================================================
TESTING EXAMPLES (cURL)
=============================================================================

# Get all logs
curl -X GET "http://localhost:8080/api/admin/logs?page=1&limit=20" \
  -H "Authorization: Bearer ADMIN_TOKEN"

# Get security logs
curl -X GET "http://localhost:8080/api/admin/logs/security?page=1&limit=20" \
  -H "Authorization: Bearer ADMIN_TOKEN"

# Get entity history
curl -X GET "http://localhost:8080/api/admin/logs/entity?entity_type=seller&entity_id=5" \
  -H "Authorization: Bearer ADMIN_TOKEN"

# Get statistics
curl -X GET "http://localhost:8080/api/admin/logs/stats" \
  -H "Authorization: Bearer ADMIN_TOKEN"

# Export logs
curl -X POST "http://localhost:8080/api/admin/logs/export" \
  -H "Authorization: Bearer ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"format": "csv", "start_date": "2025-11-01", "end_date": "2025-11-30"}' \
  --output audit_logs.csv

=============================================================================
BEST PRACTICES
=============================================================================

1. Always log security-sensitive actions (login, password changes)
2. Include old/new values for data modifications
3. Use appropriate severity levels
4. Don't log sensitive data (passwords, tokens)
5. Implement log retention policy (cleanup old logs)
6. Monitor security alerts regularly
7. Set up alerts for critical severity events
8. Export logs periodically for compliance
