package selleraccountsettinghandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// GET NOTIFICATION PREFERENCES
// GET /api/seller/notifications/preferences
// ==========================================

func (h *Selleraccountsettinghandler) GetNotificationPreferences(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	prefs, err := h.service.GetNotificationPreferences(sellerID)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, prefs, "")
}
