package selleraccounthandler

import (
	"encoding/json"
	"net/http"

	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
	util "github.com/saadahmedbd/Treestore/Util"
)

// // POST /api/seller/register
func (h *SellerRegistrationHandler) RegisterSeller(w http.ResponseWriter, r *http.Request) {
	var req selleraccount.SellerRegistationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if !req.AgreeToTerms {
		http.Error(w, "you must agree to the terms and conditions", http.StatusBadRequest)
		return
	}
	response, err := h.service.RegisterSeller(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, response, http.StatusCreated)
}
