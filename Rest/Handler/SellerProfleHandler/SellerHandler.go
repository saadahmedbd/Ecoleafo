package sellerproflehandler

import sellerprofileservice "github.com/saadahmedbd/Treestore/Rest/Service/sellerProfileService"

type SellerProfileHandler struct {
	service sellerprofileservice.SellerService
}

func NewSellerProfileHandler(service sellerprofileservice.SellerService) *SellerProfileHandler {
	return &SellerProfileHandler{
		service: service,
	}
}
