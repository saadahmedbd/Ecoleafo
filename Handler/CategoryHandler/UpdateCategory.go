package categoryHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Please provide valid method", http.StatusBadRequest)
		return
	}
	categoryId := r.PathValue("categoryId")
	id, err := strconv.Atoi(categoryId)
	fmt.Println("Params:", id)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	var existingCategory models.Category
	if err := Config.DB.First(&existingCategory, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	var updateCategory models.Category
	if err := json.NewDecoder(r.Body).Decode(&updateCategory); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Update existing seller with new values
	existingCategory.Name = updateCategory.Name
	existingCategory.Slug = updateCategory.Slug

	if err := Config.DB.Save(&existingCategory).Error; err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}
	util.SendData(w, updateCategory, 200)
}
