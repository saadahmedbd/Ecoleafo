package producthandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"
	util "github.com/saadahmedbd/Treestore/Util"
)

// now i do not use this file
func (h *Handler) AddProductImage(w http.ResponseWriter, r *http.Request) {
	// Get seller ID from JWT token
	userIDStr := r.Header.Get("X-User-ID")
	userType := r.Header.Get("X-User-Type")

	if userType != "seller" {
		http.Error(w, "Only sellers can add product images", http.StatusForbidden)
		return
	}

	sellerID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	product, err := h.service.GetProduct(uint(id))
	if err != nil || product.SellerID != uint(sellerID) {
		http.Error(w, "Product not found or not owned by the seller", http.StatusNotFound)
		return
	}
	var req productservice.ProductImageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Create product image
	image := models.ProductImage{
		ProductID: uint(id),
		ImageURL:  req.ImageURL,
		AltText:   req.AltText,
		IsPrimary: req.IsPrimary,
		SortOrder: req.SortOrder,
	}

	if err := Config.DB.Create(&image).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, image, http.StatusCreated)

}
