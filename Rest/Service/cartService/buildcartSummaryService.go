package cartservice

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
)

func (s *cartService) buildCartSummary(buyerID uint, items []models.CartItem) (*cartitem.CartSummaryResponse, error) {
	var cartItems []cartitem.CartItemResponse
	var savedItems []cartitem.CartItemResponse

	var subtotal, discount, totalSavings float64
	hasUnavailable := false
	var unavailableItems []string

	for _, item := range items {
		imageURL := ""
		if len(item.Product.Images) > 0 {
			imageURL = item.Product.Images[0].ImageURL
		}

		sellerName := ""
		sellerID := uint(0)
		if item.Product.Seller.ID > 0 {
			sellerName = item.Product.Seller.StoreName
			sellerID = item.Product.Seller.ID
		}

		inStock := item.Product.IsActive && item.Product.Quantity >= item.Quantity
		available := inStock && item.Product.IsApproved

		availMsg := "In Stock"
		if !available {
			availMsg = "Out of Stock"
			hasUnavailable = true
			unavailableItems = append(unavailableItems, item.Product.Name)
		} else if item.Product.Quantity < item.Quantity {
			availMsg = fmt.Sprintf("Only %d available", item.Product.Quantity)
		}

		itemSubtotal := item.Price * float64(item.Quantity)
		itemDiscount := 0.0
		if item.Product.DiscountPrice > 0 {
			itemDiscount = (item.Price - item.Product.DiscountPrice) * float64(item.Quantity)
		}

		itemResponse := cartitem.CartItemResponse{
			ID:              item.ID,
			ProductID:       item.ProductID,
			ProductName:     item.Product.Name,
			ProductSlug:     item.Product.Slug,
			Price:           item.Price,
			OriginalPrice:   item.Product.Price,
			DiscountPercent: item.Product.DiscountPercent,
			Quantity:        item.Quantity,
			Subtotal:        itemSubtotal,
			Image:           imageURL,
			SellerName:      sellerName,
			SellerID:        sellerID,
			InStock:         inStock,
			StockQuantity:   item.Product.Quantity,
			IsAvailable:     available,
			AvailabilityMsg: availMsg,
			IsGift:          item.IsGift,
			GiftMessage:     item.GiftMessage,
			IsSavedForLater: item.IsSavedForLater,
			CanIncreaseQty:  item.Quantity < item.Product.Quantity,
		}

		if item.IsSavedForLater {
			savedItems = append(savedItems, itemResponse)
		} else {
			cartItems = append(cartItems, itemResponse)
			subtotal += itemSubtotal
			discount += itemDiscount
		}
	}

	totalSavings = discount
	total := subtotal - discount

	return &cartitem.CartSummaryResponse{
		Items:               cartItems,
		SavedForLater:       savedItems,
		Subtotal:            subtotal,
		Discount:            discount,
		TotalAmount:         total,
		TotalSavings:        totalSavings,
		ItemCount:           len(cartItems),
		SavedItemCount:      len(savedItems),
		HasUnavailableItems: hasUnavailable,
		UnavailableItems:    unavailableItems,
	}, nil
}
