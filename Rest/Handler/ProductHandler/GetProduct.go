package producthandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	//perse query parameter
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 1 || limit > 10 {
		limit = 10
	}
	includeRelations := r.URL.Query().Get("include_relations") == "true"

	// Build filters
	filters := make(map[string]interface{})
	if sellerID := r.URL.Query().Get("seller_id"); sellerID != "" {
		if id, err := strconv.Atoi(sellerID); err == nil {
			filters["seller_id"] = uint(id)
		}
	}
	if categoryID := r.URL.Query().Get("category_id"); categoryID != "" {
		if id, err := strconv.Atoi(categoryID); err == nil {
			filters["category_id"] = uint(id)
		}
	}
	if isActive := r.URL.Query().Get("is_active"); isActive != "" {
		filters["is_active"] = isActive == "true"
	}
	if isApproved := r.URL.Query().Get("is_approved"); isApproved != "" {
		filters["is_approved"] = isApproved == "true"
	}
	if isFeatured := r.URL.Query().Get("is_featured"); isFeatured != "" {
		filters["is_featured"] = isFeatured == "true"
	}
	if treeType := r.URL.Query().Get("tree_type"); treeType != "" {
		filters["tree_type"] = treeType
	}
	if search := r.URL.Query().Get("search"); search != "" {
		filters["search"] = search
	}
	result, err := h.service.GetProducts(page, limit, filters, includeRelations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, result, 200)
}
