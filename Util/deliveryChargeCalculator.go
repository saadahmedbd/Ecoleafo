package util

import (
	"math"
	"strings"
)

type DeliveryChargeCalculator struct{}

type DeliveryOption struct {
	Type        string  `json:"type"`        // "home_delivery" or "pickup_point"
	Charge      float64 `json:"charge"`
	Available   bool    `json:"available"`
	Description string  `json:"description"`
	Savings     float64 `json:"savings,omitempty"`
}

type DeliveryCalculationResult struct {
	TotalWeight     float64          `json:"total_weight"`
	DeliveryOptions []DeliveryOption `json:"delivery_options"`
	HomeDelivery    float64          `json:"home_delivery_charge"`
	PickupPoint     float64          `json:"pickup_point_charge"`
}

func NewDeliveryChargeCalculator() *DeliveryChargeCalculator {
	return &DeliveryChargeCalculator{}
}

// CalculateDeliveryCharge calculates delivery charge based on location and weight
func (d *DeliveryChargeCalculator) CalculateDeliveryCharge(address string, totalWeight float64) float64 {
	district := d.extractDistrict(address)

	// Sirajganj: 50 taka per kg, +20 taka per kg after 1kg, max 150 for 4kg+
	if district == "sirajganj" {
		return d.calculateSirajganjCharge(totalWeight)
	}

	// Dhaka: 110 taka base, +20 taka per kg after 1kg, max 150 for 4kg+
	if district == "dhaka" {
		return d.calculateDhakaCharge(totalWeight)
	}

	// Other districts: 110 taka per 0.5kg, +20 taka per 0.5kg increment
	return d.calculateOtherDistrictCharge(totalWeight)
}

// GetDeliveryOptions returns available delivery options with pricing
func (d *DeliveryChargeCalculator) GetDeliveryOptions(address string, totalWeight float64) *DeliveryCalculationResult {
	homeDeliveryCharge := d.CalculateDeliveryCharge(address, totalWeight)
	pickupCharge := 150.0

	result := &DeliveryCalculationResult{
		TotalWeight:  totalWeight,
		HomeDelivery: homeDeliveryCharge,
		PickupPoint:  pickupCharge,
	}

	// Weight >= 4kg: Only pickup available
	if totalWeight >= 4.0 {
		result.DeliveryOptions = []DeliveryOption{
			{
				Type:        "pickup_point",
				Charge:      pickupCharge,
				Available:   true,
				Description: "Collect from nearest pickup point (Heavy order)",
			},
		}
		return result
	}

	// Weight < 4kg: Both options available
	options := []DeliveryOption{
		{
			Type:        "home_delivery",
			Charge:      homeDeliveryCharge,
			Available:   true,
			Description: "Delivered to your doorstep",
		},
		{
			Type:        "pickup_point",
			Charge:      pickupCharge,
			Available:   true,
			Description: "Collect from nearest pickup point",
		},
	}

	// Add savings if pickup is cheaper
	if homeDeliveryCharge > pickupCharge {
		options[1].Savings = homeDeliveryCharge - pickupCharge
	}

	result.DeliveryOptions = options
	return result
}

func (d *DeliveryChargeCalculator) calculateSirajganjCharge(weight float64) float64 {
	// For 4kg+: 150 taka (pickup point delivery)
	if weight >= 4.0 {
		return 150.0
	}

	if weight <= 1.0 {
		return 50.0
	}

	// After 1kg: 50 + (additional kg * 20)
	additionalKg := math.Ceil(weight - 1.0)
	charge := 50.0 + (additionalKg * 20.0)

	return charge
}

func (d *DeliveryChargeCalculator) calculateDhakaCharge(weight float64) float64 {
	// For 4kg+: 150 taka (pickup point delivery)
	if weight >= 4.0 {
		return 150.0
	}

	if weight <= 1.0 {
		return 110.0
	}

	// After 1kg: 110 + (additional kg * 20)
	additionalKg := math.Ceil(weight - 1.0)
	charge := 110.0 + (additionalKg * 20.0)

	return charge
}

func (d *DeliveryChargeCalculator) calculateOtherDistrictCharge(weight float64) float64 {
	// For 4kg+: 150 taka (pickup point delivery)
	if weight >= 4.0 {
		return 150.0
	}

	// 110 taka per 0.5kg
	halfKgUnits := math.Ceil(weight / 0.5)

	if halfKgUnits <= 1 {
		return 110.0
	}

	// First 0.5kg: 110, then +20 per additional 0.5kg
	charge := 110.0 + ((halfKgUnits - 1) * 20.0)

	return charge
}

func (d *DeliveryChargeCalculator) extractDistrict(address string) string {
	addressLower := strings.ToLower(address)

	if strings.Contains(addressLower, "sirajganj") || strings.Contains(addressLower, "sirajgonj") {
		return "sirajganj"
	}

	if strings.Contains(addressLower, "dhaka") || strings.Contains(addressLower, "daka") {
		return "dhaka"
	}

	return "other"
}
