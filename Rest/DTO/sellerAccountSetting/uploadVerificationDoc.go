package selleraccountsetting

// UploadVerificationDocumentRequest - Document upload request
type UploadVerificationDocumentRequest struct {
	DocumentType string `json:"document_type" validate:"required,oneof=identity tax bank"`
}
