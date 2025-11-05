package buyerprofilehandler

import (
	"encoding/json"

	"net/http"
	"strconv"
)

// DeleteBuyerProfilePicture handles profile picture deletion

// ProfilePictureResponse defines the response structure

// DeleteBuyerProfilePicture deletes the buyer’s profile picture (Cloudinary + DB)
type DeleteProfileRequest struct {
	ImageURL string `json:"image_url"`
}


func (h *Buyerprofilehandler) DeleteBuyerProfilePicture(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req DeleteProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.ImageURL == "" {
		http.Error(w, "image_url is required", http.StatusBadRequest)
		return
	}

	if err := h.buyerservice.DeleteBuyerProfilePicture(uint(userID), req.ImageURL); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Profile picture deleted successfully",
	})
}
