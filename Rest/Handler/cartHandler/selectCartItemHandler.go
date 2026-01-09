package carthandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

type SelectCartItemRequest struct {
	IsSelected bool `json:"is_selected"`
}

func (h *CartHandler) SelectCartItem(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value(constants.ContextKeyUserID)
	ctxRoles := r.Context().Value(constants.ContextKeyRole)

	if ctxUserID == nil || ctxRoles == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	roles, ok := ctxRoles.([]interface{})
	if !ok {
		http.Error(w, "Invalid role type", http.StatusInternalServerError)
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
		http.Error(w, "Only buyers can access cart", http.StatusForbidden)
		return
	}

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

	cartItemIDStr := r.PathValue("productId")
	cartItemID, err := strconv.ParseUint(cartItemIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"Invalid cart item ID"}`, http.StatusBadRequest)
		return
	}

	var req SelectCartItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.SelectCartItemByID(buyerID, uint(cartItemID), req.IsSelected); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]interface{}{
		"message": "Cart item selection updated",
		"selected": req.IsSelected,
	}, http.StatusOK)
}

func (h *CartHandler) SelectAllCartItems(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value(constants.ContextKeyUserID)
	ctxRoles := r.Context().Value(constants.ContextKeyRole)

	if ctxUserID == nil || ctxRoles == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	roles, ok := ctxRoles.([]interface{})
	if !ok {
		http.Error(w, "Invalid role type", http.StatusInternalServerError)
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
		http.Error(w, "Only buyers can access cart", http.StatusForbidden)
		return
	}

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

	var req SelectCartItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if err := h.service.SelectAllCartItems(buyerID, req.IsSelected); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]interface{}{
		"message": "All cart items selection updated",
		"selected": req.IsSelected,
	}, http.StatusOK)
}
