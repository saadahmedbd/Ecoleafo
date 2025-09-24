package categoryHandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	categoryId := r.PathValue("categoryId")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}
	//try to delete
	var categories models.Category
	if err := Config.DB.First(&categories, id).Error; err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}
	if err := Config.DB.Delete(&categories).Error; err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}
	// util.SendData(w, seller, 200)
	w.Write([]byte(fmt.Sprintf("category %d deleted successfully", id)))

}
