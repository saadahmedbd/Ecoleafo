package CMD

import (
	"github.com/saadahmedbd/Treestore/Config"
	rest "github.com/saadahmedbd/Treestore/Rest"
	adminhandler "github.com/saadahmedbd/Treestore/Rest/Handler/AdminHandler"
	adminreviewhandler "github.com/saadahmedbd/Treestore/Rest/Handler/AdminReviewHandler"
	adminmangementhandler "github.com/saadahmedbd/Treestore/Rest/Handler/AdminmangementHandler"
	auditloghandler "github.com/saadahmedbd/Treestore/Rest/Handler/AuditLogHandler"
	authhandler "github.com/saadahmedbd/Treestore/Rest/Handler/AuthHandler"
	buyerhandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerHandler"
	buyerprofilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerProfileHandler"
	cartitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/CartItemHandler"
	categoryHandler "github.com/saadahmedbd/Treestore/Rest/Handler/CategoryHandler"
	commissionpayoutearningshandler "github.com/saadahmedbd/Treestore/Rest/Handler/CommissionPayoutEarningsHandler"
	guestcarthandler "github.com/saadahmedbd/Treestore/Rest/Handler/GuestcartHandler"
	inventoryhandler "github.com/saadahmedbd/Treestore/Rest/Handler/InventoryHandler"
	orderhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderHandler"
	orderitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderItemHandler"
	producthandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProductHandler"
	profilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProfileHandler"
	reviewhandler "github.com/saadahmedbd/Treestore/Rest/Handler/ReviewHandler"
	rolehandler "github.com/saadahmedbd/Treestore/Rest/Handler/RoleHandler"
	searchhandler "github.com/saadahmedbd/Treestore/Rest/Handler/SearchHandler"
	selleraccountsettinghandler "github.com/saadahmedbd/Treestore/Rest/Handler/SellerAccountSettingHandler"
	sellerdashboardhandler "github.com/saadahmedbd/Treestore/Rest/Handler/SellerDashboardHandler"
	sellerproflehandler "github.com/saadahmedbd/Treestore/Rest/Handler/SellerProfleHandler"
	userhandler "github.com/saadahmedbd/Treestore/Rest/Handler/UserHandler"
	carthandler "github.com/saadahmedbd/Treestore/Rest/Handler/cartHandler"
	selleraccounthandler "github.com/saadahmedbd/Treestore/Rest/Handler/sellerAccountHandler"
	repository "github.com/saadahmedbd/Treestore/Rest/Repository"
	adminmangement "github.com/saadahmedbd/Treestore/Rest/Repository/AdminMangement"
	adminrepo "github.com/saadahmedbd/Treestore/Rest/Repository/AdminRepo"
	auditlogrepo "github.com/saadahmedbd/Treestore/Rest/Repository/AuditLogRepo"
	buyercompleterepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerCompleteRepo"
	buyerProfilerepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerProfileRepo"
	cartitemrepo "github.com/saadahmedbd/Treestore/Rest/Repository/CartItemRepo"
	categoryrepo "github.com/saadahmedbd/Treestore/Rest/Repository/CategoryRepo"
	commissionpayoutearningrepo "github.com/saadahmedbd/Treestore/Rest/Repository/CommissionPayoutEarningRepo"
	guestcartrepo "github.com/saadahmedbd/Treestore/Rest/Repository/GuestCartRepo"
	inventoryrepo "github.com/saadahmedbd/Treestore/Rest/Repository/InventoryRepo"
	orderrepo "github.com/saadahmedbd/Treestore/Rest/Repository/OrderRepo"
	productrepo "github.com/saadahmedbd/Treestore/Rest/Repository/ProductRepo"
	reguserrepo "github.com/saadahmedbd/Treestore/Rest/Repository/RegUserRepo"
	reviewrepo "github.com/saadahmedbd/Treestore/Rest/Repository/ReviewRepo"
	selleraccountrepo "github.com/saadahmedbd/Treestore/Rest/Repository/SellerAccountRepo"
	sellerdashboardrepo "github.com/saadahmedbd/Treestore/Rest/Repository/SellerDashboardRepo"
	adminreviewrepo "github.com/saadahmedbd/Treestore/Rest/Repository/adminReviewRepo"
	selleraccountsettingrepo "github.com/saadahmedbd/Treestore/Rest/Repository/sellerAccountSettingRepo"
	adminmangementservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminMangementService"
	adminreviewservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminReviewService"
	adminservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminService"
	audithelper "github.com/saadahmedbd/Treestore/Rest/Service/AuditHelper"
	auditlogservice "github.com/saadahmedbd/Treestore/Rest/Service/AuditLogService"
	commissionpayoutearningservice "github.com/saadahmedbd/Treestore/Rest/Service/CommissionPayoutEarningService"
	inventoryservice "github.com/saadahmedbd/Treestore/Rest/Service/InventoryService"
	orderservice "github.com/saadahmedbd/Treestore/Rest/Service/OrderService"
	productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"
	reviewservice "github.com/saadahmedbd/Treestore/Rest/Service/ReviewService"
	selleraccountservice "github.com/saadahmedbd/Treestore/Rest/Service/SellerAccountService"
	selleraccountsettingservice "github.com/saadahmedbd/Treestore/Rest/Service/SellerAccountSettingService"
	sellerdashboardservice "github.com/saadahmedbd/Treestore/Rest/Service/SellerDashboardService"
	buyerservice "github.com/saadahmedbd/Treestore/Rest/Service/buyerService"
	cartservice "github.com/saadahmedbd/Treestore/Rest/Service/cartService"
	categoryservice "github.com/saadahmedbd/Treestore/Rest/Service/categoryService"
	guestcartservice "github.com/saadahmedbd/Treestore/Rest/Service/guestCartService"
	sellerprofileservice "github.com/saadahmedbd/Treestore/Rest/Service/sellerProfileService"
)

