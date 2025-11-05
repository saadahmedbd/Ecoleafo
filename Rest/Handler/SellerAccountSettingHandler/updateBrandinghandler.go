package selleraccountsettinghandler

import (
	"net/http"
	"strconv"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// UPDATE BRANDING (LOGO/BANNER)
// PUT /api/seller/store/branding
// ==========================================

func (h *Selleraccountsettinghandler) UpdateBranding(w http.ResponseWriter, r *http.Request) {
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

	// Parse multipart form (max 5MB for banner)
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		util.RespondError(w, http.StatusBadRequest, "File too large")
		return
	}

	var logoURL, bannerURL string
	var uploadErrors []error

	// Handle logo upload
	if logoFile, logoHeader, err := r.FormFile("logo"); err == nil {
		defer logoFile.Close()

		logoURL, err = util.UploadStoreLogo(logoFile, logoHeader)
		if err != nil {
			uploadErrors = append(uploadErrors, err)
		}
	}

	// Handle banner upload
	if bannerFile, bannerHeader, err := r.FormFile("banner"); err == nil {
		defer bannerFile.Close()

		bannerURL, err = util.UploadStoreBanner(bannerFile, bannerHeader)
		if err != nil {
			uploadErrors = append(uploadErrors, err)
		}
	}

	// Check if any uploads failed
	if len(uploadErrors) > 0 {
		// Clean up successfully uploaded files
		if logoURL != "" {
			util.DeleteImageFromCloudinary(logoURL)
		}
		if bannerURL != "" {
			util.DeleteImageFromCloudinary(bannerURL)
		}
		util.RespondError(w, http.StatusBadRequest, uploadErrors[0].Error())
		return
	}

	// Update database
	response, err := h.service.UpdateBranding(uint(userID), logoURL, bannerURL)
	if err != nil {
		// Clean up uploaded files on database error
		if logoURL != "" {
			util.DeleteImageFromCloudinary(logoURL)
		}
		if bannerURL != "" {
			util.DeleteImageFromCloudinary(bannerURL)
		}
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, response, "upload seller logo and banner succesful")

}
