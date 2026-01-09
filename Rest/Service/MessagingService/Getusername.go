package messagingservice

func (s *messagingServiceImpl) getUserName(userID uint, userType string) string {
	switch userType {
	case "buyer":
		buyer, err := s.messagerepo.GetBuyerInfo(userID)
		if err == nil && buyer.RegUser != nil {
			return buyer.RegUser.FirstName + " " + buyer.RegUser.LastName
		}
	case "seller":
		seller, err := s.messagerepo.GetSellerInfo(userID)
		if err == nil && seller.RegUser != nil {
			return seller.StoreName
		}
	case "admin":
		admin, err := s.messagerepo.GetAdminInfo(userID)
		if err == nil {
			return admin.FullName
		}
	}
	return "Unknown User"
}
