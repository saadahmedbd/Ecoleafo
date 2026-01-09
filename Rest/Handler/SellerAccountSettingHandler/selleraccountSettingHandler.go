package selleraccountsettinghandler

import selleraccountsettingservice "github.com/saadahmedbd/Treestore/Rest/Service/SellerAccountSettingService"

type Selleraccountsettinghandler struct {
	service *selleraccountsettingservice.SellerAccountSettingService
}

func NewSelleraccountsettinghandler(service *selleraccountsettingservice.SellerAccountSettingService) *Selleraccountsettinghandler {
	return &Selleraccountsettinghandler{
		service: service,
	}
}
