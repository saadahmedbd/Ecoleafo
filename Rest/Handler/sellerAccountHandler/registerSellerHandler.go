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
		util.SendError(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	
	if req.Email == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" || req.StoreName == "" || req.Phone == "" {
		util.SendError(w, "All fields are required", http.StatusBadRequest)
		return
	}
	
	if !req.AgreeToTerms {
		util.SendError(w, "You must agree to the terms and conditions", http.StatusBadRequest)
		return
	}
	
	if req.Commission < 15 && req.Commission != 0 {
		util.SendError(w, "Commission must be at least 15%", http.StatusBadRequest)
		return
	}
	
	response, err := h.service.RegisterSeller(&req)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	util.SendData(w, response, http.StatusCreated)
}
