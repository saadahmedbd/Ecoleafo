package selleraccountsettingservice

import (
	"errors"

	models "github.com/saadahmedbd/Treestore/Models"
	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
)

// UploadVerificationDocument uploads verification document
func (s *SellerAccountSettingService) UploadVerificationDocument(sellerID uint, docType, documentURL string) (*selleraccountsetting.UploadVerificationDocumentResponse, error) {
	// Validate document type
	validTypes := map[string]bool{
		"identity": true,
		"tax":      true,
		"bank":     true,
	}

	if !validTypes[docType] {
		return nil, errors.New("invalid document type")
	}

	// Create document record
	doc := &models.SellerVerificationDocument{
		SellerID:     sellerID,
		DocumentType: docType,
		DocumentURL:  documentURL,
		Status:       "pending",
	}

	if err := s.repo.CreateVerificationDocument(doc); err != nil {
		return nil, errors.New("failed to upload document")
	}

	return &selleraccountsetting.UploadVerificationDocumentResponse{
		DocumentURL:  documentURL,
		DocumentType: docType,
		Status:       "pending",
		Message:      "Document uploaded successfully and pending review",
	}, nil
}
