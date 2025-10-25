package buyerprofilehandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	cloudinarydto "github.com/saadahmedbd/Treestore/Rest/DTO/CloudinaryDTO"
)

// ProfilePictureResponse defines the response structure
type ProfilePictureResponse struct {
	Success  bool   `json:"success"`
	URL      string `json:"url"`
	PublicID string `json:"public_id,omitempty"`
	Message  string `json:"message,omitempty"`
	Error    string `json:"error,omitempty"`
}

// UploadBuyerProfilePicture handles profile picture upload
func (h *Buyerprofilehandler) UploadBuyerProfilePicture(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract and validate user ID
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

	// Parse multipart form with size limit (10MB)
	const maxUploadSize = 10 << 20 // 10MB
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "File too large or invalid form data", http.StatusBadRequest)
		return
	}

	// Get the file from form
	file, handler, err := r.FormFile("profile_picture")
	if err != nil {
		http.Error(w, "Could not get file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file size
	if handler.Size > maxUploadSize {
		http.Error(w, "File size exceeds 10MB limit", http.StatusBadRequest)
		return
	}

	// TODO: Delete old profile picture if exists
	// oldPublicID := h.getOldProfilePicturePublicID(userID)
	// if oldPublicID != "" {
	//     h.cloudinaryService.DeleteImage(ctx, oldPublicID)
	// }

	// Upload to Cloudinary
	uploadOpts := cloudinarydto.UploadImageOptions{
		Folder:           "buyer_profiles",
		PublicID:         fmt.Sprintf("buyer_%d_%d", userID, time.Now().Unix()),
		AllowedFormats:   []string{".jpg", ".jpeg", ".png", ".webp"},
		MaxFileSizeBytes: maxUploadSize,
		Tags:             []string{"buyer", "profile_picture"},
		// Optional: Add transformation for optimization
		Transformation: "c_fill,g_face,h_500,w_500,q_auto,f_auto",
	}

	result, err := h.cloudniaryservice.UploadImage(ctx, file, handler.Filename, uploadOpts)
	if err != nil {
		http.Error(w, fmt.Sprintf("Upload failed: %v", err), http.StatusInternalServerError)
		return
	}

	// TODO: Save the secure URL to your database
	// err = h.saveProfilePictureURL(userID, result.SecureURL, result.PublicID)
	// if err != nil {
	//     // Rollback: delete uploaded image
	//     h.cloudinaryService.DeleteImage(ctx, result.PublicID)
	//     h.sendErrorResponse(w, "Failed to save profile picture", http.StatusInternalServerError)
	//     return
	// }

	// Send success response
	response := ProfilePictureResponse{
		Success:  true,
		URL:      result.SecureURL, // Always use secure URL (HTTPS)
		PublicID: result.PublicID,
		Message:  "Profile picture uploaded successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// getServerBaseURL returns the base URL of the server from the request
