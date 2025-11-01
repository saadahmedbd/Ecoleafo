package selleraccountsettinghandler

import (
	"encoding/json"
	"net/http"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// UPDATE ACCOUNT
// PUT /api/seller/account
// ==========================================

func (h *Selleraccountsettinghandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	var req selleraccountsetting.UpdateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.UpdateAccount(sellerID, &req); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Account updated successfully",
	},
		"update account successfully",
	)
}
