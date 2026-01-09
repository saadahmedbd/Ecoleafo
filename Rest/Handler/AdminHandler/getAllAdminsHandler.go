package adminhandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Adminhandler) GetAllAdmins(w http.ResponseWriter, r *http.Request) {

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	admins, total, err := h.adminService.GetAllAdmins(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"admins": admins,
		"total":  total,
		"page":   page,
		"limit":  limit,
	}
	util.SendData(w, response, 200)
}
