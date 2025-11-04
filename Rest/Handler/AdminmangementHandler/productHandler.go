package adminmangementhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	adminmangementservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminMangementService"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

type ProductHandler struct {
	productService *adminmangementservice.ProductService
}

func NewProductHandler(productService *adminmangementservice.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// GetAllProducts handles GET /api/products
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
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

	products, total, err := h.productService.GetAllProducts(page, limit, status)
	if err != nil {
		util.SendError(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, products, page, limit, total)
}

// GetProductByID handles GET /api/products/{id}
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	util.RespondJSON(w, http.StatusOK, product, "Product retrieved successfully")
}

// GetPendingApprovals handles GET /api/products/pending
func (h *ProductHandler) GetPendingApprovals(w http.ResponseWriter, r *http.Request) {
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

	products, total, err := h.productService.GetPendingApprovals(page, limit)
	if err != nil {
		util.SendError(w, "Failed to fetch pending approvals", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, products, page, limit, total)
}

// ApproveProduct handles POST /api/products/{id}/approve
func (h *ProductHandler) ApproveProduct(w http.ResponseWriter, r *http.Request) {
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
		util.SendError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	err = h.productService.ApproveProduct(uint(id), adminUserID)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Product approved successfully")
}

// RejectProduct handles POST /api/products/{id}/reject
func (h *ProductHandler) RejectProduct(w http.ResponseWriter, r *http.Request) {
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
		util.SendError(w, "Invalid product ID", http.StatusBadRequest)
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

	err = h.productService.RejectProduct(uint(id), reqBody.Reason, adminUserID)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Product rejected successfully")
}

// SearchProducts handles GET /api/products/search
func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
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

	products, total, err := h.productService.SearchProducts(query, page, limit)
	if err != nil {
		util.SendError(w, "Failed to search products", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, products, page, limit, total)
}

// GetProductStats handles GET /api/products/stats
func (h *ProductHandler) GetProductStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.productService.GetProductStats()
	if err != nil {
		util.SendError(w, "Failed to fetch statistics", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, stats, "Statistics retrieved successfully")
}

// DeleteProduct handles DELETE /api/products/{id}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
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
		util.SendError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	err = h.productService.DeleteProduct(uint(id), adminUserID)
	if err != nil {
		util.SendError(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Product deleted successfully")
}
