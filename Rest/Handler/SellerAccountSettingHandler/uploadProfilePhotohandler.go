package selleraccountsettinghandler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// UPLOAD PROFILE PHOTO
// POST /api/seller/account/photo
// ==========================================

func (h *Selleraccountsettinghandler) UploadProfilePhoto(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("user_id")
	userType := r.Header.Get("user_role")

	if userIDStr == "" || userType == "" {
		http.Error(w, `{"error":"authentication required"}`, http.StatusUnauthorized)
		return
	}

	if !strings.Contains(userType, "seller") {
		http.Error(w, `{"error":"only sellers can access this"}`, http.StatusForbidden)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
		return
	}

	// Parse multipart form (max 2MB)
	if err := r.ParseMultipartForm(2 << 20); err != nil {
		util.RespondError(w, http.StatusBadRequest, "File too large (max 2MB)")
		return
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		log.Printf("File upload error: %v", err)
		util.RespondError(w, http.StatusBadRequest, "Photo file required")
		return
	}
	defer file.Close()

	// Upload to Cloudinary
	photoURL, err := util.UploadProfilePhoto(file, header)
	if err != nil {
		// log.Printf("File upload error: %v", err)
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Update database
	response, err := h.service.UpdateProfilePhoto(uint(userID), photoURL)
	if err != nil {
		// If database update fails, try to delete uploaded image
		util.DeleteImageFromCloudinary(photoURL)
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, response, "upload profile photo successful")

}
