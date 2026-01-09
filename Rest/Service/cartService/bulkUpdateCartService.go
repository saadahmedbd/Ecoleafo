package cartservice

import cartitem "github.com/saadahmedbd/Treestore/Rest/DTO/CartItem"

func (s *cartService) BulkUpdateCart(buyerID uint, req cartitem.BulkUpdateCartRequest) error {
	updates := make(map[uint]int)
	for _, item := range req.Item {
		updates[item.ProductID] = item.Quantity
	}
	return s.cartRepo.BulkUpdateQuantities(buyerID, updates)
}
