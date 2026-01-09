package cartservice

import (
	"fmt"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
)

func (s *cartService) AddToWishlist(buyerID uint, productID uint) error {
	var product models.Product
	if err := Config.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		return fmt.Errorf("product not found")
	}

	wishlist := &models.Wishlist{
		BuyerID:   buyerID,
		ProductID: productID,
		AddedFrom: "browse",
	}

	return s.cartRepo.AddToWishlist(wishlist)
}
