package producthandler

import (
	"net/http"
	"strconv"

	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

// LoadMoreResponse - Optimized for e-commerce UX
type LoadMoreResponse struct {
	Data       []models.Product `json:"data"`
	Total      int64            `json:"total"`
	Loaded     int              `json:"loaded"`      // Products loaded so far
	HasMore    bool             `json:"has_more"`    // Are there more products?
	NextOffset int              `json:"next_offset"` // Offset for next load
	PerPage    int              `json:"per_page"`
}

// GetProductsLoadMore - Best UX for e-commerce homepage
func (h *Handler) GetProductsLoadMore(w http.ResponseWriter, r *http.Request) {
	// Parse parameters
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if offset < 0 {
		offset = 0
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 || limit > 50 {
		limit = 20 // Optimal for homepage: 20 products per load
	}

	includeRelations := r.URL.Query().Get("include_relations") == "true"

	// Build filters
	filters := make(map[string]interface{})

	// Only apply filters if explicitly provided
	if featured := r.URL.Query().Get("is_featured"); featured != "" {
		filters["is_featured"] = featured == "true"
	}
	if isActive := r.URL.Query().Get("is_active"); isActive != "" {
		filters["is_active"] = isActive == "true"
	}
	if isApproved := r.URL.Query().Get("is_approved"); isApproved != "" {
		filters["is_approved"] = isApproved == "true"
	}
	if categoryID := r.URL.Query().Get("category_id"); categoryID != "" {
		if id, err := strconv.Atoi(categoryID); err == nil {
			filters["category_id"] = uint(id)
		}
	}
	if search := r.URL.Query().Get("search"); search != "" {
		filters["search"] = search
	}

	result, err := h.service.GetProductsLoadMore(offset, limit, filters, includeRelations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, result, 200)
}

// func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
// 	//perse query parameter
// 	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
// 	if page < 0 {
// 		page = 1
// 	}
// 	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
// 	if limit <= 1 || limit > 10 {
// 		limit = 10
// 	}
// 	includeRelations := r.URL.Query().Get("include_relations") == "true"

// 	// Build filters
// 	filters := make(map[string]interface{})
// 	if sellerID := r.URL.Query().Get("seller_id"); sellerID != "" {
// 		if id, err := strconv.Atoi(sellerID); err == nil {
// 			filters["seller_id"] = uint(id)
// 		}
// 	}
// 	if categoryID := r.URL.Query().Get("category_id"); categoryID != "" {
// 		if id, err := strconv.Atoi(categoryID); err == nil {
// 			filters["category_id"] = uint(id)
// 		}
// 	}
// 	if isActive := r.URL.Query().Get("is_active"); isActive != "" {
// 		filters["is_active"] = isActive == "true"
// 	}
// 	if isApproved := r.URL.Query().Get("is_approved"); isApproved != "" {
// 		filters["is_approved"] = isApproved == "true"
// 	}
// 	if isFeatured := r.URL.Query().Get("is_featured"); isFeatured != "" {
// 		filters["is_featured"] = isFeatured == "true"
// 	}
// 	if treeType := r.URL.Query().Get("tree_type"); treeType != "" {
// 		filters["tree_type"] = treeType
// 	}
// 	if search := r.URL.Query().Get("search"); search != "" {
// 		filters["search"] = search
// 	}
// 	result, err := h.service.GetProducts(page, limit, filters, includeRelations)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	util.SendData(w, result, 200)
// }
