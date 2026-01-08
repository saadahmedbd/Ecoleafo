package sellerdashboardservice

import (
	"errors"
)

// UpdateOrderStatus - Update order status by seller
func (s *DashboardService) UpdateOrderStatus(regUserID uint, orderID uint, status string, trackingNumber string, notes string) error {
	if regUserID == 0 || orderID == 0 {
		return errors.New("invalid parameters")
	}

	validStatuses := map[string]bool{
		"processing": true,
		"shipped":    true,
		"delivered":  true,
	}

	if !validStatuses[status] {
		return errors.New("invalid status. Allowed: processing, shipped, delivered")
	}

	return s.DashboardRepository.UpdateOrderStatus(regUserID, orderID, status, trackingNumber, notes)
}
