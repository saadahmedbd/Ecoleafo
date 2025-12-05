package carthandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

func (h *CartHandler) ValidateCart(w http.ResponseWriter, r *http.Request) {
	// Extract JWT info safely
	ctxUserID := r.Context().Value(constants.ContextKeyUserID)
	ctxRoles := r.Context().Value(constants.ContextKeyRole)

	if ctxUserID == nil || ctxRoles == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Convert to slice of string
	roles, ok := ctxRoles.([]interface{}) // or []string if you stored it that way
	if !ok {
		http.Error(w, "Invalid role type", http.StatusInternalServerError)
		return
	}

	// Check if "buyer" is in roles
	isBuyer := false
	for _, r := range roles {
		if roleStr, ok := r.(string); ok && roleStr == "buyer" {
			isBuyer = true
			break
		}
	}

	if !isBuyer {
		http.Error(w, "Only buyers can access cart", http.StatusForbidden)
		return
	}

	// Convert JWT reguser.id → buyer.id
	uid, ok := ctxUserID.(uint)
	if !ok {
		http.Error(w, `{"error":"Invalid user ID type"}`, http.StatusInternalServerError)
		return
	}

	buyerID, err := getBuyerID(uid)
	if err != nil {
		http.Error(w, `{"error":"Buyer account not found"}`, http.StatusNotFound)
		return
	}

	// Get address from query parameter (address_id or address string)
	addressID := r.URL.Query().Get("address_id")
	addressStr := r.URL.Query().Get("address")

	var address string
	if addressID != "" {
		// Fetch address by ID
		address = getAddressByID(addressID)
	} else if addressStr != "" {
		address = addressStr
	} else {
		// Use buyer's default address
		address = getBuyerAddress(buyerID)
	}

	var cart interface{}
	if address != "" {
		cart, err = h.service.ValidateCartWithAddress(buyerID, address)
	} else {
		cart, err = h.service.ValidateCart(buyerID)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, cart, 200)
}
