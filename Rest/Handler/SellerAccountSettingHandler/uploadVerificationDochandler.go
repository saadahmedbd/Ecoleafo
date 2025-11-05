package selleraccountsettinghandler

import (
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// UPLOAD VERIFICATION DOCUMENT
// POST /api/seller/verification/upload
// ==========================================

func (h *Selleraccountsettinghandler) UploadVerificationDocument(w http.ResponseWriter, r *http.Request) {
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

	// Upload to Cloudinary
	documentURL, err := util.UploadVerificationDocument(file, header)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Save to database
	response, err := h.service.UploadVerificationDocument(uint(userID), documentType, documentURL)
	if err != nil {
		// Clean up uploaded file on database error
		util.DeleteImageFromCloudinary(documentURL)
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, response, "upload document succesful")

}
