package adminmangementhandler

import (
	"log"
	"net/http"
	"strconv"

	adminmangementservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminMangementService"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

type BuyerHandler struct {
	buyerService *adminmangementservice.BuyerService
}

func NewBuyerHandler(buyerService *adminmangementservice.BuyerService) *BuyerHandler {
	return &BuyerHandler{
		buyerService: buyerService,
	}
}

// GetAllBuyers handles GET /api/Buyers
func (h *BuyerHandler) GetAllBuyers(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	status := r.URL.Query().Get("status")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// Get Buyers
	Buyers, total, err := h.buyerService.GetAllBuyers(page, limit, status)
	if err != nil {
		util.SendError(w, "Failed to fetch Buyers", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, Buyers, page, limit, total)
}

// GetBuyerByID handles GET /api/Buyers/{id}
func (h *BuyerHandler) GetBuyerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid Buyer ID", http.StatusBadRequest)
		return
	}

	Buyer, err := h.buyerService.GetBuyerByID(uint(id))
	if err != nil {
		util.SendError(w, "Buyer not found", http.StatusNotFound)
		return
	}
	util.RespondJSON(w, http.StatusOK, Buyer, "Buyer retrieved successfully")

}

// SearchBuyers handles GET /api/Buyers/search
func (h *BuyerHandler) SearchBuyers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		util.SendError(w, "Search query is required", http.StatusBadRequest)
		return
	}

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	Buyers, total, err := h.buyerService.SearchBuyers(query, page, limit)
	if err != nil {
		util.SendError(w, "Failed to search Buyers", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, Buyers, page, limit, total)
}

// ActivateBuyer handles POST /api/Buyers/{id}/activate
func (h *BuyerHandler) ActivateBuyer(w http.ResponseWriter, r *http.Request) {
	// Get user_id and role from context
	userIDVal := r.Context().Value(constants.ContextKeyUserID)

	if userIDVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	adminID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get buyer ID from URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid buyer ID", http.StatusBadRequest)
		return
	}
	log.Println("buyer id", id)

	// Call service to activate buyer
	err = h.buyerService.ActivateBuyer(uint(id), adminID)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Buyer activated successfully")
}

// DeactivateBuyer handles POST /api/Buyers/{id}/deactivate
func (h *BuyerHandler) DeactivateBuyer(w http.ResponseWriter, r *http.Request) {
	// Get user_id and role from context
	userIDVal := r.Context().Value(constants.ContextKeyUserID)

	if userIDVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	adminID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get buyer ID from URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid buyer ID", http.StatusBadRequest)
		return
	}
	log.Println("buyer id", id)
	err = h.buyerService.DeactivateBuyer(uint(id), uint(adminID))
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Buyer deactivated successfully")
}

// SuspendBuyer handles POST /api/Buyers/{id}/suspend
func (h *BuyerHandler) SuspendBuyer(w http.ResponseWriter, r *http.Request) {
	// Get user_id and role from context
	userIDVal := r.Context().Value(constants.ContextKeyUserID)

	if userIDVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	adminID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get buyer ID from URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid buyer ID", http.StatusBadRequest)
		return
	}
	log.Println("buyer id", id)
	err = h.buyerService.SuspendBuyer(uint(id), uint(adminID))
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Buyer suspended successfully")
}

// GetBuyerStats handles GET /api/Buyers/stats
func (h *BuyerHandler) GetBuyerStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.buyerService.GetBuyerStats()
	if err != nil {
		util.SendError(w, "Failed to fetch statistics", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, stats, "Statistics retrieved successfully")
}
