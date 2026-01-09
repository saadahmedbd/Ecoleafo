package selleraccounthandler

import selleraccountservice "github.com/saadahmedbd/Treestore/Rest/Service/SellerAccountService"

type SellerRegistrationHandler struct {
	service selleraccountservice.SellerRegistrationService
}

func NewSellerRegistrationHandler(service selleraccountservice.SellerRegistrationService) *SellerRegistrationHandler {
	return &SellerRegistrationHandler{service: service}
}
