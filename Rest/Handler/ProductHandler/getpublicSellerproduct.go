package producthandler

import (
	"fmt"
	"net/http"
	"strconv"

	productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"
	util "github.com/saadahmedbd/Treestore/Util"
)

// ============= GET SPECIFIC SELLER'S PRODUCTS (PUBLIC) =============
func (h *Handler) GetPublicSellerProducts(w http.ResponseWriter, r *http.Request) {
	// Extract seller ID from URL
	sellerId := r.PathValue("sellerID")
	sid, err := strconv.Atoi(sellerId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	// Get query parameters
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	limit, _ := strconv.Atoi(queryParams.Get("limit"))
	category := queryParams.Get("category")
	search := queryParams.Get("search")
	sortBy := queryParams.Get("sort")
	order := queryParams.Get("order")

	// Set defaults
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 20
	}
	if order != "asc" && order != "desc" {
		order = "desc"
	}
	if sortBy == "" {
		sortBy = "created_at"
	}

	filters := productservice.ProductFilters{
		Page:     page,
		Limit:    limit,
		Status:   "active", // Only show active products to public
		Category: category,
		Search:   search,
		SortBy:   sortBy,
		Order:    order,
	}

	// Get public seller products (by seller primary key ID)
	result, err := h.service.GetPublicSellerProducts(uint(sid), filters)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	util.SendData(w, result, http.StatusOK)
}
