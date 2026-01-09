package selleraccountsettingservice

import selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"

// GetLoginActivity retrieves recent login activities
func (s *SellerAccountSettingService) GetLoginActivity(sellerID uint) (*selleraccountsetting.LoginActivityResponse, error) {
	activities, err := s.repo.GetLoginActivities(sellerID, 10) // Last 10 activities
	if err != nil {
		return nil, err
	}

	var dtoActivities []selleraccountsetting.LoginActivity
	for _, activity := range activities {
		dtoActivities = append(dtoActivities, selleraccountsetting.LoginActivity{
			ID:        activity.ID,
			Device:    activity.Device,
			Browser:   activity.Browser,
			Location:  activity.Location,
			IPAddress: activity.IPAddress,
			TimeAgo:   s.calculateTimeAgo(activity.CreatedAt),
			IsCurrent: activity.IsCurrent,
			CreatedAt: activity.CreatedAt,
		})
	}

	return &selleraccountsetting.LoginActivityResponse{
		Activities: dtoActivities,
	}, nil
}
