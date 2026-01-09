package order

type UpdateOrderStatusRequest struct {
	Status         string `json:"status" validate:"required,oneof=pending processing shipped delivered cancelled"`
	TrackingNumber string `json:"tracking_number"`
	Comment        string `json:"comment"`
}
