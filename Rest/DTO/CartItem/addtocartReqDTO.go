package cartitem

type AddToCartRequest struct {
	ProductID   uint   `json:"product_id" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required,min=1"`
	IsGift      bool   `json:"is_gift"`
	GiftMessage string `json:"gift_message"`
}
