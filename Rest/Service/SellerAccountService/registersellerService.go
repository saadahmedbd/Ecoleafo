package selleraccountservice

import (
	"fmt"
	"time"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	selleraccount "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccount"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (s *sellerRegistrationService) RegisterSeller(req *selleraccount.SellerRegistationRequest) (*selleraccount.SellerRegistrationResponse, error) {
	//validate password match
	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("passwords do not match")
	}
	//check if email already exists
	existingSeller, _ := s.sellerRepo.GetSellerByEmail(req.Email)
	if existingSeller != nil {
		return nil, fmt.Errorf("email already Registered")
	}
	//check if email exists in regUser
	var existingRegUser models.RegUser
	if err := Config.DB.Where("email = ?", req.Email).First(&existingRegUser).Error; err == nil {
		return nil, fmt.Errorf("email already Registered")
	}
	//hashPassword
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}
	// get seller role
	var sellerRole models.Role
	if err := Config.DB.Where("name = ?", "seller").First(&sellerRole).Error; err != nil {
		return nil, fmt.Errorf("seller role not found")
	}
	//create regUser
	regUser := models.RegUser{
		Email:     req.Email,
		Password:  hashPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsActive:  true,
	}
	if err := s.sellerRepo.CreateRegUser(&regUser); err != nil {
		return nil, fmt.Errorf("failed to create regUser: %v", err)
	}
	//generate store slug
	storeSlug := util.GenerateSlug(req.StoreName)
	//validate and set commission
	commission := req.Commission
	if commission == 0 {
		commission = 15 // Default to 15% if not provided
	}
	if commission < 15 {
		return nil, fmt.Errorf("commission must be at least 15%%")
	}
	//create seller profile
	seller := models.User{
		RoleID:        sellerRole.ID,
		UserId:        regUser.ID,
		BusinessEmail: req.Email,
		Password:      hashPassword,
		Phone:         req.Phone,
		StoreName:     req.StoreName,
		StoreSlug:     storeSlug,
		Commission:    commission,
		IsActive:      true,
		IsVerified:    false,
		IsApproved:    false,
	}
	if err := s.sellerRepo.CreateSeller(&seller); err != nil {
		return nil, fmt.Errorf("failed to create seller profile: %v", err)
	}
	//generate jwt token
	token, err := util.CreateJwt(regUser.ID, req.FirstName, req.LastName, []string{"seller"}, seller.ID, 24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}
	//check profile status
	profileStatus, _ := s.CheckProfileStatus(regUser.ID)
	return &selleraccount.SellerRegistrationResponse{
		Message:       "Seller Registered successfully,please complete your profile to start selling. ",
		SellerID:      seller.ID,
		UserID:        regUser.ID,
		Token:         token,
		ProfileStatus: profileStatus,
		NeedsSetup:    true,
	}, nil

}
