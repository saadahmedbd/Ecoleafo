package selleraccountsetting

import models "github.com/saadahmedbd/Treestore/Models"

// ToSellerProfileResponse - Convert User model to SellerProfileResponse
func ToSellerProfileResponse(user *models.User, regUser *models.RegUser) *SellerProfileResponse {
	response := &SellerProfileResponse{
		// Basic Info
		ID:           user.ID,
		FirstName:    regUser.FirstName,
		LastName:     regUser.LastName,
		Email:        regUser.Email,
		Phone:        user.Phone,
		ProfilePhoto: regUser.ProfilePhoto,

		// Store Info
		StoreName:        user.StoreName,
		StoreSlug:        user.StoreSlug,
		StoreDescription: user.StoreDesc,
		StoreLogo:        user.StoreLogo,
		StoreBanner:      user.StoreBanner,

		// Business Info
		BusinessEmail:   user.BusinessEmail,
		BusinessType:    user.BusinessType,
		TaxNumber:       user.TaxNumber,
		BusinessLicense: user.BusinessLicense,

		// Address
		Address:    user.Address,
		City:       user.City,
		State:      user.State,
		Country:    user.Country,
		PostalCode: user.PostalCode,

		// Statistics
		TotalSales:    user.TotalSales,
		TotalEarnings: user.TotalEarnings,
		TotalOrders:   user.TotalOrders,
		AverageRating: user.AverageRating,
		Commission:    user.Commission,

		// Status
		IsActive:        user.IsActive,
		IsVerified:      user.IsVerified,
		IsApproved:      user.IsApproved,
		ApprovedAt:      user.ApprovedAt,
		RejectedAt:      user.RejectedAt,
		RejectionReason: user.RejectionReason,

		// Timestamps
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response
}
