package cartservice

import (
	cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"
)

func (s *cartService) GetWishlist(buyerID uint) (*cartitem.WishlistResponse, error) {
	items, err := s.cartRepo.GetWishlist(buyerID)
	if err != nil {
		return nil, err
	}

	var itemResponses []cartitem.CartItemResponse
	for _, item := range items {
		var imageURLs []string
		for _, img := range item.Product.Images {
			imageURLs = append(imageURLs, img.ImageURL)
		}

		sellerName := ""
		sellerID := uint(0)
		if item.Product.Seller.ID > 0 {
			sellerName = item.Product.Seller.StoreName
			sellerID = item.Product.Seller.ID
		}

		inStock := item.Product.IsActive && item.Product.Quantity >= 1
		available := inStock //that time i can't check product approveness i wil check is future

		availMsg := "In Stock"
		if !available {
			availMsg = "Out of Stock"
		}

		currentPrice := item.Product.Price
		if item.Product.DiscountPrice > 0 {
			currentPrice = item.Product.DiscountPrice
		}

		discountPercent := float64(0)
		if item.Product.DiscountPrice > 0 && item.Product.Price > 0 {
			discountPercent = ((item.Product.Price - item.Product.DiscountPrice) / item.Product.Price) * 100
		}

		itemResponses = append(itemResponses, cartitem.CartItemResponse{
			ID:              item.ID,
			ProductID:       item.ProductID,
			ProductName:     item.Product.Name,
			ProductSlug:     item.Product.Slug,
			Price:           currentPrice,
			DiscountPrice:   float32(item.Product.DiscountPrice),
			OriginalPrice:   item.Product.Price,
			DiscountPercent: discountPercent,
			Quantity:        item.Product.Quantity,
			Subtotal:        currentPrice,
			Image:           imageURLs,
			SellerName:      sellerName,
			SellerID:        sellerID,
			InStock:         inStock,
			StockQuantity:   item.Product.Quantity,
			IsAvailable:     available,
			AvailabilityMsg: availMsg,
			IsGift:          false,
			GiftMessage:     "",
			IsSavedForLater: false,
			CanIncreaseQty:  item.Product.Quantity > 1,
			MaxQuantity:     item.Product.Quantity,
		})
	}

	return &cartitem.WishlistResponse{
		Items:      itemResponses,
		TotalItems: len(items),
	}, nil
}
