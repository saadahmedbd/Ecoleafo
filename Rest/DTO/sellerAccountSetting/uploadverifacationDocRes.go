package selleraccountsetting

// UploadVerificationDocumentResponse - Document upload response
type UploadVerificationDocumentResponse struct {
	DocumentURL  string `json:"document_url"`
	DocumentType string `json:"document_type"`
	Status       string `json:"status"`
	Message      string `json:"message"`
}
