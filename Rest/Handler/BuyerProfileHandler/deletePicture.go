package buyerprofilehandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DeleteBuyerProfilePicture handles profile picture deletion
func (h *Buyerprofilehandler) DeleteBuyerProfilePicture(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Validate user authentication
	userIDStr := r.Header.Get("user_id")
	if userIDStr == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	// // Convert userID to uint
	// userID, err := strconv.ParseUint(userIDStr, 10, 64)
	// if err != nil {
	// 	http.Error(w, "Invalid user ID", http.StatusBadRequest)
	// 	return
	// }

	// Parse request body
	var request struct {
		PublicID string `json:"public_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if request.PublicID == "" {
		http.Error(w, "Public ID is required", http.StatusBadRequest)
		return
	}

	// Delete from Cloudinary
	if err := h.cloudniaryservice.DeleteImage(ctx, request.PublicID); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete image: %v", err), http.StatusInternalServerError)
		return
	}

	// TODO: Update database to remove profile_picture_url and profile_picture_public_id
	// userID, _ := strconv.ParseUint(userIDStr, 10, 64)
	// err := h.clearProfilePicture(userID)
	// if err != nil {
	//     h.sendErrorResponse(w, "Failed to update database", http.StatusInternalServerError)
	//     return
	// }

	// Send success response
	response := ProfilePictureResponse{
		Success: true,
		Message: "Profile picture deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
