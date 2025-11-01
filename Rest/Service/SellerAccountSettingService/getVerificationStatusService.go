package selleraccountsettingservice

import selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"

// GetVerificationStatus retrieves verification status
func (s *SellerAccountSettingService) GetVerificationStatus(sellerID uint) (*selleraccountsetting.VerificationStatus, error) {
	docs, err := s.repo.GetVerificationDocuments(sellerID)
	if err != nil {
		return nil, err
	}

	status := &selleraccountsetting.VerificationStatus{
		Identity: "not-submitted",
		Tax:      "not-submitted",
		Bank:     "not-submitted",
	}

	for _, doc := range docs {
		switch doc.DocumentType {
		case "identity":
			status.Identity = doc.Status
		case "tax":
			status.Tax = doc.Status
		case "bank":
			status.Bank = doc.Status
		}
	}

	return status, nil
}
