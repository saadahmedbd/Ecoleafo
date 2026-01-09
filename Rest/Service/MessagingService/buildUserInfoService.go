package messagingservice

import (
	"fmt"
	
	messagingdto "github.com/saadahmedbd/Treestore/Rest/DTO/MessagingDTO"
)

func (s *messagingServiceImpl) buildUserInfo(userID uint, userType string) messagingdto.ConversationUserInfo {
	info := messagingdto.ConversationUserInfo{
		ID:       userID,
		UserType: userType,
	}

	switch userType {
	case "buyer":
		buyer, err := s.messagerepo.GetBuyerInfo(userID)
		fmt.Printf("[DEBUG] GetBuyerInfo - ID: %d, Error: %v, Buyer: %+v\n", userID, err, buyer)
		if err == nil && buyer != nil {
			fmt.Printf("[DEBUG] Buyer RegUser: %+v\n", buyer.RegUser)
			if buyer.RegUser != nil {
				info.Name = buyer.RegUser.FirstName + " " + buyer.RegUser.LastName
				info.Email = buyer.RegUser.Email
				info.ProfilePicture = buyer.ProfilePictureUrl
			}
		}
	case "seller":
		seller, err := s.messagerepo.GetSellerInfo(userID)
		if err == nil && seller != nil {
			info.StoreName = seller.StoreName
			info.ProfilePicture = seller.StoreLogo
			info.Name = seller.StoreName
			info.Email = seller.BusinessEmail
		}
	case "admin":
		admin, err := s.messagerepo.GetAdminInfo(userID)
		if err == nil && admin != nil && admin.RegUser.Email != "" {
			info.Name = admin.FullName
			info.Email = admin.RegUser.Email
		}
	}

	return info
}
