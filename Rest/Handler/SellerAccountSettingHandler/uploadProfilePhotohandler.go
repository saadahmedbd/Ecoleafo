package selleraccountsettinghandler

import (
	"net/http"
	"path/filepath"
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
		util.RespondError(w, http.StatusBadRequest, "Photo file required")
		return
	}
	defer file.Close()

	// Validate file type
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		util.RespondError(w, http.StatusBadRequest, "Only JPG and PNG files allowed")
		return
	}

	// Upload file (implement your file upload logic)
	photoURL, err := util.UploadFile(file, header, "seller-photos")
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to upload photo")
		return
	}

	response, err := h.service.UpdateProfilePhoto(uint(userID), photoURL)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, response, "upload profile photo succesful")
}
