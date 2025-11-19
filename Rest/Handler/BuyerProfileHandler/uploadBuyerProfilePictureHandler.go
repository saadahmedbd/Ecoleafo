package buyerprofilehandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
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
	//Get user data from context
	ctxUserID := r.Context().Value(constants.ContextKeyUserID)
	ctxRoles := r.Context().Value(constants.ContextKeyRole)

	if ctxUserID == nil || ctxRoles == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Convert roles
	roleList, ok := ctxRoles.([]interface{})
	if !ok {
		http.Error(w, "Invalid role type", http.StatusInternalServerError)
		return
	}

	// Check if buyer
	isBuyer := false
	for _, raw := range roleList {
		if roleStr, ok := raw.(string); ok && roleStr == "buyer" {
			isBuyer = true
			break
		}
	}

	if !isBuyer {
		http.Error(w, "Only buyer can access this", http.StatusForbidden)
		return
	}

	// Convert user id
	uid, ok := ctxUserID.(uint)
	if !ok {
		http.Error(w, `{"error":"Invalid user ID type"}`, http.StatusInternalServerError)
		return
	}

	// Convert reguser.id → buyer.id
	buyerID, err := getBuyerID(uid)
	if err != nil {
		http.Error(w, `{"error":"seller account not found"}`, http.StatusNotFound)
		return
	}

	//  Step 2: Limit upload size to 10MB
	const maxUploadSize = 10 << 20 // 10MB
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	// Step 3: Parse multipart form safely
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "File too large or invalid form data", http.StatusBadRequest)
		return
	}

	//  Step 4: Extract file
	file, header, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Failed to read uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Step 5: Validate file size
	if header.Size > maxUploadSize {
		http.Error(w, "File too large (max 10MB)", http.StatusBadRequest)
		return
	}

	//  6: Validate file type (security check)
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		http.Error(w, "Failed to read file buffer", http.StatusInternalServerError)
		return
	}
	fileType := http.DetectContentType(buff)
	if fileType != "image/jpeg" && fileType != "image/png" && fileType != "image/webp" {
		http.Error(w, "Only JPEG, PNG, or WEBP images are allowed", http.StatusBadRequest)
		return
	}
	file.Seek(0, 0) // reset pointer before upload

	//  7: Upload to Cloudinary (or wherever)
	photoURL, publicID, err := util.UploadBuyerProfilePhoto(file, header)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	//  Step 8: Update database
	response, err := h.buyerservice.UpdateProfilePhoto(uint(buyerID), photoURL, publicID)
	if err != nil {
		// Rollback uploaded image if DB update fails
		util.DeleteImageFromCloudinary(photoURL)
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//  Step 9: Send success response
	util.RespondJSON(w, http.StatusOK, response, "Profile photo uploaded successfully")
}

// getServerBaseURL returns the base URL of the server from the request
