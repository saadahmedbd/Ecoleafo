package carthandler

import cartservice "github.com/saadahmedbd/Treestore/Rest/Service/cartService"

type CartHandler struct {
	service cartservice.CartService
}

func NewCartService(service cartservice.CartService) *CartHandler {
	return &CartHandler{
		service: service,
	}
}
