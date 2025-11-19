package carthandler

import (
	"log"
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

func (h *CartHandler) RemoveFromWishlist(w http.ResponseWriter, r *http.Request) {
	// Create a standard error response map
	errResponse := func(message string) map[string]string {
		return map[string]string{"error": message}
	}

	ctxUserID := r.Context().Value(constants.ContextKeyUserID)
	ctxRoles := r.Context().Value(constants.ContextKeyRole)

	if ctxUserID == nil || ctxRoles == nil {
		// FIX: Use SendData for errors
		util.SendData(w, errResponse("unauthorized"), http.StatusUnauthorized)
		return
	}

	roles, ok := ctxRoles.([]interface{})
	if !ok {
		// FIX: Use SendData for errors
		util.SendData(w, errResponse("Invalid role type"), http.StatusInternalServerError)
		return
	}

	isBuyer := false
	for _, r := range roles {
		if roleStr, ok := r.(string); ok && roleStr == "buyer" {
			isBuyer = true
			break
		}
	}

	if !isBuyer {
		// FIX: Use SendData for errors
		util.SendData(w, errResponse("Only buyers can access cart"), http.StatusForbidden)
		return
	}

	uid, ok := ctxUserID.(uint)
	if !ok {
		// FIX: Use SendData for errors
		util.SendData(w, errResponse("Invalid user ID type"), http.StatusInternalServerError)
		return
	}

	buyerID, err := getBuyerID(uid) // Assuming getBuyerID is in this package or imported
	if err != nil {
		// FIX: Use SendData for errors
		util.SendData(w, errResponse("Buyer account not found"), http.StatusNotFound)
		return
	}

	productIDStr := r.PathValue("productId")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		// FIX: Use SendData for errors
		util.SendData(w, errResponse("Invalid product ID"), http.StatusBadRequest)
		return
	}

	if err := h.service.RemoveFromWishlist(buyerID, uint(productID)); err != nil {
		// FIX: Use SendData for errors AND log the real error for yourself
		log.Printf("Failed to remove from wishlist: %v", err) // Log the internal error
		util.SendData(w, errResponse("Could not remove item from wishlist"), http.StatusInternalServerError)
		return
	}

	// This was already correct
	util.SendData(w, map[string]string{"message": "Item removed from wishlist"}, 200)
}
