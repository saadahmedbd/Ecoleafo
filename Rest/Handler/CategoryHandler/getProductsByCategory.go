package categoryHandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *CategoryHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		util.SendError(w, "Invalid category ID", http.StatusBadRequest)
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
		limit = 20
	}

	products, total, err := h.categoryService.GetProductsByCategory(uint(id), page, limit)
	if err != nil {
		util.SendError(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedResponse(w, http.StatusOK, products, page, limit, total)
}
