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
		imageURL := ""
		if len(item.Product.Images) > 0 {
			imageURL = item.Product.Images[0].ImageURL
		}

		itemResponses = append(itemResponses, cartitem.CartItemResponse{
			ProductID:   item.ProductID,
			ProductName: item.Product.Name,
			ProductSlug: item.Product.Slug,
			Price:       item.Product.Price,
			Image:       imageURL,
			InStock:     item.Product.Quantity > 0,
		})
	}

	return &cartitem.WishlistResponse{
		Items:      itemResponses,
		TotalItems: len(items),
	}, nil
}
