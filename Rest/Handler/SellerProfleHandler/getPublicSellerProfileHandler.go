package sellerproflehandler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *SellerProfileHandler) GetPublicSellerProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sellerIDStr := vars["id"]

	if sellerIDStr == "" {
		http.Error(w, `{"error":"seller ID is required"}`, http.StatusBadRequest)
		return
	}

	sellerID, err := strconv.ParseUint(sellerIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid seller ID"}`, http.StatusBadRequest)
		return
	}
	profile, err := h.service.GetPublicSellerProfile(uint(sellerID))
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	util.SendData(w, profile, http.StatusOK)

}
