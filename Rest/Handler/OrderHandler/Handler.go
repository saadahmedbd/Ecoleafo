package orderhandler

import orderservice "github.com/saadahmedbd/Treestore/Rest/Service/OrderService"

type OrderHandler struct {
	orderService orderservice.OrderService
}

func NewOrderHandler(orederService orderservice.OrderService) *OrderHandler {

	return &OrderHandler{
		orderService: orederService,
	}
}
