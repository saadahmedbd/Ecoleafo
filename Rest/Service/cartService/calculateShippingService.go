package cartservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

// calculateShipping calculates shipping cost based on address and cart weight
func (s *cartService) calculateShipping(address string, items []models.CartItem) float64 {
	// If no address provided, return default shipping
	if address == "" {
		return 100.0
	}

	// Calculate total weight from cart items
	var totalWeight float64
	for _, item := range items {
		if item.Product.Weight > 0 {
			totalWeight += item.Product.Weight * float64(item.Quantity)
		}
	}

	// Use delivery charge calculator
	deliveryCalc := util.NewDeliveryChargeCalculator()
	shippingCost := deliveryCalc.CalculateDeliveryCharge(address, totalWeight)

	return shippingCost
}
