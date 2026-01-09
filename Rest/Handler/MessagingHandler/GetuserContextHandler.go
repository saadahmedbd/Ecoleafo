package messaginghandler

import (
	"net/http"

	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// Helper to extract user context from request
func (h *MessagingHandler) getUserContext(r *http.Request) (*messagingdto.UserContext, error) {
	// Extract user_id (RegUser.ID)
	userID, ok := r.Context().Value(constants.ContextKeyUserID).(uint)
	if !ok {
		return nil, util.NewAppError("user_id not found in context", http.StatusUnauthorized)
	}

	// Extract role
	roleVal := r.Context().Value(constants.ContextKeyRole)
	var userType string
	switch v := roleVal.(type) {
	case string:
		userType = v
	case []interface{}:
		if len(v) > 0 {
			if s, ok := v[0].(string); ok {
				userType = s
			}
		}
	}

	if userType == "" {
		return nil, util.NewAppError("user role not found in context", http.StatusUnauthorized)
	}

	// Extract role_id (Buyer.ID, User.ID for seller, or Admin.ID) ⭐ SIMPLIFIED
	roleID, ok := r.Context().Value(constants.ContextKeyRoleID).(uint)
	if !ok || roleID == 0 {
		return nil, util.NewAppError("role_id not found in context", http.StatusUnauthorized)
	}

	return &messagingdto.UserContext{
		UserID:   userID,   // RegUser.ID
		UserType: userType, // "buyer", "seller", "admin"
		RoleID:   roleID,   // Buyer.ID, User.ID (seller), Admin.ID
	}, nil
}
