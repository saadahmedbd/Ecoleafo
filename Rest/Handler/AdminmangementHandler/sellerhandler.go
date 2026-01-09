package adminmangementhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	adminmangementservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminMangementService"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

type SellerHandler struct {
	sellerService *adminmangementservice.SellerService
}

func NewSellerHandler(sellerService *adminmangementservice.SellerService) *SellerHandler {
	return &SellerHandler{
		sellerService: sellerService,
	}
}

// GetAllSellers handles GET /api/sellers
func (h *SellerHandler) GetAllSellers(w http.ResponseWriter, r *http.Request) {
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

	sellers, total, err := h.sellerService.GetAllSellers(page, limit, status)
	if err != nil {
		util.SendError(w, "Failed to fetch sellers", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, sellers, page, limit, total)
}

// GetSellerByID handles GET /api/sellers/{id}
func (h *SellerHandler) GetSellerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	seller, err := h.sellerService.GetSellerByID(uint(id))
	if err != nil {
		util.SendError(w, "Seller not found", http.StatusNotFound)
		return
	}

	util.RespondJSON(w, http.StatusOK, seller, "Seller retrieved successfully")
}

// GetPendingApprovals handles GET /api/sellers/pending
func (h *SellerHandler) GetPendingApprovals(w http.ResponseWriter, r *http.Request) {
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

	sellers, total, err := h.sellerService.GetPendingApprovals(page, limit)
	if err != nil {
		util.SendError(w, "Failed to fetch pending approvals", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, sellers, page, limit, total)
}

// ApproveSeller handles POST /api/sellers/{id}/approve
func (h *SellerHandler) ApproveSeller(w http.ResponseWriter, r *http.Request) {
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
	err = h.sellerService.ApproveSeller(uint(id), adminID)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Seller approved successfully")
}

// RejectSeller handles POST /api/sellers/{id}/reject
func (h *SellerHandler) RejectSeller(w http.ResponseWriter, r *http.Request) {
	// Get user_id and role from context
	userIDVal := r.Context().Value(constants.ContextKeyUserID)

	if userIDVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	adminUserID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get buyer ID from URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	var reqBody struct {
		Reason string `json:"reason"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		util.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if reqBody.Reason == "" {
		util.SendError(w, "Rejection reason is required", http.StatusBadRequest)
		return
	}

	err = h.sellerService.RejectSeller(uint(id), reqBody.Reason, adminUserID)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Seller rejected successfully")
}

// SuspendSeller handles POST /api/sellers/{id}/suspend
func (h *SellerHandler) SuspendSeller(w http.ResponseWriter, r *http.Request) {
	// Get user_id and role from context
	userIDVal := r.Context().Value(constants.ContextKeyUserID)

	if userIDVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	adminUserID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get buyer ID from URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	var reqBody struct {
		Reason string `json:"reason"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		util.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get admin ID from context

	err = h.sellerService.SuspendSeller(uint(id), reqBody.Reason, adminUserID)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Seller suspended successfully")
}

// ReactivateSeller handles POST /api/sellers/{id}/reactivate
func (h *SellerHandler) ReactivateSeller(w http.ResponseWriter, r *http.Request) {
	// Get user_id and role from context
	userIDVal := r.Context().Value(constants.ContextKeyUserID)

	if userIDVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	adminUserID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get buyer ID from URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}
	err = h.sellerService.ReactivateSeller(uint(id), adminUserID)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Seller reactivated successfully")
}

// SearchSellers handles GET /api/sellers/search
func (h *SellerHandler) SearchSellers(w http.ResponseWriter, r *http.Request) {
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

	sellers, total, err := h.sellerService.SearchSellers(query, page, limit)
	if err != nil {
		util.SendError(w, "Failed to search sellers", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, sellers, page, limit, total)
}

// GetSellerStats handles GET /api/sellers/stats
func (h *SellerHandler) GetSellerStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.sellerService.GetSellerStats()
	if err != nil {
		util.SendError(w, "Failed to fetch statistics", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, stats, "Statistics retrieved successfully")
}

// GetTopSellers handles GET /api/sellers/top
func (h *SellerHandler) GetTopSellers(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limitStr)

	if limit < 1 {
		limit = 10
	}

	sellers, err := h.sellerService.GetTopSellers(limit)
	if err != nil {
		util.SendError(w, "Failed to fetch top sellers", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, sellers, "Top sellers retrieved successfully")
}
