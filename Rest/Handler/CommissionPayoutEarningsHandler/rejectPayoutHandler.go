package commissionpayoutearningshandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	commissionearningpayoutdto "github.com/saadahmedbd/Treestore/Rest/DTO/CommissionEarningPayoutDTO"
	util "github.com/saadahmedbd/Treestore/Util"
)

// POST /api/admin/payouts/reject?id=5
func (h *CommissionHandler) RejectPayout(w http.ResponseWriter, r *http.Request) {
	payoutIDStr := r.URL.Query().Get("id")
	if payoutIDStr == "" {
		util.RespondError(w, http.StatusBadRequest, "payout id is required")
		return
	}

	payoutID, err := strconv.ParseUint(payoutIDStr, 10, 64)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid payout id")
		return
	}

	var req commissionearningpayoutdto.RejectPayoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Get admin ID from context
	adminID := r.Context().Value("user_id").(uint)

	if err := h.commissionService.RejectPayout(uint(payoutID), &req, adminID); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Payout rejected successfully")
}
