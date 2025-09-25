package producthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Please provide valid method", http.StatusBadRequest)
		return
	}
	ProductId := r.PathValue("productId")
	id, err := strconv.Atoi(ProductId)
	fmt.Println("Params:", id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var existingProduct models.Product
	if err := Config.DB.First(&existingProduct, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	var updateProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updateProduct); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Update existing seller with new values
	existingProduct.Name = updateProduct.Name
	existingProduct.Slug = updateProduct.Slug
	existingProduct.Description = updateProduct.Description
	existingProduct.SKU = updateProduct.Slug

	existingProduct.CategoryID = updateProduct.CategoryID
	existingProduct.Price = updateProduct.Price
	existingProduct.Quantity = updateProduct.Quantity
	existingProduct.IsActive = updateProduct.IsActive

	if err := Config.DB.Save(&existingProduct).Error; err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	util.SendData(w, updateProduct, 200)

}
