package selleraccountsettinghandler

import (
	"encoding/json"
	"net/http"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// UPDATE NOTIFICATION PREFERENCES
// PUT /api/seller/notifications/preferences
// ==========================================

func (h *Selleraccountsettinghandler) UpdateNotificationPreferences(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	var req selleraccountsetting.NotificationPreferences
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.UpdateNotificationPreferences(sellerID, &req); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{
		"message": "Notification preferences updated successfully",
	}, "")
}
