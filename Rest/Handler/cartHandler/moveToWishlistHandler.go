package carthandler

import (
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

func (h *CartHandler) MoveToWishlist(w http.ResponseWriter, r *http.Request) {
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

	productIDStr := r.PathValue("productId")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if err := h.service.MoveToWishlist(buyerID, uint(productID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]string{"message": "Item moved to wishlist"}, 200)
}
