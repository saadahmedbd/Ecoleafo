package selleraccountsettinghandler

import (
	"encoding/json"
	"net/http"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// UPDATE STORE
// PUT /api/seller/store
// ==========================================

func (h *Selleraccountsettinghandler) UpdateStore(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	var req selleraccountsetting.UpdateStoreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.UpdateStore(sellerID, &req); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{
		"message": "Store information updated successfully",
	},
		"",
	)
}
