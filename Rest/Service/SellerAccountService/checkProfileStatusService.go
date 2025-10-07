package selleraccountservice

import selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"

func (s *sellerRegistrationService) CheckProfileStatus(userID uint) (*selleraccount.SellerProfileCompletionStatus, error) {
	seller, err := s.sellerRepo.GetSellerByUserId(userID)
	if err != nil {
		return nil, err
	}
	isComplete, missingFields, err := s.sellerRepo.CheckProfileCompletion(seller.ID)
	if err != nil {
		return nil, err
	}
	status := &selleraccount.SellerProfileCompletionStatus{
		IsProfileComplete: isComplete,
		HasBusinessInfo:   seller.Phone != "" && seller.BusinessEmail != "N/A",
		HasAddress:        seller.Address != "" && seller.Address != "N/A",
		IsApproved:        seller.IsApproved,
		IsVerified:        seller.IsVerified,
		MissingFields:     missingFields,
		RejectionReason:   seller.RejectionReason,
	}
	// Check payment method
	methods, _ := s.sellerRepo.GetPaymentMethods(seller.ID)
	status.HasPaymentMethod = len(methods) > 0

	// Determine approval status
	if seller.RejectedAt != nil {
		status.ApprovalStatus = "rejected"
	} else if seller.IsApproved {
		status.ApprovalStatus = "approved"
	} else {
		status.ApprovalStatus = "pending"
	}
	// Can add products only if approved
	status.CanAddProducts = seller.IsApproved && isComplete

	// Determine next step
	if !status.HasBusinessInfo || !status.HasAddress {
		status.NextStep = "complete_profile"
	} else if !status.HasPaymentMethod {
		status.NextStep = "add_payment"
	} else if !seller.IsApproved {
		status.NextStep = "wait_approval"
	} else {
		status.NextStep = "can_sell"
	}

	return status, nil

}
