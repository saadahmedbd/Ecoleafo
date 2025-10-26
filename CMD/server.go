package CMD

import (
	"log"

	"github.com/saadahmedbd/Treestore/Config"
	rest "github.com/saadahmedbd/Treestore/Rest"
	adminhandler "github.com/saadahmedbd/Treestore/Rest/Handler/AdminHandler"
	authhandler "github.com/saadahmedbd/Treestore/Rest/Handler/AuthHandler"
	buyerhandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerHandler"
	buyerprofilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerProfileHandler"
	cartitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/CartItemHandler"
	categoryHandler "github.com/saadahmedbd/Treestore/Rest/Handler/CategoryHandler"
	guestcarthandler "github.com/saadahmedbd/Treestore/Rest/Handler/GuestcartHandler"
	orderhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderHandler"
	orderitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderItemHandler"
	producthandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProductHandler"
	profilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProfileHandler"
	reviewhandler "github.com/saadahmedbd/Treestore/Rest/Handler/ReviewHandler"
	rolehandler "github.com/saadahmedbd/Treestore/Rest/Handler/RoleHandler"
	searchhandler "github.com/saadahmedbd/Treestore/Rest/Handler/SearchHandler"
	sellerproflehandler "github.com/saadahmedbd/Treestore/Rest/Handler/SellerProfleHandler"
	userhandler "github.com/saadahmedbd/Treestore/Rest/Handler/UserHandler"
	carthandler "github.com/saadahmedbd/Treestore/Rest/Handler/cartHandler"
	selleraccounthandler "github.com/saadahmedbd/Treestore/Rest/Handler/sellerAccountHandler"
	repository "github.com/saadahmedbd/Treestore/Rest/Repository"
	adminrepo "github.com/saadahmedbd/Treestore/Rest/Repository/AdminRepo"
	buyercompleterepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerCompleteRepo"
	buyerProfilerepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerProfileRepo"
	cartitemrepo "github.com/saadahmedbd/Treestore/Rest/Repository/CartItemRepo"
	categoryrepo "github.com/saadahmedbd/Treestore/Rest/Repository/CategoryRepo"
	guestcartrepo "github.com/saadahmedbd/Treestore/Rest/Repository/GuestCartRepo"
	orderrepo "github.com/saadahmedbd/Treestore/Rest/Repository/OrderRepo"
	productrepo "github.com/saadahmedbd/Treestore/Rest/Repository/ProductRepo"
	reguserrepo "github.com/saadahmedbd/Treestore/Rest/Repository/RegUserRepo"
	selleraccountrepo "github.com/saadahmedbd/Treestore/Rest/Repository/SellerAccountRepo"
	adminservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminService"
	orderservice "github.com/saadahmedbd/Treestore/Rest/Service/OrderService"
	productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"
	selleraccountservice "github.com/saadahmedbd/Treestore/Rest/Service/SellerAccountService"
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

	//service
	productservice := productservice.NewProductService(Config.DB)
	sellerprofileservice := sellerprofileservice.NewSellerService(repository)
	buyerProfileservice := buyerservice.NewBuyerService(buyerProfilerepo)
	guestcartservice := guestcartservice.NewGuestCartService(guestcartrepo, buyercompleterepo, buyerProfilerepo)
	selleracoubtService := selleraccountservice.NewSellerRegistrationService(selleraccountrepo)
	cartservice := cartservice.NewCartService(cartitemrepo)
	orderservice := orderservice.NewOrderService(orderrepo, cartitemrepo, buyerProfilerepo, *productservice)
	adminservice := adminservice.NewAdminService(adminrepo, reguserrepo)
	cloudniaryservice := Config.InitializeCloudinary()
	if cloudniaryservice == nil {
		log.Fatalf("Failed to initialize Cloudinary service. Check your credentials.")
	}
	categoryService := categoryservice.NewCategoryService(categoryrepo, productrepo)

	//handler
	buyerhandler := buyerhandler.NewHandler()
	cartitemhandler := cartitemhandler.NewHandler()
	orderitemhandler := orderitemhandler.NewHandler()
	producthandler := producthandler.NewHandler(productservice)
	profilehandler := profilehandler.NewHandler()
	reviewhandler := reviewhandler.NewHandler()
	rolehandler := rolehandler.NewHandler()
	searchhandler := searchhandler.NewHandler()
	userhandler := userhandler.NewHandler()
	authhandler := authhandler.NewHandler(authhandler.NewProductService(Config.DB))
	sellerprofilehandler := sellerproflehandler.NewSellerProfileHandler(sellerprofileservice)
	buyerprofilehandler := buyerprofilehandler.NewBuyerProfileHandler(buyerProfileservice, *cloudniaryservice)
	guestcarthandler := guestcarthandler.NewGuestCartHandler(guestcartservice)
	selleraccounthandler := selleraccounthandler.NewSellerRegistrationHandler(selleracoubtService)
	carthandler := carthandler.NewCartService(cartservice)
	orderhandler := orderhandler.NewOrderHandler(orderservice)
	adminhandler := adminhandler.NewAdminHandler(adminservice)
	categoryHandler := categoryHandler.NewCategoryHandler(categoryService)

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
	)
	server.Start(cnf)

}
