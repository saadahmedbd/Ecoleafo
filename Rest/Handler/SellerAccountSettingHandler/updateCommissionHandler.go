package selleraccountsettinghandler

// import (
// 	"encoding/json"
// 	"net/http"

// 	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
// 	util "github.com/saadahmedbd/Treestore/Util"
// 	"github.com/saadahmedbd/Treestore/constants"
// )

// func (h *SellerAccountSettingHandler) UpdateCommission(w http.ResponseWriter, r *http.Request) {
// 	userIDVal := r.Context().Value(constants.ContextKeyUserID)
// 	if userIDVal == nil {
// 		util.SendError(w, "unauthorized", http.StatusUnauthorized)
// 		return
// 	}

// 	userID, ok := userIDVal.(uint)
// 	if !ok {
// 		util.SendError(w, "invalid user_id type", http.StatusInternalServerError)
// 		return
// 	}

// 	seller, err := h.service.GetSellerByUserID(userID)
// 	if err != nil {
// 		util.SendError(w, "seller not found", http.StatusNotFound)
// 		return
// 	}

// 	var req selleraccountsetting.UpdateCommissionRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		util.SendError(w, "invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	if err := h.service.UpdateCommission(seller.ID, req.Commission); err != nil {
// 		util.SendError(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	util.RespondJSON(w, http.StatusOK, nil, "Commission updated successfully")
// }
