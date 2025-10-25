package buyerprofilehandler

import (
	cloudniaryservice "github.com/saadahmedbd/Treestore/Rest/Service/CloudniaryService"
	buyerservice "github.com/saadahmedbd/Treestore/Rest/Service/buyerService"
)

type Buyerprofilehandler struct {
	buyerservice      buyerservice.BuyerService
	cloudniaryservice cloudniaryservice.CloudniaryService
}

func NewBuyerProfileHandler(buyerservice buyerservice.BuyerService, cloudniaryservice cloudniaryservice.CloudniaryService) *Buyerprofilehandler {
	return &Buyerprofilehandler{buyerservice: buyerservice,
		cloudniaryservice: cloudniaryservice,
	}
}
