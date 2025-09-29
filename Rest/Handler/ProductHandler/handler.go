package producthandler

import productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"

type Handler struct {
	service *productservice.ProductService
}

func NewHandler(service *productservice.ProductService) *Handler {
	return &Handler{service: service}
}
