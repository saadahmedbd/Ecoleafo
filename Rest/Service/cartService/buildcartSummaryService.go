package cartservice

import (
	"fmt"
	"math"

	models "github.com/saadahmedbd/Treestore/Models"
	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
)

func (s *cartService) buildCartSummary(buyerID uint, items []models.CartItem, address string) (*cartitem.CartSummaryResponse, error) {
	var cartItems []cartitem.CartItemResponse
	var savedItems []cartitem.CartItemResponse

	var subtotal, discount, totalSavings, shipping_cost, giftCharge float64

	hasUnavailable := false
	var unavailableItems []string
	hasGiftItems := false

	for _, item := range items {
		var imageURLs []string
		// Find the primary image
		for _, img := range item.Product.Images {
			if img.IsPrimary {
				imageURLs = append(imageURLs, img.ImageURL)
				break // Only include the primary image
			}
		}
		// If no primary image, use the first available image
		if len(imageURLs) == 0 && len(item.Product.Images) > 0 {
			imageURLs = append(imageURLs, item.Product.Images[0].ImageURL)
		}

		sellerName := ""
		sellerID := uint(0)
		if item.Product.Seller.ID > 0 {
			sellerName = item.Product.Seller.StoreName
			sellerID = item.Product.Seller.ID
		}

		inStock := item.Product.IsActive && item.Product.Quantity >= item.Quantity
		available := inStock //this time i can;t check is product approved or not(funture i will implement item.product.approved)

		availMsg := "In Stock"
		if !available {
			availMsg = "Out of Stock"
			hasUnavailable = true
			unavailableItems = append(unavailableItems, item.Product.Name)
		} else if item.Product.Quantity < item.Quantity {
			availMsg = fmt.Sprintf("Only %d available", item.Product.Quantity)
		}

		originalPrice := item.Product.Price
		itemSubtotal := originalPrice * float64(item.Quantity)
		itemDiscount := 0.0
		if item.Product.DiscountPrice > 0 {
			itemDiscount = (originalPrice - item.Product.DiscountPrice) * float64(item.Quantity)
		}

		itemResponse := cartitem.CartItemResponse{
			ID:            item.ID,
			ProductID:     item.ProductID,
			ProductName:   item.Product.Name,
			ProductSlug:   item.Product.Slug,
			Price:         item.Product.DiscountPrice,
			OriginalPrice: math.Round(originalPrice),
			DiscountPrice: (float32(item.Product.DiscountPrice)),
			DiscountPercent: func() float64 {
				if item.Product.DiscountPrice > 0 && originalPrice > 0 {
					return ((originalPrice - item.Product.DiscountPrice) / originalPrice) * 100
				}
				return 0
			}(),

			Quantity:        item.Quantity,
			Subtotal:        math.Round(itemSubtotal - itemDiscount),
			Image:           imageURLs,
			SellerName:      sellerName,
			SellerID:        sellerID,
			InStock:         inStock,
			StockQuantity:   item.Product.Quantity,
			IsAvailable:     available,
			AvailabilityMsg: availMsg,
			IsGift:          item.IsGift,
			GiftMessage:     item.GiftMessage,
			IsSavedForLater: item.IsSavedForLater,
			IsSelected:      item.IsSelected,
			CanIncreaseQty:  item.Quantity < item.Product.Quantity,
		}

		if item.IsSavedForLater {
			savedItems = append(savedItems, itemResponse)
		} else {
			cartItems = append(cartItems, itemResponse)
			subtotal += itemSubtotal
			discount += itemDiscount
			if item.IsGift && item.IsSelected {
				hasGiftItems = true
			}
		}
	}

	if hasGiftItems {
		giftCharge = 50.0
	}

	totalSavings = discount
	total := subtotal - discount

	// Calculate shipping cost based on address
	shipping_cost = s.calculateShipping(address, total)

	return &cartitem.CartSummaryResponse{
		Items:               cartItems,
		SavedForLater:       savedItems,
		Subtotal:            math.Round(subtotal),
		Discount:            math.Round(discount),
		ShippingCost:        shipping_cost,
		GiftCharge:          giftCharge,
		TotalAmount:         math.Round(shipping_cost + total + giftCharge),
		TotalSavings:        math.Round(totalSavings),
		ItemCount:           len(cartItems),
		SavedItemCount:      len(savedItems),
		HasUnavailableItems: hasUnavailable,
		UnavailableItems:    unavailableItems,
	}, nil
}
