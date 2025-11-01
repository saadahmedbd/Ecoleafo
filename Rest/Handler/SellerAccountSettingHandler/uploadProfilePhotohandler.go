package selleraccountsettinghandler

import (
	"net/http"
	"path/filepath"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// UPLOAD PROFILE PHOTO
// POST /api/seller/account/photo
// ==========================================

func (h *Selleraccountsettinghandler) UploadProfilePhoto(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

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

	response, err := h.service.UpdateProfilePhoto(sellerID, photoURL)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, response, "upload profile photo succesful")
}