func Server() {
	cnf := Config.GetConfig()

	//repo
	repository := repository.NewSellerRepostory(Config.DB)
	buyerProfilerepo := buyerProfilerepo.NewBuyerRepository(Config.DB)
	guestcartrepo := guestcartrepo.NewGuestCartRepository(Config.DB)
	buyercompleterepo := buyercompleterepo.NewProfileCompleteRepository(Config.DB)
	selleraccountrepo := selleraccountrepo.NewSellerRegistrationRepository(Config.DB)
	cartitemrepo := cartitemrepo.NewCartRepository(Config.DB)
	orderrepo := orderrepo.NewOrderRepository(Config.DB)
	reguserrepo := reguserrepo.NewRegUserRepository(Config.DB)
	adminrepo := adminrepo.NewAdminRepository(Config.DB)
	categoryrepo := categoryrepo.NewCategoryRepository(Config.DB)
	productrepo := productrepo.NewProductRepository(Config.DB)
	selleraccountsettingrepo := selleraccountsettingrepo.NewSellerAccountSettingRepositoryOptions(Config.DB)
	//admin management repository
	adminbuyerrepo := adminmangement.NewBuyerRepository(Config.DB)
	adminsellerrepo := adminmangement.NewSellerRepository(Config.DB)
	adminproductrepo := adminmangement.NewProductRepository(Config.DB)
	adminorderrepo := adminmangement.NewOrderRepository(Config.DB)
	adminauditlogrepo := adminmangement.NewAuditLogRepository(Config.DB)

	sellerdashboardrepo := sellerdashboardrepo.NewDashboardRepository(Config.DB)
	inventoryrepo := inventoryrepo.NewInventoryRepository(Config.DB)
	reviewrepo := reviewrepo.NewReviewRepository(Config.DB)
	commissionrepo := commissionpayoutearningrepo.NewCommissionRepository(Config.DB)
	adminreviewrepo := adminreviewrepo.NewAdminReviewRepository(Config.DB)
	auditlogrepo := auditlogrepo.NewAuditLogRepository(Config.DB)

	//service
	productservice := productservice.NewProductService(Config.DB)
	sellerprofileservice := sellerprofileservice.NewSellerService(repository)
	buyerProfileservice := buyerservice.NewBuyerService(buyerProfilerepo)
	guestcartservice := guestcartservice.NewGuestCartService(guestcartrepo, buyercompleterepo, buyerProfilerepo)
	selleracoubtService := selleraccountservice.NewSellerRegistrationService(selleraccountrepo)
	cartservice := cartservice.NewCartService(cartitemrepo)
	categoryService := categoryservice.NewCategoryService(categoryrepo, productrepo)
	selleraccountsettingservice := selleraccountsettingservice.NewSellerAccountSettingService(selleraccountsettingrepo)
	commissionService := commissionpayoutearningservice.NewCommissionService(commissionrepo, orderrepo, *selleraccountsettingrepo)
	
	// Order service configuration
	orderConfig := &orderservice.Config{
		FreeShippingThreshold: 5000.0,  // Free shipping above 5000 Taka
		CommissionRate:        0.15,     // 15% platform commission
	}
	orderservice := orderservice.NewOrderService(orderrepo, cartitemrepo, buyerProfilerepo, *productservice, commissionService, Config.DB, orderConfig)

	adminservice := adminservice.NewAdminService(adminrepo, reguserrepo)
	auditlogservice := auditlogservice.NewAuditLogService(auditlogrepo)

	audithelper := audithelper.NewAuditHelper(auditlogservice)
	// cloudniaryservice := Config.InitializeCloudinary()
	// if cloudniaryservice == nil {
	// 	log.Fatalf("Failed to initialize Cloudinary service. Check your credentials.")
	// }
	//admin management service
	adminbuyerservice := adminmangementservice.NewBuyerService(adminbuyerrepo, adminauditlogrepo)
	adminsellerservice := adminmangementservice.NewSellerService(adminsellerrepo, adminauditlogrepo, audithelper)
	adminproductservice := adminmangementservice.NewProductService(adminproductrepo, adminauditlogrepo)
	adminorderservice := adminmangementservice.NewOrderService(adminorderrepo, adminauditlogrepo, audithelper)
	admindashboardservice := adminmangementservice.NewDashboardService(adminbuyerservice, adminsellerservice, adminproductservice, adminorderservice)

	sellerdashboardservice := sellerdashboardservice.NewDashboardService(sellerdashboardrepo)
	inventoryservice := inventoryservice.NewInventoryService(inventoryrepo, selleraccountsettingrepo)
	reviewservice := reviewservice.NewReviewService(reviewrepo)
	adminreviewservice := adminreviewservice.NewAdminReviewService(adminreviewrepo, productrepo, auditlogrepo)
	//handler
	buyerhandler := buyerhandler.NewHandler()
	cartitemhandler := cartitemhandler.NewHandler()
	orderitemhandler := orderitemhandler.NewHandler()
	producthandler := producthandler.NewHandler(productservice)
	profilehandler := profilehandler.NewHandler()
	reviewhandler := reviewhandler.NewReviewHandler(reviewservice)
	rolehandler := rolehandler.NewHandler()
	searchhandler := searchhandler.NewHandler()
	userhandler := userhandler.NewHandler()
	authhandler := authhandler.NewHandler(authhandler.NewProductService(Config.DB))
	sellerprofilehandler := sellerproflehandler.NewSellerProfileHandler(sellerprofileservice)
	buyerprofilehandler := buyerprofilehandler.NewBuyerProfileHandler(buyerProfileservice)
	guestcarthandler := guestcarthandler.NewGuestCartHandler(guestcartservice)
	selleraccounthandler := selleraccounthandler.NewSellerRegistrationHandler(selleracoubtService)
	carthandler := carthandler.NewCartService(cartservice)
	orderhandler := orderhandler.NewOrderHandler(orderservice)
	adminhandler := adminhandler.NewAdminHandler(adminservice)
	categoryHandler := categoryHandler.NewCategoryHandler(categoryService)
	selleraccountsettinghandler := selleraccountsettinghandler.NewSelleraccountsettinghandler(selleraccountsettingservice)
	//handler of admin management
	adminbuyerhandler := adminmangementhandler.NewBuyerHandler(adminbuyerservice)
	adminsellerhandler := adminmangementhandler.NewSellerHandler(adminsellerservice)
	adminproducthandler := adminmangementhandler.NewProductHandler(adminproductservice)
	adminorderhandler := adminmangementhandler.NewOrderHandler(adminorderservice)
	admindashboardhandler := adminmangementhandler.NewDashboardHandler(admindashboardservice)

	sellerdashboardhandler := sellerdashboardhandler.NewDashboardService(sellerdashboardservice)
	inventoryhandler := inventoryhandler.NewInventoryHandler(inventoryservice)
	commissionHandler := commissionpayoutearningshandler.NewCommissionHandler(commissionService)
	adminreviewhandler := adminreviewhandler.NewAdminReviewService(adminreviewservice)
	auditloghandler := auditloghandler.NewAuditLogHandler(auditlogservice)

	server := rest.NewServer(buyerhandler,
		cartitemhandler,
		orderitemhandler,
		producthandler,
		profilehandler,
		reviewhandler,
		rolehandler,
		searchhandler,
		userhandler,
		authhandler,
		sellerprofilehandler,
		buyerprofilehandler,
		guestcarthandler,
		selleraccounthandler,
		carthandler,
		orderhandler,
		adminhandler,
		categoryHandler,
		selleraccountsettinghandler,
		//admin management
		adminbuyerhandler,
		adminsellerhandler,
		adminproducthandler,
		adminorderhandler,
		admindashboardhandler,

		sellerdashboardhandler,
		inventoryhandler,
		commissionHandler,
		adminreviewhandler,
		auditloghandler,
	)
	server.Start(cnf)

}
