package producthandler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) GetSellerProduct(w http.ResponseWriter, r *http.Request) {
	// Get user info from JWT (set by middleware)
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")

	// Validation
	if userIDStr == "" {
		http.Error(w, `{"error":"user_id not found in headers"}`, http.StatusUnauthorized)
		return
	}
	if userType == "" {
		http.Error(w, `{"error":"user_role not found in headers"}`, http.StatusUnauthorized)
		return
	}

	// Check if user is a seller
	if !strings.Contains(userType, "seller") {
		http.Error(w, `{"error":"only sellers can view their products"}`, http.StatusForbidden)
		return
	}

	// Parse user ID from JWT
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id format"}`, http.StatusBadRequest)
		return
	}

	// Get query parameters for filtering and pagination
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	limit, _ := strconv.Atoi(queryParams.Get("limit"))
	status := queryParams.Get("status") // active, inactive, approved, pending
	category := queryParams.Get("category")
	search := queryParams.Get("search")
	sortBy := queryParams.Get("sort") // name, price, created_at
	order := queryParams.Get("order") // asc, desc

	// Set defaults
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20 // Default limit
	}
	if order != "asc" && order != "desc" {
		order = "desc"
	}
	if sortBy == "" {
		sortBy = "created_at"
	}

	// Create filter options
	filters := productservice.ProductFilters{
		Page:     page,
		Limit:    limit,
		Status:   status,
		Category: category,
		Search:   search,
		SortBy:   sortBy,
		Order:    order,
	}

	// Get seller products
	result, err := h.service.GetSellerProducts(uint(userID), filters)
	if err != nil {
		fmt.Printf("DEBUG: GetSellerProducts error: %v\n", err)
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	util.SendData(w, result, http.StatusOK)
}
