package adminmangementservice

type DashboardService struct {
	buyerService   *BuyerService
	sellerService  *SellerService
	productService *ProductService
	orderService   *OrderService
}

func NewDashboardService(
	buyerService *BuyerService,
	sellerService *SellerService,
	productService *ProductService,
	orderService *OrderService,
) *DashboardService {
	return &DashboardService{
		buyerService:   buyerService,
		sellerService:  sellerService,
		productService: productService,
		orderService:   orderService,
	}
}

func (s *DashboardService) GetDashboardStats() (map[string]interface{}, error) {
	buyerStats, _ := s.buyerService.GetBuyerStats()
	sellerStats, _ := s.sellerService.GetSellerStats()
	productStats, _ := s.productService.GetProductStats()
	orderStats, _ := s.orderService.GetOrderStats()

	return map[string]interface{}{
		"users":    buyerStats,
		"sellers":  sellerStats,
		"products": productStats,
		"orders":   orderStats,
	}, nil
}
