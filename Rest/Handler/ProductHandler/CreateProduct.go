package producthandler

import (
	"encoding/json"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"

	util "github.com/saadahmedbd/Treestore/Util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Get claims from context
	claims, ok := r.Context().Value("claims").(map[string]interface{})
	if !ok {
		http.Error(w, `{"error":"Unauthorized: No claims found"}`, http.StatusUnauthorized)
		return
	}

	// Check if user has "seller" role
	roles, ok := claims["roles"].([]interface{})
	if !ok {
		http.Error(w, `{"error":"Unauthorized: Invalid roles format"}`, http.StatusUnauthorized)
		return
	}
	hasSellerRole := false
	for _, role := range roles {
		if role == "seller" {
			hasSellerRole = true
			break
		}
	}
	if !hasSellerRole {
		http.Error(w, `{"error":"Unauthorized: Only sellers can post products"}`, http.StatusForbidden)
		return
	}

	// Get user_id from claims
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, `{"error":"Unauthorized: Invalid user_id format"}`, http.StatusUnauthorized)
		return
	}
	userID := uint(userIDFloat)

	if r.Method != "POST" {
		http.Error(w, "Please provide valid request", http.StatusBadRequest)
		return
	}
	var products models.Product
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&products)
	if err != nil {
		http.Error(w, "please provide valid json", http.StatusBadRequest)
		return
	}
	// products.SellerID = sellerID
	result := Config.DB.Create(&products)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, products, 200)

}
