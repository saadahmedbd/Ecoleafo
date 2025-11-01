package selleraccountsettinghandler

import (
	"net/http"
	"path/filepath"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// UPLOAD VERIFICATION DOCUMENT
// POST /api/seller/verification/upload
// ==========================================

func (h *Selleraccountsettinghandler) UploadVerificationDocument(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	// Parse multipart form (max 5MB)
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		util.RespondError(w, http.StatusBadRequest, "File too large (max 5MB)")
		return
	}

	documentType := r.FormValue("document_type")
	if documentType == "" {
		util.RespondError(w, http.StatusBadRequest, "document_type required")
		return
	}

	file, header, err := r.FormFile("document")
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Document file required")
		return
	}
	defer file.Close()

	// Validate file type
	ext := strings.ToLower(filepath.Ext(header.Filename))
	validExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".pdf": true}
	if !validExts[ext] {
		util.RespondError(w, http.StatusBadRequest, "Only JPG, PNG, and PDF files allowed")
		return
	}

	// Upload file
	documentURL, err := util.UploadFile(file, header, "verification-docs")
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Failed to upload document")
		return
	}

	response, err := h.service.UploadVerificationDocument(sellerID, documentType, documentURL)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, response, "upload successfully")
}
