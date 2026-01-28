package util

import "testing"

func TestDeliveryChargeCalculator_Sirajganj(t *testing.T) {
	calc := NewDeliveryChargeCalculator()
	
	tests := []struct {
		weight   float64
		expected float64
	}{
		{0.5, 50.0},
		{1.0, 50.0},
		{1.5, 70.0},
		{2.0, 70.0},
		{3.0, 90.0},
		{3.5, 110.0},  // 50 + (3 * 20)
		{4.0, 150.0},  // pickup point
		{5.0, 150.0},  // pickup point
	}
	
	for _, tt := range tests {
		result := calc.CalculateDeliveryCharge("Sirajganj, Bangladesh", tt.weight)
		if result != tt.expected {
			t.Errorf("Sirajganj %.1fkg: expected %.2f, got %.2f", tt.weight, tt.expected, result)
		}
	}
}

func TestDeliveryChargeCalculator_Dhaka(t *testing.T) {
	calc := NewDeliveryChargeCalculator()
	
	tests := []struct {
		weight   float64
		expected float64
	}{
		{0.5, 110.0},
		{1.0, 110.0},
		{1.5, 130.0},
		{2.0, 130.0},
		{3.0, 150.0},  // 110 + (2 * 20)
		{3.5, 170.0},  // 110 + (3 * 20)
		{4.0, 150.0},  // pickup point
		{5.0, 150.0},  // pickup point
	}
	
	for _, tt := range tests {
		result := calc.CalculateDeliveryCharge("Dhaka, Bangladesh", tt.weight)
		if result != tt.expected {
			t.Errorf("Dhaka %.1fkg: expected %.2f, got %.2f", tt.weight, tt.expected, result)
		}
	}
}

func TestDeliveryChargeCalculator_OtherDistricts(t *testing.T) {
	calc := NewDeliveryChargeCalculator()
	
	tests := []struct {
		weight   float64
		expected float64
	}{
		{0.3, 110.0},  // rounds to 1 unit (0.5kg)
		{0.5, 110.0},  // 1 unit
		{0.8, 130.0},  // rounds to 2 units (1.0kg)
		{1.0, 130.0},  // 2 units
		{1.5, 150.0},  // 3 units
		{2.0, 170.0},  // 4 units
		{3.0, 210.0},  // 6 units: 110 + (5 × 20)
		{3.5, 230.0},  // 7 units: 110 + (6 × 20)
		{4.0, 150.0},  // pickup point
		{5.0, 150.0},  // pickup point
	}
	
	for _, tt := range tests {
		result := calc.CalculateDeliveryCharge("Chittagong, Bangladesh", tt.weight)
		if result != tt.expected {
			t.Errorf("Other District %.1fkg: expected %.2f, got %.2f", tt.weight, tt.expected, result)
		}
	}
}
