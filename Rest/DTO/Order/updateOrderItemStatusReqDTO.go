package order

type UpdateOrderItemStatusRequest struct {
	ItemID  uint   `json:"item_id" validate:"required"`
	Status  string `json:"status" validate:"required,oneof=pending processing shipped delivered cancelled"`
	Comment string `json:"comment"`
}
