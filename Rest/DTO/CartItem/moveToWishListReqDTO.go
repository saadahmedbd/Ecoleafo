package cartitem

type MoveToWishlistRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
}
