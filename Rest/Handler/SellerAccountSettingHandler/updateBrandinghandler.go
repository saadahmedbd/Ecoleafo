package selleraccountsettinghandler

import (
	"net/http"
	"path/filepath"
	"strings"

	util "github.com/saadahmedbd/Treestore/Util"
)

// ==========================================
// UPDATE BRANDING (LOGO/BANNER)
// PUT /api/seller/store/branding
// ==========================================

func (h *Selleraccountsettinghandler) UpdateBranding(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value("seller_id").(uint)

	// Parse multipart form (max 5MB for banner)
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		util.RespondError(w, http.StatusBadRequest, "File too large")
		return
	}

	var logoURL, bannerURL string

	// Handle logo upload
	if logoFile, logoHeader, err := r.FormFile("logo"); err == nil {
		defer logoFile.Close()

		// Validate logo
		ext := strings.ToLower(filepath.Ext(logoHeader.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			util.RespondError(w, http.StatusBadRequest, "Invalid logo file type")
			return
		}

		logoURL, err = util.UploadFile(logoFile, logoHeader, "store-logos")
		if err != nil {
			util.RespondError(w, http.StatusInternalServerError, "Failed to upload logo")
			return
		}
	}

	// Handle banner upload
	if bannerFile, bannerHeader, err := r.FormFile("banner"); err == nil {
		defer bannerFile.Close()

		// Validate banner
		ext := strings.ToLower(filepath.Ext(bannerHeader.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			util.RespondError(w, http.StatusBadRequest, "Invalid banner file type")
			return
		}

		bannerURL, err = util.UploadFile(bannerFile, bannerHeader, "store-banners")
		if err != nil {
			util.RespondError(w, http.StatusInternalServerError, "Failed to upload banner")
			return
		}
	}

	response, err := h.service.UpdateBranding(sellerID, logoURL, bannerURL)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondJSON(w, http.StatusOK, response, "update seller branding photo")
}
