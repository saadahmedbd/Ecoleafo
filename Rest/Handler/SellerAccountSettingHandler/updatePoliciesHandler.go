package selleraccountsettinghandler

import (
	"encoding/json"
	"net/http"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// UPDATE POLICIES
// PUT /api/seller/store/policies
// ==========================================

func (h *Selleraccountsettinghandler) UpdatePolicies(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	var req selleraccountsetting.UpdatePoliciesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.UpdatePolicies(sellerID, &req); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{
		"message": "Store policies updated successfully",
	},
		"",
	)
}
