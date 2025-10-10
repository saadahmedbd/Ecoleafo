package cartservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
)

func (s *cartService) GetSavedForLater(buyerID uint) ([]models.CartItem, error) {
	return s.cartRepo.GetSavedForLater(buyerID)
}
