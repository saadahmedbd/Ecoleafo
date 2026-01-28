# Delivery Charge System Documentation

## Overview
District-based delivery charge calculation system with weight-based pricing.

## Pricing Rules

### 1. Sirajganj District
- **Base Rate**: 50 taka for up to 1 kg
- **Additional Weight**: +20 taka per kg after 1 kg
- **4kg or More**: 150 taka (pickup point delivery)

**Examples:**
- 0.5 kg → 50 taka
- 1.0 kg → 50 taka
- 1.5 kg → 70 taka (50 + 20)
- 2.0 kg → 70 taka (50 + 20)
- 3.0 kg → 90 taka (50 + 40)
- 3.5 kg → 110 taka (50 + 60)
- 4.0 kg+ → 150 taka (pickup point)

### 2. Dhaka District
- **Base Rate**: 110 taka for up to 1 kg
- **Additional Weight**: +20 taka per kg after 1 kg
- **4kg or More**: 150 taka (pickup point delivery)

**Examples:**
- 0.5 kg → 110 taka
- 1.0 kg → 110 taka
- 1.5 kg → 130 taka (110 + 20)
- 2.0 kg → 130 taka (110 + 20)
- 3.0 kg → 150 taka (110 + 40)
- 3.5 kg → 170 taka (110 + 60)
- 4.0 kg+ → 150 taka (pickup point)

### 3. Other Districts
- **Base Rate**: 110 taka per 0.5 kg
- **Additional Weight**: +20 taka per 0.5 kg increment
- **4kg or More**: 150 taka (pickup point delivery)

**Examples:**
- 0.3 kg → 110 taka (rounded to 0.5 kg)
- 0.5 kg → 110 taka
- 0.8 kg → 130 taka (110 + 20 for second 0.5kg)
- 1.0 kg → 130 taka
- 1.5 kg → 150 taka (110 + 20 + 20)
- 2.0 kg → 170 taka (110 + 60)
- 3.0 kg → 210 taka (110 + 100)
- 3.5 kg → 230 taka (110 + 120)
- 4.0 kg+ → 150 taka (pickup point)

## Implementation

### Location Detection
The system automatically detects the district from the shipping address:
- Searches for "sirajganj" or "sirajgonj" in address
- Searches for "dhaka" or "daka" in address
- All other addresses treated as "other districts"

### Weight Calculation
- Total weight = Sum of (Product Weight × Quantity) for all items
- Weight is in kilograms (kg)
- Ensure all products have accurate weight values

## Integration Points

### 1. Order Creation
- Automatically calculated during order creation
- Based on shipping address and cart items
- Stored in `Order.ShippingCost` field

### 2. Buyer View
- Displayed in cart summary
- Shown in order confirmation
- Included in order history

### 3. Seller View
- Visible in order details
- Included in seller earnings calculation
- Shown in order management

### 4. Admin View
- Displayed in order management
- Included in financial reports
- Visible in order analytics

## Database Fields

### Order Model
```go
ShippingCost float64 `json:"shipping_cost" gorm:"type:decimal(10,2);default:0"`
```

### Product Model (Required)
```go
Weight float64 `json:"weight" gorm:"type:decimal(8,2)"` // in kg
```

## Testing

### Test Cases
1. **Sirajganj - Light**: 0.5 kg → 50 taka ✓
2. **Sirajganj - Medium**: 2 kg → 70 taka ✓
3. **Sirajganj - Heavy**: 5 kg → 150 taka (capped) ✓
4. **Dhaka - Light**: 0.5 kg → 110 taka ✓
5. **Dhaka - Medium**: 2 kg → 130 taka ✓
6. **Dhaka - Heavy**: 5 kg → 150 taka (capped) ✓
7. **Other - Light**: 0.5 kg → 110 taka ✓
8. **Other - Medium**: 1.5 kg → 150 taka ✓
9. **Other - Heavy**: 3 kg → 230 taka ✓

## Important Notes

1. **Product Weight Required**: All products MUST have weight values set
2. **Address Format**: Ensure shipping addresses include district names
3. **Case Insensitive**: District detection is case-insensitive
4. **Rounding**: Weight is rounded up to nearest unit (kg or 0.5kg)
5. **No Free Shipping**: Delivery charges always apply

## Future Enhancements

- [ ] Add delivery charge preview in cart
- [ ] Support for custom delivery zones
- [ ] Bulk order discounts
- [ ] Express delivery options
- [ ] Delivery charge calculator API endpoint
