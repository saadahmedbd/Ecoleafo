package commissionpayoutearningservice

import (
	"time"

	commissionearningpayoutdto "github.com/saadahmedbd/Treestore/Rest/DTO/CommissionEarningPayoutDTO"
)

func (s *commissionService) GetCommissionSettings() (*commissionearningpayoutdto.CommissionSettingResponse, error) {
	settings, err := s.commissionRepo.GetCommissionSettings()
	if err != nil {
		return nil, err
	}

	return &commissionearningpayoutdto.CommissionSettingResponse{
		ID:          settings.ID,
		DefaultRate: settings.DefaultRate,
		Description: settings.Description,
		IsActive:    settings.IsActive,
		UpdatedBy:   settings.UpdatedBy,
		UpdatedAt:   settings.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *commissionService) UpdateCommissionSettings(req *commissionearningpayoutdto.UpdateCommissionSettingRequest, adminID uint) error {
	settings, err := s.commissionRepo.GetCommissionSettings()
	if err != nil {
		return err
	}

	settings.DefaultRate = req.DefaultRate
	settings.Description = req.Description
	settings.UpdatedBy = adminID

	return s.commissionRepo.UpdateCommissionSettings(settings)
}
