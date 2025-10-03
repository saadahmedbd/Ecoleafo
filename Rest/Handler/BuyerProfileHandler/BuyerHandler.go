package buyerprofilehandler

import buyerservice "github.com/saadahmedbd/Treestore/Rest/Service/buyerService"

type Buyerprofilehandler struct {
	buyerservice buyerservice.BuyerService
}

func NewBuyerProfileHandler(buyerservice buyerservice.BuyerService) *Buyerprofilehandler {
	return &Buyerprofilehandler{buyerservice: buyerservice}
}
