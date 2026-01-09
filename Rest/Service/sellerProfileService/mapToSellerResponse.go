package sellerprofileservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	sellerprofile "github.com/saadahmedbd/Treestore/Rest/DTO/SellerProfile"
)

func (s *sellerService) mapToSellerProfileResponse(seller *models.User) *sellerprofile.SellerProfileResponse {
	return &sellerprofile.SellerProfileResponse{
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
		Address:         seller.Address,
		City:            seller.City,
		State:           seller.State,
		Country:         seller.Country,
		PostalCode:      seller.PostalCode,
		IsActive:        seller.IsActive,
		IsVerified:      seller.IsVerified,
		IsApproved:      seller.IsApproved,
		CreatedAt:       seller.CreatedAt,
		UpdatedAt:       seller.UpdatedAt,
		ProductCount:    len(seller.Products),
		RegUser: sellerprofile.RegUserBasicInfo{
			FirstName: seller.RegUser.FirstName,
			LastName:  seller.RegUser.LastName,
			Email:     seller.RegUser.Email,
		},
	}
}
