package cartitem

type UpdateCartItemRequest struct {
	Quantity    int    `json:"quantity" validate:"required,min=1"`
	IsGift      bool   `json:"gift"`
	GiftMessage string `json:"gift_message"`
}
