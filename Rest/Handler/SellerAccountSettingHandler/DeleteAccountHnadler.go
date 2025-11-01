package selleraccountsettinghandler

import (
	"encoding/json"
	"net/http"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// DELETE ACCOUNT
// DELETE /api/seller/account/delete
// ==========================================

func (h *Selleraccountsettinghandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	var req selleraccountsetting.DeleteAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := h.service.DeleteAccount(sellerID, req.Password)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, response, "")
}
