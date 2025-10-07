package selleraccountservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
)

func (s *sellerRegistrationService) mapToSellerProfileResponse(seller *models.User) *selleraccount.SellerProfileResponse {
	response := &selleraccount.SellerProfileResponse{
		ID:              seller.ID,
		UserId:          seller.UserId,
		BusinessEmail:   seller.BusinessEmail,
		Phone:           seller.Phone,
		StoreName:       seller.StoreName,
		StoreSlug:       seller.StoreSlug,
		StoreDesc:       seller.StoreDesc,
		StoreLogo:       seller.StoreLogo,
		StoreBanner:     seller.StoreBanner,
		BusinessType:    seller.BusinessType,
		TaxNumber:       seller.TaxNumber,
		BusinessLicense: seller.BusinessLicense,
		TotalSales:      seller.TotalSales,
		TotalEarnings:   seller.TotalEarnings,
		TotalOrders:     seller.TotalOrders,
		AverageRating:   seller.AverageRating,
		Commission:      seller.Commission,
		Address:         seller.Address,
		City:            seller.City,
		State:           seller.State,
		Country:         seller.Country,
		PostalCode:      seller.PostalCode,
		IsActive:        seller.IsActive,
		IsVerified:      seller.IsVerified,
		IsApproved:      seller.IsApproved,
		ApprovedAt:      seller.ApprovedAt,
		RejectedAt:      seller.RejectedAt,
		RejectionReason: seller.RejectionReason,
		CreatedAt:       seller.CreatedAt,
		UpdatedAt:       seller.UpdatedAt,
		ProductCount:    len(seller.Products),
		RegUser: sellerprofile.RegUserBasicInfo{
			FirstName: seller.RegUser.FirstName,
			LastName:  seller.RegUser.LastName,
			Email:     seller.RegUser.Email,
		},
	}
	for _, method := range seller.PaymentMethods {
		response.PaymentMethods = append(response.PaymentMethods, selleraccount.PaymentMethodInfo{
			ID:            method.ID,
			Type:          method.Type,
			AccountName:   method.AccountName,
			AccountNumber: method.AccountNumber,
			BankName:      method.BankName,
			IsDefault:     method.IsDefault,
			IsActive:      method.IsActive,
		})
	}
	return response
}
