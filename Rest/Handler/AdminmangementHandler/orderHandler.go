package adminmangementhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	adminmangementservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminMangementService"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

type OrderHandler struct {
	orderService *adminmangementservice.OrderService
}

func NewOrderHandler(orderService *adminmangementservice.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// GetAllOrders handles GET /api/orders
func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
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

	orders, total, err := h.orderService.GetAllOrders(page, limit, status)
	if err != nil {
		util.SendError(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, orders, page, limit, total)
}

// GetOrderByID handles GET /api/orders/{id}
func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {

	// Get admin ID from context
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	order, err := h.orderService.GetOrderByID(uint(id))
	if err != nil {
		util.SendError(w, "Order not found", http.StatusNotFound)
		return
	}

	util.RespondJSON(w, http.StatusOK, order, "Order retrieved successfully")
}

// SearchOrders handles GET /api/orders/search
func (h *OrderHandler) SearchOrders(w http.ResponseWriter, r *http.Request) {
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

	orders, total, err := h.orderService.SearchOrders(query, page, limit)
	if err != nil {
		util.SendError(w, "Failed to search orders", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, orders, page, limit, total)
}

// UpdateOrderStatus handles PUT /api/orders/{id}/status
func (h *OrderHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	// Get user info from context
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
		util.SendError(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	var reqBody struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		util.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if reqBody.Status == "" {
		util.SendError(w, "Status is required", http.StatusBadRequest)
		return
	}

	err = h.orderService.UpdateOrderStatus(uint(id), reqBody.Status, adminUserID)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Order status updated successfully")
}

// CancelOrder handles POST /api/orders/{id}/cancel
func (h *OrderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
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
		util.SendError(w, "Invalid order ID", http.StatusBadRequest)
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
		util.SendError(w, "Cancellation reason is required", http.StatusBadRequest)
		return
	}

	err = h.orderService.CancelOrder(uint(id), reqBody.Reason, adminUserID)
	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, nil, "Order cancelled successfully")
}

// GetOrderStats handles GET /api/orders/stats
func (h *OrderHandler) GetOrderStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.orderService.GetOrderStats()
	if err != nil {
		util.SendError(w, "Failed to fetch statistics", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, stats, "Statistics retrieved successfully")
}

// GetRecentOrders handles GET /api/orders/recent
func (h *OrderHandler) GetRecentOrders(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limitStr)

	if limit < 1 {
		limit = 10
	}

	orders, err := h.orderService.GetRecentOrders(limit)
	if err != nil {
		util.SendError(w, "Failed to fetch recent orders", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, orders, "Recent orders retrieved successfully")
}
