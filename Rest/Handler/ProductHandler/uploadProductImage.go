package producthandler

import (
	"net/http"

	util "github.com/saadahmedbd/Treestore/Util"
)

// UploadProductImage handles product image upload to Cloudinary
func (h *Handler) UploadProductImage(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form (max 10MB)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		util.RespondError(w, http.StatusBadRequest, "Failed to parse form data")
		return
	}

	// Get the file from form
	file, header, err := r.FormFile("image")
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, "Image file is required")
		return
	}
	defer file.Close()

	// Upload to Cloudinary
	imageURL, err := util.UploadProductImage(file, header)
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Return the uploaded image URL
	util.RespondJSON(w, http.StatusOK, map[string]string{
		"image_url": imageURL,
	}, "Image uploaded successfully")
}
